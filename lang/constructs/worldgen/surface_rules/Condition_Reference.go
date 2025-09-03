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
		reflect.TypeFor[grammar.SurfaceCondition_ReferenceContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			ref := ctx.(*grammar.SurfaceCondition_ReferenceContext)

			var refVal *traversal.Reference
			if rr := ref.ResourceReference(); rr != nil {
				if cons := traversal.ConstructRegistry.Construct(rr, ns, scope); cons != nil {
					if r, ok := cons.(*traversal.Reference); ok {
						refVal = r
					}
				}
			}
			return &ReferenceCondition{
				Ref: refVal,
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
	return json.MarshalIndent(struct {
		Type    SurfaceConditionKind `json:"type"`
		Ref     traversal.Reference  `json:"ref"`
		Comment string               `json:"__comment"`
	}{
		Type:    ReferenceConditionKind,
		Ref:     *c.Ref,
		Comment: "This should not appear in the JSON output of MMS. If you are seeing this, something went wrong.",
	}, "", "  ")
}
