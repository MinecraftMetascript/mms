package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestLocalNames covers hover, go-to-definition, references and rename for
// the names a Template or Blueprint binds, and hover on type names.
func TestLocalNames(t *testing.T) {
	src := `/**
 * One item, rolled.
 * @param rolls how many times
 */
Template single(item: Id<item>, rolls: int = 1): LootPool = Pool(rolls, item)

Blueprint Wand {
  /** The wand's name. */
  name: string
  cooldown: int = 20
  emit {
    primary CustomItem ${self.path} {
      base = stick
      name = name
    }
    LootTable t { pools = [single(stick, cooldown)] }
  }
}

Wand w { name = "x" }

Function greet(who: string) {
  let greeting = "Hello, " + who
  tell(@a, "${greeting}")
}
`
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("main.mms", src)
	r, _ := p.Analyze(context.Background())
	s := &Session{P: p, R: r}
	// at is the rune offset of the n-th (from 1) sub in src, plus d.
	at := func(sub string, n, d int) int {
		i := -1
		for k := 0; k < n; k++ {
			i = i + 1 + strings.Index(src[i+1:], sub)
		}
		return len([]rune(src[:i])) + d
	}
	hover := func(off int) string { h, _ := s.Hover("main.mms", off); return h }

	rollsUse := at("rolls,", 1, 1)
	if h := hover(rollsUse); !strings.Contains(h, "parameter rolls: int = 1") || !strings.Contains(h, "how many times") {
		t.Errorf("hover on a parameter use = %q", h)
	}
	if d := s.Definition("main.mms", rollsUse); len(d) != 1 || d[0].Span.Start != at("rolls:", 1, 0) {
		t.Errorf("definition of a parameter use = %v", d)
	}
	nameUse := at("= name", 2, 3)
	if h := hover(nameUse); !strings.Contains(h, "field name: string") || !strings.Contains(h, "The wand's name.") {
		t.Errorf("hover on a field use = %q", h)
	}
	if refs := s.References("main.mms", nameUse, true); len(refs) != 2 {
		t.Errorf("references of a field = %v, want the field and its use", refs)
	}
	if _, ok := s.PrepareRename("main.mms", at("cooldown)", 1, 1)); ok {
		t.Errorf("a field can be renamed, but its instances would keep the old name")
	}
	letUse := at("{greeting}", 1, 2)
	if h := hover(letUse); !strings.Contains(h, "let greeting") {
		t.Errorf("hover on a let inside ${...} = %q", h)
	}
	if edits, msg := s.Rename("main.mms", letUse, "hello"); msg != "" || len(edits) != 2 {
		t.Errorf("rename of a let = %v %q, want two edits", edits, msg)
	}
	if h := hover(at("LootPool", 1, 1)); !strings.Contains(h, "`rolls`") {
		t.Errorf("hover on a schema type = %q", h)
	}
	if h := hover(at("Id<", 1, 1)); !strings.Contains(h, "Id<registry>") {
		t.Errorf("hover on Id = %q", h)
	}
}
