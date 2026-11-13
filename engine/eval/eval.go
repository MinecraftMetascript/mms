// Package eval expands the compile-time parts of MMS: const values, the
// top-level for, Template calls and Blueprint declarations. It turns a
// project's files into plain declarations that the checker and the build
// already understand, and remembers for each piece of output the user
// span that caused it (its provenance), so diagnostics, go to definition
// and previews can lead back to the line the user wrote.
//
// Evaluation is pure and deterministic: no I/O, no clock, and a step
// budget, so a runaway Template is a diagnostic rather than a hang in the
// browser.
package eval

import (
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// DefaultBudget is how many evaluation steps one top-level item may
// take. Each expression, statement and loop iteration is one step; the
// largest example's biggest item takes a few hundred.
const DefaultBudget = 200_000

// maxDepth bounds nested Template calls and Blueprint expansions.
const maxDepth = 100

// Unit is one parsed file.
type Unit struct {
	Path string
	File *syntax.File
}

// Options configure a run.
type Options struct {
	// Namespace is the namespace before a file's first Namespace
	// declaration.
	Namespace string
	// Budget is the step budget; 0 means DefaultBudget.
	Budget int
	// Cache, when set, remembers across runs which declarations need no
	// expansion, so an edit does not walk every unchanged declaration.
	Cache *Cache
	// Vanilla, when set, returns a file of the game's own resource
	// pack: Vanilla("models", "block/stone") is
	// assets/minecraft/models/block/stone.json. block_models reads it.
	Vanilla func(folder, path string) ([]byte, bool)
}

// Cache remembers which declarations pass through unexpanded. It is safe
// for concurrent use.
type Cache struct {
	mu   sync.Mutex
	key  string // the const, Template and Blueprint names it holds for
	pass map[*syntax.Decl]bool
}

// lookup returns whether d is known to pass through, under names key.
func (k *Cache) lookup(key string, d *syntax.Decl) (pass, ok bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.pass == nil || k.key != key {
		k.key, k.pass = key, map[*syntax.Decl]bool{}
	}
	pass, ok = k.pass[d]
	return pass, ok
}

func (k *Cache) store(key string, d *syntax.Decl, pass bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.key == key {
		k.pass[d] = pass
	}
}

// keep drops entries for declarations no longer in use.
func (k *Cache) keep(live map[*syntax.Decl]bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	for d := range k.pass {
		if !live[d] {
			delete(k.pass, d)
		}
	}
}

// Result is the expansion of a set of files.
//
// Consts, Templates, Blueprints and Types hold a public declaration by its
// name and a private one (`private const x`) by PrivateKey(file, name).
// Look a name up with Const, Template, Blueprint and Type, which know
// that a file's own private declaration hides a public one.
type Result struct {
	Files      map[string]*FileOut
	Consts     map[string]*Const
	Templates  map[string]*Template
	Blueprints map[string]*Blueprint
	Types      map[string]*TypeDef
	// Instances are the Blueprint instances, in the order they were
	// expanded.
	Instances []*Instance
	// Calls are the Template calls written in user code, whose
	// arguments the type checker checks against the parameters.
	Calls []*Call
	// Nodes holds the provenance of output nodes whose text was written
	// somewhere other than their span: inside a Template, a Blueprint or
	// a const.
	Nodes map[syntax.Node]Prov
	// Lines holds the provenance of the statements a Blueprint or
	// Template body wrote: each sits at the span of its use, and its Def
	// is the line in the body.
	Lines map[syntax.Node]Prov
}

// PrivateKey is the key a private const, Template, Blueprint or Type
// declared in file path has in a Result's maps.
func PrivateKey(path, name string) string { return name + "@" + path }

// Const returns the const that name means in file path, or nil.
func (r *Result) Const(path, name string) *Const {
	if k := r.Consts[PrivateKey(path, name)]; k != nil {
		return k
	}
	return r.Consts[name]
}

// Template returns the Template that name means in file path, or nil.
func (r *Result) Template(path, name string) *Template {
	if t := r.Templates[PrivateKey(path, name)]; t != nil {
		return t
	}
	return r.Templates[name]
}

// Blueprint returns the Blueprint that name means in file path, or nil.
func (r *Result) Blueprint(path, name string) *Blueprint {
	if b := r.Blueprints[PrivateKey(path, name)]; b != nil {
		return b
	}
	return r.Blueprints[name]
}

// Type returns the Type that name means in file path, or nil.
func (r *Result) Type(path, name string) *TypeDef {
	if t := r.Types[PrivateKey(path, name)]; t != nil {
		return t
	}
	return r.Types[name]
}

// TypeKey returns the key in Types of the Type that name means in file
// path: its private key when the file declares it privately, else name.
func (r *Result) TypeKey(path, name string) string {
	if r != nil && r.Types[PrivateKey(path, name)] != nil {
		return PrivateKey(path, name)
	}
	return name
}

// ConstsIn lists the consts file path can name, by name: its own
// private ones and the public ones they do not hide.
func (r *Result) ConstsIn(path string) []*Const {
	return visibleIn(r.Consts, path, func(k *Const) (bool, string, string) { return k.S.Private, k.Path, k.Name() })
}

// TemplatesIn lists the Templates file path can name, by name.
func (r *Result) TemplatesIn(path string) []*Template {
	return visibleIn(r.Templates, path, func(t *Template) (bool, string, string) { return t.D.Private, t.Path, t.D.Name.Text })
}

// BlueprintsIn lists the Blueprints file path can name, by name.
func (r *Result) BlueprintsIn(path string) []*Blueprint {
	return visibleIn(r.Blueprints, path, func(b *Blueprint) (bool, string, string) { return b.D.Private, b.Path, b.D.Name.Text })
}

// TypesIn lists the Types file path can name, by name.
func (r *Result) TypesIn(path string) []*TypeDef {
	return visibleIn(r.Types, path, func(t *TypeDef) (bool, string, string) { return t.D.Private, t.Path, t.D.Name.Text })
}

func visibleIn[T any](m map[string]T, path string, info func(T) (private bool, def, name string)) []T {
	names := map[string]T{}
	for _, v := range m {
		private, def, name := info(v)
		if !Visible(private, def, path) {
			continue
		}
		if _, taken := names[name]; taken && !private {
			continue // the file's private one hides it
		}
		names[name] = v
	}
	keys := make([]string, 0, len(names))
	for n := range names {
		keys = append(keys, n)
	}
	sort.Strings(keys)
	out := make([]T, len(keys))
	for i, n := range keys {
		out[i] = names[n]
	}
	return out
}

// Visible reports whether a declaration made in file def can be named
// from file path.
func Visible(private bool, def, path string) bool { return !private || def == path }

// FileOut is one file's expansion.
type FileOut struct {
	// Decls are the file's declarations after expansion, in order,
	// Namespace declarations included. Template, Blueprint and Type
	// declarations are not in it.
	Decls []*Decl
	Diags []syntax.Diagnostic
	// Missing are the Blueprint instances that leave out required fields,
	// by the span their mms/missing-field error is at, for the quick fix
	// that adds them.
	Missing map[syntax.Span][]string
	// Changed is false when Decls are exactly the file's own
	// declarations, so a caller can keep using the parsed file.
	Changed bool
}

// Decl is one declaration of the output.
type Decl struct {
	D    *syntax.Decl
	Prov Prov
	// From is the declaration the user wrote that this one came from:
	// D itself, the declaration in a top-level for, or a Blueprint
	// instance.
	From *syntax.Decl
	// Primary marks the declaration a Blueprint instance forwards to.
	Primary bool
	// Instance is the Blueprint instance that emitted it, if any.
	Instance *Instance
	// Emit is the declaration in the Blueprint's emit block it was
	// written from, when a Blueprint emitted it.
	Emit *syntax.Decl
}

// Const is a const declared at the top of a file.
type Const struct {
	Path  string
	S     *syntax.ConstStmt
	state int // 0 not evaluated, 1 evaluating, 2 done
	val   Value
}

// Template is a Template declaration.
type Template struct {
	Path string
	D    *syntax.Decl
}

// TypeDef is a Type declaration. The type checker reads it.
type TypeDef struct {
	Path string
	D    *syntax.Decl
}

// Blueprint is a Blueprint declaration, split into its parts.
type Blueprint struct {
	Path   string
	D      *syntax.Decl
	Fields []*syntax.FieldDecl
	Lets   []*syntax.LetStmt
	Checks []*syntax.Block
	As     []*syntax.AsStmt
	Emits  []*syntax.Block
}

// Instance is one declaration of a Blueprint's kind.
type Instance struct {
	Path      string
	D         *syntax.Decl
	Name      *syntax.Name
	Blueprint *Blueprint
	// Self holds id, namespace and path.
	Self *Record
	// Fields are the instance's field values, defaults filled in.
	Fields *Record
	// Primary is the emitted declaration the name forwards to, or nil.
	Primary *syntax.Decl
	// Written holds the field values the user wrote, evaluated and at
	// the user's spans, when the declaration is user code (not emitted
	// by another Blueprint).
	Written map[string]syntax.Expr
	// User is set when the declaration is user code.
	User bool
	// Conversions are the instance's `as` conversions, in order: what
	// its name means where each type is expected.
	Conversions []*Conversion
	// Forward is set instead when the primary declaration is itself a
	// Blueprint instance: the name means what that instance's name means.
	Forward *Instance
	// Consts are the consts at the top of the Blueprint's emit blocks,
	// which other code reads as name.const, and ConstDefs where each is
	// written in the Blueprint.
	Consts    *Record
	ConstDefs map[string]*syntax.ConstStmt
	fieldSp   map[string]syntax.Span
}

// Conversion is one `as T = expr hint(...)` of a Blueprint instance.
type Conversion struct {
	Type syntax.TypeExpr
	Path string // the Blueprint's file, where Type is written
	// X is the converted value, built at the instance's name.
	X syntax.Expr
	// Rule and Hint are the hint's rule name and message, if any.
	Rule, Hint string
}

// Call is a Template call in user code.
type Call struct {
	Path     string
	Template *Template
	X        *syntax.Call
	// Args maps each parameter given to its argument, evaluated and at
	// the user's spans.
	Args map[string]syntax.Expr
	// Runtime are the parameters given a function value (a variable, a
	// Global or a Storage field), which the function compiler checks.
	Runtime map[string]bool
	// Expanded are the statements a Commands Template wrote at the call.
	Expanded []syntax.Stmt
	// Result is the value a value Template gave, lowered, or nil.
	Result syntax.Expr
}

type evaluator struct {
	r      *Result
	budget int
	steps  int
	over   bool
	cache  *Cache
	key    string
	// emitted holds the declarations Blueprints emitted, by namespace
	// and text.
	emitted map[string]bool
	// values are the names of Globals and Storages, which function code
	// passes to a Template as the values they hold when the code runs.
	values map[string]bool
	// vanilla reads the game's resource pack files (Options.Vanilla).
	vanilla func(folder, path string) ([]byte, bool)
	// sites are the declarations written at the top of a file, by name,
	// and siteOf by the name as written: the ones of a Blueprint's kind
	// are instances whose emitted consts are read as name.const.
	sites  map[string][]*site
	siteOf map[*syntax.Name]*site
}

// site is one name of a declaration written at the top of a file. When
// it is a Blueprint instance it is expanded once, when it is first
// needed: in its place in the file, or earlier, when something reads one
// of its consts.
type site struct {
	path  string
	ns    string
	d     *syntax.Decl
	n     *syntax.Name
	state int // 0 not expanded, 1 expanding, 2 done
	out   []*Decl
	// inst is the instance, and insts the instances its expansion made,
	// held back when it was expanded early so that Result.Instances
	// keeps the order they are written in.
	inst   *Instance
	insts  []*Instance
	placed bool
}

// Run expands units, which must be in a stable order.
func Run(units []Unit, opt Options) *Result {
	e := &evaluator{r: &Result{Files: map[string]*FileOut{}, Consts: map[string]*Const{}, Templates: map[string]*Template{},
		Blueprints: map[string]*Blueprint{}, Types: map[string]*TypeDef{}, Nodes: map[syntax.Node]Prov{}, Lines: map[syntax.Node]Prov{}}, budget: opt.Budget, values: map[string]bool{}, vanilla: opt.Vanilla}
	if e.budget <= 0 {
		e.budget = DefaultBudget
	}
	for _, u := range units {
		e.r.Files[u.Path] = &FileOut{}
	}
	for _, u := range units {
		e.collect(u, opt.Namespace)
	}
	e.shadows()
	if e.cache = opt.Cache; e.cache != nil {
		var names []string
		for n := range e.r.Consts {
			names = append(names, "c"+n)
		}
		for n := range e.r.Templates {
			names = append(names, "t"+n)
		}
		for n := range e.r.Blueprints {
			names = append(names, "b"+n)
		}
		for n := range e.values {
			names = append(names, "v"+n)
		}
		for n := range e.sites {
			if e.instanceSites(n) != nil {
				names = append(names, "i"+n)
			}
		}
		sort.Strings(names)
		e.key = strings.Join(names, "\x00")
	}
	live := map[*syntax.Decl]bool{}
	for _, u := range units {
		for _, d := range u.File.Decls {
			live[d] = true
		}
		e.expand(u, opt.Namespace)
	}
	if e.cache != nil {
		e.cache.keep(live)
	}
	for _, f := range e.r.Files {
		sort.SliceStable(f.Diags, func(i, j int) bool { return f.Diags[i].Sp.Start < f.Diags[j].Sp.Start })
	}
	return e.r
}

// emittedBefore reports whether a Blueprint already emitted the same
// declaration in namespace ns, and records it.
func (e *evaluator) emittedBefore(ns string, d *syntax.Decl) bool {
	key := ns + "\x00" + format.File(&syntax.File{Items: []syntax.Stmt{d}, Decls: []*syntax.Decl{d}})
	if e.emitted[key] {
		return true
	}
	if e.emitted == nil {
		e.emitted = map[string]bool{}
	}
	e.emitted[key] = true
	return false
}

func (e *evaluator) report(path string, d syntax.Diagnostic) {
	f := e.r.Files[path]
	if f == nil {
		f = &FileOut{}
		e.r.Files[path] = f
	}
	for _, o := range f.Diags {
		if o.Sp == d.Sp && o.Message == d.Message && sameRelated(o.Related, d.Related) {
			return // the same problem, met again in another iteration
		}
	}
	f.Diags = append(f.Diags, d)
}

// sameRelated reports whether two problems point at the same places:
// one problem met on two lines of a Blueprint body is two problems.
func sameRelated(a, b []syntax.Related) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Path != b[i].Path || a[i].Sp != b[i].Sp {
			return false
		}
	}
	return true
}

// collect records the file's const, Template, Type and Blueprint
// declarations, which every file can use.
func (e *evaluator) collect(u Unit, ns string) {
	errorf := func(sp syntax.Span, msg string) {
		e.report(u.Path, syntax.Diagnostic{Code: "mms/duplicate", Severity: syntax.SeverityError, Message: msg, Sp: sp})
	}
	// keyOf is the key a declaration takes in its map. dup reports whether
	// that map already has it in this scope: a public name anywhere, or a
	// private name in this file, or the same name the other way in this
	// file.
	keyOf := func(private bool, name string) string { return key(private, u.Path, name) }
	dup := func(has func(string) (string, bool), private bool, name string) bool {
		if _, ok := has(keyOf(private, name)); ok {
			return true
		}
		if private {
			path, ok := has(name)
			return ok && path == u.Path
		}
		_, ok := has(PrivateKey(u.Path, name))
		return ok
	}
	constOrTemplate := func(k string) (string, bool) {
		if c := e.r.Consts[k]; c != nil {
			return c.Path, true
		}
		if t := e.r.Templates[k]; t != nil {
			return t.Path, true
		}
		return "", false
	}
	for _, it := range u.File.Items {
		switch it := it.(type) {
		case *syntax.ConstStmt:
			if it.Name == nil {
				continue
			}
			if dup(constOrTemplate, it.Private, it.Name.Text) {
				errorf(it.Name.Sp, "'"+it.Name.Text+"' is already declared")
				continue
			}
			e.r.Consts[keyOf(it.Private, it.Name.Text)] = &Const{Path: u.Path, S: it}
		case *syntax.Decl:
			if it.Kind == nil || it.Name == nil {
				continue
			}
			name := it.Name.Text
			if it.Kind.Name == "Namespace" {
				ns = name
			}
			if !definitionKind(it.Kind.Name) {
				e.addSites(u.Path, ns, it)
			}
			if it.Private && !definitionKind(it.Kind.Name) {
				e.report(u.Path, syntax.Diagnostic{Code: "mms/private", Severity: syntax.SeverityError, Sp: it.Kind.Sp,
					Message: "only a const, Template, Type or Blueprint can be private; a " + it.Kind.Name + " is part of the data pack, which every file shares"})
			}
			k := keyOf(it.Private, name)
			switch it.Kind.Name {
			case "Template":
				if dup(constOrTemplate, it.Private, name) {
					errorf(it.Name.Sp, "'"+name+"' is already declared")
					continue
				}
				e.r.Templates[k] = &Template{Path: u.Path, D: it}
			case "Global", "Storage":
				e.values[name] = true
			case "Type":
				if dup(func(k string) (string, bool) {
					t := e.r.Types[k]
					if t == nil {
						return "", false
					}
					return t.Path, true
				}, it.Private, name) {
					errorf(it.Name.Sp, "Type '"+name+"' is already declared")
					continue
				}
				e.r.Types[k] = &TypeDef{Path: u.Path, D: it}
			case "Blueprint":
				if it.Extern {
					continue // a kind the compiler builds in Go
				}
				if dup(func(k string) (string, bool) {
					b := e.r.Blueprints[k]
					if b == nil {
						return "", false
					}
					return b.Path, true
				}, it.Private, name) {
					errorf(it.Name.Sp, "Blueprint '"+name+"' is already declared")
					continue
				}
				e.r.Blueprints[k] = e.blueprint(u.Path, u.File.Src, it)
			}
		}
	}
}

// addSites records the plainly named declarations written at the top of
// a file, which may be Blueprint instances.
func (e *evaluator) addSites(path, ns string, d *syntax.Decl) {
	for _, n := range d.Names() {
		if n == nil || n.Templated() || n.Text == "" {
			continue
		}
		if e.sites == nil {
			e.sites, e.siteOf = map[string][]*site{}, map[*syntax.Name]*site{}
		}
		s := &site{path: path, ns: ns, d: d, n: n}
		e.sites[n.Text] = append(e.sites[n.Text], s)
		e.siteOf[n] = s
	}
}

// instanceSites are the Blueprint instances written at the top of a file
// with this name.
// tagID is the id an unqualified tag means, as a reference does: the
// project's own tag of that name in namespace ns when one is written,
// and otherwise vanilla's. A declaration's name would mean ns:path
// instead, so a tag passed in as a name piece keeps what it refers to.
func (e *evaluator) tagID(tag, ns string) string {
	if strings.Contains(tag, ":") {
		return tag
	}
	for _, s := range e.sites[tag] {
		if s.ns == ns && s.d.Kind != nil && strings.HasSuffix(s.d.Kind.Name, "Tag") {
			return ns + ":" + tag
		}
	}
	return "minecraft:" + tag
}

func (e *evaluator) instanceSites(name string) []*site {
	var out []*site
	for _, s := range e.sites[name] {
		if e.r.Blueprint(s.path, s.d.Kind.Name) != nil {
			out = append(out, s)
		}
	}
	return out
}

// shadows reports, as hints, the private declarations that hide a public
// one of the same name declared in another file.
func (e *evaluator) shadows() {
	// No related location: the project echoes a problem at each one, and
	// the public declaration has nothing wrong with it.
	hint := func(path string, sp syntax.Span, what, name, other string) {
		e.report(path, syntax.Diagnostic{Code: "mms/shadowed-private", Severity: syntax.SeverityHint, Sp: sp,
			Message: "this private " + what + " hides the " + what + " '" + name + "' declared in " + other + "; in this file '" + name + "' means this one"})
	}
	for _, k := range e.r.Consts {
		if k.S.Private {
			if o := e.r.Consts[k.Name()]; o != nil {
				hint(k.Path, k.S.Name.Sp, "const", k.Name(), o.Path)
			} else if o := e.r.Templates[k.Name()]; o != nil {
				hint(k.Path, k.S.Name.Sp, "const", k.Name(), o.Path)
			}
		}
	}
	for _, t := range e.r.Templates {
		if t.D.Private {
			if o := e.r.Templates[t.D.Name.Text]; o != nil {
				hint(t.Path, t.D.Name.Sp, "Template", t.D.Name.Text, o.Path)
			} else if o := e.r.Consts[t.D.Name.Text]; o != nil {
				hint(t.Path, t.D.Name.Sp, "Template", t.D.Name.Text, o.Path)
			}
		}
	}
	for _, t := range e.r.Types {
		if o := e.r.Types[t.D.Name.Text]; t.D.Private && o != nil {
			hint(t.Path, t.D.Name.Sp, "Type", t.D.Name.Text, o.Path)
		}
	}
	for _, b := range e.r.Blueprints {
		if o := e.r.Blueprints[b.D.Name.Text]; b.D.Private && o != nil {
			hint(b.Path, b.D.Name.Sp, "Blueprint", b.D.Name.Text, o.Path)
		}
	}
}

// blueprint splits a Blueprint body into its parts.
func (e *evaluator) blueprint(path string, src []rune, d *syntax.Decl) *Blueprint {
	bp := &Blueprint{Path: path, D: d}
	e.blueprintShape(path, src, d)
	if d.Body == nil {
		if d.Value != nil {
			e.report(path, syntax.Diagnostic{Code: "mms/blueprint", Severity: syntax.SeverityError,
				Message: "a Blueprint takes a { } body", Sp: d.Value.Span()})
		}
		return bp
	}
	for _, it := range d.Body.Items {
		switch it := it.(type) {
		case *syntax.FieldDecl:
			bp.Fields = append(bp.Fields, it)
		case *syntax.LetStmt:
			bp.Lets = append(bp.Lets, it)
		case *syntax.AsStmt:
			bp.As = append(bp.As, it)
			if id, ok := it.Value.(*syntax.Ident); ok && typeWords[id.Name] && it.Type != nil {
				t := strings.TrimSpace(string(src[it.Type.Span().Start:it.Type.Span().End]))
				e.report(path, syntax.Diagnostic{Code: "mms/blueprint", Severity: syntax.SeverityError, Sp: id.Sp,
					Message: "'" + id.Name + "' is a type, not a value. `as " + t + " = value` says what the declaration's name means where a " + t +
						" is expected. For a " + t + " that gives " + id.Name + ", emit it with the declaration's own name and mark it primary: primary " + t + " ${self.path}: " + id.Name + " { ... }"})
			}
		case *syntax.SectionStmt:
			switch it.Word.Name {
			case "check":
				bp.Checks = append(bp.Checks, it.Body)
			case "emit":
				bp.Emits = append(bp.Emits, it.Body)
			default:
				e.report(path, syntax.Diagnostic{Code: "mms/blueprint", Severity: syntax.SeverityError,
					Message: "a Blueprint has no '" + it.Word.Name + "' section", Sp: it.Word.Sp})
			}
		default:
			e.report(path, syntax.Diagnostic{Code: "mms/blueprint", Severity: syntax.SeverityError,
				Message: "a Blueprint body holds fields, let, check, as and emit", Sp: it.Span()})
		}
	}
	return bp
}

// typeWords are the value types, which an `as` value is not.
var typeWords = map[string]bool{"int": true, "float": true, "double": true, "bool": true, "string": true,
	"byte": true, "short": true, "long": true, "any": true}

// blueprintShape reports the parts of a Template a Blueprint does not
// have: parameters in ( ), which are its fields, and a result type.
// Both would otherwise be dropped without a word, and the body's use of
// a parameter would fail far from the cause.
func (e *evaluator) blueprintShape(path string, src []rune, d *syntax.Decl) {
	text := func(sp syntax.Span) string {
		if sp.Start < 0 || sp.End > len(src) || sp.Start > sp.End {
			return ""
		}
		return string(src[sp.Start:sp.End])
	}
	if d.Params != nil {
		var fields []string
		for _, p := range d.Params.Params {
			if t := strings.TrimSpace(text(p.Sp)); t != "" {
				fields = append(fields, t)
			}
		}
		msg := "a Blueprint has no parameters: its fields go in its body, one per line"
		if len(fields) > 0 {
			msg += ", as in " + fields[0]
		}
		dg := syntax.Diagnostic{Code: "mms/blueprint", Severity: syntax.SeverityError, Message: msg, Sp: d.Params.Sp}
		if b := d.Body; b != nil && len(fields) > 0 && b.Sp.Start < len(src) && src[b.Sp.Start] == '{' {
			indent := lineIndent(src, d.Sp.Start) + "  "
			ins := ""
			for _, f := range fields {
				ins += "\n" + indent + f
			}
			dg.Fixes = []syntax.Fix{{Title: "Move the parameters into the body as fields", Safe: true, Edits: []syntax.TextEdit{
				{Sp: d.Params.Sp},
				{Sp: syntax.Span{Start: b.Sp.Start + 1, End: b.Sp.Start + 1}, Text: ins},
			}}}
		}
		e.report(path, dg)
	}
	if d.Type != nil {
		t := strings.TrimSpace(text(d.Type.Span()))
		sp := d.Type.Span()
		// Take the ':' before the type along.
		for i := sp.Start - 1; i >= 0 && i < len(src); i-- {
			if src[i] == ':' {
				sp.Start = i
				break
			}
			if src[i] != ' ' && src[i] != '\t' {
				break
			}
		}
		e.report(path, syntax.Diagnostic{Code: "mms/blueprint", Severity: syntax.SeverityError, Sp: sp,
			Message: "a Blueprint has no result type. For its name to stand for a " + t + ", emit one with the declaration's own name and mark it primary: primary " + t + " ${self.path} { ... }",
			Fixes:   []syntax.Fix{{Title: "Remove the result type", Safe: true, Edits: []syntax.TextEdit{{Sp: sp}}}}})
	}
}

// lineIndent is the indentation of the line holding off.
func lineIndent(src []rune, off int) string {
	start := off
	for start > 0 && src[start-1] != '\n' {
		start--
	}
	end := start
	for end < len(src) && (src[end] == ' ' || src[end] == '\t') {
		end++
	}
	return string(src[start:end])
}

// definitionKind reports whether a kind declares something this package
// keeps for itself rather than output.
func definitionKind(k string) bool { return k == "Template" || k == "Blueprint" || k == "Type" }

// expand writes one file's output declarations.
func (e *evaluator) expand(u Unit, ns string) {
	out := e.r.Files[u.Path]
	same := true
	for _, it := range u.File.Items {
		// Each item has a budget of its own, so one runaway Template
		// does not stop the rest of the project.
		e.steps, e.over = 0, false
		c := &ctx{e: e, path: u.Path, defPath: u.Path, sc: &scope{}, ns: ns, compile: true}
		switch it := it.(type) {
		case *syntax.Decl:
			if it.Kind == nil {
				continue
			}
			if it.Kind.Name == "Namespace" && it.Name != nil {
				ns = it.Name.Text
			}
			if definitionKind(it.Kind.Name) {
				same = false
				if it.Kind.Name == "Template" && it.Extern && intrinsics[it.NameText()] == nil && !kept[it.NameText()] {
					c.errorf("mms/unknown-intrinsic", it.Name.Sp, "there is no intrinsic named '"+it.NameText()+"'")
				}
				continue
			}
			if !e.touches(u.Path, it) {
				out.Decls = append(out.Decls, &Decl{D: it, From: it, Prov: Prov{At: Origin{u.Path, it.Sp}}})
				continue
			}
			same = false
			out.Decls = append(out.Decls, c.decl(it, it)...)
		case *syntax.ConstStmt:
			// Evaluate it even when unused, so its problems show.
			if it.Name == nil {
				continue
			}
			if k := e.r.Const(u.Path, it.Name.Text); k != nil && k.S == it {
				c.constValue(k, it.Sp)
			}
		case *syntax.LetStmt:
			c.errorf("mms/top-level-let", it.Sp, "a value at the top of a file is a const: write 'const' instead of 'let'")
		case *syntax.ForStmt:
			same = false
			var emitted []*Decl
			c.emit = &emitted
			c.forStmt(it, nil)
			out.Decls = append(out.Decls, emitted...)
		}
	}
	out.Changed = !same
}

// Name is the const's name.
func (s *Const) Name() string { return s.S.Name.Text }

// Key is the const's key in Result.Consts.
func (s *Const) Key() string { return key(s.S.Private, s.Path, s.S.Name.Text) }

// Key is the Template's key in Result.Templates.
func (t *Template) Key() string { return key(t.D.Private, t.Path, t.D.Name.Text) }

// Key is the Blueprint's key in Result.Blueprints.
func (b *Blueprint) Key() string { return key(b.D.Private, b.Path, b.D.Name.Text) }

// Key is the Type's key in Result.Types.
func (t *TypeDef) Key() string { return key(t.D.Private, t.Path, t.D.Name.Text) }

func key(private bool, path, name string) string {
	if private {
		return PrivateKey(path, name)
	}
	return name
}

// touches reports whether a declaration uses anything this package
// expands. A declaration that does not is passed through as it is, so the
// checker can keep its earlier results.
func (e *evaluator) touches(path string, d *syntax.Decl) bool {
	if e.r.Blueprint(path, d.Kind.Name) != nil {
		return true
	}
	if e.cache == nil {
		return e.walkTouches(path, d)
	}
	if pass, ok := e.cache.lookup(e.key, d); ok {
		return !pass
	}
	t := e.walkTouches(path, d)
	e.cache.store(e.key, d, !t)
	return t
}

func (e *evaluator) walkTouches(path string, d *syntax.Decl) bool {
	for _, n := range d.Names() {
		if n.Templated() {
			return true
		}
	}
	found := false
	visit := func(n syntax.Node) bool {
		if found {
			return false
		}
		switch n := n.(type) {
		case *syntax.Ident:
			if e.r.Const(path, n.Name) != nil || e.r.Template(path, n.Name) != nil {
				found = true
			}
		case *syntax.Member:
			if id, ok := n.X.(*syntax.Ident); ok && e.instanceSites(id.Name) != nil {
				found = true
			}
		case *syntax.ConstStmt, *syntax.Comprehension, *syntax.Spread, *syntax.IfExpr, *syntax.TemplatedName:
			found = true
		case *syntax.Binary:
			if n.Op == "??" {
				found = true
			}
		case *syntax.ForStmt:
			if _, rng := n.Iter.(*syntax.Range); !rng {
				found = true
			}
		case *syntax.String:
			found = e.interpTouches(path, n.Value)
		case *syntax.CommandStmt:
			found = e.interpTouches(path, n.Text)
		}
		return !found
	}
	syntax.Walk(d, visit)
	return found
}

// interpTouches reports whether a `${...}` in s names a const or Template.
func (e *evaluator) interpTouches(path, s string) bool {
	if !strings.Contains(s, "${") {
		return false
	}
	for _, p := range pieces(s) {
		if p.expr == "" {
			continue
		}
		x, diags := syntax.ParseExpr(p.expr, 0)
		if len(diags) > 0 {
			continue
		}
		hit := false
		syntax.Walk(x, func(n syntax.Node) bool {
			if id, ok := n.(*syntax.Ident); ok && (e.r.Const(path, id.Name) != nil || e.r.Template(path, id.Name) != nil) {
				hit = true
			}
			if m, ok := n.(*syntax.Member); ok {
				if id, ok := m.X.(*syntax.Ident); ok && e.instanceSites(id.Name) != nil {
					hit = true
				}
			}
			return !hit
		})
		if hit {
			return true
		}
	}
	return false
}

// piece is a run of literal text, or the inside of one `${...}`.
type piece struct {
	lit  string
	expr string
	off  int // offset of the piece in the string
}

// pieces splits s at its `${...}` pieces. An unclosed `${` stays text.
func pieces(s string) []piece {
	var out []piece
	lit := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '$' || i+1 >= len(s) || s[i+1] != '{' {
			continue
		}
		depth, j := 0, i+1
		for ; j < len(s); j++ {
			if s[j] == '{' {
				depth++
			} else if s[j] == '}' {
				if depth--; depth == 0 {
					break
				}
			}
		}
		if j >= len(s) {
			break
		}
		if i > lit {
			out = append(out, piece{lit: s[lit:i], off: lit})
		}
		out = append(out, piece{expr: s[i+2 : j], off: i})
		lit = j + 1
		i = j
	}
	if lit < len(s) {
		out = append(out, piece{lit: s[lit:], off: lit})
	}
	return out
}

// Lowered returns the const's value written back as an expression at
// the spans of its own text, or nil when it has none or was not
// evaluated.
func (s *Const) Lowered() syntax.Expr {
	if s.state != 2 || s.val == nil {
		return nil
	}
	c := &ctx{e: &evaluator{r: &Result{Nodes: map[syntax.Node]Prov{}, Lines: map[syntax.Node]Prov{}}}, path: s.Path, defPath: s.Path, sc: &scope{}}
	return c.lower(s.val)
}
