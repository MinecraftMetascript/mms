package sema

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Book markup: the text of a Book declaration, turned into the pages of
// a written_book_content component.
//
//	# Day 12                      a bold line
//	**bold** *italic* __underlined__ ~~struck~~ ||obfuscated||
//	{gold}gold text{/}            colors, styles, shadow:, font:, combined
//	{key:jump}                    the key the reader has bound to jump
//	{sprite:item/diamond}         an inline icon; {head:Notch} a player head
//	{t:block.minecraft.stone}     a translated text
//	[Next](page:2)                a link: page:N, run:/command,
//	                              suggest:/command, copy:text, hover:text
//	                              or https://...; [text](target "hover")
//	${kills[@s]}                  a score, a Global or a constant
//	---                           a line of its own: the next page
//
// A backslash makes the next character plain text.

// bookSrc is a piece of markup and where it came from.
type bookSrc struct {
	text []rune
	// at maps a rune range of text to the span a problem is reported at.
	at func(start, end int) syntax.Span
	// file is the BookText file the text came from, "" for a string in
	// the declaration. Problems in a file are reported at the BookText
	// call with the file position in the message.
	file string
	// lines holds the offset each line of a file starts at.
	lines []int
	// extern is the path of the .mmsbook file the text is, when its
	// problems are reported in that file rather than at the call.
	extern string
}

// where is the file position a message starts with, for file sources.
func (s *bookSrc) where(off int) string {
	if s.file == "" {
		return ""
	}
	line := 0
	for i, st := range s.lines {
		if st <= off {
			line = i
		}
	}
	return fmt.Sprintf("%s:%d:%d: ", s.file, line+1, off-s.lines[line]+1)
}

type bookStyle struct {
	color, font                                         string
	shadow                                              *int64
	bold, italic, underlined, strikethrough, obfuscated bool
}

// bookRun is a piece of a page with one style: text, or a component such
// as a keybind or a score.
type bookRun struct {
	style bookStyle
	text  string
	obj   *jsonv.Object // the component without its style, when not text
	width int           // obj's width in pixels
	click *jsonv.Object
	hover string
	glyph []bookGlyph
}

type bookPage struct {
	runs []*bookRun
	src  *bookSrc
	off  int // where the page starts
}

type bookMark struct {
	src *bookSrc
	off int
}

type bookGroup struct {
	style bookStyle
	at    bookMark
}

// bookLink is a page:N link, checked once the page count is known.
type bookLink struct {
	page int
	at   bookMark
	end  int
}

type bookParser struct {
	c       *Checker
	pages   []*bookPage
	groups  []bookGroup
	toggles map[string]bookMark
	links   []bookLink
	// lineStart is set at the start of each line; a heading line is bold.
	heading bool
	started bool // a line has been written on the current page
}

func newBookParser(c *Checker) *bookParser {
	return &bookParser{c: c, toggles: map[string]bookMark{}}
}

func (p *bookParser) errorf(src *bookSrc, start, end int, format string, args ...any) {
	p.report(src, "mms/book", src.at(start, end), src.where(start)+fmt.Sprintf(format, args...))
}

// report adds an error at sp, a span src.at gave.
func (p *bookParser) report(src *bookSrc, code string, sp syntax.Span, msg string) {
	if src.extern == "" {
		p.c.errorf(code, sp, "%s", msg)
		return
	}
	c := p.c
	c.errs++
	if c.trial {
		return
	}
	if c.Extern == nil {
		c.Extern = map[string][]syntax.Diagnostic{}
	}
	c.Extern[src.extern] = append(c.Extern[src.extern], syntax.Diagnostic{Code: code, Severity: syntax.SeverityError, Message: msg, Sp: sp})
}

func (p *bookParser) page() *bookPage {
	if len(p.pages) == 0 {
		p.pages = append(p.pages, &bookPage{})
	}
	return p.pages[len(p.pages)-1]
}

// style is the style text gets now: the innermost {...} group and the
// open ** and similar marks.
func (p *bookParser) style() bookStyle {
	var s bookStyle
	if n := len(p.groups); n > 0 {
		s = p.groups[n-1].style
	}
	for m := range p.toggles {
		switch m {
		case "**":
			s.bold = true
		case "*":
			s.italic = true
		case "__":
			s.underlined = true
		case "~~":
			s.strikethrough = true
		case "||":
			s.obfuscated = true
		}
	}
	if p.heading {
		s.bold = true
	}
	return s
}

// add appends text at offset off of src to the current page.
func (p *bookParser) add(src *bookSrc, off int, text string, click *jsonv.Object, hover string) {
	pg := p.page()
	st := p.style()
	var r *bookRun
	if n := len(pg.runs); n > 0 {
		last := pg.runs[n-1]
		if last.obj == nil && last.style == st && sameClick(last.click, click) && last.hover == hover {
			r = last
		}
	}
	if r == nil {
		r = &bookRun{style: st, click: click, hover: hover}
		pg.runs = append(pg.runs, r)
	}
	r.text += text
	for i, ch := range text {
		pos := bookPos{src: src, off: -1}
		if off >= 0 {
			pos.off = off + i
		}
		r.glyph = append(r.glyph, bookGlyph{w: glyphAdvance(ch, st.bold), space: ch == ' ', newline: ch == '\n', at: pos})
	}
}

// addObj appends a component such as a keybind to the current page.
func (p *bookParser) addObj(src *bookSrc, off int, obj *jsonv.Object, width int, click *jsonv.Object, hover string) {
	pg := p.page()
	r := &bookRun{style: p.style(), obj: obj, width: width, click: click, hover: hover}
	r.glyph = []bookGlyph{{w: width, at: bookPos{src: src, off: off}}}
	pg.runs = append(pg.runs, r)
}

func sameClick(a, b *jsonv.Object) bool {
	if a == nil || b == nil {
		return a == b
	}
	return string(jsonv.Marshal(a)) == string(jsonv.Marshal(b))
}

// text parses a whole source: lines separated by '\n'.
func (p *bookParser) text(src *bookSrc) {
	t := src.text
	start := 0
	for i := 0; i <= len(t); i++ {
		if i < len(t) && t[i] != '\n' {
			continue
		}
		end := i
		if end > start && t[end-1] == '\r' {
			end--
		}
		p.line(src, start, end)
		start = i + 1
	}
}

// line parses one line of markup, the runes [start, end) of src.
func (p *bookParser) line(src *bookSrc, start, end int) {
	t := src.text
	if strings.TrimSpace(string(t[start:end])) == "---" {
		p.pages = append(p.pages, &bookPage{src: src, off: start})
		p.started = false
		return
	}
	if len(p.pages) == 0 {
		p.pages = append(p.pages, &bookPage{src: src, off: start})
	}
	if p.started {
		p.add(src, -1, "\n", nil, "")
	}
	p.started = true
	if end-start >= 2 && t[start] == '#' && t[start+1] == ' ' {
		p.heading = true
		start += 2
	}
	p.inline(src, start, end, nil, "")
	p.heading = false
}

// marks are the style marks that open and close the same way, longest
// first so that ** wins over *.
var bookMarks = []string{"**", "__", "~~", "||", "*"}

// inline parses the runes [start, end) of src, adding them to the page
// with a click event and hover text when they are a link's text.
func (p *bookParser) inline(src *bookSrc, start, end int, click *jsonv.Object, hover string) {
	t := src.text
	var plain strings.Builder
	plainAt := -1
	flush := func() {
		if plain.Len() > 0 {
			p.add(src, plainAt, plain.String(), click, hover)
			plain.Reset()
		}
		plainAt = -1
	}
	lit := func(i int, r rune) {
		if plainAt < 0 {
			plainAt = i
		} else if plainAt+len([]rune(plain.String())) != i {
			// An escape skipped a character, so positions no longer line
			// up; start a new piece.
			flush()
			plainAt = i
		}
		plain.WriteRune(r)
	}
	has := func(i int, s string) bool {
		rs := []rune(s)
		if i+len(rs) > end {
			return false
		}
		for k, r := range rs {
			if t[i+k] != r {
				return false
			}
		}
		return true
	}
	for i := start; i < end; i++ {
		r := t[i]
		switch {
		case r == '\\' && i+1 < end:
			i++
			lit(i, t[i])
			continue
		case r == '$' && has(i, "${"):
			j := matchBrace(t, i+1, end)
			if j < 0 {
				p.errorf(src, i, end, "this ${ is not closed")
				return
			}
			flush()
			p.expr(src, i+2, j, click, hover)
			i = j
			continue
		case r == '{':
			j := indexRune(t, '}', i+1, end)
			if j < 0 {
				p.errorf(src, i, i+1, "this { is not closed; write \\{ for a brace")
				return
			}
			flush()
			p.tag(src, i, j, click, hover)
			i = j
			continue
		case r == '[' && click == nil:
			if close, tStart, tEnd, ok := linkAt(t, i, end); ok {
				flush()
				p.link(src, i, close, tStart, tEnd, hover)
				i = tEnd
				continue
			}
		}
		marked := false
		for _, m := range bookMarks {
			if has(i, m) {
				flush()
				if _, open := p.toggles[m]; open {
					delete(p.toggles, m)
				} else {
					p.toggles[m] = bookMark{src: src, off: i}
				}
				i += len(m) - 1
				marked = true
				break
			}
		}
		if !marked {
			lit(i, r)
		}
	}
	flush()
}

func indexRune(t []rune, r rune, from, end int) int {
	for i := from; i < end; i++ {
		if t[i] == '\\' {
			i++
			continue
		}
		if t[i] == r {
			return i
		}
	}
	return -1
}

// matchBrace returns the index of the '}' that closes the '{' at open.
func matchBrace(t []rune, open, end int) int {
	depth := 0
	for i := open; i < end; i++ {
		switch t[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

// linkAt finds a link `[text](target)` starting at i: the index of the
// ']' and of the target's '(' and ')'.
func linkAt(t []rune, i, end int) (close, tStart, tEnd int, ok bool) {
	close = indexRune(t, ']', i+1, end)
	if close < 0 || close+1 >= end || t[close+1] != '(' {
		return 0, 0, 0, false
	}
	depth, quoted := 0, false
	for j := close + 1; j < end; j++ {
		switch {
		case t[j] == '\\':
			j++
		case t[j] == '"':
			quoted = !quoted
		case quoted:
		case t[j] == '(':
			depth++
		case t[j] == ')':
			depth--
			if depth == 0 {
				return close, close + 1, j, true
			}
		}
	}
	return 0, 0, 0, false
}

// link handles `[text](target)`: [open, close] is the text in brackets
// and (tStart, tEnd) the target in parentheses.
func (p *bookParser) link(src *bookSrc, open, close, tStart, tEnd int, outerHover string) {
	t := src.text
	ts, te := tStart+1, tEnd
	for ts < te && t[ts] == ' ' {
		ts++
	}
	for te > ts && t[te-1] == ' ' {
		te--
	}
	hover := outerHover
	// A quoted title at the end is hover text, as in Markdown.
	if te > ts && t[te-1] == '"' {
		for q := te - 2; q > ts; q-- {
			if t[q] == '"' && t[q-1] == ' ' {
				hover = string(t[q+1 : te-1])
				te = q - 1
				for te > ts && t[te-1] == ' ' {
					te--
				}
				break
			}
		}
	}
	target := string(t[ts:te])
	kind, rest, _ := strings.Cut(target, ":")
	restAt := ts + len([]rune(kind)) + 1
	var click *jsonv.Object
	mk := func(action, key string, v jsonv.Value) {
		click = &jsonv.Object{}
		click.Set("action", jsonv.String(action))
		click.Set(key, v)
	}
	switch {
	case kind == "page":
		n, err := strconv.Atoi(strings.TrimSpace(rest))
		if err != nil || n < 1 {
			p.errorf(src, ts, te, "page:%s is not a page number; pages count from 1", rest)
			return
		}
		mk("change_page", "page", jsonv.Number(strconv.Itoa(n)))
		p.links = append(p.links, bookLink{page: n, at: bookMark{src: src, off: ts}, end: te})
	case kind == "run":
		cmd := strings.TrimPrefix(rest, "/")
		cmdAt := restAt + len([]rune(rest)) - len([]rune(cmd))
		if strings.TrimSpace(cmd) == "" {
			p.errorf(src, ts, te, "run: needs a command, such as run:/function demo:start")
			return
		}
		at := func(s, e int) syntax.Span {
			if e <= s {
				e = s + 1
			}
			return src.at(cmdAt+s, cmdAt+e)
		}
		if src.extern == "" {
			p.c.commandText(cmd, at)
		} else {
			sub := p.c.bookSub()
			sub.commandText(cmd, at)
			for _, d := range sub.Diags {
				p.report(src, d.Code, d.Sp, d.Message)
			}
		}
		mk("run_command", "command", jsonv.String(cmd))
	case kind == "suggest":
		mk("suggest_command", "command", jsonv.String(rest))
	case kind == "copy":
		mk("copy_to_clipboard", "value", jsonv.String(rest))
	case kind == "hover":
		hover = rest
	case kind == "https" || kind == "http":
		mk("open_url", "url", jsonv.String(target))
	default:
		p.errorf(src, ts, te, "'%s' is not a link target; use page:N, run:/command, suggest:/command, copy:text, hover:text or an https:// address", target)
		return
	}
	if click == nil {
		click = &jsonv.Object{} // hover only: still a link, for merging
	}
	p.inline(src, open+1, close, click, hover)
}

var bookStyleWords = map[string]string{
	"bold": "bold", "italic": "italic", "underlined": "underlined", "underline": "underlined",
	"strikethrough": "strikethrough", "strike": "strikethrough", "obfuscated": "obfuscated",
}

// tag handles `{...}` at [open, close].
func (p *bookParser) tag(src *bookSrc, open, close int, click *jsonv.Object, hover string) {
	body := string(src.text[open+1 : close])
	bodyAt := open + 1
	if strings.TrimSpace(body) == "/" {
		if len(p.groups) == 0 {
			p.errorf(src, open, close+1, "{/} closes a {...} style, but none is open")
			return
		}
		p.groups = p.groups[:len(p.groups)-1]
		return
	}
	key, val, hasVal := strings.Cut(body, ":")
	val = strings.TrimSpace(val)
	switch {
	case hasVal && key == "key":
		k := val
		if !strings.HasPrefix(k, "key.") {
			k = "key." + k
		}
		if !p.c.isKeybind(k) {
			p.errorf(src, bodyAt+4, close, "'%s' is not a key binding%s", val, didYouMean(strings.TrimPrefix(k, "key."), p.c.keybindNames()))
			return
		}
		o := &jsonv.Object{}
		o.Set("keybind", jsonv.String(k))
		p.addObj(src, open, o, keyAdvance(k, p.style().bold), click, hover)
	case hasVal && key == "sprite":
		fields := strings.Fields(val)
		if len(fields) == 0 {
			p.errorf(src, open, close+1, "sprite: needs a texture, such as {sprite:item/diamond}")
			return
		}
		sprite, atlas := fields[0], ""
		for _, f := range fields[1:] {
			if a, ok := strings.CutPrefix(f, "atlas:"); ok {
				atlas = a
			} else {
				p.errorf(src, open, close+1, "'%s' is not a sprite option; write {sprite:<texture> atlas:<atlas>}", f)
				return
			}
		}
		if atlas == "" && strings.HasPrefix(sprite, "item/") {
			atlas = "minecraft:items"
		}
		o := &jsonv.Object{}
		o.Set("type", jsonv.String("object"))
		o.Set("object", jsonv.String("atlas"))
		if atlas != "" {
			if !strings.Contains(atlas, ":") {
				atlas = "minecraft:" + atlas
			}
			o.Set("atlas", jsonv.String(atlas))
		}
		if !strings.Contains(sprite, ":") {
			sprite = "minecraft:" + sprite
		}
		o.Set("sprite", jsonv.String(sprite))
		p.addObj(src, open, o, 9, click, hover)
	case hasVal && key == "head":
		if val == "" {
			p.errorf(src, open, close+1, "head: needs a player name, such as {head:Notch}")
			return
		}
		o := &jsonv.Object{}
		o.Set("type", jsonv.String("object"))
		o.Set("object", jsonv.String("player"))
		pl := &jsonv.Object{}
		pl.Set("name", jsonv.String(val))
		o.Set("player", pl)
		p.addObj(src, open, o, 9, click, hover)
	case hasVal && key == "t":
		if val == "" {
			p.errorf(src, open, close+1, "t: needs a translation key, such as {t:block.minecraft.stone}")
			return
		}
		o := &jsonv.Object{}
		o.Set("translate", jsonv.String(val))
		last := val[strings.LastIndexByte(val, '.')+1:]
		p.addObj(src, open, o, textAdvance(last, p.style().bold), click, hover)
	default:
		if strings.TrimSpace(body) == "" {
			p.errorf(src, open, close+1, "{} is empty; write a style such as {gold}, or \\{ for a brace")
			return
		}
		st := p.style()
		for _, w := range strings.Fields(body) {
			if !p.styleWord(&st, w) {
				at := bodyAt + strings.Index(body, w)
				// The group still opens, so its {/} is not reported too.
				p.errorf(src, at, at+len([]rune(w)), "'%s' is not a style; use a color such as gold or #ff8800, bold, italic, underlined, strikethrough, obfuscated, shadow:#rrggbb or font:<font>%s", w, didYouMean(w, p.c.styleNames()))
			}
		}
		// Styles inside a {...} group are the group's own, not the marks'.
		p.groups = append(p.groups, bookGroup{style: st, at: bookMark{src: src, off: open}})
	}
}

func (p *bookParser) styleWord(st *bookStyle, w string) bool {
	if s, ok := bookStyleWords[w]; ok {
		switch s {
		case "bold":
			st.bold = true
		case "italic":
			st.italic = true
		case "underlined":
			st.underlined = true
		case "strikethrough":
			st.strikethrough = true
		case "obfuscated":
			st.obfuscated = true
		}
		return true
	}
	if isHexColor(w, 6) {
		st.color = strings.ToUpper(w)
		return true
	}
	if p.c.isTextColor(w) {
		st.color = w
		return true
	}
	if v, ok := strings.CutPrefix(w, "shadow:"); ok {
		switch {
		case v == "none":
			z := int64(0)
			st.shadow = &z
		case isHexColor(v, 6) || isHexColor(v, 8):
			n, _ := strconv.ParseUint(v[1:], 16, 32)
			if len(v) == 7 {
				n |= 0xFF000000 // opaque
			} else {
				n = n>>8 | (n&0xFF)<<24 // #rrggbbaa to ARGB
			}
			s := int64(int32(uint32(n)))
			st.shadow = &s
		default:
			return false
		}
		return true
	}
	if v, ok := strings.CutPrefix(w, "font:"); ok && v != "" {
		if !strings.Contains(v, ":") {
			v = "minecraft:" + v
		}
		st.font = v
		return true
	}
	return false
}

func isHexColor(s string, digits int) bool {
	if len(s) != digits+1 || s[0] != '#' {
		return false
	}
	for _, r := range s[1:] {
		if !unicode.Is(unicode.ASCII_Hex_Digit, r) {
			return false
		}
	}
	return true
}

// expr handles ${...} over the runes [start, end): a score, a Global or a
// constant, which the game fills in when the book is first opened.
func (p *bookParser) expr(src *bookSrc, start, end int, click *jsonv.Object, hover string) {
	text := string(src.text[start:end])
	c := p.c
	const head = "X _ = "
	file := syntax.Parse(head + text)
	report := func(sp syntax.Span, code, msg string) {
		s, e := sp.Start-len(head)+start, sp.End-len(head)+start
		if s < start || e > end || e <= s {
			s, e = start, end
		}
		p.report(src, code, src.at(s, e), src.where(s)+msg)
	}
	if strings.TrimSpace(text) == "" || len(file.Diags) > 0 || len(file.Decls) != 1 || file.Decls[0].Value == nil {
		report(syntax.Span{}, "mms/syntax", "expected a score, a Global or a constant inside ${...}")
		return
	}
	sub := c.bookSub()
	sub.Src = []rune(head + text)
	f := &fc{c: sub, src: sub.Src, ns: c.NS, obj: fn.Objective(c.NS), names: map[string]int{}}
	f.push()
	var pre []fn.Op
	v := f.expr(&pre, file.Decls[0].Value)
	for _, d := range sub.Diags {
		report(d.Sp, d.Code, d.Message)
	}
	if v == nil || len(sub.Diags) > 0 {
		return
	}
	if len(pre) > 0 {
		report(syntax.Span{}, "mms/book", "a book can show a score, a Global or a constant, not a computed value; work it out in a function and keep it in a Global")
		return
	}
	var o *jsonv.Object
	width := 18 // a three digit number
	switch {
	case v.str != nil:
		p.add(src, -1, *v.str, click, hover)
		return
	case v.typ == tString:
		loc, path, _ := strings.Cut(v.path, " ")
		o = &jsonv.Object{}
		o.Set("nbt", jsonv.String(path))
		o.Set("storage", jsonv.String(loc))
		width = 48
	case v.op.IsConst:
		p.add(src, -1, strconv.Itoa(int(v.op.Const)), click, hover)
		return
	default:
		o = &jsonv.Object{}
		sc := &jsonv.Object{}
		sc.Set("name", jsonv.String(v.op.S.Holder))
		sc.Set("objective", jsonv.String(v.op.S.Obj))
		o.Set("score", sc)
	}
	p.addObj(src, start-2, o, width, click, hover)
}

// finish reports marks and styles left open, and links to pages that do
// not exist.
func (p *bookParser) finish() {
	for m, at := range p.toggles {
		p.errorf(at.src, at.off, at.off+len(m), "this %s is not closed; write \\%s for the character itself", m, m)
	}
	for _, g := range p.groups {
		e := indexRune(g.at.src.text, '}', g.at.off, len(g.at.src.text)) + 1
		p.errorf(g.at.src, g.at.off, e, "this style is never closed with {/}")
	}
	for _, l := range p.links {
		if l.page > len(p.pages) {
			p.errorf(l.at.src, l.at.off, l.end, "page %d does not exist: the book has %d", l.page, len(p.pages))
		}
	}
}

// json is a page as a text component.
func (pg *bookPage) json() jsonv.Value {
	var parts jsonv.Array
	for _, r := range pg.runs {
		if r.obj == nil && r.text == "" {
			continue
		}
		var o *jsonv.Object
		if r.obj != nil {
			o = &jsonv.Object{Members: append([]jsonv.Member(nil), r.obj.Members...)}
		} else {
			o = &jsonv.Object{}
			o.Set("text", jsonv.String(r.text))
		}
		s := r.style
		if s.color != "" {
			o.Set("color", jsonv.String(s.color))
		}
		if s.font != "" {
			o.Set("font", jsonv.String(s.font))
		}
		if s.shadow != nil {
			o.Set("shadow_color", jsonv.Number(strconv.FormatInt(*s.shadow, 10)))
		}
		for _, b := range []struct {
			key string
			on  bool
		}{{"bold", s.bold}, {"italic", s.italic}, {"underlined", s.underlined}, {"strikethrough", s.strikethrough}, {"obfuscated", s.obfuscated}} {
			if b.on {
				o.Set(b.key, jsonv.Bool(true))
			}
		}
		if r.click != nil && len(r.click.Members) > 0 {
			o.Set("click_event", r.click)
		}
		if r.hover != "" {
			h := &jsonv.Object{}
			h.Set("action", jsonv.String("show_text"))
			h.Set("value", jsonv.String(r.hover))
			o.Set("hover_event", h)
		}
		if len(o.Members) == 1 && r.obj == nil {
			parts = append(parts, jsonv.String(r.text))
			continue
		}
		parts = append(parts, o)
	}
	switch len(parts) {
	case 0:
		return jsonv.String("")
	case 1:
		if s, ok := parts[0].(jsonv.String); ok {
			return s
		}
	}
	// The first element of a list is the parent of the rest; an empty
	// string keeps it from passing its style on.
	return append(jsonv.Array{jsonv.String("")}, parts...)
}

// glyphs lists a page's layout units in order.
func (pg *bookPage) glyphs() []bookGlyph {
	var out []bookGlyph
	for _, r := range pg.runs {
		out = append(out, r.glyph...)
	}
	return out
}
