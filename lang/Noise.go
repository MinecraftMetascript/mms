package lang

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lang/spec"
	"github.com/minecraftmetascript/mms/lib"
)

var amplitudesBuilder = spec.NewFunctionSpec(
	"Amplitudes",
	spec.NewOverloadSpec([]spec.ValueSpec{
		spec.NewNumberSpec(true), // require at least 1
	}, spec.NewNumberSpec(true), nil),
).SetHelp("Defines the amplitudes for each noise")

var noiseFn = spec.NewFunctionSpec("Noise",
	spec.NewOverloadSpec(
		[]spec.ValueSpec{spec.NewNumberSpec(false)},
		nil,
		[]spec.FunctionSpec{amplitudesBuilder},
	),
).
	SetFileExporter(NoiseExporter).
	SetOutputFn(NoiseSerializer).
	SetHelp("Defines a noise function.").
	SetKind(ast.SymbolNoise)

var NoiseBlock = spec.NewBlockSpec(
	"Noise", []spec.ValueSpec{
		noiseFn,
	},
)

func init() {
	Blocks.Add(&NoiseBlock)
}

var NoiseSerializer = func(fn spec.FunctionNode) any {
	if fn.Name != "Noise" {
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}

	out := struct {
		FirstOctave int       `json:"firstOctave"`
		Amplitudes  []float64 `json:"amplitudes"`
	}{
		Amplitudes: make([]float64, 0),
	}

	if len(fn.Arguments) < 1 {
		// TODO: ERROR
		return "{ \"__\": \"MMS: Unable to serialize\"}"
	}
	if firstOctave, ok := fn.Arguments[0].(*spec.NumberNode); ok {
		out.FirstOctave = int(firstOctave.Value)
	}
	if len(fn.Builders) > 0 {
		amplitudes := fn.Builders[0]
		for _, ampArg := range amplitudes.Arguments {
			amp := ampArg.(*spec.NumberNode)
			out.Amplitudes = append(out.Amplitudes, amp.Value)
		}
	}

	return out
}

var NoiseExporter = func(fn spec.FunctionNode, name string) *lib.FileTreeLike {
	content := NoiseSerializer(fn)
	contentBytes, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		log.Println("Error marshalling noise: ", err)
		return nil
	}

	root := lib.
		NewDirLike("worldgen", nil)
	root.
		MkDir("noise", nil).
		MkFile(
			fmt.Sprintf("%s.json", name),
			string(contentBytes),
			nil,
		)
	return root
}
