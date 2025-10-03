package spec

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/samber/lo"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ReferenceSpec struct {
	Kind ast.SymbolKind
}

func NewReferenceSpec(kind ast.SymbolKind) *ReferenceSpec {
	return &ReferenceSpec{Kind: kind}
}
func (r ReferenceSpec) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	// Extract any prefix the user has already typed
	prefix := ExtractPrefixAtPosition(fileSource, position)

	start := protocol.Position{
		Line:      position.Line,
		Character: protocol.UInteger(int(position.Character) - len(prefix)),
	}

	out := make([]protocol.CompletionItem, 0)
	for ns, nsSymbols := range symbols {
		for n, s := range nsSymbols.AllDecls() {
			if s.GetKind() == r.Kind {
				label := fmt.Sprintf("%s:%s", ns, n)
				out = append(out, protocol.CompletionItem{
					Label: label,
					TextEdit: protocol.TextEdit{
						Range: protocol.Range{
							Start: start,
							End:   position,
						},
						NewText: label,
					},
					Kind: &ReferenceKind,
				})
			}
		}
	}

	// Filter by prefix if user has typed something
	return lo.Filter(out, func(item protocol.CompletionItem, index int) bool {
		parts := strings.Split(strings.ToLower(item.Label), ":")
		return strings.HasPrefix(parts[0], strings.ToLower(prefix)) || strings.HasPrefix(parts[1], strings.ToLower(prefix))
	})
}

func (r ReferenceSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	refCtx := valueCtx.ResourceReference()
	if refCtx == nil {
		return nil, nil
	}

	parts := refCtx.AllIdentifier()
	refL := ast.RuleLocation(refCtx)
	if len(parts) == 1 {
		return &ReferenceNode{
			BaseNode: ast.BaseNode{
				Location: &refL,
			},
			Namespace: "",
			Name:      parts[0].GetText(),
			Kind:      r.Kind,
		}, nil

	} else if len(parts) == 2 {
		return &ReferenceNode{
			BaseNode: ast.BaseNode{
				Location: &refL,
			},
			Namespace: parts[0].GetText(),
			Name:      parts[1].GetText(),
			Kind:      r.Kind,
		}, nil

	} else {
		// broken
		return nil, []ast.Diagnostic{
			{
				Location: ast.RuleLocation(valueCtx),
				Message:  "Invalid Reference",
				Severity: ast.Warning,
			},
		}
	}
}

type ReferenceNode struct {
	ast.BaseNode
	Namespace string
	Name      string
	Kind      ast.SymbolKind
}

func (r *ReferenceNode) Children() []ast.Node {
	return []ast.Node{}
}

func (r *ReferenceNode) String() string { return fmt.Sprintf("%s:%s", r.Namespace, r.Name) }

func GetReferenceNodeValue(n ast.Node) *string {
	if n == nil {
		return nil
	}
	if n, ok := n.(*ReferenceNode); ok {
		r := n.String()
		return &r
	}
	return nil
}
