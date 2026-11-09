package sema

import (
	"slices"
	"strings"

	"github.com/minecraftmetascript/mms/engine/command"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// CheckSelectors checks every target selector written in f, wherever it
// is: in a const, a Template, a Blueprint field or function code. A raw
// command's selectors are checked with the command. Checking the source,
// rather than each place a const's value lands, reports a problem once,
// where it is written.
func CheckSelectors(f *syntax.File) []syntax.Diagnostic {
	var out []syntax.Diagnostic
	errorf := func(sp syntax.Span, msg string, fixes ...syntax.Fix) {
		out = append(out, syntax.Diagnostic{Code: "mms/selector", Severity: syntax.SeverityError, Message: msg, Sp: sp, Fixes: fixes})
	}
	names := command.SelectorOptionNames()
	syntax.Walk(f, func(n syntax.Node) bool {
		var ix *syntax.Index
		switch x := n.(type) {
		case *syntax.Selector:
			if len(x.Name) != 1 || !strings.Contains("aenprs", x.Name) {
				errorf(x.Sp, "unknown selector '@"+x.Name+"'; expected @a, @e, @n, @p, @r or @s")
			}
			return true
		case *syntax.Index:
			if _, ok := x.X.(*syntax.Selector); !ok {
				return true
			}
			ix = x
		default:
			return true
		}
		for _, a := range ix.Args {
			if a.Name == nil {
				errorf(a.Sp, "expected a selector option such as tag=boss")
				continue
			}
			name := a.Name.Text
			if slices.Contains(names, name) || strings.Contains(name, "${") {
				continue
			}
			msg := "unknown selector option '" + name + "'"
			var fixes []syntax.Fix
			if best := closest(name, names); best != "" {
				msg += "; did you mean '" + best + "'?"
				fixes = append(fixes, syntax.Fix{Title: "Change to " + best, Edits: []syntax.TextEdit{{Sp: a.Name.Sp, Text: best}}})
			}
			errorf(a.Name.Sp, msg, fixes...)
		}
		return true
	})
	return out
}

// SelectorType is the language's Selector type: a string that is a target
// selector, which the game's schema marks #[entity]. A selector written
// as @e[...] is checked where that type is expected.
func SelectorType() *spec.Type {
	return &spec.Type{Kind: "string", Attributes: []spec.Attribute{{Name: "entity"}}}
}

// isSelectorType tells whether t is a target selector's string.
func isSelectorType(t *spec.Type) bool {
	if t == nil || t.Kind != "string" {
		return false
	}
	_, ok := t.Attr("entity")
	return ok
}
