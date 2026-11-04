package project

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Apply says what a player in the world has to do for a rebuilt pack to
// take effect.
type Apply int

const (
	// ApplyNone: the build wrote the same files as before.
	ApplyNone Apply = iota
	// ApplyReload: /reload picks the changes up.
	ApplyReload
	// ApplyRejoin: the changes are read only when the world is opened, so
	// the player has to save and quit to title and open the world again.
	ApplyRejoin
)

// ApplyPlan is what a world build changed and what that asks of the
// player. Reload and Rejoin name the kinds of content that changed, such
// as "function" or "worldgen/biome", sorted.
type ApplyPlan struct {
	Need   Apply
	Reload []string
	Rejoin []string
}

// Message is the line `mms build --world` prints last. Tools read the
// start of it: "next step: run /reload", "next step: exit and rejoin" or
// "next step: none".
func (a ApplyPlan) Message() string {
	switch a.Need {
	case ApplyRejoin:
		return "next step: exit and rejoin the world; " + strings.Join(a.Rejoin, ", ") + " changed, which Minecraft only reads when the world is opened"
	case ApplyReload:
		return "next step: run /reload in the world; " + strings.Join(a.Reload, ", ") + " changed"
	}
	return "next step: none; the pack is the same as the last build"
}

// ResourcePackKind is the kind PlanApply reports for a changed world
// resource pack.
const ResourcePackKind = "resource pack"

// reloadable are the data pack folders /reload reads again. Everything
// else in a data pack is a registry entry (worldgen, dimensions,
// enchantments, variants, damage types, dialogs...), which a world loads
// once, when it is opened.
var reloadable = map[string]bool{
	"advancement":   true,
	"function":      true,
	"item_modifier": true,
	"loot_table":    true,
	"predicate":     true,
	"recipe":        true,
	// Structure templates are a cache /reload clears; template pools are
	// worldgen and need a rejoin.
	"structure": true,
}

// staticTags are the registries whose tags /reload reads again: the
// built-in ones. Tags of registries a data pack can add to (worldgen,
// enchantment, damage_type...) are loaded with the registry, when the
// world opens.
var staticTags = map[string]bool{
	"block":                  true,
	"entity_type":            true,
	"fluid":                  true,
	"function":               true,
	"game_event":             true,
	"item":                   true,
	"potion":                 true,
	"mob_effect":             true,
	"point_of_interest_type": true,
	"villager_profession":    true,
	"villager_type":          true,
	"attribute":              true,
	"block_entity_type":      true,
	"particle_type":          true,
	"sound_event":            true,
	"stat_type":              true,
	"menu":                   true,
}

// dataKind is the kind of content a data pack path holds, and whether
// /reload is enough for it: "function" for data/ns/function/a.mcfunction,
// "tags/block" for a block tag, "worldgen/biome" for a biome. Files at
// the pack root (pack.mcmeta, pack.png) are "pack.mcmeta" and so on.
func dataKind(p string) (kind string, reload bool) {
	parts := strings.Split(p, "/")
	if len(parts) < 4 || parts[0] != "data" {
		return p, true
	}
	dir := parts[2]
	switch dir {
	case "tags":
		reg := parts[3]
		if reg == "worldgen" && len(parts) > 5 {
			return "tags/worldgen/" + parts[4], false
		}
		return "tags/" + reg, staticTags[reg]
	case "worldgen":
		if len(parts) > 4 {
			return "worldgen/" + parts[3], false
		}
	}
	return dir, reloadable[dir]
}

// ReadPackDir reads the files a data pack folder written by WriteDir
// holds: pack.mcmeta, pack.png and everything under data/. A folder that
// does not exist reads as empty.
func ReadPackDir(dir string) (map[string][]byte, error) {
	out := map[string][]byte{}
	for _, f := range []string{"pack.mcmeta", "pack.png"} {
		b, err := os.ReadFile(filepath.Join(dir, f))
		if err == nil {
			out[f] = b
		} else if !os.IsNotExist(err) {
			return nil, err
		}
	}
	root := filepath.Join(dir, "data")
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			if os.IsNotExist(err) && p == root {
				return nil
			}
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		out[filepath.ToSlash(rel)] = b
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlanApply compares the data pack in the world before a build (old, from
// ReadPackDir) with the files the build writes, and says whether the
// player needs /reload or a rejoin. resourcesChanged is whether the
// world's resource pack changed; Minecraft applies a world's pack only
// when the world is opened.
func PlanApply(old map[string][]byte, files []OutFile, resourcesChanged bool) ApplyPlan {
	reload, rejoin := map[string]bool{}, map[string]bool{}
	note := func(p string) {
		if kind, ok := dataKind(p); ok {
			reload[kind] = true
		} else {
			rejoin[kind] = true
		}
	}
	seen := map[string]bool{}
	for _, f := range files {
		seen[f.Path] = true
		if b, ok := old[f.Path]; !ok || string(b) != string(f.Data) {
			note(f.Path)
		}
	}
	for p := range old {
		if !seen[p] {
			note(p)
		}
	}
	if resourcesChanged {
		rejoin[ResourcePackKind] = true
	}
	plan := ApplyPlan{Reload: sortedKeys(reload), Rejoin: sortedKeys(rejoin)}
	switch {
	case len(rejoin) > 0:
		plan.Need = ApplyRejoin
	case len(reload) > 0:
		plan.Need = ApplyReload
	}
	return plan
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
