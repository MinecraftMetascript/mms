package ide

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spline"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// SplineView is the spline at a position, for the spline editor: the
// outermost spline of the nesting the position is in, and the point
// indices from it to the innermost spline there.
type SplineView struct {
	Root  *SplineInfo
	Focus []int
}

// SplineInfo is one spline as the editor shows it.
type SplineInfo struct {
	Sp         syntax.Span
	Coordinate string
	// Editable is set when points are written as a list in place, so the
	// editor can add and remove them.
	Editable bool
	// Template is the source of a new point, with $location,
	// $derivative and $value in place of its numbers.
	Template string
	// Inserts are where a new point goes, one per gap: Inserts[i] puts it
	// before point i, and the last one after the last point.
	Inserts []SplineInsert
	Points  []*SplinePointInfo
}

// SplineInsert adds a point: Before, the point and After go at At.
type SplineInsert struct {
	At            int
	Before, After string
}

// SplinePointInfo is one point; Value is unset when Spline is.
type SplinePointInfo struct {
	Sp                          syntax.Span
	Location, Derivative, Value SplineNum
	Spline                      *SplineInfo
	// Remove is the span that deleting the point removes.
	Remove syntax.Span
}

// SplineNum is a number of a point. Sp is set (non-empty) only for a
// literal the editor can rewrite.
type SplineNum struct {
	Value float64
	Known bool
	Sp    syntax.Span
}

// SplineAt returns the spline at off in a file's source, as written.
func (s *Session) SplineAt(path string, off int) (SplineView, bool) {
	f := s.R.Source[path]
	if f == nil {
		return SplineView{}, false
	}
	root, focus := spline.At(syntax.PathAt(f, off), off, f.Src, s.foldNumber)
	if root == nil {
		return SplineView{}, false
	}
	return SplineView{Root: splineInfo(root, f.Src), Focus: focus}, true
}

func splineInfo(sp *spline.Spline, src []rune) *SplineInfo {
	out := &SplineInfo{Sp: sp.X.Sp, Coordinate: sp.Coordinate, Editable: sp.List != nil}
	for i, p := range sp.Points {
		pi := &SplinePointInfo{Sp: p.X.Span(), Location: splineNum(p.Location), Derivative: splineNum(p.Derivative)}
		if p.Spline != nil {
			pi.Spline = splineInfo(p.Spline, src)
		} else {
			pi.Value = splineNum(p.Value)
		}
		if sp.List != nil {
			pi.Remove = removeSpan(sp.List, i, src)
		}
		out.Points = append(out.Points, pi)
	}
	if sp.List != nil {
		out.Template = pointTemplate(sp, src)
		out.Inserts = inserts(sp.List, src)
	}
	return out
}

func splineNum(n spline.Num) SplineNum {
	out := SplineNum{Value: n.Value, Known: n.Known}
	if n.Literal {
		out.Sp = n.Sp
	}
	return out
}

// pointTemplate copies the first point written with literal numbers, so
// a new point looks like its neighbours; without one it is a
// SplinePoint call.
func pointTemplate(sp *spline.Spline, src []rune) string {
	for _, p := range sp.Points {
		if p.Spline != nil || !p.Location.Literal || !p.Derivative.Literal || !p.Value.Literal {
			continue
		}
		whole := p.X.Span()
		nums := []struct {
			sp   syntax.Span
			name string
		}{{p.Location.Sp, "$location"}, {p.Derivative.Sp, "$derivative"}, {p.Value.Sp, "$value"}}
		// Replace from the end so the earlier spans stay put.
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[j].sp.Start > nums[i].sp.Start {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}
		}
		t := []rune(string(src[whole.Start:whole.End]))
		for _, n := range nums {
			a, b := n.sp.Start-whole.Start, n.sp.End-whole.Start
			t = append(append(append([]rune{}, t[:a]...), []rune(n.name)...), t[b:]...)
		}
		return string(t)
	}
	return spline.SplinePoint + "($location, $derivative, $value)"
}

// separator is what goes between two points: as written between the
// first two, or for a single point ", " (on one line) or a new line at
// its indent.
func separator(l *syntax.List, src []rune) string {
	if len(l.Elems) >= 2 {
		sep := string(src[l.Elems[0].Span().End:l.Elems[1].Span().Start])
		if !strings.Contains(sep, "//") && !strings.Contains(sep, "/*") {
			return sep
		}
	}
	if len(l.Elems) >= 1 {
		before := string(src[l.Sp.Start+1 : l.Elems[0].Span().Start])
		if i := strings.LastIndexByte(before, '\n'); i >= 0 {
			return "," + before[i:]
		}
	}
	return ", "
}

func inserts(l *syntax.List, src []rune) []SplineInsert {
	n := len(l.Elems)
	if n == 0 {
		return []SplineInsert{{At: l.Sp.Start + 1}}
	}
	sep := separator(l, src)
	var out []SplineInsert
	for _, e := range l.Elems {
		out = append(out, SplineInsert{At: e.Span().Start, After: sep})
	}
	return append(out, SplineInsert{At: l.Elems[n-1].Span().End, Before: sep})
}

// removeSpan is what deleting point i removes: the point and the
// separator after it, or before it for the last point.
func removeSpan(l *syntax.List, i int, src []rune) syntax.Span {
	es := l.Elems
	switch {
	case i < len(es)-1:
		return syntax.Span{Start: es[i].Span().Start, End: es[i+1].Span().Start}
	case i > 0:
		return syntax.Span{Start: es[i-1].Span().End, End: es[i].Span().End}
	}
	sp := es[i].Span()
	// A lone point's trailing comma goes too.
	for j := sp.End; j < l.Sp.End && j < len(src); j++ {
		if src[j] == ',' {
			sp.End = j + 1
			break
		}
		if src[j] != ' ' && src[j] != '\t' && src[j] != '\n' && src[j] != '\r' {
			break
		}
	}
	return sp
}

// foldNumber works out a number expression of literals, arithmetic and
// consts.
func (s *Session) foldNumber(e syntax.Expr) (float64, bool) {
	return s.fold(e, 0)
}

func (s *Session) fold(e syntax.Expr, depth int) (float64, bool) {
	if depth > 16 {
		return 0, false
	}
	switch x := e.(type) {
	case *syntax.Number:
		v, err := strconv.ParseFloat(strings.ReplaceAll(x.Text, "_", ""), 64)
		return v, err == nil
	case *syntax.Paren:
		return s.fold(x.X, depth)
	case *syntax.Unary:
		v, ok := s.fold(x.X, depth)
		switch x.Op {
		case "-":
			return -v, ok
		case "+":
			return v, ok
		}
	case *syntax.Binary:
		a, ok1 := s.fold(x.L, depth)
		b, ok2 := s.fold(x.R, depth)
		if !ok1 || !ok2 {
			return 0, false
		}
		switch x.Op {
		case "+":
			return a + b, true
		case "-":
			return a - b, true
		case "*":
			return a * b, true
		case "/":
			return a / b, b != 0
		}
	case *syntax.Ident:
		if s.R.Eval == nil {
			return 0, false
		}
		if c := s.R.Eval.Consts[x.Name]; c != nil {
			if v := c.Lowered(); v != nil {
				return s.fold(v, depth+1)
			}
		}
	}
	return 0, false
}
