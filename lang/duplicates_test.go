package lang

import "testing"

func TestDuplicateSymbol_SameFile_ReportsSemanticDiagnostic(t *testing.T) {
    p := NewProject()
    src := `namespace dup;
Rule := SurfaceRule { Block minecraft:stone }
Rule := SurfaceRule { Block minecraft:dirt }`
    f := p.AddFile("dup1.mms", src)
    if err := f.Parse(); err != nil {
        t.Fatalf("parse returned error: %v", err)
    }
    found := false
    for _, d := range f.Diagnostics {
        if d.Source == "semantic" && d.Severity == "error" && d.Where.Filename == f.Path {
            if d.Message != "" && (contains(d.Message, "already exists") || contains(d.Message, "dup:Rule")) {
                found = true
                break
            }
        }
    }
    if !found {
        t.Fatalf("expected duplicate declaration diagnostic; diags: %+v", f.Diagnostics)
    }
}

func TestDuplicateSymbol_CrossFile_ReportsOnSecondFile(t *testing.T) {
    p := NewProject()
    f1 := p.AddFile("file1.mms", `namespace x; A := SurfaceRule { Block minecraft:stone }`)
    if err := f1.Parse(); err != nil {
        t.Fatalf("parse error on f1: %v", err)
    }
    f2 := p.AddFile("file2.mms", `namespace x; A := SurfaceRule { Block minecraft:dirt }`)
    if err := f2.Parse(); err != nil {
        t.Fatalf("parse error on f2: %v", err)
    }
    if len(f2.Diagnostics) == 0 {
        t.Fatalf("expected diagnostics on f2 for duplicate symbol, got none")
    }
}

// small helper to avoid importing strings for trivial contains
func contains(s, sub string) bool {
    return len(s) >= len(sub) && (indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
    // naive search is fine for tests
    for i := 0; i+len(sub) <= len(s); i++ {
        if s[i:i+len(sub)] == sub {
            return i
        }
    }
    return -1
}
