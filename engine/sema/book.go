package sema

import (
	"encoding/json"
	"hash/fnv"
	"io"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Books: a written book whose pages are written in a small markup (see
// bookmarkup.go) rather than as text components.
//
//	Book guide = {
//	  title = "Cave Siege Guide"
//	  author = "The Garrison"
//	  text = [
//	    "# Welcome",
//	    "Plant the **Siege Banner** on gold.",
//	    "---",
//	    "[Start](run:/function siege:start)",
//	  ]
//	}
//
// Its name stands for the written book, like a CustomItem's: in item
// stacks, loot entries and ${...} in commands. It writes no files.

const bookKind = "Book"

// isItemKind reports whether a language kind's name stands for an item
// stack: a CustomItem or a Book.
func isItemKind(k string) bool { return k == bookKind }

// BookPageStart is where a page of a Book's text starts: File is the
// .mmsbook unit it is in, empty for the declaring file.
type BookPageStart struct {
	File string
	Sp   syntax.Span
}

// BookOverflow is a page that runs past the bottom of the book screen.
type BookOverflow struct {
	Page, Lines int
	// Sp is where the first hidden line starts, or the text's span.
	Sp syntax.Span
	// Where is the BookText file position for a message, when the
	// problem is reported at the call.
	Where string
	// File is the .mmsbook file Sp is in, "" for the declaration's file.
	File string
}

// IsBookFile reports whether a project file holds book markup.
func IsBookFile(path string) bool { return strings.HasSuffix(path, BookFileExt) }

// BookFileExt is the extension of book markup files.
const BookFileExt = ".mmsbook"

// bookSub is a checker for parts of book text, sharing c's lookups.
func (c *Checker) bookSub() *Checker {
	sub := NewChecker(c.Pack, c.Kinds, c.Syms, c.NS)
	sub.Files, sub.Texts, sub.Deps = c.Files, c.Texts, c.Deps
	return sub
}

// TextDepsHold reports whether every .mmsbook text a check read is
// still the same.
func TextDepsHold(texts func(string) (string, string, bool), deps map[string]bool) bool {
	for k := range deps {
		rest, ok := strings.CutPrefix(k, "t|")
		if !ok {
			continue
		}
		i := strings.LastIndexByte(rest, '|')
		src, sig := rest[:i], rest[i+1:]
		cur := ""
		if texts != nil {
			if _, t, ok := texts(src); ok {
				cur = textSig(t)
			}
		}
		if cur != sig {
			return false
		}
	}
	return true
}

func textSig(t string) string {
	h := fnv.New64a()
	h.Write([]byte(t))
	return strconv.FormatUint(h.Sum64(), 36)
}

var bookTypes sync.Map // *spec.Pack -> *spec.Type

var bookGenerations = []string{"original", "copy", "copy_of_copy", "tattered"}

func bookType(p *spec.Pack) *spec.Type {
	if t, ok := bookTypes.Load(p); ok {
		return t.(*spec.Type)
	}
	str := &spec.Type{Kind: "string"}
	gen := &spec.Type{Kind: "enum", EnumKind: "string"}
	descs := []string{"The original: it can be copied.", "A copy of the original.", "A copy of a copy: it cannot be copied again.", "Tattered: it cannot be copied, and shows as worn."}
	for i, g := range bookGenerations {
		v, _ := json.Marshal(g)
		gen.Values = append(gen.Values, &spec.EnumValue{Identifier: g, Value: v, Desc: descs[i]})
	}
	text := &spec.Type{Kind: "reference", Path: "::java::util::text::Text"}
	t := &spec.Type{Kind: "struct", Fields: []*spec.Field{
		{Kind: "pair", Key: "title", Type: &spec.Type{Kind: "string", LengthRange: &spec.Range{Max: ptr(32.0)}},
			Desc: "The title on the cover and the item's name, at most 32 characters."},
		{Kind: "pair", Key: "author", Type: str,
			Desc: "The author shown under the title: \"by <author>\"."},
		{Kind: "pair", Key: "generation", Optional: true, Type: gen,
			Desc: "How many times the book has been copied: `original` (the default), `copy`, `copy_of_copy` or `tattered`. The last two cannot be copied."},
		{Kind: "pair", Key: "text", Optional: true, Type: bookTextType,
			Desc: "The book's text in book markup, one string per line, or `BookText(\"books/guide.mmsbook\")` to read it from a file. A line of `---` starts the next page.\n\n" + bookMarkupDoc},
		{Kind: "pair", Key: "pages", Optional: true, Type: &spec.Type{Kind: "list", Item: text},
			Desc: "The pages as text components, for full control. Use `text` instead for book markup."},
		{Kind: "pair", Key: "components", Optional: true, Type: &spec.Type{Kind: "reference", Path: "::java::world::component::DataComponentPatch"},
			Desc: "More item components, such as `lore` or `custom_data`. Every use of the book gets them."},
	}}
	got, _ := bookTypes.LoadOrStore(p, t)
	return got.(*spec.Type)
}

// bookTextType is the type of a Book's text field. Completion knows it by
// identity to offer BookText("...") there.
var bookTextType = &spec.Type{Kind: "list", Item: &spec.Type{Kind: "string"}}

// bookTextDoc documents BookText in completion.
const bookTextDoc = "```mms\nBookText(path: string)\n```\n\nReads the book's text from a `.mmsbook` file in the project, such as `BookText(\"books/guide.mmsbook\")`. The editor checks and completes the markup in that file."

func ptr[T any](v T) *T { return &v }

// bookMarkupDoc is the short markup reference in hovers.
const bookMarkupDoc = "Markup: `# heading`, `**bold**`, `*italic*`, `__underlined__`, `~~struck~~`, `||obfuscated||`, " +
	"`{gold}...{/}` (colors, `#ff8800`, `bold`, `shadow:#000000`, `font:uniform`, combined as `{bold red}`), " +
	"`{key:jump}`, `{sprite:item/diamond}`, `{head:Notch}`, `{t:translation.key}`, " +
	"`[text](page:2)` (also `run:/command`, `suggest:/command`, `copy:text`, `hover:text`, `https://...`; add `\"hover text\"` after the target), " +
	"`${kills[@s]}` for a score or Global, `\\` before a character to write it as is."

// Book checks a Book declaration. id is its "ns:path". The result has
// only Stack set, and Overflows for the page-overflow lint.
func (c *Checker) Book(d *syntax.Decl, id string) *CustomItemOut {
	const usage = "Book needs a body: Book guide = { title = \"Guide\", author = \"Me\", text = [\"Hello!\"] }"
	body := customItemBody(d)
	if body == nil || d.Name == nil {
		if d.Value != nil {
			c.errorf("mms/type-mismatch", d.Value.Span(), "%s", usage)
			return nil
		}
		c.errorf("mms/missing-value", d.Kind.Sp, "%s", usage)
		return nil
	}
	fields := customItemFields(d)
	// BookText(...) is read here; the rest of the body is checked as usual.
	checked := body
	var fileText *syntax.Call
	if call, ok := fields["text"].(*syntax.Call); ok {
		if fnID, ok := call.Fn.(*syntax.Ident); ok && fnID.Name == "BookText" {
			fileText = call
			cp := *body
			cp.Items = nil
			for _, it := range body.Items {
				if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil && f.Key.Text == "text" {
					continue
				}
				cp.Items = append(cp.Items, it)
			}
			checked = &cp
		}
	}
	prevOf := c.customOf
	c.customOf = id
	v := c.check(&syntax.Object{Body: checked, Sp: body.Sp}, tview{T: bookType(c.Pack), Name: bookKind})
	c.customOf = prevOf
	o, ok := v.(*jsonv.Object)
	if !ok {
		return nil
	}
	if fields["text"] != nil && fields["pages"] != nil {
		c.errorf("mms/book", fields["pages"].Span(), "a book takes text or pages, not both")
		return nil
	}

	p := newBookParser(c)
	switch {
	case fileText != nil:
		src := c.bookFile(fileText)
		if src == nil {
			return nil
		}
		p.text(src)
	case fields["text"] != nil:
		list, ok := fields["text"].(*syntax.List)
		if !ok {
			break // the type check reported it
		}
		for _, e := range list.Elems {
			s, ok := e.(*syntax.String)
			if !ok {
				continue
			}
			p.text(c.stringSrc(s))
		}
	}
	p.finish()

	out := &CustomItemOut{}
	var pages jsonv.Array
	if fields["text"] != nil {
		for i, pg := range p.pages {
			pages = append(pages, pg.json())
			if pg.src != nil {
				out.PageStarts = append(out.PageStarts, BookPageStart{File: pg.src.extern, Sp: pg.src.at(pg.off, pg.off+1)})
			} else {
				out.PageStarts = append(out.PageStarts, BookPageStart{Sp: fields["text"].Span()})
			}
			lines, cut := layoutPage(pg.glyphs())
			if lines <= bookLines {
				continue
			}
			ov := BookOverflow{Page: i + 1, Lines: lines}
			gs := pg.glyphs()
			at := bookPos{src: pg.src, off: pg.off}
			for j := cut; j >= 0 && j < len(gs); j++ {
				if gs[j].at.off >= 0 {
					at = gs[j].at
					break
				}
			}
			if at.src != nil {
				ov.Sp = at.src.at(at.off, at.off+1)
				ov.Where = at.src.where(at.off)
				ov.File = at.src.extern
			} else {
				ov.Sp = fields["text"].Span()
			}
			out.Overflows = append(out.Overflows, ov)
		}
	} else if ps, ok := o.Get("pages").(jsonv.Array); ok {
		pages = ps
	}
	content := &jsonv.Object{}
	if len(pages) > 0 {
		content.Set("pages", pages)
	}
	content.Set("title", o.Get("title"))
	content.Set("author", o.Get("author"))
	if g, ok := o.Get("generation").(jsonv.String); ok {
		for i, name := range bookGenerations {
			if string(g) == name && i > 0 {
				content.Set("generation", jsonv.Number(string(rune('0'+i))))
			}
		}
	}
	comps := &jsonv.Object{}
	user, _ := o.Get("components").(*jsonv.Object)
	if hasComponent(user, "written_book_content") {
		c.errorf("mms/book", fields["components"].Span(), "the book writes written_book_content itself; set title, author and text instead")
	} else {
		comps.Set("written_book_content", content)
	}
	if user != nil {
		comps.Members = append(comps.Members, user.Members...)
	}
	stack := &jsonv.Object{}
	stack.Set("id", jsonv.String("minecraft:written_book"))
	stack.Set("components", comps)
	out.Stack = stack
	return out
}

// stringSrc is the markup source of a string literal. Problems point
// into the string when its text is as written; with escapes they point
// at the whole string.
func (c *Checker) stringSrc(s *syntax.String) *bookSrc {
	val := []rune(s.Value)
	raw := []rune(s.Raw)
	exact := len(raw) == len(val)+2 && string(raw[1:len(raw)-1]) == s.Value
	return &bookSrc{text: val, at: func(start, end int) syntax.Span {
		if !exact || start < 0 || end > len(val) || end <= start {
			return s.Sp
		}
		return syntax.Span{Start: s.Sp.Start + 1 + start, End: s.Sp.Start + 1 + end}
	}}
}

// bookFile reads BookText("path"): a text file in the project.
func (c *Checker) bookFile(call *syntax.Call) *bookSrc {
	if len(call.Args) != 1 || call.Args[0].Name != nil {
		c.errorf("mms/type-mismatch", call.Sp, "BookText takes one file path: BookText(\"books/guide.md\")")
		return nil
	}
	s, ok := call.Args[0].Value.(*syntax.String)
	if !ok {
		c.errorf("mms/type-mismatch", call.Args[0].Value.Span(), "expected a quoted file path")
		return nil
	}
	src := path.Clean(strings.ReplaceAll(s.Value, "\\", "/"))
	if path.IsAbs(src) || src == ".." || strings.HasPrefix(src, "../") {
		c.errorf("mms/asset", s.Sp, "file paths are relative to the project folder and stay inside it")
		return nil
	}
	if !IsBookFile(src) {
		c.errorf("mms/asset", s.Sp, "BookText reads a %s file, not '%s'", BookFileExt, path.Base(src))
		return nil
	}
	if c.Texts != nil {
		unit, t, ok := c.Texts(src)
		if c.Deps != nil {
			sig := ""
			if ok {
				sig = textSig(t)
			}
			c.Deps["t|"+src+"|"+sig] = ok
		}
		if ok {
			// Problems are shown in the file itself.
			text := []rune(strings.TrimRight(t, "\r\n"))
			return &bookSrc{text: text, extern: unit, at: func(start, end int) syntax.Span {
				if end <= start {
					end = start + 1
				}
				return syntax.Span{Start: start, End: min(end, max(len(text), start+1))}
			}}
		}
	}
	data, ok := c.readFile(src)
	if !ok && c.Files == nil {
		c.errorf("mms/asset", s.Sp, "project files are not available here yet; build with the mms command line to read '%s'", src)
		return nil
	}
	if !ok {
		c.errorf("mms/asset", s.Sp, "no file '%s' in the project", src)
		return nil
	}
	text := []rune(strings.TrimRight(string(data), "\r\n"))
	lines := []int{0}
	for i, r := range text {
		if r == '\n' {
			lines = append(lines, i+1)
		}
	}
	return &bookSrc{text: text, file: src, lines: lines, at: func(int, int) syntax.Span { return s.Sp }}
}

// readFile reads a whole project file, recording it in Deps like an
// asset so that editing the file checks its users again.
func (c *Checker) readFile(src string) ([]byte, bool) {
	var data []byte
	ok, sig := false, ""
	if c.Files != nil {
		if f, err := c.Files.Open(src); err == nil {
			if st, err := f.Stat(); err == nil && !st.IsDir() {
				if b, err := io.ReadAll(f); err == nil {
					ok, sig, data = true, fileSig(st), b
				}
			}
			f.Close()
		}
	}
	if c.Deps != nil {
		c.Deps["a|"+src+"|"+sig] = ok
	}
	return data, ok
}

func (c *Checker) specEnum(path string) []string {
	t := c.Pack.Resolve(c.Pack.Types[path])
	if t == nil {
		return nil
	}
	var out []string
	for _, v := range t.Values {
		out = append(out, v.String())
	}
	return out
}

func (c *Checker) keybindNames() []string {
	var out []string
	for _, k := range c.specEnum("::java::util::text::Keybind") {
		out = append(out, strings.TrimPrefix(k, "key."))
	}
	sort.Strings(out)
	return out
}

func (c *Checker) isKeybind(k string) bool {
	for _, v := range c.specEnum("::java::util::text::Keybind") {
		if v == k {
			return true
		}
	}
	return false
}

func (c *Checker) isTextColor(w string) bool {
	for _, v := range c.specEnum("::java::util::text::TextColor") {
		if v == w {
			return true
		}
	}
	return false
}

func (c *Checker) styleNames() []string {
	out := c.specEnum("::java::util::text::TextColor")
	for w := range bookStyleWords {
		out = append(out, w)
	}
	return out
}

func didYouMean(s string, cands []string) string {
	if b := closest(s, cands); b != "" {
		return "; did you mean '" + b + "'?"
	}
	return ""
}
