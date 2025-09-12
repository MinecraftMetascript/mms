package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.RegisterHelp[*grammar.SurfaceCondition_AboveSurfaceContext](
		func(construct traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
			out := "Checks if the current position is above the preliminary surface level, which is a Y-level usually a few blocks below the main surface, ignoring noise caves. This is used to prevent grass blocks from being placed in noise caves."
			return &out
		},
	)
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
