package parser

import "github.com/antlr4-go/antlr/v4"

// valueFollows reports whether the token after '=' can start the value.
// It is false when the next line starts a new top-level declaration
// (`Kind name` at column 0), so a missing value does not swallow it.
func (p *MMSParser) valueFollows() bool {
	ts := p.GetTokenStream()
	next := ts.LT(1)
	if next == nil || next.GetTokenType() == antlr.TokenEOF {
		return false
	}
	prev := ts.LT(-1)
	return !(prev != nil && StartsDeclaration(ts, 1) && next.GetLine() > prev.GetLine())
}

// StartsDeclaration reports whether the token at LT(k) looks like the start
// of a top-level declaration: an identifier in column 0 followed by a name.
func StartsDeclaration(ts antlr.TokenStream, k int) bool {
	t := ts.LT(k)
	if t == nil || t.GetTokenType() != MMSLexerIDENT || t.GetColumn() != 0 {
		return false
	}
	switch ts.LT(k + 1).GetTokenType() {
	case MMSLexerIDENT, MMSLexerRESLOC:
		return true
	}
	return false
}
