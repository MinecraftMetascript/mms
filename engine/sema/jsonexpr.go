package sema

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

var (
	jsonIdentRe  = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	jsonResLocRe = regexp.MustCompile(`^[a-z0-9_.-]+:[a-z0-9_./-]+$`)
)

// JSONExpr prints a JSON value as an MMS expression, with plain object
// literals, so it checks back to the same JSON.
func JSONExpr(v jsonv.Value) string {
	var sb strings.Builder
	jsonExpr(&sb, v)
	return sb.String()
}

func jsonExpr(sb *strings.Builder, v jsonv.Value) {
	switch x := v.(type) {
	case *jsonv.Object:
		if len(x.Members) == 0 {
			sb.WriteString("{}")
			return
		}
		sb.WriteString("{ ")
		for i, m := range x.Members {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(JSONKey(m.Key))
			sb.WriteString(" = ")
			jsonExpr(sb, m.Value)
		}
		sb.WriteString(" }")
	case jsonv.Array:
		sb.WriteString("[")
		for i, e := range x {
			if i > 0 {
				sb.WriteString(", ")
			}
			jsonExpr(sb, e)
		}
		sb.WriteString("]")
	case jsonv.String:
		if jsonResLocRe.MatchString(string(x)) {
			// Resource locations read better bare (minecraft:stone).
			sb.WriteString(string(x))
			return
		}
		sb.WriteString(strconv.Quote(string(x)))
	case jsonv.Number:
		sb.WriteString(string(x))
	case jsonv.Bool:
		sb.WriteString(strconv.FormatBool(bool(x)))
	case jsonv.Null, nil:
		sb.WriteString("null")
	}
}

// JSONKey prints an object key, quoted when it is not a plain name.
func JSONKey(k string) string {
	if jsonIdentRe.MatchString(k) && !syntax.ReservedKey(k) {
		return k
	}
	return strconv.Quote(k)
}
