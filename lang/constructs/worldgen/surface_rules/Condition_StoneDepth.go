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
		reflect.TypeFor[grammar.SurfaceCondition_StoneDepthContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			stoneDepth := ctx.(*grammar.SurfaceCondition_StoneDepthContext)
			var offset int
			if stoneDepth.Int(0) != nil {
				if v, err := strconv.Atoi(stoneDepth.Int(0).GetText()); err == nil {
					offset = v
				}
			}
			var secondaryDepthRange int
			if stoneDepth.Int(1) != nil {
				if v, err := strconv.Atoi(stoneDepth.Int(1).GetText()); err == nil {
					secondaryDepthRange = v
				}
			}
			surface := "floor"
			if stoneDepth.Keyword_Ceiling() != nil {
				surface = "ceiling"
			}
			return &StoneDepthCondition{
				Depth:   offset,
				Add:     stoneDepth.Keyword_Add() != nil, // if this is nil, then Keyword_Sub MUST exist per the grammar
				Range:   secondaryDepthRange,
				Surface: surface,
			}
		},
	)
}

type StoneDepthCondition struct {
	traversal.Construct
	Depth   int
	Add     bool
	Range   int
	Surface string
}

func (c StoneDepthCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c StoneDepthCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type    SurfaceConditionKind `json:"type"`
		Depth   int                  `json:"offset"`
		Surface string               `json:"surface_type"`
		Add     bool                 `json:"add_surface_depth"`
		Range   int                  `json:"secondary_depth_range"`
	}{
		Type:    StoneDepthConditionKind,
		Depth:   c.Depth,
		Surface: c.Surface,
		Add:     c.Add,
		Range:   c.Range,
	}, "", "  ")
}
