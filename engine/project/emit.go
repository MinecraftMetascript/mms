package project

import (
	"archive/zip"
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
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
	// Pack is the pack the file belongs to; resource pack files exist
	// only when the project uses Texture or Sound.
	Pack sema.PackKind
	// Asset is the project file a copied asset comes from. Its Data is
	// nil until ReadAssets.
	Asset string
}

// Emit lays out the data pack: one file per resource, merged tags, raw
// command functions, generated load and tick functions, and pack.mcmeta;
// then the asset files the project uses, and the resource pack when it
// uses textures or sounds. Output is sorted by pack, then path, and is
// byte-identical across runs. Copied assets need ReadAssets.
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
	objectives := map[string][]string{}
	globals := map[string][]string{}
	load := map[string][]string{}
	tick := map[string][]string{}
	consts := map[string]map[int32]bool{}
	ret := map[string]bool{}
	used := map[string]bool{} // namespaces whose objective the code uses
	fnFiles := map[string]bool{}
	addFn := func(d *Decl) {
		for id, ls := range d.Fn.Files {
			if fnFiles[id] {
				continue // identical OnLoad bodies share their functions
			}
			fnFiles[id] = true
			fns, fpath := splitID(id)
			out = append(out, OutFile{Path: "data/" + fns + "/function/" + fpath + ".mcfunction", Data: lines(ls), Unit: d.Unit.Path, Span: d.D.Sp})
			markUsed(used, ls)
		}
		ns, _ := splitID(d.Fn.ID())
		if consts[ns] == nil {
			consts[ns] = map[int32]bool{}
		}
		for v := range d.Fn.Consts {
			consts[ns][v] = true
		}
		if d.Fn.Ret {
			ret[ns] = true
		}
	}
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
		case d.Fn != nil:
			addFn(d)
			switch d.Kind.MMS {
			case "OnLoad":
				load[d.NS] = append(load[d.NS], d.Lines...)
			case "OnTick":
				tick[d.NS] = append(tick[d.NS], d.Lines...)
			}
		case d.Kind.MMS == "Objective":
			objectives[d.NS] = append(objectives[d.NS], d.Lines...)
		case d.Kind.MMS == "Global":
			gns := d.NS
			if strings.Contains(d.D.NameText(), ":") {
				gns, _ = splitID(d.D.NameText())
			}
			globals[gns] = append(globals[gns], d.Lines...)
			used[gns] = true
		case d.Kind.MMS == "OnLoad":
			load[d.NS] = append(load[d.NS], d.Lines...)
			markUsed(used, d.Lines)
		case d.Kind.MMS == "OnTick":
			tick[d.NS] = append(tick[d.NS], d.Lines...)
			markUsed(used, d.Lines)
		}
	}
	// The load function sets up objectives first, then the backend's
	// objective and constants, then globals, then OnLoad bodies.
	nss := map[string]bool{}
	for _, m := range []map[string][]string{objectives, globals, load} {
		for ns := range m {
			nss[ns] = true
		}
	}
	for ns := range used {
		nss[ns] = true
	}
	for ns := range ret {
		out = append(out, OutFile{Path: "data/" + ns + "/function/mms_gen/ret.mcfunction", Data: lines(fn.RetHelper(ns))})
	}
	fullLoad := map[string][]string{}
	for ns := range nss {
		ls := append([]string(nil), objectives[ns]...)
		if used[ns] || len(consts[ns]) > 0 || ret[ns] {
			ls = append(ls, fn.Setup(ns, consts[ns])...)
		}
		ls = append(ls, globals[ns]...)
		ls = append(ls, load[ns]...)
		if len(ls) > 0 {
			fullLoad[ns] = ls
		}
	}
	for _, gen := range []struct {
		name string
		fns  map[string][]string
	}{{"load", fullLoad}, {"tick", tick}} {
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
	out = append(out, OutFile{Path: "pack.mcmeta", Data: packMeta(r.cfg.Description, r.pack.PackFormat)})
	out = append(out, r.emitAssets()...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Pack != out[j].Pack {
			return out[i].Pack < out[j].Pack
		}
		return out[i].Path < out[j].Path
	})
	return out
}

func packMeta(description string, packFormat int) []byte {
	format := jsonv.Int(packFormat)
	pack := &jsonv.Object{}
	pack.Set("description", jsonv.String(description))
	pack.Set("min_format", format)
	pack.Set("max_format", format)
	return jsonv.Marshal(&jsonv.Object{Members: []jsonv.Member{{Key: "pack", Value: pack}}})
}

// markUsed notes the namespaces whose backend objective (mms.<ns>) the
// lines use.
func markUsed(used map[string]bool, ls []string) {
	for _, l := range ls {
		for _, w := range strings.Fields(l) {
			if ns, ok := strings.CutPrefix(w, "mms."); ok && ns != "" {
				used[ns] = true
			}
		}
	}
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

// WriteDir writes one pack's files under dir, removing what a previous
// build left there first (data/, assets/ and pack.png) so deleted
// declarations and unused assets do not linger.
func WriteDir(dir string, files []OutFile) error {
	for _, old := range []string{"data", "assets", "pack.png"} {
		if err := os.RemoveAll(filepath.Join(dir, old)); err != nil {
			return err
		}
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
