package ide

import (
	"strings"
	"testing"
)

// TestCompletionMatrix walks the places completion should open, one row
// per context, so a gap does not come back unnoticed. Each row lists
// labels that must be offered (want), labels that must not (not), and the
// text the chosen item replaces (replace, "" for an empty span). A row
// with none set expects no items at all.
func TestCompletionMatrix(t *testing.T) {
	cases := []struct {
		name, src string
		want, not []string
		none      bool
		replace   string
	}{
		// The top level: declaration kinds.
		{name: "empty file", src: "|", want: []string{"Feature", "Recipe", "Biome", "Function", "OnTick", "Namespace"}},
		{name: "kind prefix", src: "Fe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind after declaration", src: "Feature f {}\nFe|", want: []string{"Feature", "PlacedFeature"}, replace: "Fe"},
		{name: "kind after function", src: "Function f {\n  /say hi\n}\nRec|", want: []string{"Recipe"}, replace: "Rec"},
		{name: "kind after namespace", src: "Namespace foo\nBio|", want: []string{"Biome"}, replace: "Bio"},
		{name: "kind before declaration", src: "Fe|\nFunction f {\n}", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind after word value", src: "DensityFunction x = foo\nFe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind after criterion", src: "Objective kills: deathCount\nFe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind mid-word", src: "Fe|ature f = {}", want: []string{"Feature"}, replace: "Feature"},
		{name: "kind after comment", src: "// ores\nFe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "declaration name", src: "Feature |", none: true},
		{name: "declaration name typed", src: "Feature f|", none: true},
		{name: "namespace name", src: "Namespace |", none: true},

		// After `=`: the kind's values.
		{name: "feature value", src: "Feature f = |", want: []string{"Ore", "RandomNeighborSpread"}},
		{name: "recipe value", src: "Recipe r = |", want: []string{"Shaped", "Smelting"}},
		{name: "surface rule value", src: "SurfaceRule r = |", want: []string{"If", "Sequence"}},
		{name: "objective criterion", src: "Objective kills: |", want: []string{"dummy", "deathCount", "trigger"}},
		{name: "objective criterion typed", src: "Objective kills: de|", want: []string{"deathCount"}, replace: "de"},
		{name: "objective criterion old form", src: "Objective kills = |", want: []string{"dummy", "deathCount", "trigger"}},
		{name: "global type", src: "Global g: |", want: []string{"int", "bool", "string"}},

		// Helper arguments.
		{name: "helper args", src: "Feature f = Ore(|)", want: []string{"size", "targets"}},
		{name: "helper named arg prefix", src: "Feature f = Ore(si|)", replace: "si"},
		{name: "helper list arg", src: "Feature f = Ore(targets = [|])", want: []string{"TargetBlock"}},
		{name: "helper enum arg", src: "Feature f = BlockColumn(direction = |)", want: []string{"up", "down"}},
		{name: "helper id arg", src: "Recipe r = Smelting(ingredient = |)", want: []string{"iron_ore"}},
		{name: "helper enum category", src: "Recipe r = Smelting(category = |)", want: []string{"blocks", "food", "misc"}},
		{name: "positional arg", src: "DensityFunction x = Noise(|)", want: []string{"ridge"}},
		{name: "method", src: "DensityFunction x = Noise(minecraft:ridge).|", want: []string{"XzScale", "YScale"}},
		{name: "surface If condition", src: "SurfaceRule r = If(|)", want: []string{"AboveWater", "StoneDepth"}, not: []string{"Sequence"}},
		{name: "surface If rule", src: "SurfaceRule r = If(AboveWater(), |)", want: []string{"Block", "Sequence"}},
		{name: "surface If condition before block", src: "SurfaceRule r = If(|) {\n}", want: []string{"AboveWater", "StoneDepth"}, not: []string{"Sequence"}},
		{name: "surface If condition before rule", src: "SurfaceRule r = If(|) { Block(stone) }", want: []string{"AboveWater"}, not: []string{"Block"}},
		{name: "surface If condition in list", src: "SurfaceRule r = [\n  If(|) { Block(stone) },\n]", want: []string{"AboveWater"}},
		{name: "surface If condition in material_rule", src: "NoiseSettings ns {\n  material_rule = [\n    If(|) {\n    }\n  ]\n}", want: []string{"AboveWater", "Biome"}, not: []string{"aquifers"}},
		{name: "surface If declared condition", src: "SurfaceCondition wet = AboveWater()\nSurfaceRule r = If(|) {\n}", want: []string{"wet", "Biome"}},
		{name: "surface If declared condition operand", src: "SurfaceCondition wet = AboveWater()\nSurfaceRule r = If(wet && |) {\n}", want: []string{"wet", "Biome"}},
		{name: "surface If block body", src: "SurfaceRule stone_top = Block(stone)\nSurfaceRule r = If(AboveWater()) {\n  |\n}", want: []string{"stone_top", "Block", "If"}, not: []string{"AboveWater"}},
		{name: "surface If block body prefix", src: "SurfaceRule stone_top = Block(stone)\nSurfaceRule r = If(AboveWater()) {\n  sto|\n}", want: []string{"stone_top"}, replace: "sto"},
		{name: "surface If condition no block in list", src: "SurfaceRule r = [\n  If(|)\n  Block(stone)\n]", want: []string{"AboveWater", "StoneDepth"}, not: []string{"Sequence"}},
		{name: "surface If condition no block in field", src: "NoiseSettings ns {\n  material_rule = If(|)\n  sea_level = 0\n}", want: []string{"AboveWater"}, not: []string{"aquifers"}},
		{name: "surface If condition no block nested", src: "SurfaceRule r = If(AboveWater()) {\n  If(|)\n}", want: []string{"AboveWater"}, not: []string{"Block"}},
		{name: "surface If rule declared", src: "SurfaceRule stone_top = Block(stone)\nSurfaceRule r = If(AboveWater(), |)", want: []string{"stone_top", "Block"}},

		// Objects and nested maps.
		{name: "object fields", src: "Biome b {\n  |\n}", want: []string{"effects", "temperature"}},
		{name: "nested object", src: "Biome b {\n  effects = {\n    |\n  }\n}", want: []string{"water_color"}, not: []string{"temperature"}},
		{name: "object in helper", src: "Recipe r = Smelting(result = { |})", want: []string{"id", "count"}},
		{name: "object in list", src: "Feature f = Ore(targets = [{ |}])", want: []string{"target", "state"}},
		{name: "dispatched on sibling", src: "Advancement a {\n  display = {\n    |\n  }\n}", want: []string{"icon", "title", "frame", "background"}, not: []string{"criteria"}},
		{name: "dispatched on sibling set", src: "Advancement a {\n  parent = minecraft:story/root\n  display = {\n    |\n  }\n}", want: []string{"icon", "title"}, not: []string{"criteria", "background"}},
		{name: "dispatched value", src: "Advancement a { display = { frame = |} }", want: []string{"challenge", "goal", "task"}},
		{name: "map keys", src: "Enchantment e {\n  effects = {\n    |\n  }\n}", want: []string{"damage", "attributes"}, not: []string{"max_level"}},
		{name: "map keys with slash", src: "Biome b {\n  attributes = {\n    vis|\n  }\n}", want: []string{"visual/sky_color"}, not: []string{"temperature"}, replace: "vis"},
		{name: "item components", src: "Recipe r = Shapeless([emerald], { id = emerald\n components = {\n |\n }\n})", want: []string{"item_name", "max_stack_size"}, not: []string{"count"}},
		{name: "map value", src: "Enchantment e {\n  effects = {\n    minecraft:damage = [|]\n  }\n}", want: []string{"DamageEnchantmentEffect"}},
		{name: "empty field value before next field", src: "NoiseSettings ns {\n  default_block = |\n  default_fluid = water\n  sea_level = 0\n}", want: []string{"stone", "FullBlockState"}, not: []string{"default_fluid", "aquifers"}},
		{name: "empty field value before close", src: "NoiseSettings ns {\n  sea_level = 0\n  default_block = |\n}", want: []string{"stone"}, not: []string{"default_fluid"}},
		{name: "empty field value before quoted key", src: "Biome b {\n  temperature = |\n  \"downfall\" = 0\n}", not: []string{"downfall", "effects"}},
		{name: "surface rule list in field", src: "NoiseSettings ns {\n  spawn_target = []\n  material_rule = [\n    |\n  ]\n}", want: []string{"Block", "If", "Sequence"}, not: []string{"aquifers", "debug_functions"}},
		{name: "surface rule list second item", src: "NoiseSettings ns {\n  material_rule = [Block(stone), |]\n}", want: []string{"Block", "If"}, not: []string{"aquifers"}},
		{name: "object where no object fits", src: "NoiseSettings ns {\n  sea_level = {\n    |\n  }\n}", not: []string{"default_block", "aquifers"}},
		{name: "inline definition object", src: "Advancement a {\n  parent = {\n    |\n  }\n}", want: []string{"criteria"}},
		{name: "function reference", src: "Function g {}\nAdvancement a { rewards = { function = |} }", want: []string{"g"}},

		// Tags.
		{name: "tag entries", src: "ItemTag gems = [|]", want: []string{"diamond", "#logs", "Optional"}, not: []string{"#gems"}},
		{name: "tag entry prefix", src: "ItemTag gems = [dia|]", want: []string{"diamond"}, replace: "dia"},
		{name: "tag entry tag", src: "BlockTag t = [#|]", want: []string{"#logs"}, replace: "#"},
		{name: "tag optional", src: "ItemTag t = [Optional(|)]", want: []string{"diamond"}},
		{name: "own tag reference", src: "ItemTag gems = [diamond]\nItemTag more = [|]", want: []string{"#gems"}},

		// Function bodies.
		{name: "statement", src: "Function f {\n  /say hi\n  |\n}", want: []string{"As", "let", "if"}},
		{name: "test line", src: "Function f {\n}\nTest t {\n  run f()\n  |\n}", want: []string{"expect", "run", "setup"}, not: []string{"let", "As"}},
		{name: "test setup", src: "Global g: int = 0\nFunction f {\n}\nTest t {\n  setup {\n    |\n  }\n  run f()\n}", want: []string{"let", "g"}, not: []string{"expect"}},
		{name: "test expect value", src: "Global g: int = 0\nFunction f {\n}\nTest t {\n  run f()\n  expect g == |\n}", want: []string{"g"}},
		{name: "statement local", src: "Function f {\n  let x = 1\n  x|\n}", want: []string{"x", "As"}, replace: "x"},
		{name: "statement global", src: "Global g: int = 0\nFunction f {\n  g|\n}", want: []string{"g"}, replace: "g"},
		{name: "statement keyword prefix", src: "Function f {\n  ret|\n}", want: []string{"return"}, replace: "ret"},
		{name: "let value", src: "Function f(n: int) {\n  let x = 1\n  let y = |\n}", want: []string{"n", "x", "true"}, not: []string{"y", "As"}},
		{name: "let value after operator", src: "Global g: int = 0\nFunction f {\n  let y = g + |\n}", want: []string{"g"}, not: []string{"y", "let"}},
		{name: "if condition", src: "Function f {\n  let x = 1\n  if x > |\n}", want: []string{"x"}, not: []string{"let"}},
		{name: "return value", src: "Function f(n: int): int {\n  return |\n}", want: []string{"n"}, not: []string{"As"}},
		{name: "assignment value", src: "Global g: int = 0\nFunction f {\n  g = |\n}", want: []string{"g"}, not: []string{"As"}},
		{name: "loop variable", src: "Function f {\n  for i in 0..10 {\n    let y = |\n  }\n}", want: []string{"i"}},
		{name: "loop variable out of scope", src: "Function f {\n  for i in 0..10 {\n  }\n  let y = |\n}", not: []string{"i"}},
		{name: "function call arg", src: "Function g(a: int): int {\n  return a\n}\nFunction f {\n  let y = g(|)\n}", want: []string{"g", "true"}},
		{name: "objective in expression", src: "Objective o: dummy\nFunction f {\n  let y = |\n}", want: []string{"o"}},
		{name: "parameter type", src: "Function f(n: |) {\n}", want: []string{"int", "bool", "string"}},
		{name: "let type", src: "Function f {\n  let s: |\n}", want: []string{"int", "string"}},
		{name: "execute selector", src: "Function f {\n  As(|)\n}", want: []string{"@a", "@s"}},
		{name: "execute selector prefix", src: "Function f {\n  As(@|)\n}", want: []string{"@a"}, replace: "@"},
		{name: "execute selector typed", src: "Function f {\n  As(@a|)\n}", want: []string{"@a"}, replace: "@a"},
		{name: "execute anchor", src: "Function f {\n  Anchored(|)\n}", want: []string{"eyes", "feet"}},
		{name: "execute dimension", src: "Function f {\n  In(|)\n}", want: []string{"the_nether"}},
		{name: "execute chain", src: "Function f {\n  As(@a).|\n}", want: []string{"At", "Positioned"}},
		{name: "hook body", src: "OnTick {\n  |\n}", want: []string{"As", "return"}},

		// Commands.
		{name: "command name", src: "Function f {\n  /scoreb|\n}", want: []string{"scoreboard"}, replace: "scoreb"},
		{name: "command criterion", src: "Function f {\n  /scoreboard objectives add x |\n}", want: []string{"dummy", "deathCount"}},
		{name: "command item", src: "Function f {\n  /give @s minecraft:dia|\n}", want: []string{"diamond"}, replace: "minecraft:dia"},
		{name: "command in execute block", src: "Function f {\n  As(@a).At(@s) {\n    /|\n  }\n}", want: []string{"say"}},
		{name: "command value name", src: "Function f {\n  let n = /da|\n}", want: []string{"data"}, replace: "da"},
		{name: "command value argument", src: "Function f {\n  let n = /data get |\n}", want: []string{"block", "entity", "storage"}},
		{name: "command value assignment", src: "Global g: int = 0\nFunction f {\n  g = /time query |\n}", want: []string{"gametime"}},
		{name: "command value interpolation", src: "Function f(k: int) {\n  let n = /data get storage demo:x ${|}\n}", want: []string{"k"}},
		{name: "command value inline", src: "CustomItem w { base = stick, on_use_on_block = run {\n  let ok: bool = /exec|\n} }", want: []string{"execute"}, replace: "exec"},

		{name: "interpolation", src: "Function f(n: int) {\n  let x = 1\n  /say ${|}\n}", want: []string{"n", "x"}, not: []string{"say", "As"}},
		{name: "interpolation prefix", src: "Global g: int = 0\nFunction f {\n  /say ${g|}\n}", want: []string{"g"}, replace: "g"},
		{name: "interpolation custom item", src: "CustomItem shard { base = stick }\nFunction f {\n  /give @s ${|}\n}", want: []string{"shard"}},
		{name: "after interpolation", src: "Function f {\n  let x = 1\n  /scoreboard players set @s ${x} |\n}", not: []string{"x"}},

		// Inline function blocks, where a function id is needed: the same
		// as a Function body.
		{name: "inline command start", src: "Advancement a { rewards = { function = run {\n  /|\n} } }", want: []string{"say", "scoreboard"}},
		{name: "inline command name", src: "Advancement a { rewards = { function = run {\n  /scoreb|\n} } }", want: []string{"scoreboard"}, replace: "scoreb"},
		{name: "inline statement", src: "Advancement a { rewards = { function = run {\n  /say hi\n  |\n} } }", want: []string{"As", "let", "return"}, not: []string{"experience", "loot"}},
		{name: "inline local", src: "Advancement a { rewards = { function = run {\n  let x = 1\n  let y = |\n} } }", want: []string{"x", "true"}, not: []string{"As"}},
		{name: "inline interpolation", src: "Advancement a { rewards = { function = run {\n  let x = 1\n  /say ${|}\n} } }", want: []string{"x"}},
		{name: "inline execute selector", src: "Advancement a { rewards = { function = run {\n  As(|)\n} } }", want: []string{"@a", "@s"}},
		{name: "inline execute chain", src: "Advancement a { rewards = { function = run {\n  As(@a).|\n} } }", want: []string{"At", "Positioned"}},
		{name: "item hook command", src: "CustomItem w {\n  base = stick\n  on_right_click = run {\n    /|\n  }\n}", want: []string{"say", "give"}},
		{name: "item hook statement", src: "CustomItem w {\n  base = stick\n  on_right_click = run {\n    |\n  }\n}", want: []string{"As", "let"}, not: []string{"on_hit", "components"}},
		{name: "item hook execute block", src: "CustomItem w {\n  base = stick\n  on_hit = run {\n    As(@a) {\n      /|\n    }\n  }\n}", want: []string{"say"}},
		{name: "item hook global", src: "Global g: int = 0\nCustomItem w {\n  base = stick\n  on_hit = run {\n    g|\n  }\n}", want: []string{"g"}, replace: "g"},
		{name: "item hook field still a field", src: "CustomItem w {\n  base = stick\n  on_hit = run {\n    /say hi\n  }\n  |\n}", want: []string{"on_right_click"}, not: []string{"As"}},

		// Selectors.
		{name: "selector option", src: "Function f {\n  As(@e[ty|]) { /say hi }\n}", want: []string{"type", "tag", "limit"}, replace: "ty"},
		{name: "selector next option", src: "Function f {\n  As(@e[type=zombie, |]) { /say hi }\n}", want: []string{"tag", "distance"}},
		{name: "selector entity type", src: "Function f {\n  As(@e[type=|]) { /say hi }\n}", want: []string{"zombie", "creeper"}},
		{name: "selector scores", src: "Objective kills\nFunction f {\n  As(@a[scores={ k| }]) { /say hi }\n}", want: []string{"kills"}, replace: "k"},
		{name: "command selector option", src: "Function f {\n  /kill @e[ty|\n}", want: []string{"type=", "tag="}, replace: "ty"},
		{name: "command selector type", src: "Function f {\n  /kill @e[tag=x,type=zom|\n}", want: []string{"zombie"}, replace: "zom"},
		{name: "command selector scores", src: "Objective kills\nFunction f {\n  /kill @e[scores={a=1,ki|\n}", want: []string{"kills"}, replace: "ki"},
		{name: "assignment target", src: "Global count: int = 0\nFunction f {\n  co| = 1\n}", want: []string{"count"}, replace: "co"},

		// No completion.
		{name: "comment", src: "// Feature |\nFeature f {}", none: true},
		{name: "string", src: "Function f {\n  let s = \"a|b\"\n}", none: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s, off := session(t, c.src)
			items, rep := s.Completion("main.mms", off)
			got := labels(items)
			if c.none {
				if len(got) != 0 {
					t.Errorf("want no items, got %d: %v", len(got), first(got))
				}
				return
			}
			if !has(got, c.want...) {
				t.Errorf("want %v among %d items: %v", c.want, len(got), first(got))
			}
			for _, n := range c.not {
				if has(got, n) {
					t.Errorf("%q offered here", n)
				}
			}
			src := []rune(strings.Replace(c.src, "|", "", 1))
			if r := string(src[rep.Start:rep.End]); r != c.replace {
				t.Errorf("replaces %q, want %q", r, c.replace)
			}
		})
	}
}

func first(xs []string) []string {
	if len(xs) > 12 {
		return xs[:12]
	}
	return xs
}
