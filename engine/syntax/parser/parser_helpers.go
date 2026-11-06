package parser

import "github.com/antlr4-go/antlr/v4"

// IsNameToken reports whether a token type can be a declaration's name.
func IsNameToken(t int) bool {
	return t == MMSLexerIDENT || t == MMSLexerRESLOC || t == MMSLexerTNAME
}

// onNewLine reports whether the next token starts a later line than the
// previous one. List items may be separated by a newline instead of a
// comma.
func (p *MMSParser) onNewLine() bool {
	ts := p.GetTokenStream()
	next, prev := ts.LT(1), ts.LT(-1)
	if next == nil || prev == nil || next.GetTokenType() == antlr.TokenEOF {
		return false
	}
	switch next.GetTokenType() {
	case MMSLexerRBRACK, MMSLexerCOMMA:
		return false
	}
	return next.GetLine() > prev.GetLine()
}

// word reports whether the next token is the identifier w, a word that
// only means something in one position (such as `hint`).
func (p *MMSParser) word(w string) bool {
	t := p.GetTokenStream().LT(1)
	return t != nil && t.GetTokenType() == MMSLexerIDENT && t.GetText() == w
}
