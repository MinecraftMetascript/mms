package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_AndContext](),
		func(_ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			compound := _ctx.(*grammar.SurfaceCondition_AndContext)
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
