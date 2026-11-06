package prelude_test

import (
	"context"
	"testing"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestClean checks that the prelude parses and evaluates without a
// problem: none of its diagnostics would ever be shown.
func TestClean(t *testing.T) {
	for _, f := range prelude.Parse() {
		for _, d := range f.File.Diags {
			t.Errorf("%s: %s", f.Path, d.Message)
		}
	}
	p := project.New(spec.MustDefault(), project.Config{Namespace: "demo"})
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range prelude.Parse() {
		for _, d := range r.Eval.Files[f.Path].Diags {
			line, col := f.File.LineCol(d.Sp.Start)
			t.Errorf("%s:%d:%d %s: %s", f.Path, line, col, d.Code, d.Message)
		}
	}
}

// TestIntrinsicsDeclared checks that every intrinsic has an extern
// Template in the prelude, which gives it its signature and docs. The
// evaluator reports an extern without an intrinsic, which TestClean
// catches.
func TestIntrinsicsDeclared(t *testing.T) {
	declared := map[string]bool{}
	for _, f := range prelude.Parse() {
		for _, d := range f.File.Decls {
			if d.Extern && d.Kind != nil && d.Kind.Name == "Template" && d.Name != nil {
				declared[d.Name.Text] = true
			}
		}
	}
	for _, n := range eval.Intrinsics() {
		if !declared[n] {
			t.Errorf("intrinsic %s has no extern Template in the prelude", n)
		}
	}
}
