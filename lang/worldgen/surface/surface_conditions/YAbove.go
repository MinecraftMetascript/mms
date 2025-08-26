package surface_conditions

import (
	"encoding/json"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lib"
)

type YAboveCondition struct {
	BaseCondition
	Anchor     lib.VerticalAnchor
	Multiplier int
	Add        bool
}

func NewYAboveCondition(ctx *grammars.SurfaceCondition_YAboveContext) (*YAboveCondition, error) {
	anchor, err := lib.ParseVerticalAnchor(ctx.VerticalAnchor())
	if err != nil {
		return nil, err
	}
	multiplier, _ := strconv.Atoi(ctx.Int().GetText())
	return &YAboveCondition{
		Anchor:     *anchor,
		Multiplier: multiplier,
		Add:        ctx.Keyword_Add() != nil,
	}, nil
}

func (c YAboveCondition) Type() ConditionType {
	return YAboveConditionType
}

func (c YAboveCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       ConditionType      `json:"type"`
		Anchor     lib.VerticalAnchor `json:"anchor"`
		Multiplier int                `json:"surface_depth_multiplier"`
		Add        bool               `json:"add_stone_depth"`
	}{
		Type:       YAboveConditionType,
		Anchor:     c.Anchor,
		Multiplier: c.Multiplier,
		Add:        c.Add,
	})
}
