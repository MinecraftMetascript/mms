package project

import (
	"context"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

const customItemSrc = `CustomItem ruby = {
  base = emerald
  name = { en_us = "Ruby", de_de = "Rubin" }
  texture = Texture("art/ruby.png")
  components = { rarity = rare }
}
CustomItem ruby_sword = {
  base = iron_sword
  name = "Ruby Sword"
  texture = Texture("art/ruby_sword.png")
  parent = item/handheld
}
CustomItem badge = { base = paper, name = "Badge" }
Lang en_us = { "item.demo.other" = "Other" }
Recipe ruby = Shapeless([emerald, redstone], ruby * 2)
Recipe ruby_sword = Shapeless([iron_sword, emerald], { id = ruby_sword, components = { rarity = epic } })
LootTable drops { pools = [Pool(1, [ruby, Item(name = badge, modifier = [SetCount(3)])])] }
Function give {
  /give @s ${ruby} 4
  /give @s ${demo:badge}
}
`

func TestCustomItem(t *testing.T) {
	files := fstest.MapFS{"art/ruby.png": {Data: pngBytes}, "art/ruby_sword.png": {Data: pngBytes}}
	_, r := assetProject(t, files, customItemSrc)
	if d := diagText(r); d != "" {
		t.Fatalf("unexpected diagnostics:\n%s", d)
	}
	got := map[string]string{}
	for _, f := range r.Emit() {
		key := f.Path
		if f.Pack == 1 {
			key = "res " + key
		}
		got[key] = string(f.Data)
	}
	want := map[string][]string{
		"res assets/demo/models/item/ruby.json":       {`"parent": "minecraft:item/generated"`, `"layer0": "demo:item/ruby"`},
		"res assets/demo/models/item/ruby_sword.json": {`"parent": "minecraft:item/handheld"`, `"layer0": "demo:item/ruby_sword"`},
		"res assets/demo/items/ruby.json":             {`"type": "minecraft:model"`, `"model": "demo:item/ruby"`},
		"res assets/demo/lang/en_us.json":             {`"item.demo.other": "Other"`, `"item.demo.ruby": "Ruby"`, `"item.demo.ruby_sword": "Ruby Sword"`, `"item.demo.badge": "Badge"`},
		"res assets/demo/lang/de_de.json":             {`"item.demo.ruby": "Rubin"`},
		"res assets/demo/textures/item/ruby.png":      {},
		"data/demo/recipe/ruby.json":                  {`"id": "minecraft:emerald"`, `"count": 2`, `"item_model": "demo:ruby"`, `"translate": "item.demo.ruby"`, `"fallback": "Ruby"`, `"rarity": "rare"`},
		"data/demo/recipe/ruby_sword.json":            {`"id": "minecraft:iron_sword"`, `"rarity": "epic"`, `"item_model": "demo:ruby_sword"`},
		"data/demo/loot_table/drops.json":             {`"name": "minecraft:emerald"`, `"name": "minecraft:paper"`, `"type": "minecraft:set_components"`, `"type": "minecraft:set_count"`},
		"data/demo/function/give.mcfunction": {
			`give @s emerald[item_model="demo:ruby",item_name={translate:"item.demo.ruby",fallback:"Ruby"},rarity="rare"] 4`,
			`give @s paper[item_name={translate:"item.demo.badge",fallback:"Badge"}]`,
		},
	}
	for path, parts := range want {
		s, ok := got[path]
		if !ok {
			t.Errorf("no %s", path)
			continue
		}
		for _, p := range parts {
			if !strings.Contains(s, p) {
				t.Errorf("%s lacks %s:\n%s", path, p, s)
			}
		}
	}
	// The badge keeps paper's look: no models, no item_model component.
	for _, p := range []string{"res assets/demo/items/badge.json", "res assets/demo/models/item/badge.json"} {
		if _, ok := got[p]; ok {
			t.Errorf("unexpected %s", p)
		}
	}
	if strings.Contains(got["data/demo/function/give.mcfunction"], "demo:badge") {
		t.Errorf("badge has an item_model:\n%s", got["data/demo/function/give.mcfunction"])
	}
	// The custom item's components come first in a loot entry, so later
	// modifiers can change them.
	drops := got["data/demo/loot_table/drops.json"]
	if strings.Index(drops, "set_components") > strings.Index(drops, "set_count") {
		t.Errorf("set_components should come before set_count:\n%s", drops)
	}
	// The sword's rarity at the use wins over nothing, and the custom
	// components stay.
	if s := got["data/demo/recipe/ruby_sword.json"]; strings.Count(s, "rarity") != 1 {
		t.Errorf("ruby_sword.json:\n%s", s)
	}
}

func TestCustomItemErrors(t *testing.T) {
	files := fstest.MapFS{"art/a.png": {Data: pngBytes}}
	cases := []struct{ src, want string }{
		{`CustomItem a = { name = "A" }`, "missing"},
		{`CustomItem a = { base = stick, texture = Texture("art/a.png"), model = item/stick }`, "model names the model to show"},
		{`CustomItem a = { base = stick, model = item/stick, item_model = stick }`, "item_model replaces the generated models"},
		{`CustomItem a = { base = stick, parent = item/handheld }`, "add a texture"},
		{`CustomItem a = { base = stick, name = { english = "A" } }`, "not a language code"},
		{"CustomItem a = { base = stick }\nItemTag t = [a]", "is a CustomItem"},
		{"CustomItem a = { base = stick }\nRecipe r = Shapeless([a], stick)", "is a CustomItem"},
		{"CustomItem a = { base = stick, texture = Texture(\"art/a.png\") }\nItemModel a = { model = Model(model = item/stick) }", "already made by CustomItem"},
		{"ItemModel a = { model = Model(model = item/stick) }\nCustomItem a = { base = stick, texture = Texture(\"art/a.png\") }", "makes ItemModel 'demo:a', which is already declared"},
		{"CustomItem a = { base = stick, texture = Texture(\"art/a.png\") }\nModel item/a = { parent = item/generated }", "already made by CustomItem"},
		{"CustomItem a = { base = stick, name = \"A\" }\nLang en_us = { \"item.demo.a\" = \"B\" }", "already sets 'item.demo.a'"},
		{"CustomItem a = { base = stick, components = { bundle_contents = [a] } }\nRecipe r = Shapeless([stick], a)", "contains itself"},
	}
	for _, tc := range cases {
		_, r := assetProject(t, files, tc.src)
		if d := diagText(r); !strings.Contains(d, tc.want) {
			t.Errorf("%s\nwant %q in:\n%s", tc.src, tc.want, d)
		}
	}
}

// A declaration that uses a custom item is checked again when the item
// changes, though the set of declared names stays the same.
func TestCustomItemRecheck(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("items.mms", `CustomItem a = { base = stick, name = "A" }`)
	p.SetFile("use.mms", "Function give {\n  /give @s ${a}\n}\n")
	if _, err := p.Analyze(context.Background()); err != nil {
		t.Fatal(err)
	}
	p.SetFile("items.mms", `CustomItem a = { base = stick, name = "B" }`)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range r.Emit() {
		if strings.HasSuffix(f.Path, "give.mcfunction") && !strings.Contains(string(f.Data), `fallback:"B"`) {
			t.Errorf("give was not rechecked:\n%s", f.Data)
		}
	}
}

// A use checked before its custom item is declared is checked again once
// it is, as when the web editor opens a project file by file.
func TestCustomItemDeclaredLater(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("use.mms", "Function give {\n  /give @s ${a}\n}\nRecipe r = Shapeless([stick], a)\n")
	if _, err := p.Analyze(context.Background()); err != nil {
		t.Fatal(err)
	}
	p.SetFile("items.mms", `CustomItem a = { base = stick, name = "A" }`)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"use.mms", "items.mms"} {
		for _, d := range r.Diagnostics(path) {
			t.Errorf("%s: %s: %s", path, d.Code, d.Message)
		}
	}
}

// A use of a custom item that holds another one is checked again when
// the inner item changes.
func TestCustomItemNestedRecheck(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("inner.mms", `CustomItem b = { base = paper, name = "B" }`)
	p.SetFile("outer.mms", `CustomItem a = { base = bundle, components = { bundle_contents = [b] } }`)
	p.SetFile("use.mms", "Function give {\n  /give @s ${a}\n}\n")
	if _, err := p.Analyze(context.Background()); err != nil {
		t.Fatal(err)
	}
	p.SetFile("inner.mms", `CustomItem b = { base = paper, name = "C" }`)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range r.Emit() {
		if strings.HasSuffix(f.Path, "give.mcfunction") && !strings.Contains(string(f.Data), `fallback:"C"`) {
			t.Errorf("give was not rechecked:\n%s", f.Data)
		}
	}
}

// A newline in a name stays on the command's line.
func TestCustomItemNameEscapes(t *testing.T) {
	_, r := assetProject(t, fstest.MapFS{}, "CustomItem a = { base = paper, name = \"A\\nB\" }\nFunction give {\n  /give @s ${a}\n}\n")
	for _, f := range r.Emit() {
		if strings.HasSuffix(f.Path, "give.mcfunction") {
			if got := strings.TrimSpace(string(f.Data)); strings.Contains(got, "\n") || !strings.Contains(got, `fallback:"A\nB"`) {
				t.Errorf("give.mcfunction:\n%s", got)
			}
		}
	}
}

// TestOnRightClick checks the right-click hook: the item's consumable,
// use_effects and marker, the using_item advancement, the once-per-press
// reward function and the objective's setup.
func TestOnRightClick(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("main.mms", `
CustomItem wand = {
  base = stick
  name = "Wand"
  components = { custom_data = { kind = "wand" } }
}
OnRightClick wand {
  /say zap
}
Function give {
  /give @s ${wand}
}
`)
	r, _ := p.Analyze(context.Background())
	if d := diagText(r); d != "" {
		t.Fatalf("unexpected diagnostics:\n%s", d)
	}
	got := dumpFiles(r.Emit())
	for _, want := range []string{
		`give @s stick[item_name={translate:"item.demo.wand",fallback:"Wand"},custom_data={kind:"wand",mms_item:"demo:wand"},consumable={consume_seconds:1000000,animation:"none",has_consume_particles:false},use_effects={can_sprint:true,speed_multiplier:1,interact_vibrations:false}]`,
		"=== data/demo/advancement/mms_gen/right_click/wand.json\n",
		`"trigger": "minecraft:using_item"`,
		`"items": "minecraft:stick"`,
		`"mms_item": "demo:wand"`,
		`"function": "demo:mms_gen/right_click/wand"`,
		"=== data/demo/function/mms_gen/right_click/wand.mcfunction\nadvancement revoke @s only demo:mms_gen/right_click/wand\nexecute if score @s mms_click.demo matches 1.. run return run scoreboard players set @s mms_click.demo 2\nscoreboard players set @s mms_click.demo 2\nsay zap\n",
		"scoreboard objectives add mms_click.demo dummy\n",
		"scoreboard players remove @a[scores={mms_click.demo=1..}] mms_click.demo 1\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}

	errs := map[string]string{
		"OnRightClick stick {\n  /say hi\n}":                                                                        "mms/unknown-id",
		"Function f {\n}\nOnRightClick f {\n  /say hi\n}":                                                           "mms/unknown-id",
		"CustomItem w = { base = stick }\nOnRightClick w":                                                           "mms/missing-value",
		"CustomItem w = { base = stick, components = { custom_data = \"{a:1}\" } }\nOnRightClick w {\n  /say hi\n}": "mms/custom-item",
		"CustomItem w = { base = stick }\nOnRightClick w {\n  /say a\n}\nOnRightClick w {\n  /say b\n}":             "mms/duplicate",
	}
	for src, want := range errs {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("main.mms", src)
		r, _ := p.Analyze(context.Background())
		var codes []string
		for _, d := range r.Diagnostics("main.mms") {
			if d.Severity == syntax.SeverityError {
				codes = append(codes, d.Code)
			}
		}
		if len(codes) == 0 || codes[0] != want {
			t.Errorf("%q: errors %v, want %s", src, codes, want)
		}
	}
}

// TestOnRightClickIncremental checks that adding and removing an
// OnRightClick in another file changes the uses of the item.
func TestOnRightClickIncremental(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("item.mms", "CustomItem wand = { base = stick }\n")
	p.SetFile("give.mms", "Function give {\n  /give @s ${wand}\n}\n")
	give := func() string {
		r, _ := p.Analyze(context.Background())
		for _, f := range r.Emit() {
			if f.Path == "data/demo/function/give.mcfunction" {
				return string(f.Data)
			}
		}
		t.Fatal("no give function")
		return ""
	}
	if g := give(); strings.Contains(g, "consumable") {
		t.Fatalf("consumable before the hook: %s", g)
	}
	p.SetFile("click.mms", "OnRightClick wand {\n  /say zap\n}\n")
	if g := give(); !strings.Contains(g, "consumable") || !strings.Contains(g, `mms_item:"demo:wand"`) {
		t.Fatalf("no consumable after adding the hook: %s", g)
	}
	p.SetFile("click.mms", "")
	if g := give(); strings.Contains(g, "consumable") {
		t.Fatalf("consumable after removing the hook: %s", g)
	}
}
