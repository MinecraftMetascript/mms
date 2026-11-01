package project

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// schemaFiles are the committed copies of ConfigSchema: the VS Code
// extension bundles one, and the docs site serves the other at SchemaURL.
var schemaFiles = []string{
	"../../editors/vscode/schemas/mms.schema.json",
	"../../web/static/mms.schema.json",
}

func TestConfigSchemaFiles(t *testing.T) {
	want := ConfigSchema()
	for _, f := range schemaFiles {
		if update {
			if err := os.MkdirAll(filepath.Dir(f), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(f, want, 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		got, err := os.ReadFile(f)
		if err != nil || !bytes.Equal(got, want) {
			t.Errorf("%s is out of date; run `MMS_UPDATE=1 go test ./project -run TestConfigSchemaFiles`", f)
		}
	}
}

func TestConfigSchemaCoversConfig(t *testing.T) {
	var s struct {
		Properties map[string]struct {
			Properties map[string]any `json:"properties"`
		} `json:"properties"`
	}
	if err := json.Unmarshal(ConfigSchema(), &s); err != nil {
		t.Fatal(err)
	}
	// Every field the loader reads must be in the schema, or editors
	// would flag it as unknown.
	full := Config{Namespace: "x", Description: "x", Output: "x", Sources: []string{"x"},
		Lint: map[string]string{"x": "x"}, Icon: "x.png",
		ResourcePack: ResourcePackConfig{Output: "x", Description: "x"}}
	b, _ := json.Marshal(full)
	var fields map[string]json.RawMessage
	json.Unmarshal(b, &fields)
	for k := range fields {
		if _, ok := s.Properties[k]; !ok {
			t.Errorf("schema lacks mms.json field %q", k)
		}
	}
	b, _ = json.Marshal(full.ResourcePack)
	fields = nil
	json.Unmarshal(b, &fields)
	for k := range fields {
		if _, ok := s.Properties["resourcePack"].Properties[k]; !ok {
			t.Errorf("schema lacks resourcePack field %q", k)
		}
	}
	for _, r := range Rules {
		if _, ok := s.Properties["lint"].Properties[r.ID]; !ok {
			t.Errorf("schema lacks lint rule %q", r.ID)
		}
	}
}

// A config with "$schema" still loads: the loader ignores the key.
func TestLoadConfigWithSchemaKey(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, ConfigFile), []byte(`{"$schema": "`+SchemaURL+`", "namespace": "demo"}`), 0o644)
	cfg, err := LoadConfig(dir)
	if err != nil || cfg.Namespace != "demo" {
		t.Fatalf("got %+v, %v", cfg, err)
	}
}
