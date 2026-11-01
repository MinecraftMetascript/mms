package project

import (
	"context"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/spec"
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
