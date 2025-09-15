package getters

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type IntCtx interface {
	antlr.ParserRuleContext
	Int() antlr.TerminalNode
}

type IntNCtx interface {
	antlr.ParserRuleContext
	GetAllInt() []antlr.TerminalNode
}

func GetInt(
	ctx IntCtx,
	mod func(int),
	scope *traversal.Scope,
	label string) {
	if raw := ctx.Int(); raw != nil {
		if v, err := strconv.Atoi(raw.GetText()); err == nil {
			mod(v)
		} else {
			scope.DiagnoseSemanticError(
				fmt.Sprintf("%s must be an integer, got \"%s\"", label, raw.GetText()),
				ctx,
			)
		}
	} else {
		scope.DiagnoseSemanticError(
			fmt.Sprintf("Missing %s value", label),
			ctx,
		)
	}
}

func GetIntN(
	ctx IntNCtx,
	mod func(int),
	scope *traversal.Scope,
	label string,
	idx int,
) {
	ints := ctx.GetAllInt()
	if idx >= len(ints) {
		return
	}

	raw := ints[idx]

	if raw != nil {
		if v, err := strconv.Atoi(raw.GetText()); err == nil {
			mod(v)
		} else {
			scope.DiagnoseSemanticError(
				fmt.Sprintf("%s must be an integer, got \"%s\"", label, raw.GetText()),
				ctx,
			)
		}
	} else {
		scope.DiagnoseSemanticError(
			fmt.Sprintf("Missing %s value", label),
			ctx,
		)
	}
}
