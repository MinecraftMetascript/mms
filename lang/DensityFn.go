package lang

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lib"
)

func DensityFnExporter(serializer func(node spec.FunctionNode) any) func(fn spec.FunctionNode, name string) *lib.FileTreeLike {
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
			MkDir("density_function", nil).
			MkFile(
				fmt.Sprintf("%s.json", name),
				string(contentBytes),
				nil,
			)
		return root
	}
}

func serializeUnary(kind string, s ast.Symbol) any {
	return struct {
		Kind     string `json:"type"`
		Argument any    `json:"argument"`
	}{
		Kind:     kind,
		Argument: s.ToSerializable(),
	}
}

func serializeBinary(kind string, s1 ast.Symbol, s2 ast.Symbol) any {
	return struct {
		Kind      string `json:"type"`
		Argument1 any    `json:"argument1"`
		Argument2 any    `json:"argument2"`
	}{
		Kind:      kind,
		Argument1: s1.ToSerializable(),
		Argument2: s2.ToSerializable(),
	}
}

func unaryDensityFn(label, kind string) spec.FunctionSpec {
	return spec.NewFunctionSpec(
		label, spec.NewOverloadSpec(
			[]spec.ValueSpec{DensityFunctions},
			nil,
			nil,
		)).SetOutputFn(func(n spec.FunctionNode) any {
		if s, ok := n.Arguments[0].(ast.Symbol); ok {
			return serializeUnary(kind, s)
		}
		return nil
	}).
		SetFileExporter(DensityFnExporter(func(n spec.FunctionNode) any {
			if s, ok := n.Arguments[0].(ast.Symbol); ok {
				return serializeUnary(kind, s)
			}
			return nil
		})).
		SetKind(ast.SymbolDensityFunction)
}

func binaryDensityFn(label, kind string) spec.FunctionSpec {
	return spec.NewFunctionSpec(
		label, spec.NewOverloadSpec(
			[]spec.ValueSpec{DensityFunctions, DensityFunctions},
			nil,
			nil,
		)).SetOutputFn(func(n spec.FunctionNode) any {
		if len(n.Arguments) < 2 {
			return nil
		}
		if arg1, ok := n.Arguments[0].(ast.Symbol); ok {
			if arg2, ok := n.Arguments[1].(ast.Symbol); ok {
				return serializeBinary(kind, arg1, arg2)
			}
		}
		return nil
	}).SetKind(ast.SymbolDensityFunction) // TODO: Serialize / Export
}

var DensityFunctions = spec.NewValueSpecList()

var Interpolated = unaryDensityFn("Interpolated", "minecraft:interpolated")
var Abs = unaryDensityFn("Abs", "minecraft:abs")
var Cube = unaryDensityFn("Cube", "minecraft:cube")
var Square = unaryDensityFn("Square", "minecraft:square")
var HalfNegative = unaryDensityFn("HalfNegative", "minecraft:half_negative")
var QuarterNegative = unaryDensityFn("QuarterNegative", "minecraft:quarter_negative")
var Squeeze = unaryDensityFn("Squeeze", "minecraft:squeeze")
var Invert = unaryDensityFn("Invert", "minecraft:invert")

var Add = binaryDensityFn("Add", "minecraft:add")
var Mul = binaryDensityFn("Mul", "minecraft:mul")
var Min = binaryDensityFn("Min", "minecraft:min")
var Max = binaryDensityFn("Max", "minecraft:max")

var Cache = spec.NewFunctionSpec(
	"Cache",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{
			spec.NewEnumSpec("flat", "_2d", "once", "cell"),
			DensityFunctions,
		},
		nil,
		nil,
	)).
	SetOutputFn(func(n spec.FunctionNode) any {
		if n.Arguments[0] == nil || len(n.Arguments) < 2 || n.Arguments[1] == nil {
			return nil
		}

		kind := "minecraft:"
		switch v := spec.GetEnumNodeValue(n.Arguments[0]); *v {
		case "flat":
			kind += "flat_cache"
		case "2d":
			kind += "cache_2d"
		case "once":
			kind += "cache_once"
		case "cell":
			kind += "cache_all_in_cell"
		}

		if s, ok := n.Arguments[1].(ast.Symbol); ok {
			return serializeUnary(kind, s)
		}
		return nil
	}).
	SetKind(ast.SymbolDensityFunction)

func serializeShift(n spec.FunctionNode) any {
	out := &struct {
		Type     string `json:"type"`
		Argument string `json:"argument"`
	}{}

	switch len(n.Arguments) {
	case 1:
		out.Type = "minecraft:shift"
		out.Argument = GetInlinedNoiseRef(n.Arguments[0])
	case 2:
		if v := spec.GetEnumNodeValue(n.Arguments[0]); v != nil {
			switch *v {
			case "a":
				out.Type = "minecraft:shift_a"
			case "b":
				out.Type = "minecraft:shift_b"
			default:
				out.Type = "minecraft:shift"
			}
		} else {
			out.Type = "minecraft:shift"
		}
		out.Argument = GetInlinedNoiseRef(n.Arguments[1])
	default:
		out.Type = "minecraft:shift"
	}

	return out
}

var Shift = spec.NewFunctionSpec(
	"Shift",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{spec.NewEnumSpec("a", "b"), InlinedNoise}, nil, nil,
	),
	spec.NewOverloadSpec(
		[]spec.ValueSpec{InlinedNoise}, nil, nil,
	),
).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeShift).
	SetFileExporter(DensityFnExporter(serializeShift)).
	// Inline noise can be at arg 1 (when enum present) or arg 0
	SetSymbolExtractor(func(fn spec.FunctionNode) []ast.Symbol {
		idx := 0
		if len(fn.Arguments) == 2 {
			idx = 1
		}
		if idx < len(fn.Arguments) {
			if inline, ok := fn.Arguments[idx].(*spec.FunctionNode); ok {
				return []ast.Symbol{inline}
			}
		}
		return nil
	})

func serializeOldBlendedNoise(n spec.FunctionNode) any {
	out := &struct {
		Type                 string  `json:"type"`
		XzScale              float64 `json:"xz_scale"`
		YScale               float64 `json:"y_scale"`
		XzFactor             float64 `json:"xz_factor"`
		YFactor              float64 `json:"y_factor"`
		SmearScaleMultiplier float64 `json:"smear_scale_multiplier"`
	}{
		Type:                 "minecraft:old_blended_noise",
		XzScale:              1.0,
		YScale:               1.0,
		XzFactor:             1.0,
		YFactor:              1.0,
		SmearScaleMultiplier: 1.0,
	}
	for _, b := range n.Builders {
		switch b.Name {
		case XzScale.Name:
			if val := spec.GetNumberNodeValue(b.Arguments[0]); !lib.IsNilInterface(val) {
				out.XzScale = *val
			}
		case YScale.Name:
			if val := spec.GetNumberNodeValue(b.Arguments[0]); !lib.IsNilInterface(val) {
				out.YScale = *val
			}
		case XzFactor.Name:
			if val := spec.GetNumberNodeValue(b.Arguments[0]); !lib.IsNilInterface(val) {
				out.XzFactor = *val
			}
		case YFactor.Name:
			if val := spec.GetNumberNodeValue(b.Arguments[0]); !lib.IsNilInterface(val) {
				out.YFactor = *val
			}
		case SmearScaleMul.Name:
			if val := spec.GetNumberNodeValue(b.Arguments[0]); !lib.IsNilInterface(val) {
				out.SmearScaleMultiplier = *val
			}
		}
	}
	return out
}

var OldBlendedNoise = spec.NewFunctionSpec("OldBlendedNoise", spec.NewOverloadSpec(nil, nil, []spec.FunctionSpec{
	XzScale, YScale, XzFactor, YFactor, SmearScaleMul,
})).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeOldBlendedNoise).
	SetFileExporter(DensityFnExporter(serializeOldBlendedNoise)).
	SetSymbolExtractor(func(fn spec.FunctionNode) []ast.Symbol {
		if len(fn.Arguments) < 1 {
			return nil
		}
		if inlineNoise, ok := fn.Arguments[0].(*spec.FunctionNode); ok {
			return []ast.Symbol{
				inlineNoise,
			}
		}
		return nil
	})

func serializeNoiseDensityFn(n spec.FunctionNode) any {
	out := &struct {
		Type    string  `json:"type"`
		Noise   string  `json:"noise"`
		XzScale float64 `json:"xz_scale" mms_builder:"XzScale"`
		YScale  float64 `json:"y_scale" mms_builder:"YScale"`
	}{
		Type:    "minecraft:noise",
		XzScale: 1.0,
		YScale:  1.0,
	}

	UnpackBuilders(out, n.Builders)
	out.Noise = GetInlinedNoiseRef(n.Arguments[0])

	return out
}

var Noise = spec.NewFunctionSpec("Noise", spec.NewOverloadSpec(
	[]spec.ValueSpec{InlinedNoise},
	nil,
	[]spec.FunctionSpec{XzScale, YScale}),
).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeNoiseDensityFn).
	SetFileExporter(DensityFnExporter(serializeNoiseDensityFn)).
	SetSymbolExtractor(ExtractInlineNoiseSymbol(0))

var EndIslands = spec.NewFunctionSpec("EndIslands", spec.NewOverloadSpec(nil, nil, nil)).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(SimpleSerializer("minecraft:end_islands")).
	SetFileExporter(DensityFnExporter(SimpleSerializer("minecraft:end_islands")))

func serializeWeirdScaledSampler(n spec.FunctionNode) any {
	out := &struct {
		Type        string `json:"type"`
		Noise       string `json:"noise"`
		RarityValue string `json:"rarity_value_mapper"`
		Input       any    `json:"input"`
	}{}
	out.Type = "minecraft:weird_scaled_sampler"
	if len(n.Arguments) > 0 {
		if val := spec.GetEnumNodeValue(n.Arguments[0]); !lib.IsNilInterface(val) {
			out.RarityValue = *val
		}
	}
	if len(n.Arguments) > 1 {
		out.Noise = GetInlinedNoiseRef(n.Arguments[1])
	}
	if len(n.Arguments) > 2 {
		if s, ok := n.Arguments[2].(ast.Symbol); ok {
			out.Input = s.ToSerializable()
		}
	}

	return out
}

var WeirdScaledSampler = spec.NewFunctionSpec("WeirdScaledSampler", spec.NewOverloadSpec(
	[]spec.ValueSpec{spec.NewEnumSpec("type_1", "type_2"), InlinedNoise, DensityFunctions}, nil, nil,
)).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeWeirdScaledSampler).
	SetFileExporter(DensityFnExporter(serializeWeirdScaledSampler)).
	SetSymbolExtractor(ExtractInlineNoiseSymbol(1))

func serializeShiftedNoise(n spec.FunctionNode) any {
	out := &struct {
		Type   string  `json:"type"`
		Noise  string  `json:"noise"`
		Xz     float64 `json:"xz_scale" mms_builder:"XzScale"`
		Y      float64 `json:"y_scale"  mms_builder:"YScale"`
		ShiftX any     `json:"shift_x"  mms_builder:"ShiftX"`
		ShiftY any     `json:"shift_y"  mms_builder:"ShiftY"`
		ShiftZ any     `json:"shift_z"  mms_builder:"ShiftZ"`
	}{
		Type: "minecraft:shifted_noise",
		Xz:   1.0,
		Y:    1.0,
	}
	UnpackBuilders(out, n.Builders)
	// arg0 = inlined noise (or ref)
	if len(n.Arguments) > 0 {
		out.Noise = GetInlinedNoiseRef(n.Arguments[0])
	}
	return out
}

var ShiftedNoise = spec.NewFunctionSpec(
	"ShiftedNoise",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{InlinedNoise},
		nil,
		[]spec.FunctionSpec{XzScale, YScale, ShiftX, ShiftY, ShiftZ},
	),
).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeShiftedNoise).
	SetFileExporter(DensityFnExporter(serializeShiftedNoise)).
	SetSymbolExtractor(ExtractInlineNoiseSymbol(0))

var RangeChoice = spec.NewFunctionSpec("RangeChoice", spec.NewOverloadSpec(nil, nil, nil)).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(func(n spec.FunctionNode) any { return nil }).
	SetFileExporter(func(n spec.FunctionNode, name string) *lib.FileTreeLike { return nil })

var Clamp = spec.NewFunctionSpec("Clamp", spec.NewOverloadSpec(nil, nil, nil)).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(func(n spec.FunctionNode) any { return nil }).
	SetFileExporter(func(n spec.FunctionNode, name string) *lib.FileTreeLike { return nil })

var Spline = spec.NewFunctionSpec("Spline", spec.NewOverloadSpec(nil, nil, nil)).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(func(n spec.FunctionNode) any { return nil }).
	SetFileExporter(func(n spec.FunctionNode, name string) *lib.FileTreeLike { return nil })

var YClampedGradient = spec.NewFunctionSpec("YClampedGradient", spec.NewOverloadSpec(nil, nil, nil)).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(func(n spec.FunctionNode) any { return nil }).
	SetFileExporter(func(n spec.FunctionNode, name string) *lib.FileTreeLike { return nil })

var DensityFnBlock spec.BlockSpec

func init() {
	DensityFunctions.Add(spec.NewNumberSpec(true))                         // constants
	DensityFunctions.Add(spec.NewReferenceSpec(ast.SymbolDensityFunction)) // constants
	DensityFunctions.Add(Interpolated)
	DensityFunctions.Add(Cache)
	DensityFunctions.Add(Shift)
	DensityFunctions.Add(Abs)
	DensityFunctions.Add(Cube)
	DensityFunctions.Add(Square)
	DensityFunctions.Add(HalfNegative)
	DensityFunctions.Add(QuarterNegative)
	DensityFunctions.Add(Squeeze)
	DensityFunctions.Add(Invert)
	DensityFunctions.Add(Add)
	DensityFunctions.Add(Mul)
	DensityFunctions.Add(Min)
	DensityFunctions.Add(Max)
	DensityFunctions.Add(OldBlendedNoise)
	DensityFunctions.Add(Noise)
	DensityFunctions.Add(EndIslands)
	DensityFunctions.Add(WeirdScaledSampler)
	DensityFunctions.Add(ShiftedNoise)
	DensityFunctions.Add(RangeChoice)
	DensityFunctions.Add(Clamp)
	DensityFunctions.Add(Spline)
	DensityFunctions.Add(YClampedGradient)

	DensityFnBlock = spec.NewBlockSpec(
		"Density",
		DensityFunctions.All(),
	)
	Blocks.Add(&DensityFnBlock)

}
