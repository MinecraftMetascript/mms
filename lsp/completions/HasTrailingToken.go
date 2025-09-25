package completions

import (
	"github.com/antlr4-go/antlr/v4"
)

func CtxHasTrailingToken(ctx antlr.ParserRuleContext, trailing string) bool {
	children := ctx.GetChildren()
	if len(children) == 0 {
		return false
	}
	if r, ok := children[len(children)-1].(antlr.TerminalNode); ok {
		return r.GetText() == trailing
	} else {
		if r, ok := children[len(children)-1].(antlr.ParserRuleContext); ok {
			return CtxHasTrailingToken(r, trailing)
		}
	}
	return false
}
