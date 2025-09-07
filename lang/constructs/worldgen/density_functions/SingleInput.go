package density_functions

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

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
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_SingleInputContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			densityFn := ctx.(*grammar.DensityFn_SingleInputContext)
			out := &SingleInputDensityFn{}

			kind := densityFn.GetChild(0)
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

			if target := densityFn.DensityFn(); target != nil {
				val := traversal.ConstructRegistry.Construct(target.(antlr.ParserRuleContext), currentNamespace, scope)
				fmt.Println(">>>", reflect.TypeOf(val).Elem().Name(), val)
				if val == nil {
					scope.DiagnoseSemanticError("Missing density function target", ctx)
				} else {
					out.Child = val
				}
			} else {
				scope.DiagnoseSemanticError("Missing density function target", ctx)
			}

			return out
		},
	)
}

type SingleInputDensityFn struct {
	Kind  SingleInputDensityFnKind
	Child traversal.Construct
}

func (s SingleInputDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     string              `json:"type"`
		Argument traversal.Construct `json:"argument"`
	}{
		Type:     string(s.Kind),
		Argument: s.Child,
	}, "", "  ")
}

func (s SingleInputDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, s)
}
