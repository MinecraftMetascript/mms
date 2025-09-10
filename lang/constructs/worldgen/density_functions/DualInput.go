package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type DualInputDensityFnKind string

const (
	DensityFn_Min DualInputDensityFnKind = "minecraft:min"
	DensityFn_Max DualInputDensityFnKind = "minecraft:max"
)

func init() {
	traversal.Register(
		func(densityFn *grammar.DensityFn_DualInputContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			out := &DualInputDensityFn{}
			kind := densityFn.GetChild(0)
			switch k := kind.(type) {
			case antlr.TerminalNode:
				if k != nil {
					switch k.GetText() {
					case "Min":
						out.Kind = DensityFn_Min
					case "Max":
						out.Kind = DensityFn_Max
					default:
						scope.DiagnoseSemanticError("Invalid density function kind", densityFn)
					}
				} else {
					scope.DiagnoseSemanticError("Missing density function kind", densityFn)
				}
			default:
				scope.DiagnoseSemanticError("Invalid density function kind", densityFn)
			}

			firstArg := densityFn.DensityFn(0)
			if firstArg != nil {
				val := traversal.ConstructRegistry.Construct(firstArg, currentNamespace, scope)
				if val == nil {
					scope.DiagnoseSemanticError("Missing density function target", densityFn)
				} else {
					out.FirstArg = val
				}
			} else {
				scope.DiagnoseSemanticError("Missing density function target", densityFn)
			}

			secondArg := densityFn.DensityFn(1)
			if secondArg != nil {
				val := traversal.ConstructRegistry.Construct(secondArg, currentNamespace, scope)
				if val == nil {
					scope.DiagnoseSemanticError("Missing density function target", densityFn)
				} else {
					out.SecondArg = val
				}
			} else {
				scope.DiagnoseSemanticError("Missing density function target", densityFn)
			}

			return out
		},
	)
}

type DualInputDensityFn struct {
	Kind      DualInputDensityFnKind
	FirstArg  traversal.Construct
	SecondArg traversal.Construct
}

func (s DualInputDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(
		struct {
			Type      DualInputDensityFnKind `json:"type"`
			FirstArg  traversal.Construct    `json:"argument1"`
			SecondArg traversal.Construct    `json:"argument2"`
		}{
			Type:      s.Kind,
			FirstArg:  s.FirstArg,
			SecondArg: s.SecondArg,
		},
		"",
		"  ",
	)
}

func (s DualInputDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, s)
}
