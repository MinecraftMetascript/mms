package density_functions

import (
	"encoding/json"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"github.com/minecraftmetascript/mms/lsp/completions"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type YClampedGradientFnFactory struct {
}

func (y YClampedGradientFnFactory) Create(ctx *grammar.DensityFn_YClampedGradientContext, namespace string, scope *traversal.Scope) *YClampedGradientDensityFn {
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

	out := &YClampedGradientDensityFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}

	for _, builderWrap := range ctx.AllDensityFn_YClampedGradientBuilder() {
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
		reflect.TypeFor[*grammar.Builder_BottomLiteralContext](),
		".Bottom",
	)
	builder_chain.Require(
		rangeChoiceBuilder,
		ctx,
		scope,
		reflect.TypeFor[*grammar.Builder_TopLiteralContext](),
		".Top",
	)
	return out
}

func (y YClampedGradientFnFactory) GetHelp(node *YClampedGradientDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Creates a gradient that runs from a specified y-coordinate to another.<br/>.`Top(int)` and `.Max(float)` define the upper elevation and value, while `.Bottom(int)` and `.Min(float)` define the lower elevation and value.<br/>Example: `YClampedGradient().Top(100).Max(0).Bottom(-100).Max(1)`",
		Position: node.GetLocation(),
	}
}

func init() {
	traversal.RegisterNodeFactory(YClampedGradientFnFactory{})
}

type YClampedGradientDensityFn struct {
	Top    int
	Bottom int
	Min    float64
	Max    float64

	location traversal.TextLocation
}

func (c YClampedGradientDensityFn) GetLocation() traversal.TextLocation {
    return c.location
}

func (c YClampedGradientDensityFn) GetCompletions(cursorPosition protocol.Position) []protocol.CompletionItem {
    return []protocol.CompletionItem{
        completions.BuilderFnCompletion("Top", "Top(${1})", c.location.Stop.ToLspPosition()),
        completions.BuilderFnCompletion("Bottom", "Bottom(${1})", c.location.Stop.ToLspPosition()),
        completions.BuilderFnCompletion("Min", "Min(${1})", c.location.Stop.ToLspPosition()),
        completions.BuilderFnCompletion("Max", "Max(${1})", c.location.Stop.ToLspPosition()),
    }
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
