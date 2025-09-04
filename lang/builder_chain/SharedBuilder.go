package builder_chain

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/constructs/primitives"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type (
	IntBuilder interface {
		antlr.ParserRuleContext
		Int() antlr.TerminalNode
	}

	FloatBuilder interface {
		antlr.ParserRuleContext
		Number() grammar.INumberContext
	}

	VerticalAnchorBuilder interface {
		antlr.ParserRuleContext
		VerticalAnchor() grammar.IVerticalAnchorContext
	}
)

func SharedBuilder_Add(
	ctx *grammar.SharedBuilder_AddContext,
	mod func(bool),
) {
	if !ctx.IsEmpty() {
		mod(true)
	}
}

func Builder_GetInt(
	ctx IntBuilder,
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

func Builder_GetFloat(
	ctx FloatBuilder,
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

func Builder_GetVerticalAnchor(
	ctx VerticalAnchorBuilder,
	currentNamespace string,
	mod func(anchor primitives.VerticalAnchor),
	scope *traversal.Scope,
	label string,
) {
	if raw := ctx.VerticalAnchor(); raw != nil {
		anchor := traversal.ConstructRegistry.Construct(raw, currentNamespace, scope)
		if a, ok := anchor.(*primitives.VerticalAnchor); ok && a != nil {
			mod(*a)
		} else {
			scope.DiagnoseSemanticError(
				fmt.Sprintf("%s must be a vertical anchor, got \"%s\"", label, raw.GetText()),
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
