package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ListSpec struct {
	ValueOptions []ValueSpec

	export func(fn ListNode, name string) *lib.FileTreeLike
	output func(fn ListNode) any
}

func NewListSpec(valueOptions ...ValueSpec) *ListSpec {
	return &ListSpec{ValueOptions: valueOptions}
}

func (l ListSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
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
		spec:   &l,
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
func (l ListSpec) SetFileExporter(exporter func(n ListNode, name string) *lib.FileTreeLike) ListSpec {
	out := &l
	out.export = exporter
	return *out
}
func (l ListSpec) SetOutputFn(outputFn func(n ListNode) any) ListSpec {
	out := &l
	out.output = outputFn
	return *out
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
