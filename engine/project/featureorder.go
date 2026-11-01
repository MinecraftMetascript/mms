package project

import (
	"fmt"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// featureSteps names the generation steps a biome's features lists run
// in, by index.
var featureSteps = []string{
	"raw_generation", "lakes", "local_modifications", "underground_structures",
	"surface_structures", "strongholds", "underground_ores", "underground_decoration",
	"fluid_springs", "vegetal_decoration", "top_layer_modification",
}

// orderBiome is one biome's features, resolved to placed feature ids.
type orderBiome struct {
	decl  *Decl
	steps [][]string
}

// maxOrderCycles bounds the cycles reported per step and biome source.
const maxOrderCycles = 4

// orderEdge is "from runs right before to" in some biome.
type orderEdge struct{ from, to string }

// checkFeatureOrder reports placed features that biomes sharing a biome
// source list in conflicting orders. The game sorts every feature of a
// biome source into one order per step, keeping each biome's own order
// (the feature sorter: each feature points at the one listed right after
// it, across all biomes), and refuses to load the world when the result
// has a cycle: "Feature order cycle found". Biomes are grouped by the
// biome sources that name them; project biomes no source names are
// checked together. Vanilla biomes are not, as their features are not
// known here.
func (r *Result) checkFeatureOrder() {
	r.orderErr = map[string][]syntax.Diagnostic{}
	tags := map[string][]string{}
	biomeTags := map[string][]string{}
	biomes := map[string]*orderBiome{}
	byID := map[string]*Decl{}
	var biomeOrder []string
	for _, d := range r.Order {
		if id := d.ID(); id != "" {
			byID[id] = d
		}
		switch d.Kind.Folder() {
		case "tags/worldgen/placed_feature":
			tags[d.ID()] = tagIDs(d.Tag)
		case "tags/worldgen/biome":
			biomeTags[d.ID()] = tagIDs(d.Tag)
		case "worldgen/biome":
			if hasErrors(d.Diags) {
				continue
			}
			if o, ok := d.Value.(*jsonv.Object); ok {
				biomes[d.ID()] = &orderBiome{decl: d}
				biomeOrder = append(biomeOrder, d.ID())
				if fs, ok := o.Get("features").(jsonv.Array); ok {
					for _, step := range fs {
						biomes[d.ID()].steps = append(biomes[d.ID()].steps, stepIDs(step))
					}
				}
			}
		}
	}
	if len(biomes) == 0 {
		return
	}
	// Expand placed feature tags in place, in tag order. A tag the project
	// does not declare is dropped: the features before and after it keep
	// their order, which the game's order implies too.
	for _, b := range biomes {
		for i, step := range b.steps {
			b.steps[i] = expandTags(step, tags, map[string]bool{})
		}
	}

	// Group biomes by the biome sources that name them.
	var groups [][]string
	grouped := map[string]bool{}
	seenGroup := map[string]bool{}
	for _, d := range r.Order {
		for _, src := range biomeSources(d.Value) {
			var ids []string
			seen := map[string]bool{}
			var add func(id string, depth int)
			add = func(id string, depth int) {
				if tag, ok := strings.CutPrefix(id, "#"); ok {
					if depth < 8 {
						for _, e := range biomeTags[tag] {
							add(e, depth+1)
						}
					}
					return
				}
				if biomes[id] != nil && !seen[id] {
					seen[id] = true
					ids = append(ids, id)
				}
			}
			for _, id := range sourceBiomes(src, byID, 0) {
				add(id, 0)
			}
			key := strings.Join(ids, ",")
			if len(ids) == 0 || seenGroup[key] {
				continue
			}
			seenGroup[key] = true
			groups = append(groups, ids)
			for _, id := range ids {
				grouped[id] = true
			}
		}
	}
	var rest []string
	for _, id := range biomeOrder {
		if !grouped[id] {
			rest = append(rest, id)
		}
	}
	if len(rest) > 0 {
		groups = append(groups, rest)
	}

	reported := map[string]bool{}
	for _, g := range groups {
		for step := range featureSteps {
			// The game names one cycle; report a few, each with edges the
			// ones before it did not use, so one check shows most fixes.
			used := map[orderEdge]bool{}
			for range maxOrderCycles {
				cyc, who := findOrderCycle(g, biomes, step, used)
				if cyc == nil {
					break
				}
				for i := range who {
					used[orderEdge{cyc[i], cyc[i+1]}] = true
				}
				key := fmt.Sprint(step, cyc)
				if reported[key] {
					continue
				}
				reported[key] = true
				r.reportOrderCycle(biomes, step, cyc, who)
			}
		}
	}
}

// tagIDs returns a tag's entry ids, "#"-prefixed for nested tags.
func tagIDs(a jsonv.Array) []string {
	var out []string
	for _, e := range a {
		switch v := e.(type) {
		case jsonv.String:
			out = append(out, string(v))
		case *jsonv.Object:
			if id, ok := v.Get("id").(jsonv.String); ok {
				out = append(out, string(id))
			}
		}
	}
	return out
}

// stepIDs reads one step of a biome's features: a list of ids, one id,
// or a tag. Inline placed features are left out: each is its own
// feature, so it cannot appear in another biome.
func stepIDs(v jsonv.Value) []string {
	switch v := v.(type) {
	case jsonv.String:
		return []string{string(v)}
	case jsonv.Array:
		var out []string
		for _, e := range v {
			if s, ok := e.(jsonv.String); ok {
				out = append(out, string(s))
			}
		}
		return out
	}
	return nil
}

func expandTags(ids []string, tags map[string][]string, open map[string]bool) []string {
	var out []string
	for _, id := range ids {
		tag, ok := strings.CutPrefix(id, "#")
		if !ok {
			out = append(out, id)
			continue
		}
		if open[tag] {
			continue
		}
		open[tag] = true
		out = append(out, expandTags(tags[tag], tags, open)...)
		delete(open, tag)
	}
	return out
}

// biomeSources finds the biome_source objects in a resource: a
// dimension's generator, or each dimension of a world preset.
func biomeSources(v jsonv.Value) []*jsonv.Object {
	var out []*jsonv.Object
	var walk func(v jsonv.Value)
	walk = func(v jsonv.Value) {
		switch v := v.(type) {
		case *jsonv.Object:
			for _, m := range v.Members {
				if o, ok := m.Value.(*jsonv.Object); ok && m.Key == "biome_source" {
					out = append(out, o)
					continue
				}
				walk(m.Value)
			}
		case jsonv.Array:
			for _, e := range v {
				walk(e)
			}
		}
	}
	walk(v)
	return out
}

// sourceBiomes lists the biome ids a biome source can place: "biome"
// fields (fixed, and each multi_noise entry), "biomes" lists and tags
// (checkerboard), and the entries of a project parameter list named as
// a multi_noise preset.
func sourceBiomes(src *jsonv.Object, byID map[string]*Decl, depth int) []string {
	var out []string
	var walk func(key string, v jsonv.Value)
	walk = func(key string, v jsonv.Value) {
		switch v := v.(type) {
		case jsonv.String:
			switch key {
			case "biome", "biomes":
				out = append(out, string(v))
			case "preset":
				if d := byID[string(v)]; d != nil && depth < 4 && d.Kind.Folder() == "worldgen/multi_noise_biome_source_parameter_list" {
					if o, ok := d.Value.(*jsonv.Object); ok {
						out = append(out, sourceBiomes(o, byID, depth+1)...)
					}
				}
			}
		case *jsonv.Object:
			for _, m := range v.Members {
				walk(m.Key, m.Value)
			}
		case jsonv.Array:
			for _, e := range v {
				walk(key, e)
			}
		}
	}
	walk("", src)
	return out
}

// findOrderCycle builds one step's order graph over a group of biomes,
// as the game does, and returns a cycle of feature ids (first id
// repeated at the end) and, per edge of it, the first biome that lists
// the pair, leaving out the edges in skip. It returns nil when the step can be ordered. The game's
// graph has a node per (step, feature) and also links each step's last
// feature to the next step's first, but those links only go forward, so
// every cycle lies within one step.
func findOrderCycle(group []string, biomes map[string]*orderBiome, step int, skip map[orderEdge]bool) ([]string, []string) {
	next := map[string][]string{}
	by := map[orderEdge]string{}
	var nodes []string
	seen := map[string]bool{}
	for _, id := range group {
		b := biomes[id]
		if step >= len(b.steps) {
			continue
		}
		fs := b.steps[step]
		for i, f := range fs {
			if !seen[f] {
				seen[f] = true
				nodes = append(nodes, f)
			}
			if i+1 < len(fs) {
				e := orderEdge{f, fs[i+1]}
				if _, ok := by[e]; !ok && !skip[e] {
					by[e] = id
					next[f] = append(next[f], fs[i+1])
				}
			}
		}
	}
	// The shortest cycle reads best: try each feature, in listing order,
	// as its start.
	var cycle []string
	for _, start := range nodes {
		prev := map[string]string{start: ""}
		queue := []string{start}
		var end string
		for len(queue) > 0 && end == "" {
			f := queue[0]
			queue = queue[1:]
			for _, n := range next[f] {
				if n == start {
					end = f
					break
				}
				if _, ok := prev[n]; !ok {
					prev[n] = f
					queue = append(queue, n)
				}
			}
		}
		if end == "" {
			continue
		}
		path := []string{start} // backwards: start, end, ..., start
		for f := end; f != start; f = prev[f] {
			path = append(path, f)
		}
		path = append(path, start)
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		if cycle == nil || len(path) < len(cycle) {
			cycle = path
		}
	}
	if cycle == nil {
		return nil, nil
	}
	who := make([]string, len(cycle)-1)
	for i := range who {
		who[i] = by[orderEdge{cycle[i], cycle[i+1]}]
	}
	return cycle, who
}

// reportOrderCycle puts one error on each biome in the cycle, at its list
// for the step.
func (r *Result) reportOrderCycle(biomes map[string]*orderBiome, step int, cyc, who []string) {
	// Rotate so the cycle starts where the biome changes, then group runs
	// of edges from the same biome into one chain.
	n := len(who)
	start := 0
	for i := range who {
		if who[i] != who[(i+n-1)%n] {
			start = i
			break
		}
	}
	type chain struct {
		biome string
		feats []string
	}
	var chains []chain
	var involved []string
	for k := 0; k < n; k++ {
		i := (start + k) % n
		if len(chains) == 0 || chains[len(chains)-1].biome != who[i] {
			chains = append(chains, chain{who[i], []string{cyc[i]}})
			if !contains(involved, who[i]) {
				involved = append(involved, who[i])
			}
		}
		c := &chains[len(chains)-1]
		c.feats = append(c.feats, cyc[i+1])
	}
	stepName := fmt.Sprintf("step %d (%s)", step+1, featureSteps[step])
	for _, id := range involved {
		b := biomes[id]
		name := func(f string) string { return shortID(f, b.decl.NS) }
		var msg string
		if len(involved) == 1 {
			msg = fmt.Sprintf("feature order cycle in %s: this biome lists %s more than once; the game will not load the world. List each feature once per step",
				stepName, name(cyc[0]))
		} else {
			// Start the cycle at this biome's own part.
			first := 0
			for i, c := range chains {
				if c.biome == id {
					first = i
					break
				}
			}
			var parts []string
			for k := range chains {
				c := chains[(first+k)%len(chains)]
				fs := make([]string, len(c.feats))
				for i, f := range c.feats {
					fs[i] = name(f)
				}
				parts = append(parts, fmt.Sprintf("%s runs %s", shortID(c.biome, b.decl.NS), strings.Join(fs, " before ")))
			}
			msg = fmt.Sprintf("feature order cycle in %s: %s; the game will not load a world with these biomes. Features that biomes share must run in the same order in each of them",
				stepName, joinParts(parts))
		}
		d := b.decl
		r.orderErr[d.Unit.Path] = append(r.orderErr[d.Unit.Path], syntax.Diagnostic{
			Code: "mms/feature-order-cycle", Severity: syntax.SeverityError, Sp: stepSpan(d, step), Message: msg,
		})
	}
}

// shortID drops the project's own namespace, as source writes it.
func shortID(id, ns string) string {
	if s, ok := strings.CutPrefix(id, ns+":"); ok {
		return s
	}
	return id
}

// stepSpan finds a biome's list for one step in its source, falling back
// to the features field or the declaration's name.
func stepSpan(d *Decl, step int) syntax.Span {
	sp := d.D.Sp
	if d.D.Name != nil {
		sp = d.D.Name.Sp
	}
	obj, ok := d.D.Value.(*syntax.Object)
	if !ok || obj.Body == nil {
		return sp
	}
	for _, it := range obj.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.Key == nil || f.Key.Text != "features" {
			continue
		}
		if l, ok := f.Value.(*syntax.List); ok && step < len(l.Elems) {
			return l.Elems[step].Span()
		}
		return f.Sp
	}
	return sp
}

// joinParts joins clauses: "a, but b" for two, "a, b, and c" for more.
func joinParts(parts []string) string {
	if len(parts) == 2 {
		return parts[0] + ", but " + parts[1]
	}
	return strings.Join(parts[:len(parts)-1], ", ") + ", and " + parts[len(parts)-1]
}
