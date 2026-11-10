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
// of the file. It reports each bracket it closes. The layout is only read
// when some bracket is never closed: a file whose brackets all match keeps
// its brackets, however it is laid out.
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
	unclosed := hasUnclosed(toks)
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
		if unclosed && len(stack) > 0 && t.GetColumn() == 0 && t.GetLine() > prevLine {
			if n := nextDefault(i); n != nil && startsTopItem(t, n.GetTokenType()) {
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

// hasUnclosed reports whether some opening bracket in toks is never
// closed, matching brackets the way closeUnbalanced does.
func hasUnclosed(toks []antlr.Token) bool {
	var stack []int
	for _, t := range toks {
		if t.GetChannel() != antlr.TokenDefaultChannel {
			continue
		}
		tt := t.GetTokenType()
		switch {
		case closerFor[tt].text != "":
			stack = append(stack, tt)
		case isCloser(tt):
			for j := len(stack) - 1; j >= 0; j-- {
				if closerFor[stack[j]].kind == tt {
					stack = stack[:j]
					break
				}
			}
		}
	}
	return len(stack) > 0
}

// startsTopItem reports whether a token pair starts a top-level item:
// `Kind name` with a PascalCase kind, `const`, `for` or `extern`
// followed by a word, or `private` followed by either.
func startsTopItem(first antlr.Token, second int) bool {
	switch first.GetTokenType() {
	case parser.MMSLexerIDENT:
		r := first.GetText()
		return r != "" && r[0] >= 'A' && r[0] <= 'Z' && parser.IsNameToken(second)
	case parser.MMSLexerCONST, parser.MMSLexerFOR, parser.MMSLexerEXTERN:
		return second == parser.MMSLexerIDENT
	case parser.MMSLexerPRIVATE:
		return second == parser.MMSLexerIDENT || second == parser.MMSLexerCONST || second == parser.MMSLexerTYPE
	}
	return false
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

// markMissingValues inserts a MISSING token after an '=' whose value was
// left out: the next token closes a bracket or is a comma, or the next
// line starts another `key = ...`. Without it the parser reads
// `default_block =` followed by `default_fluid = water` as
// `default_block = default_fluid` and then fails on the second '='. The
// same goes for an '=' at the end of the file, or followed by a line that
// starts a new top-level item (`Kind name` in column 0).
// No expression contains a bare '=', so a key followed by '=' on a later
// line always starts a new field or statement.
func markMissingValues(toks []antlr.Token, synthetic map[antlr.Token]bool) ([]antlr.Token, []Diagnostic) {
	var def []int // indices of default-channel tokens
	for i, t := range toks {
		if t.GetChannel() == antlr.TokenDefaultChannel {
			def = append(def, i)
		}
	}
	kind := func(k int) int {
		if k < len(def) {
			return toks[def[k]].GetTokenType()
		}
		return antlr.TokenEOF
	}
	missing := map[int]bool{} // toks index of an '=' missing its value
	depth := 0                // open brackets before the token
	for k, i := range def {
		eq := toks[i]
		switch tt := eq.GetTokenType(); {
		case closerFor[tt].text != "":
			depth++
		case isCloser(tt) && depth > 0:
			depth--
		}
		if eq.GetTokenType() != parser.MMSLexerASSIGN || k+1 >= len(def) {
			continue
		}
		next := toks[def[k+1]]
		switch next.GetTokenType() {
		case parser.MMSLexerCOMMAND:
			// A command on the line of the '=' is its value; one on the
			// next line is a statement of its own.
			if next.GetLine() != eq.GetLine() {
				missing[i] = true
			}
		case parser.MMSLexerRBRACE, parser.MMSLexerRPAREN, parser.MMSLexerRBRACK, parser.MMSLexerCOMMA,
			// Statements never start a value.
			parser.MMSLexerLET, parser.MMSLexerCONST, parser.MMSLexerWHILE,
			parser.MMSLexerRETURN, parser.MMSLexerBREAK, parser.MMSLexerCONTINUE:
			missing[i] = true
		case parser.MMSLexerIDENT, parser.MMSLexerRESLOC, parser.MMSLexerSTRING:
			// `a = b = 1` is never a value followed by more: b starts the
			// next field.
			if kind(k+2) == parser.MMSLexerASSIGN {
				missing[i] = true
			}
		case antlr.TokenEOF:
			missing[i] = true
		}
		// At the top level, `Kind name` is a new item, never the value:
		// no top-level item starts with a value. When the pair after it
		// starts an item too (`x = Floor` then `Noise n ...`), that later
		// pair is the item.
		if depth == 0 && k+2 < len(def) && startsTopItem(next, kind(k+2)) &&
			!(k+3 < len(def) && startsTopItem(toks[def[k+2]], kind(k+3))) {
			missing[i] = true
		}
	}
	if len(missing) == 0 {
		return toks, nil
	}
	var (
		out   []antlr.Token
		diags []Diagnostic
	)
	for i, t := range toks {
		out = append(out, t)
		if !missing[i] {
			continue
		}
		pos := t.GetStop() + 1
		tok := antlr.CommonTokenFactoryDEFAULT.Create(&antlr.TokenSourceCharStreamPair{},
			parser.MMSLexerMISSING, "", antlr.TokenDefaultChannel, pos, pos-1, t.GetLine(), t.GetColumn()+1)
		out = append(out, tok)
		synthetic[tok] = true
		diags = append(diags, Diagnostic{
			Code: "syntax/missing-value", Severity: SeverityError,
			Message: "expected a value after '='", Sp: tokSpan(t),
		})
	}
	return out, diags
}

// separateWords inserts a SEP token between two words inside braces when
// the first cannot be a declaration's kind: `x` then `y = 1`, or
// `orange_edge` then `Bandlands()`. Kinds are PascalCase, so a lowercase
// word, a resource location or a templated name always ends a statement
// there, and without the SEP the parser would read the pair as one
// declaration, `x y = 1`. The rule looks only at the tokens, never at line
// breaks, so `a b` and `a` then `b` on the next line parse the same.
// Inside brackets and parentheses nothing is inserted.
func separateWords(toks []antlr.Token, synthetic map[antlr.Token]bool) []antlr.Token {
	var (
		out   []antlr.Token
		stack []int
		prev  antlr.Token
	)
	for i, t := range toks {
		if t.GetChannel() != antlr.TokenDefaultChannel {
			out = append(out, t)
			continue
		}
		tt := t.GetTokenType()
		if prev != nil && len(stack) > 0 && stack[len(stack)-1] == parser.MMSLexerLBRACE &&
			notKind(prev) && startsName(tt) && !isHint(toks, i) {
			sep := antlr.CommonTokenFactoryDEFAULT.Create(&antlr.TokenSourceCharStreamPair{},
				parser.MMSLexerSEP, "", antlr.TokenDefaultChannel, prev.GetStop()+1, prev.GetStop(), prev.GetLine(), prev.GetColumn())
			synthetic[sep] = true
			out = append(out, sep)
		}
		switch {
		case closerFor[tt].text != "":
			stack = append(stack, tt)
		case isCloser(tt):
			for j := len(stack) - 1; j >= 0; j-- {
				if closerFor[stack[j]].kind == tt {
					stack = stack[:j]
					break
				}
			}
		}
		out = append(out, t)
		prev = t
	}
	return out
}

// notKind reports whether a token is a word that ends a statement and can
// never be a declaration's kind: a lowercase identifier, a resource
// location or a templated name.
func notKind(t antlr.Token) bool {
	switch t.GetTokenType() {
	case parser.MMSLexerRESLOC, parser.MMSLexerTNAME:
		return true
	case parser.MMSLexerIDENT:
		r := t.GetText()
		return r != "" && (r[0] == '_' || r[0] >= 'a' && r[0] <= 'z')
	}
	return false
}

// startsName reports whether a token of kind t can be a declaration's
// name (the declName rule).
func startsName(t int) bool {
	switch t {
	case parser.MMSLexerIDENT, parser.MMSLexerRESLOC, parser.MMSLexerTNAME, parser.MMSLexerTYPE,
		parser.MMSLexerPRIMARY, parser.MMSLexerAS, parser.MMSLexerCHECK, parser.MMSLexerEMIT,
		parser.MMSLexerSETUP, parser.MMSLexerEXPECT, parser.MMSLexerRUN:
		return true
	}
	return false
}

// isHint reports whether toks[i] starts the `hint(...)` that may follow
// the value of an `as` conversion.
func isHint(toks []antlr.Token, i int) bool {
	if toks[i].GetTokenType() != parser.MMSLexerIDENT || toks[i].GetText() != "hint" {
		return false
	}
	for _, t := range toks[i+1:] {
		if t.GetChannel() == antlr.TokenDefaultChannel {
			return t.GetTokenType() == parser.MMSLexerLPAREN
		}
	}
	return false
}
