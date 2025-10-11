package spec

import (
	"strconv"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type NumberSpec struct {
	floating bool
	kind     ast.SymbolKind
}

func (n NumberSpec) UsageStr() string {
	if n.floating {
		return "float"
	} else {
		return "int"
	}
}

func NewNumberSpec(floating bool) *NumberSpec {
	return &NumberSpec{
		floating: floating,
	}
}

func (n NumberSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if numCtx := valueCtx.Number(); numCtx != nil {
		numTxt := numCtx.GetText()
		numLocation := ast.RuleLocation(numCtx)
		if n.floating {
			val, err := strconv.ParseFloat(numTxt, 64)
			if err != nil {
				return nil, []ast.Diagnostic{{
					Location: ast.RuleLocation(valueCtx),
					Message:  "Invalid Float Value",
					Severity: ast.Warning,
				}}
			}
			return &NumberNode{
				Value: val,
				BaseSymbol: ast.BaseSymbol{
					Location: &numLocation,
					BaseNode: ast.BaseNode{},
				},
				spec: &n,
			}, nil
		} else {
			val, err := strconv.Atoi(numTxt)
			if err != nil {
				return nil, []ast.Diagnostic{{
					Location: ast.RuleLocation(valueCtx),
					Message:  "Invalid Int Value",
					Severity: ast.Warning,
				}}
			}
			return &NumberNode{
				Value: float64(val),
				BaseSymbol: ast.BaseSymbol{
					Location: &numLocation,
					BaseNode: ast.BaseNode{},
				},
				spec: &n,
			}, nil
		}
	}
	return nil, nil
}

type NumberNode struct {
	ast.BaseSymbol
	Value float64
	spec  *NumberSpec
}

func GetNumberNodeValue(n ast.Node) *float64 {
	if n == nil {
		return nil
	}
	if n, ok := n.(*NumberNode); ok {
		return &n.Value
	}
	return nil
}

func (n *NumberNode) ToSerializable() any {
	return n.Value
}

func (n *NumberNode) GetKind() ast.SymbolKind {
	return n.spec.kind
}

func (n *NumberSpec) SetKind(kind ast.SymbolKind) *NumberSpec {
	n.kind = kind
	return n
}
