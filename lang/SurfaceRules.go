package lang

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
)

var (
	Bandlands = spec.NewFunctionSpec(
		"Bandlands",
		spec.NewOverloadSpec(nil, nil, nil),
	).SetKind(ast.SymbolSurfaceRule).SetHelp("Used in badlands to place terracotta.")

	Block = spec.NewFunctionSpec(
		"Block",
		spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNever),
		}, nil, nil),
	).SetKind(ast.SymbolSurfaceRule).SetHelp("Selects a block to place.")

	/// Conditions
	AboveSurface = spec.NewFunctionSpec(
		"AboveSurface",
		spec.NewOverloadSpec(nil, nil, nil),
	).SetKind(ast.SymbolSurfaceCondition).SetHelp("Checks if the current position is above the preliminary surface level, which is a Y-level usually a few blocks below the main surface, ignoring noise caves.")

	Biome = spec.NewFunctionSpec(
		"Biome",
		spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNever),
		}, spec.NewReferenceSpec(ast.SymbolNever), nil))

	Hole = spec.NewFunctionSpec(
		"Hole",
		spec.NewOverloadSpec(nil, nil, nil),
	).
		SetHelp("Passes for columns where the surface depth is 0.").
		SetKind(ast.SymbolSurfaceCondition)

	NoiseThreshold = spec.NewFunctionSpec(
		"NoiseThreshold",
		spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNoise),
			noiseFn,
		}, nil, []spec.FunctionSpec{MinBuilder, MaxBuilder}),
	).
		SetHelp("Passes for columns where the input noise is between the minimum and maximum values.").
		SetKind(ast.SymbolSurfaceCondition)
)

// TODO: Properly handle IF, SEQUENCE, and NOT

var SurfaceRuleBlock = spec.NewBlockSpec("Surface", []spec.ValueSpec{
	Bandlands, Block,

	AboveSurface, Biome, Hole, NoiseThreshold,
})
