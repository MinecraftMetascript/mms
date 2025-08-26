package surface_conditions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lib"
)

func NewVerticalGradientCondition(ctx *grammars.SurfaceCondition_VerticalGradientContext) (*VerticalGradientCondition, error) {
	trueAnchor, err := lib.ParseVerticalAnchor(ctx.VerticalAnchor(0))
	if err != nil {
		return nil, err
	}
	falseAnchor, err := lib.ParseVerticalAnchor(ctx.VerticalAnchor(1))
	if err != nil {
		return nil, err
	}
	return &VerticalGradientCondition{
		SeedText:        ctx.String_().GetText(),
		TrueAtAndBelow:  *trueAnchor,
		FalseAtAndAbove: *falseAnchor,
	}, nil
}

type VerticalGradientCondition struct {
	BaseCondition
	SeedText        string
	TrueAtAndBelow  lib.VerticalAnchor
	FalseAtAndAbove lib.VerticalAnchor
}

func (c VerticalGradientCondition) Type() ConditionType {
	return VerticalGradientConditionType
}

func (c VerticalGradientCondition) MarshalJSON() ([]byte, error) {
	escapedSeed := ""
	json.Unmarshal([]byte(c.SeedText), &escapedSeed)
	json, err := json.Marshal(struct {
		Type            ConditionType      `json:"type"`
		SeedText        string             `json:"random_name"`
		TrueAtAndBelow  lib.VerticalAnchor `json:"true_at_and_below"`
		FalseAtAndAbove lib.VerticalAnchor `json:"false_at_and_above"`
	}{
		Type:            VerticalGradientConditionType,
		SeedText:        escapedSeed,
		TrueAtAndBelow:  c.TrueAtAndBelow,
		FalseAtAndAbove: c.FalseAtAndAbove,
	})
	if err != nil {
		return nil, err
	}
	return json, nil
}
