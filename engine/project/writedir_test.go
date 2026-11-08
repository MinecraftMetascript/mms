package project

import (
	"os"
	"path/filepath"
	"testing"
)

// TestWriteDirRefusesEscapes checks the last guard before disk: a pack
// path that leaves the pack folder is refused before anything is removed.
func TestWriteDirRefusesEscapes(t *testing.T) {
	for _, p := range []string{"../evil.txt", "data/../../evil.txt", "/abs.txt", "data//x.json", `data\..\..\evil.txt`} {
		dir := t.TempDir()
		keep := filepath.Join(dir, "data", "keep.txt")
		if err := os.MkdirAll(filepath.Dir(keep), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(keep, nil, 0o644); err != nil {
			t.Fatal(err)
		}
		if err := WriteDir(dir, []OutFile{{Path: p, Data: []byte("x")}}); err == nil {
			t.Errorf("WriteDir accepted %q", p)
		}
		if _, err := os.Stat(keep); err != nil {
			t.Errorf("%q: cleanup ran before the check: %v", p, err)
		}
		if _, err := Zip([]OutFile{{Path: p}}); err == nil {
			t.Errorf("Zip accepted %q", p)
		}
	}
}

// TestWriteDirReplacesSymlinks checks that a file kept from the last
// build that is now a symlink is replaced, not written through.
func TestWriteDirReplacesSymlinks(t *testing.T) {
	dir, outside := t.TempDir(), t.TempDir()
	target := filepath.Join(outside, "victim.txt")
	if err := os.WriteFile(target, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "pack.mcmeta")
	if err := os.Symlink(target, link); err != nil {
		t.Skip("symlinks not supported:", err)
	}
	if err := WriteDir(dir, []OutFile{{Path: "pack.mcmeta", Data: []byte("{}")}}); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(target); string(b) != "keep" {
		t.Errorf("wrote through the symlink: target now %q", b)
	}
	if fi, err := os.Lstat(link); err != nil || fi.Mode()&os.ModeSymlink != 0 {
		t.Errorf("pack.mcmeta is still a symlink (err %v)", err)
	}
}
