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

// TestMissingListComma checks that list elements need commas, that the
// list still parses whole, and that the fix inserts each comma.
func TestMissingListComma(t *testing.T) {
	src := "const x = [a, b\n  c d]\n"
	f := Parse(src)
	var got []string
	for _, d := range f.Diags {
		if d.Code != "syntax/missing-comma" {
			t.Errorf("unexpected %s: %s", d.Code, d.Message)
			continue
		}
		got = append(got, d.Fixes[0].Edits[0].Text)
	}
	if len(got) != 2 {
		t.Fatalf("got %d missing-comma errors, want 2: %v", len(got), f.Diags)
	}
	fixed := []rune(src)
	for i := len(f.Diags) - 1; i >= 0; i-- {
		e := f.Diags[i].Fixes[0].Edits[0]
		fixed = append(fixed[:e.Sp.Start], append([]rune(e.Text), fixed[e.Sp.End:]...)...)
	}
	if string(fixed) != "const x = [a, b,\n  c, d]\n" {
		t.Errorf("fixed = %q", string(fixed))
	}
	if d := Parse(string(fixed)).Diags; len(d) != 0 {
		t.Errorf("fixed source still has %v", d)
	}
}
