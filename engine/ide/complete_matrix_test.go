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
// markupDecls are the names the ${...} rows complete, after the shape of
// a spellbook: a const list, a Storage record, a Global and an objective.
// bpHead is the start of a Blueprint the Blueprint rows complete in:
// two fields and a let.
const bpHead = "const tier_starts = [128, -128]\nBlueprint Tier {\n  tier: int = 1\n  label: string\n  let doubled = tier * 2\n"

const markupDecls = "Type S {\n  mode: int\n}\nconst modes: [string] = [\"a\", \"b\"]\nStorage st: S\nGlobal g: int = 0\nObjective kills: dummy\n"

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

		// Template arguments: a processor list's id goes where the whole
		// list does, and inside an inline list its processors go.
		{name: "template arg list items", src: "JigsawStructure v {\n  pools = {\n    o = [Piece(\"a\", processors = [|])]\n  }\n}", want: []string{"Rule", "BlockRot"}, not: []string{"projection", "bottom_rampart", "empty"}},
		{name: "template arg list items in pool", src: "JigsawStructure v {\n  pools = {\n    o = { elements = [Piece(\"a\", processors=[|])] }\n  }\n}", want: []string{"Rule"}, not: []string{"projection", "bottom_rampart"}},
		{name: "template arg list item prefix", src: "JigsawStructure v {\n  pools = {\n    o = [Piece(\"a\", processors = [Ru|])]\n  }\n}", want: []string{"Rule"}, replace: "Ru"},
		{name: "template arg list second item", src: "JigsawStructure v {\n  pools = {\n    o = [Piece(\"a\", processors = [Nop(), |])]\n  }\n}", want: []string{"Rule"}, not: []string{"projection"}},
		{name: "template arg nested call", src: "JigsawStructure v {\n  pools = {\n    o = [Piece(\"a\", processors = [BlockRot(|)])]\n  }\n}", want: []string{"integrity", "rottable_blocks"}, not: []string{"projection"}},
		{name: "template named arg value", src: "JigsawStructure v {\n  pools = {\n    o = [Piece(\"a\", processors = |)]\n  }\n}", want: []string{"bottom_rampart", "empty"}, not: []string{"projection", "weight"}},
		{name: "template args", src: "JigsawStructure v {\n  pools = {\n    o = [Piece(\"a\", |)]\n  }\n}", want: []string{"weight", "processors", "projection"}},
		{name: "inline processor list in pool", src: "TemplatePool p = { fallback = minecraft:empty, elements = [{ weight = 1, element = { element_type = minecraft:single_pool_element, location = minecraft:a, projection = rigid, processors = [|] } }] }", want: []string{"Rule", "BlockRot"}},

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
		{name: "if predicate", src: "Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })\nFunction f {\n  if (|)\n}", want: []string{"sneaking", "exists"}},
		{name: "if predicate operand", src: "Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })\nFunction f {\n  if (!sne|)\n}", want: []string{"sneaking"}, replace: "sne"},
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
		{name: "execute coordinate value", src: "Global gz: int = 0\nFunction f(x: int) {\n  Positioned(1, |)\n}", want: []string{"~ ~ ~", "x", "gz"}, not: []string{"true"}},
		{name: "execute coordinate offset", src: "Global gz: int = 0\nFunction f(dx: int) {\n  Positioned(~d|)\n}", want: []string{"dx"}, not: []string{"@a", "true", "exists"}, replace: "d"},
		{name: "execute local coordinate offset", src: "Function f(dx: int) {\n  Positioned(^ ^ ^-d|)\n}", want: []string{"dx"}, replace: "d"},
		{name: "execute selector no values", src: "Function f(x: int) {\n  As(|)\n}", want: []string{"@a"}, not: []string{"x"}},
		{name: "execute anchor", src: "Function f {\n  Anchored(|)\n}", want: []string{"eyes", "feet"}},
		{name: "execute dimension", src: "Function f {\n  In(|)\n}", want: []string{"the_nether"}},
		{name: "execute chain", src: "Function f {\n  As(@a).|\n}", want: []string{"At", "Positioned"}},
		{name: "hook body", src: "OnTick {\n  |\n}", want: []string{"As", "return"}},

		// Commands.
		{name: "command name", src: "Function f {\n  /scoreb|\n}", want: []string{"scoreboard"}, replace: "scoreb"},
		{name: "command criterion", src: "Function f {\n  /scoreboard objectives add x |\n}", want: []string{"dummy", "deathCount"}},
		{name: "command item", src: "Function f {\n  /give @s minecraft:dia|\n}", want: []string{"diamond"}, replace: "minecraft:dia"},
		{name: "command playsound", src: "Function f {\n  /playsound |\n}", want: []string{"entity.experience_orb.pickup", "ambient.cave"}},
		{name: "command playsound typed", src: "Function f {\n  /playsound minecraft:entity.pla|\n}", want: []string{"entity.player.levelup"}, replace: "minecraft:entity.pla"},
		{name: "command playsound in execute", src: "Function f {\n  /execute as @a run playsound |\n}", want: []string{"block.note_block.bell"}},
		{name: "command stopsound", src: "Function f {\n  /stopsound @a master |\n}", want: []string{"music.game"}},
		{name: "command stopsound any source", src: "Function f {\n  /stopsound @a * |\n}", want: []string{"music.game"}},
		{name: "command storage", src: "Type S { wave: int = 0 }\nStorage game: S\nFunction f {\n  /data get storage other:x\n  /data modify storage |\n}", want: []string{"test:mms", "other:x"}},
		{name: "command storage in execute", src: "Type S { wave: int = 0 }\nStorage game: S\nFunction f {\n  /execute store result storage |\n}", want: []string{"test:mms"}},
		{name: "command storage path", src: "Type S { wave: int = 0 }\nStorage game: S\nFunction f {\n  /data modify storage test:mms |\n}", want: []string{"game", "game.wave"}},
		{name: "command nested storage path", src: "Type P { x: int }\nType S { blah: { y: int }\n home: P }\nStorage game: S\nFunction f {\n  /data modify storage test:mms |\n}", want: []string{"game.blah", "game.blah.y", "game.home.x"}},
		{name: "nested storage field", src: "Type S {\n  blah: {\n    x: int\n    y: int\n  }\n}\nStorage s: S\nFunction f {\n  s.blah.|\n}", want: []string{"x", "y"}, not: []string{"blah"}},
		{name: "record field type", src: "Type S {\n  blah: {\n    x: in|\n  }\n}", want: []string{"int"}, replace: "in"},
		{name: "command in execute block", src: "Function f {\n  As(@a).At(@s) {\n    /|\n  }\n}", want: []string{"say"}},
		{name: "command value slash", src: "Function f {\n  let n = /|\n}", want: []string{"data", "execute", "time"}},
		{name: "command value slash assignment", src: "Global g: int = 0\nFunction f {\n  g = /|\n}", want: []string{"data"}},
		{name: "command value slash inline", src: "CustomItem w { base = stick, on_use_on_block = run {\n  let n = /|\n} }", want: []string{"data"}},
		{name: "division is not a command", src: "Function f {\n  let a = 4\n  let n = a /|\n}", not: []string{"data", "execute"}},
		{name: "command value name", src: "Function f {\n  let n = /da|\n}", want: []string{"data"}, replace: "da"},
		{name: "command value argument", src: "Function f {\n  let n = /data get |\n}", want: []string{"block", "entity", "storage"}},
		{name: "command value assignment", src: "Global g: int = 0\nFunction f {\n  g = /time query |\n}", want: []string{"gametime"}},
		{name: "command value interpolation", src: "Function f(k: int) {\n  let n = /data get storage demo:x ${|}\n}", want: []string{"k"}},
		{name: "command value inline", src: "CustomItem w { base = stick, on_use_on_block = run {\n  let ok: bool = /exec|\n} }", want: []string{"execute"}, replace: "exec"},

		{name: "interpolation", src: "Function f(n: int) {\n  let x = 1\n  /say ${|}\n}", want: []string{"n", "x"}, not: []string{"say", "As"}},
		{name: "interpolation prefix", src: "Global g: int = 0\nFunction f {\n  /say ${g|}\n}", want: []string{"g"}, replace: "g"},
		{name: "interpolation custom item", src: "CustomItem shard { base = stick }\nFunction f {\n  /give @s ${|}\n}", want: []string{"shard"}},
		{name: "interpolation storage field", src: "Type S { mode: int }\nStorage st: S\nGlobal g: int = 0\nFunction f {\n  /say ${st.|}\n}", want: []string{"mode"}, not: []string{"g", "st"}},
		{name: "interpolation storage field in index", src: "const modes: [string] = [\"a\", \"b\"]\nType S { mode: int }\nStorage st: S\nFunction f {\n  /say ${modes[st.mo|]}\n}", want: []string{"mode"}, not: []string{"modes"}, replace: "mo"},
		{name: "interpolation const", src: "const modes: [string] = [\"a\", \"b\"]\nFunction f {\n  /say ${mod|}\n}", want: []string{"modes"}, replace: "mod"},

		// ${...} in text markup, as in a spellbook item: the same names
		// as code there, and Storage fields after a dot.
		{name: "markup interpolation", src: markupDecls + "Function f {\n  let x = 1\n  actionbar(@s, \"Mode ${|}\")\n}", want: []string{"x", "g", "modes", "st", "kills"}},
		{name: "markup interpolation prefix", src: markupDecls + "Function f {\n  actionbar(@s, \"Mode ${mod|}\")\n}", want: []string{"modes"}, replace: "mod"},
		{name: "markup interpolation index", src: markupDecls + "Function f {\n  actionbar(@s, \"Mode ${modes[st|]}\")\n}", want: []string{"st"}, replace: "st"},
		{name: "markup interpolation index empty", src: markupDecls + "Function f {\n  let x = 1\n  actionbar(@s, \"Mode ${modes[|]}\")\n}", want: []string{"st", "x", "g"}},
		{name: "markup interpolation storage field", src: markupDecls + "Function f {\n  actionbar(@s, \"Mode ${modes[st.|]}\")\n}", want: []string{"mode"}, not: []string{"g", "st", "modes"}},
		{name: "markup interpolation storage field typed", src: markupDecls + "Function f {\n  actionbar(@s, \"${st.mo|de}\")\n}", want: []string{"mode"}, replace: "mode"},
		{name: "markup interpolation after call", src: markupDecls + "Function f {\n  let x = 1\n  actionbar(@s, \"${len(modes) - |}\")\n}", want: []string{"x", "st"}},
		{name: "markup interpolation item hook", src: markupDecls + "CustomItem w {\n  base = stick\n  on_right_click = run {\n    let x = 1\n    actionbar(@s, \"Mode ${modes[st.|]}\")\n  }\n}", want: []string{"mode"}, not: []string{"x"}},
		{name: "markup interpolation item hook names", src: markupDecls + "CustomItem w {\n  base = stick\n  on_right_click = run {\n    let x = 1\n    actionbar(@s, \"Mode ${|}\")\n  }\n}", want: []string{"x", "modes", "st", "g"}},
		{name: "markup interpolation outside code", src: markupDecls + "Function fn {\n}\nAdvancement a { display = { title = \"${|}\", description = \"d\", icon = { id = stone } }, criteria = { x = Tick() } }", want: []string{"g", "modes", "st", "kills"}, not: []string{"fn", "true"}},
		{name: "markup interpolation outside code field", src: markupDecls + "Advancement a { display = { title = \"${st.|}\", description = \"d\", icon = { id = stone } }, criteria = { x = Tick() } }", want: []string{"mode"}, not: []string{"g"}},
		{name: "markup after interpolation", src: markupDecls + "Function f {\n  let x = 1\n  actionbar(@s, \"${x} th|\")\n}", none: true},
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
		{name: "const selector option", src: "const boss = @e[ty|]", want: []string{"type", "tag", "limit"}, replace: "ty"},
		{name: "typed const selector", src: "const boss: Selector = |", want: []string{"@a", "@e", "@s"}},
		{name: "selector const as target", src: "const boss = @e[tag=b]\nFunction f {\n  tell(b|, \"hi\")\n}", want: []string{"boss"}, replace: "b"},
		{name: "selector const in command", src: "const boss = @e[tag=b]\nFunction f {\n  /kill ${b|}\n}", want: []string{"boss"}, replace: "b"},
		{name: "extended selector option", src: "const guards = @e[tag=guard]\nconst awake = guards[ta|]", want: []string{"tag", "limit"}, replace: "ta"},
		{name: "extended template selector", src: "Template near(r: int): Selector = @a[distance=..${r}]\nFunction f {\n  tell(near(3)[li|], \"hi\")\n}", want: []string{"limit"}, replace: "li"},
		{name: "extended selector in function", src: "const guards = @e[tag=guard]\nFunction f {\n  tell(guards[li|], \"hi\")\n}", want: []string{"limit"}, replace: "li"},
		{name: "literal selector in tell", src: "Function f {\n  tell(@e[li|], \"hi\")\n}", want: []string{"limit"}, replace: "li"},
		{name: "objective score not a selector", src: "Objective hp\nFunction f {\n  hp[|] = 1\n}", not: []string{"limit"}},
		{name: "assignment target", src: "Global count: int = 0\nFunction f {\n  co| = 1\n}", want: []string{"count"}, replace: "co"},

		// A Blueprint's own body: its keywords, the lines of its check
		// block, and the values its lets, checks and emit block use.
		{name: "blueprint item", src: bpHead + "  |\n}", want: []string{"let", "check", "emit", "as"}, not: []string{"tier"}},
		{name: "blueprint item prefix", src: bpHead + "  ch|\n}", want: []string{"check"}, replace: "ch"},
		{name: "blueprint item after section", src: bpHead + "  check {\n  }\n  em|\n}", want: []string{"emit"}, replace: "em"},
		{name: "blueprint field name", src: "Blueprint B {\n  ti|\n}", want: []string{"let", "check"}, replace: "ti"},
		{name: "blueprint as type", src: bpHead + "  as |\n}", want: []string{"ItemStack", "int"}},
		{name: "blueprint let value", src: bpHead + "  let x = |\n}", want: []string{"tier", "label", "doubled", "self", "tier_starts", "len"}, not: []string{"x", "As"}},
		{name: "blueprint let value prefix", src: bpHead + "  let x = ti|\n}", want: []string{"tier", "tier_starts"}, replace: "ti"},
		{name: "blueprint let index", src: bpHead + "  let x = tier_starts[|]\n}", want: []string{"tier", "doubled"}},
		{name: "blueprint let operand", src: bpHead + "  let x = tier - |\n}", want: []string{"tier", "doubled"}},
		{name: "blueprint later let", src: bpHead + "  let x = |\n  let y = 1\n}", want: []string{"doubled"}, not: []string{"x", "y"}},
		{name: "blueprint self", src: bpHead + "  let x = self.|\n}", want: []string{"id", "namespace", "path"}, not: []string{"tier"}},
		{name: "blueprint typed let", src: bpHead + "  let g: DensityFunction = {\n    |\n  }\n}", want: []string{"type"}},
		{name: "blueprint typed let field", src: bpHead + "  let g: DensityFunction = {\n    type = gradient\n    from_coordinate = tier_starts[|]\n  }\n}", want: []string{"tier"}},
		{name: "blueprint string interpolation", src: bpHead + "  let key = \"item.${|}\"\n}", want: []string{"tier", "label", "self"}, not: []string{"key"}},
		{name: "blueprint string self", src: bpHead + "  let key = \"item.${self.na|}\"\n}", want: []string{"namespace"}, replace: "na"},
		{name: "check line", src: bpHead + "  check {\n    |\n  }\n}", want: []string{"error", "warn", "if", "tier", "label"}},
		{name: "check line in if", src: bpHead + "  check {\n    if (tier < 1) {\n      |\n    }\n  }\n}", want: []string{"error", "warn", "tier"}},
		{name: "check condition", src: bpHead + "  check {\n    if (|) {\n    }\n  }\n}", want: []string{"tier", "label", "self", "contains"}, not: []string{"error"}},
		{name: "check condition operand", src: bpHead + "  check {\n    if (tier < |) {\n    }\n  }\n}", want: []string{"doubled"}},
		{name: "check error arg", src: bpHead + "  check {\n    if (tier < 1) {\n      error(|)\n    }\n  }\n}", want: []string{"tier", "label"}},
		{name: "check error message", src: bpHead + "  check {\n    error(tier, \"tier ${|} is too low\")\n  }\n}", want: []string{"tier", "doubled"}},
		{name: "check for", src: bpHead + "  check {\n    for t in |\n  }\n}", want: []string{"tier_starts", "tier"}},
		{name: "emit statement", src: bpHead + "  emit {\n    |\n  }\n}", want: []string{"Objective", "Function", "primary", "if", "for", "Tier"}, not: []string{"Template", "error"}},
		{name: "emit primary kind", src: bpHead + "  emit {\n    primary |\n  }\n}", want: []string{"DensityFunction", "CustomItem"}, not: []string{"primary", "if"}},
		{name: "emit primary kind prefix", src: bpHead + "  emit {\n    primary Dens|\n  }\n}", want: []string{"DensityFunction"}, replace: "Dens"},
		{name: "emit name interpolation", src: bpHead + "  emit {\n    Objective ${|}: dummy\n  }\n}", want: []string{"self", "label", "tier"}},
		{name: "emit name self", src: bpHead + "  emit {\n    primary Objective ${self.|}: dummy\n  }\n}", want: []string{"path", "id"}, not: []string{"tier"}},
		{name: "emit name self after text", src: bpHead + "  emit {\n    Objective ${label}_${self.pa|}: dummy\n  }\n}", want: []string{"path"}, replace: "pa"},
		{name: "emit if condition", src: bpHead + "  emit {\n    if (|) {\n    }\n  }\n}", want: []string{"tier", "label"}},
		{name: "emit for values", src: bpHead + "  emit {\n    for t in |\n  }\n}", want: []string{"tier_starts"}},
		{name: "emit const value", src: bpHead + "  emit {\n    const id = |\n  }\n}", want: []string{"self", "label"}},
		{name: "emit body field", src: bpHead + "  emit {\n    primary DensityFunction ${self.path} = {\n      |\n    }\n  }\n}", want: []string{"type"}},
		{name: "emit body value", src: bpHead + "  emit {\n    primary Objective ${self.path}: dummy\n    CustomItem ${self.path}_item {\n      base = stick\n      name = |\n    }\n  }\n}", want: []string{"label"}},
		{name: "emit body string", src: bpHead + "  emit {\n    CustomItem ${self.path} {\n      base = stick\n      name = \"${|}\"\n    }\n  }\n}", want: []string{"label", "self"}},
		{name: "emit function code", src: bpHead + "  emit {\n    Function ${self.path} {\n      let y = |\n    }\n  }\n}", want: []string{"tier", "label", "self", "true"}},
		{name: "emit function command", src: bpHead + "  emit {\n    Function ${self.path} {\n      /say ${|}\n    }\n  }\n}", want: []string{"label", "self"}},
		{name: "emit function command self", src: bpHead + "  emit {\n    Function ${self.path} {\n      /say ${self.|}\n    }\n  }\n}", want: []string{"path"}},
		{name: "blueprint instance field", src: bpHead + "}\nTier t {\n  |\n}", want: []string{"tier", "label"}, not: []string{"doubled"}},

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
