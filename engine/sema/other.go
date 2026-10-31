package sema

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Tag checks a tag declaration's entries: a list of ids and tags of the
// tag's registry. Declarations of the same tag are merged by the emitter.
func (c *Checker) Tag(d *syntax.Decl, k Kind) (values jsonv.Array, replace bool) {
	var elems []syntax.Expr
	switch x := d.Value.(type) {
	case *syntax.List:
		elems = x.Elems
	case *syntax.Object:
		// The whole tag file: { replace = true, values = [...] }.
		for _, it := range x.Body.Items {
			f, ok := it.(*syntax.FieldStmt)
			if !ok {
				c.errorf("mms/type-mismatch", it.Span(), "expected 'values' or 'replace'")
				continue
			}
			switch f.Key.Text {
			case "values":
				if l, ok := f.Value.(*syntax.List); ok {
					elems = l.Elems
				} else {
					c.errorf("mms/type-mismatch", f.Value.Span(), "expected a list of %s ids", k.TagOf)
				}
			case "replace":
				if b, ok := f.Value.(*syntax.Bool); ok {
					replace = b.Value
				} else {
					c.errorf("mms/type-mismatch", f.Value.Span(), "expected true or false")
				}
			default:
				c.errorf("mms/unknown-field", f.Key.Sp, "a tag has only 'values' and 'replace'")
			}
		}
	case nil:
		if d.Body != nil {
			for _, it := range d.Body.Items {
				if s, ok := it.(*syntax.ExprStmt); ok {
					elems = append(elems, s.X)
				} else {
					c.errorf("mms/type-mismatch", it.Span(), "expected a %s id", k.TagOf)
				}
			}
		} else {
			c.errorf("mms/missing-value", d.Kind.Sp, "%s needs a list of ids: = [a, b]", d.Kind.Name)
		}
	default:
		elems = []syntax.Expr{x}
	}
	id := idInfo{Registry: k.TagOf, Tags: "allowed", Present: true}
	values = jsonv.Array{}
	for _, e := range elems {
		if v := c.tagEntry(e, k, id); v != nil {
			values = append(values, v)
		}
	}
	return values, replace
}

// tagEntry checks one tag entry: an id, a #tag, either as a string, or
// { id = ..., required = false }.
func (c *Checker) tagEntry(e syntax.Expr, k Kind, id idInfo) jsonv.Value {
	c.info(e).Type = tview{}
	switch x := e.(type) {
	case *syntax.Ident:
		return jsonv.String(c.resolveID(x.Name, id, e))
	case *syntax.ResLoc:
		return jsonv.String(c.resolveID(x.Text, id, e))
	case *syntax.Tag:
		return jsonv.String(c.resolveTag(x.Name, id, e))
	case *syntax.String:
		if rest, ok := strings.CutPrefix(x.Value, "#"); ok {
			return jsonv.String(c.resolveTag(rest, id, e))
		}
		return jsonv.String(c.resolveID(x.Value, id, e))
	case *syntax.Call:
		// Optional(x) is { id = x, required = false }: the tag still
		// loads when x does not exist.
		if fn, ok := x.Fn.(*syntax.Ident); ok && fn.Name == "Optional" {
			if len(x.Args) != 1 || x.Args[0].Name != nil {
				c.errorf("mms/type-mismatch", x.Sp, "Optional takes one id or tag: Optional(%s)", "some_id")
				return nil
			}
			v := c.tagEntry(x.Args[0].Value, k, id)
			if v == nil {
				return nil
			}
			o := &jsonv.Object{}
			o.Set("id", v)
			o.Set("required", jsonv.Bool(false))
			return o
		}
	case *syntax.Object:
		o := &jsonv.Object{}
		for _, it := range x.Body.Items {
			f, ok := it.(*syntax.FieldStmt)
			if !ok {
				continue
			}
			switch f.Key.Text {
			case "id":
				if v := c.tagEntry(f.Value, k, id); v != nil {
					o.Set("id", v)
				}
			case "required":
				if b, ok := f.Value.(*syntax.Bool); ok {
					o.Set("required", jsonv.Bool(b.Value))
				} else {
					c.errorf("mms/type-mismatch", f.Value.Span(), "expected true or false")
				}
			default:
				c.errorf("mms/unknown-field", f.Key.Sp, "a tag entry has only 'id' and 'required'")
			}
		}
		if o.Get("id") == nil {
			c.errorf("mms/missing-field", x.Span(), "a tag entry needs an 'id'")
			return nil
		}
		return o
	}
	c.errorf("mms/type-mismatch", e.Span(), "expected a %s id or tag, got %s", k.TagOf, exprKind(e))
	return nil
}

// Function checks a function body. Until the function backend lands
// (milestone 5), a body may hold only raw commands, which are emitted as
// written.
func (c *Checker) Function(d *syntax.Decl) []string {
	if d.Body == nil {
		c.errorf("mms/missing-value", d.Kind.Sp, "%s needs a body: { ... }", d.Kind.Name)
		return nil
	}
	if d.Params != nil && len(d.Params.Params) > 0 {
		c.errorf("mms/unsupported", d.Params.Sp, "function parameters arrive with the function backend")
	}
	var lines []string
	for _, it := range d.Body.Items {
		cmd, ok := it.(*syntax.CommandStmt)
		if !ok {
			c.errorf("mms/unsupported", it.Span(), "only raw commands (/...) are supported in functions for now; statements arrive with the function backend")
			continue
		}
		if strings.Contains(cmd.Text, "${") {
			c.errorf("mms/unsupported", cmd.Sp, "${...} interpolation arrives with the function backend")
		}
		lines = append(lines, cmd.Text)
	}
	return lines
}
