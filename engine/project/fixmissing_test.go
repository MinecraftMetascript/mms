package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// missingFix applies the fix of the first missing-field diagnostic.
func missingFix(t *testing.T, src string) (string, syntax.Fix) {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	for _, d := range r.Diagnostics("a.mms") {
		if d.Code != "mms/missing-field" {
			continue
		}
		if len(d.Fixes) == 0 {
			t.Fatalf("%q: no fix for %s", src, d.Message)
		}
		f := d.Fixes[0]
		rs := []rune(src)
		e := f.Edits[0]
		return string(rs[:e.Sp.Start]) + e.Text + string(rs[e.Sp.End:]), f
	}
	t.Fatalf("%q: no missing-field diagnostic", src)
	return "", syntax.Fix{}
}

func TestMissingFieldFix(t *testing.T) {
	cases := []struct{ src, want string }{
		{"Noise n = {}\n", "Noise n = { base_octave = 0 }\n"},
		{"Noise n = { amplitudes = [1] }\n", "Noise n = { amplitudes = [1], base_octave = 0 }\n"},
		{"Noise n = {\n  amplitudes = [1]\n}\n", "Noise n = {\n  amplitudes = [1]\n  base_octave = 0\n}\n"},
		{"Noise n = {\n}\n", "Noise n = {\n  base_octave = 0\n}\n"},
		{"Noise n = {\n    amplitudes = [1] }\n", "Noise n = {\n    amplitudes = [1]\n    base_octave = 0\n}\n"},
		{"DensityFunction d = Clamp(1)\n", "DensityFunction d = Clamp(1, min = 0, max = 0)\n"},
		{"DensityFunction d = Noise()\n", `DensityFunction d = Noise(noise = "", xz_scale = 0, y_scale = 0)` + "\n"},
		// Unions start from their simplest member: a number provider is a number.
		{"LootTable t = { pools = [{ entries = [] }] }\n", "LootTable t = { pools = [{ entries = [], rolls = 0 }] }\n"},
		{"Biome b {\n  temperature = 1\n}\n", "Biome b {\n  temperature = 1\n  downfall = 0\n  has_precipitation = false\n  effects = {}\n  carvers = []\n  features = []\n}\n"},
	}
	for _, c := range cases {
		got, f := missingFix(t, c.src)
		if got != c.want {
			t.Errorf("fix of %q:\n got %q\nwant %q", c.src, got, c.want)
		}
		if f.Safe {
			t.Errorf("%q: placeholder fix marked safe", c.src)
		}
	}

	// The placeholders type-check, so only nested fields remain.
	got, f := missingFix(t, "DensityFunction d = Clamp(1)\nNoise n = {}\n")
	if f.Title != "Add missing fields 'min', 'max'" {
		t.Errorf("title = %q", f.Title)
	}
	got, _ = missingFix(t, got)
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", got)
	r, _ := p.Analyze(context.Background())
	for _, d := range r.Diagnostics("a.mms") {
		if d.Severity == syntax.SeverityError {
			t.Errorf("after fixes %q: %s", got, d.Message)
		}
	}
	if !strings.Contains(got, "base_octave = 0") {
		t.Errorf("second fix not applied: %q", got)
	}

	// Fix all applies only safe fixes, so placeholders are never guessed.
	p.SetFile("a.mms", "Noise n = {}\n")
	files, _, err := p.FixAll(context.Background(), "a.mms")
	if err != nil || strings.Contains(files["a.mms"], "base_octave") {
		t.Errorf("fix all = %q, %v", files["a.mms"], err)
	}
}
