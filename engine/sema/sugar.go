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
	tMaterialRule = "::java::data::worldgen::material_rule::MaterialRule"
	tMaterialCond = "::java::data::worldgen::material_condition::MaterialCondition"
	tAnchor       = "::java::data::worldgen::VerticalAnchor"
	tFullBlock    = "::java::util::block_state::FullBlockState"
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
	}
	return nil
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
