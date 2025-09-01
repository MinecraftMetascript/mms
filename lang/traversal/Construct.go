package traversal

import (
	"encoding/json"
	"fmt"
	"github.com/minecraftmetascript/mms/lib"
	"reflect"

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
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if factory, ok := r.constructs[t]; ok {
		return factory(ctx, currentNamespace, scope)
	} else {
		fmt.Println("No factory found for", t)
	}

	return nil
}

var ConstructRegistry = &constructRegistryImpl{
	constructs: make(map[reflect.Type]ConstructFactory),
}
