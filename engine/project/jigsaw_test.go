package project

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/nbt"
)

// piece builds a structure file with a stone block and the given jigsaw
// blocks, each as pool|name|target.
func piece(t *testing.T, jigsaws ...string) []byte {
	t.Helper()
	palette := []any{
		map[string]any{"Name": "minecraft:stone"},
		map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": "north_up"}},
	}
	blocks := []any{map[string]any{"pos": []any{int32(0), int32(0), int32(0)}, "state": int32(0)}}
	for i, j := range jigsaws {
		f := strings.Split(j, "|")
		blocks = append(blocks, map[string]any{
			"pos": []any{int32(i + 1), int32(0), int32(0)}, "state": int32(1),
			"nbt": map[string]any{"id": "minecraft:jigsaw", "pool": f[0], "name": f[1], "target": f[2], "final_state": "minecraft:air", "joint": "rollable"},
		})
	}
	data, err := nbt.Encode(map[string]any{"DataVersion": int32(4440), "size": []any{int32(3), int32(1), int32(1)}, "palette": palette, "blocks": blocks, "entities": []any{}})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func pool(name string, locations ...string) string {
	var els []string
	for _, l := range locations {
		els = append(els, `{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = `+l+`, projection = "rigid", processors = "minecraft:empty" } }`)
	}
	return "TemplatePool " + name + " = {\n  fallback = minecraft:empty\n  elements = [" + strings.Join(els, ", ") + "]\n}\n"
}

func jigsawDiags(r *Result) string {
	var sb strings.Builder
	for _, d := range r.Diagnostics("main.mms") {
		if strings.HasPrefix(d.Code, "lint/jigsaw") {
			sb.WriteString(d.Code + ": " + d.Message + "\n")
		}
	}
	return sb.String()
}

func TestJigsawPoolMissing(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png": {Data: pngBytes},
		"s/base.nbt":   {Data: piece(t, "demo:tops|demo:bottom|demo:top", "minecraft:village/plains/streets|minecraft:a|minecraft:b", "other:x|demo:y|demo:z")},
		"s/notnbt.nbt": {Data: nbtBytes},
	}
	_, r := assetProject(t, files, pool("starts", `Structure("s/base.nbt", "base")`)+pool("junk", `Structure("s/notnbt.nbt")`))
	got := jigsawDiags(r)
	want := "lint/jigsaw-pool: 's/base.nbt' has a jigsaw block at 1 0 0 that draws from the pool 'demo:tops', which no TemplatePool declares; the structure stops growing there\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	// The fix declares the pool, which then has no pieces.
	var fixed string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/jigsaw-pool" {
			if len(d.Fixes) == 0 || d.Fixes[0].Title != "Declare TemplatePool tops" {
				t.Fatalf("fixes: %+v", d.Fixes)
			}
			src := string(r.Unit("main.mms").File.Src)
			fixed = src + d.Fixes[0].Edits[0].Text
		}
	}
	_, r = assetProject(t, files, fixed)
	got = jigsawDiags(r)
	want = "lint/jigsaw-target: 's/base.nbt' has a jigsaw block at 1 0 0 that draws from the pool 'demo:tops', which has no structure pieces yet, so nothing joins there\n"
	if got != want {
		t.Fatalf("after fix, got:\n%s\nwant:\n%s", got, want)
	}
}

func TestJigsawTarget(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png": {Data: pngBytes},
		"s/base.nbt":   {Data: piece(t, "demo:tops|demo:bottom|demo:top", "demo:tops|demo:side|demo:sideways")},
		"s/top.nbt":    {Data: piece(t, "minecraft:empty|demo:top|minecraft:empty")},
		"s/roof.nbt":   {Data: piece(t, "minecraft:empty|demo:roof|minecraft:empty")},
	}
	src := pool("starts", `Structure("s/base.nbt", "base")`) + pool("tops", `Structure("s/top.nbt", "top")`, `Structure("s/roof.nbt", "roof")`)
	_, r := assetProject(t, files, src)
	got := jigsawDiags(r)
	want := "lint/jigsaw-target: 's/base.nbt' has a jigsaw block at 2 0 0 that targets 'demo:sideways' in the pool 'demo:tops', but no piece of that pool has a jigsaw block named 'demo:sideways', so the pieces never connect (the pieces' jigsaw names: demo:roof, demo:top)\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	// A pool with a piece from elsewhere is not checked: its jigsaws are unknown.
	src = pool("starts", `Structure("s/base.nbt", "base")`) + pool("tops", `Structure("s/top.nbt", "top")`, `minecraft:shipwreck/rightsideup_backhalf`)
	_, r = assetProject(t, files, src)
	if got := jigsawDiags(r); got != "" {
		t.Fatalf("unknown piece: got %s", got)
	}
}
