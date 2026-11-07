package project

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/spec"
)

func TestUnusedAssets(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png":  {Data: pngBytes}, // the icon
		"art/used.png":  {Data: pngBytes},
		"art/stray.png": {Data: pngBytes},
		"s/base.nbt":    {Data: piece(t, "demo:tops|demo:bottom|demo:top")},
		"s/top.nbt":     {Data: piece(t, "minecraft:empty|demo:top|minecraft:empty")},
		"s/old.nbt":     {Data: piece(t)},
		"build/x.png":   {Data: pngBytes}, // the output folder is not the project's
		"notes/a.txt":   {Data: []byte("x")},
	}
	src := `PaintingVariant p { width = 1, height = 1, asset_id = Texture("art/used.png", "used") }
Structure s = Jigsaw(biomes = #minecraft:is_overworld, step = surface_structures, spawn_overrides = {}, start_pool = starts, size = 2, start_height = Constant({ absolute = 0 }), max_distance_from_center = 80, use_expansion_hack = false)
` + pool("starts", `StructureFile("s/base.nbt", "base")`) + pool("tops", `StructureFile("s/top.nbt", "top")`) + pool("lonely")
	p := New(spec.MustDefault(), Config{Namespace: "demo", Description: "d", Icon: "art/icon.png", Output: "build"})
	p.Files = files
	p.SetFile("main.mms", src)
	r := analyze(t, p)
	got, _ := r.UnusedAssets()
	if fmt.Sprint(got) != "[art/stray.png s/old.nbt]" {
		t.Errorf("unused assets = %v", got)
	}
	// A pool counts as used when a jigsaw block draws from it.
	var unused []string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/unused" {
			unused = append(unused, d.Message)
		}
	}
	if s := strings.Join(unused, "\n"); s != "TemplatePool 'demo:lonely' is never used" {
		t.Errorf("unused:\n%s", s)
	}

	// The rule can be turned off, and says nothing while the project has errors.
	p.Config.Lint = map[string]string{"unused-asset": "off"}
	if r := analyze(t, p); len(first(r.UnusedAssets())) != 0 {
		t.Error("reported with the rule off")
	}
	p.Config.Lint = nil
	p.SetFile("main.mms", src+"Nope x = 1\n")
	if r := analyze(t, p); len(first(r.UnusedAssets())) != 0 {
		t.Error("reported while the project has errors")
	}
}

func first[A, B any](a A, _ B) A { return a }

func analyze(t *testing.T, p *Project) *Result {
	t.Helper()
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return r
}
