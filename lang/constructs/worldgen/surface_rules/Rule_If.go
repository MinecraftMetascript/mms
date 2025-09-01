package surface_rules

import (
	"encoding/json"
	"fmt"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRule_ConditionContext](),
		func(ctx_ antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			ctx := ctx_.(*grammar.SurfaceRule_ConditionContext)
			out := &IfRule{
				Negate: ctx.SurfaceCondition().Bang() != nil,
				scope:  scope,
			}

			var condition traversal.Construct

			if cond := ctx.SurfaceCondition(); cond != nil {
				var inner antlr.ParserRuleContext
				if len(cond.GetChildren()) == 1 {
					inner = cond.GetChild(0).(antlr.ParserRuleContext) // not negated
				} else {
					inner = cond.GetChild(1).(antlr.ParserRuleContext) // has bang
				}

				condition = traversal.ConstructRegistry.Construct(inner, ns, scope)
			} else {
				fmt.Println("[ERROR] Condition isn't found!", ctx.SurfaceCondition())
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
				inner := rule.GetChild(0).(antlr.ParserRuleContext)
				action = traversal.ConstructRegistry.Construct(inner, ns, scope)
			} else {
				fmt.Println("[ERROR] Rule isn't found!", ctx.SurfaceRule())
			}

			out.Action = action
			out.Condition = condition

			return out
		})
}

type IfRule struct {
	traversal.Construct
	Negate    bool
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

	if c.Negate {
		c.Negate = false
		condition := c.Condition
		c.Condition = InvertCondition(condition)
	}

	current := c
	if and, ok := c.Condition.(*CompoundAndCondition); ok {
		for i, cond := range and.Conditions {
			next := c.Action
			if i > 0 {
				next = current
			}
			current = IfRule{
				Negate:    false,
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
					Negate:    false,
					Action:    c.Action,
					Condition: cond,
				},
			)
		}
		return sequenceRule.MarshalJSON()
	}

	return json.Marshal(struct {
		Type    SurfaceRuleKind     `json:"type"`
		IfTrue  traversal.Construct `json:"if_true"`
		ThenRun traversal.Construct `json:"then_run"`
	}{
		Type:    SurfaceRule_Condition,
		IfTrue:  current.Condition,
		ThenRun: current.Action,
	})
}
