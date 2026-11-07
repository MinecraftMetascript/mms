package sema

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// Item stacks a declaration's name stands for. CustomItem and Book are
// Blueprints in the prelude (prelude/items.mms, books.mms), whose `as`
// conversions give their names a stack; what is here describes that
// stack for the editor and renders stacks as SNBT for commands.

// CustomItemOut is a checked custom item.
type CustomItemOut struct {
	// Model is the generated item model under models/, nil when the item
	// names its own model or item_model.
	Model jsonv.Value
	// ModelPath is Model's path under models/: "item/<name>".
	ModelPath string
	// ItemModel is the generated item model definition under items/, nil
	// when item_model names an existing one.
	ItemModel jsonv.Value
	// LangKey is the translation key the item_name component uses, and
	// Names its text by language; both are empty without a name.
	LangKey string
	Names   map[string]string
	// Stack is the item stack the name stands for: {id, components}.
	Stack *jsonv.Object
	// Book is a Book's layout: its pages that run past the bottom of the
	// page, and where each starts.
	Book *BookLayout
}

// Languages lists the item's languages, sorted.
func (o *CustomItemOut) Languages() []string {
	var out []string
	for l := range o.Names {
		out = append(out, l)
	}
	sort.Strings(out)
	return out
}

func idType(registry string) *spec.Type {
	lit, _ := json.Marshal(registry)
	return &spec.Type{Kind: "string", Attributes: []spec.Attribute{{Name: "id", Value: &spec.Type{Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: lit}}}}}
}

// StackSNBT writes an item stack as a command's item argument:
// id[component=value,...].
func StackSNBT(st *jsonv.Object) string {
	id, _ := st.Get("id").(jsonv.String)
	var sb strings.Builder
	sb.WriteString(strings.TrimPrefix(string(id), "minecraft:"))
	if comps, ok := st.Get("components").(*jsonv.Object); ok && len(comps.Members) > 0 {
		sb.WriteByte('[')
		for i, m := range comps.Members {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(m.Key)
			sb.WriteByte('=')
			writeSNBT(&sb, m.Value)
		}
		sb.WriteByte(']')
	}
	return sb.String()
}

var bareSNBTKey = regexp.MustCompile(`^[A-Za-z0-9_\-.+]+$`)

func writeSNBT(sb *strings.Builder, v jsonv.Value) {
	switch x := v.(type) {
	case *jsonv.Object:
		sb.WriteByte('{')
		for i, m := range x.Members {
			if i > 0 {
				sb.WriteByte(',')
			}
			if bareSNBTKey.MatchString(m.Key) {
				sb.WriteString(m.Key)
			} else {
				sb.WriteString(quoteSNBT(m.Key))
			}
			sb.WriteByte(':')
			writeSNBT(sb, m.Value)
		}
		sb.WriteByte('}')
	case jsonv.Array:
		sb.WriteByte('[')
		for i, e := range x {
			if i > 0 {
				sb.WriteByte(',')
			}
			writeSNBT(sb, e)
		}
		sb.WriteByte(']')
	case jsonv.String:
		sb.WriteString(quoteSNBT(string(x)))
	case jsonv.Number:
		sb.WriteString(string(x))
	case jsonv.Bool:
		if x {
			sb.WriteString("true")
		} else {
			sb.WriteString("false")
		}
	}
}

// quoteSNBT quotes a string for SNBT, escaping control characters so
// that a newline in a name cannot split the command line.
func quoteSNBT(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch {
		case r == '\\':
			b.WriteString(`\\`)
		case r == '"':
			b.WriteString(`\"`)
		case r == '\n':
			b.WriteString(`\n`)
		case r == '\t':
			b.WriteString(`\t`)
		case r < 0x20 || (r >= 0x7f && r <= 0x9f):
			fmt.Fprintf(&b, `\u%04x`, r)
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
	return b.String()
}
