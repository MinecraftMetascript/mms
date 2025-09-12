package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.RegisterHelp[*grammar.DensityFn_NoiseContext](func(construct traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
		out := "References a defined noise function.\nUses .XzScale( densityFn ) and .YScale( densityFn ) to scale the noise function."
		return &out
	})
	traversal.RegisterHelp[*grammar.DensityFn_InlineNoiseContext](func(construct traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
		out := "Defines a noise inline. This will automatically be extracted the mms_inline namespace.<br/>  Uses `.XzScale(_densityFn_)` and `.YScale(_densityFn_)` to scale the noise function."
		return &out
	})
	traversal.Register(
		func(densityFn *grammar.DensityFn_NoiseContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			noiseFnBuilder := builder_chain.NewBuilderChain(
				builder_chain.Build(
					func(ctx *grammar.Builder_XZScaleContext, target *NoiseDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.XzScale = v }, scope, "XzScale")
					},
				), builder_chain.Build(
					func(ctx *grammar.Builder_YScaleContext, target *NoiseDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.YScale = v }, scope, "XzScale")
					},
				),
			)

			out := &NoiseDensityFn{
				XzScale: 1,
				YScale:  1,
			}

			if noiseInlineCtx := densityFn.DensityFn_InlineNoise(); noiseInlineCtx != nil {
				noiseRef := traversal.ConstructRegistry.Construct(noiseInlineCtx, currentNamespace, scope)
				if noiseRef == nil {
					// TODO: Diagnose
				} else if ref, ok := noiseRef.(*traversal.Reference); ok {
					out.Noise = *ref
				}
			}

			if refCtx := densityFn.ResourceReference(); refCtx != nil {
				cons := traversal.ConstructRegistry.Construct(refCtx, currentNamespace, scope)
				if r, ok := cons.(*traversal.Reference); ok {
					out.Noise = *r
				}
			}

			for _, r := range densityFn.AllDensityFn_NoiseBuilder() {
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

type NoiseDensityFn struct {
	Noise   traversal.Reference
	XzScale float64
	YScale  float64
}

func (c NoiseDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type    string  `json:"type"`
		Noise   string  `json:"noise"`
		XzScale float64 `json:"xz_scale"`
		YScale  float64 `json:"y_scale"`
	}{
		Type:    "minecraft:noise",
		Noise:   c.Noise.String(),
		XzScale: c.XzScale,
		YScale:  c.YScale,
	}, "", "  ")
}

func (c NoiseDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
