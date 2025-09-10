package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.Register(

		func(densityFn *grammar.DensityFn_WierdScaledSamplerContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			out := &WeirdScaledSamplerFn{}

			input := densityFn.DensityFn()
			if input == nil {
				scope.DiagnoseSemanticError("Missing input density function", densityFn)
			} else {
				inputFn := traversal.ConstructRegistry.Construct(input.(antlr.ParserRuleContext), currentNamespace, scope)
				if inputFn == nil {
					scope.DiagnoseSemanticError("Invalid input density function", densityFn)
				} else {
					out.Input = inputFn
				}
			}

			for _, builderWrap := range densityFn.AllDensityFn_WierdScaledSamplerBuilder() {
				builder := builderWrap.GetChild(0)
				if builder == nil {
					// TODO: Diagnose
					continue
				}
				if _, ok := builder.(*grammar.Builder_Type1Context); ok {
					out.ValueMapper = "type_1"
				}
				if _, ok := builder.(*grammar.Builder_Type2Context); ok {
					out.ValueMapper = "type_2"
				}
				if noise, ok := builder.(*grammar.Builder_NoiseContext); ok {
					if noise.ResourceReference() != nil {
						ref := traversal.ConstructRegistry.Construct(noise.ResourceReference(), currentNamespace, scope)
						if r, ok := ref.(*traversal.Reference); ok {
							out.Noise = *r
						} else {
							scope.DiagnoseSemanticError("Invalid noise reference", noise)
							// TODO: Diagnose
							continue
						}
					} else if inline := noise.NoiseDefinition(); inline != nil {
						ref := ExtractInlinedNoise(inline.(*grammar.NoiseDefinitionContext), currentNamespace, scope)
						if ref != nil {
							out.Noise = *ref
						} else {
							// TODO: Diagnose
							scope.DiagnoseSemanticError("Invalid inline noise definition.", noise)
							continue
						}

					}
				}
			}

			return out
		},
	)
}

type WeirdScaledSamplerFn struct {
	Noise       traversal.Reference
	ValueMapper string
	Input       traversal.Construct
}

func (c WeirdScaledSamplerFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type              string              `json:"type"`
		RarityValueMapper string              `json:"rarity_value_mapper"`
		Noise             string              `json:"noise"`
		Input             traversal.Construct `json:"input"`
	}{
		Type:              "minecraft:weird_scaled_sampler",
		RarityValueMapper: c.ValueMapper,
		Noise:             c.Noise.String(),
		Input:             c.Input,
	}, "", "  ")
}

func (c WeirdScaledSamplerFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
