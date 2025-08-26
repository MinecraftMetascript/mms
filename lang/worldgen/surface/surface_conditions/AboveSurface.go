package surface_conditions

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewAboveSurfaceCondition(ctx *grammars.SurfaceCondition_AboveSurfaceContext) (*AboveSurfaceCondition, error) {
	return &AboveSurfaceCondition{}, nil
}

type AboveSurfaceCondition struct {
	BaseCondition
}

func (c AboveSurfaceCondition) Type() ConditionType { return AboveSurfaceConditionType }

func (c AboveSurfaceCondition) String() string {
	return fmt.Sprintf("surfaceCondition(%s)", c.Type())
}

func (c AboveSurfaceCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type ConditionType `json:"type"`
	}{
		Type: c.Type(),
	})
}
