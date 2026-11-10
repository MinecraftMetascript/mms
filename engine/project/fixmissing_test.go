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
		{"Noise n {}\n", "Noise n { base_octave = 0 }\n"},
		{"Noise n { amplitudes = [1] }\n", "Noise n { amplitudes = [1], base_octave = 0 }\n"},
		{"Noise n {\n  amplitudes = [1]\n}\n", "Noise n {\n  amplitudes = [1]\n  base_octave = 0\n}\n"},
		{"Noise n {\n}\n", "Noise n {\n  base_octave = 0\n}\n"},
		{"Noise n {\n    amplitudes = [1] }\n", "Noise n {\n    amplitudes = [1]\n    base_octave = 0\n}\n"},
		{"DensityFunction d = Clamp(1)\n", "DensityFunction d = Clamp(1, min = 0, max = 0)\n"},
		{"DensityFunction d = Noise()\n", `DensityFunction d = Noise(noise = "", xz_scale = 0, y_scale = 0)` + "\n"},
		// Unions start from their simplest member: a number provider is a number.
		{"LootTable t { pools = [{ entries = [] }] }\n", "LootTable t { pools = [{ entries = [], rolls = 0 }] }\n"},
		// Colors start as opaque white, in the "#hex" form.
		{"Biome b { temperature = 0, downfall = 0, has_precipitation = false, carvers = [], features = [], effects = {} }\n", "Biome b { temperature = 0, downfall = 0, has_precipitation = false, carvers = [], features = [], effects = { water_color = \"#FFFFFF\" } }\n"},
		{"Biome b {\n  temperature = 1\n}\n", "Biome b {\n  temperature = 1\n  downfall = 0\n  has_precipitation = false\n  effects = {}\n  carvers = []\n  features = []\n}\n"},
		// Typed consts, Template arguments and Blueprint fields get the
		// fix too, where their values are written.
		{"const src: BiomeSource = { type = checkerboard }\n", "const src: BiomeSource = { type = checkerboard, biomes = [] }\n"},
		{"const src: BiomeSource = Checkerboard()\n", "const src: BiomeSource = Checkerboard(biomes = [])\n"},
		{"const n: Noise = {\n}\n", "const n: Noise = {\n  base_octave = 0\n}\n"},
		{"Type P {\n  x: int\n  y: string\n}\nconst p: P = {\n  x = 1\n}\n", "Type P {\n  x: int\n  y: string\n}\nconst p: P = {\n  x = 1\n  y = \"\"\n}\n"},
		{"Type P {\n  x: int\n  q: { a: int, b: int }\n}\nconst p: P = { x = 1, q = { a = 1 } }\n", "Type P {\n  x: int\n  q: { a: int, b: int }\n}\nconst p: P = { x = 1, q = { a = 1, b = 0 } }\n"},
		{"Type P {\n  x: int\n  y: string\n}\nconst ps: [P] = [{ x = 1 }]\n", "Type P {\n  x: int\n  y: string\n}\nconst ps: [P] = [{ x = 1, y = \"\" }]\n"},
		{"Type P {\n  x: int\n  y: string\n}\nTemplate f(p: P): int = 1\nconst z = f({ x = 1 })\n", "Type P {\n  x: int\n  y: string\n}\nTemplate f(p: P): int = 1\nconst z = f({ x = 1, y = \"\" })\n"},
		{"Type P {\n  x: int\n  y: string\n}\nBlueprint B {\n  p: P\n}\nB b {\n  p = { x = 1 }\n}\n", "Type P {\n  x: int\n  y: string\n}\nBlueprint B {\n  p: P\n}\nB b {\n  p = { x = 1, y = \"\" }\n}\n"},
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
	got, f := missingFix(t, "DensityFunction d = Clamp(1)\nNoise n {}\n")
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
	p.SetFile("a.mms", "Noise n {}\n")
	files, _, err := p.FixAll(context.Background(), "a.mms")
	if err != nil || strings.Contains(files["a.mms"], "base_octave") {
		t.Errorf("fix all = %q, %v", files["a.mms"], err)
	}
}

// A value evaluated from elsewhere keeps the spans of where it is
// written, so a typed const naming another const gets no fix that would
// edit the other one's text.
func TestMissingFieldFixStaysInValue(t *testing.T) {
	for _, src := range []string{
		"Type P {\n  x: int\n  y: string\n}\nconst a = { x = 1 }\nconst p: P = a\n",
		"Type P {\n  x: int\n  y: string\n}\nconst a = { x = 1 }\nconst p: P = { ...a }\n",
	} {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("a.mms", src)
		r, _ := p.Analyze(context.Background())
		found := false
		for _, d := range r.Diagnostics("a.mms") {
			if d.Code != "mms/missing-field" {
				continue
			}
			found = true
			if len(d.Fixes) > 0 {
				t.Errorf("%q: fix %q edits another value", src, d.Fixes[0].Title)
			}
		}
		if !found {
			t.Errorf("%q: no missing-field error", src)
		}
	}
}
