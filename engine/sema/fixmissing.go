package sema

import (
	"math"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// fixSite is where missing entries of a struct can be written: the
// braces of an object or the parentheses of a helper call.
type fixSite struct {
	block *syntax.Block
	call  *syntax.Call
}

// missingFix builds the quick fix that adds the missing fields with
// placeholder values. Placeholders only start the value, so the fix is
// not safe and "fix all" leaves it alone.
func (c *Checker) missingFix(site fixSite, fs []field, missing []string) []syntax.Fix {
	if c.Src == nil || len(missing) == 0 {
		return nil
	}
	types := map[string]tview{}
	for _, f := range fs {
		types[f.Key] = f.Type
	}
	entries := make([]string, len(missing))
	for i, k := range missing {
		entries[i] = k + " = " + c.placeholder(types[k], 0)
	}
	var e syntax.TextEdit
	var ok bool
	switch {
	case site.block != nil:
		e, ok = c.blockInsert(site.block, entries)
	case site.call != nil:
		e, ok = c.callInsert(site.call, entries)
	}
	if !ok {
		return nil
	}
	title := "Add missing field '" + missing[0] + "'"
	if len(missing) > 1 {
		title = "Add missing fields " + quoteList(missing)
	}
	return []syntax.Fix{{Title: title, Edits: []syntax.TextEdit{e}}}
}

// blockInsert writes entries before the closing brace: one per line in a
// block that spans lines, comma-separated in a one-line block.
func (c *Checker) blockInsert(b *syntax.Block, entries []string) (syntax.TextEdit, bool) {
	src := c.Src
	open, close := b.Sp.Start, b.Sp.End-1
	if !b.Closed || open < 0 || close >= len(src) || close <= open || src[open] != '{' || src[close] != '}' {
		return syntax.TextEdit{}, false
	}
	inner := string(src[open+1 : close])
	if !strings.Contains(inner, "\n") {
		if len(b.Items) == 0 {
			return syntax.TextEdit{Sp: syntax.Span{Start: open, End: close + 1}, Text: "{ " + strings.Join(entries, ", ") + " }"}, true
		}
		end := b.Items[len(b.Items)-1].Span().End
		return syntax.TextEdit{Sp: syntax.Span{Start: end, End: end}, Text: ", " + strings.Join(entries, ", ")}, true
	}
	indent := lineIndent(src, open) + "  "
	if n := len(b.Items); n > 0 {
		indent = lineIndent(src, b.Items[n-1].Span().Start)
	}
	var sb strings.Builder
	for _, e := range entries {
		sb.WriteString(indent + e + "\n")
	}
	ls := lineStart(src, close)
	if strings.TrimSpace(string(src[ls:close])) == "" {
		return syntax.TextEdit{Sp: syntax.Span{Start: ls, End: ls}, Text: sb.String()}, true
	}
	// `x = 1 }`: break the line before the brace.
	at := close
	for at > open+1 && (src[at-1] == ' ' || src[at-1] == '\t') {
		at--
	}
	return syntax.TextEdit{Sp: syntax.Span{Start: at, End: close}, Text: "\n" + sb.String() + lineIndent(src, open)}, true
}

// callInsert adds entries as named arguments before the closing
// parenthesis.
func (c *Checker) callInsert(call *syntax.Call, entries []string) (syntax.TextEdit, bool) {
	src := c.Src
	close := call.Sp.End - 1
	if close < 0 || close >= len(src) || src[close] != ')' {
		return syntax.TextEdit{}, false
	}
	text := strings.Join(entries, ", ")
	if n := len(call.Args); n > 0 {
		end := call.Args[n-1].Span().End
		// Keep a trailing comma where it is: `f(a,)` becomes `f(a, b,)`.
		return syntax.TextEdit{Sp: syntax.Span{Start: end, End: end}, Text: ", " + text}, true
	}
	return syntax.TextEdit{Sp: syntax.Span{Start: close, End: close}, Text: text}, true
}

func lineStart(src []rune, at int) int {
	for at > 0 && src[at-1] != '\n' {
		at--
	}
	return at
}

func lineIndent(src []rune, at int) string {
	i := lineStart(src, at)
	j := i
	for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
		j++
	}
	return string(src[i:j])
}

// placeholder is source text for a value of type v: the smallest number
// in range, false, the first enum value, or an empty string, list or
// object for the user to fill in.
func (c *Checker) placeholder(v tview, depth int) string {
	v = c.norm(v)
	if v.T == nil || depth > 8 {
		return "{}"
	}
	t := v.T
	switch t.Kind {
	case "boolean":
		return "false"
	case "byte", "short", "int", "long":
		return number(t, true)
	case "float", "double":
		return number(t, false)
	case "string":
		return `""`
	case "list", "byte_array", "int_array", "long_array", "tuple":
		return "[]"
	case "literal":
		if t.Literal != nil {
			if s := t.Literal.String(); s != "" || string(t.Literal.Value) == `""` {
				return enumText(s)
			}
			return strings.TrimSpace(string(t.Literal.Value))
		}
	case "enum":
		if len(t.Values) > 0 {
			ev := t.Values[0]
			if t.EnumKind == "string" {
				return enumText(ev.String())
			}
			return ev.String()
		}
	case "union":
		if m, _ := c.simplest(v, depth); m.T != nil && m.T.Kind != "union" {
			return c.placeholder(m, depth+1)
		}
	}
	return "{}"
}

// number is 0, or the bound nearest 0 when the range excludes it.
func number(t *spec.Type, integer bool) string {
	r := t.ValueRange
	x := 0.0
	if r != nil {
		if r.Min != nil && (*r.Min > 0 || *r.Min == 0 && r.Kind&1 != 0) {
			x = *r.Min
			if r.Kind&1 != 0 {
				x++
				if !integer {
					x = *r.Min + 0.1
				}
			}
		} else if r.Max != nil && (*r.Max < 0 || *r.Max == 0 && r.Kind&2 != 0) {
			x = *r.Max
			if r.Kind&2 != 0 {
				x--
				if !integer {
					x = *r.Max - 0.1
				}
			}
		}
	}
	if integer {
		return strconv.FormatInt(int64(math.Ceil(x)), 10)
	}
	return strconv.FormatFloat(x, 'f', -1, 64)
}

// simplest picks the union member (nested unions flattened) that is
// easiest to write: a number, then true/false, a list, text, an object.
func (c *Checker) simplest(v tview, depth int) (tview, int) {
	v = c.norm(v)
	if v.T == nil || depth > 8 {
		return tview{}, len(unionPref)
	}
	if v.T.Kind == "union" {
		best, rank := tview{}, len(unionPref)
		for _, m := range v.T.Members {
			if mv, r := c.simplest(tview{T: m, Env: v.Env}, depth+1); r < rank {
				best, rank = mv, r
			}
		}
		return best, rank
	}
	for i, k := range unionPref {
		if strings.Contains(" "+k+" ", " "+v.T.Kind+" ") {
			return v, i
		}
	}
	return v, len(unionPref) - 1
}

var unionPref = []string{"byte short int long float double", "boolean", "list tuple", "literal enum", "string", "struct"}

func enumText(s string) string {
	if lexesBare(s) {
		return s
	}
	return strconv.Quote(s)
}

// removeFix deletes the line holding sp, when nothing else is on it.
func (c *Checker) removeFix(sp syntax.Span, title string) []syntax.Fix {
	src := c.Src
	if src == nil || sp.Start < 0 || sp.End > len(src) || sp.Start >= sp.End {
		return nil
	}
	ls := sp.Start
	for ls > 0 && (src[ls-1] == ' ' || src[ls-1] == '\t') {
		ls--
	}
	le := sp.End
	for le < len(src) && (src[le] == ' ' || src[le] == '\t' || src[le] == ',') {
		le++
	}
	if !(ls == 0 || src[ls-1] == '\n') || !(le == len(src) || src[le] == '\n') {
		return nil // shares its line: removing it could leave stray commas
	}
	start, end := ls, le
	if end < len(src) {
		end++ // the newline
	}
	return []syntax.Fix{{Title: title, Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: start, End: end}, Text: ""}}}}
}
