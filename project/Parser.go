package project

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang"
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/spec"
)

type Parser struct {
	grammar.BaseMinecraftMetascriptListener
	lexer  *grammar.MinecraftMetascriptLexer
	parser *grammar.MinecraftMetascriptParser

	namespaces  map[string]*ast.Namespace
	blocks      []ast.Node
	diagnostics *ast.Diagnostics
	filename    string
}

func (p *Parser) ExitNamedBlock(ctx *grammar.NamedBlockContext) {
	kind := ctx.Identifier(0)
	if kind == nil {
		return
	}
	if kind.GetText() != "Namespace" {
		// TODO: ✏ Diagnose -- Invalid GetKind
		return
	}

	nsIdCtx := ctx.Identifier(1)
	if nsIdCtx == nil {
		p.diagnostics.Add(ast.Diagnostic{
			Location: ast.RuleLocation(ctx),
			Message:  "Namespace must have a name.",
			Severity: ast.Error,
		})
	}
	namespace := "__unknown"
	if nsIdCtx != nil {
		namespace = nsIdCtx.GetText()
	}

	ns := ast.NewNamespace(namespace)
	inlinedNs := ast.NewNamespace("mms_inline")

	for _, b := range ctx.AllBlock() {
		blockKind := b.Identifier()
		if blockKind == nil {
			p.diagnostics.Add(ast.Diagnostic{
				Location: ast.RuleLocation(b),
				Message:  "Block must have a kind.",
				Severity: ast.Error,
			})
			continue
		}

		blockSpec := lang.Blocks.Get(blockKind.GetText())
		if blockSpec == nil {
			p.diagnostics.Add(ast.Diagnostic{
				Location: ast.TerminalLocation(blockKind),
				Message:  fmt.Sprintf("Unknown block kind: %s", blockKind),
				Severity: ast.Error,
			})
			continue
		}
		block, diags := blockSpec.Match(b)

		p.diagnostics.Add(diags...)
		if block == nil {
			continue
		}
		p.blocks = append(p.blocks, block)

		spec.SetDefaultNamespaces(block, namespace)
		spec.SetFilenames(block, p.filename)
		for name, decl := range block.Declarations {
			if symbol, ok := decl.(ast.Symbol); ok {
				symbol.SetRef(fmt.Sprintf("%s:%s", namespace, name))
				ns.Declare(
					name, symbol,
				)
			}
		}
		for _, decl := range block.ExtractInlineSymbols() {
			decl.SetRef(ast.InlineSymbolId(decl))
			inlinedNs.Declare(
				ast.InlineSymbolId(decl), decl,
			)
		}
		if existing, ok := p.namespaces[namespace]; ok {
			existing.Merge(ns)
		} else {
			p.namespaces[namespace] = ns
		}
		if existing, ok := p.namespaces["mms_inline"]; ok {
			existing.Merge(inlinedNs)
		} else {
			p.namespaces["mms_inline"] = inlinedNs
		}
	}

	if p.namespaces["mms_inline"].Size() == 0 {
		delete(p.namespaces, "mms_inline")
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
		blocks:      make([]ast.Node, 0),
		filename:    filename,
	}

	out.parser.RemoveErrorListeners()
	out.parser.AddErrorListener(ast.NewDiagnosticsErrorListener(content, filename, out.diagnostics))

	out.lexer.RemoveErrorListeners()
	out.lexer.AddErrorListener(ast.NewDiagnosticsErrorListener(content, filename, out.diagnostics))

	out.parser.AddParseListener(out)

	return out
}
