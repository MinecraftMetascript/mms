package ide

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

const splineSrc = `const low = -0.2

DensityFunction offset = Spline({
  coordinate = overworld/continents
  points = [
    SplinePoint(-1.1, 0, 0.044),
    SplinePoint(location = 0.3, derivative = 0, value = {
      coordinate = overworld/erosion
      points = [
        { location = -0.9, derivative = 0, value = 0.2 },
        { location = 0.5, derivative = 0.1, value = low },
      ]
    }),
  ]
})
`

func TestSplineAt(t *testing.T) {
	at := strings.Index(splineSrc, "0.2 }")
	s, _ := session(t, splineSrc[:at]+"|"+splineSrc[at:])
	v, ok := s.SplineAt("main.mms", at)
	if !ok {
		t.Fatal("no spline")
	}
	if v.Root.Coordinate != "overworld/continents" || len(v.Root.Points) != 2 || len(v.Focus) != 1 || v.Focus[0] != 1 {
		t.Fatalf("root %q, %d points, focus %v", v.Root.Coordinate, len(v.Root.Points), v.Focus)
	}
	src := []rune(splineSrc)
	text := func(sp syntax.Span) string { return string(src[sp.Start:sp.End]) }
	p0 := v.Root.Points[0]
	if text(p0.Location.Sp) != "-1.1" || p0.Value.Value != 0.044 || text(p0.Derivative.Sp) != "0" {
		t.Errorf("point 0: %+v", p0)
	}
	in := v.Root.Points[1].Spline
	if in == nil || in.Coordinate != "overworld/erosion" || len(in.Points) != 2 {
		t.Fatalf("nested: %+v", in)
	}
	// A const is worked out, but it is not a literal to rewrite.
	if lo := in.Points[1].Value; !lo.Known || lo.Value != -0.2 || lo.Sp.End != lo.Sp.Start {
		t.Errorf("const value: %+v", lo)
	}
	if v.Root.Template != "SplinePoint($location, $derivative, $value)" {
		t.Errorf("template: %q", v.Root.Template)
	}
	if in.Template != "{ location = $location, derivative = $derivative, value = $value }" {
		t.Errorf("nested template: %q", in.Template)
	}
	// Adding after the last point and removing the first keep the list tidy.
	last := in.Inserts[len(in.Inserts)-1]
	added := string(src[:last.At]) + last.Before + "X" + last.After + string(src[last.At:])
	if !strings.Contains(added, "value = low },\n        X,\n      ]") {
		t.Errorf("insert:\n%s", added)
	}
	rm := v.Root.Points[0].Remove
	if got := string(src[:rm.Start]) + string(src[rm.End:]); !strings.Contains(got, "points = [\n    SplinePoint(location = 0.3") {
		t.Errorf("remove:\n%s", got)
	}
	if v, ok := s.SplineAt("main.mms", strings.Index(splineSrc, "Spline(")+2); !ok || len(v.Focus) != 0 || v.Root.Coordinate != "overworld/continents" {
		t.Errorf("at the call: %v %+v", ok, v)
	}
	if _, ok := s.SplineAt("main.mms", strings.Index(splineSrc, "low = ")); ok {
		t.Error("a spline outside any spline")
	}
}

// An empty Spline() is a spline to start: the first point writes it.
func TestSplineAtEmpty(t *testing.T) {
	src := "DensityFunction d = Spline()\n"
	at := strings.Index(src, "()") + 1
	s, _ := session(t, src[:at]+"|"+src[at:])
	v, ok := s.SplineAt("main.mms", at)
	if !ok || !v.Root.Editable || len(v.Root.Points) != 0 || len(v.Root.Inserts) != 1 {
		t.Fatalf("empty spline: %v %+v", ok, v.Root)
	}
	in := v.Root.Inserts[0]
	point := strings.NewReplacer("$location", "0.5", "$derivative", "0", "$value", "1").Replace(v.Root.Template)
	got := src[:in.At] + in.Before + point + in.After + src[in.At:]
	if want := "DensityFunction d = Spline({ coordinate = 0, points = [SplinePoint(0.5, 0, 1)] })\n"; got != want {
		t.Fatalf("first point:\n%s", got)
	}
	// What it writes is a spline the editor reads.
	s2, _ := session(t, got[:at+3]+"|"+got[at+3:])
	if v, ok := s2.SplineAt("main.mms", at+3); !ok || len(v.Root.Points) != 1 || v.Root.Coordinate != "0" {
		t.Fatalf("after the first point: %v %+v", ok, v.Root)
	}
	for _, d := range s2.R.Diagnostics("main.mms") {
		if d.Code != "lint/unused" {
			t.Errorf("diagnostic: %+v", d)
		}
	}
}
