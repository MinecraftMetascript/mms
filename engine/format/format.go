package format

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

const (
	syntaxWS          = syntax.TokenKind(parser.MMSLexerWS)
	syntaxLineComment = syntax.TokenKind(parser.MMSLexerLINE_COMMENT)
	syntaxComma       = syntax.TokenKind(parser.MMSLexerCOMMA)
)

// Width is the line width the formatter fills to (docs/plan.md).
const Width = 100

// ErrSyntax is returned for files with syntax errors; they are left alone.
var ErrSyntax = errors.New("format: file has syntax errors")

// ErrUnsafe is returned when formatting would change the program or drop a
// comment. It means the formatter has a bug; the input is left alone.
var ErrUnsafe = errors.New("format: output would not match input")

// Source formats MMS source text.
func Source(src string) (string, error) {
	f := syntax.Parse(src)
	for _, d := range f.Diags {
		if d.Severity == syntax.SeverityError {
			return "", ErrSyntax
		}
	}
	out := File(f)
	if err := verify(f, out); err != nil {
		return "", err
	}
	return out, nil
}

// File formats a parsed file. Callers should use Source, which also checks
// the result; File assumes the file has no syntax errors.
func File(f *syntax.File) string {
	p := newPrinter(f)
	// Print every item first, in order, so comments are taken in source
	// order; then join them, with a blank line around any item that
	// prints on more than one line.
	type printed struct {
		lead       leadingComments
		body, tail doc
		multi      bool
	}
	items := make([]printed, len(f.Items))
	for i, it := range f.Items {
		lead := p.leading(it.Span().Start, i == 0)
		body := p.stmt(it)
		items[i] = printed{lead: lead, body: body, tail: p.trailing(it.Span().End),
			multi: strings.Contains(render(body, Width), "\n")}
	}
	var parts []doc
	for i, it := range items {
		if i > 0 {
			parts = append(parts, hard)
			if p.blankBefore(it.lead, f.Items[i].Span().Start) || items[i-1].multi || it.multi {
				parts = append(parts, hard)
			}
		}
		parts = append(parts, it.lead.doc, it.body, it.tail)
	}
	rest := p.leading(len(f.Src)+1, len(f.Items) == 0)
	if len(f.Items) > 0 && len(rest.comments) > 0 {
		parts = append(parts, hard)
		if rest.blank {
			parts = append(parts, hard)
		}
	}
	parts = append(parts, rest.docTrimmed())
	out := strings.TrimRight(render(concat(parts), Width), " \n")
	if out == "" {
		return ""
	}
	return out + "\n"
}

// Expr prints one expression on a single line, as in a `${...}` piece
// that the evaluator writes back into a command.
func Expr(x syntax.Expr) string {
	p := newPrinter(&syntax.File{})
	return render(p.expr(x), 1<<30)
}

// Stmts prints statements one after another, as the code a Commands
// Template writes at its call.
func Stmts(items []syntax.Stmt) string {
	p := newPrinter(&syntax.File{})
	var out []string
	for _, it := range items {
		out = append(out, strings.TrimRight(render(p.stmt(it), Width), " \n"))
	}
	return strings.Join(out, "\n")
}

// srcRange returns src[a:b], clipped to src. Nodes the evaluator builds
// can carry spans of another file, or of none.
func srcRange(src []rune, a, b int) []rune {
	a, b = max(a, 0), min(b, len(src))
	if a >= b {
		return nil
	}
	return src[a:b]
}

// oneLine reports whether a span of the source holds no newline.
func oneLine(f *syntax.File, sp syntax.Span) bool {
	for _, r := range srcRange(f.Src, sp.Start, sp.End) {
		if r == '\n' {
			return false
		}
	}
	return true
}

// flatBody reports whether a declaration body may stay on one line: it
// was written on one line and holds only fields, as in
// `LootTable chests/gems { pools = [gem_pool] }`.
func flatBody(f *syntax.File, b *syntax.Block) bool {
	return len(b.Items) > 0 && oneLine(f, b.Span()) && !hasStatements(b)
}

type printer struct {
	f        *syntax.File
	comments []syntax.Token // pending comments, in source order
	lines    []int          // rune offset of each line start
}

func newPrinter(f *syntax.File) *printer {
	p := &printer{f: f, lines: []int{0}}
	for _, t := range f.Tokens {
		if t.IsComment() {
			p.comments = append(p.comments, t)
		}
	}
	for i, r := range f.Src {
		if r == '\n' {
			p.lines = append(p.lines, i+1)
		}
	}
	return p
}

func (p *printer) lineOf(off int) int {
	return sort.Search(len(p.lines), func(i int) bool { return p.lines[i] > off }) - 1
}

// leadingComments are comments printed on their own lines before an item.
type leadingComments struct {
	comments []syntax.Token
	blank    bool // a blank line followed the last comment, or preceded the item
	doc      doc
}

func (l leadingComments) docTrimmed() doc {
	if len(l.comments) == 0 {
		return nil
	}
	return l.doc
}

// leading takes the pending comments that start before off and renders
// them one per line, keeping single blank lines between them.
func (p *printer) leading(off int, first bool) leadingComments {
	var lc leadingComments
	var parts []doc
	for len(p.comments) > 0 && p.comments[0].Sp.Start < off {
		c := p.comments[0]
		p.comments = p.comments[1:]
		if len(lc.comments) > 0 && p.blankBetween(lc.comments[len(lc.comments)-1].Sp.End, c.Sp.Start) {
			parts = append(parts, hard)
		}
		lc.comments = append(lc.comments, c)
		parts = append(parts, p.comment(c), hard)
	}
	if len(lc.comments) > 0 && off <= len(p.f.Src) && p.blankBetween(lc.comments[len(lc.comments)-1].Sp.End, off) {
		parts = append(parts, hard)
	}
	if len(lc.comments) > 0 {
		lc.blank = !first && p.blankBetween(p.prevContentEnd(lc.comments[0].Sp.Start), lc.comments[0].Sp.Start)
	}
	lc.doc = concat(parts)
	return lc
}

// blankBefore reports whether the author left a blank line before an item
// (or before its leading comments).
func (p *printer) blankBefore(lc leadingComments, start int) bool {
	if len(lc.comments) > 0 {
		return lc.blank
	}
	return p.blankBetween(p.prevContentEnd(start), start)
}

func (p *printer) prevContentEnd(off int) int {
	for i := p.tokenFrom(off) - 1; i >= 0; i-- {
		if t := p.f.Tokens[i]; t.Kind != syntaxWS && !t.IsNewline() {
			return t.Sp.End
		}
	}
	return 0
}

// tokenFrom returns the index of the first token starting at or after off.
func (p *printer) tokenFrom(off int) int {
	toks := p.f.Tokens
	return sort.Search(len(toks), func(i int) bool { return toks[i].Sp.Start >= off })
}

// blankBetween reports whether a blank line (two newlines with only
// whitespace between them) lies between offsets a and b.
func (p *printer) blankBetween(a, b int) bool {
	newlines := 0
	for _, r := range srcRange(p.f.Src, a, b) {
		switch r {
		case '\n':
			newlines++
			if newlines == 2 {
				return true
			}
		case ' ', '\t', '\r':
		default:
			newlines = 0
		}
	}
	return false
}

// dangling takes the pending comments before off, the closing bracket of
// an enclosing node, and prints them one per line with no newline after the
// last. It returns nil when there are none.
func (p *printer) dangling(off int) doc {
	var parts []doc
	var last syntax.Token
	for len(p.comments) > 0 && p.comments[0].Sp.Start < off {
		c := p.comments[0]
		p.comments = p.comments[1:]
		if len(parts) > 0 {
			parts = append(parts, hard)
			if p.blankBetween(last.Sp.End, c.Sp.Start) {
				parts = append(parts, hard)
			}
		}
		parts = append(parts, p.comment(c))
		if c.Kind == syntaxLineComment {
			parts = append(parts, forceBreak{})
		}
		last = c
	}
	if len(parts) == 0 {
		return nil
	}
	return concat(parts)
}

// trailing takes a comment that sits on the same line right after end.
func (p *printer) trailing(end int) doc {
	if len(p.comments) == 0 {
		return nil
	}
	c := p.comments[0]
	if c.Sp.Start < end || p.lineOf(c.Sp.Start) != p.lineOf(max(end-1, 0)) {
		return nil
	}
	// The comment trails this node only if nothing but whitespace and one
	// separating comma sits in between; otherwise it belongs to an
	// enclosing node that ends later.
	commas := 0
	for _, t := range p.f.Tokens[p.tokenFrom(end):] {
		if t.Sp.Start >= c.Sp.Start {
			break
		}
		if t.Hidden || t.Synthetic {
			continue
		}
		if t.Kind != syntaxComma || commas > 0 {
			return nil
		}
		commas++
	}
	p.comments = p.comments[1:]
	d := cat(txt(" "), p.comment(c))
	if c.Kind == syntaxLineComment {
		// A line comment ends the line, so whatever holds it must break.
		return cat(d, forceBreak{})
	}
	return d
}

func (p *printer) comment(c syntax.Token) doc {
	text := strings.TrimRight(c.Text, " \t\r")
	if !strings.Contains(text, "\n") {
		return txt(text)
	}
	// Re-indent the continuation lines of a block comment: ` *` lines line
	// up under the opener, other lines keep their indent relative to it.
	col := c.Sp.Start - (p.lineStart(c.Sp.Start))
	lines := strings.Split(text, "\n")
	parts := []doc{txt(lines[0])}
	for _, l := range lines[1:] {
		l = strings.TrimRight(l, " \t\r")
		trimmed := strings.TrimLeft(l, " \t")
		switch {
		case strings.HasPrefix(trimmed, "*"):
			trimmed = " " + trimmed
		case len(l)-len(trimmed) > col:
			trimmed = l[col:]
		}
		parts = append(parts, hard, txt(trimmed))
	}
	return concat(parts)
}

// lineStart returns the offset of the first rune on off's line.
func (p *printer) lineStart(off int) int {
	for off > 0 && p.f.Src[off-1] != '\n' {
		off--
	}
	return off
}

func (p *printer) decl(d *syntax.Decl) doc {
	head := d.Kind.Name
	switch {
	case d.Extern:
		head = "extern " + head
	case d.Primary:
		head = "primary " + head
	}
	parts := []doc{txt(head)}
	for i, n := range d.Names() {
		if i > 0 {
			parts = append(parts, txt(","))
		}
		parts = append(parts, txt(" "+n.Text))
	}
	if d.Params != nil {
		parts = append(parts, p.params(d.Params))
	}
	if d.Type != nil {
		parts = append(parts, txt(": "), p.typ(d.Type))
	}
	if d.TypeValue != nil {
		parts = append(parts, txt(" = "), p.typ(d.TypeValue))
	}
	if d.Value != nil {
		parts = append(parts, p.assigned(d.Value))
	}
	if d.Body != nil {
		parts = append(parts, txt(" "), p.block(d.Body, !flatBody(p.f, d.Body)))
	}
	return concat(parts)
}

// assigned prints ` = value`. A value that cannot break, such as a long
// string, moves to its own indented line when it does not fit after the
// '='.
func (p *printer) assigned(v syntax.Expr) doc {
	switch v.(type) {
	case *syntax.String, *syntax.Number, *syntax.Ident, *syntax.ResLoc, *syntax.TemplatedName:
		return grp(txt(" ="), ind(ln, p.expr(v)))
	}
	return cat(txt(" = "), p.expr(v))
}

func (p *printer) params(pl *syntax.ParamList) doc {
	var items []item
	for _, prm := range pl.Params {
		items = append(items, p.item(prm, func() doc { return p.param(prm) }))
	}
	return p.bracketed("(", ")", items, pl.Span().End, false)
}

func (p *printer) param(prm *syntax.Param) doc {
	s := ""
	if prm.Name != nil {
		s = prm.Name.Text
	}
	if prm.Type != nil {
		s += ": " + syntax.TypeText(prm.Type)
	}
	if prm.Default != nil {
		return cat(txt(s+" = "), p.expr(prm.Default))
	}
	return txt(s)
}

// typed prints `name: Type = value`, the shape of const, let and field
// declarations.
func (p *printer) typed(word string, name *syntax.Name, t syntax.TypeExpr, v syntax.Expr) doc {
	s := word
	if name != nil {
		s += name.Text
	}
	if t == nil {
		if v == nil {
			return txt(s)
		}
		return cat(txt(s), p.assigned(v))
	}
	d := cat(txt(s+": "), p.typ(t))
	if v == nil {
		return d
	}
	return cat(d, p.assigned(v))
}

// typ prints a type. A record written in place, `{ x: int }`, prints as
// a block of field declarations, keeping their doc comments and defaults.
func (p *printer) typ(t syntax.TypeExpr) doc {
	if !hasRecord(t) {
		return txt(syntax.TypeText(t))
	}
	switch t := t.(type) {
	case *syntax.TypeRecord:
		items := make([]syntax.Stmt, len(t.Fields))
		for i, f := range t.Fields {
			items[i] = f
		}
		// One line when the author wrote it so, like a value block.
		return p.blockDoc(&syntax.Block{Items: items, Sp: t.Sp}, !oneLine(p.f, t.Sp))
	case *syntax.TypeOptional:
		return cat(p.typ(t.X), txt("?"))
	case *syntax.TypeList:
		return cat(txt("["), p.typ(t.Elem), txt("]"))
	case *syntax.TypeParen:
		return cat(txt("("), p.typ(t.X), txt(")"))
	case *syntax.TypeMap:
		return cat(txt("{"+syntax.TypeText(t.Key)+": "), p.typ(t.Value), txt("}"))
	case *syntax.TypeUnion:
		var parts []doc
		for i, a := range t.Arms {
			if i > 0 {
				parts = append(parts, txt(" | "))
			}
			parts = append(parts, p.typ(a))
		}
		return cat(parts...)
	}
	return txt(syntax.TypeText(t))
}

// hasRecord reports whether a record written in place is anywhere in t.
func hasRecord(t syntax.TypeExpr) bool {
	found := false
	syntax.Walk(t, func(n syntax.Node) bool {
		if _, ok := n.(*syntax.TypeRecord); ok {
			found = true
		}
		return !found
	})
	return found
}

// block prints `{ ... }`. Statement blocks and declaration bodies always
// break; value blocks stay on one line when they fit.
func (p *printer) block(b *syntax.Block, forceBreak bool) doc {
	if !forceBreak {
		// Statements always go one per line. A value block with several
		// items stays on one line only when the author wrote it so.
		forceBreak = hasStatements(b) || len(b.Items) > 1 && !oneLine(p.f, b.Span())
	}
	return p.blockDoc(b, forceBreak)
}

// blockDoc prints `{ ... }`, on one line when it fits unless forceBreak.
func (p *printer) blockDoc(b *syntax.Block, forceBreak bool) doc {
	if len(b.Items) == 0 {
		inner := p.dangling(b.Span().End)
		if inner == nil {
			return txt("{}")
		}
		return cat(txt("{"), ind(hard, inner), hard, txt("}"))
	}
	var items []doc
	for i, it := range b.Items {
		lead := p.leading(it.Span().Start, i == 0)
		var sep doc
		if i > 0 {
			if p.blankBefore(lead, it.Span().Start) {
				sep = cat(whenBroken(hard), lnOrComma())
			} else {
				sep = lnOrComma()
			}
		}
		items = append(items, cat(sep, lead.doc, p.stmt(it), p.trailing(it.Span().End)))
	}
	if tail := p.dangling(b.Span().End); tail != nil {
		items = append(items, hard, tail)
	}
	if forceBreak {
		return cat(txt("{"), ind(append([]doc{hard}, items...)...), hard, txt("}"))
	}
	return grp(txt("{"), ind(append([]doc{ln}, items...)...), ln, txt("}"))
}

// lnOrComma separates block items: ", " when flat, a newline when broken.
func lnOrComma() doc { return ifBreak{broken: ln, flat: txt(", ")} }

func hasStatements(b *syntax.Block) bool {
	for _, it := range b.Items {
		switch it := it.(type) {
		case *syntax.FieldStmt:
		case *syntax.ExprStmt:
			if _, ok := it.X.(*syntax.RunExpr); ok {
				return true
			}
		default:
			return true
		}
	}
	return false
}

func (p *printer) stmt(s syntax.Stmt) doc {
	switch s := s.(type) {
	case *syntax.Decl:
		return p.decl(s)
	case *syntax.LetStmt:
		return p.typed("let ", s.Name, s.Type, s.Value)
	case *syntax.ConstStmt:
		return p.typed("const ", s.Name, s.Type, s.Value)
	case *syntax.FieldDecl:
		return p.typed("", s.Name, s.Type, s.Default)
	case *syntax.SectionStmt:
		return cat(txt(s.Word.Name+" "), p.block(s.Body, true))
	case *syntax.AsStmt:
		d := cat(txt("as "+syntax.TypeText(s.Type)), p.assigned(s.Value))
		if s.Hint != nil {
			d = cat(d, txt(" hint"), p.args(s.Hint.Args, s.Hint.ArgsSp.End))
		}
		return d
	case *syntax.ExpectStmt:
		return cat(txt("expect "), p.expr(s.X))
	case *syntax.RunStmt:
		return cat(txt("run "), p.expr(s.X))
	case *syntax.IfStmt:
		return p.ifStmt(s)
	case *syntax.WhileStmt:
		return cat(txt("while ("), p.expr(s.Cond), txt(") "), p.block(s.Body, true))
	case *syntax.ForStmt:
		return cat(txt("for "+forVars(s.Var, s.Var2)+" in "), p.expr(s.Iter), txt(" "), p.block(s.Body, true))
	case *syntax.ReturnStmt:
		if s.Value == nil {
			return txt("return")
		}
		return cat(txt("return "), p.expr(s.Value))
	case *syntax.BreakStmt:
		return txt("break")
	case *syntax.ContinueStmt:
		return txt("continue")
	case *syntax.CommandStmt:
		// Command text is kept verbatim: Brigadier treats extra spaces as
		// part of an argument (say) or as an error, never as layout.
		return txt("/" + strings.TrimRight(s.Text, " \t"))
	case *syntax.FieldStmt:
		return cat(txt(fieldKey(s)+" = "), p.expr(s.Value))
	case *syntax.AssignStmt:
		return cat(p.expr(s.Target), txt(" "+s.Op+" "), p.expr(s.Value))
	case *syntax.ExprStmt:
		return p.expr(s.X)
	case *syntax.Block:
		return p.block(s, true)
	}
	return nil
}

func (p *printer) ifStmt(s *syntax.IfStmt) doc {
	parts := []doc{txt("if ("), p.expr(s.Cond), txt(") "), p.branch(s.Then)}
	switch e := s.Else.(type) {
	case *syntax.IfStmt:
		parts = append(parts, txt(" else "), p.ifStmt(e))
	case *syntax.Block:
		parts = append(parts, txt(" else "), p.branch(e))
	}
	return concat(parts)
}

// branch prints an if statement's branch: a block, or the single call of
// a bare one, as in `if (x) error(...)`.
func (p *printer) branch(b *syntax.Block) doc {
	if b.Bare && len(b.Items) == 1 {
		return p.stmt(b.Items[0])
	}
	return p.block(b, true)
}

func forVars(a, b *syntax.Name) string {
	s := "?"
	if a != nil {
		s = a.Text
	}
	if b != nil {
		s += ", " + b.Text
	}
	return s
}

var bareKey = regexp.MustCompile(`^([a-zA-Z_][a-zA-Z0-9_]*|[a-z0-9_.-]+:[a-z0-9_./-]+|[a-zA-Z_][a-zA-Z0-9_]*(/[a-z0-9_./-]+)+)$`)

func fieldKey(s *syntax.FieldStmt) string {
	if s.KeyQuoted || !bareKey.MatchString(s.Key.Text) {
		return strconv.Quote(s.Key.Text)
	}
	return s.Key.Text
}

// precedence of binary operators, higher binds tighter.
var precedence = map[string]int{
	"??": 2, "||": 3, "&&": 4, "==": 5, "!=": 5,
	"<": 6, "<=": 6, ">": 6, ">=": 6, "..": 7,
	"+": 8, "-": 8, "*": 9, "/": 9, "%": 9,
}

// Precedence of the other forms, on the same scale.
const (
	precRange = 7
	precUnary = 10
	precChain = 11
)

func (p *printer) expr(e syntax.Expr) doc {
	switch e := e.(type) {
	case *syntax.CommandStmt:
		// A command as a value runs to the end of its line, as on a
		// line of its own.
		return txt("/" + strings.TrimRight(e.Text, " \t"))
	case *syntax.Number:
		return txt(e.Text)
	case *syntax.String:
		if strings.HasPrefix(e.Raw, `"`) {
			return txt(e.Raw)
		}
		return txt(strconv.Quote(e.Value))
	case *syntax.Bool:
		return txt(strconv.FormatBool(e.Value))
	case *syntax.Ident:
		return txt(e.Name)
	case *syntax.ResLoc:
		return txt(e.Text)
	case *syntax.Tag:
		return txt("#" + e.Name)
	case *syntax.Selector:
		return txt("@" + e.Name)
	case *syntax.Coord:
		return txt(strings.Join(strings.Fields(e.Text), " "))
	case *syntax.Paren:
		return cat(txt("("), p.expr(e.X), txt(")"))
	case *syntax.List:
		var elems []item
		for _, x := range e.Elems {
			elems = append(elems, p.item(x, func() doc { return p.expr(x) }))
		}
		return p.bracketed("[", "]", elems, e.Span().End, p.listBreaks(e))
	case *syntax.Object:
		return p.block(e.Body, false)
	case *syntax.Unary:
		return cat(txt(e.Op), p.operand(e.X, precUnary))
	case *syntax.TemplatedName:
		return txt(e.Text)
	case *syntax.Spread:
		return cat(txt("..."), p.expr(e.X))
	case *syntax.RunExpr:
		return cat(txt("run "), p.block(e.Body, true))
	case *syntax.IfExpr:
		d := cat(txt("if ("), p.expr(e.Cond), txt(")"), ind(ln, p.expr(e.Then)))
		if e.Else != nil {
			d = cat(d, ln, txt("else"), ind(ln, p.expr(e.Else)))
		}
		return grp(d)
	case *syntax.Comprehension:
		head := []doc{txt("for " + forVars(e.Var, e.Var2) + " in "), p.expr(e.Iter)}
		if e.Cond != nil {
			head = append(head, txt(" if "), p.expr(e.Cond))
		}
		head = append(head, txt(":"))
		return grp(txt("["), ind(soft, concat(head), ind(ln, p.expr(e.Body))), soft, txt("]"))
	case *syntax.Binary:
		return p.binary(e)
	case *syntax.Range:
		var lo, hi doc
		if e.Lo != nil {
			lo = p.operand(e.Lo, precRange+1)
		}
		if e.Hi != nil {
			hi = p.operand(e.Hi, precRange+1)
		}
		return cat(lo, txt(".."), hi)
	case *syntax.Member, *syntax.Call, *syntax.Index, *syntax.Trailing:
		return p.chain(e, 0)
	}
	return nil
}

// item is an element of a bracketed list with the comments around it.
type item struct{ lead, body, trail doc }

// item renders n's leading comments, then n, then its trailing comment, in
// that order so comments are consumed in source order.
func (p *printer) item(n syntax.Node, body func() doc) item {
	var it item
	it.lead = p.leading(n.Span().Start, true).doc
	it.body = body()
	it.trail = p.trailing(n.Span().End)
	return it
}

// listBreaks: a list of objects or of calls with blocks (such as surface
// rules) reads better one per line, so it always breaks.
func (p *printer) listBreaks(l *syntax.List) bool {
	if len(l.Elems) < 2 {
		return false
	}
	for _, x := range l.Elems {
		switch x.(type) {
		case *syntax.Object, *syntax.Trailing:
			return true
		}
	}
	return false
}

// bracketed prints items between open and close, separated by commas, with
// a trailing comma when broken. end is the offset just past the closer;
// comments before it that no item took are printed inside, at the end.
func (p *printer) bracketed(open, close string, items []item, end int, force bool) doc {
	tail := p.dangling(end)
	if len(items) == 0 && tail == nil {
		return txt(open + close)
	}
	sep, comma := doc(ln), whenBroken(txt(","))
	if force {
		sep, comma = hard, txt(",")
	}
	var inner []doc
	for i, it := range items {
		if i > 0 {
			inner = append(inner, sep)
		}
		// The comma goes before a trailing comment, never after it.
		c := doc(txt(","))
		if i == len(items)-1 {
			c = comma
		}
		inner = append(inner, it.lead, it.body, c, it.trail)
	}
	if tail != nil {
		if len(items) > 0 {
			inner = append(inner, hard)
		}
		inner = append(inner, tail)
	}
	if force {
		return cat(txt(open), ind(append([]doc{hard}, inner...)...), hard, txt(close))
	}
	return grp(txt(open), ind(append([]doc{soft}, inner...)...), soft, txt(close))
}

// operand prints a subexpression, adding parentheses only where the tree
// needs them (the AST keeps the author's own parentheses as Paren nodes).
func (p *printer) operand(x syntax.Expr, prec int) doc {
	if b, ok := x.(*syntax.Binary); ok && precedence[b.Op] < prec {
		return cat(txt("("), p.expr(x), txt(")"))
	}
	if r, ok := x.(*syntax.Range); ok && precRange < prec {
		return cat(txt("("), p.expr(r), txt(")"))
	}
	if s, ok := x.(*syntax.Spread); ok && prec > 0 {
		return cat(txt("("), p.expr(s), txt(")"))
	}
	return p.expr(x)
}

func (p *printer) binary(b *syntax.Binary) doc {
	prec := precedence[b.Op]
	switch b.Op {
	case "&&", "||":
		// Flatten a chain of the same operator; break before each operator,
		// all or nothing.
		var ops []syntax.Expr
		var walk func(e syntax.Expr)
		walk = func(e syntax.Expr) {
			if x, ok := e.(*syntax.Binary); ok && x.Op == b.Op {
				walk(x.L)
				walk(x.R)
				return
			}
			ops = append(ops, e)
		}
		walk(b)
		parts := []doc{p.operand(ops[0], prec+1)}
		for _, o := range ops[1:] {
			parts = append(parts, ln, txt(b.Op+" "), p.operand(o, prec+1))
		}
		return grp(parts[0], ind(parts[1:]...))
	}
	// An object or list on the right hugs the operator and breaks inside
	// itself: `components + {` on one line.
	switch b.R.(type) {
	case *syntax.Object, *syntax.List:
		return cat(p.operand(b.L, prec), txt(" "+b.Op+" "), p.operand(b.R, prec+1))
	}
	// Left-associative arithmetic and comparisons: break after the operator.
	return grp(p.operand(b.L, prec), txt(" "+b.Op), ind(ln, p.operand(b.R, prec+1)))
}

// chain prints calls, member access, indexes and trailing blocks. A chain
// with two or more `.method(...)` steps breaks before every dot when it
// does not fit.
//
// maxFlat, when set, caps the flat width of the outermost call's argument
// list once the enclosing list has broken (see nestedArgWidth).
func (p *printer) chain(e syntax.Expr, maxFlat int) doc {
	type step struct {
		member bool
		d      doc
	}
	// Collect the steps outside in, then render them inside out so that
	// comments are consumed in source order.
	var nodes []syntax.Expr
	var base syntax.Expr = e
loop:
	for {
		nodes = append(nodes, base)
		switch x := base.(type) {
		case *syntax.Call:
			base = x.Fn
		case *syntax.Index:
			base = x.X
		case *syntax.Member:
			base = x.X
		case *syntax.Trailing:
			base = x.X
		default:
			nodes = nodes[:len(nodes)-1]
			break loop
		}
	}
	head := []doc{p.operand(base, precChain)}
	var steps []step
	members := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		switch x := nodes[i].(type) {
		case *syntax.Call:
			d := p.args(x.Args, x.ArgsSp.End)
			if g, ok := d.(group); ok && i == 0 && maxFlat > 0 {
				g.maxFlat = maxFlat
				d = g
			}
			steps = append(steps, step{d: d})
		case *syntax.Index:
			var items []item
			for _, a := range x.Args {
				items = append(items, p.item(a, func() doc { return p.indexArg(a) }))
			}
			steps = append(steps, step{d: p.bracketed("[", "]", items, x.Span().End, false)})
		case *syntax.Member:
			members++
			steps = append(steps, step{member: true, d: txt("." + x.Name.Text)})
		case *syntax.Trailing:
			steps = append(steps, step{d: cat(txt(" "), p.block(x.Block, false))})
		}
	}
	if members < 2 {
		for _, s := range steps {
			head = append(head, s.d)
		}
		return concat(head)
	}
	// Keep everything up to the first member access on the head line.
	i := 0
	for ; i < len(steps) && !steps[i].member; i++ {
		head = append(head, steps[i].d)
	}
	var tail []doc
	for ; i < len(steps); i++ {
		if steps[i].member {
			tail = append(tail, soft)
		}
		tail = append(tail, steps[i].d)
	}
	return grp(concat(head), ind(tail...))
}

func (p *printer) args(args []*syntax.Arg, end int) doc {
	var items []item
	for _, a := range args {
		items = append(items, p.item(a, func() doc { return p.arg(a) }))
	}
	return p.bracketed("(", ")", items, end, false)
}

// indexArg prints a block state or selector argument, which reads like the
// command syntax it lowers to: `[type=zombie, limit=1]`.
func (p *printer) indexArg(a *syntax.Arg) doc {
	if a.Name != nil {
		return cat(txt(a.Name.Text+"="), p.expr(a.Value))
	}
	return p.expr(a.Value)
}

func (p *printer) arg(a *syntax.Arg) doc {
	v := p.argValue(a.Value)
	if a.Name != nil {
		return cat(txt(a.Name.Text+" = "), v)
	}
	return v
}

// nestedArgWidth is the widest a nested call may print on one line as an
// argument of a call that already broke one argument per line. A nested
// call is one with two or more arguments, at least one of which holds
// another call with two or more arguments. Past the width it breaks too,
// so chains such as RangeChoice(..., RangeChoice(..., RangeChoice(...)))
// step down one level per line instead of leaving a long tail of calls on
// one line. Short ones like Add(Mul(a, b), c) and wrappers with one
// argument like Not(...) keep the plain fits-or-breaks rule.
const nestedArgWidth = 40

// argValue prints an argument value, applying nestedArgWidth to a call
// with calls in its arguments.
func (p *printer) argValue(e syntax.Expr) doc {
	if c, ok := e.(*syntax.Call); ok && len(c.Args) >= 2 && argsHaveCall(c.Args) {
		return p.chain(c, nestedArgWidth)
	}
	return p.expr(e)
}

// argsHaveCall reports whether any argument contains a call with two or
// more arguments.
func argsHaveCall(args []*syntax.Arg) bool {
	for _, a := range args {
		if hasCall(a.Value) {
			return true
		}
	}
	return false
}

func hasCall(e syntax.Expr) bool {
	switch e := e.(type) {
	case *syntax.Call:
		return len(e.Args) >= 2 || hasCall(e.Fn) || argsHaveCall(e.Args)
	case *syntax.Paren:
		return hasCall(e.X)
	case *syntax.Unary:
		return hasCall(e.X)
	case *syntax.Binary:
		return hasCall(e.L) || hasCall(e.R)
	case *syntax.Range:
		return e.Lo != nil && hasCall(e.Lo) || e.Hi != nil && hasCall(e.Hi)
	case *syntax.Member:
		return hasCall(e.X)
	case *syntax.Index:
		return hasCall(e.X)
	case *syntax.List:
		for _, x := range e.Elems {
			if hasCall(x) {
				return true
			}
		}
	}
	return false
}
