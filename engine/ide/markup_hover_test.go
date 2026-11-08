package ide

import (
	"strings"
	"testing"
)

// `${...}` in text markup names what it reads, and says when it is
// filled in, which depends on where the text is.
func TestMarkupExprHover(t *testing.T) {
	for _, tc := range []struct {
		src, want, not string
	}{
		{
			"Type S {\n  /** Which spell is ready. */\n  mode: \"x\" | \"y\"\n}\nStorage s: S\nFunction f {\n  tell(@a, \"Mode ${s.mo^de}\")\n}",
			"s.mode: \"x\" | \"y\"\n```\nWhich spell is ready.\n\nA value the function computes, filled in when this line runs", "book",
		},
		{"Global g: int = 0\nFunction f {\n  tell(@a, \"Wave ${^g}\")\n}", "Global test:g: int", "book"},
		// The name under the cursor in a longer expression.
		{markupDecls + "Function f {\n  actionbar(@s, \"Mode ${mo^des[st.mode]}\")\n}", "const modes: [string]", "book"},
		{markupDecls + "Function f {\n  actionbar(@s, \"Mode ${modes[s^t.mode]}\")\n}", "Storage test:st", "book"},
		{markupDecls + "Function f {\n  actionbar(@s, \"Mode ${modes[st.mo^de]}\")\n}", "st.mode: int", "book"},
		{markupDecls + "Function f {\n  actionbar(@s, \"${modes[st.mode] + ^g}\")\n}", "Global test:g: int", "book"},
		{"Global g: int = 0\nAdvancement a { display = { title = \"Hi ${^g}\", description = \"d\", icon = { id = stone } }, criteria = { x = Tick() } }", "filled in when the text is shown", "book"},
	} {
		off := strings.Index(tc.src, "^")
		src := tc.src[:off] + tc.src[off+1:]
		// The `|` session marks the cursor with goes first, since the
		// union type has one of its own.
		s, _ := session(t, "|"+src)
		md, _ := s.Hover("main.mms", off)
		if !strings.Contains(md, tc.want) || strings.Contains(md, tc.not) {
			t.Errorf("%q: hover %q, want %q and not %q", src, md, tc.want, tc.not)
		}
	}
}
