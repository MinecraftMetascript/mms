package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(
		func(not *grammar.SurfaceCondition_NotContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			if not.SurfaceCondition() == nil {
				*scope.Diagnostics = append(*scope.Diagnostics, traversal.Diagnostic{
					Message:  "Expected a surface condition",
					Where:    traversal.RuleLocation(not, scope.CurrentFile),
					Severity: traversal.SeverityError,
					Source:   "semantic",
					File:     scope.CurrentFile,
				})
				return nil
			}
			return InvertCondition(
				traversal.ConstructRegistry.Construct(
					not.SurfaceCondition(),
					currentNamespace,
					scope,
				),
			)
		},
	)
}

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
