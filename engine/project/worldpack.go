package project

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Minecraft applies one bundled resource pack per world, the file
// <world>/resourcepacks/resources.zip, so every project built into a world
// shares that zip. MergeWorldPack updates it in place for one project: it
// keeps the files other projects (or a hand-made pack) put there, replaces
// the files this project wrote last time, and merges the JSON files that
// several packs commonly add to (lang, sounds.json, atlases, fonts).
//
// Ownership is recorded in WorldPackManifest at the root of the zip, which
// Minecraft ignores: it only reads pack.mcmeta, pack.png and assets/.

// WorldPackManifest is the file in a world resource pack that records
// which project wrote which files.
const WorldPackManifest = "mms-world-pack.json"

// ExistingOwner owns the files that were in the zip before MMS first
// merged into it, such as a hand-made pack.
const ExistingOwner = "(existing)"

type worldManifest struct {
	// Version lets a later MMS change the format.
	Version  int                      `json:"version"`
	Projects map[string]*worldProject `json:"projects"`
}

type worldProject struct {
	Description string `json:"description,omitempty"`
	// Files are the whole files the project owns.
	Files []string `json:"files,omitempty"`
	// Shared holds the project's own version of each merged file; the
	// file in the zip is rebuilt from every project's version.
	Shared map[string]json.RawMessage `json:"shared,omitempty"`
}

// sharedFile matches the merged files and names how they merge.
var sharedFiles = []struct {
	re *regexp.Regexp
	// list is the array member whose entries are appended; "" merges
	// top-level keys, later projects winning.
	list string
}{
	{regexp.MustCompile(`^assets/[^/]+/lang/[^/]+\.json$`), ""},
	{regexp.MustCompile(`^assets/[^/]+/sounds\.json$`), ""},
	{regexp.MustCompile(`^assets/[^/]+/atlases/.+\.json$`), "sources"},
	{regexp.MustCompile(`^assets/[^/]+/font/.+\.json$`), "providers"},
}

func sharedRule(path string) (list string, ok bool) {
	for _, s := range sharedFiles {
		if s.re.MatchString(path) {
			return s.list, true
		}
	}
	return "", false
}

// WorldPackResult describes a merge.
type WorldPackResult struct {
	Zip []byte
	// Others are the other projects whose files the zip still holds,
	// sorted, with ExistingOwner for files MMS did not write.
	Others []string
	// Replaced lists files of other projects that this build overwrote,
	// as "path (owner)".
	Replaced []string
	// Adopted says a zip without a manifest looked like this project's
	// own earlier build, so its files were replaced rather than kept.
	Adopted bool
}

// MergeWorldPack returns the world resource pack after building owner's
// resource pack files into old, the current zip (nil when there is none).
// owner is the project's namespace, the name its data pack has in
// <world>/datapacks.
func MergeWorldPack(old []byte, owner string, files []OutFile) (*WorldPackResult, error) {
	entries := map[string][]byte{}
	// Zip writes every entry with a zero (1980) timestamp; archivers
	// and file managers stamp the time the file was added.
	zeroTimes := true
	if len(old) > 0 {
		zr, err := zip.NewReader(bytes.NewReader(old), int64(len(old)))
		if err != nil {
			return nil, fmt.Errorf("reading the world's resources.zip: %w", err)
		}
		for _, f := range zr.File {
			if f.FileInfo().IsDir() {
				continue
			}
			if f.ModifiedDate != 0 || f.ModifiedTime != 0 {
				zeroTimes = false
			}
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			b, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, err
			}
			entries[f.Name] = b
		}
	}
	res := &WorldPackResult{}
	man := &worldManifest{Version: 1, Projects: map[string]*worldProject{}}
	if b, ok := entries[WorldPackManifest]; ok {
		if err := json.Unmarshal(b, man); err != nil {
			return nil, fmt.Errorf("reading %s in the world's resources.zip: %w", WorldPackManifest, err)
		}
		if man.Projects == nil {
			man.Projects = map[string]*worldProject{}
		}
		delete(entries, WorldPackManifest)
	} else if len(entries) > 0 {
		// A zip MMS has not merged into: an older MMS build (which wrote
		// the whole zip) or a pack from elsewhere.
		if zeroTimes && ownEarlierBuild(entries, owner, files) {
			res.Adopted = true
			entries = map[string][]byte{}
		} else {
			adoptExisting(man, entries)
		}
	}

	// Take back what this project wrote last time.
	if p := man.Projects[owner]; p != nil {
		for _, f := range p.Files {
			delete(entries, f)
		}
		// Merged files are rebuilt below from what is left.
		for f := range p.Shared {
			delete(entries, f)
		}
	}
	me := &worldProject{Shared: map[string]json.RawMessage{}}
	man.Projects[owner] = me
	owners := map[string]string{} // whole file -> its owner
	for name, p := range man.Projects {
		for _, f := range p.Files {
			owners[f] = name
		}
	}

	var meta, icon []byte
	for _, f := range files {
		switch f.Path {
		case "pack.mcmeta":
			meta = f.Data
			me.Description = packDescription(f.Data)
			continue
		case "pack.png":
			icon = f.Data
			continue
		}
		if _, ok := sharedRule(f.Path); ok && json.Valid(f.Data) {
			me.Shared[f.Path] = json.RawMessage(f.Data)
			continue
		}
		if o, ok := owners[f.Path]; ok && o != owner {
			res.Replaced = append(res.Replaced, f.Path+" ("+o+")")
			p := man.Projects[o]
			p.Files = without(p.Files, f.Path)
		}
		entries[f.Path] = f.Data
		me.Files = append(me.Files, f.Path)
	}
	sort.Strings(me.Files)
	// Drop projects that no longer own anything.
	for name, p := range man.Projects {
		if name != owner && len(p.Files) == 0 && len(p.Shared) == 0 {
			delete(man.Projects, name)
		}
	}

	// Rebuild every merged file from the projects' versions: the existing
	// pack first, then the other projects, then this one, so the project
	// built last wins a clash.
	order := projectOrder(man, owner)
	shared := map[string]bool{}
	for _, name := range order {
		for path := range man.Projects[name].Shared {
			shared[path] = true
		}
	}
	for path := range shared {
		list, _ := sharedRule(path)
		var parts []json.RawMessage
		for _, name := range order {
			if b, ok := man.Projects[name].Shared[path]; ok {
				parts = append(parts, b)
			}
		}
		b, err := mergeJSON(parts, list)
		if err != nil {
			return nil, fmt.Errorf("merging %s: %w", path, err)
		}
		entries[path] = b
	}

	for name := range man.Projects {
		if name != owner {
			res.Others = append(res.Others, name)
		}
	}
	sort.Strings(res.Others)
	sort.Strings(res.Replaced)

	if meta != nil {
		entries["pack.mcmeta"] = worldPackMeta(meta, man, order)
	}
	if icon != nil {
		entries["pack.png"] = icon
	}
	mb, err := json.MarshalIndent(man, "", "  ")
	if err != nil {
		return nil, err
	}
	entries[WorldPackManifest] = append(mb, '\n')

	paths := make([]string, 0, len(entries))
	for p := range entries {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	out := make([]OutFile, len(paths))
	for i, p := range paths {
		out[i] = OutFile{Path: p, Data: entries[p]}
	}
	res.Zip, err = Zip(out)
	return res, err
}

// WriteWorldPack merges owner's resource pack files into the zip at path,
// creating it when it does not exist. The new zip replaces the old one in
// a single rename, so a failed build never leaves half a pack.
func WriteWorldPack(path, owner string, files []OutFile) (*WorldPackResult, error) {
	old, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	res, err := MergeWorldPack(old, owner, files)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".resources-*.zip")
	if err != nil {
		return nil, err
	}
	_, err = tmp.Write(res.Zip)
	if cerr := tmp.Close(); err == nil {
		err = cerr
	}
	// CreateTemp makes the file 0600; the game may run as another user.
	if err == nil {
		err = os.Chmod(tmp.Name(), 0o644)
	}
	if err == nil {
		err = os.Rename(tmp.Name(), path)
	}
	if err != nil {
		os.Remove(tmp.Name())
		return nil, err
	}
	return res, nil
}

// ownEarlierBuild says a zip without a manifest, whose entries all carry
// the zero timestamp Zip writes, is most likely this project's output
// from an MMS that overwrote the whole zip: its pack.mcmeta has the keys
// packMeta writes and nothing else, it has
// files in the project's namespace, and every other namespace in it,
// apart from minecraft, is one the build writes too.
func ownEarlierBuild(entries map[string][]byte, owner string, files []OutFile) bool {
	ours := map[string]bool{}
	for _, f := range files {
		if ns := assetNS(f.Path); ns != "" {
			ours[ns] = true
		}
	}
	var meta map[string]map[string]json.RawMessage
	if json.Unmarshal(entries["pack.mcmeta"], &meta) != nil || len(meta) != 1 || len(meta["pack"]) != 3 {
		return false
	}
	for _, k := range []string{"description", "min_format", "max_format"} {
		if _, ok := meta["pack"][k]; !ok {
			return false
		}
	}
	seen := false
	for p := range entries {
		ns := assetNS(p)
		switch {
		case ns == "" || ns == "minecraft":
		case !ours[ns]:
			return false
		case ns == owner:
			seen = true
		}
	}
	return seen
}

func assetNS(path string) string {
	rest, ok := strings.CutPrefix(path, "assets/")
	if !ok {
		return ""
	}
	ns, _, _ := strings.Cut(rest, "/")
	return ns
}

// adoptExisting records the files of a pack MMS did not write as
// ExistingOwner's, so merging keeps them.
func adoptExisting(man *worldManifest, entries map[string][]byte) {
	p := &worldProject{Shared: map[string]json.RawMessage{}}
	for path, b := range entries {
		if path == "pack.mcmeta" || path == "pack.png" {
			continue
		}
		if _, ok := sharedRule(path); ok && json.Valid(b) {
			p.Shared[path] = json.RawMessage(b)
			continue
		}
		p.Files = append(p.Files, path)
	}
	sort.Strings(p.Files)
	if meta, ok := entries["pack.mcmeta"]; ok {
		p.Description = packDescription(meta)
	}
	man.Projects[ExistingOwner] = p
}

// projectOrder is the merge order: the existing pack, the other projects
// by name, and owner last.
func projectOrder(man *worldManifest, owner string) []string {
	var names []string
	for name := range man.Projects {
		if name != owner && name != ExistingOwner {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	if _, ok := man.Projects[ExistingOwner]; ok && owner != ExistingOwner {
		names = append([]string{ExistingOwner}, names...)
	}
	return append(names, owner)
}

// mergeJSON merges objects key by key, later parts winning; with list set,
// the list member's arrays are appended instead.
func mergeJSON(parts []json.RawMessage, list string) ([]byte, error) {
	if len(parts) == 1 {
		return parts[0], nil
	}
	merged := map[string]json.RawMessage{}
	var items []json.RawMessage
	for _, part := range parts {
		var o map[string]json.RawMessage
		if err := json.Unmarshal(part, &o); err != nil {
			return nil, err
		}
		for k, v := range o {
			if k == list {
				var arr []json.RawMessage
				if err := json.Unmarshal(v, &arr); err != nil {
					return nil, fmt.Errorf("%q is not a list: %w", k, err)
				}
				items = append(items, arr...)
				continue
			}
			merged[k] = v
		}
	}
	if list != "" && items != nil {
		b, err := marshal(items)
		if err != nil {
			return nil, err
		}
		merged[list] = b
	}
	return marshal(merged)
}

func marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// worldPackMeta is this build's pack.mcmeta, with a description naming
// every pack in the zip when there is more than one.
func worldPackMeta(meta []byte, man *worldManifest, order []string) []byte {
	if len(order) < 2 {
		return meta
	}
	var descs []string
	for _, name := range order {
		d := man.Projects[name].Description
		if d == "" {
			d = name
		}
		descs = append(descs, d)
	}
	var o map[string]map[string]json.RawMessage
	if err := json.Unmarshal(meta, &o); err != nil || o["pack"] == nil {
		return meta
	}
	d, _ := marshal(strings.Join(descs, " + "))
	o["pack"]["description"] = bytes.TrimSpace(d)
	b, err := marshal(o)
	if err != nil {
		return meta
	}
	return b
}

// packDescription is a pack.mcmeta's description when it is plain text.
func packDescription(meta []byte) string {
	var o struct {
		Pack struct {
			Description any `json:"description"`
		} `json:"pack"`
	}
	if json.Unmarshal(meta, &o) != nil {
		return ""
	}
	s, _ := o.Pack.Description.(string)
	return s
}

func without(list []string, s string) []string {
	out := list[:0]
	for _, x := range list {
		if x != s {
			out = append(out, x)
		}
	}
	return out
}
