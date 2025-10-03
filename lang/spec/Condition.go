package spec

import (
	"log"
	"reflect"
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
)

type ConditionalSpec struct {
	ConditionOptions []ValueSpec
	ValueOptions     []ValueSpec
}

func NewConditionSpec() *ConditionalSpec {
	return &ConditionalSpec{
		ConditionOptions: []ValueSpec{},
		ValueOptions:     []ValueSpec{},
	}
}

func (c *ConditionalSpec) AddConditionOption(spec ...ValueSpec) *ConditionalSpec {
	c.ConditionOptions = slices.Concat(c.ConditionOptions, spec)
	return c
}

func (c *ConditionalSpec) AddValueOption(spec ...ValueSpec) *ConditionalSpec {
	c.ValueOptions = slices.Concat(c.ValueOptions, spec)
	return c
}

func (c *ConditionalSpec) matchConditionCtx(ctx grammar.IConditionContext) (ast.Node, []ast.Diagnostic) {
	switch condition := ctx.(type) {
	// TODO: If condition != primary, we need to match the children against the same spec
	case *grammar.CondAndContext:
		log.Println("AND", condition)
	case *grammar.CondOrContext:
		log.Println("OR", condition)
	case *grammar.CondGroupedOrContext:
		log.Println("GROUPED OR", condition)
	case *grammar.CondGroupedAndContext:
		log.Println("GROUPED AND", condition)
	case *grammar.CondNegateContext:
		log.Println("NEGATE", condition)
	case *grammar.CondPrimaryContext:
		for _, opt := range c.ConditionOptions {
			r, diags := opt.Match(condition.RootCondition().Value())
			if !lib.IsNilInterface(r) {
				return r, diags
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
		BaseNode: ast.BaseNode{
			Location: &l,
		},
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
				out.Value = valueNode
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
	ast.BaseNode
	Condition ast.Node
	Value     ast.Node
}

func (c ConditionalNode) Children() []ast.Node {
	return []ast.Node{c.Condition, c.Value}
}
