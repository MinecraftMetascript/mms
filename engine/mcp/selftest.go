package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// SelfTest talks to the server the way a client does, through Run: it
// starts a session, lists the tools and checks a small file. It returns
// a line saying what answered, or what went wrong. `mms mcp -check`
// prints it, to tell a broken install from a client that is not set up.
func (s *Server) SelfTest(ctx context.Context) (string, error) {
	in := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"` + protocolVersions[0] + `","capabilities":{},"clientInfo":{"name":"mms mcp -check","version":"0"}}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized"}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list"}`,
		`{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"check","arguments":{"source":"Function greet {\n  /say hi\n}\n"}}}`,
	}, "\n") + "\n"
	var out bytes.Buffer
	if err := s.Run(ctx, strings.NewReader(in), &out); err != nil {
		return "", err
	}
	type answer struct {
		Error  *rpcError `json:"error"`
		Result struct {
			ProtocolVersion string            `json:"protocolVersion"`
			Tools           []json.RawMessage `json:"tools"`
			IsError         bool              `json:"isError"`
			Content         []struct {
				Text string `json:"text"`
			} `json:"content"`
		} `json:"result"`
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 3 {
		return "", fmt.Errorf("the server gave %d answers to 3 requests", len(lines))
	}
	var got [3]answer
	for i, line := range lines {
		if err := json.Unmarshal([]byte(line), &got[i]); err != nil {
			return "", fmt.Errorf("answer %d is not JSON: %w", i+1, err)
		}
		if got[i].Error != nil {
			return "", fmt.Errorf("answer %d is an error: %s", i+1, got[i].Error.Message)
		}
	}
	if len(got[1].Result.Tools) == 0 {
		return "", fmt.Errorf("the server lists no tools")
	}
	if c := got[2].Result; c.IsError || len(c.Content) == 0 || !strings.HasPrefix(c.Content[0].Text, "ok:") {
		return "", fmt.Errorf("the check tool did not pass a correct file: %+v", c.Content)
	}
	return fmt.Sprintf("%d tools answer, protocol %s", len(got[1].Result.Tools), got[0].Result.ProtocolVersion), nil
}

// ConfigJSON is the entry an MCP client's JSON file of servers needs to
// start the server with command, on one line.
func ConfigJSON(command string) string {
	type entry struct {
		Command string   `json:"command"`
		Args    []string `json:"args"`
	}
	b, _ := json.Marshal(map[string]any{"mcpServers": map[string]entry{"mms": {command, []string{"mcp"}}}})
	return string(b)
}
