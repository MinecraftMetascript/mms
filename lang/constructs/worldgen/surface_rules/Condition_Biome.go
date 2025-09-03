package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

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
				if cons := traversal.ConstructRegistry.Construct(ref, "minecraft", scope); cons != nil {
					if r, ok := cons.(*traversal.Reference); ok && r != nil {
						refs = append(refs, *r)
					}
				}
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
	return json.MarshalIndent(struct {
		Type   SurfaceConditionKind  `json:"type"`
		Biomes []traversal.Reference `json:"biome_is"`
	}{
		Type:   BiomeConditionKind,
		Biomes: c.biomes,
	}, "", "  ")
}
