package sema

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

var update = os.Getenv("MMS_UPDATE") != ""

// checkSource checks every resource declaration in src and returns a
// readable dump of the JSON and diagnostics.
func checkSource(src string) string {
	p := spec.MustDefault()
	kinds := NewKinds(p)
	f := syntax.Parse(src)
	u := &Unit{Path: "test.mms", File: f}
	syms, diags := Collect([]*Unit{u}, kinds, "test")
	var sb strings.Builder
	all := append([]syntax.Diagnostic{}, f.Diags...)
	all = append(all, diags[u.Path]...)
	for _, d := range f.Decls {
		k, ok := syms.DeclKind[d]
		if !ok || !k.IsResource() || k.TagOf != "" {
			continue
		}
		c := NewChecker(p, kinds, syms, syms.DeclNS[d])
		v := c.Resource(d, k)
		all = append(all, c.Diags...)
		fmt.Fprintf(&sb, "== %s %s:%s\n", k.Resource, syms.DeclNS[d], d.NameText())
		if v != nil {
			sb.Write(jsonv.Marshal(v))
		}
	}
	for _, d := range all {
		line := 1 + strings.Count(string([]rune(src)[:min(d.Sp.Start, len([]rune(src)))]), "\n")
		fmt.Fprintf(&sb, "! %d: %s %s\n", line, d.Code, d.Message)
	}
	return sb.String()
}

// TestGolden checks testdata/*.mms against *.golden. MMS_UPDATE=1 rewrites
// the golden files.
func TestGolden(t *testing.T) {
	files, _ := filepath.Glob("testdata/*.mms")
	for _, path := range files {
		t.Run(filepath.Base(path), func(t *testing.T) {
			src, _ := os.ReadFile(path)
			got := checkSource(string(src))
			golden := strings.TrimSuffix(path, ".mms") + ".golden"
			if update {
				os.WriteFile(golden, []byte(got), 0o644)
				return
			}
			want, _ := os.ReadFile(golden)
			if got != string(want) {
				t.Errorf("mismatch for %s:\n%s", path, got)
			}
		})
	}
}
