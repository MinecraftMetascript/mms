package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type SingleInputDensityFnFactory struct {
	BaseDensityFnFactory
}

func (s SingleInputDensityFnFactory) Create(ctx *grammar.DensityFn_SingleInputContext, namespace string, scope *traversal.Scope) *SingleInputDensityFn {
	out := &SingleInputDensityFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}

	kind := ctx.GetChild(0)
	switch k := kind.(type) {
	case antlr.TerminalNode:
		if k != nil {
			switch k.GetText() {
			case "Interpolated":
				out.Kind = DensityFn_Interpolated
			case "FlatCache":
				out.Kind = DensityFn_FlatCache
			case "Abs":
				out.Kind = DensityFn_Abs
			case "Square":
				out.Kind = DensityFn_Square
			case "Cube":
				out.Kind = DensityFn_Cube
			case "HalfNeg":
				out.Kind = DensityFn_HalfNeg
			case "QuarterNeg":
				out.Kind = DensityFn_QuarterNeg
			case "Squeeze":
				out.Kind = DensityFn_Squeeze
			case "Shift":
				out.Kind = DensityFn_Shift
			case "ShiftA":
				out.Kind = DensityFn_ShiftA
			case "ShiftB":
				out.Kind = DensityFn_ShiftB
			default:
				scope.DiagnoseSemanticError("Invalid density function kind", ctx)
			}
		} else {
			scope.DiagnoseSemanticError("Missing density function kind", ctx)
		}
	default:
		scope.DiagnoseSemanticError("Invalid density function kind", ctx)
	}

	if target := ctx.DensityFn(); target != nil {

		val := traversal.ConstructNode(target, namespace, scope)
		if val == nil {
			scope.DiagnoseSemanticError("Missing density function target", ctx)
		} else {
			out.Child = val
		}
	} else {
		scope.DiagnoseSemanticError("Missing density function target", ctx)
	}

	return out
}

func (s SingleInputDensityFnFactory) GetHelp(node *SingleInputDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	var out string
	switch node.Kind {
	case DensityFn_Interpolated:
		out = "Interpolates at each block in one cell based on the input density function value of some cells around. The size of each cell is size_horizontal * 4 and size_vertical * 4. Used often in combination with flat_cache."
	case DensityFn_FlatCache:
		out = "Calculate the value per 4×4 column (Value at each block in one column is the same). And it is calculated only once per column, at Y=0. Used often in combination with interpolated."
	case DensityFn_Abs:
		out = "Modifies input function by taking the absolute value.<br/>Example: `Abs(1)`"
	case DensityFn_Cube:
		out = "Modifies input function by taking the cube of the value. (x^3)<br/>Example: `Cube(1)`"
	case DensityFn_Square:
		out = "Modifies input function by taking the square of the value. (x^2)<br/>Example: `Square(1)`"
	case DensityFn_HalfNeg:
		out = "If the input is negative, returns half of the input. Otherwise returns the input. (x < 0 ? x/2 : x)"
	case DensityFn_QuarterNeg:
		out = "If the input is negative, returns a quarter of the input. Otherwise returns the input. (x < 0 ? x/4 : x)"
	case DensityFn_Squeeze:
		out = "First clamps the input between −1 and 1, then transforms it using x/2 - x*x*x/24."
	case DensityFn_Shift:
		out = "Samples a noise at (x/4, y/4, z/4), then multiplies it by 4."
	case DensityFn_ShiftA:
		out = "Samples a noise at (x/4, 0, z/4), then multiplies it by 4."
	case DensityFn_ShiftB:
		out = "Samples a noise at (z/4, x/4, 0), then multiplies it by 4."
	}
	return &traversal.Help{
		Content:  out,
		Position: node.GetLocation(),
	}

}

type SingleInputDensityFnKind string

const (
	DensityFn_Interpolated SingleInputDensityFnKind = "minecraft:interpolated"
	DensityFn_FlatCache    SingleInputDensityFnKind = "minecraft:flat_cache"
	DensityFn_Abs          SingleInputDensityFnKind = "minecraft:abs"
	DensityFn_Square       SingleInputDensityFnKind = "minecraft:square"
	DensityFn_Cube         SingleInputDensityFnKind = "minecraft:cube"
	DensityFn_HalfNeg      SingleInputDensityFnKind = "minecraft:half_negative"
	DensityFn_QuarterNeg   SingleInputDensityFnKind = "minecraft:quarter_negative"
	DensityFn_Squeeze      SingleInputDensityFnKind = "minecraft:squeeze"
	DensityFn_Shift        SingleInputDensityFnKind = "minecraft:shift"
	DensityFn_ShiftA       SingleInputDensityFnKind = "minecraft:shift_a"
	DensityFn_ShiftB       SingleInputDensityFnKind = "minecraft:shift_b"
)

func init() {

}

type SingleInputDensityFn struct {
	Kind     SingleInputDensityFnKind
	Child    traversal.Node
	location traversal.TextLocation
}

func (s SingleInputDensityFn) GetLocation() traversal.TextLocation {
	return s.location
}

func (s SingleInputDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     string         `json:"type"`
		Argument traversal.Node `json:"argument"`
	}{
		Type:     string(s.Kind),
		Argument: s.Child,
	}, "", "  ")
}

func (s SingleInputDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, s)
}
