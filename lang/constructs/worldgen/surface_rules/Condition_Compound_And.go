package surface_rules

import (
	"encoding/json"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_AndContext](),
		func(_ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			compound := _ctx.(*grammar.SurfaceCondition_AndContext)
			conditions := make([]traversal.Construct, 0)

			for _, child := range compound.AllSurfaceCondition() {
				var condition antlr.ParserRuleContext

				if child.GetChildCount() == 1 {
					condition = child.GetChild(0).(antlr.ParserRuleContext)
				} else {
					condition = child.GetChild(1).(antlr.ParserRuleContext)
				}

				conditions = append(conditions,
					traversal.ConstructRegistry.Construct(condition, currentNamespace, scope),
				)
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
	return json.Marshal(struct {
		Conditions []traversal.Construct `json:"conditions"`
	}{
		Conditions: c.Conditions,
	})
}

//func init() {
//	traversal.ConstructRegistry.Register(
//		reflect.TypeFor[grammar.SurfaceCondition_CompoundContext](),
//		func(ctx antlr.ParserRuleContext, ns string) traversal.Construct {
//			compound := ctx.(*grammar.SurfaceCondition_CompoundContext)
//			items := make([]traversal.Construct, 0)
//			for _, item := range compound.AllSurfaceCondition_Compound__Item() {
//				items = append(items, traversal.ConstructRegistry.Construct(item, ns))
//			}
//
//			return &CompoundCondition{
//				Conditions: items,
//			}
//		},
//	)
//}
//
//type CompoundCondition struct {
//	traversal.Construct
//	Conditions []traversal.Construct
//}
//
//func (c CompoundCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
//	return exportSurfaceCondition(symbol, rootDir, c)
//}
//
//func (c CompoundCondition) MarshalJSON() ([]byte, error) {
//	if len(c.Conditions) == 1 {
//		return json.Marshal(c.Conditions[0])
//	}
//	return json.Marshal(struct {
//		Type       SurfaceConditionKind         `json:"type"`
//		Conditions []traversal.Construct `json:"conditions"`
//	}{
//		Type:       CompoundConditionType,
//		Conditions: c.Conditions,
//	})
//}
//
//func init() {
//	traversal.ConstructRegistry.Register(
//		reflect.TypeFor[grammar.SurfaceCondition_Compound__ItemContext](),
//		func(ctx antlr.ParserRuleContext, ns string) traversal.Construct {
//			compoundItem := ctx.(*grammar.SurfaceCondition_Compound__ItemContext)
//
//			condition := traversal.ConstructRegistry.Construct(compoundItem.SurfaceConditionDefinition(), ns)
//			negated := compoundItem.Bang() != nil
//			return &CompoundConditionItem{
//				Condition: condition,
//				Negate:    negated,
//			}
//		},
//	)
//}
//
//type CompoundConditionItem struct {
//	json.Marshaler
//	Condition traversal.Construct
//	Negate    bool
//}
//
//func (c CompoundConditionItem) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
//	return nil
//}
//
//func (c CompoundConditionItem) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		Type      SurfaceConditionKind       `json:"type"`
//		Condition traversal.Construct `json:"condition"`
//		Negate    bool                `json:"negate"`
//	}{
//		Type:      CompoundConditionType,
//		Condition: c.Condition,
//		Negate:    c.Negate,
//	})
//}
