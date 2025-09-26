package surface_rules

import (
	"encoding/json"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

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
	NotConditionKind              SurfaceConditionKind = "minecraft:not"
	ReferenceConditionKind        SurfaceConditionKind = "mms:__reference"
)

func init() {
	traversal.Register(func(blockCtx *grammar.SurfaceBlockContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
		for _, statementCtx := range blockCtx.GetChildren() {
			statement, ok := statementCtx.(antlr.ParserRuleContext)
			if !ok {
				continue
			}
			targetCtx, ok := statement.GetChild(0).(antlr.ParserRuleContext)
			if !ok {
				continue
			}
			traversal.ConstructRegistry.Construct(targetCtx, currentNamespace, scope)
		}
		return nil
	})
	traversal.Register(func(conditionCtx *grammar.SurfaceConditionContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
		if conditionCtx.GetChildCount() > 0 {
			return traversal.ConstructRegistry.Construct(
				conditionCtx.GetChild(0).(antlr.ParserRuleContext),
				currentNamespace,
				scope,
			)
		}
		return nil
	})
	traversal.Register(func(declarationContext *grammar.SurfaceConditionDeclarationContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
		s := traversal.ProcessDeclaration(declarationContext.Declare(), declarationContext.SurfaceCondition(), scope, currentNamespace, "SurfaceCondition")
		return s.GetValue()
	})
}

func exportSurfaceCondition(symbol traversal.Symbol, rootDir *lib.FileTreeLike, condition json.Marshaler) error {
	data, err := json.MarshalIndent(condition, "", "  ")
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
