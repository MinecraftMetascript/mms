package project

import (
	"sort"
	"strings"
	"testing"
)

const customBiomeSrc = `const climate_zero: ClimateParameters = { temperature = 0, humidity = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }

CustomBiome ash_fields {
  climate = { temperature = [0.5, 1], humidity = [-1, 0], continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }
  biome = { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [], features = [] }
  surface_rule = If(StoneDepth(Floor)) { Block(basalt) }
  tags = [#is_overworld, #hot]
}

BiomeTag hot = [minecraft:desert]

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

	// Each tag gets the biome, merged with the tag's own values.
	if tag := emitted(t, r, "data/minecraft/tags/worldgen/biome/is_overworld.json"); !strings.Contains(tag, `"demo:ash_fields"`) {
		t.Errorf("is_overworld:\n%s", tag)
	}
	if tag := emitted(t, r, "data/demo/tags/worldgen/biome/hot.json"); !strings.Contains(tag, `"demo:ash_fields"`) || !strings.Contains(tag, `"minecraft:desert"`) {
		t.Errorf("hot:\n%s", tag)
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

// An instance whose noise entries no biome source lists, or whose
// surface rule no material_rule names, is reported by CustomBiome's
// check, which asks used() once the whole project is read.
func TestCustomBiomeUsed(t *testing.T) {
	src := customBiomeSrc + `
CustomBiome lost {
  climate = climate_zero
  biome = { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [], features = [] }
  surface_rule = Block(stone)
}

Function show {
  /say ${frost_fields.noise_entries}
}
`
	_, r := assetProject(t, nil, src)
	var got []string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/unplaced-biome" || d.Code == "lint/unused-surface-rule" {
			got = append(got, src[d.Sp.Start:d.Sp.End]+": "+d.Message)
		}
	}
	sort.Strings(got)
	want := []string{
		"lost: nothing reads lost.noise_entries, so no biome source places the biome",
		"surface_rule = Block(stone): nothing references the surface rule demo:lost, so its ground is not made until a noise settings' material_rule lists it",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("used:\n%s", strings.Join(got, "\n"))
	}
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/unused" && strings.Contains(d.Message, "lost") {
			t.Errorf("checked output also reported as unused: %s", d.Message)
		}
	}
}

// A mistake in a climate is reported where it is written, and a
// CustomBiome without a surface_rule has no surface rule to name.
func TestCustomBiomeErrors(t *testing.T) {
	src := strings.Replace(customBiomeSrc, "humidity = [-1, 0]", "humidty = [-1, 0]", 1)
	src = strings.Replace(src, "[ash_fields, Block(stone)]", "[ash_fields, frost_fields]", 1)
	src = strings.Replace(src, "#is_overworld, #hot", "#is_overworld, plains", 1)
	_, r := assetProject(t, nil, src)
	got := diagText(r)
	for _, want := range []string{
		"mms/unknown-field: unknown field 'humidty' in ClimateParameters; did you mean 'humidity'?",
		"mms/type-mismatch: 'frost_fields' is a Biome, not a SurfaceRule",
		"mms/type-mismatch: expected a worldgen/biome tag (#name) here",
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

// A mistake in a typed const's value is reported once, where it is
// written, not again at each use or in each const built from it. A
// problem that only a use has is still reported there.
func TestConstErrorNotRepeated(t *testing.T) {
	src := `Biome ash { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [], features = [] }
const c: ClimateParameters = { temperature = [0.5, 1], humidty = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }
const src: BiomeSource = MultiNoise(biomes = [{ biome = ash, parameters = c }])
Dimension d { type = overworld, generator = Noise(settings = overworld, biome_source = src) }
const n: int = 3
Biome hot { has_precipitation = n, temperature = 2, downfall = n, effects = { water_color = "#3F76E4" }, carvers = [], features = [] }
`
	_, r := assetProject(t, nil, src)
	count := map[string]int{}
	for _, d := range r.Diagnostics("main.mms") {
		if !d.Echo {
			count[d.Code]++
		}
	}
	if count["mms/unknown-field"] != 1 || count["mms/missing-field"] != 1 {
		t.Errorf("want each climate problem once:\n%s", diagText(r))
	}
	if count["mms/type-mismatch"] != 1 {
		t.Errorf("want the misused const reported where it is used:\n%s", diagText(r))
	}
}

// Only a problem in the const's own value makes a use's problem a
// repeat: the same problem elsewhere in the instance does not, and a
// local that hides the instance's name is not a use of its const.
func TestConstRepeatNeedsTheConstsValue(t *testing.T) {
	src := `CustomBiome a {
  climate = { temperature = 0, humidity = 0, continentalness = 0, erosion = 0, weirdness = 0, depth = 0, offset = 0 }
  biome = { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [{ type = canyon, biome = 1 }], features = [] }
}
Dimension d { type = overworld, generator = Noise(settings = overworld, biome_source = MultiNoise(biomes = [...a.noise_entries])) }
Biome hot { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = a.noise_entries, features = [] }
Biome cold { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [for a in [{ noise_entries = { biome = 1 } }]: a.noise_entries], features = [] }
Template carvers_of(a: { noise_entries: [ConfiguredCarver] }): [ConfiguredCarver] = a.noise_entries
`
	_, r := assetProject(t, nil, src)
	n := 0
	for _, d := range r.Diagnostics("main.mms") {
		if !d.Echo && d.Message == "unknown field 'biome' in ConfiguredCarver" {
			n++
		}
	}
	// The instance's carver, the misused const in hot, and cold's own
	// carver, whose loop variable a hides the instance.
	if n != 3 {
		t.Errorf("want each carver problem reported, got %d:\n%s", n, diagText(r))
	}
	uses := r.typedConstUses()["main.mms"]
	for _, u := range uses {
		if text := src[u.sp.Start:u.sp.End]; u.sp.Start > strings.Index(src, "Biome cold") {
			t.Errorf("a loop variable or param a hides the instance, but %q counts as a use", text)
		}
	}
	if len(uses) != 2 {
		t.Errorf("want 2 uses of a.noise_entries, got %d", len(uses))
	}
}

// used() names a const of the emit block or a kind it emits, and only a
// check block may call it.
func TestUsedArgs(t *testing.T) {
	src := `Blueprint Ore {
  check {
    if (!used(entries)) { warn(self, "x", "a") }
    if (!used(Feature)) { warn(self, "y", "b") }
  }
  emit {
    primary Biome ${self.path} = { has_precipitation = false, temperature = 2, downfall = 0, effects = { water_color = "#3F76E4" }, carvers = [], features = [] }
    const noise_entries: [BiomeNoiseEntry] = []
  }
}
Ore iron {}
`
	_, r := assetProject(t, nil, src)
	got := diagText(r)
	for _, want := range []string{
		"Ore's emit block has no const or declaration kind 'entries'",
		"Ore's emit block has no const or declaration kind 'Feature'",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("want %q in:\n%s", want, got)
		}
	}
}

// An unqualified tag in a CustomBiome's tags is the project's own only
// when the project declares a biome tag of that name, written plainly
// or with its namespace; a tag of another registry does not count.
func TestCustomBiomeTagNamespace(t *testing.T) {
	src := strings.Replace(customBiomeSrc, "BiomeTag hot = [minecraft:desert]", "BiomeTag demo:hot = [minecraft:desert]\nBlockTag is_overworld = [stone]", 1)
	_, r := assetProject(t, nil, src)
	files := emittedFiles(r)
	if _, ok := files["data/demo/tags/worldgen/biome/is_overworld.json"]; ok {
		t.Error("a BlockTag of the name moved #is_overworld into the project's namespace")
	}
	if tag := emitted(t, r, "data/minecraft/tags/worldgen/biome/is_overworld.json"); !strings.Contains(tag, `"demo:ash_fields"`) {
		t.Errorf("is_overworld:\n%s", tag)
	}
	if tag := emitted(t, r, "data/demo/tags/worldgen/biome/hot.json"); !strings.Contains(tag, `"demo:ash_fields"`) {
		t.Errorf("hot:\n%s", tag)
	}
	if _, ok := files["data/minecraft/tags/worldgen/biome/hot.json"]; ok {
		t.Error("#hot went to minecraft although the project declares demo:hot")
	}
}
