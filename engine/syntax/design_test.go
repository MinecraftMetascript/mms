package syntax

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDesignSnippets parses every code snippet from the language design
// doc (testdata/design) and compares the outline with its golden file.
func TestDesignSnippets(t *testing.T) {
	files, _ := filepath.Glob("testdata/design/*.mms")
	if len(files) == 0 {
		t.Fatal("no design snippets")
	}
	for _, path := range files {
		t.Run(filepath.Base(path), func(t *testing.T) {
			src, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			f := Parse(string(src))
			for _, d := range f.Diags {
				line, col := f.LineCol(d.Sp.Start)
				t.Errorf("%d:%d: %s", line, col, d.Message)
			}
			got := Dump(f)
			golden := strings.TrimSuffix(path, ".mms") + ".golden"
			if *update {
				if err := os.WriteFile(golden, []byte(got), 0o644); err != nil {
					t.Fatal(err)
				}
				return
			}
			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("missing golden file (run with -update): %v", err)
			}
			if got != string(want) {
				t.Errorf("outline mismatch\n--- got\n%s--- want\n%s", got, want)
			}
		})
	}
}
