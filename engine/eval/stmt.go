package eval

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// block expands a block's statements.
func (c *ctx) block(b *syntax.Block) *syntax.Block {
	if b == nil {
		return nil
	}
	out := &syntax.Block{Closed: b.Closed, Bare: b.Bare, Sp: c.span(b.Sp)}
	bc := c.with()
	bc.stmts(b.Items, &out.Items)
	return out
}

// stmts expands statements into out, in c's scope.
func (c *ctx) stmts(items []syntax.Stmt, out *[]syntax.Stmt) {
	for _, it := range items {
		if c.e.over {
			return
		}
		c.stmt(it, out)
	}
}

func (c *ctx) stmt(it syntax.Stmt, out *[]syntax.Stmt) {
	if !c.step(it.Span()) {
		return
	}
	add := func(st syntax.Stmt) { *out = append(*out, st) }
	switch st := it.(type) {
	case *syntax.Decl:
		if c.emit != nil {
			first := len(c.e.r.Instances)
			res := c.decl(st, nil)
			if st.Primary {
				c.markPrimary(st, res, first)
			}
			plain := st.Kind != nil && c.e.r.Blueprints[st.Kind.Name] == nil
			for _, d := range res {
				// Emitting the same declaration twice writes it once:
				// every custom item with a click hook needs the click
				// objective, but the namespace gets one.
				// A Blueprint instance's own output was weeded where
				// it was emitted.
				if plain && !d.Primary && c.inst != nil && c.e.emittedBefore(c.ns, d.D) {
					continue
				}
				*c.emit = append(*c.emit, d)
			}
			return
		}
		add(c.cloneStmt(st))
	case *syntax.ConstStmt:
		if st.Name == nil {
			return
		}
		if st.Value == nil {
			c.errorf("mms/const-value", st.Sp, "a const needs a value: const "+st.Name.Text+" = ...")
			return
		}
		c.sc.bind(st.Name.Text, c.reduce(st.Value))
	case *syntax.LetStmt:
		if st.Name == nil {
			return
		}
		if c.compile {
			if st.Value == nil {
				c.errorf("mms/let-value", st.Sp, "a let here needs a value")
				return
			}
			c.sc.bind(st.Name.Text, c.reduce(st.Value))
			return
		}
		v := c.lower(c.reduce(st.Value))
		c.sc.hide(st.Name.Text)
		add(&syntax.LetStmt{Name: c.name(st.Name), Type: c.typ(st.Type), Value: v, Doc: st.Doc, Sp: c.span(st.Sp)})
	case *syntax.FieldStmt:
		if st.Key == nil {
			return
		}
		key := &syntax.Name{Text: st.Key.Text, Sp: c.span(st.Key.Sp)}
		if st.KeyQuoted {
			key.Text, _ = c.interp(st.Key.Text, st.Key.Sp)
		}
		v := c.reduce(st.Value)
		if dropped(v) {
			return // an unset optional field
		}
		add(&syntax.FieldStmt{Key: key, KeyQuoted: st.KeyQuoted, Value: c.lower(v), Sp: c.span(st.Sp)})
	case *syntax.IfStmt:
		c.ifStmt(st, out)
	case *syntax.ForStmt:
		c.forStmt(st, out)
	case *syntax.WhileStmt:
		add(&syntax.WhileStmt{Cond: c.lower(c.reduce(st.Cond)), Body: c.block(st.Body), Sp: c.span(st.Sp)})
	case *syntax.ReturnStmt:
		add(&syntax.ReturnStmt{Value: c.lower(c.reduce(st.Value)), Sp: c.span(st.Sp)})
	case *syntax.CommandStmt:
		t, _ := c.interp(st.Text, st.Sp)
		add(&syntax.CommandStmt{Text: t, Sp: c.span(st.Sp)})
	case *syntax.AssignStmt:
		add(&syntax.AssignStmt{Target: c.lower(c.target(st.Target)), Op: st.Op, Value: c.lower(c.reduce(st.Value)), Sp: c.span(st.Sp)})
	case *syntax.ExprStmt:
		// Calling a Function value that holds `run { ... }` runs its
		// statements in place: `action()` in a Blueprint's emitted code.
		if call, ok := st.X.(*syntax.Call); ok && len(call.Args) == 0 && !c.compile {
			if id, ok := call.Fn.(*syntax.Ident); ok {
				if v, ok, _ := c.sc.lookup(id.Name); ok {
					if code, ok := v.(*Code); ok {
						if run, ok := code.X.(*syntax.RunExpr); ok && run.Body != nil {
							for _, s := range run.Body.Items {
								add(syntax.CloneStmt(s))
							}
							return
						}
					}
				}
			}
		}
		v := c.reduce(st.X)
		switch v := v.(type) {
		case *Stmts:
			for _, s := range v.Items {
				add(syntax.CloneStmt(s))
			}
		case *None:
		default:
			if x := c.lower(v); x != nil {
				add(&syntax.ExprStmt{X: x, Sp: c.span(st.Sp)})
			}
		}
	case *syntax.Block:
		add(c.block(st))
	case *syntax.ExpectStmt:
		add(&syntax.ExpectStmt{X: c.lower(c.reduce(st.X)), Sp: c.span(st.Sp)})
	case *syntax.RunStmt:
		add(&syntax.RunStmt{X: c.lower(c.reduce(st.X)), Sp: c.span(st.Sp)})
	case *syntax.SectionStmt:
		sc := c.with()
		sc.compile = false
		add(&syntax.SectionStmt{Word: &syntax.Ident{Name: st.Word.Name, Sp: c.span(st.Word.Sp)}, Body: sc.block(st.Body), Sp: c.span(st.Sp)})
	default:
		add(c.cloneStmt(st))
	}
}

func (c *ctx) cloneStmt(st syntax.Stmt) syntax.Stmt {
	if c.site != nil {
		return syntax.CloneStmtAt(st, *c.site)
	}
	return syntax.CloneStmt(st)
}

func (c *ctx) name(n *syntax.Name) *syntax.Name {
	if n == nil {
		return nil
	}
	return &syntax.Name{Text: n.Text, Sp: c.span(n.Sp)}
}

func (c *ctx) typ(t syntax.TypeExpr) syntax.TypeExpr {
	if t == nil {
		return nil
	}
	if c.site != nil {
		return syntax.CloneTypeAt(t, *c.site)
	}
	return syntax.CloneType(t)
}

// target evaluates an assignment target: names in it are left alone, so
// `x = 1` assigns x even where x is also a field, but indexes and
// members are expanded.
func (c *ctx) target(x syntax.Expr) Value {
	if id, ok := x.(*syntax.Ident); ok {
		return &Code{c.pos(id.Sp), c.clone(id)}
	}
	return c.reduce(x)
}

func (c *ctx) ifStmt(st *syntax.IfStmt, out *[]syntax.Stmt) {
	cond := c.reduce(st.Cond)
	if b, ok := cond.(*Bool); ok {
		// Known at compile time: keep only the branch taken.
		var branch syntax.Stmt
		if b.V {
			branch = st.Then
		} else {
			branch = st.Else
		}
		switch br := branch.(type) {
		case *syntax.Block:
			c.with().stmts(br.Items, out)
		case *syntax.IfStmt:
			c.ifStmt(br, out)
		}
		return
	}
	if known(cond) {
		c.errorf("mms/condition", st.Cond.Span(), "a condition is true or false, not "+kindName(cond))
		return
	}
	if c.check || c.emit != nil {
		c.errorf("mms/condition", st.Cond.Span(), "this condition must be known at compile time, but it is "+kindName(cond))
		return
	}
	n := &syntax.IfStmt{Cond: c.lower(cond), Then: c.block(st.Then), Sp: c.span(st.Sp)}
	switch e := st.Else.(type) {
	case *syntax.Block:
		n.Else = c.block(e)
	case *syntax.IfStmt:
		var more []syntax.Stmt
		c.ifStmt(e, &more)
		switch len(more) {
		case 0:
		case 1:
			n.Else = more[0]
		default:
			n.Else = &syntax.Block{Items: more, Closed: true, Sp: c.span(e.Sp)}
		}
	}
	*out = append(*out, n)
}

// forStmt unrolls a loop over a list or record, and in compile-time code
// over a range. A range loop in function code stays a loop.
func (c *ctx) forStmt(st *syntax.ForStmt, out *[]syntax.Stmt) {
	if st.Var == nil || st.Iter == nil {
		return
	}
	iter := c.reduce(st.Iter)
	if _, ok := iter.(*RangeV); ok && !c.compile {
		lc := c.with()
		lc.sc.hide(st.Var.Text)
		if out != nil {
			*out = append(*out, &syntax.ForStmt{Var: c.name(st.Var), Var2: c.name(st.Var2), Iter: c.lower(iter),
				Body: lc.block(st.Body), Sp: c.span(st.Sp)})
		}
		return
	}
	if cd, ok := iter.(*Code); ok && !c.compile && c.emit == nil {
		// A loop the function compiler handles (and reports).
		lc := c.with()
		lc.sc.hide(st.Var.Text)
		if out != nil {
			*out = append(*out, &syntax.ForStmt{Var: c.name(st.Var), Var2: c.name(st.Var2), Iter: cd.X,
				Body: lc.block(st.Body), Sp: c.span(st.Sp)})
		}
		return
	}
	items, ok := c.iterate(iter, st.Iter.Span(), true)
	if !ok || st.Body == nil {
		return
	}
	for _, it := range items {
		if c.e.over {
			return
		}
		ic := c.with()
		ic.bindLoop(st.Var, st.Var2, it)
		if ic.emit != nil {
			bind := st.Var.Text + " = " + showValue(it.one)
			if st.Var2 != nil {
				bind = st.Var.Text + ", " + st.Var2.Text + " = " + showValue(it.key) + ", " + showValue(it.val)
			}
			ic.fr = append(append([]Frame(nil), c.fr...), Frame{Kind: FrameFor, Name: bind,
				Call: Origin{c.path, c.span(st.Sp)}, Def: Origin{c.defPath, st.Sp}})
		}
		var discard []syntax.Stmt
		target := out
		if target == nil {
			target = &discard
		}
		ic.stmts(st.Body.Items, target)
		if len(discard) > 0 {
			ic.errorf("mms/top-level-for", st.Body.Sp, "a top-level for holds declarations, const and nested for and if")
		}
	}
}

// showValue spells a loop value for the editor: `oak`, `3`, `"en_us"`.
func showValue(v Value) string {
	switch v := v.(type) {
	case *Str:
		return strconv.Quote(v.S)
	case *Word, *Num, *Bool:
		t, _ := text(v)
		return t
	case *None:
		return "none"
	case *List:
		return "[...]"
	case *Record:
		return "{...}"
	}
	return "..."
}

// decl expands one declaration: a top-level one (from == d), one in a
// top-level for, or one emitted by a Blueprint.
func (c *ctx) decl(d *syntax.Decl, from *syntax.Decl) []*Decl {
	if d.Kind == nil {
		return nil
	}
	if from == nil {
		from = d
		if c.inst != nil {
			from = c.inst.D
		}
	}
	if definitionKind(d.Kind.Name) {
		c.errorf("mms/nested-definition", d.Kind.Sp, "a "+d.Kind.Name+" is declared at the top of a file")
		return nil
	}
	if bp := c.e.r.Blueprints[d.Kind.Name]; bp != nil {
		var out []*Decl
		for _, n := range d.Names() {
			out = append(out, c.instance(d, n, bp, from)...)
		}
		return out
	}
	nd := &syntax.Decl{Kind: &syntax.Ident{Name: d.Kind.Name, Sp: c.span(d.Kind.Sp)}, Type: c.typ(d.Type),
		TypeValue: c.typ(d.TypeValue), Extern: d.Extern, Doc: d.Doc, Sp: c.span(d.Sp)}
	if c.inst != nil {
		nd.Sp = c.inst.D.Sp
	}
	// An objective's criterion written as a string, as in a top-level
	// for: `Objective mined_${w}: "minecraft.mined:minecraft.${w}_log"`.
	if ts, ok := nd.Type.(*syntax.TypeString); ok && d.Kind.Name == "Objective" {
		ts.Value, _ = c.interp(ts.Value, ts.Sp)
	}
	if d.Name != nil {
		nd.Name = c.declName(d.Name)
		if nd.Name == nil {
			return nil
		}
	}
	for _, m := range d.More {
		if n := c.declName(m); n != nil {
			nd.More = append(nd.More, n)
		}
	}
	bc := c.with()
	bc.emit = nil
	if fnKinds[d.Kind.Name] {
		bc.compile = false
	}
	if d.Params != nil {
		nd.Params = &syntax.ParamList{Sp: c.span(d.Params.Sp)}
		for _, p := range d.Params.Params {
			np := &syntax.Param{Name: c.name(p.Name), Type: c.typ(p.Type), Default: c.lower(c.reduce(p.Default)), Sp: c.span(p.Sp)}
			nd.Params.Params = append(nd.Params.Params, np)
			if p.Name != nil {
				bc.sc.hide(p.Name.Text)
			}
		}
	}
	if d.Value != nil {
		v := bc.reduce(d.Value)
		nd.Value = bc.lower(v)
		if nd.Value == nil {
			nd.Value = &syntax.Ident{Name: "none", Sp: c.span(d.Value.Span())}
		}
	}
	if d.Body != nil {
		nd.Body = bc.block(d.Body)
	}
	at := Origin{c.path, nd.Sp}
	if d.Name != nil {
		at = Origin{c.path, nd.Name.Sp}
	}
	prov := Prov{At: at, Frames: c.fr}
	if c.site != nil {
		prov.Def = &Origin{c.defPath, d.Sp}
	}
	return []*Decl{{D: nd, Prov: prov, From: from, Instance: c.inst}}
}

// declName spells a declaration's name, filling a templated one.
func (c *ctx) declName(n *syntax.Name) *syntax.Name {
	if !n.Templated() {
		return &syntax.Name{Text: n.Text, Sp: c.span(n.Sp)}
	}
	t, ok := c.nameParts(n.Parts, n.Sp)
	if !ok {
		return nil
	}
	return &syntax.Name{Text: t, Sp: c.span(n.Sp)}
}

// call evaluates a call: a Template's body, an intrinsic, or a built-in
// helper left for the checker.
func (c *ctx) call(x *syntax.Call) Value {
	p := c.pos(x.Sp)
	if id, ok := x.Fn.(*syntax.Ident); ok {
		_, bound, hidden := c.sc.lookup(id.Name)
		if !bound && !hidden {
			if t := c.e.r.Templates[id.Name]; t != nil {
				if t.D.Extern {
					if fn := intrinsics[id.Name]; fn != nil {
						return fn(c, x, c.argValues(x.Args))
					}
					return &Code{p, &syntax.Call{Fn: c.clone(id), Args: c.args(x.Args), ArgsSp: c.span(x.ArgsSp), Sp: c.span(x.Sp)}}
				}
				return c.apply(t, x)
			}
			if fn := builtins[id.Name]; fn != nil && (c.check || !checkOnly[id.Name]) {
				return fn(c, x, c.argValues(x.Args))
			}
		}
	}
	fn := c.reduce(x.Fn)
	return &Code{p, &syntax.Call{Fn: c.lower(fn), Args: c.args(x.Args), ArgsSp: c.span(x.ArgsSp), Sp: c.span(x.Sp)}}
}

// argValue is one evaluated argument.
type argValue struct {
	name string
	v    Value
	x    *syntax.Arg
}

func (c *ctx) argValues(as []*syntax.Arg) []argValue {
	out := make([]argValue, len(as))
	for i, a := range as {
		out[i] = argValue{v: c.reduce(a.Value), x: a}
		if a.Name != nil {
			out[i].name = a.Name.Text
		}
	}
	return out
}

// apply runs a Template: its parameters bound to the arguments, its body
// evaluated where it is written, and everything it builds placed at the
// call.
func (c *ctx) apply(t *Template, x *syntax.Call) Value {
	p := c.pos(x.Sp)
	name := t.D.NameText()
	if c.depth >= maxDepth {
		c.errorf("mms/eval-depth", x.Sp, "Template '"+name+"' calls itself too deeply")
		return &Code{p, &syntax.Bad{Sp: c.span(x.Sp)}}
	}
	// Arguments are compile-time values: a bare name is a name, not a
	// function variable, unless a variable of that name is in scope.
	ac := *c
	ac.compile = true
	args := ac.argValues(x.Args)
	if !c.compile {
		// A Global or Storage field given in function code is its value
		// when the code runs, as a local is.
		for i, a := range x.Args {
			if c.namesValue(a.Value) {
				args[i].v = c.reduce(a.Value)
			}
		}
	}
	site := c.span(x.Sp)
	tc := &ctx{e: c.e, path: c.path, defPath: t.Path, site: &site, sc: &scope{}, ns: c.ns, compile: true,
		depth: c.depth + 1, fr: append(append([]Frame(nil), c.fr...), Frame{Kind: FrameTemplate, Name: name,
			Call: Origin{c.path, site}, Def: Origin{t.Path, t.D.Sp}})}
	var call *Call
	if c.site == nil {
		call = &Call{Path: c.path, Template: t, X: x, Args: map[string]syntax.Expr{}}
		for param, i := range argIndex(t.D, x) {
			if a := c.lower(args[i].v); a != nil {
				call.Args[param] = a
			}
			if _, code := args[i].v.(*Code); code && !c.compile && isName(args[i].x.Value) {
				if call.Runtime == nil {
					call.Runtime = map[string]bool{}
				}
				call.Runtime[param] = true
			}
		}
		c.e.r.Calls = append(c.e.r.Calls, call)
	}
	if !c.bindParams(tc, t.D, "Template '"+name+"'", args, x) {
		return &Code{p, &syntax.Bad{Sp: site}}
	}
	switch {
	case t.D.Value != nil:
		v := tc.reduce(t.D.Value)
		if c.e.over {
			return &Code{p, &syntax.Bad{Sp: site}}
		}
		if call != nil {
			call.Result = tc.lower(v)
		}
		return v
	case t.D.Body != nil:
		st := &Stmts{pos: tc.pos(t.D.Body.Sp)}
		tc.stmts(t.D.Body.Items, &st.Items)
		if call != nil {
			call.Expanded = st.Items
		}
		return st
	}
	c.errorf("mms/template-body", x.Sp, "Template '"+name+"' has no body")
	return &Code{p, &syntax.Bad{Sp: site}}
}

// bindParams binds a Template's parameters in tc's scope: positional
// arguments in order, then named ones, then defaults, which may use the
// parameters before them.
func (c *ctx) bindParams(tc *ctx, d *syntax.Decl, what string, args []argValue, call *syntax.Call) bool {
	var params []*syntax.Param
	if d.Params != nil {
		params = d.Params.Params
	}
	got := map[string]Value{}
	ok := true
	i := 0
	for _, a := range args {
		if a.name != "" {
			found := false
			for _, pr := range params {
				if pr.Name != nil && pr.Name.Text == a.name {
					found = true
				}
			}
			if !found {
				c.errorf("mms/unknown-arg", a.x.Name.Sp, what+" has no parameter '"+a.name+"'")
				ok = false
				continue
			}
			if _, dup := got[a.name]; dup {
				c.errorf("mms/duplicate-arg", a.x.Name.Sp, "'"+a.name+"' is given twice")
				ok = false
			}
			got[a.name] = a.v
			continue
		}
		for i < len(params) && (params[i].Name == nil || got[params[i].Name.Text] != nil) {
			i++
		}
		if i >= len(params) {
			c.errorf("mms/too-many-args", a.x.Sp, what+" takes "+strconv.Itoa(len(params))+" argument"+plural(len(params)))
			ok = false
			continue
		}
		got[params[i].Name.Text] = a.v
		i++
	}
	var missing []string
	for _, pr := range params {
		if pr.Name == nil {
			continue
		}
		if v, have := got[pr.Name.Text]; have {
			tc.sc.bind(pr.Name.Text, tc.withDefaults(v, pr.Type))
			continue
		}
		switch {
		case pr.Default != nil:
			tc.sc.bind(pr.Name.Text, tc.reduce(pr.Default))
		case isOptional(pr.Type):
			tc.sc.bind(pr.Name.Text, &None{pos: tc.pos(pr.Sp)})
		default:
			missing = append(missing, pr.Name.Text)
		}
	}
	if len(missing) > 0 {
		c.errorf("mms/missing-arg", call.Sp, what+" needs "+strings.Join(quoteAll(missing), ", "))
		ok = false
	}
	return ok
}

func isOptional(t syntax.TypeExpr) bool {
	switch t := t.(type) {
	case *syntax.TypeOptional:
		return true
	case *syntax.TypeParen:
		return isOptional(t.X)
	case *syntax.TypeUnion:
		for _, a := range t.Arms {
			if isOptional(a) {
				return true
			}
		}
	}
	return false
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func quoteAll(ss []string) []string {
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = "'" + s + "'"
	}
	return out
}

// argIndex matches a call's arguments to the parameters they give.
func argIndex(d *syntax.Decl, x *syntax.Call) map[string]int {
	out := map[string]int{}
	if d.Params == nil {
		return out
	}
	ps := d.Params.Params
	i := 0
	for j, a := range x.Args {
		if a.Name != nil {
			out[a.Name.Text] = j
			continue
		}
		for i < len(ps) && ps[i].Name != nil {
			if _, taken := out[ps[i].Name.Text]; !taken {
				break
			}
			i++
		}
		for i < len(ps) && ps[i].Name == nil {
			i++
		}
		if i < len(ps) {
			out[ps[i].Name.Text] = j
			i++
		}
	}
	return out
}

// namesValue reports whether x is a Global or Storage, or a field of
// one, that no compile-time name hides.
func (c *ctx) namesValue(x syntax.Expr) bool {
	switch x := x.(type) {
	case *syntax.Ident:
		if _, ok, _ := c.sc.lookup(x.Name); ok || c.e.r.Consts[x.Name] != nil {
			return false
		}
		return c.e.values[x.Name]
	case *syntax.Member:
		return c.namesValue(x.X)
	case *syntax.Paren:
		return c.namesValue(x.X)
	}
	return false
}

// isName reports whether x is a name or a field of one: `y`, `game.wave`.
func isName(x syntax.Expr) bool {
	switch x := x.(type) {
	case *syntax.Ident:
		return true
	case *syntax.Member:
		return isName(x.X)
	case *syntax.Paren:
		return isName(x.X)
	}
	return false
}
