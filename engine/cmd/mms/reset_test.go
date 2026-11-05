package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// copyExample copies an example project into a temporary folder.
func copyExample(t *testing.T, name string) string {
	t.Helper()
	dir := t.TempDir()
	src := filepath.Join("..", "..", "..", "examples", name)
	entries, err := os.ReadDir(src)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		b, err := os.ReadFile(filepath.Join(src, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		os.WriteFile(filepath.Join(dir, e.Name()), b, 0o644)
	}
	return dir
}

// TestBuildResetDimensions checks that -reset-dimensions moves a
// dimension's chunks out only when its worldgen changed, or when named.
func TestBuildResetDimensions(t *testing.T) {
	dir, world := copyExample(t, "sky-islands"), t.TempDir()
	os.WriteFile(filepath.Join(world, "level.dat"), nil, 0o644)
	region := filepath.Join(world, "dimensions", "sky", "islands", "region")
	build := func(args ...string) string {
		t.Helper()
		var out, errOut bytes.Buffer
		if code := runBuild(append(append([]string{"-world", world}, args...), dir), &out, &errOut); code != 0 {
			t.Fatalf("exit %d: %s", code, errOut.String())
		}
		return out.String()
	}
	build()
	os.MkdirAll(region, 0o755)
	if out := build("-reset-dimensions"); !strings.Contains(out, "reset: no dimension's worldgen changed") {
		t.Errorf("unchanged build:\n%s", out)
	}
	if _, err := os.Stat(region); err != nil {
		t.Fatal("an unchanged build reset the dimension")
	}
	terrain := filepath.Join(dir, "terrain.mms")
	b, _ := os.ReadFile(terrain)
	os.WriteFile(terrain, bytes.Replace(b, []byte("base_octave = -6"), []byte("base_octave = -5"), 1), 0o644)
	if out := build("-reset-dimensions"); !strings.Contains(out, "reset sky:islands: it generates again when visited") {
		t.Errorf("changed build:\n%s", out)
	}
	if _, err := os.Stat(region); !os.IsNotExist(err) {
		t.Error("the region folder is still there")
	}
	if _, err := os.Stat(filepath.Join(world, "mms-reset", "dimensions", "sky", "islands", "region")); err != nil {
		t.Error("the old chunks were not kept in mms-reset")
	}
	os.MkdirAll(region, 0o755)
	if out := build("-reset-dimensions=sky:islands"); !strings.Contains(out, "reset sky:islands:") {
		t.Errorf("named reset:\n%s", out)
	}

	var out, errOut bytes.Buffer
	if code := runBuild([]string{"-reset-dimensions", dir}, &out, &errOut); code != 2 || !strings.Contains(errOut.String(), "needs -world") {
		t.Errorf("without -world: exit %d, %s", code, errOut.String())
	}
	if code := runBuild([]string{"-world", world, "-reset-dimensions=sky:../x", dir}, &out, &errOut); code != 2 {
		t.Errorf("a bad id should be a usage error, got exit %d", code)
	}
}
