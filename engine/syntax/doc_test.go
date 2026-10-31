package syntax

import "testing"

func TestCleanDoc(t *testing.T) {
	cases := []struct{ in, want string }{
		{"/** One line. */", "One line."},
		{"/**\n * Two\n * lines.\n */", "Two\nlines."},
		{"/**\n * Code:\n *\n *     let x = 1\n *   - nested\n */", "Code:\n\n    let x = 1\n  - nested"},
		{"/**\n   No stars,\n     indented.\n */", "No stars,\n  indented."},
		{"/**\r\n * CRLF\r\n */", "CRLF"},
		{"/***/", ""},
	}
	for _, tc := range cases {
		if got := cleanDoc(tc.in); got != tc.want {
			t.Errorf("cleanDoc(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestParseDoc(t *testing.T) {
	d := ParseDoc("Adds two numbers.\n\n```mms\n@notatag\n```\n@param a the first\n  number\n@param [b] - the second\n@returns the sum\n@deprecated use `plus`\n@see plus\n@example\nadd(1, 2)\n@since 1.0")
	if d.Text != "Adds two numbers.\n\n```mms\n@notatag\n```" {
		t.Errorf("text = %q", d.Text)
	}
	if len(d.Params) != 2 || d.Param("a") != "the first\nnumber" || d.Param("b") != "the second" {
		t.Errorf("params = %+v", d.Params)
	}
	if d.Returns != "the sum" || d.Deprecated == nil || *d.Deprecated != "use `plus`" {
		t.Errorf("returns = %q, deprecated = %v", d.Returns, d.Deprecated)
	}
	if len(d.See) != 1 || len(d.Examples) != 1 || d.Examples[0] != "add(1, 2)" || len(d.Other) != 1 || d.Other[0].Name != "since" {
		t.Errorf("see = %v, examples = %q, other = %v", d.See, d.Examples, d.Other)
	}
	want := "**Deprecated.** use `plus`\n\nAdds two numbers.\n\n```mms\n@notatag\n```\n\n**Parameters**\n- `a`: the first number\n- `b`: the second\n\n**Returns** the sum\n\n**Example**\n```mms\nadd(1, 2)\n```\n\n**See** plus\n\n*@since* 1.0"
	if got := d.Markdown(); got != want {
		t.Errorf("markdown =\n%s\nwant\n%s", got, want)
	}
	if got := ParseDoc("Plain **Markdown** text.").Markdown(); got != "Plain **Markdown** text." {
		t.Errorf("plain = %q", got)
	}
	if d := ParseDoc("@deprecated"); d.Deprecated == nil || *d.Deprecated != "" {
		t.Errorf("bare deprecated = %v", d.Deprecated)
	}
}

func TestDocAttaches(t *testing.T) {
	src := "/**\n * Multi-line doc.\n * @param a x\n */\nFunction f(a: int) {\n  /** A local. */\n  let y = a\n}\n/** Not attached */\n// a line comment in between\nGlobal g: int = 0\n"
	f := Parse(src)
	if f.Decls[0].Doc != "Multi-line doc.\n@param a x" {
		t.Errorf("function doc = %q", f.Decls[0].Doc)
	}
	if let := f.Decls[0].Body.Items[0].(*LetStmt); let.Doc != "A local." {
		t.Errorf("let doc = %q", let.Doc)
	}
	if f.Decls[1].Doc != "" {
		t.Errorf("global doc = %q", f.Decls[1].Doc)
	}
}
