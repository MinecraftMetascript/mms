package project

import (
	"context"
	"fmt"
	"hash/fnv"
	"runtime/debug"
	"sort"
	"strings"
	"sync"

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

	mu    sync.Mutex
	files map[string]*sema.Unit
	texts map[string]string
	last  *Result
}

// New returns an empty project.
func New(pack *spec.Pack, cfg Config) *Project {
	return &Project{Pack: pack, Kinds: sema.NewKinds(pack), Config: cfg, files: map[string]*sema.Unit{}, texts: map[string]string{}}
}

// SetFile adds or replaces a file. Unchanged text keeps the old parse, so
// cached results for the file stay valid.
func (p *Project) SetFile(path, text string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if old, ok := p.texts[path]; ok && old == text {
		return
	}
	p.texts[path] = text
	p.files[path] = &sema.Unit{Path: path, File: syntax.Parse(text)}
}

// RemoveFile drops a file.
func (p *Project) RemoveFile(path string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.texts, path)
	delete(p.files, path)
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

	collect map[string][]syntax.Diagnostic
	symKey  uint64
	pack    *spec.Pack
	cfg     Config
}

// Decl is the checked form of one declaration.
type Decl struct {
	Unit  *sema.Unit
	D     *syntax.Decl
	Kind  sema.Kind
	NS    string
	Value jsonv.Value // JSON resources
	Tag   jsonv.Array // tag entries
	Lines []string    // function commands
	Info  map[syntax.Node]*sema.Info
	Diags []syntax.Diagnostic
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
	for _, d := range r.Order {
		if d.Unit == u {
			out = append(out, d.Diags...)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Sp.Start < out[j].Sp.Start })
	return out
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
// on its own text, its namespace and which names exist.
func (p *Project) Analyze(ctx context.Context) (*Result, error) {
	p.mu.Lock()
	units := make([]*sema.Unit, 0, len(p.files))
	for _, u := range p.files {
		units = append(units, u)
	}
	prev := p.last
	p.mu.Unlock()
	sort.Slice(units, func(i, j int) bool { return units[i].Path < units[j].Path })

	syms, collect := sema.Collect(units, p.Kinds, p.Config.Namespace)
	r := &Result{Syms: syms, Units: units, Decls: map[*syntax.Decl]*Decl{}, collect: collect, pack: p.Pack, cfg: p.Config}
	r.symKey = symbolKey(syms)
	reuse := prev != nil && prev.symKey == r.symKey
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
			if reuse {
				dr = prev.Decls[d]
			}
			if dr == nil {
				dr = p.checkDecl(u, d, k, syms)
			}
			r.Decls[d] = dr
			r.Order = append(r.Order, dr)
		}
	}
	p.mu.Lock()
	p.last = r
	p.mu.Unlock()
	return r, nil
}

// symbolKey fingerprints the declared names and declaration namespaces.
func symbolKey(t *sema.Table) uint64 {
	var keys []string
	for _, s := range t.All() {
		keys = append(keys, s.Kind.Folder()+"|"+s.ID())
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
	defer func() {
		if rec := recover(); rec != nil {
			dr.Value, dr.Tag, dr.Lines = nil, nil, nil
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
		dr.Tag = c.Tag(d, k)
	case k.IsResource():
		dr.Value = c.Resource(d, k)
	case k.MMS == "Function" || k.MMS == "OnLoad" || k.MMS == "OnTick":
		dr.Lines = c.Function(d)
	case k.MMS == "Objective":
		dr.Lines = objective(c, d)
	case k.MMS == "Global":
		c.Diags = append(c.Diags, syntax.Diagnostic{Code: "mms/unsupported", Severity: syntax.SeverityError, Sp: d.Kind.Sp,
			Message: "Global arrives with the function backend"})
	}
	dr.Info = c.Info
	dr.Diags = c.Diags
	// Re-validate lowered output: a value that passed the checker but not
	// the schema means a lowering bug, reported as an internal error
	// rather than written as bad JSON.
	if dr.Value != nil && !hasErrors(dr.Diags) {
		if errs := sema.Validate(p.Pack, p.Kinds, k.Resource, dr.Value); len(errs) > 0 {
			dr.Diags = append(dr.Diags, syntax.Diagnostic{
				Code: "mms/internal", Severity: syntax.SeverityError, Sp: d.Kind.Sp,
				Message: "internal error: the generated JSON does not match the schema (" + errs[0] + "). Please report this with the declaration's source.",
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
