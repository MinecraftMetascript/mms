package spec

import (
	"strings"
	"unicode"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var SnippetFormat = protocol.InsertTextFormatSnippet
var SnippetKind = protocol.CompletionItemKindSnippet
var MethodKind = protocol.CompletionItemKindMethod
var ReferenceKind = protocol.CompletionItemKindReference

// ExtractPrefixAtPosition extracts the identifier-like prefix before the cursor position.
// It walks backwards from the position until it hits a non-identifier character.
// Returns the extracted prefix (lowercase for case-insensitive matching).
func ExtractPrefixAtPosition(fileSource string, position protocol.Position) string {
	idx := position.IndexIn(fileSource)
	if idx <= 0 || idx > len(fileSource) {
		return ""
	}

	// Walk backwards to find the start of the identifier
	start := idx
	for start > 0 {
		ch := rune(fileSource[start-1])
		// Identifier characters: letters, numbers, underscore, colon (for namespaced refs)
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' && ch != ':' {
			break
		}
		start--
	}

	prefix := fileSource[start:idx]
	return strings.ToLower(prefix)
}

// FilterCompletionsByPrefix filters completion items based on a prefix.
// It performs case-insensitive prefix matching on the Label field.
func FilterCompletionsByPrefix(items []protocol.CompletionItem, prefix string) []protocol.CompletionItem {
	if prefix == "" {
		return items
	}

	filtered := make([]protocol.CompletionItem, 0)
	lowerPrefix := strings.ToLower(prefix)

	for _, item := range items {
		labelLower := strings.ToLower(item.Label)
		if strings.HasPrefix(labelLower, lowerPrefix) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

type ValueSpec interface {
	Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic)
}

type ValueSpecList struct {
	specs []ValueSpec
}

func (bsl *ValueSpecList) Add(bs ValueSpec) {
	bsl.specs = append(bsl.specs, bs)
}

func (bsl *ValueSpecList) Match(ctx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if ctx == nil {
		return nil, nil
	}

	allDiags := make([]ast.Diagnostic, 0)

	for _, bs := range bsl.specs {
		val, diags := bs.Match(ctx)

		if !lib.IsNilInterface(val) {
			// Found a match - return it with any diagnostics
			return val, diags
		}

		// Collect diagnostics from failed matches
		if diags != nil && len(diags) > 0 {
			allDiags = append(allDiags, diags...)
		}
	}

	// No spec matched - return collected diagnostics if any
	if len(allDiags) > 0 {
		return nil, allDiags
	}

	return nil, nil
}

func NewValueSpecList(
	specs ...ValueSpec,
) ValueSpecList {
	out := ValueSpecList{
		specs: specs,
	}
	return out
}
