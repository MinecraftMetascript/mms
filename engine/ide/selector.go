package ide

import (
	"github.com/minecraftmetascript/mms/engine/command"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// selectorCompletion completes inside a selector's brackets in MMS code,
// `@e[type=zombie, tag=x]`: the option names, and the values of the
// options that take a known set, such as an entity type or, inside
// scores={...}, an objective. ok is false outside a selector.
func (s *Session) selectorCompletion(a *at, off int) ([]sema.Candidate, bool) {
	var ix *syntax.Index
	at := -1
	for i := len(a.path) - 1; i >= 0; i-- {
		if n, ok := a.path[i].(*syntax.Index); ok {
			if s.isSelector(a.u.Path, n.X) && off > n.X.Span().End {
				ix, at = n, i
				break
			}
		}
	}
	if ix == nil {
		return nil, false
	}
	var arg *syntax.Arg
	if at+1 < len(a.path) {
		arg, _ = a.path[at+1].(*syntax.Arg)
	}
	f := a.u.File
	start := off
	if sp, ok := wordAt(f, off); ok {
		start = sp.Start
	}
	prev := prevToken(f, start)
	switch {
	case arg == nil:
		return selectorOptionCandidates(ix), true
	case arg.Name == nil:
		// A bare word is an option being named: `@e[ty|`.
		if _, ok := arg.Value.(*syntax.Ident); ok || arg.Value == nil {
			return selectorOptionCandidates(ix), true
		}
		return nil, true
	case within(arg.Name.Sp, off) && !(prev != nil && prev.Kind == parser.MMSLexerASSIGN):
		return selectorOptionCandidates(ix), true
	}
	return s.selectorValues(a, arg, at+1), true
}

// isSelector tells whether e is a selector: @e, or the name of a const or
// the call of a Template whose value is one, as in guards[tag=awake].
// isSelector reports whether e, written in file path, is a selector.
func (s *Session) isSelector(path string, e syntax.Expr) bool {
	switch x := e.(type) {
	case *syntax.Selector:
		return true
	case *syntax.Index:
		// Function code is read expanded: the const is its value.
		_, ok := x.X.(*syntax.Selector)
		return ok
	case *syntax.Ident:
		if s.R == nil || s.R.Eval == nil {
			return false
		}
		k := s.R.Eval.Const(path, x.Name)
		if k == nil {
			return false
		}
		if k.S.Type != nil {
			return syntax.TypeText(k.S.Type) == "Selector"
		}
		switch v := k.Lowered().(type) {
		case *syntax.Selector:
			return true
		case *syntax.Index:
			_, ok := v.X.(*syntax.Selector)
			return ok
		}
	case *syntax.Call:
		id, ok := x.Fn.(*syntax.Ident)
		if !ok || s.R == nil || s.R.Eval == nil {
			return false
		}
		t := s.R.Eval.Template(path, id.Name)
		return t != nil && t.D.Type != nil && syntax.TypeText(t.D.Type) == "Selector"
	}
	return false
}

// selectorOptionCandidates lists the options a selector does not set yet.
func selectorOptionCandidates(ix *syntax.Index) []sema.Candidate {
	set := map[string]bool{}
	for _, a := range ix.Args {
		if a.Name != nil {
			set[a.Name.Text] = true
		}
	}
	var out []sema.Candidate
	for i, o := range command.SelectorOptionNames() {
		// Options that may repeat are offered again.
		if set[o] && o != "tag" && o != "type" && o != "predicate" && o != "nbt" {
			continue
		}
		out = append(out, sema.Candidate{Label: o, Kind: sema.CandField, Insert: o + "=$1", Detail: "selector option", Rank: i})
	}
	return out
}

// selectorValues lists the values of one selector option.
func (s *Session) selectorValues(a *at, arg *syntax.Arg, i int) []sema.Candidate {
	var words []string
	kind := sema.CandEnum
	switch arg.Name.Text {
	case "type":
		words = s.P.Pack.RegistryIDs("entity_type")
		kind = sema.CandID
	case "gamemode":
		words = []string{"survival", "creative", "adventure", "spectator"}
	case "sort":
		words = []string{"nearest", "furthest", "random", "arbitrary"}
	case "scores":
		// Inside the braces: the objectives.
		for _, n := range a.path[i:] {
			if _, ok := n.(*syntax.Block); ok {
				return s.symbolNames(a, "Objective", "objective")
			}
		}
		return nil
	case "predicate":
		return s.symbolNames(a, "Predicate", "predicate")
	}
	var out []sema.Candidate
	for j, w := range words {
		out = append(out, sema.Candidate{Label: w, Kind: kind, Rank: j})
	}
	return out
}

// symbolNames lists the project's declarations of one kind, by the name
// code writes for them.
func (s *Session) symbolNames(a *at, kind, detail string) []sema.Candidate {
	ns := s.P.Config.Namespace
	if a.decl != nil {
		ns = a.decl.NS
	}
	var out []sema.Candidate
	for _, sym := range s.R.Syms.All() {
		if sym.Kind.MMS != kind {
			continue
		}
		name := sym.ID()
		if sym.NS == ns {
			name = sym.Path
		}
		out = append(out, sema.Candidate{Label: name, Kind: sema.CandID, Detail: detail, Doc: sema.DocMarkdown(sym.Decl.Doc)})
	}
	return out
}
