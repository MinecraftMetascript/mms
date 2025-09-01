package surface_rules

import (
	"encoding/json"
	"errors"

	"mms2/lang/grammar"
	"mms2/lang/traversal"
	"mms2/lib"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

type SurfaceRuleKind string

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRuleDefinitionContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			def := ctx.(*grammar.SurfaceRuleDefinitionContext)

			if v, ok := def.SurfaceRule().GetChild(0).(antlr.ParserRuleContext); ok {
				return traversal.ConstructRegistry.Construct(
					v,
					currentNamespace,
					scope,
				)
			} else {
				return nil
			}
		},
	)
}

const (
	SurfaceRule_Block     SurfaceRuleKind = "minecraft:block"
	SurfaceRule_Bandlands SurfaceRuleKind = "minecraft:bandlands"
	SurfaceRule_Condition SurfaceRuleKind = "minecraft:condition"
	SurfaceRule_Sequence  SurfaceRuleKind = "minecraft:sequence"
	SurfaceRule_Reference SurfaceRuleKind = "__mms:reference"
)

func exportSurfaceRule(symbol traversal.Symbol, rootDir *lib.FileTreeLike, rule json.Marshaler) error {
	data, err := json.Marshal(rule)
	if err != nil {
		return err
	}
	ref := symbol.GetReference()
	if ref.String() == ":" {
		return errors.New("symbol has no reference")
	}
	rootDir.
		MkDir(symbol.GetReference().GetNamespace(), nil).
		MkDir("_debug", nil).
		MkDir("surface_rule", nil).
		MkFile(symbol.GetReference().GetName()+".json", string(data), nil)

	return nil
}
