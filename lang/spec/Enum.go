package spec

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/samber/lo"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type EnumSpec struct {
	Options []string
	Help    string
	Label   string
}

func (e EnumSpec) UsageStr() string {
	if e.Label != "" {
		return e.Label
	}
	return strings.Join(e.Options, " | ")
}

func (e EnumSpec) Complete(
	fileSource string,
	position protocol.Position,
	_ *string,
	_ map[string]*ast.Namespace,
) []protocol.CompletionItem {
	prefix := ExtractPrefixAtPosition(fileSource, position)

	start := protocol.Position{
		Line:      position.Line,
		Character: protocol.UInteger(int(position.Character) - len(prefix)),
	}
	// Build completion items for all enum options
	items := make([]protocol.CompletionItem, 0)
	for _, option := range e.Options {
		detail := e.Help
		items = append(items, protocol.CompletionItem{
			Label:      option,
			FilterText: &option,
			SortText:   &option,
			Kind:       &EnumKind,
			Detail:     &detail,
			TextEdit: protocol.TextEdit{
				Range: protocol.Range{
					Start: start,
					End:   position,
				},
				NewText: option,
			},
		})
	}
	return items
}

func NewEnumSpec(options ...string) EnumSpec {
	return EnumSpec{Options: options}
}

func (e EnumSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	// Enums will be parsed in as resource references
	if valueCtx.ResourceReference() == nil {
		return nil, nil
	}
	rr := valueCtx.ResourceReference()

	if len(rr.AllIdentifier()) == 0 {
		// Not a match
		return nil, nil
	}
	value := rr.GetText()
	l := ast.RuleLocation(valueCtx)

	out := &EnumNode{
		spec:  e,
		Value: value,
		BaseSymbol: ast.BaseSymbol{
			BaseNode: ast.BaseNode{
				Location: &l,
			},
		},
	}

	valueValid := false
	for _, option := range e.Options {
		if option == value {
			valueValid = true
			break
		}
	}
	if !valueValid {
		// Build a helpful error message with valid options
		msg := fmt.Sprintf("Invalid enum value '%s'. Expected one of: %v", value, e.Options)
		return nil, []ast.Diagnostic{
			{
				Location: ast.RuleLocation(valueCtx),
				Message:  msg,
				Severity: ast.Error,
			},
		}
	}
	return out, nil
}

func (e EnumSpec) SetHelp(help string) EnumSpec {
	out := &e
	out.Help = help
	return *out
}

func (e EnumSpec) SetLabel(label string) EnumSpec {
	out := &e
	out.Label = label
	return *out
}

type EnumNode struct {
	ast.BaseSymbol
	Value string
	spec  EnumSpec
}

func (e EnumNode) GetHelp() string {
	return e.spec.Help
}

func (e EnumNode) Complete(_ string, position protocol.Position, _ *string, _ map[string]*ast.Namespace) []protocol.CompletionItem {
	if lo.IndexOf(e.spec.Options, e.Value) != -1 {
		// The user has already typed a valid value for this enum
		return make([]protocol.CompletionItem, 0)
	}

	// Build completion items for all enum options
	items := make([]protocol.CompletionItem, 0)
	for _, option := range e.spec.Options {
		detail := e.spec.Help
		items = append(items, protocol.CompletionItem{
			Label:      option,
			FilterText: &option,
			SortText:   &option,
			Kind:       &EnumKind,
			Detail:     &detail,
			TextEdit: protocol.TextEdit{
				Range: protocol.Range{
					Start: position,
					End:   position,
				},
				NewText: option,
			},
		})
	}

	return items
}

func GetEnumNodeValue(n ast.Node) *string {
	if n == nil {
		return nil
	}
	if n, ok := n.(*EnumNode); ok {
		return &n.Value
	}
	return nil
}
