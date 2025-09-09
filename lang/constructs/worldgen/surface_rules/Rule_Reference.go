package surface_rules

import (
	"errors"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRule_ReferenceContext](),
		func(ctx antlr.ParserRuleContext, ns string, scope *traversal.Scope) traversal.Construct {
			refCtx := ctx.(*grammar.SurfaceRule_ReferenceContext)
			out := &ReferenceRule{}
			var ref *traversal.Reference
			if rr := refCtx.ResourceReference(); rr != nil {
				if cons := traversal.ConstructRegistry.Construct(rr, ns, scope); cons != nil {
					if r, ok := cons.(*traversal.Reference); ok {
						ref = r
					}
				}
			}
			if ref != nil {
				ref.SetResolver(func() error {
					if next, ok := scope.Get(*ref); ok {
						val := next.GetValue().(traversal.Construct)
						out.Value = &val
					}
					return nil
				})
			}
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
	if c.Ref == nil {
		return nil, errors.New("reference is nil in ReferenceRule")
	}
	if err := c.Ref.Resolve(); err != nil {
		return nil, err
	}
	if c.Value == nil {
		return c.Ref.MarshalJSON()
	}
	return (*c.Value).MarshalJSON()
}
