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
		label,
		spec.NewOverloadSpec(
			[]spec.ValueSpec{DensityFunctions, DensityFunctions},
			nil,
			nil,
		)).
		SetKind(ast.SymbolDensityFunction).
		SetOutputFn(func(n spec.FunctionNode) any {
			if len(n.Arguments) < 2 {
				return nil
			}
			if arg1, ok := n.Arguments[0].(ast.Symbol); ok {
				if arg2, ok := n.Arguments[1].(ast.Symbol); ok {
					return serializeBinary(kind, arg1, arg2)
				}
			}
			return nil
		}).
		SetFileExporter(DensityFnExporter(func(n spec.FunctionNode) any {
			if len(n.Arguments) < 2 {
				return nil
			}
			if arg1, ok := n.Arguments[0].(ast.Symbol); ok {
				if arg2, ok := n.Arguments[1].(ast.Symbol); ok {
					return serializeBinary(kind, arg1, arg2)
				}
			}
			return nil
		}))
}

var DensityFunctions = spec.NewValueSpecList().SetLabel("DensityFunction")

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

	unpack.Builders(out, n.Builders)
	out.Noise = GetInlinedNoiseRef(n.Arguments[0])

	return out
}

var Noise = spec.NewFunctionSpec("Noise",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{InlinedNoise},
		nil,
		[]spec.FunctionSpec{XzScale, YScale},
	),
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
	unpack.Builders(out, n.Builders)
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

func serializeRangeChoice(n spec.FunctionNode) any {
	out := &struct {
		Type     string  `json:"type"`
		Input    any     `json:"input" mms_arg:"0" mms_type:"symbol,DensityFn|float"`
		Min      float64 `json:"min_inclusive" mms_builder:"Min"`
		Max      float64 `json:"max_exclusive" mms_builder:"Max"`
		InRange  any     `json:"when_in_range" mms_builder:"InRange" mms_type:"symbol,DensityFn"`
		OutRange any     `json:"when_out_of_range" mms_builder:"OutRange" mms_type:"symbol,DensityFn"`
	}{
		Type: "minecraft:range_choice",
	}
	unpack.Builders(out, n.Builders)
	unpack.Args(out, n.Arguments)
	return out
}

func getDensityFn(n spec.FunctionNode) any {
	switch arg := n.Arguments[0].(type) {
	case ast.Symbol:
		return arg.ToSerializable()
	case *spec.NumberNode:
		return spec.GetNumberNodeValue(arg)
	default:
		return nil
	}
}

var RangeChoice = spec.NewFunctionSpec("RangeChoice", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil,
	[]spec.FunctionSpec{
		Min, Max,
		spec.NewFunctionSpec("InRange", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil)).
			SetOutputFn(getDensityFn),
		spec.NewFunctionSpec("OutRange", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil)).
			SetOutputFn(getDensityFn),
	})).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeRangeChoice).
	SetFileExporter(DensityFnExporter(serializeRangeChoice))

func serializeClamp(n spec.FunctionNode) any {
	out := &struct {
		Type     string  `json:"type"`
		Argument any     `json:"argument" mms_arg:"0" mms_type:"symbol,DensityFn|float"`
		Min      float64 `json:"min" mms_builder:"Min"`
		Max      float64 `json:"max" mms_builder:"Max"`
	}{}
	out.Type = "minecraft:clamp"
	unpack.Builders(out, n.Builders)
	unpack.Args(out, n.Arguments)
	return out
}

var Clamp = spec.NewFunctionSpec("Clamp", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, []spec.FunctionSpec{MinBuilder, MaxBuilder})).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeClamp).
	SetFileExporter(DensityFnExporter(serializeClamp))

type splinePoint struct {
	Location   float64 `json:"location" mms_builder:"Location"`
	Derivative float64 `json:"derivative" mms_builder:"Derivative"`
	Value      any     `json:"value"` // We have to process these by hand?
}

type spline struct {
	Coordinate any           `json:"coordinate" mms_arg:"0" mms_type:"symbol,DensityFn|float"`
	Points     []splinePoint `json:"points"`
}

func serializeSpline(n spec.FunctionNode) any {
	out := &spline{
		Points: []splinePoint{},
	}
	// Get Coordinate
	unpack.Args(out, n.Arguments)

	// Get Points
	for _, builder := range n.Builders {
		point := &splinePoint{}

		unpack.Args(point, builder.Arguments)
		unpack.Builders(point, builder.Builders)
		out.Points = append(out.Points, *point)
	}

	return out
}

var Spline spec.FunctionSpec

type yClampedGradientStop struct {
	Y float64 `json:"y" mms_arg:"0"`
	V float64 `json:"v" mms_arg:"1"`
}

func serializeYClampedGradient(n spec.FunctionNode) any {
	out := &struct {
		Type  string  `json:"type"`
		FromY float64 `json:"from_y"`
		ToY   float64 `json:"to_y"`
		From  float64 `json:"from_value"`
		To    float64 `json:"to_value"`
	}{}
	out.Type = "minecraft:y_clamped_gradient"
	for _, builders := range n.Builders {
		v := &yClampedGradientStop{}
		unpack.Args(v, builders.Arguments)
		switch builders.Name {
		case "From":
			out.FromY = v.Y
			out.From = v.V
		case "To":
			out.ToY = v.Y
			out.To = v.V
		}
	}

	return out
}

var YClampedGradient = spec.NewFunctionSpec("YClampedGradient", spec.NewOverloadSpec(nil, nil, []spec.FunctionSpec{
	spec.NewFunctionSpec("From", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(false), spec.NewNumberSpec(true)}, nil, nil)),
	spec.NewFunctionSpec("To", spec.NewOverloadSpec([]spec.ValueSpec{spec.NewNumberSpec(false), spec.NewNumberSpec(true)}, nil, nil)),
})).
	SetKind(ast.SymbolDensityFunction).
	SetOutputFn(serializeYClampedGradient).
	SetFileExporter(DensityFnExporter(serializeYClampedGradient))

var DensityFnBlock spec.BlockSpec

func init() {
	splineRef := spec.NewValueSpecList()
	Spline = spec.NewFunctionSpec(
		"Spline",
		spec.NewOverloadSpec(
			[]spec.ValueSpec{DensityFunctions}, // Coordinate
			nil,
			[]spec.FunctionSpec{ // Points
				spec.NewFunctionSpec(
					"Point",
					spec.NewOverloadSpec(
						[]spec.ValueSpec{
							spec.NewNumberSpec(true), // Location
							spec.NewNumberSpec(true), // Derivative
							splineRef,                // Value
						},
						nil,
						nil,
					),
				),
			},
		)).
		SetKind(ast.SymbolDensityFunction).
		SetOutputFn(serializeSpline).
		SetFileExporter(DensityFnExporter(serializeSpline))
	splineRef.Add(Spline)

	DensityFunctions.Add(spec.NewNumberSpec(true).SetKind(ast.SymbolDensityFunction)) // constants
	DensityFunctions.Add(spec.NewReferenceSpec(ast.SymbolDensityFunction))            // constants
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
