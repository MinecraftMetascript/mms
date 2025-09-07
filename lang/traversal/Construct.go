package traversal

import (
	"encoding/json"
	"reflect"

	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

type Construct interface {
	json.Marshaler

	ExportSymbol(symbol Symbol, rootDir *lib.FileTreeLike) error
}

type ConstructFactory func(
	ctx antlr.ParserRuleContext,
	currentNamespace string,
	scope *Scope,
) Construct

type BaseConstruct struct {
	Construct
}

type constructRegistryImpl struct {
	constructs map[reflect.Type]ConstructFactory
}

func (r *constructRegistryImpl) Register(ctxType reflect.Type, factory ConstructFactory) {
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem() // Dereference pointers for simplicity
	}
	r.constructs[ctxType] = factory
}

func (r *constructRegistryImpl) Construct(ctx antlr.ParserRuleContext, currentNamespace string, scope *Scope) Construct {
	t := reflect.TypeOf(ctx)
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if factory, ok := r.constructs[t]; ok {
		val := factory(ctx, currentNamespace, scope)
		return val
	}
	return nil
}

var ConstructRegistry = &constructRegistryImpl{
	constructs: make(map[reflect.Type]ConstructFactory),
}
