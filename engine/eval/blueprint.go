package eval

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// instance expands one declaration of a Blueprint's kind: it binds the
// fields, runs the check rules and returns what the emit blocks write.
func (c *ctx) instance(d *syntax.Decl, n *syntax.Name, bp *Blueprint, from *syntax.Decl) []*Decl {
	name := c.declName(n)
	if name == nil {
		return nil
	}
	bpName := bp.D.NameText()
	if c.depth >= maxDepth {
		c.errorf("mms/eval-depth", n.Sp, "Blueprint '"+bpName+"' emits itself too deeply")
		return nil
	}
	ns, path := c.ns, name.Text
	if i := strings.IndexByte(path, ':'); i >= 0 {
		ns, path = path[:i], path[i+1:]
	}
	id := ns + ":" + path
	site := c.span(n.Sp)
	inst := &Instance{Path: c.path, D: from, Name: name, Blueprint: bp, fieldSp: map[string]syntax.Span{}}
	c.e.r.Instances = append(c.e.r.Instances, inst)
	if d.Value != nil {
		c.errorf("mms/blueprint-body", d.Value.Span(), "a "+bpName+" takes a { } body, not '='")
	}

	inst.User = c.site == nil
	// The fields the declaration sets, evaluated where it is written.
	given := &Record{pos: c.pos(n.Sp)}
	if d.Body != nil {
		for _, it := range d.Body.Items {
			f, ok := it.(*syntax.FieldStmt)
			if !ok {
				c.errorf("mms/blueprint-body", it.Span(), "a "+bpName+" body holds its fields: name = value")
				continue
			}
			if f.Key == nil {
				continue
			}
			key := f.Key.Text
			if fieldDecl(bp, key) == nil {
				c.errorf("mms/unknown-field", f.Key.Sp, bpName+" has no field '"+key+"'")
				continue
			}
			if _, dup := inst.fieldSp[key]; dup {
				c.errorf("mms/duplicate-field", f.Key.Sp, "'"+key+"' is set twice")
				continue
			}
			inst.fieldSp[key] = c.span(f.Sp)
			v := c.reduce(f.Value)
			given.set(key, v)
			if c.site == nil {
				if inst.Written == nil {
					inst.Written = map[string]syntax.Expr{}
				}
				if x := c.lower(v); x != nil {
					inst.Written[key] = x
				}
			}
		}
	}

	bc := &ctx{e: c.e, path: c.path, defPath: bp.Path, site: &site, sc: &scope{}, ns: ns, compile: true, inst: inst,
		depth: c.depth + 1, fr: append(append([]Frame(nil), c.fr...), Frame{Kind: FrameBlueprint, Name: bpName,
			Call: Origin{c.path, site}, Def: Origin{bp.Path, bp.D.Sp}})}
	sp := bc.pos(bp.D.Sp)
	inst.Self = &Record{pos: sp, Keys: []string{"id", "namespace", "path"},
		Vals: []Value{&Word{sp, id}, &Str{sp, ns}, &Str{sp, path}}, id: &Word{sp, id}}
	bc.sc.bind("self", inst.Self)

	inst.Fields = &Record{pos: sp}
	var missing []string
	for _, f := range bp.Fields {
		if f.Name == nil {
			continue
		}
		v, ok := given.get(f.Name.Text)
		switch {
		case ok:
		case f.Default != nil:
			v = bc.reduce(f.Default)
		case isOptional(f.Type):
			v = &None{pos: bc.pos(f.Sp)}
		default:
			missing = append(missing, f.Name.Text)
			v = &None{pos: bc.pos(f.Sp)}
		}
		bc.sc.bind(f.Name.Text, v)
		inst.Fields.set(f.Name.Text, v)
	}
	if len(missing) > 0 {
		c.e.report(c.path, syntax.Diagnostic{Code: "mms/missing-field", Severity: syntax.SeverityError, Sp: site,
			Message: bpName + " " + name.Text + " needs " + strings.Join(quoteAll(missing), ", ")})
		return nil
	}
	for _, l := range bp.Lets {
		var discard []syntax.Stmt
		bc.stmt(l, &discard)
	}

	for _, as := range bp.As {
		if as.Type == nil || as.Value == nil {
			continue
		}
		conv := &Conversion{Type: as.Type, Path: bp.Path, X: bc.lower(bc.reduce(as.Value))}
		if conv.X == nil {
			continue
		}
		if as.Hint != nil {
			args := bc.argValues(as.Hint.Args)
			if len(args) == 2 {
				r, ok1 := args[0].v.(*Str)
				m, ok2 := args[1].v.(*Str)
				if ok1 && ok2 {
					conv.Rule, conv.Hint = r.S, m.S
				}
			}
			if conv.Hint == "" {
				bc.errorf("mms/hint-args", as.Hint.Sp, "hint takes a rule name and a message: hint(\"rule\", \"message\")")
			}
		}
		inst.Conversions = append(inst.Conversions, conv)
	}

	for _, blk := range bp.Checks {
		cc := bc.with()
		cc.check = true
		var discard []syntax.Stmt
		cc.stmts(blk.Items, &discard)
	}

	var emitted []*Decl
	for _, blk := range bp.Emits {
		ec := bc.with()
		ec.emit = &emitted
		var discard []syntax.Stmt
		ec.stmts(blk.Items, &discard)
		if len(discard) > 0 {
			bc.errorf("mms/emit-body", blk.Sp, "an emit block holds declarations, let, if and for")
		}
	}
	var pid string
	switch {
	case inst.Primary != nil:
		pid = inst.Primary.NameText()
	case inst.Forward != nil:
		pid = inst.Forward.Name.Text
	}
	if pid != "" {
		if !strings.Contains(pid, ":") {
			pid = ns + ":" + pid
		}
		if pid != id {
			c.e.report(c.path, syntax.Diagnostic{Code: "mms/primary-id", Severity: syntax.SeverityError, Sp: site,
				Message: "the primary declaration of a " + bpName + " has its id: " + id + ", not " + pid})
		}
	}
	return emitted
}

func fieldDecl(bp *Blueprint, name string) *syntax.FieldDecl {
	for _, f := range bp.Fields {
		if f.Name != nil && f.Name.Text == name {
			return f
		}
	}
	return nil
}

// markPrimary records the declaration a `primary` statement emitted.
// When that statement is itself a Blueprint instance (the first one
// expanded from index first on), the name forwards on to it.
func (c *ctx) markPrimary(st *syntax.Decl, out []*Decl, first int) {
	if c.inst == nil {
		c.errorf("mms/primary", st.Kind.Sp, "'primary' marks a declaration in a Blueprint's emit block")
		return
	}
	if c.inst.Primary != nil || c.inst.Forward != nil {
		c.errorf("mms/primary", st.Kind.Sp, "a Blueprint has at most one primary declaration")
		return
	}
	if c.e.r.Blueprints[st.Kind.Name] != nil {
		if first < len(c.e.r.Instances) {
			c.inst.Forward = c.e.r.Instances[first]
		}
		return
	}
	for _, d := range out {
		if d.Instance == c.inst {
			c.inst.Primary = d.D
			d.Primary = true
			return
		}
	}
}

// checkRule reports a Blueprint check: error(field, message),
// warn(field, rule, message) or hint(...). The finding goes on the field
// the user wrote, or on the declaration's name when the field took its
// default.
func checkRule(sev syntax.Severity) intrinsic {
	return func(c *ctx, x *syntax.Call, args []argValue) Value {
		p := c.pos(x.Sp)
		if len(args) < 2 || len(args) > 3 {
			c.errorf("mms/check-args", x.Sp, "a check rule takes a field and a message, with a rule name between them: warn(field, \"rule\", \"message\")")
			return &None{pos: p}
		}
		sp := c.span(x.Sp)
		if id, ok := args[0].x.Value.(*syntax.Ident); ok && c.inst != nil {
			if fsp, ok := c.inst.fieldSp[id.Name]; ok {
				sp = fsp
			}
		}
		// A rule without a name is reported under the Blueprint's:
		// mms/custom-item for CustomItem.
		code := "check"
		if c.inst != nil {
			code = "mms/" + kebab(c.inst.Blueprint.D.NameText())
		}
		if len(args) == 3 {
			r, ok := args[1].v.(*Str)
			if !ok {
				c.errorf("mms/check-args", x.Sp, "a rule name is a string")
				return &None{pos: p}
			}
			code = r.S
		}
		msg, ok := args[len(args)-1].v.(*Str)
		if !ok {
			c.errorf("mms/check-args", x.Sp, "a check message is a string")
			return &None{pos: p}
		}
		c.e.report(c.path, syntax.Diagnostic{Code: code, Severity: sev, Message: msg.S, Sp: sp})
		return &None{pos: p}
	}
}

// kebab spells a PascalCase name in kebab case: CustomItem is custom-item.
func kebab(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('-')
			}
			r += 'a' - 'A'
		}
		b.WriteRune(r)
	}
	return b.String()
}
