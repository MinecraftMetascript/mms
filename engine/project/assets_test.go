package project

import (
	"context"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

var (
	pngBytes = []byte("\x89PNG\r\n\x1a\nrest")
	oggBytes = []byte("OggS rest")
	nbtBytes = []byte{0x1f, 0x8b, 8, 0}
)

func assetProject(t *testing.T, files fstest.MapFS, src string) (*Project, *Result) {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo", Description: "d", Icon: "art/icon.png"})
	p.Files = files
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return p, r
}

func diagText(r *Result) string {
	var sb strings.Builder
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityHint {
			continue
		}
		sb.WriteString(d.Code + ": " + d.Message + "\n")
	}
	return sb.String()
}

const assetSrc = `PaintingVariant sunset { width = 2, height = 1, asset_id = Texture("art/Sunset Final.png", "sunset") }
PaintingVariant again { width = 2, height = 1, asset_id = demo:sunset }
JukeboxSong tune { sound_event = Sound("music/tune.ogg"), description = "Tune", length_in_seconds = 30, comparator_output = 3 }
TemplatePool houses {
  fallback = minecraft:empty
  elements = [{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = StructureFile("builds/house.nbt"), projection = "rigid", processors = "minecraft:empty" } }]
}
`

func TestAssetsEmit(t *testing.T) {
	files := fstest.MapFS{
		"art/Sunset Final.png": {Data: pngBytes},
		"art/icon.png":         {Data: pngBytes},
		"art/unused.png":       {Data: pngBytes},
		"music/tune.ogg":       {Data: oggBytes},
		"builds/house.nbt":     {Data: nbtBytes},
	}
	_, r := assetProject(t, files, assetSrc)
	if d := diagText(r); d != "" {
		t.Fatalf("unexpected diagnostics:\n%s", d)
	}
	out := r.Emit()
	if err := ReadAssets(out, files); err != nil {
		t.Fatal(err)
	}
	data, res := SplitPacks(out)
	var got []string
	for _, f := range data {
		got = append(got, strings.TrimSpace("data "+f.Path+" <- "+f.Asset))
	}
	for _, f := range res {
		got = append(got, strings.TrimSpace("res "+f.Path+" <- "+f.Asset))
	}
	want := `data data/demo/jukebox_song/tune.json <-
data data/demo/painting_variant/again.json <-
data data/demo/painting_variant/sunset.json <-
data data/demo/structure/house.nbt <- builds/house.nbt
data data/demo/worldgen/template_pool/houses.json <-
data pack.mcmeta <-
data pack.png <- art/icon.png
res assets/demo/sounds.json <-
res assets/demo/sounds/tune.ogg <- music/tune.ogg
res assets/demo/textures/painting/sunset.png <- art/Sunset Final.png
res pack.mcmeta <-
res pack.png <- art/icon.png`
	if g := strings.Join(got, "\n"); g != want {
		t.Errorf("files:\n%s\nwant:\n%s", g, want)
	}
	for _, f := range out {
		switch f.Path {
		case "data/demo/painting_variant/sunset.json":
			if !strings.Contains(string(f.Data), `"asset_id": "demo:sunset"`) {
				t.Errorf("sunset: %s", f.Data)
			}
		case "data/demo/jukebox_song/tune.json":
			// The game's sound_event registry is fixed; a pack's own sound
			// only loads written inline.
			if s := strings.Join(strings.Fields(string(f.Data)), ""); !strings.Contains(s, `"sound_event":{"sound_id":"demo:tune"}`) {
				t.Errorf("tune: %s", s)
			}
		case "data/demo/worldgen/template_pool/houses.json":
			if !strings.Contains(string(f.Data), `"location": "demo:house"`) {
				t.Errorf("houses: %s", f.Data)
			}
		case "assets/demo/sounds.json":
			if s := strings.Join(strings.Fields(string(f.Data)), ""); s != `{"tune":{"sounds":[{"name":"demo:tune"}]}}` {
				t.Errorf("sounds.json: %s", s)
			}
		case "assets/demo/sounds/tune.ogg":
			if string(f.Data) != string(oggBytes) {
				t.Errorf("tune.ogg not copied: %q", f.Data)
			}
		case "pack.mcmeta":
			if f.Pack == sema.ResourcePack && !strings.Contains(string(f.Data), `"min_format": 97`) {
				t.Errorf("resource pack.mcmeta: %s", f.Data)
			}
		}
	}
}

func TestAssetsNoResourcePack(t *testing.T) {
	files := fstest.MapFS{"builds/house.nbt": {Data: nbtBytes}, "art/icon.png": {Data: pngBytes}}
	_, r := assetProject(t, files, `TemplatePool houses {
  fallback = minecraft:empty
  elements = [{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = StructureFile("builds/house.nbt"), projection = "rigid", processors = "minecraft:empty" } }]
}
`)
	if _, res := SplitPacks(r.Emit()); len(res) != 0 {
		t.Errorf("resource pack without textures or sounds: %v", res)
	}
}

func TestAssetsDiagnostics(t *testing.T) {
	files := fstest.MapFS{
		"art/fake.png": {Data: []byte("text")},
		"art/a.png":    {Data: pngBytes},
		"art/b.png":    {Data: pngBytes},
		"music/t.ogg":  {Data: oggBytes},
	}
	_, r := assetProject(t, files, `PaintingVariant a { width = 1, height = 1, asset_id = Texture("art/fake.png") }
PaintingVariant b { width = 1, height = 1, asset_id = Texture("art/nope.png") }
PaintingVariant c { width = 1, height = 1, asset_id = Sound("music/t.ogg") }
PaintingVariant d { width = 1, height = 1, asset_id = Texture("music/t.ogg") }
PaintingVariant e { width = 1, height = 1, asset_id = Texture("../x.png") }
PaintingVariant f { width = 1, height = 1, asset_id = Texture("art/a.png", "same") }
PaintingVariant g { width = 1, height = 1, asset_id = Texture("art/b.png", "same") }
PaintingVariant h { width = 1, height = 1, asset_id = "demo:elsewhere" }
PaintingVariant i { width = 1, height = 1, asset_id = "other:elsewhere" }
`)
	want := `mms/asset: 'art/fake.png' is not a PNG image
mms/asset: no file 'art/nope.png' in the project
mms/type-mismatch: Sound(...) makes a sound_event id, but this field expects a texture id
mms/asset: Texture expects a .png file, not 't.ogg'
mms/asset: asset paths are relative to the project folder and stay inside it
mms/asset-conflict: 'art/b.png' and 'art/a.png' both become texture 'demo:same'; give this one its own id name, e.g. Texture("art/b.png", "same_2")
lint/external-asset: this project does not provide the texture 'demo:elsewhere'; it must come from another pack, or include the file with Texture("path")
`
	if got := diagText(r); got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

// TestAssetsIncremental: a declaration's cached result is dropped when a
// file it reads appears, disappears or changes.
func TestAssetsIncremental(t *testing.T) {
	files := fstest.MapFS{"art/icon.png": {Data: pngBytes}}
	p, r := assetProject(t, files, `PaintingVariant a { width = 1, height = 1, asset_id = Texture("art/a.png") }`+"\n")
	if !strings.Contains(diagText(r), "no file 'art/a.png'") {
		t.Fatalf("want a missing-file error, got:\n%s", diagText(r))
	}
	files["art/a.png"] = &fstest.MapFile{Data: []byte("nope"), ModTime: time.Unix(1, 0)}
	r, _ = p.Analyze(context.Background())
	if !strings.Contains(diagText(r), "is not a PNG image") {
		t.Fatalf("want a not-PNG error after adding the file, got:\n%s", diagText(r))
	}
	files["art/a.png"] = &fstest.MapFile{Data: pngBytes, ModTime: time.Unix(2, 0)}
	r, _ = p.Analyze(context.Background())
	if d := diagText(r); d != "" {
		t.Fatalf("want no diagnostics after fixing the file, got:\n%s", d)
	}
}

// TestAssetsExtension: ids of files leave the extension out, since the
// game adds it. One written with it is fixed in the output and warned
// about, with a fix that removes it from the source.
func TestAssetsExtension(t *testing.T) {
	files := fstest.MapFS{"art/a.png": {Data: pngBytes}, "art/icon.png": {Data: pngBytes}}
	src := `PaintingVariant a { width = 1, height = 1, asset_id = Texture("art/a.png", "sun.png") }
PaintingVariant b { width = 1, height = 1, asset_id = "demo:sun.png" }
Model item/ruby { parent = "item/generated.json", textures = { layer0 = demo:sun.png, layer1 = "item/diamond.png" } }
`
	_, r := assetProject(t, files, src)
	want := `mms/asset-extension: texture ids leave out the file extension, which the game adds itself; MMS writes 'sun'
mms/asset-extension: texture ids leave out the file extension, which the game adds itself; MMS writes 'sun'
mms/asset-extension: model ids leave out the file extension, which the game adds itself; MMS writes 'item/generated'
mms/asset-extension: texture ids leave out the file extension, which the game adds itself; MMS writes 'sun'
lint/texture-outside-atlas: no atlas stitches the texture 'demo:sun', so models show the missing texture; models find textures under textures/block/ and textures/item/, such as 'demo:block/sun', or add it to an Atlas blocks declaration in the minecraft namespace
mms/asset-extension: texture ids leave out the file extension, which the game adds itself; MMS writes 'item/diamond'
`
	if got := diagText(r); got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
	fixed := src
	ds := r.Diagnostics("main.mms")
	for i := len(ds) - 1; i >= 0; i-- {
		if ds[i].Code != "mms/asset-extension" {
			continue
		}
		if len(ds[i].Fixes) != 1 {
			t.Fatalf("%s: fixes %v", ds[i].Message, ds[i].Fixes)
		}
		e := ds[i].Fixes[0].Edits[0]
		rs := []rune(fixed)
		fixed = string(rs[:e.Sp.Start]) + e.Text + string(rs[e.Sp.End:])
	}
	if want := strings.NewReplacer(`"sun.png"`, `"sun"`, ":sun.png", ":sun", "generated.json", "generated", "diamond.png", "diamond").Replace(src); fixed != want {
		t.Errorf("fixed source:\n%s", fixed)
	}
	for _, f := range r.Emit() {
		if strings.Contains(string(f.Data), ".png") || strings.Contains(string(f.Data), ".json") {
			t.Errorf("%s holds an extension: %s", f.Path, f.Data)
		}
		if strings.HasSuffix(f.Path, ".png.png") {
			t.Errorf("output path %s", f.Path)
		}
	}
}

// TestAssetsExtensionDeclName: a declaration named with its file's
// extension would be written as ruby.json.json, so it is an error.
func TestAssetsExtensionDeclName(t *testing.T) {
	_, r := assetProject(t, fstest.MapFS{"art/icon.png": {Data: pngBytes}}, "Model item/ruby.json { parent = \"item/generated\" }\n")
	ds := r.Diagnostics("main.mms")
	if len(ds) == 0 || ds[0].Code != "mms/invalid-name" || !strings.Contains(ds[0].Message, "use 'item/ruby'") || len(ds[0].Fixes) != 1 {
		t.Fatalf("diagnostics: %+v", ds)
	}
	if e := ds[0].Fixes[0].Edits[0]; e.Sp.End-e.Sp.Start != len(".json") {
		t.Errorf("fix edit %+v", e)
	}
}

// A Texture(...) in a Model goes under textures/block/, which the blocks
// atlas stitches; without that the game never loads the sprite.
func TestAssetsModelTextureAtlas(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png":      {Data: pngBytes},
		"tools/select.png":  {Data: pngBytes},
		"tools/wand.png":    {Data: pngBytes},
		"tools/stray.png":   {Data: pngBytes},
		"tools/own.png":     {Data: pngBytes},
		"tools/handled.png": {Data: pngBytes},
	}
	src := `Model spellbook/selection {
  parent = block/cube_all
  textures = { all = Texture("tools/select.png") }
}
CustomItem wand {
  base = stick
  texture = Texture("tools/wand.png")
}
Model stray {
  parent = item/generated
  textures = {
    layer0 = Texture("tools/stray.png", "gui/stray")
    layer1 = demo:loose
    layer2 = Texture("tools/own.png", "custom/own")
    layer3 = Texture("tools/handled.png", "item/handled")
  }
}
Atlas minecraft:blocks { sources = [{ type = "single", resource = "demo:custom/own" }] }
`
	_, r := assetProject(t, files, src)
	var atlas []string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/texture-outside-atlas" {
			atlas = append(atlas, d.Message)
		} else if d.Severity <= syntax.SeverityWarning && d.Code != "lint/external-asset" {
			t.Errorf("unexpected: %s: %s", d.Code, d.Message)
		}
	}
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/texture-outside-atlas" && strings.Contains(d.Message, "gui/stray") {
			if len(d.Fixes) == 0 || d.Fixes[0].Edits[0].Text != `"block/stray"` || d.Fixes[0].Safe {
				t.Errorf("fix: %+v", d.Fixes)
			}
		}
	}
	if len(atlas) != 2 || !strings.Contains(atlas[0], "'demo:gui/stray'") || !strings.Contains(atlas[0], `Texture("tools/stray.png", "block/stray")`) || !strings.Contains(atlas[1], "'demo:loose'") {
		t.Errorf("atlas warnings: %q", atlas)
	}
	out := r.Emit()
	paths := map[string]string{}
	for _, f := range out {
		paths[f.Path] = strings.Join(strings.Fields(string(f.Data)), "")
	}
	if s := paths["assets/demo/models/spellbook/selection.json"]; !strings.Contains(s, `"all":"demo:block/select"`) {
		t.Errorf("selection model: %s", s)
	}
	for _, p := range []string{"assets/demo/textures/block/select.png", "assets/demo/textures/item/wand.png"} {
		if _, ok := paths[p]; !ok {
			t.Errorf("no %s", p)
		}
	}
}

// pngOf returns the start of a PNG of the given size: its signature and
// IHDR chunk, which is all MMS reads.
func pngOf(w, h int) []byte {
	b := []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\x0dIHDR")
	b = append(b, byte(w>>24), byte(w>>16), byte(w>>8), byte(w), byte(h>>24), byte(h>>16), byte(h>>8), byte(h))
	return append(b, "rest"...)
}

func TestAssetsTextureAnimation(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png":        {Data: pngBytes},
		"art/lamp.png":        {Data: pngOf(16, 64)},
		"art/odd.png":         {Data: pngOf(16, 40)},
		"art/strip.png":       {Data: pngOf(16, 48)},
		"art/glow.png":        {Data: pngOf(16, 32)},
		"art/glow.png.mcmeta": {Data: []byte(`{"animation": {"frametime": 3, "interpolate": true}}`)},
		"art/stone.png":       {Data: pngOf(16, 16)},
	}
	src := `Model block/lamp {
  parent = block/cube_all
  textures = { all = Texture("art/lamp.png", animation = { frametime = 4, frames = [0, 1, { index = 3, time = 8 }, 4] }) }
}
Model block/odd { parent = block/cube_all, textures = { all = Texture("art/odd.png", animation = {}) } }
Model block/strip { parent = block/cube_all, textures = { all = Texture("art/strip.png") } }
Model block/glow { parent = block/cube_all, textures = { all = Texture("art/glow.png") } }
Model block/stone_again { parent = block/cube_all, textures = { all = Texture("art/stone.png", "minecraft:block/stone", texture = { blur = true }) } }
`
	_, r := assetProject(t, files, src)
	var got []string
	for _, d := range r.Diagnostics("main.mms") {
		if strings.HasPrefix(d.Code, "mms/texture-") {
			got = append(got, d.Code+": "+d.Message)
			switch {
			case strings.Contains(d.Message, "strip.png"):
				if len(d.Fixes) != 1 || d.Fixes[0].Edits[0].Text != ", animation = { frametime = 2 }" {
					t.Errorf("strip fix: %+v", d.Fixes)
				}
			case d.Code == "mms/texture-mcmeta":
				if len(d.Fixes) != 1 || d.Fixes[0].Edits[0].Text != ", animation = { frametime = 3, interpolate = true }" {
					t.Errorf("mcmeta fix: %+v", d.Fixes)
				}
			}
		} else if d.Severity <= syntax.SeverityWarning {
			t.Errorf("unexpected: %s: %s", d.Code, d.Message)
		}
	}
	want := []string{"frame 4 is not there", "16x40, which is not a whole number of 16x16 frames", "a strip of 3 square frames", "does not copy 'art/glow.png.mcmeta'"}
	if len(got) != len(want) {
		t.Fatalf("got %d texture diagnostics, want %d:\n%s", len(got), len(want), strings.Join(got, "\n"))
	}
	for i, w := range want {
		if !strings.Contains(got[i], w) {
			t.Errorf("diagnostic %d: %s\nwant it to say %q", i, got[i], w)
		}
	}
	paths := map[string]string{}
	for _, f := range r.Emit() {
		paths[f.Path] = strings.Join(strings.Fields(string(f.Data)), "")
	}
	if s := paths["assets/demo/textures/block/lamp.png.mcmeta"]; s != `{"animation":{"frametime":4,"frames":[0,1,{"index":3,"time":8},4]}}` {
		t.Errorf("lamp mcmeta: %q", s)
	}
	if _, ok := paths["assets/demo/textures/block/strip.png.mcmeta"]; ok {
		t.Error("strip has no animation, so no mcmeta")
	}
	if _, ok := paths["assets/minecraft/textures/block/stone.png"]; !ok {
		t.Error("no replaced vanilla texture")
	}
	if s := paths["assets/minecraft/textures/block/stone.png.mcmeta"]; s != `{"texture":{"blur":true}}` {
		t.Errorf("stone mcmeta: %q", s)
	}
	if s := paths["assets/demo/models/block/stone_again.json"]; !strings.Contains(s, `"all":"minecraft:block/stone"`) {
		t.Errorf("stone model: %s", s)
	}
}

func TestAssetsTextureMetaConflict(t *testing.T) {
	files := fstest.MapFS{"art/icon.png": {Data: pngBytes}, "art/lamp.png": {Data: pngOf(16, 32)}}
	src := `Model block/a { parent = block/cube_all, textures = { all = Texture("art/lamp.png", animation = { frametime = 2 }) } }
Model block/b { parent = block/cube_all, textures = { all = Texture("art/lamp.png", animation = { frametime = 4 }) } }
Model block/c { parent = block/cube_all, textures = { all = Texture("art/lamp.png", animation = { frametime = 2 }) } }
`
	_, r := assetProject(t, files, src)
	if s := diagText(r); strings.Count(s, "mms/asset-conflict") != 1 || !strings.Contains(s, "different named arguments") {
		t.Errorf("diagnostics:\n%s", s)
	}
}

func TestModelMissingTexture(t *testing.T) {
	files := fstest.MapFS{"art/icon.png": {Data: pngBytes}}
	src := `Model block/log { parent = block/cube_column, textures = { side = block/oak_log } }
Model block/full { parent = block/cube_column, textures = { side = block/oak_log, end = "#side" } }
Model block/mine { parent = block/log, textures = { end = block/stone } }
Model block/own {
  textures = { a = block/stone }
  elements = [{ from = [0, 0, 0], to = [16, 16, 16], faces = { up = { texture = "#a" }, down = { texture = "#b" } } }]
}
Model block/elsewhere { parent = other:block/thing }
Model item/flat { parent = item/generated, textures = { layer0 = item/stick } }
`
	_, r := assetProject(t, files, src)
	var got []string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/model-missing-texture" {
			got = append(got, d.Message)
		}
	}
	want := []string{
		"nothing fills the texture slot #end, which the faces of 'block/cube_column' use",
		"nothing fills the texture slot #b, which its elements use",
	}
	if len(got) != len(want) {
		t.Fatalf("got %q", got)
	}
	for i := range want {
		if !strings.HasPrefix(got[i], want[i]) {
			t.Errorf("got %q, want %q", got[i], want[i])
		}
	}
	for _, d := range r.Order {
		if d.ID() == "demo:block/mine" {
			slots, ok := r.ModelSlots(d)
			if !ok || len(slots) < 2 || slots[0].Name != "end" || slots[0].Filled != "minecraft:block/stone" {
				t.Errorf("slots: %+v", slots)
			}
		}
	}
}
