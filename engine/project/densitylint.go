package project

import (
	"fmt"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Density function evaluation-count lints. Every node under a noise
// router is computed at every block position it is asked for, unless a
// cache or interpolated node changes that, so these rules walk the
// checked density functions the routers of the project's noise settings
// reach.

// dfNonRef are density function fields whose strings are not density
// function ids.
var dfNonRef = map[string]bool{
	"type": true, "noise": true, "axis": true, "metric": true, "rarity_value_mapper": true,
}

// dfCostly are the node types worth caching or interpolating: noises,
// and splines over them.
var dfCostly = map[string]bool{
	"noise": true, "shifted_noise": true, "weird_scaled_sampler": true,
	"old_blended_noise": true, "end_outer_islands": true, "spline": true,
}

// dfType is a density function object's type without "minecraft:".
func dfType(v jsonv.Value) string {
	o, ok := v.(*jsonv.Object)
	if !ok {
		return ""
	}
	t, _ := o.Get("type").(jsonv.String)
	return strings.TrimPrefix(string(t), "minecraft:")
}

// dfMarker reports whether a node is computed once per position however
// many times it is used: a cache or an interpolated node, which the game
// shares between every use of the same id.
func dfMarker(v jsonv.Value) bool {
	t := dfType(v)
	return t == "cache" || t == "interpolated"
}

// dfModel is the project's density functions by id.
type dfModel struct {
	vals  map[string]jsonv.Value
	decls map[string]*Decl
	// costly is whether an id's function samples a noise or a spline,
	// itself or through uncached functions it references.
	costly map[string]bool
}

func (r *Result) densityModel() *dfModel {
	m := &dfModel{vals: map[string]jsonv.Value{}, decls: map[string]*Decl{}, costly: map[string]bool{}}
	for _, d := range r.Order {
		if d.Kind.Folder() != "worldgen/density_function" || d.Value == nil || hasErrors(d.Diags) {
			continue
		}
		if id := d.ID(); id != "" {
			m.vals[id] = d.Value
			m.decls[id] = d
		}
	}
	done := map[string]bool{}
	var costly func(v jsonv.Value) bool
	costly = func(v jsonv.Value) bool {
		found := false
		m.walk(v, func(v jsonv.Value, _ bool) {
			if dfCostly[dfType(v)] {
				found = true
			}
		}, func(id string, _ bool) {
			if dfMarker(m.vals[id]) {
				return // computed once, wherever it is used
			}
			if !done[id] {
				done[id] = true // cycle guard
				m.costly[id] = costly(m.vals[id])
			}
			found = found || m.costly[id]
		})
		return found
	}
	for id := range m.vals {
		if !done[id] {
			done[id] = true
			m.costly[id] = costly(m.vals[id])
		}
	}
	return m
}

// walk calls node for every object of a density function value, and ref
// for every reference to a project density function, without following
// it. inCache says whether the position is under an inline cache or
// interpolated node of this value.
func (m *dfModel) walk(v jsonv.Value, node func(jsonv.Value, bool), ref func(string, bool)) {
	var rec func(v jsonv.Value, inCache bool)
	rec = func(v jsonv.Value, inCache bool) {
		switch x := v.(type) {
		case jsonv.String:
			if _, ok := m.vals[string(x)]; ok {
				ref(string(x), inCache)
			}
		case jsonv.Array:
			for _, e := range x {
				rec(e, inCache)
			}
		case *jsonv.Object:
			node(x, inCache)
			in := inCache || dfMarker(x)
			for _, mb := range x.Members {
				if !dfNonRef[mb.Key] {
					rec(mb.Value, in)
				}
			}
		}
	}
	rec(v, false)
}

// knownRoot reports whether every reference in a router slot's graph is
// to a project density function, so what it lacks can be told.
func (m *dfModel) knownRoot(v jsonv.Value) bool {
	known := true
	seen := map[string]bool{}
	var rec func(v jsonv.Value)
	rec = func(v jsonv.Value) {
		switch x := v.(type) {
		case jsonv.String:
			id := string(x)
			if _, ok := m.vals[id]; !ok {
				known = false
				return
			}
			if !seen[id] {
				seen[id] = true
				rec(m.vals[id])
			}
		case jsonv.Array:
			for _, e := range x {
				rec(e)
			}
		case *jsonv.Object:
			for _, mb := range x.Members {
				if !dfNonRef[mb.Key] {
					rec(mb.Value)
				}
			}
		}
	}
	rec(v)
	return known
}

func (l *lintRun) densities() *dfModel {
	if l.r.dfs == nil {
		l.r.dfs = l.r.densityModel()
	}
	return l.r.dfs
}

// router is one noise settings declaration's noise_router, by slot.
type router struct {
	decl  *Decl
	slots []jsonv.Member
}

func (r *Result) routers() []router {
	var out []router
	for _, d := range r.Order {
		if d.Kind.Folder() != "worldgen/noise_settings" || hasErrors(d.Diags) {
			continue
		}
		o, ok := d.Value.(*jsonv.Object)
		if !ok {
			continue
		}
		if nr, ok := o.Get("noise_router").(*jsonv.Object); ok {
			out = append(out, router{decl: d, slots: nr.Members})
		}
	}
	return out
}

// routerSpan finds a noise router slot in a noise settings declaration's
// source, falling back to the declaration's name.
func routerSpan(d *Decl, slot string) syntax.Span {
	sp := d.D.Sp
	if d.D.Name != nil {
		sp = d.D.Name.Sp
	}
	field := func(e syntax.Expr, key string) *syntax.FieldStmt {
		o, ok := e.(*syntax.Object)
		if !ok || o.Body == nil {
			return nil
		}
		for _, it := range o.Body.Items {
			if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil && f.Key.Text == key {
				return f
			}
		}
		return nil
	}
	if nr := field(d.D.Value, "noise_router"); nr != nil {
		sp = nr.Key.Sp
		if f := field(nr.Value, slot); f != nil {
			sp = f.Key.Sp
		}
	}
	return sp
}

// evalCounts returns how many times each project density function is
// computed per block position by one router: the sum, over the paths
// that reach it, of the uses along the path. A cache or interpolated
// function counts once however it is reached, and so does a function
// under an inline cache node.
func (m *dfModel) evalCounts(slots []jsonv.Member) map[string]int {
	type edge struct {
		to      string
		inCache bool
	}
	edges := map[string][]edge{}
	for id, v := range m.vals {
		// Under a cache or interpolated node, what it uses is computed
		// once (walk reports those uses as in a cache).
		m.walk(v, func(jsonv.Value, bool) {}, func(to string, c bool) {
			edges[id] = append(edges[id], edge{to, c})
		})
	}
	const root = "\x00root"
	for _, s := range slots {
		m.walk(s.Value, func(jsonv.Value, bool) {}, func(to string, c bool) {
			edges[root] = append(edges[root], edge{to, c})
		})
	}
	// Topological order from the root; a cycle (which the game rejects)
	// is cut where it closes.
	var order []string
	state := map[string]int{}
	var visit func(string)
	visit = func(id string) {
		if state[id] != 0 {
			return
		}
		state[id] = 1
		for _, e := range edges[id] {
			if state[e.to] != 1 {
				visit(e.to)
			}
		}
		state[id] = 2
		order = append(order, id)
	}
	visit(root)
	count := map[string]int{root: 1}
	for i := len(order) - 1; i >= 0; i-- {
		id := order[i]
		n := count[id]
		if id != root && dfMarker(m.vals[id]) && n > 1 {
			n = 1
		}
		for _, e := range edges[id] {
			add := n
			if e.inCache && add > 1 {
				add = 1
			}
			if count[e.to] < 1<<30 {
				count[e.to] += add
			}
		}
	}
	delete(count, root)
	return count
}

func lintDensityUncachedShared(l *lintRun) {
	m := l.densities()
	type hit struct {
		n      int
		router *Decl
	}
	worst := map[string]hit{}
	for _, rt := range l.r.routers() {
		for id, n := range m.evalCounts(rt.slots) {
			if n < 2 || dfMarker(m.vals[id]) || !m.costly[id] {
				continue
			}
			if h, ok := worst[id]; !ok || n > h.n {
				worst[id] = hit{n, rt.decl}
			}
		}
	}
	ids := make([]string, 0, len(worst))
	for id := range worst {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		d, h := m.decls[id], worst[id]
		if d.D.Name == nil || d.D.Value == nil {
			continue
		}
		vsp := d.D.Value.Span()
		msg := fmt.Sprintf("%s is computed %d times per block by the noise router of %s; wrap it in Cache(...) so it is computed once",
			shortID(id, d.NS), h.n, shortID(h.router.ID(), d.NS))
		l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{
			Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: d.D.Name.Sp, Message: msg,
			Fixes: []syntax.Fix{{Title: "Wrap in Cache(...)", Edits: []syntax.TextEdit{
				{Sp: syntax.Span{Start: vsp.Start, End: vsp.Start}, Text: "Cache("},
				{Sp: syntax.Span{Start: vsp.End, End: vsp.End}, Text: ")"},
			}}},
		})
	}
}

// finalDensity returns a router's final_density slot.
func (rt router) finalDensity() (jsonv.Value, bool) {
	for _, s := range rt.slots {
		if s.Key == "final_density" {
			return s.Value, true
		}
	}
	return nil, false
}

// finalFacts walks final_density: whether it has a node of type t, and
// whether it samples a noise or spline anywhere. ok is false when it
// references a density function the project does not declare, whose
// content is not known here.
func (m *dfModel) finalFacts(v jsonv.Value, t string) (has, costly, ok bool) {
	if !m.knownRoot(v) {
		return false, false, false
	}
	seen := map[string]bool{}
	var rec func(v jsonv.Value)
	rec = func(v jsonv.Value) {
		m.walk(v, func(n jsonv.Value, _ bool) {
			switch dfType(n) {
			case t:
				has = true
			}
			if dfCostly[dfType(n)] {
				costly = true
			}
		}, func(id string, _ bool) {
			if !seen[id] {
				seen[id] = true
				rec(m.vals[id])
			}
		})
	}
	rec(v)
	return has, costly, true
}

func lintDensityNoInterpolation(l *lintRun) {
	m := l.densities()
	for _, rt := range l.r.routers() {
		v, ok := rt.finalDensity()
		if !ok {
			continue
		}
		if has, costly, known := m.finalFacts(v, "interpolated"); known && costly && !has {
			d := rt.decl
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{
				Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: routerSpan(d, "final_density"),
				Message: "final_density samples its noise at every block; wrap the terrain in Interpolated(..., 4, 8) as vanilla does, and keep thresholds outside it",
			})
		}
	}
}

func lintDensityNoBeardifier(l *lintRun) {
	m := l.densities()
	for _, rt := range l.r.routers() {
		v, ok := rt.finalDensity()
		if !ok {
			continue
		}
		if has, costly, known := m.finalFacts(v, "beardifier"); known && costly && !has {
			d := rt.decl
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{
				Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: routerSpan(d, "final_density"),
				Message: "final_density has no Beardifier(), so structures that adapt the terrain (villages, strongholds, ancient cities) are not blended in; 26.3 no longer adds it by itself. Add `+ Beardifier()` unless the dimension has none of them",
			})
		}
	}
}

func lintDensityNestedInterpolation(l *lintRun) {
	m := l.densities()
	reported := map[string]bool{} // id, or "" for the noise settings itself
	for _, rt := range l.r.routers() {
		for _, s := range rt.slots {
			seen := map[string]bool{}
			var rec func(v jsonv.Value, owner string, inside bool)
			rec = func(v jsonv.Value, owner string, inside bool) {
				var walk func(v jsonv.Value, inside bool)
				walk = func(v jsonv.Value, inside bool) {
					switch x := v.(type) {
					case jsonv.String:
						id := string(x)
						if _, ok := m.vals[id]; !ok {
							return
						}
						key := fmt.Sprint(id, inside)
						if !seen[key] {
							seen[key] = true
							rec(m.vals[id], id, inside)
						}
					case jsonv.Array:
						for _, e := range x {
							walk(e, inside)
						}
					case *jsonv.Object:
						in := inside
						if dfType(x) == "interpolated" {
							if inside {
								l.reportNested(rt, s.Key, owner, reported)
							}
							in = true
						}
						for _, mb := range x.Members {
							if !dfNonRef[mb.Key] {
								walk(mb.Value, in)
							}
						}
					}
				}
				walk(v, inside)
			}
			rec(s.Value, "", false)
		}
	}
}

func (l *lintRun) reportNested(rt router, slot, owner string, reported map[string]bool) {
	key := owner
	if owner == "" {
		key = rt.decl.ID() + "|" + slot
	}
	if reported[key] {
		return
	}
	reported[key] = true
	msg := "Interpolated inside another Interpolated: the inner one is interpolated again from the outer cell's corners. Interpolate once, around the whole terrain"
	d, sp := rt.decl, routerSpan(rt.decl, slot)
	if owner != "" {
		if od := l.densities().decls[owner]; od != nil && od.D.Name != nil {
			d, sp = od, od.D.Name.Sp
			msg = fmt.Sprintf("%s contains an Interpolated, and is used inside another Interpolated (from the noise router of %s). Interpolate once, around the whole terrain",
				shortID(owner, od.NS), shortID(rt.decl.ID(), od.NS))
		}
	}
	l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{
		Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: sp, Message: msg,
	})
}

// lintDensitySquare reports `x * x` in a density function, which
// computes x twice where Square(x) computes it once.
func lintDensitySquare(l *lintRun, d *Decl) {
	switch d.Kind.Folder() {
	case "worldgen/density_function", "worldgen/noise_settings":
	default:
		return
	}
	src := d.Unit.File.Src
	text := func(e syntax.Expr) string {
		sp := e.Span()
		return string(src[sp.Start:sp.End])
	}
	squash := func(s string) string { return strings.Join(strings.Fields(s), "") }
	syntax.Walk(d.D, func(n syntax.Node) bool {
		b, ok := n.(*syntax.Binary)
		if !ok || b.Op != "*" || b.L == nil || b.R == nil {
			return true
		}
		if _, ok := constValue(b.L); ok {
			return true
		}
		if squash(text(b.L)) != squash(text(b.R)) {
			return true
		}
		inner := b.L
		for {
			p, ok := inner.(*syntax.Paren)
			if !ok {
				break
			}
			inner = p.X
		}
		repl := "Square(" + text(inner) + ")"
		l.report(d, b.Sp, "this computes "+text(inner)+" twice; Square(...) computes it once",
			syntax.Fix{Title: "Replace with " + repl, Edits: []syntax.TextEdit{{Sp: b.Sp, Text: repl}}})
		return true
	})
}
