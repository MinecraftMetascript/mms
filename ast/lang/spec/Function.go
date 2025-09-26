package spec

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/ast"
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

func getFnValue(ctx grammar.IFnContext, possibilities []Function) *FunctionNode {
	funcName := ctx.Identifier()
	if funcName == nil {
		// TODO: ✏ Diagnose -- missing function kind
		return nil
	}

	for _, possible := range possibilities {
		if possible.Name != funcName.GetText() {
			continue // not a match
		}
		// found a match
		out := &FunctionNode{
			Args:     make([]ValueNode, 0),
			Builders: make([]FunctionNode, 0),
		}

		args := ctx.AllValue()
		for i, arg := range args {
			spec := possible.argSpec(i)
			valid := false
			// TODO: We need to modify this to allow for error messages, but also
			// 			fallthrough when something doesn't match?
			for _, opt := range spec {
				ok, e := opt.Validate(arg)
				if ok {
					valid = true
					value := GetValue(arg, spec)
					fmt.Println("VALUE:::", value)
					break
				} else if e != nil {
					fmt.Println("Invalid Argument:", e)
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

		for _, builder := range builders {
			node := getFnValue(builder, possible.Builders)
			if node != nil {
				out.Builders = append(out.Builders, *node)
			}
		}
		return out
	}
	return nil
}

type FunctionNode struct {
	location ast.SourceLocation
	Args     []ValueNode
	Builders []FunctionNode
}

func (f FunctionNode) GetLocation() ast.SourceLocation {
	return f.location
}
