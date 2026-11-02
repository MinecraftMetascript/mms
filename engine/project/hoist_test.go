package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// emitText builds src as demo's only file and returns the pack by path,
// failing on any error.
func emitText(t *testing.T, src string) map[string]string {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range r.Diagnostics("a.mms") {
		if d.Severity == syntax.SeverityError {
			t.Fatalf("unexpected error: %s %s", d.Code, d.Message)
		}
	}
	out := map[string]string{}
	for _, f := range r.Emit() {
		out[f.Path] = string(f.Data)
	}
	return out
}

func compact(s string) string { return strings.Join(strings.Fields(s), "") }

// TestHoistInlineNoise checks that an inline noise in a density function,
// which the game cannot seed without an id, is written as its own noise
// and referenced by id. Hoisted values are numbered in the order the
// schema lists their fields.
func TestHoistInlineNoise(t *testing.T) {
	out := emitText(t, `
Noise simple = { base_octave = -4, amplitude_modifiers = [1] }

NoiseSettings custom {
  default_block = stone
  default_fluid = water
  disable_mob_generation = false
  material_rule = [
    If(NoiseThreshold(simple, 0.5, 1)) { Block(glowstone) }
  ]
  sea_level = 0
  noise = { min_y = -256, height = 512 }
  legacy_random_source = false
  noise_router = {
    temperature = Noise(simple, 1, 1)
    continents = 0
    erosion = 0
    depth = 0
    ridges = 0
    chunk_surface_level = 0
    final_density = Max(
      left = Gradient(axis = y, from_coordinate = 0, to_coordinate = 256, from_value = 0, to_value = 1),
      right = Noise({
        normalize = false
        base_octave = 4
        base_amplitude = 6
      }, 1, 1),
    )
    vegetation = Noise({ base_octave = -2, amplitude_modifiers = [1, 1] }, 1, 1)
  }
  aquifers = { barrier = 0, fluid_level_floodedness = 0, fluid_level_spread = 0, lava = 0, exclusion = 0, surface_level = 0 }
  spawn_target = []
}
`)
	settings := compact(out["data/demo/worldgen/noise_settings/custom.json"])
	for _, want := range []string{
		`"right":{"type":"minecraft:noise","noise":"demo:mms_gen/custom/noise_2"`,
		`"vegetation":{"type":"minecraft:noise","noise":"demo:mms_gen/custom/noise"`,
		`"temperature":{"type":"minecraft:noise","noise":"demo:simple"`,
	} {
		if !strings.Contains(settings, want) {
			t.Errorf("noise settings lack %s:\n%s", want, settings)
		}
	}
	if got := compact(out["data/demo/worldgen/noise/mms_gen/custom/noise_2.json"]); !strings.Contains(got, `"base_octave":4`) || !strings.Contains(got, `"normalize":false`) {
		t.Errorf("hoisted noise = %q", got)
	}
	if got := compact(out["data/demo/worldgen/noise/mms_gen/custom/noise.json"]); !strings.Contains(got, `"base_octave":-2`) {
		t.Errorf("second hoisted noise = %q", got)
	}
	// Inline density functions stay inline: the game takes them.
	for path := range out {
		if strings.Contains(path, "density_function/mms_gen") || strings.Contains(path, "material_rule/mms_gen") {
			t.Errorf("unexpected hoisted file %s", path)
		}
	}
}

// TestHoistIDOnly checks a field that only takes an id: noise_threshold's
// noise.
func TestHoistIDOnly(t *testing.T) {
	out := emitText(t, `
SurfaceRule surface = [
  If(NoiseThreshold({ base_octave = -5, amplitude_modifiers = [1, 0.5] }, 0, 1)) { Block(sand) }
  Block(stone)
]
`)
	rule := compact(out["data/demo/worldgen/material_rule/surface.json"])
	if !strings.Contains(rule, `"noise":"demo:mms_gen/surface/noise"`) {
		t.Errorf("surface rule = %s", rule)
	}
	if got := compact(out["data/demo/worldgen/noise/mms_gen/surface/noise.json"]); !strings.Contains(got, `"base_octave":-5`) {
		t.Errorf("hoisted noise = %q", got)
	}
}

// TestHoistErrors checks that a bad inline value reports against the
// hoisted kind.
func TestHoistErrors(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", "DensityFunction d = Noise({ base_octave = \"x\", amplitude_modifiers = [1] }, 1, 1)\n")
	r, _ := p.Analyze(context.Background())
	var errs []string
	for _, d := range r.Diagnostics("a.mms") {
		if d.Severity == syntax.SeverityError {
			errs = append(errs, d.Message)
		}
	}
	if len(errs) != 1 || !strings.Contains(errs[0], "integer") {
		t.Fatalf("errors = %q", errs)
	}
}
