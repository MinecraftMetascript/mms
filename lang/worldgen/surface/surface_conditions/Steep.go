package surface_conditions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewSteepCondition(ctx *grammars.SurfaceCondition_SteepContext) (*SteepCondition, error) {
	return &SteepCondition{}, nil
}

type SteepCondition struct {
	BaseCondition
}

func (c SteepCondition) Type() ConditionType {
	return SteepConditionType
}

func (c SteepCondition) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(struct {
		Type ConditionType `json:"type"`
	}{
		Type: SteepConditionType,
	})
	if err != nil {
		return nil, err
	}
	return json, nil
}
