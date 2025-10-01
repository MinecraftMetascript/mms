package spec

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type SequenceSpec struct{}

func (c SequenceSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	//TODO implement me
	panic("implement me")
}

type SequenceNode struct{}

func (c SequenceNode) Children() []ast.Node {
	//TODO implement me
	panic("implement me")
}

func (c SequenceNode) GetLocation() *ast.SourceLocation {
	//TODO implement me
	panic("implement me")
}
