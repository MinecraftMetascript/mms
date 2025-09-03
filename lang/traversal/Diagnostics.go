package traversal

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Severity indicates the diagnostic level
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

// Diagnostic describes an issue found during lexing/parsing/analysis
// Start/Stop positions are inclusive of the underlying text segment
// via TextLocation
// Source can be "lexer", "parser", or "semantic"
type Diagnostic struct {
	Message  string       `json:"message"`
	Where    TextLocation `json:"where"`
	Severity Severity     `json:"severity"`
	Source   string       `json:"source"`
	File     string       `json:"file"`
}

func (d Diagnostic) String() string {
	return fmt.Sprintf("%s (%s): %s", d.Severity, d.Where.String(), d.Message)
}

// DiagnosticsErrorListener collects syntax errors from ANTLR and stores them
// into the provided slice. It uses the file content to populate the
// TextLocation text and accurate columns.
type DiagnosticsErrorListener struct {
	antlr.DefaultErrorListener
	content     string
	filename    string
	diagnostics *[]Diagnostic
}

func NewDiagnosticsErrorListener(content string, filename string, diags *[]Diagnostic) *DiagnosticsErrorListener {
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

	loc := TextLocation{
		Start:    Location{Line: line, Col: startCol},
		StartIdx: startIdx,
		Stop:     Location{Line: line, Col: stopCol},
		StopIdx:  stopIdx,
		Text:     snippetAt(l.content, line, startCol, stopCol),
		Filename: l.filename,
	}

	*l.diagnostics = append(*l.diagnostics, Diagnostic{
		Message:  fmt.Sprintf("syntax error: %s", msg),
		Where:    loc,
		Severity: SeverityError,
		Source:   "parser",
		File:     l.filename,
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
