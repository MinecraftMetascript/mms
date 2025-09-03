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
		reflect.TypeFor[grammar.SurfaceCondition_OrContext](),
		func(_ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			compound := _ctx.(*grammar.SurfaceCondition_OrContext)
			conditions := make([]traversal.Construct, 0)
			for _, child := range compound.AllSurfaceCondition() {
				next := traversal.ConstructRegistry.Construct(child.GetChild(0).(antlr.ParserRuleContext), currentNamespace, scope)
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
