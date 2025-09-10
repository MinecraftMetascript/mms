package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(
		func(compound *grammar.SurfaceCondition_AndContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			conditions := make([]traversal.Construct, 0)

			for _, child := range compound.AllSurfaceCondition() {
				condition := traversal.ConstructRegistry.Construct(child, currentNamespace, scope)
				conditions = append(conditions, condition)
			}

			return &CompoundAndCondition{
				Conditions: conditions,
			}
		},
	)
}

type CompoundAndCondition struct {
	traversal.Construct
	Conditions []traversal.Construct
}

func (c *CompoundAndCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c *CompoundAndCondition) MarshalJSON() (data []byte, err error) {
	return json.MarshalIndent(struct {
		Conditions []traversal.Construct `json:"conditions"`
	}{
		Conditions: c.Conditions,
	}, "", "  ")
}
