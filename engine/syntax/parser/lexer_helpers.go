package parser

import "github.com/antlr4-go/antlr/v4"

// atLineStart reports whether only spaces and tabs separate the next
// character from the start of its line. The COMMAND rule uses it so that
// '/' starts a raw command only at the beginning of a line.
func (l *MMSLexer) atLineStart() bool {
	in := l.GetInputStream()
	for i := in.Index() - 1; i >= 0; i-- {
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
