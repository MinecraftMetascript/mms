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
		{name: "kind after declaration", src: "Feature f = {}\nFe|", want: []string{"Feature", "PlacedFeature"}, replace: "Fe"},
		{name: "kind after function", src: "Function f {\n  /say hi\n}\nRec|", want: []string{"Recipe"}, replace: "Rec"},
		{name: "kind after namespace", src: "Namespace foo\nBio|", want: []string{"Biome"}, replace: "Bio"},
		{name: "kind before declaration", src: "Fe|\nFunction f {\n}", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind after word value", src: "DensityFunction x = foo\nFe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind after criterion", src: "Objective kills = deathCount\nFe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "kind mid-word", src: "Fe|ature f = {}", want: []string{"Feature"}, replace: "Feature"},
		{name: "kind after comment", src: "// ores\nFe|", want: []string{"Feature"}, replace: "Fe"},
		{name: "declaration name", src: "Feature |", none: true},
		{name: "declaration name typed", src: "Feature f|", none: true},
		{name: "namespace name", src: "Namespace |", none: true},

		// After `=`: the kind's values.
		{name: "feature value", src: "Feature f = |", want: []string{"Ore", "RandomNeighborSpread"}},
		{name: "recipe value", src: "Recipe r = |", want: []string{"Shaped", "Smelting"}},
		{name: "surface rule value", src: "SurfaceRule r = |", want: []string{"If", "Sequence"}},
		{name: "objective criterion", src: "Objective kills = |", want: []string{"dummy", "deathCount", "trigger"}},
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

		// Objects and nested maps.
		{name: "object fields", src: "Biome b = {\n  |\n}", want: []string{"effects", "temperature"}},
		{name: "nested object", src: "Biome b = {\n  effects = {\n    |\n  }\n}", want: []string{"water_color"}, not: []string{"temperature"}},
		{name: "object in helper", src: "Recipe r = Smelting(result = { |})", want: []string{"id", "count"}},
		{name: "object in list", src: "Feature f = Ore(targets = [{ |}])", want: []string{"target", "state"}},
		{name: "dispatched on sibling", src: "Advancement a = {\n  display = {\n    |\n  }\n}", want: []string{"icon", "title", "frame", "background"}, not: []string{"criteria"}},
		{name: "dispatched on sibling set", src: "Advancement a = {\n  parent = minecraft:story/root\n  display = {\n    |\n  }\n}", want: []string{"icon", "title"}, not: []string{"criteria", "background"}},
		{name: "dispatched value", src: "Advancement a = { display = { frame = |} }", want: []string{"challenge", "goal", "task"}},
		{name: "map keys", src: "Enchantment e = {\n  effects = {\n    |\n  }\n}", want: []string{"damage", "attributes"}, not: []string{"max_level"}},
		{name: "map keys with slash", src: "Biome b = {\n  attributes = {\n    vis|\n  }\n}", want: []string{"visual/sky_color"}, not: []string{"temperature"}, replace: "vis"},
		{name: "item components", src: "Recipe r = Shapeless([emerald], { id = emerald\n components = {\n |\n }\n})", want: []string{"item_name", "max_stack_size"}, not: []string{"count"}},
		{name: "map value", src: "Enchantment e = {\n  effects = {\n    minecraft:damage = [|]\n  }\n}", want: []string{"DamageEnchantmentEffect"}},
		{name: "function reference", src: "Function g {}\nAdvancement a = { rewards = { function = |} }", want: []string{"g"}},

		// Tags.
		{name: "tag entries", src: "ItemTag gems = [|]", want: []string{"diamond", "#logs", "Optional"}, not: []string{"#gems"}},
		{name: "tag entry prefix", src: "ItemTag gems = [dia|]", want: []string{"diamond"}, replace: "dia"},
		{name: "tag entry tag", src: "BlockTag t = [#|]", want: []string{"#logs"}, replace: "#"},
		{name: "tag optional", src: "ItemTag t = [Optional(|)]", want: []string{"diamond"}},
		{name: "own tag reference", src: "ItemTag gems = [diamond]\nItemTag more = [|]", want: []string{"#gems"}},

		// Function bodies.
		{name: "statement", src: "Function f {\n  /say hi\n  |\n}", want: []string{"As", "let", "if"}},
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
		{name: "objective in expression", src: "Objective o = dummy\nFunction f {\n  let y = |\n}", want: []string{"o"}},
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

		// No completion.
		{name: "comment", src: "// Feature |\nFeature f = {}", none: true},
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
