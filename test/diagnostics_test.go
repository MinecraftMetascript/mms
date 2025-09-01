package test

import (
	"testing"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func TestSemanticDuplicateSymbolDiagnostic_CrossNamespace(t *testing.T) {
	p := lang.NewProject()
	f1 := p.AddFile("one.mms", "namespace test; x := SurfaceRule.Block minecraft:stone")
	f2 := p.AddFile("two.mms", "namespace alternate; x := SurfaceRule.Block minecraft:dirt")
	f1.Parse()
	f2.Parse()

	// Prefer to see the semantic error on the second file, but assert at project level too
	if len(f1.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", f1.Diagnostics)
	}
	if len(f2.Diagnostics) != 0 {
		t.Errorf("expected no diagnostics, got: %+v", f2.Diagnostics)
	}
	if _, ok := p.GlobalScope.Get(*traversal.Ref("test", "x")); !ok {
		t.Errorf("expected symbol to be defined in test namespace")
	}
	if _, ok := p.GlobalScope.Get(*traversal.Ref("alternate", "x")); !ok {
		t.Errorf("expected symbol to be defined in alternate namespace")
	}
}

func TestSemanticDuplicateSymbolDiagnostic(t *testing.T) {
	p := lang.NewProject()
	f1 := p.AddFile("one.mms", "namespace test; x := SurfaceRule.Block minecraft:stone\nx := SurfaceRule.Block minecraft:dirt")
	f1.Parse()

	if len(f1.Diagnostics) == 0 {
		t.Fatalf("expected at least 1 semantic diagnostic, got 0")
	}
	found := false
	for _, d := range f1.Diagnostics {
		if d.Source == "semantic" && d.Severity == "error" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected a semantic error diagnostic, got: %+v", f1.Diagnostics)
	}

}
func TestSemanticDuplicateSymbolDiagnosticMultiFile(t *testing.T) {
	p := lang.NewProject()
	f1 := p.AddFile("one.mms", "namespace test; x := SurfaceRule.Block minecraft:stone")
	f2 := p.AddFile("two.mms", "namespace test; x := SurfaceRule.Block minecraft:dirt")
	f1.Parse()
	f2.Parse()

	// Prefer to see the semantic error on the second file, but assert at project level too
	found := false
	for _, d := range f2.Diagnostics {
		if d.Source == "semantic" && d.Severity == "error" {
			found = true
			break
		}
	}
	if !found {
		for _, d := range p.Diagnostics {
			if d.Source == "semantic" && d.Severity == "error" {
				found = true
				break
			}
		}
	}
	if !found {
		t.Errorf("expected a semantic error diagnostic for duplicate symbol, got: f2=%+v project=%+v", f2.Diagnostics, p.Diagnostics)
	}
}

func TestSyntaxErrorDiagnostic(t *testing.T) {
	p := lang.NewProject()
	// Intentionally malformed: missing assignment and definition
	file := p.AddFile("bad.mms", "this is not valid mms syntax")
	file.Parse()

	if len(file.Diagnostics) == 0 {
		t.Fatalf("expected at least 1 syntax diagnostic, got 0")
	}
	found := false
	for _, d := range file.Diagnostics {
		if d.Source == "parser" && d.Severity == "error" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected a parser error diagnostic, got: %+v", file.Diagnostics)
	}
}

func TestProjectAggregatesDiagnostics(t *testing.T) {
	p := lang.NewProject()
	f1 := p.AddFile("one.mms", "a := SurfaceRule.Block minecraft:stone\n a := SurfaceRule.Block minecraft:stone")
	f2 := p.AddFile("two.mms", "not valid")
	f1.Parse()
	f2.Parse()

	if len(f1.Diagnostics) == 0 {
		t.Fatalf("expected diagnostics in f1, got 0")
	}
	if len(f2.Diagnostics) == 0 {
		t.Fatalf("expected diagnostics in f2, got 0")
	}
	if len(p.Diagnostics) < len(f1.Diagnostics)+len(f2.Diagnostics) {
		// Must be at least the sum; duplicates aren't introduced by aggregation
		t.Errorf("project diagnostics not aggregated: project=%d f1=%d f2=%d", len(p.Diagnostics), len(f1.Diagnostics), len(f2.Diagnostics))
	}
}
