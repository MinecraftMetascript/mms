package builder_chain

import (
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type BuilderChain[Target any] struct {
	fns map[reflect.Type]func(ctx any, target *Target, scope *traversal.Scope, namespace string)

	seen map[reflect.Type]bool
}

// Invoke looks up and calls the function registered for context type C.
// Returns (nil, false) if no function is registered for that type.
func Invoke[Target any](bc *BuilderChain[Target], ctx antlr.ParserRuleContext, target *Target, scope *traversal.Scope, namespace string) bool {
	ctxType := reflect.TypeOf(ctx)
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem()
	}

	fmt.Println(ctxType.Name())
	fn, ok := bc.fns[ctxType]
	if !ok {
		return false
	}
	if bc.seen[ctxType] {
		scope.DiagnoseSemanticWarning(
			fmt.Sprintf("Ignoring duplicate builder call %s", grammar.MinecraftMetascriptParserStaticData.RuleNames[ctx.GetRuleIndex()]),
			ctx,
		)
		return false
	} else {
		fn(ctx, target, scope, namespace)
		bc.seen[ctxType] = true
	}
	return true
}

func Require[Target any](bc *BuilderChain[Target], ctx antlr.ParserRuleContext, scope *traversal.Scope, required reflect.Type, label string) {
	if required.Kind() == reflect.Ptr {
		required = required.Elem()
	}
	if !bc.seen[required] {
		scope.DiagnoseSemanticError(
			fmt.Sprintf("Missing required builder call %s", label),
			ctx,
		)
	}
}

// A type-safe registration that pairs a context type with its handler for a given Target.
type Registration[Target any] struct {
	t  reflect.Type
	fn func(ctx any, target *Target, scope *traversal.Scope, namespace string)
}

// Build creates a type-safe Registration for a context type C.
func Build[C antlr.ParserRuleContext, Target any](fn func(ctx C, target *Target, scope *traversal.Scope, namespace string)) Registration[Target] {
	ctxType := reflect.TypeFor[C]()
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem()
	}
	return Registration[Target]{
		t: ctxType,
		fn: func(ctx any, target *Target, scope *traversal.Scope, namespace string) {
			fn(ctx.(C), target, scope, namespace)
		},
	}
}

func NewBuilderChain[Target any](regs ...Registration[Target]) *BuilderChain[Target] {
	bc := &BuilderChain[Target]{
		fns:  make(map[reflect.Type]func(ctx any, target *Target, scope *traversal.Scope, namespace string)),
		seen: make(map[reflect.Type]bool),
	}
	for _, r := range regs {
		bc.fns[r.t] = r.fn
	}
	return bc
}
