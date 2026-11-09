package spec

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"io"
	"strings"
	"sync"
)

// Vanilla holds the JSON files of the game's own data pack, as mcmeta's
// summary lists them: Vanilla[folder][path] is the file
// data/minecraft/<folder>/<path>.json. Folders are data pack folders such
// as "worldgen/noise" and "tags/block". Editors show these files when go to
// definition lands on a vanilla id.
type Vanilla map[string]map[string]json.RawMessage

// VanillaPrefix starts the path of a vanilla file, as go to definition
// reports it: @vanilla/data/minecraft/worldgen/noise/ridge.json.
const VanillaPrefix = "@vanilla/"

//go:embed data/26.3.vanilla.json.gz
var embeddedVanilla263 []byte

var (
	vanillaOnce sync.Once
	vanillaData Vanilla
	vanillaErr  error
)

// DefaultVanilla returns the vanilla files of the default target version.
// They are decoded on first use, not with the pack, since only go to
// definition needs them.
func DefaultVanilla() (Vanilla, error) {
	vanillaOnce.Do(func() {
		vanillaData, vanillaErr = DecodeVanilla(bytes.NewReader(embeddedVanilla263))
	})
	return vanillaData, vanillaErr
}

// DecodeVanilla reads gzipped vanilla files.
func DecodeVanilla(r io.Reader) (Vanilla, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var v Vanilla
	if err := json.NewDecoder(zr).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// Encode writes the files gzipped with deterministic JSON.
func (v Vanilla) Encode(w io.Writer) error {
	zw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(zw)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return err
	}
	return zw.Close()
}

// Vanilla returns the vanilla files of the pack's version, or nil when
// there are none (a pack that is not the embedded default).
func (p *Pack) Vanilla() Vanilla {
	if p == nil || !p.vanilla {
		return nil
	}
	v, _ := DefaultVanilla()
	return v
}

// VanillaPath returns the path go to definition reports for a vanilla
// file, or "" when vanilla has no such file. folder is a data pack folder
// ("worldgen/noise", "tags/block") and id a minecraft id, with or without
// its namespace and a tag's '#'.
func (v Vanilla) VanillaPath(folder, id string) string {
	id = strings.TrimPrefix(id, "#")
	if ns, path, ok := strings.Cut(id, ":"); ok {
		if ns != "minecraft" {
			return ""
		}
		id = path
	}
	if _, ok := v[folder][id]; !ok {
		return ""
	}
	return VanillaPrefix + "data/minecraft/" + folder + "/" + id + ".json"
}

// File returns the text of the vanilla file at a path VanillaPath gave,
// indented for reading.
func (v Vanilla) File(path string) (string, bool) {
	rest, ok := strings.CutPrefix(path, VanillaPrefix+"data/minecraft/")
	if !ok {
		return "", false
	}
	rest, ok = strings.CutSuffix(rest, ".json")
	if !ok {
		return "", false
	}
	// Folders and ids both hold slashes; try every split.
	for i := strings.Index(rest, "/"); i >= 0; {
		if raw, ok := v[rest[:i]][rest[i+1:]]; ok {
			var b bytes.Buffer
			if err := json.Indent(&b, raw, "", "  "); err != nil {
				return string(raw), true
			}
			b.WriteByte('\n')
			return b.String(), true
		}
		j := strings.Index(rest[i+1:], "/")
		if j < 0 {
			break
		}
		i += j + 1
	}
	return "", false
}
