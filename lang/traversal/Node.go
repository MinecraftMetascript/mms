package traversal

import (
	"reflect"
	"slices"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Node interface {
	GetLocation() TextLocation
}
type CompletableNode interface {
	GetCompletions(cursorPosition protocol.Position) []protocol.CompletionItem
}

type HelpfulNodeFactory[Target Node] interface {
	GetHelp(Target, Symbol, TextLocation) *Help
}

type NodeFactory[Result Node, Ctx antlr.ParserRuleContext] interface {
	Create(ctx Ctx, namespace string, scope *Scope) Result
}

type DeclarableFactory[DeclareCtx antlr.ParserRuleContext] interface {
	CreateDeclaration(ctx DeclareCtx, namespace string, scope *Scope) (Symbol, bool)
}

type RegisterDeclarableFactory[Result Node, Ctx antlr.ParserRuleContext, DeclareCtx antlr.ParserRuleContext] interface {
	NodeFactory[Result, Ctx]
	DeclarableFactory[DeclareCtx]
}

type ExportableNodeFactory interface {
	Export(symbol Symbol, rootDir *lib.FileTreeLike) error
}

type DeclarableContext interface {
	antlr.ParserRuleContext
	Declare() grammar.IDeclareContext
}

var factoriesByValueCtx = make(map[reflect.Type]NodeFactory[Node, antlr.ParserRuleContext])
var factoriesByDeclarationCtx = make(map[reflect.Type]NodeFactory[Node, antlr.ParserRuleContext])
var factoriesByResultType = make(map[reflect.Type]NodeFactory[Node, antlr.ParserRuleContext])

var nodesByLocation = make(map[TextLocation]Node)
var symbolsByLocation = make(map[TextLocation]Symbol)

func RemoveFile(filename string) {
	for loc := range nodesByLocation {
		if loc.Filename == filename {
			delete(nodesByLocation, loc)
		}
	}
	for loc := range symbolsByLocation {
		if loc.Filename == filename {
			delete(symbolsByLocation, loc)
		}
	}
}

func RegisterNodeFactory[Result Node, Ctx antlr.ParserRuleContext](f NodeFactory[Result, Ctx]) {
	ctxType := lib.DerefType(reflect.TypeFor[Ctx]())
	val := innerNodeFactory{
		create: func(ctx antlr.ParserRuleContext, namespace string, scope *Scope) Node {
			return f.Create(ctx.(Ctx), namespace, scope)
		},
		export: func(s Symbol, root *lib.FileTreeLike) error {
			if c, ok := f.(ExportableNodeFactory); ok {
				return c.Export(s, root)
			}
			return nil
		},
		declare: func(ctx DeclarableContext, namespace string, scope *Scope) (Symbol, bool) {
			return nil, false
		},
		getCompletions: func(node Node, location protocol.Position) []protocol.CompletionItem {
			if c, ok := node.(CompletableNode); ok {
				return c.GetCompletions(location)
			}
			return nil
		},
	}

	if helpful, ok := f.(HelpfulNodeFactory[Result]); ok {
		val.getHelp = func(node Node, symbol Symbol, location TextLocation) *Help {
			return helpful.GetHelp(node.(Result), symbol, location)
		}
	}

	resultType := reflect.TypeFor[Result]()
	if resultType.Kind() == reflect.Ptr {
		resultType = resultType.Elem()
	}

	factoriesByResultType[resultType] = val
	factoriesByValueCtx[ctxType] = val
}

func RegisterDeclarableNodeFactory[
	Result Node,
	Ctx antlr.ParserRuleContext,
	DeclareCtx DeclarableContext,
	F interface {
		NodeFactory[Result, Ctx]
		DeclarableFactory[DeclareCtx]
	},
](f F) {
	ctxType := lib.DerefType(reflect.TypeFor[Ctx]())
	val := innerNodeFactory{
		create: func(ctx antlr.ParserRuleContext, namespace string, scope *Scope) Node {
			return f.Create(ctx.(Ctx), namespace, scope)
		},
		export: func(s Symbol, root *lib.FileTreeLike) error {
			if c, ok := any(f).(ExportableNodeFactory); ok {
				return c.Export(s, root)
			}
			return nil
		},
		declare: func(ctx DeclarableContext, namespace string, scope *Scope) (Symbol, bool) {
			return f.CreateDeclaration(ctx.(DeclareCtx), namespace, scope)
		},
		getCompletions: func(node Node, location protocol.Position) []protocol.CompletionItem {
			if c, ok := node.(CompletableNode); ok {
				return c.GetCompletions(location)
			}
			return nil
		},
	}

	if helpful, ok := any(f).(HelpfulNodeFactory[Result]); ok {
		val.getHelp = func(node Node, symbol Symbol, location TextLocation) *Help {
			return helpful.GetHelp(node.(Result), symbol, location)
		}
	}

	declareCtxType := reflect.TypeFor[DeclareCtx]()
	if declareCtxType.Kind() == reflect.Ptr {
		declareCtxType = declareCtxType.Elem()
	}
	factoriesByDeclarationCtx[declareCtxType] = val

	resultType := reflect.TypeFor[Result]()
	if resultType.Kind() == reflect.Ptr {
		resultType = resultType.Elem()
	}

	factoriesByResultType[resultType] = val
	factoriesByValueCtx[ctxType] = val
}

func DeclareNode(ctx DeclarableContext, namespace string, scope *Scope) (Symbol, bool) {
	t := lib.DerefType(reflect.TypeOf(ctx))

	if factory, ok := factoriesByDeclarationCtx[t]; ok {
		res, ok := factory.(innerNodeFactory).CreateDeclaration(ctx, namespace, scope)
		if ok {
			if err := scope.Register(res); err == nil {
				location := RuleLocation(ctx, scope.CurrentFile)
				symbolsByLocation[location] = res
				nodesByLocation[res.GetContentLocation()] = res.GetValue()

				return res, true
			} else {
				scope.DiagnoseSemanticError(err.Error(), ctx)
			}
		}
	}
	return nil, false
}

func ConstructNode(ctx antlr.ParserRuleContext, namespace string, scope *Scope) Node {
	t := lib.DerefType(reflect.TypeOf(ctx))

	if factory, ok := factoriesByValueCtx[t]; ok {
		val := factory.(innerNodeFactory).Create(ctx, namespace, scope)
		if val != nil {
			location := RuleLocation(ctx, scope.CurrentFile)
			nodesByLocation[location] = val

			return val
		}
	}
	return nil
}

func GetNodeHelp(location TextLocation, filename string) *Help {
	symbol, ok := symbolsByLocation[location]
	if ok {
		node := symbol.GetValue()
		nodeType := lib.DerefType(reflect.TypeOf(node))

		if factory, ok := factoriesByResultType[nodeType]; ok {
			return factory.(innerNodeFactory).GetHelp(node, symbol, location)
		}

	}

	node, ok := nodesByLocation[location]
	if ok {
		nodeType := lib.DerefType(reflect.TypeOf(node))
		if factory, ok := factoriesByResultType[nodeType]; ok {
			return factory.(innerNodeFactory).GetHelp(node, nil, location)
		}
	}

	return nil
}

func getNodesAtPosition(line int, character int) []Node {
	out := make([]Node, 0)
	nodeLocations := make(map[Node]TextLocation)

	for loc, node := range nodesByLocation {
		lineMatch := loc.Start.Line <= line && loc.Stop.Line >= line
		colMatch := loc.Start.Col <= character && loc.Stop.Col >= character
		if lineMatch && colMatch {
			out = append(out, node)
			nodeLocations[node] = loc
		}
	}

	// Sort nodes by location span size (smaller spans first)
	slices.SortStableFunc(out, func(a, b Node) int {
		locA := nodeLocations[a]
		locB := nodeLocations[b]
		spanA := locA.Stop.Col - locA.Start.Col
		spanB := locB.Stop.Col - locB.Start.Col
		return spanA - spanB
	})

	return out
}

func GetCompletions(location protocol.Position, filename string) []protocol.CompletionItem {
	for col := int(location.Character); col >= 0; col-- {
		line := int(location.Line + 1)
		for _, node := range getNodesAtPosition(line, col) {
			if c, ok := node.(CompletableNode); ok {
				return c.GetCompletions(location)
			}
		}

	}
	// TODO: Implement Me
	return nil
}

func ExportNode(symbol Symbol, rootDir *lib.FileTreeLike) error {
	val := symbol.GetValue()
	t := lib.DerefType(reflect.TypeOf(val))

	factory, ok := factoriesByResultType[t]
	if !ok {
		return nil
	}
	exportable, ok := factory.(ExportableNodeFactory)
	if !ok {
		return nil
	}
	return exportable.Export(symbol, rootDir)
}

type innerNodeFactory struct {
	declare        func(DeclarableContext, string, *Scope) (Symbol, bool)
	create         func(antlr.ParserRuleContext, string, *Scope) Node
	getHelp        func(Node, Symbol, TextLocation) *Help
	export         func(Symbol, *lib.FileTreeLike) error
	getCompletions func(Node, protocol.Position) []protocol.CompletionItem
}

func (i innerNodeFactory) CreateDeclaration(ctx DeclarableContext, namespace string, scope *Scope) (Symbol, bool) {
	return i.declare(ctx, namespace, scope)
}
func (i innerNodeFactory) Create(ctx antlr.ParserRuleContext, namespace string, scope *Scope) Node {
	return i.create(ctx, namespace, scope)
}
func (i innerNodeFactory) GetHelp(res Node, s Symbol, l TextLocation) *Help {
	if i.getHelp == nil {
		return nil
	}
	return i.getHelp(res, s, l)
}
func (i innerNodeFactory) Export(symbol Symbol, rootDir *lib.FileTreeLike) error {
	return i.export(symbol, rootDir)
}
func (i innerNodeFactory) GetCompletions(res Node, cursorLocation protocol.Position) []protocol.CompletionItem {
	if i.getCompletions == nil {
		return nil
	}
	return i.getCompletions(res, cursorLocation)
}
