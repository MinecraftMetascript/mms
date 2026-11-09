package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// selectorErrors analyzes src and returns its errors.
func selectorErrors(src string) []syntax.Diagnostic {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("main.mms", src)
	r, _ := p.Analyze(context.Background())
	var out []syntax.Diagnostic
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			out = append(out, d)
		}
	}
	return out
}

// A selector const or Template is a selector wherever one goes, a raw
// command's ${...} among them.
func TestNamedSelectors(t *testing.T) {
	src := `Objective hp
const boss = @e[tag=siege.boss, limit=1]
const players: Selector = @a[tag=siege.player]
Template near(r: int): Selector = @a[distance=..${r}]
Function f {
  /kill ${boss}
  /execute as ${players} at @s run say hi
  /tp ${near(4)} ~ ~1 ~
  tell(players, "hello ${boss}")
  hp[boss] = 3
  As(players).At(boss) {
    /say x
  }
  if (exists(near(8))) {
    /say y
  }
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	got := emitted(t, r, "data/demo/function/f.mcfunction")
	for _, want := range []string{
		"kill @e[tag=siege.boss,limit=1]",
		"execute as @a[tag=siege.player] at @s run say hi",
		"tp @a[distance=..4] ~ ~1 ~",
		`tellraw @a[tag=siege.player] ["","hello ",{selector:"@e[tag=siege.boss,limit=1]"}]`,
		"scoreboard players set @e[tag=siege.boss,limit=1] hp 3",
		"execute as @a[tag=siege.player] at @e[tag=siege.boss,limit=1] run say x",
		"execute if entity @a[distance=..8] run say y",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output lacks %s:\n%s", want, got)
		}
	}
}

// Selectors are checked where they are written, whatever holds them.
func TestSelectorErrors(t *testing.T) {
	cases := []struct{ src, code, at string }{
		{"const boss = @e[tagg=x]", "mms/selector", "tagg"},
		{"Function f {\n  tell(@q, \"hi\")\n}", "mms/selector", "@q"},
		{"Template t(): Selector = @a[bogus=1]", "mms/selector", "bogus"},
		{"const x: Selector = 5", "mms/type-mismatch", "5"},
		{"Function f {\n  As(@e[tag=a, distanse=..3]) {\n    /say hi\n  }\n}", "mms/selector", "distanse"},
	}
	for _, c := range cases {
		ds := selectorErrors(c.src)
		if len(ds) != 1 {
			t.Errorf("%q: %d problems %v, want one", c.src, len(ds), ds)
			continue
		}
		d := ds[0]
		if d.Code != c.code || c.src[d.Sp.Start:d.Sp.End] != c.at {
			t.Errorf("%q: %s at %q (%s), want %s at %q", c.src, d.Code, c.src[d.Sp.Start:d.Sp.End], d.Message, c.code, c.at)
		}
	}
}

// An unknown option offers the closest known one.
func TestSelectorOptionFix(t *testing.T) {
	ds := selectorErrors("const boss = @e[tagg=x]")
	if len(ds) != 1 || len(ds[0].Fixes) != 1 || ds[0].Fixes[0].Edits[0].Text != "tag" {
		t.Fatalf("got %+v, want one fix to tag", ds)
	}
}
