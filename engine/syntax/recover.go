package syntax

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// errorStrategy is ANTLR's default recovery with one change: it never
// deletes a closing bracket to resynchronise. Deleting `)` in
// `If(a && ) { ... }` would make the trailing block the missing operand;
// keeping it leaves a Bad operand and the call intact.
type errorStrategy struct {
	*antlr.DefaultErrorStrategy
}

func (s *errorStrategy) Sync(r antlr.Parser) {
	_, topLevel := r.GetParserRuleContext().(*parser.FileContext)
	if !topLevel && isCloser(r.GetTokenStream().LA(1)) {
		return
	}
	s.DefaultErrorStrategy.Sync(r)
}

func isCloser(t int) bool {
	return t == parser.MMSLexerRPAREN || t == parser.MMSLexerRBRACK || t == parser.MMSLexerRBRACE
}

var closerFor = map[int]struct {
	kind int
	text string
}{
	parser.MMSLexerLPAREN: {parser.MMSLexerRPAREN, ")"},
	parser.MMSLexerLBRACK: {parser.MMSLexerRBRACK, "]"},
	parser.MMSLexerLBRACE: {parser.MMSLexerRBRACE, "}"},
}

// closeUnbalanced inserts synthetic closing brackets where the author most
// likely meant them: before a line that starts a new top-level declaration
// (`Kind name` in column 0) while brackets are still open, and at the end
// of the file. It reports each bracket it closes.
func closeUnbalanced(toks []antlr.Token) ([]antlr.Token, map[antlr.Token]bool, []Diagnostic) {
	var (
		out       []antlr.Token
		synthetic = map[antlr.Token]bool{}
		diags     []Diagnostic
		stack     []antlr.Token
		prevLine  = 0
	)
	closeAll := func(at antlr.Token) {
		for i := len(stack) - 1; i >= 0; i-- {
			open := stack[i]
			c := closerFor[open.GetTokenType()]
			pos := at.GetStart()
			tok := antlr.CommonTokenFactoryDEFAULT.Create(&antlr.TokenSourceCharStreamPair{},
				c.kind, c.text, antlr.TokenDefaultChannel, pos, pos-1, at.GetLine(), at.GetColumn())
			out = append(out, tok)
			synthetic[tok] = true
			diags = append(diags, Diagnostic{
				Code: "syntax/unclosed", Severity: SeverityError,
				Message: fmt.Sprintf("'%s' is never closed", open.GetText()),
				Sp:      tokSpan(open),
			})
		}
		stack = stack[:0]
	}
	nextDefault := func(i int) antlr.Token {
		for j := i + 1; j < len(toks); j++ {
			if toks[j].GetChannel() == antlr.TokenDefaultChannel {
				return toks[j]
			}
		}
		return nil
	}
	for i, t := range toks {
		if t.GetChannel() != antlr.TokenDefaultChannel {
			out = append(out, t)
			continue
		}
		tt := t.GetTokenType()
		if tt == antlr.TokenEOF {
			closeAll(t)
			out = append(out, t)
			continue
		}
		if len(stack) > 0 && tt == parser.MMSLexerIDENT && t.GetColumn() == 0 && t.GetLine() > prevLine {
			if n := nextDefault(i); n != nil && (n.GetTokenType() == parser.MMSLexerIDENT || n.GetTokenType() == parser.MMSLexerRESLOC) {
				closeAll(t)
			}
		}
		switch {
		case closerFor[tt].text != "":
			stack = append(stack, t)
		case isCloser(tt):
			// Pop to the matching opener; a stray closer is left for the
			// parser to report.
			for j := len(stack) - 1; j >= 0; j-- {
				if closerFor[stack[j].GetTokenType()].kind == tt {
					stack = stack[:j]
					break
				}
			}
		}
		prevLine = t.GetLine()
		out = append(out, t)
	}
	return out, synthetic, diags
}

// sliceSource replays a token slice. It embeds the lexer only to satisfy
// antlr.TokenSource's unexported methods.
type sliceSource struct {
	*parser.MMSLexer
	toks []antlr.Token
	i    int
}

func (s *sliceSource) NextToken() antlr.Token {
	t := s.toks[s.i]
	if s.i < len(s.toks)-1 {
		s.i++
	}
	return t
}
