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
		{"object field", "Noise n {\n  |\n}", []string{"base_octave", "amplitude_modifiers"}, false},
		{"local ref", "Noise n { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction x = Noise(|)", []string{"n"}, false},
		{"comment", "// Noise |\nNoise n {}", nil, true},
		{"string", "Noise n { a = \"x|y\" }", nil, true},
		{"command name", "Function f {\n  /scoreb|\n}", []string{"scoreboard", "say"}, false},
		{"command literal", "Function f {\n  /scoreboard players |\n}", []string{"set", "operation"}, false},
		{"command selector", "Function f {\n  /kill |\n}", []string{"@e", "@s"}, false},
		{"command function", "Function f {\n  /function |\n}\nFunction g {\n  /say hi\n}", []string{"test:g", "test:f"}, false},
		{"command objective", "Objective kills: deathCount\nFunction f {\n  /scoreboard players set @s |\n}", []string{"kills"}, false},
		{"command after execute run", "Function f {\n  /execute as @a run gam|\n}", []string{"gamemode"}, false},
		{"command registry", "Function f {\n  /effect give @s |\n}", []string{"speed"}, false},
		{"resource pack kind", "BlockSt|", []string{"BlockState"}, false},
		{"block state property", "BlockState minecraft:oak_fence {\n  multipart = [{ when = {\n    |\n  }, apply = { model = minecraft:block/oak_fence_post } }]\n}", []string{"north", "waterlogged"}, false},
		{"execute helper", "Function f {\n  A|\n}", []string{"As", "At", "Positioned"}, false},
		{"empty statement", "Function f {\n  /say hi\n  |\n}", []string{"As", "Positioned", "let", "if"}, false},
		{"empty statement nested", "Function f {\n  As(@a) {\n    |\n  }\n}", []string{"At", "while"}, false},
		{"empty statement hook", "OnTick {\n  |\n}", []string{"As", "return"}, false},
		{"command slash", "Function f {\n  /|\n}", []string{"say", "scoreboard"}, false},
		{"slash in math", "DensityFunction x = 1 / C|", []string{"Clamp"}, false},
		{"after divide", "DensityFunction x = 1 / |", []string{"Clamp", "Noise"}, false},
		{"after times", "DensityFunction x = 1 * |", []string{"Clamp"}, false},
		{"after plus", "DensityFunction x = 1 + |", []string{"Clamp"}, false},
		{"execute chain", "Function f {\n  As(@a).|\n}", []string{"At", "Facing"}, false},
		{"execute chain prefix", "Function f {\n  As(@a).A| {\n    /say hi\n  }\n}", []string{"At", "Anchored"}, false},
		{"execute nested", "OnTick {\n  As(@a) {\n    P|\n  }\n}", []string{"Positioned"}, false},
		{"advancement trigger", "OnAdvancement a = Pla|", []string{"PlacedBlock"}, false},
		{"advancement condition", "OnAdvancement a = PlacedBlock(|) {\n  /say hi\n}", []string{"location"}, false},
		{"advancement body value", "Global held: int = 0\nOnAdvancement a = PlacedBlock() {\n  held = h|\n}", []string{"held"}, false},
		{"advancement body condition", "Global held: int = 0\nOnAdvancement a = PlacedBlock() {\n  if (h| == 1) {\n    /say hi\n  }\n}", []string{"held"}, false},
		{"hook kind", "OnFirst|", []string{"OnFirstJoin"}, false},
		{"model parent", "Model block/x { parent = minecraft:block/cube_a| }", []string{"block/cube_all"}, false},
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
	src := "Noise ridges { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction x = Noise(rid|ges)\nDensityFunction y = Noise(test:ridges)\n"
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
	s, _ := session(t, "Noise a { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction b = 1\n")
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

// Templates, Blueprints, Types, consts and their parameters are
// classified as written, not as they expand.
func TestSemanticTokensSource(t *testing.T) {
	src := "Type Size { w: int }\nconst amount: int = 3\nTemplate twice(x: int): int = x * 2\n" +
		"Blueprint Wand(name: string) { CustomItem ${self.path} { base = stick, name = name } }\n" +
		"Wand ruby { name = \"Ruby\" }\nDensityFunction d = twice(amount)\n"
	s, _ := session(t, src)
	toks := s.SemanticTokens("main.mms")
	r := []rune(src)
	at := func(sub string, nth int) (TokenType, int, bool) {
		off := -1
		for range nth + 1 {
			off += 1 + strings.Index(src[off+1:], sub)
		}
		for _, tk := range toks {
			if tk.Span.Start == len([]rune(src[:off])) && string(r[tk.Span.Start:tk.Span.End]) == sub {
				return tk.Type, tk.Mods, true
			}
		}
		return 0, 0, false
	}
	for _, c := range []struct {
		sub  string
		nth  int
		want TokenType
		mods int
	}{
		{"Size", 0, TokType, ModDeclaration},
		{"int", 0, TokType, 0},
		{"amount", 0, TokVariable, ModDeclaration},
		{"amount", 1, TokVariable, 0},
		{"twice", 0, TokFunction, ModDeclaration},
		{"x", 0, TokParameter, ModDeclaration},
		{"x", 1, TokParameter, 0},
		{"Wand", 0, TokType, ModDeclaration},
		{"self", 0, TokVariable, 0},
		{"stick", 0, TokEnumMember, 0},
		{"Wand", 1, TokType, 0},
		{"twice", 1, TokFunction, 0}, // the user's own Template is not the library's
	} {
		typ, mods, ok := at(c.sub, c.nth)
		if !ok || typ != c.want || mods != c.mods {
			t.Errorf("%s #%d: got %s/%d (found %v), want %s/%d", c.sub, c.nth, TokenTypes[typ], mods, ok, TokenTypes[c.want], c.mods)
		}
	}
	for i := 1; i < len(toks); i++ {
		if toks[i].Span.Start < toks[i-1].Span.End {
			t.Errorf("overlap at %v", toks[i].Span)
		}
	}
}

func TestResourcePackNavigation(t *testing.T) {
	src := "Model block/glow { parent = minecraft:block/cube_all }\nBlockState minecraft:glowstone { variants = { \"\" = [{ model = block/gl|ow, weight = 3 }] } }\n"
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

func TestExecuteHover(t *testing.T) {
	for _, tc := range []struct{ src, want string }{
		{"Function f {\n  A|s(@a).At(@s) {\n    /say hi\n  }\n}", "execute as"},
		{"Function f {\n  As(@a).A|t(@s) {\n    /say hi\n  }\n}", "execute at"},
		{"Function f {\n  /advancement grant @s only test:ki|lled\n}\nOnAdvancement killed = PlayerKilledEntity() {\n  /say hi\n}", "OnAdvancement test:killed"},
	} {
		s, off := session(t, tc.src)
		h, _ := s.Hover("main.mms", off)
		if !strings.Contains(h, tc.want) {
			t.Errorf("%q: hover = %q, want %q", tc.src, h, tc.want)
		}
	}
}

// Item ids and #item tags complete wherever a recipe takes an ingredient,
// including right after `=` with nothing typed yet and after `#`.
func TestRecipeIngredientCompletion(t *testing.T) {
	const result = ", { id = stick })"
	cases := []struct {
		name, src, replaced string
	}{
		{"shaped key empty", "Recipe r = Shaped([\"A\"], {\n  A = |\n}" + result, ""},
		{"shaped key after =", "Recipe r = Shaped([\"A\"], {\n  A =|\n}" + result, ""},
		{"shaped key inline", "Recipe r = Shaped([\"A\"], { A = stick, B = | }" + result, ""},
		{"shaped key prefix", "Recipe r = Shaped([\"A\"], {\n  A = amethyst|\n}" + result, "amethyst"},
		{"shaped key tag", "Recipe r = Shaped([\"A\"], {\n  A = #|\n}" + result, "#"},
		{"shaped key tag prefix", "Recipe r = Shaped([\"A\"], {\n  A = #minecraft:creeper|\n}" + result, "#minecraft:creeper"},
		{"object key empty", "Recipe r { type = crafting_shaped, pattern = [\"A\"], key = { A = | }, result = { id = stick } }", ""},
		{"shapeless empty", "Recipe r = Shapeless([|]" + result, ""},
		{"shapeless second", "Recipe r = Shapeless([stick, |]" + result, ""},
		{"shapeless tag", "Recipe r = Shapeless([#|]" + result, "#"},
		{"smelting empty", "Recipe r = Smelting(ingredient = |, result = { id = stick })", ""},
		{"smelting tag", "Recipe r = Smelting(ingredient = #|, result = { id = stick })", "#"},
		{"stonecutting", "Recipe r = Stonecutting(ingredient = |, result = { id = stick })", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s, off := session(t, tc.src)
			items, sp := s.Completion("main.mms", off)
			got := labels(items)
			if !has(got, "amethyst_shard", "#planks") {
				if len(got) > 20 {
					got = got[:20]
				}
				t.Fatalf("want item ids and tags, got %d items: %v", len(items), got)
			}
			src := []rune(strings.Replace(tc.src, "|", "", 1))
			if r := string(src[sp.Start:sp.End]); r != tc.replaced {
				t.Errorf("replaces %q, want %q", r, tc.replaced)
			}
		})
	}
}

// A string read as text markup is split into its marks, tags and
// ${...} reads; a mark left open is text.
func TestSemanticTokensMarkup(t *testing.T) {
	plain := "Advancement a { display = { title = \"**Ruby** 5 * 3\", description = \"d\" } }\n"
	ps, _ := session(t, plain)
	var pg []string
	for _, tk := range ps.SemanticTokens("main.mms") {
		pg = append(pg, TokenTypes[tk.Type]+":"+string([]rune(plain)[tk.Span.Start:tk.Span.End]))
	}
	if g := strings.Join(pg, " "); !strings.Contains(g, `string:" operator:** string:Ruby operator:** string: 5 * 3" property:description`) {
		t.Errorf("plain tokens: %s", g)
	}

	src := "const who = \"x\"\nAdvancement a { display = { title = \"**Ruby** {gold}gem ${who} 5 * 3\", description = \"d\", icon = { id = stone } }, criteria = { x = Tick() } }\n"
	s, _ := session(t, src)
	r := []rune(src)
	var got []string
	for _, tk := range s.SemanticTokens("main.mms") {
		if tk.Span.Start > len([]rune(src[:strings.Index(src, "title =")])) {
			got = append(got, TokenTypes[tk.Type]+":"+string(r[tk.Span.Start:tk.Span.End]))
		}
	}
	want := `string:" operator:** string:Ruby operator:** string:  keyword:{gold} string:gem  variable:${who} string: 5 * 3" property:description`
	if g := strings.Join(got, " "); !strings.Contains(g, want) {
		t.Errorf("tokens:\n%s\nwant:\n%s", g, want)
	}
}

// A Predicate named in a condition hovers, jumps and renames like any
// other reference to it.
func TestPredicateConditionReference(t *testing.T) {
	src := "Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })\nFunction f {\n  if (!snea|king) {\n    /say up\n  }\n}\n"
	s, off := session(t, src)
	if h, _ := s.Hover("main.mms", off); !strings.Contains(h, "sneaking") {
		t.Errorf("hover = %q", h)
	}
	if defs := s.Definition("main.mms", off); len(defs) != 1 || defs[0].Span.Start != strings.Index(src, "sneaking") {
		t.Errorf("definition = %+v", defs)
	}
	edits, msg := s.Rename("main.mms", off, "crouching")
	if msg != "" {
		t.Fatal(msg)
	}
	if len(edits) != 2 {
		t.Errorf("rename edits = %+v", edits)
	}
}

// A variable in an execute helper's coordinates, `x` or `~dx`, is the
// variable for hover and rename.
func TestExecuteCoordinateVariables(t *testing.T) {
	for _, src := range []string{
		"Function f {\n  let dx = 1\n  Positioned(d|x, 64, 0) {\n    /say hi\n  }\n}",
		"Function f {\n  let dx = 1\n  Positioned(~d|x ~ ~-dx) {\n    /say hi\n  }\n}",
	} {
		s, off := session(t, src)
		h, _ := s.Hover("main.mms", off)
		if !strings.Contains(h, "dx") || !strings.Contains(h, "int") {
			t.Errorf("%q: hover = %q", src, h)
		}
		edits, msg := s.Rename("main.mms", off, "step")
		if msg != "" {
			t.Errorf("%q: rename: %s", src, msg)
		}
		if want := strings.Count(strings.Replace(src, "|", "", 1), "dx"); len(edits) != want {
			t.Errorf("%q: rename edits %+v, want %d", src, edits, want)
		}
	}
}
