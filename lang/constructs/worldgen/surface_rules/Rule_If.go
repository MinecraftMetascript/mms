package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRule_IfContext](),
		func(ctx_ antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			ctx := ctx_.(*grammar.SurfaceRule_IfContext)
			out := &IfRule{
				scope: scope,
			}

			var condition traversal.Construct

			if cond := ctx.SurfaceCondition(); cond != nil {
				condition = traversal.ConstructRegistry.Construct(cond, ns, scope)
			} else {
				if scope != nil && scope.Diagnostics != nil {
					loc := traversal.RuleLocation(ctx, scope.CurrentFile)
					*scope.Diagnostics = append(*scope.Diagnostics, traversal.Diagnostic{
						Message:  "Expected a surface condition",
						Where:    loc,
						Severity: traversal.SeverityError,
						Source:   "semantic",
						File:     scope.CurrentFile,
					})
				}
			}

			if ref, ok := condition.(*ReferenceCondition); ok {
				ref.Ref.SetResolver(func() error {
					if next, ok := scope.Get(*ref.Ref); ok {
						out.Condition = next.GetValue()
					}
					return nil
				})
			}

			var action traversal.Construct
			if rule := ctx.SurfaceRule(); rule != nil {
				if rule.GetChildCount() > 0 {
					if inner, ok := rule.GetChild(0).(antlr.ParserRuleContext); ok {
						action = traversal.ConstructRegistry.Construct(inner, ns, scope)
					}
				}
			} else {
				if scope != nil && scope.Diagnostics != nil {
					loc := traversal.RuleLocation(ctx, scope.CurrentFile)
					*scope.Diagnostics = append(*scope.Diagnostics, traversal.Diagnostic{
						Message:  "missing nested rule in if rule",
						Where:    loc,
						Severity: traversal.SeverityError,
						Source:   "semantic",
						File:     scope.CurrentFile,
					})
				}
			}

			out.Action = action
			out.Condition = condition

			return out
		})
}

type IfRule struct {
	traversal.Construct
	Condition traversal.Construct
	Action    traversal.Construct
	scope     *traversal.Scope
}

func (c IfRule) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceRule(symbol, rootDir, c)
}

func (c IfRule) MarshalJSON() ([]byte, error) {
	if refCond, ok := c.Condition.(*ReferenceCondition); ok {
		if cond, ok := c.scope.Get(*refCond.Ref); ok {
			c.Condition = cond.GetValue()
		}
	}

	current := c
	if and, ok := c.Condition.(*CompoundAndCondition); ok {
		for i, cond := range and.Conditions {
			next := c.Action
			if i > 0 {
				next = current
			}
			current = IfRule{
				Condition: cond,
				Action:    next,
			}
		}
	}

	if or, ok := c.Condition.(*CompoundOrCondition); ok {
		sequenceRule := &Sequence{
			Rules: make([]traversal.Construct, 0),
		}
		for _, cond := range or.Conditions {
			sequenceRule.Rules = append(sequenceRule.Rules,
				&IfRule{
					Action:    c.Action,
					Condition: cond,
					scope:     c.scope,
				},
			)
		}
		return sequenceRule.MarshalJSON()
	}

	return json.MarshalIndent(struct {
		Type    SurfaceRuleKind     `json:"type"`
		IfTrue  traversal.Construct `json:"if_true"`
		ThenRun traversal.Construct `json:"then_run"`
	}{
		Type:    SurfaceRule_Condition,
		IfTrue:  current.Condition,
		ThenRun: current.Action,
	}, "", "  ")
}
