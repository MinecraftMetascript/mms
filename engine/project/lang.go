package project

import (
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// checkLangKeys reports a translation key that two Lang declarations of
// the same language set, such as a CustomItem's name and a Lang written
// by hand: one would silently replace the other.
func (r *Result) checkLangKeys() {
	type setBy struct{ d *Decl }
	keys := map[string]map[string]setBy{} // ns:language -> key -> first Lang
	for _, d := range r.Order {
		if !d.Kind.Merges() {
			continue
		}
		o, ok := d.Value.(*jsonv.Object)
		if !ok {
			continue
		}
		id := d.ID()
		if keys[id] == nil {
			keys[id] = map[string]setBy{}
		}
		for _, m := range o.Members {
			first, dup := keys[id][m.Key]
			if !dup {
				keys[id][m.Key] = setBy{d}
				continue
			}
			if r.langErr == nil {
				r.langErr = map[string][]syntax.Diagnostic{}
			}
			r.langErr[d.Unit.Path] = append(r.langErr[d.Unit.Path], syntax.Diagnostic{
				Code: "mms/duplicate", Severity: syntax.SeverityError, Sp: langKeySpan(d.D, m.Key),
				Message: "Lang " + d.D.NameText() + " in " + first.d.Unit.Path + " already sets '" + m.Key + "'; remove one of them",
			})
		}
	}
}

// langKeySpan is where a Lang declaration writes a key, or its name.
func langKeySpan(d *syntax.Decl, key string) syntax.Span {
	var items []syntax.Stmt
	if d.Body != nil {
		items = d.Body.Items
	} else if o, ok := d.Value.(*syntax.Object); ok {
		items = o.Body.Items
	}
	for _, it := range items {
		if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil && f.Key.Text == key {
			return f.Key.Sp
		}
	}
	if d.Name != nil {
		return d.Name.Sp
	}
	return d.Sp
}
