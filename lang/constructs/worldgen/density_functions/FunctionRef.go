package density_functions

import (
	"errors"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

func init() {
	traversal.RegisterNodeFactory(ReferenceFunctionFactory{}, false)
}

type ReferenceFunctionFactory struct {
	BaseDensityFnFactory
}

func (r ReferenceFunctionFactory) Create(ctx *grammar.DensityFn_ReferenceContext, namespace string, scope *traversal.Scope) *ReferenceFunction {
	out := &ReferenceFunction{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}
	if rr := ctx.ResourceReference(); rr != nil {
		// TODO: Kill the construct registry
		if cons := traversal.ConstructRegistry.Construct(rr, namespace, scope); cons != nil {
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
}

func (r ReferenceFunctionFactory) GetHelp(_ *ReferenceFunction, _ traversal.Symbol, _ traversal.TextLocation) *traversal.Help {
	return nil
}

type ReferenceFunction struct {
	Ref      *traversal.Reference
	Value    *traversal.Construct
	location traversal.TextLocation
}

func (c ReferenceFunction) GetLocation() traversal.TextLocation {
	return c.location
}

func (c ReferenceFunction) MarshalJSON() ([]byte, error) {
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
