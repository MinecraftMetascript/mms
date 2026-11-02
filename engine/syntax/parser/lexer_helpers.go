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

// afterAssign reports whether the current token follows '=' (an
// assignment, also a compound one) or the keyword return, with only
// spaces and tabs between. INLINE_COMMAND uses it so that '/' starts a
// command used as a value, while `a / b` stays a division. The command
// runs to the end of the line, so it only fits where the statement ends
// with it.
func (l *MMSLexer) afterAssign() bool {
	in := l.GetInputStream()
	i := l.TokenStartCharIndex - 1
	for ; i >= 0; i-- {
		switch in.GetTextFromInterval(antlr.NewInterval(i, i)) {
		case " ", "\t":
			continue
		case "=":
			if i == 0 {
				return true
			}
			switch in.GetTextFromInterval(antlr.NewInterval(i-1, i-1)) {
			case "=", "!", "<", ">":
				return false
			}
			return true
		case "n":
			return i >= 5 && in.GetTextFromInterval(antlr.NewInterval(i-5, i)) == "return" &&
				(i == 5 || !isWordChar(in.GetTextFromInterval(antlr.NewInterval(i-6, i-6))))
		default:
			return false
		}
	}
	return false
}

func isWordChar(s string) bool {
	return s != "" && (s[0] == '_' || s[0] >= 'a' && s[0] <= 'z' || s[0] >= 'A' && s[0] <= 'Z' || s[0] >= '0' && s[0] <= '9')
}
