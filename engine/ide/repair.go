package ide

import (
	"context"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// closers maps each bracket token to its partner.
var closers = map[syntax.TokenKind]syntax.TokenKind{
	parser.MMSLexerLPAREN: parser.MMSLexerRPAREN,
	parser.MMSLexerLBRACK: parser.MMSLexerRBRACK,
	parser.MMSLexerLBRACE: parser.MMSLexerRBRACE,
}

var closeText = map[syntax.TokenKind]string{
	parser.MMSLexerLPAREN: ")",
	parser.MMSLexerLBRACK: "]",
	parser.MMSLexerLBRACE: "}",
}

// repaired returns a copy of f's text for completing at off when the
// declaration being typed does not parse: the declaration is cut after
// the cursor's line and the brackets still open there are closed. The
// grammar is whitespace agnostic, so a half-typed line such as
// `blocks = [gra|` would otherwise read on into the lines below it and
// completion would lose its place. What follows the cursor's line in that
// declaration matters little to what fits at the cursor.
func repaired(f *syntax.File, off int) (string, bool) {
	src := f.Src
	lineStart, lineEnd := off, off
	for lineStart > 0 && src[lineStart-1] != '\n' {
		lineStart--
	}
	for lineEnd < len(src) && src[lineEnd] != '\n' {
		lineEnd++
	}
	// The declaration: from the last line before the cursor's that starts
	// at the left margin to the next one after it.
	itemStart := lineStart
	for itemStart > 0 && !topLine(src, itemStart) {
		itemStart--
		for itemStart > 0 && src[itemStart-1] != '\n' {
			itemStart--
		}
	}
	next := lineEnd
	for next < len(src) {
		next++
		if next < len(src) && src[next-1] == '\n' && topLine(src, next) {
			break
		}
	}
	next = min(next, len(src))
	broken := false
	for _, d := range f.Diags {
		if strings.HasPrefix(d.Code, "syntax/") && d.Sp.End >= itemStart && d.Sp.Start <= next {
			broken = true
			break
		}
	}
	if !broken {
		return "", false
	}
	var stack []syntax.Token
	for _, t := range f.Tokens {
		if t.Hidden || t.Synthetic || t.Sp.Start < itemStart {
			continue
		}
		if t.Sp.Start >= lineEnd {
			break
		}
		if _, open := closers[t.Kind]; open {
			stack = append(stack, t)
			continue
		}
		for i := len(stack) - 1; i >= 0; i-- {
			if closers[stack[i].Kind] == t.Kind {
				stack = stack[:i]
				break
			}
		}
	}
	// Before a comment that ends the line.
	at := lineEnd
	for _, t := range f.Tokens {
		if t.Sp.Start >= off && t.Sp.Start < lineEnd && t.IsComment() {
			at = t.Sp.Start
			break
		}
	}
	var sb strings.Builder
	sb.WriteString(string(src[:at]))
	for i := len(stack) - 1; i >= 0; i-- {
		sb.WriteString(closeText[stack[i].Kind])
	}
	sb.WriteString("\n")
	if next < len(src) {
		sb.WriteString(string(src[next:]))
	}
	return sb.String(), true
}

// topLine reports whether the line starting at i starts a top-level item:
// it begins at the left margin with something other than a closing
// bracket or a comment's continuation.
func topLine(src []rune, i int) bool {
	if i >= len(src) {
		return true
	}
	switch c := src[i]; {
	case c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '}' || c == ']' || c == ')' || c == '*':
		return false
	}
	return true
}

// repairedSession is a session on a copy of the project in which the
// line being typed at off in path is repaired, or nil.
func (s *Session) repairedSession(path string, off int) *Session {
	if s.repair || s.P == nil {
		return nil
	}
	u := s.R.Unit(path)
	if u == nil || u.File == nil {
		return nil
	}
	text, ok := repaired(u.File, off)
	if !ok {
		return nil
	}
	q := s.P.Fork()
	q.SetFile(path, text)
	r, err := q.Analyze(context.Background())
	if err != nil {
		return nil
	}
	return &Session{P: q, R: r, repair: true}
}
