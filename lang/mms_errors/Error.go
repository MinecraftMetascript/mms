package mms_errors

import (
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lib"
)

type ErrorLevel string

const (
	ErrorLevelWarning ErrorLevel = "warning"
	ErrorLevelError   ErrorLevel = "error"
)

type ErrorType string

const (
	ErrorTypeUnknown         ErrorType = "unknown"
	ErrorTypeSyntaxError     ErrorType = "syntax_error"
	ErrorTypeDuplicateSymbol ErrorType = "duplicate_symbol"
)

type MMSError struct {
	Message  string
	Level    ErrorLevel
	Type     ErrorType
	Location lib.TextLocation
	// TODO: End Column, needed for annotations in LSP
	Filename string
}

func (e MMSError) String() string {
	return fmt.Sprintf("%s:%d:%d %s", e.Filename, e.Location.Start.Line, e.Location.Start.Column, e.Message)
}
func (e MMSError) Error() string {
	return e.String()
}

func DuplicateSymbolError(filename string, location lib.TextLocation, msg string) MMSError {
	return MMSError{
		Message:  msg,
		Level:    ErrorLevelError,
		Location: location,
		Filename: filename,
		Type:     ErrorTypeDuplicateSymbol,
	}
}

func SyntaxWarning(filename string, location lib.TextLocation, msg string) MMSError {
	return MMSError{
		Message:  msg,
		Level:    ErrorLevelWarning,
		Location: location,
		Filename: filename,
		Type:     ErrorTypeSyntaxError,
	}
}

func SyntaxError(filename string, location lib.TextLocation, msg string) MMSError {
	return MMSError{
		Message:  msg,
		Level:    ErrorLevelError,
		Location: location,
		Filename: filename,
		Type:     ErrorTypeSyntaxError,
	}
}

type ErrorListener struct {
	*antlr.ConsoleErrorListener
	filename string

	emitError func(MMSError)
}

func NewErrorListener(filename string, emitError func(MMSError)) *ErrorListener {
	return &ErrorListener{
		filename:  filename,
		emitError: emitError,
	}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, symbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	switch e.(type) {
	case *antlr.NoViableAltException:
		token := symbol.(antlr.Token)

		l.emitError(SyntaxError(
			l.filename,
			lib.TextLocation{
				Start: lib.Location{
					Column: token.GetColumn() + 1,
					Line:   token.GetLine(),
				},
			},
			"Syntax Error", // TODO: Provide more specifics by looking up token or rule type?
		))
	default:
		l.emitError(SyntaxError(
			l.filename,
			lib.TextLocation{
				Start: lib.Location{
					Column: column + 1,
					Line:   line,
				},
			},
			msg,
		))
	}
}

func (l *ErrorListener) ReportAmbiguity(parser antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	alternativeRules := make([]string, 0)

	for _, state := range configs.GetStates().Values() {
		alternativeRules = append(alternativeRules, parser.GetRuleNames()[state.GetRuleIndex()])
	}

	log.Println(fmt.Sprintf("[WARN]: Ambiguity detected %-s", alternativeRules))

}

func (l *ErrorListener) ReportAttemptingFullContext(parser antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact *antlr.BitSet, configs *antlr.ATNConfigSet) {

	alternativeRules := make([]string, 0)

	for _, state := range configs.GetStates().Values() {
		alternativeRules = append(alternativeRules, parser.GetRuleNames()[state.GetRuleIndex()])
	}

	log.Println(fmt.Sprintf("[WARN]: Attempting Full Context %-s", alternativeRules))
}

func (l *ErrorListener) ReportContextSensitivity(parser antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	log.Println(">>>", "ReportContextSensitivity")
}
