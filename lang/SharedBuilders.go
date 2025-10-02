package lang

import "github.com/minecraftmetascript/mms/lang/spec"

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

