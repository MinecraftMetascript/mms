package main

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestImportZip imports a zipped pack with a top-level folder, a second
// namespace and a function, then builds the project.
func TestImportZip(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{
		"ruby/pack.mcmeta":                                   `{"pack":{"pack_format":121,"description":"Ruby pack"}}`,
		"ruby/data/ruby/function/load.mcfunction":            "# Load\nsay hi\n",
		"ruby/data/ruby/tags/item/gems.json":                 `{"values":["minecraft:diamond"]}`,
		"ruby/data/ruby/loot_table/blocks/ruby_ore.json":     `{"type":"minecraft:block","pools":[{"rolls":1,"entries":[{"type":"minecraft:item","name":"minecraft:diamond"}]}]}`,
		"ruby/data/minecraft/tags/function/load.json":        `{"values":["ruby:load"]}`,
		"ruby/data/ruby/worldgen/density_function/hill.json": `{"type":"minecraft:mul","left":"minecraft:overworld/continents","right":0.5}`,
	}
	zpath := filepath.Join(dir, "ruby.zip")
	zf, _ := os.Create(zpath)
	zw := zip.NewWriter(zf)
	for name, text := range files {
		w, _ := zw.Create(name)
		w.Write([]byte(text))
	}
	zw.Close()
	zf.Close()

	proj := filepath.Join(dir, "proj")
	var out, errOut bytes.Buffer
	if code := runImport([]string{zpath, proj}, &out, &errOut); code != 0 {
		t.Fatalf("import exit %d: %s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "imported 5 resources") {
		t.Errorf("import output %q", out.String())
	}
	cfg, _ := os.ReadFile(filepath.Join(proj, "mms.json"))
	if !strings.Contains(string(cfg), `"namespace": "ruby"`) || !strings.Contains(string(cfg), "Ruby pack") {
		t.Errorf("mms.json = %s", cfg)
	}
	tag, _ := os.ReadFile(filepath.Join(proj, "minecraft", "tags_function.mms"))
	if !strings.HasPrefix(string(tag), "Namespace minecraft\n") {
		t.Errorf("minecraft tag file = %q", tag)
	}
	if code := runBuild([]string{proj}, &out, &errOut); code != 0 {
		t.Fatalf("build exit %d: %s", code, errOut.String())
	}
	got, err := os.ReadFile(filepath.Join(proj, "build", "data", "ruby", "function", "load.mcfunction"))
	if err != nil || !strings.Contains(string(got), "say hi") {
		t.Errorf("load.mcfunction = %q, %v", got, err)
	}
	if code := runImport([]string{zpath, proj}, &out, &errOut); code != 1 {
		t.Errorf("second import exit %d, want 1 (mms.json exists)", code)
	}
}
