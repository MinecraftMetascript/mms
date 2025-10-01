package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type ConditionSpec struct{}

func (c ConditionSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	//TODO implement me
	panic("implement me")
}

type ConditionNode struct{}

func (c ConditionNode) Children() []ast.Node {
	//TODO implement me
	panic("implement me")
}

func (c ConditionNode) GetLocation() *ast.SourceLocation {
	//TODO implement me
	panic("implement me")
}
