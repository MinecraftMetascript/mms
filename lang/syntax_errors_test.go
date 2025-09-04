package lang

import (
    "testing"
)

func TestSyntaxError_MissingDeclareOperator(t *testing.T) {
    p := NewProject()
    src := `namespace n;
X SurfaceRule { Block minecraft:stone }`
    f := p.AddFile("missing_declare.mms", src)
    if err := f.Parse(); err != nil {
        t.Fatalf("parse returned error: %v", err)
    }
    if len(f.Diagnostics) == 0 {
        t.Fatalf("expected diagnostics for missing ':=', got none")
    }
    d := f.Diagnostics[0]
    if d.Source != "parser" || d.Severity != "error" {
        t.Fatalf("expected parser error, got source=%s severity=%s", d.Source, d.Severity)
    }
    if d.File != f.Path || d.Where.Filename != f.Path {
        t.Fatalf("expected diagnostic filenames set to file path")
    }
}

func TestSyntaxError_MissingDefinitionAfterDeclare(t *testing.T) {
    p := NewProject()
    src := `namespace n;
X := \n` // intentionally incomplete
    f := p.AddFile("missing_definition.mms", src)
    if err := f.Parse(); err != nil {
        t.Fatalf("parse returned error: %v", err)
    }
    if len(f.Diagnostics) == 0 {
        t.Fatalf("expected diagnostics for missing definition, got none")
    }
}
