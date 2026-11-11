package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestInitWritesMCPJSON checks that init registers the mms and nbt-mcp
// servers in a new .mcp.json.
func TestInitWritesMCPJSON(t *testing.T) {
	dir := t.TempDir()
	var out, errOut bytes.Buffer
	if code := runInit([]string{dir}, &out, &errOut); code != 0 {
		t.Fatalf("init exit %d: %s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "created mms.json, main.mms and .mcp.json;") {
		t.Errorf("init output %q", out.String())
	}
	got, _ := os.ReadFile(filepath.Join(dir, ".mcp.json"))
	want := `{
  "mcpServers": {
    "mms": {
      "command": "mms",
      "args": [
        "mcp"
      ]
    },
    "nbt": {
      "command": "nbt-mcp",
      "args": [
        "-dir",
        "structures"
      ]
    }
  }
}
`
	if string(got) != want {
		t.Errorf(".mcp.json =\n%s\nwant\n%s", got, want)
	}
}

// TestInitMergesMCPJSON checks that init adds only the servers an
// existing .mcp.json lacks, keeps the rest as written, and leaves a file
// it cannot read alone.
func TestInitMergesMCPJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".mcp.json")
	write(t, path, `{"mcpServers": {"mms": {"command": "/opt/mms", "args": ["mcp"]}, "other": {"command": "x"}}, "extra": 1}`)
	var out, errOut bytes.Buffer
	if code := runInit([]string{dir}, &out, &errOut); code != 0 {
		t.Fatalf("init exit %d: %s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "added nbt to .mcp.json") || strings.Contains(out.String(), "and .mcp.json") {
		t.Errorf("init output %q", out.String())
	}
	var doc struct {
		Servers map[string]struct {
			Command string
		} `json:"mcpServers"`
		Extra int `json:"extra"`
	}
	raw, _ := os.ReadFile(path)
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}
	if doc.Servers["mms"].Command != "/opt/mms" || doc.Servers["other"].Command != "x" || doc.Servers["nbt"].Command != "nbt-mcp" || doc.Extra != 1 {
		t.Errorf(".mcp.json = %s", raw)
	}

	// Every server already there: nothing to say, file unchanged.
	if _, note, err := writeMCPJSON(dir); err != nil || note != "" {
		t.Errorf("second write: %q %v", note, err)
	}
	if again, _ := os.ReadFile(path); !bytes.Equal(again, raw) {
		t.Errorf("rewrote .mcp.json:\n%s", again)
	}

	bad := t.TempDir()
	write(t, filepath.Join(bad, ".mcp.json"), "{ not json")
	out.Reset()
	if code := runInit([]string{bad}, &out, &errOut); code != 0 {
		t.Fatalf("init exit %d: %s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "left .mcp.json alone") {
		t.Errorf("init output %q", out.String())
	}
	if got, _ := os.ReadFile(filepath.Join(bad, ".mcp.json")); string(got) != "{ not json" {
		t.Errorf("bad .mcp.json was changed: %s", got)
	}
}
