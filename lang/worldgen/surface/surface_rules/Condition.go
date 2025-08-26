package surface_rules

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_conditions"
)

func NewConditionalRule(ctx *grammars.SurfaceRule_ConditionalContext, ruleFactory RuleFactory, conditionFactory surface_conditions.ConditionFactory) (*ConditionalRule, []error) {
	condition, err := conditionFactory.ConstructSurfaceCondition(ctx.SurfaceCondition().(*grammars.SurfaceConditionContext))
	if err != nil {
		return nil, []error{err}
	}
	if condition == nil {
		return nil, []error{fmt.Errorf("condition is nil")}
	}

	action, errs := ruleFactory.ConstructSurfaceRule(ctx.SurfaceRule().(*grammars.SurfaceRuleContext))
	if errs != nil {
		return nil, errs
	}
	if action == nil {
		return nil, []error{fmt.Errorf("action is nil")}
	}

	return &ConditionalRule{
		Negate:    ctx.Bang() != nil,
		Condition: condition,
		Action:    action,
	}, nil
}

type ConditionalRule struct {
	BaseRule
	Negate    bool
	Condition surface_conditions.SurfaceCondition
	Action    SurfaceRule
}

func (r ConditionalRule) Type() SurfaceRuleKind {
	return ConditionalRuleKind
}

func (r ConditionalRule) String() string {
	return fmt.Sprintf("surfaceRule(%s)", r.Type())
}

func (r ConditionalRule) MarshalJSON() ([]byte, error) {
	if compound, ok := r.Condition.(*surface_conditions.CompoundCondition); ok {
		// We need to split out the conditions into a deep nest
		rules := make([]SurfaceRule, len(compound.Conditions))
		for i, condition := range compound.Conditions {
			action := r.Action
			if i > 0 {
				action = rules[i-1]
			}
			rules[i] = &ConditionalRule{Condition: condition.Condition, Action: action, Negate: condition.Negate}
		}
		return json.Marshal(rules[len(rules)-1])
	}

	if r.Negate {
		r.Negate = false
		condition := r.Condition
		r.Condition = surface_conditions.InvertCondition(condition)

	}

	return json.Marshal(struct {
		Type    SurfaceRuleKind                     `json:"type"`
		IfTrue  surface_conditions.SurfaceCondition `json:"if_true"`
		ThenRun SurfaceRule                         `json:"then_run"`
	}{
		Type:    ConditionalRuleKind,
		IfTrue:  r.Condition,
		ThenRun: r.Action,
	})

}
