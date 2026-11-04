package sema

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Colors: the mcdoc color types are unions of an int, a list of 0 to 1
// floats and, for some, a "#hex" string. The schema names them (`RGBA`)
// without saying how to write one, so hovers, errors and completion spell
// the forms out, and a "#hex" string is accepted for the int-only types
// too and written as the int.

const (
	tRGB        = "::java::util::color::RGB"
	tRGBA       = "::java::util::color::RGBA"
	tStringRGB  = "::java::util::color::StringRGB"
	tStringARGB = "::java::util::color::StringARGB"
)

// colorType is what a color type accepts.
type colorType struct {
	alpha bool // ARGB int, [r, g, b, a] and "#AARRGGBB"
	hex   bool // the game itself reads "#hex" strings
}

var colorTypes = map[string]colorType{
	tRGB:        {},
	tRGBA:       {alpha: true},
	tStringRGB:  {hex: true},
	tStringARGB: {alpha: true, hex: true},
}

// colorOf reports whether v is one of the color types, by name or by an
// unresolved reference to one.
func colorOf(v tview) (colorType, bool) {
	if ct, ok := colorTypes[v.Name]; ok {
		return ct, true
	}
	t := v.T
	for t != nil && t.Kind == "union" && len(t.Members) == 1 {
		t = t.Members[0] // `(StringRGB)`, as in a biome's water_color
	}
	if t != nil && t.Kind == "reference" {
		ct, ok := colorTypes[t.Path]
		return ct, ok
	}
	return colorType{}, false
}

func (ct colorType) name() string {
	if ct.alpha {
		return "an RGBA color"
	}
	return "an RGB color"
}

// help says how to write the color, as one sentence of Markdown.
func (ct colorType) help() string {
	if ct.alpha {
		return "Write a color as `\"#RRGGBB\"` or `\"#AARRGGBB\"` (alpha first: `\"#80FF0000\"` is half-transparent red), " +
			"an ARGB integer such as `-65536` (`0xFFFF0000`), or `[r, g, b, a]` with each part from 0 to 1, such as `[1, 0, 0, 0.5]`."
	}
	return "Write a color as `\"#RRGGBB\"` (`\"#FF0000\"` is red), " +
		"an RGB integer such as `16711680` (`0xFF0000`), or `[r, g, b]` with each part from 0 to 1, such as `[1, 0, 0]`."
}

// colorPlaceholder is the value a quick fix fills in: opaque white.
const colorPlaceholder = `"#FFFFFF"`

// colorHex checks a "#hex" string where a color is expected. For the
// int-only types the result is the color's int; for types the game reads
// as hex strings it is the string itself.
func colorHex(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	ct, _ := colorOf(v)
	s := e.(*syntax.String).Value
	n, ok := parseHexColor(s, ct.alpha)
	if !ok {
		want := `"#RRGGBB"`
		if ct.alpha {
			want = `"#RRGGBB" or "#AARRGGBB"`
		}
		c.errorf("mms/type-mismatch", e.Span(), "expected %s such as %s, got string %q", ct.name(), want, s)
		return nil
	}
	if ct.hex {
		return jsonv.String(s)
	}
	return jsonv.Number(strconv.FormatInt(n, 10))
}

// parseHexColor reads "#RRGGBB", or "#AARRGGBB" when alpha is allowed, as
// the int the game stores: ARGB as a signed 32-bit int (opaque when the
// alpha is left out), or plain RGB.
func parseHexColor(s string, alpha bool) (int64, bool) {
	if !isHexColor(s, 6) && !(alpha && isHexColor(s, 8)) {
		return 0, false
	}
	u, err := strconv.ParseUint(s[1:], 16, 32)
	if err != nil {
		return 0, false
	}
	if !alpha {
		return int64(u), true
	}
	if len(s) == 7 {
		u |= 0xFF000000
	}
	return int64(int32(uint32(u))), true
}

// isColorHexString reports whether e is a string that starts like a hex
// color, which the color sugar takes.
func isColorHexString(e syntax.Expr) bool {
	s, ok := e.(*syntax.String)
	return ok && strings.HasPrefix(s.Value, "#")
}

// withColorHelp adds how to write a color to a field's docs.
func withColorHelp(desc string, v tview) string {
	ct, ok := colorOf(v)
	if !ok {
		return desc
	}
	if desc == "" {
		return ct.help()
	}
	return desc + "\n\n" + ct.help()
}
