// Command specgen builds the embedded spec pack from vanilla-mcdoc and
// mcmeta. Run it from engine/ with `go run ./specgen`; it downloads the
// pinned inputs (cached under specgen/.cache) and writes
// spec/data/<version>.json.gz plus a short report.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Pinned inputs (decision D6). Bump these, rerun, and review the report.
const (
	version    = "26.3"
	packFormat = 121
	// From mcmeta's version.json for 26.3: resource_pack_version 97.1
	// (the minor part needs no pin; max_format 97 covers it) and
	// data_version.
	resourcePackFormat = 97
	dataVersion        = 5023

	mcdocRepo   = "SpyglassMC/vanilla-mcdoc"
	mcdocCommit = "5aab1da2ff00955cdb66f0d6c2a1154ef691a015" // generated branch
	mcmetaRepo  = "misode/mcmeta"
	mcmetaRef   = "d96c75fec200c4580dd75e033a76341521165461" // tag 26.3-summary
)

func main() {
	out := flag.String("out", "spec/data/"+version+".json.gz", "output file")
	cache := flag.String("cache", "specgen/.cache", "download cache directory")
	flag.Parse()

	if err := run(*out, *cache); err != nil {
		fmt.Fprintln(os.Stderr, "specgen:", err)
		os.Exit(1)
	}
}

func run(out, cache string) error {
	var symbols struct {
		Types       map[string]*spec.Type            `json:"mcdoc"`
		Dispatchers map[string]map[string]*spec.Type `json:"mcdoc/dispatcher"`
	}
	if err := fetchJSON(cache, mcdocRepo, mcdocCommit, "symbols.json", &symbols); err != nil {
		return err
	}
	var registries map[string][]string
	if err := fetchJSON(cache, mcmetaRepo, mcmetaRef, "registries/data.min.json", &registries); err != nil {
		return err
	}
	var blocks map[string][2]json.RawMessage
	if err := fetchJSON(cache, mcmetaRepo, mcmetaRef, "blocks/data.min.json", &blocks); err != nil {
		return err
	}
	var commands spec.CommandNode
	if err := fetchJSON(cache, mcmetaRepo, mcmetaRef, "commands/data.min.json", &commands); err != nil {
		return err
	}

	f := filter{version: version}
	for path, t := range symbols.Types {
		symbols.Types[path] = f.typ(t)
		if symbols.Types[path] == nil {
			delete(symbols.Types, path)
		}
	}
	for id, cases := range symbols.Dispatchers {
		for key, t := range cases {
			if t = f.typ(t); t == nil {
				delete(cases, key)
			} else {
				cases[key] = t
			}
		}
		if len(cases) == 0 {
			delete(symbols.Dispatchers, id)
		}
	}

	pack := &spec.Pack{
		Version:            version,
		PackFormat:         packFormat,
		ResourcePackFormat: resourcePackFormat,
		DataVersion:        dataVersion,
		Sources: map[string]string{
			mcdocRepo:  mcdocCommit,
			mcmetaRepo: mcmetaRef,
		},
		Resources:   map[string]*spec.Type{},
		Registries:  registries,
		Blocks:      map[string]*spec.BlockStates{},
		Commands:    &commands,
		Types:       map[string]*spec.Type{},
		Dispatchers: map[string]map[string]*spec.Type{},
	}

	// Data-pack resource kinds: every kind whose type does not point into
	// ::java::assets:: (decisions D5, D41).
	for kind, t := range symbols.Dispatchers["minecraft:resource"] {
		if b, _ := json.Marshal(t); !strings.Contains(string(b), `"::java::assets::`) {
			pack.Resources[kind] = t
		}
	}

	// Keep every non-asset type, plus whatever those reach.
	r := reach{src: symbols.Types, srcDisp: symbols.Dispatchers, pack: pack}
	for path := range symbols.Types {
		if !strings.HasPrefix(path, "::java::assets::") {
			r.path(path)
		}
	}
	for _, t := range pack.Resources {
		r.typ(t)
	}

	for id, raw := range blocks {
		bs := &spec.BlockStates{}
		if err := json.Unmarshal(raw[0], &bs.Properties); err != nil {
			return fmt.Errorf("block %s: %w", id, err)
		}
		if err := json.Unmarshal(raw[1], &bs.Defaults); err != nil {
			return fmt.Errorf("block %s: %w", id, err)
		}
		pack.Blocks[id] = bs
	}

	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := pack.Encode(&buf); err != nil {
		return err
	}
	if err := os.WriteFile(out, buf.Bytes(), 0o644); err != nil {
		return err
	}

	start := time.Now()
	if _, err := spec.Decode(bytes.NewReader(buf.Bytes())); err != nil {
		return fmt.Errorf("round trip: %w", err)
	}
	fmt.Printf("spec pack %s (format %d): %d types, %d dispatchers, %d resource kinds, %d registries, %d blocks\n",
		version, packFormat, len(pack.Types), len(pack.Dispatchers), len(pack.Resources), len(pack.Registries), len(pack.Blocks))
	fmt.Printf("wrote %s: %d KB gzipped, decodes in %s\n", out, buf.Len()/1024, time.Since(start).Round(time.Millisecond))
	return nil
}

func fetchJSON(cache, repo, ref, file string, v any) error {
	local := filepath.Join(cache, strings.ReplaceAll(repo, "/", "_"), ref, file)
	b, err := os.ReadFile(local)
	if err != nil {
		url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, ref, file)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("GET %s: %s", url, resp.Status)
		}
		if b, err = io.ReadAll(resp.Body); err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(local), 0o755); err == nil {
			_ = os.WriteFile(local, b, 0o644)
		}
	}
	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("%s/%s: %w", repo, file, err)
	}
	return nil
}

// filter removes everything gated out of the target version and strips the
// since/until attributes from what remains.
type filter struct{ version string }

func (f filter) attrs(as []spec.Attribute) ([]spec.Attribute, bool) {
	if !spec.InVersion(as, f.version) {
		return nil, false
	}
	var kept []spec.Attribute
	for _, a := range as {
		if a.Name == "since" || a.Name == "until" {
			continue
		}
		a.Value = f.typ(a.Value)
		kept = append(kept, a)
	}
	return kept, true
}

func (f filter) typ(t *spec.Type) *spec.Type {
	if t == nil {
		return nil
	}
	attrs, ok := f.attrs(t.Attributes)
	if !ok {
		return nil
	}
	t.Attributes = attrs

	var fields []*spec.Field
	for _, fl := range t.Fields {
		as, ok := f.attrs(fl.Attributes)
		if !ok {
			continue
		}
		fl.Attributes = as
		fl.KeyType = f.typ(fl.KeyType)
		if fl.Type = f.typ(fl.Type); fl.Type == nil {
			continue
		}
		fields = append(fields, fl)
	}
	t.Fields = fields

	var members []*spec.Type
	for _, m := range t.Members {
		if m = f.typ(m); m != nil {
			members = append(members, m)
		}
	}
	t.Members = members

	var values []*spec.EnumValue
	for _, v := range t.Values {
		as, ok := f.attrs(v.Attributes)
		if !ok {
			continue
		}
		v.Attributes = as
		values = append(values, v)
	}
	t.Values = values

	t.Item = f.typ(t.Item)
	t.Child = f.typ(t.Child)
	for i := range t.TypeArgs {
		t.TypeArgs[i] = f.typ(t.TypeArgs[i])
	}
	for i := range t.Items {
		t.Items[i] = f.typ(t.Items[i])
	}
	for k, v := range t.Tree {
		t.Tree[k] = f.typ(v)
	}
	return t
}

// reach copies the types and dispatchers a set of roots depends on.
type reach struct {
	src     map[string]*spec.Type
	srcDisp map[string]map[string]*spec.Type
	pack    *spec.Pack
}

func (r *reach) path(p string) {
	if _, done := r.pack.Types[p]; done {
		return
	}
	t, ok := r.src[p]
	if !ok {
		return
	}
	r.pack.Types[p] = t
	r.typ(t)
}

func (r *reach) dispatcher(id string) {
	if _, done := r.pack.Dispatchers[id]; done {
		return
	}
	cases, ok := r.srcDisp[id]
	if !ok {
		return
	}
	r.pack.Dispatchers[id] = cases
	keys := make([]string, 0, len(cases))
	for k := range cases {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		r.typ(cases[k])
	}
}

func (r *reach) typ(t *spec.Type) {
	if t == nil {
		return
	}
	switch t.Kind {
	case "reference":
		r.path(t.Path)
	case "dispatcher":
		r.dispatcher(t.Registry)
	}
	for _, a := range t.Attributes {
		r.typ(a.Value)
	}
	for _, f := range t.Fields {
		for _, a := range f.Attributes {
			r.typ(a.Value)
		}
		r.typ(f.KeyType)
		r.typ(f.Type)
	}
	for _, m := range t.Members {
		r.typ(m)
	}
	r.typ(t.Item)
	r.typ(t.Child)
	for _, a := range t.TypeArgs {
		r.typ(a)
	}
	for _, i := range t.Items {
		r.typ(i)
	}
	for _, v := range t.Tree {
		r.typ(v)
	}
	for _, p := range t.TypeParams {
		r.path(p.Path)
	}
}
