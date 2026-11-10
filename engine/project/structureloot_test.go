package project

import (
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// lootPiece builds a structure file with a chest and a barrel, each with
// a loot table, and a chest minecart with one.
func lootPiece(t *testing.T, chest, barrel, cart string) []byte {
	t.Helper()
	palette := []any{
		map[string]any{"Name": "minecraft:chest", "Properties": map[string]any{"facing": "north"}},
		map[string]any{"Name": "minecraft:barrel"},
	}
	blocks := []any{
		map[string]any{"pos": []any{int32(0), int32(0), int32(0)}, "state": int32(0), "nbt": map[string]any{"id": "minecraft:chest", "LootTable": chest}},
		// No id in the block entity: the palette names the block.
		map[string]any{"pos": []any{int32(1), int32(0), int32(0)}, "state": int32(1), "nbt": map[string]any{"LootTable": barrel}},
	}
	entities := []any{map[string]any{
		"pos": []any{2.5, 0.5, 0.5}, "blockPos": []any{int32(2), int32(0), int32(0)},
		"nbt": map[string]any{"id": "minecraft:chest_minecart", "LootTable": cart},
	}}
	data, err := encodeStructure(map[string]any{"DataVersion": int32(4440), "size": []any{int32(3), int32(1), int32(1)}, "palette": palette, "blocks": blocks, "entities": entities})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func lootDiags(r *Result) string {
	var sb strings.Builder
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/structure-unknown-loot-table" {
			sb.WriteString(d.Message + "\n")
		}
	}
	return sb.String()
}

func TestStructureLootTable(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png": {Data: pngBytes},
		"s/house.nbt":  {Data: lootPiece(t, "demo:chests/hous", "chests/simple_dungon", "other:chests/cart")},
	}
	src := "LootTable chests/house {\n  pools = []\n}\n" + pool("starts", `StructureFile("s/house.nbt", "house")`)
	_, r := assetProject(t, files, src)
	want := "'s/house.nbt' has a chest at 0 0 0 with the loot table 'demo:chests/hous', which no LootTable declares (did you mean 'demo:chests/house'?), so it is empty when opened\n" +
		"'s/house.nbt' has a barrel at 1 0 0 with the loot table 'minecraft:chests/simple_dungon', which is not a vanilla loot table (did you mean 'minecraft:chests/simple_dungeon'?), so it is empty when opened\n"
	if got := lootDiags(r); got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}

	// Declared, vanilla and other packs' tables are fine; so is a minecart's.
	// (Each file gets its own time: structure files are cached by size and time.)
	files["s/house.nbt"] = &fstest.MapFile{Data: lootPiece(t, "demo:chests/house", "minecraft:chests/simple_dungeon", "demo:chests/house"), ModTime: time.Unix(1, 0)}
	_, r = assetProject(t, files, src)
	if got := lootDiags(r); got != "" {
		t.Fatalf("known tables: got %s", got)
	}

	// The minecart's missing table is reported at the block it is in, and
	// the fix declares it.
	files["s/house.nbt"] = &fstest.MapFile{Data: lootPiece(t, "demo:chests/house", "chests/simple_dungeon", "demo:carts/supply"), ModTime: time.Unix(2, 0)}
	_, r = assetProject(t, files, src)
	want = "'s/house.nbt' has a chest_minecart at 2 0 0 with the loot table 'demo:carts/supply', which no LootTable declares, so it is empty when opened\n"
	if got := lootDiags(r); got != want {
		t.Fatalf("minecart: got:\n%s\nwant:\n%s", got, want)
	}
	var fixed string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/structure-unknown-loot-table" {
			if len(d.Fixes) == 0 || d.Fixes[0].Title != "Declare LootTable carts/supply" {
				t.Fatalf("fixes: %+v", d.Fixes)
			}
			fixed = applyEdits(src, d.Fixes[0].Edits)
		}
	}
	_, r = assetProject(t, files, fixed)
	if got := lootDiags(r); got != "" {
		t.Fatalf("after fix: got %s", got)
	}
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			t.Fatalf("after fix, error: %s", d.Message)
		}
	}
}
