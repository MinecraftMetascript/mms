package ide

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

const evalNamesSrc = "Type Size { w: int }\n" +
	"const amount: int = 3\n" +
	"Template twice(x: int): int = x * 2\n" +
	"Template shadow(amount: int): int = amount\n" +
	"Blueprint Wand(name: string) { CustomItem ${self.path} { base = stick, name = name } }\n" +
	"Wand ruby { name = \"Ruby ${amount}\" }\n" +
	"DensityFunction d = twice(amount)\n" +
	"Template big(): Size = { w = amount }\n" +
	"Function f { /say ${amount} }\n"

// apply applies edits to src, last first.
func apply(src string, edits []Edit) string {
	r := []rune(src)
	sort.Slice(edits, func(i, j int) bool { return edits[i].Span.Start > edits[j].Span.Start })
	for _, e := range edits {
		r = append(r[:e.Span.Start], append([]rune(e.Text), r[e.Span.End:]...)...)
	}
	return string(r)
}

// Consts, Templates, Blueprints and Types rename everywhere they are
// written, ${...} pieces included, but not where a parameter of the same
// name hides them.
func TestRenameEvalNames(t *testing.T) {
	s, _ := session(t, evalNamesSrc)
	for _, c := range []struct{ at, to, want string }{
		{"amount: int = 3", "count2", strings.NewReplacer(
			"const amount", "const count2", "Ruby ${amount}", "Ruby ${count2}", "twice(amount)", "twice(count2)",
			"w = amount", "w = count2", "/say ${amount}", "/say ${count2}").Replace(evalNamesSrc)},
		{"twice(amount)", "double", strings.NewReplacer("Template twice", "Template double", "twice(amount)", "double(amount)").Replace(evalNamesSrc)},
		{"Wand ruby", "Staff", strings.NewReplacer("Blueprint Wand", "Blueprint Staff", "Wand ruby", "Staff ruby").Replace(evalNamesSrc)},
		{"Size =", "Dims", strings.NewReplacer("Type Size", "Type Dims", "): Size =", "): Dims =").Replace(evalNamesSrc)},
	} {
		off := strings.Index(evalNamesSrc, c.at) + 1
		if _, ok := s.PrepareRename("main.mms", off); !ok {
			t.Errorf("%s: cannot rename", c.at)
			continue
		}
		edits, msg := s.Rename("main.mms", off, c.to)
		if msg != "" {
			t.Errorf("%s: %s", c.at, msg)
			continue
		}
		if got := apply(evalNamesSrc, edits); got != c.want {
			t.Errorf("%s -> %s:\n%s", c.at, c.to, got)
		}
		if defs := s.Definition("main.mms", off); len(defs) != 1 {
			t.Errorf("%s: definition = %v", c.at, defs)
		}
	}
	// The parameter amount is not the const: its references are the
	// parameter's.
	constAt := strings.Index(evalNamesSrc, "const amount") + len("const ")
	for _, r := range s.References("main.mms", strings.Index(evalNamesSrc, "= amount\n")+3, true) {
		if r.Span.Start == constAt {
			t.Errorf("shadowed use finds the const: %v", r)
		}
	}
	off := strings.Index(evalNamesSrc, "twice(amount)") + 1
	if _, msg := s.Rename("main.mms", off, "amount"); msg == "" {
		t.Error("rename onto a taken name")
	}
	if _, msg := s.Rename("main.mms", off, "2x"); msg == "" {
		t.Error("rename to a bad name")
	}
}

// A prelude Template has no definition in the project and cannot be
// renamed.
func TestRenamePreludeTemplate(t *testing.T) {
	src := "LootTable l { pools = [Pool(1, [diamond])] }\n"
	s, _ := session(t, src)
	off := strings.Index(src, "Pool(") + 1
	if _, ok := s.PrepareRename("main.mms", off); ok {
		t.Error("prelude Template is renamable")
	}
}

// A call of a Template has signature help and its name a hover, though
// expansion replaced the call.
func TestTemplateSignature(t *testing.T) {
	src := "/** Doubles.\n@param x what to double */\nTemplate twice(x: int, y: int = 2): int = x * y\n" +
		"const amount = 4\nDensityFunction d = twice(3, amount)\n"
	s, _ := session(t, src)
	h, ok := s.Signature("main.mms", strings.Index(src, "3, ")+3)
	if !ok || h.Label != "twice(x: int, y: int = 2): int" || h.Active != 1 || h.Doc != "Doubles." || h.ParamDocs[0] != "what to double" {
		t.Errorf("signature = %+v", h)
	}
	md, _ := s.Hover("main.mms", strings.Index(src, "twice(3")+1)
	if !strings.Contains(md, "Template twice(x: int, y: int = 2): int") || !strings.Contains(md, "Doubles.") {
		t.Errorf("template hover = %q", md)
	}
	md, _ = s.Hover("main.mms", strings.LastIndex(src, "amount")+1)
	if !strings.Contains(md, "const amount: int = 4") {
		t.Errorf("const hover = %q", md)
	}
}

// labelsAt completes at the ^ in src, whose other text may hold `|`.
func labelsAt(t *testing.T, src string) []string {
	t.Helper()
	i := strings.Index(src, "^")
	src = src[:i] + src[i+1:]
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	s := &Session{P: p, R: r}
	items, _ := s.Completion("main.mms", len([]rune(src[:i])))
	var out []string
	for _, it := range items {
		out = append(out, it.Label)
	}
	return out
}

// A Template argument completes from the parameter's type, then the
// consts and the value Templates whose result fits it; the top level
// offers the new keywords.
func TestEvalCompletion(t *testing.T) {
	src := "const amount = 3\nTemplate pick(mode: \"rigid\" | \"terrain_matching\", n: int): int = n\nDensityFunction d = pick(^)\n"
	got := strings.Join(labelsAt(t, src), " ")
	for _, want := range []string{"rigid", "terrain_matching", "amount"} {
		if !strings.Contains(" "+got+" ", " "+want+" ") {
			t.Errorf("pick(|) lacks %s: %s", want, got)
		}
	}
	for _, unwanted := range []string{"Clamp", "pick", "cos"} {
		if strings.Contains(" "+got+" ", " "+unwanted+" ") {
			t.Errorf("pick(|) offers %s, which is not a mode: %s", unwanted, got)
		}
	}
	got = strings.Join(labelsAt(t, strings.Replace(src, "pick(^)", "pick(\"rigid\", ^)", 1)), " ")
	for _, want := range []string{"amount", "pick", "len"} {
		if !strings.Contains(" "+got+" ", " "+want+" ") {
			t.Errorf("pick(\"rigid\", |) lacks %s: %s", want, got)
		}
	}
	got = strings.Join(labelsAt(t, "Te^"), " ")
	for _, want := range []string{"const", "Template", "Type", "Blueprint", "for", "CustomItem"} {
		if !strings.Contains(" "+got+" ", " "+want+" ") {
			t.Errorf("top level lacks %s", want)
		}
	}
}

// The outline lists consts, Templates, Types and Blueprints too, and a
// Blueprint instance as written rather than what it emits.
func TestEvalSymbols(t *testing.T) {
	s, _ := session(t, evalNamesSrc)
	var got []string
	for _, sym := range s.Symbols("main.mms") {
		got = append(got, sym.Kind+":"+sym.Name)
	}
	want := "type:Size constant:amount template:twice template:shadow blueprint:Wand resource:ruby"
	if !strings.HasPrefix(strings.Join(got, " "), want) {
		t.Errorf("symbols = %v", got)
	}
}

// A helper's named arguments include the fields a secondary key brings in:
// a multi-noise source without a preset lists its biomes, also while the
// first argument is still being typed.
func TestCompletionSecondaryDispatchFields(t *testing.T) {
	for _, src := range []string{
		"Dimension d {\n  type = overworld\n  generator = Noise(settings = overworld, biome_source = MultiNoise(^))\n}\n",
		"Dimension d {\n  type = overworld\n  generator = Noise(settings = overworld, biome_source = MultiNoise(b^))\n}\n",
	} {
		got := labelsAt(t, src)
		if !slices.Contains(got, "biomes") || !slices.Contains(got, "preset") {
			t.Errorf("%q offers %v", src, got)
		}
	}
}

// A field of an object the evaluator worked out completes where it is
// written: its key keeps its own place, not the whole object's. The
// cave-siege world preset, whose overworld uses a const generator, once
// offered the dimensions map's keys there.
func TestCompletionInEvaluatedObject(t *testing.T) {
	dir := "../../examples/cave-siege"
	cfg, err := project.LoadConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	paths, err := cfg.SourceFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	p := project.New(spec.MustDefault(), cfg)
	p.Files = os.DirFS(dir)
	for _, rel := range paths {
		b, err := os.ReadFile(filepath.Join(dir, rel))
		if err != nil {
			t.Fatal(err)
		}
		p.SetFile(rel, string(b))
	}
	b, _ := os.ReadFile(filepath.Join(dir, "worldgen", "dimension.mms"))
	src := string(b)
	at := strings.Index(src, "      type = the_nether") + len("      t")
	if at < len("      t") {
		t.Fatal("no the_nether entry in the cave-siege world preset")
	}
	p.SetFile("worldgen/dimension.mms", src[:at]+src[at+len("ype"):])
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	items, _ := (&Session{P: p, R: r}).Completion("worldgen/dimension.mms", at)
	if got := labels(items); !slices.Contains(got, "type") {
		t.Errorf("offers %v, want type", got)
	}
}
