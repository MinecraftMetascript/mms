package syntax

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "rewrite golden files")

// TestRecovery parses half-typed files and compares the AST outline with a
// golden file. Every case must parse without panicking and keep the
// declarations that follow the broken one.
func TestRecovery(t *testing.T) {
	files, _ := filepath.Glob("testdata/recovery/*.mms")
	for _, path := range files {
		t.Run(filepath.Base(path), func(t *testing.T) {
			src, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			got := Dump(Parse(string(src)))
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
