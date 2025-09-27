package ast

type DiagnosticSeverity string

const (
	Error   DiagnosticSeverity = "error"
	Warning DiagnosticSeverity = "warning"
	Info    DiagnosticSeverity = "info"
)

type Diagnostic struct {
	Location SourceLocation
	Message  string
	Severity DiagnosticSeverity
}
