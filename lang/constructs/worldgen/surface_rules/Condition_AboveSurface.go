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
		reflect.TypeFor[grammar.SurfaceCondition_AboveSurfaceContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			return &AboveSurfaceCondition{}
		},
	)
}

type AboveSurfaceCondition struct {
	traversal.Construct
}

func (c AboveSurfaceCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c AboveSurfaceCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: AboveSurfaceConditionKind,
	})
}
