package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Returns: []DocumentSymbol | []SymbolInformation | nil

func (ls *LanguageServer) DocumentSymbols(*glsp.Context, *protocol.DocumentSymbolParams) (any, error) {
	out := make([]protocol.DocumentSymbol, 0)

	for name, s := range ls.project.GlobalScope.Symbols() {
		out = append(out, protocol.DocumentSymbol{
			Name: name,
			Kind: protocol.SymbolKindVariable,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      protocol.UInteger(s.GetNameLocation().Start.Line),
					Character: protocol.UInteger(s.GetNameLocation().Start.Col),
				},
				End: protocol.Position{
					Line:      protocol.UInteger(s.GetNameLocation().Stop.Line),
					Character: protocol.UInteger(s.GetNameLocation().Stop.Col),
				},
			},
		})
	}

	return out, nil
}
