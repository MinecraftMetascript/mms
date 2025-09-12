package lsp

import (
	"slices"

	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type document struct {
	uri     protocol.URI
	file    lang.File
	project *lang.Project
}

func newDocument(
	content string,
	uri protocol.URI,
	project *lang.Project,
) (*document, error) {
	file := project.AddFile(uri, content)
	err := file.Parse()
	if err != nil {
		return nil, err
	}

	return &document{
		uri:     uri,
		project: project,
		file:    *file,
	}, err
}

func (d *document) Update() {

}

func (d *document) PublishDiagnostics(context *glsp.Context) {
	errSeverity := protocol.DiagnosticSeverityError
	diags := make([]protocol.Diagnostic, 0)
	for _, diag := range d.file.Diagnostics {
		diags = append(diags, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      protocol.UInteger(diag.Where.Start.Line) - 1,
					Character: protocol.UInteger(diag.Where.Start.Col),
				},
				End: protocol.Position{
					Line:      protocol.UInteger(diag.Where.Stop.Line) - 1,
					Character: protocol.UInteger(diag.Where.Stop.Col),
				},
			},
			Severity: &errSeverity,
			Source:   &diag.Source,
			Message:  diag.Message,
		})
	}

	go context.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         d.uri,
		Diagnostics: diags,
	})
}

func (d *document) ApplyChanges(changes []interface{}) error {
	content := d.file.Content
	for _, change := range changes {
		switch c := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			startIndex, endIndex := c.Range.IndexesIn(content)
			content = content[:startIndex] + c.Text + content[endIndex:]
		case protocol.TextDocumentContentChangeEventWhole:
			content = c.Text
		}
	}
	d.file = *d.project.AddFile(d.uri, content)

	return d.file.Parse()
}

func (d *document) HelpAtPosition(position protocol.Position, filename string) *traversal.Help {
	for _, rule := range slices.Backward(d.file.GetRulesAtPosition(position.IndexIn(d.file.Content))) {
		if helpContent := traversal.GetHelp(rule, filename); helpContent != nil {
			return helpContent
		}
	}
	return nil
}
