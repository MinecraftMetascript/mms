package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type NoiseFnFactory struct {
	BaseDensityFnFactory
}

func (n NoiseFnFactory) Create(ctx *grammar.DensityFn_NoiseContext, namespace string, scope *traversal.Scope) *NoiseDensityFn {
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
		XzScale:  1,
		YScale:   1,
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}

	if noiseInlineCtx := ctx.DensityFn_InlineNoise(); noiseInlineCtx != nil {
		noiseRef := traversal.ConstructRegistry.Construct(noiseInlineCtx, namespace, scope)
		if noiseRef == nil {
			// TODO: Diagnose
		} else if ref, ok := noiseRef.(*traversal.Reference); ok {
			out.Noise = *ref
		}
	}

	if refCtx := ctx.ResourceReference(); refCtx != nil {
		cons := traversal.ConstructRegistry.Construct(refCtx, namespace, scope)
		if r, ok := cons.(*traversal.Reference); ok {
			out.Noise = *r
		}
	}

	for _, r := range ctx.AllDensityFn_NoiseBuilder() {
		child := r.GetChild(0)
		if child == nil {
			continue
		}
		builder_chain.Invoke(noiseFnBuilder, child.(antlr.ParserRuleContext), out, scope, namespace)
	}

	return out

}

func (n NoiseFnFactory) GetHelp(node *NoiseDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Uses a noise function to determine density.<br/>Uses .XzScale( densityFn ) and .YScale( densityFn ) to scale the noise function.",
		Position: node.GetLocation(),
	}
}

func init() {
	traversal.RegisterNodeFactory(NoiseFnFactory{}, false)
}

type NoiseDensityFn struct {
	Noise    traversal.Reference
	XzScale  float64
	YScale   float64
	location traversal.TextLocation
}

func (c NoiseDensityFn) GetLocation() traversal.TextLocation {
	return c.location
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
