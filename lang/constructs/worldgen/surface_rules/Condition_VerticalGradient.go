package surface_rules

import (
	"encoding/json"
	"github.com/minecraftmetascript/mms/lang/constructs/primitives"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_VerticalGradientContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			verticalGradient := ctx.(*grammar.SurfaceCondition_VerticalGradientContext)
			trueAnchor := traversal.ConstructRegistry.Construct(verticalGradient.VerticalAnchorDefinition(0), ns, scope).(*primitives.VerticalAnchor)
			falseAnchor := traversal.ConstructRegistry.Construct(verticalGradient.VerticalAnchorDefinition(1), ns, scope).(*primitives.VerticalAnchor)

			return &VerticalGradientCondition{
				SeedText:        verticalGradient.String_().GetText(),
				TrueAtAndBelow:  *trueAnchor,
				FalseAtAndAbove: *falseAnchor,
			}
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
	return json.Marshal(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: VerticalGradientConditionKind,
	})
}
