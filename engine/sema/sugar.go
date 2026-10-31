package sema

import (
	"strconv"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Sugar: hand-written rules for MMS syntax with no direct schema shape,
// keyed by the type the position expects. Most rules rewrite the
// expression into the call it stands for and check that, so the schema
// still validates every field.

const (
	tDensity      = "::java::data::worldgen::density_function::DensityFunction"
	tDensityRef   = "::java::data::worldgen::density_function::DensityFunctionRef"
	tMaterialRule = "::java::data::worldgen::material_rule::MaterialRule"
	tMaterialCond = "::java::data::worldgen::material_condition::MaterialCondition"
	tAnchor       = "::java::data::worldgen::VerticalAnchor"
	tFullBlock    = "::java::util::block_state::FullBlockState"

	tLootCondition = "::java::data::loot::condition::LootCondition"
	// tLootConditionRef is an alias of tLootCondition used by loot tables.
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
	switch v.Name {
	case tDensity:
		switch x := e.(type) {
		case *syntax.Binary:
			if _, ok := densityOps[x.Op]; ok {
				return &sugarRule{densityBinary}
			}
		case *syntax.Unary:
			if x.Op == "-" && !isNumber(x) {
				return &sugarRule{densityNeg}
			}
		}
	case tMaterialRule:
		switch e.(type) {
		case *syntax.List:
			return &sugarRule{ruleSequence}
		case *syntax.Call, *syntax.Trailing:
			if callName(e) == "If" {
				return &sugarRule{ruleIf}
			}
		}
	case tMaterialCond:
		switch x := e.(type) {
		case *syntax.Unary:
			if x.Op == "!" {
				return &sugarRule{condNot}
			}
		case *syntax.Binary:
			if x.Op == "&&" || x.Op == "||" {
				return &sugarRule{condLogic}
			}
		}
	case tAnchor:
		if isNumber(e) {
			return &sugarRule{anchorAbsolute}
		}
	case tFullBlock:
		if _, ok := e.(*syntax.Index); ok {
			return &sugarRule{blockWithState}
		}
	case tLootCondition, tLootConditionRef:
		switch x := e.(type) {
		case *syntax.Unary:
			if x.Op == "!" {
				return &sugarRule{lootNot}
			}
		case *syntax.Binary:
			if x.Op == "&&" || x.Op == "||" {
				return &sugarRule{lootLogic}
			}
		}
	case tLootEntry:
		switch e.(type) {
		case *syntax.Ident, *syntax.ResLoc:
			return &sugarRule{entryItem}
		case *syntax.Tag:
			return &sugarRule{entryTag}
		}
	case tIntProvider, tFloatProvider:
		if r, ok := e.(*syntax.Range); ok && r.Lo != nil && r.Hi != nil {
			return &sugarRule{providerRange}
		}
	case tIntRange, tFloatRange, tMinMax:
		if _, ok := e.(*syntax.Range); ok {
			return &sugarRule{boundsRange}
		}
	case tItemStack:
		if b, ok := e.(*syntax.Binary); ok && b.Op == "*" && isID(b.L) && isNumber(b.R) {
			return &sugarRule{stackCount}
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

func word(text string, sp syntax.Span) *syntax.Ident { return &syntax.Ident{Name: text, Sp: sp} }

// lootLogic lowers `a && b && c` to all_of and `a || b` to any_of, with
// chains of the same operator flattened into one list of terms.
// Parenthesised groups stay nested.
func lootLogic(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	b := e.(*syntax.Binary)
	var terms []syntax.Expr
	var add func(x syntax.Expr)
	add = func(x syntax.Expr) {
		// Parentheses keep a nested group: (a && b) && c is all_of(all_of(a, b), c).
		if inner, ok := x.(*syntax.Binary); ok && inner.Op == b.Op {
			add(inner.L)
			add(inner.R)
			return
		}
		terms = append(terms, x)
	}
	add(b)
	kind := "all_of"
	if b.Op == "||" {
		kind = "any_of"
	}
	return c.check(mkObject(b.Sp, "type", word(kind, b.Sp), "terms", &syntax.List{Elems: terms, Sp: b.Sp}), v)
}

func lootNot(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	u := e.(*syntax.Unary)
	return c.check(mkObject(u.Sp, "type", word("inverted", u.Sp), "term", u.X), v)
}

// entryItem lowers a bare id in a pool's entries to an item entry.
func entryItem(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	return c.check(mkObject(e.Span(), "type", word("item", e.Span()), "name", e), v)
}

// entryTag lowers `#tag` in a pool's entries to a tag entry that expands
// to one entry per item, so each item is equally likely.
func entryTag(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	t := e.(*syntax.Tag)
	return c.check(mkObject(t.Sp, "type", word("tag", t.Sp), "items", t, "expand", &syntax.Bool{Value: true, Sp: t.Sp}), v)
}

// providerRange lowers `a..b` to a uniform number provider.
func providerRange(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	r := e.(*syntax.Range)
	return c.check(mkObject(r.Sp, "type", word("uniform", r.Sp), "min", r.Lo, "max", r.Hi), v)
}

// boundsRange lowers `a..b`, `a..` or `..b` to `{min, max}` bounds.
func boundsRange(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	r := e.(*syntax.Range)
	var lo, hi syntax.Expr
	if r.Lo != nil {
		lo = r.Lo
	}
	if r.Hi != nil {
		hi = r.Hi
	}
	return c.check(mkObject(r.Sp, "min", lo, "max", hi), v)
}

// stackCount lowers `diamond * 4` to an item stack of four.
func stackCount(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	b := e.(*syntax.Binary)
	return c.check(mkObject(b.Sp, "id", b.L, "count", b.R), v)
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
	var conditions syntax.Expr
	if len(cond.Items) > 0 {
		conditions = &syntax.Object{Body: cond, Sp: call.ArgsSp}
	}
	in := c.info(name)
	in.Type = v
	return c.check(mkObject(call.Sp, "trigger", &syntax.Ident{Name: Snake(name.Name), Sp: name.Sp}, "conditions", conditions), v)
}

func mkCall(name string, sp syntax.Span, args ...syntax.Expr) *syntax.Call {
	call := &syntax.Call{Fn: &syntax.Ident{Name: name, Sp: sp}, Sp: sp, ArgsSp: sp}
	for _, a := range args {
		call.Args = append(call.Args, &syntax.Arg{Value: a, Sp: a.Span()})
	}
	return call
}

var densityOps = map[string]string{"+": "Add", "-": "Sub", "*": "Mul", "/": "Div"}

// densityBinary lowers `a + b` to add, `a - b` to sub and so on, folding
// constant operands.
func densityBinary(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	b := e.(*syntax.Binary)
	if f, ok := fold(b); ok {
		return c.check(&syntax.Number{Text: strconv.FormatFloat(f, 'g', -1, 64), Sp: b.Sp}, v)
	}
	return c.check(mkCall(densityOps[b.Op], b.Sp, b.L, b.R), v)
}

func densityNeg(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	u := e.(*syntax.Unary)
	return c.check(mkCall("Mul", u.Sp, &syntax.Number{Text: "-1", Sp: u.Sp}, u.X), v)
}

// fold evaluates an arithmetic expression made only of numbers.
func fold(e syntax.Expr) (float64, bool) {
	switch x := e.(type) {
	case *syntax.Paren:
		return fold(x.X)
	case *syntax.Number, *syntax.Unary:
		if t, ok := numberText(e); ok {
			f, err := strconv.ParseFloat(t, 64)
			return f, err == nil
		}
	case *syntax.Binary:
		l, ok1 := fold(x.L)
		r, ok2 := fold(x.R)
		if !ok1 || !ok2 {
			return 0, false
		}
		switch x.Op {
		case "+":
			return l + r, true
		case "-":
			return l - r, true
		case "*":
			return l * r, true
		case "/":
			if r != 0 {
				return l / r, true
			}
		}
	}
	return 0, false
}

// ruleSequence lowers a list of rules to a sequence. Nested lists (written,
// or produced by || in If) are flattened: a sequence inside a sequence
// tries the same rules in the same order.
func ruleSequence(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	l := e.(*syntax.List)
	flat := &syntax.List{Sp: l.Sp}
	var add func(l *syntax.List)
	add = func(l *syntax.List) {
		for _, x := range l.Elems {
			if inner, ok := x.(*syntax.List); ok {
				add(inner)
				continue
			}
			flat.Elems = append(flat.Elems, x)
		}
	}
	add(l)
	return c.check(mkCall("Sequence", e.Span(), flat), v)
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
	if len(args) != 2 {
		c.errorf("mms/missing-value", name.Sp, "If needs a condition and a rule: If(cond) { rule }")
		return nil
	}
	return c.check(lowerIf(args[0], args[1], e.Span()), v)
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

func condNot(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	u := e.(*syntax.Unary)
	return c.check(mkCall("Not", u.Sp, u.X), v)
}

func condLogic(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	b := e.(*syntax.Binary)
	c.errorf("mms/unsupported-operator", b.Sp, "surface conditions cannot be combined with '%s' on their own; use it inside If(...), which nests or sequences the rule", b.Op)
	return nil
}

func anchorAbsolute(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	return c.check(mkCall("Absolute", e.Span(), e), v)
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
