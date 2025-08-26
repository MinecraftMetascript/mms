package surface

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_conditions"
	"github.com/minecraftmetascript/mms/lib"
)

func mkConditionSymbol(ctx *grammars.SurfaceConditionDeclarationContext, rule surface_conditions.SurfaceCondition, ref lib.Reference, file string) lib.Symbol[surface_conditions.SurfaceCondition] {
	return lib.Symbol[surface_conditions.SurfaceCondition]{
		Location: lib.GetRuleLocation(ctx),

		Ref:   ref,
		File:  file,
		Value: rule,
		Kind:  lib.SymbolKindSurfaceCondition,
	}
}

func (v *Visitor) ConstructSurfaceCondition(ctx *grammars.SurfaceConditionContext) (surface_conditions.SurfaceCondition, error) {
	var cond surface_conditions.SurfaceCondition
	var err error
	switch ctx := ctx.GetChild(0).(type) {
	case *grammars.SurfaceCondition_AboveWaterContext:
		cond, err = surface_conditions.NewAboveWaterCondition(ctx)
	case *grammars.SurfaceCondition_AboveSurfaceContext:
		cond, err = surface_conditions.NewAboveSurfaceCondition(ctx)
	case *grammars.SurfaceCondition_BiomeContext:
		cond, err = surface_conditions.NewBiomeCondition(ctx)
	case *grammars.SurfaceCondition_FreezingContext:
		cond, err = surface_conditions.NewFreezingCondition(ctx)
	case *grammars.SurfaceCondition_HoleContext:
		cond, err = surface_conditions.NewHoleCondition(ctx)
	case *grammars.SurfaceCondition_NoiseContext:
		cond, err = surface_conditions.NewNoiseCondition(ctx)
	case *grammars.SurfaceConditionReferenceContext:
		cond, err = surface_conditions.NewSurfaceConditionReference(ctx, v.namespace)
	case *grammars.SurfaceCondition_SteepContext:
		cond, err = surface_conditions.NewSteepCondition(ctx)
	case *grammars.SurfaceCondition_StoneDepthContext:
		cond, err = surface_conditions.NewStoneDepthCondition(ctx)
	case *grammars.SurfaceCondition_VerticalGradientContext:
		cond, err = surface_conditions.NewVerticalGradientCondition(ctx)
	case *grammars.SurfaceCondition_YAboveContext:
		cond, err = surface_conditions.NewYAboveCondition(ctx)
	case *grammars.SurfaceCondition_CompoundContext:
		cond, err = surface_conditions.NewCompoundCondition(ctx, v)
	}

	if cond == nil {
		return nil, nil
	}

	cond.SetLocation(
		lib.GetRuleLocation(ctx.GetRuleContext().(antlr.ParserRuleContext)),
	)

	if err != nil {
		return cond, err
	}

	return cond, nil
}
