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
		reflect.TypeFor[grammar.SurfaceCondition_HoleContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			return &HoleCondition{}
		},
	)
}

type HoleCondition struct {
	traversal.Construct
}

func (c HoleCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c HoleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: HoleConditionKind,
	})
}
