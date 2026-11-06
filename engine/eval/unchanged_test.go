package eval

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TestPlainFilesPassThrough checks that files using none of the new
// forms come out as they went in: the same declarations, unchanged, so
// the checker keeps reusing its results.
func TestPlainFilesPassThrough(t *testing.T) {
	var paths []string
	for _, glob := range []string{"../../examples/*/*.mms", "../../examples/*/*/*.mms", "../sema/testdata/*.mms", "../project/testdata/*.mms"} {
		m, _ := filepath.Glob(glob)
		paths = append(paths, m...)
	}
	if len(paths) < 20 {
		t.Fatalf("found only %d files", len(paths))
	}
	var units []Unit
	for _, p := range paths {
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		units = append(units, Unit{Path: p, File: syntax.Parse(string(b))})
	}
	r := Run(units, Options{Namespace: "demo"})
	plain := 0
	for _, u := range units {
		out := r.Files[u.Path]
		if len(out.Diags) > 0 {
			t.Errorf("%s: %s", u.Path, out.Diags[0].Message)
		}
		// A top-level const, let or for is one of the new forms, so the
		// file is expected to change.
		if len(u.File.Items) != len(u.File.Decls) {
			continue
		}
		plain++
		if out.Changed {
			t.Errorf("%s: changed", u.Path)
		}
		if len(out.Decls) != len(u.File.Decls) {
			t.Errorf("%s: %d declarations, want %d", u.Path, len(out.Decls), len(u.File.Decls))
			continue
		}
		for i, d := range out.Decls {
			if d.D != u.File.Decls[i] {
				t.Errorf("%s: declaration %d was rebuilt", u.Path, i)
			}
		}
	}
	if plain < 20 {
		t.Errorf("only %d files use none of the new forms", plain)
	}
}

// FuzzRun checks that evaluation never panics on parsed input.
func FuzzRun(f *testing.F) {
	seeds, _ := filepath.Glob("testdata/*.txt")
	for _, s := range seeds {
		b, _ := os.ReadFile(s)
		f.Add(string(b))
	}
	f.Fuzz(func(t *testing.T, src string) {
		Run([]Unit{{Path: "a.mms", File: syntax.Parse(src)}}, Options{Namespace: "demo", Budget: 20000})
	})
}
