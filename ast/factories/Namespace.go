package factories

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type NamespaceFactory struct{}

func (n NamespaceFactory) Match(ctx antlr.ParserRuleContext) bool {
	namedBlockCtx, ok := ctx.(*grammar.NamedBlockContext)
	if !ok {
		return false
	}
	kind := namedBlockCtx.Identifier(0)
	return kind.GetText() == "Namespace"
}

func (n NamespaceFactory) ConstructNode(ctx antlr.ParserRuleContext) ast.Node {
	namedBlockCtx, ok := ctx.(*grammar.NamedBlockContext)
	if !ok {
		panic("Expected NamedBlockContext")
	}
	name := namedBlockCtx.Identifier(1)
	if name == nil {
		// TODO: ✏ Diagnose
	}
	log.Println("Beginning to parse namespace", name)

	out := Namespace{
		Name:   name.GetText(),
		Blocks: make([]ast.Node, 0),
	}
	for _, block := range namedBlockCtx.AllBlock() {
		log.Println("Found child block", block.Identifier().GetText())

		out.Blocks = append(out.Blocks, ast.Registry.Construct(block))
	}

	return out
}

type Namespace struct {
	Name   string
	Blocks []ast.Node
}

func init() {
	ast.Registry.Register(NamespaceFactory{})
}
