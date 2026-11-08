package syntax

import (
	"os"
	"strings"
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
		"Type Ore", "DensityFunction ridges", "DensityFunction terrain", "Template gem_block",
		"Recipe ruby_block", "Objective kills", "Global deaths_total", "Function on_kill", "OnLoad ",
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
	var loops int
	for _, it := range f.Items {
		if _, ok := it.(*ForStmt); ok {
			loops++
		}
	}
	if loops != 1 {
		t.Errorf("top-level for loops = %d, want 1", loops)
	}
}

// TestCommandValue checks that a command after a lone '=' is the value,
// and that '/' anywhere else keeps its old meaning.
func TestCommandValue(t *testing.T) {
	cases := map[string]string{
		"let n = /data get block ~ ~ ~ Items":    "    Let n\n      Command \"data get block ~ ~ ~ Items\"\n",
		"g = /time query gametime":               "    Field g\n      Command \"time query gametime\"\n",
		"kills[@s] = /data get entity @s Health": "Command \"data get entity @s Health\"\n",
		"n = 1 /2":                               "      Binary /\n",
		"n = 6/2":                                "      Binary /\n",
	}
	for line, want := range cases {
		f := Parse("Function f {\n  " + line + "\n}\n")
		if len(f.Diags) > 0 {
			t.Errorf("%q: %v", line, f.Diags)
			continue
		}
		if got := Dump(f); !strings.Contains(got, want) {
			t.Errorf("%q: want %q in\n%s", line, want, got)
		}
	}
	// A command on the line after '=' is a statement of its own, and the
	// '=' is missing its value.
	f := Parse("Function f {\n  let n =\n  /say hi\n}\n")
	if len(f.Diags) != 1 || f.Diags[0].Code != "syntax/missing-value" {
		t.Errorf("diags = %v", f.Diags)
	}
	if got := Dump(f); !strings.Contains(got, "    Command \"say hi\"\n") {
		t.Errorf("the command should stay a statement:\n%s", got)
	}
	// After ==, <=, += and the like, '/' is not a command.
	for _, line := range []string{"let b = n == /x", "n += /x"} {
		f := Parse("Function f {\n  " + line + "\n}\n")
		if strings.Contains(Dump(f), "Command") {
			t.Errorf("%q lexed a command:\n%s", line, Dump(f))
		}
	}
}

// TestTemplatedScoreAssign checks that a score of an objective with a
// templated name can be set, as in a Blueprint body.
func TestTemplatedScoreAssign(t *testing.T) {
	f := Parse("Function f {\n  ${label}_cd[@s] = rate\n}\n")
	if len(f.Diags) > 0 {
		t.Fatalf("diags = %v", f.Diags)
	}
	as, ok := f.Decls[0].Body.Items[0].(*AssignStmt)
	if !ok {
		t.Fatalf("got %T", f.Decls[0].Body.Items[0])
	}
	ix, ok := as.Target.(*Index)
	if !ok {
		t.Fatalf("target %T", as.Target)
	}
	if n, ok := ix.X.(*TemplatedName); !ok || n.Text != "${label}_cd" || len(n.Parts) != 2 {
		t.Errorf("name = %+v", ix.X)
	}
}
