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

const assetSrc = `PaintingVariant sunset = { width = 2, height = 1, asset_id = Texture("art/Sunset Final.png", "sunset") }
PaintingVariant again = { width = 2, height = 1, asset_id = demo:sunset }
JukeboxSong tune = { sound_event = Sound("music/tune.ogg"), description = "Tune", length_in_seconds = 30, comparator_output = 3 }
TemplatePool houses = {
  fallback = minecraft:empty
  elements = [{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = Structure("builds/house.nbt"), projection = "rigid", processors = "minecraft:empty" } }]
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
	_, r := assetProject(t, files, `TemplatePool houses = {
  fallback = minecraft:empty
  elements = [{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = Structure("builds/house.nbt"), projection = "rigid", processors = "minecraft:empty" } }]
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
	_, r := assetProject(t, files, `PaintingVariant a = { width = 1, height = 1, asset_id = Texture("art/fake.png") }
PaintingVariant b = { width = 1, height = 1, asset_id = Texture("art/nope.png") }
PaintingVariant c = { width = 1, height = 1, asset_id = Sound("music/t.ogg") }
PaintingVariant d = { width = 1, height = 1, asset_id = Texture("music/t.ogg") }
PaintingVariant e = { width = 1, height = 1, asset_id = Texture("../x.png") }
PaintingVariant f = { width = 1, height = 1, asset_id = Texture("art/a.png", "same") }
PaintingVariant g = { width = 1, height = 1, asset_id = Texture("art/b.png", "same") }
PaintingVariant h = { width = 1, height = 1, asset_id = "demo:elsewhere" }
PaintingVariant i = { width = 1, height = 1, asset_id = "other:elsewhere" }
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
	p, r := assetProject(t, files, `PaintingVariant a = { width = 1, height = 1, asset_id = Texture("art/a.png") }`+"\n")
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
