package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(
		func(_ *grammar.SurfaceCondition_SteepContext, _ string, _ *traversal.Scope) traversal.Construct {
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
	return json.MarshalIndent(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: SteepConditionKind,
	}, "", "  ")
}
