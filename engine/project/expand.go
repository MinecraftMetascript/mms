package project

import (
	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// expand runs the evaluator over the project's MMS files. A file whose
// declarations it changed gets a new unit holding the expanded
// declarations, with the same text and tokens; every other unit is
// returned as it was, so its checked results carry over.
func expand(units []*sema.Unit, ns string, cache *eval.Cache) ([]*sema.Unit, *eval.Result) {
	in := make([]eval.Unit, 0, len(units))
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
