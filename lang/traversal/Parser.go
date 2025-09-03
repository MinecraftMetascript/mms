package traversal

import (
	"errors"

	"github.com/minecraftmetascript/mms/lang/grammar"

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
	filename  string
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
	if ns != nil {
		p.namespace = ns.GetText()
	} else {
		*p.diagnostics = append(*p.diagnostics, Diagnostic{
			Message:  "Missing Namespace Declaration",
			Where:    RuleLocation(ctx, p.filename),
			Severity: SeverityError,
			Source:   "parser",
		})
		p.namespace = "_unnamed_"
	}
}

func (p *Parser) ExitDeclaration(ctx *grammar.DeclarationContext) {
	nameCtx := ctx.Identifier()
	definitionCtx := ctx.Definition()
	// Validate contexts are present
	if nameCtx == nil || definitionCtx == nil {
		if p.diagnostics != nil {
			*p.diagnostics = append(*p.diagnostics, Diagnostic{
				Message:  "invalid declaration",
				Where:    RuleLocation(ctx, p.filename),
				Severity: SeverityError,
				Source:   "parser",
			})
		}
		return
	}
	name := nameCtx.GetText()

	var val Construct
	if definitionCtx.GetChildCount() > 0 {
		if inner, ok := definitionCtx.GetChild(0).(antlr.ParserRuleContext); ok {
			val = ConstructRegistry.Construct(inner, p.namespace, p.scope)
		}
	}

	if val == nil {
		if p.diagnostics != nil {
			*p.diagnostics = append(*p.diagnostics, Diagnostic{
				Message:  "invalid definition for " + name,
				Where:    RuleLocation(definitionCtx, p.filename),
				Severity: SeverityError,
				Source:   "parser",
			})
		}
		return
	}

	symbol := NewSymbol(
		TerminalNodeLocation(nameCtx, p.filename),
		RuleLocation(definitionCtx, p.filename),
		val,
		NewReference(name, p.namespace),
	)
	if err := p.scope.Register(symbol); err != nil {
		if p.diagnostics != nil {
			*p.diagnostics = append(*p.diagnostics, Diagnostic{
				Message:  err.Error(),
				Where:    TerminalNodeLocation(nameCtx, p.filename),
				Severity: SeverityError,
				Source:   "semantic",
			})
		}
	}
}

func NewParser(content string, filename string, globalScope *Scope, diagnostics *[]Diagnostic) *Parser {
	input := antlr.NewInputStream(content)
	lexer := grammar.NewMain_Lexer(input)

	out := &Parser{
		lexer:       lexer,
		parser:      grammar.NewMain_Parser(antlr.NewCommonTokenStream(lexer, 0)),
		content:     content,
		diagnostics: diagnostics,
		scope:       globalScope,
		filename:    filename,
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
