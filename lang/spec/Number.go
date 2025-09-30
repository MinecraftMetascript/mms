package spec

import (
	"strconv"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type NumberSpec struct {
	floating bool
}

func NewNumberSpec(floating bool) *NumberSpec {
	return &NumberSpec{
		floating: floating,
	}
}

func (n NumberSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if numCtx := valueCtx.Number(); numCtx != nil {
		numTxt := numCtx.GetText()
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
			l := ast.RuleLocation(numCtx)
			return &NumberNode{
				Value: float64(val),
				BaseNode: ast.BaseNode{
					Location: &l,
				},
			}, nil
		}
	}
	return nil, nil
}

type NumberNode struct {
	ast.BaseNode
	Value float64
}

func (n NumberNode) GetLocation() *ast.SourceLocation {
	return n.Location
}

func (n NumberNode) Children() []ast.Node {
	return []ast.Node{}
}
