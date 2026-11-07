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
	toks, missingDiags := markMissingValues(toks, synthetic)
	toks = separateWords(toks, synthetic)
	closeDiags = append(closeDiags, missingDiags...)
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
	for _, it := range tree.AllTopItem() {
		if item := b.topItem(it); item != nil {
			f.Items = append(f.Items, item)
			if d, ok := item.(*Decl); ok {
				f.Decls = append(f.Decls, d)
			}
		}
	}
	return f
}

// parseExpr parses the text of a `${...}` piece in a name, with spans
// shifted by off. Problems are added to into.
func parseExpr(text string, off int, into *File) Expr {
	sub := &File{Src: []rune(text)}
	lexer := parser.NewMMSLexer(antlr.NewInputStream(text))
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()
	for _, t := range stream.GetAllTokens() {
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		sub.Tokens = append(sub.Tokens, Token{Kind: TokenKind(t.GetTokenType()), Text: t.GetText(),
			Sp: Span{t.GetStart(), t.GetStop() + 1}, Hidden: t.GetChannel() != antlr.TokenDefaultChannel})
	}
	p := parser.NewMMSParser(stream)
	p.RemoveErrorListeners()
	el := &errorListener{synthetic: map[antlr.Token]bool{}}
	p.AddErrorListener(el)
	p.SetErrorHandler(&errorStrategy{antlr.NewDefaultErrorStrategy()})
	tree := p.ExprOnly()
	b := builder{file: sub}
	x := b.expr(tree.Expr(), Span{0, len(sub.Src)})
	s := shifter{d: off}
	for _, d := range append(el.diags, sub.Diags...) {
		into.Diags = append(into.Diags, d.Shift(off))
	}
	return s.expr(x)
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
		// MISSING is only ever inserted by recovery; never ask for it.
		"{MISSING, ", "{",
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

func (b *builder) topItem(c parser.ITopItemContext) Stmt {
	sp := span(c.(antlr.ParserRuleContext))
	switch it := c.(type) {
	case *parser.DeclItemContext:
		d := b.decl(it.Declaration())
		if d != nil && it.EXTERN() != nil {
			d.Extern = true
			d.Sp = sp
			d.Doc = b.docBefore(it.GetStart().GetTokenIndex())
		}
		if d == nil {
			return nil
		}
		return d
	case *parser.ConstItemContext:
		return b.constDecl(it.ConstDecl())
	case *parser.LetItemContext:
		return b.letDecl(it.LetDecl())
	case *parser.ForItemContext:
		st := b.forStmt(sp, it.AllIDENT(), it.Expr(), it.Block())
		nestDecls(st.Body)
		return st
	}
	return nil
}

// declContext is what declaration and innerDecl contexts share.
type declContext interface {
	antlr.ParserRuleContext
	GetKind() antlr.Token
	AllDeclName() []parser.IDeclNameContext
	TypeAnnotation() parser.ITypeAnnotationContext
	Params() parser.IParamsContext
}

func (b *builder) decl(c parser.IDeclarationContext) *Decl {
	if c == nil {
		return nil
	}
	ctx := c.(*parser.DeclarationContext)
	d := b.declFrom(ctx)
	if d != nil && ctx.TypeExpr() != nil {
		d.TypeValue = b.typeExpr(ctx.TypeExpr())
	}
	return d
}

func (b *builder) innerDecl(c parser.IInnerDeclContext) *Decl {
	if c == nil {
		return nil
	}
	return b.declFrom(c.(*parser.InnerDeclContext))
}

func (b *builder) declFrom(ctx declContext) *Decl {
	if ctx.GetKind() == nil {
		return nil
	}
	d := &Decl{
		Kind: &Ident{Name: ctx.GetKind().GetText(), Sp: tokSpan(ctx.GetKind())},
		Sp:   span(ctx),
	}
	for i, n := range ctx.AllDeclName() {
		name := b.declName(n.(*parser.DeclNameContext))
		if i == 0 {
			d.Name = name
		} else {
			d.More = append(d.More, name)
		}
	}
	if t := ctx.TypeAnnotation(); t != nil {
		d.Type = b.typeAnn(t)
	}
	if p := ctx.Params(); p != nil {
		d.Params = b.params(p.(*parser.ParamsContext))
	}
	var body interface {
		antlr.ParserRuleContext
		Block() parser.IBlockContext
		Value() parser.IValueContext
	}
	switch c := ctx.(type) {
	case *parser.DeclarationContext:
		if blk := c.Block(); blk != nil {
			d.Body = b.block(blk)
		}
		if c.DeclBody() != nil {
			body = c.DeclBody().(*parser.DeclBodyContext)
		}
	case *parser.InnerDeclContext:
		if c.InnerBody() != nil {
			body = c.InnerBody().(*parser.InnerBodyContext)
		}
	}
	if body != nil {
		bc := body
		switch {
		case bc.Block() != nil:
			d.Body = b.block(bc.Block())
		case bc.Value() != nil:
			d.Value = b.value(bc.Value(), span(bc))
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

// declName builds a declaration name, splitting a templated one into its
// pieces.
func (b *builder) declName(n *parser.DeclNameContext) *Name {
	name := &Name{Text: n.GetText(), Sp: span(n)}
	if n.TNAME() != nil {
		name.Parts = b.nameParts(name.Text, name.Sp.Start)
	}
	return name
}

// nameParts splits a templated name such as `blocks/${w}_crate` into
// literal and `${...}` pieces. off is the name's offset in the file.
func (b *builder) nameParts(text string, off int) []NamePart {
	var parts []NamePart
	rs := []rune(text)
	lit := 0
	for i := 0; i < len(rs); i++ {
		if rs[i] != '$' || i+1 >= len(rs) || rs[i+1] != '{' {
			continue
		}
		end := i + 2
		for end < len(rs) && rs[end] != '}' {
			end++
		}
		if lit < i {
			parts = append(parts, NamePart{Lit: string(rs[lit:i]), Sp: Span{off + lit, off + i}})
		}
		inner := string(rs[i+2 : end])
		part := NamePart{Sp: Span{off + i, off + end + 1}}
		if strings.TrimSpace(inner) == "" {
			part.X = &Bad{Sp: Span{off + i + 2, off + i + 2}}
			b.file.Diags = append(b.file.Diags, Diagnostic{
				Code: "syntax/empty-interpolation", Severity: SeverityError,
				Message: "expected an expression inside '${}'", Sp: part.Sp,
			})
		} else {
			part.X = parseExpr(inner, off+i+2, b.file)
		}
		parts = append(parts, part)
		i = end
		lit = end + 1
	}
	if lit < len(rs) {
		parts = append(parts, NamePart{Lit: string(rs[lit:]), Sp: Span{off + lit, off + len(rs)}})
	}
	return parts
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

func (b *builder) typeAnn(c parser.ITypeAnnotationContext) TypeExpr {
	tc := c.(*parser.TypeAnnotationContext)
	if tc.TypeExpr() == nil {
		sp := span(tc)
		return &TypeName{Sp: Span{sp.End, sp.End}}
	}
	return b.typeExpr(tc.TypeExpr())
}

func (b *builder) typeExpr(c parser.ITypeExprContext) TypeExpr {
	ctx := c.(*parser.TypeExprContext)
	sp := span(ctx)
	arms := ctx.AllTypeArm()
	if len(arms) == 0 {
		return &TypeName{Sp: Span{sp.Start, sp.Start}}
	}
	if len(arms) == 1 {
		return b.typeArm(arms[0].(*parser.TypeArmContext))
	}
	u := &TypeUnion{Sp: sp}
	for _, a := range arms {
		u.Arms = append(u.Arms, b.typeArm(a.(*parser.TypeArmContext)))
	}
	return u
}

func (b *builder) typeArm(c *parser.TypeArmContext) TypeExpr {
	var t TypeExpr
	if c.TypeAtom() != nil {
		t = b.typeAtom(c.TypeAtom())
	} else {
		t = &TypeName{Sp: span(c)}
	}
	if c.QUESTION() != nil {
		return &TypeOptional{X: t, Sp: span(c)}
	}
	return t
}

func (b *builder) typeAtom(c parser.ITypeAtomContext) TypeExpr {
	sp := span(c.(antlr.ParserRuleContext))
	switch t := c.(type) {
	case *parser.NamedTypeContext:
		n := &TypeName{Name: t.TypeName().GetText(), Sp: sp}
		for _, a := range t.AllTypeExpr() {
			n.Args = append(n.Args, b.typeExpr(a))
		}
		return n
	case *parser.LiteralTypeContext:
		raw := t.STRING().GetText()
		return &TypeString{Value: unquote(raw), Raw: raw, Sp: sp}
	case *parser.ListTypeContext:
		return &TypeList{Elem: b.typeOrEmpty(t.TypeExpr(), sp), Sp: sp}
	case *parser.RecordTypeContext:
		r := &TypeRecord{Sp: sp}
		for _, fc := range t.AllTypeField() {
			f := fc.(*parser.TypeFieldContext)
			fd := &FieldDecl{Name: &Name{Text: f.IDENT().GetText(), Sp: nodeSpan(f.IDENT())}, Sp: span(f)}
			if f.TypeAnnotation() != nil {
				fd.Type = b.typeAnn(f.TypeAnnotation())
			}
			if f.ASSIGN() != nil {
				fd.Default = b.value(f.Value(), fd.Sp)
			}
			fd.Doc = b.docBefore(f.GetStart().GetTokenIndex())
			r.Fields = append(r.Fields, fd)
		}
		// `{string: T}` and `{Kind: T}` are maps: a record's fields are
		// lowercase names, and no field is called string.
		if len(r.Fields) == 1 {
			f := r.Fields[0]
			if n := f.Name.Text; f.Default == nil && f.Type != nil && (n == "string" || n[0] >= 'A' && n[0] <= 'Z') {
				return &TypeMap{Key: &TypeName{Name: n, Sp: f.Name.Sp}, Value: f.Type, Sp: sp}
			}
		}
		return r
	case *parser.MapTypeContext:
		return &TypeMap{Key: b.mapKey(t.MapKey().(*parser.MapKeyContext)), Value: b.typeOrEmpty(t.TypeExpr(), sp), Sp: sp}
	case *parser.ParenTypeContext:
		return &TypeParen{X: b.typeOrEmpty(t.TypeExpr(), sp), Sp: sp}
	}
	return &TypeName{Sp: sp}
}

func (b *builder) mapKey(c *parser.MapKeyContext) TypeExpr {
	sp := span(c)
	switch {
	case c.TypeName() != nil:
		n := &TypeName{Name: c.TypeName().GetText(), Sp: sp}
		for _, a := range c.AllTypeExpr() {
			n.Args = append(n.Args, b.typeExpr(a))
		}
		return n
	case c.RESLOC() != nil:
		return &TypeName{Name: c.RESLOC().GetText(), Sp: sp}
	case len(c.AllTypeExpr()) == 1:
		return &TypeParen{X: b.typeExpr(c.TypeExpr(0)), Sp: sp}
	}
	return &TypeName{Sp: sp}
}

func (b *builder) typeOrEmpty(c parser.ITypeExprContext, around Span) TypeExpr {
	if c == nil {
		return &TypeName{Sp: Span{around.End, around.End}}
	}
	return b.typeExpr(c)
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
		if pc.ASSIGN() != nil {
			param.Default = b.value(pc.Value(), param.Sp)
		}
		pl.Params = append(pl.Params, param)
	}
	return pl
}

func (b *builder) constDecl(c parser.IConstDeclContext) *ConstStmt {
	ctx := c.(*parser.ConstDeclContext)
	sp := span(ctx)
	st := &ConstStmt{Sp: sp}
	if real(ctx.IDENT()) {
		st.Name = &Name{Text: ctx.IDENT().GetText(), Sp: nodeSpan(ctx.IDENT())}
	}
	if t := ctx.TypeAnnotation(); t != nil {
		st.Type = b.typeAnn(t)
	}
	if ctx.ASSIGN() != nil {
		st.Value = b.value(ctx.Value(), sp)
	}
	st.Doc = b.docBefore(ctx.GetStart().GetTokenIndex())
	return st
}

func (b *builder) letDecl(c parser.ILetDeclContext) *LetStmt {
	ctx := c.(*parser.LetDeclContext)
	sp := span(ctx)
	st := &LetStmt{Sp: sp}
	if real(ctx.IDENT()) {
		st.Name = &Name{Text: ctx.IDENT().GetText(), Sp: nodeSpan(ctx.IDENT())}
	}
	if t := ctx.TypeAnnotation(); t != nil {
		st.Type = b.typeAnn(t)
	}
	if ctx.ASSIGN() != nil {
		st.Value = b.value(ctx.Value(), sp)
	}
	st.Doc = b.docBefore(ctx.GetStart().GetTokenIndex())
	return st
}

func (b *builder) forStmt(sp Span, vars []antlr.TerminalNode, iter parser.IExprContext, body parser.IBlockContext) *ForStmt {
	st := &ForStmt{Iter: b.expr(iter, sp), Body: b.blockOrEmpty(body, sp), Sp: sp}
	if len(vars) > 0 && real(vars[0]) {
		st.Var = &Name{Text: vars[0].GetText(), Sp: nodeSpan(vars[0])}
	}
	if len(vars) > 1 && real(vars[1]) {
		st.Var2 = &Name{Text: vars[1].GetText(), Sp: nodeSpan(vars[1])}
	}
	return st
}

func (b *builder) block(c parser.IBlockContext) *Block {
	ctx := c.(*parser.BlockContext)
	blk := &Block{Sp: span(ctx), Closed: real(ctx.RBRACE())}
	for _, item := range ctx.AllBlockItem() {
		ic := item.(*parser.BlockItemContext)
		if s := ic.Statement(); s != nil {
			st := b.stmt(s)
			if d, ok := st.(*Decl); ok && !d.Primary && lowerWords(d) {
				// `say hi`: two words, not a declaration, whose kinds
				// are PascalCase.
				blk.Items = append(blk.Items,
					&ExprStmt{X: &Ident{Name: d.Kind.Name, Sp: d.Kind.Sp}, Sp: d.Kind.Sp},
					&ExprStmt{X: &Ident{Name: d.Name.Text, Sp: d.Name.Sp}, Sp: d.Name.Sp})
				continue
			}
			if st != nil {
				blk.Items = append(blk.Items, st)
			}
		}
	}
	return blk
}

// lowerWords reports whether a nested declaration is just two lowercase
// words, as in `say hi` after a '/' that does not start a line.
func lowerWords(d *Decl) bool {
	k := d.Kind.Name
	return k != "" && k[0] >= 'a' && k[0] <= 'z' && d.Name != nil && !d.Name.Templated() &&
		len(d.More) == 0 && d.Type == nil && d.Params == nil && d.Value == nil && d.Body == nil
}

// nestDecls turns `Kind { ... }` statements into declarations in a block
// that holds declarations: a top-level for or an emit block, and the
// if and for blocks inside them. The parser reads them as an expression
// with a trailing block, since in a function body that is what they are.
func nestDecls(blk *Block) {
	if blk == nil {
		return
	}
	for i, it := range blk.Items {
		switch st := it.(type) {
		case *ExprStmt:
			if t, ok := st.X.(*Trailing); ok {
				if id, ok := t.X.(*Ident); ok && id.Name != "" && id.Name[0] >= 'A' && id.Name[0] <= 'Z' {
					blk.Items[i] = &Decl{Kind: id, Body: t.Block, Sp: st.Sp}
				}
			}
		case *IfStmt:
			for st != nil {
				nestDecls(st.Then)
				switch e := st.Else.(type) {
				case *IfStmt:
					st = e
					continue
				case *Block:
					nestDecls(e)
				}
				st = nil
			}
		case *ForStmt:
			nestDecls(st.Body)
		}
	}
}

func (b *builder) stmt(c parser.IStatementContext) Stmt {
	sp := span(c.(antlr.ParserRuleContext))
	switch s := c.(type) {
	case *parser.LetStmtContext:
		st := b.letDecl(s.LetDecl())
		st.Sp = sp
		return st
	case *parser.ConstStmtContext:
		st := b.constDecl(s.ConstDecl())
		st.Sp = sp
		return st
	case *parser.IfStmtContext:
		return b.ifChain(sp, s.Expr(), s.ThenPart(), s.ElsePart())
	case *parser.WhileStmtContext:
		return &WhileStmt{Cond: b.expr(s.Expr(), sp), Body: b.blockOrEmpty(s.Block(), sp), Sp: sp}
	case *parser.ForStmtContext:
		return b.forStmt(sp, s.AllIDENT(), s.Expr(), s.Block())
	case *parser.ReturnStmtContext:
		st := &ReturnStmt{Sp: sp}
		if s.Value() != nil {
			st.Value = b.value(s.Value(), sp)
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
			Value:     b.value(s.Value(), sp),
			Sp:        sp,
		}
	case *parser.CompoundAssignStmtContext:
		if s.GetOp() == nil {
			return &ExprStmt{X: b.expr(s.Expr(0), sp), Sp: sp}
		}
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
		return &AssignStmt{Target: target, Op: "=", Value: b.value(s.Value(), sp), Sp: sp}
	case *parser.IndexAssignStmtContext:
		id := s.IDENT()
		target := &Index{X: &Ident{Name: id.GetText(), Sp: tokSpan(id.GetSymbol())}, Args: b.args([]parser.IArgContext{s.Arg()}), Sp: Span{sp.Start, tokSpan(s.RBRACK().GetSymbol()).End}}
		return &AssignStmt{Target: target, Op: "=", Value: b.value(s.Value(), sp), Sp: sp}
	case *parser.ExprStmtContext:
		return &ExprStmt{X: b.expr(s.Expr(), sp), Sp: sp}
	case *parser.PrimaryStmtContext:
		d := b.innerDecl(s.InnerDecl())
		if d == nil {
			return nil
		}
		d.Primary = true
		d.Sp = sp
		d.Doc = b.docBefore(s.GetStart().GetTokenIndex())
		return d
	case *parser.DeclStmtContext:
		if d := b.innerDecl(s.InnerDecl()); d != nil {
			return d
		}
		return nil
	case *parser.FieldDeclStmtContext:
		st := &FieldDecl{Name: &Name{Text: s.IDENT().GetText(), Sp: nodeSpan(s.IDENT())}, Sp: sp}
		st.Type = b.typeAnn(s.TypeAnnotation())
		if s.ASSIGN() != nil {
			st.Default = b.value(s.Value(), sp)
		}
		st.Doc = b.docBefore(s.GetStart().GetTokenIndex())
		return st
	case *parser.SectionStmtContext:
		st := &SectionStmt{
			Word: &Ident{Name: s.GetSection().GetText(), Sp: tokSpan(s.GetSection())},
			Body: b.blockOrEmpty(s.Block(), sp),
			Sp:   sp,
		}
		if st.Word.Name == "emit" {
			nestDecls(st.Body)
		}
		return st
	case *parser.AsStmtContext:
		st := &AsStmt{Type: b.typeExpr(s.TypeExpr()), Value: b.value(s.Value(), sp), Sp: sp}
		if h := s.Hint(); h != nil {
			hc := h.(*parser.HintContext)
			ac := hc.Args().(*parser.ArgsContext)
			st.Hint = &Call{
				Fn:     &Ident{Name: hc.IDENT().GetText(), Sp: nodeSpan(hc.IDENT())},
				Args:   b.args(ac.AllArg()),
				ArgsSp: span(ac),
				Sp:     span(hc),
			}
		}
		st.Doc = b.docBefore(s.GetStart().GetTokenIndex())
		return st
	case *parser.ExpectStmtContext:
		return &ExpectStmt{X: b.expr(s.Expr(), sp), Sp: sp}
	case *parser.RunStmtContext:
		return &RunStmt{X: b.bareCall(s.BareCall(), sp), Sp: sp}
	}
	return nil
}

func (b *builder) ifChain(sp Span, cond parser.IExprContext, then parser.IThenPartContext, els parser.IElsePartContext) *IfStmt {
	st := &IfStmt{Cond: b.expr(cond, sp), Sp: sp}
	if then != nil {
		tc := then.(*parser.ThenPartContext)
		st.Then = b.part(tc.Block(), tc.BareCall(), span(tc))
	} else {
		st.Then = &Block{Sp: Span{sp.End, sp.End}}
	}
	if els != nil {
		ec := els.(*parser.ElsePartContext)
		if tail := ec.IfTail(); tail != nil {
			tc := tail.(*parser.IfTailContext)
			st.Else = b.ifChain(span(tc), tc.Expr(), tc.ThenPart(), tc.ElsePart())
		} else {
			st.Else = b.part(ec.Block(), ec.BareCall(), span(ec))
		}
	}
	return st
}

// part builds the branch of an if statement: a block, or a single
// expression kept as a bare block.
func (b *builder) part(blk parser.IBlockContext, call parser.IBareCallContext, sp Span) *Block {
	if blk != nil {
		return b.block(blk)
	}
	e := b.bareCall(call, sp)
	return &Block{Items: []Stmt{&ExprStmt{X: e, Sp: e.Span()}}, Closed: true, Bare: true, Sp: e.Span()}
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
	case *parser.RangeExprContext:
		r := &Range{Lo: b.expr(e.Expr(0), sp), Sp: sp}
		if len(e.AllExpr()) > 1 {
			r.Hi = b.expr(e.Expr(1), sp)
		}
		return r
	case *parser.RangeToExprContext:
		return &Range{Hi: b.expr(e.Expr(), sp), Sp: sp}
	case *parser.CoalesceExprContext:
		return b.binary("??", e.AllExpr(), sp)
	case *parser.SpreadExprContext:
		return &Spread{X: b.expr(e.Expr(), sp), Sp: sp}
	}
	return &Bad{Sp: sp}
}

// bareCall converts a call written without braces around it, as in
// `if (x) error(...)` or a Test's `run f()`.
func (b *builder) bareCall(call parser.IBareCallContext, around Span) Expr {
	if call == nil {
		return &Bad{Sp: Span{around.End, around.End}}
	}
	cc := call.(*parser.BareCallContext)
	var e Expr = &Ident{Name: cc.IDENT().GetText(), Sp: nodeSpan(cc.IDENT())}
	for _, pf := range cc.AllPostfix() {
		psp := span(pf.(antlr.ParserRuleContext))
		e = b.postfix(e, pf, Span{e.Span().Start, psp.End})
	}
	return e
}

// value converts a value: an if expression, or an expression.
func (b *builder) value(c parser.IValueContext, around Span) Expr {
	if c == nil {
		return &Bad{Sp: Span{around.End, around.End}}
	}
	ctx := c.(*parser.ValueContext)
	if t := ctx.COMMAND(); t != nil {
		sp := span(ctx)
		return &CommandStmt{Text: strings.TrimRight(strings.TrimPrefix(t.GetText(), "/"), " \t\r"), Sp: sp}
	}
	if ctx.IF() == nil {
		return b.expr(ctx.Expr(), around)
	}
	sp := span(ctx)
	x := &IfExpr{Cond: b.expr(ctx.Expr(), sp), Sp: sp}
	vs := ctx.AllValue()
	if len(vs) > 0 {
		x.Then = b.value(vs[0], sp)
	} else {
		x.Then = &Bad{Sp: Span{sp.End, sp.End}}
	}
	if len(vs) > 1 {
		x.Else = b.value(vs[1], sp)
	}
	return x
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
		if p.GetName() != nil && p.GetName().GetStart().GetStart() >= 0 {
			m.Name = &Name{Text: p.GetName().GetText(), Sp: span(p.GetName())}
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
		if ac.ASSIGN() != nil && ac.Word() != nil {
			a.Name = &Name{Text: ac.Word().GetText(), Sp: span(ac.Word())}
		}
		a.Value = b.value(ac.Value(), a.Sp)
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
	case *parser.TemplatedRefContext:
		return &TemplatedName{Text: text, Parts: b.nameParts(text, sp.Start), Sp: sp}
	case *parser.TagRefContext:
		return &Tag{Name: strings.TrimPrefix(text, "#"), Sp: sp}
	case *parser.RunExprContext:
		return &RunExpr{Body: b.blockOrEmpty(p.Block(), sp), Sp: sp}
	case *parser.ComprehensionContext:
		c := &Comprehension{Sp: sp}
		ids := p.AllIDENT()
		if len(ids) > 0 && real(ids[0]) {
			c.Var = &Name{Text: ids[0].GetText(), Sp: nodeSpan(ids[0])}
		}
		if len(ids) > 1 && real(ids[1]) {
			c.Var2 = &Name{Text: ids[1].GetText(), Sp: nodeSpan(ids[1])}
		}
		xs := p.AllExpr()
		c.Iter = b.exprAt(xs, 0, sp)
		if p.IF() != nil {
			c.Cond = b.exprAt(xs, 1, sp)
		}
		if p.Value() != nil {
			c.Body = b.value(p.Value(), sp)
		} else {
			c.Body = &Bad{Sp: Span{sp.End, sp.End}}
		}
		return c
	case *parser.SelectorContext:
		return &Selector{Name: strings.TrimPrefix(text, "@"), Sp: sp}
	case *parser.CoordinateContext:
		return &Coord{Text: coordText(p), Sp: sp}
	case *parser.ParenExprContext:
		return &Paren{X: b.value(p.Value(), sp), Sp: sp}
	case *parser.ListLitContext:
		l := &List{Sp: sp}
		for _, e := range p.AllValue() {
			l.Elems = append(l.Elems, b.value(e, sp))
		}
		return l
	case *parser.ObjectLitContext:
		return &Object{Body: b.block(p.Block()), Sp: sp}
	}
	return &Bad{Sp: sp}
}

// exprAt converts the i'th expression of a list, or a Bad one when
// recovery left it out.
func (b *builder) exprAt(xs []parser.IExprContext, i int, around Span) Expr {
	if i < len(xs) {
		return b.expr(xs[i], around)
	}
	return &Bad{Sp: Span{around.End, around.End}}
}

// unquote reads a string literal's value. It resolves the language's own
// escapes (\" \' \\ \n \t \r \uHHHH) and keeps every other backslash as
// written, since \{, \$, \* and the like belong to the markup and
// interpolation layers, which read the value after this.
func unquote(s string) string {
	if len(s) < 2 {
		return s
	}
	body := []rune(s[1 : len(s)-1])
	var b strings.Builder
	for i := 0; i < len(body); i++ {
		r := body[i]
		if r != '\\' || i+1 == len(body) {
			b.WriteRune(r)
			continue
		}
		i++
		switch c := body[i]; c {
		case 'n':
			b.WriteRune('\n')
		case 't':
			b.WriteRune('\t')
		case 'r':
			b.WriteRune('\r')
		case '"', '\'', '\\':
			b.WriteRune(c)
		case 'u':
			if i+4 < len(body) {
				if v, err := strconv.ParseUint(string(body[i+1:i+5]), 16, 32); err == nil {
					b.WriteRune(rune(v))
					i += 4
					continue
				}
			}
			b.WriteString(`\u`)
		default:
			b.WriteRune('\\')
			b.WriteRune(c)
		}
	}
	return b.String()
}

// coordText spells a coordinate group from its tokens, one space between
// coordinates and none inside one, so `^ ^-1.7` reads the same however it
// is spaced or broken across lines.
func coordText(p *parser.CoordinateContext) string {
	var b strings.Builder
	for _, c := range p.GetChildren() {
		t, ok := c.(antlr.TerminalNode)
		if !ok {
			continue
		}
		switch tt := t.GetSymbol().GetTokenType(); tt {
		case parser.MMSLexerTILDE, parser.MMSLexerCARET:
			if b.Len() > 0 {
				b.WriteByte(' ')
			}
		}
		b.WriteString(t.GetText())
	}
	return b.String()
}
