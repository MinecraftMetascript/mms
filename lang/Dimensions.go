package lang

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lang/unpack"
	"github.com/minecraftmetascript/mms/lib"
)

func exporter[T any](serializer func(node T) any, rootPath []string, extension string) func(fn T, name string) *lib.FileTreeLike {
	return func(fn T, name string) *lib.FileTreeLike {
		content := serializer(fn)
		contentBytes, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			log.Println("Error marshalling: ", err)
			return nil
		}

		if len(rootPath) == 0 {
			log.Println("exporter requires at least one path segment")
			return nil
		}
		root := lib.NewDirLike(rootPath[0], nil)
		current := root
		for _, part := range rootPath[1:] {
			current = current.MkDir(part, nil)
		}
		current.MkFile(fmt.Sprintf("%s.%s", name, extension), string(contentBytes), nil)

		return root
	}
}

func exportWorldgen(serializer func(node spec.FunctionNode) any, subdir string) func(fn spec.FunctionNode, name string) *lib.FileTreeLike {
	return func(fn spec.FunctionNode, name string) *lib.FileTreeLike {
		content := serializer(fn)
		contentBytes, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			log.Println("Error marshalling: ", err)
			return nil
		}

		root := lib.
			NewDirLike("worldgen", nil)
		root.
			MkDir(subdir, nil).
			MkFile(
				fmt.Sprintf("%s.json", name),
				string(contentBytes),
				nil,
			)
		return root
	}
}

func serializeDimensionType(n spec.FunctionNode) any {
	out := &struct {
		Ultrawarm   bool `json:"ultrawarm" mms_builder:"Ultrawarm"`
		Natural     bool `json:"natural" mms_builder:"Natural"`
		Skylight    bool `json:"skylight" mms_builder:"Skylight"`
		Ceiling     bool `json:"ceiling" mms_builder:"Ceiling"`
		PiglinSafe  bool `json:"piglin_safe" mms_builder:"PiglinSafe"`
		BedsWork    bool `json:"beds_work" mms_builder:"BedsWork"`
		AnchorsWork bool `json:"anchors_work" mms_builder:"AnchorsWork"`
		HasRaids    bool `json:"has_raids" mms_builder:"HasRaids"`

		CoordinateScale float64 `json:"coordinate_scale" mms_builder:"CoordinateScale"`
		AmbientLight    float64 `json:"ambient_light" mms_builder:"AmbientLight"`
		FixedTime       float64 `json:"fixed_time,omitempty" mms_builder:"FixedTime"`
		// TODO: "Int Provider"
		MonsterLightLevel float64 `json:"monster_light_level" mms_builder:"MonsterLightLevel"`
		MonsterLightLimit float64 `json:"monster_light_limit" mms_builder:"MonsterLightLimit"`
		LogicalHeight     float64 `json:"logical_height" mms_builder:"LogicalHeight"`
		CloudHeight       float64 `json:"cloud_height" mms_builder:"CloudHeight"`
		MinY              float64 `json:"min_y" mms_builder:"MinY"`
		Height            float64 `json:"height" mms_builder:"Height"`

		Infiniburn string `json:"infiniburn,omitempty" mms_builder:"Infiniburn"`
		Effects    string `json:"effects,omitempty" mms_builder:"Effects"`
	}{}

	unpack.Builders(out, n.Builders)

	return out
}

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
			spec.NewFunctionSpec("CloudHeight", SimpleIntFn),
			spec.NewFunctionSpec("MinY", SimpleIntFn),
			spec.NewFunctionSpec("Height", SimpleIntFn),
			spec.NewFunctionSpec("Infiniburn", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewTagSpec(ast.TagBlock)}, nil, nil)),
			spec.NewFunctionSpec("Effects", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("minecraft:overworld", "minecraft:the_nether", "minecraft:the_end")}, nil, nil)),
		},
	),
).
	SetKind(ast.SymbolDimensionType).
	SetHelp("Defines a dimension type").
	SetOutputFn(serializeDimensionType).
	SetFileExporter(exportWorldgen(serializeDimensionType, "dimension_type"))

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
