package surface_rules

import (
	"errors"
	"fmt"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRule_ReferenceContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			refCtx := ctx.(*grammar.SurfaceRule_ReferenceContext)
			out := &ReferenceRule{}
			ref := traversal.ConstructRegistry.Construct(refCtx.ResourceReference(), ns, scope).(*traversal.Reference)
			ref.SetResolver(func() error {
				if next, ok := scope.Get(*ref); ok {
					val := next.GetValue().(traversal.Construct)
					out.Value = &val
				}
				return nil
			})
			out.Ref = ref

			return out
		},
	)
}

type ReferenceRule struct {
	traversal.Construct
	Ref   *traversal.Reference
	Value *traversal.Construct
}

func (c *ReferenceRule) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceRule(symbol, rootDir, c)
}

func (c *ReferenceRule) MarshalJSON() ([]byte, error) {
	if err := c.Ref.Resolve(); err != nil {
		return nil, err
	}
	if c.Value == nil {
		return nil, errors.New(
			fmt.Sprintf("unable to resolve reference %s", c.Ref.String()),
		)
	}
	return (*c.Value).MarshalJSON()
}
