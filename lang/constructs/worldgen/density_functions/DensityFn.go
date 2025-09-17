package density_functions

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/constructs/primitives"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DensityFnFactory struct {
}

func (d DensityFnFactory) CreateDeclaration(ctx *grammar.DensityFnDeclarationContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	if ctx.Declare() != nil {
		declaration := primitives.DeclarationFactory{}.Create(ctx.Declare().(*grammar.DeclareContext), namespace, scope)
		if declaration == nil {
			scope.DiagnoseSemanticError("Missing identifier", ctx.Declare())
		}

		fn := ctx.DensityFn()
		if fn == nil {
			scope.DiagnoseSemanticError("Missing density function definition", ctx)
		}
		val := traversal.ConstructNode(fn.(*grammar.DensityFnContext), namespace, scope)

		return traversal.NewSymbol(
			declaration.GetNameLocation(),
			traversal.RuleLocation(fn, scope.CurrentFile),
			val,
			declaration.GetReference(),
			traversal.VerticalAnchor,
		), true

	} else {
		return nil, false
	}
}

func (d DensityFnFactory) Create(ctx *grammar.DensityFnContext, namespace string, scope *traversal.Scope) traversal.Node {
	targetCtx := ctx.GetChild(0)
	if targetCtx == nil {
		return nil
	}

	if target, ok := targetCtx.(antlr.ParserRuleContext); !ok {
		return nil
	} else {
		value := traversal.ConstructNode(target, namespace, scope)
		if math := ctx.DensityFn_Math(); math != nil {
			// This has attached maths
			arg2 := traversal.ConstructNode(math.DensityFn(), namespace, scope)

			var kind MathDensityFnKind
			if operation := math.GetChild(0); operation != nil {
				switch operation.(antlr.TerminalNode).GetText() {
				case "*":
					kind = MathDensityFn_Mul
				case "+":
					kind = MathDensityFn_Add
				default:
					scope.DiagnoseSemanticError("Unknown operator", ctx)
				}
			}
			return &MathDensityFn{
				Type: kind,
				Arg1: value,
				Arg2: arg2,
			}
		}

		return value
	}
}

func (d DensityFnFactory) Export(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, symbol.GetValue())
}

type InlineNoiseDensityFnFactory struct {
}

func (i InlineNoiseDensityFnFactory) Create(ctx *grammar.DensityFn_InlineNoiseContext, namespace string, scope *traversal.Scope) traversal.Node {
	if noiseCtx := ctx.Noise(); noiseCtx != nil {
		if node, ref := traversal.ExtractInlineConstruct(noiseCtx.NoiseDefinition(), namespace, scope, traversal.Noise); ref != nil {
			return node.GetValue()
		}
	}
	return nil
}

type DensityFnBlockFactory struct{}

func (n DensityFnBlockFactory) CreateDeclaration(ctx traversal.DeclarableContext, namespace string, scope *traversal.Scope) (traversal.Symbol, bool) {
	return nil, false
}

func (n DensityFnBlockFactory) Create(ctx *grammar.DensityFnBlockContext, namespace string, scope *traversal.Scope) *DensityFnBlock {
	out := &DensityFnBlock{
		location:     traversal.RuleLocation(ctx, scope.CurrentFile),
		declarations: make([]traversal.Symbol, 0),
	}
	for _, fn := range ctx.AllDensityFnDeclaration() {
		res, ok := traversal.DeclareNode(fn, namespace, scope)
		if ok {
			out.declarations = append(out.declarations, res)
		}
	}
	return out
}

type DensityFnBlock struct {
	location     traversal.TextLocation
	declarations []traversal.Symbol
}

func (d DensityFnBlock) GetLocation() traversal.TextLocation {
	return d.location
}

func init() {
	traversal.RegisterDeclarableNodeFactory(DensityFnFactory{})
	traversal.RegisterNodeFactory(InlineNoiseDensityFnFactory{})
	traversal.RegisterNodeFactory(DensityFnBlockFactory{})
}

func exportDensityFunction(symbol traversal.Symbol, rootDir *lib.FileTreeLike, condition traversal.Node) error {
	data, err := json.MarshalIndent(condition, "", "  ")
	if err != nil {
		return err
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("worldgen", nil).
		MkDir("density_function", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)

	return nil
}
