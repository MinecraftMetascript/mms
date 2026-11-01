package project

import (
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// checkCustomItemLang reports a CustomItem name whose translation key a
// Lang declaration of the same language also sets: one would silently
// replace the other.
func (r *Result) checkCustomItemLang() {
	type langDecl struct {
		d *Decl
		o *jsonv.Object
	}
	langs := map[string]langDecl{} // ns:language
	for _, d := range r.Order {
		if d.Kind.Resource == "lang" && d.Kind.Assets != "" {
			if o, ok := d.Value.(*jsonv.Object); ok {
				langs[d.ID()] = langDecl{d, o}
			}
		}
	}
	for _, d := range r.Order {
		it := d.Item
		if it == nil || it.LangKey == "" {
			continue
		}
		for _, lc := range it.Languages() {
			ns, _ := splitID(d.ID())
			l, ok := langs[ns+":"+lc]
			if !ok || l.o.Get(it.LangKey) == nil {
				continue
			}
			sp := d.D.Name.Sp
			for _, st := range customItemBody(d.D) {
				if f, ok := st.(*syntax.FieldStmt); ok && f.Key != nil && f.Key.Text == "name" {
					sp = f.Value.Span()
				}
			}
			if r.langErr == nil {
				r.langErr = map[string][]syntax.Diagnostic{}
			}
			r.langErr[d.Unit.Path] = append(r.langErr[d.Unit.Path], syntax.Diagnostic{
				Code: "mms/duplicate", Severity: syntax.SeverityError, Sp: sp,
				Message: "Lang " + lc + " in " + l.d.Unit.Path + " already sets '" + it.LangKey + "'; remove that entry, or leave out this name",
			})
		}
	}
}

func customItemBody(d *syntax.Decl) []syntax.Stmt {
	if d.Body != nil {
		return d.Body.Items
	}
	if o, ok := d.Value.(*syntax.Object); ok {
		return o.Body.Items
	}
	return nil
}
