package spec

import (
	"strings"
	"testing"
)

func TestVanillaFiles(t *testing.T) {
	p := MustDefault()
	v := p.Vanilla()
	if v == nil {
		t.Fatal("the default pack has no vanilla files")
	}
	for _, tc := range []struct{ folder, id, want string }{
		{"worldgen/noise", "minecraft:ridge", "data/minecraft/worldgen/noise/ridge.json"},
		{"worldgen/template_pool", "village/plains/houses", "data/minecraft/worldgen/template_pool/village/plains/houses.json"},
		{"tags/worldgen/biome", "#minecraft:is_overworld", "data/minecraft/tags/worldgen/biome/is_overworld.json"},
		{"worldgen/noise", "other:ridge", ""},
		{"worldgen/noise", "no_such_noise", ""},
	} {
		got := v.VanillaPath(tc.folder, tc.id)
		if tc.want == "" {
			if got != "" {
				t.Errorf("VanillaPath(%s, %s) = %s, want none", tc.folder, tc.id, got)
			}
			continue
		}
		if got != VanillaPrefix+tc.want {
			t.Errorf("VanillaPath(%s, %s) = %s", tc.folder, tc.id, got)
			continue
		}
		if text, ok := v.File(got); !ok || !strings.HasPrefix(text, "{\n  \"") {
			t.Errorf("File(%s) = %q, %v", got, text, ok)
		}
	}
	// Every registry with ids in the vanilla files has a file for each id.
	for folder, files := range v {
		reg := folder
		if t, ok := strings.CutPrefix(folder, "tags/"); ok {
			reg = "tag/" + t
		}
		for _, id := range p.Registries[reg] {
			if _, ok := files[id]; !ok {
				t.Errorf("no vanilla file for %s %s", folder, id)
				break
			}
		}
	}
	// A decoded pack that is not the embedded one has none.
	if (&Pack{}).Vanilla() != nil {
		t.Error("an empty pack has vanilla files")
	}
}
