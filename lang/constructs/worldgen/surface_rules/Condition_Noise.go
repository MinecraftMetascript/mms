package surface_rules

import (
	"encoding/json"
	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_NoiseContext](),
		func(ctx_ antlr.ParserRuleContext, namespace string, scope *traversal.Scope) traversal.Construct {
			ctx := ctx_.(*grammar.SurfaceCondition_NoiseContext)
			minThresh, err := strconv.ParseFloat(ctx.Number(0).GetText(), 64)
			if err != nil {
				return nil
			}
			maxThresh, err := strconv.ParseFloat(ctx.Number(1).GetText(), 64)
			if err != nil {
				return nil
			}
			return &NoiseCondition{
				NoiseRef: *traversal.ConstructRegistry.Construct(ctx.ResourceReference(), namespace, scope).(*traversal.Reference),
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
	return json.Marshal(struct {
		Type     SurfaceConditionKind `json:"type"`
		NoiseRef traversal.Reference  `json:"noise_threshold"`
		Min      float64              `json:"min_threshold"`
		Max      float64              `json:"max_threshold"`
	}{
		Type:     NoiseConditionKind,
		NoiseRef: c.NoiseRef,
		Min:      c.Min,
		Max:      c.Max,
	})
}
