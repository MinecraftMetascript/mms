package traversal

import (
	"errors"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammar"

	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
	grammar.BaseMinecraftMetascriptListener

	lexer       *grammar.MinecraftMetascriptLexer
	parser      *grammar.MinecraftMetascriptParser
	content     string
	diagnostics *[]Diagnostic

	scope     *Scope
	namespace string // Current namespace
	filename  string
}

func (p *Parser) GetInternalParser(namespace string) *grammar.MinecraftMetascriptParser {
	p.namespace = namespace
	return p.parser
}

func (p *Parser) Parse() (*grammar.ScriptContext, error) {
	fmt.Println("Parsing...")
	if res := p.parser.Script(); res == nil {
		return nil, errors.New("failed to parse")
	} else {
		fmt.Println("Done parsing")
		return res.(*grammar.ScriptContext), nil
	}
}

func (p *Parser) ExitNamespace(ctx *grammar.NamespaceContext) {
	for _, contentBlockCtx := range ctx.AllContentBlocks() {
		if contentBlock, ok := contentBlockCtx.(*grammar.ContentBlocksContext); ok {
			inner := contentBlock.GetChild(0)
			ConstructRegistry.Construct(inner.(antlr.ParserRuleContext), p.namespace, p.scope)

		}
	}
}

func (p *Parser) ExitNamespaceDeclaration(ctx *grammar.NamespaceDeclarationContext) {
	ns := ctx.Identifier()
	if ns != nil {
		p.namespace = ns.GetText()
	} else {
		*p.diagnostics = append(*p.diagnostics, Diagnostic{
			Message:  "Missing Namespace Declaration",
			Where:    RuleLocation(ctx, p.filename),
			Severity: SeverityError,
			Source:   "parser",
			File:     p.filename,
		})
		p.namespace = "[[missing]]"
	}
}

func NewParser(content string, filename string, globalScope *Scope, diagnostics *[]Diagnostic) *Parser {
	input := antlr.NewInputStream(content)
	lexer := grammar.NewMinecraftMetascriptLexer(input)

	out := &Parser{
		lexer:       lexer,
		parser:      grammar.NewMinecraftMetascriptParser(antlr.NewCommonTokenStream(lexer, 0)),
		content:     content,
		diagnostics: diagnostics,
		scope:       globalScope,
		filename:    filename,
	}

	// Attach diagnostics listener(s) instead of console listeners
	if out.diagnostics == nil {
		out.diagnostics = &[]Diagnostic{}
	}
	diagListener := NewDiagnosticsErrorListener(content, out.filename, out.diagnostics)
	lexer.RemoveErrorListeners()
	parser := out.parser
	parser.RemoveErrorListeners()
	lexer.AddErrorListener(diagListener)
	parser.AddErrorListener(diagListener)
	parser.AddParseListener(out)

	// Wire diagnostics sink into scope for construct factories
	if out.scope != nil {
		out.scope.Diagnostics = out.diagnostics
		out.scope.CurrentFile = out.filename
	}

	return out
}
