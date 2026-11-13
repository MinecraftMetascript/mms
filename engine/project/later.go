package project

import (
	"regexp"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// memberRe finds `name.const` inside a ${...} piece of a string or command.
var memberRe = regexp.MustCompile(`([A-Za-z_][\w]*)\.([A-Za-z_][\w]*)`)

// laterDiagnostics runs the check blocks of the project's Blueprint
// instances that call used(), now that the whole project is read, and
// returns what they report, by file.
func (r *Result) laterDiagnostics() map[string][]syntax.Diagnostic {
	r.laterOnce.Do(func() {
		r.later = map[string][]syntax.Diagnostic{}
		if r.Eval == nil {
			return
		}
		var reads, used map[string]bool
		for _, in := range r.Eval.Instances {
			if !in.HasLater() || in.Name == nil || r.Unit(in.Path) == nil {
				continue
			}
			if reads == nil {
				reads, used = r.constReads(), (&lintRun{r: r}).usedIDs()
			}
			ds := in.RunLater(func(u eval.Use) bool {
				if u.Const != "" {
					return reads[in.Name.Text+"."+u.Const]
				}
				return r.emitsUsed(in, u.Kind, used)
			})
			r.later[in.Path] = append(r.later[in.Path], ds...)
		}
	})
	return r.later
}

// constReads are the instance consts the project's files read, as
// name.const: `ash.noise_entries`, or ${ash.noise_entries} in a string or
// command. They are taken from the source as written, so they hold
// however the files were expanded.
func (r *Result) constReads() map[string]bool {
	reads := map[string]bool{}
	for _, f := range r.Source {
		syntax.Walk(f, func(n syntax.Node) bool {
			switch n := n.(type) {
			case *syntax.Member:
				if id, ok := n.X.(*syntax.Ident); ok && n.Name != nil {
					reads[id.Name+"."+n.Name.Text] = true
				}
			case *syntax.String, *syntax.CommandStmt:
				sp := n.Span()
				if sp.End <= len(f.Src) {
					for _, piece := range evalPieceRe.FindAllString(string(f.Src[sp.Start:sp.End]), -1) {
						for _, m := range memberRe.FindAllStringSubmatch(piece, -1) {
							reads[m[1]+"."+m[2]] = true
						}
					}
				}
			}
			return true
		})
	}
	return reads
}

// emitsUsed reports whether something references a declaration of kind
// that in emitted. One that emitted none of that kind has nothing to
// use, and counts as used.
func (r *Result) emitsUsed(in *eval.Instance, kind string, used map[string]bool) bool {
	found := false
	for _, pd := range r.emittedOf(in, kind) {
		found = true
		if id := pd.ID(); id == "" || used[pd.Kind.Folder()+"|"+id] || used[id] {
			return true
		}
	}
	return !found
}

// emittedOf are the declarations of kind that in emitted.
func (r *Result) emittedOf(in *eval.Instance, kind string) []*Decl {
	var out []*Decl
	if f := r.Eval.Files[in.Path]; f != nil {
		for _, od := range f.Decls {
			if od.Instance == in && od.D.Kind != nil && od.D.Kind.Name == kind {
				if pd := r.Decls[od.D]; pd != nil {
					out = append(out, pd)
				}
			}
		}
	}
	return out
}

// checkedUse are the declarations whose use a Blueprint instance's check
// block asks about with used(Kind): the Blueprint reports them itself,
// so lint/unused does not.
func (r *Result) checkedUse() map[*syntax.Decl]bool {
	out := map[*syntax.Decl]bool{}
	if r.Eval == nil {
		return out
	}
	for _, in := range r.Eval.Instances {
		for _, u := range in.Uses() {
			if u.Kind == "" {
				continue
			}
			for _, pd := range r.emittedOf(in, u.Kind) {
				out[pd.D] = true
			}
		}
	}
	return out
}
