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
	// A declaration of a project Blueprint's kind is one of the new forms.
	blueprints := map[string]bool{}
	for _, u := range units {
		for _, d := range u.File.Decls {
			if d.Kind != nil && d.Kind.Name == "Blueprint" {
				blueprints[d.NameText()] = true
			}
		}
	}
	plain := 0
	for _, u := range units {
		out := r.Files[u.Path]
		// The prelude, which declares helpers such as len and count, is
		// not loaded here.
		if callsPrelude(u.File) {
			continue
		}
		if len(out.Diags) > 0 {
			t.Errorf("%s: %s", u.Path, out.Diags[0].Message)
		}
		// A top-level const, let or for, or a Type, Template or Blueprint
		// declaration, is one of the new forms, so the file is expected
		// to change.
		if len(u.File.Items) != len(u.File.Decls) || declaresForms(u.File) || usesKinds(u.File, blueprints) {
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

// callsPrelude reports whether f calls an intrinsic the prelude
// declares, such as len, or count in a Test.
func callsPrelude(f *syntax.File) bool {
	found := false
	syntax.Walk(f, func(n syntax.Node) bool {
		if c, ok := n.(*syntax.Call); ok {
			if id, ok := c.Fn.(*syntax.Ident); ok && (id.Name == "count" || id.Name == "exists" || id.Name == "len") {
				found = true
			}
		}
		return !found
	})
	return found
}

func usesKinds(f *syntax.File, kinds map[string]bool) bool {
	for _, d := range f.Decls {
		if d.Kind != nil && kinds[d.Kind.Name] {
			return true
		}
	}
	return false
}

func declaresForms(f *syntax.File) bool {
	for _, d := range f.Decls {
		if d.Kind != nil && (d.Kind.Name == "Type" || d.Kind.Name == "Template" || d.Kind.Name == "Blueprint") {
			return true
		}
	}
	return false
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
