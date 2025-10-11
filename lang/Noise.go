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

var InlinedNoise = spec.NewValueSpecList(spec.NewReferenceSpec(ast.SymbolNoise), noiseFn)

func init() {
	unpack.RegisterParser("InlineNoise", func(node ast.Node) any { return GetInlinedNoiseRef(node) })
}

func GetInlinedNoiseRef(v ast.Node) string {
	if v == nil {
		return ""
	}
	if val := spec.GetReferenceNodeValue(v, ast.SymbolNoise); !lib.IsNilInterface(val) {
		return *val
	} else if fn, ok := v.(*spec.FunctionNode); ok && fn != nil {
		return fn.Ref()
	}
	return ""
}

func ExtractInlineNoiseSymbol(argIdx int) func(spec.FunctionNode) []ast.Symbol {
	return func(fn spec.FunctionNode) []ast.Symbol {
		if argIdx < 0 || argIdx >= len(fn.Arguments) {
			return nil
		}
		if inlineNoise, ok := fn.Arguments[argIdx].(*spec.FunctionNode); ok {
			return []ast.Symbol{
				inlineNoise,
			}
		}
		return nil
	}
}
