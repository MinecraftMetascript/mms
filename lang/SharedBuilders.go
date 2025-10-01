package lang

import "github.com/minecraftmetascript/mms/lang/spec"

var MinBuilder = spec.NewFunctionSpec("Min", spec.NewOverloadSpec([]spec.ValueSpec{
	spec.NewNumberSpec(true),
}, nil, nil))

var MaxBuilder = spec.NewFunctionSpec("Max", spec.NewOverloadSpec([]spec.ValueSpec{
	spec.NewNumberSpec(true),
}, nil, nil))
