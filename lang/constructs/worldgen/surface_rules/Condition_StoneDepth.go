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
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_StoneDepthContext](),
		func(ctx antlr.ParserRuleContext, namespace string, scope *traversal.Scope) traversal.Construct {
			stoneDepthBuildChain := builder_chain.NewBuilderChain[StoneDepthCondition](
				builder_chain.Build(
					func(ctx *grammar.SurfaceCondition_StoneDepthBuilder_SecondaryDepthRangeContext, target *StoneDepthCondition, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetInt(ctx, func(v int) { target.Range = v }, scope, "Secondary Depth Range")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_OffsetContext, target *StoneDepthCondition, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetInt(ctx, func(v int) { target.Offset = v }, scope, "Offset")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_AddContext, target *StoneDepthCondition, scope *traversal.Scope, namespace string) {
						builder_chain.SharedBuilder_Add(ctx, func(v bool) { target.Add = v })
					},
				),
			)

			stoneDepth := ctx.(*grammar.SurfaceCondition_StoneDepthContext)
			out := &StoneDepthCondition{}
			if mode := stoneDepth.StoneDepthMode(); mode != nil {
				out.Surface = mode.GetText()
			} else {
				scope.DiagnoseSemanticError("Missing stone depth mode, expected Floor or Ceiling", ctx)
			}

			for _, r := range stoneDepth.AllSurfaceCondition_StoneDepthBuilder() {
				child := r.GetChild(0)
				if child == nil {
					continue
				}
				builder_chain.Invoke(stoneDepthBuildChain, child.(antlr.ParserRuleContext), out, scope, namespace)
			}

			return out
		},
	)
}

type StoneDepthCondition struct {
	traversal.Construct
	Offset  int
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
		Offset  int                  `json:"offset"`
		Surface string               `json:"surface_type"`
		Add     bool                 `json:"add_surface_depth"`
		Range   int                  `json:"secondary_depth_range"`
	}{
		Type:    StoneDepthConditionKind,
		Offset:  c.Offset,
		Surface: c.Surface,
		Add:     c.Add,
		Range:   c.Range,
	}, "", "  ")
}
