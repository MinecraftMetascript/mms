package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// recurse is a Blueprint that emits itself one level down, as
// `Recurse <inner>`, until layer reaches 0.
func recurse(inner string, layer int) string {
	return `Blueprint Recurse {
  layer: int
  let next = layer - 1
  emit {
    if (layer > 0) {
      Recurse ` + inner + ` { layer = next }
      primary DensityFunction ${self.path} = Add(` + inner + `, layer)
    } else {
      primary DensityFunction ${self.path} = 1
    }
  }
}
Recurse x { layer = ` + string(rune('0'+layer)) + ` }
DimensionNoise n = Noise(minecraft:jagged, 1, 1)
NoiseSettings s { final_density = x }
`
}

func analyzeOne(t *testing.T, src string) *Result {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return r
}

// TestEmitClash checks that a Blueprint emitting itself under a fixed
// name is told the levels clash, at the line that names it, instead of
// writing a density function that refers to itself (which ran the
// preview out of stack).
func TestEmitClash(t *testing.T) {
	// One level down is fine: only one instance is called inner.
	r := analyzeOne(t, recurse("inner", 1))
	if ds := codes(r, "a.mms"); hasCode(ds, "mms/emit-clash") || hasCode(ds, "mms/duplicate") {
		t.Errorf("layer 1: %v", ds)
	}

	r = analyzeOne(t, recurse("inner", 2))
	ds := codes(r, "a.mms")
	if !hasCode(ds, "mms/emit-clash") {
		t.Fatalf("layer 2: no clash in %v", ds)
	}
	for _, code := range []string{"mms/duplicate", "mms/density-cycle", "mms/unknown-id"} {
		if hasCode(ds, code) {
			t.Errorf("layer 2: the clash should be the only problem, got %v", ds)
		}
	}
	var echoed bool
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code != "mms/emit-clash" {
			continue
		}
		if !strings.Contains(d.Message, "`Recurse ${self.path}_inner`") {
			t.Errorf("message = %q", d.Message)
		}
		if d.Echo {
			echoed = true
			if got := strings.Split(recurse("inner", 2)[d.Sp.Start:], "\n")[0]; got != "inner { layer = next }" {
				t.Errorf("echo at %q, want the nested instance's name", got)
			}
		}
	}
	if !echoed {
		t.Error("the clash is not shown at the line in the Blueprint")
	}
	// The first instance called inner keeps its files whole: no density
	// function refers to itself.
	for _, f := range r.Emit() {
		if strings.HasSuffix(f.Path, "density_function/inner.json") && strings.Contains(string(f.Data), "demo:inner") {
			t.Errorf("inner refers to itself:\n%s", f.Data)
		}
	}

	// Named from the instance, every level has its own name.
	r = analyzeOne(t, recurse("${self.path}_inner", 3))
	if ds := codes(r, "a.mms"); hasCode(ds, "mms/emit-clash") || hasCode(ds, "mms/duplicate") || hasCode(ds, "mms/density-cycle") {
		t.Errorf("named from self: %v", ds)
	}
}

// TestEmitSameInstance checks that a Blueprint instance emitted twice
// with the same output is no clash: its files are written once.
func TestEmitSameInstance(t *testing.T) {
	r := analyzeOne(t, `Blueprint Shared {
  emit {
    Function ${self.path}/hi {
      /say hi
    }
  }
}
Blueprint Item {
  emit {
    Shared shared { }
    Function ${self.path} {
      /say item
    }
  }
}
Item a { }
Item b { }
`)
	if ds := codes(r, "a.mms"); hasCode(ds, "mms/emit-clash") || hasCode(ds, "mms/duplicate") {
		t.Errorf("%v", ds)
	}
}

// TestDensityCycle checks that density functions that refer back to
// themselves are an error at each of them.
func TestDensityCycle(t *testing.T) {
	r := analyzeOne(t, "DensityFunction g = Add(h, 0)\nDensityFunction h = Add(g, 1)\nDensityFunction self_ref = Add(self_ref, 1)\nDensityFunction fine = Add(g, 2)\n")
	var got []string
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code == "mms/density-cycle" {
			got = append(got, d.Message)
		}
	}
	want := []string{"density function g refers to itself through h (g → h → g)", "density function h refers to itself through g (h → g → h)", "density function self_ref refers to itself (self_ref → self_ref)"}
	if len(got) != len(want) {
		t.Fatalf("got %q", got)
	}
	for i := range want {
		if !strings.HasPrefix(got[i], want[i]) {
			t.Errorf("got %q, want %q...", got[i], want[i])
		}
	}
}
