package sema

import (
	"strconv"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Shorthand the prelude cannot write yet, keyed by the type the position
// expects, and the per-helper data the schema lacks. The forms of
// expression (`1..3`, `!a`, `a + b`, bare ids in loot pools) are in the
// prelude instead: see forms.go and prelude/forms.mms. What stays here
// needs more than a conversion can say: If(cond) { rule } on surface
// rules pushes && and || into nested rules, `block[state=value]` is
// checked against the block's states, colors are parsed from hex text,
// a trigger call names its trigger, and a Book's name is a stack. Each
// rule rewrites the expression into what it stands for and checks that,
// so the schema still validates every field.

const (
	tDensity      = "::java::data::worldgen::density_function::DensityFunction"
	tDensityRef   = "::java::data::worldgen::density_function::DensityFunctionRef"
	tMaterialRule = "::java::data::worldgen::material_rule::MaterialRule"
	tMaterialCond = "::java::data::worldgen::material_condition::MaterialCondition"
	// The Ref types are a declared id or an inline value, as the
	// condition rule's if_true and then_run take them.
	tMaterialRuleRef = "::java::data::worldgen::material_rule::MaterialRuleRef"
	tMaterialCondRef = "::java::data::worldgen::material_condition::MaterialConditionRef"
	tAnchor          = "::java::data::worldgen::VerticalAnchor"
	tFullBlock       = "::java::util::block_state::FullBlockState"

	// tLootConditionRef is the loot condition type loot tables use.
	tLootConditionRef = "::java::data::loot::LootCondition"
	tPredicateRef     = "::java::data::predicate::PredicateRef"
	tPredicate        = "::java::data::predicate::Predicate"
	tLootEntry        = "::java::data::loot::LootPoolEntry"
	tIntProvider      = "::java::data::number_provider::IntNumberProviderRef"
	tFloatProvider    = "::java::data::number_provider::FloatNumberProviderRef"
	tIntRange         = "::java::data::loot::IntRange"
	tFloatRange       = "::java::data::loot::FloatRange"
	tMinMax           = "::java::data::util::MinMaxBounds"
	tItemStack        = "::java::world::item::ItemStackTemplate"
	tCriterion        = "::java::data::advancement::AdvancementCriterion"
)

// overlay is per-helper data the schema lacks: builder methods that are
// not field names, defaults for required fields, and argument order.
type overlay struct {
	methods    map[string]methodAlias
	defaults   map[string]jsonv.Value
	positional []string
}

type methodAlias struct {
	field string
	val   jsonv.Value // set for flag methods such as .Add()
}

var (
	zero = jsonv.Number("0")
	no   = jsonv.Bool(false)
)

// callAliases rename helpers per dispatcher, mostly names the old language
// used.
var callAliases = map[string]map[string]string{
	"minecraft:material_condition": {
		"AboveWater":   "water",
		"AboveSurface": "above_preliminary_surface",
	},
	"minecraft:recipe_serializer": {
		"Shaped":    "crafting_shaped",
		"Shapeless": "crafting_shapeless",
	},
}

// structAliases name non-dispatched structs by a shorter helper name.
var structAliases = map[string]string{
	"Pool": "LootPool",
}

// sugarCallables are helper names that sugar rules accept at a type, for
// completion.
var sugarCallables = map[string][]string{
	tMaterialRule: {"If"},
}

var overlays = map[string]*overlay{
	"LootPool": {positional: []string{"rolls", "entries"}},
	"minecraft:recipe_serializer/crafting_shaped": {
		positional: []string{"pattern", "key", "result"},
	},
	"minecraft:recipe_serializer/crafting_shapeless": {
		positional: []string{"ingredients", "result"},
	},
	"minecraft:material_condition/water": {
		methods: map[string]methodAlias{
			"Mul": {field: "surface_depth_multiplier"},
			"Add": {field: "add_stone_depth", val: jsonv.Bool(true)},
		},
		defaults: map[string]jsonv.Value{"offset": zero, "surface_depth_multiplier": zero, "add_stone_depth": no},
	},
	"minecraft:material_condition/y_above": {
		methods: map[string]methodAlias{
			"Mul": {field: "surface_depth_multiplier"},
			"Add": {field: "add_stone_depth", val: jsonv.Bool(true)},
		},
		defaults: map[string]jsonv.Value{"surface_depth_multiplier": zero, "add_stone_depth": no},
	},
	"minecraft:material_condition/stone_depth": {
		methods: map[string]methodAlias{
			"Add": {field: "add_surface_depth", val: jsonv.Bool(true)},
		},
		defaults:   map[string]jsonv.Value{"offset": zero, "add_surface_depth": no, "secondary_depth_range": zero},
		positional: []string{"surface_type", "offset", "add_surface_depth", "secondary_depth_range"},
	},
	"minecraft:material_condition/noise_threshold": {
		methods: map[string]methodAlias{
			"Min": {field: "min_threshold"},
			"Max": {field: "max_threshold"},
		},
	},
}

type sugarRule struct {
	apply func(c *Checker, e syntax.Expr, v tview) jsonv.Value
}

func (c *Checker) sugar(e syntax.Expr, v tview) *sugarRule {
	if v.T == bookTextType {
		return &sugarRule{bookTextRule}
	}
	if v.T != nil && v.T.Kind == "list" && callName(e) == "book_pages" {
		return &sugarRule{bookPagesRule}
	}
	switch v.Name {
	case tMaterialRule:
		switch e.(type) {
		case *syntax.Call, *syntax.Trailing:
			if callName(e) == "If" {
				return &sugarRule{ruleIf}
			}
		}
	case tMaterialCond:
		if x, ok := e.(*syntax.Binary); ok && (x.Op == "&&" || x.Op == "||") {
			return &sugarRule{condLogic}
		}
	case tFullBlock:
		if _, ok := e.(*syntax.Index); ok {
			return &sugarRule{blockWithState}
		}
	case tRGB, tRGBA, tStringRGB, tStringARGB:
		if isColorHexString(e) {
			return &sugarRule{colorHex}
		}
	case tText:
		switch x := e.(type) {
		case *syntax.String:
			return &sugarRule{textString}
		case *syntax.Call:
			if callName(x) == "TextFile" {
				return &sugarRule{textFile}
			}
		}
	case tCriterion:
		if x, ok := e.(*syntax.Call); ok {
			if id, ok := x.Fn.(*syntax.Ident); ok && id.Name != "AdvancementCriterion" {
				return &sugarRule{criterionCall}
			}
		}
	}
	return nil
}

func isID(e syntax.Expr) bool {
	switch e.(type) {
	case *syntax.Ident, *syntax.ResLoc:
		return true
	}
	return false
}

// mkObject builds `{ k1 = v1, k2 = v2, ... }`; nil values are skipped.
func mkObject(sp syntax.Span, kv ...any) *syntax.Object {
	b := &syntax.Block{Closed: true, Sp: sp}
	for i := 0; i+1 < len(kv); i += 2 {
		v, _ := kv[i+1].(syntax.Expr)
		if v == nil || isNilExpr(v) {
			continue
		}
		b.Items = append(b.Items, &syntax.FieldStmt{Key: &syntax.Name{Text: kv[i].(string), Sp: sp}, Value: v, Sp: v.Span()})
	}
	return &syntax.Object{Body: b, Sp: sp}
}

func isNilExpr(e syntax.Expr) bool {
	switch x := e.(type) {
	case *syntax.Ident:
		return x == nil
	case *syntax.Number:
		return x == nil
	}
	return false
}

// criterionCall lowers `InventoryChanged(items = [...])` to a criterion
// whose trigger is inventory_changed and whose conditions are the named
// arguments.
func criterionCall(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	call := e.(*syntax.Call)
	name := call.Fn.(*syntax.Ident)
	cond := &syntax.Block{Closed: true, Sp: call.ArgsSp}
	for _, a := range call.Args {
		if a.Name == nil {
			c.errorf("mms/missing-field", a.Sp, "criterion conditions are named: %s(player = ...)", name.Name)
			continue
		}
		cond.Items = append(cond.Items, &syntax.FieldStmt{Key: a.Name, Value: a.Value, Sp: a.Sp})
	}
	// The conditions are checked even when there are none, so that the
	// call records their fields for completion; an empty object is then
	// left out.
	conditions := &syntax.Object{Body: cond, Sp: call.ArgsSp}
	in := c.info(name)
	in.Type = v
	out := c.check(mkObject(call.Sp, "trigger", &syntax.Ident{Name: Snake(name.Name), Sp: name.Sp}, "conditions", conditions), v)
	if ci := c.Info[cond]; ci != nil {
		callIn := c.info(call)
		callIn.Fields, callIn.Disp = ci.Fields, ci.Disp
	}
	if o, ok := out.(*jsonv.Object); ok && len(cond.Items) == 0 {
		if co, ok := o.Get("conditions").(*jsonv.Object); ok && len(co.Members) == 0 {
			kept := o.Members[:0]
			for _, m := range o.Members {
				if m.Key != "conditions" {
					kept = append(kept, m)
				}
			}
			o.Members = kept
		}
	}
	return out
}

func mkCall(name string, sp syntax.Span, args ...syntax.Expr) *syntax.Call {
	call := &syntax.Call{Fn: &syntax.Ident{Name: name, Sp: sp}, Sp: sp, ArgsSp: sp}
	for _, a := range args {
		call.Args = append(call.Args, &syntax.Arg{Value: a, Sp: a.Span()})
	}
	return call
}

// ruleIf lowers `If(cond) { rule }` (or `If(cond, rule)`) to nested
// condition rules. Conditions have no `and` or `or` type, so && nests
// conditions, || becomes a sequence that tries each branch, and ! of a
// compound condition is pushed inward with De Morgan's laws.
func ruleIf(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	base, name, trailing, steps := unchain(e)
	if len(steps) > 0 {
		c.errorf("mms/unknown-method", steps[0].sp, "If(...) has no methods")
	}
	var args []syntax.Expr
	for _, a := range base.Args {
		if a.Name != nil {
			c.errorf("mms/unknown-field", a.Name.Sp, "If takes a condition and a rule, not named arguments")
			continue
		}
		args = append(args, a.Value)
	}
	if trailing != nil {
		args = append(args, blockExpr(trailing))
	}
	in := c.info(name)
	in.Type = v
	if trailing != nil && base != nil {
		// Completion inside `If(|) { ... }` reads the call's type.
		bi := c.info(base)
		if !bi.typed {
			bi.Type, bi.typed = v, true
		}
	}
	if len(args) != 2 {
		// Still check what is there, so hover, completion and signature
		// help work while the If is half written: the condition, or the
		// trailing rule of an If whose condition is empty.
		switch {
		case len(args) == 1 && trailing != nil && len(base.Args) == 0:
			c.check(args[0], c.norm(tview{T: &spec.Type{Kind: "reference", Path: tMaterialRuleRef}}))
		case len(args) == 1:
			c.checkIfCond(args[0])
		}
		c.errorf("mms/missing-value", name.Sp, "If needs a condition and a rule: If(cond) { rule }")
		return nil
	}
	return c.check(lowerIf(args[0], args[1], e.Span()), v)
}

// checkIfCond checks the condition of an If whose rule is missing,
// operand by operand for `&&` and `||`, which only If lowers.
func (c *Checker) checkIfCond(e syntax.Expr) {
	switch x := e.(type) {
	case *syntax.Paren:
		c.checkIfCond(x.X)
		return
	case *syntax.Binary:
		if x.Op == "&&" || x.Op == "||" {
			c.checkIfCond(x.L)
			c.checkIfCond(x.R)
			return
		}
	}
	c.check(e, c.norm(tview{T: &spec.Type{Kind: "reference", Path: tMaterialCondRef}}))
}

func lowerIf(cond, rule syntax.Expr, sp syntax.Span) syntax.Expr {
	switch x := cond.(type) {
	case *syntax.Paren:
		return lowerIf(x.X, rule, sp)
	case *syntax.Binary:
		switch x.Op {
		case "&&":
			return lowerIf(x.L, lowerIf(x.R, rule, sp), sp)
		case "||":
			return &syntax.List{Elems: []syntax.Expr{lowerIf(x.L, rule, sp), lowerIf(x.R, rule, sp)}, Sp: sp}
		}
	case *syntax.Unary:
		if x.Op == "!" {
			inner := x.X
			for {
				p, ok := inner.(*syntax.Paren)
				if !ok {
					break
				}
				inner = p.X
			}
			switch y := inner.(type) {
			case *syntax.Unary:
				if y.Op == "!" {
					return lowerIf(y.X, rule, sp)
				}
			case *syntax.Binary:
				not := func(e syntax.Expr) syntax.Expr { return &syntax.Unary{Op: "!", X: e, Sp: e.Span()} }
				switch y.Op {
				case "&&":
					return lowerIf(&syntax.Binary{Op: "||", L: not(y.L), R: not(y.R), Sp: y.Sp}, rule, sp)
				case "||":
					return lowerIf(&syntax.Binary{Op: "&&", L: not(y.L), R: not(y.R), Sp: y.Sp}, rule, sp)
				}
			}
		}
	}
	return mkCall("Condition", sp, cond, rule)
}

func condLogic(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	b := e.(*syntax.Binary)
	c.errorf("mms/unsupported-operator", b.Sp, "surface conditions cannot be combined with '%s' on their own; use it inside If(...), which nests or sequences the rule", b.Op)
	return nil
}

// blockWithState lowers `oak_log[axis=y]` to { Name, Properties } and
// checks the properties against the block's states.
func blockWithState(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	ix := e.(*syntax.Index)
	o := &jsonv.Object{}
	var name string
	switch x := ix.X.(type) {
	case *syntax.Ident:
		name = x.Name
	case *syntax.ResLoc:
		name = x.Text
	default:
		c.errorf("mms/type-mismatch", ix.X.Span(), "expected a block id before '['")
		return nil
	}
	id := c.resolveID(name, idInfo{Registry: "block", Present: true}, ix.X)
	o.Set("Name", jsonv.String(id))
	_, path := splitID(id)
	states := c.Pack.Blocks[path]
	props := &jsonv.Object{}
	for _, a := range ix.Args {
		if a.Name == nil {
			c.errorf("mms/expected-field", a.Sp, "expected `property=value`")
			continue
		}
		var val string
		switch x := a.Value.(type) {
		case *syntax.Ident:
			val = x.Name
		case *syntax.Number:
			val = x.Text
		case *syntax.String:
			val = x.Value
		case *syntax.Bool:
			val = strconv.FormatBool(x.Value)
		default:
			c.errorf("mms/type-mismatch", a.Value.Span(), "expected a property value")
			continue
		}
		if states != nil {
			allowed, ok := states.Properties[a.Name.Text]
			if !ok {
				c.errorf("mms/unknown-field", a.Name.Sp, "%s has no property '%s'", id, a.Name.Text)
				continue
			}
			if !contains(allowed, val) {
				c.errorf("mms/type-mismatch", a.Value.Span(), "%s.%s must be one of %s", path, a.Name.Text, quoteList(allowed))
				continue
			}
		}
		props.Set(a.Name.Text, jsonv.String(val))
	}
	if len(props.Members) > 0 {
		o.Set("Properties", props)
	}
	return o
}
