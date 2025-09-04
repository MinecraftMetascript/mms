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
	yAboveBuilder := builder_chain.NewBuilderChain(
		builder_chain.Build(
			func(ctx *grammar.SharedBuilder_MulIntContext, target *YAboveCondition, scope *traversal.Scope, namespace string) {
				builder_chain.Builder_GetInt(ctx, func(v int) { target.Multiplier = v }, scope, "Top")
			},
		),
		builder_chain.Build(
			func(ctx *grammar.SharedBuilder_AddContext, target *YAboveCondition, scope *traversal.Scope, namespace string) {
				builder_chain.SharedBuilder_Add(ctx, func(v bool) { target.Add = v })
			},
		),
	)

	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_YAboveContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			yAbove := ctx.(*grammar.SurfaceCondition_YAboveContext)
			out := &YAboveCondition{}

			if vad := yAbove.VerticalAnchor(); vad != nil {
				if cons := traversal.ConstructRegistry.Construct(vad, ns, scope); cons != nil {
					if a, ok := cons.(*primitives.VerticalAnchor); ok && a != nil {
						out.Anchor = *a
					}
				}
			}

			for _, r := range yAbove.AllSurfaceCondition_YAboveBuilder() {
				builder_chain.Invoke(yAboveBuilder, r, out, scope, ns)
			}

			return out
		},
	)
}

type YAboveCondition struct {
	traversal.Construct
	Anchor     primitives.VerticalAnchor
	Multiplier int
	Add        bool
}

func (c YAboveCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c YAboveCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type       SurfaceConditionKind      `json:"type"`
		Anchor     primitives.VerticalAnchor `json:"anchor"`
		Multiplier int                       `json:"surface_depth_multiplier"`
		Add        bool                      `json:"add_stone_depth"`
	}{
		Type:       YAboveConditionKind,
		Anchor:     c.Anchor,
		Multiplier: c.Multiplier,
		Add:        c.Add,
	}, "", "  ")
}
