package sema

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Value types in function bodies.
const (
	tInt    = "int"
	tBool   = "bool"
	tString = "string"
)

// FuncSig is a function's parameter and return types.
type FuncSig struct {
	Params []FuncParam
	Ret    string // "" for none
}

type FuncParam struct {
	Name, Type string
}

func (s FuncSig) String() string {
	var ps []string
	for _, p := range s.Params {
		ps = append(ps, p.Name+": "+p.Type)
	}
	out := "(" + strings.Join(ps, ", ") + ")"
	if s.Ret != "" {
		out += ": " + s.Ret
	}
	return out
}

// SigOf reads a function declaration's signature. Unknown types are kept
// as written; the declaration's own check reports them.
func SigOf(d *syntax.Decl) FuncSig {
	var s FuncSig
	if d.Params != nil {
		for _, p := range d.Params.Params {
			fp := FuncParam{Type: tInt}
			if p.Name != nil {
				fp.Name = p.Name.Text
			}
			if p.Type != nil {
				fp.Type = p.Type.Text
			}
			s.Params = append(s.Params, fp)
		}
	}
	if d.Type != nil {
		s.Ret = d.Type.Text
	}
	return s
}

// SymSig is the signature part of a symbol that other declarations
// depend on; it is "" for most kinds.
func SymSig(s *Symbol) string {
	if s.Kind.MMS == "Function" || s.Kind.MMS == "Global" || s.Kind.MMS == customItemKind || s.Kind.MMS == "Const" {
		return symSig(s)
	}
	return ""
}

// symSig is what a dependent declaration relies on in a symbol: a
// function's signature or a global's type.
func symSig(s *Symbol) string {
	if s == nil || s.Decl == nil {
		return ""
	}
	if s.Kind.MMS == customItemKind {
		return customItemSig(s)
	}
	if s.Kind.MMS == "Const" {
		if s.Const == nil {
			return "?"
		}
		return s.Const.String()
	}
	if s.Kind.MMS == "Global" {
		if s.Decl.Type == nil {
			return tInt
		}
		return s.Decl.Type.Text
	}
	// Callers also depend on what the function calls, which decides
	// whether a call is recursive.
	return SigOf(s.Decl).String() + " calls " + strings.Join(callees(s.Decl, s.NS), ",")
}

// callees lists the functions a body calls, by id, sorted: calls in
// statements and in ${...} interpolations.
func callees(d *syntax.Decl, ns string) []string {
	seen := map[string]bool{}
	var walk func(n syntax.Node)
	walk = func(n syntax.Node) {
		switch x := n.(type) {
		case *syntax.Call:
			switch fx := x.Fn.(type) {
			case *syntax.Ident:
				seen[ns+":"+fx.Name] = true
			case *syntax.ResLoc:
				seen[fx.Text] = true
			}
		case *syntax.CommandStmt:
			for _, e := range interpolations(x.Text) {
				if f := syntax.Parse("X _ = " + e); len(f.Decls) == 1 && f.Decls[0].Value != nil {
					walk(f.Decls[0].Value)
				}
			}
		}
		for _, c := range syntax.Children(n) {
			walk(c)
		}
	}
	if d.Body != nil {
		walk(d.Body)
	}
	out := make([]string, 0, len(seen))
	for id := range seen {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

// interpolations returns the expressions inside ${...} in a command.
func interpolations(text string) []string {
	var out []string
	for {
		i := strings.Index(text, "${")
		if i < 0 {
			return out
		}
		depth, j := 0, i+1
		for ; j < len(text); j++ {
			if text[j] == '{' {
				depth++
			} else if text[j] == '}' {
				if depth--; depth == 0 {
					break
				}
			}
		}
		if j >= len(text) {
			return out
		}
		out = append(out, text[i+2:j])
		text = text[j+1:]
	}
}

// reaches reports whether calling from can lead back to target.
func (f *fc) reaches(from, target string) bool {
	seen := map[string]bool{}
	queue := []string{from}
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		if id == target {
			return true
		}
		if seen[id] {
			continue
		}
		seen[id] = true
		sym := f.c.findSig("function", id)
		if sym == nil || sym.Decl == nil {
			continue
		}
		queue = append(queue, callees(sym.Decl, sym.NS)...)
	}
	return false
}

// findSig looks a symbol up and records a dependency on its signature,
// so callers are rechecked when it changes.
func (c *Checker) findSig(folder, id string) *Symbol {
	s := c.find(folder, id)
	if s != nil && c.Deps != nil {
		c.Deps["s|"+folder+"|"+id+"|"+symSig(s)] = true
	}
	return s
}

// GeneratedID names the function an OnLoad, OnTick, OnFirstJoin or
// OnRespawn body compiles to when it needs one: by content, so it is
// stable while the code is.
func GeneratedID(ns, kind string, d *syntax.Decl, src []rune) string {
	h := sha256.Sum256([]byte(string(src[d.Sp.Start:d.Sp.End])))
	return ns + ":mms_gen/" + HookKinds[kind] + "/" + hex.EncodeToString(h[:4])
}

// IsRaw reports whether a body holds only raw commands, which need no
// backend.
func IsRaw(b *syntax.Block) bool {
	if b == nil {
		return true
	}
	for _, it := range b.Items {
		cmd, ok := it.(*syntax.CommandStmt)
		if !ok || strings.Contains(cmd.Text, "${") {
			return false
		}
	}
	return true
}

// local is a variable: ints and bools live in a score, strings at a
// storage path.
type local struct {
	s    fn.Score
	path string
	typ  string
	info *Local
}

// Local describes a function parameter, local variable or loop variable,
// for hovers on its uses.
type Local struct {
	Name, Type string
	What       string // "parameter", "let" or "for"
	Doc        string // the doc comment, or the function's @param text
}

// fc compiles one function body.
type fc struct {
	c      *Checker
	src    []rune
	ns     string
	path   string
	obj    string
	sig    FuncSig
	scopes []map[string]local
	names  map[string]int // holder names used, to keep shadowed locals apart
	temps  int
	loops  int
	// exec counts the execute blocks being compiled; control flow cannot
	// leave them.
	exec   int
	selfID string
	// loopVar is set while declaring a for loop's variable.
	loopVar bool
	// env is the build-time scope: consts, the loop variables of the
	// build-time loop that made the declaration, and the variables of
	// unrolled loops in this body.
	env *Env
	// unroll holds, per enclosing loop from the outside in, the flag a
	// break in an unrolled loop sets; nil for a scoreboard loop.
	unroll []*unrolled
	// frame lists the scores the function has used so far; a recursive
	// call saves and restores them. strings is set once it uses storage.
	frame   []fn.Score
	strings bool
}

// CompileFunction checks a function body and lowers it to the IR. id is
// the function's "ns:path"; src is the unit's source, for selector text.
func (c *Checker) CompileFunction(d *syntax.Decl, id string, src []rune) *fn.Func {
	if d.Body == nil {
		c.errorf("mms/missing-value", d.Kind.Sp, "%s needs a body: { ... }", d.Kind.Name)
		return nil
	}
	ns, path, _ := strings.Cut(id, ":")
	f := &fc{c: c, src: src, ns: ns, path: path, obj: fn.Objective(ns), sig: SigOf(d), names: map[string]int{}, selfID: id, env: c.Env}
	f.push()
	if d.Params != nil {
		for i, p := range d.Params.Params {
			if p.Name == nil {
				continue
			}
			t := f.sig.Params[i].Type
			if !f.valueType(t, p.Type) {
				t = tInt
			}
			if _, dup := f.scopes[0][p.Name.Text]; dup {
				c.errorf("mms/duplicate", p.Name.Sp, "parameter '%s' is already declared", p.Name.Text)
			}
			f.names[p.Name.Text]++
			l := local{s: f.score(fn.Holder(path, p.Name.Text)), typ: t,
				info: &Local{Name: p.Name.Text, Type: t, What: "parameter", Doc: syntax.ParseDoc(d.Doc).Param(p.Name.Text)}}
			f.c.info(p.Name).Local = l.info
			if t == tString {
				l.path = fn.Storage(ns) + " " + fn.LocalPath(path, p.Name.Text)
				f.strings = true
			} else {
				f.frame = append(f.frame, l.s)
			}
			f.scopes[0][p.Name.Text] = l
		}
	}
	if f.sig.Ret != "" && !f.valueType(f.sig.Ret, d.Type) {
		f.sig.Ret = tInt
	}
	body := &fn.Block{}
	f.block(body, d.Body)
	if f.sig.Ret != "" && !terminates(d.Body) {
		c.errorf("mms/missing-return", d.Body.Sp, "this function returns %s, but it can reach its end without a return", f.sig.Ret)
	}
	return &fn.Func{ID: id, Body: body}
}

func (f *fc) valueType(t string, at *syntax.Name) bool {
	switch t {
	case tInt, tBool, tString:
		return true
	}
	sp := syntax.Span{}
	if at != nil {
		sp = at.Sp
	}
	f.c.errorf("mms/unknown-type", sp, "unknown type '%s'; function values are int, bool or string", t)
	return false
}

// terminates reports whether a block always ends in a return.
func terminates(b *syntax.Block) bool {
	if b == nil || len(b.Items) == 0 {
		return false
	}
	switch s := b.Items[len(b.Items)-1].(type) {
	case *syntax.ReturnStmt:
		return true
	case *syntax.IfStmt:
		return ifTerminates(s)
	}
	return false
}

func ifTerminates(s *syntax.IfStmt) bool {
	if !terminates(s.Then) {
		return false
	}
	switch e := s.Else.(type) {
	case *syntax.Block:
		return terminates(e)
	case *syntax.IfStmt:
		return ifTerminates(e)
	}
	return false
}

func (f *fc) score(holder string) fn.Score { return fn.Score{Holder: holder, Obj: f.obj} }

func (f *fc) push() { f.scopes = append(f.scopes, map[string]local{}) }
func (f *fc) pop()  { f.scopes = f.scopes[:len(f.scopes)-1] }

func (f *fc) lookup(name string) (local, bool) {
	for i := len(f.scopes) - 1; i >= 0; i-- {
		if l, ok := f.scopes[i][name]; ok {
			return l, true
		}
	}
	return local{}, false
}

// declare adds a local. A name shadowing an outer one gets its own holder.
func (f *fc) declare(n *syntax.Name, typ string) local {
	if _, dup := f.scopes[len(f.scopes)-1][n.Text]; dup {
		f.c.errorf("mms/duplicate", n.Sp, "'%s' is already declared in this block", n.Text)
	}
	holder := n.Text
	if k := f.names[n.Text]; k > 0 {
		holder = n.Text + "_" + strconv.Itoa(k)
	}
	f.names[n.Text]++
	what := "let"
	if len(f.scopes) > 0 && f.loopVar {
		what = "for"
	}
	l := local{s: f.score(fn.Holder(f.path, holder)), typ: typ, info: &Local{Name: n.Text, Type: typ, What: what}}
	f.c.info(n).Local = l.info
	if typ == tString {
		l.path = fn.Storage(f.ns) + " " + fn.LocalPath(f.path, holder)
		f.strings = true
	} else {
		f.frame = append(f.frame, l.s)
	}
	f.scopes[len(f.scopes)-1][n.Text] = l
	return l
}

func (f *fc) temp() fn.Score {
	s := f.score(fn.Holder(f.path, "_t"+strconv.Itoa(f.temps)))
	f.temps++
	f.frame = append(f.frame, s)
	return s
}

func (f *fc) block(out *fn.Block, b *syntax.Block) {
	f.push()
	defer f.pop()
	for _, it := range b.Items {
		f.stmt(out, it)
	}
}

func (f *fc) stmt(out *fn.Block, s syntax.Stmt) {
	c := f.c
	switch s := s.(type) {
	case *syntax.CommandStmt:
		c.Command(s)
		if strings.Contains(s.Text, "${") {
			f.interpolate(out, s)
			return
		}
		out.Add(fn.Cmd{Text: s.Text})
	case *syntax.LetStmt:
		if s.Name == nil || f.buildLet(s) {
			return
		}
		typ := ""
		if s.Type != nil {
			typ = s.Type.Text
			if !f.valueType(typ, s.Type) {
				typ = tInt
			}
		}
		var v *val
		if s.Value != nil {
			v = f.expr(&out.Ops, s.Value)
			if v != nil && typ == "" {
				typ = v.typ
			}
		}
		if typ == "" {
			typ = tInt
		}
		l := f.declare(s.Name, typ)
		l.info.Doc = s.Doc
		if v == nil {
			if typ == tString {
				empty := ""
				out.Add(f.setString(l.path, &val{typ: tString, str: &empty}))
				return
			}
			out.Add(fn.Set{Dst: l.s, Src: fn.Const(0)})
			return
		}
		f.store(&out.Ops, l, v, s.Value)
	case *syntax.FieldStmt:
		if s.Key == nil || s.Value == nil {
			return
		}
		target := f.variable(&syntax.Ident{Name: s.Key.Text, Sp: s.Key.Sp})
		if target == nil {
			return
		}
		f.store(&out.Ops, *target, f.expr(&out.Ops, s.Value), s.Value)
	case *syntax.AssignStmt:
		target := f.target(&out.Ops, s.Target)
		if target == nil || s.Value == nil {
			return
		}
		if s.Op == "=" {
			f.store(&out.Ops, *target, f.expr(&out.Ops, s.Value), s.Value)
			return
		}
		if target.typ == tString && s.Op == "+=" {
			v := f.expr(&out.Ops, s.Value)
			if v == nil {
				return
			}
			joined := f.concat(&out.Ops, &val{typ: tString, path: target.path}, v)
			out.Add(f.setString(target.path, joined))
			return
		}
		if target.typ != tInt {
			c.errorf("mms/type-mismatch", s.Sp, "%s needs an int, but this is %s", s.Op, target.typ)
			return
		}
		v := f.intExpr(&out.Ops, s.Value)
		if v == nil {
			return
		}
		out.Add(fn.Arith{Dst: target.s, Op: s.Op, Src: *v})
	case *syntax.IfStmt:
		f.ifStmt(out, s)
	case *syntax.WhileStmt:
		cond := f.cond(s.Cond)
		body := &fn.Block{}
		f.loops++
		f.unroll = append(f.unroll, nil)
		f.block(body, s.Body)
		f.unroll = f.unroll[:len(f.unroll)-1]
		f.loops--
		out.Add(fn.Loop{Cond: cond, Body: body})
	case *syntax.ForStmt:
		if f.unrollFor(out, s) {
			return
		}
		f.forStmt(out, s)
	case *syntax.ReturnStmt:
		f.returnStmt(out, s)
	case *syntax.BreakStmt:
		if f.loops == 0 && f.exec > 0 {
			c.errorf("mms/misplaced", s.Sp, "break cannot leave an execute block, which can run once per entity; put the loop inside the block")
			return
		}
		if f.loops == 0 {
			c.errorf("mms/misplaced", s.Sp, "break is only allowed inside a loop")
			return
		}
		if u := f.innerUnrolled(); u != nil {
			out.Add(fn.Set{Dst: *u.flag, Src: fn.Const(1)})
		}
		out.Add(fn.Break{})
	case *syntax.ContinueStmt:
		if f.loops == 0 && f.exec > 0 {
			c.errorf("mms/misplaced", s.Sp, "continue cannot leave an execute block, which can run once per entity; put the loop inside the block")
			return
		}
		if f.loops == 0 {
			c.errorf("mms/misplaced", s.Sp, "continue is only allowed inside a loop")
			return
		}
		if len(f.unroll) > 0 && f.unroll[len(f.unroll)-1] != nil {
			// The iteration runs as a loop of its own; leaving it goes on
			// to the next one.
			out.Add(fn.Break{})
			return
		}
		out.Add(fn.Continue{})
	case *syntax.ExprStmt:
		if t, ok := s.X.(*syntax.Trailing); ok {
			f.execBlock(out, t)
			return
		}
		if IsExecChain(s.X) {
			name, _ := execHeadName(s.X)
			c.errorf("mms/missing-value", s.Sp, "an execute chain needs a block to run: %s(...) { ... }", name)
			return
		}
		call, ok := s.X.(*syntax.Call)
		if !ok {
			c.errorf("mms/unused-value", s.Sp, "this expression does nothing; did you mean to assign it?")
			return
		}
		f.call(&out.Ops, call, false)
	case *syntax.Block:
		f.block(out, s)
	default:
		c.errorf("mms/unsupported", s.Span(), "this is not allowed in a function body")
	}
}

func (f *fc) ifStmt(out *fn.Block, s *syntax.IfStmt) {
	cond := f.cond(s.Cond)
	then := &fn.Block{}
	f.block(then, s.Then)
	var els *fn.Block
	switch e := s.Else.(type) {
	case *syntax.Block:
		els = &fn.Block{}
		f.block(els, e)
	case *syntax.IfStmt:
		els = &fn.Block{}
		f.ifStmt(els, e)
	}
	out.Add(fn.If{Cond: cond, Then: then, Else: els})
}

// forStmt lowers `for i in a..b` (both ends included) to a loop.
func (f *fc) forStmt(out *fn.Block, s *syntax.ForStmt) {
	r, ok := s.Iter.(*syntax.Range)
	if !ok || r.Lo == nil || r.Hi == nil || s.Var == nil {
		f.c.errorf("mms/unsupported", s.Sp, "for loops go over a range: for i in 1..10 { ... }")
		return
	}
	f.push()
	defer f.pop()
	lo := f.intExpr(&out.Ops, r.Lo)
	hi := f.intExpr(&out.Ops, r.Hi)
	if lo == nil || hi == nil {
		return
	}
	f.loopVar = true
	i := f.declare(s.Var, tInt)
	f.loopVar = false
	out.Add(fn.Set{Dst: i.s, Src: *lo})
	var cond fn.Cond
	if hi.IsConst {
		cond.Parts = []string{fmt.Sprintf("if score %s matches ..%d", i.s, hi.Const)}
	} else {
		end := f.temp()
		out.Add(fn.Set{Dst: end, Src: *hi})
		cond.Parts = []string{fmt.Sprintf("if score %s <= %s", i.s, end)}
	}
	body := &fn.Block{}
	f.loops++
	f.unroll = append(f.unroll, nil)
	f.block(body, s.Body)
	f.unroll = f.unroll[:len(f.unroll)-1]
	f.loops--
	out.Add(fn.Loop{Cond: cond, Body: body, Post: &fn.Block{Ops: []fn.Op{fn.Arith{Dst: i.s, Op: "+=", Src: fn.Const(1)}}}})
}

func (f *fc) returnStmt(out *fn.Block, s *syntax.ReturnStmt) {
	if f.exec > 0 {
		f.c.errorf("mms/misplaced", s.Sp, "return cannot leave an execute block, which can run once per entity; set a variable in the block and return after it")
		return
	}
	if s.Value == nil {
		if f.sig.Ret != "" {
			f.c.errorf("mms/missing-value", s.Sp, "this function returns %s; return a value", f.sig.Ret)
		}
		out.Add(fn.Return{})
		return
	}
	if f.sig.Ret == "" {
		f.c.errorf("mms/type-mismatch", s.Value.Span(), "this function has no return type; declare one, as in Function f(): int")
		return
	}
	v := f.expr(&out.Ops, s.Value)
	if v == nil {
		return
	}
	if v.typ == tString {
		f.c.errorf("mms/unsupported", s.Value.Span(), "functions return int or bool; pass strings back through a Global")
		return
	}
	if v.typ != f.sig.Ret {
		f.c.errorf("mms/type-mismatch", s.Value.Span(), "this function returns %s, but this is %s", f.sig.Ret, v.typ)
		return
	}
	op := f.operand(&out.Ops, v)
	out.Add(fn.Return{Val: &op})
}

// val is an expression's value: an int or bool operand, or a condition
// for a bool not yet stored.
type val struct {
	typ  string
	op   fn.Operand
	cond *fn.Cond
	str  *string // a constant string
	path string  // a string in storage
}

// operand stores a value as an operand, turning a condition into 0 or 1.
func (f *fc) operand(pre *[]fn.Op, v *val) fn.Operand {
	if v.cond == nil {
		return v.op
	}
	if v.cond.Never {
		*pre = append(*pre, v.cond.Pre...)
		return fn.Const(0)
	}
	if v.cond.Always() {
		*pre = append(*pre, v.cond.Pre...)
		return fn.Const(1)
	}
	t := f.temp()
	*pre = append(*pre, fn.Set{Dst: t, Src: fn.Const(0)}, fn.If{Cond: *v.cond, Then: &fn.Block{Ops: []fn.Op{fn.Set{Dst: t, Src: fn.Const(1)}}}})
	return fn.Of(t)
}

func (f *fc) store(pre *[]fn.Op, l local, v *val, e syntax.Expr) {
	if v == nil {
		return
	}
	if v.typ != l.typ {
		f.c.errorf("mms/type-mismatch", e.Span(), "expected %s, but this is %s", l.typ, v.typ)
		return
	}
	if l.typ == tString {
		*pre = append(*pre, f.setString(l.path, v))
		return
	}
	if v.cond != nil && !v.cond.Never && !v.cond.Always() {
		*pre = append(*pre, fn.Set{Dst: l.s, Src: fn.Const(0)}, fn.If{Cond: *v.cond, Then: &fn.Block{Ops: []fn.Op{fn.Set{Dst: l.s, Src: fn.Const(1)}}}})
		return
	}
	*pre = append(*pre, fn.Set{Dst: l.s, Src: f.operand(pre, v)})
}

// variable resolves a name to a local, parameter or Global.
func (f *fc) variable(e syntax.Expr) *local {
	switch e := e.(type) {
	case *syntax.Ident:
		if l, ok := f.lookup(e.Name); ok {
			if l.info != nil {
				f.c.info(e).Local = l.info
			}
			return &l
		}
		if g := f.global(f.ns+":"+e.Name, e); g != nil {
			return g
		}
		if f.c.find("Objective", f.ns+":"+e.Name) != nil {
			f.c.errorf("mms/type-mismatch", e.Sp, "'%s' is an objective; read a score with %s[@s]", e.Name, e.Name)
			return nil
		}
		f.c.errorf("mms/unknown-name", e.Sp, "unknown variable '%s'", e.Name)
	case *syntax.ResLoc:
		if g := f.global(e.Text, e); g != nil {
			return g
		}
		f.c.errorf("mms/unknown-name", e.Sp, "no Global named '%s' is declared", e.Text)
	default:
		f.c.errorf("mms/unsupported", e.Span(), "expected a variable")
	}
	return nil
}

// customItem returns the custom item an interpolated name refers to,
// when no variable has that name.
func (f *fc) customItem(e syntax.Expr) *Symbol {
	if id, ok := e.(*syntax.Ident); ok {
		if _, ok := f.lookup(id.Name); ok {
			return nil
		}
		if f.c.Syms.Find("Global", f.ns+":"+id.Name) != nil {
			return nil
		}
	}
	if !f.c.isCustomItem(e) {
		return nil
	}
	return f.c.customItemRef(e)
}

func (f *fc) global(id string, e syntax.Expr) *local {
	sym := f.c.findSig("Global", id)
	if sym == nil {
		return nil
	}
	in := f.c.info(e)
	in.Symbol = sym
	in.ID = id
	typ := symSig(sym)
	switch typ {
	case tString:
		return &local{typ: tString, path: fn.Storage(sym.NS) + " " + fn.GlobalPath(sym.Path)}
	case tInt, tBool:
	default:
		typ = tInt
	}
	return &local{s: fn.Score{Holder: "$" + sym.Path, Obj: fn.Objective(sym.NS)}, typ: typ}
}

// target resolves an assignment target: a variable or an objective score.
func (f *fc) target(pre *[]fn.Op, e syntax.Expr) *local {
	if ix, ok := e.(*syntax.Index); ok {
		s := f.objScore(ix)
		if s == nil {
			return nil
		}
		return &local{s: *s, typ: tInt}
	}
	return f.variable(e)
}

// objScore resolves `kills[@s]` to a score.
func (f *fc) objScore(ix *syntax.Index) *fn.Score {
	var name, id string
	switch x := ix.X.(type) {
	case *syntax.Ident:
		name, id = x.Name, f.ns+":"+x.Name
	case *syntax.ResLoc:
		id = x.Text
		_, name, _ = strings.Cut(x.Text, ":")
	default:
		f.c.errorf("mms/unsupported", ix.Sp, "expected an objective's score, such as kills[@s]")
		return nil
	}
	sym := f.c.find("Objective", id)
	if sym == nil {
		f.c.errorf("mms/unknown-name", ix.X.Span(), "no Objective named '%s' is declared", name)
		return nil
	}
	in := f.c.info(ix.X)
	in.Symbol = sym
	in.ID = id
	if len(ix.Args) != 1 || ix.Args[0].Name != nil {
		f.c.errorf("mms/arity", ix.Sp, "an objective takes one score holder: %s[@s]", name)
		return nil
	}
	var holder string
	switch h := ix.Args[0].Value.(type) {
	case *syntax.String:
		holder = h.Value
	case *syntax.Selector, *syntax.Index:
		sp := h.Span()
		holder = string(f.src[sp.Start:sp.End])
	default:
		f.c.errorf("mms/type-mismatch", h.Span(), "expected a selector such as @s, or a name in quotes")
		return nil
	}
	if holder == "" || strings.ContainsAny(holder, " \t") {
		f.c.errorf("mms/type-mismatch", ix.Args[0].Value.Span(), "a score holder cannot contain spaces")
		return nil
	}
	return &fn.Score{Holder: holder, Obj: sym.Path}
}

func (f *fc) intExpr(pre *[]fn.Op, e syntax.Expr) *fn.Operand {
	v := f.expr(pre, e)
	if v == nil {
		return nil
	}
	if v.typ != tInt {
		f.c.errorf("mms/type-mismatch", e.Span(), "expected an int, but this is %s", v.typ)
		return nil
	}
	return &v.op
}

// cond lowers a bool expression to a condition.
func (f *fc) cond(e syntax.Expr) fn.Cond {
	var pre []fn.Op
	v := f.expr(&pre, e)
	if v == nil {
		return fn.Cond{Pre: pre, Never: true}
	}
	if v.typ != tBool {
		f.c.errorf("mms/type-mismatch", e.Span(), "expected a condition, but this is %s; compare it, as in x > 0", v.typ)
		return fn.Cond{Pre: pre, Never: true}
	}
	return f.asCond(pre, v)
}

func (f *fc) asCond(pre []fn.Op, v *val) fn.Cond {
	if v.cond != nil {
		c := *v.cond
		c.Pre = append(pre, c.Pre...)
		return c
	}
	if v.op.IsConst {
		return fn.Cond{Pre: pre, Never: v.op.Const == 0}
	}
	return fn.Cond{Pre: pre, Parts: []string{fmt.Sprintf("if score %s matches 1", v.op.S)}}
}

func floorDiv(a, b int64) int64 {
	q := a / b
	if (a%b != 0) && ((a < 0) != (b < 0)) {
		q--
	}
	return q
}

func wrap(v int64) int32 { return int32(uint32(v)) }

func (f *fc) expr(pre *[]fn.Op, e syntax.Expr) *val {
	c := f.c
	if v, ok := f.buildTime(e); ok {
		return v
	}
	switch e := e.(type) {
	case *syntax.Number:
		n, err := strconv.ParseInt(e.Text, 10, 64)
		if err != nil || n > math.MaxInt32 || n < math.MinInt32 {
			c.errorf("mms/type-mismatch", e.Sp, "scores are whole numbers from -2147483648 to 2147483647")
			return nil
		}
		return &val{typ: tInt, op: fn.Const(int32(n))}
	case *syntax.Bool:
		v := int32(0)
		if e.Value {
			v = 1
		}
		return &val{typ: tBool, op: fn.Const(v)}
	case *syntax.Paren:
		return f.expr(pre, e.X)
	case *syntax.Ident, *syntax.ResLoc:
		l := f.variable(e)
		if l == nil {
			return nil
		}
		if l.typ == tString {
			return &val{typ: tString, path: l.path}
		}
		return &val{typ: l.typ, op: fn.Of(l.s)}
	case *syntax.String:
		v := e.Value
		return &val{typ: tString, str: &v}
	case *syntax.Index:
		s := f.objScore(e)
		if s == nil {
			return nil
		}
		return &val{typ: tInt, op: fn.Of(*s)}
	case *syntax.Call:
		return f.call(pre, e, true)
	case *syntax.Unary:
		return f.unary(pre, e)
	case *syntax.Binary:
		return f.binary(pre, e)
	}
	c.errorf("mms/unsupported", e.Span(), "this is not a value a function can compute with")
	return nil
}

func (f *fc) unary(pre *[]fn.Op, e *syntax.Unary) *val {
	switch e.Op {
	case "+":
		v := f.intExpr(pre, e.X)
		if v == nil {
			return nil
		}
		return &val{typ: tInt, op: *v}
	case "-":
		v := f.intExpr(pre, e.X)
		if v == nil {
			return nil
		}
		if v.IsConst {
			return &val{typ: tInt, op: fn.Const(wrap(-int64(v.Const)))}
		}
		t := f.temp()
		*pre = append(*pre, fn.Set{Dst: t, Src: fn.Const(0)}, fn.Arith{Dst: t, Op: "-=", Src: *v})
		return &val{typ: tInt, op: fn.Of(t)}
	case "!":
		var p []fn.Op
		v := f.expr(&p, e.X)
		if v == nil {
			return nil
		}
		if v.typ != tBool {
			f.c.errorf("mms/type-mismatch", e.X.Span(), "! needs a condition, but this is %s", v.typ)
			return nil
		}
		cd := f.asCond(p, v)
		return &val{typ: tBool, cond: f.not(cd)}
	}
	return nil
}

func (f *fc) not(cd fn.Cond) *fn.Cond {
	switch {
	case cd.Never:
		return &fn.Cond{Pre: cd.Pre}
	case cd.Always():
		return &fn.Cond{Pre: cd.Pre, Never: true}
	}
	if neg, ok := fn.Negate(cd.Parts); ok {
		return &fn.Cond{Pre: cd.Pre, Parts: neg}
	}
	t := f.temp()
	pre := append(cd.Pre, fn.Set{Dst: t, Src: fn.Const(0)}, fn.If{Cond: fn.Cond{Parts: cd.Parts}, Then: &fn.Block{Ops: []fn.Op{fn.Set{Dst: t, Src: fn.Const(1)}}}})
	return &fn.Cond{Pre: pre, Parts: []string{fmt.Sprintf("unless score %s matches 1", t)}}
}

func (f *fc) binary(pre *[]fn.Op, e *syntax.Binary) *val {
	switch e.Op {
	case "&&", "||":
		return f.logic(pre, e)
	case "<", "<=", ">", ">=", "==", "!=":
		return f.compare(pre, e)
	case "+":
		l, r := f.expr(pre, e.L), f.expr(pre, e.R)
		if l == nil || r == nil {
			return nil
		}
		if l.typ == tString || r.typ == tString {
			return f.concat(pre, l, r)
		}
		return f.arith(pre, e, l, r)
	case "-", "*", "/", "%":
	default:
		f.c.errorf("mms/unsupported", e.Sp, "'%s' is not supported in functions", e.Op)
		return nil
	}
	l, r := f.expr(pre, e.L), f.expr(pre, e.R)
	if l == nil || r == nil {
		return nil
	}
	return f.arith(pre, e, l, r)
}

func (f *fc) arith(pre *[]fn.Op, e *syntax.Binary, lv, rv *val) *val {
	for _, x := range []struct {
		v *val
		e syntax.Expr
	}{{lv, e.L}, {rv, e.R}} {
		if x.v.typ != tInt {
			f.c.errorf("mms/type-mismatch", x.e.Span(), "%s needs ints, but this is %s", e.Op, x.v.typ)
			return nil
		}
	}
	l, r := &lv.op, &rv.op
	if l.IsConst && r.IsConst {
		a, b := int64(l.Const), int64(r.Const)
		var v int64
		switch e.Op {
		case "+":
			v = a + b
		case "-":
			v = a - b
		case "*":
			v = a * b
		case "/", "%":
			if b == 0 {
				f.c.errorf("mms/division-by-zero", e.Sp, "division by zero")
				return nil
			}
			q := floorDiv(a, b)
			v = q
			if e.Op == "%" {
				v = a - q*b
			}
		}
		return &val{typ: tInt, op: fn.Const(wrap(v))}
	}
	if r.IsConst && r.Const == 0 && (e.Op == "/" || e.Op == "%") {
		f.c.errorf("mms/division-by-zero", e.Sp, "division by zero")
		return nil
	}
	t := f.temp()
	*pre = append(*pre, fn.Set{Dst: t, Src: *l}, fn.Arith{Dst: t, Op: e.Op + "=", Src: *r})
	return &val{typ: tInt, op: fn.Of(t)}
}

func (f *fc) compare(pre *[]fn.Op, e *syntax.Binary) *val {
	var p []fn.Op
	l := f.expr(&p, e.L)
	r := f.expr(&p, e.R)
	if l == nil || r == nil {
		return nil
	}
	if l.typ != r.typ || (l.typ != tInt && e.Op != "==" && e.Op != "!=") {
		f.c.errorf("mms/type-mismatch", e.Sp, "cannot compare %s with %s using %s", l.typ, r.typ, e.Op)
		return nil
	}
	if l.typ == tString {
		return f.compareStrings(p, e.Op, l, r)
	}
	lo, ro := f.operand(&p, l), f.operand(&p, r)
	if lo.IsConst && ro.IsConst {
		a, b := lo.Const, ro.Const
		var ok bool
		switch e.Op {
		case "<":
			ok = a < b
		case "<=":
			ok = a <= b
		case ">":
			ok = a > b
		case ">=":
			ok = a >= b
		case "==":
			ok = a == b
		case "!=":
			ok = a != b
		}
		return &val{typ: tBool, cond: &fn.Cond{Pre: p, Never: !ok}}
	}
	op := e.Op
	if lo.IsConst { // put the score on the left
		lo, ro = ro, lo
		op = map[string]string{"<": ">", "<=": ">=", ">": "<", ">=": "<=", "==": "==", "!=": "!="}[op]
	}
	var part string
	if ro.IsConst {
		v := int64(ro.Const)
		var rng string
		switch op {
		case "<":
			if v == math.MinInt32 {
				return &val{typ: tBool, cond: &fn.Cond{Pre: p, Never: true}}
			}
			rng = fmt.Sprintf("..%d", v-1)
		case "<=":
			rng = fmt.Sprintf("..%d", v)
		case ">":
			if v == math.MaxInt32 {
				return &val{typ: tBool, cond: &fn.Cond{Pre: p, Never: true}}
			}
			rng = fmt.Sprintf("%d..", v+1)
		case ">=":
			rng = fmt.Sprintf("%d..", v)
		default:
			rng = strconv.FormatInt(v, 10)
		}
		kw := "if"
		if op == "!=" {
			kw = "unless"
		}
		part = fmt.Sprintf("%s score %s matches %s", kw, lo.S, rng)
	} else {
		switch op {
		case "==":
			part = fmt.Sprintf("if score %s = %s", lo.S, ro.S)
		case "!=":
			part = fmt.Sprintf("unless score %s = %s", lo.S, ro.S)
		default:
			part = fmt.Sprintf("if score %s %s %s", lo.S, op, ro.S)
		}
	}
	return &val{typ: tBool, cond: &fn.Cond{Pre: p, Parts: []string{part}}}
}

// logic lowers && and ||. The right side runs only when it must, as its
// calls may have effects.
func (f *fc) logic(pre *[]fn.Op, e *syntax.Binary) *val {
	var lp, rp []fn.Op
	l := f.expr(&lp, e.L)
	r := f.expr(&rp, e.R)
	if l == nil || r == nil {
		return nil
	}
	if l.typ != tBool || r.typ != tBool {
		f.c.errorf("mms/type-mismatch", e.Sp, "%s needs conditions on both sides", e.Op)
		return nil
	}
	lc, rc := f.asCond(lp, l), f.asCond(rp, r)
	if e.Op == "&&" {
		switch {
		case lc.Never:
			return &val{typ: tBool, cond: &fn.Cond{Pre: lc.Pre, Never: true}}
		case lc.Always():
			return &val{typ: tBool, cond: &fn.Cond{Pre: append(lc.Pre, rc.Pre...), Parts: rc.Parts, Never: rc.Never}}
		case len(rc.Pre) == 0:
			return &val{typ: tBool, cond: &fn.Cond{Pre: lc.Pre, Parts: append(append([]string(nil), lc.Parts...), rc.Parts...), Never: rc.Never}}
		}
		t := f.temp()
		inner := &fn.Block{Ops: append(append([]fn.Op(nil), rc.Pre...), fn.If{Cond: fn.Cond{Parts: rc.Parts, Never: rc.Never}, Then: &fn.Block{Ops: []fn.Op{fn.Set{Dst: t, Src: fn.Const(1)}}}})}
		p := append(lc.Pre, fn.Set{Dst: t, Src: fn.Const(0)}, fn.If{Cond: fn.Cond{Parts: lc.Parts}, Then: inner})
		return &val{typ: tBool, cond: &fn.Cond{Pre: p, Parts: []string{fmt.Sprintf("if score %s matches 1", t)}}}
	}
	switch {
	case lc.Always():
		return &val{typ: tBool, cond: &fn.Cond{Pre: lc.Pre}}
	case lc.Never:
		return &val{typ: tBool, cond: &fn.Cond{Pre: append(lc.Pre, rc.Pre...), Parts: rc.Parts, Never: rc.Never}}
	}
	t := f.temp()
	set1 := &fn.Block{Ops: []fn.Op{fn.Set{Dst: t, Src: fn.Const(1)}}}
	p := append(lc.Pre, fn.Set{Dst: t, Src: fn.Const(0)}, fn.If{Cond: fn.Cond{Parts: lc.Parts}, Then: set1})
	right := &fn.Block{Ops: append(append([]fn.Op(nil), rc.Pre...), fn.If{Cond: fn.Cond{Parts: rc.Parts, Never: rc.Never}, Then: set1})}
	p = append(p, fn.If{Cond: fn.Cond{Parts: []string{fmt.Sprintf("unless score %s matches 1", t)}}, Then: right})
	return &val{typ: tBool, cond: &fn.Cond{Pre: p, Parts: []string{fmt.Sprintf("if score %s matches 1", t)}}}
}

// call lowers a call to an MMS function: arguments go to the callee's
// parameter scores, and the result comes back through `execute store`.
func (f *fc) call(pre *[]fn.Op, e *syntax.Call, wantValue bool) *val {
	c := f.c
	var id string
	switch x := e.Fn.(type) {
	case *syntax.Ident:
		id = f.ns + ":" + x.Name
	case *syntax.ResLoc:
		id = x.Text
		if !strings.Contains(id, ":") {
			id = "minecraft:" + id
		}
	default:
		c.errorf("mms/unsupported", e.Fn.Span(), "only functions can be called here")
		return nil
	}
	sym := c.findSig("function", id)
	if sym == nil || sym.Kind.MMS != "Function" {
		c.errorf("mms/unknown-name", e.Fn.Span(), "no Function named '%s' is declared", id)
		return nil
	}
	in := c.info(e.Fn)
	in.Symbol = sym
	in.ID = id
	sig := SigOf(sym.Decl)
	if len(e.Args) != len(sig.Params) {
		c.errorf("mms/arity", e.ArgsSp, "%s takes %d %s, but got %d", id, len(sig.Params), plural(len(sig.Params), "argument"), len(e.Args))
		return nil
	}
	calleeNS, calleePath, _ := strings.Cut(id, ":")
	calleeObj := fn.Objective(calleeNS)
	// A call that can come back to this function would overwrite its
	// scores: they are saved on a storage stack around the call.
	recursive := f.selfID != "" && f.reaches(id, f.selfID)
	// Arguments are evaluated before any is stored, so one argument
	// cannot see another's value when a function calls itself.
	var stores []fn.Op
	for i, a := range e.Args {
		if a.Name != nil {
			c.errorf("mms/unsupported", a.Name.Sp, "arguments are passed by position")
			return nil
		}
		v := f.expr(pre, a.Value)
		if v == nil {
			return nil
		}
		if v.typ != sig.Params[i].Type {
			c.errorf("mms/type-mismatch", a.Value.Span(), "parameter '%s' is %s, but this is %s", sig.Params[i].Name, sig.Params[i].Type, v.typ)
			return nil
		}
		if v.typ == tString {
			if recursive && v.str == nil {
				tmp := f.tempPath()
				*pre = append(*pre, f.setString(tmp, v))
				v = &val{typ: tString, path: tmp}
			}
			stores = append(stores, f.setString(fn.Storage(calleeNS)+" "+fn.LocalPath(calleePath, sig.Params[i].Name), v))
			continue
		}
		op := f.operand(pre, v)
		if !op.IsConst && recursive {
			t := f.temp()
			*pre = append(*pre, fn.Set{Dst: t, Src: op})
			op = fn.Of(t)
		}
		stores = append(stores, fn.Set{Dst: fn.Score{Holder: fn.Holder(calleePath, sig.Params[i].Name), Obj: calleeObj}, Src: op})
	}
	call := fn.Call{Fn: id, Args: stores, Recursive: recursive}
	if recursive {
		call.Frame = append([]fn.Score(nil), f.frame...)
		if f.strings {
			call.Strings = strings.TrimSuffix(fn.LocalPath(f.path, ""), ".")
		}
	}
	if !wantValue {
		*pre = append(*pre, call)
		return nil
	}
	if sig.Ret == "" {
		c.errorf("mms/type-mismatch", e.Sp, "%s returns nothing", id)
		return nil
	}
	t := f.temp()
	call.Dst = &t
	*pre = append(*pre, call)
	return &val{typ: sig.Ret, op: fn.Of(t)}
}

// Global checks `Global name: type = value` and returns the load-function
// line that sets the initial value, once: a reload keeps the value.
func (c *Checker) Global(d *syntax.Decl) []string {
	if d.Name == nil {
		return nil
	}
	typ := tInt
	if d.Type != nil {
		typ = d.Type.Text
		if typ != tInt && typ != tBool && typ != tString {
			c.errorf("mms/unknown-type", d.Type.Sp, "unknown type '%s'; globals are int, bool or string", typ)
			return nil
		}
	}
	ns, path := c.NS, d.Name.Text
	if i := strings.IndexByte(path, ':'); i >= 0 {
		ns, path = path[:i], path[i+1:]
	}
	s := fn.Score{Holder: "$" + path, Obj: fn.Objective(ns)}
	init := int32(0)
	str := ""
	if d.Value != nil {
		f := &fc{c: c, ns: ns, obj: fn.Objective(ns), names: map[string]int{}}
		f.push()
		var pre []fn.Op
		v := f.expr(&pre, d.Value)
		if v == nil {
			return nil
		}
		if v.typ != typ {
			c.errorf("mms/type-mismatch", d.Value.Span(), "expected %s, but this is %s", typ, v.typ)
			return nil
		}
		if typ == tString {
			if v.str == nil || len(pre) > 0 {
				c.errorf("mms/not-constant", d.Value.Span(), "a global's initial value must be a constant")
				return nil
			}
			str = *v.str
		} else {
			op := f.operand(&pre, v)
			if !op.IsConst || len(pre) > 0 {
				c.errorf("mms/not-constant", d.Value.Span(), "a global's initial value must be a constant")
				return nil
			}
			init = op.Const
		}
	}
	if typ == tString {
		loc := fn.Storage(ns) + " " + fn.GlobalPath(path)
		return []string{fmt.Sprintf("execute unless data storage %s run data modify storage %s set value %s", loc, loc, fn.QuoteSNBT(str))}
	}
	return []string{fmt.Sprintf("execute unless score %s = %s run scoreboard players set %s %d", s, s, s, init)}
}

// setString writes a string value to a storage location.
func (f *fc) setString(dst string, v *val) fn.Op {
	if v.str != nil {
		return fn.Cmd{Text: "data modify storage " + dst + " set value " + fn.QuoteSNBT(*v.str)}
	}
	return fn.Cmd{Text: "data modify storage " + dst + " set from storage " + v.path}
}

func (f *fc) tempPath() string {
	p := fn.Storage(f.ns) + " " + fn.LocalPath(f.path, "_s"+strconv.Itoa(f.temps))
	f.temps++
	f.strings = true
	return p
}

// text turns a value into a piece of text: fixed when it is constant,
// otherwise a macro argument.
func (f *fc) text(pre *[]fn.Op, v *val, key string) (string, *fn.MacroArg) {
	switch {
	case v.str != nil:
		return *v.str, nil
	case v.typ == tString:
		return "", &fn.MacroArg{Key: key, Path: v.path}
	case v.typ == tBool:
		cd := f.asCond(nil, v)
		switch {
		case cd.Never:
			*pre = append(*pre, cd.Pre...)
			return "false", nil
		case cd.Always():
			*pre = append(*pre, cd.Pre...)
			return "true", nil
		}
		p := f.tempPath()
		no, yes := "false", "true"
		*pre = append(*pre, f.setString(p, &val{str: &no}), fn.If{Cond: cd, Then: &fn.Block{Ops: []fn.Op{f.setString(p, &val{str: &yes})}}})
		return "", &fn.MacroArg{Key: key, Path: p}
	}
	op := f.operand(pre, v)
	if op.IsConst {
		return strconv.Itoa(int(op.Const)), nil
	}
	s := op.S
	return "", &fn.MacroArg{Key: key, Score: &s}
}

// concat joins two values into a string, through a macro when either is
// only known at run time.
func (f *fc) concat(pre *[]fn.Op, l, r *val) *val {
	lt, la := f.text(pre, l, "a0")
	rt, ra := f.text(pre, r, "a1")
	if la == nil && ra == nil {
		s := lt + rt
		return &val{typ: tString, str: &s}
	}
	esc := func(s string) string {
		q := fn.QuoteSNBT(s)
		return q[1 : len(q)-1]
	}
	var args []fn.MacroArg
	piece := func(t string, a *fn.MacroArg) string {
		if a == nil {
			return esc(t)
		}
		args = append(args, *a)
		return "$(" + a.Key + ")"
	}
	dst := f.tempPath()
	line := "data modify storage " + dst + " set value \"" + piece(lt, la) + piece(rt, ra) + "\""
	*pre = append(*pre, fn.Macro{Line: line, Args: args})
	return &val{typ: tString, path: dst}
}

// compareStrings tests two strings for equality: copying one over a copy
// of the other only succeeds when they differ.
func (f *fc) compareStrings(pre []fn.Op, op string, l, r *val) *val {
	if l.str != nil && r.str != nil {
		eq := *l.str == *r.str
		return &val{typ: tBool, cond: &fn.Cond{Pre: pre, Never: eq != (op == "==")}}
	}
	if l.str != nil {
		l, r = r, l
	}
	tmp := f.tempPath()
	t := f.temp()
	pre = append(pre, f.setString(tmp, l))
	src := "from storage " + r.path
	if r.str != nil {
		src = "value " + fn.QuoteSNBT(*r.str)
	}
	pre = append(pre, fn.Cmd{Text: fmt.Sprintf("execute store success score %s run data modify storage %s set %s", t, tmp, src)})
	kw := "if"
	if op == "!=" {
		kw = "unless"
	}
	return &val{typ: tBool, cond: &fn.Cond{Pre: pre, Parts: []string{fmt.Sprintf("%s score %s matches 0", kw, t)}}}
}

// interpolate lowers a command holding ${expr} parts. Constant parts are
// written into the line; the rest become macro arguments.
func (f *fc) interpolate(out *fn.Block, cmd *syntax.CommandStmt) {
	text := []rune(cmd.Text)
	base := cmd.Sp.Start + 1
	var line strings.Builder
	var args []fn.MacroArg
	var pre []fn.Op
	for i := 0; i < len(text); i++ {
		if text[i] != '$' || i+1 >= len(text) || text[i+1] != '{' {
			line.WriteRune(text[i])
			continue
		}
		depth, j := 0, i+1
		for ; j < len(text); j++ {
			if text[j] == '{' {
				depth++
			} else if text[j] == '}' {
				depth--
				if depth == 0 {
					break
				}
			}
		}
		if j >= len(text) {
			f.c.errorf("mms/syntax", syntax.Span{Start: base + i, End: base + len(text)}, "this ${ is not closed")
			return
		}
		e := f.parseExpr(string(text[i+2:j]), base+i+2)
		if e == nil {
			return
		}
		if s := f.customItem(e); s != nil {
			// A custom item is its base item with its components, as a
			// command's item argument: /give @s ${ember_shard} 4
			if st := f.c.customStack(s, e.Span()); st != nil {
				line.WriteString(StackSNBT(st))
			}
			i = j
			continue
		}
		v := f.expr(&pre, e)
		if v == nil {
			return
		}
		t, a := f.text(&pre, v, "a"+strconv.Itoa(len(args)))
		if a != nil {
			args = append(args, *a)
			line.WriteString("$(" + a.Key + ")")
		} else {
			line.WriteString(t)
		}
		i = j
	}
	out.Add(pre...)
	if len(args) == 0 {
		out.Add(fn.Cmd{Text: line.String()})
		return
	}
	out.Add(fn.Macro{Line: line.String(), Args: args})
}

// parseExpr parses an interpolated expression at offset off in the file,
// so its spans are the file's.
func (f *fc) parseExpr(text string, off int) syntax.Expr {
	const head = "X _ = "
	sp := syntax.Span{Start: off, End: off + len([]rune(text))}
	if strings.TrimSpace(text) == "" || off < len(head) {
		f.c.errorf("mms/syntax", sp, "expected an expression inside ${...}")
		return nil
	}
	file := syntax.Parse(strings.Repeat(" ", off-len(head)) + head + text)
	if len(file.Diags) > 0 || len(file.Decls) != 1 || file.Decls[0].Value == nil {
		f.c.errorf("mms/syntax", sp, "cannot read this expression")
		return nil
	}
	return file.Decls[0].Value
}
