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

	namespaces  map[string]*ast.Namespace
	diagnostics *ast.Diagnostics
	filename    string
}

func (p *Parser) ExitNamedBlock(ctx *grammar.NamedBlockContext) {
	kind := ctx.Identifier(0)
	if kind == nil {
		// TODO: ✏ Diagnose -- No GetKind
		return
	}
	if kind.GetText() != "Namespace" {
		// TODO: ✏ Diagnose -- Invalid GetKind
		return
	}

	nsIdCtx := ctx.Identifier(1)
	if nsIdCtx == nil {
		// TODO: ✏ Diagnose -- Missing namespace name
	}
	namespace := "__unknown"
	if nsIdCtx != nil {
		namespace = nsIdCtx.GetText()
	}

	ns := ast.NewNamespace(namespace)

	for _, b := range ctx.AllBlock() {
		blockKind := b.Identifier()
		if blockKind == nil {
			// TODO: ✏ Diagnose -- No GetKind
			continue
		}

		blockSpec := lang.Blocks.Get(blockKind.GetText())
		if blockSpec == nil {
			// TODO: ✏ Diagnose -- Invalid GetKind
			continue
		}
		block, diags := blockSpec.Match(b)
		if block == nil {
			// TODO: ✏ Diagnose -- Invalid GetKind
			continue
		}

		p.diagnostics.Add(diags...)

		if diags != nil {
			fmt.Println("Diagnostics:")
			for i, diag := range diags {
				fmt.Println("\t", i, ":", diag)
			}
		}
		for name, decl := range block.Declarations {
			if symbol, ok := decl.(ast.Symbol); ok {
				symbol.SetFilename(p.filename)
				ns.Declare(
					name, symbol,
				)
			}
		}
		p.namespaces[namespace] = ns
	}
}

func NewParser(content, filename string) *Parser {
	input := antlr.NewInputStream(content)
	lexer := grammar.NewMinecraftMetascriptLexer(input)
	out := &Parser{
		lexer:       lexer,
		parser:      grammar.NewMinecraftMetascriptParser(antlr.NewCommonTokenStream(lexer, 0)),
		namespaces:  make(map[string]*ast.Namespace),
		diagnostics: ast.NewDiagnostics(),
		filename:    filename,
	}

	out.parser.RemoveErrorListeners()
	out.parser.AddErrorListener(ast.NewDiagnosticsErrorListener(content, filename, out.diagnostics))

	out.lexer.RemoveErrorListeners()
	out.lexer.AddErrorListener(ast.NewDiagnosticsErrorListener(content, filename, out.diagnostics))

	out.parser.AddParseListener(out)

	return out
}
