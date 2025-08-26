package surface

import (
	"errors"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_conditions"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_rules"
	"github.com/minecraftmetascript/mms/lib"
)

func ReplaceConditionReferences(
	condition surface_conditions.SurfaceCondition,
	conditions map[string]map[string]lib.Symbol[surface_conditions.SurfaceCondition],
	rules map[string]map[string]lib.Symbol[surface_rules.SurfaceRule],
	defaultNamespace string,
) (surface_conditions.SurfaceCondition, []error) {
	switch condition := condition.(type) {
	case *surface_conditions.CompoundCondition:
		errs := make([]error, 0)
		for i, child := range condition.Conditions {
			next, err := ReplaceConditionReferences(child.Condition, conditions, rules, defaultNamespace)
			if err != nil {
				errs = append(errs, err...)
				continue
			}
			child.Condition = next
			condition.Conditions[i] = child
		}
		if len(errs) > 0 {
			return condition, errs
		}
		return condition, nil
	case *surface_conditions.ConditionReference:
		namespace := condition.Namespace
		if namespace == "" {
			if condition.Reference.Namespace != "" {
				namespace = condition.Reference.Namespace
			} else {
				namespace = defaultNamespace
			}
		}
		if namespaceConditions, ok := conditions[namespace]; ok {
			if r, ok := namespaceConditions[condition.Name]; ok {
				return r.Value, nil
			}
		}

		return nil, []error{errors.New(fmt.Sprintf("unknown condition reference: \"%s:%s\"", condition.Namespace, condition.Name))}

	default:
		return condition, nil
	}
}

func ReplaceRuleReferences(
	rule surface_rules.SurfaceRule,
	conditions map[string]map[string]lib.Symbol[surface_conditions.SurfaceCondition],
	rules map[string]map[string]lib.Symbol[surface_rules.SurfaceRule],
	defaultNamespace string,
) (surface_rules.SurfaceRule, []error) {
	switch rule := rule.(type) {
	case *surface_rules.BlockRule:
		return rule, nil
	case *surface_rules.BandlandsRule:
		return rule, nil
	case *surface_rules.SequenceRule:
		errs := make([]error, 0)
		for i, r := range rule.Rules {
			if r.Type() == surface_rules.ReferenceRuleKind {
				resolvedRule, err := ReplaceRuleReferences(
					r,
					conditions, rules, defaultNamespace,
				)
				if err != nil {
					errs = append(errs, err...)
					continue
				}
				rule.Rules[i] = resolvedRule
			}
		}
		return rule, nil
	case *surface_rules.ConditionalRule:
		resolvedCondition, err := ReplaceConditionReferences(
			rule.Condition,
			conditions, rules, defaultNamespace,
		)
		if err != nil {
			return nil, err
		}
		rule.Condition = resolvedCondition

		resolvedAction, err := ReplaceRuleReferences(
			rule.Action,
			conditions, rules, defaultNamespace,
		)
		if err != nil {
			return nil, err
		}
		rule.Action = resolvedAction

		return rule, nil
	case *surface_rules.ReferenceRule:
		namespace := rule.Namespace
		if namespace == "" {
			if rule.Reference.Namespace != "" {
				namespace = rule.Reference.Namespace
			} else {
				namespace = defaultNamespace
			}
		}
		if rules, ok := rules[namespace]; ok {
			if r, ok := rules[rule.Name]; ok {
				return r.Value, nil
			}
		}
		return nil, []error{errors.New(fmt.Sprintf("unknown rule reference: \"%s\"", rule.Name))}
	}
	return nil, []error{errors.New("unknown surface rule type")}
}
