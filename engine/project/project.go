package project

import (
	"context"
	"fmt"
	"hash/fnv"
	"io/fs"
	"runtime/debug"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/types"
)

// Project is a set of source files plus configuration. It is safe for
// concurrent use; Analyze returns immutable results.
type Project struct {
	Pack   *spec.Pack
	Kinds  *sema.Kinds
	Config Config
	// Files are the project's files, which asset helpers such as
	// Texture("a.png") read. Nil means the host has none; set it before
	// the first Analyze.
	Files fs.FS
	// Root is what file paths start with before their path in the
	// project, such as "file:///home/me/pack/" when a host names files by
	// URI; "" when they are named by project path.
	Root string

	mu     sync.Mutex
	files  map[string]*sema.Unit
	texts  map[string]string
	chunks map[string][]syntax.Chunk
	// moved records declarations reparsed at a new offset since the last
	// analysis, so their checked results can be carried over.
	moved map[*syntax.Decl]syntax.Moved
	last  *Result
	// evalCache remembers which declarations need no expansion.
	evalCache eval.Cache
}

// New returns an empty project.
func New(pack *spec.Pack, cfg Config) *Project {
	return &Project{Pack: pack, Kinds: sema.NewKinds(pack), Config: cfg, files: map[string]*sema.Unit{}, texts: map[string]string{},
		chunks: map[string][]syntax.Chunk{}, moved: map[*syntax.Decl]syntax.Moved{}}
}

// SetFile adds or replaces a file. Parsing is incremental: declarations
// whose text did not change keep their parse, and their checked results
// carry over to the next analysis.
func (p *Project) SetFile(path, text string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if old, ok := p.texts[path]; ok && old == text {
		return
	}
	if sema.IsBookFile(path) {
		// Book text is not MMS: a Book declaration reads it with
		// TextFile, and its problems are reported in the file.
		p.texts[path] = text
		p.files[path] = &sema.Unit{Path: path, File: &syntax.File{Src: []rune(text)}}
		return
	}
	f, chunks, moved := syntax.Reparse(text, p.chunks[path])
	for d, m := range moved {
		// A declaration moved twice since the last analysis: chain the
		// records back to the node the analysis knows.
		if first, ok := p.moved[m.Old]; ok {
			nodes := make(map[syntax.Node]syntax.Node, len(first.Nodes))
			for oldest, mid := range first.Nodes {
				if n, ok := m.Nodes[mid]; ok {
					nodes[oldest] = n
				}
			}
			moved[d] = syntax.Moved{Old: first.Old, Nodes: nodes, Delta: first.Delta + m.Delta}
		}
	}
	for _, d := range f.Decls {
		if m, ok := p.moved[d]; ok {
			if _, again := moved[d]; !again {
				moved[d] = m // reused in place after an earlier move
			}
		}
	}
	for d, m := range moved {
		p.moved[d] = m
	}
	p.texts[path] = text
	p.chunks[path] = chunks
	p.files[path] = &sema.Unit{Path: path, File: f}
}

// RemoveFile drops a file.
func (p *Project) RemoveFile(path string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.texts, path)
	delete(p.files, path)
	delete(p.chunks, path)
}

// bookText returns a .mmsbook file's current text, by its path in the
// project. Hosts that name files by URI, such as the language server,
// have it as the URI that ends in the path.
func (p *Project) bookText(path string) (unit, text string, ok bool) {
	if !sema.IsBookFile(path) {
		return "", "", false
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if t, ok := p.texts[path]; ok {
		return path, t, true
	}
	if p.Root != "" {
		if t, ok := p.texts[p.Root+path]; ok {
			return p.Root + path, t, true
		}
	}
	for k, t := range p.texts {
		if strings.HasSuffix(k, "/"+path) && strings.Contains(k, "://") && (unit == "" || len(k) < len(unit)) {
			unit, text, ok = k, t, true
		}
	}
	return unit, text, ok
}

// Text returns a file's current text.
func (p *Project) Text(path string) (string, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	t, ok := p.texts[path]
	return t, ok
}

// Paths lists the files, sorted.
func (p *Project) Paths() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]string, 0, len(p.files))
	for path := range p.files {
		out = append(out, path)
	}
	sort.Strings(out)
	return out
}

// Result is one analysis of the whole project.
type Result struct {
	// Eval is the expansion of const, for, Template and Blueprint
	// declarations, with the provenance of what they wrote.
	Eval  *eval.Result
	Syms  *sema.Table
	Units []*sema.Unit
	// TypeEnv resolves type expressions: Template parameters, fields.
	TypeEnv *types.Env
	// Source holds each file as parsed, before expansion: with its
	// const, Template, Blueprint and for items and the calls as written.
	Source map[string]*syntax.File
	Decls  map[*syntax.Decl]*Decl
	// Consts are the file-level consts with a type, each value checked
	// for the editor against it. D stands in for the const statement.
	Consts map[*syntax.ConstStmt]*Decl
	// Order lists declarations in file and source order.
	Order []*Decl
	// Instances are the Blueprint instances the user wrote, in order,
	// checked for the editor (see instanceDecls). They write nothing.
	Instances []*Decl

	collect  map[string][]syntax.Diagnostic
	typeErr  map[string][]syntax.Diagnostic // Template arguments, Blueprint fields, typed consts
	assets   []assetUse
	assetErr map[string][]syntax.Diagnostic
	orderErr map[string][]syntax.Diagnostic // feature order cycles
	emitted  map[*syntax.Decl]bool          // declarations Blueprint instances emitted
	langErr  map[string][]syntax.Diagnostic // keys two Lang declarations set
	formErr  map[string][]syntax.Diagnostic // declaration forms the language dropped
	selErr   map[string][]syntax.Diagnostic // selectors written with unknown options
	jigsaw   *jigsawModel                   // structure files' jigsaw blocks, for the jigsaw rules
	files    fs.FS
	lintMu   sync.Mutex
	// models are the project's models' slot facts, by full id, built
	// on first use by modelInfo.
	// atlasSources are the sources of the atlases models read, by
	// name, built on first use by modelAtlasesOf.
	atlasOnce    sync.Once
	atlasSources map[string][]spec.AtlasSource
	modelsOnce   sync.Once
	models       map[string]*spec.ModelInfo
	relOnce      sync.Once
	written      map[eval.Origin]eval.Prov // see relate
	diagOnce     sync.Once
	own          map[string][]syntax.Diagnostic // each file's own problems, see Diagnostics
	echoes       map[string][]syntax.Diagnostic // problems shown again where they were written
	lintDone     bool
	lints        map[string][]syntax.Diagnostic
	symKey       uint64
	pack         *spec.Pack
	kinds        *sema.Kinds
	cfg          Config
}

// Decl is the checked form of one declaration.
type Decl struct {
	Unit  *sema.Unit
	D     *syntax.Decl
	Kind  sema.Kind
	NS    string
	Value jsonv.Value // JSON resources
	Tag   jsonv.Array // tag entries
	// TagReplace is set when a tag replaces lower packs' entries.
	TagReplace bool
	Lines      []string // raw function commands, or load lines
	// Fn is a compiled function's output: its files and what the load
	// function must set up for it.
	Fn *fn.Output
	// Item is a checked CustomItem: the files it writes and the stack
	// its name stands for.
	Item *sema.CustomItemOut
	// Hoisted are inline values written where an id is required, each
	// written as its own resource.
	Hoisted []sema.Hoisted
	// Shown are the messages its text commands show, for a chat preview.
	Shown []sema.ShownText
	Info  map[syntax.Node]*sema.Info
	// Blueprint is set on a Blueprint instance's body, checked for the
	// editor only.
	Blueprint *eval.Blueprint
	// Template is set on a Template's body, checked for the editor only
	// against its return type, or read as function code for Commands.
	Template *eval.Template
	// Emitter is set on a declaration in a Blueprint's emit block,
	// checked for the editor only.
	Emitter *eval.Blueprint
	Diags   []syntax.Diagnostic
	// Extern are problems in other files, such as a Book's .mmsbook
	// file, by path.
	Extern map[string][]syntax.Diagnostic
	// Expects are a Test's expect lines, in order.
	Expects []sema.TestExpect
	// Deps are the symbol lookups the check made (see sema.DepsHold).
	Deps map[string]bool
	// id is a Blueprint instance's id, which a top-level for gives each
	// copy of one declaration.
	id string

	lints    []syntax.Diagnostic            // cached lint findings, valid for lintKey
	extLints map[string][]syntax.Diagnostic // the same, in other files
	lintKey  string
	refs     []string
	refsDone bool
}

// ID returns "ns:path" for a named declaration.
func (d *Decl) ID() string {
	if d.id != "" {
		return d.id
	}
	if d.D.Name == nil {
		return ""
	}
	if strings.Contains(d.D.Name.Text, ":") {
		return d.D.Name.Text
	}
	return d.NS + ":" + d.D.Name.Text
}

// Unit returns a file's parsed unit from this result.
func (r *Result) Unit(path string) *sema.Unit {
	for _, u := range r.Units {
		if u.Path == path {
			return u
		}
	}
	return nil
}

// Diagnostics returns every diagnostic for one file, sorted by position.
// A problem in text a Blueprint, Template or const wrote is reported
// where it is used, and shown again at the line of the body that wrote
// it, marked Echo.
func (r *Result) Diagnostics(path string) []syntax.Diagnostic {
	r.diagOnce.Do(r.diagnose)
	own, echoes := r.own[path], r.echoes[path]
	if own == nil && echoes == nil {
		return nil
	}
	out := make([]syntax.Diagnostic, 0, len(own)+len(echoes))
	out = append(append(out, own...), echoes...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Sp.Start < out[j].Sp.Start })
	return out
}

// diagnose works out every file's problems, and the echoes of those
// whose text was written in another place.
func (r *Result) diagnose() {
	r.own, r.echoes = map[string][]syntax.Diagnostic{}, map[string][]syntax.Diagnostic{}
	type key struct {
		path string
		sp   syntax.Span
		code string
		msg  string
	}
	at := map[key]int{}
	whole := r.wholeDefs()
	for _, u := range r.Units {
		r.own[u.Path] = r.fileDiagnostics(u.Path)
	}
	r.dropConstRepeats()
	for _, u := range r.Units {
		ds := r.own[u.Path]
		for _, d := range ds {
			if d.Echo {
				continue
			}
			for _, rel := range d.Related {
				if r.Unit(rel.Path) == nil || (rel.Path == u.Path && rel.Sp == d.Sp) {
					continue // the prelude, say, which nobody edits
				}
				if whole[eval.Origin{Path: rel.Path, Sp: rel.Sp}] {
					// No line of the body wrote it, only the Blueprint
					// or Template as a whole, such as an unused
					// instance: shown over the whole body, it reads as
					// the Blueprint's own problem.
					continue
				}
				back := syntax.Related{Path: u.Path, Sp: d.Sp, Message: usedAt(u, d.Sp)}
				k := key{rel.Path, rel.Sp, d.Code, d.Message}
				if i, ok := at[k]; ok {
					// Every use of the Blueprint met the same problem:
					// one echo, pointing at each use.
					e := &r.echoes[rel.Path][i]
					if len(e.Related) < maxEchoUses {
						e.Related = append(e.Related, back)
					}
					continue
				}
				at[k] = len(r.echoes[rel.Path])
				r.echoes[rel.Path] = append(r.echoes[rel.Path], syntax.Diagnostic{Code: d.Code, Severity: d.Severity,
					Message: d.Message, Sp: rel.Sp, Related: []syntax.Related{back}, Echo: true})
			}
		}
	}
}

// dropConstRepeats drops a problem found in a typed const's value where
// the const is used, when the same problem is reported elsewhere: the
// const's own check found it where the value is written, and each use,
// and each const built from it, would only say it again.
func (r *Result) dropConstRepeats() {
	if r.Eval == nil {
		return
	}
	uses := r.typedConstUses()
	repeat := func(path string, d syntax.Diagnostic) bool {
		for _, sp := range uses[path] {
			if d.Sp.Start >= sp.Start && d.Sp.End <= sp.End {
				return true
			}
		}
		return false
	}
	type key struct{ code, msg string }
	root := map[key]bool{}
	for path, ds := range r.own {
		for _, d := range ds {
			if !d.Echo && !repeat(path, d) {
				root[key{d.Code, d.Message}] = true
			}
		}
	}
	for path, ds := range r.own {
		kept := ds[:0:0]
		for _, d := range ds {
			if !d.Echo && root[key{d.Code, d.Message}] && repeat(path, d) {
				continue
			}
			kept = append(kept, d)
		}
		r.own[path] = kept
	}
}

// typedConstUses are the spans, by file, of the names of typed consts
// where they are used: `c` for `const c: T = ...`, and `ash.noise_entries`
// for an instance's const that its Blueprint gives a type.
func (r *Result) typedConstUses() map[string][]syntax.Span {
	emitTyped := map[string]bool{}
	for _, b := range r.Eval.Blueprints {
		for _, e := range b.Emits {
			for _, it := range e.Items {
				if k, ok := it.(*syntax.ConstStmt); ok && k.Type != nil && k.Name != nil {
					emitTyped[k.Name.Text] = true
				}
			}
		}
	}
	out := map[string][]syntax.Span{}
	for path, f := range r.Source {
		syntax.Walk(f, func(n syntax.Node) bool {
			switch n := n.(type) {
			case *syntax.Ident:
				if k := r.Eval.Const(path, n.Name); k != nil && k.S != nil && k.S.Type != nil && k.S.Name != nil && k.S.Name.Sp != n.Sp {
					out[path] = append(out[path], n.Sp)
				}
			case *syntax.Member:
				if _, ok := n.X.(*syntax.Ident); ok && n.Name != nil && emitTyped[n.Name.Text] {
					out[path] = append(out[path], n.Sp)
					return false
				}
			}
			return true
		})
	}
	return out
}

// wholeDefs are the spans of the Blueprints and Templates, which a
// problem points at when no single line of the body wrote it.
func (r *Result) wholeDefs() map[eval.Origin]bool {
	out := map[eval.Origin]bool{}
	if r.Eval == nil {
		return out
	}
	for _, b := range r.Eval.Blueprints {
		out[eval.Origin{Path: b.Path, Sp: b.D.Sp}] = true
	}
	for _, t := range r.Eval.Templates {
		out[eval.Origin{Path: t.Path, Sp: t.D.Sp}] = true
	}
	return out
}

// maxEchoUses is how many uses an echoed problem points back at.
const maxEchoUses = 5

// usedAt describes the use at sp, by its line: "found at throttle
// spellbook {}".
func usedAt(u *sema.Unit, sp syntax.Span) string {
	line, _ := u.File.LineCol(sp.Start)
	text := strings.TrimSpace(u.File.LineText(line))
	if r := []rune(text); len(r) > 60 {
		text = string(r[:57]) + "..."
	}
	return "found where it is used: " + text
}

// fileDiagnostics is one file's own problems, sorted by position.
func (r *Result) fileDiagnostics(path string) []syntax.Diagnostic {
	u := r.Unit(path)
	if u == nil {
		return nil
	}
	var out []syntax.Diagnostic
	out = append(out, u.File.Diags...)
	if r.Eval != nil {
		if f := r.Eval.Files[path]; f != nil {
			out = append(out, f.Diags...)
		}
	}
	out = append(out, r.collect[path]...)
	out = append(out, r.typeErr[path]...)
	out = append(out, r.assetErr[path]...)
	out = append(out, r.orderErr[path]...)
	out = append(out, r.langErr[path]...)
	out = append(out, r.formErr[path]...)
	out = append(out, r.selErr[path]...)
	typeErr := r.typeErr[path]
	for _, d := range r.Order {
		if d.Unit == u {
			for _, dg := range d.Diags {
				// A Template argument or Blueprint field with a type
				// error was reported where it is written; the
				// expanded declaration's own complaint about it,
				// which covers that spot, adds nothing.
				if !coversAny(dg.Sp, typeErr) {
					out = append(out, dg)
				}
			}
		}
		out = append(out, d.Extern[path]...)
	}
	for _, d := range r.Instances {
		if d.Unit == u {
			out = append(out, d.Diags...)
		}
		out = append(out, d.Extern[path]...)
	}
	out = dropSelectorEchoes(out, r.selErr[path])
	out = r.completeTagListFixes(u, out)
	out = r.configure(path, out)
	out = append(out, r.Lints()[path]...)
	r.relate(path, out)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Sp.Start < out[j].Sp.Start })
	return out
}

// Lints returns lint findings by file, computed on first use.
func (r *Result) Lints() map[string][]syntax.Diagnostic {
	r.lintMu.Lock()
	defer r.lintMu.Unlock()
	if !r.lintDone {
		r.lints, r.lintDone = r.lint(), true
	}
	return r.lints
}

// HasErrors reports whether any file has an error.
func (r *Result) HasErrors() bool {
	for _, u := range r.Units {
		for _, d := range r.Diagnostics(u.Path) {
			if d.Severity == syntax.SeverityError {
				return true
			}
		}
	}
	return false
}

// Analyze checks the whole project. Declarations in files whose text did
// not change since the last analysis reuse their results, as long as the
// set of declared names is the same: checking a declaration depends only
// on its own text, its namespace and which names exist; when the names
// change, results are kept for declarations whose lookups still resolve
// the same way.
func (p *Project) Analyze(ctx context.Context) (*Result, error) {
	p.mu.Lock()
	units := make([]*sema.Unit, 0, len(p.files))
	for _, u := range p.files {
		units = append(units, u)
	}
	prev := p.last
	// Move records point at nodes of the previous analysis; edits from
	// now on are recorded relative to this one.
	moved := p.moved
	p.moved = map[*syntax.Decl]syntax.Moved{}
	p.mu.Unlock()
	sort.Slice(units, func(i, j int) bool { return units[i].Path < units[j].Path })

	source := make(map[string]*syntax.File, len(units))
	for _, u := range units {
		source[u.Path] = u.File
	}
	units, ev := expand(units, p.Config.Namespace, &p.evalCache, p.Pack)
	syms, collect := sema.Collect(units, p.Kinds, p.Config.Namespace)
	tenv := types.New(p.Pack, p.Kinds, ev)
	syms.Types = tenv.Types
	syms.TypeDecls = map[string]*syntax.Decl{}
	for name, td := range ev.Types {
		syms.TypeDecls[name] = td.D
	}
	syms.TypeKey = ev.TypeKey
	syms.Hidden = preludeEmitted(ev)
	syms.Conversions = tenv.Conversions(ev, p.Config.Namespace)
	r := &Result{Syms: syms, Units: units, Source: source, Decls: map[*syntax.Decl]*Decl{}, collect: collect, pack: p.Pack, kinds: p.Kinds, cfg: p.Config, files: p.Files, Eval: ev}
	r.TypeEnv = tenv
	r.selErr = map[string][]syntax.Diagnostic{}
	for path, f := range source {
		r.selErr[path] = sema.CheckSelectors(f)
	}
	r.typeErr = tenv.Check(ev, p.Config.Namespace, func(ns string) *sema.Checker {
		c := sema.NewChecker(p.Pack, p.Kinds, syms, ns)
		c.Files, c.Texts = p.Files, p.bookText
		return c
	}, func(path string) []rune {
		if f := source[path]; f != nil {
			return f.Src
		}
		return nil
	})
	r.symKey = symbolKey(syms)
	sameSyms := prev != nil && prev.symKey == r.symKey
	for _, u := range units {
		for _, d := range u.File.Decls {
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			k, ok := syms.DeclKind[d]
			if !ok || k.MMS == "Namespace" {
				continue
			}
			var dr *Decl
			if prev != nil {
				dr = prev.reuse(u, d, syms.DeclNS[d], moved)
				// With different declared names, a result still holds if
				// every name it looked up resolves the same way.
				if dr != nil && !sameSyms && (dr.Deps == nil || !sema.DepsHold(syms, dr.Deps)) {
					dr = nil
				}
				// A Storage reads its Type's defaults, which are cheap to
				// read again and not among its dependencies.
				if dr != nil && k.MMS == "Storage" {
					dr = nil
				}
				// An asset file it read was added, removed or changed.
				if dr != nil && !sema.AssetDepsHold(p.Files, dr.Deps) {
					dr = nil
				}
				// A .mmsbook file it read was edited.
				if dr != nil && !sema.TextDepsHold(p.bookText, dr.Deps) {
					dr = nil
				}
			}
			if dr == nil {
				dr = p.checkDecl(u, d, k, syms)
			}
			r.Decls[d] = dr
			r.Order = append(r.Order, dr)
		}
	}
	r.instanceDecls(tenv, func(ns string) *sema.Checker {
		c := sema.NewChecker(p.Pack, p.Kinds, syms, ns)
		c.Files, c.Texts = p.Files, p.bookText
		return c
	})
	r.templateDecls(tenv, p.Kinds, func(ns string) *sema.Checker {
		c := sema.NewChecker(p.Pack, p.Kinds, syms, ns)
		c.Files, c.Texts = p.Files, p.bookText
		return c
	})
	// A declaration in a Blueprint's emit block is checked as written,
	// for the editor: its Blueprint's fields are names the checker does
	// not know, so its diagnostics are dropped, and it writes nothing.
	for _, bp := range ev.Blueprints {
		u := r.Unit(bp.Path)
		if u == nil || prelude.Is(bp.Path) {
			continue
		}
		// The Blueprint's own head: each `as` value checked against its
		// type.
		c := sema.NewChecker(p.Pack, p.Kinds, syms, p.Config.Namespace)
		c.Files, c.Texts, c.Src = p.Files, p.bookText, u.File.Src
		for _, as := range bp.As {
			if as.Value == nil {
				continue
			}
			if want := tenv.Lookup(bp.Path, as.Type); want != nil && want.Kind != "any" {
				c.CheckType(as.Value, want, syntax.TypeText(as.Type))
			}
		}
		// Typed lets and field defaults too, so their values complete
		// like the type's: `let g: DensityFunction = { ty| }`.
		for _, l := range bp.Lets {
			if l.Type == nil || l.Value == nil {
				continue
			}
			if want := tenv.Lookup(bp.Path, l.Type); want != nil && want.Kind != "any" {
				c.CheckType(l.Value, want, syntax.TypeText(l.Type))
			}
		}
		for _, f := range bp.Fields {
			if f.Type == nil || f.Default == nil {
				continue
			}
			if want := tenv.Lookup(bp.Path, f.Type); want != nil && want.Kind != "any" {
				c.CheckType(f.Default, want, syntax.TypeText(f.Type))
			}
		}
		r.Decls[bp.D] = &Decl{Unit: u, D: bp.D, NS: p.Config.Namespace, Info: c.Info, Emitter: bp}
		for _, b := range bp.Emits {
			syntax.Walk(b, func(n syntax.Node) bool {
				// A typed const the emit block computes with completes
				// like its type, as a typed let does.
				if k, ok := n.(*syntax.ConstStmt); ok && k.Type != nil && k.Value != nil {
					if want := tenv.Lookup(bp.Path, k.Type); want != nil && want.Kind != "any" {
						c.CheckType(k.Value, want, syntax.TypeText(k.Type))
					}
					return false
				}
				d, ok := n.(*syntax.Decl)
				if !ok {
					return true
				}
				if d.Kind == nil || r.Decls[d] != nil {
					return false
				}
				// Another Blueprint's instance, such as a CustomItem.
				if inner := ev.Blueprint(bp.Path, d.Kind.Name); inner != nil && d.Body != nil {
					c := sema.NewChecker(p.Pack, p.Kinds, syms, p.Config.Namespace)
					c.Files, c.Texts, c.Src = p.Files, p.bookText, u.File.Src
					c.CheckType(&syntax.Object{Body: d.Body, Sp: d.Body.Sp}, tenv.InstanceType(inner), d.Kind.Name)
					r.Decls[d] = &Decl{Unit: u, D: d, NS: p.Config.Namespace, Info: c.Info, Emitter: bp, Blueprint: inner}
					return false
				}
				k, ok := p.Kinds.Lookup(d.Kind.Name)
				if !ok {
					return false
				}
				dr := p.checkDecl(u, d, k, syms)
				dr.NS, dr.Diags, dr.Extern, dr.Emitter = p.Config.Namespace, nil, nil, bp
				r.Decls[d] = dr
				return false
			})
		}
	}
	r.Consts = map[*syntax.ConstStmt]*Decl{}
	for _, k := range ev.Consts {
		u := r.Unit(k.Path)
		if u == nil || prelude.Is(k.Path) || k.S.Type == nil || k.S.Value == nil {
			continue
		}
		c := sema.NewChecker(p.Pack, p.Kinds, syms, p.Config.Namespace)
		c.Files, c.Texts, c.Src = p.Files, p.bookText, u.File.Src
		if want := tenv.Lookup(k.Path, k.S.Type); want != nil && want.Kind != "any" {
			c.CheckType(k.S.Value, want, syntax.TypeText(k.S.Type))
		}
		d := &syntax.Decl{Name: k.S.Name, Type: k.S.Type, Value: k.S.Value, Doc: k.S.Doc, Sp: k.S.Sp}
		r.Consts[k.S] = &Decl{Unit: u, D: d, NS: p.Config.Namespace, Info: c.Info}
	}
	r.missingFieldFixes(tenv, p.Kinds, source)
	r.formErr = map[string][]syntax.Diagnostic{}
	for path, f := range source {
		if ds := checkForms(path, f); ds != nil {
			r.formErr[path] = ds
		}
	}
	r.collectAssets()
	r.checkFeatureOrder()
	r.checkLangKeys()
	p.mu.Lock()
	p.last = r
	p.mu.Unlock()
	return r, nil
}

// reuse returns d's result from this (previous) analysis, if d was not
// changed: the same node, or a node reparsed at a new offset from
// unchanged text, whose Info and diagnostics are moved over.
func (r *Result) reuse(u *sema.Unit, d *syntax.Decl, ns string, moved map[*syntax.Decl]syntax.Moved) *Decl {
	if old := r.Decls[d]; old != nil && old.NS == ns {
		if old.Unit == u {
			return old
		}
		cp := *old
		cp.Unit = u
		return &cp
	}
	m, ok := moved[d]
	if !ok {
		return nil
	}
	old := r.Decls[m.Old]
	if old == nil || old.NS != ns {
		return nil
	}
	cp := *old
	cp.Unit, cp.D = u, d
	cp.Info = make(map[syntax.Node]*sema.Info, len(old.Info))
	for n, in := range old.Info {
		if nn, ok := m.Nodes[n]; ok {
			cp.Info[nn] = in
		} else if nn := syntax.ShiftLeaf(n, m.Delta); nn != nil {
			cp.Info[nn] = in // a reference inside a command
		}
	}
	cp.Diags = make([]syntax.Diagnostic, len(old.Diags))
	for i, dg := range old.Diags {
		cp.Diags[i] = dg.Shift(m.Delta)
	}
	cp.lints, cp.lintKey = nil, "" // spans moved; lint again
	return &cp
}

// symbolKey fingerprints the declared names and declaration namespaces.
func symbolKey(t *sema.Table) uint64 {
	var keys []string
	for _, s := range t.All() {
		keys = append(keys, s.Kind.Folder()+"|"+s.ID()+"|"+sema.SymSig(s))
	}
	for id, cs := range t.Conversions {
		keys = append(keys, "as|"+id+"|"+sema.ConversionsSig(cs))
	}
	sort.Strings(keys)
	h := fnv.New64a()
	for _, k := range keys {
		h.Write([]byte(k))
		h.Write([]byte{0})
	}
	return h.Sum64()
}

// checkDecl checks one declaration; a panic becomes an internal-error
// diagnostic so the rest of the project keeps working.
func (p *Project) checkDecl(u *sema.Unit, d *syntax.Decl, k sema.Kind, syms *sema.Table) (dr *Decl) {
	ns := syms.DeclNS[d]
	dr = &Decl{Unit: u, D: d, Kind: k, NS: ns}
	c := sema.NewChecker(p.Pack, p.Kinds, syms, ns)
	c.Files = p.Files
	c.Texts = p.bookText
	c.Src = u.File.Src
	c.Owner = dr.ID()
	if c.Owner == "" {
		c.Owner = ns + ":" + sema.Snake(k.Name)
	}
	defer func() {
		if rec := recover(); rec != nil {
			dr.Value, dr.Tag, dr.Lines, dr.Fn, dr.Item, dr.Hoisted = nil, nil, nil, nil, nil, nil
			dr.Info = c.Info
			dr.Diags = append(c.Diags, syntax.Diagnostic{
				Code: "mms/internal", Severity: syntax.SeverityError, Sp: d.Kind.Sp,
				Message: fmt.Sprintf("internal error while checking this declaration: %v. Please report this with the declaration's source.", rec),
			})
			if debugPanics {
				panic(fmt.Sprintf("%v\n%s", rec, debug.Stack()))
			}
		}
	}()
	switch {
	case k.TagOf != "":
		dr.Tag, dr.TagReplace = c.Tag(d, k)
	case k.IsResource():
		dr.Value = c.Resource(d, k)
	case k.MMS == "Function":
		if f := c.CompileFunction(d, dr.ID(), u.File.Src); f != nil && d.Name != nil {
			dr.Fn = fn.Generate(*f)
		}
	case k.MMS == "OnAdvancement":
		adv, f := c.OnAdvancement(d, dr.ID(), u.File.Src)
		if f != nil {
			dr.Value, dr.Fn = adv, fn.Generate(*f)
		}
	case sema.HookKinds[k.MMS] != "":
		if sema.IsRaw(d.Body) {
			dr.Lines = c.Function(d)
			break
		}
		id := sema.GeneratedID(ns, k.MMS, d, u.File.Src)
		if f := c.CompileFunction(d, id, u.File.Src); f != nil {
			dr.Fn = fn.Generate(*f)
			dr.Lines = []string{"function " + id}
		}
	case k.MMS == "Objective":
		dr.Lines = objective(c, d)
	case k.MMS == "Global":
		dr.Lines = c.Global(d)
	case k.MMS == "Storage":
		dr.Lines = c.Storage(d)
	case k.MMS == "Test":
		if d.Name != nil {
			f, ex := c.CompileTest(d, sema.TestID(ns, d.Name.Text), u.File.Src)
			if f != nil {
				dr.Fn, dr.Expects = fn.Generate(*f), ex
			}
		} else {
			c.CompileTest(d, "", u.File.Src)
		}
	case k.MMS == sema.JigsawHoleKind:
		c.JigsawHole(d)
	}
	dr.Info = c.Info
	dr.Diags = c.Diags
	dr.Extern = c.Extern
	dr.Deps = c.Deps
	dr.Hoisted = c.Hoisted
	dr.Shown = c.Shown
	// Re-validate lowered output: a value that passed the checker but not
	// the schema means a lowering bug, reported as an internal error
	// rather than written as bad JSON.
	if dr.Value != nil && !hasErrors(dr.Diags) {
		res := k.Resource
		if k.MMS == "OnAdvancement" {
			res = "advancement"
		}
		if errs := sema.Validate(p.Pack, p.Kinds, res, dr.Value); len(errs) > 0 {
			dr.Diags = append(dr.Diags, syntax.Diagnostic{
				Code: "mms/internal", Severity: syntax.SeverityError, Sp: d.Kind.Sp,
				Message: "internal error: the generated JSON does not match the schema (" + errs[0] + "). Please report this with the declaration's source.",
			})
		}
	}
	for _, h := range dr.Hoisted {
		if hasErrors(dr.Diags) {
			break
		}
		if h.Value == nil {
			continue
		}
		if errs := sema.Validate(p.Pack, p.Kinds, h.Kind.Resource, h.Value); len(errs) > 0 {
			dr.Diags = append(dr.Diags, syntax.Diagnostic{
				Code: "mms/internal", Severity: syntax.SeverityError, Sp: h.Sp,
				Message: "internal error: the generated " + h.Kind.Resource + " JSON for this inline value does not match the schema (" + errs[0] + "). Please report this with the declaration's source.",
			})
		}
	}
	return dr
}

func hasErrors(ds []syntax.Diagnostic) bool {
	for _, d := range ds {
		if d.Severity == syntax.SeverityError {
			return true
		}
	}
	return false
}

// debugPanics re-raises checker panics, for tests.
var debugPanics = false

// objective turns `Objective kills: deathCount` into its setup command.
// The old `= deathCount` form still compiles, next to its error.
func objective(c *sema.Checker, d *syntax.Decl) []string {
	if d.Name == nil {
		return nil
	}
	criterion := "dummy"
	if ts, ok := d.Type.(*syntax.TypeString); ok {
		criterion = ts.Value
	} else if d.Type != nil {
		criterion = syntax.TypeText(d.Type)
	}
	switch v := d.Value.(type) {
	case nil:
	case *syntax.Ident:
		criterion = v.Name
	case *syntax.ResLoc:
		criterion = v.Text
	case *syntax.String:
		criterion = v.Value
	default:
		c.Diags = append(c.Diags, syntax.Diagnostic{Code: "mms/type-mismatch", Severity: syntax.SeverityError, Sp: v.Span(),
			Message: "expected an objective criterion such as dummy or deathCount"})
	}
	return []string{"scoreboard objectives add " + d.Name.Text + " " + criterion}
}

// dropSelectorEchoes drops a command's complaint about a selector that
// the selector check already reported, with its fix, at the option or
// head that is wrong: an execute helper's chain is checked as a command
// too.
func dropSelectorEchoes(ds, sel []syntax.Diagnostic) []syntax.Diagnostic {
	if len(sel) == 0 {
		return ds
	}
	out := ds[:0]
	for _, d := range ds {
		if d.Code == "mms/command" && overlapsAny(d.Sp, sel) {
			continue
		}
		out = append(out, d)
	}
	return out
}

func overlapsAny(sp syntax.Span, ds []syntax.Diagnostic) bool {
	for _, d := range ds {
		if d.Sp.Start < sp.End && sp.Start < d.Sp.End {
			return true
		}
	}
	return false
}

// coversAny reports whether sp holds the span of one of ds.
func coversAny(sp syntax.Span, ds []syntax.Diagnostic) bool {
	for _, d := range ds {
		if sp.Start <= d.Sp.Start && d.Sp.End <= sp.End {
			return true
		}
	}
	return false
}

// Fork returns a copy of the project whose files can be changed without
// touching this one, such as to complete in a repaired copy of a file
// being typed. It starts from this project's last analysis.
func (p *Project) Fork() *Project {
	p.mu.Lock()
	defer p.mu.Unlock()
	q := &Project{Pack: p.Pack, Kinds: p.Kinds, Config: p.Config, Files: p.Files, Root: p.Root,
		files: make(map[string]*sema.Unit, len(p.files)), texts: make(map[string]string, len(p.texts)),
		chunks: make(map[string][]syntax.Chunk, len(p.chunks)), moved: map[*syntax.Decl]syntax.Moved{}, last: p.last}
	for k, v := range p.files {
		q.files[k] = v
	}
	for k, v := range p.texts {
		q.texts[k] = v
	}
	for k, v := range p.chunks {
		q.chunks[k] = v
	}
	return q
}
