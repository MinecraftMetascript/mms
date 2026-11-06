package eval

import (
	"math"
	"regexp"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// intrinsic is a Template written in Go. It gets the call and its
// evaluated arguments.
type intrinsic func(c *ctx, x *syntax.Call, args []argValue) Value

// intrinsics back `extern Template` declarations, which give them their
// signatures and docs. An extern with no entry here is an error.
var intrinsics = map[string]intrinsic{
	"cos":         trig(math.Cos),
	"sin":         trig(math.Sin),
	"len":         length,
	"by_language": byLanguage,
	"english":     english,
	"contains":    contains,
	"matches":     matches,
	"replace":     replace,
	"named":       named,
	"kind_of":     kindOf,
}

// Intrinsics lists the names of the intrinsics, for the test that each
// has an extern declaration in the prelude.
func Intrinsics() []string {
	out := make([]string, 0, len(intrinsics))
	for n := range intrinsics {
		out = append(out, n)
	}
	return out
}

// builtins are always in scope.
var builtins = map[string]intrinsic{
	"error": checkRule(syntax.SeverityError),
	"warn":  checkRule(syntax.SeverityWarning),
	"hint":  checkRule(syntax.SeverityHint),
}

// checkOnly builtins are only callable in a Blueprint's check block;
// elsewhere a call of that name is left alone.
var checkOnly = map[string]bool{"error": true, "warn": true, "hint": true}

func oneArg(c *ctx, x *syntax.Call, args []argValue, name string) (Value, bool) {
	if len(args) != 1 {
		c.errorf("mms/intrinsic-args", x.Sp, name+" takes one argument")
		return nil, false
	}
	return args[0].v, true
}

// trig works in degrees, as rotations in commands do.
func trig(f func(float64) float64) intrinsic {
	return func(c *ctx, x *syntax.Call, args []argValue) Value {
		v, ok := oneArg(c, x, args, "this")
		if !ok {
			return &Num{c.pos(x.Sp), 0, false}
		}
		n, ok := v.(*Num)
		if !ok {
			c.errorf("mms/intrinsic-args", x.Sp, "takes a number known at compile time, not "+kindName(v))
			return &Num{c.pos(x.Sp), 0, false}
		}
		return &Num{c.pos(x.Sp), f(n.F * math.Pi / 180), false}
	}
}

func length(c *ctx, x *syntax.Call, args []argValue) Value {
	v, ok := oneArg(c, x, args, "len")
	p := c.pos(x.Sp)
	if !ok {
		return &Num{p, 0, true}
	}
	switch v := v.(type) {
	case *List:
		return &Num{p, float64(len(v.Elems)), true}
	case *Record:
		return &Num{p, float64(len(v.Keys)), true}
	case *Str:
		return &Num{p, float64(len([]rune(v.S))), true}
	case *None:
		return &Num{p, 0, true}
	}
	c.errorf("mms/intrinsic-args", x.Sp, "len takes a list, object or string, not "+kindName(v))
	return &Num{p, 0, true}
}

// byLanguage turns a LangText (a string, or one string per language)
// into an object of language to text. A plain string is English.
func byLanguage(c *ctx, x *syntax.Call, args []argValue) Value {
	v, ok := oneArg(c, x, args, "by_language")
	p := c.pos(x.Sp)
	if !ok {
		return &Record{pos: p}
	}
	switch v := v.(type) {
	case *None:
		return &Record{pos: p}
	case *Record:
		return v
	default:
		return &Record{pos: p, Keys: []string{"en_us"}, Vals: []Value{v}}
	}
}

// english picks the English text of a LangText.
func english(c *ctx, x *syntax.Call, args []argValue) Value {
	v, ok := oneArg(c, x, args, "english")
	if !ok {
		return &None{pos: c.pos(x.Sp)}
	}
	if r, ok := v.(*Record); ok {
		if en, ok := r.get("en_us"); ok {
			return en
		}
		if len(r.Vals) > 0 {
			return r.Vals[0]
		}
		return &None{pos: c.pos(x.Sp)}
	}
	return v
}

// args checks that a call has n arguments known at compile time.
func knownArgs(c *ctx, x *syntax.Call, args []argValue, name string, n int) ([]Value, bool) {
	if len(args) != n {
		c.errorf("mms/intrinsic-args", x.Sp, name+" takes "+map[int]string{1: "one argument", 2: "two arguments", 3: "three arguments"}[n])
		return nil, false
	}
	out := make([]Value, n)
	for i, a := range args {
		if !known(a.v) && !isData(a.v) {
			c.errorf("mms/intrinsic-args", a.x.Value.Span(), name+" takes values known at compile time, not "+kindName(a.v))
			return nil, false
		}
		out[i] = a.v
	}
	return out, true
}

// contains reports whether a list holds a value.
func contains(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	vs, ok := knownArgs(c, x, args, "contains", 2)
	if !ok {
		return &Bool{p, false}
	}
	l, isList := vs[0].(*List)
	if !isList {
		if _, none := vs[0].(*None); none {
			return &Bool{p, false}
		}
		c.errorf("mms/intrinsic-args", args[0].x.Value.Span(), "contains takes a list, not "+kindName(vs[0]))
		return &Bool{p, false}
	}
	for _, e := range l.Elems {
		if known(e) && equal(e, vs[1]) {
			return &Bool{p, true}
		}
	}
	return &Bool{p, false}
}

var patterns = map[string]*regexp.Regexp{}

// matches reports whether a string matches a regular expression as a
// whole.
func matches(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	vs, ok := knownArgs(c, x, args, "matches", 2)
	if !ok {
		return &Bool{p, false}
	}
	s, ok1 := text(vs[0])
	pat, ok2 := vs[1].(*Str)
	if !ok1 || !ok2 {
		c.errorf("mms/intrinsic-args", x.Sp, "matches takes a string and a pattern string")
		return &Bool{p, false}
	}
	re := patterns[pat.S]
	if re == nil {
		var err error
		if re, err = regexp.Compile("^(?:" + pat.S + ")$"); err != nil {
			c.errorf("mms/intrinsic-args", args[1].x.Value.Span(), "not a regular expression: "+err.Error())
			return &Bool{p, false}
		}
		patterns[pat.S] = re
	}
	return &Bool{p, re.MatchString(s)}
}

// replace replaces every old in a string with new.
func replace(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	vs, ok := knownArgs(c, x, args, "replace", 3)
	if !ok {
		return &Str{p, ""}
	}
	s, ok1 := text(vs[0])
	o, ok2 := text(vs[1])
	n, ok3 := text(vs[2])
	if !ok1 || !ok2 || !ok3 {
		c.errorf("mms/intrinsic-args", x.Sp, "replace takes three strings")
		return &Str{p, ""}
	}
	return &Str{p, strings.ReplaceAll(s, o, n)}
}

// named gives an asset call that has no id name one: a CustomItem's
// Texture("art/ruby.png") becomes the texture item/ruby.
func named(c *ctx, x *syntax.Call, args []argValue) Value {
	if len(args) != 2 {
		c.errorf("mms/intrinsic-args", x.Sp, "named takes an asset and an id name")
		return &None{pos: c.pos(x.Sp)}
	}
	n, ok := text(args[1].v)
	if !ok {
		c.errorf("mms/intrinsic-args", args[1].x.Value.Span(), "an id name is a string")
		return args[0].v
	}
	code, ok := args[0].v.(*Code)
	if !ok {
		return args[0].v
	}
	call, ok := code.X.(*syntax.Call)
	if !ok || len(call.Args) != 1 {
		return code
	}
	nc := *call
	nc.Args = append(append([]*syntax.Arg(nil), call.Args...), &syntax.Arg{Value: &syntax.String{Value: n, Sp: call.Sp}, Sp: call.Sp})
	return &Code{code.pos, &nc}
}

// kindOf names the kind of a value: "none", "bool", "number", "string",
// "name", "list", "object", "range", "commands", or "runtime" for a
// value only known when the pack runs.
func kindOf(c *ctx, x *syntax.Call, args []argValue) Value {
	p := c.pos(x.Sp)
	v, ok := oneArg(c, x, args, "kind_of")
	if !ok {
		return &Str{p, "none"}
	}
	k := "runtime"
	switch v.(type) {
	case *None:
		k = "none"
	case *Bool:
		k = "bool"
	case *Num:
		k = "number"
	case *Str:
		k = "string"
	case *Word:
		k = "name"
	case *List:
		k = "list"
	case *Record:
		k = "object"
	case *RangeV:
		k = "range"
	case *Stmts:
		k = "commands"
	}
	return &Str{p, k}
}
