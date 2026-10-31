package format

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// verify checks that out parses to the same tree as the input and keeps
// every comment. Formatting must never change meaning (docs/plan.md,
// "Safety").
func verify(in *syntax.File, out string) error {
	re := syntax.Parse(out)
	if len(re.Diags) > 0 || syntax.Dump(re) != syntax.Dump(in) {
		return ErrUnsafe
	}
	if strings.Join(commentTexts(in), "\x00") != strings.Join(commentTexts(re), "\x00") {
		return ErrUnsafe
	}
	return nil
}

func commentTexts(f *syntax.File) []string {
	var out []string
	for _, t := range f.Tokens {
		if t.IsComment() {
			// Block comments may be re-indented, so compare their words.
			out = append(out, strings.Join(strings.Fields(strings.ReplaceAll(t.Text, "*", " ")), " "))
		}
	}
	return out
}
