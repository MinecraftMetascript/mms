package lang

import (
	"encoding/json"
	"testing"
)

func TestSurfaceRule_Sequence_ExportJSON(t *testing.T) {
	p := NewProject()
	src := `namespace seq;
SeqRule := SurfaceRule {
    Sequence [
        Block minecraft:stone
        Block minecraft:dirt
    ]
}`
	f := p.AddFile("seq.mms", src)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse returned error: %v", err)
	}
	if len(f.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got: %+v", f.Diagnostics)
	}
	content, ok := getExportContent(t, p, "seq/_debug/surface_rule/SeqRule.json")
	if !ok {
		t.Fatalf("expected sequence export")
	}
	var obj map[string]any
	if err := json.Unmarshal([]byte(content), &obj); err != nil {
		t.Fatalf("invalid sequence JSON: %v\n%s", err, content)
	}
	if typ, _ := obj["type"].(string); typ != "minecraft:sequence" {
		t.Fatalf("expected type minecraft:sequence, got %q", typ)
	}
	seq, ok := obj["sequence"].([]any)
	if !ok || len(seq) != 2 {
		t.Fatalf("expected sequence of 2 items, got %v", obj["sequence"])
	}
	// spot check second entry is a block dirt
	if second, ok := seq[1].(map[string]any); ok {
		if typ, _ := second["type"].(string); typ != "minecraft:block" {
			t.Fatalf("expected second item type minecraft:block, got %q", typ)
		}
		if rs, ok := second["result_state"].(map[string]any); ok {
			if name, _ := rs["Name"].(string); name != "minecraft:dirt" {
				t.Fatalf("expected second item result_state.Name minecraft:dirt, got %q", name)
			}
		} else {
			t.Fatalf("missing result_state in second sequence item")
		}
	} else {
		t.Fatalf("unexpected JSON for second sequence item: %v", seq[1])
	}
}

func TestSurfaceRule_If_AboveWater_JSON(t *testing.T) {
	p := NewProject()
	src := `namespace n;
Rule := SurfaceRule {
    If ( AboveWater 2 1.5 Add )
        Block minecraft:stone
}`
	f := p.AddFile("if_water.mms", src)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse returned error: %v", err)
	}
	if len(f.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got: %+v", f.Diagnostics)
	}
	content, ok := getExportContent(t, p, "n/_debug/surface_rule/Rule.json")
	if !ok {
		t.Fatalf("expected export for If rule")
	}
	var obj map[string]any
	if err := json.Unmarshal([]byte(content), &obj); err != nil {
		t.Fatalf("invalid If JSON: %v\n%s", err, content)
	}
	if typ, _ := obj["type"].(string); typ != "minecraft:condition" {
		t.Fatalf("expected type minecraft:condition, got %q", typ)
	}
	ifTrue, ok := obj["if_true"].(map[string]any)
	if !ok {
		t.Fatalf("expected if_true object in JSON")
	}
	if ctyp, _ := ifTrue["type"].(string); ctyp != "minecraft:water" {
		t.Fatalf("expected if_true.type minecraft:water, got %q", ctyp)
	}
	if off, _ := ifTrue["offset"].(float64); int(off) != 2 { // JSON numbers are float64
		t.Fatalf("expected offset 2, got %v", ifTrue["offset"])
	}
	if mult, _ := ifTrue["surface_depth_multiplier"].(float64); mult != 1.5 {
		t.Fatalf("expected surface_depth_multiplier 1.5, got %v", ifTrue["surface_depth_multiplier"])
	}
	if add, _ := ifTrue["add_stone_depth"].(bool); add != true {
		t.Fatalf("expected add_stone_depth true, got %v", ifTrue["add_stone_depth"])
	}
	thenRun, ok := obj["then_run"].(map[string]any)
	if !ok {
		t.Fatalf("expected then_run object in JSON")
	}
	if trtyp, _ := thenRun["type"].(string); trtyp != "minecraft:block" {
		t.Fatalf("expected then_run.type minecraft:block, got %q", trtyp)
	}
}

func TestIf_WithConditionReference_Resolves(t *testing.T) {
	p := NewProject()
	src := `namespace n;
Cond := SurfaceCondition { AboveSurface }
Rule := SurfaceRule {
    If ( Cond )
        Block minecraft:stone
}`
	f := p.AddFile("cond_ref.mms", src)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse returned error: %v", err)
	}
	if len(f.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got: %+v", f.Diagnostics)
	}
	content, ok := getExportContent(t, p, "n/_debug/surface_rule/Rule.json")
	if !ok {
		t.Fatalf("expected export for If rule with condition ref")
	}
	var obj map[string]any
	if err := json.Unmarshal([]byte(content), &obj); err != nil {
		t.Fatalf("invalid JSON: %v\n%s", err, content)
	}
	ifTrue, ok := obj["if_true"].(map[string]any)
	if !ok {
		t.Fatalf("expected if_true object")
	}
	if ctyp, _ := ifTrue["type"].(string); ctyp != "minecraft:above_preliminary_surface" {
		t.Fatalf("expected resolved condition type minecraft:above_preliminary_surface, got %q", ctyp)
	}
}

func TestProjectDiagnostics_AggregatesAcrossFiles(t *testing.T) {
	p := NewProject()
	f1 := p.AddFile("bad1.mms", "namespace a\nX := ")
	if err := f1.Parse(); err != nil {
		t.Fatalf("parse f1: %v", err)
	}
	f2 := p.AddFile("bad2.mms", "namespace b X := SurfaceRule { Block minecraft:stone }")
	if err := f2.Parse(); err != nil {
		t.Fatalf("parse f2: %v", err)
	}
	// Expect diagnostics in both files
	if len(f1.Diagnostics) == 0 || len(f2.Diagnostics) == 0 {
		t.Fatalf("expected diagnostics in both files, got f1=%d f2=%d", len(f1.Diagnostics), len(f2.Diagnostics))
	}
	total := len(f1.Diagnostics) + len(f2.Diagnostics)
	if got := len(p.Diagnostics()); got != total {
		t.Fatalf("expected project diagnostics %d, got %d", total, got)
	}
}
