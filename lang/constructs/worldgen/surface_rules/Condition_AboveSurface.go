package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(func(ctx *grammar.SurfaceCondition_AboveSurfaceContext, _ string, _ *traversal.Scope) traversal.Construct {
		return &AboveSurfaceCondition{}
	})
}

type AboveSurfaceCondition struct {
	traversal.Construct
}

func (c AboveSurfaceCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c AboveSurfaceCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type SurfaceConditionKind `json:"type"`
	}{
		Type: AboveSurfaceConditionKind,
	}, "", "  ")
}
