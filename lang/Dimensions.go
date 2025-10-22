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

func serializeFlatSettings(n spec.FunctionNode) any {
	out := &struct {
		Type       string   `json:"type"`
		Layers     []any    `json:"layers,omitempty"`
		Structures []string `json:"structures,omitempty"`
		Lakes      []string `json:"lakes,omitempty"`
		Features   []string `json:"features,omitempty"`
		Biome      string   `json:"biome,omitempty"`
	}{
		Type: "minecraft:flat",
	}

	// Handle layers
	if layers := findBuilderByName(n.Builders, "Layers"); layers != nil {
		for _, layerNode := range layers.Arguments {
			if layer, ok := layerNode.(*spec.FunctionNode); ok && layer.Name == "FlatLayer" {
				layerObj := struct {
					Block  string `json:"block"`
					Height int    `json:"height"`
				}{}
				if len(layer.Arguments) > 0 {
					if ref := spec.GetReferenceNodeValue(layer.Arguments[0], ast.SymbolNever); ref != nil {
						layerObj.Block = *ref
					}
				}
				if len(layer.Arguments) > 1 {
					if height := spec.GetNumberNodeValue(layer.Arguments[1]); height != nil {
						layerObj.Height = int(*height)
					}
				}
				out.Layers = append(out.Layers, layerObj)
			}
		}
	}

	// Handle other settings
	if structures := findBuilderByName(n.Builders, "Structures"); structures != nil {
		for _, arg := range structures.Arguments {
			if ref := spec.GetReferenceNodeValue(arg, ast.SymbolNever); ref != nil {
				out.Structures = append(out.Structures, *ref)
			}
		}
	}

	if lakes := findBuilderByName(n.Builders, "Lakes"); lakes != nil {
		for _, arg := range lakes.Arguments {
			if ref := spec.GetReferenceNodeValue(arg, ast.SymbolNever); ref != nil {
				out.Lakes = append(out.Lakes, *ref)
			}
		}
	}

	if features := findBuilderByName(n.Builders, "Features"); features != nil {
		for _, arg := range features.Arguments {
			if ref := spec.GetReferenceNodeValue(arg, ast.SymbolNever); ref != nil {
				out.Features = append(out.Features, *ref)
			}
		}
	}

	if biome := findBuilderByName(n.Builders, "Biome"); biome != nil {
		if len(biome.Arguments) > 0 {
			if ref := spec.GetReferenceNodeValue(biome.Arguments[0], ast.SymbolNever); ref != nil {
				out.Biome = *ref
			}
		}
	}

	return out
}

func serializeBiomeSource(n spec.FunctionNode) any {
	if len(n.Arguments) == 0 {
		return nil
	}

	generatorType := spec.GetEnumNodeValue(n.Arguments[0])
	if generatorType == nil {
		return nil
	}

	switch *generatorType {
	case "checkerboard":
		return serializeCheckerboardBiomeSource(n)
	case "fixed":
		return serializeFixedBiomeSource(n)
	case "multi_noise":
		return serializeMultiNoiseBiomeSource(n)
	case "the_end":
		return struct {
			Type string `json:"type"`
		}{Type: "minecraft:the_end"}
	default:
		return nil
	}
}

func serializeCheckerboardBiomeSource(n spec.FunctionNode) any {
	out := &struct {
		Type   string   `json:"type"`
		Biomes []string `json:"biomes"`
		Scale  int      `json:"scale"`
	}{
		Type: "minecraft:checkerboard",
	}

	// Handle biomes
	if biomes := findBuilderByName(n.Builders, "Biomes"); biomes != nil {
		for _, arg := range biomes.Arguments {
			if ref := spec.GetReferenceNodeValue(arg, ast.SymbolNever); ref != nil {
				out.Biomes = append(out.Biomes, *ref)
			}
		}
	}

	// Handle scale
	if scale := findBuilderByName(n.Builders, "Scale"); scale != nil {
		if len(scale.Arguments) > 0 {
			if val := spec.GetNumberNodeValue(scale.Arguments[0]); val != nil {
				out.Scale = int(*val)
			}
		}
	}

	return out
}

func serializeFixedBiomeSource(n spec.FunctionNode) any {
	out := &struct {
		Type  string `json:"type"`
		Biome string `json:"biome"`
	}{
		Type: "minecraft:fixed",
	}

	if biome := findBuilderByName(n.Builders, "Biome"); biome != nil {
		if len(biome.Arguments) > 0 {
			if ref := spec.GetReferenceNodeValue(biome.Arguments[0], ast.SymbolNever); ref != nil {
				out.Biome = *ref
			}
		}
	}

	return out
}

func serializeNoiseSettings(n spec.FunctionNode) any {
	out := &struct {
		SeaLevel       *float64   `json:"sea_level,omitempty" mms_builder:"SeaLevel"`
		DisableMobGen  bool       `json:"disable_mob_generation,omitempty" mms_builder:"DisableMobGen"`
		EnableOreVeins bool       `json:"ore_veins_enabled,omitempty" mms_builder:"EnableOreVeins"`
		Aquifers       bool       `json:"aquifers_enabled,omitempty" mms_builder:"Aquifers"`
		DefaultBlock   *string    `json:"default_block,omitempty" mms_builder:"DefaultBlock"`
		DefaultFluid   *string    `json:"default_fluid,omitempty" mms_builder:"DefaultFluid"`
		MinY           *float64   `json:"min_y,omitempty" mms_builder:"MinY"`
		Height         *float64   `json:"height,omitempty" mms_builder:"Height"`
		Size           [2]float64 `json:"size,omitempty" mms_builder:"Size"`
		NoiseRouter    any        `json:"noise_router,omitempty" mms_builder:"NoiseRouter"`
		SurfaceRule    any        `json:"surface_rule,omitempty" mms_builder:"SurfaceRule"`
	}{}

	// Manual mapping
	for _, builder := range n.Builders {
		switch builder.Name {
		case "SeaLevel":
			if len(builder.Arguments) > 0 {
				if val := spec.GetNumberNodeValue(builder.Arguments[0]); val != nil {
					out.SeaLevel = val
				}
			}
		case "DisableMobGen":
			out.DisableMobGen = true
		case "EnableOreVeins":
			out.EnableOreVeins = true
		case "Aquifers":
			out.Aquifers = true
		case "DefaultBlock":
			if len(builder.Arguments) > 0 {
				if ref := spec.GetReferenceNodeValue(builder.Arguments[0], ast.SymbolNever); ref != nil {
					out.DefaultBlock = ref
				}
			}
		case "DefaultFluid":
			if len(builder.Arguments) > 0 {
				if ref := spec.GetReferenceNodeValue(builder.Arguments[0], ast.SymbolNever); ref != nil {
					out.DefaultFluid = ref
				}
			}
		case "MinY":
			if len(builder.Arguments) > 0 {
				if val := spec.GetNumberNodeValue(builder.Arguments[0]); val != nil {
					out.MinY = val
				}
			}
		case "Height":
			if len(builder.Arguments) > 0 {
				if val := spec.GetNumberNodeValue(builder.Arguments[0]); val != nil {
					out.Height = val
				}
			}
		case "Size":
			if len(builder.Arguments) >= 2 {
				if val1 := spec.GetNumberNodeValue(builder.Arguments[0]); val1 != nil {
					out.Size[0] = *val1
				}
				if val2 := spec.GetNumberNodeValue(builder.Arguments[1]); val2 != nil {
					out.Size[1] = *val2
				}
			}
		case "NoiseRouter":
			if len(builder.Arguments) > 0 {
				if ref := spec.GetReferenceNodeValue(builder.Arguments[0], ast.SymbolNoiseRouter); ref != nil {
					out.NoiseRouter = *ref
				} else if fn, ok := builder.Arguments[0].(*spec.FunctionNode); ok {
					out.NoiseRouter = serializeNoiseRouter(*fn)
				}
			}
		case "SurfaceRule":
			// For now, leave as nil
		}
	}

	return out
}

func serializeNoiseRouter(n spec.FunctionNode) any {
	out := map[string]any{}
	for _, builder := range n.Builders {
		if len(builder.Arguments) > 0 {
			if ref := spec.GetReferenceNodeValue(builder.Arguments[0], ast.SymbolDensityFunction); ref != nil {
				out[builder.Name] = *ref
			} else {
				out[builder.Name] = "placeholder" // Placeholder for density function serialization
			}
		}
	}
	return out
}

func serializeDimension(n spec.FunctionNode) any {
	out := &struct {
		Type      string `json:"type"`
		Generator any    `json:"generator"`
	}{}

	// Get type from first argument
	if len(n.Arguments) > 0 {
		if ref := spec.GetReferenceNodeValue(n.Arguments[0], ast.SymbolDimensionType); ref != nil {
			out.Type = *ref
		} else if enum := spec.GetEnumNodeValue(n.Arguments[0]); enum != nil {
			out.Type = *enum
		}
	}

	// Find generator builder
	if generator := findBuilderByName(n.Builders, "Generator"); generator != nil {
		if len(generator.Arguments) > 0 {
			genType := spec.GetEnumNodeValue(generator.Arguments[0])
			if genType != nil {
				switch *genType {
				case "debug":
					out.Generator = struct {
						Type string `json:"type"`
					}{Type: "minecraft:debug"}
				case "flat":
					flatOut := &struct {
						Type     string `json:"type"`
						Settings any    `json:"settings,omitempty"`
					}{Type: "minecraft:flat"}
					if settings := findBuilderByName(generator.Builders, "Settings"); settings != nil {
						if len(settings.Arguments) > 0 {
							if fn, ok := settings.Arguments[0].(*spec.FunctionNode); ok {
								flatOut.Settings = serializeFlatSettings(*fn)
							} else if ref := spec.GetReferenceNodeValue(settings.Arguments[0], ast.SymbolNoiseSettings); ref != nil {
								flatOut.Settings = *ref
							}
						}
					}
					out.Generator = flatOut
				case "noise":
					noiseOut := &struct {
						Type        string `json:"type"`
						Settings    any    `json:"settings,omitempty"`
						BiomeSource any    `json:"biome_source,omitempty"`
					}{Type: "minecraft:noise"}
					if settings := findBuilderByName(generator.Builders, "Settings"); settings != nil {
						if len(settings.Arguments) > 0 {
							if fn, ok := settings.Arguments[0].(*spec.FunctionNode); ok {
								noiseOut.Settings = serializeNoiseSettings(*fn)
							} else if ref := spec.GetReferenceNodeValue(settings.Arguments[0], ast.SymbolNoiseSettings); ref != nil {
								noiseOut.Settings = *ref
							}
						}
					}
					if biomeSource := findBuilderByName(generator.Builders, "BiomeSource"); biomeSource != nil {
						if len(biomeSource.Arguments) > 0 {
							if fn, ok := biomeSource.Arguments[0].(*spec.FunctionNode); ok {
								noiseOut.BiomeSource = serializeBiomeSource(*fn)
							} else if ref := spec.GetReferenceNodeValue(biomeSource.Arguments[0], ast.SymbolNever); ref != nil {
								noiseOut.BiomeSource = *ref
							}
						}
					}
					out.Generator = noiseOut
				}
			}
		}
	}

	return out
}

func serializeMultiNoiseBiomeSource(n spec.FunctionNode) any {
	out := &struct {
		Type   string `json:"type"`
		Preset string `json:"preset,omitempty"`
		Biomes []any  `json:"biomes,omitempty"`
	}{
		Type: "minecraft:multi_noise",
	}

	// Check for preset
	if preset := findBuilderByName(n.Builders, "Preset"); preset != nil {
		if len(preset.Arguments) > 0 {
			if val := spec.GetEnumNodeValue(preset.Arguments[0]); val != nil {
				out.Preset = *val
			}
		}
	}

	// Handle custom biomes
	if biomes := findBuilderByName(n.Builders, "Biomes"); biomes != nil {
		for _, arg := range biomes.Arguments {
			if biomeNode, ok := arg.(*spec.FunctionNode); ok && biomeNode.Name == "MultiNoiseBiome" {
				biomeObj := struct {
					Biome      string `json:"biome"`
					Parameters any    `json:"parameters,omitempty"`
				}{}

				if len(biomeNode.Arguments) > 0 {
					if ref := spec.GetReferenceNodeValue(biomeNode.Arguments[0], ast.SymbolNever); ref != nil {
						biomeObj.Biome = *ref
					}
				}

				if params := findBuilderByName(biomeNode.Builders, "Parameters"); params != nil {
					if len(params.Arguments) > 0 {
						if paramNode, ok := params.Arguments[0].(*spec.FunctionNode); ok && paramNode.Name == "MultiNoiseParameters" {
							biomeObj.Parameters = serializeMultiNoiseParameters(*paramNode)
						}
					}
				}

				out.Biomes = append(out.Biomes, biomeObj)
			}
		}
	}

	return out
}

func serializeMultiNoiseParameters(n spec.FunctionNode) any {
	out := &struct {
		Temperature     *float64 `json:"temperature,omitempty"`
		Humidity        *float64 `json:"humidity,omitempty"`
		Continentalness *float64 `json:"continentalness,omitempty"`
		Erosion         *float64 `json:"erosion,omitempty"`
		Weirdness       *float64 `json:"weirdness,omitempty"`
		Depth           *float64 `json:"depth,omitempty"`
		Offset          *float64 `json:"offset,omitempty"`
	}{}

	for _, builder := range n.Builders {
		if len(builder.Arguments) > 0 {
			if val := spec.GetNumberNodeValue(builder.Arguments[0]); val != nil {
				switch builder.Name {
				case "Temperature":
					out.Temperature = val
				case "Humidity":
					out.Humidity = val
				case "Continentalness":
					out.Continentalness = val
				case "Erosion":
					out.Erosion = val
				case "Weirdness":
					out.Weirdness = val
				case "Depth":
					out.Depth = val
				case "Offset":
					out.Offset = val
				}
			}
		}
	}

	return out
}

func findBuilderByName(builders []spec.FunctionNode, name string) *spec.FunctionNode {
	for _, builder := range builders {
		if builder.Name == name {
			return &builder
		}
	}
	return nil
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
			// Debug generator
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("debug").SetHelp("Define a debug mode world")}, nil, nil),
			// Flat generator
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("flat").SetHelp("Define a superflat world")}, nil, []spec.FunctionSpec{
				spec.NewFunctionSpec("Settings", spec.NewOverloadSpec([]spec.ValueSpec{flatSettings}, nil, nil)).SetHelp("Superflat settings"),
			}),
			// Noise generator
			spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("noise").SetHelp("Define a normal world")}, nil, []spec.FunctionSpec{
				spec.NewFunctionSpec("Settings", spec.NewOverloadSpec([]spec.ValueSpec{noiseSettings}, nil, nil)).SetHelp("Noise Settings - can be a reference or inline definition"),
				spec.NewFunctionSpec("BiomeSource", spec.NewOverloadSpec([]spec.ValueSpec{biomeSource}, nil, nil)).SetHelp("Biome source - defines how biomes are distributed"),
			}),
		),
	}),
).
	SetKind(ast.SymbolDimension).
	SetHelp("Defines a Minecraft dimension with a type and generator settings").SetOutputFn(serializeDimension).SetFileExporter(exportWorldgen(serializeDimension, "dimension"))

// TODO: Define proper AST symbols for flat settings and layers
// For now, using basic function specs without custom symbols
var flatSettings = spec.NewFunctionSpec(
	"FlatSettings",
	spec.NewOverloadSpec(nil, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Layers", spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewValueSpecList(
				spec.NewReferenceSpec(ast.SymbolNever),
			)}, nil, nil)),
		spec.NewFunctionSpec("Structures", spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewTagSpec(ast.TagStructure)}, nil, nil)),
		spec.NewFunctionSpec("Lakes", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewTagSpec(ast.TagBiome)}, nil, nil)),
		spec.NewFunctionSpec("Features", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewTagSpec(ast.TagStructure)}, nil, nil)), // Using Structure tag for now
		spec.NewFunctionSpec("Biome", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNever)}, nil, nil)),
	}),
).
	SetOutputFn(serializeFlatSettings).
	SetFileExporter(exportWorldgen(serializeFlatSettings, "flat"))

var flatLayer = spec.NewFunctionSpec(
	"FlatLayer",
	spec.NewOverloadSpec([]spec.ValueSpec{
		spec.NewReferenceSpec(ast.SymbolNever),
		spec.NewNumberSpec(false),
	}, nil, nil),
)

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
).SetKind(ast.SymbolNoiseSettings).SetOutputFn(serializeNoiseSettings).SetFileExporter(exportWorldgen(serializeNoiseSettings, "noise_settings"))

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
).SetOutputFn(serializeNoiseRouter)

func serializeMultiNoiseBiome(n spec.FunctionNode) any {
	out := &struct {
		Biome      string `json:"biome"`
		Parameters any    `json:"parameters,omitempty"`
	}{}

	if len(n.Arguments) > 0 {
		if ref := spec.GetReferenceNodeValue(n.Arguments[0], ast.SymbolNever); ref != nil {
			out.Biome = *ref
		}
	}

	if params := findBuilderByName(n.Builders, "Parameters"); params != nil {
		if len(params.Arguments) > 0 {
			if paramNode, ok := params.Arguments[0].(*spec.FunctionNode); ok && paramNode.Name == "MultiNoiseParameters" {
				out.Parameters = serializeMultiNoiseParameters(*paramNode)
			}
		}
	}

	return out
}

var multiNoiseBiome = spec.NewFunctionSpec(
	"MultiNoiseBiome",
	spec.NewOverloadSpec([]spec.ValueSpec{
		spec.NewReferenceSpec(ast.SymbolNever), // Biome
	}, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Parameters", spec.NewOverloadSpec([]spec.ValueSpec{multiNoiseParameters}, nil, nil)),
	}),
).
	SetOutputFn(serializeMultiNoiseBiome).
	SetFileExporter(exportWorldgen(serializeMultiNoiseBiome, "biome_source"))

var multiNoiseParameters = spec.NewFunctionSpec(
	"MultiNoiseParameters",
	spec.NewOverloadSpec(nil, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Temperature", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
		spec.NewFunctionSpec("Humidity", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
		spec.NewFunctionSpec("Continentalness", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
		spec.NewFunctionSpec("Erosion", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
		spec.NewFunctionSpec("Weirdness", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
		spec.NewFunctionSpec("Depth", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
		spec.NewFunctionSpec("Offset", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(true)}, nil, nil)),
	}),
).
	SetOutputFn(serializeMultiNoiseParameters).
	SetFileExporter(exportWorldgen(serializeMultiNoiseParameters, "biome_source"))

var biomeSource = spec.NewFunctionSpec("BiomeSource",
	// Checkerboard biome source
	spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("checkerboard")}, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Biomes", spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewValueSpecList(
				spec.NewReferenceSpec(ast.SymbolNever),
				spec.NewTagSpec(ast.TagBiome),
			)}, nil, nil)).SetHelp("Biome IDs or #biome tags to place in checkerboard pattern"),
		spec.NewFunctionSpec("Scale", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(false)}, nil, nil)).SetHelp("Size of checkerboard squares (0-62, defaults to 2)"),
	}),
	// Fixed biome source
	spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("fixed")}, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Biome", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNever)}, nil, nil)),
	}),
	// Multi-noise biome source
	spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("multi_noise")}, nil, []spec.FunctionSpec{
		spec.NewFunctionSpec("Preset", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("overworld", "nether")}, nil, nil)).SetHelp("Use a predefined biome layout (overworld or nether)"),
		spec.NewFunctionSpec("Biomes", spec.NewOverloadSpec([]spec.ValueSpec{
			spec.NewValueSpecList(multiNoiseBiome),
		}, nil, nil)).SetHelp("Custom list of biomes with their noise parameters"),
	}),
	// The End biome source
	spec.NewOverloadSpec([]spec.ValueSpec{spec.NewEnumSpec("the_end")}, nil, nil),
).
	SetOutputFn(serializeBiomeSource).
	SetFileExporter(exportWorldgen(serializeBiomeSource, "biome_source"))

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
