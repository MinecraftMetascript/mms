package syntax

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// Parse parses src. It always returns a file; problems are in File.Diags.
func Parse(src string) *File {
	f := &File{Src: []rune(src)}

	lexer := parser.NewMMSLexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	var raw []antlr.Token
	for {
		t := lexer.NextToken()
		raw = append(raw, t)
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
	toks, synthetic, closeDiags := closeUnbalanced(raw)
	stream := antlr.NewCommonTokenStream(&sliceSource{MMSLexer: lexer, toks: toks}, antlr.TokenDefaultChannel)
	stream.Fill()
	for _, t := range stream.GetAllTokens() {
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		tok := Token{
			Kind:      TokenKind(t.GetTokenType()),
			Text:      t.GetText(),
			Sp:        Span{t.GetStart(), t.GetStop() + 1},
			Hidden:    t.GetChannel() != antlr.TokenDefaultChannel,
			Synthetic: synthetic[t],
		}
		f.Tokens = append(f.Tokens, tok)
		switch tok.Kind {
		case parser.MMSLexerERROR_CHAR:
			f.Diags = append(f.Diags, Diagnostic{
				Code: "syntax/unexpected-char", Severity: SeverityError,
				Message: fmt.Sprintf("unexpected character %q", tok.Text), Sp: tok.Sp,
			})
		case parser.MMSLexerUNTERMINATED_COMMENT:
			f.Diags = append(f.Diags, Diagnostic{
				Code: "syntax/unterminated-comment", Severity: SeverityError,
				Message: "comment is never closed", Sp: Span{tok.Sp.Start, tok.Sp.Start + 2},
			})
		}
	}
	f.Diags = append(f.Diags, closeDiags...)

	// Two-stage parse: fast SLL prediction first, full LL only when SLL
	// reports a syntax error (the standard ANTLR recipe).
	tree, errs := parseWith(stream, antlr.PredictionModeSLL, synthetic)
	if len(errs) > 0 {
		stream.Seek(0)
		tree, errs = parseWith(stream, antlr.PredictionModeLL, synthetic)
	}
	f.Diags = append(f.Diags, errs...)

	b := builder{file: f}
	for _, d := range tree.AllDeclaration() {
		if decl := b.decl(d); decl != nil {
			f.Decls = append(f.Decls, decl)
		}
	}
	return f
}

func parseWith(stream *antlr.CommonTokenStream, mode int, synthetic map[antlr.Token]bool) (parser.IFileContext, []Diagnostic) {
	p := parser.NewMMSParser(stream)
	p.RemoveErrorListeners()
	el := &errorListener{synthetic: synthetic}
	p.AddErrorListener(el)
	p.GetInterpreter().SetPredictionMode(mode)
	p.SetErrorHandler(&errorStrategy{antlr.NewDefaultErrorStrategy()})
	return p.File(), el.diags
}

type errorListener struct {
	*antlr.DefaultErrorListener
	synthetic map[antlr.Token]bool
	diags     []Diagnostic
}

func (e *errorListener) SyntaxError(_ antlr.Recognizer, offending any, _, _ int, msg string, _ antlr.RecognitionException) {
	sp := Span{}
	if t, ok := offending.(antlr.Token); ok && e.synthetic[t] {
		return // already reported as an unclosed bracket
	}
	if t, ok := offending.(antlr.Token); ok && t != nil {
		sp = Span{t.GetStart(), t.GetStop() + 1}
		if t.GetTokenType() == antlr.TokenEOF {
			sp = Span{t.GetStart(), t.GetStart()}
		}
	}
	if t, ok := offending.(antlr.Token); ok && t != nil && t.GetTokenType() == parser.MMSLexerSLASH && commandFollows(t) {
		msg = "a command must start on its own line"
	} else {
		msg = cleanMessage(msg)
	}
	e.diags = append(e.diags, Diagnostic{
		Code: "syntax/error", Severity: SeverityError, Message: msg, Sp: sp,
	})
}

// commandFollows reports whether a '/' token is followed by a letter, so it
// was meant as a command but is not at the start of a line.
func commandFollows(t antlr.Token) bool {
	in := t.GetInputStream()
	next := t.GetStop() + 1
	if in == nil || next >= in.Size() {
		return false
	}
	c := in.GetText(next, next)
	return c >= "a" && c <= "z"
}

// cleanMessage turns ANTLR's token-name messages into readable ones.
func cleanMessage(msg string) string {
	r := strings.NewReplacer(
		"<EOF>", "end of file",
		"extraneous input", "unexpected",
		"mismatched input", "unexpected",
		"no viable alternative at input", "cannot parse",
	)
	return r.Replace(msg)
}

type builder struct{ file *File }

func span(ctx antlr.ParserRuleContext) Span {
	start := ctx.GetStart()
	stop := ctx.GetStop()
	if start == nil {
		return Span{}
	}
	s := start.GetStart()
	if s < 0 {
		s = 0
	}
	if stop == nil || stop.GetStop() < s || stop.GetTokenType() == antlr.TokenEOF {
		return Span{s, s}
	}
	return Span{s, stop.GetStop() + 1}
}

func tokSpan(t antlr.Token) Span {
	if t == nil || t.GetStart() < 0 {
		return Span{}
	}
	return Span{t.GetStart(), t.GetStop() + 1}
}

func nodeSpan(n antlr.TerminalNode) Span {
	if n == nil {
		return Span{}
	}
	return tokSpan(n.GetSymbol())
}

// real reports whether a terminal was actually in the input rather than
// conjured up by error recovery.
func real(n antlr.TerminalNode) bool {
	return n != nil && n.GetSymbol() != nil && n.GetSymbol().GetStart() >= 0 &&
		n.GetSymbol().GetTokenIndex() >= 0
}

func (b *builder) decl(c parser.IDeclarationContext) *Decl {
	ctx := c.(*parser.DeclarationContext)
	if ctx.GetKind() == nil {
		return nil
	}
	d := &Decl{
		Kind: &Ident{Name: ctx.GetKind().GetText(), Sp: tokSpan(ctx.GetKind())},
		Sp:   span(ctx),
	}
	if n := ctx.DeclName(); n != nil {
		d.Name = &Name{Text: n.GetText(), Sp: span(n.(antlr.ParserRuleContext))}
	}
	if t := ctx.TypeAnnotation(); t != nil {
		d.Type = b.typeAnn(t)
	}
	if p := ctx.Params(); p != nil {
		d.Params = b.params(p.(*parser.ParamsContext))
	}
	if blk := ctx.Block(); blk != nil {
		d.Body = b.block(blk)
	}
	if body := ctx.DeclBody(); body != nil {
		bc := body.(*parser.DeclBodyContext)
		switch {
		case bc.Block() != nil:
			d.Body = b.block(bc.Block())
		case bc.Expr() != nil:
			d.Value = b.expr(bc.Expr(), span(bc))
		default:
			sp := span(bc)
			d.Value = &Bad{Sp: Span{sp.End, sp.End}}
			b.file.Diags = append(b.file.Diags, Diagnostic{
				Code: "syntax/missing-value", Severity: SeverityError,
				Message: "expected a value after '='", Sp: sp,
			})
		}
	}
	d.Doc = b.docBefore(ctx.GetStart().GetTokenIndex())
	return d
}

// docBefore returns the doc comment directly above the token at index i,
// allowing only whitespace in between.
func (b *builder) docBefore(tokIndex int) string {
	// Token indexes in the stream match positions in f.Tokens because both
	// list every token in order.
	for i := tokIndex - 1; i >= 0 && i < len(b.file.Tokens); i-- {
		t := b.file.Tokens[i]
		switch {
		case t.Kind == parser.MMSLexerWS || t.Kind == parser.MMSLexerNEWLINE:
			continue
		case t.Kind == parser.MMSLexerDOC_COMMENT:
			return cleanDoc(t.Text)
		default:
			return ""
		}
	}
	return ""
}

func cleanDoc(text string) string {
	text = strings.TrimSuffix(strings.TrimPrefix(text, "/**"), "*/")
	var lines []string
	for _, l := range strings.Split(text, "\n") {
		l = strings.TrimSpace(l)
		l = strings.TrimPrefix(l, "* ")
		if l == "*" {
			l = ""
		}
		lines = append(lines, l)
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func (b *builder) typeAnn(c parser.ITypeAnnotationContext) *Name {
	t := c.(*parser.TypeAnnotationContext).TypeRef()
	if t == nil {
		return &Name{Sp: span(c.(antlr.ParserRuleContext))}
	}
	return &Name{Text: t.GetText(), Sp: span(t.(antlr.ParserRuleContext))}
}

func (b *builder) params(c *parser.ParamsContext) *ParamList {
	pl := &ParamList{Sp: span(c)}
	for _, p := range c.AllParam() {
		pc := p.(*parser.ParamContext)
		param := &Param{Sp: span(pc)}
		if real(pc.IDENT()) {
			param.Name = &Name{Text: pc.IDENT().GetText(), Sp: nodeSpan(pc.IDENT())}
		}
		if t := pc.TypeAnnotation(); t != nil {
			param.Type = b.typeAnn(t)
		}
		pl.Params = append(pl.Params, param)
	}
	return pl
}

func (b *builder) block(c parser.IBlockContext) *Block {
	ctx := c.(*parser.BlockContext)
	blk := &Block{Sp: span(ctx), Closed: real(ctx.RBRACE())}
	for _, item := range ctx.AllBlockItem() {
		ic := item.(*parser.BlockItemContext)
		if s := ic.Statement(); s != nil {
			if st := b.stmt(s); st != nil {
				blk.Items = append(blk.Items, st)
			}
		}
	}
	return blk
}

func (b *builder) stmt(c parser.IStatementContext) Stmt {
	sp := span(c.(antlr.ParserRuleContext))
	switch s := c.(type) {
	case *parser.LetStmtContext:
		st := &LetStmt{Sp: sp}
		if real(s.IDENT()) {
			st.Name = &Name{Text: s.IDENT().GetText(), Sp: nodeSpan(s.IDENT())}
		}
		if t := s.TypeAnnotation(); t != nil {
			st.Type = b.typeAnn(t)
		}
		if s.ASSIGN() != nil {
			st.Value = b.expr(s.Expr(), sp)
		}
		return st
	case *parser.IfStmtContext:
		return b.ifChain(sp, s.Expr(), s.AllBlock(), s.IfTail())
	case *parser.WhileStmtContext:
		return &WhileStmt{Cond: b.expr(s.Expr(), sp), Body: b.blockOrEmpty(s.Block(), sp), Sp: sp}
	case *parser.ForStmtContext:
		st := &ForStmt{Iter: b.expr(s.Expr(), sp), Body: b.blockOrEmpty(s.Block(), sp), Sp: sp}
		if real(s.IDENT()) {
			st.Var = &Name{Text: s.IDENT().GetText(), Sp: nodeSpan(s.IDENT())}
		}
		return st
	case *parser.ReturnStmtContext:
		st := &ReturnStmt{Sp: sp}
		if s.Expr() != nil {
			st.Value = b.expr(s.Expr(), sp)
		}
		return st
	case *parser.BreakStmtContext:
		return &BreakStmt{Sp: sp}
	case *parser.ContinueStmtContext:
		return &ContinueStmt{Sp: sp}
	case *parser.CommandStmtContext:
		text := strings.TrimRight(strings.TrimPrefix(s.COMMAND().GetText(), "/"), " \t\r")
		return &CommandStmt{Text: text, Sp: sp}
	case *parser.FieldStmtContext:
		key := s.FieldKey().(*parser.FieldKeyContext)
		name := key.GetText()
		if key.STRING() != nil {
			name = unquote(name)
		}
		return &FieldStmt{
			Key:       &Name{Text: name, Sp: span(key)},
			KeyQuoted: key.STRING() != nil,
			Value:     b.expr(s.Expr(), sp),
			Sp:        sp,
		}
	case *parser.CompoundAssignStmtContext:
		return &AssignStmt{
			Target: b.expr(s.Expr(0), sp),
			Op:     s.GetOp().GetText(),
			Value:  b.expr(s.Expr(1), sp),
			Sp:     sp,
		}
	case *parser.PathAssignStmtContext:
		ids := s.AllIDENT()
		var target Expr = &Ident{Name: ids[0].GetText(), Sp: tokSpan(ids[0].GetSymbol())}
		for _, id := range ids[1:] {
			n := &Name{Text: id.GetText(), Sp: tokSpan(id.GetSymbol())}
			target = &Member{X: target, Name: n, Sp: Span{sp.Start, n.Sp.End}}
		}
		return &AssignStmt{Target: target, Op: "=", Value: b.expr(s.Expr(), sp), Sp: sp}
	case *parser.IndexAssignStmtContext:
		id := s.IDENT()
		target := &Index{X: &Ident{Name: id.GetText(), Sp: tokSpan(id.GetSymbol())}, Args: b.args([]parser.IArgContext{s.Arg()}), Sp: Span{sp.Start, tokSpan(s.RBRACK().GetSymbol()).End}}
		return &AssignStmt{Target: target, Op: "=", Value: b.expr(s.Expr(), sp), Sp: sp}
	case *parser.ExprStmtContext:
		return &ExprStmt{X: b.expr(s.Expr(), sp), Sp: sp}
	}
	return nil
}

func (b *builder) ifChain(sp Span, cond parser.IExprContext, blocks []parser.IBlockContext, tail parser.IIfTailContext) *IfStmt {
	st := &IfStmt{Cond: b.expr(cond, sp), Sp: sp}
	if len(blocks) > 0 {
		st.Then = b.block(blocks[0])
	} else {
		st.Then = &Block{Sp: Span{sp.End, sp.End}}
	}
	if tail != nil {
		tc := tail.(*parser.IfTailContext)
		st.Else = b.ifChain(span(tc), tc.Expr(), tc.AllBlock(), tc.IfTail())
	} else if len(blocks) > 1 {
		st.Else = b.block(blocks[1])
	}
	return st
}

func (b *builder) blockOrEmpty(c parser.IBlockContext, sp Span) *Block {
	if c == nil {
		return &Block{Sp: Span{sp.End, sp.End}}
	}
	return b.block(c)
}

// expr converts an expression; a missing one becomes Bad at the end of
// the surrounding span.
func (b *builder) expr(c parser.IExprContext, around Span) Expr {
	if c == nil {
		return &Bad{Sp: Span{around.End, around.End}}
	}
	ctx := c.(antlr.ParserRuleContext)
	sp := span(ctx)
	if ctx.GetStart() == nil || ctx.GetStart().GetTokenType() == antlr.TokenEOF {
		return &Bad{Sp: sp}
	}
	switch e := c.(type) {
	case *parser.PrimaryExprContext:
		return b.primary(e.Primary(), sp)
	case *parser.PostfixExprContext:
		return b.postfix(b.expr(e.Expr(), sp), e.Postfix(), sp)
	case *parser.UnaryExprContext:
		return &Unary{Op: e.GetOp().GetText(), X: b.expr(e.Expr(), sp), Sp: sp}
	case *parser.MulExprContext:
		return b.binary(e.GetOp().GetText(), e.AllExpr(), sp)
	case *parser.AddExprContext:
		return b.binary(e.GetOp().GetText(), e.AllExpr(), sp)
	case *parser.CompareExprContext:
		return b.binary(e.GetOp().GetText(), e.AllExpr(), sp)
	case *parser.EqualExprContext:
		return b.binary(e.GetOp().GetText(), e.AllExpr(), sp)
	case *parser.AndExprContext:
		return b.binary("&&", e.AllExpr(), sp)
	case *parser.OrExprContext:
		return b.binary("||", e.AllExpr(), sp)
	case *parser.ArrowExprContext:
		return b.binary("->", e.AllExpr(), sp)
	case *parser.RangeExprContext:
		r := &Range{Lo: b.expr(e.Expr(0), sp), Sp: sp}
		if len(e.AllExpr()) > 1 {
			r.Hi = b.expr(e.Expr(1), sp)
		}
		return r
	case *parser.RangeToExprContext:
		return &Range{Hi: b.expr(e.Expr(), sp), Sp: sp}
	}
	return &Bad{Sp: sp}
}

func (b *builder) binary(op string, xs []parser.IExprContext, sp Span) Expr {
	bin := &Binary{Op: op, Sp: sp}
	if len(xs) > 0 {
		bin.L = b.expr(xs[0], sp)
	} else {
		bin.L = &Bad{Sp: Span{sp.Start, sp.Start}}
	}
	if len(xs) > 1 {
		bin.R = b.expr(xs[1], sp)
	} else {
		bin.R = &Bad{Sp: Span{sp.End, sp.End}}
	}
	return bin
}

func (b *builder) postfix(x Expr, c parser.IPostfixContext, sp Span) Expr {
	switch p := c.(type) {
	case *parser.MemberPostfixContext:
		m := &Member{X: x, Sp: sp}
		if p.GetName() != nil && p.GetName().GetStart() >= 0 {
			m.Name = &Name{Text: p.GetName().GetText(), Sp: tokSpan(p.GetName())}
		} else {
			// Kept in the tree so completion after a '.' knows the receiver.
			b.file.Diags = append(b.file.Diags, Diagnostic{
				Code: "syntax/missing-name", Severity: SeverityError,
				Message: "expected a name after '.'", Sp: Span{sp.End, sp.End},
			})
		}
		return m
	case *parser.CallPostfixContext:
		ac := p.Args().(*parser.ArgsContext)
		return &Call{Fn: x, Args: b.args(ac.AllArg()), ArgsSp: span(ac), Sp: sp}
	case *parser.IndexPostfixContext:
		return &Index{X: x, Args: b.args(p.AllArg()), Sp: sp}
	case *parser.BlockPostfixContext:
		return &Trailing{X: x, Block: b.block(p.Block()), Sp: sp}
	}
	return x
}

func (b *builder) args(cs []parser.IArgContext) []*Arg {
	var out []*Arg
	for _, c := range cs {
		ac := c.(*parser.ArgContext)
		a := &Arg{Sp: span(ac)}
		if ac.ASSIGN() != nil && real(ac.IDENT()) {
			a.Name = &Name{Text: ac.IDENT().GetText(), Sp: nodeSpan(ac.IDENT())}
		}
		a.Value = b.expr(ac.Expr(), a.Sp)
		out = append(out, a)
	}
	return out
}

func (b *builder) primary(c parser.IPrimaryContext, sp Span) Expr {
	text := c.GetText()
	switch p := c.(type) {
	case *parser.NumberLitContext:
		return &Number{Text: text, Sp: sp}
	case *parser.StringLitContext:
		return &String{Value: unquote(text), Raw: text, Sp: sp}
	case *parser.BoolLitContext:
		return &Bool{Value: text == "true", Sp: sp}
	case *parser.IdentRefContext:
		return &Ident{Name: text, Sp: sp}
	case *parser.ReslocRefContext:
		return &ResLoc{Text: text, Sp: sp}
	case *parser.TagRefContext:
		return &Tag{Name: strings.TrimPrefix(text, "#"), Sp: sp}
	case *parser.SelectorContext:
		return &Selector{Name: strings.TrimPrefix(text, "@"), Sp: sp}
	case *parser.CoordinateContext:
		return &Coord{Text: b.sourceText(sp), Sp: sp}
	case *parser.ParenExprContext:
		return &Paren{X: b.expr(p.Expr(), sp), Sp: sp}
	case *parser.ListLitContext:
		l := &List{Sp: sp}
		for _, e := range p.AllExpr() {
			l.Elems = append(l.Elems, b.expr(e, sp))
		}
		return l
	case *parser.ObjectLitContext:
		return &Object{Body: b.block(p.Block()), Sp: sp}
	}
	return &Bad{Sp: sp}
}

func (b *builder) sourceText(sp Span) string {
	if sp.Start < 0 || sp.End > len(b.file.Src) || sp.Start > sp.End {
		return ""
	}
	return string(b.file.Src[sp.Start:sp.End])
}

func unquote(s string) string {
	if len(s) >= 2 && s[0] == '\'' {
		s = `"` + strings.ReplaceAll(strings.ReplaceAll(s[1:len(s)-1], `\'`, `'`), `"`, `\"`) + `"`
	}
	if u, err := strconv.Unquote(s); err == nil {
		return u
	}
	return strings.Trim(s, `"'`)
}
