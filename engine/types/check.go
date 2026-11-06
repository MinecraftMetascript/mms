package types

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/eval"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Checkers makes a schema checker for declarations in namespace ns.
type Checkers func(ns string) *sema.Checker

// Check checks what the evaluator was given against the types written
// for it: each Blueprint field the user set, each Template argument in
// user code, and each const with a type. It returns diagnostics by file,
// on the user's spans.
func (e *Env) Check(r *eval.Result, ns string, newChecker Checkers) map[string][]syntax.Diagnostic {
	out := map[string][]syntax.Diagnostic{}
	for p, ds := range e.Diags {
		out[p] = append(out[p], ds...)
	}
	add := func(path string, c *sema.Checker) {
		out[path] = append(out[path], c.Diags...)
	}
	for _, inst := range r.Instances {
		if len(inst.Written) == 0 {
			continue
		}
		bpNS := ns
		if v, ok := inst.Self.Vals[1].(*eval.Str); ok {
			bpNS = v.S
		}
		c := newChecker(bpNS)
		for _, f := range inst.Blueprint.Fields {
			if f.Name == nil {
				continue
			}
			x, ok := inst.Written[f.Name.Text]
			if !ok {
				continue
			}
			t, _ := e.Resolve(inst.Blueprint.Path, f.Type)
			check(c, x, t, syntax.TypeText(f.Type))
		}
		add(inst.Path, c)
	}
	for _, call := range r.Calls {
		d := call.Template.D
		if d.Params == nil {
			continue
		}
		c := newChecker(ns)
		for _, p := range d.Params.Params {
			if p.Name == nil || p.Type == nil {
				continue
			}
			x, ok := call.Args[p.Name.Text]
			if !ok {
				continue
			}
			t, _ := e.Resolve(call.Template.Path, p.Type)
			check(c, x, t, syntax.TypeText(p.Type))
		}
		add(call.Path, c)
	}
	names := make([]string, 0, len(r.Consts))
	for n := range r.Consts {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		k := r.Consts[n]
		if k.S.Type == nil || k.S.Value == nil {
			continue
		}
		x := k.Lowered()
		if x == nil {
			continue
		}
		t, _ := e.Resolve(k.Path, k.S.Type)
		c := newChecker(ns)
		check(c, x, t, syntax.TypeText(k.S.Type))
		add(k.Path, c)
	}
	for p, ds := range out {
		out[p] = dedupe(ds)
	}
	return out
}

// check runs the schema checker on x at type t, unless t is any: the
// values of functions, commands and asset files are checked where they
// are used.
func check(c *sema.Checker, x syntax.Expr, t *spec.Type, name string) {
	if t == nil || t.Kind == "any" {
		return
	}
	c.CheckType(x, t, name)
}

// dedupe drops a diagnostic met again: a call in a loop is checked once
// per iteration.
func dedupe(ds []syntax.Diagnostic) []syntax.Diagnostic {
	type key struct {
		sp  syntax.Span
		msg string
	}
	seen := map[key]bool{}
	out := ds[:0]
	for _, d := range ds {
		k := key{d.Sp, d.Message}
		if seen[k] {
			continue
		}
		seen[k] = true
		out = append(out, d)
	}
	return out
}

// Conversions resolves the `as` conversions of every Blueprint instance,
// by the instance's id. An instance whose primary declaration is another
// instance also converts as that one does.
func (e *Env) Conversions(r *eval.Result, ns string) map[string][]*sema.Conversion {
	out := map[string][]*sema.Conversion{}
	for _, inst := range r.Instances {
		if inst.Name == nil {
			continue
		}
		id := inst.Name.Text
		if !strings.Contains(id, ":") {
			instNS := ns
			if v, ok := inst.Self.Vals[1].(*eval.Str); ok {
				instNS = v.S
			}
			id = instNS + ":" + id
		}
		seen := map[*eval.Instance]bool{}
		for in := inst; in != nil && !seen[in]; in = in.Forward {
			seen[in] = true
			for _, cv := range in.Conversions {
				t, _ := e.Resolve(cv.Path, cv.Type)
				out[id] = append(out[id], &sema.Conversion{Type: t, X: cv.X, Rule: cv.Rule, Hint: cv.Hint,
					Sig:  syntax.TypeText(cv.Type) + "=" + format.Expr(cv.X) + "|" + cv.Rule + "|" + cv.Hint,
					Kind: inst.Blueprint.D.NameText(), Decl: inst.D, Path: inst.Path})
			}
		}
	}
	return out
}

// InstanceType is the struct type of a Blueprint's fields, for the editor
// to complete and describe an instance body. A Function field takes a
// function id there, or a `run { }` body, which the editor reads as code.
func (e *Env) InstanceType(bp *eval.Blueprint) *spec.Type {
	t := &spec.Type{Kind: "struct"}
	for _, f := range bp.Fields {
		if f.Name == nil {
			continue
		}
		ft, opt := e.Resolve(bp.Path, f.Type)
		if n, ok := unwrapOptional(f.Type).(*syntax.TypeName); ok && n.Name == "Function" {
			ft = idType("function")
		}
		t.Fields = append(t.Fields, &spec.Field{Kind: "pair", Key: f.Name.Text, Type: ft, Desc: f.Doc, Optional: opt || f.Default != nil})
	}
	return t
}

func unwrapOptional(t syntax.TypeExpr) syntax.TypeExpr {
	if o, ok := t.(*syntax.TypeOptional); ok {
		return o.X
	}
	return t
}
