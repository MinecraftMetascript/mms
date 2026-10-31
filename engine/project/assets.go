package project

import (
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/sema"
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

// providedAssets returns the ids asset helpers give files, by helper.
func (r *Result) providedAssets() map[string]bool {
	out := map[string]bool{}
	for _, u := range r.assets {
		out[u.Helper+"|"+u.ID] = true
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
	helperFor := map[string]string{"texture": "Texture", "sound_event": "Sound", "weighed_sound_event": "Sound", "structure": "Structure"}
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
				msg += fmt.Sprintf(", or be declared here: %s %s = { ... }", k, name)
			}
			hits = append(hits, hit{n.Span(), msg})
		}
		sort.Slice(hits, func(i, j int) bool { return hits[i].sp.Start < hits[j].sp.Start })
		for _, h := range hits {
			l.out[d.Unit.Path] = append(l.out[d.Unit.Path], syntax.Diagnostic{Code: "lint/" + l.rule.ID, Severity: l.sev, Message: h.msg, Sp: h.sp})
		}
	}
}
