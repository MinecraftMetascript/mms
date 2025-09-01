package surface_rules

import (
	"encoding/json"
	"mms2/lang/constructs/primitives"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_YAboveContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			yAbove := ctx.(*grammar.SurfaceCondition_YAboveContext)

			x := traversal.ConstructRegistry.Construct(yAbove.VerticalAnchor(), ns, scope)

			anchor := x.(*primitives.VerticalAnchor)
			multiplier, _ := strconv.Atoi(yAbove.Int().GetText())
			return &YAboveCondition{
				Anchor:     *anchor,
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
	return json.Marshal(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: YAboveConditionKind,
	})
}
