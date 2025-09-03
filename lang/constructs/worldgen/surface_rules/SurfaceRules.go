package surface_rules

import (
	"encoding/json"
	"errors"

	"reflect"

	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"

	"github.com/antlr4-go/antlr/v4"
)

type SurfaceRuleKind string

func init() {
	traversal.ConstructRegistry.Register(
		reflect.TypeFor[grammar.SurfaceRuleDefinitionContext](),
		func(ctx antlr.ParserRuleContext, currentNamespace string, scope *traversal.Scope) traversal.Construct {
			def := ctx.(*grammar.SurfaceRuleDefinitionContext)
			rule := def.SurfaceRule()
			if rule == nil {
				return nil
			}

			if rule.GetChildCount() > 0 {
				if v, ok := rule.GetChild(0).(antlr.ParserRuleContext); ok {
					return traversal.ConstructRegistry.Construct(
						v,
						currentNamespace,
						scope,
					)
				}
			}
			return nil
		},
	)
}

const (
	SurfaceRule_Block     SurfaceRuleKind = "minecraft:block"
	SurfaceRule_Bandlands SurfaceRuleKind = "minecraft:bandlands"
	SurfaceRule_Condition SurfaceRuleKind = "minecraft:condition"
	SurfaceRule_Sequence  SurfaceRuleKind = "minecraft:sequence"
)

func exportSurfaceRule(symbol traversal.Symbol, rootDir *lib.FileTreeLike, rule json.Marshaler) error {
	data, err := json.MarshalIndent(rule, "", "  ")
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
