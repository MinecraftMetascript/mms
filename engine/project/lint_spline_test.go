package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

func TestLintSplineOrder(t *testing.T) {
	src := `DensityFunction d = Spline({
  coordinate = overworld/continents
  points = [
    SplinePoint(-1, 0, 0),
    { location = 0.5, derivative = 0, value = { coordinate = overworld/erosion, points = [SplinePoint(1, 0, 0), SplinePoint(0, 0, 1)] } },
    SplinePoint(0.2, 0, 1),
  ]
})
`
	var codes []string
	for _, c := range lintCodes(t, Config{Namespace: "demo"}, src) {
		if strings.HasPrefix(c, "lint/spline-order") {
			codes = append(codes, c)
		}
	}
	if got := strings.Join(codes, " "); got != "lint/spline-order@5 lint/spline-order@6" {
		t.Fatalf("lint = %s", got)
	}
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var fix *syntax.Fix
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code == "lint/spline-order" && len(d.Fixes) > 0 && strings.HasPrefix(d.Fixes[0].Title, "Sort") {
			line, _ := r.Unit("a.mms").File.LineCol(d.Sp.Start)
			if line == 6 {
				fix = &d.Fixes[0]
			}
		}
	}
	if fix == nil {
		t.Fatal("no sort fix on the outer spline")
	}
	got := applyEdits(src, fix.Edits)
	want := "    SplinePoint(-1, 0, 0),\n    SplinePoint(0.2, 0, 1),\n    { location = 0.5"
	if !strings.Contains(got, want) {
		t.Errorf("sorted:\n%s", got)
	}
}
