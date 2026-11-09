package spec

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
)

// Pack is everything MMS knows about one Minecraft version.
type Pack struct {
	Version    string `json:"version"`
	PackFormat int    `json:"packFormat"`
	// ResourcePackFormat is the resource pack format of the same version.
	ResourcePackFormat int `json:"resourcePackFormat"`
	// DataVersion is the version's world data version, which structure
	// files record.
	DataVersion int `json:"dataVersion"`
	// Sources records the upstream commits the pack was generated from.
	Sources map[string]string `json:"sources"`

	// Types maps mcdoc paths such as
	// "::java::data::worldgen::density_function::Clamp" to their type.
	Types map[string]*Type `json:"types"`
	// Dispatchers maps a dispatcher id such as "minecraft:density_function"
	// to its cases. Case keys are ids without namespace, or the special keys
	// "%none" and "%unknown".
	Dispatchers map[string]map[string]*Type `json:"dispatchers"`
	// Resources maps each data-pack resource kind (the folder under
	// data/<ns>/, e.g. "worldgen/noise") to its root type.
	Resources map[string]*Type `json:"resources"`
	// Assets maps the resource-pack kinds among Resources ("model",
	// "block_definition") to their folder under assets/<ns>/ ("models",
	// "blockstates").
	Assets map[string]string `json:"assets,omitempty"`

	// Registries maps a registry name ("block", "worldgen/biome") to its ids,
	// without the "minecraft:" prefix.
	Registries map[string][]string `json:"registries"`
	// Blocks maps a block id to its state properties and defaults.
	Blocks map[string]*BlockStates `json:"blocks"`
	// Commands is the root of the Brigadier command tree.
	Commands *CommandNode `json:"commands"`
	// Atlases maps each vanilla texture atlas ("blocks", "items") to the
	// sources its assets/minecraft/atlases/<name>.json lists.
	Atlases map[string][]AtlasSource `json:"atlases,omitempty"`

	registrySets map[string]map[string]struct{}
	// vanilla is set on the embedded default pack, whose version has
	// vanilla files (see Vanilla).
	vanilla bool
}

// AtlasSource is one entry of an atlas's sources list. Only the fields
// that say which texture files it stitches are kept: a "directory"
// source's Source folder and Prefix, and a "single" source's Resource
// and Sprite.
type AtlasSource struct {
	Type     string `json:"type"`
	Source   string `json:"source,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
	Resource string `json:"resource,omitempty"`
	Sprite   string `json:"sprite,omitempty"`
}

type BlockStates struct {
	Properties map[string][]string `json:"properties"`
	Defaults   map[string]string   `json:"defaults"`
}

type CommandNode struct {
	Type       string                  `json:"type"` // root, literal, argument
	Children   map[string]*CommandNode `json:"children,omitempty"`
	Executable bool                    `json:"executable,omitempty"`
	Redirect   []string                `json:"redirect,omitempty"`
	Parser     string                  `json:"parser,omitempty"`
	Properties map[string]any          `json:"properties,omitempty"`
}

//go:embed data/26.3.json.gz
var embedded263 []byte

var (
	defaultOnce sync.Once
	defaultPack *Pack
	defaultErr  error
)

// Default returns the embedded pack for the default target version. It is
// decoded once and shared; callers must treat it as read-only.
func Default() (*Pack, error) {
	defaultOnce.Do(func() {
		defaultPack, defaultErr = Decode(bytes.NewReader(embedded263))
		if defaultPack != nil {
			defaultPack.vanilla = true
		}
	})
	return defaultPack, defaultErr
}

// MustDefault is Default for callers that cannot continue without a pack.
func MustDefault() *Pack {
	p, err := Default()
	if err != nil {
		panic(fmt.Sprintf("spec: embedded pack is invalid: %v", err))
	}
	return p
}

// Decode reads a gzipped pack.
func Decode(r io.Reader) (*Pack, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var p Pack
	if err := json.NewDecoder(zr).Decode(&p); err != nil {
		return nil, err
	}
	p.index()
	return &p, nil
}

// Encode writes the pack gzipped with deterministic JSON.
func (p *Pack) Encode(w io.Writer) error {
	zw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(zw)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(p); err != nil {
		return err
	}
	return zw.Close()
}

func (p *Pack) index() {
	p.registrySets = make(map[string]map[string]struct{}, len(p.Registries))
	for name, ids := range p.Registries {
		set := make(map[string]struct{}, len(ids))
		for _, id := range ids {
			set[id] = struct{}{}
		}
		p.registrySets[name] = set
	}
}

// Resolve follows references until it reaches a non-reference type. It
// returns nil for unknown paths.
func (p *Pack) Resolve(t *Type) *Type {
	for i := 0; t != nil && t.Kind == "reference" && i < 64; i++ {
		t = p.Types[t.Path]
	}
	return t
}

// Dispatch returns the case of a dispatcher for a key such as "clamp" or
// "minecraft:clamp".
func (p *Pack) Dispatch(dispatcher, key string) (*Type, bool) {
	cases, ok := p.Dispatchers[dispatcher]
	if !ok {
		return nil, false
	}
	t, ok := cases[strings.TrimPrefix(key, "minecraft:")]
	return t, ok
}

// DispatchKeys returns the sorted case keys of a dispatcher, without the
// special "%" keys.
func (p *Pack) DispatchKeys(dispatcher string) []string {
	var keys []string
	for k := range p.Dispatchers[dispatcher] {
		if !strings.HasPrefix(k, "%") {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasID reports whether a registry contains an id. The id may carry the
// "minecraft:" prefix.
func (p *Pack) HasID(registry, id string) bool {
	set, ok := p.registrySets[strings.TrimPrefix(registry, "minecraft:")]
	if !ok {
		return false
	}
	_, ok = set[strings.TrimPrefix(id, "minecraft:")]
	return ok
}

// HasRegistry reports whether the pack lists a registry's ids.
func (p *Pack) HasRegistry(registry string) bool {
	_, ok := p.registrySets[strings.TrimPrefix(registry, "minecraft:")]
	return ok
}

// RegistryIDs returns a registry's ids, without the "minecraft:" prefix.
func (p *Pack) RegistryIDs(registry string) []string {
	return p.Registries[strings.TrimPrefix(registry, "minecraft:")]
}

// ResourceKinds returns the sorted data-pack resource kinds.
func (p *Pack) ResourceKinds() []string {
	kinds := make([]string, 0, len(p.Resources))
	for k := range p.Resources {
		if _, asset := p.Assets[k]; !asset {
			kinds = append(kinds, k)
		}
	}
	sort.Strings(kinds)
	return kinds
}

// AssetKinds returns the sorted resource-pack kinds.
func (p *Pack) AssetKinds() []string {
	kinds := make([]string, 0, len(p.Assets))
	for k := range p.Assets {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)
	return kinds
}

// SpriteOf returns the sprite the source stitches from the texture file
// assets/<ns>/textures/<path>.png, named "<ns>:<path>"; ok is false when
// the source does not take that file.
func (s AtlasSource) SpriteOf(file string) (sprite string, ok bool) {
	ns, p := splitNS(file)
	switch strings.TrimPrefix(s.Type, "minecraft:") {
	case "directory":
		if rest, ok := strings.CutPrefix(p, strings.TrimSuffix(s.Source, "/")+"/"); ok {
			return ns + ":" + s.Prefix + rest, true
		}
	case "single":
		if qualify(s.Resource) == ns+":"+p {
			if s.Sprite != "" {
				return qualify(s.Sprite), true
			}
			return ns + ":" + p, true
		}
	}
	return "", false
}

// Makes reports whether the source can stitch a sprite named id, from
// some file: a directory source names each of its files with its prefix.
func (s AtlasSource) Makes(id string) bool {
	switch strings.TrimPrefix(s.Type, "minecraft:") {
	case "directory":
		_, p := splitNS(id)
		return strings.HasPrefix(p, s.Prefix)
	case "single":
		name := s.Sprite
		if name == "" {
			name = s.Resource
		}
		return qualify(name) == qualify(id)
	}
	return false
}

func splitNS(id string) (ns, path string) {
	if ns, path, ok := strings.Cut(id, ":"); ok {
		return ns, path
	}
	return "minecraft", id
}

func qualify(id string) string {
	if strings.Contains(id, ":") {
		return id
	}
	return "minecraft:" + id
}
