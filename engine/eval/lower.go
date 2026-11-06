package eval

import (
	"regexp"
	"strconv"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

var identRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// keyRe matches a field key that needs no quotes.
var keyRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

// lower writes a value back as an expression, with fresh nodes at the
// value's span. None lowers to nil; callers drop it.
func (c *ctx) lower(v Value) syntax.Expr {
	if v == nil {
		return nil
	}
	p := v.at()
	var out syntax.Expr
	switch v := v.(type) {
	case *None:
		if v.lit {
			out = &syntax.Ident{Name: "none", Sp: p.Sp}
			break
		}
		return nil
	case *Bool:
		out = &syntax.Bool{Value: v.V, Sp: p.Sp}
	case *Num:
		out = &syntax.Number{Text: numText(v), Sp: p.Sp}
	case *Str:
		out = &syntax.String{Value: v.S, Raw: strconv.Quote(v.S), Sp: p.Sp}
	case *Word:
		if identRe.MatchString(v.Text) {
			out = &syntax.Ident{Name: v.Text, Sp: p.Sp}
		} else {
			out = &syntax.ResLoc{Text: v.Text, Sp: p.Sp}
		}
	case *List:
		l := &syntax.List{Elems: []syntax.Expr{}, Sp: p.Sp}
		for _, el := range v.Elems {
			if x := c.lower(el); x != nil {
				l.Elems = append(l.Elems, x)
			}
		}
		out = l
	case *Record:
		if v.id != nil {
			return c.lower(&Word{*p, v.id.Text})
		}
		out = &syntax.Object{Body: c.recordBlock(v), Sp: p.Sp}
	case *RangeV:
		out = &syntax.Range{Lo: c.lower(v.Lo), Hi: c.lower(v.Hi), Sp: p.Sp}
	case *Code:
		out = syntax.CloneExpr(v.X)
	case *Stmts:
		b := &syntax.Block{Closed: true, Sp: p.Sp}
		for _, st := range v.Items {
			b.Items = append(b.Items, syntax.CloneStmt(st))
		}
		out = &syntax.RunExpr{Body: b, Sp: p.Sp}
	}
	if p.def != nil {
		c.e.r.Nodes[out] = Prov{At: Origin{c.path, p.Sp}, Def: p.def, Frames: p.fr}
	}
	return out
}

// recordBlock writes a record's fields as a block, leaving out none.
func (c *ctx) recordBlock(r *Record) *syntax.Block {
	b := &syntax.Block{Closed: true, Sp: r.Sp}
	for i, k := range r.Keys {
		if dropped(r.Vals[i]) {
			continue
		}
		x := c.lower(r.Vals[i])
		b.Items = append(b.Items, &syntax.FieldStmt{Key: &syntax.Name{Text: k, Sp: r.Sp}, KeyQuoted: !keyRe.MatchString(k),
			Value: x, Sp: r.Sp})
	}
	return b
}

// dropped reports whether a field with value v is left out: v is none
// that evaluation produced, not the word `none` as written.
func dropped(v Value) bool {
	n, ok := v.(*None)
	return ok && !n.lit
}
