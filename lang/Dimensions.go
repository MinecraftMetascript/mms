package lang

import (
	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
)

var dimensionType = spec.NewFunctionSpec(
	"DimensionType",
	spec.NewOverloadSpec(
		nil,
		nil,
		[]spec.FunctionSpec{
			spec.NewFunctionSpec("Ultrawarm", EmptyFn),
			spec.NewFunctionSpec("Natural", EmptyFn),
			spec.NewFunctionSpec("Skylight", EmptyFn),
			spec.NewFunctionSpec("Ceiling", EmptyFn),
			spec.NewFunctionSpec("PiglinSafe", EmptyFn),
			spec.NewFunctionSpec("BedsWork", EmptyFn),
			spec.NewFunctionSpec("AnchorsWork", EmptyFn),
			spec.NewFunctionSpec("HasRaids", EmptyFn),

			spec.NewFunctionSpec("CoordinateScale", SimpleIntFn),
			spec.NewFunctionSpec("AmbientLight", SimpleFloatFn),
			spec.NewFunctionSpec("FixedTime", SimpleIntFn),
			spec.NewFunctionSpec("MonsterLightLevel", SimpleIntFn),
			spec.NewFunctionSpec("MonsterLightLimit", SimpleIntFn),
			spec.NewFunctionSpec("LogicalHeight", SimpleIntFn),
			spec.NewFunctionSpec("CLoudHeight", SimpleIntFn),
			spec.NewFunctionSpec("MinY", SimpleIntFn),
			spec.NewFunctionSpec("Height", SimpleIntFn),
			// TODO: This is meant to be a tag, not a reference
			spec.NewFunctionSpec("Infiniburn", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNever)}, nil, nil)),
			spec.NewFunctionSpec("Effects", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("minecraft:overworld", "minecraft:the_nether", "minecraft:the_end")}, nil, nil)),
		},
	),
).
	SetKind(ast.SymbolDimensionType)

var dimension = spec.NewFunctionSpec(
	"Dimension",
	spec.NewOverloadSpec([]spec.ValueSpec{
		spec.NewValueSpecList(
			spec.NewReferenceSpec(ast.SymbolDimensionType),
			spec.NewEnumSpec("overworld", "the_nether", "the_end", "overworld_caves"),
		),
	}, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Generator",
			// TODO: Modify Function.go so that it presents all 3 enums as a value
			// We probably need to modify the completion behavior entirely to support all valid overloads, rather than only using one
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("debug").SetHelp("Define a debug mode world")}, nil, nil),
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("flat").SetHelp("Define a superflat world")}, nil, []spec.FunctionSpec{
				// TODO: Define
				spec.NewFunctionSpec("Settings").SetHelp("Superflat settings"),
			}),
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("noise").SetHelp("Define a normal world")}, nil, []spec.FunctionSpec{
				spec.NewFunctionSpec("Settings", spec.NewOverloadSpec([]spec.ValueSpec{noiseSettings}, nil, nil)).SetHelp("Noise Settings"),
				spec.NewFunctionSpec("Biomes", spec.NewOverloadSpec([]spec.ValueSpec{biomeSource}, nil, nil)).SetHelp("Biome source"),
			}),
		),
	}),
).
	SetKind(ast.SymbolDimension)

var noiseSettings = spec.NewFunctionSpec(
	"NoiseSettings",
	spec.NewOverloadSpec(nil, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("SeaLevel", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(false)}, nil, nil)),
		spec.NewFunctionSpec("DisableMobGen", EmptyFn),
		spec.NewFunctionSpec("EnableOreVeins", EmptyFn),
		spec.NewFunctionSpec("Aquifers", EmptyFn),
		spec.NewFunctionSpec("DefaultBlock", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNever)}, nil, nil)),
		spec.NewFunctionSpec("DefaultFluid", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNever)}, nil, nil)),
		// TODO: Requires "Noise Parameter"
		spec.NewFunctionSpec("SpawnTarget", spec.NewOverloadSpec(nil, nil, nil)),
		spec.NewFunctionSpec("MinY", SimpleFloatFn),
		spec.NewFunctionSpec("Height", SimpleFloatFn),
		spec.NewFunctionSpec("Size", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(false), spec.NewNumberSpec(false)}, nil, nil)),
		spec.NewFunctionSpec("NoiseRouter",
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNoiseRouter)}, nil, nil),
			noiseRouterDef,
		),
		spec.NewFunctionSpec("SurfaceRule", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewValueSpecList(SurfaceRules, spec.NewReferenceSpec(ast.SymbolSurfaceRule))}, nil, nil)),
	}),
).SetKind(ast.SymbolNoiseSettings)

func noiseRouterFn(name string) spec.FunctionSpec {
	return spec.NewFunctionSpec(name,
		spec.NewOverloadSpec([]spec.ValueSpec{
			DensityFunctions,
		}, nil, nil),
	).SetKind(ast.SymbolNoiseRouter)
}

var noiseRouterDef = spec.NewOverloadSpec(nil, nil, []spec.FunctionSpec{
	noiseRouterFn("PreliminarySurfaceLevel"),
	noiseRouterFn("FinalDensity"),
	noiseRouterFn("Barrier"),
	noiseRouterFn("FluidLevelFloodedness"),
	noiseRouterFn("FluidLevelSpread"),
	noiseRouterFn("Lava"),
	noiseRouterFn("VeinToggle"),
	noiseRouterFn("VeinRidged"),
	noiseRouterFn("VeinGap"),
	noiseRouterFn("Temperature"),
	noiseRouterFn("Vegetation"),
	noiseRouterFn("Continents"),
	noiseRouterFn("Erosion"),
	noiseRouterFn("Depth"),
	noiseRouterFn("Ridges"),
})
var noiseRouter = spec.NewFunctionSpec(
	"NoiseRouter",
	noiseRouterDef,
)

var biomeSource = spec.NewFunctionSpec("BiomeSource", EmptyFn)

var DimensionStatements = spec.NewValueSpecList()
var DimensionBlock spec.BlockSpec

func init() {
	DimensionStatements.Add(dimension)
	DimensionStatements.Add(noiseSettings)
	DimensionStatements.Add(dimensionType)
	DimensionStatements.Add(noiseRouter)
	DimensionBlock = spec.NewBlockSpec("Dimension", []spec.ValueSpec{DimensionStatements})

	Blocks.Add(&DimensionBlock)
}
