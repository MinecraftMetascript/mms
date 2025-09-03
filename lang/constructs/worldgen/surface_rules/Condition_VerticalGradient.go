package surface_rules

import (
	"encoding/json"
	"reflect"

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
			verticalGradient := ctx.(*grammar.SurfaceCondition_VerticalGradientContext)
			var trueAnchor primitives.VerticalAnchor
			if vad := verticalGradient.VerticalAnchor(0); vad != nil {
				if cons := traversal.ConstructRegistry.Construct(vad, ns, scope); cons != nil {
					if a, ok := cons.(*primitives.VerticalAnchor); ok && a != nil {
						trueAnchor = *a
					}
				}
			}
			var falseAnchor primitives.VerticalAnchor
			if vad := verticalGradient.VerticalAnchor(1); vad != nil {
				if cons := traversal.ConstructRegistry.Construct(vad, ns, scope); cons != nil {
					if a, ok := cons.(*primitives.VerticalAnchor); ok && a != nil {
						falseAnchor = *a
					}
				}
			}

			seed := ""
			if s := verticalGradient.String_(); s != nil {
				seed = s.GetText()
			}

			return &VerticalGradientCondition{
				SeedText:        seed,
				TrueAtAndBelow:  trueAnchor,
				FalseAtAndAbove: falseAnchor,
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
