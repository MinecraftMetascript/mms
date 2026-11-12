package project

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

func retextureFiles() fstest.MapFS {
	return fstest.MapFS{"art/icon.png": {Data: pngBytes}, "art/side.png": {Data: pngOf(16, 16)}, "art/end.png": {Data: pngOf(16, 16)}}
}

func emittedFiles(r *Result) map[string]string {
	out := map[string]string{}
	for _, f := range r.Emit() {
		out[f.Path] = string(f.Data)
	}
	return out
}

func TestRetexture(t *testing.T) {
	src := `Retexture minecraft:oak_log { textures = { side = Texture("art/side.png"), end = Texture("art/end.png") } }
Retexture minecraft:oak_door { textures = { top = Texture("art/side.png"), bottom = Texture("art/end.png") } }
`
	_, r := assetProject(t, retextureFiles(), src)
	if d := diagText(r); d != "" {
		t.Fatalf("diagnostics:\n%s", d)
	}
	out := emittedFiles(r)
	var paths []string
	for p := range out {
		if strings.HasPrefix(p, "assets/") {
			paths = append(paths, p)
		}
	}
	sort.Strings(paths)
	want := []string{
		"assets/demo/models/block/oak_door/oak_door_bottom_left.json",
		"assets/demo/models/block/oak_door/oak_door_bottom_left_open.json",
		"assets/demo/models/block/oak_door/oak_door_bottom_right.json",
		"assets/demo/models/block/oak_door/oak_door_bottom_right_open.json",
		"assets/demo/models/block/oak_door/oak_door_top_left.json",
		"assets/demo/models/block/oak_door/oak_door_top_left_open.json",
		"assets/demo/models/block/oak_door/oak_door_top_right.json",
		"assets/demo/models/block/oak_door/oak_door_top_right_open.json",
		"assets/demo/models/block/oak_log/oak_log.json",
		"assets/demo/models/block/oak_log/oak_log_horizontal.json",
		"assets/demo/textures/block/end.png",
		"assets/demo/textures/block/side.png",
		// A door's item is a flat picture of its own, so only the log
		// gets an item model.
		"assets/minecraft/blockstates/oak_door.json",
		"assets/minecraft/blockstates/oak_log.json",
		"assets/minecraft/items/oak_log.json",
	}
	if strings.Join(paths, "\n") != strings.Join(want, "\n") {
		t.Fatalf("files:\n%s", strings.Join(paths, "\n"))
	}
	model := out["assets/demo/models/block/oak_log/oak_log_horizontal.json"]
	for _, s := range []string{`"parent": "minecraft:block/oak_log_horizontal"`, `"side": "demo:block/side"`, `"end": "demo:block/end"`} {
		if !strings.Contains(model, s) {
			t.Errorf("model lacks %s:\n%s", s, model)
		}
	}
	state := out["assets/minecraft/blockstates/oak_log.json"]
	if !strings.Contains(state, `"model": "demo:block/oak_log/oak_log_horizontal"`) || !strings.Contains(state, `"x": 90`) || strings.Contains(state, "minecraft:block/") {
		t.Errorf("block state:\n%s", state)
	}
	if item := out["assets/minecraft/items/oak_log.json"]; !strings.Contains(item, `"model": "demo:block/oak_log/oak_log"`) {
		t.Errorf("item model:\n%s", item)
	}
}

// Blocks that share vanilla models, as copper and waxed copper do, are
// retextured each on its own.
func TestRetextureSharedModels(t *testing.T) {
	src := `Retexture minecraft:copper_block { textures = { all = Texture("art/side.png") } }
Retexture minecraft:waxed_copper_block { textures = { all = Texture("art/end.png") } }
`
	_, r := assetProject(t, retextureFiles(), src)
	if d := diagText(r); d != "" {
		t.Fatalf("diagnostics:\n%s", d)
	}
	out := emittedFiles(r)
	if s := out["assets/minecraft/blockstates/waxed_copper_block.json"]; !strings.Contains(s, "demo:block/waxed_copper_block/copper_block") {
		t.Errorf("waxed block state:\n%s", s)
	}
	if m := out["assets/demo/models/block/waxed_copper_block/copper_block.json"]; !strings.Contains(m, "demo:block/end") {
		t.Errorf("waxed model:\n%s", m)
	}
}

func TestRetextureErrors(t *testing.T) {
	src := `Retexture minecraft:oak_stairs { textures = { sid = Texture("art/side.png") } }
Retexture oak_log { textures = { side = Texture("art/side.png") } }
Retexture minecraft:oak_planks { textures = { particle = Texture("art/side.png") } }
`
	_, r := assetProject(t, retextureFiles(), src)
	got := diagText(r)
	want := "mms/retexture: minecraft:oak_stairs has no texture slot 'sid'; its faces show bottom, side, top\n" +
		"mms/retexture: Retexture names a vanilla block, such as minecraft:oak_log, and 'demo:oak_log' is not one\n"
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

// Every vanilla block state retextures into models, block states and
// item models the checker accepts.
func TestRetextureEveryBlock(t *testing.T) {
	v, err := spec.DefaultVanilla()
	if err != nil {
		t.Fatal(err)
	}
	var b strings.Builder
	for k := range v["assets/blockstates"] {
		fmt.Fprintf(&b, "Retexture minecraft:%s { textures = { particle = Texture(\"art/side.png\") } }\n", k)
	}
	_, r := assetProject(t, retextureFiles(), b.String())
	for _, d := range r.Diagnostics("main.mms") {
		// Item frames have block states but are entities.
		if d.Code == "mms/unknown-block" && strings.Contains(d.Message, "item_frame") {
			continue
		}
		if d.Severity == syntax.SeverityError {
			t.Errorf("%s: %s", d.Code, d.Message)
		}
	}
}
