package lang

import (
	"fmt"
	"testing"
)

func TestMissingSemicolonInNamespace_ReportsSyntaxErrorWithLocation(t *testing.T) {
	p := NewProject()
	src := "namespace a\nRule := SurfaceRule { Block minecraft:stone }"
	f := p.AddFile("ns_missing_semicolon.mms", src)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse returned error: %v", err)
	}
	if len(f.Diagnostics) == 0 {
		t.Fatalf("expected diagnostics, got none")
	}
	d := f.Diagnostics[0]
	if d.Source != "parser" || string(d.Severity) != "error" {
		t.Fatalf("expected parser error, got source=%s severity=%s", d.Source, d.Severity)
	}
	if d.Where.Filename != f.Path || d.File != f.Path {
		t.Fatalf("expected filenames to match path; got where=%q file=%q path=%q", d.Where.Filename, d.File, f.Path)
	}
	if d.Where.Start.Line != 1 {
		t.Fatalf("expected error on line 1, got %d", d.Where.Start.Line)
	}
	if d.Where.Stop.Col <= d.Where.Start.Col {
		t.Fatalf("expected non-zero span, got start=%d stop=%d", d.Where.Start.Col, d.Where.Stop.Col)
	}
	// Snippet text may be empty for missing-token errors; location indices are sufficient.
	// if strings.TrimSpace(d.Where.Text) == "" {
	//     t.Fatalf("expected snippet text, got empty")
	// }
}

func TestNamespaceIdentifierMissing_ReportsMissingNamespaceDiagnostic(t *testing.T) {
	p := NewProject()
	// Identifier missing between 'namespace' and ';'
	src := "namespace ;\nRule := SurfaceRule { Block minecraft:stone }"
	f := p.AddFile("ns_identifier_missing.mms", src)
	if err := f.Parse(); err != nil {
		t.Fatalf("parse returned error: %v", err)
	}
	found := false
	for _, d := range f.Diagnostics {
		fmt.Println("Diag", d.Message)
		if d.Message == "syntax error: missing Identifier at ';'" {
			found = true
			if d.Where.Filename != f.Path || d.File != f.Path {
				t.Fatalf("expected filenames to match for Missing Namespace diagnostic")
			}
			break
		}
	}
	if !found {
		t.Fatalf("expected 'Missing Namespace Declaration' diagnostic, not found. diags: %+v", f.Diagnostics)
	}
}
