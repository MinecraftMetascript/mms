package lang

import (
	"math"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lang/unpack"
)

var EmptyFn = spec.NewOverloadSpec(nil, nil, nil)
var SimpleFloatFn = spec.NewOverloadSpec(
	[]spec.ValueSpec{
		spec.NewNumberSpec(true),
	},
	nil,
	nil,
)
var SimpleIntFn = spec.NewOverloadSpec(
	[]spec.ValueSpec{
		spec.NewNumberSpec(true),
	},
	nil,
	nil,
)

var ShiftX = spec.NewFunctionSpec("ShiftX", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil))
var ShiftY = spec.NewFunctionSpec("ShiftY", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil))
var ShiftZ = spec.NewFunctionSpec("ShiftZ", spec.NewOverloadSpec([]spec.ValueSpec{DensityFunctions}, nil, nil))

var XzScale = spec.NewFunctionSpec("XzScale", SimpleFloatFn)
var YScale = spec.NewFunctionSpec("YScale", SimpleFloatFn)
var XzFactor = spec.NewFunctionSpec("XzFactor", SimpleFloatFn)
var YFactor = spec.NewFunctionSpec("YFactor", SimpleFloatFn)
var SmearScaleMul = spec.NewFunctionSpec("SmearScaleMul", SimpleFloatFn)
var MinBuilder = spec.NewFunctionSpec("Min", SimpleFloatFn)
var MaxBuilder = spec.NewFunctionSpec("Max", SimpleFloatFn)
var OffsetBuilder = spec.NewFunctionSpec("Offset", SimpleFloatFn)

var VerticalAnchor = spec.NewValueSpecList(
	spec.NewNumberSpec(false),
	spec.NewFunctionSpec("Abs", SimpleFloatFn),
)

func init() {
	unpack.RegisterParser("VerticalAnchor", ParseAnchor)
}
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
