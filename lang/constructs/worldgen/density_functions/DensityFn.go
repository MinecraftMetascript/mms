package density_functions

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

type MathDensityFnKind string

const (
	MathDensityFn_Add MathDensityFnKind = "add"
	MathDensityFn_Mul MathDensityFnKind = "mul"
)

type MathDensityFn struct {
	Type       MathDensityFnKind
	Arg1, Arg2 traversal.Construct
}

func (m MathDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type MathDensityFnKind   `json:"type"`
		Arg1 traversal.Construct `json:"argument1"`
		Arg2 traversal.Construct `json:"argument2"`
	}{
		Type: m.Type,
		Arg1: m.Arg1,
		Arg2: m.Arg2,
	}, "", "  ")
}

func (m MathDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, m)
}

func ExtractInlinedNoise(noiseCtx *grammar.NoiseDefinitionContext, currentNamespace string, scope *traversal.Scope) *traversal.Reference {
	noiseDef := traversal.ConstructRegistry.Construct(noiseCtx, currentNamespace, scope)
	noiseRef := traversal.NewReference(
		fmt.Sprintf("densityfn_noise_%d_%d", noiseCtx.GetStart().GetLine(), noiseCtx.GetStart().GetColumn()),
		"mms_inline",
	)
	noiseSymbol := traversal.NewSymbol(
		traversal.RuleLocation(noiseCtx, scope.CurrentFile),
		traversal.RuleLocation(noiseCtx, scope.CurrentFile),
		noiseDef,
		noiseRef,
		"Noise",
	)
	if _, ok := scope.Get(*noiseRef); !ok {
		if err := scope.Register(noiseSymbol); err != nil {
			// TODO: Improve behavior
			panic(err)
		}
	}
	return noiseRef
}

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_InlineNoiseContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			inlineCtx, ok := ctx.(*grammar.DensityFn_InlineNoiseContext)

			if !ok {
				// TODO: Diagnose?
				return nil
			}

			if noiseInlineCtx := inlineCtx.Noise(); noiseInlineCtx != nil {
				return ExtractInlinedNoise(noiseInlineCtx.NoiseDefinition().(*grammar.NoiseDefinitionContext), currentNamespace, scope)
			}
			return nil
		},
	)

	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFnBlockContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			blockCtx := ctx.(*grammar.DensityFnBlockContext)
			for _, child := range blockCtx.GetChildren() {
				if rule, ok := child.(antlr.ParserRuleContext); ok {
					traversal.ConstructRegistry.Construct(rule, currentNamespace, scope)
				}
			}
			return nil
		},
	)
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFnDeclarationContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			declaration := ctx.(*grammar.DensityFnDeclarationContext)

			if child := declaration.DensityFn(); child == nil {
				scope.DiagnoseSemanticError("Missing density function definition", ctx)
			} else if targetCtx := child; targetCtx != nil {
				s := traversal.ProcessDeclaration(declaration, targetCtx.(antlr.ParserRuleContext), scope, currentNamespace, "DensityFunction")
				if s == nil {
					fmt.Println("Failed to process declaration")
					return nil
				}
				return s.GetValue()
			}
			return nil
		},
	)

	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFnContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			fnCtx := ctx.(*grammar.DensityFnContext)
			targetCtx := fnCtx.GetChild(0)
			if targetCtx == nil {
				return nil
			}

			if target, ok := targetCtx.(antlr.ParserRuleContext); !ok {
				return nil
			} else {
				value := traversal.ConstructRegistry.Construct(target, currentNamespace, scope)
				fmt.Println(ctx.GetText(), value, reflect.TypeOf(value).Elem().Name(), fnCtx.DensityFn_Math() != nil)
				if math := fnCtx.DensityFn_Math(); math != nil {
					// This has attached maths
					arg2 := traversal.ConstructRegistry.Construct(math.DensityFn(), currentNamespace, scope)

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
		})
}

func exportDensityFunction(symbol traversal.Symbol, rootDir *lib.FileTreeLike, condition json.Marshaler) error {
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
