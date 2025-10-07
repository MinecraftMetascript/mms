package lang

import (
	"math"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
)

var EmptyFn = spec.NewOverloadSpec(nil, nil, nil)
var SimpleNumberFn = spec.NewOverloadSpec(
	[]spec.ValueSpec{
		spec.NewNumberSpec(true),
	},
	nil,
	nil,
)

var MinBuilder = spec.NewFunctionSpec(
	"Min",
	SimpleNumberFn,
)

var MaxBuilder = spec.NewFunctionSpec(
	"Max",
	SimpleNumberFn,
)

var OffsetBuilder = spec.NewFunctionSpec(
	"Offset",
	SimpleNumberFn,
)

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
