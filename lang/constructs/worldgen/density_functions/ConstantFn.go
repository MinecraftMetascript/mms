package density_functions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(
		func(densityFn *grammar.DensityFn_ConstantContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			out := &ConstantDensityFn{}
			builder_chain.Builder_GetFloat(
				densityFn, func(f float64) { out.Value = f }, scope, "Value",
			)
			return out
		},
	)
}

type ConstantDensityFn struct {
	Value float64
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

func (c ConstantDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
