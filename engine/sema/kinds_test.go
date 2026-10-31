package sema

import (
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

func TestKinds(t *testing.T) {
	ks := NewKinds(spec.MustDefault())
	for name, want := range map[string]string{
		"DensityFunction": "worldgen/density_function",
		"LootTable":       "loot_table",
		"SurfaceRule":     "worldgen/material_rule",
		"NoiseSettings":   "worldgen/noise_settings",
		"Noise":           "worldgen/noise",
		"Recipe":          "recipe",
		"BlockTag":        "tags/block",
		"BiomeTag":        "tags/worldgen/biome",
		"FunctionTag":     "tags/function",
		"Function":        "",
	} {
		k, ok := ks.Lookup(name)
		if !ok || k.Resource != want {
			t.Errorf("Lookup(%s) = %+v, %v; want resource %q", name, k, ok, want)
		}
	}
	if _, ok := ks.Lookup("Bogus"); ok {
		t.Error("Bogus resolved")
	}
}
