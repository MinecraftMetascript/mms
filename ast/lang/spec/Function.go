package spec

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Function struct {
	Name string
	// Args 2d array to allow overrides. (e.g. [["x", "y"], ["z"]] allows ("x", "z") and ("y", "z"))
	// Args nils are allowed but must be last (e.g. [["x", "y"], ["z", "nil"]] allows the above as well as ("x") and ("y"))
	// Args [[nil],["x"]] is invalid, because the nils must be last
	Args [][]Value
	// RestArg is a special argument that when non-nil and non-0-length allows any number of that argument to appear
	// RestArg at the end of the function.
	RestArg []Value

	Builders []Function
}

func (fc Function) GetLabel() string {
	return "Function"
}

func (fc Function) Validate(ctx grammar.IValueContext) (bool, error) {
	return false, nil
}

func (fc Function) argSpec(i int) []Value {
	if len(fc.Args) <= i {
		if fc.RestArg != nil && len(fc.RestArg) > 0 {
			return fc.RestArg
		}
		return nil
	}
	fmt.Println(fc.Args, fc.RestArg, i)
	return fc.Args[i]
}

func (fc Function) Process(ctx grammar.IFnContext) {
	args := ctx.AllValue()
	for i, arg := range args {
		spec := fc.argSpec(i)
		valid := false
		for _, opt := range spec {
			ok, _ := opt.Validate(arg)
			if ok {
				valid = true
				break
			}
		}
		if !valid {
			validTypes := make([]string, 0)
			for _, opt := range spec {
				validTypes = append(validTypes, opt.GetLabel())
			}
			fmt.Println("Invalid Argument, expected one of:", strings.Join(validTypes, ", "))
		}
	}

	builders := ctx.AllFn()

	builderSpecsAsValue := make([]Value, len(fc.Builders))
	for i, builder := range fc.Builders {
		builderSpecsAsValue[i] = builder
	}

	for _, builder := range builders {
		fmt.Println("Attempting to process a builder!")
		getFnValue(builder, builderSpecsAsValue)
	}
}
