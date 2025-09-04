package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	waterBuildChain := builder_chain.NewBuilderChain[AboveWaterCondition](
		builder_chain.Build(
			func(ctx *grammar.SharedBuilder_OffsetContext, out *AboveWaterCondition, scope *traversal.Scope, _ string) {
				builder_chain.Builder_GetInt(ctx, func(v int) { out.Offset = v }, scope, "Offset")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.SharedBuilder_AddContext, out *AboveWaterCondition, _ *traversal.Scope, _ string) {
				builder_chain.SharedBuilder_Add(ctx, func(v bool) { out.Add = v })
			},
		),
		builder_chain.Build(
			func(ctx *grammar.SharedBuilder_MulContext, out *AboveWaterCondition, scope *traversal.Scope, _ string) {
				builder_chain.Builder_GetFloat(
					ctx,
					func(v float64) { out.DepthMultiplier = v },
					scope,
					"Multiplier",
				)
			},
		),
	)

	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_AboveWaterContext](),
		func(ctx antlr.ParserRuleContext, namespace string, scope *traversal.Scope) traversal.Construct {
			aboveWater := ctx.(*grammar.SurfaceCondition_AboveWaterContext)
			out := &AboveWaterCondition{}
			for _, r := range aboveWater.AllSurfaceCondition_AboveWaterBuilder() {
				builder_chain.Invoke(waterBuildChain, r, out, scope, namespace)
			}

			return out
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
