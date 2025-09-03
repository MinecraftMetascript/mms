package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type NotCondition struct {
	traversal.Construct
	Invert traversal.Construct
}

func (c NotCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c NotCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type   SurfaceConditionKind `json:"type"`
		Invert traversal.Construct  `json:"invert"`
	}{
		Type:   NotConditionKind,
		Invert: c.Invert,
	}, "", "  ")
}

func InvertCondition(c traversal.Construct) NotCondition {
	return NotCondition{Invert: c}
}
