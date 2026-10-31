package project

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// OutFile is one generated pack file with the source it came from.
type OutFile struct {
	Path string // slash-separated, relative to the pack root
	Data []byte
	// Unit and Span locate the declaration that produced the file; Unit
	// is "" for generated files such as pack.mcmeta.
	Unit string
	Span syntax.Span
}

// Emit lays out the data pack: one file per resource, merged tags, raw
// command functions, generated load and tick functions, and pack.mcmeta.
// Output is sorted by path and byte-identical across runs.
func (r *Result) Emit() []OutFile {
	var out []OutFile
	type tagFile struct {
		values  jsonv.Array
		seen    map[string]bool
		decl    *Decl
		replace bool
	}
	tags := map[string]*tagFile{}
	addTag := func(path string, values jsonv.Array, d *Decl) {
		t := tags[path]
		if t == nil {
			t = &tagFile{seen: map[string]bool{}, decl: d}
			tags[path] = t
		}
		for _, v := range values {
			s := jsonv.Compact(v)
			if !t.seen[s] {
				t.seen[s] = true
				t.values = append(t.values, v)
			}
		}
	}
	load := map[string][]string{}
	tick := map[string][]string{}
	for _, d := range r.Order {
		id := d.ID()
		ns, path := splitID(id)
		switch {
		case d.Kind.TagOf != "":
			p := "data/" + ns + "/" + d.Kind.Resource + "/" + path + ".json"
			addTag(p, d.Tag, d)
			if d.TagReplace {
				tags[p].replace = true
			}
		case d.Kind.IsResource():
			if d.Value == nil {
				continue
			}
			out = append(out, OutFile{Path: "data/" + ns + "/" + d.Kind.Resource + "/" + path + ".json", Data: jsonv.Marshal(d.Value), Unit: d.Unit.Path, Span: d.D.Sp})
		case d.Kind.MMS == "Function":
			out = append(out, OutFile{Path: "data/" + ns + "/function/" + path + ".mcfunction", Data: lines(d.Lines), Unit: d.Unit.Path, Span: d.D.Sp})
		case d.Kind.MMS == "Objective":
			load[d.NS] = append(load[d.NS], d.Lines...)
		case d.Kind.MMS == "OnLoad":
			load[d.NS] = append(load[d.NS], d.Lines...)
		case d.Kind.MMS == "OnTick":
			tick[d.NS] = append(tick[d.NS], d.Lines...)
		}
	}
	for _, gen := range []struct {
		name string
		fns  map[string][]string
	}{{"load", load}, {"tick", tick}} {
		for ns, ls := range gen.fns {
			out = append(out, OutFile{Path: "data/" + ns + "/function/mms_gen/" + gen.name + ".mcfunction", Data: lines(ls)})
			addTag("data/minecraft/tags/function/"+gen.name+".json", jsonv.Array{jsonv.String(ns + ":mms_gen/" + gen.name)}, nil)
		}
	}
	for path, t := range tags {
		o := &jsonv.Object{}
		if t.replace {
			o.Set("replace", jsonv.Bool(true))
		}
		o.Set("values", t.values)
		f := OutFile{Path: path, Data: jsonv.Marshal(o)}
		if t.decl != nil {
			f.Unit, f.Span = t.decl.Unit.Path, t.decl.D.Sp
		}
		out = append(out, f)
	}
	out = append(out, OutFile{Path: "pack.mcmeta", Data: r.packMeta()})
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out
}

func (r *Result) packMeta() []byte {
	format := jsonv.Int(r.pack.PackFormat)
	pack := &jsonv.Object{}
	pack.Set("description", jsonv.String(r.cfg.Description))
	pack.Set("min_format", format)
	pack.Set("max_format", format)
	return jsonv.Marshal(&jsonv.Object{Members: []jsonv.Member{{Key: "pack", Value: pack}}})
}

func lines(ls []string) []byte {
	if len(ls) == 0 {
		return nil
	}
	return []byte(strings.Join(ls, "\n") + "\n")
}

func splitID(id string) (string, string) {
	if i := strings.IndexByte(id, ':'); i >= 0 {
		return id[:i], id[i+1:]
	}
	return "minecraft", id
}

// WriteDir writes files under dir, removing what a previous build left
// there first so deleted declarations do not linger.
func WriteDir(dir string, files []OutFile) error {
	if err := os.RemoveAll(filepath.Join(dir, "data")); err != nil {
		return err
	}
	for _, f := range files {
		p := filepath.Join(dir, filepath.FromSlash(f.Path))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(p, f.Data, 0o644); err != nil {
			return err
		}
	}
	return nil
}

// Zip returns the files as a zip archive with fixed timestamps, so the
// archive is reproducible.
func Zip(files []OutFile) ([]byte, error) {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for _, f := range files {
		w, err := zw.CreateHeader(&zip.FileHeader{Name: f.Path, Method: zip.Deflate})
		if err != nil {
			return nil, err
		}
		if _, err := w.Write(f.Data); err != nil {
			return nil, err
		}
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
