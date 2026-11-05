package project

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// pack is a data pack written as "path=json" lines.
func pack(lines ...string) []OutFile {
	var out []OutFile
	for _, l := range lines {
		p, data, _ := strings.Cut(l, "=")
		out = append(out, OutFile{Path: p, Data: []byte(data)})
	}
	return out
}

func asOld(files []OutFile) map[string][]byte {
	m := map[string][]byte{}
	for _, f := range files {
		m[f.Path] = f.Data
	}
	return m
}

// TestChangedDimensions checks which dimensions a change resets: those
// whose files lead to it by ids, and those a changed structure set can
// reach by biome.
func TestChangedDimensions(t *testing.T) {
	base := []string{
		`data/a/dimension/caves.json={"type":"a:caves","generator":{"type":"minecraft:noise","settings":"a:caves","biome_source":{"type":"minecraft:fixed","biome":"a:deep"}}}`,
		`data/a/dimension_type/caves.json={"height":256}`,
		`data/a/worldgen/noise_settings/caves.json={"noise_router":{"final_density":"a:shape"}}`,
		`data/a/worldgen/density_function/shape.json={"type":"minecraft:noise","noise":"a:wiggle"}`,
		`data/a/worldgen/noise/wiggle.json={"firstOctave":-4}`,
		`data/a/worldgen/biome/deep.json={"features":[["a:glow"]]}`,
		`data/a/worldgen/placed_feature/glow.json={"feature":"minecraft:glow_lichen"}`,
		`data/a/dimension/sky.json={"type":"minecraft:overworld","generator":{"type":"minecraft:noise","settings":"minecraft:floating_islands","biome_source":{"type":"minecraft:fixed","biome":"a:high"}}}`,
		`data/a/worldgen/biome/high.json={"features":[]}`,
		`data/a/tags/worldgen/biome/has_tower.json={"values":["a:high"]}`,
		`data/a/worldgen/structure/tower.json={"biomes":"#a:has_tower","start_pool":"a:tower"}`,
		`data/a/worldgen/template_pool/tower.json={"elements":[{"element":{"location":"a:tower"}}]}`,
		`data/a/structure/tower.nbt=v1`,
		`data/a/worldgen/structure_set/tower.json={"structures":[{"structure":"a:tower"}]}`,
		`data/minecraft/dimension/overworld.json={"type":"minecraft:overworld","generator":{"settings":"a:caves"}}`,
		`data/a/function/hi.mcfunction=say hi`,
	}
	old := asOld(pack(base...))
	with := func(path, data string) []OutFile {
		var lines []string
		for _, l := range base {
			if strings.HasPrefix(l, path+"=") {
				if data == "" {
					continue
				}
				l = path + "=" + data
			}
			lines = append(lines, l)
		}
		return pack(lines...)
	}
	for _, c := range []struct {
		name, path, data string
		want             []string
	}{
		{"nothing", "data/a/function/hi.mcfunction", "say hi", nil},
		{"function", "data/a/function/hi.mcfunction", "say bye", nil},
		{"noise under the router", "data/a/worldgen/noise/wiggle.json", `{"firstOctave":-5}`, []string{"a:caves"}},
		{"feature of a biome", "data/a/worldgen/placed_feature/glow.json", `{"feature":"minecraft:vines"}`, []string{"a:caves"}},
		{"dimension type", "data/a/dimension_type/caves.json", `{"height":320}`, []string{"a:caves"}},
		{"structure piece in a tagged biome", "data/a/structure/tower.nbt", "v2", []string{"a:sky"}},
		{"removed feature file", "data/a/worldgen/placed_feature/glow.json", "", []string{"a:caves"}},
	} {
		if got := ChangedDimensions(old, with(c.path, c.data)); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
	if got := ChangedDimensions(nil, pack(base...)); !reflect.DeepEqual(got, []string{"a:caves", "a:sky"}) {
		t.Errorf("first build: got %v", got)
	}
}

// TestChangedDimensionsVanillaBiomes checks that a structure in vanilla
// biomes resets a dimension that uses vanilla biomes.
func TestChangedDimensionsVanillaBiomes(t *testing.T) {
	dim := `data/a/dimension/copy.json={"type":"minecraft:overworld","generator":{"biome_source":{"type":"minecraft:multi_noise","preset":"minecraft:overworld"}}}`
	other := `data/a/dimension/flat.json={"type":"minecraft:overworld","generator":{"biome_source":{"type":"minecraft:fixed","biome":"a:only"}}}`
	set := `data/a/worldgen/structure_set/hut.json={"structures":[{"structure":"a:hut"}]}`
	before := pack(dim, other, set, `data/a/worldgen/structure/hut.json={"biomes":"#minecraft:is_forest","size":1}`)
	after := pack(dim, other, set, `data/a/worldgen/structure/hut.json={"biomes":"#minecraft:is_forest","size":2}`)
	if got := ChangedDimensions(asOld(before), after); !reflect.DeepEqual(got, []string{"a:copy"}) {
		t.Errorf("got %v, want [a:copy]", got)
	}
}

// TestResetDimensions checks that a reset moves only the chunk folders,
// keeps the latest reset, and finds vanilla dimensions in old worlds.
func TestResetDimensions(t *testing.T) {
	world := t.TempDir()
	mk := func(rel string) {
		os.MkdirAll(filepath.Join(world, filepath.FromSlash(rel)), 0o755)
		os.WriteFile(filepath.Join(world, filepath.FromSlash(rel), "r.0.0.mca"), []byte("x"), 0o644)
	}
	mk("dimensions/a/deep/caves/region")
	mk("dimensions/a/deep/caves/poi")
	mk("dimensions/a/deep/caves/data")
	mk("DIM-1/region")
	got, err := ResetDimensions(world, []string{"a:deep/caves", "a:never", "minecraft:the_nether"})
	if err != nil {
		t.Fatal(err)
	}
	want := []DimensionReset{
		{ID: "a:deep/caves", Moved: []string{"dimensions/a/deep/caves/region", "dimensions/a/deep/caves/poi"}},
		{ID: "a:never"},
		{ID: "minecraft:the_nether", Moved: []string{"DIM-1/region"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v\nwant %+v", got, want)
	}
	for _, rel := range []string{"mms-reset/dimensions/a/deep/caves/region/r.0.0.mca", "mms-reset/DIM-1/region/r.0.0.mca", "dimensions/a/deep/caves/data/r.0.0.mca"} {
		if _, err := os.Stat(filepath.Join(world, rel)); err != nil {
			t.Errorf("missing %s", rel)
		}
	}
	mk("dimensions/a/deep/caves/region")
	if _, err := ResetDimensions(world, []string{"a:deep/caves"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(world, "mms-reset/DIM-1")); !os.IsNotExist(err) {
		t.Errorf("an earlier reset was kept")
	}
	for _, bad := range []string{"caves", "a:../x", "#a:caves", "a:"} {
		if err := CheckDimensionID(bad); err == nil {
			t.Errorf("%q passed", bad)
		}
	}
}
