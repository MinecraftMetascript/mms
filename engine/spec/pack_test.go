package spec

import (
	"testing"
	"time"
)

func TestDefaultPackLoads(t *testing.T) {
	start := time.Now()
	p, err := Default()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("decoded %d types in %s", len(p.Types), time.Since(start))
	if p.Version != "26.3" || p.PackFormat != 121 {
		t.Fatalf("got version %s format %d", p.Version, p.PackFormat)
	}
}

// The old hand-written specs covered these; the generated pack must too.
func TestCoversOldConstructs(t *testing.T) {
	p := MustDefault()
	for _, kind := range []string{
		"worldgen/noise", "worldgen/density_function", "worldgen/noise_settings",
		"dimension", "dimension_type", "loot_table", "recipe", "predicate",
		"item_modifier", "advancement", "worldgen/biome",
	} {
		if _, ok := p.Resources[kind]; !ok {
			t.Errorf("missing resource kind %s", kind)
		}
	}
	df := p.DispatchKeys("minecraft:density_function")
	// 26.3 reworked density functions: cache_2d and friends became `cache`,
	// y_clamped_gradient became `gradient`, shifted_noise folded into noise.
	for _, want := range []string{
		"abs", "add", "blend_density", "cache", "clamp", "constant", "cube", "div",
		"gradient", "half_negative", "interpolated", "max", "min", "mul", "noise",
		"old_blended_noise", "quarter_negative", "range_choice", "shift", "shift_a",
		"shift_b", "spline", "square", "squeeze", "sub",
	} {
		if _, ok := p.Dispatch("minecraft:density_function", want); !ok {
			t.Errorf("missing density function type %s", want)
		}
	}
	t.Logf("%d density function types: %v", len(df), df)
	for _, d := range []string{"minecraft:material_rule", "minecraft:material_condition"} {
		if len(p.DispatchKeys(d)) == 0 {
			t.Errorf("dispatcher %s is empty", d)
		}
	}
	if !p.HasID("block", "minecraft:orange_terracotta") || !p.HasID("worldgen/biome", "badlands") {
		t.Error("registries are missing vanilla ids")
	}
	if p.Blocks["oak_log"].Defaults["axis"] != "y" {
		t.Error("block states are missing")
	}
	if p.Commands.Children["scoreboard"] == nil {
		t.Error("command tree is missing")
	}
}

// Clamp.input became reference-or-inline in 26.2, so 26.3 must keep only
// the newer member (decision D4).
func TestVersionFilter(t *testing.T) {
	p := MustDefault()
	clamp := p.Types["::java::data::worldgen::density_function::Clamp"]
	if clamp == nil {
		t.Fatal("Clamp type missing")
	}
	for _, f := range clamp.Fields {
		if f.Key != "input" {
			continue
		}
		in := f.Type
		if in.Kind == "union" {
			if len(in.Members) != 1 {
				t.Fatalf("input union has %d members, want 1", len(in.Members))
			}
			in = in.Members[0]
		}
		if in.Path != "::java::data::worldgen::density_function::DensityFunctionRef" {
			t.Errorf("input resolves to %s", in.Path)
		}
		return
	}
	t.Error("Clamp has no input field")
}
