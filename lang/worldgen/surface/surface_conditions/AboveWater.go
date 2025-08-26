package surface_conditions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewAboveWaterCondition(ctx *grammars.SurfaceCondition_AboveWaterContext) (*AboveWaterCondition, error) {
	offset, err := strconv.Atoi(ctx.Int().GetText())
	if err != nil {
		return nil, err
	}
	depthMultiplier, err := strconv.ParseFloat(ctx.Number().GetText(), 64)
	if err != nil {
		return nil, err
	}
	return &AboveWaterCondition{
		Offset:          offset,
		DepthMultiplier: depthMultiplier,
		Add:             ctx.Keyword_Add() != nil,
	}, nil
}

type AboveWaterCondition struct {
	BaseCondition
	Offset          int
	DepthMultiplier float64
	Add             bool
}

func (c AboveWaterCondition) Type() ConditionType { return AboveWaterConditionType }

func (c AboveWaterCondition) String() string {
	return fmt.Sprintf("surfaceCondition(%s)", c.Type())
}

func (c AboveWaterCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       ConditionType `json:"type"`
		Offset     int           `json:"offset"`
		Multiplier float64       `json:"surface_depth_multiplier"`
		Add        bool          `json:"add_stone_depth"`
	}{
		Type:       c.Type(),
		Offset:     c.Offset,
		Multiplier: c.DepthMultiplier,
		Add:        c.Add,
	})
}
