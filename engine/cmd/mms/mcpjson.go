package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// mcpServers are the MCP servers `mms init` registers in .mcp.json: the
// engine's own server, and nbt-mcp from nbt-go for building the
// structure files (structures/<path>.nbt) that pools load.
var mcpServers = []struct {
	name  string
	entry string
}{
	{"mms", `{"command": "mms", "args": ["mcp"]}`},
	{"nbt", `{"command": "nbt-mcp", "args": ["-dir", "structures"]}`},
}

// writeMCPJSON writes dir/.mcp.json with mcpServers, or adds to the one
// already there the servers it does not name yet. It keeps every other
// entry and field as written, and leaves a file it cannot read alone. It
// reports whether it made a new file, and otherwise a line for the user
// ("" when every server was already there).
func writeMCPJSON(dir string) (created bool, note string, err error) {
	path := filepath.Join(dir, ".mcp.json")
	doc := map[string]json.RawMessage{}
	servers := map[string]json.RawMessage{}
	mode, old, rerr := readWithMode(path)
	switch err := rerr; {
	case err == nil:
		if err := json.Unmarshal(old, &doc); err != nil || doc == nil {
			return false, "left .mcp.json alone: it is not a JSON object", nil
		}
		if raw, ok := doc["mcpServers"]; ok {
			if err := json.Unmarshal(raw, &servers); err != nil || servers == nil {
				return false, "left .mcp.json alone: its mcpServers is not an object", nil
			}
		}
	case !os.IsNotExist(err):
		return false, "", err
	}
	var added []string
	for _, s := range mcpServers {
		if _, ok := servers[s.name]; !ok {
			servers[s.name] = json.RawMessage(s.entry)
			added = append(added, s.name)
		}
	}
	if len(added) == 0 {
		return false, "", nil
	}
	raw, err := json.Marshal(servers)
	if err != nil {
		return false, "", err
	}
	doc["mcpServers"] = raw
	flat, err := json.Marshal(doc)
	if err != nil {
		return false, "", err
	}
	var out bytes.Buffer
	if err := json.Indent(&out, flat, "", "  "); err != nil {
		return false, "", err
	}
	out.WriteByte('\n')
	// Write beside it and rename, so a failed write never truncates the
	// user's file.
	tmp, err := os.CreateTemp(dir, ".mcp.json.*")
	if err != nil {
		return false, "", err
	}
	_, werr := tmp.Write(out.Bytes())
	cerr := tmp.Close()
	if err := errors.Join(werr, cerr, os.Chmod(tmp.Name(), mode)); err != nil {
		os.Remove(tmp.Name())
		return false, "", err
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		os.Remove(tmp.Name())
		return false, "", err
	}
	if old != nil {
		return false, fmt.Sprintf("added %s to .mcp.json", joinAnd(added)), nil
	}
	return true, "", nil
}

func joinAnd(s []string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0]
	}
	out := s[0]
	for _, x := range s[1 : len(s)-1] {
		out += ", " + x
	}
	return out + " and " + s[len(s)-1]
}

// readWithMode reads path and its permission bits from one open file, so
// a merge keeps a private (0600) .mcp.json private: it can hold secrets in
// env. A missing file reads as 0644 and a not-exist error.
func readWithMode(path string) (os.FileMode, []byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0o644, nil, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return 0, nil, err
	}
	data, err := io.ReadAll(f)
	return fi.Mode().Perm(), data, err
}
