package project

import (
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// checkDensityCycles reports density functions that refer back to
// themselves, directly or through others: `DensityFunction a = b` and
// `DensityFunction b = a`. Computing one never ends, so the game can't
// load the world, and the previews would run out of stack.
func (r *Result) checkDensityCycles() {
	r.cycleErr = map[string][]syntax.Diagnostic{}
	decls := map[string]*Decl{}
	var ids []string
	for _, d := range r.Order {
		if d.Kind.Folder() != "worldgen/density_function" || d.Value == nil {
			continue
		}
		if _, dup := decls[d.ID()]; !dup {
			decls[d.ID()] = d
			ids = append(ids, d.ID())
		}
	}
	refs := map[string][]string{}
	for _, id := range ids {
		refs[id] = densityRefs(decls[id].Value, "", decls, nil)
	}
	const (
		visiting = 1
		done     = 2
	)
	state := map[string]int{}
	var path []string
	var visit func(id string)
	visit = func(id string) {
		state[id] = visiting
		path = append(path, id)
		for _, to := range refs[id] {
			switch state[to] {
			case visiting:
				i := 0
				for path[i] != to {
					i++
				}
				r.reportDensityCycle(append(append([]string(nil), path[i:]...), to), decls)
			case 0:
				visit(to)
			}
		}
		path = path[:len(path)-1]
		state[id] = done
	}
	for _, id := range ids {
		if state[id] == 0 {
			visit(id)
		}
	}
}

// densityRefs lists the project's density functions v refers to by id.
// Noises are ids too, but only under a "noise" key.
func densityRefs(v jsonv.Value, key string, decls map[string]*Decl, out []string) []string {
	switch v := v.(type) {
	case jsonv.String:
		if key != "type" && key != "noise" && decls[string(v)] != nil {
			out = append(out, string(v))
		}
	case jsonv.Array:
		for _, e := range v {
			out = densityRefs(e, key, decls, out)
		}
	case *jsonv.Object:
		for _, m := range v.Members {
			out = densityRefs(m.Value, m.Key, decls, out)
		}
	}
	return out
}

// reportDensityCycle reports cyc (a -> b -> a, the first id repeated
// last) on each density function in it.
func (r *Result) reportDensityCycle(cyc []string, decls map[string]*Decl) {
	ring := cyc[:len(cyc)-1]
	for i, id := range ring {
		d := decls[id]
		var chain []string
		for k := 0; k <= len(ring); k++ {
			chain = append(chain, shortID(ring[(i+k)%len(ring)], d.NS))
		}
		msg := "density function " + shortID(id, d.NS) + " refers to itself"
		if len(ring) > 1 {
			msg += " through " + strings.Join(chain[1:len(chain)-1], ", ")
		}
		msg += " (" + strings.Join(chain, " → ") + "), so computing it never ends and the game can't load the world"
		sp := d.D.Sp
		if d.D.Name != nil {
			sp = d.D.Name.Sp
		}
		r.cycleErr[d.Unit.Path] = append(r.cycleErr[d.Unit.Path], syntax.Diagnostic{
			Code: "mms/density-cycle", Severity: syntax.SeverityError, Sp: sp, Message: msg,
		})
	}
}
