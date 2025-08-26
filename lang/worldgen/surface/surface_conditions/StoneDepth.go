package surface_conditions

import (
	"encoding/json"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

func NewStoneDepthCondition(ctx *grammars.SurfaceCondition_StoneDepthContext) (*StoneDepthCondition, error) {
	offset, err := strconv.Atoi(ctx.Int(0).GetText())
	if err != nil {
		return nil, err
	}
	secondary_depth_range, err := strconv.Atoi(ctx.Int(1).GetText())
	if err != nil {
		return nil, err
	}
	surface := "floor"
	if ctx.Keyword_Ceiling() != nil {
		surface = "ceiling"
	}
	return &StoneDepthCondition{
		Depth:   offset,
		Add:     ctx.Keyword_Add() != nil, // if this is nil, then Keyword_Sub MUST exist per the grammar
		Range:   secondary_depth_range,
		Surface: surface,
	}, nil
}

type StoneDepthCondition struct {
	BaseCondition
	Depth   int
	Add     bool
	Range   int
	Surface string
}

func (c StoneDepthCondition) Type() ConditionType {
	return StoneDepthConditionType
}

func (c StoneDepthCondition) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(struct {
		Type    ConditionType `json:"type"`
		Depth   int           `json:"offset"`
		Surface string        `json:"surface_type"`
		Add     bool          `json:"add_surface_depth"`
		Range   int           `json:"secondary_depth_range"`
	}{
		Type:    StoneDepthConditionType,
		Depth:   c.Depth,
		Surface: c.Surface,
		Add:     c.Add,
		Range:   c.Range,
	})
	if err != nil {
		return nil, err
	}
	return json, nil
}
