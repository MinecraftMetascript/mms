package mcp

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// The preview tools run the embedded script with a JavaScript runtime;
// without one they say how to get one, and these tests skip.
func needJS(t *testing.T) {
	t.Helper()
	if _, _, err := jsRuntime(); err != nil {
		t.Skip("no JavaScript runtime:", err)
	}
}

func TestPreviewJigsaw(t *testing.T) {
	needJS(t)
	s := server(t)
	dir := "../../examples/jigsaw"
	text, failed := call(t, s, "jigsaw", map[string]any{"dir": dir})
	if failed {
		t.Fatal(text)
	}
	wantAll(t, text, "Jigsaw structure dungeon:dungeon", "Seed 0.", "dungeon:start", "Pieces (index, depth")
	text, failed = call(t, s, "jigsaw", map[string]any{"dir": dir, "runs": 10, "seed": 5})
	if failed {
		t.Fatal(text)
	}
	wantAll(t, text, "Seeds 5..14 (10 runs).", "Pieces per run:", "dungeon:start  ")
	if text, failed := call(t, s, "jigsaw", map[string]any{"dir": dir, "structure": "nope"}); !failed {
		t.Errorf("an unknown structure should fail: %s", text)
	}
}

func TestPreviewBiomes(t *testing.T) {
	needJS(t)
	s := server(t)
	dir := "../../examples/cave-siege"
	text, failed := call(t, s, "biomes", map[string]any{"dir": dir})
	if !failed {
		t.Errorf("two dimensions: should ask which: %s", text)
	}
	wantAll(t, text, "siege:caves")
	text, failed = call(t, s, "biomes", map[string]any{"dir": dir, "dimension": "caves", "biome": "deep_dark", "radius": 512})
	if failed {
		t.Fatal(text)
	}
	// At y 64 depth is 0.25, out of deep_dark's 1.1..1.5.
	wantAll(t, text, "Dimension siege:caves", "Biomes in the area", "Nearest entries", "minecraft:deep_dark wins nowhere", "depth: area 0.25..0.25", "never in range here", "try another y")
}

func TestPreviewDensity(t *testing.T) {
	needJS(t)
	s := server(t)
	dir := "../../examples/cave-siege"
	text, failed := call(t, s, "density", map[string]any{"dir": dir, "id": "minecraft:overworld", "mode": "column"})
	if failed {
		t.Fatal(text)
	}
	wantAll(t, text, "noise settings minecraft:overworld", "the column at x 0, z 0", "Crosses 0")
	text, failed = call(t, s, "density", map[string]any{"dir": dir, "id": "caves"})
	if !failed {
		t.Errorf("siege:caves is several kinds: should ask which: %s", text)
	}
	text, failed = call(t, s, "density", map[string]any{"dir": dir, "id": "caves", "kind": "noise_settings", "radius": 64, "size": 4})
	if failed {
		t.Fatal(text)
	}
	wantAll(t, text, "noise settings siege:caves", "above 0 at", "Values on a 4×4 grid")
	text, failed = call(t, s, "density", map[string]any{"dir": dir, "id": "minecraft:overworld", "mode": "surface", "radius": 8, "size": 2})
	if failed {
		t.Fatal(text)
	}
	wantAll(t, text, "Sea level is 63", "Top blocks:")
}

// The embedded script must be what packages/preview builds; skipped
// where the npm workspace isn't installed (CI checks it in the web job).
func TestAgentInSync(t *testing.T) {
	node, err := exec.LookPath("node")
	if err != nil {
		t.Skip("no node")
	}
	if err := exec.Command(node, "-e", "require.resolve('esbuild', {paths: ['../../packages/preview']})").Run(); err != nil {
		t.Skip("the npm workspace is not installed")
	}
	out := t.TempDir() + "/agent.mjs"
	if b, err := exec.Command(node, "../../packages/preview/scripts/build-agent.mjs", out).CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, b)
	}
	built, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(built, agentJS) {
		t.Fatal("engine/mcp/preview/agent.mjs is stale: run `npm run build:agent -w packages/preview`")
	}
}

// A note only some runs give, such as a pool only some starts reach, is
// reported over many runs with how many gave it.
func TestPreviewJigsawNotesFromEveryRun(t *testing.T) {
	needJS(t)
	s := server(t)
	dir := t.TempDir()
	if err := os.CopyFS(dir, os.DirFS("../../examples/jigsaw")); err != nil {
		t.Fatal(err)
	}
	src, err := os.ReadFile(filepath.Join(dir, "dungeon.mms"))
	if err != nil {
		t.Fatal(err)
	}
	// Some runs start with a cap, which leads nowhere, so only the others
	// reach the halls pool, which lists a vanilla piece the simulation lacks.
	text := strings.Replace(string(src), `start = [Piece("start")]`, `start = [Piece("start"), Piece("caps/hall_cap")]`, 1)
	text = strings.Replace(text, `Piece("halls/short_hall"),`, `Piece("halls/short_hall"), WeightedElement(1, SinglePoolElement(rigid, minecraft:village/plains/houses/plains_small_house_1, minecraft:empty)),`, 1)
	if err := os.WriteFile(filepath.Join(dir, "dungeon.mms"), []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	out, failed := call(t, s, "jigsaw", map[string]any{"dir": dir, "runs": 20, "depth": 1})
	if failed {
		t.Fatal(out)
	}
	if strings.Contains(out, "has errors") {
		t.Fatal(out)
	}
	// Seed 0 starts with the cap, so before, the report showed none of the halls run's notes.
	wantAll(t, out, "Note: minecraft:village/plains/houses/plains_small_house_1 is not one of this project's structure files, so it is never placed here. (in 9 of 20 runs)")
}
