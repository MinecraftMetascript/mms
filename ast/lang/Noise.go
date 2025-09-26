package lang

import (
	"math"

	"github.com/minecraftmetascript/mms/ast/lang/spec"
)

var NoiseBlock = spec.Block{
	Kind: "Noise",
	Values: []spec.Value{
		spec.Function{
			Name: "Noise",
			Args: [][]spec.Value{
				{spec.Int{
					Min: math.MinInt32,
					Max: 1,
				}},
			},
			Builders: []spec.Function{
				{
					Name:    "Amplitudes",
					RestArg: []spec.Value{spec.Float{}},
				},
			},
		},
	},
}
