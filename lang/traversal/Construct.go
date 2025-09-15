package traversal

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lib"
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
	factories map[reflect.Type]ConstructFactory
	help      map[reflect.Type]func(construct Construct, symbol Symbol, location TextLocation) *string

	constructs map[TextLocation]Construct
	symbols    map[TextLocation]Symbol
}

type Factory[C antlr.ParserRuleContext] func(ctx C, ns string, scope *Scope) Construct

func RegisterHelp[C antlr.ParserRuleContext](f func(construct Construct, symbol Symbol, location TextLocation) *string) {
	ctxType := reflect.TypeFor[C]()
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem()
	}
	ConstructRegistry.help[ctxType] = f
}

type Help struct {
	Content  string
	Position TextLocation
}

func GetHelp(ctx antlr.ParserRuleContext, filename string) *Help {
	t := reflect.TypeOf(ctx)
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	helper, ok := ConstructRegistry.help[t]
	if !ok {
		return nil
	}

	ctxLocation := RuleLocation(ctx, filename)
	constructs := FilterByLocation(ctxLocation, ConstructRegistry.constructs)
	//symbols := getByLocation(ctxLocation, ConstructRegistry.symbols)

	//for location, symbol := range symbols {
	//	construct := symbol.GetValue()
	//	if construct == nil {
	//		continue
	//	}
	//	val := helper(construct, symbol, location)
	//	if val == nil {
	//		continue
	//	}
	//
	//	return &Help{
	//		Content:  *val,
	//		Position: location,
	//	}
	//}

	for location, construct := range constructs {
		if construct == nil {
			continue
		}
		val := helper(construct, nil, location)
		if val == nil {
			continue
		}

		return &Help{
			Content:  *val,
			Position: location,
		}
	}

	return nil
}

func Register[C antlr.ParserRuleContext](f Factory[C]) {
	ctxType := reflect.TypeFor[C]()
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem()
	}
	ConstructRegistry.factories[ctxType] = func(base antlr.ParserRuleContext, ns string, scope *Scope) Construct {
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

	if factory, ok := r.factories[t]; ok {
		val := factory(ctx, currentNamespace, scope)
		r.constructs[RuleLocation(ctx, scope.CurrentFile)] = val
		return val
	}
	return nil
}

var ConstructRegistry = &constructRegistryImpl{
	factories:  make(map[reflect.Type]ConstructFactory),
	help:       make(map[reflect.Type]func(construct Construct, symbol Symbol, location TextLocation) *string),
	constructs: make(map[TextLocation]Construct),
	symbols:    make(map[TextLocation]Symbol),
}

func ExtractInlineConstruct(
	ctx antlr.ParserRuleContext,
	namespace string,
	scope *Scope,
	kind SymbolKind,
) (Symbol, *Reference) {
	def := ConstructNode(ctx, namespace, scope)
	if def == nil {
		log.Printf("[Error] Failed to extract inline construct -- ConstructRegistry returned nil for %s\n", reflect.TypeOf(ctx).Elem().Name())
		scope.DiagnoseSemanticError(fmt.Sprintf("Failed to extract inline %s", kind), ctx)
		return nil, nil
	}
	ref :=
		NewReference(
			fmt.Sprintf(
				"%s_%d_%d",
				kind,
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
		kind,
	)
	if _, ok := scope.Get(*ref); !ok {
		if err := scope.Register(s); err != nil {
			log.Println("Failed to extract inline construct -- ", err)
			scope.DiagnoseSemanticError(fmt.Sprintf("Failed to extract inline %s", kind), ctx)
			return nil, nil
		}
	} else {
		scope.DiagnoseSemanticError(fmt.Sprintf("Found duplicate inline values. This should not be possible %s", kind), ctx)
		return nil, nil
	}
	ConstructRegistry.symbols[RuleLocation(ctx, scope.CurrentFile)] = s
	return s, ref
}
