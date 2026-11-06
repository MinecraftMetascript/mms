package ide

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestEditorOnExpandedCode runs every editor feature at every offset of
// the design doc's snippets, which use const, for, Template and
// Blueprint, to check none of them fails on expanded declarations.
func TestEditorOnExpandedCode(t *testing.T) {
	paths, _ := filepath.Glob("../syntax/testdata/design/*.mms")
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	for _, path := range paths {
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		p.SetFile(filepath.Base(path), string(b))
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	s := &Session{P: p, R: r}
	for _, path := range paths {
		name := filepath.Base(path)
		n := len(r.Unit(name).File.Src)
		s.SemanticTokens(name)
		s.InlayHints(name)
		s.Symbols(name)
		r.Diagnostics(name)
		for off := 0; off <= n; off++ {
			s.Completion(name, off)
			s.Hover(name, off)
			s.Signature(name, off)
			s.Definition(name, off)
			s.References(name, off, true)
			s.PrepareRename(name, off)
		}
	}
}
