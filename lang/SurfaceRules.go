package lang

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lib"
)

func SurfaceExporter(serializer func(node spec.FunctionNode) any) func(fn spec.FunctionNode, name string) *lib.FileTreeLike {
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
			MkDir("debug", nil).
			MkDir("surface", nil).
			MkFile(
				fmt.Sprintf("%s.json", name),
				string(contentBytes),
				nil,
			)
		return root
	}
}

func SimpleSerializer(name, t string) func(fn spec.FunctionNode) any {
	return func(fn spec.FunctionNode) any {
		if fn.Name != name {
			return "{ \"__\": \"MMS: Unable to serialize\"}"
		}
		out := struct {
			Type string `json:"type"`
		}{
			Type: t,
		}
		return out
	}
}

var DepthMultiplier = spec.NewFunctionSpec("DepthMultiplier", SimpleNumberFn)
var AddStoneDepth = spec.NewFunctionSpec("AddStoneDepth", EmptyFn)

var Bandlands = spec.NewFunctionSpec(
	"Bandlands",
	EmptyFn,
).
	SetKind(ast.SymbolSurfaceRule).
	SetHelp("Used in badlands to place terracotta.").
	SetFileExporter(
		SurfaceExporter(SimpleSerializer("Bandlands", "minecraft:bandlands")),
	).SetOutputFn(SimpleSerializer("Bandlands", "minecraft:bandlands"))

var Block = spec.NewFunctionSpec(
	"Block",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNever),
		},
		nil,
		nil,
	),
).
	SetKind(ast.SymbolSurfaceRule).
	SetHelp("Selects a block to place.").
	SetOutputFn(BlockSerializer).
	SetFileExporter(SurfaceExporter(BlockSerializer))

func BlockSerializer(node spec.FunctionNode) any {
	out := struct {
		Type        string `json:"type"`
		ResultState struct {
			Name string `json:"Name"`
		} `json:"result_state"`
	}{
		Type: "minecraft:block",
		ResultState: struct {
			Name string `json:"Name"`
		}{
			Name: "minecraft:air",
		},
	}
	if len(node.Arguments) > 0 {
		if val := spec.GetStringNodeValue(node.Arguments[0]); val != nil {
			out.ResultState.Name = *val
		}
	}
	return out
}

var AboveSurface = spec.NewFunctionSpec(
	"AboveSurface",
	EmptyFn,
).
	SetKind(ast.SymbolSurfaceCondition).
	SetHelp("Checks if the current position is above the preliminary surface level, which is a Y-level usually a few blocks below the main surface, ignoring noise caves.").
	SetOutputFn(SimpleSerializer("AboveSurface", "minecraft:above_surface")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("AboveSurface", "minecraft:above_surface")))

var Biome = spec.NewFunctionSpec(
	"Biome",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNever),
		}, spec.NewReferenceSpec(ast.SymbolNever),
		nil,
	)).
	SetHelp("Passes for columns where the biome matches the specified biomes.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(BiomeSerializer).
	SetFileExporter(
		SurfaceExporter(BiomeSerializer),
	)

func BiomeSerializer(node spec.FunctionNode) any {
	if node.Name != "Biome" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	filters := make([]string, 0)
	for _, arg := range node.Arguments {
		if val := spec.GetReferenceNodeValue(arg); val != nil {
			filters = append(filters, *val)
		}
	}
	out := struct {
		Type   string   `json:"type"`
		Biomes []string `json:"biomes"`
	}{
		Type:   "minecraft:biome",
		Biomes: filters,
	}
	return out
}

var Hole = spec.NewFunctionSpec(
	"Hole",
	EmptyFn,
).
	SetHelp("Passes for columns where the surface depth is 0.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(SimpleSerializer("Hole", "minecraft:hole")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("Hole", "minecraft:hole")))

var NoiseThreshold = spec.NewFunctionSpec(
	"NoiseThreshold",
	// Completion will try to use the first overload, so we want to prioritize that one
	spec.NewOverloadSpec(
		[]spec.ValueSpec{spec.NewReferenceSpec(ast.SymbolNoise)}, nil, []spec.FunctionSpec{MinBuilder, MaxBuilder},
	),
	spec.NewOverloadSpec(
		[]spec.ValueSpec{noiseFn}, nil, []spec.FunctionSpec{MinBuilder, MaxBuilder},
	),
).
	SetHelp("Passes for columns where the input noise is between the minimum and maximum values.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(NoiseThresholdSerializer).
	SetFileExporter(
		SurfaceExporter(NoiseThresholdSerializer),
	)

func NoiseThresholdSerializer(node spec.FunctionNode) any {
	if node.Name != "NoiseThreshold" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	if len(node.Arguments) < 1 {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	out := struct {
		Type      string  `json:"type"`
		Noise     string  `json:"noise"`
		MinThresh float64 `json:"min_threshold"`
		MaxThresh float64 `json:"max_threshold"`
	}{
		Type: "minecraft:noise_threshold",
	}
	switch a := node.Arguments[0].(type) {
	case *spec.FunctionNode:
		log.Println("Found inline noise?")
		break
	case *spec.ReferenceNode:
		if a.Kind != ast.SymbolNoise {
			// TODO: Diagnose (?)
			break
		}
		out.Noise = a.String()
	}

	for _, b := range node.Builders {
		switch b.Name {
		case "Min":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.MinThresh = *val
			}
		case "Max":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.MaxThresh = *val
			}
		}
	}

	return out
}

var Steep = spec.NewFunctionSpec(
	"Steep",
	EmptyFn,
).
	SetHelp("Checks if the current position is a steep face on the north or east sides of a mountain.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(SimpleSerializer("Steep", "minecraft:steep")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("Steep", "minecraft:steep")))

var Frozen = spec.NewFunctionSpec(
	"Frozen",
	EmptyFn,
).
	SetHelp("Checks if the current position is in a biome that can snow.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(SimpleSerializer("Frozen", "minecraft:temperature")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("Frozen", "minecraft:temperature")))

var StoneDepth = spec.NewFunctionSpec(
	"StoneDepth",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewEnumSpec("floor", "ceiling"),
		},
		nil,
		[]spec.FunctionSpec{
			OffsetBuilder,
			spec.NewFunctionSpec("AddSurfaceDepth", EmptyFn),
			spec.NewFunctionSpec("SecondaryDepthRange", SimpleNumberFn),
		},
	),
).
	SetHelp("Checks if the current position is within a specified distance from the surface, either upward or downward, using terrain depth.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(StoneDepthSerializer).
	SetFileExporter(
		SurfaceExporter(StoneDepthSerializer),
	)

func StoneDepthSerializer(node spec.FunctionNode) any {
	if node.Name != "StoneDepth" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	out := struct {
		Type                string  `json:"type"`
		SurfaceType         string  `json:"surface_type"`
		Offset              float64 `json:"offset"`
		AddDepth            bool    `json:"add_surface_depth"`
		SecondaryDepthRange float64 `json:"secondary_depth_range"`
	}{
		Type: "minecraft:stone_depth",
	}
	if len(node.Arguments) > 0 {
		if val := spec.GetEnumNodeValue(node.Arguments[0]); val != nil {
			out.SurfaceType = *val
		}
	}
	for _, b := range node.Builders {
		switch b.Name {
		case "Offset":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.Offset = *val
			}
		case "AddSurfaceDepth":
			out.AddDepth = true
		case "SecondaryDepthRange":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.SecondaryDepthRange = *val
			}
		}
	}

	return out
}

var Water = spec.NewFunctionSpec(
	"Water",
	spec.NewOverloadSpec(
		nil,
		nil,
		[]spec.FunctionSpec{
			OffsetBuilder,
			DepthMultiplier,
			AddStoneDepth,
		},
	),
).
	SetHelp("Checks if the current position is above water, based on terrain depth.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(WaterSerializer).
	SetFileExporter(
		SurfaceExporter(WaterSerializer),
	)

func WaterSerializer(node spec.FunctionNode) any {
	if node.Name != "Water" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	out := struct {
		Type            string  `json:"type"`
		AddStoneDepth   bool    `json:"add_stone_depth"`
		Offset          float64 `json:"offset"`
		DepthMultiplier float64 `json:"surface_depth_multiplier"`
	}{
		Type: "minecraft:water",
	}
	for _, b := range node.Builders {
		switch b.Name {
		case "Offset":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.Offset = *val
			}
		case "DepthMultiplier":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.DepthMultiplier = *val
			}
		case "AddStoneDepth":
			out.AddStoneDepth = true
		}
	}

	return out
}

var YAbove = spec.NewFunctionSpec(
	"YAbove",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			/* TODO: Vertical Anchor */
		},
		nil,
		[]spec.FunctionSpec{
			DepthMultiplier,
			AddStoneDepth,
		},
	)).SetHelp("Checks if the current position is above a specified height (exclusive)").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(YAboveSerializer).
	SetFileExporter(
		SurfaceExporter(YAboveSerializer),
	)

func YAboveSerializer(node spec.FunctionNode) any {
	if node.Name != "YAbove" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	out := struct {
		Type            string  `json:"type"`
		Anchor          float64 `json:"anchor"`
		AddStoneDepth   bool    `json:"add_stone_depth"`
		DepthMultiplier float64 `json:"surface_depth_multiplier"`
	}{
		Type: "minecraft:y_above",
	}

	for _, b := range node.Builders {
		switch b.Name {
		case "AddStoneDepth":
			out.AddStoneDepth = true
		case "DepthMultiplier":
			if val := spec.GetNumberNodeValue(b.Arguments[0]); val != nil {
				out.DepthMultiplier = *val
			}
		}
	}
	return out
}

// TODO: Properly handle SEQUENCE
func init() {
	// TODO: Conditional serialization (AND / OR / NOT) :(
	Conditional = spec.NewConditionSpec()
	SurfaceConditions = []spec.ValueSpec{AboveSurface, Biome, Hole, NoiseThreshold, Steep, StoneDepth, Frozen, Water, YAbove, Conditional}
	SurfaceRules = []spec.ValueSpec{Bandlands, Block}
	ListRule := spec.NewListSpec(SurfaceRules...).SetOutputFn(
		func(node spec.ListNode) any {
			out := struct {
				Type     string `json:"type"`
				Sequence []any  `json:"sequence"`
			}{
				Type:     "minecraft:sequence",
				Sequence: make([]any, 0),
			}

			for _, val := range node.Values {
				if s, ok := val.(ast.Symbol); ok && !lib.IsNilInterface(s) {
					out.Sequence = append(out.Sequence, s.ToSerializable())
				}
			}

			return out
		})
	SurfaceRules = append(SurfaceRules, ListRule)
	ListRule.ValueOptions = append(ListRule.ValueOptions, spec.NewReferenceSpec(ast.SymbolSurfaceRule))

	Conditional.
		AddConditionOption(SurfaceConditions...).
		AddConditionOption(spec.NewReferenceSpec(ast.SymbolSurfaceCondition)).
		AddValueOption(SurfaceRules...).
		AddValueOption(spec.NewReferenceSpec(ast.SymbolSurfaceRule)).
		SetOutputFn(func(n spec.ConditionalNode) any {

			return SerializeConditional(n.Condition, n.Value.ToSerializable())

		})

	SurfaceRuleBlock = spec.NewBlockSpec(
		"Surface",
		slices.Concat(SurfaceRules, SurfaceConditions),
	)

	Blocks.Add(&SurfaceRuleBlock)
}

var SurfaceConditions []spec.ValueSpec
var SurfaceRules []spec.ValueSpec
var Conditional *spec.ConditionalSpec
var SurfaceRuleBlock spec.BlockSpec

type SurfaceRuleConditional struct {
	Type      string `json:"type"`
	Condition any    `json:"if_true"`
	Value     any    `json:"then_run"`
}

// Keep existing helpers:
func mkSurfaceRuleConditional(cond any, value any) SurfaceRuleConditional {
	return SurfaceRuleConditional{
		Type:      "minecraft:condition",
		Condition: cond,
		Value:     value,
	}
}

// NEW helper: build a minecraft:not condition wrapper for a single atomic condition
func notCondition(inner any) any {
	return struct {
		Type   string `json:"type"`
		Invert any    `json:"invert"`
	}{
		Type:   "minecraft:not",
		Invert: inner,
	}
}

// NEW helper: turn !(...) into an equivalent structure without top-level AND/OR,
// by applying De Morgan’s and recursing back through SerializeConditional.
func serializeNegated(cond ast.Symbol, then any) any {
	switch c := cond.(type) {
	case *spec.ConditionAndNode:
		// !(A && B) == (!A) || (!B)
		return SerializeConditional(&spec.ConditionOrNode{
			Left:  &spec.ConditionNegateNode{Condition: c.Left},
			Right: &spec.ConditionNegateNode{Condition: c.Right},
		}, then)

	case *spec.ConditionOrNode:
		// !(A || B) == (!A) && (!B)
		return SerializeConditional(&spec.ConditionAndNode{
			Left:  &spec.ConditionNegateNode{Condition: c.Left},
			Right: &spec.ConditionNegateNode{Condition: c.Right},
		}, then)

	case *spec.ConditionNegateNode:
		// !!A == A
		return SerializeConditional(c.Condition, then)

	default:
		// Atomic: wrap with minecraft:not and emit a single condition rule
		return mkSurfaceRuleConditional(notCondition(c.ToSerializable()), then)
	}
}

func SerializeConditional(cond ast.Symbol, value any) any {
	switch c := cond.(type) {
	case *spec.ConditionAndNode:
		// A && B  -> condition(A, condition(B, then))
		// Note: order (left then right) matches the tree; swap if you prefer right-associative.
		inner := SerializeConditional(c.Right, value)
		return SerializeConditional(c.Left, inner)

	case *spec.ConditionOrNode:
		// A || B  -> sequence(condition(A, then), condition(B, then))
		first := SerializeConditional(c.Left, value)
		second := SerializeConditional(c.Right, value)
		return struct {
			Type     string `json:"type"`
			Sequence []any  `json:"sequence"`
		}{
			Type:     "minecraft:sequence",
			Sequence: []any{first, second},
		}

	case *spec.ConditionNegateNode:
		// Push NOT down until it hits atomics
		return serializeNegated(c.Condition, value)

	default:
		// Atomic condition: single condition node
		return mkSurfaceRuleConditional(c.ToSerializable(), value)
	}
}
