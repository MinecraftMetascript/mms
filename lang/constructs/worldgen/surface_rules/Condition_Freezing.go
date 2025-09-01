package surface_rules

import (
	"encoding/json"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_FreezingContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			return &FreezingCondition{}
		},
	)
}

type FreezingCondition struct {
	traversal.Construct
}

func (c FreezingCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c FreezingCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: FreezingConditionKind,
	})
}
