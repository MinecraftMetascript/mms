package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Execute blocks: `As(@e[tag=x]).At(@s) { ... }` in a function body runs
// the block under `execute as @e[tag=x] at @s run ...`. Each helper is
// one execute subcommand; its arguments are written into the command as
// they appear in the source, and the whole chain is checked against the
// command tree.

// ExecHelper is one execute block helper.
type ExecHelper struct {
	Name   string // "As"
	Sub    string // "as"
	Params string // "targets", for the signature
	Doc    string
	Insert string // completion snippet
	// render builds the subcommand from the arguments' text; nil means
	// Sub followed by the arguments.
	render func(args []execArg) string
	min    int // fewest arguments
}

type execArg struct {
	text     string
	selector bool
}

func joinArgs(args []execArg) string {
	ts := make([]string, len(args))
	for i, a := range args {
		ts[i] = a.text
	}
	return strings.Join(ts, " ")
}

// orAs writes `<sub> as <selector>` when the first argument is a selector,
// as positioned and rotated do.
func orAs(sub string) func([]execArg) string {
	return func(args []execArg) string {
		if args[0].selector {
			return sub + " as " + joinArgs(args)
		}
		return sub + " " + joinArgs(args)
	}
}

// ExecHelpers lists the execute block helpers, in completion order.
var ExecHelpers = []*ExecHelper{
	{Name: "As", Sub: "as", Params: "targets", min: 1, Insert: "As(${1:@a}) {\n\t$0\n}",
		Doc: "Runs the block once for each target, as that entity (`execute as`). Position and rotation stay as they were; chain `.At(@s)` to move there too."},
	{Name: "At", Sub: "at", Params: "targets", min: 1, Insert: "At(${1:@s}) {\n\t$0\n}",
		Doc: "Runs the block at each target's position, rotation and dimension (`execute at`)."},
	{Name: "Positioned", Sub: "positioned", Params: "pos | targets", min: 1, render: orAs("positioned"), Insert: "Positioned(${1:~ ~ ~}) {\n\t$0\n}",
		Doc: "Moves the position (`execute positioned`). Coordinates are written `~ ~1 ~` or `0, 64, 0`; a selector uses `positioned as`."},
	{Name: "Rotated", Sub: "rotated", Params: "rot | targets", min: 1, render: orAs("rotated"), Insert: "Rotated(${1:~ ~}) {\n\t$0\n}",
		Doc: "Sets the rotation (`execute rotated`): `~ ~`, `90, 0`, or a selector for `rotated as`."},
	{Name: "Facing", Sub: "facing", Params: "pos | targets, anchor?", min: 1, Insert: "Facing(${1:@p}) {\n\t$0\n}",
		Doc: "Turns to face a position or an entity (`execute facing`). With a selector the second argument is the anchor, `eyes` (the default) or `feet`.",
		render: func(args []execArg) string {
			if args[0].selector {
				anchor := "eyes"
				if len(args) > 1 {
					anchor = joinArgs(args[1:])
				}
				return "facing entity " + args[0].text + " " + anchor
			}
			return "facing " + joinArgs(args)
		}},
	{Name: "In", Sub: "in", Params: "dimension", min: 1, Insert: "In(${1:the_nether}) {\n\t$0\n}",
		Doc: "Runs the block in another dimension (`execute in`), at the same coordinates."},
	{Name: "Anchored", Sub: "anchored", Params: "eyes | feet", min: 1, Insert: "Anchored(${1:eyes}) {\n\t$0\n}",
		Doc: "Sets the anchor that local coordinates (`^ ^ ^`) and facing use (`execute anchored`)."},
	{Name: "Align", Sub: "align", Params: "axes", min: 1, Insert: "Align(${1:xyz}) {\n\t$0\n}",
		Doc: "Rounds the position down on the given axes, such as `xz` (`execute align`)."},
	{Name: "On", Sub: "on", Params: "relation", min: 1, Insert: "On(${1:passengers}) {\n\t$0\n}",
		Doc: "Runs the block as related entities: `attacker`, `controller`, `leasher`, `origin`, `owner`, `passengers`, `target` or `vehicle` (`execute on`)."},
	{Name: "Summon", Sub: "summon", Params: "entity", min: 1, Insert: "Summon(${1:marker}) {\n\t$0\n}",
		Doc: "Summons an entity and runs the block as it (`execute summon`)."},
	{Name: "Execute", Sub: "", Params: "\"subcommands\"", min: 1, Insert: "Execute(\"$1\") {\n\t$0\n}",
		Doc:    "Any execute subcommands, written as text: `Execute(\"if entity @e[tag=boss] positioned ~ ~1 ~\")`. Use it for the conditions and subcommands that have no helper.",
		render: joinArgs},
}

var execByName = func() map[string]*ExecHelper {
	m := map[string]*ExecHelper{}
	for _, h := range ExecHelpers {
		m[h.Name] = h
	}
	return m
}()

// Signature is the helper as it is written, for hover.
func (h *ExecHelper) Signature() string { return h.Name + "(" + h.Params + ") { ... }" }

func (h *ExecHelper) help() string {
	return "```mms\n" + h.Signature() + "\n```\n" + h.Doc
}

// ExecCandidates are the execute helpers, for completion where a
// statement starts or after `.` in a chain.
func ExecCandidates(block bool) []Candidate {
	var out []Candidate
	for i, h := range ExecHelpers {
		c := Candidate{Label: h.Name, Kind: CandCallable, Detail: "execute " + strings.TrimSpace(h.Sub), Doc: h.help(), Insert: h.Insert, Rank: 10 + i}
		if !block {
			c.Insert = h.Name + "($1)"
		}
		out = append(out, c)
	}
	return out
}

// execStep is one helper call of a chain.
type execStep struct {
	name  *syntax.Ident // nil for a chained step, which names with nameN
	nameN *syntax.Name
	call  *syntax.Call
}

func (s execStep) text() string {
	if s.name != nil {
		return s.name.Name
	}
	return s.nameN.Text
}

func (s execStep) node() syntax.Node {
	if s.name != nil {
		return s.name
	}
	return s.nameN
}

func (s execStep) nameSpan() syntax.Span { return s.node().Span() }

// execChain splits `As(x).At(y)` into its steps; ok is false when e is
// not a chain of calls.
func execChain(e syntax.Expr) (steps []execStep, ok bool) {
	call, isCall := e.(*syntax.Call)
	if !isCall {
		return nil, false
	}
	switch fx := call.Fn.(type) {
	case *syntax.Ident:
		return []execStep{{name: fx, call: call}}, true
	case *syntax.Member:
		if fx.Name == nil {
			return nil, false
		}
		head, ok := execChain(fx.X)
		if !ok {
			return nil, false
		}
		return append(head, execStep{nameN: fx.Name, call: call}), true
	}
	return nil, false
}

// IsExecChain reports whether e starts with an execute helper, such as
// `As(@a)` or `As(@a).At(@s).` while the next step is being typed.
func IsExecChain(e syntax.Expr) bool {
	for {
		switch x := e.(type) {
		case *syntax.Member:
			e = x.X
		case *syntax.Call:
			if id, ok := x.Fn.(*syntax.Ident); ok {
				return execByName[id.Name] != nil
			}
			e = x.Fn
		case *syntax.Trailing:
			e = x.X
		default:
			return false
		}
	}
}

// execHeadName returns the first helper name of a trailing-block
// statement, for the error when it is not an execute chain.
func execHeadName(e syntax.Expr) (string, syntax.Span) {
	for {
		switch x := e.(type) {
		case *syntax.Member:
			e = x.X
		case *syntax.Call:
			e = x.Fn
		case *syntax.Ident:
			return x.Name, x.Sp
		default:
			return "", e.Span()
		}
	}
}

func execNames() []string {
	names := make([]string, len(ExecHelpers))
	for i, h := range ExecHelpers {
		names[i] = h.Name
	}
	return names
}

// piece maps a part of the generated execute text back to the source.
type piece struct {
	start, end int // runes in the text
	src        syntax.Span
	verbatim   bool // the text is the source as written
}

// execParts checks an execute chain and returns its subcommands, or ok
// false after reporting a problem.
func (f *fc) execParts(e syntax.Expr) (parts []string, ok bool) {
	c := f.c
	steps, isChain := execChain(e)
	if !isChain {
		name, sp := execHeadName(e)
		f.unknownExec(name, sp)
		return nil, false
	}
	text := "execute"
	var pieces []piece
	ok = true
	for _, st := range steps {
		h := execByName[st.text()]
		if h == nil {
			f.unknownExec(st.text(), st.nameSpan())
			ok = false
			continue
		}
		c.info(st.node()).Help = h.help()
		var args []execArg
		var argPieces []piece
		for _, a := range st.call.Args {
			if a.Name != nil {
				c.errorf("mms/unsupported", a.Sp, "%s takes its arguments by position", h.Name)
				ok = false
				continue
			}
			if a.Value == nil {
				continue
			}
			var ea execArg
			verbatim := true
			switch v := a.Value.(type) {
			case *syntax.String:
				ea.text, verbatim = v.Value, false
				if strings.Contains(v.Value, "${") {
					c.errorf("mms/unsupported", v.Sp, "execute helpers take constant text; ${...} is not supported here")
					ok = false
				}
			case *syntax.Bad:
				ok = false
				continue
			default:
				sp := a.Value.Span()
				ea.text = string(f.src[sp.Start:sp.End])
				switch x := a.Value.(type) {
				case *syntax.Selector:
					ea.selector = true
				case *syntax.Index:
					_, ea.selector = x.X.(*syntax.Selector)
				}
			}
			args = append(args, ea)
			argPieces = append(argPieces, piece{src: a.Value.Span(), verbatim: verbatim})
		}
		if len(args) < h.min {
			c.errorf("mms/missing-value", st.call.ArgsSp, "%s needs %s", h.Name, h.Params)
			ok = false
			continue
		}
		var part string
		if h.render != nil {
			part = h.render(args)
		} else {
			part = h.Sub + " " + joinArgs(args)
		}
		part = strings.TrimSpace(part)
		if part == "" {
			c.errorf("mms/missing-value", st.call.ArgsSp, "%s needs %s", h.Name, h.Params)
			ok = false
			continue
		}
		base := len([]rune(text)) + 1
		// Find each argument's text in the part, in order, to map
		// command errors back onto it.
		at := 0
		pr := []rune(part)
		for i, ap := range argPieces {
			at = runeIndex(pr, []rune(args[i].text), at)
			if at < 0 {
				break
			}
			n := len([]rune(args[i].text))
			ap.start, ap.end = base+at, base+at+n
			pieces = append(pieces, ap)
			at += n
		}
		pieces = append(pieces, piece{start: base, end: base + len(pr), src: st.call.Sp})
		text += " " + part
		parts = append(parts, part)
	}
	if !ok || len(parts) == 0 {
		return nil, false
	}
	whole := e.Span()
	errs := len(c.Diags)
	c.commandText(text+" run say mms", func(start, end int) syntax.Span {
		for _, p := range pieces {
			if start >= p.start && start < p.end || start == p.end && start == end && p.verbatim {
				if p.verbatim {
					s := p.src.Start + start - p.start
					e := p.src.Start + end - p.start
					if e > p.src.End {
						e = p.src.End
					}
					if e <= s {
						e = s + 1
					}
					return syntax.Span{Start: s, End: e}
				}
				return p.src
			}
		}
		return whole
	})
	if hasErrorsSince(c.Diags, errs) {
		return nil, false
	}
	return parts, true
}

func hasErrorsSince(ds []syntax.Diagnostic, from int) bool {
	for _, d := range ds[from:] {
		if d.Severity == syntax.SeverityError {
			return true
		}
	}
	return false
}

func runeIndex(s, sub []rune, from int) int {
	for i := from; i+len(sub) <= len(s); i++ {
		if string(s[i:i+len(sub)]) == string(sub) {
			return i
		}
	}
	return -1
}

func (f *fc) unknownExec(name string, sp syntax.Span) {
	names := execNames()
	msg := "a block here is an execute block, and '" + name + "' is not an execute helper"
	if name == "" {
		msg = "a block here is an execute block, such as As(@a) { ... }"
	} else if s := closest(name, names); s != "" {
		msg += "; did you mean '" + s + "'?"
	}
	sorted := append([]string(nil), names...)
	sort.Strings(sorted)
	f.c.errorf("mms/unknown-helper", sp, "%s (helpers: %s)", msg, strings.Join(sorted, ", "))
}

// execBlock lowers `As(@a).At(@s) { ... }`.
func (f *fc) execBlock(out *fn.Block, t *syntax.Trailing) {
	parts, ok := f.execParts(t.X)
	body := &fn.Block{}
	loops := f.loops
	f.loops = 0
	f.exec++
	f.block(body, t.Block)
	f.exec--
	f.loops = loops
	if ok {
		out.Add(fn.Exec{Parts: parts, Body: body})
	}
}

// Selectors are the target selectors, for completion.
var Selectors = []string{"@a", "@e", "@n", "@p", "@r", "@s"}

func words(kind CandidateKind, ws ...string) []Candidate {
	out := make([]Candidate, len(ws))
	for i, w := range ws {
		out[i] = Candidate{Label: w, Kind: kind, Rank: i}
	}
	return out
}

// ExecArgCandidates lists values for argument index of the execute helper
// name, such as the selectors of As(|), or nil when it has no list.
func (c *Checker) ExecArgCandidates(name string, index int) []Candidate {
	sel := words(CandKeyword, Selectors...)
	switch name {
	case "As", "At":
		return sel
	case "Positioned":
		return append(words(CandKeyword, "~ ~ ~", "^ ^ ^"), sel...)
	case "Rotated":
		return append(words(CandKeyword, "~ ~"), sel...)
	case "Facing":
		if index > 0 {
			return words(CandEnum, "eyes", "feet")
		}
		return append(words(CandKeyword, "~ ~ ~"), sel...)
	case "Anchored":
		return words(CandEnum, "eyes", "feet")
	case "Align":
		return words(CandEnum, "xyz", "xz", "x", "y", "z")
	case "On":
		return words(CandEnum, "attacker", "controller", "leasher", "origin", "owner", "passengers", "target", "vehicle")
	case "In":
		return dedupe(c.idCandidates(idInfo{Registry: "dimension", Present: true}))
	case "Summon":
		return dedupe(c.idCandidates(idInfo{Registry: "entity_type", Present: true}))
	}
	return nil
}

// ExecHelper looks up an execute helper by name.
func ExecHelperNamed(name string) *ExecHelper { return execByName[name] }
