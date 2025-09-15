package getters

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type FloatCtx interface {
	antlr.ParserRuleContext
	Number() grammar.INumberContext
}

type FloatNCtx interface {
	antlr.ParserRuleContext
	GetAllNumber() []grammar.INumberContext
}

func GetFloat(
	ctx FloatCtx,
	mod func(float64),
	scope *traversal.Scope,
	label string,
) {
	if raw := ctx.Number(); raw != nil {
		if v, err := strconv.ParseFloat(raw.GetText(), 64); err == nil {
			mod(v)
		} else {
			scope.DiagnoseSemanticError(
				fmt.Sprintf("%s must be a decimal, got \"%s\"", label, raw.GetText()),
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

func GetFloatN(
	ctx FloatNCtx,
	mod func(float64),
	scope *traversal.Scope,
	label string,
	idx int,
) {
	floats := ctx.GetAllNumber()
	if idx >= len(floats) {
		return
	}
	raw := floats[idx]

	if raw != nil {
		if v, err := strconv.ParseFloat(raw.GetText(), 64); err == nil {
			mod(v)
		} else {
			scope.DiagnoseSemanticError(
				fmt.Sprintf("%s must be a decimal, got \"%s\"", label, raw.GetText()),
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
