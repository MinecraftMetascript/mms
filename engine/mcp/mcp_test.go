package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

var update = flag.Bool("update", false, "copy site/ into mcp/docs")

func server(t *testing.T) *Server {
	t.Helper()
	pack, err := spec.Default()
	if err != nil {
		t.Fatal(err)
	}
	return NewServer(pack)
}

// call runs a tool and returns its text and whether it failed.
func call(t *testing.T, s *Server, name string, args map[string]any) (string, bool) {
	t.Helper()
	raw, _ := json.Marshal(args)
	for _, tl := range s.tools {
		if tl.name == name {
			res := s.call(context.Background(), tl, raw)
			text := res["content"].([]map[string]any)[0]["text"].(string)
			failed, _ := res["isError"].(bool)
			return text, failed
		}
	}
	t.Fatalf("no tool %s", name)
	return "", false
}

func wantAll(t *testing.T, text string, parts ...string) {
	t.Helper()
	for _, p := range parts {
		if !strings.Contains(text, p) {
			t.Errorf("missing %q in:\n%s", p, text)
		}
	}
}

// The wire: one JSON message per line in, one per line out, nothing for
// a notification.
func TestProtocol(t *testing.T) {
	s := server(t)
	in := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-03-26","capabilities":{},"clientInfo":{"name":"t","version":"0"}}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized"}`,
		``,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list"}`,
		`{"jsonrpc":"2.0","id":"three","method":"tools/call","params":{"name":"check","arguments":{"source":"Function greet {\n  /say hi\n}\n"}}}`,
		`{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"nope"}}`,
		`{"jsonrpc":"2.0","id":5,"method":"no/such"}`,
		`not json`,
		`{"jsonrpc":"2.0","id":6,"method":"ping"}`,
		`{"jsonrpc":"2.0","id":7,"method":"resources/read","params":{"uri":"mms://docs/guides/functions"}}`,
		`{"jsonrpc":"2.0","id":8,"method":"resources/read","params":{"uri":"mms://docs/nope"}}`,
	}, "\n") // no newline after the last message
	var out bytes.Buffer
	if err := s.Run(context.Background(), strings.NewReader(in), &out); err != nil {
		t.Fatal(err)
	}
	var got []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(out.String()), "\n") {
		var m map[string]any
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			t.Fatalf("not one JSON message per line: %q", line)
		}
		got = append(got, m)
	}
	if len(got) != 9 {
		t.Fatalf("got %d answers, want 9:\n%s", len(got), out.String())
	}
	result := func(i int) map[string]any {
		t.Helper()
		r, ok := got[i]["result"].(map[string]any)
		if !ok {
			t.Fatalf("answer %d has no result: %v", i, got[i])
		}
		return r
	}
	code := func(i int) float64 {
		t.Helper()
		e, ok := got[i]["error"].(map[string]any)
		if !ok {
			t.Fatalf("answer %d is not an error: %v", i, got[i])
		}
		return e["code"].(float64)
	}
	if v := result(0)["protocolVersion"]; v != "2025-03-26" {
		t.Errorf("protocolVersion = %v, want the client's", v)
	}
	tools := result(1)["tools"].([]any)
	names := map[string]bool{}
	for _, tl := range tools {
		m := tl.(map[string]any)
		names[m["name"].(string)] = true
		if m["inputSchema"].(map[string]any)["type"] != "object" {
			t.Errorf("%s: inputSchema is not an object", m["name"])
		}
	}
	for _, n := range []string{"check", "explain", "test", "fix", "fmt", "convert_json", "lookup", "vanilla", "docs"} {
		if !names[n] {
			t.Errorf("tools/list lacks %s", n)
		}
	}
	if got[2]["id"] != "three" {
		t.Errorf("id not echoed: %v", got[2]["id"])
	}
	text := result(2)["content"].([]any)[0].(map[string]any)["text"].(string)
	wantAll(t, text, "ok: no problems in 1 file")
	if c := code(3); c != codeInvalidParams {
		t.Errorf("unknown tool: code %v", c)
	}
	if c := code(4); c != codeMethodNotFound {
		t.Errorf("unknown method: code %v", c)
	}
	if c := code(5); c != codeParseError {
		t.Errorf("bad JSON: code %v", c)
	}
	result(6)
	page := result(7)["contents"].([]any)[0].(map[string]any)
	wantAll(t, page["text"].(string), "# Functions")
	if c := code(8); c != codeNotFound {
		t.Errorf("missing resource: code %v", c)
	}
}

func TestCheck(t *testing.T) {
	s := server(t)
	src := "LootTable gems = LootTable(pools = [LootPool(rolls = 1, entries = [Item(diamond)])])\nRecipe r = Nope()\n"
	text, failed := call(t, s, "check", map[string]any{"source": src})
	if failed {
		t.Fatalf("problems are a result, not a failure:\n%s", text)
	}
	wantAll(t, text,
		"main.mms:1:37: error: 'LootPool' is written Pool [mms/unknown-helper]",
		"^^^^^^^^",
		"safe fix: Write Pool",
		`replace 1:37 to 1:45 with "Pool"`,
		"main.mms:2:",
		"2 errors, 0 warnings in 1 file",
		"1 problem has a safe fix",
	)
	if text, failed := call(t, s, "check", map[string]any{}); !failed {
		t.Errorf("no dir and no source should fail:\n%s", text)
	}
	if text, failed := call(t, s, "check", map[string]any{"sauce": "x"}); !failed || !strings.Contains(text, "sauce") {
		t.Errorf("an unknown argument should fail and be named:\n%s", text)
	}
}

func TestCheckDir(t *testing.T) {
	s := server(t)
	text, failed := call(t, s, "check", map[string]any{"dir": "../../examples/shop"})
	if failed || !strings.Contains(text, "ok: no problems") {
		t.Errorf("examples/shop:\n%s", text)
	}
	if text, failed := call(t, s, "check", map[string]any{"dir": "no/such/folder"}); !failed {
		t.Errorf("a missing folder should fail:\n%s", text)
	}
}

func TestExplain(t *testing.T) {
	s := server(t)
	src := "Global n: int = 0\nFunction bump {\n  n += 2\n  tell(@a, \"now ${n}\")\n  shout(3)\n}\nTemplate shout(times: int): Commands {\n  for i in 1..times {\n    /say ${i}\n  }\n}\nLootTable gems = LootTable(pools = [Pool(rolls = 1, entries = [Item(diamond)])])\n"
	list, _ := call(t, s, "explain", map[string]any{"source": src})
	wantAll(t, list,
		"data/mms/function/bump.mcfunction  (from main.mms:2)",
		"data/mms/loot_table/gems.json  (from main.mms:12)",
		"pack.mcmeta  (generated)",
	)
	one, _ := call(t, s, "explain", map[string]any{"source": src, "name": "bump"})
	wantAll(t, one, "=== data/mms/function/bump.mcfunction", "scoreboard players add", "tellraw @a")
	if strings.Contains(one, "loot_table") {
		t.Errorf("explain bump shows another declaration's file:\n%s", one)
	}
	// A Template call inside the declaration is shown expanded.
	wantAll(t, one, "=== Template shout at main.mms:5 expands to", "/say 3")
	byPath, _ := call(t, s, "explain", map[string]any{"source": src, "name": "loot_table/gems"})
	wantAll(t, byPath, `"minecraft:diamond"`)
	none, _ := call(t, s, "explain", map[string]any{"source": src, "name": "zzz"})
	wantAll(t, none, "Nothing is written", "Function mms:bump")
}

func TestTest(t *testing.T) {
	s := server(t)
	src := "Global n: int = 0\nFunction bump {\n  n += 2\n  /say bumped\n}\nTest right {\n  run bump()\n  expect n == 2\n}\nTest wrong {\n  run bump()\n  expect n == 3\n}\n"
	text, failed := call(t, s, "test", map[string]any{"source": src})
	if failed {
		t.Fatalf("a failing test is a result:\n%s", text)
	}
	wantAll(t, text, "ok    right", "FAIL  wrong", "main.mms:12:3: expected n == 3, but n is 2", "said: ", "1 of 2 tests failed")
	text, _ = call(t, s, "test", map[string]any{"source": src, "run": "^right$"})
	wantAll(t, text, "ok: 1 test passed")
	text, failed = call(t, s, "test", map[string]any{"source": "Function f = 1\n"})
	if !failed || !strings.Contains(text, "no tests run") {
		t.Errorf("errors should stop the tests:\n%s", text)
	}
}

func TestFixAndFmt(t *testing.T) {
	s := server(t)
	text, _ := call(t, s, "fix", map[string]any{"source": "LootTable gems = LootTable(pools = [LootPool(rolls = 1, entries = [Item(diamond)])])\n"})
	wantAll(t, text, "1 problem fixed", "[Pool(rolls")
	text, _ = call(t, s, "fmt", map[string]any{"source": "Function   greet{\n/say hi\n}"})
	if text != "Function greet {\n  /say hi\n}\n" {
		t.Errorf("fmt: %q", text)
	}

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.mms"), []byte("Function   greet{\n/say hi\n}\nLootTable gems = LootTable(pools = [LootPool(rolls = 1, entries = [Item(diamond)])])\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	text, failed := call(t, s, "fix", map[string]any{"dir": dir})
	if failed {
		t.Fatal(text)
	}
	wantAll(t, text, "fixed a.mms", "1 problem fixed in 1 file")
	b, _ := os.ReadFile(filepath.Join(dir, "a.mms"))
	wantAll(t, string(b), "Function greet {\n", "[Pool(rolls")

	if err := os.WriteFile(filepath.Join(dir, "b.mms"), []byte("Function   wave{\n/say hi\n}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	text, _ = call(t, s, "fmt", map[string]any{"dir": dir})
	wantAll(t, text, "formatted b.mms", "1 file formatted")
	b, _ = os.ReadFile(filepath.Join(dir, "b.mms"))
	wantAll(t, string(b), "Function wave {\n")
}

func TestConvertJSON(t *testing.T) {
	s := server(t)
	text, failed := call(t, s, "convert_json", map[string]any{"kind": "tags/block", "name": "soft", "json": `{"values":["minecraft:dirt","#minecraft:sand"]}`})
	if failed || strings.TrimSpace(text) != "BlockTag soft = [dirt, #sand]" {
		t.Errorf("tag: %q", text)
	}
	text, _ = call(t, s, "convert_json", map[string]any{"kind": "LootTable", "name": "x:gems", "json": `{"pools":[{"rolls":1,"entries":[{"type":"minecraft:item","name":"minecraft:diamond"}]}]}`})
	wantAll(t, text, "LootTable x:gems", "diamond")
	// What it writes checks clean.
	if out, _ := call(t, s, "check", map[string]any{"source": text}); !strings.Contains(out, "ok: no problems") {
		t.Errorf("converted source has problems:\n%s\n%s", text, out)
	}
	if text, failed := call(t, s, "convert_json", map[string]any{"kind": "Function", "name": "f", "json": "{}"}); !failed {
		t.Errorf("Function is not a JSON kind:\n%s", text)
	}
	if text, failed := call(t, s, "convert_json", map[string]any{"kind": "LootTable", "name": "f", "json": "{"}); !failed {
		t.Errorf("bad JSON should fail:\n%s", text)
	}
}

func TestLookup(t *testing.T) {
	s := server(t)
	text, _ := call(t, s, "lookup", map[string]any{"query": "customitem"})
	wantAll(t, text, "Blueprint CustomItem", "base: Id<item>", "on_right_click")
	text, _ = call(t, s, "lookup", map[string]any{"query": "Clamp"})
	wantAll(t, text, "helper Clamp(input", "example:", "Also matching:", "Clamped")
	text, _ = call(t, s, "lookup", map[string]any{"query": "noise", "in": "kind"})
	wantAll(t, text, "kind Noise: writes data/<namespace>/worldgen/noise/<name>.json")
	if strings.Contains(text, "helper ") {
		t.Errorf("in=kind lists a helper:\n%s", text)
	}
	text, _ = call(t, s, "lookup", map[string]any{"query": "unused", "in": "rule"})
	wantAll(t, text, "lint rule unused (default: hint)", "unused-asset")
	text, _ = call(t, s, "lookup", map[string]any{"query": "density", "limit": 3})
	wantAll(t, text, "more; narrow the query")
	text, _ = call(t, s, "lookup", map[string]any{"query": "qqqqzz"})
	wantAll(t, text, "Nothing in the reference matches")
}

func TestVanilla(t *testing.T) {
	s := server(t)
	text, _ := call(t, s, "vanilla", nil)
	wantAll(t, text, "  block: ", "  sound_event: ", "  loot_table: ")
	text, _ = call(t, s, "vanilla", map[string]any{"registry": "block", "query": "minecraft:oak_stairs"})
	wantAll(t, text, "minecraft:oak_stairs", "minecraft:dark_oak_stairs", "facing: north, south, west, east (default north)")
	text, _ = call(t, s, "vanilla", map[string]any{"query": "amethyst_shard"})
	wantAll(t, text, "item (1):", "minecraft:amethyst_shard")
	text, _ = call(t, s, "vanilla", map[string]any{"file": "tags/block/logs"})
	wantAll(t, text, `"#minecraft:logs_that_burn"`)
	text, _ = call(t, s, "vanilla", map[string]any{"file": "data/minecraft/loot_table/chests/simple_dungeon.json"})
	wantAll(t, text, `"pools"`)
	text, failed := call(t, s, "vanilla", map[string]any{"file": "loot_table/chests/simple"})
	if !failed {
		t.Error("a missing file should fail")
	}
	wantAll(t, text, "loot_table/chests/simple_dungeon")
	text, failed = call(t, s, "vanilla", map[string]any{"registry": "biome"})
	if !failed {
		t.Error("a missing registry should fail")
	}
	wantAll(t, text, "worldgen/biome")
}

func TestDocs(t *testing.T) {
	s := server(t)
	text, _ := call(t, s, "docs", nil)
	wantAll(t, text, "- guides/functions: Functions.", "- start/getting-started: Getting started.")
	if strings.Index(text, "- start/getting-started") > strings.Index(text, "- guides/functions") {
		t.Errorf("pages are not in reading order:\n%s", text)
	}
	text, _ = call(t, s, "docs", map[string]any{"page": "guides/functions"})
	wantAll(t, text, "# Functions")
	text, _ = call(t, s, "docs", map[string]any{"search": "storage records"})
	wantAll(t, text, "guides/functions", "§ Storage records")
	if _, failed := call(t, s, "docs", map[string]any{"page": "nope"}); !failed {
		t.Error("a missing page should fail")
	}
	found := map[string]bool{}
	for _, r := range resourceList() {
		found[r["uri"].(string)] = true
	}
	for _, uri := range []string{"mms://docs/guides/resources", "mms://prelude/items.mms"} {
		if !found[uri] {
			t.Errorf("resources lack %s", uri)
		}
		if _, ok := readResource(uri); !ok {
			t.Errorf("cannot read %s", uri)
		}
	}
}

// A long answer is cut at a rune, with a note.
func TestClip(t *testing.T) {
	long := strings.Repeat("é", maxResult)
	got := clip(long)
	if len(got) > maxResult+200 || !strings.Contains(got, "more bytes left out") {
		t.Errorf("clip left %d bytes", len(got))
	}
	if !json.Valid([]byte(`"` + strings.SplitN(got, "\n", 2)[0] + `"`)) {
		t.Error("clip cut a rune in half")
	}
}

// mcp/docs is a copy of site/, which go:embed cannot reach. Run
// `go test ./mcp -run TestDocsInSync -update` after changing a page.
func TestDocsInSync(t *testing.T) {
	const site, copyDir = "../../site", "docs"
	want := map[string]string{}
	err := filepath.WalkDir(site, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".md") {
			return err
		}
		b, err := os.ReadFile(p)
		rel, _ := filepath.Rel(site, p)
		want[filepath.ToSlash(rel)] = string(b)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	if *update {
		if err := os.RemoveAll(copyDir); err != nil {
			t.Fatal(err)
		}
		for rel, text := range want {
			p := filepath.Join(copyDir, filepath.FromSlash(rel))
			if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(p, []byte(text), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		return
	}
	const hint = "; run `go test ./mcp -run TestDocsInSync -update` in engine/"
	got := map[string]bool{}
	_ = filepath.WalkDir(copyDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(copyDir, p)
		rel = filepath.ToSlash(rel)
		got[rel] = true
		b, _ := os.ReadFile(p)
		if w, ok := want[rel]; !ok {
			t.Errorf("mcp/docs/%s is not in site/%s", rel, hint)
		} else if w != string(b) {
			t.Errorf("mcp/docs/%s differs from site/%s%s", rel, rel, hint)
		}
		return nil
	})
	for rel := range want {
		if !got[rel] {
			t.Errorf("site/%s is not in mcp/docs%s", rel, hint)
		}
	}
}
