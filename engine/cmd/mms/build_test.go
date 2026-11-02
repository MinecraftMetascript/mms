package main

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestBuildAssets builds a project with a texture, a sound and a
// structure into a folder, a zip and a world, and checks both packs.
func TestBuildAssets(t *testing.T) {
	dir := t.TempDir()
	png := "\x89PNG\r\n\x1a\nimage"
	files := map[string]string{
		"mms.json":         `{"namespace":"demo","icon":"art/icon.png"}`,
		"art/icon.png":     png,
		"art/sunset.png":   png,
		"art/unused.png":   png,
		"music/tune.ogg":   "OggS sound",
		"builds/house.nbt": "\x1f\x8b\x08\x00",
		"main.mms": `PaintingVariant sunset = { width = 2, height = 1, asset_id = Texture("art/sunset.png") }
JukeboxSong tune = { sound_event = Sound("music/tune.ogg"), description = "Tune", length_in_seconds = 30, comparator_output = 3 }
TemplatePool houses = {
  fallback = minecraft:empty
  elements = [{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = Structure("builds/house.nbt"), projection = "rigid", processors = "minecraft:empty" } }]
}
`,
	}
	for name, text := range files {
		p := filepath.Join(dir, filepath.FromSlash(name))
		os.MkdirAll(filepath.Dir(p), 0o755)
		if err := os.WriteFile(p, []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	run := func(args ...string) {
		t.Helper()
		var out, errOut bytes.Buffer
		if code := runBuild(append(args, dir), &out, &errOut); code != 0 {
			t.Fatalf("build %v: exit %d: %s", args, code, errOut.String())
		}
	}
	listDir := func(root string) string {
		var got []string
		filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				rel, _ := filepath.Rel(root, p)
				got = append(got, filepath.ToSlash(rel))
			}
			return nil
		})
		sort.Strings(got)
		return strings.Join(got, " ")
	}
	listZip := func(path string) string {
		zr, err := zip.OpenReader(path)
		if err != nil {
			t.Fatal(err)
		}
		defer zr.Close()
		var got []string
		for _, f := range zr.File {
			got = append(got, f.Name)
		}
		sort.Strings(got)
		return strings.Join(got, " ")
	}
	const data = "data/demo/jukebox_song/tune.json data/demo/painting_variant/sunset.json data/demo/structure/house.nbt data/demo/worldgen/template_pool/houses.json pack.mcmeta pack.png"
	const res = "assets/demo/sounds.json assets/demo/sounds/tune.ogg assets/demo/textures/painting/sunset.png pack.mcmeta pack.png"

	run()
	if got := listDir(filepath.Join(dir, "build")); got != data {
		t.Errorf("build/:\n got %s\nwant %s", got, data)
	}
	if got := listDir(filepath.Join(dir, "build-resources")); got != res {
		t.Errorf("build-resources/:\n got %s\nwant %s", got, res)
	}
	if b, _ := os.ReadFile(filepath.Join(dir, "build-resources/assets/demo/textures/painting/sunset.png")); string(b) != png {
		t.Errorf("texture not copied byte for byte: %q", b)
	}

	zpath := filepath.Join(t.TempDir(), "pack.zip")
	run("-zip", zpath)
	if got := listZip(zpath); got != data {
		t.Errorf("pack.zip:\n got %s\nwant %s", got, data)
	}
	if got := listZip(strings.TrimSuffix(zpath, ".zip") + "-resources.zip"); got != res {
		t.Errorf("pack-resources.zip:\n got %s\nwant %s", got, res)
	}

	game := t.TempDir()
	world := filepath.Join(game, "saves", "New World")
	os.MkdirAll(world, 0o755)
	os.WriteFile(filepath.Join(world, "level.dat"), nil, 0o644)
	run("-world", world)
	if got := listDir(filepath.Join(world, "datapacks", "demo")); got != data {
		t.Errorf("world data pack:\n got %s\nwant %s", got, data)
	}
	if got := listZip(filepath.Join(game, "resourcepacks", "demo-resources.zip")); got != res {
		t.Errorf("world resource pack:\n got %s\nwant %s", got, res)
	}
}
