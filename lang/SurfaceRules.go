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

func SimpleSerializer(t string) func(fn spec.FunctionNode) any {
	return func(fn spec.FunctionNode) any {
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
		SurfaceExporter(SimpleSerializer("minecraft:bandlands")),
	).SetOutputFn(SimpleSerializer("minecraft:bandlands"))

var Block = spec.NewFunctionSpec(
	"Block",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNever).SetDefaultNamespace("minecraft"),
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
		if val := spec.GetReferenceNodeValue(node.Arguments[0], ast.SymbolNever); val != nil {
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
	SetOutputFn(SimpleSerializer("minecraft:above_preliminary_surface")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("minecraft:above_preliminary_surface")))

var Biome = spec.NewFunctionSpec(
	"Biome",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewReferenceSpec(ast.SymbolNever).SetDefaultNamespace("minecraft"),
		}, spec.NewReferenceSpec(ast.SymbolNever).SetDefaultNamespace("minecraft"),
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
		return "{ \"__\": \"MMS: Unable to serialisze\"}"
	}
	filters := make([]string, 0)
	for _, arg := range node.Arguments {
		if val := spec.GetReferenceNodeValue(arg, ast.SymbolNever); val != nil {
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
	SetOutputFn(SimpleSerializer("minecraft:hole")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("minecraft:hole")))

var NoiseThreshold = spec.NewFunctionSpec(
	"NoiseThreshold",
	// Completion will try to use the first overload, so we want to prioritize that one
	spec.NewOverloadSpec(
		[]spec.ValueSpec{spec.NewValueSpecList(spec.NewReferenceSpec(ast.SymbolNoise), noiseFn)}, nil, []spec.FunctionSpec{MinBuilder, MaxBuilder},
	),
).
	SetHelp("Passes for columns where the input noise is between the minimum and maximum values.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(noiseThresholdSerializer).
	SetFileExporter(
		SurfaceExporter(noiseThresholdSerializer),
	).
	SetSymbolExtractor(ExtractInlineNoiseSymbol(0))

func noiseThresholdSerializer(node spec.FunctionNode) any {
	if node.Name != "NoiseThreshold" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	if len(node.Arguments) < 1 {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	out := &struct {
		Type      string  `json:"type"`
		Noise     string  `json:"noise"`
		MinThresh float64 `json:"min_threshold" mms_builder:"Min"`
		MaxThresh float64 `json:"max_threshold" mms_builder:"Max"`
	}{
		Type: "minecraft:noise_threshold",
	}
	UnpackBuilders(out, node.Builders)
	out.Noise = GetInlinedNoiseRef(node.Arguments[0])

	return out
}

var Steep = spec.NewFunctionSpec(
	"Steep",
	EmptyFn,
).
	SetHelp("Checks if the current position is a steep face on the north or east sides of a mountain.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(SimpleSerializer("minecraft:steep")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("minecraft:steep")))

var Frozen = spec.NewFunctionSpec(
	"Frozen",
	EmptyFn,
).
	SetHelp("Checks if the current position is in a biome that can snow.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(SimpleSerializer("minecraft:temperature")).
	SetFileExporter(SurfaceExporter(SimpleSerializer("minecraft:temperature")))

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
	out := &struct {
		Type                string  `json:"type"`
		SurfaceType         string  `json:"surface_type"`
		Offset              float64 `json:"offset" mms_builder:"Offset"`
		AddDepth            bool    `json:"add_surface_depth" mms_builder:"AddSurfaceDepth"`
		SecondaryDepthRange float64 `json:"secondary_depth_range" mms_builder:"SecondaryDepthRange"`
	}{
		Type: "minecraft:stone_depth",
	}
	UnpackBuilders(out, node.Builders)
	if len(node.Arguments) > 0 {
		if val := spec.GetEnumNodeValue(node.Arguments[0]); val != nil {
			out.SurfaceType = *val
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
	out := &struct {
		Type            string  `json:"type"`
		AddStoneDepth   bool    `json:"add_stone_depth" mms_builder:"AddStoneDepth"`
		Offset          float64 `json:"offset" mms_builder:"Offset"`
		DepthMultiplier float64 `json:"surface_depth_multiplier" mms_builder:"DepthMultiplier"`
	}{
		Type: "minecraft:water",
	}
	UnpackBuilders(out, node.Builders)

	return out
}

var VerticalGradient = spec.NewFunctionSpec(
	"VerticalGradient",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewStringSpec(),
		},
		nil,
		[]spec.FunctionSpec{
			spec.NewFunctionSpec("Lower", spec.NewOverloadSpec([]spec.ValueSpec{VerticalAnchor}, nil, nil)),
			spec.NewFunctionSpec("Upper", spec.NewOverloadSpec([]spec.ValueSpec{VerticalAnchor}, nil, nil)),
		},
	),
).
	SetHelp("Compares the current Y position, with a messy transition, just like the deepslate and bedrock transitions.").
	SetKind(ast.SymbolSurfaceCondition).
	SetOutputFn(VerticalGradientSerializer).
	SetFileExporter(
		SurfaceExporter(VerticalGradientSerializer),
	)

func VerticalGradientSerializer(node spec.FunctionNode) any {
	if node.Name != "VerticalGradient" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	out := struct {
		Type       string `json:"type"`
		RandomName string `json:"random_name"`
		Lower      any    `json:"true_at_and_below" mms_builder:"Lower"`
		Upper      any    `json:"false_at_and_above" mms_builder:"Upper"`
	}{
		Type: "minecraft:vertical_gradient",
	}

	UnpackBuilders(&out, node.Builders)
	if len(node.Arguments) > 0 {
		if val := spec.GetStringNodeValue(node.Arguments[0]); val != nil {
			out.RandomName = *val
		}
	}

	return out
}

var YAbove = spec.NewFunctionSpec(
	"YAbove",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			VerticalAnchor,
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
	out := &struct {
		Type            string  `json:"type"`
		Anchor          any     `json:"anchor"`
		AddStoneDepth   bool    `json:"add_stone_depth" mms_builder:"AddStoneDepth"`
		DepthMultiplier float64 `json:"surface_depth_multiplier" mms_builder:"DepthMultiplier"`
	}{
		Type: "minecraft:y_above",
	}
	UnpackBuilders(out, node.Builders)
	if len(node.Arguments) > 0 {
		anchor := node.Arguments[0]

		out.Anchor = ParseAnchor(anchor)
	}

	return out
}

func init() {
	Conditional = spec.NewConditionSpec().
		SetKind(ast.SymbolSurfaceRule).
		SetHelp("Applies a rule based on a surface condition")
	SurfaceConditions = []spec.ValueSpec{AboveSurface, Biome, Hole, NoiseThreshold, Steep, StoneDepth, Frozen, Water, VerticalGradient, YAbove}
	SurfaceRules = []spec.ValueSpec{Bandlands, Block, Conditional}
	ListRule := spec.NewListSpec().
		SetOutputFn(
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
			}).
		SetKind(ast.SymbolSurfaceRule).
		SetHelp("Creates a list of rules, the first valid rule will be used.s")
	SurfaceRules = append(SurfaceRules, ListRule)

	Conditional.
		AddConditionOption(SurfaceConditions...).
		AddConditionOption(spec.NewReferenceSpec(ast.SymbolSurfaceCondition)).
		AddValueOption(SurfaceRules...).
		AddValueOption(spec.NewReferenceSpec(ast.SymbolSurfaceRule)).
		SetOutputFn(func(n spec.ConditionalNode) any {
			var val any = nil
			if n.Value != nil && n.Value.ToSerializable() != nil {
				val = n.Value.ToSerializable()
			}
			return SerializeConditional(n.Condition, val)
		})

	ListRule.
		AddValueOption(spec.NewReferenceSpec(ast.SymbolSurfaceRule)).
		AddValueOption(SurfaceRules...)

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
	case nil:
		// There is no condition yet
		return mkSurfaceRuleConditional(nil, value)
	default:
		// Atomic condition: single condition node
		return mkSurfaceRuleConditional(c.ToSerializable(), value)
	}
}
