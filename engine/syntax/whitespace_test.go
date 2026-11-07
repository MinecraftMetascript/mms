package syntax

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// TestWhitespaceAgnostic checks that line breaks and indentation never
// change a parse: every design snippet, example and checker fixture parses
// to the same tree with its tokens all on one line, and with every token
// on a line of its own. Comments are left out of all three versions, and a
// raw command keeps a line of its own, since it runs to the end of its
// line.
func TestWhitespaceAgnostic(t *testing.T) {
	var paths []string
	for _, glob := range []string{"testdata/design/*.mms", "../../examples/*/*.mms", "../../examples/*/*/*.mms", "../sema/testdata/*.mms"} {
		m, _ := filepath.Glob(glob)
		paths = append(paths, m...)
	}
	if len(paths) < 20 {
		t.Fatalf("found only %d sources", len(paths))
	}
	for _, path := range paths {
		src, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		f := Parse(string(src))
		if len(f.Diags) > 0 {
			continue // recovery fixtures; their recovery may use the layout
		}
		base := Dump(Parse(rejoin(f, " ", "\n")))
		for _, v := range []struct{ name, sep, cmd string }{{"one line", " ", "\n"}, {"one token per line", "\n", "\n"}, {"indented", "\n    ", "\n"}} {
			text := rejoin(f, v.sep, v.cmd)
			g := Parse(text)
			if len(g.Diags) > 0 {
				t.Errorf("%s, %s: %s at %d\n%s", path, v.name, g.Diags[0].Message, g.Diags[0].Sp.Start, text)
				continue
			}
			if got := Dump(g); got != base {
				t.Errorf("%s, %s: tree differs\n%s", path, v.name, firstDiff(base, got))
			}
			// The editor and the CLI read files in chunks (Reparse); that
			// must agree with Parse on every layout too.
			if r, _, _ := Reparse(text, nil); len(r.Diags) > 0 {
				t.Errorf("%s, %s: Reparse: %s at %d", path, v.name, r.Diags[0].Message, r.Diags[0].Sp.Start)
			} else if got := Dump(r); got != base {
				t.Errorf("%s, %s: Reparse tree differs\n%s", path, v.name, firstDiff(base, got))
			}
		}
	}
}

// rejoin writes f's tokens with sep between them and no comments. A raw
// command goes on a line of its own, with cmd before and after it.
func rejoin(f *File, sep, cmd string) string {
	var b strings.Builder
	for _, tk := range f.Tokens {
		if tk.Hidden || tk.Synthetic {
			continue
		}
		if tk.Kind == parser.MMSLexerCOMMAND {
			b.WriteString(cmd + tk.Text + cmd)
			continue
		}
		b.WriteString(tk.Text + sep)
	}
	return b.String()
}

func firstDiff(a, b string) string {
	al, bl := strings.Split(a, "\n"), strings.Split(b, "\n")
	for i := 0; i < len(al) && i < len(bl); i++ {
		if al[i] != bl[i] {
			lo := max(0, i-3)
			return "want:\n" + strings.Join(al[lo:min(len(al), i+3)], "\n") + "\ngot:\n" + strings.Join(bl[lo:min(len(bl), i+3)], "\n")
		}
	}
	return "lengths differ"
}

// TestReparseContinuationLines checks that a line which only continues
// the item above, such as an `else` arm, is not read as a new chunk.
func TestReparseContinuationLines(t *testing.T) {
	for _, src := range []string{
		"Global flag: bool = true\nconst x: int = if (flag) 1\nelse 2\n",
		"const xs = [1,\n2]\nconst y = 3\n",
		"Function f {\n  /say hi\n}\nlet_me = 1\n",
	} {
		whole := Parse(src)
		r, _, _ := Reparse(src, nil)
		if (len(whole.Diags) == 0) != (len(r.Diags) == 0) {
			t.Errorf("%q: Parse has %d problems, Reparse %d", src, len(whole.Diags), len(r.Diags))
			continue
		}
		if len(whole.Diags) == 0 && Dump(whole) != Dump(r) {
			t.Errorf("%q: trees differ\n%s", src, firstDiff(Dump(whole), Dump(r)))
		}
	}
}
