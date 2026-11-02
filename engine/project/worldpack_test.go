package project

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"sort"
	"strings"
	"testing"
)

func unzip(t *testing.T, b []byte) map[string]string {
	t.Helper()
	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, f := range zr.File {
		rc, _ := f.Open()
		d, _ := io.ReadAll(rc)
		rc.Close()
		out[f.Name] = string(d)
	}
	return out
}

func names(m map[string]string) string {
	var s []string
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	return strings.Join(s, " ")
}

func meta(desc string) OutFile {
	return OutFile{Path: "pack.mcmeta", Data: packMeta(desc, 97)}
}

func merge(t *testing.T, old []byte, owner string, files ...OutFile) *WorldPackResult {
	t.Helper()
	r, err := MergeWorldPack(old, owner, files)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func TestMergeWorldPackKeepsOtherProjects(t *testing.T) {
	a := merge(t, nil, "embers",
		meta("Embers"),
		OutFile{Path: "assets/embers/textures/item/ember.png", Data: []byte("png")},
		OutFile{Path: "assets/minecraft/lang/en_us.json", Data: []byte(`{"item.embers.ember":"Ember"}`)},
		OutFile{Path: "assets/minecraft/atlases/blocks.json", Data: []byte(`{"sources":[{"type":"directory","source":"embers"}]}`)},
	)
	b := merge(t, a.Zip, "barrels",
		meta("Barrels"),
		OutFile{Path: "assets/barrels/models/block/barrel.json", Data: []byte(`{}`)},
		OutFile{Path: "assets/minecraft/lang/en_us.json", Data: []byte(`{"block.barrels.barrel":"Barrel"}`)},
		OutFile{Path: "assets/minecraft/atlases/blocks.json", Data: []byte(`{"sources":[{"type":"directory","source":"barrels"}]}`)},
	)
	if strings.Join(b.Others, ",") != "embers" {
		t.Errorf("others = %v", b.Others)
	}
	z := unzip(t, b.Zip)
	want := "assets/barrels/models/block/barrel.json assets/embers/textures/item/ember.png assets/minecraft/atlases/blocks.json assets/minecraft/lang/en_us.json mms-world-pack.json pack.mcmeta"
	if got := names(z); got != want {
		t.Fatalf("files:\n got %s\nwant %s", got, want)
	}
	var lang map[string]string
	json.Unmarshal([]byte(z["assets/minecraft/lang/en_us.json"]), &lang)
	if lang["item.embers.ember"] != "Ember" || lang["block.barrels.barrel"] != "Barrel" {
		t.Errorf("lang not merged: %s", z["assets/minecraft/lang/en_us.json"])
	}
	if s := z["assets/minecraft/atlases/blocks.json"]; !strings.Contains(s, `"embers"`) || !strings.Contains(s, `"barrels"`) {
		t.Errorf("atlas sources not merged: %s", s)
	}
	if s := z["pack.mcmeta"]; !strings.Contains(s, `"Barrels + Embers"`) && !strings.Contains(s, `"Embers + Barrels"`) {
		t.Errorf("pack.mcmeta description: %s", s)
	}

	// Rebuilding embers without its texture and lang drops them, and
	// leaves barrels alone.
	c := merge(t, b.Zip, "embers",
		meta("Embers"),
		OutFile{Path: "assets/embers/textures/item/coal.png", Data: []byte("png2")},
	)
	z = unzip(t, c.Zip)
	want = "assets/barrels/models/block/barrel.json assets/embers/textures/item/coal.png assets/minecraft/atlases/blocks.json assets/minecraft/lang/en_us.json mms-world-pack.json pack.mcmeta"
	if got := names(z); got != want {
		t.Fatalf("after rebuild:\n got %s\nwant %s", got, want)
	}
	if s := z["assets/minecraft/lang/en_us.json"]; strings.Contains(s, "embers") || !strings.Contains(s, "barrels") {
		t.Errorf("lang after rebuild: %s", s)
	}
	if s := z["assets/minecraft/atlases/blocks.json"]; strings.Contains(s, `"embers"`) {
		t.Errorf("atlas after rebuild: %s", s)
	}
}

func TestMergeWorldPackConflictLastBuildWins(t *testing.T) {
	a := merge(t, nil, "one", meta("One"), OutFile{Path: "assets/minecraft/textures/block/stone.png", Data: []byte("one")})
	b := merge(t, a.Zip, "two", meta("Two"), OutFile{Path: "assets/minecraft/textures/block/stone.png", Data: []byte("two")})
	if strings.Join(b.Replaced, ",") != "assets/minecraft/textures/block/stone.png (one)" {
		t.Errorf("replaced = %v", b.Replaced)
	}
	if got := unzip(t, b.Zip)["assets/minecraft/textures/block/stone.png"]; got != "two" {
		t.Errorf("stone = %q", got)
	}
	// "one" owns nothing now, so it is dropped from the manifest.
	if len(b.Others) != 0 {
		t.Errorf("others = %v", b.Others)
	}
}

func TestMergeWorldPackKeepsHandMadePack(t *testing.T) {
	old, _ := Zip([]OutFile{
		{Path: "pack.mcmeta", Data: packMeta("Hand made", 97)},
		{Path: "pack.png", Data: []byte("icon")},
		{Path: "assets/mypack/textures/x.png", Data: []byte("x")},
		{Path: "assets/minecraft/lang/en_us.json", Data: []byte(`{"a":"A"}`)},
	})
	r := merge(t, old, "embers", meta("Embers"), OutFile{Path: "assets/minecraft/lang/en_us.json", Data: []byte(`{"b":"B"}`)})
	if r.Adopted || strings.Join(r.Others, ",") != ExistingOwner {
		t.Errorf("adopted=%v others=%v", r.Adopted, r.Others)
	}
	z := unzip(t, r.Zip)
	if z["assets/mypack/textures/x.png"] != "x" || z["pack.png"] != "icon" {
		t.Errorf("hand-made files lost: %s", names(z))
	}
	if s := z["assets/minecraft/lang/en_us.json"]; !strings.Contains(s, `"a"`) || !strings.Contains(s, `"b"`) {
		t.Errorf("lang: %s", s)
	}
}

func TestMergeWorldPackReplacesOlderMMSBuild(t *testing.T) {
	old, _ := Zip([]OutFile{
		meta("Embers"),
		{Path: "assets/embers/textures/old.png", Data: []byte("old")},
	})
	r := merge(t, old, "embers", meta("Embers"), OutFile{Path: "assets/embers/textures/new.png", Data: []byte("new")})
	if !r.Adopted || len(r.Others) != 0 {
		t.Errorf("adopted=%v others=%v", r.Adopted, r.Others)
	}
	if got := names(unzip(t, r.Zip)); got != "assets/embers/textures/new.png mms-world-pack.json pack.mcmeta" {
		t.Errorf("files: %s", got)
	}
}
