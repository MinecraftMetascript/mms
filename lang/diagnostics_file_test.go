package lang

import "testing"

func TestDiagnosticFileSetsFilename(t *testing.T) {
    p := NewProject()
    file := p.AddFile("smoke.mms", "namespace smoke blockRule := SurfaceRule.Block minecraft:stone")
    if err := file.Parse(); err != nil {
        t.Fatalf("parse returned error: %v", err)
    }
    if len(file.Diagnostics) == 0 {
        t.Fatalf("expected diagnostics, got none")
    }
    d := file.Diagnostics[0]
    if d.File == "" {
        t.Fatalf("expected Diagnostic.File to be set, got empty")
    }
    if d.Where.Filename == "" {
        t.Fatalf("expected TextLocation.Filename to be set, got empty")
    }
    if d.File != d.Where.Filename {
        t.Fatalf("expected Diagnostic.File (%s) to equal TextLocation.Filename (%s)", d.File, d.Where.Filename)
    }
    if d.File != "smoke.mms" {
        t.Fatalf("expected Diagnostic.File to be 'smoke.mms', got %q", d.File)
    }
}
