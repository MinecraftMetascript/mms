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
		reflect.TypeFor[grammar.SurfaceCondition_NoiseThresholdContext](),
		func(ctx antlr.ParserRuleContext, namespace string, scope *traversal.Scope) traversal.Construct {
			noiseBuildChain := builder_chain.NewBuilderChain[NoiseCondition](
				builder_chain.Build(
					func(ctx *grammar.SurfaceCondition_NoiseThresholdBuilder_MinContext, out *NoiseCondition, scope *traversal.Scope, _ string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { out.Min = v }, scope, "Min")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.SurfaceCondition_NoiseThresholdBuilder_MaxContext, out *NoiseCondition, scope *traversal.Scope, _ string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { out.Max = v }, scope, "Max")
					},
				),
			)

			noise := ctx.(*grammar.SurfaceCondition_NoiseThresholdContext)
			var noiseRef traversal.Reference
			if rr := noise.ResourceReference(); rr != nil {
				if cons := traversal.ConstructRegistry.Construct(rr, namespace, scope); cons != nil {
					if r, ok := cons.(*traversal.Reference); ok && r != nil {
						noiseRef = *r
					}
				}
			}
			out := &NoiseCondition{
				NoiseRef: noiseRef,
			}
			for _, r := range noise.AllSurfaceCondition_NoiseThresholdBuilder() {
				child := r.GetChild(0)
				if child == nil {
					continue
				}
				builder_chain.Invoke(noiseBuildChain, child.(antlr.ParserRuleContext), out, scope, namespace)
			}

			builder_chain.Require(
				noiseBuildChain,
				noise,
				scope,
				reflect.TypeFor[*grammar.SurfaceCondition_NoiseThresholdBuilder_MinContext](),
				".Min",
			)
			builder_chain.Require(
				noiseBuildChain,
				noise,
				scope,
				reflect.TypeFor[*grammar.SurfaceCondition_NoiseThresholdBuilder_MaxContext](),
				".Max",
			)

			return out
		},
	)
}

type NoiseCondition struct {
	traversal.Construct
	NoiseRef traversal.Reference
	Min      float64
	Max      float64
}

func (c NoiseCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c NoiseCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     SurfaceConditionKind `json:"type"`
		NoiseRef traversal.Reference  `json:"noise_threshold"`
		Min      float64              `json:"min_threshold"`
		Max      float64              `json:"max_threshold"`
	}{
		Type:     NoiseConditionKind,
		NoiseRef: c.NoiseRef,
		Min:      c.Min,
		Max:      c.Max,
	}, "", "  ")
}
