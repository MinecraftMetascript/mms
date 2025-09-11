package traversal

import (
	"encoding/json"
	"fmt"
	"log"
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

type Factory[C antlr.ParserRuleContext] func(ctx C, ns string, scope *Scope) Construct

func Register[C antlr.ParserRuleContext](f Factory[C]) {
	ctxType := reflect.TypeFor[C]()
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem()
	}
	ConstructRegistry.constructs[ctxType] = func(base antlr.ParserRuleContext, ns string, scope *Scope) Construct {
		return f(base.(C), ns, scope)
	}
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

func ExtractInlineConstruct(
	ctx antlr.ParserRuleContext,
	namespace string,
	scope *Scope,
	label string,
) (Symbol, *Reference) {
	def := ConstructRegistry.Construct(ctx, namespace, scope)
	if def == nil {
		log.Printf("[Error] Failed to extract inline construct -- ConstructRegistry returned nil for %s\n", reflect.TypeOf(ctx).Elem().Name())
		scope.DiagnoseSemanticError(fmt.Sprintf("Failed to extract inline %s", label), ctx)
		return nil, nil
	}
	ref :=
		NewReference(
			fmt.Sprintf(
				"%s_%d_%d",
				label,
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
			),
			"mms_inline",
		)
	s := NewSymbol(
		RuleLocation(ctx.GetParent().(antlr.ParserRuleContext), scope.CurrentFile),
		RuleLocation(ctx, scope.CurrentFile),
		def,
		ref,
		label,
	)
	if _, ok := scope.Get(*ref); !ok {
		if err := scope.Register(s); err != nil {
			log.Println("Failed to extract inline construct -- ", err)
			scope.DiagnoseSemanticError(fmt.Sprintf("Failed to extract inline %s", label), ctx)
			return nil, nil
		}
	} else {
		scope.DiagnoseSemanticError(fmt.Sprintf("Found duplicate inline values. This should not be possible %s", label), ctx)
		return nil, nil
	}
	return s, ref
}
