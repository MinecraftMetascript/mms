package lang

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lib"
)

var amplitudesBuilder = *spec.NewFunctionSpec("Amplitudes", spec.NewOverloadSpec(nil, spec.NewNumberSpec(true), nil))

var NoiseBlock = spec.NewBlockSpec(
	"Noise", []spec.ValueSpec{
		spec.NewFunctionSpec("Noise",
			spec.NewOverloadSpec(
				[]spec.ValueSpec{spec.NewNumberSpec(false)},
				nil,
				[]spec.FunctionSpec{amplitudesBuilder},
			),
		).SetExporter(NoiseExporter).SetKind("Noise"),
	},
)

var NoiseExporter = func(fn spec.FunctionNode, name string) *lib.FileTreeLike {
	if fn.Name != "Noise" {
		return nil
	}

	out := struct {
		FirstOctave int       `json:"first_octave"`
		Amplitudes  []float64 `json:"amplitudes"`
	}{}

	if len(fn.Arguments) < 1 {
		// TODO: ERROR
		return nil
	}
	if firstOctave, ok := fn.Arguments[0].(*spec.NumberNode); ok {
		out.FirstOctave = int(firstOctave.Value)
	}
	out.Amplitudes = make([]float64, len(fn.Arguments))
	if len(fn.Builders) < 1 {
		return nil
	}
	amplitudes := fn.Builders[0]
	for _, ampArg := range amplitudes.Arguments {
		amp := ampArg.(*spec.NumberNode)
		out.Amplitudes = append(out.Amplitudes, amp.Value)
	}

	x, _ := json.Marshal(out)
	fmt.Println(string(x))
	root := lib.
		NewDirLike("worldgen", nil)
	root.
		MkDir("noise", nil).
		MkFile(
			fmt.Sprintf("%s.json", name),
			string(x),
			nil,
		)
	return root
}
