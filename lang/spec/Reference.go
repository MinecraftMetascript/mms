package spec

import (
	"fmt"
	"log"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/samber/lo"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ReferenceSpec struct {
	Kind             ast.SymbolKind
	defaultNamespace string
}

func (r *ReferenceSpec) UsageStr() string {
	if r.Kind == ast.SymbolNever {
		return "ref"
	} else {
		return fmt.Sprintf("ref(%s)", r.Kind)
	}
}

func NewReferenceSpec(kind ast.SymbolKind) *ReferenceSpec {
	return &ReferenceSpec{Kind: kind}
}
func (r *ReferenceSpec) SetDefaultNamespace(defaultNamespace string) *ReferenceSpec {
	r.defaultNamespace = defaultNamespace
	return r
}

func (r *ReferenceSpec) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	// Extract any prefix the user has already typed
	prefix := ExtractPrefixAtPosition(fileSource, position)

	start := protocol.Position{
		Line:      position.Line,
		Character: protocol.UInteger(int(position.Character) - len(prefix)),
	}

	out := make([]protocol.CompletionItem, 0)
	for ns, nsSymbols := range symbols {
		for n, s := range nsSymbols.AllDecls() {
			log.Printf("%T %s %s %s", s, s.GetKind(), ns, n)
			if s.GetKind() == r.Kind {
				out = append(out, protocol.CompletionItem{
					Label: fmt.Sprintf("[%s] %s:%s", s.GetKind(), ns, n),
					TextEdit: protocol.TextEdit{
						Range: protocol.Range{
							Start: start,
							End:   position,
						},
						NewText: fmt.Sprintf("%s:%s", ns, n),
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

func (r *ReferenceSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	refCtx := valueCtx.ResourceReference()
	if refCtx == nil {
		return nil, nil
	}

	parts := refCtx.AllIdentifier()
	refL := ast.RuleLocation(refCtx)
	if len(parts) == 1 {
		return &ReferenceNode{
			BaseSymbol: ast.BaseSymbol{
				Location: &refL,
				BaseNode: ast.BaseNode{},
			},
			Namespace: r.defaultNamespace,
			Name:      parts[0].GetText(),
			Kind:      r.Kind,
		}, nil

	} else if len(parts) == 2 {
		return &ReferenceNode{
			BaseSymbol: ast.BaseSymbol{
				Location: &refL,
				BaseNode: ast.BaseNode{},
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
	ast.BaseSymbol
	Namespace string
	Name      string
	Kind      ast.SymbolKind
}

func (r *ReferenceNode) Children() []ast.Node {
	return []ast.Node{}
}

func (r *ReferenceNode) String() string { return fmt.Sprintf("%s:%s", r.Namespace, r.Name) }

func GetReferenceNodeValue(node ast.Node, kind ast.SymbolKind) *string {
	if node == nil {
		return nil
	}
	if n, ok := node.(*ReferenceNode); ok {
		if n.Kind != kind {
			return nil
		}
		r := n.String()
		return &r
	}
	return nil
}
