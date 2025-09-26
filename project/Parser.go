package project

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/ast/lang"
	"github.com/minecraftmetascript/mms/ast/lang/spec"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Parser struct {
	grammar.BaseMinecraftMetascriptListener
	lexer  *grammar.MinecraftMetascriptLexer
	parser *grammar.MinecraftMetascriptParser
}

func (p Parser) ExitNamedBlock(ctx *grammar.NamedBlockContext) {
	kind := ctx.Identifier(0)
	if kind == nil {
		// TODO: ✏ Diagnose -- No Kind
		return
	}
	if kind.GetText() != "Namespace" {
		// TODO: ✏ Diagnose -- Invalid Kind
		return
	}

	for _, b := range ctx.AllBlock() {
		blockKind := b.Identifier()
		if blockKind == nil {
			// TODO: ✏ Diagnose -- No Kind
			continue
		}
		blockSpec := lang.GetBlockSpec(blockKind.GetText())
		if blockSpec == nil {
			// TODO: ✏ Diagnose -- Invalid Kind
			continue
		}
		for _, v := range b.AllVarDecl() {
			varName := v.Identifier()
			if varName == nil {
				// TODO: ✏ Diagnose -- Missing Name
				continue
			}
			log.Println("Identified variable: ", varName)
			spec.GetValue(v.Value(), blockSpec.Values)
		}
	}
}

func NewParser(content, filename string) *Parser {
	input := antlr.NewInputStream(content)
	lexer := grammar.NewMinecraftMetascriptLexer(input)
	out := &Parser{
		lexer:  lexer,
		parser: grammar.NewMinecraftMetascriptParser(antlr.NewCommonTokenStream(lexer, 0)),
	}

	out.parser.AddParseListener(out)

	return out
}
