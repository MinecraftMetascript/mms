package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TestWrongKindDeps checks that a reference naming a declaration of the
// wrong kind is rechecked when that declaration changes in another file.
func TestWrongKindDeps(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", "DensityFunction terrain = Abs(ridges)\n")
	p.SetFile("b.mms", "Noise ridges = { base_octave = -7, octave_count = 1 }\n")
	msg := func() string {
		r, _ := p.Analyze(context.Background())
		var out []string
		for _, d := range r.Diagnostics("a.mms") {
			if d.Severity == syntax.SeverityError {
				out = append(out, d.Code+" "+d.Message)
			}
		}
		return strings.Join(out, "\n")
	}
	if got := msg(); !strings.Contains(got, "'ridges' is a Noise, not a DensityFunction") {
		t.Fatalf("want a wrong-kind error, got %q", got)
	}
	p.SetFile("b.mms", "DensityFunction ridges = 0.5\n")
	if got := msg(); got != "" {
		t.Fatalf("after declaring the density function, got %q", got)
	}
	p.SetFile("b.mms", "// nothing\n")
	if got := msg(); !strings.Contains(got, "unknown worldgen/density_function 'minecraft:ridges'") {
		t.Fatalf("after removing it, got %q", got)
	}
}
