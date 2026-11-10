package sema

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TagListFixTitle starts the title of the quick fix offered where a list
// of ids holds a tag but the field also takes a single tag: the list can
// become a tag declaration of its own, which may hold tags. The checker
// knows the field takes a tag; it does not know the file, so it gives
// the fix with one edit, the list, and the project completes it: the
// declaration's name and where it goes (see completeTagListFixes in
// package project). The title reads "Move the list into a BiomeTag"
// until then.
const TagListFixTitle = "Move the list into a "

// tagListKind is the tag kind, such as BiomeTag, that list e may become:
// v is the union being checked, m the list member chosen, and some other
// member takes one id of the same registry or a tag of it.
func (c *Checker) tagListKind(e syntax.Expr, v, m tview) string {
	if _, ok := e.(*syntax.List); !ok || m.T == nil || m.T.Kind != "list" || m.T.Item == nil {
		return ""
	}
	item := idOf(c.norm(tview{T: m.T.Item, Env: m.Env}).T)
	if !item.Present || item.Tags != "" || item.Registry == "" {
		return ""
	}
	reg := c.registryName(item.Registry)
	for _, o := range v.T.Members {
		ov := c.norm(tview{T: o, Env: v.Env})
		if ov.T == nil || ov.T.Kind != "string" {
			continue
		}
		id := idOf(ov.T)
		if !id.Present || (id.Tags != "allowed" && id.Tags != "required") || c.registryName(id.Registry) != reg {
			continue
		}
		name := Pascal(lastSeg(reg)) + "Tag"
		if k, ok := c.Kinds.Lookup(name); ok && k.TagOf == reg {
			return name
		}
	}
	return ""
}

// tagListFix gives the errors on the tags in list l, which the checker
// has just reported, the fix that moves the list into a tag of kind.
func (c *Checker) tagListFix(l *syntax.List, kind string, from int) {
	fix := syntax.Fix{Title: TagListFixTitle + kind, Edits: []syntax.TextEdit{{Sp: l.Sp}}}
	for i := from; i < len(c.Diags); i++ {
		d := &c.Diags[i]
		if d.Code != "mms/type-mismatch" || !strings.HasPrefix(d.Message, "a tag is not allowed here") {
			continue
		}
		for _, x := range l.Elems {
			if _, ok := x.(*syntax.Tag); ok && x.Span() == d.Sp {
				d.Fixes = append(d.Fixes, fix)
			}
		}
	}
}

// TagListKind returns the tag kind a fix from tagListFix names, such as
// "BiomeTag".
func TagListKind(f syntax.Fix) (string, bool) {
	k, ok := strings.CutPrefix(f.Title, TagListFixTitle)
	return k, ok && len(f.Edits) == 1 && f.Edits[0].Text == ""
}
