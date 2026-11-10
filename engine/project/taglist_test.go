package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// tagListFix applies the tag-list fix of the first "a tag is not allowed
// here" error in src, or returns ok false when it has none.
func tagListFix(t *testing.T, src string) (got string, f syntax.Fix, ok bool) {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", src)
	r, _ := p.Analyze(context.Background())
	for _, d := range r.Diagnostics("a.mms") {
		if !strings.HasPrefix(d.Message, "a tag is not allowed here") {
			continue
		}
		for _, f := range d.Fixes {
			if strings.HasPrefix(f.Title, "Move the list into ") {
				return applyEdits(src, append([]syntax.TextEdit(nil), f.Edits...)), f, true
			}
		}
		return "", syntax.Fix{}, false
	}
	t.Fatalf("%q: no tag error", src)
	return "", syntax.Fix{}, false
}

const tagListJigsaw = `/** Houses on the plains. */
JigsawStructure village {
  biomes = [#minecraft:is_forest, desert, savanna]
  step = surface_structures
  spacing = 24
  separation = 8
  height = Constant(Absolute(0))
  start = start
  pools = { start = [] }
}
`

func TestTagListFix(t *testing.T) {
	got, f, ok := tagListFix(t, tagListJigsaw)
	if !ok {
		t.Fatal("no fix")
	}
	want := "BiomeTag village_biomes = [#minecraft:is_forest, desert, savanna]\n\n" +
		strings.Replace(tagListJigsaw, "[#minecraft:is_forest, desert, savanna]", "#village_biomes", 1)
	if got != want {
		t.Errorf("fix:\n got %q\nwant %q", got, want)
	}
	if f.Title != "Move the list into BiomeTag village_biomes" || f.Safe {
		t.Errorf("fix = %q, safe %v", f.Title, f.Safe)
	}
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", got)
	r, _ := p.Analyze(context.Background())
	for _, d := range r.Diagnostics("a.mms") {
		if d.Severity == syntax.SeverityError {
			t.Errorf("after the fix: %s", d.Message)
		}
	}

	// A name already taken gets a number.
	taken := "BiomeTag village_biomes = [desert]\n\n" + tagListJigsaw
	got, _, _ = tagListFix(t, taken)
	if !strings.Contains(got, "BiomeTag village_biomes_2 = [#minecraft:is_forest, desert, savanna]\n\n/** Houses") ||
		!strings.Contains(got, "biomes = #village_biomes_2\n") {
		t.Errorf("taken name: %q", got)
	}

	// A Structure written out in full has the same field.
	structure := "Structure s {\n  type = jigsaw\n  biomes = [#minecraft:is_ocean, plains]\n}\n"
	got, _, ok = tagListFix(t, structure)
	if !ok || !strings.HasPrefix(got, "BiomeTag s_biomes = [#minecraft:is_ocean, plains]\n\nStructure s {\n  type = jigsaw\n  biomes = #s_biomes\n") {
		t.Errorf("Structure: %v %q", ok, got)
	}

	// A list over several lines keeps its lines, moved to the left.
	multi := "// Comment.\nStructure s {\n  type = jigsaw\n  biomes = [\n    #minecraft:is_ocean,\n    plains,\n  ]\n}\n"
	got, _, _ = tagListFix(t, multi)
	if want := "BiomeTag s_biomes = [\n  #minecraft:is_ocean,\n  plains,\n]\n\n// Comment.\nStructure s {\n  type = jigsaw\n  biomes = #s_biomes\n}\n"; got != want {
		t.Errorf("multi-line:\n got %q\nwant %q", got, want)
	}
}
