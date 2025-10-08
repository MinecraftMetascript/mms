package lang

import (
	"math"
	"reflect"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/samber/lo"
)

func UnpackBuilders[T any](v T, builders []spec.FunctionNode) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	mutableV := reflect.ValueOf(v)
	if mutableV.Kind() == reflect.Ptr {
		mutableV = mutableV.Elem()
	}
	for _, field := range reflect.VisibleFields(t) {
		builderKind := field.Tag.Get("mms_builder")
		if builderKind == "" {
			continue
		}
		mutableField := mutableV.FieldByName(field.Name)
		builder, ok := lo.Find(builders, func(item spec.FunctionNode) bool {
			return item.Name == builderKind
		})
		if !ok {
			continue
		}

		switch field.Type.Kind() {
		case reflect.Float64:
			if val := spec.GetNumberNodeValue(builder.Arguments[0]); val != nil {
				mutableField.SetFloat(*val)
			}
		case reflect.Bool:
			mutableField.SetBool(true)
		// TODO: Support for ranges and string variant types
		default:
			break
		}
	}

}

var EmptyFn = spec.NewOverloadSpec(nil, nil, nil)
var SimpleNumberFn = spec.NewOverloadSpec(
	[]spec.ValueSpec{
		spec.NewNumberSpec(true),
	},
	nil,
	nil,
)

var ShiftX = spec.NewFunctionSpec("ShiftX", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil))
var ShiftY = spec.NewFunctionSpec("ShiftY", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil))
var ShiftZ = spec.NewFunctionSpec("ShiftZ", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil))

var XzScale = spec.NewFunctionSpec("XzScale", SimpleNumberFn)
var YScale = spec.NewFunctionSpec("YScale", SimpleNumberFn)
var XzFactor = spec.NewFunctionSpec("XzFactor", SimpleNumberFn)
var YFactor = spec.NewFunctionSpec("YFactor", SimpleNumberFn)
var SmearScaleMul = spec.NewFunctionSpec("SmearScaleMul", SimpleNumberFn)
var MinBuilder = spec.NewFunctionSpec("Min", SimpleNumberFn)
var MaxBuilder = spec.NewFunctionSpec("Max", SimpleNumberFn)
var OffsetBuilder = spec.NewFunctionSpec("Offset", SimpleNumberFn)

var VerticalAnchor = spec.NewValueSpecList(
	spec.NewNumberSpec(false),
	spec.NewFunctionSpec("Abs", SimpleNumberFn),
)

func ParseAnchor(anchor ast.Node) any {
	if val := spec.GetNumberNodeValue(anchor); val != nil {
		if *val > 0 {
			return struct {
				V float64 `json:"above_bottom"`
			}{
				V: *val,
			}
		} else {
			return struct {
				V float64 `json:"below_top"`
			}{
				V: math.Abs(*val),
			}

		}
	} else if fn, ok := anchor.(*spec.FunctionNode); ok {
		if len(fn.Arguments) > 0 {
			if inner := spec.GetNumberNodeValue(fn.Arguments[0]); inner != nil {
				return struct {
					V float64 `json:"absolute"`
				}{
					V: *inner,
				}
			}
		}

	}
	return nil
}
