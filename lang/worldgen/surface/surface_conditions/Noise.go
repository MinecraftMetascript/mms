package surface_conditions

import (
	"encoding/json"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lib"
)

func NewNoiseCondition(ctx *grammars.SurfaceCondition_NoiseContext) (*NoiseCondition, error) {
	min, err := strconv.ParseFloat(ctx.Number(0).GetText(), 64)
	if err != nil {
		return nil, err
	}
	max, err := strconv.ParseFloat(ctx.Number(1).GetText(), 64)
	if err != nil {
		return nil, err
	}
	return &NoiseCondition{
		NoiseRef: lib.ParseReferential("minecraft", ctx.ResourceReference()),
		Min:      min,
		Max:      max,
	}, nil
}

type NoiseCondition struct {
	BaseCondition
	NoiseRef lib.Reference
	Min      float64
	Max      float64
}

func (c NoiseCondition) Type() ConditionType {
	return NoiseConditionType
}

func (c NoiseCondition) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(struct {
		Type     ConditionType `json:"type"`
		NoiseRef lib.Reference `json:"noise_threshold"`
		Min      float64       `json:"min_threshold"`
		Max      float64       `json:"max_threshold"`
	}{
		Type:     NoiseConditionType,
		NoiseRef: c.NoiseRef,
		Min:      c.Min,
		Max:      c.Max,
	})
	if err != nil {
		return nil, err
	}
	return json, nil
}
