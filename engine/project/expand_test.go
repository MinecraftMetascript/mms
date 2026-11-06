package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestExpandEndToEnd checks that const, a top-level for and a Template
// reach the build, and that a problem inside a Template is reported at
// the call the user wrote.
func TestExpandEndToEnd(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("lib.mms", "const woods = [oak, birch]\nTemplate Drop(item: string, n: int = 1) = Pool(n, ${item})\n")
	p.SetFile("a.mms", `for w in woods {
  LootTable blocks/${w}_crate { pools = [Drop("${w}_planks", 2)] }
}
LootTable bad { pools = [Drop()] }
`)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var msgs []string
	for _, d := range r.Diagnostics("a.mms") {
		msgs = append(msgs, d.Code+" "+d.Message)
		if d.Code == "mms/missing-arg" {
			if got := string(r.Unit("a.mms").File.Src[d.Sp.Start:d.Sp.End]); got != "Drop()" {
				t.Errorf("missing-arg reported at %q, want the call", got)
			}
		}
	}
	if len(msgs) != 1 || !strings.Contains(msgs[0], "needs 'item'") {
		t.Fatalf("diagnostics: %q", msgs)
	}
	files := map[string]string{}
	for _, f := range r.Emit() {
		files[f.Path] = string(f.Data)
	}
	for _, w := range []string{"oak", "birch"} {
		got, ok := files["data/demo/loot_table/blocks/"+w+"_crate.json"]
		if !ok {
			t.Fatalf("no loot table for %s; files: %v", w, keys(files))
		}
		if !strings.Contains(got, `"minecraft:`+w+`_planks"`) || !strings.Contains(got, `"rolls": 2`) {
			t.Errorf("%s crate:\n%s", w, got)
		}
	}
	if r.Eval.Files["lib.mms"].Changed != true {
		t.Error("lib.mms declares a Template, so its declarations change")
	}
}

func keys(m map[string]string) []string {
	var out []string
	for k := range m {
		out = append(out, k)
	}
	return out
}

// TestConversionEdit checks that a use of a Blueprint instance in another
// file follows an edit to the instance: the cached check of the user
// must not survive the change of what its name converts to.
func TestConversionEdit(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("gem.mms", `Blueprint Gem {
  base: Id<item>
  as Id<item> = base hint("gem-base", "Matches any ${base}")
  emit { Function gems/${self.path} { /say ${base} } }
}
Gem ruby { base = emerald }
`)
	p.SetFile("tag.mms", "ItemTag shiny = [ruby]\n")
	hint := func() string {
		r, err := p.Analyze(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		var msgs []string
		for _, d := range r.Diagnostics("tag.mms") {
			msgs = append(msgs, d.Code+" "+d.Message)
		}
		return strings.Join(msgs, "; ")
	}
	if got := hint(); got != "gem-base Matches any emerald" {
		t.Fatalf("before the edit: %q", got)
	}
	p.SetFile("gem.mms", strings.Replace(p.texts["gem.mms"], "emerald", "diamond", 1))
	if got := hint(); got != "gem-base Matches any diamond" {
		t.Fatalf("after the edit: %q", got)
	}
}
