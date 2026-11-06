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
}

// FileOut is one file's expansion.
type FileOut struct {
	// Decls are the file's declarations after expansion, in order,
	// Namespace declarations included. Template, Blueprint and Type
	// declarations are not in it.
	Decls []*Decl
	Diags []syntax.Diagnostic
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
	fieldSp map[string]syntax.Span
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
}

// Run expands units, which must be in a stable order.
func Run(units []Unit, opt Options) *Result {
	e := &evaluator{r: &Result{Files: map[string]*FileOut{}, Consts: map[string]*Const{}, Templates: map[string]*Template{},
		Blueprints: map[string]*Blueprint{}, Types: map[string]*TypeDef{}, Nodes: map[syntax.Node]Prov{}}, budget: opt.Budget}
	if e.budget <= 0 {
		e.budget = DefaultBudget
	}
	for _, u := range units {
		e.r.Files[u.Path] = &FileOut{}
	}
	for _, u := range units {
		e.collect(u)
	}
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
		if o.Sp == d.Sp && o.Message == d.Message {
			return // the same problem, met again in another iteration
		}
	}
	f.Diags = append(f.Diags, d)
}

// collect records the file's const, Template, Type and Blueprint
// declarations, which every file can use.
func (e *evaluator) collect(u Unit) {
	errorf := func(sp syntax.Span, msg string) {
		e.report(u.Path, syntax.Diagnostic{Code: "mms/duplicate", Severity: syntax.SeverityError, Message: msg, Sp: sp})
	}
	taken := func(name string) bool {
		_, c := e.r.Consts[name]
		_, t := e.r.Templates[name]
		return c || t
	}
	for _, it := range u.File.Items {
		switch it := it.(type) {
		case *syntax.ConstStmt:
			if it.Name == nil {
				continue
			}
			if taken(it.Name.Text) {
				errorf(it.Name.Sp, "'"+it.Name.Text+"' is already declared")
				continue
			}
			e.r.Consts[it.Name.Text] = &Const{Path: u.Path, S: it}
		case *syntax.Decl:
			if it.Kind == nil || it.Name == nil {
				continue
			}
			name := it.Name.Text
			switch it.Kind.Name {
			case "Template":
				if taken(name) {
					errorf(it.Name.Sp, "'"+name+"' is already declared")
					continue
				}
				e.r.Templates[name] = &Template{Path: u.Path, D: it}
			case "Type":
				if _, dup := e.r.Types[name]; dup {
					errorf(it.Name.Sp, "Type '"+name+"' is already declared")
					continue
				}
				e.r.Types[name] = &TypeDef{Path: u.Path, D: it}
			case "Blueprint":
				if it.Extern {
					continue // a kind the compiler builds in Go
				}
				if _, dup := e.r.Blueprints[name]; dup {
					errorf(it.Name.Sp, "Blueprint '"+name+"' is already declared")
					continue
				}
				e.r.Blueprints[name] = e.blueprint(u.Path, it)
			}
		}
	}
}

// blueprint splits a Blueprint body into its parts.
func (e *evaluator) blueprint(path string, d *syntax.Decl) *Blueprint {
	bp := &Blueprint{Path: path, D: d}
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

// definitionKind reports whether a kind declares something this package
// keeps for itself rather than output.
func definitionKind(k string) bool { return k == "Template" || k == "Blueprint" || k == "Type" }

// fnKinds are the declarations whose body is function code.
var fnKinds = map[string]bool{"Function": true, "OnLoad": true, "OnTick": true, "OnFirstJoin": true,
	"OnRespawn": true, "OnAdvancement": true, "Test": true}

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
				if it.Kind.Name == "Template" && it.Extern && intrinsics[it.NameText()] == nil {
					c.errorf("mms/unknown-intrinsic", it.Name.Sp, "there is no intrinsic named '"+it.NameText()+"'")
				}
				continue
			}
			if !e.touches(it) {
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
			if k := e.r.Consts[it.Name.Text]; k != nil && k.S == it {
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

// touches reports whether a declaration uses anything this package
// expands. A declaration that does not is passed through as it is, so the
// checker can keep its earlier results.
func (e *evaluator) touches(d *syntax.Decl) bool {
	if e.r.Blueprints[d.Kind.Name] != nil {
		return true
	}
	if e.cache == nil {
		return e.walkTouches(d)
	}
	if pass, ok := e.cache.lookup(e.key, d); ok {
		return !pass
	}
	t := e.walkTouches(d)
	e.cache.store(e.key, d, !t)
	return t
}

func (e *evaluator) walkTouches(d *syntax.Decl) bool {
	for _, n := range d.Names() {
		if n.Templated() {
			return true
		}
	}
	found := false
	var visit func(n syntax.Node) bool
	visit = func(n syntax.Node) bool {
		if found {
			return false
		}
		switch n := n.(type) {
		case *syntax.Ident:
			if e.r.Consts[n.Name] != nil || e.r.Templates[n.Name] != nil {
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
			found = e.interpTouches(n.Value)
		case *syntax.CommandStmt:
			found = e.interpTouches(n.Text)
		}
		return !found
	}
	syntax.Walk(d, visit)
	return found
}

// interpTouches reports whether a `${...}` in s names a const or Template.
func (e *evaluator) interpTouches(s string) bool {
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
			if id, ok := n.(*syntax.Ident); ok && (e.r.Consts[id.Name] != nil || e.r.Templates[id.Name] != nil) {
				hit = true
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
	c := &ctx{e: &evaluator{r: &Result{Nodes: map[syntax.Node]Prov{}}}, path: s.Path, defPath: s.Path, sc: &scope{}}
	return c.lower(s.val)
}
