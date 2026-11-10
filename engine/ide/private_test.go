package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// files: a.mms and b.mms each keep a private const gap, and lib.mms has a
// public one, which c.mms uses.
var privateFiles = map[string]string{
	"lib.mms": "const gap: int = 1\nconst shared: int = 5\n",
	"a.mms":   "private const gap: int = 2\nprivate Template twice(x: int): int = x * 2\nLootTable a { pools = [Pool(twice(gap), stone)] }\n",
	"b.mms":   "private const gap: int = 3\nLootTable b { pools = [Pool(gap, stone)] }\n",
	"c.mms":   "LootTable c { pools = [Pool(gap, stone)] }\n",
}

func privateSession(t *testing.T) *Session {
	t.Helper()
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	for path, src := range privateFiles {
		p.SetFile(path, src)
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return &Session{P: p, R: r}
}

func offIn(path, at string) int {
	return len([]rune(privateFiles[path][:strings.Index(privateFiles[path], at)]))
}

// Go to definition, references and rename follow the name a file means.
func TestPrivateNavigation(t *testing.T) {
	s := privateSession(t)
	for _, c := range []struct{ path, at, defPath string }{
		{"a.mms", "twice(gap)", "a.mms"},
		{"a.mms", "gap)", "a.mms"},
		{"b.mms", "gap, stone", "b.mms"},
		{"c.mms", "gap, stone", "lib.mms"},
	} {
		defs := s.Definition(c.path, offIn(c.path, c.at)+1)
		if len(defs) != 1 || defs[0].Path != c.defPath {
			t.Errorf("%s %q: definition = %v, want in %s", c.path, c.at, defs, c.defPath)
		}
	}
	refs := s.References("b.mms", offIn("b.mms", "gap: int = 3")+1, true)
	if len(refs) != 2 {
		t.Errorf("b's private gap: references = %v", refs)
	}
	for _, r := range refs {
		if r.Path != "b.mms" {
			t.Errorf("b's private gap found in %s", r.Path)
		}
	}
	refs = s.References("lib.mms", offIn("lib.mms", "gap: int = 1")+1, true)
	paths := map[string]bool{}
	for _, r := range refs {
		paths[r.Path] = true
	}
	if !paths["lib.mms"] || !paths["c.mms"] || paths["a.mms"] || paths["b.mms"] {
		t.Errorf("lib's gap: references in %v", paths)
	}
	edits, msg := s.Rename("b.mms", offIn("b.mms", "gap: int = 3")+1, "space")
	if msg != "" || len(edits) != 2 {
		t.Errorf("rename b's gap: %v %q", edits, msg)
	}
	for _, e := range edits {
		if e.Path != "b.mms" {
			t.Errorf("rename b's gap edits %s", e.Path)
		}
	}
	// A private name may take a name another file keeps private, but not
	// one its own file can already name.
	if _, msg := s.Rename("b.mms", offIn("b.mms", "gap: int = 3")+1, "twice"); msg != "" {
		t.Errorf("rename to another file's private name: %q", msg)
	}
	if _, msg := s.Rename("b.mms", offIn("b.mms", "gap: int = 3")+1, "shared"); msg == "" {
		t.Error("rename onto a public name the file can see")
	}
	if hover, _ := s.Hover("b.mms", offIn("b.mms", "gap, stone")+1); !strings.Contains(hover, "private const gap") || !strings.Contains(hover, "= 3") {
		t.Errorf("hover = %q", hover)
	}
}

// Completion offers a private name only in its own file.
func TestPrivateCompletion(t *testing.T) {
	s := privateSession(t)
	at := func(path string) []string {
		items, _ := s.Completion(path, offIn(path, "gap, stone"))
		return labels(items)
	}
	if got := at("c.mms"); has(got, "twice") || !has(got, "gap", "shared") {
		t.Errorf("c.mms offers %v", got)
	}
	items, _ := s.Completion("a.mms", offIn("a.mms", "twice(gap)"))
	if got := labels(items); !has(got, "twice", "gap") {
		t.Errorf("a.mms offers %v", got)
	}
}
