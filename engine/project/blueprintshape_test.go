package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestBlueprintShape checks the mistakes of writing a Blueprint like a
// Template: they are reported where they are written, with the way to
// write it.
func TestBlueprintShape(t *testing.T) {
	for _, c := range []struct{ src, want string }{
		{"Blueprint T(label: string, rate: int = 20) {\n  emit {\n  }\n}\n", "its fields go in its body, one per line, as in label: string"},
		{"Blueprint T: Function {\n  emit {\n  }\n}\n", "primary Function ${self.path} { ... }"},
		{"Blueprint T {\n  as Function = bool\n  emit {\n  }\n}\n", "'bool' is a type, not a value"},
		{"Blueprint T {\n  let title = \"x\"\n  emit {\n    Objective ${self.title}: dummy\n  }\n}\nT a { }\n",
			"self has no field 'title'; it has id, namespace, path. A let or field of the Blueprint is read by its name: title"},
	} {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("a.mms", c.src)
		r, _ := p.Analyze(context.Background())
		found := false
		for _, d := range r.Diagnostics("a.mms") {
			found = found || strings.Contains(d.Message, c.want)
		}
		if !found {
			t.Errorf("%q: no %q in %+v", c.src, c.want, r.Diagnostics("a.mms"))
		}
	}
}

// TestBlueprintParamsFix checks that the fix for a Blueprint's
// parameters moves them into its body as fields, which then work.
func TestBlueprintParamsFix(t *testing.T) {
	src := "Blueprint Wait(label: string, rate: int = 20): Function {\n  emit {\n    Objective ${label}_cd: dummy\n  }\n}\nWait w { label = \"x\" }\n"
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	// The two fixes are applied in two rounds: the fields go in after
	// the result type, which the first fix's edits surround.
	out := src
	for range 2 {
		out, _ = ApplyFixes([]rune(out), r.Diagnostics("a.mms"))
		p.SetFile("a.mms", out)
		r, _ = p.Analyze(context.Background())
	}
	want := "Blueprint Wait {\n  label: string\n  rate: int = 20\n  emit {"
	if !strings.HasPrefix(out, want) {
		t.Fatalf("fixed = %q, want it to start %q", out, want)
	}
	if ds := r.Diagnostics("a.mms"); len(ds) > 0 {
		t.Errorf("after the fix: %+v", ds)
	}
}
