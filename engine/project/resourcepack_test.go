package project

import (
	"strings"
	"testing"
	"testing/fstest"
)

const resourcePackSrc = `BlockState minecraft:glowstone = {
  variants = {
    "" = [
      { model = block/glowstone_1, weight = 3 },
      { model = block/glowstone_2, y = 90 },
    ]
  }
}
BlockState minecraft:oak_fence = {
  multipart = [
    { apply = { model = minecraft:block/oak_fence_post } },
    { when = { OR = [{ north = "true" }, { east = "true|false" }] }, apply = { model = minecraft:block/oak_fence_side, uvlock = true } },
  ]
}
Model block/glowstone_1 = { parent = minecraft:block/cube_all, textures = { all = Texture("art/glow1.png", "block/glowstone_1") } }
Model block/glowstone_2 = { parent = minecraft:block/cube_all, textures = { all = Texture("art/glow2.png", "block/glowstone_2"), particle = "#all" } }
ItemModel banner = { model = Model(model = block/glowstone_1) }
Lang en_us = { "item.demo.banner" = "Siege Banner" }
`

func TestResourcePackKinds(t *testing.T) {
	files := fstest.MapFS{"art/glow1.png": {Data: pngBytes}, "art/glow2.png": {Data: pngBytes}}
	_, r := assetProject(t, files, resourcePackSrc)
	if d := diagText(r); d != "" {
		t.Fatalf("unexpected diagnostics:\n%s", d)
	}
	data, res := SplitPacks(r.Emit())
	var got []string
	for _, f := range data {
		got = append(got, "data "+f.Path)
	}
	for _, f := range res {
		got = append(got, "res "+f.Path)
	}
	want := `data pack.mcmeta
data pack.png
res assets/demo/items/banner.json
res assets/demo/lang/en_us.json
res assets/demo/models/block/glowstone_1.json
res assets/demo/models/block/glowstone_2.json
res assets/demo/textures/block/glowstone_1.png
res assets/demo/textures/block/glowstone_2.png
res assets/minecraft/blockstates/glowstone.json
res assets/minecraft/blockstates/oak_fence.json
res pack.mcmeta
res pack.png`
	if g := strings.Join(got, "\n"); g != want {
		t.Errorf("files:\n%s\nwant:\n%s", g, want)
	}
	for _, f := range res {
		switch f.Path {
		case "assets/minecraft/blockstates/glowstone.json":
			s := string(f.Data)
			if !strings.Contains(s, `"model": "demo:block/glowstone_1"`) || !strings.Contains(s, `"weight": 3`) {
				t.Errorf("glowstone.json:\n%s", s)
			}
		case "assets/demo/models/block/glowstone_2.json":
			s := string(f.Data)
			if !strings.Contains(s, `"all": "demo:block/glowstone_2"`) || !strings.Contains(s, `"particle": "#all"`) {
				t.Errorf("glowstone_2.json:\n%s", s)
			}
		case "pack.mcmeta":
			if !strings.Contains(string(f.Data), "97") {
				t.Errorf("resource pack.mcmeta:\n%s", f.Data)
			}
		}
	}
}

// A project with only declarations, no copied files, still gets a
// resource pack.
func TestResourcePackWithoutAssets(t *testing.T) {
	_, r := assetProject(t, nil, `Lang en_us = { "a" = "b" }`)
	_, res := SplitPacks(r.Emit())
	var paths []string
	for _, f := range res {
		paths = append(paths, f.Path)
	}
	if strings.Join(paths, " ") != "assets/demo/lang/en_us.json pack.mcmeta pack.png" {
		t.Errorf("resource pack = %v", paths)
	}
}

func TestBlockStateChecks(t *testing.T) {
	cases := []struct{ src, want string }{
		{`BlockState minecraft:furnace = { variants = { "facing=up" = { model = minecraft:block/furnace } } }`,
			"mms/block-state: 'up' is not a value of minecraft:furnace's 'facing'; it is one of north, south, west, east"},
		{`BlockState minecraft:furnace = { variants = { "lt=true" = { model = minecraft:block/furnace } } }`,
			"mms/block-state: 'minecraft:furnace' has no property 'lt'; did you mean 'lit'?"},
		{`BlockState minecraft:furnace = { variants = { "facing" = { model = minecraft:block/furnace } } }`,
			`mms/block-state: a variant key is "" or property=value pairs separated by ',', such as "facing=north,lit=true"`},
		{`BlockState minecraft:furnace = { variants = { "lit=true,lit=false" = { model = minecraft:block/furnace } } }`,
			"mms/block-state: property 'lit' is set twice"},
		{`BlockState minecraft:glowstone = { variants = { "lit=true" = { model = minecraft:block/glowstone } } }`,
			`mms/block-state: minecraft:glowstone has no state properties, so its only variant is ""`},
		{`BlockState minecraft:oak_fence = { multipart = [{ when = { AND = [{ nort = "true" }] }, apply = { model = minecraft:block/oak_fence_post } }] }`,
			"mms/block-state: 'minecraft:oak_fence' has no property 'nort'; did you mean 'north'?"},
		{`BlockState minecraft:oak_fence = { multipart = [{ when = { north = "true|maybe" }, apply = { model = minecraft:block/oak_fence_post } }] }`,
			"mms/block-state: 'maybe' is not a value of minecraft:oak_fence's 'north'; it is one of true, false"},
		{`BlockState minecraft:glowstne = { variants = { "" = { model = minecraft:block/glowstone } } }`,
			"mms/unknown-block: the game has no block 'minecraft:glowstne', so it never reads this block state; did you mean 'glowstone'?"},
		{`BlockState glowstone = { variants = { "" = { model = minecraft:block/glowstone } } }`,
			"mms/unknown-block: the game only reads block states of its own blocks, and 'demo:glowstone' is not one; name it after a block, such as minecraft:glowstone"},
		{`BlockState minecraft:glowstone = { variants = { "" = { model = minecraft:block/glowstne } } }`,
			"mms/unknown-id: unknown model 'minecraft:block/glowstne'; did you mean 'block/glowstone'?"},
		{`Model block/x = { parent = minecraft:block/cube_all, textures = { all = demo:block/x } }`,
			"lint/external-asset: this project does not provide the texture 'demo:block/x'; it must come from another pack, or include the file with Texture(\"path\")"},
		{`ItemModel x = { model = Model(model = demo:block/missing) }`,
			"lint/external-asset: this project does not provide the model 'demo:block/missing'; it must come from another pack, or be declared here: Model block/missing = { ... }"},
	}
	for _, tc := range cases {
		_, r := assetProject(t, nil, tc.src)
		if got := strings.TrimSpace(diagText(r)); got != tc.want {
			t.Errorf("%s\ngot:  %s\nwant: %s", tc.src, got, tc.want)
		}
	}
}
