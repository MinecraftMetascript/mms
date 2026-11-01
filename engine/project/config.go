// Package project ties files, configuration and the checker together: it
// analyses a set of .mms files, caches results per file, and emits the
// data pack.
package project

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ConfigFile is the project file name.
const ConfigFile = "mms.json"

// Config is the contents of mms.json.
type Config struct {
	// Namespace is the default namespace for files without a
	// `Namespace` declaration.
	Namespace   string `json:"namespace"`
	Description string `json:"description,omitempty"`
	// Output is the folder `mms build` writes, relative to the project.
	Output string `json:"output,omitempty"`
	// Sources are glob patterns, relative to the project; ** matches any
	// number of folders. The default is every .mms file.
	Sources []string `json:"sources,omitempty"`
	// Lint sets rule severities by rule id: "error", "warning", "info",
	// "hint" or "off".
	Lint map[string]string `json:"lint,omitempty"`
	// Icon is a PNG in the project that becomes pack.png in both packs.
	Icon string `json:"icon,omitempty"`
	// ResourcePack configures the resource pack built when the project
	// uses Texture or Sound.
	ResourcePack ResourcePackConfig `json:"resourcePack,omitzero"`
}

// ResourcePackConfig is the "resourcePack" object in mms.json.
type ResourcePackConfig struct {
	// Output is the folder `mms build` writes the resource pack to,
	// relative to the project. The default is the data pack's output
	// with "-resources" appended.
	Output string `json:"output,omitempty"`
	// Description defaults to the data pack's description.
	Description string `json:"description,omitempty"`
}

// ResourceOutput is the resource pack's output folder.
func (c Config) ResourceOutput() string {
	if c.ResourcePack.Output != "" {
		return c.ResourcePack.Output
	}
	return c.Output + "-resources"
}

// DefaultConfig is used when there is no mms.json.
func DefaultConfig() Config {
	return Config{Namespace: "mms", Description: "Built with MMS", Output: "build"}
}

var nsPattern = regexp.MustCompile(`^[a-z0-9_.\-]+$`)

// Validate reports a problem with the config, or nil.
func (c Config) Validate() error {
	if !nsPattern.MatchString(c.Namespace) {
		return fmt.Errorf("mms.json: namespace %q must be lowercase letters, digits, '_', '-' or '.'", c.Namespace)
	}
	if c.Icon != "" && !strings.HasSuffix(strings.ToLower(c.Icon), ".png") {
		return fmt.Errorf("mms.json: icon %q must be a .png file", c.Icon)
	}
	for id, sev := range c.Lint {
		if RuleByID(id) == nil {
			return fmt.Errorf("mms.json: lint: unknown rule %q", id)
		}
		if _, ok := ParseSeverity(sev); !ok {
			return fmt.Errorf("mms.json: lint: %s: severity must be error, warning, info, hint or off, not %q", id, sev)
		}
	}
	return nil
}

// LoadConfig reads dir/mms.json, filling defaults for missing fields. A
// missing file yields the default config.
func LoadConfig(dir string) (Config, error) {
	cfg := DefaultConfig()
	b, err := os.ReadFile(filepath.Join(dir, ConfigFile))
	if os.IsNotExist(err) {
		return cfg, nil
	}
	if err != nil {
		return cfg, err
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		return cfg, fmt.Errorf("mms.json: %w", err)
	}
	if cfg.Output == "" {
		cfg.Output = "build"
	}
	return cfg, cfg.Validate()
}

// SourceFiles lists the project's .mms files, sorted, relative to dir.
// Folders holding their own mms.json are other projects and left out.
func (c Config) SourceFiles(dir string) ([]string, error) {
	patterns := c.Sources
	if len(patterns) == 0 {
		patterns = []string{"**/*.mms"}
	}
	var res []*regexp.Regexp
	for _, p := range patterns {
		res = append(res, globRegexp(p))
	}
	var out []string
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(dir, path)
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			if rel != "." && (rel == c.Output || rel == c.ResourceOutput() || strings.HasPrefix(d.Name(), ".") || d.Name() == "node_modules") {
				return filepath.SkipDir
			}
			if rel != "." {
				// A folder with its own mms.json is a separate project.
				if st, err := os.Stat(filepath.Join(path, ConfigFile)); err == nil && !st.IsDir() {
					return filepath.SkipDir
				}
			}
			return nil
		}
		for _, re := range res {
			if re.MatchString(rel) {
				out = append(out, rel)
				break
			}
		}
		return nil
	})
	return out, err
}

// globRegexp compiles a glob with *, ** and ? into an anchored regexp.
func globRegexp(glob string) *regexp.Regexp {
	re := "^"
	for i := 0; i < len(glob); i++ {
		switch ch := glob[i]; ch {
		case '*':
			if i+1 < len(glob) && glob[i+1] == '*' {
				i++
				if i+1 < len(glob) && glob[i+1] == '/' {
					i++
					re += "(?:.*/)?"
				} else {
					re += ".*"
				}
			} else {
				re += "[^/]*"
			}
		case '?':
			re += "[^/]"
		default:
			re += regexp.QuoteMeta(string(ch))
		}
	}
	return regexp.MustCompile(re + "$")
}
