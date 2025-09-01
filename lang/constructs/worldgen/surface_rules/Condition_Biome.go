package surface_rules

import (
	"encoding/json"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_BiomeContext](),
		func(ctx_ antlr.ParserRuleContext, namespace string, scope *traversal.Scope) traversal.Construct {
			ctx, ok := ctx_.(*grammar.SurfaceCondition_BiomeContext)
			if !ok {
				return nil
			}
			refs := make([]traversal.Reference, 0)
			for _, ref := range ctx.AllResourceReference() {
				// Biome references should default to minecraft when unqualified
				r := traversal.ConstructRegistry.Construct(ref, "minecraft", scope).(*traversal.Reference)
				refs = append(refs, *r)
			}
			return &BiomeCondition{
				biomes: refs,
			}
		},
	)
}

type BiomeCondition struct {
	traversal.Construct

	biomes []traversal.Reference
}

func (c BiomeCondition) Biomes() []traversal.Reference {
	return c.biomes
}

func (c BiomeCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c BiomeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type   SurfaceConditionKind  `json:"type"`
		Biomes []traversal.Reference `json:"biomes"`
	}{
		Type:   BiomeConditionKind,
		Biomes: c.biomes,
	})
}
