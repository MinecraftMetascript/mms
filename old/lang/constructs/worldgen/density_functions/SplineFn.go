package density_functions

import (
	"encoding/json"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/builder_chain"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type SplineFactory struct {
}

func (s SplineFactory) Create(ctx *grammar.DensityFn_SplineFnContext, namespace string, scope *traversal.Scope) traversal.Node {
	splineDef := ctx.DensityFn_Spline()
	return buildSpline(splineDef.(*grammar.DensityFn_SplineContext), namespace, scope)
}

func (s SplineFactory) GetHelp(node *SplineDensityFn, symbol traversal.Symbol, location traversal.TextLocation) *traversal.Help {
	return &traversal.Help{
		Content:  "Defines a cubic spline",
		Position: node.GetLocation(),
	}
}
func buildSpline(ctx *grammar.DensityFn_SplineContext, currentNamespace string, scope *traversal.Scope) traversal.Node {
	if constant := ctx.Number(); constant != nil {
		var spline float64
		builder_chain.Builder_GetFloat(ctx, func(f float64) { spline = f }, scope, "Spline")
		return &SplineConst{
			Spline:   spline,
			location: traversal.RuleLocation(ctx, scope.CurrentFile),
		}
	}

	out := Spline{
		Points: make([]SplinePoint, 0),
	}
	inputCtx := ctx.DensityFn()
	if inputCtx == nil {
		scope.DiagnoseSemanticError("Missing input to spline", ctx)
	} else {
		input := traversal.ConstructRegistry.Construct(inputCtx, currentNamespace, scope)
		out.Coordinate = input
	}

	for _, p := range ctx.AllDensityFn_SplinePoint() {
		point := buildSplinePoint(p.(*grammar.DensityFn_SplinePointContext), currentNamespace, scope)
		if point != nil {
			out.Points = append(out.Points, *point)
		}
	}

	return &SplineDensityFn{
		Spline:   out,
		location: traversal.RuleLocation(ctx, scope.CurrentFile),
	}
}

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

}

type SplineDensityFn struct {
	Spline   Spline
	location traversal.TextLocation
}

func (c SplineDensityFn) GetLocation() traversal.TextLocation {
	return c.location
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

// Spline Constant Form

type SplineConst struct {
	Spline   float64
	location traversal.TextLocation
}

func (s SplineConst) GetLocation() traversal.TextLocation {
	return s.location
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
