package completions

import (
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func BuilderFnCompletion(label, snippet string, insertPosition protocol.Position) protocol.CompletionItem {
	return protocol.CompletionItem{
		Label: label,
		Kind:  lib.Ptr(protocol.CompletionItemKindMethod),
		TextEdit: protocol.TextEdit{
			Range: protocol.Range{
				Start: lib.AddCols(insertPosition, 2),
				End:   lib.AddCols(insertPosition, 2),
			},
			NewText: snippet,
		},
		InsertTextFormat: lib.Ptr(protocol.InsertTextFormatSnippet),
	}
}
