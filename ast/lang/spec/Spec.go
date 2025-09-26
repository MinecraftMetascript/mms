package spec

import (
	"fmt"
	"log"

	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Value interface {
	Validate(ctx grammar.IValueContext) (bool, error)
	GetLabel() string
}

func getFnValue(ctx grammar.IFnContext, specs []Value) {
	funcName := ctx.Identifier()
	if funcName == nil {
		// TODO: ✏ Diagnose -- missing function identifier
		return
	}

	for _, v := range specs {
		if fn, ok := v.(Function); ok {
			if funcName.GetText() != fn.Name {
				fmt.Println(funcName.GetText(), "not", fn.Name)
				continue
			}
			log.Println("Found func:", funcName)
			fn.Process(ctx)
		}
	}
}

func GetValue(ctx grammar.IValueContext, specs []Value) {
	switch valueCtx := ctx.GetChild(0).(type) {
	case *grammar.FnContext:
		getFnValue(valueCtx, specs)
	}
}
