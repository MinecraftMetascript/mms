package spec

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/ast"
	"github.com/minecraftmetascript/mms/lang/grammar"
)

type Float struct {
	Min float64
	Max float64
}

func (f Float) GetLabel() string {
	return "Float"
}

func (f Float) Validate(ctx grammar.IValueContext) (bool, error) {
	var rawVal string
	switch c := ctx.GetChild(0).(type) {
	case *grammar.NumberContext:
		rawVal = c.GetText()
	case antlr.TerminalNode:
		rawVal = c.GetText()
	default:
		return false, nil
	}
	_, e := strconv.ParseFloat(rawVal, 64)
	return e == nil, e
}

func getFloatValue() {

}

func getNumberValue(ctx grammar.INumberContext, possibilities []Float) *FloatNode {
	out := &FloatNode{}



}

type FloatNode struct {
	location ast.SourceLocation
	value    float64
}

func (fn FloatNode) GetLocation() ast.SourceLocation {
	return fn.location
}
func (fn FloatNode) GetValue() float64 {
	return fn.value
}
