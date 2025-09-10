package noise_router

import (
	"encoding/json"
	"errors"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
 traversal.Register(func(blockCtx *grammar.NoiseRouterBlockContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
	for _, child := range blockCtx.GetChildren() {
		if rule, ok := child.(antlr.ParserRuleContext); ok {
			traversal.ConstructRegistry.Construct(rule, currentNamespace, scope)
		}
	}
	return nil
})
	traversal.Register(func(declaration *grammar.NoiseRouterDeclarationContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {

			routerDef := declaration.NoiseRouter()
			if routerDef == nil {
				scope.DiagnoseSemanticError("Missing noise router definition", declaration)
				return nil
			}

			s := traversal.ProcessDeclaration(declaration.Declare(), routerDef, scope, currentNamespace, "NoiseRouter")
			if s == nil {
				return nil
			}
			return s.GetValue()
		})

	traversal.Register(func(router *grammar.NoiseRouterContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
		out := &NoiseRouter{}

			finalDensityPresent := false
			for _, builderCtx := range router.AllNoiseRouter_Builder() {
				builderKindCtx := builderCtx.GetChild(1)
				if builderKindCtx == nil {
					scope.DiagnoseSemanticError("Missing builder", router)
					continue
				}
				builderKind := builderKindCtx.(antlr.TerminalNode).GetText()

				value := traversal.ConstructRegistry.Construct(builderCtx.DensityFn(), currentNamespace, scope)
				switch NoiseRouterBuilderKind(builderKind) {
				case FinalDensity:
					finalDensityPresent = true
					out.FinalDensity = value
				case Barrier:
					out.Barrier = value
				case FluidLevelFloodedness:
					out.FluidLevelFloodedness = value
				case FluidLevelSpread:
					out.FluidLevelSpread = value
				case Lava:
					out.Lava = value
				case VeinToggle:
					out.VeinToggle = value
				case VeinRidged:
					out.VeinRidged = value
				case VeinGap:
					out.VeinGap = value
				case Temperature:
					out.Temperature = value
				case Vegetation:
					out.Vegetation = value
				case Continents:
					out.Continents = value
				case Erosion:
					out.Erosion = value
				case Depth:
					out.Depth = value
				case Ridges:
					out.Ridges = value
				}
			}

			if !finalDensityPresent {
				scope.DiagnoseSemanticError("Missing required FinalDensity()", router)
			}
			return out
		})
}

type NoiseRouter struct {
	FinalDensity          traversal.Construct `json:"final_density,omitempty"`
	Barrier               traversal.Construct `json:"barrier,omitempty"`
	FluidLevelFloodedness traversal.Construct `json:"fluid_level_floodedness,omitempty"`
	FluidLevelSpread      traversal.Construct `json:"fluid_level_spread,omitempty"`
	Lava                  traversal.Construct `json:"lava,omitempty"`
	VeinToggle            traversal.Construct `json:"vein_toggle,omitempty"`
	VeinRidged            traversal.Construct `json:"vein_ridged,omitempty"`
	VeinGap               traversal.Construct `json:"vein_gap,omitempty"`
	Temperature           traversal.Construct `json:"temperature,omitempty"`
	Vegetation            traversal.Construct `json:"vegetation,omitempty"`
	Continents            traversal.Construct `json:"continents,omitempty"`
	Erosion               traversal.Construct `json:"erosion,omitempty"`
	Depth                 traversal.Construct `json:"depth,omitempty"`
	Ridges                traversal.Construct `json:"ridges,omitempty"`
}

func (n *NoiseRouter) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		FinalDensity          traversal.Construct `json:"final_density,omitempty"`
		Barrier               traversal.Construct `json:"barrier,omitempty"`
		FluidLevelFloodedness traversal.Construct `json:"fluid_level_floodedness,omitempty"`
		FluidLevelSpread      traversal.Construct `json:"fluid_level_spread,omitempty"`
		Lava                  traversal.Construct `json:"lava,omitempty"`
		VeinToggle            traversal.Construct `json:"vein_toggle,omitempty"`
		VeinRidged            traversal.Construct `json:"vein_ridged,omitempty"`
		VeinGap               traversal.Construct `json:"vein_gap,omitempty"`
		Temperature           traversal.Construct `json:"temperature,omitempty"`
		Vegetation            traversal.Construct `json:"vegetation,omitempty"`
		Continents            traversal.Construct `json:"continents,omitempty"`
		Erosion               traversal.Construct `json:"erosion,omitempty"`
		Depth                 traversal.Construct `json:"depth,omitempty"`
		Ridges                traversal.Construct `json:"ridges,omitempty"`
	}{
		FinalDensity:          n.FinalDensity,
		Barrier:               n.Barrier,
		FluidLevelFloodedness: n.FluidLevelFloodedness,
		FluidLevelSpread:      n.FluidLevelSpread,
		Lava:                  n.Lava,
		VeinToggle:            n.VeinToggle,
		VeinRidged:            n.VeinRidged,
		VeinGap:               n.VeinGap,
		Temperature:           n.Temperature,
		Vegetation:            n.Vegetation,
		Continents:            n.Continents,
		Erosion:               n.Erosion,
		Depth:                 n.Depth,
		Ridges:                n.Ridges,
	}, "", "  ")
}

func (n *NoiseRouter) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	ref := symbol.GetReference()
	if ref.String() == ":" {
		return errors.New("symbol has no reference")
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("_debug", nil).
		MkDir("noise_router", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)

	return nil
}

type NoiseRouterBuilderKind string

const (
	FinalDensity          NoiseRouterBuilderKind = "FinalDensity"
	Barrier               NoiseRouterBuilderKind = "Barrier"
	FluidLevelFloodedness NoiseRouterBuilderKind = "FluidLevelFloodedness"
	FluidLevelSpread      NoiseRouterBuilderKind = "FluidLevelSpread"
	Lava                  NoiseRouterBuilderKind = "Lava"
	VeinToggle            NoiseRouterBuilderKind = "VeinToggle"
	VeinRidged            NoiseRouterBuilderKind = "VeinRidged"
	VeinGap               NoiseRouterBuilderKind = "VeinGap"
	Temperature           NoiseRouterBuilderKind = "Temperature"
	Vegetation            NoiseRouterBuilderKind = "Vegetation"
	Continents            NoiseRouterBuilderKind = "Continents"
	Erosion               NoiseRouterBuilderKind = "Erosion"
	Depth                 NoiseRouterBuilderKind = "Depth"
	Ridges                NoiseRouterBuilderKind = "Ridges"
)

type NoiseRouterBuilder struct {
	Type NoiseRouterBuilderKind
	Fn   *traversal.Construct
}
