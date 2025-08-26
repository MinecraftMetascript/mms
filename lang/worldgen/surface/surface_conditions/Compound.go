package surface_conditions

import (
	"encoding/json"
	"slices"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewCompoundCondition(ctx *grammars.SurfaceCondition_CompoundContext, factory ConditionFactory) (SurfaceCondition, error) {
	out := &CompoundCondition{}
	items := ctx.AllSurfaceCondition_Compound__Item()
	slices.Reverse(items)

	for _, item := range items {
		negated := item.Bang() != nil
		condition, err := factory.ConstructSurfaceCondition(item.SurfaceCondition().(*grammars.SurfaceConditionContext))
		if err != nil {
			return nil, err
		}
		out.Conditions = append(out.Conditions, CompoundConditionItem{Condition: condition, Negate: negated})
	}

	return out, nil
}

type CompoundCondition struct {
	BaseCondition
	Conditions []CompoundConditionItem
}

func (c CompoundCondition) Type() ConditionType {
	return CompoundConditionType
}

func (c CompoundCondition) MarshalJSON() ([]byte, error) {
	if len(c.Conditions) == 1 {
		return json.Marshal(c.Conditions[0])
	}
	return json.Marshal(struct {
		Type       ConditionType           `json:"type"`
		Conditions []CompoundConditionItem `json:"conditions"`
	}{
		Type:       CompoundConditionType,
		Conditions: c.Conditions,
	})
}

type CompoundConditionItem struct {
	json.Marshaler
	Condition SurfaceCondition
	Negate    bool
}

func (c CompoundConditionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type      ConditionType    `json:"type"`
		Condition SurfaceCondition `json:"condition"`
		Negate    bool             `json:"negate"`
	}{
		Type:      CompoundConditionType,
		Condition: c.Condition,
		Negate:    c.Negate,
	})
}
