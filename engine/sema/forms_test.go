package sema

import (
	"testing"

	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Every conversion of a built-in form names a vanilla type, and every
// form the parser gives parts for is declared.
func TestFormsResolve(t *testing.T) {
	p := spec.MustDefault()
	got := Forms(p)
	want := 0
	for _, f := range prelude.Parse() {
		for _, d := range f.File.Decls {
			if !d.Extern || d.Kind == nil || d.Kind.Name != "Type" || d.Body == nil {
				continue
			}
			for _, it := range d.Body.Items {
				if _, ok := it.(*syntax.AsStmt); ok {
					want++
				}
			}
		}
	}
	n := 0
	for _, cs := range got {
		n += len(cs)
	}
	if n != want {
		t.Errorf("%d of %d form conversions resolve", n, want)
	}
	for _, form := range formOrder {
		if len(got[form]) == 0 {
			t.Errorf("form %s has no conversion", form)
		}
	}
}
