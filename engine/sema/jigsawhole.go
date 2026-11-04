package sema

import "github.com/minecraftmetascript/mms/engine/syntax"

// JigsawHole checks a `JigsawHole name` declaration: a name and nothing
// else, since a hole is only a name that jigsaw blocks share.
func (c *Checker) JigsawHole(d *syntax.Decl) {
	if d.Name == nil {
		c.errorf("mms/syntax", d.Kind.Sp, "JigsawHole needs a name: `JigsawHole road_end`")
		return
	}
	var extra syntax.Node
	switch {
	case d.Type != nil:
		extra = d.Type
	case d.Params != nil:
		extra = d.Params
	case d.Value != nil:
		extra = d.Value
	case d.Body != nil:
		extra = d.Body
	}
	if extra == nil {
		return
	}
	sp := syntax.Span{Start: d.Name.Sp.End, End: d.Sp.End}
	c.Diags = append(c.Diags, syntax.Diagnostic{
		Code: "mms/syntax", Severity: syntax.SeverityError, Sp: extra.Span(),
		Message: "a JigsawHole is only a name; write `JigsawHole " + d.Name.Text + "` with nothing after it",
		Fixes:   []syntax.Fix{{Title: "Remove what follows the name", Safe: true, Edits: []syntax.TextEdit{{Sp: sp}}}},
	})
	c.errs++
}
