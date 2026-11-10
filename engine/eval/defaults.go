package eval

import "github.com/minecraftmetascript/mms/engine/syntax"

// withDefaults fills in the defaults of a record type's fields that a
// value leaves out, so `const waves: [Wave] = [{ mob = skeleton }]` gives
// the skeleton wave Wave's `count = 1`, and `w.count` reads it. It
// follows lists, maps, optionals and nested records; other values are
// returned as they are.
func (c *ctx) withDefaults(v Value, t syntax.TypeExpr) Value {
	return c.defaultsAt(v, t, 0)
}

func (c *ctx) defaultsAt(v Value, t syntax.TypeExpr, depth int) Value {
	if v == nil || t == nil || depth > 16 {
		return v
	}
	switch t := t.(type) {
	case *syntax.TypeParen:
		return c.defaultsAt(v, t.X, depth+1)
	case *syntax.TypeOptional:
		return c.defaultsAt(v, t.X, depth+1)
	case *syntax.TypeList:
		l, ok := v.(*List)
		if !ok {
			return v
		}
		out := &List{pos: l.pos, Elems: make([]Value, len(l.Elems))}
		for i, el := range l.Elems {
			out.Elems[i] = c.defaultsAt(el, t.Elem, depth+1)
		}
		return out
	case *syntax.TypeMap:
		r, ok := v.(*Record)
		if !ok || r.id != nil {
			return v
		}
		out := r.copy()
		for i := range out.Vals {
			out.Vals[i] = c.defaultsAt(out.Vals[i], t.Value, depth+1)
		}
		return out
	case *syntax.TypeRecord:
		return c.recordDefaults(v, t.Fields, c.defPath, depth)
	case *syntax.TypeName:
		td := c.e.r.Type(c.defPath, t.Name)
		if td == nil {
			return v
		}
		if td.D.TypeValue != nil {
			return c.defaultsAt(v, td.D.TypeValue, depth+1)
		}
		if td.D.Body == nil {
			return v
		}
		var fs []*syntax.FieldDecl
		for _, it := range td.D.Body.Items {
			if f, ok := it.(*syntax.FieldDecl); ok {
				fs = append(fs, f)
			}
		}
		return c.recordDefaults(v, fs, td.Path, depth)
	}
	return v
}

// recordDefaults fills a record's missing fields from their defaults,
// worked out where the type is written.
func (c *ctx) recordDefaults(v Value, fs []*syntax.FieldDecl, path string, depth int) Value {
	r, ok := v.(*Record)
	if !ok || r.id != nil {
		return v
	}
	out := r.copy()
	for _, f := range fs {
		if f.Name == nil {
			continue
		}
		if cur, ok := out.get(f.Name.Text); ok {
			out.set(f.Name.Text, c.defaultsAt(cur, f.Type, depth+1))
			continue
		}
		if f.Default == nil {
			continue
		}
		// The default is written with the type, maybe in another file:
		// what it builds belongs to the record.
		site := r.Sp
		if c.site != nil {
			site = *c.site
		}
		dc := &ctx{e: c.e, path: c.path, defPath: path, site: &site, sc: &scope{}, compile: true, ns: c.ns, depth: c.depth}
		out.set(f.Name.Text, dc.defaultsAt(dc.reduce(f.Default), f.Type, depth+1))
	}
	return out
}

// copy returns a record whose fields can be set without changing r.
func (r *Record) copy() *Record {
	return &Record{pos: r.pos, Keys: append([]string(nil), r.Keys...), Vals: append([]Value(nil), r.Vals...), id: r.id, keySp: r.keySp}
}
