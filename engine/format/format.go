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
	var parts []doc
	for i, d := range f.Decls {
		lead := p.leading(d.Span().Start, i == 0)
		if i > 0 {
			parts = append(parts, hard)
			if p.blankBefore(lead, d.Span().Start) || multiLine(f, f.Decls[i-1]) || multiLine(f, d) {
				parts = append(parts, hard)
			}
		}
		parts = append(parts, lead.doc, p.decl(d), p.trailing(d.Span().End))
	}
	rest := p.leading(len(f.Src)+1, len(f.Decls) == 0)
	if len(f.Decls) > 0 && len(rest.comments) > 0 {
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

func multiLine(f *syntax.File, d *syntax.Decl) bool {
	sp := d.Span()
	for _, r := range f.Src[sp.Start:min(sp.End, len(f.Src))] {
		if r == '\n' {
			return true
		}
	}
	return d.Body != nil && len(d.Body.Items) > 0
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
	for _, r := range p.f.Src[max(a, 0):min(b, len(p.f.Src))] {
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
	parts := []doc{txt(d.Kind.Name)}
	if d.Name != nil {
		parts = append(parts, txt(" "+d.Name.Text))
	}
	if d.Params != nil {
		parts = append(parts, p.params(d.Params))
	}
	if d.Type != nil {
		parts = append(parts, txt(": "+d.Type.Text))
	}
	if d.Value != nil {
		parts = append(parts, txt(" = "), p.expr(d.Value))
	}
	if d.Body != nil {
		parts = append(parts, txt(" "), p.block(d.Body, true))
	}
	return concat(parts)
}

func (p *printer) params(pl *syntax.ParamList) doc {
	var ps []doc
	for _, prm := range pl.Params {
		s := prm.Name.Text
		if prm.Type != nil {
			s += ": " + prm.Type.Text
		}
		ps = append(ps, txt(s))
	}
	if len(ps) == 0 {
		return txt("()")
	}
	return grp(txt("("), ind(soft, join(cat(txt(","), ln), ps), whenBroken(txt(","))), soft, txt(")"))
}

// block prints `{ ... }`. Statement blocks and declaration bodies always
// break; value blocks stay on one line when they fit.
func (p *printer) block(b *syntax.Block, forceBreak bool) doc {
	if !forceBreak {
		forceBreak = hasStatements(b)
	}
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
func lnOrComma() doc { return ifBreak{broken: hard, flat: txt(", ")} }

func hasStatements(b *syntax.Block) bool {
	for _, it := range b.Items {
		switch it.(type) {
		case *syntax.FieldStmt, *syntax.ExprStmt:
		default:
			return true
		}
	}
	return false
}

func (p *printer) stmt(s syntax.Stmt) doc {
	switch s := s.(type) {
	case *syntax.LetStmt:
		parts := []doc{txt("let " + s.Name.Text)}
		if s.Type != nil {
			parts = append(parts, txt(": "+s.Type.Text))
		}
		if s.Value != nil {
			parts = append(parts, txt(" = "), p.expr(s.Value))
		}
		return concat(parts)
	case *syntax.IfStmt:
		return p.ifStmt(s)
	case *syntax.WhileStmt:
		return cat(txt("while ("), p.expr(s.Cond), txt(") "), p.block(s.Body, true))
	case *syntax.ForStmt:
		return cat(txt("for "+s.Var.Text+" in "), p.expr(s.Iter), txt(" "), p.block(s.Body, true))
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
	parts := []doc{txt("if ("), p.expr(s.Cond), txt(") "), p.block(s.Then, true)}
	switch e := s.Else.(type) {
	case *syntax.IfStmt:
		parts = append(parts, txt(" else "), p.ifStmt(e))
	case *syntax.Block:
		parts = append(parts, txt(" else "), p.block(e, true))
	}
	return concat(parts)
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
	"->": 1, "||": 2, "&&": 3, "==": 4, "!=": 4,
	"<": 5, "<=": 5, ">": 5, ">=": 5, "..": 6,
	"+": 7, "-": 7, "*": 8, "/": 8, "%": 8,
}

func (p *printer) expr(e syntax.Expr) doc {
	switch e := e.(type) {
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
		return cat(txt(e.Op), p.operand(e.X, 9))
	case *syntax.Binary:
		return p.binary(e)
	case *syntax.Range:
		var lo, hi doc
		if e.Lo != nil {
			lo = p.operand(e.Lo, 7)
		}
		if e.Hi != nil {
			hi = p.operand(e.Hi, 7)
		}
		return cat(lo, txt(".."), hi)
	case *syntax.Member, *syntax.Call, *syntax.Index, *syntax.Trailing:
		return p.chain(e)
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
	if r, ok := x.(*syntax.Range); ok && 6 < prec {
		return cat(txt("("), p.expr(r), txt(")"))
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
	case "->":
		return cat(p.operand(b.L, prec+1), txt(" -> "), p.operand(b.R, prec))
	}
	// Left-associative arithmetic and comparisons: break after the operator.
	return grp(p.operand(b.L, prec), txt(" "+b.Op), ind(ln, p.operand(b.R, prec+1)))
}

// chain prints calls, member access, indexes and trailing blocks. A chain
// with two or more `.method(...)` steps breaks before every dot when it
// does not fit.
func (p *printer) chain(e syntax.Expr) doc {
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
	head := []doc{p.operand(base, 10)}
	var steps []step
	members := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		switch x := nodes[i].(type) {
		case *syntax.Call:
			steps = append(steps, step{d: p.args(x.Args, x.ArgsSp.End)})
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
	if a.Name != nil {
		return cat(txt(a.Name.Text+" = "), p.expr(a.Value))
	}
	return p.expr(a.Value)
}
