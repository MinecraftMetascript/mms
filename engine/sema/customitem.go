package sema

import (
	"strings"

	"github.com/MinecraftMetascript/nbt-go"

	"github.com/minecraftmetascript/mms/engine/jsonv"
)

// Item stacks a declaration's name stands for. CustomItem and Book are
// Blueprints in the prelude (prelude/items.mms, books.mms), whose `as`
// conversions give their names a stack; what is here describes that
// stack for the editor and renders stacks as SNBT for commands.

// CustomItemOut is a checked custom item.
type CustomItemOut struct {
	// Stack is the item stack the name stands for: {id, components}.
	Stack *jsonv.Object
	// Book is a Book's layout: its pages that run past the bottom of the
	// page, and where each starts.
	Book *BookLayout
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

func writeSNBT(sb *strings.Builder, v jsonv.Value) {
	switch x := v.(type) {
	case *jsonv.Object:
		sb.WriteByte('{')
		for i, m := range x.Members {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(nbt.QuoteKey(m.Key))
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
		sb.WriteString(nbt.Quote(string(x)))
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
