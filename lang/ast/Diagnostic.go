package ast

type DiagnosticSeverity string

const (
	Error   DiagnosticSeverity = "error"
	Warning DiagnosticSeverity = "warning"
	Info    DiagnosticSeverity = "info"
)

/*

type Diagnostic struct {
	Message  string       `json:"message"`
	Where    TextLocation `json:"where"`
	Severity Severity     `json:"severity"`
	Source   string       `json:"source"`
	File     string       `json:"file"`
}

*/

type Diagnostic struct {
	Location SourceLocation     `json:"location"`
	Message  string             `json:"message"`
	Severity DiagnosticSeverity `json:"severity"`
}

type Diagnostics struct {
	diagnostics []Diagnostic
}

func (d *Diagnostics) Add(diagnostic ...Diagnostic) {
	d.diagnostics = append(d.diagnostics, diagnostic...)
}

func (d *Diagnostics) Get() []Diagnostic {
	seen := make(map[Diagnostic]bool) // A map to store seen elements
	result := make([]Diagnostic, 0)   // The array to store unique elements

	for _, v := range d.diagnostics {
		if _, ok := seen[v]; !ok { // If element not in map (not seen yet)
			seen[v] = true             // Mark as seen
			result = append(result, v) // Add to result
		}
	}
	return result
}

func NewDiagnostics() *Diagnostics {
	return &Diagnostics{
		diagnostics: make([]Diagnostic, 0),
	}
}
