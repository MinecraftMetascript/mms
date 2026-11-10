package main

import (
	"bytes"
	"strings"
	"testing"
)

// `mms mcp -check` tests the server and prints how to connect a client;
// `mms mcp` answers a client on stdout and keeps its own words on stderr.
func TestMCP(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := runMCP([]string{"-check"}, strings.NewReader(""), &out, &errOut); code != 0 {
		t.Fatalf("-check exit %d: %s", code, errOut.String())
	}
	for _, want := range []string{"ok: mms ", "9 tools answer", "claude mcp add mms -- ", `{"mcpServers":{"mms":{"command":"mms","args":["mcp"]}}}`, "Pending approval"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("-check output lacks %q:\n%s", want, out.String())
		}
	}

	out.Reset()
	in := `{"jsonrpc":"2.0","id":1,"method":"ping"}` + "\n"
	if code := runMCP(nil, strings.NewReader(in), &out, &errOut); code != 0 {
		t.Fatalf("exit %d: %s", code, errOut.String())
	}
	if got := strings.TrimSpace(out.String()); got != `{"jsonrpc":"2.0","id":1,"result":{}}` {
		t.Errorf("stdout = %q; only protocol messages may go there", got)
	}
	if !strings.Contains(errOut.String(), "waiting for an MCP client") {
		t.Errorf("stderr = %q", errOut.String())
	}

	if code := runMCP([]string{"extra"}, strings.NewReader(""), &out, &errOut); code != 2 {
		t.Errorf("an unknown argument: exit %d", code)
	}
}
