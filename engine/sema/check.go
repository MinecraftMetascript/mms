package sema

import (
	"fmt"
	"io/fs"
	"math"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/internal/near"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Info is what the checker learned about one node, for editor features.
type Info struct {
	Type     tview   // the type the position expects
	Symbol   *Symbol // the declaration an id resolved to
	ID       string  // the id an identifier resolved to
	Field    *field  // the field a key or named argument names
	Fields   []field // for object bodies and calls: every field in scope
	Disp     *dispInfo
	Callable *Callable
	Local    *Local // a function's parameter or local variable
	// Help is Markdown for a language helper such as an execute block's
	// As(...), shown on hover.
	Help string
	// Bare says how a string literal here could be written unquoted:
	// BareID for an id (Folder is its registry folder, "tags/..." for a
	// tag), BareKey for a dispatch key, BareEnum for a string enum name.
	Bare   BareKind
	Folder string
	// Asset is set on an asset helper call such as Texture("a.png").
	Asset *Asset
	// ModelTexture is set on a texture id or Texture(...) call in a
	// model, whose sprite the game looks up in the blocks and items
	// atlases.
	ModelTexture bool
	// Converted is set where a Blueprint instance's name is checked as
	// one of its `as` conversions.
	Converted *Conversion
	// Markup is set on a string read as text markup, and Accidental holds
	// the style marks in it read as the characters themselves, such as
	// the * of "5 * 3", for the accidental-markup lint.
	Markup     bool
	Accidental []syntax.Span
	// Props are the block state properties a multipart `when` object
	// body can test, for completion.
	Props map[string][]string
	// InlineFn is set on a `run { ... }` block written where a function id is
	// needed: its body is function code, compiled as a hoisted Function.
	InlineFn bool
	// Extract is the declaration kind (as written, e.g. "Noise") of an
	// inline value that could be its own declaration referenced by id
	// here: one that is hoisted, or one at a position that also takes an
	// id of that kind. Editors offer to extract it.
	Extract string

	typed bool
}

// BareKind classifies positions where a string literal has a bare form.
type BareKind uint8

const (
	BareNone BareKind = iota
	BareID
	BareKey
	BareEnum
)

// Callable is a resolved call: which struct (and dispatch case) it builds.
type Callable struct {
	Name   string // as written
	Case   string // dispatch key, "" for a plain struct
	Struct tview
	Fields []field
	Disp   *dispInfo
	Ov     *overlay
}

// Checker checks expressions against expected types and lowers them to
// JSON in the same walk. One Checker serves one declaration.
type Checker struct {
	// Typed are values the editor offers wherever their type fits, such
	// as the project's Templates by their result type: Piece where a
	// WeightedElement goes.
	Typed []TypedCandidate
	Pack  *spec.Pack
	Kinds *Kinds
	Syms  *Table
	NS    string
	Diags []syntax.Diagnostic
	// Info is filled unless the checker is running a trial for union
	// narrowing.
	Info map[syntax.Node]*Info
	// Deps records every symbol lookup and its outcome, so a cached
	// result can be reused while the outcomes stay the same.
	Deps map[string]bool
	// Files are the project's files, read by asset helpers; nil when
	// the host has none.
	Files fs.FS
	// Src is the checked file's text; quick fixes need it for layout.
	Src []rune
	// Texts returns the text of a project file the host has as a
	// source, such as a .mmsbook file, by its path in the project, and
	// the unit path it has; nil when there are none.
	Texts func(path string) (unit, text string, ok bool)
	// Book is the layout of the last Book text checked, for the editor
	// and the book-page-overflow lint.
	Book *BookLayout
	// Shown are the messages the function's text commands show.
	Shown []ShownText
	// BookDiags are the problems in the Book text checked.
	BookDiags []syntax.Diagnostic
	// Extern holds problems found in other files, such as a .mmsbook
	// file a Book reads, by path.
	Extern map[string][]syntax.Diagnostic

	// Owner is the id of the declaration being checked; inline values
	// hoisted out of it are named after it.
	Owner string
	// Hoisted are inline values written where an id is required, each
	// to be written as its own resource (see hoist.go).
	Hoisted []Hoisted
	hoistN  map[string]int

	// resource is the resource kind of the declaration being checked,
	// such as "model".
	resource string

	trial bool
	errs  int
	depth int
	memo  map[memoKey]int
	// merging holds the objects mergeConvert is checking on their own.
	merging map[syntax.Node]bool
	// converting holds the instances whose conversions are being
	// checked, to catch one that contains itself.
	converting map[string]bool
	parent     *Checker
}

type memoKey struct {
	n syntax.Node
	t *spec.Type
	e *env
}

// NewChecker returns a checker for declarations in namespace ns.
func NewChecker(p *spec.Pack, kinds *Kinds, syms *Table, ns string) *Checker {
	return &Checker{Pack: p, Kinds: kinds, Syms: syms, NS: ns, Info: map[syntax.Node]*Info{}, Deps: map[string]bool{}, memo: map[memoKey]int{}}
}

// find looks up a symbol, recording the lookup in Deps.
func (c *Checker) find(folder, id string) *Symbol {
	s := c.Syms.Find(folder, id)
	if c.Deps != nil {
		c.Deps["f|"+folder+"|"+id] = s != nil
	}
	return s
}

// otherKind looks up a declaration of id in a folder other than folder,
// recording the lookup for incremental checking.
func (c *Checker) otherKind(folder, id string) *Symbol {
	s := c.Syms.OtherKind(folder, id)
	if c.Deps != nil {
		c.Deps["o|"+folder+"|"+id+"|"+otherFolder(s)] = true
	}
	return s
}

func otherFolder(s *Symbol) string {
	if s == nil {
		return ""
	}
	return s.Kind.Folder()
}

// DepsHold reports whether every lookup in deps has the same outcome in t.
func DepsHold(t *Table, deps map[string]bool) bool {
	for k, found := range deps {
		var now bool
		if strings.HasPrefix(k, "a|") {
			continue // asset files: see AssetDepsHold
		} else if ns, ok := strings.CutPrefix(k, "n|"); ok {
			now = t.hasNamespace(ns)
		} else if rest, ok := strings.CutPrefix(k, "o|"); ok {
			// o|folder|id|other: the other-kind declaration is in that folder.
			parts := strings.SplitN(rest, "|", 3)
			now = len(parts) == 3 && otherFolder(t.OtherKind(parts[0], parts[1])) == parts[2]
		} else if rest, ok := strings.CutPrefix(k, "c|"); ok {
			// c|id|signature: the instance's conversions are unchanged.
			i := strings.IndexByte(rest, '|')
			now = ConversionsSig(t.Conversions[rest[:i]]) == rest[i+1:]
		} else if rest, ok := strings.CutPrefix(k, "s|"); ok {
			// s|folder|id|signature: the symbol still has that signature.
			parts := strings.SplitN(rest, "|", 3)
			now = len(parts) == 3 && symSig(t.Find(parts[0], parts[1])) == parts[2]
		} else {
			rest := k[2:]
			i := strings.IndexByte(rest, '|')
			now = t.Find(rest[:i], rest[i+1:]) != nil
		}
		if now != found {
			return false
		}
	}
	return true
}

func (c *Checker) errorf(code string, sp syntax.Span, format string, args ...any) {
	c.errs++
	if c.trial {
		return
	}
	c.Diags = append(c.Diags, syntax.Diagnostic{Code: code, Severity: syntax.SeverityError, Message: fmt.Sprintf(format, args...), Sp: sp})
}

// respelled reports a helper written with a spelling an alias replaces.
func (c *Checker) respelled(name string, sp syntax.Span) bool {
	want, ok := Spelled(name)
	if !ok {
		return false
	}
	c.errs++
	if !c.trial {
		c.Diags = append(c.Diags, respell("mms/unknown-helper", sp, name, want))
	}
	return true
}

func (c *Checker) warnf(code string, sp syntax.Span, format string, args ...any) {
	if c.trial {
		return
	}
	c.Diags = append(c.Diags, syntax.Diagnostic{Code: code, Severity: syntax.SeverityWarning, Message: fmt.Sprintf(format, args...), Sp: sp})
}

func (c *Checker) info(n syntax.Node) *Info {
	if c.trial || n == nil {
		return &Info{}
	}
	in := c.Info[n]
	if in == nil {
		in = &Info{}
		c.Info[n] = in
	}
	return in
}

// Resource checks a declaration's value against a resource kind's root
// type and returns its JSON.
func (c *Checker) Resource(d *syntax.Decl, k Kind) jsonv.Value {
	root := tview{T: c.Pack.Resources[k.Resource]}
	if root.T != nil && root.T.Kind == "reference" {
		root.Name = root.T.Path
	}
	c.resource = k.Resource
	defer func() { c.resource = "" }()
	var v jsonv.Value
	switch {
	case d.Body != nil:
		obj := &syntax.Object{Body: d.Body, Sp: d.Body.Sp}
		v = c.check(obj, root)
		// The body's object is not in the tree; its Info is kept on the
		// declaration, for lints that read the body as a value.
		if in := c.Info[obj]; in != nil {
			c.Info[d] = in
		}
		if k.Resource == "block_definition" {
			c.blockState(d, d.Body)
		}
		return v
	case d.Value != nil:
		v = c.check(d.Value, root)
		if k.Resource == "block_definition" {
			c.blockState(d, objectBody(d.Value))
		}
		return v
	}
	c.errorf("mms/missing-value", d.Kind.Sp, "%s needs a value: `= ...` or `{ ... }`", d.Kind.Name)
	return nil
}

// check is the core walk: e at type v.
func (c *Checker) check(e syntax.Expr, v tview) jsonv.Value {
	if e == nil {
		return nil
	}
	c.depth++
	defer func() { c.depth-- }()
	if c.depth > 200 {
		c.errorf("mms/too-deep", e.Span(), "expression nests too deeply")
		return nil
	}
	v = c.norm(v)
	in := c.info(e)
	// Keep the outermost expected type: completion should offer every
	// member of a union, not just the one narrowing picked.
	if !in.typed {
		in.Type, in.typed = v, true
	}
	if p, ok := e.(*syntax.Paren); ok {
		return c.check(p.X, v)
	}
	if _, ok := e.(*syntax.Bad); ok {
		return nil // already reported by the parser
	}
	if r, ok := c.convert(e, v); ok {
		return r
	}
	if r, ok := c.mergeConvert(e, v); ok {
		return r
	}
	if r, ok := c.formConvert(e, v); ok {
		return r
	}
	if r := c.shorthand(e, v); r != nil {
		return r.apply(c, e, v)
	}
	if v.T == nil {
		return c.anyValue(e)
	}
	switch v.T.Kind {
	case "union":
		return c.union(e, v)
	case "struct":
		return c.structValue(e, v)
	case "list":
		return c.listValue(e, v)
	case "tuple":
		return c.tupleValue(e, v)
	case "byte_array", "int_array", "long_array":
		elem := strings.TrimSuffix(v.T.Kind, "_array")
		return c.listValue(e, tview{T: &spec.Type{Kind: "list", Item: &spec.Type{Kind: elem}}})
	case "string":
		return c.stringValue(e, v)
	case "boolean":
		if b, ok := e.(*syntax.Bool); ok {
			return jsonv.Bool(b.Value)
		}
		// NBT stores booleans as bytes, and data packs often write 0 or 1.
		if n, ok := e.(*syntax.Number); ok && (n.Text == "0" || n.Text == "1") {
			return jsonv.Number(n.Text)
		}
	case "byte", "short", "int", "long", "float", "double":
		return c.number(e, v)
	case "literal":
		return c.literal(e, v)
	case "enum":
		return c.enumValue(e, v)
	case "any":
		return c.anyValue(e)
	default:
		// Dynamic dispatchers and other types MMS cannot resolve
		// statically are passed through unchecked.
		return c.anyValue(e)
	}
	c.mismatch(e, v)
	return nil
}

func (c *Checker) mismatch(e syntax.Expr, v tview) {
	c.errorf("mms/type-mismatch", e.Span(), "expected %s, got %s%s", describe(v), exprKind(e), colorHint(v))
}

// colorHint is how to write a color, appended to an error at a color
// type; "" elsewhere.
func colorHint(v tview) string {
	if ct, ok := colorOf(v); ok {
		return ". " + strings.ReplaceAll(ct.help(), "`", "")
	}
	return ""
}

func describe(v tview) string {
	if ct, ok := colorOf(v); ok {
		return ct.name()
	}
	if v.Name != "" {
		return spec.ShortPath(v.Name)
	}
	if v.T == nil {
		return "any value"
	}
	switch v.T.Kind {
	case "string":
		if isSelectorType(v.T) {
			return "a selector, such as @e[tag=boss]"
		}
		if id := idOf(v.T); id.Registry != "" {
			return "a " + id.Registry + " id"
		}
		return "a string"
	case "boolean":
		return "true or false"
	case "int", "long", "short", "byte":
		return "an integer" + rangeText(v.T.ValueRange)
	case "float", "double":
		return "a number" + rangeText(v.T.ValueRange)
	case "list":
		return "a list"
	case "struct":
		return "an object"
	}
	return spec.Format(v.T, 0)
}

func rangeText(r *spec.Range) string {
	if r == nil {
		return ""
	}
	return " in " + strings.TrimPrefix(spec.Format(&spec.Type{Kind: "int", ValueRange: r}, 0), "int @ ")
}

// lengthText is a list length range in words: "4", "1 to 3", "at least 1".
func lengthText(r *spec.Range) string {
	if r.Min != nil && r.Max != nil && *r.Min == *r.Max && r.Kind == 0 {
		return strconv.FormatFloat(*r.Min, 'f', -1, 64)
	}
	return strings.TrimPrefix(rangeText(r), " in ")
}

func exprKind(e syntax.Expr) string {
	switch e := e.(type) {
	case *syntax.Number:
		return "number " + e.Text
	case *syntax.String:
		return "string " + e.Raw
	case *syntax.Bool:
		return strconv.FormatBool(e.Value)
	case *syntax.Ident:
		return "'" + e.Name + "'"
	case *syntax.ResLoc:
		return "'" + e.Text + "'"
	case *syntax.Tag:
		return "tag #" + e.Name
	case *syntax.List:
		return "a list"
	case *syntax.Object:
		return "an object"
	case *syntax.Call, *syntax.Trailing:
		if n := callName(e); n != "" {
			return n + "(...)"
		}
		return "a call"
	case *syntax.Binary:
		return "'" + e.Op + "' expression"
	case *syntax.Unary:
		return "'" + e.Op + "' expression"
	case *syntax.Range:
		return "a range"
	case *syntax.Member:
		return "a member access"
	case *syntax.Paren:
		return exprKind(e.X)
	}
	return "an expression"
}

// union narrows a union by the expression's shape, then by trial checks,
// and reports against the closest member when nothing fits.
func (c *Checker) union(e syntax.Expr, v tview) jsonv.Value {
	if r, ok := c.inlineSound(e, v); ok {
		return r
	}
	if id, ok := c.keyRequiredID(v); ok {
		if r, ok := c.hoist(e, id); ok {
			return r
		}
	}
	c.markExtract(e, v)
	var cands []tview
	for _, m := range v.T.Members {
		mv := c.norm(tview{T: m, Env: v.Env})
		if c.fits(e, mv) {
			cands = append(cands, mv)
		}
	}
	cands = c.preferInline(e, cands)
	cands = c.preferKnownKeys(e, cands)
	switch len(cands) {
	case 0:
		if name := callName(e); name != "" {
			_, id, _, _ := unchain(e)
			if c.respelled(name, id.Sp) {
				return nil
			}
			names := c.callableNames(v)
			msg := "'" + name + "' does not make a " + unionText(v)
			if s := closest(name, names); s != "" {
				msg += "; did you mean " + s + "(...)?"
			}
			c.errorf("mms/unknown-helper", id.Sp, "%s", msg)
			return nil
		}
		c.errorf("mms/type-mismatch", e.Span(), "expected %s, got %s%s", unionText(v), exprKind(e), colorHint(v))
		return nil
	case 1:
		return c.checkMember(e, v, cands[0])
	}
	best, bestErrs := 0, math.MaxInt
	for i, m := range cands {
		n := c.trialErrors(e, m)
		if n < bestErrs {
			best, bestErrs = i, n
		}
		if n == 0 {
			break
		}
	}
	return c.checkMember(e, v, cands[best])
}

// checkMember checks e as m, the member of union v it was narrowed to.
// A list of ids where the union also takes a tag gets, on the tags in
// it, the fix that moves the list into a tag declaration.
func (c *Checker) checkMember(e syntax.Expr, v, m tview) jsonv.Value {
	kind := ""
	if !c.trial {
		kind = c.tagListKind(e, v, m)
	}
	if kind == "" {
		return c.check(e, m)
	}
	from := len(c.Diags)
	r := c.check(e, m)
	c.tagListFix(e.(*syntax.List), kind, from)
	return r
}

// preferKnownKeys keeps, of the struct members an object might be, those
// that know every key it writes, so a half-written entry still narrows to
// the member its other keys name: `{ modifier = overlay, argument = ... }`
// is a modifier however broken the argument is.
func (c *Checker) preferKnownKeys(e syntax.Expr, cands []tview) []tview {
	o, ok := e.(*syntax.Object)
	if !ok || o.Body == nil || len(cands) < 2 {
		return cands
	}
	var keys []string
	for _, it := range o.Body.Items {
		if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil {
			keys = append(keys, f.Key.Text)
		}
	}
	if len(keys) == 0 {
		return cands
	}
	var keep []tview
	for _, m := range cands {
		if m.T == nil || m.T.Kind != "struct" {
			return cands
		}
		fs, disp := c.fields(m, "")
		if disp != nil {
			return cands
		}
		known := map[string]bool{}
		for _, f := range fs {
			if f.KeyType != nil {
				return cands // a map takes any key
			}
			known[f.Key] = true
		}
		all := true
		for _, k := range keys {
			if !known[k] {
				all = false
			}
		}
		if all {
			keep = append(keep, m)
		}
	}
	if len(keep) == 0 {
		return cands
	}
	return keep
}

func unionText(v tview) string {
	if ct, ok := colorOf(v); ok {
		return ct.name()
	}
	if v.Name != "" {
		return spec.ShortPath(v.Name)
	}
	var parts []string
	for _, m := range v.T.Members {
		parts = append(parts, describe(tview{T: m}))
	}
	return strings.Join(parts, " or ")
}

// trialErrors checks e at v without recording anything and returns the
// number of errors. Results are memoised per node and type.
func (c *Checker) trialErrors(e syntax.Expr, v tview) int {
	k := memoKey{e, v.T, v.Env}
	root := c
	for root.parent != nil {
		root = root.parent
	}
	if n, ok := root.memo[k]; ok {
		return n
	}
	t := &Checker{Pack: c.Pack, Kinds: c.Kinds, Syms: c.Syms, NS: c.NS, trial: true, depth: c.depth, parent: root, Deps: c.Deps, Files: c.Files, converting: c.converting}
	t.check(e, v)
	root.memo[k] = t.errs
	return t.errs
}

// fits is a cheap shape test: could e possibly be a value of v?
func (c *Checker) fits(e syntax.Expr, v tview) bool {
	if p, ok := e.(*syntax.Paren); ok {
		return c.fits(p.X, v)
	}
	if _, ok := e.(*syntax.Bad); ok {
		return true
	}
	v = c.norm(v)
	if r := c.shorthand(e, v); r != nil || c.formConversion(e, v) != nil {
		return true
	}
	if v.T == nil {
		return true
	}
	switch v.T.Kind {
	case "union":
		for _, m := range v.T.Members {
			if c.fits(e, tview{T: m, Env: v.Env}) {
				return true
			}
		}
		return false
	case "any":
		return true
	case "struct":
		switch e.(type) {
		case *syntax.Object:
			return true
		case *syntax.Call, *syntax.Trailing:
			return c.callable(e, v) != nil
		}
		return false
	case "list", "tuple", "byte_array", "int_array", "long_array":
		_, ok := e.(*syntax.List)
		return ok
	case "string":
		switch e.(type) {
		case *syntax.String:
			return true
		case *syntax.Ident:
			id := idOf(v.T)
			return id.Present && id.Tags != "required"
		case *syntax.ResLoc:
			// A qualified resource location is also a plain string.
			id := idOf(v.T)
			return (id.Present && id.Tags != "required") || (!id.Present && qualified(e))
		case *syntax.Tag:
			id := idOf(v.T)
			return id.Present && id.Tags != ""
		case *syntax.Call:
			return c.assetFits(e, idOf(v.T)) || c.hoistFits(e, idOf(v.T))
		}
		return c.hoistFits(e, idOf(v.T))
	case "boolean":
		if n, ok := e.(*syntax.Number); ok {
			return n.Text == "0" || n.Text == "1"
		}
		_, ok := e.(*syntax.Bool)
		return ok
	case "byte", "short", "int", "long", "float", "double":
		return isNumber(e)
	case "literal":
		if v.T.Literal == nil {
			return false
		}
		switch e.(type) {
		case *syntax.String, *syntax.Ident:
			return v.T.Literal.Kind == "string"
		case *syntax.Bool:
			return v.T.Literal.Kind == "boolean"
		}
		return isNumber(e) && v.T.Literal.Kind != "string" && v.T.Literal.Kind != "boolean"
	case "enum":
		if v.T.EnumKind == "string" {
			switch e.(type) {
			case *syntax.String, *syntax.Ident:
				return true
			case *syntax.ResLoc:
				return qualified(e)
			}
			return false
		}
		return isNumber(e)
	}
	return true
}

// qualified reports whether e is a resource location with a namespace
// (minecraft:stone), which is also accepted as a plain string.
func qualified(e syntax.Expr) bool {
	r, ok := e.(*syntax.ResLoc)
	return ok && strings.Contains(r.Text, ":")
}

func isNumber(e syntax.Expr) bool {
	switch e := e.(type) {
	case *syntax.Number:
		return true
	case *syntax.Unary:
		_, ok := e.X.(*syntax.Number)
		return ok && (e.Op == "-" || e.Op == "+")
	}
	return false
}

func (c *Checker) listValue(e syntax.Expr, v tview) jsonv.Value {
	l, ok := e.(*syntax.List)
	if !ok {
		c.mismatch(e, v)
		return nil
	}
	item := tview{T: v.T.Item, Env: v.Env}
	out := jsonv.Array{}
	for _, x := range l.Elems {
		if val := c.check(x, item); val != nil {
			out = append(out, val)
		}
	}
	if r := v.T.LengthRange; r != nil {
		n := float64(len(l.Elems))
		if (r.Min != nil && n < *r.Min) || (r.Max != nil && n > *r.Max) {
			c.errorf("mms/out-of-range", l.Sp, "list must have %s items, got %d", lengthText(r), len(l.Elems))
		}
	}
	return out
}

func (c *Checker) tupleValue(e syntax.Expr, v tview) jsonv.Value {
	l, ok := e.(*syntax.List)
	if !ok {
		c.mismatch(e, v)
		return nil
	}
	if len(l.Elems) != len(v.T.Items) {
		c.errorf("mms/type-mismatch", l.Sp, "expected %d items, got %d", len(v.T.Items), len(l.Elems))
	}
	out := jsonv.Array{}
	for i, x := range l.Elems {
		if i < len(v.T.Items) {
			out = append(out, c.check(x, tview{T: v.T.Items[i], Env: v.Env}))
		}
	}
	return out
}

func (c *Checker) number(e syntax.Expr, v tview) jsonv.Value {
	text, ok := numberText(e)
	if !ok {
		c.mismatch(e, v)
		return nil
	}
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		c.errorf("mms/type-mismatch", e.Span(), "invalid number %s", text)
		return nil
	}
	switch v.T.Kind {
	case "byte", "short", "int", "long":
		if f != math.Trunc(f) {
			c.errorf("mms/type-mismatch", e.Span(), "expected an integer, got %s", text)
		}
	}
	if !inRange(f, v.T.ValueRange) {
		c.errorf("mms/out-of-range", e.Span(), "value must be%s, got %s", rangeText(v.T.ValueRange), text)
	}
	if a, ok := v.T.Attr("divisible_by"); ok && a.Value != nil && a.Value.Literal != nil {
		// e.g. a dimension type's min_y and height: the game will not load
		// a world whose height is not a whole number of 16-block sections.
		if n, err := strconv.ParseFloat(a.Value.Literal.String(), 64); err == nil && n != 0 && math.Mod(f, n) != 0 {
			c.errorf("mms/out-of-range", e.Span(), "value must be a multiple of %s, got %s", a.Value.Literal.String(), text)
		}
	}
	return jsonv.Number(text)
}

func numberText(e syntax.Expr) (string, bool) {
	switch e := e.(type) {
	case *syntax.Number:
		return e.Text, true
	case *syntax.Unary:
		if n, ok := e.X.(*syntax.Number); ok {
			switch e.Op {
			case "-":
				return "-" + n.Text, true
			case "+":
				return n.Text, true
			}
		}
	case *syntax.Paren:
		return numberText(e.X)
	}
	return "", false
}

func inRange(f float64, r *spec.Range) bool {
	if r == nil {
		return true
	}
	minEx, maxEx := r.Kind&1 != 0, r.Kind&2 != 0
	if r.Min != nil && (f < *r.Min || (minEx && f == *r.Min)) {
		return false
	}
	if r.Max != nil && (f > *r.Max || (maxEx && f == *r.Max)) {
		return false
	}
	return true
}

func (c *Checker) literal(e syntax.Expr, v tview) jsonv.Value {
	lit := v.T.Literal
	if lit == nil {
		return c.anyValue(e)
	}
	want := lit.String()
	var got string
	switch e := e.(type) {
	case *syntax.String:
		got = e.Value
	case *syntax.Ident:
		got = e.Name
	case *syntax.Bool:
		got = strconv.FormatBool(e.Value)
	default:
		got, _ = numberText(e)
	}
	if got != want {
		c.errorf("mms/type-mismatch", e.Span(), "expected %s, got %s", string(lit.Value), exprKind(e))
		return nil
	}
	switch lit.Kind {
	case "string":
		return jsonv.String(want)
	case "boolean":
		return jsonv.Bool(want == "true")
	}
	return jsonv.Number(want)
}

func (c *Checker) enumValue(e syntax.Expr, v tview) jsonv.Value {
	var got string
	switch x := e.(type) {
	case *syntax.String:
		got = x.Value
	case *syntax.Ident:
		got = x.Name
	case *syntax.ResLoc:
		got = x.Text
	default:
		if t, ok := numberText(e); ok {
			got = t
		}
	}
	if v.T.EnumKind == "string" {
		got = strings.TrimPrefix(got, "minecraft:")
	}
	for _, val := range v.T.Values {
		s := val.String()
		if strings.TrimPrefix(s, "minecraft:") == got || (isPascal(got) && s == Snake(got)) {
			if v.T.EnumKind == "string" {
				c.info(e).Bare = BareEnum
				if v.RefID && !strings.Contains(s, ":") {
					s = "minecraft:" + s
				}
				return jsonv.String(s)
			}
			return jsonv.Number(s)
		}
	}
	var opts []string
	for _, val := range v.T.Values {
		opts = append(opts, val.String())
	}
	c.errorf("mms/type-mismatch", e.Span(), "expected one of %s, got %s", strings.Join(limit(opts, 12), ", "), exprKind(e))
	return nil
}

func limit(s []string, n int) []string {
	if len(s) > n {
		return append(s[:n:n], "...")
	}
	return s
}

// anyValue converts an expression to JSON without a schema.
func (c *Checker) anyValue(e syntax.Expr) jsonv.Value {
	switch x := e.(type) {
	case *syntax.Number, *syntax.Unary:
		if t, ok := numberText(e); ok {
			return jsonv.Number(t)
		}
	case *syntax.String:
		return jsonv.String(x.Value)
	case *syntax.Bool:
		return jsonv.Bool(x.Value)
	case *syntax.Ident:
		return jsonv.String(x.Name)
	case *syntax.ResLoc:
		return jsonv.String(x.Text)
	case *syntax.Tag:
		return jsonv.String("#" + x.Name)
	case *syntax.Paren:
		return c.anyValue(x.X)
	case *syntax.List:
		out := jsonv.Array{}
		for _, el := range x.Elems {
			out = append(out, c.anyValue(el))
		}
		return out
	case *syntax.Object:
		o := &jsonv.Object{}
		for _, it := range x.Body.Items {
			fs, ok := it.(*syntax.FieldStmt)
			if !ok {
				c.errorf("mms/expected-field", it.Span(), "expected `key = value`")
				continue
			}
			o.Set(fs.Key.Text, c.anyValue(fs.Value))
		}
		return o
	}
	c.errorf("mms/type-mismatch", e.Span(), "cannot use %s here", exprKind(e))
	return nil
}

// stringValue checks strings, and ids when the string has an #[id].
func (c *Checker) stringValue(e syntax.Expr, v tview) jsonv.Value {
	if isSelectorType(v.T) {
		if sel := selectorText(e); sel != "" {
			return jsonv.String(sel)
		}
	}
	id := idOf(v.T)
	switch x := e.(type) {
	case *syntax.String:
		if id.Slot && strings.HasPrefix(x.Value, "#") {
			return jsonv.String(x.Value) // another texture slot
		}
		if id.Present && id.Registry != "" && x.Value != "" {
			if strings.HasPrefix(x.Value, "#") {
				return jsonv.String(c.tagValue(x.Value[1:], id, e))
			}
			return jsonv.String(c.resolveID(x.Value, id, e))
		}
		if r := v.T.LengthRange; r != nil && !inRange(float64(len([]rune(x.Value))), r) {
			c.errorf("mms/out-of-range", e.Span(), "string length must be%s", rangeText(r))
		}
		return jsonv.String(x.Value)
	case *syntax.Ident, *syntax.ResLoc:
		if id.Tags == "required" {
			c.errorf("mms/type-mismatch", e.Span(), "expected a %s tag (#name) here", c.registryName(id.Registry))
			return nil
		}
		if id.Present {
			text := ""
			if i, ok := x.(*syntax.Ident); ok {
				text = i.Name
			} else {
				text = x.(*syntax.ResLoc).Text
			}
			return jsonv.String(c.resolveID(text, id, e))
		}
		if r, ok := x.(*syntax.ResLoc); ok && qualified(r) {
			return jsonv.String(r.Text)
		}
	case *syntax.Call:
		if assetCallName(x) != "" {
			return c.assetCall(e, v)
		}
	case *syntax.Tag:
		if id.Present && id.Tags != "" {
			return jsonv.String(c.tagValue(x.Name, id, e))
		}
		if id.Present {
			c.errorf("mms/type-mismatch", e.Span(), "a tag is not allowed here; expected a single %s id", id.Registry)
			return nil
		}
	}
	if r, ok := c.hoist(e, id); ok {
		return r
	}
	c.mismatch(e, v)
	return nil
}

// resolveID applies the name resolution rule: a qualified id is exact; an
// unqualified one names a declaration in the current namespace if there is
// one, and otherwise the registry's default namespace (minecraft).
func (c *Checker) resolveID(text string, id idInfo, e syntax.Expr) string {
	if id.Tags == "implicit" {
		// The id names a tag, written without '#'.
		return strings.TrimPrefix(c.resolveTag(text, id, e), "#")
	}
	reg := c.registryName(id.Registry)
	ns, path := splitID(text)
	if path != strings.ToLower(path) {
		c.errorf("mms/invalid-name", e.Span(), "resource ids are lowercase: did you mean '%s'?", Snake(path))
		return text
	}
	if IsAssetRegistry(reg) {
		path = c.stripAssetExt(path, reg, e)
	}
	var sym *Symbol
	if ns == "" {
		if reg != "" {
			sym = c.find(reg, c.NS+":"+path)
		}
		if sym != nil {
			ns = c.NS
		} else {
			ns = "minecraft"
		}
	} else if reg != "" {
		sym = c.find(reg, ns+":"+path)
	}
	full := ns + ":" + path
	in := c.info(e)
	in.ID = full
	in.Symbol = sym
	in.Bare, in.Folder = BareID, reg
	if id.excluded(full) {
		c.excludedID(full, id, e.Span())
		return full
	}
	if sym != nil || reg == "" {
		return full
	}
	if ns == "minecraft" {
		if _, known := c.Pack.Registries[reg]; known && !c.Pack.HasID(reg, id.Path+path) && !builtinIDs[reg+"|"+path] {
			if text == path && id.Path == "" {
				// A bare name that matches a local declaration of another
				// kind: say so instead of blaming minecraft:<name>.
				if other := c.otherKind(reg, c.NS+":"+path); other != nil {
					c.errorf("mms/type-mismatch", e.Span(), "%s", c.wrongKindMessage(path, other, reg))
					return c.NS + ":" + path
				}
			}
			msg := fmt.Sprintf("unknown %s '%s'", reg, full)
			if s := closest(path, withPrefix(c.Pack.Registries[reg], id.Path)); s != "" {
				msg += fmt.Sprintf("; did you mean '%s'?", s)
			}
			if id.Slot {
				// Atlases also make sprites that are not files, such as
				// armor trim colours, so this may still work.
				c.warnf("mms/unknown-id", e.Span(), "%s, unless an atlas makes it", msg)
				return full
			}
			c.errorf("mms/unknown-id", e.Span(), "%s", msg)
		}
		return full
	}
	if reg == "texture" && c.resource == "model" {
		in.ModelTexture = true
	}
	if IsAssetRegistry(reg) {
		// A file, not a declaration: the external-asset lint rule checks
		// that something provides it.
		return full
	}
	if c.declaresNamespace(ns) {
		c.errorf("mms/unknown-id", e.Span(), "no %s named '%s' is declared", reg, full)
	}
	return full
}

// wrongKindMessage explains that name is declared, but as other's kind
// rather than one whose files live in folder.
func (c *Checker) wrongKindMessage(name string, other *Symbol, folder string) string {
	want := folder
	if k, ok := c.Kinds.ForFolder(folder); ok {
		want = k.Name
	}
	msg := fmt.Sprintf("'%s' is a %s, not a %s", name, other.Kind.Name, want)
	if other.Kind.Folder() == "worldgen/noise" && folder == "worldgen/density_function" {
		msg += fmt.Sprintf("; sample it with Noise(%s, 1, 1)", name)
	}
	return msg
}

// excludedID reports an id the game rejects at this spot although it is in
// the registry (mcdoc's #[id(exclude=[...])]).
func (c *Checker) excludedID(full string, id idInfo, sp syntax.Span) {
	reg := c.registryName(id.Registry)
	switch {
	case reg == "environment_attribute":
		// Biomes take only positional attributes; the game fails to load
		// the biome with "The following attributes cannot be positional".
		c.errorf("mms/not-allowed-here", sp, "'%s' cannot be set on a biome: Minecraft reads it only from a dimension type's attributes; set it there instead", full)
	case full == "minecraft:air":
		c.errorf("mms/not-allowed-here", sp, "'%s' is not allowed here: it is an empty item, and the game rejects it", full)
	default:
		c.errorf("mms/not-allowed-here", sp, "%s '%s' is not allowed here", reg, full)
	}
}

// builtinIDs are ids the game knows that are not in its registries.
var builtinIDs = map[string]bool{
	"model|builtin/generated": true, // item models made from their layer textures
}

func (c *Checker) resolveTag(text string, id idInfo, e syntax.Expr) string {
	reg := c.registryName(id.Registry)
	ns, path := splitID(text)
	var sym *Symbol
	if ns == "" {
		sym = c.find("tags/"+reg, c.NS+":"+path)
		if sym != nil {
			ns = c.NS
		} else {
			ns = "minecraft"
		}
	} else {
		sym = c.find("tags/"+reg, ns+":"+path)
	}
	full := ns + ":" + path
	in := c.info(e)
	in.ID = "#" + full
	in.Symbol = sym
	in.Bare, in.Folder = BareID, "tags/"+reg
	if sym == nil && ns == "minecraft" {
		if tags, known := c.Pack.Registries["tag/"+reg]; known && !contains(tags, path) {
			c.errorf("mms/unknown-id", e.Span(), "unknown %s tag '#%s'", reg, full)
		}
	} else if sym == nil && c.declaresNamespace(ns) {
		c.errorf("mms/unknown-id", e.Span(), "no %s tag named '#%s' is declared", reg, full)
	}
	return "#" + full
}

// tagValue resolves a #tag and writes it the way the field expects: fields
// marked tags="implicit" always hold a tag, so the game wants it without '#'.
func (c *Checker) tagValue(text string, id idInfo, e syntax.Expr) string {
	full := c.resolveTag(text, id, e)
	if id.Tags == "implicit" {
		return strings.TrimPrefix(full, "#")
	}
	return full
}

func (c *Checker) declaresNamespace(ns string) bool {
	if ns == c.NS {
		return true
	}
	has := c.Syms.hasNamespace(ns)
	if c.Deps != nil {
		c.Deps["n|"+ns] = has
	}
	return has
}

// withPrefix returns the ids that start with prefix, without it.
func withPrefix(ids []string, prefix string) []string {
	if prefix == "" {
		return ids
	}
	var out []string
	for _, id := range ids {
		if rest, ok := strings.CutPrefix(id, prefix); ok {
			out = append(out, rest)
		}
	}
	return out
}

func contains(xs []string, x string) bool {
	for _, s := range xs {
		if s == x {
			return true
		}
	}
	return false
}

// closest returns the candidate with the smallest edit distance to s, if
// it is close enough to be a likely typo.
func closest(s string, cands []string) string { return near.Closest(s, cands) }

// inlineSound writes Sound("x.ogg") where a sound event may be given
// inline (mcdoc's SoundEventRef: an id or { sound_id, range }) as
// { "sound_id": "ns:x" }. The minecraft:sound_event registry is built
// into the game, so data packs cannot add to it and a pack's own sound
// only works inline; the id form fails to load (decision D91).
func (c *Checker) inlineSound(e syntax.Expr, v tview) (jsonv.Value, bool) {
	if assetCallName(e) != "Sound" {
		return nil, false
	}
	for _, m := range v.T.Members {
		mv := c.norm(tview{T: m, Env: v.Env})
		if mv.T == nil || mv.T.Kind != "struct" {
			continue
		}
		for _, f := range mv.T.Fields {
			if f.Key != "sound_id" || f.Type == nil || !c.assetFits(e, idOf(f.Type)) {
				continue
			}
			id := c.check(e, c.norm(tview{T: f.Type, Env: mv.Env}))
			if id == nil {
				return nil, true // already reported
			}
			o := &jsonv.Object{}
			o.Set("sound_id", id)
			return o, true
		}
	}
	return nil, false
}
