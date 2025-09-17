package density_functions

import (
	"encoding/json"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type RangeChoiceDensityFnFactory struct {
}

func (r RangeChoiceDensityFnFactory) Create(ctx *grammar.DensityFn_RangeChoiceContext, namespace string, scope *traversal.Scope) *RangeChoiceDensityFn {
	rangeChoiceBuilder := builder_chain.NewBuilderChain[RangeChoiceDensityFn](
		builder_chain.Build(
			func(ctx *grammar.Builder_MinContext, target *RangeChoiceDensityFn, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetFloat(ctx, func(v float64) { target.Min = v }, scope, "Min")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_MaxContext, target *RangeChoiceDensityFn, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetFloat(ctx, func(v float64) { target.Max = v }, scope, "Max")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_InRangeContext, target *RangeChoiceDensityFn, scope *traversal.Scope, namespace string) {
				fn := traversal.ConstructRegistry.Construct(ctx.DensityFn(), namespace, scope)
				if fn != nil {
					target.InRange = fn
				} else {
					scope.DiagnoseSemanticError("Missing or invalid InRange Function", ctx)
				}
			},
		),
		builder_chain.Build(
			func(ctx *grammar.Builder_OutRangeContext, target *RangeChoiceDensityFn, scope *traversal.Scope, namespace string) {
				fn := traversal.ConstructRegistry.Construct(ctx.DensityFn(), namespace, scope)
				if fn != nil {
					target.OutRange = fn
				} else {
					scope.DiagnoseSemanticError("Missing or invalid InRange Function", ctx)
				}
			},
		),
	)

	out := &RangeChoiceDensityFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}
	input := ctx.DensityFn()
	if input == nil {
		scope.DiagnoseSemanticError("Missing input to range choice", ctx)
	} else {
		out.Input = traversal.ConstructRegistry.Construct(input.(antlr.ParserRuleContext), namespace, scope)
	}

	for _, builderWrap := range ctx.AllDensityFn_RangeChoiceBuilder() {
		builder_chain.Invoke(rangeChoiceBuilder, builderWrap.GetChild(0).(antlr.ParserRuleContext), out, scope, namespace)
	}

	builder_chain.Require(
		rangeChoiceBuilder,
		ctx,
		scope,
		reflect.TypeFor[*grammar.Builder_MinContext](),
		".Min",
	)
	builder_chain.Require(
		rangeChoiceBuilder,
		ctx,
		scope,
		reflect.TypeFor[*grammar.Builder_MaxContext](),
		".Max",
	)
	builder_chain.Require(
		rangeChoiceBuilder,
		ctx,
		scope,
		reflect.TypeFor[*grammar.Builder_InRangeContext](),
		".InRange",
	)
	builder_chain.Require(
		rangeChoiceBuilder,
		ctx,
		scope,
		reflect.TypeFor[*grammar.Builder_OutRangeContext](),
		".OutRange",
	)

	return out
}

func (r RangeChoiceDensityFnFactory) GetHelp(node *RangeChoiceDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Selects from 2 density functions based on an input function and a range.",
		Position: node.GetLocation(),
	}
}

func init() {
	traversal.RegisterNodeFactory(RangeChoiceDensityFnFactory{})
}

type RangeChoiceDensityFn struct {
	Input    traversal.Construct
	Min      float64
	Max      float64
	InRange  traversal.Construct
	OutRange traversal.Construct
	location traversal.TextLocation
}

func (c RangeChoiceDensityFn) GetLocation() traversal.TextLocation {
	return c.location
}

func (c RangeChoiceDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     string              `json:"type"`
		Input    traversal.Construct `json:"input"`
		Min      float64             `json:"min_inclusive"`
		Max      float64             `json:"max_inclusive"`
		InRange  traversal.Construct `json:"when_in_range"`
		OutRange traversal.Construct `json:"when_out_of_range"`
	}{
		Type:     "minecraft:range_choice",
		Input:    c.Input,
		Min:      c.Min,
		Max:      c.Max,
		InRange:  c.InRange,
		OutRange: c.OutRange,
	}, "", "  ")
}
