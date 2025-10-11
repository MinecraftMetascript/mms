package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// DiagnosticsErrorListener collects syntax errors from ANTLR and stores them
// into the provided slice. It uses the file content to populate the
// TextLocation text and accurate columns.
type DiagnosticsErrorListener struct {
	antlr.DefaultErrorListener
	content     string
	filename    string
	diagnostics *Diagnostics
}

func NewDiagnosticsErrorListener(content string, filename string, diags *Diagnostics) *DiagnosticsErrorListener {
	return &DiagnosticsErrorListener{
		content:     content,
		filename:    filename,
		diagnostics: diags,
	}
}

// SyntaxError implements antlr.ANTLRErrorListener
func (l *DiagnosticsErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Build a minimal TextLocation; try to get token span if available
	startCol := column
	stopCol := column + 1

	var startIdx, stopIdx int
	if tok, ok := offendingSymbol.(antlr.Token); ok {
		if tok != nil && tok.GetStart() >= 0 && tok.GetStop() >= tok.GetStart() {
			startCol = tok.GetColumn()

			// tok.GetStop() is inclusive index over content; convert to column span conservatively
			stopCol = startCol + len(tok.GetText())
		}

		startIdx = tok.GetStart()
		stopIdx = tok.GetStop()
		if stopIdx >= startIdx {
			stopIdx++
		}
	}

	loc := SourceLocation{
		Start:    Location{Line: line, Column: startCol, Index: startIdx},
		Stop:     Location{Line: line, Column: stopCol, Index: stopIdx},
		Filename: l.filename,
	}
	l.diagnostics.Add(Diagnostic{
		Message:  fmt.Sprintf("syntax error: %s", msg),
		Location: loc,
		Severity: Error,
	})

}

// helper: get text snippet between columns on a given 1-based line
func snippetAt(content string, line, startCol, stopCol int) string {
	if line < 1 {
		return ""
	}
	curLine := 1
	lineStart := 0
	for i, r := range content {
		if r == '\n' {
			if curLine == line {
				// i is index of '\n'; segment is content[lineStart:i]
				segment := content[lineStart:i]
				return sliceColumns(segment, startCol, stopCol)
			}
			curLine++
			lineStart = i + 1
		}
	}
	// handle last line (no trailing newline)
	if curLine == line {
		segment := content[lineStart:]
		return sliceColumns(segment, startCol, stopCol)
	}
	return ""
}

// sliceColumns attempts to slice a UTF-8 string by visual columns;
// for simplicity we approximate by bytes assuming ASCII-based input typical for code.
func sliceColumns(s string, startCol, stopCol int) string {
	if startCol < 0 {
		startCol = 0
	}
	if stopCol < startCol {
		stopCol = startCol
	}
	if startCol > len(s) {
		return ""
	}
	if stopCol > len(s) {
		stopCol = len(s)
	}
	return s[startCol:stopCol]
}
