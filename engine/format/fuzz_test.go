package format

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// FuzzSource checks that the formatter never panics, never reports an
// unsafe result (which would mean a printer bug), and is idempotent.
func FuzzSource(f *testing.F) {
	seeds, _ := filepath.Glob("testdata/*.mms")
	for _, path := range seeds {
		src, _ := os.ReadFile(path)
		f.Add(string(src))
	}
	f.Fuzz(func(t *testing.T, src string) {
		out, err := Source(src)
		if errors.Is(err, ErrSyntax) {
			return
		}
		if err != nil {
			t.Fatalf("Source(%q): %v", src, err)
		}
		again, err := Source(out)
		if err != nil || again != out {
			t.Fatalf("not idempotent for %q:\nfirst:\n%s\nsecond:\n%s (err %v)", src, out, again, err)
		}
	})
}
