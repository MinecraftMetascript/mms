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

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
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

	mu     sync.Mutex
	files  map[string]*sema.Unit
	texts  map[string]string
	chunks map[string][]syntax.Chunk
	// moved records declarations reparsed at a new offset since the last
	// analysis, so their checked results can be carried over.
	moved map[*syntax.Decl]syntax.Moved
	last  *Result
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
	Syms  *sema.Table
	Units []*sema.Unit
	Decls map[*syntax.Decl]*Decl
	// Order lists declarations in file and source order.
	Order []*Decl

	collect  map[string][]syntax.Diagnostic
	assets   []assetUse
	assetErr map[string][]syntax.Diagnostic
	orderErr map[string][]syntax.Diagnostic // feature order cycles
	langErr  map[string][]syntax.Diagnostic // CustomItem names a Lang also has
	files    fs.FS
	lintMu   sync.Mutex
	lintDone bool
	lints    map[string][]syntax.Diagnostic
	symKey   uint64
	pack     *spec.Pack
	cfg      Config
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
	Info    map[syntax.Node]*sema.Info
	Diags   []syntax.Diagnostic
	// Deps are the symbol lookups the check made (see sema.DepsHold).
	Deps map[string]bool

	lints    []syntax.Diagnostic // cached lint findings, valid for lintKey
	lintKey  string
	refs     []string
	refsDone bool
}

// ID returns "ns:path" for a named declaration.
func (d *Decl) ID() string {
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
func (r *Result) Diagnostics(path string) []syntax.Diagnostic {
	u := r.Unit(path)
	if u == nil {
		return nil
	}
	var out []syntax.Diagnostic
	out = append(out, u.File.Diags...)
	out = append(out, r.collect[path]...)
	out = append(out, r.assetErr[path]...)
	out = append(out, r.orderErr[path]...)
	out = append(out, r.langErr[path]...)
	for _, d := range r.Order {
		if d.Unit == u {
			out = append(out, d.Diags...)
		}
	}
	out = append(out, r.Lints()[path]...)
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

	syms, collect := sema.Collect(units, p.Kinds, p.Config.Namespace)
	r := &Result{Syms: syms, Units: units, Decls: map[*syntax.Decl]*Decl{}, collect: collect, pack: p.Pack, cfg: p.Config, files: p.Files}
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
				// An asset file it read was added, removed or changed.
				if dr != nil && !sema.AssetDepsHold(p.Files, dr.Deps) {
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
	r.collectAssets()
	r.checkFeatureOrder()
	r.checkCustomItemLang()
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
	case k.MMS == "CustomItem":
		dr.Item = c.CustomItem(d, dr.ID())
	case k.MMS == "Objective":
		dr.Lines = objective(c, d)
	case k.MMS == "Global":
		dr.Lines = c.Global(d)
	}
	dr.Info = c.Info
	dr.Diags = c.Diags
	dr.Deps = c.Deps
	dr.Hoisted = c.Hoisted
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
	if it := dr.Item; it != nil && !hasErrors(dr.Diags) {
		for _, f := range []struct {
			res string
			v   jsonv.Value
		}{{"model", it.Model}, {"item_definition", it.ItemModel}} {
			if f.v == nil {
				continue
			}
			if errs := sema.Validate(p.Pack, p.Kinds, f.res, f.v); len(errs) > 0 {
				dr.Diags = append(dr.Diags, syntax.Diagnostic{
					Code: "mms/internal", Severity: syntax.SeverityError, Sp: d.Kind.Sp,
					Message: "internal error: the generated " + f.res + " JSON does not match the schema (" + errs[0] + "). Please report this with the declaration's source.",
				})
			}
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

// objective turns `Objective kills = deathCount` into its setup command.
func objective(c *sema.Checker, d *syntax.Decl) []string {
	if d.Name == nil {
		return nil
	}
	criterion := "dummy"
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
