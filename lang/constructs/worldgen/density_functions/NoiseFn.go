package density_functions

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_NoiseContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			noiseFnBuilder := builder_chain.NewBuilderChain(
				builder_chain.Build(
					func(ctx *grammar.SharedBuilder_XzScaleContext, target *NoiseDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.XzScale = v }, scope, "XzScale")
					},
				), builder_chain.Build(
					func(ctx *grammar.SharedBuilder_YScaleContext, target *NoiseDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.YScale = v }, scope, "XzScale")
					},
				),
			)

			densityFn := ctx.(*grammar.DensityFn_NoiseContext)

			out := &NoiseDensityFn{
				XzScale: 1,
				YScale:  1,
			}

			if noiseInlineCtx := densityFn.Noise(); noiseInlineCtx != nil {
				noiseDef := traversal.ConstructRegistry.Construct(noiseInlineCtx, currentNamespace, scope)
				noiseRef := traversal.NewReference(
					fmt.Sprintf("densityfn_noise_%d_%d", noiseInlineCtx.GetStart().GetLine(), noiseInlineCtx.GetStart().GetColumn()),
					"mms_inline",
				)
				noiseSymbol := traversal.NewSymbol(
					traversal.RuleLocation(noiseInlineCtx, scope.CurrentFile),
					traversal.RuleLocation(noiseInlineCtx, scope.CurrentFile),
					noiseDef,
					noiseRef,
					"Noise",
				)
				if _, ok := scope.Get(*noiseRef); !ok {
					if err := scope.Register(noiseSymbol); err != nil {
						// TODO: Improve behavior
						panic(err)
					}
				}
				out.Noise = *noiseRef
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
