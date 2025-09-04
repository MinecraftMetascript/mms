package noise

import (
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.NoiseDefinitionContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			noise := ctx.(*grammar.NoiseDefinitionContext)
			out := &Noise{}
			builder_chain.Builder_GetInt(
				noise, func(i int) { out.FirstOctave = i }, scope, "FirstOctave",
			)

			for _, r := range noise.AllNoiseDefinition_Builder() {
				r.NoiseDefinition_Builder_Octaves()
			}

			return out
		},
	)
}

type Noise struct {
	traversal.Construct
	FirstOctave       int
	AdditionalOctaves []float64
}

func (n Noise) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (n Noise) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	//TODO implement me
	panic("implement me")
}
