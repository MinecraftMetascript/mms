package traversal

import (
	"errors"
	"mms2/lang/grammar"

	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
	grammar.BaseMain_ParserListener

	lexer       *grammar.Main_Lexer
	parser      *grammar.Main_Parser
	content     string
	diagnostics *[]Diagnostic

	scope     *Scope
	namespace string // Current namespace
}

func (p *Parser) GetInternalParser(namespace string) *grammar.Main_Parser {
	p.namespace = namespace
	return p.parser
}

func (p *Parser) Parse() (err error) {
	if p.parser.Script() == nil {
		return errors.New("failed to parse")
	}
	return nil
}

func (p *Parser) ExitNamespaceDefinition(ctx *grammar.NamespaceDefinitionContext) {
	ns := ctx.Identifier()
	p.namespace = ns.GetText()
}

func (p *Parser) ExitDeclaration(ctx *grammar.DeclarationContext) {
	nameCtx := ctx.Identifier()
	definitionCtx := ctx.Definition()
	// Validate contexts are present
	if nameCtx == nil || definitionCtx == nil {
		if p.diagnostics != nil {
			*p.diagnostics = append(*p.diagnostics, Diagnostic{
				Message:  "invalid declaration",
				Where:    RuleLocation(ctx, p.content),
				Severity: SeverityError,
				Source:   "parser",
			})
		}
		return
	}
	name := nameCtx.GetText()

	val := ConstructRegistry.Construct(definitionCtx.GetChild(0).(antlr.ParserRuleContext), p.namespace, p.scope)

	if val == nil {
		if p.diagnostics != nil {
			*p.diagnostics = append(*p.diagnostics, Diagnostic{
				Message:  "invalid definition for " + name,
				Where:    RuleLocation(definitionCtx, p.content),
				Severity: SeverityError,
				Source:   "parser",
			})
		}
		return
	}

	symbol := NewSymbol(
		TerminalNodeLocation(nameCtx, p.content),
		RuleLocation(definitionCtx, p.content),
		val,
		NewReference(name, p.namespace),
	)

	if err := p.scope.Register(symbol); err != nil {
		if p.diagnostics != nil {
			*p.diagnostics = append(*p.diagnostics, Diagnostic{
				Message:  err.Error(),
				Where:    TerminalNodeLocation(nameCtx, p.content),
				Severity: SeverityError,
				Source:   "semantic",
			})
		}
	}
}

func NewParser(content string, globalScope *Scope, diagnostics *[]Diagnostic) *Parser {
	input := antlr.NewInputStream(content)
	lexer := grammar.NewMain_Lexer(input)

	out := &Parser{
		lexer:       lexer,
		parser:      grammar.NewMain_Parser(antlr.NewCommonTokenStream(lexer, 0)),
		content:     content,
		diagnostics: diagnostics,
		scope:       globalScope,
	}

	// Attach diagnostics listener(s) instead of console listeners
	if out.diagnostics == nil {
		out.diagnostics = &[]Diagnostic{}
	}
	diagListener := NewDiagnosticsErrorListener(content, out.diagnostics)
	lexer.RemoveErrorListeners()
	parser := out.parser
	parser.RemoveErrorListeners()
	lexer.AddErrorListener(diagListener)
	parser.AddErrorListener(diagListener)
	parser.AddParseListener(out)

	return out
}
