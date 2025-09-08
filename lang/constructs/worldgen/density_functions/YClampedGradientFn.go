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
		reflect.TypeFor[*grammar.DensityFn_YClampedGradientContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			rangeChoiceBuilder := builder_chain.NewBuilderChain[YClampedGradientDensityFn](
				builder_chain.Build(
					func(ctx *grammar.Builder_BottomLiteralContext, target *YClampedGradientDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetInt(ctx, func(v int) { target.Bottom = v }, scope, "Bottom")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_TopLiteralContext, target *YClampedGradientDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetInt(ctx, func(v int) { target.Top = v }, scope, "Top")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_MinContext, target *YClampedGradientDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.Min = v }, scope, "Min")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_MaxContext, target *YClampedGradientDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.Max = v }, scope, "Max")
					},
				),
			)

			densityFn := ctx.(*grammar.DensityFn_YClampedGradientContext)
			out := &YClampedGradientDensityFn{}

			for _, builderWrap := range densityFn.AllDensityFn_YClampedGradientBuilder() {
				builder_chain.Invoke(rangeChoiceBuilder, builderWrap.GetChild(0).(antlr.ParserRuleContext), out, scope, currentNamespace)
			}

			builder_chain.Require(
				rangeChoiceBuilder,
				densityFn,
				scope,
				reflect.TypeFor[*grammar.Builder_MinContext](),
				".Min",
			)
			builder_chain.Require(
				rangeChoiceBuilder,
				densityFn,
				scope,
				reflect.TypeFor[*grammar.Builder_MaxContext](),
				".Max",
			)
			builder_chain.Require(
				rangeChoiceBuilder,
				densityFn,
				scope,
				reflect.TypeFor[*grammar.Builder_BottomLiteralContext](),
				".Bottom",
			)
			builder_chain.Require(
				rangeChoiceBuilder,
				densityFn,
				scope,
				reflect.TypeFor[*grammar.Builder_TopLiteralContext](),
				".Top",
			)
			return out
		},
	)
}

type YClampedGradientDensityFn struct {
	Top    int
	Bottom int
	Min    float64
	Max    float64
}

func (c YClampedGradientDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type   string  `json:"type"`
		Min    float64 `json:"from_value"`
		Max    float64 `json:"to_value"`
		Bottom int     `json:"from_y"`
		Top    int     `json:"to_y"`
	}{
		Type:   "minecraft:y_clamped_gradient",
		Min:    c.Min,
		Max:    c.Max,
		Top:    c.Top,
		Bottom: c.Bottom,
	}, "", "  ")
}

func (c YClampedGradientDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
