package lsp

import (
	"log"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (ctx *LanguageServerCtx) completionItemResolve(context *glsp.Context, item *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	log.Println("Resolve Completion Item", item.Kind, item.InsertText)
	return item, nil
}

// Returns: []CompletionItem | CompletionList | nil
func (ctx *LanguageServerCtx) textDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	file := ctx.project.GetFile(params.TextDocument.URI[7:])

	list := protocol.CompletionList{}

	structItemKind := protocol.CompletionItemKindStruct

	if file != nil {
		// File contextual completions
		if file.GetNamespace() == "" {
			namespaceInsert := "namespace ;"
			list.Items = append(list.Items, protocol.CompletionItem{
				Label:      "Namespace",
				Kind:       &structItemKind,
				InsertText: &namespaceInsert,
			})
		}

		log.Println(params.Position.IndexIn(file.GetRawContent()))

	}

	surfaceInsert := "surface {}"
	list.Items = append(list.Items, protocol.CompletionItem{
		Label:      "Surface Rules",
		Kind:       &structItemKind,
		InsertText: &surfaceInsert,
	})

	return list, nil
}
