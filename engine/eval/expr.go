package eval

import (
	"math"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// ctx is where an expression is evaluated.
type ctx struct {
	e *evaluator
	// path is the user file: diagnostics and output belong to it.
	path string
	// defPath is the file the text being evaluated is in: path itself,
	// or the file of a Template or Blueprint.
	defPath string
	// site is set inside text written elsewhere (a Template body, a
	// Blueprint): every node built there gets this user span.
	site *syntax.Span
	fr   []Frame
	sc   *scope
	ns   string
	// compile is set where everything runs at compile time: Templates,
	// Blueprints, const values and top-level for bodies. In function code
	// a `let` is a run-time variable and a range loop stays a loop.
	compile bool
	// emit collects declarations met in a block that holds them.
	emit *[]*Decl
	// inst is the Blueprint instance being expanded, for check rules.
	inst *Instance
	// check allows error, warn and hint.
	check bool
	depth int
}

type scope struct {
	up   *scope
	vars map[string]Value
	// runtime names are function variables: they hide consts and outer
	// bindings, and are left for the function compiler.
	runtime map[string]bool
}

func (s *scope) child() *scope { return &scope{up: s} }

func (s *scope) bind(name string, v Value) {
	if s.vars == nil {
		s.vars = map[string]Value{}
	}
	s.vars[name] = v
	delete(s.runtime, name)
}

func (s *scope) hide(name string) {
	if s.runtime == nil {
		s.runtime = map[string]bool{}
	}
	s.runtime[name] = true
	delete(s.vars, name)
}

// lookup finds a binding. hidden is set when a run-time variable of that
// name is in scope.
func (s *scope) lookup(name string) (v Value, ok, hidden bool) {
	for ; s != nil; s = s.up {
		if v, ok := s.vars[name]; ok {
			return v, true, false
		}
		if s.runtime[name] {
			return nil, false, true
		}
	}
	return nil, false, false
}

func (c *ctx) with() *ctx {
	n := *c
	n.sc = c.sc.child()
	return &n
}

// span is where a node built from text at sp goes in the output.
func (c *ctx) span(sp syntax.Span) syntax.Span {
	if c.site != nil {
		return *c.site
	}
	return sp
}

// pos is the position of a value built from text at sp.
func (c *ctx) pos(sp syntax.Span) pos {
	if c.site != nil {
		return pos{Sp: *c.site, def: &Origin{c.defPath, sp}, fr: c.fr}
	}
	return pos{Sp: sp, fr: c.fr}
}

func (c *ctx) diag(sev syntax.Severity, code string, sp syntax.Span, msg string) {
	c.e.report(c.path, syntax.Diagnostic{Code: code, Severity: sev, Message: msg, Sp: c.span(sp)})
}

func (c *ctx) errorf(code string, sp syntax.Span, msg string) {
	c.diag(syntax.SeverityError, code, sp, msg)
}

// step spends one step of the budget, reporting at sp once it runs out.
func (c *ctx) step(sp syntax.Span) bool {
	if c.e.over {
		return false
	}
	c.e.steps++
	if c.e.steps > c.e.budget {
		c.e.over = true
		c.errorf("mms/eval-budget", sp, "evaluation took too many steps; is a Template calling itself without end?")
		return false
	}
	return true
}

// reduce evaluates an expression as far as compile time allows.
func (c *ctx) reduce(x syntax.Expr) Value {
	if x == nil {
		return nil
	}
	if !c.step(x.Span()) {
		return &Code{c.pos(x.Span()), &syntax.Bad{Sp: c.span(x.Span())}}
	}
	p := c.pos(x.Span())
	switch x := x.(type) {
	case *syntax.Number:
		f, err := strconv.ParseFloat(x.Text, 64)
		if err != nil {
			return &Code{p, c.clone(x)}
		}
		return &Num{p, f, !strings.ContainsAny(x.Text, ".eE")}
	case *syntax.String:
		s, _ := c.interp(x.Value, x.Sp)
		return &Str{p, s}
	case *syntax.Bool:
		return &Bool{p, x.Value}
	case *syntax.Ident:
		return c.ident(x)
	case *syntax.ResLoc:
		return &Word{p, x.Text}
	case *syntax.TemplatedName:
		t, ok := c.nameParts(x.Parts, x.Sp)
		if !ok {
			return &Code{p, &syntax.Bad{Sp: c.span(x.Sp)}}
		}
		return &Word{p, t}
	case *syntax.Paren:
		v := c.reduce(x.X)
		if cd, ok := v.(*Code); ok {
			return &Code{p, &syntax.Paren{X: cd.X, Sp: c.span(x.Sp)}}
		}
		return v
	case *syntax.List:
		l := &List{pos: p}
		for _, el := range x.Elems {
			if sp, ok := el.(*syntax.Spread); ok {
				v := c.reduce(sp.X)
				if inner, ok := v.(*List); ok {
					l.Elems = append(l.Elems, inner.Elems...)
					continue
				}
				if _, none := v.(*None); none {
					continue
				}
				l.Elems = append(l.Elems, &Code{c.pos(sp.Sp), &syntax.Spread{X: c.lower(v), Sp: c.span(sp.Sp)}})
				continue
			}
			l.Elems = append(l.Elems, c.reduce(el))
		}
		return l
	case *syntax.Object:
		return c.object(x)
	case *syntax.Unary:
		return c.unary(x)
	case *syntax.Binary:
		return c.binary(x)
	case *syntax.Range:
		return &RangeV{p, c.reduce(x.Lo), c.reduce(x.Hi)}
	case *syntax.Member:
		return c.member(x)
	case *syntax.Call:
		return c.call(x)
	case *syntax.Index:
		base := c.reduce(x.X)
		if l, ok := base.(*List); ok && len(x.Args) == 1 && x.Args[0].Name == nil {
			if n, ok := c.reduce(x.Args[0].Value).(*Num); ok && n.Int {
				i := int(n.F)
				if i < 0 {
					i += len(l.Elems)
				}
				if i >= 0 && i < len(l.Elems) {
					return l.Elems[i]
				}
				c.errorf("mms/index", x.Sp, "index "+numText(n)+" is outside the list of "+strconv.Itoa(len(l.Elems)))
				return &None{pos: p}
			}
		}
		return &Code{p, &syntax.Index{X: c.lower(base), Args: c.args(x.Args), Sp: c.span(x.Sp)}}
	case *syntax.Trailing:
		fn := c.reduce(x.X)
		return &Code{p, &syntax.Trailing{X: c.lower(fn), Block: c.block(x.Block), Sp: c.span(x.Sp)}}
	case *syntax.IfExpr:
		cond := c.reduce(x.Cond)
		if b, ok := cond.(*Bool); ok {
			if b.V {
				return c.reduce(x.Then)
			}
			if x.Else == nil {
				return &None{pos: p}
			}
			return c.reduce(x.Else)
		}
		if known(cond) {
			c.errorf("mms/condition", x.Cond.Span(), "a condition is true or false, not "+kindName(cond))
			return &None{pos: p}
		}
		out := &syntax.IfExpr{Cond: c.lower(cond), Then: c.lower(c.reduce(x.Then)), Sp: c.span(x.Sp)}
		if x.Else != nil {
			out.Else = c.lower(c.reduce(x.Else))
		}
		return &Code{p, out}
	case *syntax.CommandStmt:
		// A command as a value, `let n = /data get ...`: its ${...}
		// parts are filled in as in a command line.
		t, _ := c.interp(x.Text, x.Sp)
		return &Code{p, &syntax.CommandStmt{Text: t, Sp: c.span(x.Sp)}}
	case *syntax.RunExpr:
		r := c.with()
		r.compile, r.emit = false, nil
		return &Code{p, &syntax.RunExpr{Body: r.block(x.Body), Sp: c.span(x.Sp)}}
	case *syntax.Comprehension:
		return c.comprehension(x)
	case *syntax.Spread:
		v := c.reduce(x.X)
		return &Code{p, &syntax.Spread{X: c.lower(v), Sp: c.span(x.Sp)}}
	}
	return &Code{p, c.clone(x)}
}

// clone copies a node that needs no evaluation into the output.
func (c *ctx) clone(x syntax.Expr) syntax.Expr {
	if c.site != nil {
		return syntax.CloneExprAt(x, *c.site)
	}
	return syntax.CloneExpr(x)
}

func (c *ctx) ident(x *syntax.Ident) Value {
	p := c.pos(x.Sp)
	if x.Name == "none" {
		return &None{pos: p, lit: true}
	}
	v, ok, hidden := c.sc.lookup(x.Name)
	if ok {
		return v
	}
	if !hidden {
		if k := c.e.r.Consts[x.Name]; k != nil {
			v := c.constValue(k, x.Sp)
			if v == nil {
				return &Code{p, c.clone(x)}
			}
			def := Origin{k.Path, k.S.Sp}
			return withPos(v, c.span(x.Sp), append(append([]Frame(nil), c.fr...),
				Frame{Kind: FrameConst, Name: x.Name, Call: Origin{c.path, c.span(x.Sp)}, Def: def}), k.Path)
		}
	}
	if c.compile && !hidden {
		return &Word{p, x.Name}
	}
	// A name the function compiler resolves: a variable, a Global, an
	// objective.
	return &Code{p, c.clone(x)}
}

// constValue evaluates a const once, in its own file.
func (c *ctx) constValue(k *Const, use syntax.Span) Value {
	switch k.state {
	case 1:
		c.errorf("mms/const-cycle", use, "const '"+k.Name()+"' depends on itself")
		return nil
	case 2:
		return k.val
	}
	k.state = 1
	kc := &ctx{e: c.e, path: k.Path, defPath: k.Path, sc: &scope{}, compile: true, ns: c.ns}
	if k.S.Value == nil {
		kc.errorf("mms/const-value", k.S.Sp, "a const needs a value: const "+k.Name()+" = ...")
		k.val = &None{}
	} else {
		k.val = kc.reduce(k.S.Value)
	}
	k.state = 2
	return k.val
}

// nameParts spells a templated name: every `${...}` piece must be known.
func (c *ctx) nameParts(parts []syntax.NamePart, sp syntax.Span) (string, bool) {
	var b strings.Builder
	for _, pt := range parts {
		if pt.X == nil {
			b.WriteString(pt.Lit)
			continue
		}
		v := c.reduce(pt.X)
		t, ok := text(v)
		if !ok || isNone(v) {
			c.errorf("mms/name-piece", pt.Sp, "a ${...} in a name must be known at compile time, but this is "+kindName(v))
			return "", false
		}
		b.WriteString(t)
	}
	return b.String(), true
}

func isNone(v Value) bool {
	_, ok := v.(*None)
	return ok
}

// interp fills the `${...}` pieces of a string or command whose values
// are known. A piece only known at run time stays, written with any
// compile-time names inside it filled in.
func (c *ctx) interp(s string, sp syntax.Span) (string, bool) {
	if !strings.Contains(s, "${") {
		return s, false
	}
	var b strings.Builder
	changed := false
	for _, pc := range pieces(s) {
		if pc.expr == "" {
			b.WriteString(pc.lit)
			continue
		}
		x, diags := syntax.ParseExpr(pc.expr, 0)
		if len(diags) > 0 || !c.refsBound(x) {
			b.WriteString("${" + pc.expr + "}")
			continue
		}
		pin := *c
		pin.site = &sp
		if c.site != nil {
			pin.site = c.site
		}
		v := pin.reduce(x)
		changed = true
		if t, ok := text(v); ok && !isNone(v) {
			b.WriteString(t)
			continue
		}
		b.WriteString("${" + format.Expr(c.lower(v)) + "}")
	}
	return b.String(), changed
}

// refsBound reports whether x names anything this evaluation binds: a
// const, a Template, a loop variable, a parameter or a field.
func (c *ctx) refsBound(x syntax.Expr) bool {
	hit := false
	syntax.Walk(x, func(n syntax.Node) bool {
		if id, ok := n.(*syntax.Ident); ok {
			_, bound, hidden := c.sc.lookup(id.Name)
			if bound || !hidden && (c.e.r.Consts[id.Name] != nil || c.e.r.Templates[id.Name] != nil) {
				hit = true
			}
		}
		return !hit
	})
	return hit
}

// object evaluates an object literal. One that holds only fields and
// spreads is a record; anything else stays code, expanded.
func (c *ctx) object(x *syntax.Object) Value {
	p := c.pos(x.Sp)
	if x.Body == nil {
		return &Record{pos: p}
	}
	for _, it := range x.Body.Items {
		switch it := it.(type) {
		case *syntax.FieldStmt:
		case *syntax.ExprStmt:
			if _, ok := it.X.(*syntax.Spread); !ok {
				return &Code{p, &syntax.Object{Body: c.block(x.Body), Sp: c.span(x.Sp)}}
			}
		default:
			return &Code{p, &syntax.Object{Body: c.block(x.Body), Sp: c.span(x.Sp)}}
		}
	}
	r := &Record{pos: p}
	for _, it := range x.Body.Items {
		if es, ok := it.(*syntax.ExprStmt); ok {
			// A spread of something only the checker can read, such as
			// a helper call: the object stays code.
			switch c.reduce(es.X.(*syntax.Spread).X).(type) {
			case *Record, *None:
			default:
				return &Code{p, &syntax.Object{Body: c.block(x.Body), Sp: c.span(x.Sp)}}
			}
		}
	}
	for _, it := range x.Body.Items {
		switch it := it.(type) {
		case *syntax.FieldStmt:
			if it.Key == nil {
				continue
			}
			key := it.Key.Text
			if it.KeyQuoted {
				key, _ = c.interp(key, it.Key.Sp)
			}
			r.set(key, c.reduce(it.Value))
		case *syntax.ExprStmt:
			v := c.reduce(it.X.(*syntax.Spread).X)
			switch v := v.(type) {
			case *Record:
				for i, k := range v.Keys {
					r.set(k, v.Vals[i])
				}
			case *None:
			default:
				c.errorf("mms/spread", it.Sp, "only an object can be spread into an object, not "+kindName(v))
			}
		}
	}
	return r
}

func (c *ctx) unary(x *syntax.Unary) Value {
	p := c.pos(x.Sp)
	v := c.reduce(x.X)
	switch x.Op {
	case "!":
		if b, ok := v.(*Bool); ok {
			return &Bool{p, !b.V}
		}
	case "-":
		if n, ok := v.(*Num); ok {
			return &Num{p, -n.F, n.Int}
		}
	case "+":
		if n, ok := v.(*Num); ok {
			return n
		}
	}
	return &Code{p, &syntax.Unary{Op: x.Op, X: c.lower(v), Sp: c.span(x.Sp)}}
}

func (c *ctx) binary(x *syntax.Binary) Value {
	p := c.pos(x.Sp)
	l := c.reduce(x.L)
	// Short circuits, so the other side may be anything.
	switch x.Op {
	case "&&", "||":
		if b, ok := l.(*Bool); ok {
			if b.V == (x.Op == "||") {
				return &Bool{p, b.V}
			}
			return c.reduce(x.R)
		}
		if c.compile {
			// `x && false` is false whatever x is: compile-time code has
			// no side effects.
			if b, ok := c.reduce(x.R).(*Bool); ok && b.V == (x.Op == "||") {
				return &Bool{p, b.V}
			}
		}
	case "??":
		if isNone(l) {
			return c.reduce(x.R)
		}
		if known(l) || isData(l) {
			return l
		}
	}
	r := c.reduce(x.R)
	code := func() Value {
		return &Code{p, &syntax.Binary{Op: x.Op, L: c.lower(l), R: c.lower(r), Sp: c.span(x.Sp)}}
	}
	switch x.Op {
	case "==", "!=":
		if known(l) && known(r) {
			return &Bool{p, equal(l, r) == (x.Op == "==")}
		}
		// None only exists at compile time, so anything else, even a
		// value only known when the pack runs, is not none.
		if isNone(l) || isNone(r) {
			return &Bool{p, (isNone(l) && isNone(r)) == (x.Op == "==")}
		}
		return code()
	}
	ln, lok := l.(*Num)
	rn, rok := r.(*Num)
	if lok && rok {
		if v, ok := arith(p, x.Op, ln, rn); ok {
			return v
		}
		if x.Op == "/" || x.Op == "%" {
			c.errorf("mms/divide-by-zero", x.Sp, "division by zero")
			return &Num{p, 0, true}
		}
	}
	if x.Op == "+" {
		switch l := l.(type) {
		case *Str:
			if r, ok := r.(*Str); ok {
				return &Str{p, l.S + r.S}
			}
		case *List:
			if r, ok := r.(*List); ok {
				return &List{p, append(append([]Value(nil), l.Elems...), r.Elems...)}
			}
		case *Record:
			// `components + { ... }`: fields on the right win, except
			// none, which leaves the left one.
			if r, ok := r.(*Record); ok {
				out := &Record{pos: p, Keys: append([]string(nil), l.Keys...), Vals: append([]Value(nil), l.Vals...)}
				for i, k := range r.Keys {
					if !isNone(r.Vals[i]) {
						out.set(k, r.Vals[i])
					}
				}
				return out
			}
		}
	}
	return code()
}

// isData reports whether v is a computed list or record.
func isData(v Value) bool {
	switch v.(type) {
	case *List, *Record, *RangeV, *Stmts:
		return true
	}
	return false
}

func arith(p pos, op string, a, b *Num) (Value, bool) {
	isInt := a.Int && b.Int
	switch op {
	case "+":
		return &Num{p, a.F + b.F, isInt}, true
	case "-":
		return &Num{p, a.F - b.F, isInt}, true
	case "*":
		return &Num{p, a.F * b.F, isInt}, true
	case "/":
		if b.F == 0 {
			return nil, false
		}
		if isInt {
			return &Num{p, math.Trunc(a.F / b.F), true}, true
		}
		return &Num{p, a.F / b.F, false}, true
	case "%":
		if b.F == 0 {
			return nil, false
		}
		return &Num{p, math.Mod(a.F, b.F), isInt}, true
	case "<":
		return &Bool{p, a.F < b.F}, true
	case "<=":
		return &Bool{p, a.F <= b.F}, true
	case ">":
		return &Bool{p, a.F > b.F}, true
	case ">=":
		return &Bool{p, a.F >= b.F}, true
	}
	return nil, false
}

func (c *ctx) member(x *syntax.Member) Value {
	p := c.pos(x.Sp)
	base := c.reduce(x.X)
	if x.Name == nil {
		return &Code{p, &syntax.Member{X: c.lower(base), Sp: c.span(x.Sp)}}
	}
	switch b := base.(type) {
	case *Record:
		if v, ok := b.get(x.Name.Text); ok {
			return v
		}
		return &None{pos: p}
	case *None:
		return &None{pos: p}
	}
	return &Code{p, &syntax.Member{X: c.lower(base), Name: &syntax.Name{Text: x.Name.Text, Sp: c.span(x.Name.Sp)}, Sp: c.span(x.Sp)}}
}

func (c *ctx) args(as []*syntax.Arg) []*syntax.Arg {
	out := make([]*syntax.Arg, 0, len(as))
	for _, a := range as {
		n := &syntax.Arg{Value: c.lower(c.reduce(a.Value)), Sp: c.span(a.Sp)}
		if a.Name != nil {
			n.Name = &syntax.Name{Text: a.Name.Text, Sp: c.span(a.Name.Sp)}
		}
		if n.Value == nil {
			n.Value = &syntax.Ident{Name: "none", Sp: c.span(a.Sp)}
		}
		out = append(out, n)
	}
	return out
}

func (c *ctx) comprehension(x *syntax.Comprehension) Value {
	p := c.pos(x.Sp)
	iter := c.reduce(x.Iter)
	items, ok := c.iterate(iter, x.Iter.Span(), true)
	if !ok {
		return &Code{p, &syntax.Bad{Sp: c.span(x.Sp)}}
	}
	out := &List{pos: p}
	for _, it := range items {
		if c.e.over {
			return &Code{p, &syntax.Bad{Sp: c.span(x.Sp)}}
		}
		ic := c.with()
		ic.bindLoop(x.Var, x.Var2, it)
		if x.Cond != nil {
			cond := ic.reduce(x.Cond)
			b, ok := cond.(*Bool)
			if !ok {
				c.errorf("mms/condition", x.Cond.Span(), "a comprehension's filter must be known at compile time, but this is "+kindName(cond))
				return &Code{p, &syntax.Bad{Sp: c.span(x.Sp)}}
			}
			if !b.V {
				continue
			}
		}
		if v := ic.reduce(x.Body); !isNone(v) {
			out.Elems = append(out.Elems, v)
		}
	}
	return out
}

// loopItem is one iteration's values. With one loop variable it is one:
// a list's element or a record's key. With two it is key and val: index
// and element, or key and value.
type loopItem struct{ one, key, val Value }

// iterate lists what a compile-time loop walks: a list's elements, a
// record's keys and values, or a range's integers (only where ranges
// unroll). ok is false, after a diagnostic, for anything else.
func (c *ctx) iterate(v Value, sp syntax.Span, ranges bool) ([]loopItem, bool) {
	switch v := v.(type) {
	case *List:
		out := make([]loopItem, len(v.Elems))
		for i, el := range v.Elems {
			out[i] = loopItem{el, &Num{v.pos, float64(i), true}, el}
		}
		return out, true
	case *Record:
		out := make([]loopItem, len(v.Keys))
		for i, k := range v.Keys {
			out[i] = loopItem{&Str{v.pos, k}, &Str{v.pos, k}, v.Vals[i]}
		}
		return out, true
	case *None:
		return nil, true
	case *RangeV:
		lo, ok1 := v.Lo.(*Num)
		hi, ok2 := v.Hi.(*Num)
		if ranges && ok1 && ok2 && lo.Int && hi.Int {
			var out []loopItem
			for i := lo.F; i <= hi.F; i++ {
				if !c.step(sp) {
					break
				}
				n := &Num{v.pos, i, true}
				out = append(out, loopItem{n, &Num{v.pos, i - lo.F, true}, n})
			}
			return out, true
		}
	}
	c.errorf("mms/loop-value", sp, "a loop here must go over a list or range known at compile time, but this is "+kindName(v))
	return nil, false
}

// bindLoop binds a loop's variables for one iteration.
func (c *ctx) bindLoop(v1, v2 *syntax.Name, it loopItem) {
	switch {
	case v1 == nil:
	case v2 == nil:
		c.sc.bind(v1.Text, it.one)
	default:
		c.sc.bind(v1.Text, it.key)
		c.sc.bind(v2.Text, it.val)
	}
}
