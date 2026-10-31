package project

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Rule is a lint rule: a check over the analysed project that reports
// style problems and likely mistakes, never errors the checker owns.
type Rule struct {
	ID       string
	Severity syntax.Severity // default; mms.json "lint" can change it
	Doc      string
	run      func(*lintRun, *Decl) // per checked declaration
	all      func(*lintRun)        // once over the project
}

// Rules lists every lint rule, sorted by id.
var Rules = []*Rule{
	{ID: "constant-expression", Severity: syntax.SeverityHint, run: lintConstant,
		Doc: "An arithmetic expression of number literals can be written as its value."},
	{ID: "overrides-vanilla", Severity: syntax.SeverityInfo, run: lintOverrides,
		Doc: "A declaration in the minecraft namespace replaces the vanilla resource with the same id."},
	{ID: "quoted-name", Severity: syntax.SeverityHint, run: lintQuoted,
		Doc: "An id or enum name in quotes can be written bare: `minecraft:stone` or `stone` rather than \"minecraft:stone\"."},
	{ID: "redundant-namespace", Severity: syntax.SeverityHint, run: lintNamespace,
		Doc: "A `minecraft:` prefix that can be dropped: `stone` means `minecraft:stone` unless the project declares its own `stone`."},
	{ID: "use-sugar", Severity: syntax.SeverityHint, run: lintSugar,
		Doc: "An object literal that has a shorter MMS form: an item entry as its id, `1..3` for a uniform provider or bounds, `diamond * 4` for a stack, a trigger call for a criterion, and `&&`, `||` and `!` for conditions."},
	{ID: "unused", Severity: syntax.SeverityHint, all: lintUnused,
		Doc: "A worldgen piece, predicate or item modifier that nothing in the project references, so the game never uses it."},
}

// RuleByID returns a rule, or nil.
func RuleByID(id string) *Rule {
	for _, r := range Rules {
		if r.ID == id {
			return r
		}
	}
	return nil
}

// ParseSeverity reads a severity name from mms.json: "error", "warning",
// "info", "hint" or "off" (0).
func ParseSeverity(s string) (syntax.Severity, bool) {
	switch s {
	case "error":
		return syntax.SeverityError, true
	case "warning", "warn":
		return syntax.SeverityWarning, true
	case "info":
		return syntax.SeverityInfo, true
	case "hint":
		return syntax.SeverityHint, true
	case "off":
		return 0, true
	}
	return 0, false
}

// unusedFolders are the kinds the game only reaches through a reference
// from another resource or a command. Loot tables, recipes, advancements
// and the like are loaded by the game itself.
var unusedFolders = map[string]bool{
	"worldgen/density_function":                        true,
	"worldgen/noise":                                   true,
	"worldgen/configured_feature":                      true,
	"worldgen/placed_feature":                          true,
	"worldgen/configured_carver":                       true,
	"worldgen/processor_list":                          true,
	"worldgen/material_rule":                           true,
	"worldgen/material_condition":                      true,
	"predicate":                                        true,
	"item_modifier":                                    true,
	"worldgen/multi_noise_biome_source_parameter_list": true,
}

type lintRun struct {
	r    *Result
	out  map[string][]syntax.Diagnostic
	rule *Rule
	sev  syntax.Severity
	cur  *[]syntax.Diagnostic // where per-declaration findings go
}

func (l *lintRun) report(d *Decl, sp syntax.Span, msg string, fixes ...syntax.Fix) {
	*l.cur = append(*l.cur, syntax.Diagnostic{
		Code: "lint/" + l.rule.ID, Severity: l.sev, Message: msg, Sp: sp, Fixes: fixes,
	})
}

func (r *Result) severity(rule *Rule) syntax.Severity {
	if s, ok := r.cfg.Lint[rule.ID]; ok {
		if v, ok := ParseSeverity(s); ok {
			return v
		}
	}
	return rule.Severity
}

// lint runs every enabled rule and drops suppressed findings. Findings of
// per-declaration rules are cached on the declaration: they depend only
// on its checked form, the declared names and the configuration.
func (r *Result) lint() map[string][]syntax.Diagnostic {
	l := &lintRun{r: r, out: map[string][]syntax.Diagnostic{}}
	key := fmt.Sprint(r.symKey, r.cfg.Lint)
	for _, d := range r.Order {
		// Rules read checked values; a declaration with errors has more
		// pressing problems.
		if hasErrors(d.Diags) {
			continue
		}
		if d.lintKey != key {
			d.lints = nil
			l.cur = &d.lints
			for _, rule := range Rules {
				if rule.run == nil {
					continue
				}
				if l.rule, l.sev = rule, r.severity(rule); l.sev != 0 {
					rule.run(l, d)
				}
			}
			d.lintKey = key
		}
		l.out[d.Unit.Path] = append(l.out[d.Unit.Path], d.lints...)
	}
	for _, rule := range Rules {
		if rule.all == nil {
			continue
		}
		if l.rule, l.sev = rule, r.severity(rule); l.sev != 0 {
			rule.all(l)
		}
	}
	for path, ds := range l.out {
		u := r.Unit(path)
		lineOf := lineIndex(u.File.Src)
		ig := ignores(u.File, lineOf)
		kept := ds[:0]
		for _, d := range ds {
			line := lineOf(d.Sp.Start)
			if ig.covers(d, line) {
				continue
			}
			// Copy: the cached finding keeps its own fix list.
			d.Fixes = append(d.Fixes[:len(d.Fixes):len(d.Fixes)], ignoreFix(u.File.Src, lineOf, d, line))
			kept = append(kept, d)
		}
		l.out[path] = kept
	}
	return l.out
}

// lineIndex returns a function from rune offset to 1-based line.
func lineIndex(src []rune) func(int) int {
	starts := []int{0}
	for i, r := range src {
		if r == '\n' {
			starts = append(starts, i+1)
		}
	}
	return func(off int) int {
		return sort.Search(len(starts), func(i int) bool { return starts[i] > off })
	}
}

// ignoreFix inserts `// mms-ignore <rule>` above the finding's line, at
// that line's indentation.
func ignoreFix(src []rune, lineOf func(int) int, d syntax.Diagnostic, line int) syntax.Fix {
	start := d.Sp.Start
	for start > 0 && src[start-1] != '\n' {
		start--
	}
	end := start
	for end < len(src) && (src[end] == ' ' || src[end] == '\t') {
		end++
	}
	indent := string(src[start:end])
	rule := strings.TrimPrefix(d.Code, "lint/")
	return syntax.Fix{
		Title: "Ignore " + rule + " on this line",
		Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: start, End: start}, Text: indent + "// mms-ignore " + rule + "\n"}},
	}
}

// ---- suppression comments

var ignoreRe = regexp.MustCompile(`^//\s*mms-ignore(-file)?(?:\s+([a-z0-9\-, ]+))?\s*$`)

type ignoreSet struct {
	file  map[string]bool // rule ids ("" = all) ignored in the whole file
	lines map[int]map[string]bool
}

// ignores reads `// mms-ignore rule, rule` comments, which silence lint
// findings on the same line or the next line, and `// mms-ignore-file`.
func ignores(f *syntax.File, lineOf func(int) int) ignoreSet {
	ig := ignoreSet{file: map[string]bool{}, lines: map[int]map[string]bool{}}
	for _, t := range f.Tokens {
		if !t.IsComment() || !strings.Contains(t.Text, "mms-ignore") {
			continue
		}
		m := ignoreRe.FindStringSubmatch(strings.TrimSpace(t.Text))
		if m == nil {
			continue
		}
		ids := map[string]bool{}
		for _, id := range strings.FieldsFunc(m[2], func(r rune) bool { return r == ',' || r == ' ' }) {
			ids[strings.TrimPrefix(id, "lint/")] = true
		}
		if len(ids) == 0 {
			ids[""] = true
		}
		if m[1] != "" {
			for id := range ids {
				ig.file[id] = true
			}
			continue
		}
		line := lineOf(t.Sp.Start)
		for _, l := range []int{line, line + 1} {
			if ig.lines[l] == nil {
				ig.lines[l] = map[string]bool{}
			}
			for id := range ids {
				ig.lines[l][id] = true
			}
		}
	}
	return ig
}

func (ig ignoreSet) covers(d syntax.Diagnostic, line int) bool {
	id := strings.TrimPrefix(d.Code, "lint/")
	if ig.file[""] || ig.file[id] {
		return true
	}
	set := ig.lines[line]
	return set[""] || set[id]
}

// ---- rules

func lintOverrides(l *lintRun, d *Decl) {
	if l.r.cfg.Namespace == "minecraft" || !d.Kind.IsResource() || d.D.Name == nil {
		return
	}
	id := d.ID()
	if !strings.HasPrefix(id, "minecraft:") {
		return
	}
	path := strings.TrimPrefix(id, "minecraft:")
	folder := d.Kind.Folder()
	vanilla := false
	if d.Kind.TagOf != "" {
		vanilla = contains(l.r.pack.Registries["tag/"+d.Kind.TagOf], path)
	} else {
		vanilla = l.r.pack.HasID(folder, path)
	}
	if vanilla {
		l.report(d, d.D.Name.Sp, fmt.Sprintf("this replaces vanilla's %s '%s'", folder, id))
	}
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func lintQuoted(l *lintRun, d *Decl) {
	syntax.Walk(d.D, func(n syntax.Node) bool {
		s, ok := n.(*syntax.String)
		if !ok {
			return true
		}
		bare, ok := sema.BareString(d.Info[s], l.r.Syms, d.NS, s)
		if !ok {
			return true
		}
		l.report(d, s.Sp, "quotes are not needed here",
			syntax.Fix{Title: "Write " + bare, Edits: []syntax.TextEdit{{Sp: s.Sp, Text: bare}}, Safe: true})
		return true
	})
}

func lintNamespace(l *lintRun, d *Decl) {
	syntax.Walk(d.D, func(n syntax.Node) bool {
		var name string
		sp := n.Span()
		switch x := n.(type) {
		case *syntax.ResLoc:
			name = x.Text
		case *syntax.Tag:
			name = x.Name
			sp.Start++ // keep the '#'
		default:
			return true
		}
		short, ok := sema.ShortName(d.Info[n], l.r.Syms, d.NS, name)
		if !ok {
			return true
		}
		l.report(d, sp, "'minecraft:' is the default namespace here",
			syntax.Fix{Title: "Write " + short, Edits: []syntax.TextEdit{{Sp: sp, Text: short}}, Safe: true})
		return true
	})
}

func lintSugar(l *lintRun, d *Decl) {
	src := d.Unit.File.Src
	// Objects that are operands of an operator need parentheses when they
	// become an operator expression themselves.
	type slot struct {
		op    string
		right bool
	}
	operands := map[syntax.Node]slot{}
	syntax.Walk(d.D, func(n syntax.Node) bool {
		switch x := n.(type) {
		case *syntax.Unary:
			operands[x.X] = slot{op: "unary"}
		case *syntax.Binary:
			operands[x.L], operands[x.R] = slot{op: x.Op}, slot{op: x.Op, right: true}
		}
		return true
	})
	syntax.Walk(d.D, func(n syntax.Node) bool {
		o, ok := n.(*syntax.Object)
		if !ok {
			return true
		}
		short, op, ok := sema.Simplify(d.Info[o], o, src)
		if !ok {
			short, ok = sema.AsCall(d.Info[o], d.Info[o.Body], o, src)
		}
		if !ok {
			return true
		}
		if at, isOperand := operands[o]; isOperand && sema.NeedsParens(at.op, at.right, op) {
			short = "(" + short + ")"
		}
		l.report(d, o.Sp, "this can be written "+clip(short),
			syntax.Fix{Title: "Write " + clip(short), Edits: []syntax.TextEdit{{Sp: o.Sp, Text: short}}, Safe: true})
		return true
	})
}

func clip(s string) string {
	if r := []rune(s); len(r) > 40 {
		return string(r[:37]) + "..."
	}
	return s
}

// lintConstant reports the largest arithmetic subtrees whose leaves are
// all number literals.
func lintConstant(l *lintRun, d *Decl) {
	syntax.Walk(d.D, func(n syntax.Node) bool {
		e, ok := n.(syntax.Expr)
		if !ok || !hasBinary(e) {
			return true
		}
		v, ok := constValue(e)
		if !ok {
			return true
		}
		text := strconv.FormatFloat(v, 'g', -1, 64)
		if math.IsInf(v, 0) || math.IsNaN(v) || len(strings.TrimLeft(text, "-")) > 12 || strings.Contains(text, "e") {
			return false // keep expressions whose value prints badly
		}
		l.report(d, e.Span(), "this expression is the constant "+text,
			syntax.Fix{Title: "Replace with " + text, Edits: []syntax.TextEdit{{Sp: e.Span(), Text: text}}, Safe: true})
		return false
	})
}

func hasBinary(e syntax.Expr) bool {
	switch x := e.(type) {
	case *syntax.Binary:
		return true
	case *syntax.Paren:
		return hasBinary(x.X)
	case *syntax.Unary:
		return hasBinary(x.X)
	}
	return false
}

func constValue(e syntax.Expr) (float64, bool) {
	switch x := e.(type) {
	case *syntax.Number:
		v, err := strconv.ParseFloat(x.Text, 64)
		return v, err == nil
	case *syntax.Paren:
		return constValue(x.X)
	case *syntax.Unary:
		v, ok := constValue(x.X)
		switch x.Op {
		case "-":
			return -v, ok
		case "+":
			return v, ok
		}
	case *syntax.Binary:
		a, ok1 := constValue(x.L)
		b, ok2 := constValue(x.R)
		if !ok1 || !ok2 {
			return 0, false
		}
		switch x.Op {
		case "+":
			return a + b, true
		case "-":
			return a - b, true
		case "*":
			return a * b, true
		case "/":
			return a / b, b != 0
		}
	}
	return 0, false
}

var cmdIDRe = regexp.MustCompile(`#?[a-z0-9_.-]+:[a-z0-9_./-]+`)

func lintUnused(l *lintRun) {
	r := l.r
	used := map[string]bool{} // folder|id, and ids named in commands
	for _, d := range r.Order {
		for _, ref := range d.references() {
			used[ref] = true
		}
	}
	var unused []*sema.Symbol
	for _, s := range r.Syms.All() {
		if s.NS == "minecraft" || !unusedFolders[s.Kind.Folder()] || used[s.Kind.Folder()+"|"+s.ID()] || used[s.ID()] || s.Decl.Name == nil {
			continue
		}
		unused = append(unused, s)
	}
	sort.Slice(unused, func(i, j int) bool { return unused[i].Decl.Sp.Start < unused[j].Decl.Sp.Start })
	for _, s := range unused {
		if d := r.Decls[s.Decl]; d == nil || hasErrors(d.Diags) {
			continue
		}
		l.out[s.Unit.Path] = append(l.out[s.Unit.Path], syntax.Diagnostic{
			Code: "lint/unused", Severity: l.sev, Sp: s.Decl.Name.Sp, Unnecessary: true,
			Message: fmt.Sprintf("%s '%s' is never used", s.Kind.Name, s.ID()),
		})
	}
}

// references lists what a declaration refers to, as "folder|id" for
// resolved references and bare ids for ids named in its commands. It is
// computed once per checked declaration.
func (d *Decl) references() []string {
	if d.refsDone {
		return d.refs
	}
	seen := map[string]bool{}
	for _, in := range d.Info {
		if s := in.Symbol; s != nil && s.Decl != d.D { // self-references do not count
			seen[s.Kind.Folder()+"|"+s.ID()] = true
		}
	}
	scan := func(ls []string) {
		for _, line := range ls {
			for _, id := range cmdIDRe.FindAllString(line, -1) {
				seen[id] = true
			}
		}
	}
	scan(d.Lines)
	if d.Fn != nil {
		for _, ls := range d.Fn.Files {
			scan(ls)
		}
	}
	d.refs = make([]string, 0, len(seen))
	for k := range seen {
		d.refs = append(d.refs, k)
	}
	d.refsDone = true
	return d.refs
}
