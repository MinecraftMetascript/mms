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

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_WierdScaledSamplerContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			densityFn := ctx.(*grammar.DensityFn_WierdScaledSamplerContext)
			out := &WeirdScaledSamplerFn{}

			input := densityFn.DensityFn()
			if input == nil {
				scope.DiagnoseSemanticError("Missing input density function", ctx)
			} else {
				inputFn := traversal.ConstructRegistry.Construct(input.(antlr.ParserRuleContext), currentNamespace, scope)
				if inputFn == nil {
					scope.DiagnoseSemanticError("Invalid input density function", ctx)
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
				if _, ok := builder.(*grammar.DensityFn_WierdScaledSamplerBuilder_Type1Context); ok {
					out.ValueMapper = "type_1"
				}
				if _, ok := builder.(*grammar.DensityFn_WierdScaledSamplerBuilder_Type2Context); ok {
					out.ValueMapper = "type_2"
				}
				if noise, ok := builder.(*grammar.DensityFn_WierdScaledSamplerBuilder_NoiseContext); ok {
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
							fmt.Println(ref)
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
