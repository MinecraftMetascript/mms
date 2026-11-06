package project

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// densityLints returns the density-* findings for src as "rule@line".
func densityLints(t *testing.T, src string) ([]string, *Result) {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var out []string
	for _, d := range r.Diagnostics("a.mms") {
		if d.Severity == syntax.SeverityError {
			t.Fatalf("unexpected error: %s %s", d.Code, d.Message)
		}
		if !strings.HasPrefix(d.Code, "lint/density-") {
			continue
		}
		line, _ := r.Unit("a.mms").File.LineCol(d.Sp.Start)
		out = append(out, fmt.Sprintf("%s@%d", strings.TrimPrefix(d.Code, "lint/"), line))
	}
	return out, r
}

// settings is a noise settings declaration with the given final_density.
func settings(final string) string {
	return `NoiseSettings land = {
  sea_level = 0
  default_block = stone
  default_fluid = air
  disable_mob_generation = false
  legacy_random_source = false
  noise = { min_y = 0, height = 256 }
  material_rule = Block(stone)
  spawn_target = []
  noise_router = {
    final_density = ` + final + `
    temperature = 0
    vegetation = 0
    continents = 0
    erosion = 0
    depth = 0
    ridges = 0
    chunk_surface_level = 0
  }
}
`
}

const densityNoise = "Noise n = { base_octave = -6, amplitude_modifiers = [1] }\n"

func TestDensityLints(t *testing.T) {
	cases := []struct {
		name, src, want string
	}{
		{"clean",
			densityNoise + "DensityFunction a = Cache(Noise(n, 1, 1))\nDensityFunction b = a + a\n" +
				settings("Interpolated(b, 4, 8) + Beardifier()"),
			""},
		{"uncached shared",
			densityNoise + "DensityFunction a = Noise(n, 1, 1)\nDensityFunction b = a + a\n" +
				settings("Interpolated(b, 4, 8) + Beardifier()"),
			"density-uncached-shared@2"},
		{"shared through two paths",
			// a is used once by b, but b is used twice: a runs twice.
			densityNoise + "DensityFunction a = Noise(n, 1, 1)\nDensityFunction b = a * 2\n" +
				settings("Interpolated(b - b, 4, 8) + Beardifier()"),
			"density-uncached-shared@2 density-uncached-shared@3"},
		{"cheap shared functions are fine",
			densityNoise + "DensityFunction a = Cache(Noise(n, 1, 1))\nDensityFunction b = a * 2\n" +
				settings("Interpolated(b - b, 4, 8) + Beardifier()"),
			""},
		{"inline cache at the use",
			densityNoise + "DensityFunction a = Noise(n, 1, 1)\n" +
				settings("Interpolated(Cache(a) + Cache(a), 4, 8) + Beardifier()"),
			"density-uncached-shared@2"},
		{"square",
			densityNoise + "DensityFunction a = Cache(Noise(n, 1, 1))\nDensityFunction b = a * a\n" +
				settings("Interpolated(b, 4, 8) + Beardifier()"),
			"density-square@3"},
		{"no interpolation or beardifier",
			densityNoise + "DensityFunction a = Noise(n, 1, 1)\n" + settings("a"),
			"density-no-beardifier@13 density-no-interpolation@13"},
		{"constant terrain needs neither",
			settings("Gradient(axis = y, from_coordinate = 0, to_coordinate = 128, from_value = 1, to_value = -1)"),
			""},
		{"nested interpolation",
			densityNoise + "DensityFunction a = Interpolated(Noise(n, 1, 1), 4, 8)\n" +
				settings("Interpolated(a * 0.5, 4, 8) + Beardifier()"),
			"density-nested-interpolation@2"},
		{"vanilla functions are not known",
			settings("minecraft:overworld/final_density"),
			""},
	}
	for _, c := range cases {
		got, _ := densityLints(t, c.src)
		if g := strings.Join(got, " "); g != c.want {
			t.Errorf("%s: got %q, want %q", c.name, g, c.want)
		}
	}
}

func TestDensityLintFixes(t *testing.T) {
	src := densityNoise + "DensityFunction a = Noise(n, 1, 1)\nDensityFunction b = (a) * (a)\n" +
		settings("Interpolated(b, 4, 8) + Beardifier()")
	_, r := densityLints(t, src)
	var edits []string
	for _, d := range r.Diagnostics("a.mms") {
		if !strings.HasPrefix(d.Code, "lint/density-") {
			continue
		}
		f := d.Fixes[0]
		if f.Safe {
			t.Errorf("%s: fix %q is marked safe, but it changes the pack", d.Code, f.Title)
		}
		for _, e := range f.Edits {
			edits = append(edits, e.Text)
		}
	}
	if got := strings.Join(edits, "|"); got != "Cache(|)|Square(a)" {
		t.Errorf("fix edits = %q", got)
	}
}
