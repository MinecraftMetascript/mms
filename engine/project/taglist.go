package project

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// completeTagListFixes finishes the fixes the checker starts where a list
// of ids holds a tag but the field takes a tag too (see
// sema.TagListFixTitle): the list moves into a new tag declaration, such
// as `BiomeTag village_biomes = [#minecraft:is_plains, desert]`, written
// above the declaration that holds the list, and the list becomes
// `#village_biomes`. A fix this file cannot complete, as for a list that
// is not written in it, is dropped. The diagnostics are copied, since
// they can be shared with other results.
func (r *Result) completeTagListFixes(u *sema.Unit, ds []syntax.Diagnostic) []syntax.Diagnostic {
	var out []syntax.Diagnostic
	for i, d := range ds {
		var fixes []syntax.Fix
		changed := false
		for _, f := range d.Fixes {
			kind, ok := sema.TagListKind(f)
			if !ok {
				fixes = append(fixes, f)
				continue
			}
			changed = true
			if done, ok := r.tagListFix(u, f.Edits[0].Sp, kind); ok {
				fixes = append(fixes, done)
			}
		}
		if !changed {
			continue
		}
		if out == nil {
			out = append([]syntax.Diagnostic(nil), ds...)
		}
		out[i].Fixes = fixes
	}
	if out == nil {
		return ds
	}
	return out
}

// tagListFix moves the list at sp in u into a new declaration of kind.
func (r *Result) tagListFix(u *sema.Unit, sp syntax.Span, kind string) (syntax.Fix, bool) {
	// The file as written: a Blueprint instance is not among the
	// expanded file's declarations.
	f := r.Source[u.Path]
	if f == nil {
		return syntax.Fix{}, false
	}
	src := f.Src
	if sp.Start < 0 || sp.End > len(src) || sp.Start >= sp.End {
		return syntax.Fix{}, false
	}
	var top *syntax.Decl
	for _, d := range f.Decls {
		if d.Sp.Start <= sp.Start && sp.End <= d.Sp.End {
			top = d
		}
	}
	if top == nil {
		return syntax.Fix{}, false
	}
	// The list as written, and the field that holds it, which names the
	// new tag: `village` with `biomes = [...]` makes village_biomes.
	var list *syntax.List
	key := ""
	syntax.Walk(top, func(n syntax.Node) bool {
		if list != nil {
			return false
		}
		switch n := n.(type) {
		case *syntax.FieldStmt:
			if l, ok := n.Value.(*syntax.List); ok && l.Sp == sp && n.Key != nil {
				list, key = l, n.Key.Text
				return false
			}
		case *syntax.List:
			if n.Sp == sp {
				list = n
				return false
			}
		}
		return true
	})
	if list == nil || string(src[sp.Start]) != "[" {
		return syntax.Fix{}, false
	}
	base := ""
	if top.Name != nil && !top.Name.Templated() {
		base = top.Name.Text
		if i := strings.LastIndexAny(base, ":/"); i >= 0 {
			base = base[i+1:]
		}
	}
	if key == "" {
		k, _ := r.kindOf(kind)
		key = lastSegment(k.TagOf) + "s"
	}
	name := sema.Snake(key)
	if base != "" {
		name = base + "_" + name
	}
	name = r.freeTagName(top, name)

	// Above the declaration and the comments right above it, at the same
	// indent, with a blank line after.
	at := lineStart(src, top.Sp.Start)
	indent := string(src[at:top.Sp.Start])
	if strings.TrimSpace(indent) != "" {
		return syntax.Fix{}, false
	}
	at = commentsAbove(src, at)
	text := fmt.Sprintf("%s%s %s = %s\n\n", indent, kind, name, unindented(src, sp))
	return syntax.Fix{
		Title: fmt.Sprintf("Move the list into %s %s", kind, name),
		Edits: []syntax.TextEdit{
			{Sp: syntax.Span{Start: at, End: at}, Text: text},
			{Sp: sp, Text: "#" + name},
		},
	}, true
}

// kindOf looks up a kind by name.
func (r *Result) kindOf(name string) (sema.Kind, bool) {
	if r.kinds == nil {
		return sema.Kind{}, false
	}
	return r.kinds.Lookup(name)
}

// freeTagName is name, or name_2, name_3 and so on, whichever nothing in
// the declaration's namespace is called yet, as the extract refactorings
// choose.
func (r *Result) freeTagName(top *syntax.Decl, name string) string {
	ns := r.cfg.Namespace
	if d := r.Decls[top]; d != nil && d.NS != "" {
		ns = d.NS
	}
	taken := map[string]bool{}
	if r.Syms != nil {
		for _, s := range r.Syms.All() {
			if s.NS == ns {
				taken[s.Path] = true
			}
		}
	}
	for i := 1; ; i++ {
		n := name
		if i > 1 {
			n = fmt.Sprintf("%s_%d", name, i)
		}
		if !taken[n] {
			return n
		}
	}
}

func lastSegment(s string) string {
	if i := strings.LastIndexByte(s, '/'); i >= 0 {
		return s[i+1:]
	}
	return s
}

// lineStart is the offset of the start of the line holding off.
func lineStart(src []rune, off int) int {
	for off > 0 && src[off-1] != '\n' {
		off--
	}
	return off
}

// commentsAbove moves at, the start of a line, up over the comment lines
// right above it, so they stay with the declaration they describe (as
// the editor's extract refactorings do).
func commentsAbove(src []rune, at int) int {
	for at > 0 {
		prev := lineStart(src, at-1)
		line := strings.TrimSpace(string(src[prev : at-1]))
		if !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "/*") && !strings.HasPrefix(line, "*") {
			break
		}
		at = prev
	}
	return at
}

// unindented is the text of sp with its continuation lines moved left by
// the indent of the line it starts on, so a list written inside a body
// reads right at the top level.
func unindented(src []rune, sp syntax.Span) string {
	ls := lineStart(src, sp.Start)
	end := ls
	for end < sp.Start && (src[end] == ' ' || src[end] == '\t') {
		end++
	}
	prefix := string(src[ls:end])
	lines := strings.Split(string(src[sp.Start:sp.End]), "\n")
	for i := 1; i < len(lines) && prefix != ""; i++ {
		lines[i] = strings.TrimPrefix(lines[i], prefix)
	}
	return strings.Join(lines, "\n")
}
