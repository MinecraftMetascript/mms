package project

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/ast"
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

		blockSpec := lang.Blocks.Get(blockKind.GetText())
		if blockSpec == nil {
			// TODO: ✏ Diagnose -- Invalid Kind
			continue
		}
		block, diags := blockSpec.Match(b)

		if diags != nil {
			fmt.Println("Diagnostics:")
			for i, diag := range diags {
				fmt.Println("\t", i, ":", diag)
			}
		}
		for name, decl := range block.Declarations {
			if exportable, ok := decl.(ast.Symbol); ok {
				res := exportable.Export(name)
				res.PrintDebug()
			}
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
