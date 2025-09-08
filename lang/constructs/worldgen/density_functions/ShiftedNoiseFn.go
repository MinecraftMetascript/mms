package density_functions

import (
	"encoding/json"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_ShiftedNoiseContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			shiftedNoiseBuilder := builder_chain.NewBuilderChain[ShiftedNoiseDensityFn](
				builder_chain.Build(
					func(ctx *grammar.Builder_XZScaleContext, target *ShiftedNoiseDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.XZScale = v }, scope, "XZScale")
					},
				), builder_chain.Build(
					func(ctx *grammar.Builder_YScaleContext, target *ShiftedNoiseDensityFn, scope *traversal.Scope, namespace string) {
						builder_chain.Builder_GetFloat(ctx, func(v float64) { target.YScale = v }, scope, "YScale")
					},
				), builder_chain.Build(
					func(ctx *grammar.Builder_ShiftContext, target *ShiftedNoiseDensityFn, scope *traversal.Scope, namespace string) {
						axis := ctx.Axis()
						if axis == nil {
							scope.DiagnoseSemanticError("Missing or invalid Axis", ctx)
							return
						}
						childFnCtx := ctx.DensityFn()
						if childFnCtx == nil {
							scope.DiagnoseSemanticError("Missing or invalid Density Function", ctx)
						}

						fn := traversal.ConstructRegistry.Construct(childFnCtx, namespace, scope)
						if fn == nil {
							scope.DiagnoseSemanticError("Missing or invalid Density Function", ctx)
						}

						switch axis.GetText() {
						case "x":
							target.ShiftX = fn
							break
						case "y":
							target.ShiftY = fn
							break
						case "z":
							target.ShiftZ = fn
							break
						default:
							scope.DiagnoseSemanticError("Invalid Axis", ctx)
							break
						}

					}),
			)
			densityFn := ctx.(*grammar.DensityFn_ShiftedNoiseContext)
			out := &ShiftedNoiseDensityFn{
				XZScale: 1,
				YScale:  1,
			}

			for _, r := range densityFn.AllDensityFn_ShiftedNoiseBuilder() {
				builder_chain.Invoke(shiftedNoiseBuilder, r.GetChild(0).(antlr.ParserRuleContext), out, scope, currentNamespace)
			}

			return out
		},
	)
}

type ShiftedNoiseDensityFn struct {
	XZScale float64
	YScale  float64
	ShiftX  traversal.Construct
	ShiftY  traversal.Construct
	ShiftZ  traversal.Construct
}

func (c ShiftedNoiseDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type    string              `json:"type"`
		XZScale float64             `json:"xz_scale"`
		YScale  float64             `json:"y_scale"`
		ShiftX  traversal.Construct `json:"shift_x,omitempty"`
		ShiftY  traversal.Construct `json:"shift_y,omitempty"`
		ShiftZ  traversal.Construct `json:"shift_z,omitempty"`
	}{
		Type:    "minecraft:shifted_noise",
		XZScale: c.XZScale,
		YScale:  c.YScale,
		ShiftX:  c.ShiftX,
		ShiftY:  c.ShiftY,
		ShiftZ:  c.ShiftZ,
	}, "", "  ")
}

func (c ShiftedNoiseDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}
