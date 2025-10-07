package spec

import (
	"fmt"
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ListSpec struct {
	ValueOptions []ValueSpec
	Kind         ast.SymbolKind
	Help         string

	export func(fn ListNode, name string) *lib.FileTreeLike
	output func(fn ListNode) any
}

func (l *ListSpec) AddValueOption(spec ...ValueSpec) *ListSpec {
	l.ValueOptions = slices.Concat(l.ValueOptions, spec)
	return l
}

func (l *ListSpec) SetKind(kind ast.SymbolKind) *ListSpec {
	l.Kind = kind
	return l
}
func (l *ListSpec) SetHelp(help string) *ListSpec {
	l.Help = help
	return l
}

func (l *ListSpec) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	return []protocol.CompletionItem{
		{
			Label:            fmt.Sprintf("[%s] List", l.Kind),
			Kind:             &MethodKind,
			Detail:           &l.Help,
			InsertTextFormat: &SnippetFormat,
			TextEdit: protocol.TextEdit{
				Range: protocol.Range{
					Start: position,
					End:   position,
				},
				NewText: "[ ${1} ]",
			},
		},
	}
}

func NewListSpec(valueOptions ...ValueSpec) *ListSpec {
	return &ListSpec{ValueOptions: valueOptions}
}

func (l *ListSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	listCtx := valueCtx.List()
	if listCtx == nil {
		return nil, nil
	}

	loc := ast.RuleLocation(listCtx)
	out := &ListNode{
		BaseSymbol: ast.BaseSymbol{
			Location: &loc,
		},
		Values: make([]ast.Node, 0),
		spec:   l,
	}

	diags := make([]ast.Diagnostic, 0)

	for _, valCtx := range listCtx.AllValue() {
		found := false
		for _, spec := range l.ValueOptions {
			val, valDiags := spec.Match(valCtx)
			diags = append(diags, valDiags...)
			if !lib.IsNilInterface(val) {
				out.Values = append(out.Values, val)
				found = true
				break
			}
		}
		if !found {
			diags = append(diags, ast.Diagnostic{
				Location: ast.RuleLocation(valCtx),
				Message:  "Invalid value",
				Severity: ast.Error,
			})
		}
	}

	return out, diags
}
func (l *ListSpec) SetFileExporter(exporter func(n ListNode, name string) *lib.FileTreeLike) *ListSpec {
	l.export = exporter
	return l
}
func (l *ListSpec) SetOutputFn(outputFn func(n ListNode) any) *ListSpec {
	l.output = outputFn
	return l
}

type ListNode struct {
	ast.BaseSymbol
	Values []ast.Node
	spec   *ListSpec
}

func (l ListNode) ToFileTreeLike(name string) *lib.FileTreeLike {
	if l.spec.export == nil {
		return nil
	}
	return l.spec.export(l, name)
}

func (l ListNode) ToSerializable() any {
	if l.spec.output == nil {
		return nil
	}
	return l.spec.output(l)
}

func (l ListNode) Children() []ast.Node {
	return l.Values
}

func (l ListNode) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	// Check if cursor is within an existing value
	for _, v := range l.Values {
		if !lib.IsNilInterface(v) && v.GetLocation() != nil && v.GetLocation().ContainsPosition(position) {
			// If the value implements CompletableNode, delegate to it
			if completable, ok := v.(ast.CompletableNode); ok {
				return completable.Complete(fileSource, position, triggerChar, symbols)
			}
			return make([]protocol.CompletionItem, 0)
		}
	}

	// Extract any prefix the user has already typed
	prefix := ExtractPrefixAtPosition(fileSource, position)

	// Provide completions for available value options
	out := make([]protocol.CompletionItem, 0)
	for _, valueSpec := range l.spec.ValueOptions {
		if c, ok := valueSpec.(ast.CompletableNode); ok {
			out = append(out, c.Complete(fileSource, position, triggerChar, symbols)...)
		}
	}
	// Filter by prefix if user has typed something
	return FilterCompletionsByPrefix(out, prefix)
}

func (l ListNode) GetKind() ast.SymbolKind {
	return l.spec.Kind
}
