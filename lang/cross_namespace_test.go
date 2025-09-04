package lang

import "testing"

func TestCrossNamespaceReference_Succeeds_AndExports(t *testing.T) {
    p := NewProject()
    fA := p.AddFile("a.mms", `namespace a; Base := SurfaceRule { Block minecraft:stone }`)
    if err := fA.Parse(); err != nil {
        t.Fatalf("parse a: %v", err)
    }
    fB := p.AddFile("b.mms", `namespace b; Use := SurfaceRule { a:Base }`)
    if err := fB.Parse(); err != nil {
        t.Fatalf("parse b: %v", err)
    }

    if len(fB.Diagnostics) != 0 {
        t.Fatalf("expected no diagnostics in referencing file, got: %+v", fB.Diagnostics)
    }

    if _, ok := getExportContent(t, p, "b/_debug/surface_rule/Use.json"); !ok {
        t.Fatalf("expected Use export in namespace b")
    }
}

func TestCrossNamespaceReference_FailsWhenMissing_NoExport(t *testing.T) {
    p := NewProject()
    // Define nothing in namespace a
    fB := p.AddFile("b2.mms", `namespace b; Use := SurfaceRule { a:Missing }`)
    if err := fB.Parse(); err != nil {
        t.Fatalf("parse b2: %v", err)
    }

    if _, ok := getExportContent(t, p, "b/_debug/surface_rule/Use.json"); ok {
        t.Fatalf("did not expect export for unresolved cross-namespace reference")
    }
}
