package density_functions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func buildSplinePoint(point *grammar.DensityFn_SplinePointContext, currentNamespace string, scope *traversal.Scope) *SplinePoint {
	out := &SplinePoint{}

	locationCtx := point.Number(0)
	if locationCtx == nil {
		scope.DiagnoseSemanticError("Missing spline location", point)
		return nil
	} else {
		if location, ok := locationCtx.(*grammar.NumberContext); ok {
			if val, err := strconv.ParseFloat(location.GetText(), 64); err == nil {
				out.Location = val
			} else {
				scope.DiagnoseSemanticError("Invalid spline location", location)
				return nil
			}
		} else {
			scope.DiagnoseSemanticError("Missing spline location", point)
			fmt.Println(locationCtx)
			return nil
		}
	}

	if value := point.Number(1); value != nil && point.Number(2) != nil { // we have 3 numbers
		if val, err := strconv.ParseFloat(value.GetText(), 64); err == nil {
			out.Value = val
		} else {
			scope.DiagnoseSemanticError("Invalid spline value", point)
			return nil
		}
	} else if splineValue := point.DensityFn_Spline(); splineValue != nil {
		out.SplineValue = traversal.ConstructRegistry.Construct(splineValue, currentNamespace, scope)
	} else if ref := point.ResourceReference(); ref != nil {
		out.SplineValue = traversal.ConstructRegistry.Construct(ref, currentNamespace, scope)
	}

	derivativeCtx := point.GetChild(2)
	if derivative, ok := derivativeCtx.(*grammar.NumberContext); ok {
		val, err := strconv.ParseFloat(derivative.GetText(), 64)
		if err != nil {
			scope.DiagnoseSemanticError("Invalid spline derivative", point)
			return nil
		}
		out.Derivative = val
	} else {
		scope.DiagnoseSemanticError("Invalid spline derivative", point)
		return nil
	}

	return out
}

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_SplineFnContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			densityFn := ctx.(*grammar.DensityFn_SplineFnContext)
			splineDef := densityFn.DensityFn_Spline()
			return traversal.ConstructRegistry.Construct(splineDef, currentNamespace, scope)
		},
	)

	traversal.ConstructRegistry.Register(
		reflect.TypeFor[*grammar.DensityFn_SplineContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			splineDef := ctx.(*grammar.DensityFn_SplineContext)
			if constant := splineDef.Number(); constant != nil {
				var spline float64
				builder_chain.Builder_GetFloat(splineDef, func(f float64) { spline = f }, scope, "Spline")
				return &SplineConst{
					Spline: spline,
				}
			}

			out := Spline{
				Points: make([]SplinePoint, 0),
			}
			inputCtx := splineDef.DensityFn()
			if inputCtx == nil {
				scope.DiagnoseSemanticError("Missing input to spline", ctx)
			} else {
				input := traversal.ConstructRegistry.Construct(inputCtx, currentNamespace, scope)
				out.Coordinate = input
			}

			for _, p := range splineDef.AllDensityFn_SplinePoint() {
				point := buildSplinePoint(p.(*grammar.DensityFn_SplinePointContext), currentNamespace, scope)
				if point != nil {
					out.Points = append(out.Points, *point)
				}
			}

			return &SplineDensityFn{
				Spline: out,
			}

		})
}

type SplineDensityFn struct {
	Spline Spline
}

type Spline struct {
	Coordinate traversal.Construct `json:"coordinate"`
	Points     []SplinePoint       `json:"points"`
}
type SplinePoint struct {
	Location    float64
	Value       float64
	SplineValue traversal.Construct
	Derivative  float64
}

func (sp SplinePoint) MarshalJSON() ([]byte, error) {
	if sp.SplineValue != nil {
		return json.MarshalIndent(struct {
			Location   float64             `json:"location"`
			Value      traversal.Construct `json:"value"`
			Derivative float64             `json:"derivative"`
		}{
			Location:   sp.Location,
			Value:      sp.SplineValue,
			Derivative: sp.Derivative,
		}, "", "  ")
	} else {
		return json.MarshalIndent(struct {
			Location   float64 `json:"location"`
			Value      float64 `json:"value"`
			Derivative float64 `json:"derivative"`
		}{
			Location:   sp.Location,
			Value:      sp.Value,
			Derivative: sp.Derivative,
		}, "", "  ")
	}
}

func (c SplineDensityFn) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type   string `json:"type"`
		Spline Spline `json:"spline"`
	}{
		Type:   "minecraft:spline",
		Spline: c.Spline,
	}, "", "  ")
}

func (c SplineDensityFn) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, c)
}

// Spline Constant Form

type SplineConst struct {
	Spline float64
}

func (s SplineConst) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type   string  `json:"type"`
		Spline float64 `json:"spline"`
	}{
		Type:   "minecraft:spline",
		Spline: s.Spline,
	}, "", "  ")
}

func (s SplineConst) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportDensityFunction(symbol, rootDir, s)
}
