package factories

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type NoiseBlockFactory struct{}

func (n NoiseBlockFactory) Match(ctx antlr.ParserRuleContext) bool {
	blockCtx, ok := ctx.(*grammar.BlockContext)
	if !ok {
		return false
	}
	kind := blockCtx.Identifier()
	if kind == nil {
		// TODO: ✏ Diagnose
	}
	return kind.GetText() == "Noise"
}

func (n NoiseBlockFactory) ConstructNode(ctx antlr.ParserRuleContext) ast.Node {
	blockCtx, ok := ctx.(*grammar.BlockContext)
	if !ok {
		panic("Expected BlockContext")
	}

	decls := blockCtx.AllVarDecl()
	for _, decl := range decls {
		log.Println("Found decl", decl.Identifier().GetText())
	}

	return nil
}

func init() {
	ast.Registry.Register(NoiseBlockFactory{})
}
