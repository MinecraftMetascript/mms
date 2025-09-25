package completions

import (
	"slices"
	"strings"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// ReferenceCompletions returns completion items for all symbols of the given kind
// visible from the provided scope. Items insert the fully-qualified reference (namespace:name)
// at the given insertPosition.
func ReferenceCompletions(scope *traversal.Scope, kind traversal.SymbolKind, insertPosition protocol.Position) []protocol.CompletionItem {
	items := make([]protocol.CompletionItem, 0)

	seen := make(map[string]struct{})
	for s := scope; s != nil; s = s.Parent() {
		for _, sym := range s.Symbols() {
			if sym == nil || sym.GetReference() == nil {
				continue
			}
			if sym.GetKind() != kind {
				continue
			}
			refStr := sym.GetReference().String()
			if _, ok := seen[refStr]; ok {
				continue
			}
			seen[refStr] = struct{}{}

			items = append(items, protocol.CompletionItem{
				Label: refStr,
				Kind:  lib.Ptr(protocol.CompletionItemKindReference),
				TextEdit: protocol.TextEdit{
					Range: protocol.Range{
						Start: insertPosition,
						End:   insertPosition,
					},
					NewText: refStr,
				},
			})
		}
	}

	return items
}

func FilterByReference(
	scope *traversal.Scope,
	kind traversal.SymbolKind,
	cursorPosition protocol.Position,
	items []protocol.CompletionItem,
	refCtx grammar.IResourceReferenceContext,
) []protocol.CompletionItem {
	if refCtx == nil {
		return items
	}
	refCons := traversal.ConstructRegistry.Construct(refCtx, "", scope)
	ref, ok := refCons.(*traversal.Reference)
	if !ok {
		return items
	}
	if _, ok := scope.Get(*ref); ok {
		// TODO: is namespacing handled properly here?
		return []protocol.CompletionItem{}

	}
	allCompletions := ReferenceCompletions(scope, kind, cursorPosition)
	filteredCompletions := slices.DeleteFunc(allCompletions, func(item protocol.CompletionItem) bool {
		nameMatch := !strings.HasPrefix(item.Label, ref.GetName())
		namespaceMatch := !strings.HasPrefix(item.Label, ref.GetNamespace())
		// if there is neither a name match nor a namespace match, it is not a valid completion, and should be filtered out
		return !nameMatch && !namespaceMatch
	})
	translatedCompletions := make([]protocol.CompletionItem, len(filteredCompletions))
	for i, c := range filteredCompletions {
		// TODO: This is brittle
		if edit, ok := c.TextEdit.(protocol.TextEdit); ok {
			prefix := strings.ReplaceAll(refCtx.GetText(), "<missing undefined>", "")
			prefixLen := len(prefix)
			edit.Range.Start.Character = protocol.UInteger(uint32(cursorPosition.Character) - uint32(prefixLen))
			translatedCompletions[i] = protocol.CompletionItem{
				Label:      c.Label,
				Kind:       c.Kind,
				Detail:     c.Detail,
				TextEdit:   edit,
				SortText:   c.SortText,
				FilterText: c.FilterText,
			}
		}
	}
	return translatedCompletions
}
