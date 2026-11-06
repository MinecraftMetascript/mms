package ide

import (
	"sort"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// declKinds lists what can start a declaration: the kinds, then the
// Blueprints in scope, except those whose doc says @internal.
func (s *Session) declKinds() []sema.Candidate {
	out := s.P.Kinds.DeclKinds()
	if s.R == nil || s.R.Eval == nil {
		return out
	}
	names := make([]string, 0, len(s.R.Eval.Blueprints))
	for n, bp := range s.R.Eval.Blueprints {
		if !internal(bp.D) {
			names = append(names, n)
		}
	}
	sort.Strings(names)
	for _, n := range names {
		bp := s.R.Eval.Blueprints[n]
		out = append(out, sema.Candidate{
			Label: n, Kind: sema.CandKind, Detail: "Blueprint",
			Doc: sema.DocMarkdown(bp.D.Doc), Deprecated: sema.Deprecated(bp.D),
			Insert: n + " ${1:name} {\n\t$0\n}",
		})
	}
	return out
}

func internal(d *syntax.Decl) bool {
	if d == nil || d.Doc == "" {
		return false
	}
	for _, t := range syntax.ParseDoc(d.Doc).Other {
		if t.Name == "internal" {
			return true
		}
	}
	return false
}

// blueprint returns the Blueprint a declaration's kind names, or nil.
func (s *Session) blueprint(d *syntax.Decl) *eval.Blueprint {
	if d == nil || d.Kind == nil || s.R == nil || s.R.Eval == nil {
		return nil
	}
	return s.R.Eval.Blueprints[d.Kind.Name]
}

// instanceHover describes a Blueprint instance at its name.
func instanceHover(d *syntax.Decl) string {
	md := "```mms\n" + d.Kind.Name + " " + d.NameText() + "\n```"
	if doc := sema.DocMarkdown(d.Doc); doc != "" {
		md += "\n" + doc
	}
	return md
}

// convertedAt returns the conversion of the Blueprint instance named at
// off, or nil: a use such as `ruby * 2` or `${ruby}`.
func (s *Session) convertedAt(path string, off int) *sema.Conversion {
	a := s.at(path, off)
	if a == nil || len(a.path) < 2 {
		return nil
	}
	if _, in := a.inCommand(off); in != nil {
		return in.Converted
	}
	for i := len(a.path) - 1; i >= 2; i-- {
		if in := a.info(a.path[i]); in != nil && in.Converted != nil {
			return in.Converted
		}
	}
	return nil
}
