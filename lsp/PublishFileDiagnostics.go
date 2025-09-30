package lsp

import (
	"github.com/minecraftmetascript/mms/project"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (ls *LanguageServer) PublishDiagnostics(context *glsp.Context, file *project.File) {

	errSeverity := protocol.DiagnosticSeverityError
	diags := make([]protocol.Diagnostic, 0)

	source := "MMS"
	for _, diag := range file.Diagnostics() {
		diags = append(diags, protocol.Diagnostic{
			Range:    diag.Location.ToLspRange(),
			Severity: &errSeverity,
			Source:   &source,
			Message:  diag.Message,
		})
	}

	ls.log.Debugf("Publishing %d diagnostics for %s", len(diags), file.Path())

	go context.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         file.Path(),
		Diagnostics: diags,
	})
}
