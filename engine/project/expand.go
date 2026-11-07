package project

import (
	"strings"

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
		// A Book's text is checked only here, for its layout and markup
		// problems. They come from the values the evaluator worked out,
		// since the body as written can hold a loop's ${...}.
		w := c
		if len(inst.Written) > 0 {
			w = newChecker(ns)
			w.Src = u.File.Src
			body := &syntax.Block{Closed: true, Sp: inst.D.Body.Sp}
			for _, it := range inst.D.Body.Items {
				f, ok := it.(*syntax.FieldStmt)
				if !ok || f.Key == nil {
					continue
				}
				// A value from elsewhere, such as a const in another file,
				// keeps the field as written: its spans are not here.
				x := inst.Written[f.Key.Text]
				if x == nil || !within(x.Span(), f.Sp) {
					x = f.Value
				}
				body.Items = append(body.Items, &syntax.FieldStmt{Key: f.Key, KeyQuoted: f.KeyQuoted, Value: x, Sp: f.Sp})
			}
			w.CheckType(&syntax.Object{Body: body, Sp: body.Sp}, tenv.InstanceType(inst.Blueprint), inst.Blueprint.D.NameText())
			for k, ok := range w.Deps {
				if c.Deps == nil {
					c.Deps = map[string]bool{}
				}
				c.Deps[k] = ok
			}
		}
		d := &Decl{Unit: u, D: inst.D, NS: ns, Info: c.Info, Deps: c.Deps, Blueprint: inst.Blueprint}
		if id, ok := inst.Self.Vals[0].(*eval.Word); ok {
			d.id = id.Text
		}
		// The item stack the name stands for, if it stands for one, for
		// an editor to list and preview.
		if inst.Name != nil {
			if st := c.ConvertedStack(&syntax.ResLoc{Text: d.ID(), Sp: inst.Name.Sp}); st != nil {
				d.Item = &sema.CustomItemOut{Stack: st, Book: w.Book}
			}
		}
		d.Diags, d.Extern = w.BookDiags, w.Extern
		r.Decls[inst.D] = d
		r.Instances = append(r.Instances, d)
	}
}

// templateDecls checks the body of each Template the user wrote against
// its return type, for the editor, as instanceDecls does an instance's
// body: completion and hover inside `= Pool(rolls, item)` read the Info.
// A Commands Template's block is function code, so its result is given
// the Function kind. Diagnostics are dropped, since its parameters are
// not names the checker knows; calls are checked where they are made.
func (r *Result) templateDecls(tenv *types.Env, kinds *sema.Kinds, newChecker types.Checkers) {
	units := map[string]*sema.Unit{}
	for _, u := range r.Units {
		units[u.Path] = u
	}
	for _, t := range r.Eval.Templates {
		u := units[t.Path]
		if u == nil || prelude.Is(t.Path) {
			continue
		}
		d := &Decl{Unit: u, D: t.D, NS: r.cfg.Namespace, Template: t}
		switch {
		case t.D.Body != nil:
			d.Kind, _ = kinds.Lookup("Function")
		case t.D.Value != nil:
			c := newChecker(r.cfg.Namespace)
			c.Src = u.File.Src
			if want := tenv.Lookup(t.Path, t.D.Type); want != nil && want.Kind != "any" {
				c.CheckType(t.D.Value, want, t.D.NameText())
			}
			d.Info = c.Info
		default:
			d.Info = map[syntax.Node]*sema.Info{}
		}
		r.Decls[t.D] = d
	}
}

// missingFieldFixes gives a Blueprint instance's missing-field error the
// quick fix that adds the fields, with placeholder values of their types.
func (r *Result) missingFieldFixes(tenv *types.Env, kinds *sema.Kinds, source map[string]*syntax.File) {
	if r.Eval == nil {
		return
	}
	for path, f := range r.Eval.Files {
		src := source[path]
		if src == nil {
			continue
		}
		for i := range f.Diags {
			dg := &f.Diags[i]
			if dg.Code != "mms/missing-field" || len(dg.Fixes) > 0 {
				continue
			}
			_, list, ok := strings.Cut(dg.Message, " needs ")
			if !ok {
				continue
			}
			var missing []string
			for _, m := range strings.Split(list, ", ") {
				missing = append(missing, strings.Trim(m, "'"))
			}
			for _, d := range src.Decls {
				if d.Kind == nil || d.Body == nil || !d.Sp.Contains(dg.Sp.Start) {
					continue
				}
				bp := r.Eval.Blueprints[d.Kind.Name]
				if bp == nil {
					continue
				}
				c := sema.NewChecker(r.pack, kinds, r.Syms, r.cfg.Namespace)
				c.Src = src.Src
				dg.Fixes = c.MissingFieldsFix(d.Body, tenv.InstanceType(bp), missing)
			}
		}
	}
}

// within reports whether span a lies inside span b.
func within(a, b syntax.Span) bool { return a.Start >= b.Start && a.End <= b.End }
