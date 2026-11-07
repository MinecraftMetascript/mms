package format

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "rewrite golden files")

// Each testdata/*.in.mms is formatted and compared with *.out.mms. The
// output must also be stable: formatting it again changes nothing. The
// language design doc's snippets (syntax/testdata/design) are checked the
// same way, with goldens in testdata/design.
func TestGolden(t *testing.T) {
	ins, _ := filepath.Glob("testdata/*.in.mms")
	design, _ := filepath.Glob("../syntax/testdata/design/*.mms")
	ins = append(ins, design...)
	for _, in := range ins {
		t.Run(filepath.Base(in), func(t *testing.T) {
			src, err := os.ReadFile(in)
			if err != nil {
				t.Fatal(err)
			}
			got, err := Source(string(src))
			if err != nil {
				t.Fatal(err)
			}
			again, err := Source(got)
			if err != nil {
				t.Fatalf("formatting the output: %v", err)
			}
			if again != got {
				t.Errorf("not idempotent\n--- first\n%s--- second\n%s", got, again)
			}
			outPath := strings.TrimSuffix(in, ".in.mms") + ".out.mms"
			if !strings.HasSuffix(in, ".in.mms") {
				outPath = filepath.Join("testdata/design", strings.TrimSuffix(filepath.Base(in), ".mms")+".out.mms")
			}
			if *update {
				os.WriteFile(outPath, []byte(got), 0o644)
				return
			}
			want, err := os.ReadFile(outPath)
			if err != nil {
				t.Fatalf("missing golden (run with -update): %v", err)
			}
			if got != string(want) {
				t.Errorf("mismatch\n--- got\n%s--- want\n%s", got, want)
			}
		})
	}
}
