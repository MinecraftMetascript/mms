package project

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// orderSrc declares placed features named in steps and one biome per
// entry of biomes, with that step list as its vegetal decoration.
func orderSrc(biomes map[string]string, extra string) string {
	var b strings.Builder
	names := map[string]bool{}
	for _, list := range biomes {
		for _, f := range strings.Split(list, ",") {
			if f = strings.TrimSpace(f); f != "" && !strings.Contains(f, ":") && !strings.HasPrefix(f, "#") {
				names[f] = true
			}
		}
	}
	for n := range names {
		fmt.Fprintf(&b, "PlacedFeature %s = { feature = minecraft:oak, placement = [] }\n", n)
	}
	for _, n := range []string{"one", "two", "three"} {
		if list, ok := biomes[n]; ok {
			if !strings.HasPrefix(list, "#") {
				list = "[" + list + "]"
			}
			fmt.Fprintf(&b, `Biome %s = {
  temperature = 0.5
  downfall = 0.5
  has_precipitation = false
  effects = { water_color = "#3f76e4" }
  carvers = []
  features = [[], [], [], [], [], [], [], [], [], %s, []]
}
`, n, list)
		}
	}
	b.WriteString(extra)
	return b.String()
}

func orderDiags(t *testing.T, src string) []string {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	var out []string
	for _, d := range r.Diagnostics("a.mms") {
		if d.Severity == 1 && d.Code != "mms/feature-order-cycle" {
			t.Fatalf("unexpected error %s: %s", d.Code, d.Message)
		}
		if d.Code == "mms/feature-order-cycle" {
			out = append(out, d.Message)
		}
	}
	return out
}

func TestFeatureOrderCycle(t *testing.T) {
	// The Cave Siege case: rootwood_hollow runs cave_vines before
	// glow_caps, darkroot_grove runs glow_caps, spore_blossom, cave_vines.
	got := orderDiags(t, orderSrc(map[string]string{
		"one": "trees, grass, minecraft:cave_vines, glow_caps, moss, minecraft:spore_blossom",
		"two": "trees2, grass, glow_caps, minecraft:spore_blossom, minecraft:cave_vines, minecraft:classic_vines_cave_feature",
	}, ""))
	if len(got) != 2 {
		t.Fatalf("got %d cycle errors, want one per biome: %q", len(got), got)
	}
	for _, m := range got {
		for _, want := range []string{"step 10 (vegetal_decoration)", "one runs", "two runs", "minecraft:cave_vines before glow_caps", "minecraft:spore_blossom before minecraft:cave_vines"} {
			if !strings.Contains(m, want) {
				t.Errorf("message %q lacks %q", m, want)
			}
		}
	}
	if !strings.HasPrefix(got[0], "feature order cycle in step 10 (vegetal_decoration): one runs") {
		t.Errorf("first biome's message should lead with itself: %q", got[0])
	}

	// The same features in a consistent order are fine.
	if got := orderDiags(t, orderSrc(map[string]string{
		"one": "trees, grass, minecraft:cave_vines, glow_caps, moss, minecraft:spore_blossom",
		"two": "trees2, grass, minecraft:cave_vines, glow_caps, minecraft:spore_blossom, minecraft:classic_vines_cave_feature",
	}, "")); len(got) != 0 {
		t.Errorf("consistent order reported: %q", got)
	}

	// Tags expand in place.
	if got := orderDiags(t, orderSrc(map[string]string{
		"one": "#pair",
		"two": "b, a",
	}, "PlacedFeatureTag pair = [a, b]\n")); len(got) != 2 {
		t.Errorf("cycle through a tag: %q", got)
	}

	// A feature listed twice in one biome is a cycle on its own.
	got = orderDiags(t, orderSrc(map[string]string{"one": "a, b, a"}, ""))
	if len(got) != 1 || !strings.Contains(got[0], "lists a more than once") {
		t.Errorf("duplicate = %q", got)
	}

	// Biomes in different dimensions never share a feature order.
	dims := `Dimension d1 = { type = minecraft:overworld, generator = { type = noise, settings = minecraft:overworld, biome_source = { type = fixed, biome = one } } }
Dimension d2 = { type = minecraft:overworld, generator = { type = noise, settings = minecraft:overworld, biome_source = { type = fixed, biome = two } } }
`
	if got := orderDiags(t, orderSrc(map[string]string{"one": "a, b", "two": "b, a"}, dims)); len(got) != 0 {
		t.Errorf("separate dimensions reported: %q", got)
	}
	dims = `Dimension d1 = { type = minecraft:overworld, generator = { type = noise, settings = minecraft:overworld, biome_source = { type = checkerboard, biomes = [one, two] } } }
`
	if got := orderDiags(t, orderSrc(map[string]string{"one": "a, b", "two": "b, a", "three": "a, b"}, dims)); len(got) != 2 {
		t.Errorf("shared dimension = %q", got)
	}
}
