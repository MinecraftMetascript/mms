package density_functions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lang/traversal/getters"
)

type ConstantFnFactory struct {
	BaseDensityFnFactory
}

func (c ConstantFnFactory) Create(ctx *grammar.DensityFn_ConstantContext, _ string, scope *traversal.Scope) *ConstantDensityFn {
	out := &ConstantDensityFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}
	getters.GetFloat(
		ctx, func(f float64) { out.Value = f }, scope, "Value",
	)
	return out
}

func (c ConstantFnFactory) GetHelp(_ *ConstantDensityFn, _ traversal.Symbol, _ traversal.TextLocation) *traversal.Help {
	return nil
}

func init() {
	traversal.RegisterNodeFactory(ConstantFnFactory{}, false)
}

type ConstantDensityFn struct {
	Value    float64
	location traversal.TextLocation
}

func (c ConstantDensityFn) GetLocation() traversal.TextLocation {
	return c.location
}

func (c ConstantDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     string  `json:"type"`
		Argument float64 `json:"argument"`
	}{
		Type:     "minecraft:constant",
		Argument: c.Value,
	}, "", "  ")
}
