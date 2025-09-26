package spec

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Int struct {
	Min int
	Max int
}

func (Int) GetLabel() string {
	return "Integer"
}

func (i Int) Validate(ctx grammar.IValueContext) (bool, error) {
	switch c := ctx.GetChild(0).(type) {
	case *grammar.NumberContext:
		if c.Float() != nil {
			return false, fmt.Errorf("int cannot be a float")
		}
		if valCtx := c.Int(); valCtx == nil {
			return false, fmt.Errorf("int cannot be nil")
		} else {
			val, e := strconv.Atoi(valCtx.GetText())
			if val < i.Min || val > i.Max {
				return false, fmt.Errorf("int must be between %d and %d", i.Min, i.Max)
			}
			return e == nil, e
		}
	case antlr.TerminalNode:
		fmt.Println(c.GetText())
	}
	return false, nil
}
