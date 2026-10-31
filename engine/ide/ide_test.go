package ide

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// session builds a one-file project from src, where "|" marks the cursor.
func session(t *testing.T, src string) (*Session, int) {
	t.Helper()
	off := strings.Index(src, "|")
	if off >= 0 {
		src = src[:off] + src[off+1:]
	}
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return &Session{P: p, R: r}, len([]rune(src[:max(off, 0)]))
}

func labels(items []Item) []string {
	var out []string
	for _, it := range items {
		out = append(out, it.Label)
	}
	return out
}

func has(xs []string, want ...string) bool {
	set := map[string]bool{}
	for _, x := range xs {
		set[x] = true
	}
	for _, w := range want {
		if !set[w] {
			return false
		}
	}
	return true
}

func TestCompletion(t *testing.T) {
	cases := []struct {
		name, src string
		want      []string
		none      bool
	}{
		{"decl kind", "Dens|", []string{"DensityFunction", "Noise"}, false},
		{"value", "DensityFunction x = |", []string{"Noise", "Clamp", "Cache"}, false},
		{"value prefix", "DensityFunction x = Cl|", []string{"Clamp"}, false},
		{"method", "DensityFunction x = Noise(minecraft:ridge).|", []string{"XzScale", "YScale"}, false},
		{"object field", "Noise n = {\n  |\n}", []string{"base_octave", "amplitude_modifiers"}, false},
		{"local ref", "Noise n = { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction x = Noise(|)", []string{"n"}, false},
		{"comment", "// Noise |\nNoise n = {}", nil, true},
		{"string", "Noise n = { a = \"x|y\" }", nil, true},
		{"command name", "Function f {\n  /scoreb|\n}", []string{"scoreboard", "say"}, false},
		{"command literal", "Function f {\n  /scoreboard players |\n}", []string{"set", "operation"}, false},
		{"command selector", "Function f {\n  /kill |\n}", []string{"@e", "@s"}, false},
		{"command function", "Function f {\n  /function |\n}\nFunction g {\n  /say hi\n}", []string{"test:g", "test:f"}, false},
		{"command objective", "Objective kills = deathCount\nFunction f {\n  /scoreboard players set @s |\n}", []string{"kills"}, false},
		{"command after execute run", "Function f {\n  /execute as @a run gam|\n}", []string{"gamemode"}, false},
		{"command registry", "Function f {\n  /effect give @s |\n}", []string{"speed"}, false},
		{"resource pack kind", "BlockSt|", []string{"BlockState"}, false},
		{"block state property", "BlockState minecraft:oak_fence = {\n  multipart = [{ when = {\n    |\n  }, apply = { model = minecraft:block/oak_fence_post } }]\n}", []string{"north", "waterlogged"}, false},
		{"model parent", "Model block/x = { parent = minecraft:block/cube_a| }", []string{"block/cube_all"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s, off := session(t, tc.src)
			items, _ := s.Completion("main.mms", off)
			got := labels(items)
			if tc.none {
				if len(got) != 0 {
					t.Fatalf("want no items, got %v", got)
				}
				return
			}
			if !has(got, tc.want...) {
				if len(got) > 40 {
					got = got[:40]
				}
				t.Fatalf("want %v among %v", tc.want, got)
			}
		})
	}
}

func TestHoverDefinitionRename(t *testing.T) {
	src := "Noise ridges = { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction x = Noise(rid|ges)\nDensityFunction y = Noise(test:ridges)\n"
	s, off := session(t, src)
	h, _ := s.Hover("main.mms", off)
	if !strings.Contains(h, "ridges") {
		t.Errorf("hover = %q", h)
	}
	defs := s.Definition("main.mms", off)
	if len(defs) != 1 || defs[0].Span.Start != strings.Index(src, "ridges") {
		t.Errorf("definition = %+v", defs)
	}
	refs := s.References("main.mms", off, true)
	if len(refs) != 3 {
		t.Errorf("references = %+v", refs)
	}
	edits, msg := s.Rename("main.mms", off, "hills")
	if msg != "" {
		t.Fatal(msg)
	}
	text := []rune(strings.Replace(src, "|", "", 1))
	sort.Slice(edits, func(i, j int) bool { return edits[i].Span.Start < edits[j].Span.Start })
	for i := len(edits) - 1; i >= 0; i-- {
		e := edits[i]
		text = append(text[:e.Span.Start], append([]rune(e.Text), text[e.Span.End:]...)...)
	}
	got := string(text)
	if strings.Contains(got, "ridges") || !strings.Contains(got, "test:hills") {
		t.Errorf("after rename:\n%s", got)
	}
}

func TestSymbols(t *testing.T) {
	s, _ := session(t, "Noise a = { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction b = 1\n")
	syms := s.Symbols("main.mms")
	if len(syms) != 2 {
		t.Fatalf("symbols = %+v", syms)
	}
}

func TestSemanticTokens(t *testing.T) {
	src := "// c\nDensityFunction x = Noise(minecraft:ridge, 1, 1).Abs()\n"
	s, _ := session(t, src)
	toks := s.SemanticTokens("main.mms")
	got := map[string]TokenType{}
	r := []rune(src)
	for _, tk := range toks {
		got[string(r[tk.Span.Start:tk.Span.End])] = tk.Type
	}
	want := map[string]TokenType{"// c": TokComment, "DensityFunction": TokType, "x": TokVariable, "Noise": TokFunction, "1": TokNumber}
	for text, typ := range want {
		if got[text] != typ {
			t.Errorf("%q: got %s, want %s", text, TokenTypes[got[text]], TokenTypes[typ])
		}
	}
	for i := 1; i < len(toks); i++ {
		if toks[i].Span.Start < toks[i-1].Span.End {
			t.Errorf("overlap at %v", toks[i].Span)
		}
	}
}

func TestResourcePackNavigation(t *testing.T) {
	src := "Model block/glow = { parent = minecraft:block/cube_all }\nBlockState minecraft:glowstone = { variants = { \"\" = [{ model = block/gl|ow, weight = 3 }] } }\n"
	s, off := session(t, src)
	defs := s.Definition("main.mms", off)
	if len(defs) != 1 || defs[0].Span.Start != strings.Index(src, "block/glow") {
		t.Errorf("definition = %+v", defs)
	}
	h, _ := s.Hover("main.mms", strings.Index(src, "BlockState")+2)
	if !strings.Contains(h, "assets/<namespace>/blockstates/<name>.json") || !strings.Contains(h, "resource pack") {
		t.Errorf("kind hover = %q", h)
	}
}
