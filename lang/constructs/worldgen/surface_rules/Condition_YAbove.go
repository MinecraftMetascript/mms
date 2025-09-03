package surface_rules

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/constructs/primitives"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_YAboveContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			yAbove := ctx.(*grammar.SurfaceCondition_YAboveContext)

			var anchor primitives.VerticalAnchor
			if vad := yAbove.VerticalAnchor(); vad != nil {
				if cons := traversal.ConstructRegistry.Construct(vad, ns, scope); cons != nil {
					if a, ok := cons.(*primitives.VerticalAnchor); ok && a != nil {
						anchor = *a
					}
				}
			}
			multiplier := 0
			if yAbove.Int() != nil {
				if m, err := strconv.Atoi(yAbove.Int().GetText()); err == nil {
					multiplier = m
				}
			}
			return &YAboveCondition{
				Anchor:     anchor,
				Multiplier: multiplier,
				Add:        yAbove.Keyword_Add() != nil,
			}
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
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: YAboveConditionKind,
	}, "", "  ")
}
