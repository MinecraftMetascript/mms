package parser

// IsNameToken reports whether a token type can be a declaration's name.
func IsNameToken(t int) bool {
	return t == MMSLexerIDENT || t == MMSLexerRESLOC || t == MMSLexerTNAME
}

// word reports whether the next token is the identifier w, a word that
// only means something in one position (such as `hint`).
func (p *MMSParser) word(w string) bool {
	t := p.GetTokenStream().LT(1)
	return t != nil && t.GetTokenType() == MMSLexerIDENT && t.GetText() == w
}
