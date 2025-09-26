package spec

import (
	"github.com/minecraftmetascript/mms/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Value interface {
	Validate(ctx grammar.IValueContext) (bool, error)
	GetLabel() string
}

type ValueNode interface {
	GetLocation() ast.SourceLocation
}
