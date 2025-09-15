package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lang/traversal/getters"
)

type DualInputDensityFnFactory struct {
	BaseDensityFnFactory
}

func (d DualInputDensityFnFactory) Create(ctx *grammar.DensityFn_DualInputContext, namespace string, scope *traversal.Scope) *DualInputDensityFn {
	out := &DualInputDensityFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}
	kind := ctx.GetChild(0)
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

	out.FirstArg = getters.GetChildNode(
		ctx.DensityFn(0),
		namespace,
		scope,
		"Missing density function target",
	)

	out.SecondArg = getters.GetChildNode(
		ctx.DensityFn(1),
		namespace,
		scope,
		"Missing density function target",
	)

	return out
}

func (d DualInputDensityFnFactory) GetHelp(node *DualInputDensityFn, _ traversal.Symbol, _ traversal.TextLocation) *traversal.Help {
	var out string
	switch node.Kind {
	case DensityFn_Min:
		out = "Takes the minimum of 2 density functions.<br/>Functions must be separated by `,`.<br/>Example: `Min(1, 2)`"
	case DensityFn_Max:
		out = "Takes the maximum of 2 density functions.<br/>Functions must be separated by `,`.<br/>Example: `Max(1, 2)`"
	}
	return &traversal.Help{
		Content:  out,
		Position: node.GetLocation(),
	}
}

type DualInputDensityFnKind string

const (
	DensityFn_Min DualInputDensityFnKind = "minecraft:min"
	DensityFn_Max DualInputDensityFnKind = "minecraft:max"
)

func init() {
	traversal.RegisterNodeFactory(DualInputDensityFnFactory{}, false)
}

type DualInputDensityFn struct {
	Kind      DualInputDensityFnKind
	FirstArg  traversal.Node
	SecondArg traversal.Node
	location  traversal.TextLocation
}

func (s DualInputDensityFn) GetLocation() traversal.TextLocation {
	return s.location
}

func (s DualInputDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(
		struct {
			Type      DualInputDensityFnKind `json:"type"`
			FirstArg  traversal.Node         `json:"argument1"`
			SecondArg traversal.Node         `json:"argument2"`
		}{
			Type:      s.Kind,
			FirstArg:  s.FirstArg,
			SecondArg: s.SecondArg,
		},
		"",
		"  ",
	)
}
