package spec

import (
	"fmt"
	"log"
	"reflect"
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ConditionalSpec struct {
	ConditionOptions []ValueSpec
	ValueOptions     []ValueSpec
	Kind             ast.SymbolKind
	Help             string
	export           func(fn ConditionalNode, name string) *lib.FileTreeLike
	output           func(fn ConditionalNode) any
}

func NewConditionSpec() *ConditionalSpec {
	return &ConditionalSpec{
		ConditionOptions: []ValueSpec{},
		ValueOptions:     []ValueSpec{},
	}
}

func (c *ConditionalSpec) SetKind(kind ast.SymbolKind) *ConditionalSpec {
	c.Kind = kind
	return c
}
func (c *ConditionalSpec) SetHelp(help string) *ConditionalSpec {
	c.Help = help
	return c
}

func (c *ConditionalSpec) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	return []protocol.CompletionItem{{
		Label:            fmt.Sprintf("[%s] Condition", c.Kind),
		Kind:             &MethodKind,
		Detail:           &c.Help,
		InsertTextFormat: &SnippetFormat,
		TextEdit: protocol.TextEdit{
			Range: protocol.Range{
				Start: position,
				End:   position,
			},
			NewText: "If(${1}) ${2}",
		},
	},
	}
}

func (c *ConditionalSpec) SetFileExporter(exporter func(n ConditionalNode, name string) *lib.FileTreeLike) *ConditionalSpec {
	c.export = exporter
	return c
}
func (c *ConditionalSpec) SetOutputFn(outputFn func(n ConditionalNode) any) *ConditionalSpec {
	c.output = outputFn
	return c
}

func (c *ConditionalSpec) AddConditionOption(spec ...ValueSpec) *ConditionalSpec {
	c.ConditionOptions = slices.Concat(c.ConditionOptions, spec)
	return c
}

func (c *ConditionalSpec) AddValueOption(spec ...ValueSpec) *ConditionalSpec {
	c.ValueOptions = slices.Concat(c.ValueOptions, spec)
	return c
}

type ConditionAndNode struct {
	ast.BaseSymbol
	Left  ast.Symbol
	Right ast.Symbol
}
type ConditionOrNode struct {
	ast.BaseSymbol
	Left  ast.Symbol
	Right ast.Symbol
}
type ConditionNegateNode struct {
	ast.BaseSymbol
	Condition ast.Symbol
}

func (c *ConditionalSpec) matchConditionCtx(ctx grammar.IConditionContext) (ast.Symbol, []ast.Diagnostic) {
	if ctx == nil {
		return nil, nil
	}
	switch condition := ctx.(type) {
	// TODO: If condition != primary, we need to match the children against the same spec
	// TODO: Create a new node type for paired conditions, as well as one for negated conditions
	// TODO: Is the distinction between grouped and ungrouped conditions necessary here? That may be handled by the parser
	case *grammar.CondAndContext:
		l, ld := c.matchConditionCtx(condition.Condition(0))
		r, rd := c.matchConditionCtx(condition.Condition(1))
		loc := ast.RuleLocation(ctx)
		return &ConditionAndNode{
			BaseSymbol: ast.BaseSymbol{
				Location: &loc,
				BaseNode: ast.BaseNode{},
			},
			Left:  l,
			Right: r,
		}, slices.Concat(ld, rd)
	case *grammar.CondOrContext:
		l, ld := c.matchConditionCtx(condition.Condition(0))
		r, rd := c.matchConditionCtx(condition.Condition(1))
		loc := ast.RuleLocation(ctx)
		return &ConditionOrNode{
			BaseSymbol: ast.BaseSymbol{
				Location: &loc,
				BaseNode: ast.BaseNode{},
			},
			Left:  l,
			Right: r,
		}, slices.Concat(ld, rd)
	case *grammar.CondGroupedContext:
		// Pass this through
		return c.matchConditionCtx(condition.Condition())
	case *grammar.CondNegateContext:
		neg, cd := c.matchConditionCtx(condition.Condition())
		loc := ast.RuleLocation(ctx)
		return &ConditionNegateNode{
			BaseSymbol: ast.BaseSymbol{
				Location: &loc,
				BaseNode: ast.BaseNode{},
			},
			Condition: neg,
		}, cd
	case *grammar.CondPrimaryContext:
		for _, opt := range c.ConditionOptions {
			r, diags := opt.Match(condition.RootCondition().Value())
			if !lib.IsNilInterface(r) {
				if out, ok := r.(ast.Symbol); ok {
					return out, diags
				}
				diags = append(diags, ast.Diagnostic{
					Location: ast.RuleLocation(condition.RootCondition().Value()),
					Message:  "Invalid condition",
					Severity: ast.Error,
				})
				return nil, diags
			}
		}
	default:
		log.Printf("UNKNOWN: %s", reflect.TypeOf(condition))
	}
	return nil, []ast.Diagnostic{
		{
			Location: ast.RuleLocation(ctx),
			Message:  "Invalid condition",
			Severity: ast.Warning,
		},
	}
}

func (c *ConditionalSpec) Match(valueCtx grammar.IValueContext) (ast.Node, []ast.Diagnostic) {
	if valueCtx.Conditional() == nil {
		return nil, nil
	}
	conditionCtx := valueCtx.Conditional()

	l := ast.RuleLocation(valueCtx)
	out := &ConditionalNode{
		BaseSymbol: ast.BaseSymbol{
			Location: &l,
			BaseNode: ast.BaseNode{},
		},
		spec: c,
	}
	allDiags := make([]ast.Diagnostic, 0)
	iCondition := conditionCtx.Condition()
	if iCondition != nil {
		conditionNode, conditionDiags := c.matchConditionCtx(iCondition)
		if conditionDiags != nil {
			allDiags = slices.Concat(allDiags, conditionDiags)
		}
		if conditionNode != nil {
			out.Condition = conditionNode
		}
	}

	if out.Condition == nil {
		allDiags = append(allDiags, ast.Diagnostic{
			Location: ast.RuleLocation(valueCtx),
			Message:  "Missing condition",
			Severity: ast.Warning,
		})
	}

	v := conditionCtx.Value()
	if v != nil {
		for _, opt := range c.ValueOptions {
			valueNode, valueDiags := opt.Match(v)
			if valueDiags != nil {
				allDiags = slices.Concat(allDiags, valueDiags)
			}
			if !lib.IsNilInterface(valueNode) {
				if value, ok := valueNode.(ast.Symbol); ok {
					out.Value = value
				}
				break
			}
		}
		if out.Value == nil {
			allDiags = append(allDiags, ast.Diagnostic{
				Location: ast.RuleLocation(v),
				Message:  "Invalid value",
				Severity: ast.Warning,
			})
		}
	} else {
		allDiags = append(allDiags, ast.Diagnostic{
			Location: ast.RuleLocation(valueCtx),
			Message:  "Missing value",
			Severity: ast.Warning,
		})
	}

	return out, allDiags
}

type ConditionalNode struct {
	ast.BaseSymbol
	Condition ast.Symbol
	Value     ast.Symbol
	spec      *ConditionalSpec
}

func (c ConditionalNode) Children() []ast.Node {
	return []ast.Node{c.Condition, c.Value}
}

func (c ConditionalNode) Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*ast.Namespace) []protocol.CompletionItem {
	if c.Condition != nil && c.Condition.GetLocation().ContainsPosition(position) {
		// Complete the condition?
		if comp, ok := c.Condition.(ast.CompletableNode); ok {
			return comp.Complete(fileSource, position, triggerChar, symbols)
		}
	} else if c.Value != nil && c.Value.GetLocation().ContainsPosition(position) {
		if comp, ok := c.Value.(ast.CompletableNode); ok {
			return comp.Complete(fileSource, position, triggerChar, symbols)
		}
	}

	fullLocation := c.GetLocation()

	mode := "NONE"
	for idx := position.IndexIn(fileSource); idx > fullLocation.Start.Index; idx-- {
		if fileSource[idx] == ')' {
			// We are in value mode
			mode = "VALUE"
			break
		} else if fileSource[idx] == '(' {
			// We are in condition mode
			mode = "CONDITION"
			break
		}
		// We don't have anything right now
	}

	out := make([]protocol.CompletionItem, 0)
	switch mode {
	case "VALUE":
		for _, spec := range c.spec.ValueOptions {
			if comp, ok := spec.(ast.CompletableNode); ok {
				out = append(out, comp.Complete(fileSource, position, triggerChar, symbols)...)
			}
		}
	case "CONDITION":
		for _, spec := range c.spec.ConditionOptions {
			if comp, ok := spec.(ast.CompletableNode); ok {
				out = append(out, comp.Complete(fileSource, position, triggerChar, symbols)...)
			}
		}
	}

	return out
}

func (c ConditionalNode) ToFileTreeLike(name string) *lib.FileTreeLike {
	if c.spec.export == nil {
		return nil
	}
	return c.spec.export(c, name)
}

func (c ConditionalNode) ToSerializable() any {
	if c.spec.output == nil {
		return nil
	}
	return c.spec.output(c)
}

func (c ConditionalNode) GetKind() ast.SymbolKind {
	return c.spec.Kind
}
