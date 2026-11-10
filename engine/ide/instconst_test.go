package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

var instConstFiles = map[string]string{
	"biomes.mms": `Blueprint Pair {
  v: int
  emit {
    /** Twice v. */
    const twice: int = v * 2
    Objective ${self.path}: dummy
  }
}

Pair alpha { v = 3 }
`,
	"use.mms": `const z: int = alpha.twice
LootTable t { pools = [Pool(alpha.twice, stone)] }
const w = alpha.
`,
}

func instConstSession(t *testing.T) *Session {
	t.Helper()
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	for path, src := range instConstFiles {
		p.SetFile(path, src)
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return &Session{P: p, R: r}
}

func instConstOff(path, at string) int {
	return len([]rune(instConstFiles[path][:strings.Index(instConstFiles[path], at)]))
}

// name.const leads to the const in the Blueprint, and name to the
// instance.
func TestInstanceConstNavigation(t *testing.T) {
	s := instConstSession(t)
	defs := s.Definition("use.mms", instConstOff("use.mms", "twice\n")+1)
	if len(defs) != 1 || defs[0].Path != "biomes.mms" || defs[0].Span.Start != instConstOff("biomes.mms", "twice:") {
		t.Errorf("definition of alpha.twice = %v", defs)
	}
	defs = s.Definition("use.mms", instConstOff("use.mms", "alpha.twice,")+1)
	if len(defs) != 1 || defs[0].Span.Start != instConstOff("biomes.mms", "alpha {") {
		t.Errorf("definition of alpha = %v", defs)
	}
	md, _ := s.Hover("use.mms", instConstOff("use.mms", "twice\n")+1)
	if !strings.Contains(md, "const twice: int = v * 2") || !strings.Contains(md, "Emitted by Pair alpha") || !strings.Contains(md, "Twice v.") {
		t.Errorf("hover of alpha.twice = %q", md)
	}
}

func TestInstanceConstCompletion(t *testing.T) {
	s := instConstSession(t)
	items, _ := s.Completion("use.mms", instConstOff("use.mms", "alpha.\n")+len("alpha."))
	if len(items) != 1 || items[0].Label != "twice" {
		var labels []string
		for _, it := range items {
			labels = append(labels, it.Label)
		}
		t.Errorf("alpha.| offers %v, want [twice]", labels)
	}
}
