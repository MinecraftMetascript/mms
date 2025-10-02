package spec

import (
	"strings"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type StringSpec struct{}

func NewStringSpec() *StringSpec {
	return &StringSpec{}
}

func (s StringSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if val := valueCtx.String_(); val != nil {
		l := ast.TerminalLocation(val)
		return &StringNode{
			Value: strings.Trim(val.GetText(), "\""),
			BaseNode: ast.BaseNode{
				Location: &l,
			},
		}, nil
	}
	return nil, nil
}

type StringNode struct {
	ast.BaseNode
	Value string
}

func GetStringNodeValue(n ast.Node) *string {
	if n == nil {
		return nil
	}
	if n, ok := n.(*StringNode); ok {
		return &n.Value
	}
	return nil
}
