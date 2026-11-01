package ide

import (
	"github.com/minecraftmetascript/mms/engine/command"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// mmsComplete completes the language-level declarations the checker's
// types do not describe: tag entries, an objective's criterion, a type
// annotation, and expressions in function code. Parse recovery drops a
// statement with a missing value, so `let x = |` has no node, and these
// are found from the tokens before the cursor. ok is false when nothing
// here applies.
func (s *Session) mmsComplete(a *at, off int) ([]sema.Candidate, bool) {
	f := a.u.File
	start := off
	if sp, ok := wordAt(f, off); ok {
		start = sp.Start
	}
	prev := prevToken(f, start)
	if prev == nil {
		return nil, false
	}
	// A type annotation: `Global g: |`, `Function f(n: |)`, `let x: |`.
	if prev.Kind == parser.MMSLexerCOLON {
		if a.decl.Kind.MMS == "Global" || a.code() {
			return typeCandidates(), true
		}
	}
	d := a.decl.D
	switch {
	case a.decl.Kind.TagOf != "":
		if d.Value != nil && d.Value.Span().Contains(off) && off > d.Value.Span().Start {
			// Not the tag itself.
			self := "#" + d.NameText()
			var out []sema.Candidate
			for _, c := range a.c.TagEntryCandidates(a.decl.Kind) {
				if c.Label != self && c.Label != "#"+a.decl.NS+":"+d.NameText() {
					out = append(out, c)
				}
			}
			return out, true
		}
	case a.decl.Kind.MMS == "Objective":
		if prev.Kind == parser.MMSLexerASSIGN {
			return criteria(), true
		}
	}
	if !a.code() || d.Body == nil || !d.Body.Sp.Contains(off) {
		return nil, false
	}
	// Arguments of an execute helper, `As(|)`, or of a function call.
	if prev.Kind == parser.MMSLexerLPAREN || prev.Kind == parser.MMSLexerCOMMA {
		name, index, ok := enclosingCall(f, prev)
		if !ok {
			return nil, false
		}
		if sema.ExecHelperNamed(name) != nil {
			return a.c.ExecArgCandidates(name, index), true
		}
		return s.codeValues(a, start), true
	}
	if exprStart(prev.Kind) {
		return s.codeValues(a, start), true
	}
	return nil, false
}

// exprStart reports whether a token of kind k is followed by an
// expression in function code.
func exprStart(k syntax.TokenKind) bool {
	switch k {
	case parser.MMSLexerASSIGN, parser.MMSLexerPLUS_EQ, parser.MMSLexerMINUS_EQ, parser.MMSLexerSTAR_EQ,
		parser.MMSLexerSLASH_EQ, parser.MMSLexerPCT_EQ, parser.MMSLexerEQEQ, parser.MMSLexerNEQ,
		parser.MMSLexerLE, parser.MMSLexerGE, parser.MMSLexerLT, parser.MMSLexerGT,
		parser.MMSLexerANDAND, parser.MMSLexerOROR, parser.MMSLexerBANG, parser.MMSLexerDOTDOT,
		parser.MMSLexerRETURN, parser.MMSLexerIF, parser.MMSLexerWHILE, parser.MMSLexerIN:
		return true
	}
	return mathOp(k)
}

// enclosingCall finds the call whose argument list holds the token open,
// a `(` or `,`, and returns its name and which argument follows open.
func enclosingCall(f *syntax.File, open *syntax.Token) (name string, index int, ok bool) {
	depth := 0
	for i := len(f.Tokens) - 1; i >= 0; i-- {
		t := &f.Tokens[i]
		if t.Sp.Start > open.Sp.Start || t.Hidden || t.Synthetic {
			continue
		}
		switch t.Kind {
		case parser.MMSLexerRPAREN, parser.MMSLexerRBRACK, parser.MMSLexerRBRACE:
			depth++
		case parser.MMSLexerLBRACK, parser.MMSLexerLBRACE:
			if depth == 0 {
				return "", 0, false
			}
			depth--
		case parser.MMSLexerCOMMA:
			if depth == 0 {
				index++
			}
		case parser.MMSLexerLPAREN:
			if depth > 0 {
				depth--
				continue
			}
			if fn := prevToken(f, t.Sp.Start); fn != nil && fn.Kind == parser.MMSLexerIDENT && fn.Sp.End == t.Sp.Start {
				return string(f.Src[fn.Sp.Start:fn.Sp.End]), index, true
			}
			// A parenthesised expression.
			return "", 0, true
		}
	}
	return "", 0, false
}

// codeValues lists what an expression in function code can name: the
// parameters and locals in scope at off, then the Globals, the
// objectives (read as `name[@s]`) and the functions, then true and false.
func (s *Session) codeValues(a *at, off int) []sema.Candidate {
	out := locals(a.u.File, a.decl.D, off)
	out = append(out, s.projectNames(a)...)
	out = append(out,
		sema.Candidate{Label: "true", Kind: sema.CandKeyword, Rank: 200},
		sema.Candidate{Label: "false", Kind: sema.CandKeyword, Rank: 201})
	return out
}

// projectNames lists the Globals, objectives and functions code can use.
func (s *Session) projectNames(a *at) []sema.Candidate {
	var out []sema.Candidate
	ns := a.decl.NS
	short := func(sym *sema.Symbol) string {
		if sym.NS == ns {
			return sym.Path
		}
		return sym.ID()
	}
	for _, sym := range s.R.Syms.All() {
		doc := sema.DocMarkdown(sym.Decl.Doc)
		switch sym.Kind.MMS {
		case "Global":
			out = append(out, sema.Candidate{Label: short(sym), Kind: sema.CandID, Detail: "Global " + sema.SymSig(sym), Doc: doc, Rank: 50})
		case "Objective":
			name := short(sym)
			out = append(out, sema.Candidate{Label: name, Kind: sema.CandID, Detail: "objective", Doc: doc, Insert: name + "[${1:@s}]", Rank: 60})
		case "Function":
			name := short(sym)
			ins := name + "($1)"
			if len(sema.SigOf(sym.Decl).Params) == 0 {
				ins = name + "()"
			}
			out = append(out, sema.Candidate{Label: name, Kind: sema.CandCallable, Detail: "Function" + sema.SymSig(sym), Doc: doc, Insert: ins, Rank: 70})
		}
	}
	return out
}

// locals lists the parameters of d and the variables declared before off
// in the blocks that hold it, innermost last.
func locals(f *syntax.File, d *syntax.Decl, off int) []sema.Candidate {
	var out []sema.Candidate
	if d.Params != nil {
		for _, p := range d.Params.Params {
			if p.Name != nil {
				out = append(out, sema.Candidate{Label: p.Name.Text, Kind: sema.CandField, Detail: "parameter" + typeSuffix(p.Type), Rank: 0})
			}
		}
	}
	if d.Body == nil {
		return out
	}
	syntax.Walk(d.Body, func(n syntax.Node) bool {
		if !n.Span().Contains(off) {
			return false
		}
		switch n := n.(type) {
		case *syntax.Block:
			for _, it := range n.Items {
				// A let is in scope after its value, so not in `let y = |`
				// (recovery can end it early: `let y = g + |` reads `let y = g`).
				if l, ok := it.(*syntax.LetStmt); ok && l.Name != nil && l.Value != nil && badAfter(l, l.Sp.Start, l.Sp.End) == nil && l.Sp.End < off && lineOf(f, l.Sp.Start) != lineOf(f, off) {
					out = append(out, sema.Candidate{Label: l.Name.Text, Kind: sema.CandField, Detail: "let" + typeSuffix(l.Type), Doc: sema.DocMarkdown(l.Doc), Rank: 0})
				}
			}
		case *syntax.ForStmt:
			if n.Var != nil && n.Body != nil && n.Body.Sp.Contains(off) {
				out = append(out, sema.Candidate{Label: n.Var.Text, Kind: sema.CandField, Detail: "loop variable", Rank: 0})
			}
		}
		return true
	})
	// A later declaration of the same name shadows an earlier one.
	seen := map[string]bool{}
	for i := len(out) - 1; i >= 0; i-- {
		if seen[out[i].Label] {
			out = append(out[:i], out[i+1:]...)
			continue
		}
		seen[out[i].Label] = true
	}
	return out
}

func typeSuffix(t *syntax.Name) string {
	if t == nil {
		return ""
	}
	return ": " + t.Text
}

// typeCandidates are the value types of Globals, parameters and locals.
func typeCandidates() []sema.Candidate {
	var out []sema.Candidate
	for i, t := range []string{"int", "bool", "string"} {
		out = append(out, sema.Candidate{Label: t, Kind: sema.CandKeyword, Rank: i})
	}
	return out
}

// criteria lists the scoreboard criteria an Objective can count.
func criteria() []sema.Candidate {
	var out []sema.Candidate
	for i, c := range command.Criteria {
		out = append(out, sema.Candidate{Label: c, Kind: sema.CandEnum, Rank: i})
	}
	return out
}
