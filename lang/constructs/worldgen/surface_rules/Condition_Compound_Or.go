package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.RegisterHelp[*grammar.SurfaceCondition_OrContext](
		func(construct traversal.Construct, symbol traversal.Symbol, location traversal.TextLocation) *string {
			out := "Combines multiple conditions with OR. <br/> Under the hood, this creates a minecraft:sequence, with each condition as a direct child, and the result of this OR is the pass rule for all conditions in the sequence"
			return &out
		},
	)

	traversal.Register(
		func(compound *grammar.SurfaceCondition_OrContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			conditions := make([]traversal.Construct, 0)
			for _, child := range compound.AllSurfaceCondition() {
				next := traversal.ConstructRegistry.Construct(child, currentNamespace, scope)
				conditions = append(conditions, next)
			}
			return &CompoundOrCondition{
				Conditions: conditions,
			}
		},
	)
}

type CompoundOrCondition struct {
	traversal.Construct
	Conditions []traversal.Construct
}

func (c *CompoundOrCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c *CompoundOrCondition) MarshalJSON() (data []byte, err error) {
	return json.MarshalIndent(struct {
		Conditions []traversal.Construct `json:"conditions"`
	}{
		Conditions: c.Conditions,
	}, "", "  ")
}
