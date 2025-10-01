package spec

import (
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ReferenceSpec struct {
	Kind ast.SymbolKind
}

func NewReferenceSpec(kind ast.SymbolKind) *ReferenceSpec {
	return &ReferenceSpec{Kind: kind}
}
func (r ReferenceSpec) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	out := make([]protocol.CompletionItem, 0)
	log.Printf("Scope: %-s", symbols)
	log.Printf("Attempting to complete a reference! %s", r.Kind)
	for ns, nsSymbols := range symbols {
		log.Printf("Scanning %s", ns)
		for n, s := range nsSymbols.AllDecls() {
			log.Printf("  Checking %s", n)
			if s.GetKind() == r.Kind {
				log.Println("    Hit")
				out = append(out, protocol.CompletionItem{
					Label: fmt.Sprintf("%s:%s", ns, n),
					TextEdit: protocol.TextEdit{
						Range: protocol.Range{
							Start: position,
							End:   position,
						},
						NewText: fmt.Sprintf("%s:%s", ns, n),
					},
					Kind: &ReferenceKind,
				})
			} else {
				log.Printf("    Miss")
			}
		}
	}

	return out
}

func (r ReferenceSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	refCtx := valueCtx.ResourceReference()
	if refCtx == nil {
		return nil, nil
	}

	parts := refCtx.AllIdentifier()
	refL := ast.RuleLocation(refCtx)
	if len(parts) == 1 {
		return ReferenceNode{
			location:  refL,
			Namespace: "",
			Name:      parts[0].GetText(),
		}, nil

	} else if len(parts) == 2 {
		return ReferenceNode{
			location:  refL,
			Namespace: parts[0].GetText(),
			Name:      parts[1].GetText(),
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
	location  ast.SourceLocation
	Namespace string
	Name      string
}

func (r ReferenceNode) Children() []ast.Node {
	return []ast.Node{}
}

func (r ReferenceNode) GetLocation() *ast.SourceLocation {
	return &r.location
}
