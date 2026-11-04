package project

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

func bookProject(t *testing.T, src string, files map[string]string) (*Project, *Result) {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.Files = fstest.MapFS{}
	p.SetFile("main.mms", src)
	for path, text := range files {
		p.SetFile(path, text)
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return p, r
}

func emitted(t *testing.T, r *Result, path string) string {
	t.Helper()
	for _, f := range r.Emit() {
		if f.Path == path {
			return string(f.Data)
		}
	}
	t.Fatalf("no %s", path)
	return ""
}

// pagesOf returns the pages of the written book in a loot table's first
// entry.
func pagesOf(t *testing.T, lootJSON string) []any {
	t.Helper()
	var lt struct {
		Pools []struct {
			Entries []struct {
				Modifier []struct {
					Components struct {
						Content struct {
							Pages []any `json:"pages"`
						} `json:"written_book_content"`
					} `json:"components"`
				} `json:"modifier"`
			} `json:"entries"`
		} `json:"pools"`
	}
	if err := json.Unmarshal([]byte(lootJSON), &lt); err != nil {
		t.Fatal(err)
	}
	return lt.Pools[0].Entries[0].Modifier[0].Components.Content.Pages
}

func TestBook(t *testing.T) {
	const src = `Objective kills = "dummy"
Global wave = 3
Global motto: string = "Hold"
Book guide = {
  title = "Guide"
  author = "The Garrison"
  generation = tattered
  components = { rarity = rare }
  text = [
    "# Welcome",
    "**Bold** *it* __u__ ~~s~~ ||o|| {gold bold}gold{/} {#ff8800}hex{/} {shadow:none}flat{/}",
    "{key:jump} {sprite:item/diamond} {head:Notch} {t:block.minecraft.stone} \\*plain\\*",
    "${kills[@s]} ${wave} ${motto} ${2}",
    "---",
    "[Back](page:1 \"Go back\") [Go](run:/function demo:go) [Say](suggest:/say hi) [Copy](copy:x) [Hint](hover:hello) [Web](https://minecraft.wiki)",
  ]
}
LootTable chest { pools = [Pool(1, [guide])] }
Recipe copy = Shapeless([paper], guide)
Function go {
  /give @s ${guide}
}
`
	_, r := bookProject(t, src, nil)
	if d := diagText(r); d != "" {
		t.Fatalf("unexpected diagnostics:\n%s", d)
	}
	loot := emitted(t, r, "data/demo/loot_table/chest.json")
	for _, want := range []string{`"name": "minecraft:written_book"`, `"generation": 3`, `"title": "Guide"`, `"rarity": "rare"`} {
		if !strings.Contains(loot, want) {
			t.Errorf("loot table lacks %s:\n%s", want, loot)
		}
	}
	pages := pagesOf(t, loot)
	if len(pages) != 2 {
		t.Fatalf("pages = %d", len(pages))
	}
	p1, _ := json.Marshal(pages[0])
	for _, want := range []string{
		`{"bold":true,"text":"Welcome"}`, `"\n"`, `{"bold":true,"text":"Bold"}`, `{"italic":true,"text":"it"}`,
		`{"text":"u","underlined":true}`, `{"strikethrough":true,"text":"s"}`, `{"obfuscated":true,"text":"o"}`,
		`{"bold":true,"color":"gold","text":"gold"}`, `{"color":"#FF8800","text":"hex"}`, `{"shadow_color":0,"text":"flat"}`,
		`{"keybind":"key.jump"}`, `{"atlas":"minecraft:items","object":"atlas","sprite":"minecraft:item/diamond","type":"object"}`,
		`{"object":"player","player":{"name":"Notch"},"type":"object"}`, `{"translate":"block.minecraft.stone"}`, `*plain*`,
		`{"score":{"name":"@s","objective":"kills"}}`, `{"score":{"name":"$wave","objective":"mms.demo"}}`,
		`{"nbt":"g.motto","storage":"demo:mms"}`, `2"`,
	} {
		if !strings.Contains(string(p1), want) {
			t.Errorf("page 1 lacks %s:\n%s", want, p1)
		}
	}
	p2, _ := json.Marshal(pages[1])
	for _, want := range []string{
		`"click_event":{"action":"change_page","page":1}`, `"hover_event":{"action":"show_text","value":"Go back"}`,
		`{"action":"run_command","command":"function demo:go"}`, `{"action":"suggest_command","command":"/say hi"}`,
		`{"action":"copy_to_clipboard","value":"x"}`, `{"hover_event":{"action":"show_text","value":"hello"},"text":"Hint"}`,
		`{"action":"open_url","url":"https://minecraft.wiki"}`,
	} {
		if !strings.Contains(string(p2), want) {
			t.Errorf("page 2 lacks %s:\n%s", want, p2)
		}
	}
	if recipe := emitted(t, r, "data/demo/recipe/copy.json"); !strings.Contains(recipe, `"id": "minecraft:written_book"`) || !strings.Contains(recipe, `"written_book_content"`) {
		t.Errorf("recipe result:\n%s", recipe)
	}
	if fn := emitted(t, r, "data/demo/function/go.mcfunction"); !strings.HasPrefix(fn, `give @s written_book[written_book_content={pages:[`) || !strings.HasSuffix(strings.TrimSpace(fn), `,rarity="rare"]`) {
		t.Errorf("give:\n%s", fn)
	}
}

func TestBookErrors(t *testing.T) {
	cases := []struct{ text, want string }{
		{`"**open"`, "this ** is not closed"},
		{`"{purpl}x{/}"`, "'purpl' is not a style"},
		{`"{/}"`, "{/} closes a {...} style, but none is open"},
		{`"{gold}never"`, "this style is never closed with {/}"},
		{`"{key:jumpp}"`, "'jumpp' is not a key binding; did you mean 'jump'?"},
		{`"[x](page:3)"`, "page 3 does not exist: the book has 1"},
		{`"[x](page:zero)"`, "page:zero is not a page number"},
		{`"[x](run:/nope)"`, "unknown command 'nope'"},
		{`"[x](run:/function demo:missing)"`, "no function named 'demo:missing' is declared"},
		{`"[x](ftp:x)"`, "'ftp:x' is not a link target"},
		{`"${nope}"`, "unknown variable 'nope'"},
		{`"${1 + w}"`, "unknown variable 'w'"},
		{`"{x"`, "this { is not closed"},
	}
	for _, tc := range cases {
		src := "Book b = { title = \"T\", author = \"A\", text = [" + tc.text + "] }\nFunction f {\n  /say hi\n}\n"
		_, r := bookProject(t, src, nil)
		d := diagText(r)
		if !strings.Contains(d, tc.want) {
			t.Errorf("%s: diagnostics %q lack %q", tc.text, d, tc.want)
		}
		if strings.Count(d, "\n") != 1 {
			t.Errorf("%s: want one diagnostic, got:\n%s", tc.text, d)
		}
	}
	// Problems point into the string.
	src := "Book b = { title = \"T\", author = \"A\", text = [\"ok {key:jumpp}\"] }\n"
	_, r := bookProject(t, src, nil)
	ds := r.Diagnostics("main.mms")
	if len(ds) != 1 || string([]rune(src)[ds[0].Sp.Start:ds[0].Sp.End]) != "jumpp" {
		t.Errorf("diagnostics = %+v", ds)
	}
	for _, tc := range []struct{ src, want string }{
		{`Book b = { title = "This title is far too long for a book", author = "A" }`, "32"},
		{`Book b = { title = "T", author = "A", text = ["a"], pages = ["b"] }`, "a book takes text or pages, not both"},
		{`Book b = { title = "T", author = "A", components = { written_book_content = { title = "x", author = "y" } } }`, "writes written_book_content itself"},
		{`Book b = { title = "T", author = "A", text = BookText("b.md") }`, "BookText reads a .mmsbook file"},
		{`Book b = { title = "T", author = "A", text = BookText("missing.mmsbook") }`, "no file 'missing.mmsbook'"},
	} {
		_, r := bookProject(t, tc.src+"\n", nil)
		if d := diagText(r); !strings.Contains(d, tc.want) {
			t.Errorf("%s: diagnostics %q lack %q", tc.src, d, tc.want)
		}
	}
}

func TestBookOverflow(t *testing.T) {
	// The first line wraps to three: 50 letters of 6 pixels, 114 a line.
	lines := []string{`"` + strings.Repeat("aaaa ", 10) + `"`}
	for i := 4; i <= 14; i++ {
		lines = append(lines, `"line"`)
	}
	lines = append(lines, `"cut"`, `"---"`, `"fine"`)
	src := "Book b = { title = \"T\", author = \"A\", text = [" + strings.Join(lines, ", ") + "] }\n"
	_, r := bookProject(t, src, nil)
	var got []syntax.Diagnostic
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/book-overflow" {
			got = append(got, d)
		}
	}
	if len(got) != 1 || !strings.Contains(got[0].Message, "page 1 runs to 15 lines") {
		t.Fatalf("overflow = %+v", got)
	}
	// It points at the first hidden line.
	if at := string([]rune(src)[got[0].Sp.Start:]); !strings.HasPrefix(at, `cut"`) {
		t.Errorf("points at %q", at[:10])
	}
	// Exactly 14 lines fit.
	src = "Book b = { title = \"T\", author = \"A\", text = [" + strings.Repeat(`"x", `, 13) + `"x"] }` + "\n"
	_, r = bookProject(t, src, nil)
	for _, d := range r.Diagnostics("main.mms") {
		t.Errorf("unexpected %s", d.Message)
	}
}

func TestBookFile(t *testing.T) {
	const src = "Book diary = { title = \"Diary\", author = \"A\", text = BookText(\"books/diary.mmsbook\") }\nLootTable l { pools = [Pool(1, [diary])] }\n"
	p, r := bookProject(t, src, map[string]string{"books/diary.mmsbook": "# Day 1\nThe {key:jumpp} lamps\n"})
	if d := diagText(r); d != "" {
		t.Errorf("main.mms has diagnostics:\n%s", d)
	}
	// Problems are shown in the book file, where they are.
	ds := r.Diagnostics("books/diary.mmsbook")
	text := []rune("# Day 1\nThe {key:jumpp} lamps")
	if len(ds) != 1 || string(text[ds[0].Sp.Start:ds[0].Sp.End]) != "jumpp" || strings.Contains(ds[0].Message, "diary.mmsbook:") {
		t.Fatalf("book file diagnostics = %+v", ds)
	}
	if !r.HasErrors() {
		t.Error("an error in the book file does not stop the build")
	}
	// Editing the file checks the Book and its uses again.
	p.SetFile("books/diary.mmsbook", "# Day 1\n{gold}The{/} lamps\n")
	r, _ = p.Analyze(context.Background())
	if ds := r.Diagnostics("books/diary.mmsbook"); len(ds) != 0 {
		t.Fatalf("after the fix: %+v", ds)
	}
	if loot := emitted(t, r, "data/demo/loot_table/l.json"); !strings.Contains(loot, `"color": "gold"`) {
		t.Errorf("loot table uses the old text:\n%s", loot)
	}
	// Overflow warnings go to the file too.
	p.SetFile("books/diary.mmsbook", strings.Repeat("line\n", 15))
	r, _ = p.Analyze(context.Background())
	ds = r.Diagnostics("books/diary.mmsbook")
	if len(ds) != 1 || ds[0].Code != "lint/book-overflow" || ds[0].Sp.Start != 70 {
		t.Fatalf("overflow in the file = %+v", ds)
	}
}

// A host that names files by URI, like the language server, finds the
// book file by the URI that ends in its path.
func TestBookFileByURI(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.Root = "file:///pack/"
	p.SetFile("file:///pack/main.mms", "Book b = { title = \"T\", author = \"A\", text = BookText(\"b.mmsbook\") }\n")
	p.SetFile("file:///pack/b.mmsbook", "{nope}x{/}\n")
	r, _ := p.Analyze(context.Background())
	if ds := r.Diagnostics("file:///pack/b.mmsbook"); len(ds) != 1 || !strings.Contains(ds[0].Message, "'nope' is not a style") {
		t.Fatalf("diagnostics = %+v", ds)
	}
}
