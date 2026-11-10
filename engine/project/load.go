package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Load reads the project in the folder dir from disk: mms.json, when
// there is one, and every source file. The command line and the MCP
// server both start here.
func Load(dir string) (*Project, error) {
	cfg, err := LoadConfig(dir)
	if err != nil {
		return nil, err
	}
	pack, err := spec.Default()
	if err != nil {
		return nil, err
	}
	p := New(pack, cfg)
	p.Files = os.DirFS(dir)
	if cfg.Icon != "" {
		if st, err := os.Stat(filepath.Join(dir, filepath.FromSlash(cfg.Icon))); err != nil || st.IsDir() {
			return nil, fmt.Errorf("mms.json: icon: no file %q in the project", cfg.Icon)
		}
	}
	paths, err := cfg.SourceFiles(dir)
	if err != nil {
		return nil, err
	}
	for _, rel := range paths {
		b, err := os.ReadFile(filepath.Join(dir, rel))
		if err != nil {
			return nil, err
		}
		p.SetFile(rel, string(b))
	}
	return p, nil
}
