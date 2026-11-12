package project

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// assetUse is one asset helper call in a checked declaration.
type assetUse struct {
	*sema.Asset
	decl *Decl
	sp   syntax.Span
}

// collectAssets gathers every asset helper use, in source order, and
// reports two different files that would land on the same output path.
func (r *Result) collectAssets() {
	r.assetErr = map[string][]syntax.Diagnostic{}
	first := map[string]assetUse{}
	for _, d := range r.Order {
		if hasErrors(d.Diags) {
			continue // not emitted, so its files are not either
		}
		var uses []assetUse
		for n, in := range d.Info {
			if in.Asset != nil {
				uses = append(uses, assetUse{in.Asset, d, n.Span()})
			}
		}
		sort.Slice(uses, func(i, j int) bool { return uses[i].sp.Start < uses[j].sp.Start })
		for _, u := range uses {
			prev, seen := first[u.Out]
			if !seen {
				first[u.Out] = u
				r.assets = append(r.assets, u)
				continue
			}
			if prev.Src == u.Src && metaText(prev.Meta) != metaText(u.Meta) {
				r.assetErr[d.Unit.Path] = append(r.assetErr[d.Unit.Path], syntax.Diagnostic{
					Code: "mms/asset-conflict", Severity: syntax.SeverityError, Sp: u.sp,
					Message: fmt.Sprintf("'%s' is used as %s '%s' with different named arguments, but the texture has one %s.mcmeta; write the same ones at every use, or give this one its own id name", u.Src, strings.ToLower(u.Helper), u.ID, path.Base(u.Out)),
				})
				continue
			}
			if prev.Src != u.Src {
				r.assetErr[d.Unit.Path] = append(r.assetErr[d.Unit.Path], syntax.Diagnostic{
					Code: "mms/asset-conflict", Severity: syntax.SeverityError, Sp: u.sp,
					Message: fmt.Sprintf("'%s' and '%s' both become %s '%s'; give this one its own id name, e.g. %s(\"%s\", \"%s_2\")",
						u.Src, prev.Src, strings.ToLower(u.Helper), u.ID, u.Helper, u.Src, path.Base(strings.SplitN(u.ID, ":", 2)[1])),
				})
			}
		}
	}
}

// metaText is a texture's .png.mcmeta as JSON text, or "" without one.
func metaText(m *jsonv.Object) string {
	if m == nil {
		return ""
	}
	return string(jsonv.Marshal(m))
}

// providedAssets returns the ids asset helpers give files, by helper.
func (r *Result) providedAssets() map[string]bool {
	out := map[string]bool{}
	for _, u := range r.assets {
		out[u.Helper+"|"+u.ID] = true
	}
	return out
}

// SoundIDs returns the sound event ids the project's Sound("x.ogg") uses
// give their files, in source order. The build writes them to
// sounds.json, so commands such as /playsound can name them.
func (r *Result) SoundIDs() []string {
	var out []string
	for _, u := range r.assets {
		if u.Helper == "Sound" {
			out = append(out, u.ID)
		}
	}
	return out
}

// emitAssets lists the copied files and the files generated for them:
// sounds.json, the resource pack's pack.mcmeta, and pack.png from the
// icon. Copies carry their source in Asset and no Data until ReadAssets.
// hasResources says the project already declares resource-pack files
// (models, block states).
func (r *Result) emitAssets(hasResources bool) []OutFile {
	var out []OutFile
	sounds := map[string][]string{}
	for _, u := range r.assets {
		f := OutFile{Path: u.Out, Asset: u.Src, Pack: u.Pack, Unit: u.decl.Unit.Path, Span: u.sp}
		out = append(out, f)
		if u.Meta != nil {
			out = append(out, OutFile{Path: u.Out + ".mcmeta", Data: jsonv.Marshal(u.Meta), Pack: u.Pack, Unit: u.decl.Unit.Path, Span: u.sp})
		}
		if u.Pack == sema.ResourcePack {
			hasResources = true
		}
		if u.Helper == "Sound" {
			ns, name, _ := strings.Cut(u.ID, ":")
			sounds[ns] = append(sounds[ns], name)
		}
	}
	for ns, names := range sounds {
		sort.Strings(names)
		o := &jsonv.Object{}
		for _, name := range names {
			snd := &jsonv.Object{}
			snd.Set("name", jsonv.String(ns+":"+name))
			ev := &jsonv.Object{}
			ev.Set("sounds", jsonv.Array{snd})
			o.Set(name, ev)
		}
		out = append(out, OutFile{Path: "assets/" + ns + "/sounds.json", Data: jsonv.Marshal(o), Pack: sema.ResourcePack})
	}
	if r.cfg.Icon != "" {
		out = append(out, OutFile{Path: "pack.png", Asset: r.cfg.Icon})
		if hasResources {
			out = append(out, OutFile{Path: "pack.png", Asset: r.cfg.Icon, Pack: sema.ResourcePack})
		}
	}
	if hasResources {
		desc := r.cfg.ResourcePack.Description
		if desc == "" {
			desc = r.cfg.Description
		}
		out = append(out, OutFile{Path: "pack.mcmeta", Data: packMeta(desc, r.pack.ResourcePackFormat), Pack: sema.ResourcePack})
	}
	return out
}

// ReadAssets fills in the Data of copied asset files from fsys.
func ReadAssets(files []OutFile, fsys fs.FS) error {
	for i := range files {
		if files[i].Asset == "" || files[i].Data != nil {
			continue
		}
		if fsys == nil {
			return fmt.Errorf("%s: no project files to copy it from", files[i].Asset)
		}
		b, err := fs.ReadFile(fsys, files[i].Asset)
		if err != nil {
			return err
		}
		files[i].Data = b
	}
	return nil
}

// SplitPacks separates the data pack's files from the resource pack's.
func SplitPacks(files []OutFile) (data, resources []OutFile) {
	for _, f := range files {
		if f.Pack == sema.ResourcePack {
			resources = append(resources, f)
		} else {
			data = append(data, f)
		}
	}
	return data, resources
}

// lintExternalAsset reports ids of asset files (textures, sounds,
// structures) in a namespace the project declares that no asset helper
// in the project provides. The pack still loads; the game shows a
// missing texture, plays nothing or skips the structure.
func lintExternalAsset(l *lintRun) {
	provided := l.r.providedAssets()
	declared := map[string]bool{l.r.cfg.Namespace: true}
	for _, d := range l.r.Order {
		declared[d.NS] = true
	}
	declarable := map[string]string{"model": "Model", "item_definition": "ItemModel", "equipment": "Equipment", "atlas": "Atlas", "font": "Font", "post_effect": "PostEffect"}
	helperFor := map[string]string{"texture": "Texture", "sound_event": "Sound", "weighed_sound_event": "Sound", "structure": "StructureFile"}
	for _, d := range l.r.Order {
		if hasErrors(d.Diags) {
			continue
		}
		type hit struct {
			sp  syntax.Span
			msg string
		}
		var hits []hit
		for n, in := range d.Info {
			if in.Asset != nil || in.Symbol != nil || in.Bare != sema.BareID || !sema.IsAssetRegistry(in.Folder) {
				continue
			}
			ns, _, _ := strings.Cut(in.ID, ":")
			if ns == "minecraft" || !declared[ns] {
				continue
			}
			h := helperFor[in.Folder]
			if h != "" && provided[h+"|"+in.ID] {
				continue
			}
			what := strings.ReplaceAll(in.Folder, "_", " ")
			msg := fmt.Sprintf("this project does not provide the %s '%s'; it must come from another pack", what, in.ID)
			if h != "" {
				msg += fmt.Sprintf(", or include the file with %s(\"path\")", h)
			} else if k := declarable[in.Folder]; k != "" {
				name := in.ID
				if ns == l.r.cfg.Namespace {
					_, name, _ = strings.Cut(in.ID, ":")
				}
				msg += fmt.Sprintf(", or be declared here: %s %s { ... }", k, name)
			}
			hits = append(hits, hit{n.Span(), msg})
		}
		sort.Slice(hits, func(i, j int) bool { return hits[i].sp.Start < hits[j].sp.Start })
		for _, h := range hits {
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Message: h.msg, Sp: h.sp})
		}
	}
}

// modelAtlases are the atlases the game looks a model's textures up in.
var modelAtlases = []string{"blocks", "items"}

// modelAtlasesOf returns the atlases among modelAtlases that stitch a
// sprite named id: from the file Texture(...) copies when file is set,
// from any file otherwise. Their sources are vanilla's plus those of the
// project's Atlas blocks and Atlas items declarations.
func (r *Result) modelAtlasesOf(id string, file bool) []string {
	r.atlasOnce.Do(func() {
		r.atlasSources = map[string][]spec.AtlasSource{}
		for _, a := range modelAtlases {
			r.atlasSources[a] = append([]spec.AtlasSource(nil), r.pack.Atlases[a]...)
		}
		for _, d := range r.Order {
			if d.Kind.Resource != "atlas" || d.Value == nil {
				continue
			}
			name, ok := strings.CutPrefix(d.ID(), "minecraft:")
			if !ok || (name != "blocks" && name != "items") {
				continue
			}
			var a struct {
				Sources []spec.AtlasSource `json:"sources"`
			}
			// The declaration checked against the atlas schema, so it decodes.
			if json.Unmarshal(jsonv.Marshal(d.Value), &a) == nil {
				r.atlasSources[name] = append(r.atlasSources[name], a.Sources...)
			}
		}
	})
	var out []string
	for _, a := range modelAtlases {
		for _, s := range r.atlasSources[a] {
			if file {
				if sp, ok := s.SpriteOf(id); ok && sp == id {
					out = append(out, a)
					break
				}
			} else if s.Makes(id) {
				out = append(out, a)
				break
			}
		}
	}
	return out
}

// lintMixedAtlas reports a model whose faces show textures from
// different atlases, one only in the blocks atlas and another only in the
// items atlas: the game cannot draw a model from two atlases. The
// textures are those its parent chain resolves the faces' slots to,
// inherited ones included; slots no face shows do not count. When the
// chain reaches a model MMS does not know, the model's own textures are
// checked instead.
func lintMixedAtlas(l *lintRun) {
	provided := l.r.providedAssets()
	for _, d := range l.r.Order {
		if d.Kind.Resource != "model" || hasErrors(d.Diags) || d.D == nil {
			continue
		}
		type use struct {
			sp syntax.Span
			id string
		}
		// Where the model itself writes each texture id, first use first.
		own := map[string]syntax.Span{}
		var ownIDs []use
		for n, in := range d.Info {
			if !in.ModelTexture || in.ID == "" {
				continue
			}
			if sp, ok := own[in.ID]; !ok || n.Span().Start < sp.Start {
				own[in.ID] = n.Span()
			}
			ownIDs = append(ownIDs, use{n.Span(), in.ID})
		}
		var shown []use
		if slots, known := l.r.ModelSlots(d); known {
			at := d.D.Name.Sp
			if f := modelField(d.D, "textures"); f != nil {
				at = f.Key.Sp
			}
			seen := map[string]bool{}
			for _, sl := range slots {
				if !sl.Used || sl.Filled == "" || seen[sl.Filled] {
					continue
				}
				seen[sl.Filled] = true
				sp, ok := own[sl.Filled]
				if !ok {
					sp = at // inherited from a parent
				}
				shown = append(shown, use{sp, sl.Filled})
			}
		} else {
			shown = ownIDs
		}
		only := map[string][]use{}
		for _, u := range shown {
			if as := l.r.modelAtlasesOf(u.id, provided["Texture|"+u.id]); len(as) == 1 {
				only[as[0]] = append(only[as[0]], u)
			}
		}
		if len(only["blocks"]) == 0 || len(only["items"]) == 0 {
			continue
		}
		for _, us := range only {
			sort.Slice(us, func(i, j int) bool {
				if us[i].sp.Start != us[j].sp.Start {
					return us[i].sp.Start < us[j].sp.Start
				}
				return us[i].id < us[j].id
			})
		}
		// Report the side with fewer textures, against the other's first.
		few, many := "items", "blocks"
		if len(only["blocks"]) < len(only["items"]) {
			few, many = many, few
		}
		other := only[many][0]
		for _, u := range only[few] {
			msg := fmt.Sprintf("'%s' is only in the %s atlas, but '%s' in this model is only in the %s atlas; the game draws a model from one atlas, so keep its textures under one of textures/block/ or textures/item/", u.id, few, other.id, many)
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Message: msg, Sp: u.sp})
		}
	}
}

// lintTextureAtlas reports a texture a model names, in a namespace other
// than minecraft, that neither the blocks nor the items atlas stitches:
// the game shows the missing texture. The atlases' sources are vanilla's
// plus those of the project's own Atlas blocks and Atlas items
// declarations in the minecraft namespace.
func lintTextureAtlas(l *lintRun) {
	covered := func(id string, file bool) bool { return len(l.r.modelAtlasesOf(id, file)) > 0 }
	for _, d := range l.r.Order {
		if hasErrors(d.Diags) {
			continue
		}
		type hit struct {
			sp    syntax.Span
			msg   string
			fixes []syntax.Fix
		}
		var hits []hit
		for n, in := range d.Info {
			// Vanilla's own textures are stitched; a file Texture(...)
			// writes into minecraft, such as minecraft:custom/foo, may not be.
			if !in.ModelTexture || in.ID == "" || (in.Asset == nil && strings.HasPrefix(in.ID, "minecraft:")) || covered(in.ID, in.Asset != nil) {
				continue
			}
			ns, p, _ := strings.Cut(in.ID, ":")
			msg := fmt.Sprintf("no atlas stitches the texture '%s', so models show the missing texture; models find textures under textures/block/ and textures/item/", in.ID)
			var fixes []syntax.Fix
			if in.Asset != nil {
				msg += fmt.Sprintf(", so name it like %s(\"%s\", \"block/%s\")", in.Asset.Helper, in.Asset.Src, path.Base(p))
				if call, ok := n.(*syntax.Call); ok && len(call.Args) >= 2 && call.Args[1].Name == nil {
					fixes = []syntax.Fix{{Title: "Name it 'block/" + path.Base(p) + "'", Edits: []syntax.TextEdit{{Sp: call.Args[1].Value.Span(), Text: `"block/` + path.Base(p) + `"`}}}}
				}
			} else {
				msg += fmt.Sprintf(", such as '%s:block/%s'", ns, path.Base(p))
			}
			msg += ", or add it to an Atlas blocks declaration in the minecraft namespace"
			hits = append(hits, hit{n.Span(), msg, fixes})
		}
		sort.Slice(hits, func(i, j int) bool { return hits[i].sp.Start < hits[j].sp.Start })
		for _, h := range hits {
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Message: h.msg, Sp: h.sp, Fixes: h.fixes})
		}
	}
}
