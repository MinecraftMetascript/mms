package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(
		func(_ *grammar.SurfaceCondition_HoleContext, _ string, _ *traversal.Scope) traversal.Construct {
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
	return json.MarshalIndent(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: HoleConditionKind,
	}, "", "  ")
}
