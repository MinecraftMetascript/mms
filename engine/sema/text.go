package sema

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Text markup anywhere a text component goes: a string there is markup,
// the same as a book's text without pages (see bookmarkup.go). A string
// with no markup in it is the same string in the output, so plain names
// stay plain.

const tText = "::java::util::text::Text"

// markupRunes are the characters that can start markup.
const markupRunes = "*_~|{[$\\#\n"

// textString checks a string where a text component goes.
func textString(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	s := e.(*syntax.String)
	in := c.info(s)
	in.Type, in.typed = v, true
	in.Markup = true
	if !strings.ContainsAny(s.Value, markupRunes) {
		return jsonv.String(s.Value)
	}
	src := c.stringSrc(s)
	src.noPages = true
	val, accidental := c.textMarkup(src)
	in.Accidental = accidental
	return val
}

// textFile checks TextFile("path") where a text component goes: the
// markup of a .mmstext file, or of a .mmsbook file without pages.
func textFile(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	call := e.(*syntax.Call)
	c.info(call).Type = v
	src := c.bookFile(call)
	if src == nil {
		return nil
	}
	src.noPages = true
	val, _ := c.textMarkup(src)
	return val
}

// textMarkup turns markup outside a book into a text component. A style
// mark left open, such as the * of "5 * 3", is the character itself; its
// spans are returned for the accidental-markup lint.
func (c *Checker) textMarkup(src *bookSrc) (jsonv.Value, []syntax.Span) {
	return c.textMarkupIn(src, nil, nil)
}

// textMarkupIn is textMarkup for text shown from function f, whose
// variables ${...} can read; values it computes are worked out in pre.
func (c *Checker) textMarkupIn(src *bookSrc, f *fc, pre *[]fn.Op) (jsonv.Value, []syntax.Span) {
	// A first pass, whose problems are dropped, finds the marks left open.
	trial := newBookParser(c.bookSub())
	trial.plainText = true
	trial.text(src)
	lit := map[int]bool{}
	var accidental []syntax.Span
	for m, at := range trial.toggles {
		n := len([]rune(m))
		for k := 0; k < n; k++ {
			lit[at.off+k] = true
		}
		if src.extern == "" {
			accidental = append(accidental, src.at(at.off, at.off+n))
		}
	}
	p := newBookParser(c)
	p.fn, p.pre = f, pre
	p.plainText = true
	p.literal = map[*bookSrc]map[int]bool{src: lit}
	p.text(src)
	p.finish()
	if len(p.pages) == 0 {
		return jsonv.String(""), accidental
	}
	return p.pages[0].json(), accidental
}
