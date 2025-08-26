package surface_conditions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewFreezingCondition(ctx *grammars.SurfaceCondition_FreezingContext) (*FreezingCondition, error) {
	return &FreezingCondition{}, nil
}

type FreezingCondition struct {
	BaseCondition
}

func (c FreezingCondition) Type() ConditionType {
	return FreezingConditionType
}

func (c FreezingCondition) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(struct {
		Type ConditionType `json:"type"`
	}{
		Type: FreezingConditionType,
	})
	if err != nil {
		return nil, err
	}
	return json, nil
}
