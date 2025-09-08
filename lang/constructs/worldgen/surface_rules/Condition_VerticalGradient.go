package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/constructs/primitives"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_VerticalGradientContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			verticalGradientBuilder := builder_chain.NewBuilderChain(
				builder_chain.Build(
					func(ctx *grammar.Builder_TopContext, target *VerticalGradientCondition, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetVerticalAnchor(ctx, namespace, func(anchor primitives.VerticalAnchor) { target.TrueAtAndBelow = anchor }, scope, "Top")
					},
				),
				builder_chain.Build(
					func(ctx *grammar.Builder_BottomContext, target *VerticalGradientCondition, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetVerticalAnchor(ctx, namespace, func(anchor primitives.VerticalAnchor) { target.FalseAtAndAbove = anchor }, scope, "Bottom")
					},
				),
			)

			verticalGradient := ctx.(*grammar.SurfaceCondition_VerticalGradientContext)
			out := &VerticalGradientCondition{}

			if s := verticalGradient.String_(); s != nil {
				out.SeedText = s.GetText()
			} else {
				scope.DiagnoseSemanticError(
					"Missing seed value",
					ctx,
				)
			}

			for _, r := range verticalGradient.AllSurfaceCondition_VerticalGradientBuilder() {
				child := r.GetChild(0)
				if child == nil {
					continue
				}
				builder_chain.Invoke(
					verticalGradientBuilder, child.(antlr.ParserRuleContext), out, scope, ns,
				)
			}

			return out

		},
	)
}

type VerticalGradientCondition struct {
	traversal.Construct
	SeedText        string
	TrueAtAndBelow  primitives.VerticalAnchor
	FalseAtAndAbove primitives.VerticalAnchor
}

func (c VerticalGradientCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c VerticalGradientCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type         SurfaceConditionKind      `json:"type"`
		TrueAtBelow  primitives.VerticalAnchor `json:"true_at_and_below"`
		FalseAtAbove primitives.VerticalAnchor `json:"false_at_and_above"`
	}{
		Type:         VerticalGradientConditionKind,
		TrueAtBelow:  c.TrueAtAndBelow,
		FalseAtAbove: c.FalseAtAndAbove,
	}, "", "  ")
}
