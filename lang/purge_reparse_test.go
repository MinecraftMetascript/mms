package lang

import (
	"testing"
)

func TestPurgeSymbolsOnReparse_RemovesOldExports(t *testing.T) {
	p := NewProject()

	// First parse with RuleA
	src1 := `namespace n;
RuleA := SurfaceRule { Block minecraft:stone }`
	f := p.AddFile("reparse.mms", src1)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse #1 error: %v", err)
	}

	if _, ok := getExportContent(t, p, "n/_debug/surface_rule/RuleA.json"); !ok {
		t.Fatalf("expected RuleA export after first parse")
	}

	// Update file content to RuleB and re-parse (same File path)
	src2 := `namespace n;
RuleB := SurfaceRule { Block minecraft:dirt }`
	f = p.AddFile("reparse.mms", src2)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse #2 error: %v", err)
	}

	// Old symbol should be purged; new one exported
	if _, ok := getExportContent(t, p, "n/_debug/surface_rule/RuleA.json"); ok {
		t.Fatalf("did not expect RuleA export to remain after reparse")
	}
	if _, ok := getExportContent(t, p, "n/_debug/surface_rule/RuleB.json"); !ok {
		t.Fatalf("expected RuleB export after reparse")
	}
}
