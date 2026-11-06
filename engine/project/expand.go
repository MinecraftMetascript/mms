package project

import (
	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/types"
)

// expand runs the evaluator over the project's MMS files. A file whose
// declarations it changed gets a new unit holding the expanded
// declarations, with the same text and tokens; every other unit is
// returned as it was, so its checked results carry over.
func expand(units []*sema.Unit, ns string, cache *eval.Cache) ([]*sema.Unit, *eval.Result) {
	in := make([]eval.Unit, 0, len(units)+4)
	// The prelude comes first, as a library would: its Templates, Types
	// and Blueprints are in scope everywhere. It writes nothing itself.
	for _, f := range prelude.Parse() {
		in = append(in, eval.Unit{Path: f.Path, File: f.File})
	}
	for _, u := range units {
		if !sema.IsBookFile(u.Path) {
			in = append(in, eval.Unit{Path: u.Path, File: u.File})
		}
	}
	r := eval.Run(in, eval.Options{Namespace: ns, Cache: cache})
	out := make([]*sema.Unit, len(units))
	for i, u := range units {
		out[i] = u
		f := r.Files[u.Path]
		if f == nil || !f.Changed {
			continue
		}
		file := *u.File
		file.Decls = make([]*syntax.Decl, len(f.Decls))
		// Items keep what the evaluator consumed (const, Template, for)
		// and then the expanded declarations, which win where both hold
		// the cursor: the editor finds the checked declaration.
		file.Items = nil
		for _, it := range u.File.Items {
			if d, ok := it.(*syntax.Decl); ok && !definition(d) {
				continue
			}
			if _, ok := it.(*syntax.ForStmt); ok {
				continue
			}
			file.Items = append(file.Items, it)
		}
		for j, d := range f.Decls {
			file.Decls[j] = d.D
			file.Items = append(file.Items, d.D)
		}
		// A Blueprint instance the user wrote comes last, so the editor
		// finds it before what it emitted: its body is checked for the
		// editor against the Blueprint's fields (see instanceDecls).
		for _, inst := range r.Instances {
			if inst.User && inst.Path == u.Path {
				file.Items = append(file.Items, inst.D)
			}
		}
		out[i] = &sema.Unit{Path: u.Path, File: &file}
	}
	return out, r
}

func definition(d *syntax.Decl) bool {
	if d.Kind == nil {
		return false
	}
	switch d.Kind.Name {
	case "Template", "Blueprint", "Type":
		return true
	}
	return false
}

// instanceDecls checks the body of each Blueprint instance the user wrote
// against the Blueprint's fields, for the editor: completion, hover and
// definitions inside it read the Info. Its diagnostics are dropped, since
// the evaluator and types.Check report the instance's problems, and the
// result writes nothing: it is in Decls but not in Order.
func (r *Result) instanceDecls(tenv *types.Env, newChecker types.Checkers) {
	units := map[string]*sema.Unit{}
	for _, u := range r.Units {
		units[u.Path] = u
	}
	for _, inst := range r.Eval.Instances {
		u := units[inst.Path]
		if !inst.User || u == nil || inst.D.Body == nil {
			continue
		}
		ns := r.cfg.Namespace
		if v, ok := inst.Self.Vals[1].(*eval.Str); ok {
			ns = v.S
		}
		c := newChecker(ns)
		obj := &syntax.Object{Body: inst.D.Body, Sp: inst.D.Body.Sp}
		c.CheckType(obj, tenv.InstanceType(inst.Blueprint), inst.Blueprint.D.NameText())
		d := &Decl{Unit: u, D: inst.D, NS: ns, Info: c.Info, Deps: c.Deps, Blueprint: inst.Blueprint}
		// The item stack the name stands for, if it stands for one, for
		// an editor to list and preview.
		if inst.Name != nil {
			if st := c.ConvertedStack(&syntax.ResLoc{Text: d.ID(), Sp: inst.Name.Sp}); st != nil {
				d.Item = &sema.CustomItemOut{Stack: st, Book: c.Book}
			}
		}
		// A Book's text is checked only here: its markup problems.
		d.Diags, d.Extern = c.BookDiags, c.Extern
		r.Decls[inst.D] = d
		r.Instances = append(r.Instances, d)
	}
}
