package surface_rules

import (
	"encoding/json"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

type SurfaceConditionKind string

const (
	AboveSurfaceConditionKind     SurfaceConditionKind = "minecraft:above_preliminary_surface"
	BiomeConditionKind            SurfaceConditionKind = "minecraft:biome"
	HoleConditionKind             SurfaceConditionKind = "minecraft:hole"
	NoiseConditionKind            SurfaceConditionKind = "minecraft:noise"
	SteepConditionKind            SurfaceConditionKind = "minecraft:steep"
	StoneDepthConditionKind       SurfaceConditionKind = "minecraft:stone_depth"
	FreezingConditionKind         SurfaceConditionKind = "minecraft:temperature"
	VerticalGradientConditionKind SurfaceConditionKind = "minecraft:vertical_gradient"
	AboveWaterConditionKind       SurfaceConditionKind = "minecraft:water"
	YAboveConditionKind           SurfaceConditionKind = "minecraft:y_above"
	ReferenceConditionKind        SurfaceConditionKind = "mms:__reference"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceConditionDefinitionContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			def := ctx.(*grammar.SurfaceConditionDefinitionContext)

			return traversal.ConstructRegistry.Construct(
				def.SurfaceCondition().GetChild(0).(antlr.ParserRuleContext),
				currentNamespace,
				scope,
			)

		},
	)
}

func exportSurfaceCondition(symbol traversal.Symbol, rootDir *lib.FileTreeLike, condition json.Marshaler) error {
	data, err := json.Marshal(condition)
	if err != nil {
		return err
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("_debug", nil).
		MkDir("surface_conditions", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)

	return nil
}

func InvertCondition(condition traversal.Construct) traversal.Construct {
	return condition
}

type BaseSurfaceCondition struct {
	traversal.Construct
	Negate bool
}
