package spec

import (
	"fmt"
	"strconv"
	"strings"
)

// Format renders a type in mcdoc-like syntax for hover and debugging.
// Nested structs are expanded up to depth levels, then shown as `struct`.
func Format(t *Type, depth int) string {
	var sb strings.Builder
	formatType(&sb, t, depth, 0)
	return sb.String()
}

// ShortPath turns "::java::data::loot::LootTable" into "LootTable".
func ShortPath(path string) string {
	if i := strings.LastIndex(path, "::"); i >= 0 {
		return path[i+2:]
	}
	return path
}

func formatType(sb *strings.Builder, t *Type, depth, indent int) {
	if t == nil {
		sb.WriteString("?")
		return
	}
	if id, ok := t.Attr("id"); ok {
		if s, ok := id.LiteralString(); ok {
			fmt.Fprintf(sb, "#[id=%q] ", s)
		} else if id.Value != nil && id.Value.Kind == "tree" {
			var opts []string
			for _, k := range []string{"registry", "tags", "path"} {
				if v := id.Value.Tree[k]; v != nil && v.Literal != nil {
					opts = append(opts, fmt.Sprintf("%s=%q", k, v.Literal.String()))
				}
			}
			sb.WriteString("#[id(" + strings.Join(opts, ", ") + ")] ")
		} else {
			sb.WriteString("#[id] ")
		}
	}
	switch t.Kind {
	case "reference":
		sb.WriteString(ShortPath(t.Path))
	case "string", "boolean", "any":
		sb.WriteString(t.Kind)
		writeRange(sb, t.LengthRange, " @ ")
	case "byte", "short", "int", "long", "float", "double":
		sb.WriteString(t.Kind)
		writeRange(sb, t.ValueRange, " @ ")
	case "byte_array", "int_array", "long_array":
		sb.WriteString(strings.TrimSuffix(t.Kind, "_array") + "[]")
	case "literal":
		if t.Literal != nil {
			sb.Write(t.Literal.Value)
		}
	case "enum":
		var vs []string
		for _, v := range t.Values {
			vs = append(vs, strconv.Quote(v.String()))
		}
		if len(vs) > 8 {
			vs = append(vs[:8], "...")
		}
		sb.WriteString("(" + strings.Join(vs, " | ") + ")")
	case "list":
		sb.WriteString("[")
		formatType(sb, t.Item, depth, indent)
		sb.WriteString("]")
		writeRange(sb, t.LengthRange, " @ ")
	case "tuple":
		sb.WriteString("[")
		for i, it := range t.Items {
			if i > 0 {
				sb.WriteString(", ")
			}
			formatType(sb, it, depth, indent)
		}
		sb.WriteString("]")
	case "union":
		if len(t.Members) == 0 {
			sb.WriteString("()")
			return
		}
		sb.WriteString("(")
		for i, m := range t.Members {
			if i > 0 {
				sb.WriteString(" | ")
			}
			formatType(sb, m, depth, indent)
		}
		sb.WriteString(")")
	case "struct":
		if depth <= 0 {
			sb.WriteString("struct {...}")
			return
		}
		sb.WriteString("struct {\n")
		pad := strings.Repeat("  ", indent+1)
		for _, f := range t.Fields {
			sb.WriteString(pad)
			if f.Kind == "spread" {
				sb.WriteString("...")
				formatType(sb, f.Type, depth-1, indent+1)
			} else {
				if f.KeyType != nil {
					sb.WriteString("[")
					formatType(sb, f.KeyType, 0, indent+1)
					sb.WriteString("]")
				} else {
					sb.WriteString(f.Key)
				}
				if f.Optional {
					sb.WriteString("?")
				}
				sb.WriteString(": ")
				formatType(sb, f.Type, depth-1, indent+1)
			}
			sb.WriteString(",\n")
		}
		sb.WriteString(strings.Repeat("  ", indent) + "}")
	case "dispatcher":
		sb.WriteString(t.Registry)
		writeIndices(sb, t.ParallelIndices)
	case "indexed":
		formatType(sb, t.Child, depth, indent)
		writeIndices(sb, t.ParallelIndices)
	case "concrete":
		formatType(sb, t.Child, depth, indent)
		sb.WriteString("<")
		for i, a := range t.TypeArgs {
			if i > 0 {
				sb.WriteString(", ")
			}
			formatType(sb, a, depth, indent)
		}
		sb.WriteString(">")
	case "template":
		var ps []string
		for _, p := range t.TypeParams {
			ps = append(ps, ShortPath(p.Path))
		}
		sb.WriteString("<" + strings.Join(ps, ", ") + "> ")
		formatType(sb, t.Child, depth, indent)
	default:
		sb.WriteString(t.Kind)
	}
}

func writeIndices(sb *strings.Builder, idx []*Index) {
	sb.WriteString("[")
	for i, x := range idx {
		if i > 0 {
			sb.WriteString(", ")
		}
		if x.Kind == "static" {
			sb.WriteString(x.Value)
		} else {
			sb.WriteString("[")
			for j, a := range x.Accessor {
				if j > 0 {
					sb.WriteString(".")
				}
				sb.WriteString(strings.Trim(string(a), `"`))
			}
			sb.WriteString("]")
		}
	}
	sb.WriteString("]")
}

func writeRange(sb *strings.Builder, r *Range, prefix string) {
	if r == nil {
		return
	}
	sb.WriteString(prefix)
	num := func(f *float64) string {
		if f == nil {
			return ""
		}
		return strconv.FormatFloat(*f, 'g', -1, 64)
	}
	op := ".."
	switch r.Kind {
	case 1:
		op = "<.."
	case 2:
		op = "..<"
	case 3:
		op = "<..<"
	}
	if r.Min != nil && r.Max != nil && *r.Min == *r.Max && r.Kind == 0 {
		sb.WriteString(num(r.Min))
		return
	}
	sb.WriteString(num(r.Min) + op + num(r.Max))
}
