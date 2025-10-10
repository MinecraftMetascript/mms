package spec

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type TagSpec struct {
	Kind             ast.TagKind
	defaultNamespace string
}

func (r *TagSpec) UsageStr() string {
	return fmt.Sprintf("tag(%s)", r.Kind)
}

func NewTagSpec(kind ast.TagKind) *TagSpec {
	return &TagSpec{Kind: kind}
}
func (r *TagSpec) SetDefaultNamespace(defaultNamespace string) *TagSpec {
	r.defaultNamespace = defaultNamespace
	return r
}

// TODO: We don't have any way to store tags right now
//func (r *TagSpec) Complete(fileSource string, position protocol.Position, _ *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
//	// Extract any prefix the user has already typed
//	prefix := ExtractPrefixAtPosition(fileSource, position)
//
//	start := protocol.Position{
//		Line:      position.Line,
//		Character: protocol.UInteger(int(position.Character) - len(prefix)),
//	}
//
//	out := make([]protocol.CompletionItem, 0)
//	for ns, nsSymbols := range symbols {
//		for n, s := range nsSymbols.AllDecls() {
//			if s.GetKind() == r.Kind {
//				out = append(out, protocol.CompletionItem{
//					Label:      fmt.Sprintf("[%s] %s:%s", s.GetKind(), ns, n),
//					FilterText: &n,
//					SortText:   &n,
//					TextEdit: protocol.TextEdit{
//						Range: protocol.Range{
//							Start: start,
//							End:   position,
//						},
//						NewText: fmt.Sprintf("%s:%s", ns, n),
//					},
//					Kind: &ReferenceKind,
//				})
//			}
//		}
//	}
//
//	// Filter by prefix if user has typed something
//	return lo.Filter(out, func(item protocol.CompletionItem, index int) bool {
//		parts := strings.Split(strings.ToLower(item.Label), ":")
//		return strings.HasPrefix(parts[0], strings.ToLower(prefix)) || strings.HasPrefix(parts[1], strings.ToLower(prefix))
//	})
//}

func (r *TagSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	tagCtx := valueCtx.ResourceTag()
	if tagCtx == nil {
		return nil, nil
	}

	refCtx := tagCtx.ResourceReference()
	if refCtx == nil {
		return nil, []ast.Diagnostic{
			{
				Location: ast.RuleLocation(valueCtx),
				Message:  "Missing tag reference",
				Severity: ast.Warning,
			},
		}
	}

	txt := strings.Trim(refCtx.GetText(), "\n \t")
	if strings.HasPrefix(txt, ":") || strings.HasSuffix(txt, ":") {
		return nil, []ast.Diagnostic{
			{
				Location: ast.RuleLocation(valueCtx),
				Message:  "Invalid tag reference",
				Severity: ast.Warning,
			},
		}
	}

	parts := refCtx.AllIdentifier()
	refL := ast.RuleLocation(refCtx)
	if len(parts) == 1 {
		return &TagNode{
			BaseSymbol: ast.BaseSymbol{
				Location: &refL,
				BaseNode: ast.BaseNode{},
			},
			Namespace: r.defaultNamespace,
			Name:      parts[0].GetText(),
			Kind:      r.Kind,
		}, nil

	} else if len(parts) == 2 {
		return &TagNode{
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
				Message:  "Invalid Tag",
				Severity: ast.Warning,
			},
		}
	}
}

type TagNode struct {
	ast.BaseSymbol
	Namespace string
	Name      string
	Kind      ast.TagKind
}

func (r *TagNode) Ref() string {
	return fmt.Sprintf("%s:%s", r.Namespace, r.Name)
}

func (r *TagNode) Children() []ast.Node {
	return []ast.Node{}
}

func (r *TagNode) String() string { return fmt.Sprintf("%s:%s", r.Namespace, r.Name) }

func GetTagNodeValue(node ast.Node, kind ast.TagKind) *string {
	if node == nil {
		return nil
	}
	if n, ok := node.(*TagNode); ok {
		if n.Kind != kind {
			return nil
		}
		r := n.String()
		return &r
	}
	return nil
}
