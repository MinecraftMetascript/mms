package noise

import (
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type NoiseBlock struct {
	location     traversal.TextLocation
	Declarations []traversal.Symbol
}

func (n NoiseBlock) GetLocation() traversal.TextLocation {
	return n.location
}

func init() {
	traversal.RegisterNodeFactory(NoiseBlockFactory{})
}

type NoiseBlockFactory struct{}

func (n NoiseBlockFactory) Create(ctx *grammar.NoiseBlockContext, namespace string, scope *traversal.Scope) *NoiseBlock {
	out := &NoiseBlock{
		location:     traversal.RuleLocation(ctx, scope.CurrentFile),
		Declarations: make([]traversal.Symbol, 0),
	}
	for _, fn := range ctx.AllNoiseDeclaration() {
		res, ok := traversal.DeclareNode(fn, namespace, scope)
		if ok {
			out.Declarations = append(out.Declarations, res)
		}
	}
	return out
}
