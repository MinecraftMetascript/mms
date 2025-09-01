package surface_rules

import (
	"encoding/json"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_SteepContext](),
		func(ctx antlr.ParserRuleContext, _ string, _ *traversal.Scope) traversal.Construct {
			return &SteepCondition{}
		},
	)
}

type SteepCondition struct {
	traversal.Construct
}

func (c SteepCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c SteepCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: SteepConditionKind,
	})
}
