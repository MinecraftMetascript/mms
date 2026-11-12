package sema

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/fs"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Asset files: Texture("art/sunset.png"), Sound("music/tune.ogg") and
// StructureFile("builds/house.nbt") name a file anywhere in the project. The
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
	Helper string   // "Texture", "Sound" or "StructureFile"
	Src    string   // project-relative source file
	ID     string   // ns:name, the value written to the JSON
	Out    string   // pack-relative output path
	Pack   PackKind // which pack Out belongs to
	// Meta is a texture's .png.mcmeta file, from the call's named
	// arguments, such as animation = { frametime = 2 }; nil without.
	Meta *jsonv.Object
	// Width and Height are a texture's size in pixels, when known.
	Width, Height int
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
	"StructureFile": {
		registries: []string{"structure"}, ext: ".nbt", magic: []byte{0x1f, 0x8b}, what: "a gzipped structure (.nbt)", pack: DataPack,
		out: func(ns, _, name string) string { return "data/" + ns + "/structure/" + name + ".nbt" },
	},
}

// AssetHelperNames lists the helpers, for completion and docs.
var AssetHelperNames = []string{"Sound", "StructureFile", "Texture"}

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

// tTextureMeta is the type of a .png.mcmeta file. Texture(...) takes
// its sections (animation, texture, gui, villager) as named arguments.
const tTextureMeta = "::java::assets::texture_meta::TextureMeta"

// pngHeader is the length of a PNG's signature and IHDR chunk, which
// hold its size.
const pngHeader = 24

// assetCall checks Texture(path, name?, section = ...), Sound(...) or
// StructureFile(...) at an id position and returns the id it stands for.
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
	var nameArg syntax.Expr
	var meta []*syntax.Arg
	pos := 0
	for _, a := range call.Args {
		if a.Name != nil {
			if name != "Texture" {
				c.errorf("mms/type-mismatch", a.Sp, "%s takes a file path and an optional id name, without argument names", name)
				return nil
			}
			meta = append(meta, a)
			continue
		}
		s, ok := a.Value.(*syntax.String)
		switch {
		case pos > 1:
			c.errorf("mms/type-mismatch", a.Sp, "%s takes at most two arguments: a file path and an id name", name)
			return nil
		case !ok:
			c.errorf("mms/type-mismatch", a.Value.Span(), "expected a quoted %s", map[bool]string{true: "file path", false: "id name"}[pos == 0])
			return nil
		case pos == 0:
			src, srcSp = s.Value, s.Sp
		default:
			idName, nameSp, nameArg = s.Value, s.Sp, s
		}
		pos++
	}
	if pos == 0 {
		c.errorf("mms/type-mismatch", call.Sp, "%s needs the path of a %s file, relative to the project: %s(\"file%s\")", name, h.ext, name, h.ext)
		return nil
	}
	in := c.info(e)
	in.Type = v
	var metaV *jsonv.Object
	if len(meta) > 0 {
		// The named arguments are the fields of the .png.mcmeta file.
		body := &syntax.Block{Closed: true, Sp: call.Sp}
		for _, a := range meta {
			body.Items = append(body.Items, &syntax.FieldStmt{Key: a.Name, Value: a.Value, Sp: a.Sp})
		}
		ref := &spec.Type{Kind: "reference", Path: tTextureMeta}
		metaV, _ = c.check(&syntax.Object{Body: body, Sp: call.Sp}, tview{T: ref, Name: tTextureMeta}).(*jsonv.Object)
	}
	src = path.Clean(strings.ReplaceAll(src, "\\", "/"))
	if path.IsAbs(src) || src == ".." || strings.HasPrefix(src, "../") {
		c.errorf("mms/asset", srcSp, "asset paths are relative to the project folder and stay inside it")
		return nil
	}
	if !strings.EqualFold(path.Ext(src), h.ext) {
		c.errorf("mms/asset", srcSp, "%s expects %s file, not '%s'", name, map[string]string{".png": "a .png", ".ogg": "an .ogg", ".nbt": "an .nbt"}[h.ext], path.Base(src))
		return nil
	}
	model := name == "Texture" && c.resource == "model"
	ns := c.NS
	if idName == "" {
		idName = assetName(src)
		if model {
			// Models only find sprites the blocks and items atlases
			// stitch, and the blocks atlas takes every file under
			// textures/block/. Its sprites work in block and item models
			// alike, so a model's own texture goes there.
			idName = "block/" + idName
		}
	} else {
		idName = c.stripAssetExt(idName, h.registries[0], nameArg)
		// An id name may name another namespace, such as
		// "minecraft:block/stone", to replace that texture.
		if n, p, ok := strings.Cut(idName, ":"); ok && name == "Texture" {
			if !assetNamePattern.MatchString(n) || strings.Contains(n, "/") || n == "." || n == ".." {
				c.errorf("mms/invalid-name", nameSp, "a namespace is lowercase letters, digits, '_', '-' and '.'")
				return nil
			}
			ns, idName = n, p
		}
		if !assetNamePattern.MatchString(idName) || strings.Contains(idName, "..") {
			c.errorf("mms/invalid-name", nameSp, "an id name is lowercase letters, digits, '_', '-', '.' and '/'%s", map[bool]string{true: ", after an optional namespace such as minecraft:", false: ""}[name == "Texture"])
			return nil
		}
	}
	n := len(h.magic)
	if name == "Texture" {
		n = pngHeader
	}
	head, ok := c.assetHead(src, n)
	if !ok && c.Files == nil {
		c.errorf("mms/asset", srcSp, "asset files are not available here yet; build with the mms command line to include '%s'", src)
		return nil
	}
	if !ok {
		c.errorf("mms/asset", srcSp, "no file '%s' in the project", src)
		return nil
	}
	if !bytes.HasPrefix(head, h.magic) {
		c.errorf("mms/asset", srcSp, "'%s' is not %s", src, h.what)
		return nil
	}
	full := ns + ":" + idName
	in.ID = full
	in.ModelTexture = model
	in.Asset = &Asset{Helper: name, Src: src, ID: full, Out: h.out(ns, id.Path, idName), Pack: h.pack}
	if name == "Texture" {
		in.Asset.Width, in.Asset.Height = pngSize(head)
		if metaV != nil {
			in.Asset.Meta = metaV
		}
		c.textureChecks(call, in.Asset, meta, model)
	}
	return jsonv.String(full)
}

// pngSize reads a PNG's width and height from its IHDR chunk, or 0, 0.
func pngSize(head []byte) (w, h int) {
	if len(head) < pngHeader || string(head[12:16]) != "IHDR" {
		return 0, 0
	}
	return int(binary.BigEndian.Uint32(head[16:20])), int(binary.BigEndian.Uint32(head[20:24]))
}

// FrameSize returns the size of one frame of an animated texture, as the
// game works it out: the animation's width and height when given, the
// image's own otherwise, and a square of the shorter side when neither
// is. ok is false when the texture has no animation or no known size.
func (a *Asset) FrameSize() (fw, fh int, ok bool) {
	anim, _ := a.metaSection("animation")
	if anim == nil || a.Width == 0 || a.Height == 0 {
		return 0, 0, false
	}
	w, wok := intField(anim, "width")
	h, hok := intField(anim, "height")
	switch {
	case wok && hok:
		return w, h, true
	case wok:
		return w, a.Height, true
	case hok:
		return a.Width, h, true
	}
	m := min(a.Width, a.Height)
	return m, m, true
}

// Frames returns the number of frames in an animated texture's image,
// read left to right and top to bottom, or 0.
func (a *Asset) Frames() int {
	fw, fh, ok := a.FrameSize()
	if !ok || fw <= 0 || fh <= 0 || a.Width%fw != 0 || a.Height%fh != 0 {
		return 0
	}
	return (a.Width / fw) * (a.Height / fh)
}

// AnimationTicks returns how long one loop of an animated texture lasts,
// in game ticks, or 0.
func (a *Asset) AnimationTicks() int {
	anim, _ := a.metaSection("animation")
	n := a.Frames()
	if anim == nil || n == 0 {
		return 0
	}
	ft, ok := intField(anim, "frametime")
	if !ok {
		ft = 1
	}
	frames, _ := anim.Get("frames").(jsonv.Array)
	if frames == nil {
		return n * ft
	}
	total := 0
	for _, f := range frames {
		t := ft
		if o, ok := f.(*jsonv.Object); ok {
			if v, ok := intField(o, "time"); ok {
				t = v
			}
		}
		total += t
	}
	return total
}

func (a *Asset) metaSection(key string) (*jsonv.Object, bool) {
	if a.Meta == nil {
		return nil, false
	}
	o, ok := a.Meta.Get(key).(*jsonv.Object)
	return o, ok
}

func intField(o *jsonv.Object, key string) (int, bool) {
	n, ok := o.Get(key).(jsonv.Number)
	if !ok {
		return 0, false
	}
	v, err := strconv.Atoi(string(n))
	return v, err == nil
}

// textureChecks compares a texture's animation with its image: a strip
// that is not a whole number of frames, a frame index past the last
// frame, a model texture that looks like an animation strip without one,
// and a .png.mcmeta file beside the image, which MMS does not copy.
func (c *Checker) textureChecks(call *syntax.Call, a *Asset, meta []*syntax.Arg, model bool) {
	if c.trial {
		return
	}
	var animArg *syntax.Arg
	for _, m := range meta {
		if m.Name.Text == "animation" {
			animArg = m
		}
	}
	if fw, fh, ok := a.FrameSize(); ok && animArg != nil {
		if fw <= 0 || fh <= 0 || a.Width%fw != 0 || a.Height%fh != 0 {
			c.warnf("mms/texture-animation", animArg.Sp, "'%s' is %dx%d, which is not a whole number of %dx%d frames, so the game shows the missing texture", a.Src, a.Width, a.Height, fw, fh)
		} else if n := a.Frames(); n > 0 {
			c.frameIndexChecks(animArg.Value, n, a.Src)
		}
	}
	side := a.Src + ".mcmeta"
	raw, hasSide := c.assetHead(side, 1<<16)
	if animArg == nil && !hasSide && model && a.Width > 0 && a.Height >= 2*a.Width && a.Height%a.Width == 0 {
		d := syntax.Diagnostic{
			Code: "mms/texture-animation", Severity: syntax.SeverityWarning, Sp: call.Sp,
			Message: fmt.Sprintf("'%s' is %dx%d, a strip of %d square frames, but it has no animation, so the game squashes the whole strip into the model; add animation = { frametime = ... }", a.Src, a.Width, a.Height, a.Height/a.Width),
		}
		if at, ok := callEnd(call); ok {
			d.Fixes = []syntax.Fix{{Title: "Animate it", Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: at, End: at}, Text: ", animation = { frametime = 2 }"}}}}
		}
		c.Diags = append(c.Diags, d)
	}
	if !hasSide {
		return
	}
	d := syntax.Diagnostic{
		Code: "mms/texture-mcmeta", Severity: syntax.SeverityWarning, Sp: call.Sp,
		Message: fmt.Sprintf("MMS does not copy '%s'; write its sections in the call instead, such as Texture(\"%s\", animation = { ... })", side, a.Src),
	}
	if v, err := jsonv.Parse(raw); err == nil && len(meta) == 0 {
		if o, ok := v.(*jsonv.Object); ok && len(o.Members) > 0 {
			if at, ok := callEnd(call); ok {
				var sb strings.Builder
				for _, m := range o.Members {
					sb.WriteString(", ")
					sb.WriteString(JSONKey(m.Key))
					sb.WriteString(" = ")
					sb.WriteString(JSONExpr(m.Value))
				}
				d.Fixes = []syntax.Fix{{Title: "Move '" + path.Base(side) + "' into the call", Edits: []syntax.TextEdit{{Sp: syntax.Span{Start: at, End: at}, Text: sb.String()}}}}
			}
		}
	}
	c.Diags = append(c.Diags, d)
}

// frameIndexChecks warns about frames entries past the last frame: the
// game then leaves the texture still.
func (c *Checker) frameIndexChecks(anim syntax.Expr, n int, src string) {
	o, ok := anim.(*syntax.Object)
	if !ok {
		return
	}
	for _, it := range o.Body.Items {
		f, ok := it.(*syntax.FieldStmt)
		if !ok || f.Key.Text != "frames" {
			continue
		}
		l, ok := f.Value.(*syntax.List)
		if !ok {
			continue
		}
		for _, el := range l.Elems {
			idx := el
			if eo, ok := el.(*syntax.Object); ok {
				idx = nil
				for _, it := range eo.Body.Items {
					if ff, ok := it.(*syntax.FieldStmt); ok && ff.Key.Text == "index" {
						idx = ff.Value
					}
				}
			}
			num, ok := idx.(*syntax.Number)
			if !ok {
				continue
			}
			if i, err := strconv.Atoi(num.Text); err == nil && i >= n {
				c.warnf("mms/texture-animation", num.Sp, "'%s' has %d frames (0 to %d), so frame %d is not there and the game does not animate the texture", src, n, n-1, i)
			}
		}
	}
}

// callEnd returns the offset of a call's closing parenthesis.
func callEnd(call *syntax.Call) (int, bool) {
	if call.Sp.End <= call.Sp.Start {
		return 0, false
	}
	return call.Sp.End - 1, true
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

// fileExt returns the file extension p ends in, from assetFileExts or
// .mcfunction, or "".
func fileExt(p string) string {
	for _, ext := range append(assetFileExts, ".mcfunction") {
		if len(p) > len(ext) && strings.HasSuffix(p, ext) {
			return ext
		}
	}
	return ""
}

// stripAssetExt drops a file extension from the path of an asset id (a
// texture, sound, model and so on), warning with a fix that removes it
// from the source. e is the string or resource location written.
func (c *Checker) stripAssetExt(p, reg string, e syntax.Expr) string {
	if ext := fileExt(p); ext != "" && ext != ".mcfunction" {
		bare := strings.TrimSuffix(p, ext)
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
