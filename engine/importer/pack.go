package importer

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// File is one generated MMS source file.
type File struct {
	Path string // e.g. "worldgen_noise.mms" or "other_ns/recipe.mms"
	Text string // unformatted until Tidy
}

// Result is an imported data pack.
type Result struct {
	Files []File
	// Namespace is the namespace with the most resources, other than
	// minecraft when there is one: the project's default namespace.
	Namespace   string
	Description string
	// Skipped lists data files with no MMS kind, and files that failed to
	// parse, each with the reason.
	Skipped []string
	Count   int // resources imported
}

// Pack imports a data pack: fsys holds pack.mcmeta and data/. Resources
// are grouped into one MMS file per namespace and folder. Files outside
// the default namespace start with a Namespace declaration.
func Pack(fsys fs.FS, kinds *sema.Kinds, sp *spec.Pack) (*Result, error) {
	res := &Result{}
	if b, err := fs.ReadFile(fsys, "pack.mcmeta"); err == nil {
		var meta struct {
			Pack struct {
				Description json.RawMessage `json:"description"`
			} `json:"pack"`
		}
		if json.Unmarshal(b, &meta) == nil {
			var s string
			if json.Unmarshal(meta.Pack.Description, &s) == nil {
				res.Description = s
			}
		}
	}
	type group struct {
		ns, folder string
		decls      []string
	}
	groups := map[string]*group{}
	perNS := map[string]int{}
	err := fs.WalkDir(fsys, "data", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			if p == "data" && !errors.Is(err, fs.ErrNotExist) {
				return err
			}
			return nil // no data folder: an empty pack
		}
		if d.IsDir() {
			return nil
		}
		parts := strings.SplitN(strings.TrimPrefix(p, "data/"), "/", 2)
		if len(parts) < 2 {
			return nil
		}
		ns, rel := parts[0], parts[1]
		b, err := fs.ReadFile(fsys, p)
		if err != nil {
			res.Skipped = append(res.Skipped, p+": "+err.Error())
			return nil
		}
		var folder, decl string
		switch {
		case strings.HasPrefix(rel, "function/") && strings.HasSuffix(rel, ".mcfunction"):
			folder = "function"
			id := ns + ":" + strings.TrimSuffix(strings.TrimPrefix(rel, "function/"), ".mcfunction")
			decl = FunctionDecl(ns, id, string(b))
		case strings.HasSuffix(rel, ".json"):
			f, id := SplitFolder(kinds, sp, ns, rel)
			if f == "" {
				res.Skipped = append(res.Skipped, p+": no MMS kind for this folder")
				return nil
			}
			v, err := jsonv.Parse(b)
			if err != nil {
				res.Skipped = append(res.Skipped, p+": "+err.Error())
				return nil
			}
			k, _ := kinds.ForFolder(f)
			folder = f
			if k.TagOf != "" {
				decl = TagDecl(k, ns, id, v)
			} else {
				decl = Decl(k, ns, id, v)
			}
		default:
			res.Skipped = append(res.Skipped, p+": not a data pack resource")
			return nil
		}
		key := ns + "|" + folder
		g := groups[key]
		if g == nil {
			g = &group{ns: ns, folder: folder}
			groups[key] = g
		}
		g.decls = append(g.decls, decl)
		perNS[ns]++
		res.Count++
		return nil
	})
	if err != nil {
		return nil, err
	}
	res.Namespace = defaultNamespace(perNS)
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		g := groups[key]
		var sb strings.Builder
		file := strings.ReplaceAll(g.folder, "/", "_") + ".mms"
		if g.ns != res.Namespace {
			file = g.ns + "/" + file
			sb.WriteString("Namespace " + g.ns + "\n\n")
		}
		// Names are qualified against the file's namespace.
		sb.WriteString(strings.Join(g.decls, "\n"))
		res.Files = append(res.Files, File{Path: file, Text: sb.String()})
	}
	sort.Strings(res.Skipped)
	return res, nil
}

func defaultNamespace(perNS map[string]int) string {
	best, n := "", -1
	for ns, c := range perNS {
		if ns == "minecraft" {
			continue
		}
		if c > n || (c == n && ns < best) {
			best, n = ns, c
		}
	}
	if best == "" {
		if perNS["minecraft"] > 0 {
			return "minecraft"
		}
		return "mms"
	}
	return best
}

// SplitFolder finds the resource folder and id of a file path relative to
// data/<ns>/, preferring the longest matching kind folder. It returns ""
// when no MMS kind owns the folder.
func SplitFolder(kinds *sema.Kinds, sp *spec.Pack, ns, rel string) (folder, id string) {
	rel = strings.TrimSuffix(rel, path.Ext(rel))
	if rest, ok := strings.CutPrefix(rel, "tags/"); ok {
		// The registry is the longest prefix that names one.
		parts := strings.Split(rest, "/")
		for i := len(parts) - 1; i >= 1; i-- {
			reg := strings.Join(parts[:i], "/")
			if k, ok := kinds.ForFolder("tags/" + reg); ok {
				return k.Resource, ns + ":" + strings.Join(parts[i:], "/")
			}
		}
		return "", ""
	}
	best := ""
	for _, r := range sp.ResourceKinds() {
		if strings.HasPrefix(rel, r+"/") && len(r) > len(best) {
			best = r
		}
	}
	if best == "" {
		return "", ""
	}
	if _, ok := kinds.ForFolder(best); !ok {
		return "", ""
	}
	return best, ns + ":" + strings.TrimPrefix(rel, best+"/")
}

// Tidy formats the imported files and applies every safe lint fix (bare
// names, folded constants), which never changes the generated pack. Files
// that fail to format are left as they are.
func (res *Result) Tidy(ctx context.Context, sp *spec.Pack) error {
	p := project.New(sp, project.Config{Namespace: res.Namespace})
	for i, f := range res.Files {
		res.Files[i].Text = formatted(f.Text)
		p.SetFile(f.Path, res.Files[i].Text)
	}
	fixed, _, err := p.FixAll(ctx)
	if err != nil {
		return err
	}
	for i, f := range res.Files {
		if text, ok := fixed[f.Path]; ok {
			res.Files[i].Text = formatted(text)
		}
	}
	return nil
}

// TidyIn formats and fixes one declaration's source as if it were in
// namespace ns of the given project, so a shortened name can never be
// captured by one of the project's own declarations.
func TidyIn(ctx context.Context, p *project.Project, ns, src string) (string, error) {
	const scratch = "\x00import.mms" // no real file has this name
	header := "Namespace " + ns + "\n\n"
	p.SetFile(scratch, header+formatted(src))
	defer p.RemoveFile(scratch)
	fixed, _, err := p.FixAll(ctx, scratch)
	if err != nil {
		return "", err
	}
	text, ok := fixed[scratch]
	if !ok {
		text, _ = p.Text(scratch)
	}
	return strings.TrimPrefix(formatted(text), header), nil
}

func formatted(src string) string {
	if out, err := format.Source(src); err == nil {
		return out
	}
	return src
}
