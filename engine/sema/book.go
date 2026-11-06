package sema

import (
	"hash/fnv"
	"io"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Book text: the markup of a Book (prelude/books.mms), laid out as pages
// (see bookmarkup.go). Its name stands for the written book, like a
// CustomItem's: in item stacks, loot entries and ${...} in commands.

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
	// Where is the TextFile file position for a message, when the
	// problem is reported at the call.
	Where string
	// File is the .mmsbook file Sp is in, "" for the declaration's file.
	File string
}

// IsBookFile reports whether a project file holds markup: a .mmsbook
// file, whose `---` lines break pages, or a .mmstext file, which is text
// for anywhere and has no pages.
func IsBookFile(path string) bool {
	return strings.HasSuffix(path, BookFileExt) || strings.HasSuffix(path, TextFileExt)
}

// BookFileExt is the extension of book text files, and TextFileExt that
// of text files for anywhere.
const (
	BookFileExt = ".mmsbook"
	TextFileExt = ".mmstext"
)

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

// bookTextType is the type of a Book's text field, BookText in the
// prelude. Completion knows it by identity to offer TextFile("...")
// there, and the checker to lay out the text.
var bookTextType = &spec.Type{Kind: "list", Item: &spec.Type{Kind: "string"}}

// BookTextType is BookText, for the type checker.
func BookTextType() *spec.Type { return bookTextType }

// bookTextDoc documents TextFile in completion.
const bookTextDoc = "```mms\nTextFile(path: string)\n```\n\nReads text markup from a file in the project: a `.mmstext` file, or a `.mmsbook` file whose `---` lines break pages, such as `TextFile(\"books/guide.mmsbook\")`. Pages work only in a book's text. The editor checks and completes the markup in that file."

// bookMarkupDoc is the short markup reference in hovers.
const bookMarkupDoc = "Markup: `# heading`, `**bold**`, `*italic*`, `__underlined__`, `~~struck~~`, `||obfuscated||`, " +
	"`{gold}...{/}` (colors, `#ff8800`, `bold`, `shadow:#000000`, `font:uniform`, combined as `{bold red}`), " +
	"`{key:jump}`, `{sprite:item/diamond}`, `{head:Notch}`, `{t:translation.key}`, " +
	"`[text](page:2)` (also `run:/command`, `suggest:/command`, `copy:text`, `hover:text`, `https://...`; add `\"hover text\"` after the target), " +
	"`${kills[@s]}` for a score or Global, `\\` before a character to write it as is."

// BookLayout is a book's text laid out as pages, for the editor and the
// book-page-overflow lint.
type BookLayout struct {
	// Overflows are pages that run past the bottom of the page.
	Overflows []BookOverflow
	// PageStarts are where each page starts, so an editor can open its
	// preview on the page at the cursor.
	PageStarts []BookPageStart
}

// layoutBook lays out book text, a list of markup lines or TextFile(...),
// as pages. It returns nil when e is neither.
func (c *Checker) layoutBook(e syntax.Expr) (jsonv.Array, *BookLayout) {
	p := newBookParser(c)
	switch x := e.(type) {
	case *syntax.Call:
		if callName(x) != "TextFile" {
			return nil, nil
		}
		c.info(x).Type = tview{T: bookTextType}
		src := c.bookFile(x)
		if src == nil {
			return jsonv.Array{}, &BookLayout{}
		}
		p.text(src)
	case *syntax.List:
		for _, el := range x.Elems {
			s, ok := el.(*syntax.String)
			if !ok {
				c.errorf("mms/type-mismatch", el.Span(), "expected a line of book markup in quotes")
				continue
			}
			in := c.info(s)
			in.Type, in.typed, in.Markup = tview{T: &spec.Type{Kind: "string"}}, true, true
			p.text(c.stringSrc(s))
		}
	default:
		return nil, nil
	}
	p.finish()
	out := &BookLayout{}
	var pages jsonv.Array
	for i, pg := range p.pages {
		pages = append(pages, pg.json())
		if pg.src != nil {
			out.PageStarts = append(out.PageStarts, BookPageStart{File: pg.src.extern, Sp: pg.src.at(pg.off, pg.off+1)})
		} else {
			out.PageStarts = append(out.PageStarts, BookPageStart{Sp: e.Span()})
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
			ov.Sp = e.Span()
		}
		out.Overflows = append(out.Overflows, ov)
	}
	if pages == nil {
		pages = jsonv.Array{}
	}
	return pages, out
}

// bookTextRule checks a Book's text where it is written: its markup
// problems, its layout for the editor, and the markup Info. The value
// is the text as written; book_pages lays it out where the pages go.
func bookTextRule(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	// Its problems are kept apart: a Book instance's body is checked for
	// the editor only, and keeps just these (see project.instanceDecls).
	sub := c.bookSub()
	sub.Src, sub.Info = c.Src, c.Info
	pages, layout := sub.layoutBook(e)
	c.BookDiags = append(c.BookDiags, sub.Diags...)
	for f, ds := range sub.Extern {
		if c.Extern == nil {
			c.Extern = map[string][]syntax.Diagnostic{}
		}
		c.Extern[f] = append(c.Extern[f], ds...)
	}
	if layout == nil {
		c.errorf("mms/type-mismatch", e.Span(), "expected the book's text: a list of lines such as [\"# Welcome\", \"Hello!\"], or TextFile(\"books/guide.mmsbook\")")
		return nil
	}
	c.Book = layout
	return pages
}

// bookPagesRule lays out book_pages(text) where a book's pages go. The
// text's problems were reported where it is written (see bookTextRule).
func bookPagesRule(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	call := e.(*syntax.Call)
	if len(call.Args) != 1 {
		c.errorf("mms/arity", call.Sp, "book_pages takes the book's text")
		return nil
	}
	sub := c.bookSub()
	sub.Src = c.Src
	pages, _ := sub.layoutBook(call.Args[0].Value)
	return pages
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

// bookFile reads TextFile("path"): a .mmstext or .mmsbook file in the
// project.
func (c *Checker) bookFile(call *syntax.Call) *bookSrc {
	if len(call.Args) != 1 || call.Args[0].Name != nil {
		c.errorf("mms/type-mismatch", call.Sp, "TextFile takes one file path: TextFile(\"text/rules.mmstext\")")
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
		c.errorf("mms/asset", s.Sp, "TextFile reads a %s or %s file, not '%s'", TextFileExt, BookFileExt, path.Base(src))
		return nil
	}
	noPages := !strings.HasSuffix(src, BookFileExt)
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
			return &bookSrc{text: text, extern: unit, noPages: noPages, at: func(start, end int) syntax.Span {
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
	return &bookSrc{text: text, file: src, lines: lines, noPages: noPages, at: func(int, int) syntax.Span { return s.Sp }}
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
