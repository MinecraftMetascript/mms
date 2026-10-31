package syntax

import (
	"os"
	"path/filepath"
	"testing"
)

// FuzzParse checks that Parse never panics and that every span it reports
// lies inside the source.
func FuzzParse(f *testing.F) {
	seeds, _ := filepath.Glob("testdata/*.mms")
	more, _ := filepath.Glob("testdata/recovery/*.mms")
	for _, path := range append(seeds, more...) {
		src, _ := os.ReadFile(path)
		f.Add(string(src))
	}
	f.Fuzz(func(t *testing.T, src string) {
		file := Parse(src)
		n := len(file.Src)
		check := func(what string, sp Span) {
			if sp.Start < 0 || sp.End > n || sp.Start > sp.End {
				t.Fatalf("%s span %v outside source of %d runes", what, sp, n)
			}
		}
		for _, d := range file.Diags {
			check("diagnostic", d.Sp)
		}
		for _, tok := range file.Tokens {
			if !tok.Synthetic {
				check("token", tok.Sp)
			}
		}
		for _, d := range file.Decls {
			check("decl", d.Sp)
		}
		_ = Dump(file)
		// Chunked parsing must agree with a whole-file parse on valid input,
		// unless the parser joined lines across a chunk boundary (such as a
		// kind and name on separate lines), where chunks are stricter.
		if !hasErrors(file) && !straddles(file) {
			cf, _, _ := Reparse(src, nil)
			if Dump(cf) != Dump(file) || spanList(cf) != spanList(file) {
				t.Fatalf("chunked parse differs")
			}
		}
	})
}

func straddles(f *File) bool {
	for _, b := range splitChunks(f.Src) {
		for _, d := range f.Decls {
			if d.Sp.Start < b && b < d.Sp.End {
				return true
			}
		}
	}
	return false
}
