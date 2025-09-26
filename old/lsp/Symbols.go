package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Returns: []DocumentSymbol | []SymbolInformation | nil

func (ls *LanguageServer) DocumentSymbols(_ *glsp.Context, _ *protocol.DocumentSymbolParams) (any, error) {
	out := make([]protocol.DocumentSymbol, 0)

	for name, s := range ls.project.GlobalScope.Symbols() {
		out = append(out, protocol.DocumentSymbol{
			Name: name,
			Kind: protocol.SymbolKindObject,
			Range: protocol.Range{
				Start: s.GetNameLocation().Start.ToLspPosition(),
				End:   s.GetContentLocation().Stop.ToLspPosition(),
			},
			SelectionRange: protocol.Range{
				Start: s.GetNameLocation().Start.ToLspPosition(),
				End:   s.GetNameLocation().Stop.ToLspPosition(),
			},
		})
		ls.log.Infof("Symbol %s @ %d:%d - %d:%d", name, s.GetNameLocation().Start.Line, s.GetNameLocation().Start.Col, s.GetNameLocation().Stop.Line, s.GetNameLocation().Stop.Col)
	}

	return out, nil
}
