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
	traversal.Register(
		func(densityFn *grammar.DensityFn_ClampContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			rangeChoiceBuilder := builder_chain.NewBuilderChain[ClampDensityFn](
				builder_chain.Build(
					func(ctx *grammar.Builder_MinContext, target *ClampDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.Min = v }, scope, "Min")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_MaxContext, target *ClampDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.Max = v }, scope, "Max")
					},
				),
			)
			out := &ClampDensityFn{}
			input := densityFn.DensityFn()
			if input == nil {
				scope.DiagnoseSemanticError("Missing input to range choice", densityFn)
			}
			out.Input = traversal.ConstructRegistry.Construct(input.(antlr.ParserRuleContext), currentNamespace, scope)
			if out.Input == nil {
				scope.DiagnoseSemanticError("Invalid input to range choice", densityFn)
			}

			for _, builderWrap := range densityFn.AllDensityFn_ClampBuilder() {
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
			return out
		},
	)
}

type ClampDensityFn struct {
	Input traversal.Construct
	Min   float64
	Max   float64
}

func (c ClampDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type  string              `json:"type"`
		Input traversal.Construct `json:"input"`
		Min   float64             `json:"min"`
		Max   float64             `json:"max"`
	}{
		Type:  "minecraft:clamp",
		Input: c.Input,
		Min:   c.Min,
		Max:   c.Max,
	}, "", "  ")
}

func (c ClampDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
