package sema

import (
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Kind is what a declaration's leading identifier names.
type Kind struct {
	Name string // as written, e.g. "DensityFunction"
	// Resource is the data-pack resource kind ("worldgen/density_function")
	// for JSON resources, "tags/<registry>" for tags, or "" for MMS kinds.
	Resource string
	// TagOf is the registry a tag kind lists ("block", "worldgen/biome").
	TagOf string
	// MMS is set for language-level kinds: Namespace, Function, Objective,
	// Global and the event hooks (OnLoad, OnTick, OnFirstJoin, OnRespawn,
	// OnAdvancement).
	MMS string
	// Assets is the folder under assets/<ns>/ for resource-pack kinds
	// ("blockstates", "models"), and "" for data-pack kinds.
	Assets string
}

// Dir is the pack-relative folder a declaration's file goes in, before
// its namespace path: "data/<ns>/loot_table" or "assets/<ns>/models".
func (k Kind) Dir(ns string) string {
	if k.Assets != "" {
		return "assets/" + ns + "/" + k.Assets
	}
	return "data/" + ns + "/" + k.Folder()
}

// IsResource reports whether the declaration emits a JSON file.
func (k Kind) IsResource() bool { return k.Resource != "" }

// Folder is the folder under data/<ns>/ that the kind's files go in.
func (k Kind) Folder() string {
	switch k.MMS {
	case "Function":
		return "function"
	case "OnAdvancement":
		// It declares an advancement, so commands such as `advancement
		// grant` find it, and an Advancement of the same name clashes.
		return "advancement"
	}
	return k.Resource
}

var mmsKinds = map[string]bool{
	"Namespace": true, "Function": true, "Objective": true, "Global": true,
	"OnLoad": true, "OnTick": true,
	"OnFirstJoin": true, "OnRespawn": true, "OnAdvancement": true,
}

// kindAliases name resource kinds whose MMS names differ from the folder,
// mostly names the old language used.
var kindAliases = map[string]string{
	"SurfaceRule":      "worldgen/material_rule",
	"SurfaceCondition": "worldgen/material_condition",
	"NoiseSettings":    "worldgen/noise_settings",
	// Resource-pack kinds named after their folder or the component that
	// uses them, rather than mcdoc's block_definition / item_definition.
	"BlockState": "block_definition",
	"ItemModel":  "item_definition",
}

// renamed are resource kinds that only an alias names.
var renamed = map[string]bool{"block_definition": true, "item_definition": true}

// Kinds resolves declaration kinds against one spec pack.
type Kinds struct {
	pack     *spec.Pack
	byName   map[string]Kind
	tagRegs  map[string]string // snake name -> registry
	resource []string
}

// NewKinds indexes the pack's resource kinds and registries.
func NewKinds(p *spec.Pack) *Kinds {
	ks := &Kinds{pack: p, byName: map[string]Kind{}, tagRegs: map[string]string{}}
	for name := range mmsKinds {
		ks.byName[name] = Kind{Name: name, MMS: name}
	}
	// A resource kind is named after its last path segment, unless two
	// kinds share it; then only the alias table can name them.
	seen := map[string]int{}
	all := append(p.ResourceKinds(), p.AssetKinds()...)
	for _, r := range all {
		seen[lastSeg(r)]++
	}
	for _, r := range all {
		if seen[lastSeg(r)] == 1 && !renamed[r] {
			name := Pascal(lastSeg(r))
			if _, taken := ks.byName[name]; !taken {
				ks.byName[name] = Kind{Name: name, Resource: r, Assets: p.Assets[r]}
			}
		}
		ks.resource = append(ks.resource, r)
	}
	for name, r := range kindAliases {
		if _, ok := p.Resources[r]; ok {
			ks.byName[name] = Kind{Name: name, Resource: r, Assets: p.Assets[r]}
		}
	}
	// Tag kinds: any registry, named by its last segment. A registry at
	// the top level wins over a nested one with the same last segment.
	regs := make([]string, 0, len(p.Registries))
	for r := range p.Registries {
		if !strings.HasPrefix(r, "tag/") {
			regs = append(regs, r)
		}
	}
	sort.Slice(regs, func(i, j int) bool {
		di, dj := strings.Count(regs[i], "/"), strings.Count(regs[j], "/")
		if di != dj {
			return di < dj
		}
		return regs[i] < regs[j]
	})
	regs = append(regs, "function") // function tags have no registry
	for _, r := range regs {
		if _, ok := ks.tagRegs[lastSeg(r)]; !ok {
			ks.tagRegs[lastSeg(r)] = r
		}
	}
	return ks
}

// Lookup resolves a kind name such as "LootTable" or "BlockTag".
func (ks *Kinds) Lookup(name string) (Kind, bool) {
	if k, ok := ks.byName[name]; ok {
		return k, true
	}
	if base, ok := strings.CutSuffix(name, "Tag"); ok && base != "" {
		if r, ok := ks.tagRegs[Snake(base)]; ok {
			return Kind{Name: name, Resource: "tags/" + r, TagOf: r}, true
		}
	}
	return Kind{}, false
}

// Names lists every kind name, for completion: MMS kinds, resource kinds
// and the common tag kinds.
func (ks *Kinds) Names() []string {
	var out []string
	for n := range ks.byName {
		out = append(out, n)
	}
	for _, r := range []string{"block", "item", "entity_type", "function", "worldgen/biome", "fluid", "damage_type"} {
		out = append(out, Pascal(lastSeg(r))+"Tag")
	}
	sort.Strings(out)
	return out
}

func lastSeg(s string) string {
	if i := strings.LastIndexByte(s, '/'); i >= 0 {
		return s[i+1:]
	}
	return s
}

// ForFolder returns the kind whose files live in folder (a resource kind
// such as "loot_table", or "tags/<registry>").
func (ks *Kinds) ForFolder(folder string) (Kind, bool) {
	if reg, ok := strings.CutPrefix(folder, "tags/"); ok {
		name := Pascal(lastSeg(reg)) + "Tag"
		if k, ok := ks.Lookup(name); ok && k.TagOf == reg {
			return k, true
		}
		return Kind{}, false
	}
	for _, k := range ks.byName {
		if k.Resource == folder && kindAliases[k.Name] == "" {
			return k, true
		}
	}
	for name, r := range kindAliases {
		if r == folder {
			return ks.byName[name], true
		}
	}
	return Kind{}, false
}
