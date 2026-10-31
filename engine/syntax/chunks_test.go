package syntax

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func chunkSources(t testing.TB) []string {
	var out []string
	for _, pat := range []string{"testdata/*.mms", "../sema/testdata/*.mms", "../../examples/*/*.mms", "../format/testdata/*.mms"} {
		paths, _ := filepath.Glob(pat)
		for _, p := range paths {
			b, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			out = append(out, string(b))
		}
	}
	return out
}

func hasErrors(f *File) bool {
	for _, d := range f.Diags {
		if d.Severity == SeverityError {
			return true
		}
	}
	return false
}

// TestReparseMatchesParse checks that chunked parsing gives the same tree
// as parsing the whole file, for files without syntax errors.
func TestReparseMatchesParse(t *testing.T) {
	for i, src := range chunkSources(t) {
		whole := Parse(src)
		if hasErrors(whole) {
			continue
		}
		f, _, _ := Reparse(src, nil)
		if got, want := Dump(f), Dump(whole); got != want {
			t.Errorf("source %d: chunked parse differs\n--- chunked\n%s\n--- whole\n%s", i, got, want)
		}
		if a, b := spanList(f), spanList(whole); a != b {
			t.Errorf("source %d: spans differ", i)
		}
		if len(f.Tokens) != len(whole.Tokens) {
			t.Errorf("source %d: %d tokens, want %d", i, len(f.Tokens), len(whole.Tokens))
		}
	}
}

// TestReparseReuse edits one declaration and checks that the others are
// reused, and that the result matches a fresh parse.
func TestReparseReuse(t *testing.T) {
	src := "/** doc */\nNoise a = { base_octave = -1 }\n\nDensityFunction b = 1\n\nFunction f {\n  /say {\n}\n\nDensityFunction c = 2\n"
	_, chunks, _ := Reparse(src, nil)
	if len(chunks) != 4 {
		t.Fatalf("got %d chunks", len(chunks))
	}
	if chunks[0].File.Decls[0].Doc != "doc" {
		t.Errorf("doc = %q", chunks[0].File.Decls[0].Doc)
	}
	edited := strings.Replace(src, "b = 1", "b = 10 + 1", 1)
	f, chunks2, moved := Reparse(edited, chunks)
	if chunks2[0].File != chunks[0].File {
		t.Error("first chunk not reused in place")
	}
	if len(moved) != 2 {
		t.Errorf("moved = %d, want 2", len(moved))
	}
	fresh, _, _ := Reparse(edited, nil)
	if Dump(f) != Dump(fresh) {
		t.Errorf("incremental parse differs:\n%s\n---\n%s", Dump(f), Dump(fresh))
	}
	var spans, freshSpans []Span
	Walk(f, func(n Node) bool { spans = append(spans, n.Span()); return true })
	Walk(fresh, func(n Node) bool { freshSpans = append(freshSpans, n.Span()); return true })
	if len(spans) != len(freshSpans) {
		t.Fatalf("%d nodes, want %d", len(spans), len(freshSpans))
	}
	for i := range spans {
		if spans[i] != freshSpans[i] {
			t.Errorf("node %d span %v, want %v", i, spans[i], freshSpans[i])
		}
	}
	for d, m := range moved {
		if m.Nodes[m.Old] != d || m.Delta != 5 {
			t.Errorf("bad move record %+v", m)
		}
	}
}

func spanList(f *File) string {
	var sb strings.Builder
	Walk(f, func(n Node) bool {
		sp := n.Span()
		sb.WriteString(strconv.Itoa(sp.Start) + "-" + strconv.Itoa(sp.End) + " ")
		return true
	})
	return sb.String()
}

// TestChunkMultiLineDoc checks that a doc comment over several lines
// stays in its declaration's chunk, and that a block comment opened after
// code never starts a chunk.
func TestChunkMultiLineDoc(t *testing.T) {
	src := "DensityFunction a = 1\n\n/**\n * Line one.\n *\n * Line two.\n */\nDensityFunction b = 2 /* trailing\n comment\n*/\nDensityFunction c = 3\n"
	f, chunks, _ := Reparse(src, nil)
	if len(chunks) != 3 {
		t.Fatalf("got %d chunks", len(chunks))
	}
	if got := f.Decls[1].Doc; got != "Line one.\n\nLine two." {
		t.Errorf("doc = %q", got)
	}
	if whole := Parse(src); whole.Decls[1].Doc != f.Decls[1].Doc {
		t.Errorf("chunked doc %q, whole-file doc %q", f.Decls[1].Doc, whole.Decls[1].Doc)
	}
	if got := string([]rune(src)[chunks[2].Off:]); !strings.HasPrefix(got, "DensityFunction c") {
		t.Errorf("third chunk starts at %q", got)
	}
}
