package surface_rules

import (
	"encoding/json"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceCondition_NotContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			not := ctx.(*grammar.SurfaceCondition_NotContext)
			if not.SurfaceCondition() == nil {
				loc := traversal.RuleLocation(ctx, scope.CurrentFile)
				*scope.Diagnostics = append(*scope.Diagnostics, traversal.Diagnostic{
					Message:  "Expected a surface condition",
					Where:    loc,
					Severity: traversal.SeverityError,
					Source:   "semantic",
					File:     scope.CurrentFile,
				})
				return nil
			}
			return InvertCondition(
				traversal.ConstructRegistry.Construct(
					not.SurfaceCondition(),
					currentNamespace,
					scope,
				),
			)
		},
	)
}

type NotCondition struct {
	traversal.Construct
	Invert traversal.Construct
}

func (c NotCondition) ExportSymbol(symbol traversal.Symbol, rootDir *lib.FileTreeLike) error {
	return exportSurfaceCondition(symbol, rootDir, c)
}

func (c NotCondition) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Type   SurfaceConditionKind `json:"type"`
		Invert traversal.Construct  `json:"invert"`
	}{
		Type:   NotConditionKind,
		Invert: c.Invert,
	}, "", "  ")
}

func InvertCondition(c traversal.Construct) NotCondition {
	return NotCondition{Invert: c}
}
