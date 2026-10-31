package lsp

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
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
			diags = nil
			for _, d := range n["params"].(map[string]any)["diagnostics"].([]any) {
				if d.(map[string]any)["severity"].(float64) != 4 {
					diags = append(diags, d)
				}
			}
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
	if text := files[0].(map[string]any)["text"]; text != "ItemTag shiny = [diamond]\n" {
		t.Errorf("text = %q", text)
	}
	if sk := res["skipped"].([]any); len(sk) != 1 {
		t.Errorf("skipped = %v", sk)
	}
	empty := c.result("mms/importPack", map[string]any{"files": []any{}}).(map[string]any)
	if len(empty["files"].([]any)) != 0 {
		t.Errorf("empty import = %v", empty)
	}
	c.notify("textDocument/didOpen", map[string]any{"textDocument": map[string]any{"uri": "file:///a.mms", "languageId": "mms", "version": 1,
		"text": "Namespace gems\nNoise ridge = { base_octave = -1, amplitude_modifiers = [1] }\n"}})
	conv := c.result("mms/convertJSON", map[string]any{"uri": "file:///a.mms", "kind": "DensityFunction", "name": "hill",
		"text": `{"type":"minecraft:noise","noise":"minecraft:ridge","xz_scale":1,"y_scale":1}`}).(map[string]any)
	// minecraft:ridge keeps its prefix: the file declares a ridge noise.
	if want := "DensityFunction hill = Noise(noise = minecraft:ridge, xz_scale = 1, y_scale = 1)\n"; conv["text"] != want {
		t.Errorf("convertJSON = %q\nwant %q", conv["text"], want)
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

func TestCodeActions(t *testing.T) {
	c := start(t)
	c.result("initialize", map[string]any{})
	c.notify("initialized", map[string]any{})
	uri := "file:///a.mms"
	src := "LootTable t = { pools = [{ rolls = 1, entries = [{ type = \"minecraft:item\", name = minecraft:diamond }] }] }\n"
	c.notify("textDocument/didOpen", map[string]any{"textDocument": map[string]any{"uri": uri, "languageId": "mms", "version": 1, "text": src}})
	rng := func(a, b int) map[string]any {
		return map[string]any{"start": map[string]any{"line": 0, "character": a}, "end": map[string]any{"line": 0, "character": b}}
	}
	acts := c.result("textDocument/codeAction", map[string]any{"textDocument": map[string]any{"uri": uri}, "range": rng(60, 60), "context": map[string]any{"diagnostics": []any{}}}).([]any)
	var titles []string
	for _, a := range acts {
		titles = append(titles, a.(map[string]any)["title"].(string))
	}
	if got := strings.Join(titles, "|"); got != "Write minecraft:diamond|Ignore use-sugar on this line|Write item|Ignore quoted-name on this line|Fix all MMS problems" {
		t.Errorf("actions = %s", got)
	}
	all := c.result("textDocument/codeAction", map[string]any{"textDocument": map[string]any{"uri": uri}, "range": rng(0, 0),
		"context": map[string]any{"diagnostics": []any{}, "only": []string{"source.fixAll"}}}).([]any)
	if len(all) != 1 {
		t.Fatalf("fix all = %v", all)
	}
	edit := all[0].(map[string]any)["edit"].(map[string]any)["changes"].(map[string]any)[uri].([]any)[0].(map[string]any)
	if want := "LootTable t = { pools = [{ rolls = 1, entries = [diamond] }] }\n"; edit["newText"] != want {
		t.Errorf("fix all text = %q", edit["newText"])
	}
}

// TestWorkspaceFiles checks that the server loads every project file from
// the workspace, not just open ones, and resolves names across them.
func TestWorkspaceFiles(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "terrain"), 0o755)
	os.WriteFile(filepath.Join(dir, "mms.json"), []byte(`{"namespace": "demo"}`), 0o644)
	os.WriteFile(filepath.Join(dir, "terrain", "base.mms"), []byte("DensityFunction base = 1\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "main.mms"), []byte("DensityFunction top = base * 2\n"), 0o644)
	a, b := pathURI(filepath.Join(dir, "terrain", "base.mms")), pathURI(filepath.Join(dir, "main.mms"))

	c := start(t)
	c.result("initialize", map[string]any{"rootUri": pathURI(dir)})
	c.notify("textDocument/didOpen", map[string]any{"textDocument": map[string]any{"uri": b, "languageId": "mms", "version": 1, "text": "DensityFunction top = base * 2\n"}})

	// base is declared in a file that was never opened.
	def := c.result("textDocument/definition", pos(b, 0, 23)).([]any)
	if len(def) != 1 || def[0].(map[string]any)["uri"] != a {
		t.Fatalf("definition = %v, want %s", def, a)
	}
	ren := c.result("textDocument/rename", map[string]any{"textDocument": map[string]any{"uri": b}, "position": map[string]any{"line": 0, "character": 23}, "newName": "ground"}).(map[string]any)
	changes := ren["changes"].(map[string]any)
	if len(changes) != 2 || changes[a] == nil || changes[b] == nil {
		t.Errorf("rename changes = %v, want edits in both files", changes)
	}
}

// TestURIs checks file URI handling, including Windows drives, which the
// workspace scan and the client must spell the same way.
func TestURIs(t *testing.T) {
	for _, c := range []struct{ uri, path string }{
		{"file:///home/me/a.mms", "/home/me/a.mms"},
		{"file:///c%3A/Users/me/a.mms", filepath.FromSlash("c:/Users/me/a.mms")},
		{"file:///C:/Users/me/a.mms", filepath.FromSlash("C:/Users/me/a.mms")},
		{"mem:///a.mms", ""},
	} {
		if got := uriPath(c.uri); got != c.path {
			t.Errorf("uriPath(%q) = %q, want %q", c.uri, got, c.path)
		}
	}
	if got := pathURI("C:/Users/me/a.mms"); got != "file:///C:/Users/me/a.mms" {
		t.Errorf("pathURI kept the drive as a host: %q", got)
	}
	// What VS Code sends and what the workspace scan builds become one key.
	if a, b := canonURI("file:///c%3A/Users/me/a.mms"), pathURI(uriPath("file:///c%3A/Users/me/a.mms")); a != b {
		t.Errorf("canonURI = %q, scan = %q", a, b)
	}
	if got := canonURI("mem:///a.mms"); got != "mem:///a.mms" {
		t.Errorf("canonURI changed a mem URI: %q", got)
	}
}

func TestRegistersWatchers(t *testing.T) {
	c := start(t)
	caps := map[string]any{"workspace": map[string]any{"didChangeWatchedFiles": map[string]any{"dynamicRegistration": true}}}
	c.result("initialize", map[string]any{"rootUri": nil, "capabilities": caps})
	c.notify("initialized", map[string]any{})
	c.result("mms/ping", nil)
	var globs []string
	for _, n := range c.notes {
		if n["method"] != "client/registerCapability" {
			continue
		}
		for _, r := range n["params"].(map[string]any)["registrations"].([]any) {
			opts := r.(map[string]any)["registerOptions"].(map[string]any)
			for _, w := range opts["watchers"].([]any) {
				globs = append(globs, w.(map[string]any)["globPattern"].(string))
			}
		}
	}
	if strings.Join(globs, " ") != "**/*.mms **/mms.json **/*.{png,ogg,nbt}" {
		t.Fatalf("watchers = %q", globs)
	}

	// Without dynamic registration the server must not send the request.
	c2 := start(t)
	c2.result("initialize", map[string]any{"rootUri": nil})
	c2.notify("initialized", map[string]any{})
	c2.result("mms/ping", nil)
	for _, n := range c2.notes {
		if n["method"] == "client/registerCapability" {
			t.Fatal("registered watchers for a client without dynamic registration")
		}
	}
}
