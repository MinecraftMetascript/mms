package eval

import (
	"math"

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
