package spec

import (
	"sort"
	"strings"
	"sync"
)

// short indexes the schema's types by the last part of their path, built
// once per pack.
var (
	shortMu    sync.Mutex
	shortIndex = map[*Pack]map[string]string{}
)

// ShortNames maps the names MMS code can write for schema types, such as
// ClimateParameter, to their paths. When two paths end in the same name,
// the shortest wins, then the first in order: Text is
// ::java::util::text::Text, not a nested one.
func (p *Pack) ShortNames() map[string]string {
	shortMu.Lock()
	defer shortMu.Unlock()
	if m, ok := shortIndex[p]; ok {
		return m
	}
	byName := map[string][]string{}
	for path := range p.Types {
		name := ShortPath(path)
		byName[name] = append(byName[name], path)
	}
	m := map[string]string{}
	for name, paths := range byName {
		sort.Slice(paths, func(i, j int) bool {
			if len(paths[i]) != len(paths[j]) {
				return len(paths[i]) < len(paths[j])
			}
			return paths[i] < paths[j]
		})
		m[name] = paths[0]
	}
	shortIndex[p] = m
	return m
}

// TypeArea is the part of the schema a type belongs to, which the
// reference's type pages are split by: "worldgen", "loot", "entity",
// "assets" and so on. Small data modules share "data".
func TypeArea(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "::java::"), "::")
	if len(parts) < 2 {
		return "data"
	}
	switch parts[0] {
	case "assets":
		return "assets"
	case "util":
		return "util"
	case "world":
		if len(parts) > 2 {
			return parts[1] // entity, block, component, item
		}
		return "world"
	case "data":
		if len(parts) > 2 {
			switch parts[1] {
			case "worldgen", "advancement", "enchantment", "recipe", "dialog", "variants":
				return parts[1]
			case "loot", "number_provider", "item_modifier", "predicate", "slot_source":
				return "loot"
			}
		}
	}
	return "data"
}
