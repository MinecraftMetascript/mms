package sema

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Tag checks a tag declaration's entries: a list of ids and tags of the
// tag's registry. Declarations of the same tag are merged by the emitter.
func (c *Checker) Tag(d *syntax.Decl, k Kind) jsonv.Array {
	var elems []syntax.Expr
	switch x := d.Value.(type) {
	case *syntax.List:
		elems = x.Elems
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
	out := jsonv.Array{}
	for _, e := range elems {
		c.info(e).Type = tview{}
		switch x := e.(type) {
		case *syntax.Ident:
			out = append(out, jsonv.String(c.resolveID(x.Name, id, e)))
		case *syntax.ResLoc:
			out = append(out, jsonv.String(c.resolveID(x.Text, id, e)))
		case *syntax.Tag:
			out = append(out, jsonv.String(c.resolveTag(x.Name, id, e)))
		default:
			c.errorf("mms/type-mismatch", e.Span(), "expected a %s id or tag, got %s", k.TagOf, exprKind(e))
		}
	}
	return out
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
