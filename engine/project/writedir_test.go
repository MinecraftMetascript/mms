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
