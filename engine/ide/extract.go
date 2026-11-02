package ide

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Extraction is a refactoring that moves an inline value into its own
// top-level declaration and refers to it by name where it was.
type Extraction struct {
	Title string
	Kind  string // the new declaration's kind, e.g. "Noise"
	Name  string // the default name, offered for renaming
	Edits []Edit
	// NameAt is the offset of the new declaration's name in the source
	// after the edits, so an editor can start a rename there.
	NameAt int
}

// Extractions lists the inline values around the range [from, to] that
// could be their own declarations, innermost first: values the build
// hoists into mms_gen/ (an inline noise in a density function) and
// values of a declarable kind at a position that also takes its id (an
// inline density function).
func (s *Session) Extractions(path string, from, to int) []Extraction {
	a := s.at(path, from)
	if a == nil || a.decl == nil || a.decl.Info == nil || len(a.path) < 2 {
		return nil
	}
	top, ok := a.path[1].(*syntax.Decl)
	if !ok {
		return nil
	}
	src := a.u.File.Src
	at := declStart(src, top.Sp.Start)
	taken := s.takenNames(a.decl.NS)
	var out []Extraction
	for i := len(a.path) - 1; i >= 2; i-- {
		e, ok := a.path[i].(syntax.Expr)
		if !ok {
			continue
		}
		in := a.decl.Info[e]
		if in == nil || in.Extract == "" {
			continue
		}
		sp := e.Span()
		if sp.Start > from || sp.End < to || usesLocals(a, e) {
			continue
		}
		name := uniqueName(taken, extractName(top, a.path[i-1], in.Extract, s))
		taken[name] = true
		body := reindent(src, sp)
		head := in.Extract + " " + name + " = "
		if in.InlineFn {
			head = in.Extract + " " + name + " "
		}
		out = append(out, Extraction{
			Title:  "Extract " + preview(e, src) + " to " + article(in.Extract) + " " + in.Extract + " declaration",
			Kind:   in.Extract,
			Name:   name,
			NameAt: at + len([]rune(in.Extract)) + 1,
			Edits: []Edit{
				{Path: path, Span: syntax.Span{Start: at, End: at}, Text: head + body + "\n\n"},
				{Path: path, Span: sp, Text: name},
			},
		})
	}
	return out
}

// declStart is where a declaration's text starts: the start of its line,
// moved up over the comments right above it so they stay with it.
func declStart(src []rune, off int) int {
	start := lineBegin(src, off)
	for start > 0 {
		prev := lineBegin(src, start-1)
		line := strings.TrimSpace(string(src[prev : start-1]))
		if !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "/*") && !strings.HasPrefix(line, "*") {
			break
		}
		start = prev
	}
	return start
}

func lineBegin(src []rune, off int) int {
	for off > 0 && src[off-1] != '\n' {
		off--
	}
	return off
}

// reindent returns the text of sp with its continuation lines moved left
// by the indentation of the line it starts on, so a value written nested
// reads right at the top level.
func reindent(src []rune, sp syntax.Span) string {
	ls := lineBegin(src, sp.Start)
	indent := ls
	for indent < sp.Start && (src[indent] == ' ' || src[indent] == '\t') {
		indent++
	}
	prefix := string(src[ls:indent])
	lines := strings.Split(string(src[sp.Start:sp.End]), "\n")
	for i := 1; i < len(lines) && prefix != ""; i++ {
		lines[i] = strings.TrimPrefix(lines[i], prefix)
	}
	return strings.Join(lines, "\n")
}

// preview is a short label for an expression in an action title.
func preview(e syntax.Expr, src []rune) string {
	switch x := e.(type) {
	case *syntax.Call:
		return string(src[x.Fn.Span().Start:x.Fn.Span().End]) + "(...)"
	case *syntax.Object:
		return "{ ... }"
	case *syntax.List:
		return "[ ... ]"
	case *syntax.Trailing:
		return preview(x.X, src) + " { ... }"
	}
	text := strings.Join(strings.Fields(string(src[e.Span().Start:e.Span().End])), " ")
	if r := []rune(text); len(r) > 24 {
		text = string(r[:23]) + "…"
	}
	return text
}

// extractName names an extracted value after its declaration and the
// field or argument it fills, such as hills_final_density, or after its
// kind when it has no name, such as hills_noise.
func extractName(top *syntax.Decl, parent syntax.Node, kind string, s *Session) string {
	owner := top.NameText()
	if i := strings.LastIndex(owner, ":"); i >= 0 {
		owner = owner[i+1:]
	}
	suffix := ""
	switch p := parent.(type) {
	case *syntax.FieldStmt:
		if p.Key != nil {
			suffix = p.Key.Text
		}
	case *syntax.Arg:
		if p.Name != nil {
			suffix = p.Name.Text
		}
	}
	if suffix == "" {
		if k, ok := s.P.Kinds.Lookup(kind); ok {
			suffix = k.Folder()
			if i := strings.LastIndex(suffix, "/"); i >= 0 {
				suffix = suffix[i+1:]
			}
		}
	}
	name := cleanName(owner)
	if suffix = cleanName(suffix); suffix != "" {
		if name != "" {
			name += "_"
		}
		name += suffix
	}
	if name == "" {
		name = "extracted"
	}
	return name
}

// cleanName lowercases s and keeps only characters a resource path takes.
func cleanName(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '_', r == '/', r == '.', r == '-':
			b.WriteRune(r)
		case unicode.IsSpace(r):
			b.WriteRune('_')
		}
	}
	return strings.Trim(b.String(), "/")
}

// takenNames are the paths declared in a namespace, of any kind: a new
// declaration should not share a name with another, even of a different
// kind, to keep references unambiguous to read.
func (s *Session) takenNames(ns string) map[string]bool {
	out := map[string]bool{}
	for _, sym := range s.R.Syms.All() {
		if sym.NS == ns {
			out[sym.Path] = true
		}
	}
	return out
}

func uniqueName(taken map[string]bool, name string) string {
	if !taken[name] {
		return name
	}
	for n := 2; ; n++ {
		if c := name + "_" + strconv.Itoa(n); !taken[c] {
			return c
		}
	}
}

// usesLocals reports whether e reads a function's parameter or local
// variable, which a top-level declaration cannot see.
func usesLocals(a *at, e syntax.Expr) bool {
	found := false
	syntax.Walk(e, func(n syntax.Node) bool {
		if in := a.decl.Info[n]; in != nil && in.Local != nil {
			found = true
		}
		return !found
	})
	return found
}

func article(word string) string {
	if word != "" && strings.ContainsRune("AEIOU", rune(word[0])) {
		return "an"
	}
	return "a"
}
