package syntax

import (
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// TokenKind is the lexer token type (parser.MMSLexer* constants).
type TokenKind int

// Token is a lexed token with its text and span.
type Token struct {
	Kind   TokenKind
	Text   string
	Sp     Span
	Hidden bool // whitespace and comments
	// Synthetic tokens were inserted by recovery (an auto-closed bracket)
	// and do not appear in the source.
	Synthetic bool
}

func (t Token) IsComment() bool {
	switch t.Kind {
	case parser.MMSLexerLINE_COMMENT, parser.MMSLexerBLOCK_COMMENT, parser.MMSLexerDOC_COMMENT, parser.MMSLexerUNTERMINATED_COMMENT:
		return true
	}
	return false
}

func (t Token) IsNewline() bool { return t.Kind == parser.MMSLexerNEWLINE }

// Severity of a diagnostic, matching LSP's numbering.
type Severity int

const (
	SeverityError Severity = iota + 1
	SeverityWarning
	SeverityInfo
	SeverityHint
)

// Diagnostic is a problem found in a source file.
type Diagnostic struct {
	Code     string
	Severity Severity
	Message  string
	Sp       Span
	// Unnecessary marks code that can be removed (editors fade it).
	Unnecessary bool
	// Fixes are edits that resolve the problem, best first.
	Fixes []Fix
}

// Shift returns the diagnostic moved by delta runes, fixes included.
func (d Diagnostic) Shift(delta int) Diagnostic {
	d.Sp = Span{Start: d.Sp.Start + delta, End: d.Sp.End + delta}
	if len(d.Fixes) > 0 {
		fixes := make([]Fix, len(d.Fixes))
		for i, f := range d.Fixes {
			edits := make([]TextEdit, len(f.Edits))
			for j, e := range f.Edits {
				edits[j] = TextEdit{Sp: Span{Start: e.Sp.Start + delta, End: e.Sp.End + delta}, Text: e.Text}
			}
			f.Edits = edits
			fixes[i] = f
		}
		d.Fixes = fixes
	}
	return d
}

// Fix is a named set of edits to the diagnostic's file.
type Fix struct {
	Title string
	Edits []TextEdit
	// Safe fixes never change the generated pack's behaviour; "fix all"
	// applies only these.
	Safe bool
}

// TextEdit replaces a span with new text.
type TextEdit struct {
	Sp   Span
	Text string
}

// TokenAt returns the index of the token containing offset off, preferring
// the token that ends at off (the one just typed) over the one starting
// there. It returns -1 for an empty file.
func (f *File) TokenAt(off int) int {
	best := -1
	for i, t := range f.Tokens {
		if t.Sp.Start <= off && off <= t.Sp.End {
			if best == -1 || (t.Sp.End == off && !t.Hidden) {
				best = i
			}
			if t.Sp.Start < off && off < t.Sp.End {
				return i
			}
		}
	}
	return best
}

// LineCol returns the 1-based line and column (in runes) of an offset.
func (f *File) LineCol(off int) (line, col int) {
	line, col = 1, 1
	for i, r := range f.Src {
		if i >= off {
			break
		}
		if r == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return line, col
}

// LineText returns the text of the 1-based line.
func (f *File) LineText(line int) string {
	n := 1
	start := 0
	for i, r := range f.Src {
		if r == '\n' {
			if n == line {
				return string(f.Src[start:i])
			}
			n++
			start = i + 1
		}
	}
	if n == line {
		return string(f.Src[start:])
	}
	return ""
}
