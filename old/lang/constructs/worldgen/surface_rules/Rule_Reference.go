package surface_rules

import (
	"errors"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(func(refCtx *grammar.SurfaceRule_ReferenceContext, ns string, scope *traversal.Scope) traversal.Construct {
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
	})
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
