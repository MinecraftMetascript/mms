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

func (p *Parser) Parse() (err error) {
	if p.parser.Script() == nil {
		return errors.New("failed to parse")
	}
	return nil
}

func (p *Parser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	ConstructRegistry.Construct(ctx, p.namespace, p.scope)
}

func (p *Parser) ExitNamespaceDeclaration(ctx *grammar.NamespaceDeclarationContext) {
	ns := ctx.Identifier()
	if ns != nil {
		p.namespace = ns.GetText()
		fmt.Println("BEGIN: ", p.namespace)
	} else {
		*p.diagnostics = append(*p.diagnostics, Diagnostic{
			Message:  "Missing Namespace Declaration",
			Where:    RuleLocation(ctx, p.filename),
			Severity: SeverityError,
			Source:   "parser",
			File:     p.filename,
		})
		p.namespace = "_unnamed_"
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
