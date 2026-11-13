package project

import (
	"strings"
	"testing"
)

const customBiomeSrc = `const climate_zero: ClimateParameters = { temperature = 0, humidity = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }

CustomBiome ash_fields {
  climate = { temperature = [0.5, 1], humidity = [-1, 0], continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }
  biome = { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [], features = [] }
  surface_rule = If(StoneDepth(Floor)) { Block(basalt) }
}

CustomBiome frost_fields {
  climate = [
    { temperature = [-1, -0.5], humidity = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 },
    climate_zero,
  ]
  biome = { has_precipitation = true, temperature = -0.5, downfall = 0.5, effects = { water_color = "#3D57D6" }, carvers = [], features = [] }
}

SurfaceRule surface = [ash_fields, Block(stone)]

Dimension fields {
  type = overworld
  generator = Noise(
    settings = overworld,
    biome_source = MultiNoise(biomes = [...ash_fields.noise_entries, ...frost_fields.noise_entries]),
  )
}
`

// CustomBiome writes the biome and its surface rule, and its noise
// entries fill a multi-noise biome source, one for each climate.
func TestCustomBiome(t *testing.T) {
	_, r := assetProject(t, nil, customBiomeSrc)
	if d := diagText(r); d != "" {
		t.Fatalf("diagnostics:\n%s", d)
	}
	if b := emitted(t, r, "data/demo/worldgen/biome/ash_fields.json"); !strings.Contains(b, `"temperature": 2`) {
		t.Errorf("biome:\n%s", b)
	}
	emitted(t, r, "data/demo/worldgen/biome/frost_fields.json")

	rule := emitted(t, r, "data/demo/worldgen/material_rule/ash_fields.json")
	for _, want := range []string{`"biome_is": "demo:ash_fields"`, `"result_state": "minecraft:basalt"`, `"type": "minecraft:stone_depth"`} {
		if !strings.Contains(rule, want) {
			t.Errorf("surface rule lacks %s:\n%s", want, rule)
		}
	}
	// No surface_rule, no surface rule file.
	if _, ok := emittedFiles(r)["data/demo/worldgen/material_rule/frost_fields.json"]; ok {
		t.Error("frost_fields has no surface_rule but wrote one")
	}
	if s := emitted(t, r, "data/demo/worldgen/material_rule/surface.json"); !strings.Contains(s, `"demo:ash_fields"`) {
		t.Errorf("surface:\n%s", s)
	}

	dim := emitted(t, r, "data/demo/dimension/fields.json")
	if n := strings.Count(dim, `"biome": "demo:frost_fields"`); n != 2 {
		t.Errorf("want 2 frost_fields entries, got %d:\n%s", n, dim)
	}
	if n := strings.Count(dim, `"biome": "demo:ash_fields"`); n != 1 {
		t.Errorf("want 1 ash_fields entry, got %d:\n%s", n, dim)
	}
	if !strings.Contains(dim, `"humidity": [`) {
		t.Errorf("climate ranges lost:\n%s", dim)
	}
}

// A mistake in a climate is reported where it is written, and a
// CustomBiome without a surface_rule has no surface rule to name.
func TestCustomBiomeErrors(t *testing.T) {
	src := strings.Replace(customBiomeSrc, "humidity = [-1, 0]", "humidty = [-1, 0]", 1)
	src = strings.Replace(src, "[ash_fields, Block(stone)]", "[ash_fields, frost_fields]", 1)
	_, r := assetProject(t, nil, src)
	got := diagText(r)
	for _, want := range []string{
		"mms/unknown-field: unknown field 'humidty' in ClimateParameters; did you mean 'humidity'?",
		"mms/type-mismatch: 'frost_fields' is a Biome, not a SurfaceRule",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("want %q in:\n%s", want, got)
		}
	}
	var at bool
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "mms/unknown-field" && strings.HasPrefix(src[d.Sp.Start:], "humidty") {
			at = true
		}
	}
	if !at {
		t.Errorf("unknown-field not reported at the climate field:\n%s", got)
	}
}
