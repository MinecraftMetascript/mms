package density_functions

import (
	"encoding/json"
	"reflect"

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
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_DualInputContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			densityFn := ctx.(*grammar.DensityFn_DualInputContext)
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
						scope.DiagnoseSemanticError("Invalid density function kind", ctx)
					}
				} else {
					scope.DiagnoseSemanticError("Missing density function kind", ctx)
				}
			default:
				scope.DiagnoseSemanticError("Invalid density function kind", ctx)
			}

			firstArg := densityFn.DensityFn(0)
			if firstArg != nil {
				val := traversal.ConstructRegistry.Construct(firstArg, currentNamespace, scope)
				if val == nil {
					scope.DiagnoseSemanticError("Missing density function target", ctx)
				} else {
					out.FirstArg = val
				}
			} else {
				scope.DiagnoseSemanticError("Missing density function target", ctx)
			}

			secondArg := densityFn.DensityFn(1)
			if secondArg != nil {
				val := traversal.ConstructRegistry.Construct(secondArg, currentNamespace, scope)
				if val == nil {
					scope.DiagnoseSemanticError("Missing density function target", ctx)
				} else {
					out.SecondArg = val
				}
			} else {
				scope.DiagnoseSemanticError("Missing density function target", ctx)
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
