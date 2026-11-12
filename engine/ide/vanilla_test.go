package ide

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Go to definition on a vanilla id opens vanilla's own JSON file.
func TestDefinitionVanilla(t *testing.T) {
	for _, tc := range []struct{ src, want string }{
		{"DensityFunction x = Noise(ri|dge)\n", "worldgen/noise/ridge.json"},
		{"DensityFunction x = Noise(minecraft:ri|dge)\n", "worldgen/noise/ridge.json"},
		{"DensityFunction x = minecraft:overworld/conti|nents * 2\n", "worldgen/density_function/overworld/continents.json"},
		{"PlacedFeatureTag ores = [ore_di|rt]\n", "worldgen/placed_feature/ore_dirt.json"},
		{"BlockTag b = [#minecraft:lo|gs]\n", "tags/block/logs.json"},
		{"BlockTag b = [st|one]\n", ""}, // blocks are not files
		{"Model m { parent = block/cube_co|lumn }\n", "@assets/models/block/cube_column.json"},
		{"BlockState minecraft:stone { variants = { \"\" = { model = block/gran|ite } } }\n", "@assets/models/block/granite.json"},
	} {
		s, off := session(t, tc.src)
		defs := s.Definition("main.mms", off)
		if tc.want == "" {
			if len(defs) != 0 {
				t.Errorf("%q: definition = %+v, want none", tc.src, defs)
			}
			continue
		}
		want := spec.VanillaPrefix + "data/minecraft/" + tc.want
		if rest, ok := strings.CutPrefix(tc.want, "@assets/"); ok {
			want = spec.VanillaPrefix + "assets/minecraft/" + rest
		}
		if len(defs) != 1 || defs[0].Path != want {
			t.Errorf("%q: definition = %+v, want %s", tc.src, defs, want)
			continue
		}
		text, ok := s.P.Pack.Vanilla().File(defs[0].Path)
		if !ok || !strings.HasPrefix(text, "{\n  ") {
			t.Errorf("%q: file text = %q", tc.src, text)
		}
	}
}

func TestDefinitionVanillaCommand(t *testing.T) {
	s, off := session(t, "Function f {\n  /loot spawn ~ ~ ~ loot minecraft:chests/simple_dun|geon\n}\n")
	defs := s.Definition("main.mms", off)
	want := spec.VanillaPrefix + "data/minecraft/loot_table/chests/simple_dungeon.json"
	if len(defs) != 1 || defs[0].Path != want {
		t.Errorf("definition = %+v, want %s", defs, want)
	}
}
