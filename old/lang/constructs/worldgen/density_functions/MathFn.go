package density_functions

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type MathDensityFnKind string

const (
	MathDensityFn_Add MathDensityFnKind = "minecraft:add"
	MathDensityFn_Mul MathDensityFnKind = "minecraft:mul"
)

type MathDensityFn struct {
	Type       MathDensityFnKind
	Arg1, Arg2 traversal.Node
	location   traversal.TextLocation
}

func (m MathDensityFn) GetLocation() traversal.TextLocation {
	return m.location
}

func (m MathDensityFn) MarshalJSON() ([]byte, error) {
	out := struct {
		Type MathDensityFnKind `json:"type"`
		Arg1 any               `json:"argument1"`
		Arg2 any               `json:"argument2"`
	}{
		Type: m.Type,
		Arg1: m.Arg1,
		Arg2: m.Arg2,
	}
	if v, ok := m.Arg1.(*ConstantDensityFn); ok {
		out.Arg1 = v.Value
	}
	if v, ok := m.Arg2.(*ConstantDensityFn); ok {
		out.Arg2 = v.Value
	}

	return json.MarshalIndent(out, "", "  ")
}

func (m MathDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, m)
}
