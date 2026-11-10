package ide

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// A const at the top of a Blueprint's emit block belongs to each
// instance: `ash_vents.noise_entry` reads the one CustomBiome ash_vents
// emitted. The evaluator replaces the use with its value, so the editor
// finds these on the source as written.

// instConst is a use of name.const in a file.
type instConst struct {
	Inst *eval.Instance
	// Const is the const in the Blueprint, nil when the name does not
	// name one of the instance's consts.
	Const *syntax.ConstStmt
	// OnName is set when the use is at the instance's name, not at
	// the const's.
	OnName bool
	Sp     syntax.Span // the word at the cursor
}

// instancesNamed are the Blueprint instances the user wrote with this
// name, which file from reads as name.const unless it binds the name
// itself or has a const of that name.
func (s *Session) instancesNamed(from, name string) []*eval.Instance {
	if s.R.Eval == nil || s.R.Eval.Const(from, name) != nil {
		return nil
	}
	var out []*eval.Instance
	for _, in := range s.R.Eval.Instances {
		if in.User && in.Name != nil && in.Name.Text == name && in.D != nil && in.D.Name != nil && !in.D.Name.Templated() {
			out = append(out, in)
		}
	}
	return out
}

// instConstAt finds the name.const use at off.
func (s *Session) instConstAt(path string, off int) (instConst, bool) {
	src := s.R.Source[path]
	if src == nil || s.R.Eval == nil {
		return instConst{}, false
	}
	var found instConst
	ok := false
	for _, it := range src.Items {
		if !within(it.Span(), off) {
			continue
		}
		bound := boundNames(it)
		syntax.Walk(it, func(n syntax.Node) bool {
			m, isMember := n.(*syntax.Member)
			if !isMember || !within(m.Sp, off) || m.Name == nil {
				return true
			}
			id, isIdent := m.X.(*syntax.Ident)
			if !isIdent || bound[id.Name] {
				return true
			}
			onName := within(id.Sp, off)
			if !onName && !within(m.Name.Sp, off) {
				return true
			}
			insts := s.instancesNamed(path, id.Name)
			if len(insts) == 0 {
				return true
			}
			found = instConst{Inst: insts[0], OnName: onName, Sp: m.Name.Sp}
			if onName {
				found.Sp = id.Sp
			}
			for _, in := range insts {
				if st := in.ConstDefs[m.Name.Text]; st != nil {
					found.Inst, found.Const = in, st
					break
				}
			}
			ok = true
			return false
		})
	}
	return found, ok
}

// instConstDefinition is where the name.const at off is declared: the
// instance, or the const in its Blueprint.
func (s *Session) instConstDefinition(path string, off int) (Location, bool) {
	u, ok := s.instConstAt(path, off)
	if !ok {
		return Location{}, false
	}
	if u.OnName || u.Const == nil {
		return Location{Path: u.Inst.Path, Span: u.Inst.Name.Sp}, true
	}
	return Location{Path: u.Inst.Blueprint.Path, Span: u.Const.Name.Sp}, true
}

// instConstHover describes the name.const at off.
func (s *Session) instConstHover(path string, off int) (string, syntax.Span, bool) {
	u, ok := s.instConstAt(path, off)
	if !ok {
		return "", syntax.Span{}, false
	}
	bp := u.Inst.Blueprint.D.NameText()
	if u.OnName || u.Const == nil {
		md := "```mms\n" + bp + " " + u.Inst.Name.Text + "\n```"
		if doc, _ := docText(u.Inst.D.Doc); doc != "" {
			md += "\n\n" + doc
		}
		return md, u.Sp, true
	}
	head := "const " + u.Const.Name.Text
	if u.Const.Type != nil {
		head += ": " + syntax.TypeText(u.Const.Type)
	}
	if u.Const.Value != nil {
		v := format.Expr(u.Const.Value)
		if len(v) > 200 || strings.Contains(v, "\n") {
			v = strings.SplitN(v, "\n", 2)[0] + " …"
		}
		head += " = " + v
	}
	md := "```mms\n" + head + "\n```\n\nEmitted by " + bp + " " + u.Inst.Name.Text + "."
	if doc, _ := docText(u.Const.Doc); doc != "" {
		md += "\n\n" + doc
	}
	return md, u.Sp, true
}

// instConstCandidates completes `name.|` when name is a Blueprint
// instance: the consts its emit blocks declare.
func (s *Session) instConstCandidates(a *at, off int) ([]sema.Candidate, bool) {
	f := a.u.File
	dot := prevToken(f, off)
	if w, ok := wordAt(f, off); ok {
		dot = prevToken(f, w.Start)
	}
	if dot == nil || dot.Kind != parser.MMSLexerDOT {
		return nil, false
	}
	recv := prevToken(f, dot.Sp.Start)
	if recv == nil || recv.Kind != parser.MMSLexerIDENT || recv.Sp.End != dot.Sp.Start {
		return nil, false
	}
	name := string(f.Src[recv.Sp.Start:recv.Sp.End])
	if len(a.path) > 1 && boundNames(a.path[1])[name] {
		return nil, false
	}
	insts := s.instancesNamed(a.u.Path, name)
	if len(insts) == 0 {
		return nil, false
	}
	var out []sema.Candidate
	seen := map[string]bool{}
	for _, in := range insts {
		if in.Consts == nil {
			continue
		}
		for _, k := range in.Consts.Keys {
			st := in.ConstDefs[k]
			if seen[k] || st == nil {
				continue
			}
			seen[k] = true
			c := sema.Candidate{Label: k, Kind: sema.CandID, Detail: "const of " + in.Blueprint.D.NameText(), Rank: 1}
			if st.Type != nil {
				c.Detail = syntax.TypeText(st.Type) + ", " + c.Detail
			}
			c.Doc, _ = docText(st.Doc)
			out = append(out, c)
		}
	}
	return out, true
}
