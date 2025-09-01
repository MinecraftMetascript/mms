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
		reflect.TypeFor[grammar.SurfaceCondition_ReferenceContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			ref := ctx.(*grammar.SurfaceCondition_ReferenceContext)

			return &ReferenceCondition{
				Ref: traversal.ConstructRegistry.Construct(ref.ResourceReference(), ns, scope).(*traversal.Reference),
			}
		},
	)
}

type ReferenceCondition struct {
	traversal.Construct
	Ref *traversal.Reference
}

func (c ReferenceCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c ReferenceCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type    SurfaceConditionKind `json:"type"`
		Ref     traversal.Reference  `json:"ref"`
		Comment string               `json:"__comment"`
	}{
		Type:    ReferenceConditionKind,
		Ref:     *c.Ref,
		Comment: "This should not appear in the JSON output of MMS. If you are seeing this, something went wrong.",
	})
}
