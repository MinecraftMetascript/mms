// Package importer turns data pack JSON into MMS source. The output uses
// plain object literals, so it builds back to the same JSON; shorthand is
// applied where it is exact.
package importer

import (
	"regexp"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

var (
	pathRe = regexp.MustCompile(`^[a-z_][a-z0-9_]*(/[a-z0-9_./-]+)*$`)
)

// Decl returns the declaration for one resource file: kind, the id's path
// (qualified when it is not in ns, or is not a valid bare name) and value.
func Decl(k sema.Kind, ns, id string, v jsonv.Value) string {
	if o, ok := v.(*jsonv.Object); ok {
		return k.Name + " " + name(ns, id) + " " + Expr(o) + "\n"
	}
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
	if n == ns && pathRe.MatchString(path) && !syntax.Reserved(path) {
		return path
	}
	return n + ":" + path
}

// Expr prints a JSON value as an MMS expression.
func Expr(v jsonv.Value) string { return sema.JSONExpr(v) }
