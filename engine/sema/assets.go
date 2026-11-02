package sema

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Asset files: Texture("art/sunset.png"), Sound("music/tune.ogg") and
// Structure("builds/house.nbt") name a file anywhere in the project. The
// call stands for the id the file gets in the pack, and the build copies
// the file to the path that id implies. Files nothing uses are never
// copied (decision D69).

// PackKind says which pack an output file belongs to.
type PackKind uint8

const (
	DataPack PackKind = iota
	ResourcePack
)

// Asset is one use of an asset helper, recorded in the call's Info.
type Asset struct {
	Helper string   // "Texture", "Sound" or "Structure"
	Src    string   // project-relative source file
	ID     string   // ns:name, the value written to the JSON
	Out    string   // pack-relative output path
	Pack   PackKind // which pack Out belongs to
}

type assetHelper struct {
	registries []string // id registries the helper may stand in for
	ext        string
	magic      []byte // the file's first bytes
	what       string
	pack       PackKind
	out        func(ns, prefix, name string) string
}

var assetHelpers = map[string]*assetHelper{
	"Texture": {
		registries: []string{"texture"}, ext: ".png", magic: []byte("\x89PNG\r\n\x1a\n"), what: "a PNG image", pack: ResourcePack,
		out: func(ns, prefix, name string) string { return "assets/" + ns + "/textures/" + prefix + name + ".png" },
	},
	"Sound": {
		registries: []string{"sound_event", "weighed_sound_event"}, ext: ".ogg", magic: []byte("OggS"), what: "an Ogg Vorbis sound", pack: ResourcePack,
		out: func(ns, _, name string) string { return "assets/" + ns + "/sounds/" + name + ".ogg" },
	},
	"Structure": {
		registries: []string{"structure"}, ext: ".nbt", magic: []byte{0x1f, 0x8b}, what: "a gzipped structure (.nbt)", pack: DataPack,
		out: func(ns, _, name string) string { return "data/" + ns + "/structure/" + name + ".nbt" },
	},
}

// AssetHelperNames lists the helpers, for completion and docs.
var AssetHelperNames = []string{"Sound", "Structure", "Texture"}

// assetRegistries are id registries whose ids name files rather than
// data pack resources: MMS cannot declare them, so an id in the
// project's namespace is not an error when no declaration has it. The
// external-asset lint rule reports the ones nothing in the project
// provides.
var assetRegistries = map[string]bool{
	"texture": true, "sound_event": true, "weighed_sound_event": true, "structure": true,
	"equipment": true, "atlas": true, "font": true, "item_definition": true, "model": true, "post_effect": true,
}

// IsAssetRegistry reports whether ids of registry reg name asset files.
func IsAssetRegistry(reg string) bool { return assetRegistries[reg] }

// assetCallName returns the helper an expression calls, or "".
func assetCallName(e syntax.Expr) string {
	call, ok := e.(*syntax.Call)
	if !ok {
		return ""
	}
	id, ok := call.Fn.(*syntax.Ident)
	if !ok || assetHelpers[id.Name] == nil {
		return ""
	}
	return id.Name
}

// assetFits reports whether e is an asset helper call that can stand for
// an id of the given registry.
func (c *Checker) assetFits(e syntax.Expr, id idInfo) bool {
	h := assetHelpers[assetCallName(e)]
	if h == nil || !id.Present || id.Tags == "required" {
		return false
	}
	reg := c.registryName(id.Registry)
	for _, r := range h.registries {
		if r == reg {
			return true
		}
	}
	return false
}

var assetNamePattern = regexp.MustCompile(`^[a-z0-9_./-]+$`)

// assetName turns a file name into an id path: the base name without
// extension, lowercased, with characters ids cannot hold replaced by '_'.
func assetName(src string) string {
	base := strings.TrimSuffix(path.Base(src), path.Ext(src))
	var sb strings.Builder
	for _, r := range strings.ToLower(base) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '_', r == '-', r == '.':
			sb.WriteRune(r)
		default:
			sb.WriteByte('_')
		}
	}
	if sb.Len() == 0 {
		return "asset"
	}
	return sb.String()
}

// assetCall checks Texture(path, name?), Sound(...) or Structure(...) at
// an id position and returns the id it stands for.
func (c *Checker) assetCall(e syntax.Expr, v tview) jsonv.Value {
	call := e.(*syntax.Call)
	name := assetCallName(e)
	h := assetHelpers[name]
	id := idOf(v.T)
	if !c.assetFits(e, id) {
		c.errorf("mms/type-mismatch", call.Fn.Span(), "%s(...) makes a %s id, but this field expects %s", name, h.registries[0], describe(v))
		return nil
	}
	var src, idName string
	var srcSp, nameSp syntax.Span
	for i, a := range call.Args {
		s, ok := a.Value.(*syntax.String)
		switch {
		case a.Name != nil:
			c.errorf("mms/type-mismatch", a.Sp, "%s takes a file path and an optional id name, without argument names", name)
			return nil
		case i > 1:
			c.errorf("mms/type-mismatch", a.Sp, "%s takes at most two arguments: a file path and an id name", name)
			return nil
		case !ok:
			c.errorf("mms/type-mismatch", a.Value.Span(), "expected a quoted %s", map[bool]string{true: "file path", false: "id name"}[i == 0])
			return nil
		case i == 0:
			src, srcSp = s.Value, s.Sp
		default:
			idName, nameSp = s.Value, s.Sp
		}
	}
	if len(call.Args) == 0 {
		c.errorf("mms/type-mismatch", call.Sp, "%s needs the path of a %s file, relative to the project: %s(\"file%s\")", name, h.ext, name, h.ext)
		return nil
	}
	in := c.info(e)
	in.Type = v
	src = path.Clean(strings.ReplaceAll(src, "\\", "/"))
	if path.IsAbs(src) || src == ".." || strings.HasPrefix(src, "../") {
		c.errorf("mms/asset", srcSp, "asset paths are relative to the project folder and stay inside it")
		return nil
	}
	if !strings.EqualFold(path.Ext(src), h.ext) {
		c.errorf("mms/asset", srcSp, "%s expects %s file, not '%s'", name, map[string]string{".png": "a .png", ".ogg": "an .ogg", ".nbt": "an .nbt"}[h.ext], path.Base(src))
		return nil
	}
	if idName == "" && name == "Texture" && c.textureName != "" {
		idName = c.textureName
	} else if idName == "" {
		idName = assetName(src)
	} else if idName = c.stripAssetExt(idName, h.registries[0], call.Args[1].Value); !assetNamePattern.MatchString(idName) || strings.Contains(idName, "..") {
		c.errorf("mms/invalid-name", nameSp, "an id name is lowercase letters, digits, '_', '-', '.' and '/'")
		return nil
	}
	head, ok := c.assetHead(src, len(h.magic))
	if !ok && c.Files == nil {
		c.errorf("mms/asset", srcSp, "asset files are not available here yet; build with the mms command line to include '%s'", src)
		return nil
	}
	if !ok {
		c.errorf("mms/asset", srcSp, "no file '%s' in the project", src)
		return nil
	}
	if !bytes.Equal(head, h.magic) {
		c.errorf("mms/asset", srcSp, "'%s' is not %s", src, h.what)
		return nil
	}
	ns := c.NS
	full := ns + ":" + idName
	in.ID = full
	in.Asset = &Asset{Helper: name, Src: src, ID: full, Out: h.out(ns, id.Path, idName), Pack: h.pack}
	return jsonv.String(full)
}

// assetHead reads a file's first n bytes; ok is false when it does not
// exist. The outcome is recorded in Deps as "a|path|version".
func (c *Checker) assetHead(src string, n int) ([]byte, bool) {
	var head []byte
	ok, sig := false, ""
	if c.Files != nil {
		if f, err := c.Files.Open(src); err == nil {
			if st, err := f.Stat(); err == nil && !st.IsDir() {
				ok, sig = true, fileSig(st)
				head = make([]byte, n)
				k, _ := io.ReadFull(f, head)
				head = head[:k]
			}
			f.Close()
		}
	}
	if c.Deps != nil {
		c.Deps["a|"+src+"|"+sig] = ok
	}
	return head, ok
}

// fileSig identifies a file's version by size and modification time.
func fileSig(st fs.FileInfo) string {
	return strconv.FormatInt(st.Size(), 10) + "." + strconv.FormatInt(st.ModTime().UnixNano(), 10)
}

// AssetDepsHold reports whether every asset file lookup in deps has the
// same outcome in fsys.
func AssetDepsHold(fsys fs.FS, deps map[string]bool) bool {
	for k, found := range deps {
		rest, ok := strings.CutPrefix(k, "a|")
		if !ok {
			continue
		}
		i := strings.LastIndexByte(rest, '|')
		src, sig := rest[:i], rest[i+1:]
		same := !found // the file is still missing
		if fsys != nil {
			if st, err := fs.Stat(fsys, src); err == nil && !st.IsDir() {
				same = found && fileSig(st) == sig
			}
		}
		if !same {
			return false
		}
	}
	return true
}

// AssetFileType returns the file extension an asset helper takes and a
// short description of the file, such as ".png" and "a PNG image"; ok is
// false when helper is not an asset helper.
func AssetFileType(helper string) (ext, what string, ok bool) {
	h := assetHelpers[helper]
	if h == nil {
		return "", "", false
	}
	return h.ext, h.what, true
}

// assetFileExts are file extensions that never belong in an id: the game
// adds them itself, so "item/ruby.png" names textures/item/ruby.png.png.
var assetFileExts = []string{".png.mcmeta", ".png", ".ogg", ".nbt", ".json", ".mcmeta"}

// stripAssetExt drops a file extension from the path of an asset id (a
// texture, sound, model and so on), warning with a fix that removes it
// from the source. e is the string or resource location written.
func (c *Checker) stripAssetExt(p, reg string, e syntax.Expr) string {
	for _, ext := range assetFileExts {
		bare, ok := strings.CutSuffix(p, ext)
		if !ok || bare == "" {
			continue
		}
		if c.trial {
			return bare
		}
		d := syntax.Diagnostic{
			Code: "mms/asset-extension", Severity: syntax.SeverityWarning, Sp: e.Span(),
			Message: fmt.Sprintf("%s ids leave out the file extension, which the game adds itself; MMS writes '%s'", strings.ReplaceAll(reg, "_", " "), bare),
		}
		end := -1
		switch x := e.(type) {
		case *syntax.String:
			if strings.HasSuffix(x.Raw, ext+x.Raw[:1]) {
				end = x.Sp.End - 1 // before the closing quote
			}
		case *syntax.ResLoc:
			if strings.HasSuffix(x.Text, ext) {
				end = x.Sp.End
			}
		}
		if end >= 0 {
			d.Fixes = []syntax.Fix{{Title: "Remove '" + ext + "'", Safe: true, Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: end - len(ext), End: end}}}}}
		}
		c.Diags = append(c.Diags, d)
		return bare
	}
	return p
}
