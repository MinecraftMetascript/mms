package parser

import "github.com/antlr4-go/antlr/v4"

// atLineStart reports whether only spaces and tabs separate the current
// token's first character from the start of its line. The COMMAND rule uses
// it so that '/' starts a raw command only at the beginning of a line. The
// predicate sits after the '/' in the rule: a predicate at the very start
// would stop ANTLR from caching the lexer's start state, making every token
// pay for a full ATN closure.
func (l *MMSLexer) atLineStart() bool {
	in := l.GetInputStream()
	for i := l.TokenStartCharIndex - 1; i >= 0; i-- {
		switch in.GetTextFromInterval(antlr.NewInterval(i, i)) {
		case " ", "\t":
			continue
		case "\n", "\r":
			return true
		default:
			return false
		}
	}
	return true
}

// atCommandStart reports whether a '/' here starts a raw command: at the
// start of a line, or as the whole value of an assignment, after a lone
// '=' as in `let n = /data get block ~ ~ ~ Items`. No expression can
// start with '/', so that '/' is never a division.
func (l *MMSLexer) atCommandStart() bool {
	if l.atLineStart() {
		return true
	}
	in := l.GetInputStream()
	at := func(i int) string { return in.GetTextFromInterval(antlr.NewInterval(i, i)) }
	i := l.TokenStartCharIndex - 1
	for i >= 0 && (at(i) == " " || at(i) == "\t") {
		i--
	}
	if i < 0 || at(i) != "=" {
		return false
	}
	if i == 0 {
		return true
	}
	switch at(i - 1) {
	case "=", "!", "<", ">", "+", "-", "*", "/", "%":
		return false
	}
	return true
}
