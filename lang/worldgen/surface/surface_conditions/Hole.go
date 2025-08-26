package surface_conditions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewHoleCondition(ctx *grammars.SurfaceCondition_HoleContext) (*HoleCondition, error) {
	return &HoleCondition{}, nil
}

type HoleCondition struct {
	BaseCondition
}

func (c HoleCondition) Type() ConditionType {
	return HoleConditionType
}

func (c HoleCondition) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(struct {
		Type ConditionType `json:"type"`
	}{
		Type: HoleConditionType,
	})
	if err != nil {
		return nil, err
	}
	return json, nil
}
