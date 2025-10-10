package spec

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
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

func (r *ReferenceSpec) Complete(fileSource string, position protocol.Position, _ *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
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
				out = append(out, protocol.CompletionItem{
					Label:      fmt.Sprintf("[%s] %s:%s", s.GetKind(), ns, n),
					FilterText: &n,
					SortText:   &n,
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

	children := refCtx.GetChildren()
	if t, ok := children[0].(antlr.TerminalNode); ok {
		if t.GetText() == ":" {
			return nil, []ast.Diagnostic{
				{
					Location: ast.RuleLocation(valueCtx),
					Message:  "Incomplete reference",
					Severity: ast.Warning,
				},
			}
		}
	} else if t, ok = children[len(children)-1].(antlr.TerminalNode); ok {
		if t.GetText() == ":" {
			return nil, []ast.Diagnostic{
				{
					Location: ast.RuleLocation(valueCtx),
					Message:  "Incomplete reference",
					Severity: ast.Warning,
				},
			}
		}
	} else {
		log.Println("All is fine", refCtx.GetText())
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

func (r *ReferenceNode) Ref() string {
	return fmt.Sprintf("%s:%s", r.Namespace, r.Name)
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

func SetDefaultNamespaces(root ast.Node, namespace string) {
	if r, ok := root.(*ReferenceNode); ok {
		if r.Namespace == "" {
			r.Namespace = namespace
		}
	} else if t, ok := root.(*TagNode); ok {
		if t.Namespace == "" {
			t.Namespace = namespace
		}
	}
	if root != nil {
		for _, child := range root.Children() {
			if child != nil {
				SetDefaultNamespaces(child, namespace)
			}
		}
	}
}

func SetFilenames(root ast.Node, filename string) {
	if r, ok := root.(ast.Symbol); ok {
		r.SetFilename(filename)
	}

	if root != nil {
		for _, child := range root.Children() {
			if child != nil {
				SetFilenames(child, filename)
			}
		}
	}
}
