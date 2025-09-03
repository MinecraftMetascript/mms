package surface_rules

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_AboveWaterContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			aboveWater := ctx.(*grammar.SurfaceCondition_AboveWaterContext)
			offset := 0
			if aboveWater.Int() != nil {
				if v, err := strconv.Atoi(aboveWater.Int().GetText()); err == nil {
					offset = v
				}
			}
			depthMultiplier := 0.0
			if aboveWater.Number() != nil {
				if v, err := strconv.ParseFloat(aboveWater.Number().GetText(), 64); err == nil {
					depthMultiplier = v
				}
			}
			return &AboveWaterCondition{
				Offset:          offset,
				DepthMultiplier: depthMultiplier,
				Add:             aboveWater.Keyword_Add() != nil,
			}

		},
	)
}

type AboveWaterCondition struct {
	traversal.Construct
	Offset          int
	DepthMultiplier float64
	Add             bool
}

func (c AboveWaterCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c AboveWaterCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type       SurfaceConditionKind `json:"type"`
		Offset     int                  `json:"offset"`
		Multiplier float64              `json:"surface_depth_multiplier"`
		Add        bool                 `json:"add_stone_depth"`
	}{
		Type:       AboveWaterConditionKind,
		Offset:     c.Offset,
		Multiplier: c.DepthMultiplier,
		Add:        c.Add,
	}, "", "  ")
}
