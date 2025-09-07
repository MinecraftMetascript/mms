package noise

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.NoiseBlockContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			blockCtx := ctx.(*grammar.NoiseBlockContext)
			for _, child := range blockCtx.GetChildren() {
				if rule, ok := child.(antlr.ParserRuleContext); ok {
					traversal.ConstructRegistry.Construct(rule, currentNamespace, scope)
				}
			}
			return nil
		},
	)
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.NoiseDeclarationContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			declaration := ctx.(*grammar.NoiseDeclarationContext)
			noise := declaration.Noise()
			if noise == nil {
				// TODO: Diagnose?
				return nil
			}
			noiseDef := noise.NoiseDefinition()
			if noiseDef == nil {
				// TODO: Diagnose?
				return nil
			}
			s := traversal.ProcessDeclaration(declaration, noiseDef, scope, currentNamespace, "Noise")
			if s == nil {
				return nil
			}
			return s.GetValue()
		},
	)
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.NoiseDefinitionContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			noise := ctx.(*grammar.NoiseDefinitionContext)
			out := &Noise{
				Amplitudes: make([]float64, 0),
			}

			builder_chain.Builder_GetInt(
				noise, func(i int) { out.FirstOctave = i }, scope, "FirstOctave",
			)

			for _, r := range noise.AllNoise_Builder() {
				if amplitudes := r.Noise_Builder_Amplitudes(); amplitudes != nil {
					for _, amplitude := range amplitudes.AllNumber() {
						if amplitude != nil {
							val, err := strconv.ParseFloat(amplitude.GetText(), 64)
							if err != nil {
								scope.DiagnoseSemanticError(
									"Invalid amplitude value",
									amplitude,
								)
							} else {
								out.Amplitudes = append(out.Amplitudes, val)
							}
						}
					}
				} else {
					scope.DiagnoseSemanticError("Missing amplitudes value", r)
				}
			}

			return out
		},
	)
}

type Noise struct {
	traversal.Construct
	FirstOctave int
	Amplitudes  []float64
}

func (n Noise) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FirstOctave int       `json:"firstOctave"`
		Amplitudes  []float64 `json:"amplitudes"`
	}{
		FirstOctave: n.FirstOctave,
		Amplitudes:  n.Amplitudes,
	})
}

func (n Noise) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("worldgen", nil).
		MkDir("noise", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)
	return nil

}
