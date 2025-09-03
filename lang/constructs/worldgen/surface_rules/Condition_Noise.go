package surface_rules

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_NoiseContext](),
		func(ctx_ antlr.ParserRuleContext, namespace string, scope *traversal.Scope) traversal.Construct {
			ctx := ctx_.(*grammar.SurfaceCondition_NoiseContext)
			minThresh := 0.0
			if ctx.Number(0) != nil {
				if v, err := strconv.ParseFloat(ctx.Number(0).GetText(), 64); err == nil {
					minThresh = v
				}
			}
			maxThresh := 0.0
			if ctx.Number(1) != nil {
				if v, err := strconv.ParseFloat(ctx.Number(1).GetText(), 64); err == nil {
					maxThresh = v
				}
			}
			var noiseRef traversal.Reference
			if rr := ctx.ResourceReference(); rr != nil {
				if cons := traversal.ConstructRegistry.Construct(rr, namespace, scope); cons != nil {
					if r, ok := cons.(*traversal.Reference); ok && r != nil {
						noiseRef = *r
					}
				}
			}
			return &NoiseCondition{
				NoiseRef: noiseRef,
				Min:      minThresh,
				Max:      maxThresh,
			}
		},
	)
}

type NoiseCondition struct {
	traversal.Construct
	NoiseRef traversal.Reference
	Min      float64
	Max      float64
}

func (c NoiseCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c NoiseCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type     SurfaceConditionKind `json:"type"`
		NoiseRef traversal.Reference  `json:"noise_threshold"`
		Min      float64              `json:"min_threshold"`
		Max      float64              `json:"max_threshold"`
	}{
		Type:     NoiseConditionKind,
		NoiseRef: c.NoiseRef,
		Min:      c.Min,
		Max:      c.Max,
	}, "", "  ")
}
