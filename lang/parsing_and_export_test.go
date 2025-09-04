package lang

import (
	"encoding/json"
	"testing"
)

// helper: get JSON content of an exported file if present
func getExportContent(t *testing.T, p *Project, path string) (string, bool) {
	t.Helper()
	tree := p.BuildFsLike("data")
	node := tree.Get(path)
	if node == nil || node.IsDir {
		return "", false
	}
	return node.Content, true
}

func TestParseValidSurfaceRule_Block_AndExport(t *testing.T) {
	p := NewProject()
	src := `namespace a;
Rule1 := SurfaceRule {
    Block minecraft:stone
}`
	f := p.AddFile("a1.mms", src)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse returned error: %v", err)
	}
	if len(f.Diagnostics) != 0 {
		t.Fatalf("expected no diagnostics, got %d: %+v", len(f.Diagnostics), f.Diagnostics)
	}

	// Expect export at a/_debug/surface_rule/Rule1.json
	content, ok := getExportContent(t, p, "a/_debug/surface_rule/Rule1.json")
	if !ok {
		t.Fatalf("expected exported file for Rule1.json, not found")
	}
	if content == "" {
		t.Fatalf("exported content is empty")
	}

	// Validate JSON content
	var obj map[string]any
	if err := json.Unmarshal([]byte(content), &obj); err != nil {
		t.Fatalf("exported content is not valid JSON: %v\n%s", err, content)
	}
	if typ, _ := obj["type"].(string); typ != "minecraft:block" {
		t.Fatalf("expected type 'minecraft:block', got %q in JSON: %s", typ, content)
	}
	res, ok := obj["result_state"].(map[string]any)
	if !ok {
		t.Fatalf("expected 'result_state' object in JSON: %s", content)
	}
	if name, _ := res["Name"].(string); name != "minecraft:stone" {
		t.Fatalf("expected result_state.Name 'minecraft:stone', got %q in JSON: %s", name, content)
	}
}
