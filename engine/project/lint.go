package project

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spline"
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
	{ID: "book-page-overflow", Severity: syntax.SeverityWarning, run: lintBookOverflow,
		Doc: "A Book page with more text than the book screen shows: 14 lines of about 19 characters. The game cuts the rest off. Widths are measured with the vanilla font, so text outside ASCII is a close guess."},
	{ID: "constant-expression", Severity: syntax.SeverityHint, run: lintConstant,
		Doc: "An arithmetic expression of number literals can be written as its value."},
	{ID: "deprecated", Severity: syntax.SeverityWarning, all: lintDeprecated,
		Doc: "A use of a declaration whose doc comment has a `@deprecated` tag."},
	{ID: "external-asset", Severity: syntax.SeverityWarning, all: lintExternalAsset,
		Doc: "An id of a texture, sound or structure in the project's namespace that no Texture, Sound or StructureFile call in the project provides. The pack still loads, but the file must come from another pack."},
	{ID: "jigsaw-facing", Severity: syntax.SeverityWarning, all: lintJigsawFacing,
		Doc: "A jigsaw block that can never join its target: every jigsaw block with its target name in its pool faces a way it cannot meet. The game turns pieces only about the y axis, so a block facing up joins only one facing down, and a sideways block only a sideways one."},
	{ID: "jigsaw-fit", Severity: syntax.SeverityWarning, all: lintJigsawFit,
		Doc: "An internal jigsaw block (the block in front of it is inside its own piece, such as an arch mount on a road, or a jigsaw inset by a block by mistake) whose pool has pieces that stick out of the piece however they are turned. The game only places a piece joined at an internal jigsaw wholly inside its parent, so those pieces are never placed there. Height is not checked for terrain_matching pieces."},
	{ID: "jigsaw-unknown-name", Severity: syntax.SeverityWarning, all: lintJigsawUnknownName,
		Doc: "A jigsaw block in a structure file that StructureFile(\"...\") includes whose name or target is in a namespace with JigsawHole declarations, but is not one of them: most likely a typo made in the game's jigsaw screen. The quick fix declares it. Projects that declare no JigsawHole in a namespace are not checked there."},
	{ID: "jigsaw-overlap", Severity: syntax.SeverityInfo, all: lintJigsawOverlap,
		Doc: "A jigsaw block that can never grow when its piece joins a given parent: every piece its pool and the pool's fallback could place there overlaps the parent or the piece itself, however it is turned, and the game never places overlapping pieces. A junction narrower than the roads around it is the usual cause. Only certain cases are reported: every piece involved is a rigid project structure file with a known size and facing, and no pool involved has an empty or feature element."},
	{ID: "jigsaw-unknown-pool", Severity: syntax.SeverityWarning, all: lintJigsawPool,
		Doc: "A jigsaw block in a structure file that StructureFile(\"...\") includes draws from a pool in the project's namespace that no TemplatePool declares. The game skips the jigsaw, so the structure stops growing there. The quick fix declares the pool."},
	{ID: "jigsaw-unreachable-target", Severity: syntax.SeverityWarning, all: lintJigsawTarget,
		Doc: "A jigsaw block whose target no piece of its pool can join: none of the pool's structure files has a jigsaw block with that name, or the pool has no pieces. Only pools whose pieces are all the project's own structure files are checked."},
	{ID: "model-mixed-atlas", Severity: syntax.SeverityWarning, all: lintMixedAtlas,
		Doc: "A model with one texture that only the blocks atlas stitches and another that only the items atlas stitches. The game draws a model's faces from a single atlas, so keep its textures under one of textures/block/ and textures/item/. A Texture(...) in a Model goes under textures/block/ unless it is given an id name."},
	{ID: "model-missing-texture", Severity: syntax.SeverityWarning, all: lintModelTexture,
		Doc: "A model whose faces show a texture slot, such as `#side`, that nothing in its parent chain fills, so the game draws the missing texture there. The faces are those of the nearest model in the chain with elements: `block/cube_column` uses `#end` and `#side`. Chains that reach a model neither the project nor vanilla has are not checked."},
	{ID: "overrides-vanilla", Severity: syntax.SeverityInfo, run: lintOverrides,
		Doc: "A declaration in the minecraft namespace replaces the vanilla resource with the same id."},
	{ID: "accidental-markup", Severity: syntax.SeverityHint, run: lintAccidental,
		Doc: "A style mark in text that is never closed, such as the `*` of \"5 * 3\", so it is written as the character itself. The fix puts a backslash before it, which says so."},
	{ID: "quoted-name", Severity: syntax.SeverityHint, run: lintQuoted,
		Doc: "An id or enum name in quotes can be written bare: `minecraft:stone` or `stone` rather than \"minecraft:stone\"."},
	{ID: "redundant-namespace", Severity: syntax.SeverityHint, run: lintNamespace,
		Doc: "A `minecraft:` prefix that can be dropped: `stone` means `minecraft:stone` unless the project declares its own `stone`."},
	{ID: "spline-order", Severity: syntax.SeverityWarning, run: lintSplineOrder,
		Doc: "A spline point whose location is lower than the point before it. The game looks points up by location assuming they go from lowest to highest, so points out of order give a broken curve. The quick fix sorts the points."},
	{ID: "structure-unknown-loot-table", Severity: syntax.SeverityWarning, all: lintStructureLoot,
		Doc: "A chest, barrel, chest minecart or other container in a structure file that StructureFile(\"...\") includes whose LootTable tag names a loot table nothing provides: in the minecraft namespace one vanilla does not have, and in a namespace the project declares in one that no LootTable declares. The game leaves the container empty without an error. The quick fix declares the table. Other namespaces may come from another pack and are not checked."},
	{ID: "texture-outside-atlas", Severity: syntax.SeverityWarning, all: lintTextureAtlas,
		Doc: "A texture a model names, outside the minecraft namespace, that neither the blocks atlas nor the items atlas stitches, so the game shows the missing texture. Vanilla's atlases take the files under textures/block/ and textures/item/; a Texture(...) in a Model goes under textures/block/ unless it is given an id name. The project's own `Atlas blocks` and `Atlas items` declarations in the minecraft namespace count too."},
	{ID: "use-shorthand", Severity: syntax.SeverityHint, run: lintShorthand,
		Doc: "An object literal that has a shorter MMS form: an item entry as its id, `1..3` for a uniform provider or bounds, `diamond * 4` for a stack, a trigger call for a criterion, and `&&`, `||` and `!` for conditions."},
	{ID: "unused", Severity: syntax.SeverityHint, all: lintUnused,
		Doc: "A worldgen piece, predicate or item modifier that nothing in the project references, so the game never uses it, or a const, Template, Type or Blueprint that nothing names. A TemplatePool counts as used when a jigsaw block in one of the project's structure files draws from it."},
	{ID: "unused-asset", Severity: syntax.SeverityInfo,
		Doc: "A .png, .ogg or .nbt file in the project that no Texture, Sound or StructureFile call and no `icon` uses, so the packs leave it out. It is reported on the file itself (in VS Code's Problems panel, by `mms check`, and in the web editor's asset list), and not while the project has errors."},
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

// SeverityName is the mms.json name of a severity, the reverse of
// ParseSeverity.
func SeverityName(s syntax.Severity) string {
	switch s {
	case syntax.SeverityError:
		return "error"
	case syntax.SeverityWarning:
		return "warning"
	case syntax.SeverityInfo:
		return "info"
	case syntax.SeverityHint:
		return "hint"
	}
	return "off"
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
	"worldgen/template_pool":                           true,
}

type lintRun struct {
	r    *Result
	out  map[string][]syntax.Diagnostic
	rule *Rule
	sev  syntax.Severity
	cur  *[]syntax.Diagnostic           // where per-declaration findings go
	ext  map[string][]syntax.Diagnostic // and those in other files
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
			d.lints, d.extLints = nil, map[string][]syntax.Diagnostic{}
			l.cur, l.ext = &d.lints, d.extLints
			// A Blueprint's output is checked, but its style is the
			// Blueprint's: the user wrote only the instance. Hints on
			// conversions still apply.
			emitted := r.emittedDecl(d.D)
			for _, rule := range Rules {
				if rule.run == nil || emitted && !isHintRule(rule) {
					continue
				}
				if l.rule, l.sev = rule, r.severity(rule); l.sev != 0 {
					rule.run(l, d)
				}
			}
			lintHints(l, d, "")
			d.lintKey = key
		}
		l.out[d.Unit.Path] = append(l.out[d.Unit.Path], d.lints...)
		for path, ds := range d.extLints {
			if r.Unit(path) != nil {
				l.out[path] = append(l.out[path], ds...)
			}
		}
	}
	// A Book's text is checked on its instance, which is not in Order.
	for _, d := range r.Instances {
		if d.Item == nil || d.Item.Book == nil {
			continue
		}
		var found []syntax.Diagnostic
		ext := map[string][]syntax.Diagnostic{}
		l.cur, l.ext = &found, ext
		if l.rule = ruleByID("book-page-overflow"); l.rule != nil {
			if l.sev = r.severity(l.rule); l.sev != 0 {
				l.rule.run(l, d)
			}
		}
		l.out[d.Unit.Path] = append(l.out[d.Unit.Path], found...)
		for path, ds := range ext {
			if r.Unit(path) != nil {
				l.out[path] = append(l.out[path], ds...)
			}
		}
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
	// A block state only does anything when it replaces vanilla's.
	if l.r.cfg.Namespace == "minecraft" || !d.Kind.IsResource() || d.D.Name == nil || d.Kind.Resource == "block_definition" {
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
		// A tag without replace adds to vanilla's entries rather than
		// replacing them.
		if !d.TagReplace {
			return
		}
		vanilla = contains(l.r.pack.Registries["tag/"+d.Kind.TagOf], path)
	} else {
		vanilla = l.r.pack.HasID(folder, path)
	}
	if vanilla {
		what := folder
		if d.Kind.Assets != "" {
			what = strings.ReplaceAll(sema.Snake(d.Kind.Name), "_", " ")
		}
		l.report(d, d.D.Name.Sp, fmt.Sprintf("this replaces vanilla's %s '%s'", what, id))
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

func lintAccidental(l *lintRun, d *Decl) {
	syntax.Walk(d.D, func(n syntax.Node) bool {
		in := d.Info[n]
		if in == nil {
			return true
		}
		seen := map[syntax.Span]bool{}
		for _, sp := range in.Accidental {
			if seen[sp] {
				continue
			}
			seen[sp] = true
			mark := string(d.Unit.File.Src[sp.Start:sp.End])
			if !styleMarks[mark] {
				// The text was rewritten (a const put in for ${name}), so
				// the mark has no place in the source to fix.
				l.report(d, sp, "a style mark in this text is never closed, so it is the character itself; write a backslash before it to say so")
				continue
			}
			l.report(d, sp, "this "+mark+" is never closed, so it is the character itself; write \\"+mark+" to say so",
				syntax.Fix{Title: "Write \\" + mark, Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: sp.Start, End: sp.Start}, Text: "\\"}}, Safe: true})
		}
		return true
	})
}

// styleMarks are the marks of text markup that open and close a style.
var styleMarks = map[string]bool{"*": true, "**": true, "__": true, "~~": true, "||": true}

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

func lintShorthand(l *lintRun, d *Decl) {
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
		short, op, ok := sema.Simplify(l.r.pack, d.Info[o], o, src)
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
	// A body, `Kind name { ... }`, is an object too: it can become
	// `Kind name = Short(...)`.
	if b := d.D.Body; b != nil && d.D.Value == nil && d.D.Name != nil && d.Kind.IsResource() {
		o := &syntax.Object{Body: b, Sp: b.Sp}
		short, _, ok := sema.Simplify(l.r.pack, d.Info[d.D], o, src)
		if !ok {
			short, ok = sema.AsCall(d.Info[d.D], d.Info[b], o, src)
		}
		if ok {
			sp := syntax.Span{Start: head(d.D), End: b.Sp.End}
			l.report(d, b.Sp, "this can be written = "+clip(short),
				syntax.Fix{Title: "Write = " + clip(short), Edits: []syntax.TextEdit{{Sp: sp, Text: " = " + short}}, Safe: true})
		}
	}
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
	// A jigsaw block in a structure file draws from a pool by id.
	for _, s := range l.jigsaws().uses {
		for _, j := range s.jigsaws {
			used["worldgen/template_pool|"+j.Pool] = true
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
	unusedEvalNames(l)
	// A Blueprint instance's primary declaration is named by the
	// instance's kind, as the user wrote it: `TierDensityFunction upper`.
	kindOf := map[*syntax.Decl]string{}
	if r.Eval != nil {
		for _, f := range r.Eval.Files {
			for _, od := range f.Decls {
				if in := od.Instance; in != nil && in.Primary == od.D && in.D != nil && in.D.Kind != nil {
					kindOf[od.D] = in.D.Kind.Name
				}
			}
		}
	}
	for _, s := range unused {
		if d := r.Decls[s.Decl]; d == nil || hasErrors(d.Diags) {
			continue
		}
		kind := s.Kind.Name
		if k, ok := kindOf[s.Decl]; ok {
			kind = k
		}
		l.out[s.Unit.Path] = append(l.out[s.Unit.Path], syntax.Diagnostic{
			Code: "lint/unused", Severity: l.sev, Sp: s.Decl.Name.Sp, Unnecessary: true,
			Message: fmt.Sprintf("%s '%s' is never used", kind, s.ID()),
		})
	}
}

// unusedEvalNames reports the project's consts, Templates, Types and
// Blueprints that nothing names. They expand away, so their uses are
// read from the source as written: a name used as a value or a call, a
// type, a declaration's kind, or a word inside ${...}.
func unusedEvalNames(l *lintRun) {
	r := l.r
	if r.Eval == nil {
		return
	}
	// named holds the names each file writes.
	named := map[string]map[string]bool{}
	for path, f := range r.Source {
		in := map[string]bool{}
		named[path] = in
		syntax.Walk(f, func(n syntax.Node) bool {
			switch n := n.(type) {
			case *syntax.Ident:
				in[n.Name] = true
			case *syntax.TypeName:
				in[n.Name] = true
			case *syntax.String, *syntax.CommandStmt:
				sp := n.Span()
				if sp.End <= len(f.Src) {
					for _, piece := range evalPieceRe.FindAllString(string(f.Src[sp.Start:sp.End]), -1) {
						for _, w := range evalWordRe.FindAllString(piece, -1) {
							in[w] = true
						}
					}
				}
			}
			return true
		})
	}
	type decl struct {
		kind, path string
		name       *syntax.Name
		private    bool
		// hidden reports whether a file has a private declaration of
		// the name, which its uses there mean instead.
		hidden func(path string) bool
	}
	var ds []decl
	ev := r.Eval
	for _, c := range ev.Consts {
		ds = append(ds, decl{"const", c.Path, c.S.Name, c.S.Private, func(p string) bool { return ev.Consts[eval.PrivateKey(p, c.Name())] != nil }})
	}
	for _, t := range ev.Templates {
		ds = append(ds, decl{"Template", t.Path, t.D.Name, t.D.Private, func(p string) bool { return ev.Templates[eval.PrivateKey(p, t.D.Name.Text)] != nil }})
	}
	for _, t := range ev.Types {
		ds = append(ds, decl{"Type", t.Path, t.D.Name, t.D.Private, func(p string) bool { return ev.Types[eval.PrivateKey(p, t.D.Name.Text)] != nil }})
	}
	for _, b := range ev.Blueprints {
		ds = append(ds, decl{"Blueprint", b.Path, b.D.Name, b.D.Private, func(p string) bool { return ev.Blueprints[eval.PrivateKey(p, b.D.Name.Text)] != nil }})
	}
	used := func(d decl) bool {
		if d.private {
			return named[d.path][d.name.Text]
		}
		for path, in := range named {
			if in[d.name.Text] && !d.hidden(path) {
				return true
			}
		}
		return false
	}
	for _, d := range ds {
		if d.name == nil || prelude.Is(d.path) || r.Source[d.path] == nil || used(d) {
			continue
		}
		l.out[d.path] = append(l.out[d.path], syntax.Diagnostic{
			Code: "lint/unused", Severity: l.sev, Sp: d.name.Sp, Unnecessary: true,
			Message: fmt.Sprintf("%s '%s' is never used", d.kind, d.name.Text),
		})
	}
}

var (
	evalPieceRe = regexp.MustCompile(`\$\{[^}]*\}`)
	evalWordRe  = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)
)

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

// lintDeprecated reports uses of declarations marked @deprecated in their
// doc comment. It reads the current declarations, so a tag added or
// removed shows up without rechecking the users.
func lintDeprecated(l *lintRun) {
	r := l.r
	why := map[*sema.Symbol]string{}
	for _, s := range r.Syms.All() {
		if s.Decl.Doc == "" {
			continue
		}
		if dep := syntax.ParseDoc(s.Decl.Doc).Deprecated; dep != nil {
			why[s] = strings.Join(strings.Fields(*dep), " ")
		}
	}
	if len(why) == 0 {
		return
	}
	type use struct {
		sp  syntax.Span
		sym *sema.Symbol
	}
	for _, d := range r.Order {
		var uses []use
		for n, in := range d.Info {
			if in.Symbol == nil {
				continue
			}
			switch n.(type) {
			case *syntax.Ident, *syntax.ResLoc, *syntax.Tag, *syntax.String:
			default:
				continue
			}
			sym := r.Syms.Same(in.Symbol)
			if _, dep := why[sym]; !dep || sym.Decl == d.D {
				continue
			}
			uses = append(uses, use{n.Span(), sym})
		}
		sort.Slice(uses, func(i, j int) bool { return uses[i].sp.Start < uses[j].sp.Start })
		for _, u := range uses {
			msg := fmt.Sprintf("%s '%s' is deprecated", u.sym.Kind.Name, u.sym.ID())
			if w := why[u.sym]; w != "" {
				msg += ": " + w
			}
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{
				Code: "lint/deprecated", Severity: l.sev, Sp: u.sp, Deprecated: true, Message: msg,
			})
		}
	}
}

func lintBookOverflow(l *lintRun, d *Decl) {
	if d.Item == nil || d.Item.Book == nil {
		return
	}
	for _, ov := range d.Item.Book.Overflows {
		msg := fmt.Sprintf("%spage %d runs to %d lines, but a book page shows 14, so the game cuts the rest off from here; start a new page with a --- line", ov.Where, ov.Page, ov.Lines)
		if ov.File != "" {
			l.ext[ov.File] = append(l.ext[ov.File], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Message: msg, Sp: ov.Sp})
			continue
		}
		l.report(d, ov.Sp, msg)
	}
}

func ruleByID(id string) *Rule {
	for _, r := range Rules {
		if r.ID == id {
			return r
		}
	}
	return nil
}

func lintSplineOrder(l *lintRun, d *Decl) {
	src := d.Unit.File.Src
	for _, s := range spline.All(d.D, src, constValue) {
		if s.List == nil {
			continue
		}
		for i, p := range s.Points {
			if !p.Location.Known {
				continue
			}
			if i == 0 || !s.Points[i-1].Location.Known || p.Location.Value >= s.Points[i-1].Location.Value {
				continue
			}
			l.report(d, p.Location.Sp, fmt.Sprintf("this point's location %s is lower than the one before it (%s); spline points must go from the lowest location to the highest",
				num(p.Location.Value), num(s.Points[i-1].Location.Value)), splineSortFix(s, src)...)
			break
		}
	}
}

// splineSortFix sorts a spline's points by location, when every location
// is known and every point is written in this file.
func splineSortFix(s *spline.Spline, src []rune) []syntax.Fix {
	sorted := append([]*spline.Point{}, s.Points...)
	for _, p := range sorted {
		sp := p.X.Span()
		if !p.Location.Known || sp.Start < s.List.Sp.Start || sp.End > s.List.Sp.End || sp.End > len(src) {
			return nil
		}
	}
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].Location.Value < sorted[j].Location.Value })
	var edits []syntax.TextEdit
	for i, p := range s.Points {
		if sorted[i] != p {
			edits = append(edits, syntax.TextEdit{Sp: p.X.Span(), Text: string(src[sorted[i].X.Span().Start:sorted[i].X.Span().End])})
		}
	}
	return []syntax.Fix{{Title: "Sort the points by location", Edits: edits}}
}

func num(v float64) string { return strconv.FormatFloat(v, 'g', -1, 64) }
