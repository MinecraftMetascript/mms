package lsp

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

type client struct {
	t    *testing.T
	pipe *Pipe
	id   int
	// notes collects notifications seen while waiting for responses.
	notes []map[string]any
	done  chan error
}

func start(t *testing.T) *client {
	t.Helper()
	p := NewPipe()
	s := NewServer(p, spec.MustDefault())
	c := &client{t: t, pipe: p, done: make(chan error, 1)}
	go func() { c.done <- s.Run(context.Background()) }()
	t.Cleanup(p.Close)
	return c
}

func (c *client) send(m map[string]any) {
	m["jsonrpc"] = "2.0"
	b, _ := json.Marshal(m)
	c.pipe.Send(b)
}

func (c *client) notify(method string, params any) {
	c.send(map[string]any{"method": method, "params": params})
}

func (c *client) call(method string, params any) map[string]any {
	c.t.Helper()
	c.id++
	c.send(map[string]any{"id": c.id, "method": method, "params": params})
	for {
		b, ok := c.pipe.Recv()
		if !ok {
			c.t.Fatal("connection closed")
		}
		var m map[string]any
		if err := json.Unmarshal(b, &m); err != nil {
			c.t.Fatal(err)
		}
		if id, ok := m["id"].(float64); ok && int(id) == c.id {
			return m
		}
		c.notes = append(c.notes, m)
	}
}

func (c *client) result(method string, params any) any {
	c.t.Helper()
	m := c.call(method, params)
	if e, ok := m["error"]; ok {
		c.t.Fatalf("%s: %v", method, e)
	}
	return m["result"]
}

func pos(uri string, line, char int) map[string]any {
	return map[string]any{"textDocument": map[string]any{"uri": uri}, "position": map[string]any{"line": line, "character": char}}
}

const uri = "file:///p/main.mms"

func TestSession(t *testing.T) {
	c := start(t)
	initRes := c.result("initialize", map[string]any{"rootUri": nil}).(map[string]any)
	if _, ok := initRes["capabilities"].(map[string]any)["semanticTokensProvider"]; !ok {
		t.Fatal("no semantic tokens capability")
	}
	c.notify("initialized", map[string]any{})
	src := "Noise ridges = { base_octave = -1, amplitude_modifiers = [1] }\nDensityFunction x = Clamp(Noise(ridges, 1, 1), 0, 1)\nDensityFunction y = Clmp(1, 0, 1)\n"
	c.notify("textDocument/didOpen", map[string]any{"textDocument": map[string]any{"uri": uri, "languageId": "mms", "version": 1, "text": src}})

	// Diagnostics arrive before the next response.
	c.result("textDocument/hover", pos(uri, 0, 0))
	var diags []any
	for _, n := range c.notes {
		if n["method"] == "textDocument/publishDiagnostics" {
			diags = n["params"].(map[string]any)["diagnostics"].([]any)
		}
	}
	if len(diags) != 1 || !strings.Contains(diags[0].(map[string]any)["message"].(string), "Clamp") {
		t.Fatalf("diagnostics = %v", diags)
	}

	comp := c.result("textDocument/completion", pos(uri, 2, 22)).(map[string]any)
	if len(comp["items"].([]any)) == 0 {
		t.Error("no completions")
	}
	hover := c.result("textDocument/hover", pos(uri, 1, 35))
	if hover == nil || !strings.Contains(hover.(map[string]any)["contents"].(map[string]any)["value"].(string), "ridges") {
		t.Errorf("hover = %v", hover)
	}
	def := c.result("textDocument/definition", pos(uri, 1, 35)).([]any)
	if len(def) != 1 {
		t.Errorf("definition = %v", def)
	}
	ren := c.result("textDocument/rename", map[string]any{"textDocument": map[string]any{"uri": uri}, "position": map[string]any{"line": 0, "character": 7}, "newName": "hills"}).(map[string]any)
	if edits := ren["changes"].(map[string]any)[uri].([]any); len(edits) != 2 {
		t.Errorf("rename edits = %v", edits)
	}
	toks := c.result("textDocument/semanticTokens/full", map[string]any{"textDocument": map[string]any{"uri": uri}}).(map[string]any)
	if n := len(toks["data"].([]any)); n == 0 || n%5 != 0 {
		t.Errorf("semantic tokens = %v", toks)
	}
	syms := c.result("textDocument/documentSymbol", map[string]any{"textDocument": map[string]any{"uri": uri}}).([]any)
	if len(syms) != 3 {
		t.Errorf("symbols = %v", syms)
	}

	// Fix the typo; the diagnostics clear and the pack has the function.
	fixed := strings.Replace(src, "Clmp", "Clamp", 1)
	c.notify("textDocument/didChange", map[string]any{"textDocument": map[string]any{"uri": uri, "version": 2}, "contentChanges": []any{map[string]any{"text": fixed}}})
	out := c.result("mms/output", map[string]any{}).(map[string]any)
	if out["hasErrors"] != false {
		t.Errorf("output has errors")
	}
	found := false
	for _, f := range out["files"].([]any) {
		if f.(map[string]any)["path"] == "data/mms/worldgen/density_function/x.json" {
			found = true
		}
	}
	if !found {
		t.Errorf("output files = %v", out["files"])
	}

	messy := "Noise   n={base_octave=-1,amplitude_modifiers=[1]}\n"
	c.notify("textDocument/didChange", map[string]any{"textDocument": map[string]any{"uri": uri, "version": 3}, "contentChanges": []any{map[string]any{"text": messy}}})
	fm := c.result("textDocument/formatting", map[string]any{"textDocument": map[string]any{"uri": uri}, "options": map[string]any{}}).([]any)
	if len(fm) != 1 {
		t.Fatalf("formatting = %v", fm)
	}

	c.result("shutdown", nil)
	c.notify("exit", nil)
	if err := <-c.done; err != nil {
		t.Fatal(err)
	}
}

func TestStream(t *testing.T) {
	var out bytes.Buffer
	in := strings.NewReader("Content-Length: 2\r\n\r\n{}Content-Length: 3\r\nContent-Type: x\r\n\r\n[1]")
	s := NewStream(in, &out)
	for _, want := range []string{"{}", "[1]"} {
		b, err := s.Read()
		if err != nil || string(b) != want {
			t.Fatalf("read %q, %v; want %q", b, err, want)
		}
	}
	if _, err := s.Read(); err == nil {
		t.Fatal("want EOF")
	}
	s.Write([]byte("{}"))
	if out.String() != "Content-Length: 2\r\n\r\n{}" {
		t.Fatalf("wrote %q", out.String())
	}
}

func TestPositions(t *testing.T) {
	l := newLines("a😀b\nxy")
	if p := l.pos(2); p != (Position{0, 3}) {
		t.Errorf("pos(2) = %v", p)
	}
	if o := l.offset(Position{0, 3}); o != 2 {
		t.Errorf("offset = %d", o)
	}
	if o := l.offset(Position{1, 9}); o != 6 {
		t.Errorf("clamped offset = %d", o)
	}
	if p := l.pos(5); p != (Position{1, 1}) {
		t.Errorf("pos(5) = %v", p)
	}
}

func TestManagedFiles(t *testing.T) {
	c := start(t)
	c.result("initialize", map[string]any{})
	a, b := "mem:///a.mms", "mem:///b.mms"
	c.notify("mms/setFiles", map[string]any{"files": []any{
		map[string]any{"uri": a, "text": "DensityFunction a = 1\n"},
		map[string]any{"uri": b, "text": "DensityFunction b = a + 1\n"},
	}})
	c.notify("textDocument/didOpen", map[string]any{"textDocument": map[string]any{"uri": a, "text": "DensityFunction a = 2\n"}})
	c.notify("textDocument/didClose", map[string]any{"textDocument": map[string]any{"uri": a}})
	out := c.result("mms/output", map[string]any{}).(map[string]any)
	texts := map[string]string{}
	for _, f := range out["files"].([]any) {
		m := f.(map[string]any)
		texts[m["path"].(string)] = m["text"].(string)
	}
	if !strings.Contains(texts["data/mms/worldgen/density_function/a.json"], "2") {
		t.Errorf("closing a managed file lost its edits: %v", texts)
	}
	c.notify("mms/setFiles", map[string]any{"remove": []any{b}})
	out = c.result("mms/output", map[string]any{}).(map[string]any)
	for _, f := range out["files"].([]any) {
		if strings.Contains(f.(map[string]any)["path"].(string), "/b.json") {
			t.Error("removed file still built")
		}
	}
	if c.result("mms/ping", nil) != "pong" {
		t.Error("ping")
	}
}

func TestImportPack(t *testing.T) {
	c := start(t)
	c.result("initialize", map[string]any{})
	c.notify("initialized", map[string]any{})
	res := c.result("mms/importPack", map[string]any{"files": []any{
		map[string]any{"path": "pack.mcmeta", "text": `{"pack":{"description":"Gems"}}`},
		map[string]any{"path": "data/gems/tags/item/shiny.json", "text": `{"values":["minecraft:diamond"]}`},
		map[string]any{"path": "data/gems/notes.txt", "text": "hi"},
	}}).(map[string]any)
	files := res["files"].([]any)
	if len(files) != 1 || res["namespace"] != "gems" || res["description"] != "Gems" {
		t.Fatalf("import = %v", res)
	}
	if text := files[0].(map[string]any)["text"]; text != "ItemTag shiny = [minecraft:diamond]\n" {
		t.Errorf("text = %q", text)
	}
	if sk := res["skipped"].([]any); len(sk) != 1 {
		t.Errorf("skipped = %v", sk)
	}
	empty := c.result("mms/importPack", map[string]any{"files": []any{}}).(map[string]any)
	if len(empty["files"].([]any)) != 0 {
		t.Errorf("empty import = %v", empty)
	}
	kinds := c.result("mms/resourceKinds", map[string]any{}).([]any)
	found := false
	for _, k := range kinds {
		if m := k.(map[string]any); m["kind"] == "LootTable" && m["folder"] == "loot_table" {
			found = true
		}
	}
	if !found {
		t.Errorf("LootTable missing from %d kinds", len(kinds))
	}
}
