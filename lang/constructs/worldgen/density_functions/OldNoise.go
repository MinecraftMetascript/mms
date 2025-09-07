package density_functions

import (
	"encoding/json"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_OldBlendedNoiseContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			noiseFnBuilder := builder_chain.NewBuilderChain(
				builder_chain.Build(
					func(ctx *grammar.SharedBuilder_XzScaleContext, target *OldBlendedNoiseFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.XzScale = v }, scope, "XzScale")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.SharedBuilder_YScaleContext, target *OldBlendedNoiseFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.YScale = v }, scope, "XzScale")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.SharedBuilder_XzFactorContext, target *OldBlendedNoiseFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.XzFactor = v }, scope, "XzScale")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.SharedBuilder_YFactorContext, target *OldBlendedNoiseFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.YFactor = v }, scope, "XzScale")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.DensityFn_OldBlendedNoiseBuilder_SmearContext, target *OldBlendedNoiseFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.SmearMultiplier = v }, scope, "XzScale")
					},
				),
			)

			out := &OldBlendedNoiseFn{
				XzScale:         1,
				YScale:          1,
				XzFactor:        1,
				YFactor:         1,
				SmearMultiplier: 1,
			}

			densityFn := ctx.(*grammar.DensityFn_OldBlendedNoiseContext)

			for _, r := range densityFn.AllDensityFn_OldBlendedNoiseBuilder() {
				child := r.GetChild(0)
				if child == nil {
					continue
				}
				builder_chain.Invoke(noiseFnBuilder, child.(antlr.ParserRuleContext), out, scope, currentNamespace)
			}
			return out
		},
	)
}

type OldBlendedNoiseFn struct {
	XzScale         float64
	YScale          float64
	XzFactor        float64
	YFactor         float64
	SmearMultiplier float64
}

func (c OldBlendedNoiseFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     string  `json:"type"`
		XzScale  float64 `json:"xz_scale"`
		YScale   float64 `json:"y_scale"`
		XzFactor float64 `json:"xz_factor"`
		YFactor  float64 `json:"y_factor"`
		Smear    float64 `json:"smear_scale_multiplier"`
	}{
		Type:     "minecraft:noise",
		XzScale:  c.XzScale,
		YScale:   c.YScale,
		XzFactor: c.XzFactor,
		YFactor:  c.YFactor,
		Smear:    c.SmearMultiplier,
	}, "", "  ")
}

func (c OldBlendedNoiseFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
