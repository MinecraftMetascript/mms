package density_functions

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type ReferenceFunction struct {
	traversal.Construct
	Ref   *traversal.Reference
	Value *traversal.Construct
}

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.DensityFn_ReferenceContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			refCtx := ctx.(*grammar.DensityFn_ReferenceContext)
			out := &ReferenceFunction{}
			if rr := refCtx.ResourceReference(); rr != nil {
				if cons := traversal.ConstructRegistry.Construct(rr, ns, scope); cons != nil {
					if r, ok := cons.(*traversal.Reference); ok {
						out.Ref = r

						out.Ref.SetResolver(func() error {
							if next, ok := scope.Get(*r); ok {
								val := next.GetValue().(traversal.Construct)

								out.Value = &val
							}
							return nil
						})
					}
				}
			}

			return out
		},
	)
}

func (c *ReferenceFunction) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}

func (c *ReferenceFunction) MarshalJSON() ([]byte, error) {
	if c.Ref == nil {
		return nil, errors.New("reference is nil in ReferenceRule")
	}
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
