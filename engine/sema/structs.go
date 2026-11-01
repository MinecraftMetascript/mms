package sema

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// assign is one field assignment collected from an object body, a call's
// arguments or a builder method.
type assign struct {
	key  string
	expr syntax.Expr
	val  jsonv.Value // a constant from an overlay instead of expr
	sp   syntax.Span
	node syntax.Node // the key's node, for hover and completion
}

func (c *Checker) structValue(e syntax.Expr, v tview) jsonv.Value {
	switch x := e.(type) {
	case *syntax.Object:
		return c.objectStruct(x.Body, x.Sp, v)
	case *syntax.Call, *syntax.Trailing, *syntax.Member:
		return c.callStruct(e, v)
	}
	c.mismatch(e, v)
	return nil
}

func (c *Checker) objectStruct(body *syntax.Block, sp syntax.Span, v tview) jsonv.Value {
	fs, disp := c.fields(v, "")
	caseKey := ""
	if disp != nil {
		for _, it := range body.Items {
			if f, ok := it.(*syntax.FieldStmt); ok && f.Key.Text == disp.KeyField {
				caseKey = c.caseKeyOf(f.Value, disp)
			}
		}
		keys := map[string]string{}
		for _, it := range body.Items {
			if f, ok := it.(*syntax.FieldStmt); ok {
				keys[f.Key.Text] = keyText(f.Value)
			}
		}
		fs, disp = c.fieldsWith(v, func(k string) string {
			if k == disp.KeyField {
				return caseKey
			}
			return secondaryKey(keys, k)
		})
		if disp != nil {
			fs = withDocs(disp.Dispatcher+"/"+caseKey, disp.Dispatcher, fs)
		}
	}
	in := c.info(body)
	in.Type, in.Fields, in.Disp = v, fs, disp
	var as []assign
	for _, it := range body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok {
			c.errorf("mms/expected-field", it.Span(), "expected `field = value` in %s", describe(v))
			continue
		}
		as = append(as, assign{key: f.Key.Text, expr: f.Value, sp: f.Key.Sp, node: f.Key})
	}
	return c.buildStruct(fs, disp, caseKey, as, syntax.Span{Start: sp.Start, End: sp.Start + 1}, nil, v, fixSite{block: body})
}

// caseKeyOf reads a dispatch key such as `type = add` or `type = "minecraft:add"`.
func (c *Checker) caseKeyOf(e syntax.Expr, disp *dispInfo) string {
	var text string
	switch x := e.(type) {
	case *syntax.Ident:
		text = x.Name
	case *syntax.ResLoc:
		text = x.Text
	case *syntax.String:
		text = x.Value
	default:
		return ""
	}
	key := strings.TrimPrefix(text, "minecraft:")
	if !c.hasDispatchKey(disp, key) {
		msg := "unknown type '" + text + "'"
		if s := closest(key, c.dispatchKeys(disp)); s != "" {
			msg += "; did you mean '" + s + "'?"
		}
		c.errorf("mms/unknown-id", e.Span(), "%s", msg)
		return ""
	}
	in := c.info(e)
	in.ID = "minecraft:" + key
	in.Bare = BareKey
	return key
}

// buildStruct turns assignments into an object in schema field order,
// checking each value at its field's type and reporting missing and
// unknown fields.
func (c *Checker) buildStruct(fs []field, disp *dispInfo, caseKey string, as []assign, at syntax.Span, ov *overlay, v tview, site fixSite) jsonv.Value {
	o := &jsonv.Object{}
	byKey := map[string]*assign{}
	for i := range as {
		a := &as[i]
		if prev, dup := byKey[a.key]; dup && prev.expr != nil && a.expr != nil {
			c.errorf("mms/duplicate", a.sp, "'%s' is set twice", a.key)
		}
		byKey[a.key] = a
	}
	used := map[string]bool{}
	var missing []string
	for i := range fs {
		f := &fs[i]
		if f.KeyType != nil {
			continue
		}
		if disp != nil && f.Key == disp.KeyField {
			used[f.Key] = true
			if caseKey != "" {
				o.Set(f.Key, jsonv.String("minecraft:"+caseKey))
			} else if a := byKey[f.Key]; a != nil {
				o.Set(f.Key, c.check(a.expr, f.Type))
			} else if !f.Optional {
				missing = append(missing, f.Key)
			}
			continue
		}
		a, ok := byKey[f.Key]
		if !ok {
			if ov != nil {
				if d, ok := ov.defaults[f.Key]; ok {
					o.Set(f.Key, d)
					continue
				}
			}
			if !f.Optional {
				missing = append(missing, f.Key)
			}
			continue
		}
		used[f.Key] = true
		if a.node != nil {
			c.info(a.node).Field = f
		}
		if a.val != nil {
			o.Set(f.Key, a.val)
			continue
		}
		if val := c.check(a.expr, c.siblingCase(f.Type, byKey)); val != nil {
			o.Set(f.Key, val)
		}
	}
	for i := range as {
		a := &as[i]
		if used[a.key] {
			continue
		}
		if mapField := c.mapFieldFor(fs, a.key); mapField != nil {
			used[a.key] = true
			if a.node != nil {
				c.info(a.node).Field = mapField
			}
			c.checkKey(a, *mapField.KeyType)
			if a.val != nil {
				o.Set(a.key, a.val)
			} else if val := c.check(a.expr, c.mapValueType(mapField, a.key)); val != nil {
				o.Set(a.key, val)
			}
			continue
		}
		if disp != nil && disp.NoCase {
			// The mcdoc does not describe this type's fields.
			if a.val != nil {
				o.Set(a.key, a.val)
			} else if val := c.anyValue(a.expr); val != nil {
				o.Set(a.key, val)
			}
			continue
		}
		var names []string
		for _, f := range fs {
			names = append(names, f.Key)
		}
		msg := "unknown field '" + a.key + "' in " + describe(v)
		if s := closest(a.key, names); s != "" {
			msg += "; did you mean '" + s + "'?"
		}
		c.errorf("mms/unknown-field", a.sp, "%s", msg)
	}
	if len(missing) > 0 {
		if disp != nil && caseKey == "" && contains(missing, disp.KeyField) {
			c.errorf("mms/missing-field", at, "%s needs a '%s' field (or use a helper like %s(...))", describe(v), disp.KeyField, exampleCallable(c, disp))
		} else {
			c.errorf("mms/missing-field", at, "missing %s %s", plural(len(missing), "field"), quoteList(missing))
			if !c.trial {
				c.Diags[len(c.Diags)-1].Fixes = c.missingFix(site, fs, missing)
			}
		}
	}
	return o
}

// siblingCase resolves a value that dispatches on a sibling field, such
// as an advancement's `display`, which is `advancement_display[[parent]]`:
// the case is the sibling's value, %unknown for one that has no case of
// its own, or %none when the sibling is not set.
func (c *Checker) siblingCase(v tview, byKey map[string]*assign) tview {
	t := v.T
	if t == nil || t.Kind != "dispatcher" || len(t.ParallelIndices) != 1 || t.ParallelIndices[0].Kind != "dynamic" {
		return v
	}
	acc := t.ParallelIndices[0].Accessor
	if len(acc) != 1 {
		return v
	}
	var sib string
	if json.Unmarshal(acc[0], &sib) != nil || sib == "" {
		return v
	}
	keys := []string{"%none"}
	if a := byKey[sib]; a != nil {
		keys = []string{keyText(a.expr), "%unknown"}
	}
	for _, k := range keys {
		if cs, ok := c.Pack.Dispatch(t.Registry, k); ok && k != "" {
			return tview{T: cs, Env: v.Env}
		}
	}
	return v
}

func exampleCallable(c *Checker, d *dispInfo) string {
	keys := c.dispatchKeys(d)
	if len(keys) == 0 {
		return "Name"
	}
	return Pascal(keys[0])
}

func plural(n int, s string) string {
	if n == 1 {
		return s
	}
	return s + "s"
}

func quoteList(xs []string) string {
	q := make([]string, len(xs))
	for i, x := range xs {
		q[i] = "'" + x + "'"
	}
	return strings.Join(q, ", ")
}

// checkKey validates a map key against the map's key type when it is an
// id, e.g. block names in block state property maps.
func (c *Checker) checkKey(a *assign, kt tview) {
	kt = c.norm(kt)
	if kt.T == nil || kt.T.Kind != "string" {
		return
	}
	id := idOf(kt.T)
	if len(id.Exclude) > 0 {
		full := a.key
		if !strings.Contains(full, ":") {
			full = "minecraft:" + full
		}
		if id.excluded(full) {
			c.excludedID(full, id, a.sp)
			if !c.trial {
				sp := a.sp
				if a.expr != nil {
					sp = sp.Join(a.expr.Span())
				}
				c.Diags[len(c.Diags)-1].Fixes = c.removeFix(sp, "Remove '"+a.key+"'")
			}
			return
		}
	}
	if id.Present && id.Registry != "" && a.node != nil {
		if n, ok := a.node.(*syntax.Name); ok {
			c.resolveID(a.key, id, &syntax.Ident{Name: a.key, Sp: n.Sp})
		}
	}
}

// chainStep is one `.Method(args)` after a call.
type chainStep struct {
	name *syntax.Name
	args []*syntax.Arg
	sp   syntax.Span
	call bool
}

// unchain splits `Base(args).A(x).B { block }` into the base call, its
// trailing block and the method steps.
func unchain(e syntax.Expr) (base *syntax.Call, name *syntax.Ident, trailing *syntax.Block, steps []chainStep) {
	if t, ok := e.(*syntax.Trailing); ok {
		trailing = t.Block
		e = t.X
	}
	for {
		switch x := e.(type) {
		case *syntax.Call:
			if m, ok := x.Fn.(*syntax.Member); ok {
				steps = append(steps, chainStep{name: m.Name, args: x.Args, sp: x.Sp, call: true})
				e = m.X
				continue
			}
			if id, ok := x.Fn.(*syntax.Ident); ok {
				base, name = x, id
			}
		case *syntax.Member:
			steps = append(steps, chainStep{name: x.Name, sp: x.Sp})
			e = x.X
			continue
		case *syntax.Ident:
			// `Hole` without parentheses is not a call; callers handle it.
		}
		break
	}
	for i, j := 0, len(steps)-1; i < j; i, j = i+1, j-1 {
		steps[i], steps[j] = steps[j], steps[i]
	}
	return base, name, trailing, steps
}

func callName(e syntax.Expr) string {
	_, name, _, _ := unchain(e)
	if name == nil {
		return ""
	}
	return name.Name
}

// callable resolves the helper a call names at type v.
func (c *Checker) callable(e syntax.Expr, v tview) *Callable {
	_, name, _, _ := unchain(e)
	if name == nil {
		return nil
	}
	return c.callableByName(name.Name, v)
}

func (c *Checker) callableByName(name string, v tview) *Callable {
	v = c.norm(v)
	if v.T == nil || v.T.Kind != "struct" {
		return nil
	}
	fs, disp := c.fields(v, "")
	if disp != nil {
		key := Snake(name)
		if a, ok := callAliases[disp.Dispatcher][name]; ok {
			key = a
		}
		if c.hasDispatchKey(disp, key) {
			cfs, _ := c.fields(v, key)
			return &Callable{Name: name, Case: key, Struct: v, Fields: withDocs(disp.Dispatcher+"/"+key, disp.Dispatcher, cfs), Disp: disp, Ov: overlays[disp.Dispatcher+"/"+key]}
		}
		return nil
	}
	if v.Name != "" {
		short := spec.ShortPath(v.Name)
		if name == short || structAliases[name] == short {
			return &Callable{Name: name, Struct: v, Fields: withDocs(short, "", fs), Ov: overlays[short]}
		}
	}
	if len(fs) == 1 && fs[0].KeyType == nil && Pascal(fs[0].Key) == name {
		return &Callable{Name: name, Struct: v, Fields: fs}
	}
	return nil
}

// callableNames lists the helpers usable at v, for messages and completion.
func (c *Checker) callableNames(v tview) []string {
	set := map[string]bool{}
	var walk func(v tview, depth int)
	walk = func(v tview, depth int) {
		v = c.norm(v)
		if v.T == nil || depth > 6 {
			return
		}
		switch v.T.Kind {
		case "union":
			for _, m := range v.T.Members {
				walk(tview{T: m, Env: v.Env}, depth+1)
			}
		case "struct":
			fs, disp := c.fields(v, "")
			if disp != nil {
				aliased := map[string]bool{}
				for alias, key := range callAliases[disp.Dispatcher] {
					set[alias] = true
					aliased[key] = true
				}
				for _, k := range c.dispatchKeys(disp) {
					if !aliased[k] {
						set[Pascal(k)] = true
					}
				}
			} else if v.Name != "" {
				short := spec.ShortPath(v.Name)
				set[short] = true
				for alias, target := range structAliases {
					if target == short {
						set[alias] = true
					}
				}
			}
			if disp == nil && len(fs) == 1 && fs[0].KeyType == nil {
				set[Pascal(fs[0].Key)] = true
			}
		}
		for _, s := range sugarCallables[v.Name] {
			set[s] = true
		}
	}
	walk(v, 0)
	out := make([]string, 0, len(set))
	for s := range set {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func (c *Checker) callStruct(e syntax.Expr, v tview) jsonv.Value {
	base, name, trailing, steps := unchain(e)
	if base == nil {
		c.mismatch(e, v)
		return nil
	}
	cl := c.callableByName(name.Name, v)
	if cl == nil {
		names := c.callableNames(v)
		msg := "'" + name.Name + "' does not make a " + describe(v)
		if s := closest(name.Name, names); s != "" {
			msg += "; did you mean " + s + "(...)?"
		} else if len(names) > 0 {
			msg += "; expected one of " + strings.Join(limit(names, 10), ", ")
		}
		c.errorf("mms/unknown-callable", name.Sp, "%s", msg)
		return nil
	}
	in := c.info(name)
	in.Callable, in.Type = cl, v
	ci := c.info(base)
	ci.Callable, ci.Fields, ci.Disp = cl, cl.Fields, cl.Disp

	var as []assign
	slots := c.positional(cl)
	named := map[string]bool{}
	for _, a := range base.Args {
		if a.Name == nil {
			continue
		}
		key := a.Name.Text
		if !hasField(cl.Fields, key) && hasField(cl.Fields, Snake(key)) {
			key = Snake(key)
		}
		named[key] = true
		as = append(as, assign{key: key, expr: a.Value, sp: a.Name.Sp, node: a.Name})
	}
	var pos []syntax.Expr
	for _, a := range base.Args {
		if a.Name == nil {
			pos = append(pos, a.Value)
		}
	}
	if trailing != nil {
		pos = append(pos, blockExpr(trailing))
	}
	var free []field
	for _, f := range slots {
		if !named[f.Key] {
			free = append(free, f)
		}
	}
	for i := 0; i < len(pos); i++ {
		if i >= len(free) {
			c.errorf("mms/too-many-args", pos[i].Span(), "%s takes %d %s", cl.Name, len(free), plural(len(free), "argument"))
			break
		}
		f := free[i]
		// Varargs: a list field takes the rest of the arguments when
		// they are not already a list.
		if c.acceptsList(f.Type) {
			if _, isList := pos[i].(*syntax.List); !isList && (len(pos)-i > len(free)-i || !c.fits(pos[i], f.Type)) {
				rest := pos[i:]
				sp := rest[0].Span().Join(rest[len(rest)-1].Span())
				as = append(as, assign{key: f.Key, expr: &syntax.List{Elems: rest, Sp: sp}, sp: sp})
				break
			}
		}
		as = append(as, assign{key: f.Key, expr: pos[i], sp: pos[i].Span()})
	}
	for _, st := range steps {
		if st.name == nil {
			continue // `Foo().` while typing; the parser reported it
		}
		as = append(as, c.method(cl, st)...)
	}
	fs, disp := cl.Fields, cl.Disp
	if disp != nil {
		// Secondary dispatchers (like `formula` beside `type`) depend on
		// values given by name or method.
		keys := map[string]string{}
		for _, a := range as {
			if a.expr != nil {
				keys[a.key] = keyText(a.expr)
			}
		}
		fs, disp = c.fieldsWith(cl.Struct, func(k string) string {
			if k == cl.Disp.KeyField {
				return cl.Case
			}
			return secondaryKey(keys, k)
		})
		fs = withDocs(cl.Disp.Dispatcher+"/"+cl.Case, cl.Disp.Dispatcher, fs)
		fs = withDocs(cl.Disp.Dispatcher+"/"+cl.Case, cl.Disp.Dispatcher, fs)
	}
	return c.buildStruct(fs, disp, cl.Case, as, name.Sp, cl.Ov, v, fixSite{call: base})
}

// keyText reads an id-like value as a dispatch key, without namespace.
func keyText(e syntax.Expr) string {
	var s string
	switch x := e.(type) {
	case *syntax.Ident:
		s = x.Name
	case *syntax.ResLoc:
		s = x.Text
	case *syntax.String:
		s = x.Value
	}
	return strings.TrimPrefix(s, "minecraft:")
}

func hasField(fs []field, key string) bool {
	for _, f := range fs {
		if f.Key == key {
			return true
		}
	}
	return false
}

// positional returns the fields positional arguments fill, in order.
func (c *Checker) positional(cl *Callable) []field {
	var out []field
	if cl.Ov != nil && len(cl.Ov.positional) > 0 {
		for _, k := range cl.Ov.positional {
			for _, f := range cl.Fields {
				if f.Key == k {
					out = append(out, f)
				}
			}
		}
		return out
	}
	for _, f := range cl.Fields {
		if f.KeyType != nil || (cl.Disp != nil && f.Key == cl.Disp.KeyField) {
			continue
		}
		out = append(out, f)
	}
	return out
}

func (c *Checker) acceptsList(v tview) bool {
	v = c.norm(v)
	if v.T == nil {
		return false
	}
	switch v.T.Kind {
	case "list":
		return true
	case "union":
		for _, m := range v.T.Members {
			if c.acceptsList(tview{T: m, Env: v.Env}) {
				return true
			}
		}
	}
	return false
}

// method applies one builder step: an overlay alias, or a field named by
// the method in snake_case.
func (c *Checker) method(cl *Callable, st chainStep) []assign {
	name := st.name.Text
	if cl.Ov != nil {
		if m, ok := cl.Ov.methods[name]; ok {
			c.info(st.name).Field = fieldByKey(cl.Fields, m.field)
			if m.val != nil {
				if len(st.args) > 0 {
					c.errorf("mms/too-many-args", st.sp, "%s() takes no arguments", name)
				}
				return []assign{{key: m.field, val: m.val, sp: st.name.Sp}}
			}
			return c.methodArgs(m.field, cl, st)
		}
	}
	key := Snake(name)
	f := fieldByKey(cl.Fields, key)
	if f == nil {
		var opts []string
		for _, f := range cl.Fields {
			if cl.Disp == nil || f.Key != cl.Disp.KeyField {
				opts = append(opts, Pascal(f.Key))
			}
		}
		if cl.Ov != nil {
			for m := range cl.Ov.methods {
				opts = append(opts, m)
			}
		}
		sort.Strings(opts)
		msg := cl.Name + " has no method " + name
		if s := closest(name, opts); s != "" {
			msg += "; did you mean ." + s + "(...)?"
		}
		c.errorf("mms/unknown-method", st.name.Sp, "%s", msg)
		return nil
	}
	c.info(st.name).Field = f
	return c.methodArgs(key, cl, st)
}

func (c *Checker) methodArgs(key string, cl *Callable, st chainStep) []assign {
	f := fieldByKey(cl.Fields, key)
	switch {
	case len(st.args) == 0:
		if f != nil {
			if t := c.norm(f.Type); t.T != nil && t.T.Kind == "boolean" {
				return []assign{{key: key, val: jsonv.Bool(true), sp: st.name.Sp, node: st.name}}
			}
		}
		c.errorf("mms/missing-value", st.sp, ".%s needs a value", st.name.Text)
		return nil
	case len(st.args) == 1 && st.args[0].Name == nil:
		return []assign{{key: key, expr: st.args[0].Value, sp: st.name.Sp, node: st.name}}
	}
	if f != nil && c.acceptsList(f.Type) {
		var el []syntax.Expr
		for _, a := range st.args {
			el = append(el, a.Value)
		}
		sp := el[0].Span().Join(el[len(el)-1].Span())
		return []assign{{key: key, expr: &syntax.List{Elems: el, Sp: sp}, sp: st.name.Sp, node: st.name}}
	}
	c.errorf("mms/too-many-args", st.sp, ".%s takes one argument", st.name.Text)
	return nil
}

func fieldByKey(fs []field, key string) *field {
	for i := range fs {
		if fs[i].Key == key {
			return &fs[i]
		}
	}
	return nil
}

// blockExpr reads a trailing block as a value: one expression stays
// itself, several become a list, and `key = value` items an object.
func blockExpr(b *syntax.Block) syntax.Expr {
	var exprs []syntax.Expr
	fields := false
	for _, it := range b.Items {
		switch s := it.(type) {
		case *syntax.ExprStmt:
			exprs = append(exprs, s.X)
		case *syntax.FieldStmt:
			fields = true
		}
	}
	if fields {
		return &syntax.Object{Body: b, Sp: b.Sp}
	}
	if len(exprs) == 1 && len(b.Items) == 1 {
		return exprs[0]
	}
	return &syntax.List{Elems: exprs, Sp: b.Sp}
}
