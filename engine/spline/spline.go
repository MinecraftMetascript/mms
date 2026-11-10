// Package spline reads the cubic splines of density functions from
// source: a `{ coordinate = ..., points = [...] }` value, as written in
// `Spline(...)` or a const, whose points are `SplinePoint(...)` calls or
// `{ location, derivative, value }` objects, and whose point values may
// be splines again. The spline editor in the preview and the
// spline-order lint read them.
package spline

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Num is one number of a point.
type Num struct {
	Value float64
	// Known is set when the value could be worked out.
	Known bool
	// Literal is set when the number is written as a literal, such as
	// `-0.5`, which the editor can rewrite.
	Literal bool
	// Sp is the expression's span; it is empty when the field is left out.
	Sp syntax.Span
	X  syntax.Expr
}

// Point is one point of a spline.
type Point struct {
	X                    syntax.Expr
	Location, Derivative Num
	// Value is the point's value when it is a number.
	Value Num
	// Spline is the point's value when it is a spline.
	Spline *Spline
}

// Spline is a multipoint cubic spline.
type Spline struct {
	X *syntax.Object
	// Coordinate is the coordinate's source text, such as `continents`.
	Coordinate string
	CoordSp    syntax.Span
	// List is the points list, or nil when points is not written as a
	// list in place.
	List   *syntax.List
	Points []*Point
}

// Fold works out a number expression; Read uses it for numbers that are
// not literals.
type Fold func(syntax.Expr) (float64, bool)

// SplinePoint is the helper that builds a point; its arguments are, in
// order, location, derivative and value.
const SplinePoint = "SplinePoint"

var pointArgs = []string{"location", "derivative", "value"}

// Read reads e as a spline, or returns nil when it is not one.
func Read(e syntax.Expr, src []rune, fold Fold) *Spline {
	o, ok := unparen(e).(*syntax.Object)
	if !ok {
		return nil
	}
	fields := objectFields(o)
	pts, ok := fields["points"]
	if !ok {
		return nil
	}
	if _, ok := fields["coordinate"]; !ok {
		return nil
	}
	s := &Spline{X: o}
	if c := fields["coordinate"]; c != nil {
		s.CoordSp = c.Span()
		s.Coordinate = strings.Join(strings.Fields(text(src, c.Span())), " ")
	}
	l, ok := unparen(pts).(*syntax.List)
	if !ok {
		return s
	}
	s.List = l
	for _, el := range l.Elems {
		s.Points = append(s.Points, readPoint(el, src, fold))
	}
	return s
}

// Of returns the spline of a `Spline(...)` call, or nil.
func Of(c *syntax.Call, src []rune, fold Fold) *Spline {
	if id, ok := c.Fn.(*syntax.Ident); !ok || id.Name != "Spline" {
		return nil
	}
	for i, a := range c.Args {
		if a.Name == nil && i == 0 || a.Name != nil && a.Name.Text == "spline" {
			return Read(a.Value, src, fold)
		}
	}
	return nil
}

func readPoint(e syntax.Expr, src []rune, fold Fold) *Point {
	p := &Point{X: e}
	args := map[string]syntax.Expr{}
	switch x := unparen(e).(type) {
	case *syntax.Call:
		if id, ok := x.Fn.(*syntax.Ident); ok && id.Name == SplinePoint {
			i := 0
			for _, a := range x.Args {
				if a.Name != nil {
					args[a.Name.Text] = a.Value
				} else if i < len(pointArgs) {
					args[pointArgs[i]] = a.Value
					i++
				}
			}
		}
	case *syntax.Object:
		args = objectFields(x)
	}
	p.Location = readNum(args["location"], fold)
	p.Derivative = readNum(args["derivative"], fold)
	if v := args["value"]; v != nil {
		if s := Read(v, src, fold); s != nil {
			p.Spline = s
		} else {
			p.Value = readNum(v, fold)
		}
	}
	return p
}

func readNum(e syntax.Expr, fold Fold) Num {
	if e == nil {
		return Num{}
	}
	n := Num{X: e, Sp: e.Span()}
	if t, ok := literal(e); ok {
		if v, err := strconv.ParseFloat(t, 64); err == nil {
			n.Value, n.Known, n.Literal = v, true, true
			return n
		}
	}
	if fold != nil {
		n.Value, n.Known = fold(e)
	}
	return n
}

// literal is the text of a number literal, `-` included.
func literal(e syntax.Expr) (string, bool) {
	switch x := e.(type) {
	case *syntax.Number:
		return strings.ReplaceAll(x.Text, "_", ""), true
	case *syntax.Unary:
		if n, ok := x.X.(*syntax.Number); ok && x.Op == "-" {
			return "-" + strings.ReplaceAll(n.Text, "_", ""), true
		}
	}
	return "", false
}

func objectFields(o *syntax.Object) map[string]syntax.Expr {
	out := map[string]syntax.Expr{}
	if o.Body == nil {
		return out
	}
	for _, it := range o.Body.Items {
		if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil {
			out[f.Key.Text] = f.Value
		}
	}
	return out
}

func unparen(e syntax.Expr) syntax.Expr {
	for {
		p, ok := e.(*syntax.Paren)
		if !ok {
			return e
		}
		e = p.X
	}
}

func text(src []rune, sp syntax.Span) string {
	if sp.Start < 0 || sp.End > len(src) || sp.Start > sp.End {
		return ""
	}
	return string(src[sp.Start:sp.End])
}

// All returns every spline in n, nested ones included, in source order.
func All(n syntax.Node, src []rune, fold Fold) []*Spline {
	var out []*Spline
	syntax.Walk(n, func(n syntax.Node) bool {
		if e, ok := n.(syntax.Expr); ok {
			if s := Read(e, src, fold); s != nil {
				out = append(out, s)
			}
		}
		return true
	})
	return out
}

// At finds the spline at off in path (as syntax.PathAt gives it): the
// outermost spline of the nesting the cursor is in, and the point
// indices from it to the innermost spline at the cursor.
func At(path []syntax.Node, off int, src []rune, fold Fold) (*Spline, []int) {
	// The innermost spline at the cursor.
	var inner *Spline
	at := -1
	for i := len(path) - 1; i >= 0; i-- {
		if e, ok := path[i].(syntax.Expr); ok {
			if s := Read(e, src, fold); s != nil {
				inner, at = s, i
				break
			}
		}
		// The cursor on a `Spline(...)` call's name.
		if c, ok := path[i].(*syntax.Call); ok {
			if s := Of(c, src, fold); s != nil {
				inner, at = s, i
				break
			}
		}
	}
	if inner == nil {
		return nil, nil
	}
	// Climb while the spline is a point's value of the spline above.
	root, focus := inner, []int(nil)
	for i := at - 1; i >= 0; i-- {
		e, ok := path[i].(syntax.Expr)
		if !ok {
			continue
		}
		up := Read(e, src, fold)
		if up == nil {
			continue
		}
		k := -1
		for j, p := range up.Points {
			if p.Spline != nil && p.Spline.X == root.X {
				k = j
			}
		}
		if k < 0 {
			break
		}
		root, focus = up, append([]int{k}, focus...)
	}
	return root, focus
}
