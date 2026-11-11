package mcp

import (
	"os/exec"
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
	b, err := exec.Command("cmp", out, "preview/agent.mjs").CombinedOutput()
	if err != nil {
		t.Fatalf("engine/mcp/preview/agent.mjs is stale: run `npm run build:agent -w packages/preview`\n%s", b)
	}
}
