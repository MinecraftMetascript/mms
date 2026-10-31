// Package importer turns data pack JSON into MMS source. The output uses
// plain object literals, so it builds back to the same JSON; sugar is
// applied where it is exact.
package importer

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
)

var (
	identRe  = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
	keywords = map[string]bool{"let": true, "if": true, "else": true, "while": true, "for": true, "in": true,
		"return": true, "break": true, "continue": true, "true": true, "false": true}
	resLocRe = regexp.MustCompile(`^[a-z0-9_.-]+:[a-z0-9_./-]+$`)
	pathRe   = regexp.MustCompile(`^[a-z_][a-z0-9_]*(/[a-z0-9_./-]+)*$`)
)

// Decl returns the declaration for one resource file: kind, the id's path
// (qualified when it is not in ns, or is not a valid bare name) and value.
func Decl(k sema.Kind, ns, id string, v jsonv.Value) string {
	return k.Name + " " + name(ns, id) + " = " + Expr(v) + "\n"
}

// TagDecl returns the declaration for a tag file.
func TagDecl(k sema.Kind, ns, id string, v jsonv.Value) string {
	o, _ := v.(*jsonv.Object)
	var values jsonv.Value = jsonv.Array{}
	if o != nil {
		if vs := o.Get("values"); vs != nil {
			values = vs
		}
		if r, ok := o.Get("replace").(jsonv.Bool); ok && bool(r) {
			// replace has no syntax yet; keep the whole object.
			return Decl(k, ns, id, v)
		}
	}
	return k.Name + " " + name(ns, id) + " = " + Expr(values) + "\n"
}

// FunctionDecl returns a Function declaration holding the commands of an
// .mcfunction file; comments and blank lines are kept as MMS comments.
func FunctionDecl(ns, id, text string) string {
	var sb strings.Builder
	sb.WriteString("Function " + name(ns, id) + " {\n")
	for _, line := range strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n") {
		line = strings.TrimSpace(line)
		switch {
		case line == "":
		case strings.HasPrefix(line, "$"):
			// Macro lines have no MMS syntax yet; keep them visible.
			sb.WriteString("  // mms-import: macro line not supported: " + line + "\n")
		case strings.HasPrefix(line, "#"):
			sb.WriteString("  // " + strings.TrimSpace(strings.TrimPrefix(line, "#")) + "\n")
		default:
			sb.WriteString("  /" + strings.TrimPrefix(line, "/") + "\n")
		}
	}
	sb.WriteString("}\n")
	return sb.String()
}

func name(ns, id string) string {
	n, path, ok := strings.Cut(id, ":")
	if !ok {
		n, path = "minecraft", id
	}
	if n == ns && pathRe.MatchString(path) && !keywords[path] {
		return path
	}
	return n + ":" + path
}

// Expr prints a JSON value as an MMS expression.
func Expr(v jsonv.Value) string {
	var sb strings.Builder
	expr(&sb, v)
	return sb.String()
}

func expr(sb *strings.Builder, v jsonv.Value) {
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
			sb.WriteString(key(m.Key))
			sb.WriteString(" = ")
			expr(sb, m.Value)
		}
		sb.WriteString(" }")
	case jsonv.Array:
		sb.WriteString("[")
		for i, e := range x {
			if i > 0 {
				sb.WriteString(", ")
			}
			expr(sb, e)
		}
		sb.WriteString("]")
	case jsonv.String:
		if resLocRe.MatchString(string(x)) {
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

func key(k string) string {
	if identRe.MatchString(k) && !keywords[k] {
		return k
	}
	return strconv.Quote(k)
}
