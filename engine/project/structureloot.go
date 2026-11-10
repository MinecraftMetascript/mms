package project

import (
	"fmt"
	"strings"

	"github.com/MinecraftMetascript/nbt-go"
	"github.com/MinecraftMetascript/nbt-go/structure"

	"github.com/minecraftmetascript/mms/engine/internal/near"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Loot tables in structure files: a chest, barrel, chest minecart or
// other container saved with a LootTable tag rolls that table the first
// time a player opens it. The game says nothing when the table does not
// exist; the container is just empty.

// LootRef is a LootTable tag of a block entity or entity in a structure
// file.
type LootRef struct {
	Pos [3]int // the block, or the block an entity is in
	// What is the block or entity id, such as minecraft:chest.
	What string
	// ID is the loot table, with the minecraft namespace added when the
	// tag has none.
	ID string
}

// LootTables lists the LootTable tags of a structure file's block
// entities and entities, blocks first, in the order the file lists them.
func LootTables(st *structure.Structure) []LootRef {
	palette := st.Palette
	if len(palette) == 0 && len(st.Palettes) > 0 {
		palette = st.Palettes[0]
	}
	var out []LootRef
	for _, b := range st.Blocks {
		id := lootID(b.NBT)
		if id == "" {
			continue
		}
		what := nbtString(b.NBT["id"])
		if what == "" && b.State >= 0 && int(b.State) < len(palette) {
			what = palette[b.State].Name
		}
		out = append(out, LootRef{Pos: [3]int{int(b.Pos[0]), int(b.Pos[1]), int(b.Pos[2])}, What: what, ID: id})
	}
	for _, e := range st.Entities {
		id := lootID(e.NBT)
		if id == "" {
			continue
		}
		out = append(out, LootRef{Pos: [3]int{int(e.BlockPos[0]), int(e.BlockPos[1]), int(e.BlockPos[2])}, What: nbtString(e.NBT["id"]), ID: id})
	}
	return out
}

// lootID reads a LootTable tag, adding the minecraft namespace, or "".
func lootID(data nbt.Compound) string {
	s := nbtString(data["LootTable"])
	if s == "" {
		return ""
	}
	if !strings.Contains(s, ":") {
		s = "minecraft:" + s
	}
	return s
}

// lintStructureLoot reports a LootTable tag in a structure file that
// names a loot table nothing provides: in the minecraft namespace, one
// that is neither vanilla nor declared; in another namespace the project
// declares in, one it does not declare. Other namespaces may come from
// another pack and are not checked.
func lintStructureLoot(l *lintRun) {
	m := l.jigsaws()
	var declared map[string]bool // loot table ids the project declares
	for _, s := range m.uses {
		reported := map[string]bool{}
		for _, ref := range s.loot {
			ns, path, _ := strings.Cut(ref.ID, ":")
			if reported[ref.ID] || l.r.Syms.Find("loot_table", ref.ID) != nil {
				continue
			}
			if declared == nil {
				declared = map[string]bool{}
				for _, sym := range l.r.Syms.All() {
					if sym.Kind.Folder() == "loot_table" {
						declared[sym.ID()] = true
					}
				}
			}
			var hint string
			switch {
			case ns == "minecraft":
				if l.r.pack == nil || !l.r.pack.HasRegistry("loot_table") || l.r.pack.HasID("loot_table", path) {
					continue
				}
				hint = "which is not a vanilla loot table"
				if near := near.Closest(path, l.r.pack.Registries["loot_table"]); near != "" {
					hint += fmt.Sprintf(" (did you mean 'minecraft:%s'?)", near)
				}
			case m.declared[ns]:
				hint = "which no LootTable declares"
				if near := near.ClosestKey(ref.ID, declared); near != "" {
					hint += fmt.Sprintf(" (did you mean '%s'?)", near)
				}
			default:
				continue
			}
			reported[ref.ID] = true
			what := strings.TrimPrefix(ref.What, "minecraft:")
			if what == "" {
				what = "container"
			}
			d := syntax.Diagnostic{
				Code: "lint/" + l.rule.ID, Severity: l.sev, Sp: s.sp,
				Message: fmt.Sprintf("'%s' has a %s at %s with the loot table '%s', %s, so it is empty when opened", s.Src, what, posText(ref.Pos), ref.ID, hint),
			}
			if ns != "minecraft" {
				d.Fixes = []syntax.Fix{l.r.newLootTableFix(s, ref.ID)}
			}
			l.out[s.decl.Unit.Path] = append(l.out[s.decl.Unit.Path], d)
		}
	}
}

// newLootTableFix adds an empty LootTable at the end of the file of the
// StructureFile call: the bare name when the file ends in the table's
// namespace, else the full id.
func (r *Result) newLootTableFix(s *structureUse, id string) syntax.Fix {
	ns, name, _ := strings.Cut(id, ":")
	decl := name
	if endNamespace(r, s.decl.Unit.File) != ns {
		decl = id
	}
	src := s.decl.Unit.File.Src
	text := fmt.Sprintf("\nLootTable %s {\n  pools = []\n}\n", decl)
	if len(src) > 0 && src[len(src)-1] != '\n' {
		text = "\n" + text
	}
	at := len(src)
	return syntax.Fix{Title: fmt.Sprintf("Declare LootTable %s", decl), Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: at, End: at}, Text: text}}}
}

// endNamespace is the namespace in force at the end of a file: the last
// Namespace declaration's, or the project's.
func endNamespace(r *Result, f *syntax.File) string {
	end := r.cfg.Namespace
	for _, d := range f.Decls {
		if d.Kind != nil && d.Kind.Name == "Namespace" && d.Name != nil {
			end = d.Name.Text
		}
	}
	return end
}
