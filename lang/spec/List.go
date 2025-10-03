package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type ListSpec struct {
	ValueOptions []ValueSpec
}

func NewListSpec(valueOptions ...ValueSpec) *ListSpec {
	return &ListSpec{ValueOptions: valueOptions}
}

func (c ListSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
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
	}

	diags := make([]ast.Diagnostic, 0)

	for _, valCtx := range listCtx.AllValue() {
		found := false
		for _, spec := range c.ValueOptions {
			val, valDiags := spec.Match(valCtx)
			diags = append(diags, valDiags...)
			if val != nil {
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

type ListNode struct {
	ast.BaseSymbol
	Values []ast.Node
}
