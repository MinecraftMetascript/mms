package syntax

import (
	"os"
	"testing"
)

func TestParseSketch(t *testing.T) {
	src, err := os.ReadFile("testdata/sketch.mms")
	if err != nil {
		t.Fatal(err)
	}
	f := Parse(string(src))
	for _, d := range f.Diags {
		t.Errorf("%v: %s", d.Sp, d.Message)
	}
	var kinds []string
	for _, d := range f.Decls {
		kinds = append(kinds, d.Kind.Name+" "+d.NameText())
	}
	want := []string{
		"Namespace example", "Noise ridges", "DensityFunction terrain", "SurfaceRule badlands_top",
		"LootTable blocks/ruby_ore", "Recipe ruby_block", "Objective kills", "Global deaths_total",
		"Function on_kill", "OnLoad ",
	}
	if len(kinds) != len(want) {
		t.Fatalf("decls = %q", kinds)
	}
	for i := range want {
		if kinds[i] != want[i] {
			t.Errorf("decl %d = %q, want %q", i, kinds[i], want[i])
		}
	}
	if f.Decls[1].Doc != "Ridges used by the custom continent shape." {
		t.Errorf("doc = %q", f.Decls[1].Doc)
	}
}
