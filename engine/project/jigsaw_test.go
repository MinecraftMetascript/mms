package project

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/nbt"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// piece builds a structure file with a stone block and the given jigsaw
// blocks, each as pool|name|target.
func piece(t *testing.T, jigsaws ...string) []byte {
	t.Helper()
	palette := []any{
		map[string]any{"Name": "minecraft:stone"},
		map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": "north_up"}},
	}
	blocks := []any{map[string]any{"pos": []any{int32(0), int32(0), int32(0)}, "state": int32(0)}}
	for i, j := range jigsaws {
		f := strings.Split(j, "|")
		blocks = append(blocks, map[string]any{
			"pos": []any{int32(i + 1), int32(0), int32(0)}, "state": int32(1),
			"nbt": map[string]any{"id": "minecraft:jigsaw", "pool": f[0], "name": f[1], "target": f[2], "final_state": "minecraft:air", "joint": "rollable"},
		})
	}
	data, err := nbt.Encode(map[string]any{"DataVersion": int32(4440), "size": []any{int32(3), int32(1), int32(1)}, "palette": palette, "blocks": blocks, "entities": []any{}})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func pool(name string, locations ...string) string {
	var els []string
	for _, l := range locations {
		els = append(els, `{ weight = 1, element = { element_type = "minecraft:single_pool_element", location = `+l+`, projection = "rigid", processors = "minecraft:empty" } }`)
	}
	return "TemplatePool " + name + " {\n  fallback = minecraft:empty\n  elements = [" + strings.Join(els, ", ") + "]\n}\n"
}

func jigsawDiags(r *Result) string {
	var sb strings.Builder
	for _, d := range r.Diagnostics("main.mms") {
		if strings.HasPrefix(d.Code, "lint/jigsaw") {
			sb.WriteString(d.Code + ": " + d.Message + "\n")
		}
	}
	return sb.String()
}

func TestJigsawPoolMissing(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png": {Data: pngBytes},
		"s/base.nbt":   {Data: piece(t, "demo:tops|demo:bottom|demo:top", "minecraft:village/plains/streets|minecraft:a|minecraft:b", "other:x|demo:y|demo:z")},
		"s/notnbt.nbt": {Data: nbtBytes},
	}
	_, r := assetProject(t, files, pool("starts", `StructureFile("s/base.nbt", "base")`)+pool("junk", `StructureFile("s/notnbt.nbt")`))
	got := jigsawDiags(r)
	want := "lint/jigsaw-unknown-pool: 's/base.nbt' has a jigsaw block at 1 0 0 that draws from the pool 'demo:tops', which no TemplatePool declares; the structure stops growing there\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	// The fix declares the pool, which then has no pieces.
	var fixed string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/jigsaw-unknown-pool" {
			if len(d.Fixes) == 0 || d.Fixes[0].Title != "Declare TemplatePool tops" {
				t.Fatalf("fixes: %+v", d.Fixes)
			}
			src := string(r.Unit("main.mms").File.Src)
			fixed = src + d.Fixes[0].Edits[0].Text
		}
	}
	_, r = assetProject(t, files, fixed)
	got = jigsawDiags(r)
	want = "lint/jigsaw-unreachable-target: 's/base.nbt' has a jigsaw block at 1 0 0 that draws from the pool 'demo:tops', which has no structure pieces yet, so nothing joins there\n"
	if got != want {
		t.Fatalf("after fix, got:\n%s\nwant:\n%s", got, want)
	}
}

func TestJigsawTarget(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png": {Data: pngBytes},
		"s/base.nbt":   {Data: piece(t, "demo:tops|demo:bottom|demo:top", "demo:tops|demo:side|demo:sideways")},
		"s/top.nbt":    {Data: piece(t, "minecraft:empty|demo:top|minecraft:empty")},
		"s/roof.nbt":   {Data: piece(t, "minecraft:empty|demo:roof|minecraft:empty")},
	}
	src := pool("starts", `StructureFile("s/base.nbt", "base")`) + pool("tops", `StructureFile("s/top.nbt", "top")`, `StructureFile("s/roof.nbt", "roof")`)
	_, r := assetProject(t, files, src)
	got := jigsawDiags(r)
	want := "lint/jigsaw-unreachable-target: 's/base.nbt' has a jigsaw block at 2 0 0 that targets 'demo:sideways' in the pool 'demo:tops', but no piece of that pool has a jigsaw block named 'demo:sideways', so the pieces never connect (the pieces' jigsaw names: demo:roof, demo:top)\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	// A pool with a piece from elsewhere is not checked: its jigsaws are unknown.
	src = pool("starts", `StructureFile("s/base.nbt", "base")`) + pool("tops", `StructureFile("s/top.nbt", "top")`, `minecraft:shipwreck/rightsideup_backhalf`)
	_, r = assetProject(t, files, src)
	if got := jigsawDiags(r); got != "" {
		t.Fatalf("unknown piece: got %s", got)
	}
}

// A jigsaw block carries its orientation and joint; without a joint tag
// the game's default applies: aligned sideways, rollable up or down.
func TestJigsawOrientationAndJoint(t *testing.T) {
	palette := []any{
		map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": "west_up"}},
		map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": "up_north"}},
	}
	jig := func(x, state int32, joint string) any {
		data := map[string]any{"id": "minecraft:jigsaw", "pool": "demo:a", "name": "demo:b", "target": "demo:c"}
		if joint != "" {
			data["joint"] = joint
		}
		return map[string]any{"pos": []any{x, int32(0), int32(0)}, "state": state, "nbt": data}
	}
	got := Jigsaws(map[string]any{"palette": palette, "blocks": []any{jig(0, 0, ""), jig(1, 1, ""), jig(2, 1, "aligned")}})
	want := []string{"west_up aligned", "up_north rollable", "up_north aligned"}
	if len(got) != len(want) {
		t.Fatalf("got %v", got)
	}
	for i, j := range got {
		if s := j.Orientation + " " + j.Joint; s != want[i] {
			t.Errorf("jigsaw %d = %q, want %q", i, s, want[i])
		}
	}
}

// orientedPiece builds a structure file of the given size whose jigsaw
// blocks are each x,y,z|orientation|pool|name|target.
func orientedPiece(t *testing.T, size [3]int32, jigsaws ...string) []byte {
	t.Helper()
	var palette, blocks []any
	for i, j := range jigsaws {
		f := strings.Split(j, "|")
		var x, y, z int32
		fmt.Sscanf(f[0], "%d,%d,%d", &x, &y, &z)
		palette = append(palette, map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": f[1]}})
		blocks = append(blocks, map[string]any{"pos": []any{x, y, z}, "state": int32(i),
			"nbt": map[string]any{"id": "minecraft:jigsaw", "pool": f[2], "name": f[3], "target": f[4]}})
	}
	data, err := nbt.Encode(map[string]any{"size": []any{size[0], size[1], size[2]}, "palette": palette, "blocks": blocks})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// Up joins only down and sideways only sideways; an inset jigsaw is
// internal, and a room too big to fit inside the hall never joins it.
func TestJigsawFacingAndSide(t *testing.T) {
	files := fstest.MapFS{
		"s/hall.nbt": {Data: orientedPiece(t, [3]int32{5, 5, 9},
			"2,1,8|south_up|demo:rooms|demo:exit|demo:entry",    // ok: rooms has a sideways entry
			"2,4,4|up_north|demo:rooms|demo:exit|demo:entry",    // up, but every entry is sideways
			"2,0,4|down_north|demo:cellar|demo:exit|demo:hatch", // ok: cellar's hatch faces up
			"0,1,4|east_up|demo:rooms|demo:exit|demo:entry",     // faces east on the west side
		)},
		"s/room.nbt":   {Data: orientedPiece(t, [3]int32{7, 5, 7}, "3,1,0|north_up|minecraft:empty|demo:entry|minecraft:empty")},
		"s/cellar.nbt": {Data: orientedPiece(t, [3]int32{5, 4, 5}, "2,3,2|up_south|minecraft:empty|demo:hatch|minecraft:empty")},
	}
	src := pool("halls", `StructureFile("s/hall.nbt", "hall")`) + pool("rooms", `StructureFile("s/room.nbt", "room")`) + pool("cellar", `StructureFile("s/cellar.nbt", "cellar")`)
	_, r := assetProject(t, files, src)
	got := jigsawDiags(r)
	want := "lint/jigsaw-facing: 's/hall.nbt' has a jigsaw block at 2 4 4 facing up that targets 'demo:entry' in the pool 'demo:rooms', but every jigsaw block named 'demo:entry' there faces sideways (north); it can only join one facing down, so the pieces never connect\n" +
		"lint/jigsaw-fit: 's/hall.nbt' has an internal jigsaw block at 0 1 4 facing east (it is not on the piece's east edge, so the block in front of it is inside the 5×5×9 piece), so a piece from 'demo:rooms' that joins it must fit inside this piece: 'demo:room' (7×5×7) sticks out 3 east, so it is never placed there; nothing fits, so nothing joins there. If it should join outward, move it to the east edge\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

// An arch on a road 6 high, mounted by an up facing jigsaw on the road's
// floor, must be 5 high or less: the game places it inside the road.
func TestJigsawFitArch(t *testing.T) {
	files := fstest.MapFS{
		"s/road.nbt":  {Data: orientedPiece(t, [3]int32{7, 6, 16}, "3,0,8|up_north|demo:arches|demo:mount|demo:arch", "3,1,0|north_up|minecraft:empty|demo:road|minecraft:empty", "3,1,15|south_up|minecraft:empty|demo:road|minecraft:empty")},
		"s/arch5.nbt": {Data: orientedPiece(t, [3]int32{7, 5, 3}, "3,0,1|down_north|minecraft:empty|demo:arch|minecraft:empty")},
		"s/arch6.nbt": {Data: orientedPiece(t, [3]int32{7, 6, 3}, "3,0,1|down_north|minecraft:empty|demo:arch|minecraft:empty")},
	}
	src := pool("roads", `StructureFile("s/road.nbt", "road")`) + pool("arches", `StructureFile("s/arch5.nbt", "arch5")`, `StructureFile("s/arch6.nbt", "arch6")`)
	_, r := assetProject(t, files, src)
	got := jigsawDiags(r)
	want := "lint/jigsaw-fit: 's/road.nbt' has an internal jigsaw block at 3 0 8 facing up (it is not on the piece's top, so the block in front of it is inside the 7×6×16 piece), so a piece from 'demo:arches' that joins it must fit inside this piece: 'demo:arch6' (7×6×3) sticks out 1 up, so it is never placed there. If it should join outward, move it to the top\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

// A doc comment above a template pool element labels it, so the same
// structure with different processors can be told apart.
func TestPoolElementLabels(t *testing.T) {
	files := fstest.MapFS{"art/icon.png": {Data: pngBytes}, "s/cap.nbt": {Data: piece(t)}}
	src := `ProcessorList worn { processors = [] }
TemplatePool caps {
  fallback = minecraft:empty
  elements = [
    { weight = 5, element = { element_type = minecraft:single_pool_element, location = StructureFile("s/cap.nbt", "cap"), projection = rigid, processors = minecraft:empty } },
    /** @label Worn road cap */
    { weight = 1, element = { element_type = minecraft:single_pool_element, location = StructureFile("s/cap.nbt", "cap"), projection = rigid, processors = worn } },
    // a plain comment is not a label
    { weight = 1, element = { element_type = minecraft:single_pool_element, location = StructureFile("s/cap.nbt", "cap"), projection = rigid, processors = worn } },
    /**
     * Heavily worn.
     * More text.
     */
    { weight = 1, element = { element_type = minecraft:single_pool_element, location = StructureFile("s/cap.nbt", "cap"), projection = rigid, processors = worn } },
  ]
}
`
	_, r := assetProject(t, files, src)
	for _, f := range r.Emit() {
		if f.Path == "data/demo/worldgen/template_pool/caps.json" {
			if got := fmt.Sprint(f.Labels); got != "map[/elements/1:Worn road cap /elements/3:Heavily worn.]" {
				t.Fatalf("labels = %s", got)
			}
			return
		}
	}
	t.Fatal("no pool file")
}

// 26.x structure files write palette entries as {id, properties}.
func TestJigsawOrientation26Palette(t *testing.T) {
	root := map[string]any{
		"palette": []any{map[string]any{"id": "minecraft:jigsaw", "properties": map[string]any{"orientation": "up_north"}}},
		"blocks": []any{map[string]any{"pos": []any{int32(2), int32(0), int32(13)}, "state": int32(0),
			"nbt": map[string]any{"id": "minecraft:jigsaw", "pool": "demo:a", "name": "demo:b", "target": "demo:c", "joint": "rollable"}}},
	}
	got := Jigsaws(root)
	if len(got) != 1 || got[0].Orientation != "up_north" || got[0].Joint != "rollable" {
		t.Fatalf("got %+v", got)
	}
}

// A 5 wide T between 7 wide roads: the roads its other arms would grow
// overlap the road it joined, so those arms never grow. A 7 wide T is fine.
func TestJigsawOverlapNarrowJunction(t *testing.T) {
	road := func(x, y, z int, f string) string {
		return fmt.Sprintf("%d,%d,%d|%s_up|demo:roads|demo:road_end|demo:road_end", x, y, z, f)
	}
	// The junction's arms grow straights only, as in a town whose roads
	// never put two junctions side by side.
	arm := func(x, y, z int, f string) string {
		return fmt.Sprintf("%d,%d,%d|%s_up|demo:straights|demo:road_end|demo:road_end", x, y, z, f)
	}
	el := func(loc string) string {
		return `{ weight = 1, element = { element_type = minecraft:single_pool_element, location = ` + loc + `, projection = rigid, processors = minecraft:empty } }`
	}
	src := func(tee string) string {
		return "TemplatePool roads {\n  fallback = caps\n  elements = [" + el(`StructureFile("s/straight.nbt", "straight")`) + ", " + el(`StructureFile("s/`+tee+`.nbt", "`+tee+`")`) + "]\n}\n" +
			"TemplatePool straights {\n  fallback = caps\n  elements = [" + el(`StructureFile("s/straight.nbt", "straight")`) + "]\n}\n" +
			"TemplatePool caps {\n  fallback = minecraft:empty\n  elements = [" + el(`StructureFile("s/cap.nbt", "cap")`) + "]\n}\n"
	}
	files := fstest.MapFS{
		"s/straight.nbt": {Data: orientedPiece(t, [3]int32{7, 6, 7}, road(0, 0, 3, "west"), road(6, 0, 3, "east"))},
		"s/tee.nbt":      {Data: orientedPiece(t, [3]int32{5, 6, 5}, arm(0, 0, 2, "west"), arm(2, 0, 4, "south"), arm(4, 0, 2, "east"))},
		"s/wide.nbt":     {Data: orientedPiece(t, [3]int32{7, 6, 7}, arm(0, 0, 3, "west"), arm(3, 0, 6, "south"), arm(6, 0, 3, "east"))},
		"s/cap.nbt":      {Data: orientedPiece(t, [3]int32{3, 6, 7}, "0,0,3|west_up|minecraft:empty|demo:road_end|minecraft:empty")},
	}
	_, r := assetProject(t, files, src("tee"))
	got := jigsawDiags(r)
	if !strings.Contains(got, "lint/jigsaw-overlap: When 's/tee.nbt' joins 's/straight.nbt' by its jigsaw block at 0 0 2 (onto the one at 0 0 3), its jigsaw block at 2 0 4 ('demo:road_end', drawing from 'demo:straights') can never grow") || strings.Contains(got, "'s/wide.nbt'") || strings.Contains(got, "When 's/straight.nbt'") {
		t.Fatalf("narrow T:\n%s", got)
	}
	_, r = assetProject(t, files, src("wide"))
	if got := jigsawDiags(r); strings.Contains(got, "jigsaw-overlap") {
		t.Fatalf("7 wide T:\n%s", got)
	}
}

// JigsawHole declares a jigsaw name: it writes nothing, and once a
// namespace declares one, a jigsaw name or target there that is not
// declared is reported, with the closest hole and a fix that declares it.
func TestJigsawHole(t *testing.T) {
	files := fstest.MapFS{
		"art/icon.png": {Data: pngBytes},
		"s/base.nbt":   {Data: piece(t, "demo:tops|demo:bottom|demo:top", "demo:tops|demo:botom|demo:top", "minecraft:empty|other:x|minecraft:empty")},
		"s/top.nbt":    {Data: piece(t, "minecraft:empty|demo:top|minecraft:empty")},
	}
	src := pool("starts", `StructureFile("s/base.nbt", "base")`) + pool("tops", `StructureFile("s/top.nbt", "top")`)
	// No holes declared: nothing is checked.
	_, r := assetProject(t, files, src)
	if got := jigsawDiags(r); got != "" {
		t.Fatalf("no holes: got %s", got)
	}
	_, r = assetProject(t, files, src+"/** The underside. */\nJigsawHole bottom\nJigsawHole top\n")
	got := jigsawDiags(r)
	want := "lint/jigsaw-unknown-name: 's/base.nbt' has a jigsaw block at 2 0 0 whose name is 'demo:botom', which is not a declared JigsawHole (did you mean 'demo:bottom'?)\n"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
	for _, f := range r.Emit() {
		if strings.Contains(f.Path, "jigsaw_hole") || strings.Contains(f.Path, "bottom") {
			t.Fatalf("a JigsawHole wrote %s", f.Path)
		}
	}
	var fixed string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/jigsaw-unknown-name" {
			if len(d.Fixes) == 0 || d.Fixes[0].Title != "Declare JigsawHole botom" {
				t.Fatalf("fixes: %+v", d.Fixes)
			}
			fixed = string(r.Unit("main.mms").File.Src) + d.Fixes[0].Edits[0].Text
		}
	}
	_, r = assetProject(t, files, fixed)
	if got := diagText(r); got != "" {
		t.Fatalf("after fix: %s", got)
	}
	// A file that ends in another namespace gets the full id, so the fix
	// still declares the missing hole; minecraft:empty is never reported,
	// even with a minecraft hole declared.
	_, r = assetProject(t, files, src+"JigsawHole bottom\nJigsawHole top\nJigsawHole minecraft:side\nNamespace other\n")
	found := 0
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code != "lint/jigsaw-unknown-name" {
			continue
		}
		if strings.Contains(d.Message, "minecraft:empty") {
			t.Fatalf("reported minecraft:empty: %s", d.Message)
		}
		if len(d.Fixes) == 0 || d.Fixes[0].Edits[0].Text != "\nJigsawHole demo:botom\n" {
			t.Fatalf("fixes: %+v", d.Fixes)
		}
		found++
	}
	if found != 1 {
		t.Fatalf("want 1 jigsaw-unknown-name finding, got %d", found)
	}
}

// A JigsawHole is only a name: a value, body or duplicate is an error.
func TestJigsawHoleErrors(t *testing.T) {
	files := fstest.MapFS{"art/icon.png": {Data: pngBytes}}
	_, r := assetProject(t, files, "JigsawHole a = 3\nJigsawHole b { }\nJigsawHole c\nJigsawHole c\n// A hole may share a name with a pool.\nJigsawHole tops\n"+pool("tops"))
	var got []string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			got = append(got, d.Code+": "+d.Message)
		}
	}
	want := []string{
		"mms/syntax: a JigsawHole is only a name; write `JigsawHole a` with nothing after it",
		"mms/syntax: a JigsawHole is only a name; write `JigsawHole b` with nothing after it",
		"mms/duplicate: JigsawHole 'demo:c' is already declared in main.mms",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("got:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

// A jigsaw block placed from a named item keeps the name among its
// components, as a plain string, a compound or (older) JSON text.
func TestJigsawLabel(t *testing.T) {
	palette := []any{map[string]any{"id": "minecraft:jigsaw", "properties": map[string]any{"orientation": "north_up"}}}
	block := func(x int32, comps map[string]any) any {
		data := map[string]any{"id": "minecraft:jigsaw", "name": "demo:a", "target": "demo:b", "pool": "demo:p"}
		if comps != nil {
			data["components"] = comps
		}
		return map[string]any{"pos": []any{x, int32(0), int32(0)}, "state": int32(0), "nbt": data}
	}
	root := map[string]any{"palette": palette, "blocks": []any{
		block(0, map[string]any{"minecraft:item_name": "Hall exit"}),
		block(1, map[string]any{"minecraft:custom_name": map[string]any{"text": "Road ", "extra": []any{"end"}}}),
		block(2, map[string]any{"minecraft:item_name": `{"text":"Arch"}`}),
		block(3, nil),
	}}
	var got []string
	for _, j := range Jigsaws(root) {
		got = append(got, j.Label)
	}
	if want := []string{"Hall exit", "Road end", "Arch", ""}; fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("labels %q, want %q", got, want)
	}
}

// JigsawStructure writes the structure set, the structure and its pools,
// and the jigsaw rules check the pieces Piece reads.
func TestJigsawStructure(t *testing.T) {
	files := fstest.MapFS{
		"structures/start.nbt": {Data: piece(t, "demo:halls|demo:exit|demo:entry")},
		"structures/hall.nbt":  {Data: piece(t, "demo:rooms|demo:entry|demo:exit")},
	}
	src := `JigsawStructure dungeon {
  biomes = #minecraft:is_overworld
  step = underground_structures
  spacing = 32
  separation = 12
  height = AboveBottom(20)
  depth = 6
  start = start
  pools = {
    start = [Piece("start")]
    halls = { fallback = minecraft:empty, elements = [Piece("hall", 3)] }
  }
}
Function f {
  /locate structure ${dungeon}
}
`
	_, r := assetProject(t, files, src)
	got := jigsawDiags(r)
	if !strings.Contains(got, "lint/jigsaw-unknown-pool: 'structures/hall.nbt' has a jigsaw block at 1 0 0 that draws from the pool 'demo:rooms'") {
		t.Errorf("want the missing rooms pool reported:\n%s", got)
	}
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			t.Errorf("%s: %s", d.Code, d.Message)
		}
		if d.Code == "lint/jigsaw-unknown-pool" && !strings.HasPrefix(src[d.Sp.Start:], `Piece("hall", 3)`) {
			t.Errorf("jigsaw-unknown-pool at %q, want the Piece call", src[d.Sp.Start:d.Sp.End])
		}
	}
	set := emitted(t, r, "data/demo/worldgen/structure_set/dungeon.json")
	for _, want := range []string{`"structure": "demo:dungeon"`, `"spacing": 32`, `"salt": `} {
		if !strings.Contains(set, want) {
			t.Errorf("structure set lacks %s:\n%s", want, set)
		}
	}
	st := emitted(t, r, "data/demo/worldgen/structure/dungeon.json")
	for _, want := range []string{`"type": "minecraft:jigsaw"`, `"start_pool": "demo:start"`, `"size": 6`, `"above_bottom": 20`, `"max_distance_from_center": 80`} {
		if !strings.Contains(st, want) {
			t.Errorf("structure lacks %s:\n%s", want, st)
		}
	}
	if strings.Contains(st, "project_start_to_heightmap") || strings.Contains(st, "terrain_adaptation") {
		t.Errorf("unset fields written:\n%s", st)
	}
	halls := emitted(t, r, "data/demo/worldgen/template_pool/halls.json")
	if !strings.Contains(halls, `"location": "demo:hall"`) || !strings.Contains(halls, `"weight": 3`) {
		t.Errorf("halls pool:\n%s", halls)
	}
	if fn := emitted(t, r, "data/demo/function/f.mcfunction"); !strings.Contains(fn, "locate structure demo:dungeon") {
		t.Errorf("function:\n%s", fn)
	}
	emitted(t, r, "data/demo/structure/hall.nbt")

	// The fix for the missing pool adds it to pools.
	var fixed string
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/jigsaw-unknown-pool" && len(d.Fixes) > 0 {
			fixed = applyEdits(src, d.Fixes[0].Edits)
		}
	}
	if !strings.Contains(fixed, "    halls = { fallback = minecraft:empty, elements = [Piece(\"hall\", 3)] }\n    rooms = []\n  }") {
		t.Errorf("fixed:\n%s", fixed)
	}
	_, r = assetProject(t, files, fixed)
	if got := jigsawDiags(r); strings.Contains(got, "jigsaw-unknown-pool") {
		t.Errorf("after the fix:\n%s", got)
	}

	// On one line too.
	one := strings.Replace(src, "pools = {\n    start = [Piece(\"start\")]\n    halls = { fallback = minecraft:empty, elements = [Piece(\"hall\", 3)] }\n  }", `pools = { start = [Piece("start")], halls = [Piece("hall")] }`, 1)
	_, r = assetProject(t, files, one)
	for _, d := range r.Diagnostics("main.mms") {
		if d.Code == "lint/jigsaw-unknown-pool" {
			if got := applyEdits(one, d.Fixes[0].Edits); !strings.Contains(got, `halls = [Piece("hall")], rooms = [] }`) {
				t.Errorf("one line fixed:\n%s", got)
			}
		}
	}
}

func TestJigsawStructureChecks(t *testing.T) {
	for _, tc := range []struct{ fields, want string }{
		{"spacing = 8\n  separation = 8", "separation must be less than spacing"},
		{"spacing = 32\n  separation = 12\n  depth = 30", "depth is from 1 to 20"},
		{"spacing = 32\n  separation = 12\n  biomes = plains", ""},
		{"spacing = 32\n  separation = 12\n  biomes = [plains, desert]", ""},
		{"spacing = 32\n  separation = 12\n  biomes = 3", "expected"},
	} {
		biomes := "biomes = #minecraft:is_overworld\n  "
		if strings.Contains(tc.fields, "biomes") {
			biomes = ""
		}
		src := "JigsawStructure d {\n  " + biomes + tc.fields + "\n  step = surface_structures\n  height = Constant(Absolute(64))\n  start = d\n  pools = { d = [] }\n}\n"
		_, r := assetProject(t, fstest.MapFS{}, src)
		got := diagText(r)
		if tc.want == "" && got != "" || tc.want != "" && !strings.Contains(got, tc.want) {
			t.Errorf("%s: got %q, want %q", tc.fields, got, tc.want)
		}
	}
}

// applyEdits applies a fix's edits to src.
func applyEdits(src string, edits []syntax.TextEdit) string {
	rs := []rune(src)
	sort.SliceStable(edits, func(i, j int) bool { return edits[i].Sp.Start < edits[j].Sp.Start })
	var sb strings.Builder
	at := 0
	for _, e := range edits {
		sb.WriteString(string(rs[at:e.Sp.Start]))
		sb.WriteString(e.Text)
		at = e.Sp.End
	}
	sb.WriteString(string(rs[at:]))
	return sb.String()
}
