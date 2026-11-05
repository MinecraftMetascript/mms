package project

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Resetting a dimension deletes the chunks the game saved for it, so the
// next time a player goes there it generates again with the pack's
// current worldgen. The seed stays the same, so the new terrain is what
// the changes make of the same place.

// vanillaDimensions are the dimensions every world has. Their chunks hold
// the player's builds, so they are never reset just because worldgen
// changed; only when named.
var vanillaDimensions = map[string]bool{
	"minecraft:overworld":  true,
	"minecraft:the_nether": true,
	"minecraft:the_end":    true,
}

// IsVanillaDimension reports whether id is the overworld, the nether or
// the end.
func IsVanillaDimension(id string) bool { return vanillaDimensions[id] }

// dimensionChunkDirs are the folders of a dimension's saved chunks:
// blocks, entities and points of interest (beds, workstations, portals).
// The rest of the dimension's folder (raids, forced chunks) is kept.
var dimensionChunkDirs = []string{"region", "entities", "poi"}

// ResetFolder is the folder in a world where ResetDimensions moves the
// chunks it takes out, so a reset can be undone by moving them back. It
// holds only the latest reset.
const ResetFolder = "mms-reset"

// packID is the id a data pack file stands for, such as "sky:meadow" for
// data/sky/worldgen/biome/meadow.json and "#sky:hot" for a biome tag
// data/sky/tags/worldgen/biome/hot.json, with its registry, such as
// "worldgen/biome". ok is false for files outside data/<ns>/.
func packID(p string) (id, registry string, ok bool) {
	parts := strings.Split(p, "/")
	if len(parts) < 4 || parts[0] != "data" {
		return "", "", false
	}
	ns, rest := parts[1], parts[2:]
	tag := ""
	if rest[0] == "tags" {
		tag, rest = "#", rest[1:]
	}
	n := 1
	if rest[0] == "worldgen" {
		n = 2
	}
	if len(rest) <= n {
		return "", "", false
	}
	registry = strings.Join(rest[:n], "/")
	name := strings.Join(rest[n:], "/")
	name = strings.TrimSuffix(name, path.Ext(name))
	return tag + ns + ":" + name, registry, true
}

// DimensionIDs are the dimensions a pack's files declare, such as
// "sky:islands" for data/sky/dimension/islands.json, sorted.
func DimensionIDs(files []OutFile) []string {
	var out []string
	for _, f := range files {
		if id, reg, ok := packID(f.Path); ok && reg == "dimension" && !strings.HasPrefix(id, "#") {
			out = append(out, id)
		}
	}
	sort.Strings(out)
	return out
}

var idPattern = regexp.MustCompile(`^#?[a-z0-9_.-]+:[a-z0-9_./-]+$`)

// jsonIDs are the namespaced ids a JSON value names anywhere in it.
func jsonIDs(v any, out map[string]bool) {
	switch v := v.(type) {
	case string:
		if idPattern.MatchString(v) {
			out[v] = true
		}
	case []any:
		for _, e := range v {
			jsonIDs(e, out)
		}
	case map[string]any:
		for k, e := range v {
			// Keys are ids too in some places, such as a dimension type's
			// attributes; they are never files, so they only cost a lookup.
			if idPattern.MatchString(k) {
				out[k] = true
			}
			jsonIDs(e, out)
		}
	}
}

// packGraph is a data pack seen as ids that name each other: which files
// stand for an id, and which ids a file names.
type packGraph struct {
	files map[string][]string        // id -> pack paths
	reg   map[string]string          // pack path -> registry
	names map[string]map[string]bool // pack path -> ids its JSON names
	json  map[string]any             // pack path -> decoded JSON
}

func newPackGraph(contents map[string][]byte) *packGraph {
	g := &packGraph{files: map[string][]string{}, reg: map[string]string{}, names: map[string]map[string]bool{}, json: map[string]any{}}
	for p, b := range contents {
		id, reg, ok := packID(p)
		if !ok {
			continue
		}
		g.files[id] = append(g.files[id], p)
		g.reg[p] = reg
		names := map[string]bool{}
		if strings.HasSuffix(p, ".json") {
			var v any
			if json.Unmarshal(b, &v) == nil {
				jsonIDs(v, names)
				g.json[p] = v
			}
		}
		g.names[p] = names
	}
	return g
}

// closure is every pack file reachable from the start files by the ids
// they name, and every id named on the way, found or not.
func (g *packGraph) closure(start ...string) (files, ids map[string]bool) {
	files, ids = map[string]bool{}, map[string]bool{}
	queue := append([]string(nil), start...)
	for _, p := range start {
		files[p] = true
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for id := range g.names[p] {
			if ids[id] {
				continue
			}
			ids[id] = true
			for _, q := range g.files[id] {
				// Nothing refers to a dimension; an id that matches one
				// is a dimension type of the same name, such as
				// minecraft:overworld.
				if !files[q] && g.reg[q] != "dimension" {
					files[q] = true
					queue = append(queue, q)
				}
			}
		}
	}
	return files, ids
}

// biomeFields collects the ids under "biome" and "biomes" keys in a JSON
// value: a biome source's biomes, or a structure's.
func biomeFields(v any, out map[string]bool) {
	switch v := v.(type) {
	case []any:
		for _, e := range v {
			biomeFields(e, out)
		}
	case map[string]any:
		for k, e := range v {
			if k == "biome" || k == "biomes" {
				if _, ok := e.(map[string]any); !ok {
					jsonIDs(e, out)
				}
			}
			biomeFields(e, out)
		}
	}
}

// biomes are the biome ids under "biome" and "biomes" keys in a JSON
// value, following biome tags in the pack. vanilla is whether it names
// one the pack does not define, such as a vanilla biome or tag, which
// might hold any vanilla biome.
func (g *packGraph) biomes(v any) (ids map[string]bool, vanilla bool) {
	ids = map[string]bool{}
	named := map[string]bool{}
	biomeFields(v, named)
	seen := map[string]bool{}
	var visit func(id string)
	visit = func(id string) {
		if seen[id] {
			return
		}
		seen[id] = true
		if !strings.HasPrefix(id, "#") {
			ids[id] = true
			if len(g.files[id]) == 0 && strings.HasPrefix(id, "minecraft:") {
				vanilla = true
			}
			return
		}
		found := false
		for _, p := range g.files[id] {
			if g.reg[p] == "worldgen/biome" {
				found = true
				for n := range g.names[p] {
					visit(n)
				}
			}
		}
		if !found {
			vanilla = true
		}
	}
	for id := range named {
		visit(id)
	}
	return ids, vanilla
}

// ChangedDimensions are the dimensions in a build (files) whose terrain
// would come out different from what the pack in the world before it
// (old, from ReadPackDir) made: the dimension's file changed, or a file it
// leads to by ids (its dimension type, noise settings, density functions,
// biomes, their features...), or a structure set that can place a
// structure in one of its biomes. The overworld, nether and end are left
// out even when the pack changes them; see IsVanillaDimension. Sorted.
func ChangedDimensions(old map[string][]byte, files []OutFile) []string {
	all := map[string][]byte{}
	for p, b := range old {
		all[p] = b
	}
	changed := map[string]bool{}
	seen := map[string]bool{}
	for _, f := range files {
		seen[f.Path] = true
		if b, ok := old[f.Path]; !ok || string(b) != string(f.Data) {
			changed[f.Path] = true
		}
		all[f.Path] = f.Data
	}
	for p := range old {
		if !seen[p] {
			changed[p] = true
		}
	}
	if len(changed) == 0 {
		return nil
	}
	g := newPackGraph(all)

	// Structure sets are not named by a dimension: the game tries each in
	// every dimension, where its structures' biomes allow.
	type set struct {
		biomes  map[string]bool
		vanilla bool
	}
	var sets []set
	for p, reg := range g.reg {
		if reg != "worldgen/structure_set" {
			continue
		}
		reach, _ := g.closure(p)
		if !anyIn(reach, changed) {
			continue
		}
		s := set{biomes: map[string]bool{}}
		for q := range reach {
			if g.reg[q] != "worldgen/structure" {
				continue
			}
			if m, ok := g.json[q].(map[string]any); ok {
				b, v := g.biomes(map[string]any{"biomes": m["biomes"]})
				for id := range b {
					s.biomes[id] = true
				}
				s.vanilla = s.vanilla || v
			}
		}
		sets = append(sets, s)
	}

	var out []string
	for _, id := range DimensionIDs(files) {
		if IsVanillaDimension(id) {
			continue
		}
		p := g.files[id]
		var dim string
		for _, q := range p {
			if g.reg[q] == "dimension" {
				dim = q
			}
		}
		reach, _ := g.closure(dim)
		hit := anyIn(reach, changed)
		if !hit && len(sets) > 0 {
			var source any
			if m, ok := g.json[dim].(map[string]any); ok {
				if gen, ok := m["generator"].(map[string]any); ok {
					source = gen["biome_source"]
				}
			}
			biomes, vanilla := g.biomes(source)
			if src, ok := source.(map[string]any); ok && src["preset"] != nil {
				vanilla = true
			}
			for _, s := range sets {
				if (vanilla && s.vanilla) || anyIn(s.biomes, biomes) {
					hit = true
					break
				}
			}
		}
		if hit {
			out = append(out, id)
		}
	}
	return out
}

func anyIn(a, b map[string]bool) bool {
	if len(b) < len(a) {
		a, b = b, a
	}
	for k := range a {
		if b[k] {
			return true
		}
	}
	return false
}

// DimensionFolder is where a world saves a dimension: <world>/dimensions/
// <ns>/<path> since Minecraft 26.1. A vanilla dimension in a world not
// yet opened in 26.1 or later is still in the old place: the world folder
// itself for the overworld, DIM-1 for the nether and DIM1 for the end.
func DimensionFolder(world, id string) (string, error) {
	if err := CheckDimensionID(id); err != nil {
		return "", err
	}
	ns, name, _ := strings.Cut(id, ":")
	dir := filepath.Join(world, "dimensions", ns, filepath.FromSlash(name))
	if _, err := os.Stat(dir); err != nil && IsVanillaDimension(id) {
		legacy := map[string]string{"minecraft:overworld": "", "minecraft:the_nether": "DIM-1", "minecraft:the_end": "DIM1"}[id]
		old := filepath.Join(world, legacy)
		if _, err := os.Stat(filepath.Join(old, "region")); err == nil {
			return old, nil
		}
	}
	return dir, nil
}

// CheckDimensionID says why id cannot name a dimension's folder, if it
// cannot.
func CheckDimensionID(id string) error {
	ns, name, ok := strings.Cut(id, ":")
	bad := !ok || ns == "" || name == "" || strings.HasPrefix(id, "#") || !idPattern.MatchString(id)
	for _, part := range strings.Split(name, "/") {
		bad = bad || part == "" || part == "." || part == ".."
	}
	if bad {
		return fmt.Errorf("%q is not a dimension id, such as mypack:caves", id)
	}
	return nil
}

// DimensionReset is what ResetDimensions did with one dimension.
type DimensionReset struct {
	ID string
	// Moved are the folders taken out, relative to the world, such as
	// dimensions/sky/islands/region. None means the dimension had not been
	// visited yet, so there was nothing to reset.
	Moved []string
}

// ResetDimensions takes the saved chunks of each dimension out of a world
// and moves them to <world>/mms-reset/, after clearing an earlier reset
// out of it. The world must not be open in the game, which would keep
// its loaded chunks and save them again; see WorldInUse.
func ResetDimensions(world string, ids []string) ([]DimensionReset, error) {
	backup := filepath.Join(world, ResetFolder)
	cleared := false
	var out []DimensionReset
	for _, id := range ids {
		dir, err := DimensionFolder(world, id)
		if err != nil {
			return out, err
		}
		r := DimensionReset{ID: id}
		for _, sub := range dimensionChunkDirs {
			from := filepath.Join(dir, sub)
			if st, err := os.Stat(from); err != nil || !st.IsDir() {
				continue
			}
			if !cleared {
				if err := os.RemoveAll(backup); err != nil {
					return out, fmt.Errorf("clearing the last reset: %w", err)
				}
				cleared = true
			}
			rel, err := filepath.Rel(world, from)
			if err != nil {
				return out, err
			}
			to := filepath.Join(backup, rel)
			if err := os.MkdirAll(filepath.Dir(to), 0o755); err != nil {
				return out, err
			}
			if err := os.Rename(from, to); err != nil {
				return out, fmt.Errorf("resetting %s: %w", id, err)
			}
			r.Moved = append(r.Moved, filepath.ToSlash(rel))
		}
		out = append(out, r)
	}
	return out, nil
}
