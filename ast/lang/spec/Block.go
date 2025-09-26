package spec

import (
	"github.com/minecraftmetascript/mms/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Block struct {
	Kind   string
	Values []Value
}

func (b Block) Validate(ctx grammar.IValueContext) (bool, error) {
	return false, nil
}

func (b Block) GetLabel() string {
	return b.Kind
}

type BlockNode struct {
	Kind         string
	Declarations map[string]ValueNode
	location     ast.SourceLocation
}

func (b BlockNode) GetLocation() ast.SourceLocation {
	return b.location
}

func getBlockValue(ctx grammar.IBlockContext, possibilities []Block) *BlockNode {
	kind := ctx.Identifier()
	if kind == nil {
		// TODO: ✏ Diagnose -- missing block kind
		return nil
	}
	for _, possible := range possibilities {
		if possible.Kind != kind.GetText() {
			continue
		}
		out := &BlockNode{
			Kind:         possible.Kind,
			Declarations: make(map[string]ValueNode),
			location:     ast.SourceLocation{},
		}
		for _, decl := range ctx.AllVarDecl() {
			varName := decl.Identifier()
			if varName == nil {
				// TODO: ✏ Diagnose -- missing variable name
				continue
			}
			val := decl.Value()
			res := GetValue(val, possible.Values)
			out.Declarations[varName.GetText()] = res

		}

		return out
	}

	return nil
}
