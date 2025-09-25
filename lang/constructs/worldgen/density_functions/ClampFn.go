package density_functions

import (
	"encoding/json"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lsp/completions"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ClampFnFactory struct {
}

func (c ClampFnFactory) Create(ctx *grammar.DensityFn_ClampContext, namespace string, scope *traversal.Scope) *ClampDensityFn {
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
	out := &ClampDensityFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
		ctx:      ctx,
	}
	input := ctx.DensityFn()
	if input == nil {
		scope.DiagnoseSemanticError("Missing input to range choice", ctx)
	}
	out.Input = traversal.ConstructNode(input, namespace, scope)
	if out.Input == nil {
		scope.DiagnoseSemanticError("Invalid input to range choice", ctx)
	}

	for _, builderWrap := range ctx.AllDensityFn_ClampBuilder() {
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
	return out
}

func (c ClampFnFactory) GetHelp(node *ClampDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Clamps the input density function to the specified range.<br/> `.Min(float)` and `.Max(float)` are required.<br/>Example: `Clamp(_densityFn_).Min(0).Max(1))`",
		Position: node.GetLocation(),
	}
}
func init() {
	traversal.RegisterNodeFactory(ClampFnFactory{})
}

type ClampDensityFn struct {
    Input    traversal.Node
    Min      float64
    Max      float64
    location traversal.TextLocation
    ctx      *grammar.DensityFn_ClampContext
}

func (c ClampDensityFn) GetLocation() traversal.TextLocation {
    return c.location
}

func (c ClampDensityFn) GetCompletions(cursorPosition protocol.Position) []protocol.CompletionItem {
    return []protocol.CompletionItem{
        completions.BuilderFnCompletion("Min", "Min(${1})", c.location.Stop.ToLspPosition()),
        completions.BuilderFnCompletion("Max", "Max(${1})", c.location.Stop.ToLspPosition()),
    }
}

func (c ClampDensityFn) MarshalJSON() ([]byte, error) {
    return json.MarshalIndent(struct {
        Type  string         `json:"type"`
        Input traversal.Node `json:"input"`
        Min   float64        `json:"min"`
        Max   float64        `json:"max"`
    }{
        Type:  "minecraft:clamp",
        Input: c.Input,
        Min:   c.Min,
        Max:   c.Max,
    }, "", "  ")
}
