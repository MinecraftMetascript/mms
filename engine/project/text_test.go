package project

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

func adv(title string) string {
	return "Advancement a = { display = { title = " + title + ", description = \"d\", icon = { id = stone }, background = \"minecraft:block/stone\" }, criteria = { x = Tick() } }\n"
}

// A string where a text component goes is markup; one with no markup
// in it is the same string in the output.
func TestTextMarkup(t *testing.T) {
	for _, tc := range []struct{ title, want string }{
		{`"Plain title"`, `"title": "Plain title"`},
		{`"**Bold** start"`, `"bold": true`},
		{`"{gold}Gold"`, `"color": "gold"`},
		{`"Open [rules](run:/say hi)"`, `"command": "say hi"`},
		{`"Cost: 5 * 3"`, `"title": "Cost: 5 * 3"`},
		{`"Cost: 5 \* 3"`, `"title": "Cost: 5 * 3"`},
	} {
		_, r := bookProject(t, adv(tc.title), nil)
		if d := diagText(r); d != "" {
			t.Errorf("%s: %s", tc.title, d)
			continue
		}
		if got := emitted(t, r, "data/demo/advancement/a.json"); !strings.Contains(got, tc.want) {
			t.Errorf("%s: output lacks %s:\n%s", tc.title, tc.want, got)
		}
	}
}

// A mark left open is the character itself, and the lint says so.
func TestAccidentalMarkup(t *testing.T) {
	src := adv(`"Cost: 5 * 3"`)
	_, r := bookProject(t, src, nil)
	var found *syntax.Diagnostic
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/accidental-markup" {
			found = &d
		}
	}
	if found == nil {
		t.Fatalf("no accidental-markup hint: %+v", r.Diagnostics("main.mms"))
	}
	if got := src[found.Sp.Start:found.Sp.End]; got != "*" {
		t.Errorf("hint at %q", got)
	}
	if len(found.Fixes) == 0 || found.Fixes[0].Edits[0].Text != `\` {
		t.Errorf("fixes = %+v", found.Fixes)
	}
}

// Pages are only for books.
func TestTextHasNoPages(t *testing.T) {
	for _, tc := range []struct {
		src   string
		files map[string]string
		path  string
		want  string
	}{
		{adv(`"One\n---\nTwo"`), nil, "main.mms", "a page break works only in a book's text"},
		{adv(`"[Next](page:2)"`), nil, "main.mms", "a page link works only in a book's text"},
		{"Book b { title = \"T\", author = \"A\", text = TextFile(\"b.mmstext\") }\n", map[string]string{"b.mmstext": "One\n---\nTwo\n"}, "b.mmstext", "a page break works only"},
	} {
		_, r := bookProject(t, tc.src, tc.files)
		var msgs []string
		for _, d := range r.Diagnostics(tc.path) {
			msgs = append(msgs, d.Message)
		}
		if !strings.Contains(strings.Join(msgs, "\n"), tc.want) {
			t.Errorf("%q: diagnostics %q lack %q", tc.src, msgs, tc.want)
		}
	}
	// In a .mmstext file, the fix removes the --- line.
	_, r := bookProject(t, "Book b { title = \"T\", author = \"A\", text = TextFile(\"b.mmstext\") }\n", map[string]string{"b.mmstext": "One\n---\nTwo\n"})
	for _, d := range r.Diagnostics("b.mmstext") {
		if d.Code != "mms/page-break-in-text" {
			continue
		}
		if len(d.Fixes) != 1 || d.Fixes[0].Edits[0].Sp != (syntax.Span{Start: 4, End: 8}) {
			t.Errorf("fix: %+v", d.Fixes)
		}
		return
	}
	t.Errorf("no page-break-in-text in b.mmstext: %v", r.Diagnostics("b.mmstext"))
}

// TextFile reads a .mmstext file where a text component goes, and a Book
// reads one too.
func TestTextFile(t *testing.T) {
	src := adv(`TextFile("text/title.mmstext")`)
	_, r := bookProject(t, src, map[string]string{"text/title.mmstext": "{red}Siege{/} rules\n"})
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	if got := emitted(t, r, "data/demo/advancement/a.json"); !strings.Contains(got, `"color": "red"`) {
		t.Errorf("output:\n%s", got)
	}
	src = "Book b { title = \"T\", author = \"A\", text = TextFile(\"b.mmstext\") }\nLootTable l { pools = [Pool(1, [b])] }\n"
	_, r = bookProject(t, src, map[string]string{"b.mmstext": "**Hello**\n"})
	if d := diagText(r); d != "" {
		t.Fatal(d)
	}
	if pages := pagesOf(t, emitted(t, r, "data/demo/loot_table/l.json")); len(pages) != 1 {
		t.Errorf("pages = %v", pages)
	}
}
