package spec

import (
	"strings"
	"unicode"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	"github.com/samber/lo"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var SnippetFormat = protocol.InsertTextFormatSnippet
var MethodKind = protocol.CompletionItemKindMethod
var ReferenceKind = protocol.CompletionItemKindReference
var StructKind = protocol.CompletionItemKindStruct
var EnumKind = protocol.CompletionItemKindEnum

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

type ValueSpec interface {
	Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic)
	UsageStr() string
}

type ValueSpecList struct {
	specs []ValueSpec
	label string
}

func (vsl *ValueSpecList) SetLabel(label string) *ValueSpecList {
	vsl.label = label
	return vsl
}

func (vsl *ValueSpecList) UsageStr() string {
	if vsl.label != "" {
		return vsl.label
	}
	return strings.Join(
		lo.Map(lo.Filter(vsl.specs, func(item ValueSpec, index int) bool {
			return item != vsl
		}), func(item ValueSpec, index int) string { return item.UsageStr() }),
		" | ",
	)
}

func (vsl *ValueSpecList) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)
	for _, spec := range vsl.specs {
		if c, ok := spec.(ast.CompletableNode); ok {
			completions = append(completions, c.Complete(fileSource, position, triggerChar, symbols)...)
		}
	}
	return completions
}

func (vsl *ValueSpecList) Add(bs ...ValueSpec) {
	vsl.specs = append(vsl.specs, bs...)
}

func (vsl *ValueSpecList) Match(ctx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if ctx == nil {
		return nil, nil
	}

	allDiags := make([]ast.Diagnostic, 0)

	for _, bs := range vsl.specs {
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

func (vsl *ValueSpecList) All() []ValueSpec {
	return vsl.specs
}

func NewValueSpecList(
	specs ...ValueSpec,
) *ValueSpecList {
	out := ValueSpecList{
		specs: specs,
	}
	return &out
}
