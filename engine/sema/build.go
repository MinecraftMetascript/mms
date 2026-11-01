package sema

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Build-time values: what a Const, a build-time loop variable or a
// top-level `let` holds. They exist only while compiling; a value used in
// a resource or a command is written out as the literal it stands for.

// ValueKind says what a build-time value is.
type ValueKind uint8

const (
	VNum    ValueKind = iota // a number; Int is set when it is whole
	VStr                     // a string
	VBool                    // true or false
	VWord                    // a bare word or id: stone, minecraft:stone, blocks/ore
	VTag                     // #minecraft:logs
	VList                    // [a, b]
	VRecord                  // { name = "ash", sea = 40 }
	VRand                    // a Rand(key) generator
)

// Value is a build-time value.
type Value struct {
	Kind   ValueKind
	Num    float64
	Int    bool
	Str    string // VStr text, VWord id, VTag name without '#'
	Bool   bool
	Elems  []*Value
	Fields []ValueField
	Rand   *randGen
}

// ValueField is one field of a record, in written order.
type ValueField struct {
	Key string
	Val *Value
}

func numValue(f float64) *Value {
	return &Value{Kind: VNum, Num: f, Int: f == math.Trunc(f) && !math.IsInf(f, 0) && math.Abs(f) < 1<<53}
}

func intValue(n int) *Value { return &Value{Kind: VNum, Num: float64(n), Int: true} }

func strValue(s string) *Value { return &Value{Kind: VStr, Str: s} }

// Field returns a record field, or nil.
func (v *Value) Field(key string) *Value {
	for _, f := range v.Fields {
		if f.Key == key {
			return f.Val
		}
	}
	return nil
}

// TypeName describes the value's kind for messages.
func (v *Value) TypeName() string {
	switch v.Kind {
	case VNum:
		if v.Int {
			return "an int"
		}
		return "a number"
	case VStr:
		return "a string"
	case VBool:
		return "a bool"
	case VWord:
		return "an id"
	case VTag:
		return "a tag"
	case VList:
		return "a list"
	case VRecord:
		return "a record"
	case VRand:
		return "a Rand generator"
	}
	return "a value"
}

// Text is the value as it appears inside a string or command: strings
// and ids without quotes, numbers in their shortest form.
func (v *Value) Text() string {
	switch v.Kind {
	case VNum:
		return numText(v.Num)
	case VStr, VWord:
		return v.Str
	case VTag:
		return "#" + v.Str
	case VBool:
		return strconv.FormatBool(v.Bool)
	}
	return v.String()
}

func numText(f float64) string {
	if f == math.Trunc(f) && math.Abs(f) < 1e15 {
		return strconv.FormatInt(int64(f), 10)
	}
	return strconv.FormatFloat(f, 'g', -1, 64)
}

// String renders the value as MMS source, for hovers and messages.
func (v *Value) String() string {
	switch v.Kind {
	case VStr:
		return strconv.Quote(v.Str)
	case VList:
		parts := make([]string, len(v.Elems))
		for i, e := range v.Elems {
			parts[i] = e.String()
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case VRecord:
		parts := make([]string, len(v.Fields))
		for i, f := range v.Fields {
			parts[i] = f.Key + " = " + f.Val.String()
		}
		return "{ " + strings.Join(parts, ", ") + " }"
	case VRand:
		return "Rand(" + strconv.Quote(v.Rand.key) + ")"
	}
	return v.Text()
}

// Equal compares two values deeply.
func (v *Value) Equal(o *Value) bool {
	if v.Kind != o.Kind {
		return false
	}
	switch v.Kind {
	case VNum:
		return v.Num == o.Num
	case VBool:
		return v.Bool == o.Bool
	case VList:
		if len(v.Elems) != len(o.Elems) {
			return false
		}
		for i := range v.Elems {
			if !v.Elems[i].Equal(o.Elems[i]) {
				return false
			}
		}
		return true
	case VRecord:
		if len(v.Fields) != len(o.Fields) {
			return false
		}
		for _, f := range v.Fields {
			g := o.Field(f.Key)
			if g == nil || !f.Val.Equal(g) {
				return false
			}
		}
		return true
	case VRand:
		return v.Rand.key == o.Rand.key
	}
	return v.Str == o.Str
}

// Expr writes the value as a literal expression at sp, for the checker.
// It returns nil for a value that has no literal form (a generator).
func (v *Value) Expr(sp syntax.Span) syntax.Expr {
	switch v.Kind {
	case VNum:
		return &syntax.Number{Text: numText(v.Num), Sp: sp}
	case VStr:
		return &syntax.String{Value: v.Str, Raw: strconv.Quote(v.Str), Sp: sp}
	case VBool:
		return &syntax.Bool{Value: v.Bool, Sp: sp}
	case VWord:
		if strings.ContainsAny(v.Str, ":/") {
			return &syntax.ResLoc{Text: v.Str, Sp: sp}
		}
		return &syntax.Ident{Name: v.Str, Sp: sp}
	case VTag:
		return &syntax.Tag{Name: v.Str, Sp: sp}
	case VList:
		l := &syntax.List{Sp: sp, Elems: []syntax.Expr{}}
		for _, e := range v.Elems {
			x := e.Expr(sp)
			if x == nil {
				return nil
			}
			l.Elems = append(l.Elems, x)
		}
		return l
	case VRecord:
		b := &syntax.Block{Closed: true, Sp: sp}
		for _, f := range v.Fields {
			x := f.Val.Expr(sp)
			if x == nil {
				return nil
			}
			b.Items = append(b.Items, &syntax.FieldStmt{Key: &syntax.Name{Text: f.Key, Sp: sp}, Value: x, Sp: sp})
		}
		return &syntax.Object{Body: b, Sp: sp}
	}
	return nil
}

// Env holds build-time names: loop variables and lets in a chain of
// scopes, then the project's consts.
type Env struct {
	parent *Env
	vars   map[string]*Value
	order  []string
	NS     string
	consts *Consts
}

// NewEnv returns the root scope for a namespace.
func NewEnv(consts *Consts, ns string) *Env {
	return &Env{consts: consts, NS: ns}
}

// With returns a child scope binding name to v.
func (e *Env) With(name string, v *Value) *Env {
	return &Env{parent: e, vars: map[string]*Value{name: v}, order: []string{name}, NS: e.NS, consts: e.consts}
}

// Lookup finds a build-time name: the innermost binding, then a Const.
// c is set when the name is a Const.
func (e *Env) Lookup(name string) (v *Value, c *constEntry, ok bool) {
	for s := e; s != nil; s = s.parent {
		if v, ok := s.vars[name]; ok {
			return v, nil, true
		}
	}
	if e != nil && e.consts != nil {
		return e.consts.get(e.NS, name)
	}
	return nil, nil, false
}

// Bindings lists the loop variables and lets in scope, outermost first,
// as `name = value` pairs.
func (e *Env) Bindings() []string {
	var out []string
	for s := e; s != nil; s = s.parent {
		for i := len(s.order) - 1; i >= 0; i-- {
			n := s.order[i]
			v := s.vars[n]
			if v.Kind == VRand {
				continue
			}
			out = append(out, n+" = "+short(v.String()))
		}
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func short(s string) string {
	if len(s) > 40 {
		return s[:37] + "..."
	}
	return s
}

// Consts are the project's Const declarations, evaluated on first use.
type Consts struct {
	byID map[string]*constEntry
}

type constEntry struct {
	sym   *Symbol
	unit  *Unit
	decl  *syntax.Decl
	ns    string
	state uint8 // 0 new, 1 evaluating, 2 done
	val   *Value
	diags []syntax.Diagnostic
}

// NewConsts starts an empty set.
func NewConsts() *Consts { return &Consts{byID: map[string]*constEntry{}} }

func (cs *Consts) add(u *Unit, d *syntax.Decl, ns, id string) *constEntry {
	e := &constEntry{unit: u, decl: d, ns: ns}
	cs.byID[id] = e
	return e
}

func (cs *Consts) get(ns, name string) (*Value, *constEntry, bool) {
	id := name
	if !strings.Contains(id, ":") {
		id = ns + ":" + name
	}
	e := cs.byID[id]
	if e == nil {
		return nil, nil, false
	}
	switch e.state {
	case 1:
		e.diags = append(e.diags, syntax.Diagnostic{Code: "mms/const-cycle", Severity: syntax.SeverityError, Sp: e.decl.Name.Sp,
			Message: "Const '" + e.decl.Name.Text + "' depends on itself"})
		e.val = &Value{Kind: VNum, Int: true}
		return e.val, e, true
	case 0:
		e.state = 1
		ev := &Evaluator{Env: NewEnv(cs, e.ns), report: func(d syntax.Diagnostic) { e.diags = append(e.diags, d) }, ctx: "const " + name}
		if e.decl.Value == nil {
			ev.errorf("mms/missing-value", e.decl.Kind.Sp, "a Const needs a value: Const %s = ...", e.decl.NameText())
			e.val = &Value{Kind: VList}
		} else {
			e.val = ev.Eval(e.decl.Value)
			if e.val == nil {
				e.val = &Value{Kind: VList}
			}
		}
		e.state = 2
	}
	return e.val, e, true
}

// Value returns a Const's value by id, evaluating it if needed.
func (cs *Consts) Value(id string) *Value {
	ns, name, _ := strings.Cut(id, ":")
	v, _, _ := cs.get(ns, name)
	return v
}

// Evaluator computes build-time values.
type Evaluator struct {
	Env    *Env
	report func(syntax.Diagnostic)
	// ctx names where unlabeled Rand draws happen, so that draws in two
	// declarations differ; counts numbers the draws per generator.
	ctx    string
	counts map[string]int
	// Uses records each Const an evaluation read.
	Uses func(c *constEntry, at syntax.Node)
	// Shadowed names are runtime variables that hide build-time ones.
	Shadowed func(name string) bool
	failed   bool
}

func (ev *Evaluator) errorf(code string, sp syntax.Span, format string, args ...any) {
	ev.failed = true
	if ev.report != nil {
		ev.report(syntax.Diagnostic{Code: code, Severity: syntax.SeverityError, Sp: sp, Message: fmt.Sprintf(format, args...)})
	}
}

// Eval computes e, reporting problems; it returns nil when it fails.
func (ev *Evaluator) Eval(e syntax.Expr) *Value {
	ev.failed = false
	v := ev.eval(e)
	if ev.failed {
		return nil
	}
	return v
}

// lookup resolves a build-time name, unless a runtime variable hides it.
func (ev *Evaluator) lookup(name string, at syntax.Node) (*Value, bool) {
	if ev.Shadowed != nil && ev.Shadowed(name) {
		return nil, false
	}
	v, c, ok := ev.Env.Lookup(name)
	if ok && c != nil && ev.Uses != nil {
		ev.Uses(c, at)
	}
	return v, ok
}

func (ev *Evaluator) eval(e syntax.Expr) *Value {
	switch x := e.(type) {
	case nil:
		return nil
	case *syntax.Number:
		f, err := strconv.ParseFloat(x.Text, 64)
		if err != nil {
			ev.errorf("mms/type-mismatch", x.Sp, "'%s' is not a number", x.Text)
			return nil
		}
		return numValue(f)
	case *syntax.String:
		s, ok := ev.interpolate(x.Value, x.Sp, true)
		if !ok {
			return nil
		}
		return strValue(s)
	case *syntax.Bool:
		return &Value{Kind: VBool, Bool: x.Value}
	case *syntax.Ident:
		if v, ok := ev.lookup(x.Name, x); ok {
			return v
		}
		return &Value{Kind: VWord, Str: x.Name}
	case *syntax.ResLoc:
		if strings.Contains(x.Text, ":") {
			if v, ok := ev.lookup(x.Text, x); ok {
				return v
			}
		}
		return &Value{Kind: VWord, Str: x.Text}
	case *syntax.Tag:
		return &Value{Kind: VTag, Str: x.Name}
	case *syntax.Paren:
		return ev.eval(x.X)
	case *syntax.List:
		out := &Value{Kind: VList, Elems: []*Value{}}
		for _, el := range x.Elems {
			v := ev.eval(el)
			if v == nil {
				return nil
			}
			out.Elems = append(out.Elems, v)
		}
		return out
	case *syntax.Object:
		out := &Value{Kind: VRecord}
		for _, it := range x.Body.Items {
			f, ok := it.(*syntax.FieldStmt)
			if !ok || f.Key == nil {
				ev.errorf("mms/unsupported", it.Span(), "a build-time record holds fields: { name = value }")
				return nil
			}
			if out.Field(f.Key.Text) != nil {
				ev.errorf("mms/duplicate", f.Key.Sp, "field '%s' is already set", f.Key.Text)
				return nil
			}
			v := ev.eval(f.Value)
			if v == nil {
				return nil
			}
			out.Fields = append(out.Fields, ValueField{Key: f.Key.Text, Val: v})
		}
		return out
	case *syntax.Unary:
		v := ev.eval(x.X)
		if v == nil {
			return nil
		}
		switch x.Op {
		case "!":
			if v.Kind != VBool {
				ev.errorf("mms/type-mismatch", x.X.Span(), "! needs a bool, but this is %s", v.TypeName())
				return nil
			}
			return &Value{Kind: VBool, Bool: !v.Bool}
		case "-", "+":
			if v.Kind != VNum {
				ev.errorf("mms/type-mismatch", x.X.Span(), "%s needs a number, but this is %s", x.Op, v.TypeName())
				return nil
			}
			if x.Op == "-" {
				return numValue(-v.Num)
			}
			return v
		}
	case *syntax.Binary:
		return ev.binary(x)
	case *syntax.Range:
		lo, hi := ev.eval(x.Lo), ev.eval(x.Hi)
		if x.Lo == nil || x.Hi == nil {
			ev.errorf("mms/not-constant", x.Sp, "a build-time range needs both ends: 0..9")
			return nil
		}
		if lo == nil || hi == nil {
			return nil
		}
		if !isInt(lo) || !isInt(hi) {
			ev.errorf("mms/type-mismatch", x.Sp, "a range goes between two whole numbers")
			return nil
		}
		n := int(hi.Num) - int(lo.Num) + 1
		if n > MaxIterations*2 {
			ev.errorf("mms/expansion-limit", x.Sp, "this range has %d values; the limit is %d", n, MaxIterations*2)
			return nil
		}
		out := &Value{Kind: VList, Elems: []*Value{}}
		for i := int(lo.Num); i <= int(hi.Num); i++ {
			out.Elems = append(out.Elems, intValue(i))
		}
		return out
	case *syntax.Member:
		v := ev.eval(x.X)
		if v == nil || x.Name == nil {
			return nil
		}
		if v.Kind != VRecord {
			ev.errorf("mms/type-mismatch", x.Sp, "%s has no fields; '.%s' needs a record", v.TypeName(), x.Name.Text)
			return nil
		}
		f := v.Field(x.Name.Text)
		if f == nil {
			keys := make([]string, len(v.Fields))
			for i, f := range v.Fields {
				keys[i] = f.Key
			}
			msg := "this record has no field '" + x.Name.Text + "'"
			if c := closest(x.Name.Text, keys); c != "" {
				msg += "; did you mean '" + c + "'?"
			}
			ev.errorf("mms/unknown-field", x.Name.Sp, "%s", msg)
			return nil
		}
		return f
	case *syntax.Index:
		v := ev.eval(x.X)
		if v == nil {
			return nil
		}
		if len(x.Args) != 1 || x.Args[0].Name != nil {
			ev.errorf("mms/arity", x.Sp, "index with one value: xs[0]")
			return nil
		}
		k := ev.eval(x.Args[0].Value)
		if k == nil {
			return nil
		}
		switch v.Kind {
		case VList:
			if k.Kind == VWord {
				ev.errorf("mms/not-constant", x.Args[0].Sp, "'%s' is not known at build time, so it cannot index a build-time list; loop over the list instead (for x in ...), or use a constant index", k.Str)
				return nil
			}
			if !isInt(k) {
				ev.errorf("mms/type-mismatch", x.Args[0].Sp, "a list index is a whole number")
				return nil
			}
			i := int(k.Num)
			if i < 0 {
				i += len(v.Elems)
			}
			if i < 0 || i >= len(v.Elems) {
				ev.errorf("mms/out-of-range", x.Args[0].Sp, "index %d is outside this list of %d", int(k.Num), len(v.Elems))
				return nil
			}
			return v.Elems[i]
		case VRecord:
			if k.Kind != VStr {
				ev.errorf("mms/type-mismatch", x.Args[0].Sp, "a record is indexed by a field name in quotes")
				return nil
			}
			if f := v.Field(k.Str); f != nil {
				return f
			}
			ev.errorf("mms/unknown-field", x.Args[0].Sp, "this record has no field '%s'", k.Str)
			return nil
		}
		ev.errorf("mms/type-mismatch", x.X.Span(), "%s cannot be indexed", v.TypeName())
		return nil
	case *syntax.Call:
		return ev.call(x)
	case *syntax.Bad:
		ev.failed = true
		return nil
	}
	ev.errorf("mms/not-constant", e.Span(), "this cannot be computed at build time")
	return nil
}

func isInt(v *Value) bool { return v.Kind == VNum && v.Int }

func (ev *Evaluator) binary(x *syntax.Binary) *Value {
	if x.Op == "&&" || x.Op == "||" {
		l := ev.eval(x.L)
		if l == nil {
			return nil
		}
		if l.Kind != VBool {
			ev.errorf("mms/type-mismatch", x.L.Span(), "%s needs bools, but this is %s", x.Op, l.TypeName())
			return nil
		}
		if (x.Op == "&&") != l.Bool {
			return l
		}
		r := ev.eval(x.R)
		if r == nil {
			return nil
		}
		if r.Kind != VBool {
			ev.errorf("mms/type-mismatch", x.R.Span(), "%s needs bools, but this is %s", x.Op, r.TypeName())
			return nil
		}
		return r
	}
	l, r := ev.eval(x.L), ev.eval(x.R)
	if l == nil || r == nil {
		return nil
	}
	switch x.Op {
	case "==":
		return &Value{Kind: VBool, Bool: l.Equal(r)}
	case "!=":
		return &Value{Kind: VBool, Bool: !l.Equal(r)}
	case "+":
		if l.Kind == VList && r.Kind == VList {
			return &Value{Kind: VList, Elems: append(append([]*Value{}, l.Elems...), r.Elems...)}
		}
		if l.Kind == VStr || r.Kind == VStr {
			if l.Kind == VStr && r.Kind == VStr || l.Kind != VList && r.Kind != VList && l.Kind != VRecord && r.Kind != VRecord {
				return strValue(l.Text() + r.Text())
			}
		}
	}
	if l.Kind != VNum || r.Kind != VNum {
		bad := x.L
		if l.Kind == VNum {
			bad = x.R
		}
		which := l
		if l.Kind == VNum {
			which = r
		}
		ev.errorf("mms/type-mismatch", bad.Span(), "%s needs numbers, but this is %s", x.Op, which.TypeName())
		return nil
	}
	a, b := l.Num, r.Num
	switch x.Op {
	case "+":
		return numValue(a + b)
	case "-":
		return numValue(a - b)
	case "*":
		return numValue(a * b)
	case "/":
		if b == 0 {
			ev.errorf("mms/division-by-zero", x.R.Span(), "division by zero")
			return nil
		}
		if l.Int && r.Int {
			return intValue(int(math.Floor(a / b)))
		}
		return numValue(a / b)
	case "%":
		if b == 0 {
			ev.errorf("mms/division-by-zero", x.R.Span(), "division by zero")
			return nil
		}
		m := math.Mod(a, b)
		if m != 0 && (m < 0) != (b < 0) {
			m += b
		}
		return numValue(m)
	case "<":
		return &Value{Kind: VBool, Bool: a < b}
	case "<=":
		return &Value{Kind: VBool, Bool: a <= b}
	case ">":
		return &Value{Kind: VBool, Bool: a > b}
	case ">=":
		return &Value{Kind: VBool, Bool: a >= b}
	}
	ev.errorf("mms/not-constant", x.Sp, "'%s' cannot be computed at build time", x.Op)
	return nil
}

// interpolate replaces each ${expr} in s with its build-time text. With
// strict unset, a name that is not a build-time value leaves s unchanged
// and reports nothing (ok false).
func (ev *Evaluator) interpolate(s string, sp syntax.Span, strict bool) (string, bool) {
	if !strings.Contains(s, "${") {
		return s, true
	}
	var out strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '$' || i+1 >= len(s) || s[i+1] != '{' {
			out.WriteByte(s[i])
			continue
		}
		depth, j := 0, i+1
		for ; j < len(s); j++ {
			if s[j] == '{' {
				depth++
			} else if s[j] == '}' {
				depth--
				if depth == 0 {
					break
				}
			}
		}
		if j >= len(s) {
			if strict {
				ev.errorf("mms/syntax", sp, "this ${ is not closed")
			}
			return s, false
		}
		inner := s[i+2 : j]
		f := syntax.Parse("X _ = " + inner)
		if len(f.Diags) > 0 || len(f.Decls) != 1 || f.Decls[0].Value == nil {
			if strict {
				ev.errorf("mms/syntax", sp, "cannot parse ${%s}", inner)
			}
			return s, false
		}
		x := f.Decls[0].Value
		if !strict && !mentionsBuildTime(x, ev) {
			return s, false
		}
		reSpan(x, sp)
		if id, ok := x.(*syntax.Ident); ok {
			if _, found := ev.lookup(id.Name, nil); !found {
				if strict {
					ev.errorf("mms/not-constant", sp, "'%s' is not a build-time value; ${...} in a name or build-time string needs a Const, a loop variable or a let", id.Name)
				}
				return s, false
			}
		}
		v := ev.eval(x)
		if v == nil {
			return s, false
		}
		switch v.Kind {
		case VList, VRecord, VRand:
			ev.errorf("mms/type-mismatch", sp, "${%s} is %s; only numbers, strings, bools and ids go into text", inner, v.TypeName())
			return s, false
		}
		out.WriteString(v.Text())
		i = j
	}
	return out.String(), true
}

// reSpan moves every node in a parsed snippet to sp, so errors in it point
// at the string that holds it.
func reSpan(e syntax.Node, sp syntax.Span) {
	syntax.Walk(e, func(n syntax.Node) bool {
		switch x := n.(type) {
		case *syntax.Ident:
			x.Sp = sp
		case *syntax.Number:
			x.Sp = sp
		case *syntax.String:
			x.Sp = sp
		case *syntax.Binary:
			x.Sp = sp
		case *syntax.Member:
			x.Sp = sp
			if x.Name != nil {
				x.Name.Sp = sp
			}
		case *syntax.Index:
			x.Sp = sp
		case *syntax.Call:
			x.Sp, x.ArgsSp = sp, sp
		case *syntax.Arg:
			x.Sp = sp
		case *syntax.Paren:
			x.Sp = sp
		case *syntax.Unary:
			x.Sp = sp
		case *syntax.List:
			x.Sp = sp
		case *syntax.Range:
			x.Sp = sp
		}
		return true
	})
}

// mentionsBuildTime reports whether e uses a build-time name, outside
// call heads, field keys and selectors.
func mentionsBuildTime(e syntax.Node, ev *Evaluator) bool {
	found := false
	var walk func(n syntax.Node)
	walk = func(n syntax.Node) {
		if found || n == nil {
			return
		}
		switch x := n.(type) {
		case *syntax.Ident:
			if _, ok := ev.peek(x.Name); ok {
				found = true
			}
			return
		case *syntax.ResLoc:
			if strings.Contains(x.Text, ":") {
				if _, ok := ev.peek(x.Text); ok {
					found = true
				}
			}
			return
		case *syntax.String:
			if strings.Contains(x.Value, "${") {
				found = true
			}
			return
		case *syntax.Call:
			if id, ok := x.Fn.(*syntax.Ident); ok && buildTimeHelpers[id.Name] {
				found = true
				return
			}
			if _, ok := x.Fn.(*syntax.Ident); !ok {
				walk(x.Fn)
			}
			for _, a := range x.Args {
				walk(a.Value)
			}
			return
		case *syntax.Index:
			if isSelectorIndex(x) {
				return
			}
		case *syntax.Selector, *syntax.Coord:
			return
		case *syntax.Member:
			walk(x.X)
			return
		case *syntax.Arg:
			walk(x.Value)
			return
		case *syntax.FieldStmt:
			walk(x.Value)
			return
		}
		for _, c := range syntax.Children(n) {
			walk(c)
		}
	}
	walk(e)
	return found
}

func isSelectorIndex(x *syntax.Index) bool {
	_, ok := x.X.(*syntax.Selector)
	return ok
}

// peek is lookup without recording a use.
func (ev *Evaluator) peek(name string) (*Value, bool) {
	if ev.Shadowed != nil && ev.Shadowed(name) {
		return nil, false
	}
	v, _, ok := ev.Env.Lookup(name)
	return v, ok
}

// buildTimeHelpers are the calls the evaluator computes.
var buildTimeHelpers = map[string]bool{"Rand": true, "Len": true, "Keys": true, "Unique": true, "Concat": true}

func (ev *Evaluator) call(x *syntax.Call) *Value {
	if m, ok := x.Fn.(*syntax.Member); ok && m.Name != nil {
		recv := ev.eval(m.X)
		if recv == nil {
			return nil
		}
		if recv.Kind != VRand {
			ev.errorf("mms/not-constant", x.Sp, "%s has no method '%s'", recv.TypeName(), m.Name.Text)
			return nil
		}
		return ev.draw(recv.Rand, m.Name, x)
	}
	id, ok := x.Fn.(*syntax.Ident)
	if !ok || !buildTimeHelpers[id.Name] {
		ev.errorf("mms/not-constant", x.Sp, "this call cannot be computed at build time")
		return nil
	}
	args := make([]*Value, 0, len(x.Args))
	for _, a := range x.Args {
		if a.Name != nil {
			ev.errorf("mms/unsupported", a.Sp, "%s takes its arguments by position", id.Name)
			return nil
		}
		v := ev.eval(a.Value)
		if v == nil {
			return nil
		}
		args = append(args, v)
	}
	need := func(n int, what string) bool {
		if len(args) != n {
			ev.errorf("mms/arity", x.ArgsSp, "%s takes %s", id.Name, what)
			return false
		}
		return true
	}
	switch id.Name {
	case "Rand":
		if !need(1, "one key: Rand(\"planet/${i}\")") {
			return nil
		}
		if args[0].Kind != VStr && args[0].Kind != VWord {
			ev.errorf("mms/type-mismatch", x.Args[0].Sp, "a Rand key is a string")
			return nil
		}
		return &Value{Kind: VRand, Rand: &randGen{ns: ev.Env.NS, key: args[0].Str}}
	case "Len":
		if !need(1, "one list, record or string") {
			return nil
		}
		switch a := args[0]; a.Kind {
		case VList:
			return intValue(len(a.Elems))
		case VRecord:
			return intValue(len(a.Fields))
		case VStr:
			return intValue(len([]rune(a.Str)))
		}
		ev.errorf("mms/type-mismatch", x.Args[0].Sp, "Len needs a list, record or string, but this is %s", args[0].TypeName())
		return nil
	case "Keys":
		if !need(1, "one record") {
			return nil
		}
		if args[0].Kind != VRecord {
			ev.errorf("mms/type-mismatch", x.Args[0].Sp, "Keys needs a record, but this is %s", args[0].TypeName())
			return nil
		}
		out := &Value{Kind: VList, Elems: []*Value{}}
		for _, f := range args[0].Fields {
			out.Elems = append(out.Elems, strValue(f.Key))
		}
		return out
	case "Unique":
		if !need(1, "one list") {
			return nil
		}
		if args[0].Kind != VList {
			ev.errorf("mms/type-mismatch", x.Args[0].Sp, "Unique needs a list, but this is %s", args[0].TypeName())
			return nil
		}
		for i, a := range args[0].Elems {
			for _, b := range args[0].Elems[:i] {
				if a.Equal(b) {
					return &Value{Kind: VBool}
				}
			}
		}
		return &Value{Kind: VBool, Bool: true}
	case "Concat":
		out := &Value{Kind: VList, Elems: []*Value{}}
		for i, a := range args {
			if a.Kind != VList {
				ev.errorf("mms/type-mismatch", x.Args[i].Sp, "Concat joins lists, but this is %s", a.TypeName())
				return nil
			}
			out.Elems = append(out.Elems, a.Elems...)
		}
		return out
	}
	return nil
}

// Rand: a build-time generator. Each draw is a pure function of the
// namespace, the generator's key and the draw's label, so builds are
// byte-identical and a new draw never changes the others. The algorithm
// is part of the language: SHA-256 of "namespace|key|label" seeds a
// SplitMix64 stream. Do not change it; projects depend on its output.

type randGen struct {
	ns, key string
}

type stream struct{ s uint64 }

func newStream(ns, key, label string) *stream {
	h := sha256.Sum256([]byte(ns + "|" + key + "|" + label))
	return &stream{s: binary.BigEndian.Uint64(h[:8])}
}

// next is SplitMix64.
func (r *stream) next() uint64 {
	r.s += 0x9e3779b97f4a7c15
	z := r.s
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

// below returns a number in [0, n).
func (r *stream) below(n int) int { return int(r.next() % uint64(n)) }

func (r *stream) float() float64 { return float64(r.next()>>11) / (1 << 53) }

// RandMethods are the generator's methods and their value arguments, for
// completion and arity messages.
var RandMethods = []struct{ Name, Params, Doc string }{
	{"Int", "lo, hi", "a whole number from lo to hi, both included"},
	{"Float", "lo, hi", "a number from lo up to, not including, hi"},
	{"Chance", "p", "true with probability p"},
	{"Pick", "xs", "one element of a list"},
	{"Weighted", "pairs", "one element by weight: [[x, 3], [y, 1]]"},
	{"Sample", "xs, n", "n distinct elements of a list"},
	{"Shuffle", "xs", "the list in a random order"},
	{"Child", "label", "a new generator for a part, keyed under this one"},
}

func (ev *Evaluator) draw(g *randGen, name *syntax.Name, x *syntax.Call) *Value {
	var params string
	for _, m := range RandMethods {
		if m.Name == name.Text {
			params = m.Params
		}
	}
	if params == "" {
		names := make([]string, len(RandMethods))
		for i, m := range RandMethods {
			names[i] = m.Name
		}
		msg := "a Rand generator has no method '" + name.Text + "'"
		if c := closest(name.Text, names); c != "" {
			msg += "; did you mean '" + c + "'?"
		}
		ev.errorf("mms/unknown-name", name.Sp, "%s", msg)
		return nil
	}
	nparams := strings.Count(params, ",") + 1
	args := make([]*Value, 0, len(x.Args))
	for _, a := range x.Args {
		v := ev.eval(a.Value)
		if v == nil {
			return nil
		}
		args = append(args, v)
	}
	if name.Text == "Child" {
		if len(args) != 1 || args[0].Kind != VStr {
			ev.errorf("mms/arity", x.ArgsSp, "Child takes a label: r.Child(\"terrain\")")
			return nil
		}
		return &Value{Kind: VRand, Rand: &randGen{ns: g.ns, key: g.key + "/" + args[0].Str}}
	}
	var label string
	switch len(args) {
	case nparams + 1:
		if args[0].Kind != VStr {
			ev.errorf("mms/type-mismatch", x.Args[0].Sp, "a draw's label is a string: r.%s(\"sea\", %s)", name.Text, params)
			return nil
		}
		label = args[0].Str
		args = args[1:]
	case nparams:
		// The nth unlabeled draw from this generator in this place.
		if ev.counts == nil {
			ev.counts = map[string]int{}
		}
		k := ev.counts[g.key]
		ev.counts[g.key]++
		label = "#" + ev.ctx + "#" + strconv.Itoa(k)
	default:
		ev.errorf("mms/arity", x.ArgsSp, "%s takes (%s), or a label first: (\"label\", %s)", name.Text, params, params)
		return nil
	}
	r := newStream(g.ns, g.key, name.Text+":"+label)
	argAt := func(i int) syntax.Span { return x.Args[len(x.Args)-len(args)+i].Sp }
	list := func(i int) []*Value {
		if args[i].Kind != VList || len(args[i].Elems) == 0 {
			ev.errorf("mms/type-mismatch", argAt(i), "%s needs a list with at least one element", name.Text)
			return nil
		}
		return args[i].Elems
	}
	switch name.Text {
	case "Int":
		if !isInt(args[0]) || !isInt(args[1]) || args[0].Num > args[1].Num {
			ev.errorf("mms/type-mismatch", x.ArgsSp, "Int takes two whole numbers, lo <= hi")
			return nil
		}
		lo, hi := int(args[0].Num), int(args[1].Num)
		return intValue(lo + r.below(hi-lo+1))
	case "Float":
		if args[0].Kind != VNum || args[1].Kind != VNum || args[0].Num > args[1].Num {
			ev.errorf("mms/type-mismatch", x.ArgsSp, "Float takes two numbers, lo <= hi")
			return nil
		}
		f := args[0].Num + r.float()*(args[1].Num-args[0].Num)
		return &Value{Kind: VNum, Num: math.Round(f*1e6) / 1e6}
	case "Chance":
		if args[0].Kind != VNum {
			ev.errorf("mms/type-mismatch", argAt(0), "Chance takes a probability from 0 to 1")
			return nil
		}
		return &Value{Kind: VBool, Bool: r.float() < args[0].Num}
	case "Pick":
		xs := list(0)
		if xs == nil {
			return nil
		}
		return xs[r.below(len(xs))]
	case "Weighted":
		xs := list(0)
		if xs == nil {
			return nil
		}
		total := 0.0
		for _, p := range xs {
			if p.Kind != VList || len(p.Elems) != 2 || p.Elems[1].Kind != VNum || p.Elems[1].Num < 0 {
				ev.errorf("mms/type-mismatch", argAt(0), "Weighted takes [value, weight] pairs with weights of 0 or more")
				return nil
			}
			total += p.Elems[1].Num
		}
		if total <= 0 {
			ev.errorf("mms/type-mismatch", argAt(0), "Weighted needs a weight above 0")
			return nil
		}
		t := r.float() * total
		for _, p := range xs {
			t -= p.Elems[1].Num
			if t < 0 {
				return p.Elems[0]
			}
		}
		return xs[len(xs)-1].Elems[0]
	case "Sample", "Shuffle":
		xs := list(0)
		if xs == nil {
			return nil
		}
		n := len(xs)
		if name.Text == "Sample" {
			if !isInt(args[1]) || args[1].Num < 0 || int(args[1].Num) > len(xs) {
				ev.errorf("mms/out-of-range", argAt(1), "Sample takes from 0 to %d elements of this list", len(xs))
				return nil
			}
			n = int(args[1].Num)
		}
		perm := append([]*Value{}, xs...)
		for i := 0; i < n; i++ {
			j := i + r.below(len(perm)-i)
			perm[i], perm[j] = perm[j], perm[i]
		}
		return &Value{Kind: VList, Elems: perm[:n]}
	}
	return nil
}

// sortedIDs returns m's keys in order.
func sortedIDs[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// buildTime returns e with its build-time value written as a literal
// when e is a build-time expression: a Const, a loop variable or let, a
// computation on them, or a string with ${...}. Other expressions come
// back unchanged; their parts are replaced as the checker reaches them.
// A failed evaluation returns Bad, its errors already reported.
func (c *Checker) buildTime(e syntax.Expr) syntax.Expr {
	root := c
	for root.parent != nil {
		root = root.parent
	}
	if root.Env == nil || e == nil {
		return e
	}
	if r, ok := root.bt[e]; ok {
		return r
	}
	if root.bt == nil {
		root.bt = map[syntax.Node]syntax.Expr{}
	}
	ev := root.evaluator()
	r := e
	if mentionsBuildTime(e, ev) && wholeBuildTime(e, ev) {
		if s, ok := e.(*syntax.String); ok {
			// Leave text alone when its ${...} are not build-time names.
			if text, ok := ev.interpolate(s.Value, s.Sp, false); ok && text != s.Value {
				r = &syntax.String{Value: text, Raw: strconv.Quote(text), Sp: s.Sp}
			} else if ev.failed {
				r = &syntax.Bad{Sp: s.Sp}
			}
			ev.failed = false
		} else if v := ev.Eval(e); v == nil {
			r = &syntax.Bad{Sp: e.Span()}
		} else if v.Kind == VRand {
			root.errorf("mms/type-mismatch", e.Span(), "a Rand generator is not a value; draw from it: r.Int(1, 6)")
			r = &syntax.Bad{Sp: e.Span()}
		} else if x := v.Expr(e.Span()); x != nil {
			r = x
			if !root.trial {
				in := root.info(e)
				if in.Help == "" && in.Symbol == nil {
					in.Help = "Build-time value: `" + short(v.String()) + "`"
				}
			}
		}
	}
	root.bt[e] = r
	return r
}

func (c *Checker) evaluator() *Evaluator {
	if c.ev != nil {
		return c.ev
	}
	c.ev = &Evaluator{Env: c.Env, ctx: c.Ctx,
		report: func(d syntax.Diagnostic) {
			c.errs++
			c.Diags = append(c.Diags, d)
		},
		Uses: func(ce *constEntry, at syntax.Node) {
			sym := ce.sym
			if sym == nil {
				return
			}
			if c.Deps != nil {
				c.Deps["s|Const|"+sym.ID()+"|"+symSig(sym)] = true
			}
			if at != nil {
				in := c.info(at)
				in.Symbol, in.ID = sym, sym.ID()
			}
		}}
	return c.ev
}

// wholeBuildTime reports whether e is computed as a whole at build time:
// a name, a string, or an operation whose every name is build time. A
// list or object with some ordinary ids in it is not; its build-time
// parts are replaced one by one.
func wholeBuildTime(e syntax.Expr, ev *Evaluator) bool {
	switch x := e.(type) {
	case *syntax.Ident, *syntax.ResLoc, *syntax.String:
		return true
	case *syntax.Member:
		return mentionsBuildTime(x.X, ev)
	case *syntax.Index:
		return !isSelectorIndex(x) && mentionsBuildTime(x.X, ev)
	case *syntax.Call:
		if m, ok := x.Fn.(*syntax.Member); ok {
			return mentionsBuildTime(m.X, ev)
		}
		id, ok := x.Fn.(*syntax.Ident)
		return ok && buildTimeHelpers[id.Name]
	case *syntax.Binary, *syntax.Unary, *syntax.Range, *syntax.Paren:
		return allBuildTime(e, ev)
	}
	return false
}

// allBuildTime reports whether every name in e is a build-time value.
func allBuildTime(e syntax.Node, ev *Evaluator) bool {
	ok := true
	syntax.Walk(e, func(n syntax.Node) bool {
		switch x := n.(type) {
		case *syntax.Ident:
			if _, found := ev.peek(x.Name); !found {
				ok = false
			}
		case *syntax.Call:
			if id, isID := x.Fn.(*syntax.Ident); isID && !buildTimeHelpers[id.Name] {
				ok = false
			}
			if id, isID := x.Fn.(*syntax.Ident); isID && buildTimeHelpers[id.Name] {
				for _, a := range x.Args {
					if !allBuildTime(a.Value, ev) {
						ok = false
					}
				}
				return false
			}
		case *syntax.Member:
			if !allBuildTime(x.X, ev) {
				ok = false
			}
			return false
		case *syntax.Selector, *syntax.Coord, *syntax.Tag, *syntax.Object, *syntax.Trailing:
			ok = false
		}
		return ok
	})
	return ok
}
