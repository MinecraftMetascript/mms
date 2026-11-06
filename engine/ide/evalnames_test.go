package ide

import (
	"sort"
	"strings"
	"testing"
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
	// The parameter amount is not the const.
	if refs := s.References("main.mms", strings.Index(evalNamesSrc, "= amount\n")+3, true); len(refs) != 0 {
		t.Errorf("shadowed use = %v", refs)
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
