package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type WeirdScaledSamplerFnFactory struct {
	BaseDensityFnFactory
}

func (w WeirdScaledSamplerFnFactory) Create(ctx *grammar.DensityFn_WierdScaledSamplerContext, namespace string, scope *traversal.Scope) *WeirdScaledSamplerFn {
	out := &WeirdScaledSamplerFn{
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}

	input := ctx.DensityFn()
	if input == nil {
		scope.DiagnoseSemanticError("Missing input density function", ctx)
	} else {
		inputFn := traversal.ConstructRegistry.Construct(input.(antlr.ParserRuleContext), namespace, scope)
		if inputFn == nil {
			scope.DiagnoseSemanticError("Invalid input density function", ctx)
		} else {
			out.Input = inputFn
		}
	}

	for _, builderWrap := range ctx.AllDensityFn_WierdScaledSamplerBuilder() {
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
				ref := traversal.ConstructRegistry.Construct(noise.ResourceReference(), namespace, scope)
				if r, ok := ref.(*traversal.Reference); ok {
					out.Noise = *r
				} else {
					scope.DiagnoseSemanticError("Invalid noise reference", noise)
					// TODO: Diagnose
					continue
				}
			} else if inline := noise.NoiseDefinition(); inline != nil {
				_, ref := traversal.ExtractInlineConstruct(inline, namespace, scope, "Noise")

				if ref != nil {
					out.Noise = *ref
				} else {
					scope.DiagnoseSemanticError("Invalid inline noise definition.", noise)
					continue
				}

			}
		}
	}

	return out
}

func (w WeirdScaledSamplerFnFactory) GetHelp(node *WeirdScaledSamplerFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "According to the input value, scales and enhances (or weakens) some regions of the specified noise, and then returns the absolute value.",
		Position: node.GetLocation(),
	}
}

func init() {
	traversal.RegisterNodeFactory(WeirdScaledSamplerFnFactory{}, false)
}

type WeirdScaledSamplerFn struct {
	Noise       traversal.Reference
	ValueMapper string
	Input       traversal.Construct

	location traversal.TextLocation
}

func (c WeirdScaledSamplerFn) GetLocation() traversal.TextLocation {
	return c.location
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
