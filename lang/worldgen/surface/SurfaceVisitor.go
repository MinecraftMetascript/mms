package surface

import (
	"encoding/json"
	"fmt"

	"github.com/minecraftmetascript/mms/lang/grammars"
	"github.com/minecraftmetascript/mms/lang/mms_errors"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_conditions"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface/surface_rules"
	"github.com/minecraftmetascript/mms/lib"
)

func NewSurfaceVisitor(
	reportError func(mmsError mms_errors.MMSError),
	filename string,
) *Visitor {
	return &Visitor{
		AddError:              reportError,
		RuleDeclarations:      make(map[string]lib.Symbol[surface_rules.SurfaceRule]),
		ConditionDeclarations: make(map[string]lib.Symbol[surface_conditions.SurfaceCondition]),
		filename:              filename,
	}
}

func (v *Visitor) Namespace() string {
	return v.namespace
}
func (v *Visitor) SetNamespace(namespace string) {
	v.namespace = namespace
}

type Visitor struct {
	grammars.BaseMMSParserListener
	AddError              func(mmsError mms_errors.MMSError)
	RuleDeclarations      map[string]lib.Symbol[surface_rules.SurfaceRule]
	ConditionDeclarations map[string]lib.Symbol[surface_conditions.SurfaceCondition]
	filename              string
	namespace             string
}

func (v *Visitor) DumpDeclarations(ns *lib.Namespace) {
	for name, rule := range v.RuleDeclarations {
		ns.Set(name, lib.Symbol[json.Marshaler]{
			File:     rule.File,
			Location: rule.Location,
			Ref:      rule.Ref,
			Value:    rule.Value,
			Kind:     rule.Kind,
		})
	}
	for name, condition := range v.ConditionDeclarations {
		ns.Set(name, lib.Symbol[json.Marshaler]{
			File:     condition.File,
			Location: condition.Location,
			Ref:      condition.Ref,
			Value:    condition.Value,
			Kind:     condition.Kind,
		})
	}
}

func (v *Visitor) ExitSurfaceConditionDeclaration(ctx *grammars.SurfaceConditionDeclarationContext) {
	condition, err := v.ConstructSurfaceCondition(ctx.SurfaceCondition().(*grammars.SurfaceConditionContext))
	if err != nil {
		v.AddError(
			mms_errors.SyntaxError(
				v.filename, lib.GetRuleLocation(ctx), err.Error(),
			),
		)
		return
	}

	id := ctx.Identifier().GetText()

	if condition, ok := v.ConditionDeclarations[id]; ok {
		v.AddError(
			mms_errors.DuplicateSymbolError(
				v.filename,
				condition.Value.GetLocation(),
				fmt.Sprintf("Duplicate rule declaration: %v", id),
			),
		)
	}
	v.ConditionDeclarations[id] = mkConditionSymbol(ctx, condition, lib.Reference{Name: id, Namespace: v.namespace}, v.filename)
}

func (v *Visitor) ExitSurfaceRuleDeclaration(ctx *grammars.SurfaceRuleDeclarationContext) {
	surfaceRuleCtx := ctx.SurfaceRule().(*grammars.SurfaceRuleContext)
	rule, errs := v.ConstructSurfaceRule(surfaceRuleCtx)
	if errs != nil {
		for _, err := range errs {
			v.AddError(
				mms_errors.SyntaxError(
					v.filename, lib.GetRuleLocation(ctx), err.Error(),
				),
			)
		}
		return
	}

	id := ctx.Identifier().GetText()
	if rule, ok := v.RuleDeclarations[id]; ok {
		v.AddError(
			mms_errors.DuplicateSymbolError(
				v.filename,
				rule.Value.GetLocation(),
				fmt.Sprintf("Duplicate rule declaration: %v", id),
			),
		)
	}
	v.RuleDeclarations[id] = mkRuleSymbol(ctx, rule, lib.Reference{Name: id, Namespace: v.Namespace()}, v.filename)
	fmt.Println("Found Rule -- ", id)
}
