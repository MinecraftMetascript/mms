package spec

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

func FilterValues[T Value](in []Value) []T {
	outSpecs := make([]T, 0)
	for _, spec := range in {
		if function, ok := spec.(T); ok {
			outSpecs = append(outSpecs, function)
		}
	}
	return outSpecs
}

func GetValue(ctx grammar.IValueContext, specs []Value) ValueNode {
	switch valueCtx := ctx.GetChild(0).(type) {
	case *grammar.FnContext:
		return getFnValue(valueCtx, FilterValues[Function](specs))
	case *grammar.BlockContext:
		return getBlockValue(valueCtx, FilterValues[Block](specs))
	case *grammar.NumberContext:
		return getNumberValue(valueCtx, FilterValues[Float](specs))
	case antlr.TerminalNode:

	}

	return nil
}
