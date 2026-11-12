package ide

import (
	"context"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

func assetSession(t *testing.T, src string) (*Session, int) {
	t.Helper()
	off := strings.Index(src, "|")
	src = src[:off] + src[off+1:]
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test", Output: "build"})
	f := func(s string) *fstest.MapFile { return &fstest.MapFile{Data: []byte(s)} }
	p.Files = fstest.MapFS{
		"art/sunset.png":         f("\x89PNG\r\n\x1a\n"),
		"art/Moon.PNG":           f("\x89PNG\r\n\x1a\n"),
		"music/tune.ogg":         f("OggS"),
		"builds/house.nbt":       f("\x1f\x8b"),
		"build/x.png":            f(""), // build output
		".git/y.png":             f(""),
		"node_modules/z/z.png":   f(""),
		"other/mms.json":         f("{}"), // another project
		"other/art/theirs.png":   f(""),
		"README.md":              f(""),
		"art/sub/deeper/far.png": f(""),
	}
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return &Session{P: p, R: r}, len([]rune(src[:off]))
}

// TestCommandSoundCompletion offers the sounds the project's Sound uses
// give their files after /playsound, before the vanilla ones.
func TestCommandSoundCompletion(t *testing.T) {
	src := "JukeboxSong s { sound_event = Sound(\"music/tune.ogg\"), description = \"x\", length_in_seconds = 1, comparator_output = 1 }\nFunction f {\n  /playsound |\n}"
	s, off := assetSession(t, src)
	items, _ := s.Completion("main.mms", off)
	got := labels(items)
	if len(got) == 0 || got[0] != "test:tune" {
		t.Fatalf("first items = %v, want test:tune first", got[:min(len(got), 3)])
	}
	if !slices.Contains(got, "ambient.cave") {
		t.Errorf("vanilla sounds missing: %v", got[:min(len(got), 5)])
	}
}

func TestAssetPathCompletion(t *testing.T) {
	pngs := []string{"art/Moon.PNG", "art/sub/deeper/far.png", "art/sunset.png"}
	cases := []struct {
		name, src string
		want      []string
		text      string // the text the item replaces
	}{
		{"empty string", `PaintingVariant p { width = 1, height = 1, asset_id = Texture("|") }`, pngs, ""},
		{"partial", `PaintingVariant p { width = 1, height = 1, asset_id = Texture("art/su|nset.png") }`, pngs, "art/sunset.png"},
		{"unterminated", "PaintingVariant p { width = 1, height = 1, asset_id = Texture(\"art/|\n", pngs, "art/"},
		{"sound", `JukeboxSong s { sound_event = Sound("|"), description = "x", length_in_seconds = 1, comparator_output = 1 }`, []string{"music/tune.ogg"}, ""},
		{"structure", `Function f { /say hi }` + "\nTemplatePool p { fallback = minecraft:empty, elements = [{ weight = 1, element = { element_type = minecraft:single_pool_element, location = StructureFile(\"|\"), projection = rigid, processors = minecraft:empty } }] }", []string{"builds/house.nbt"}, ""},
		{"second argument", `PaintingVariant p { width = 1, height = 1, asset_id = Texture("art/sunset.png", "|") }`, nil, ""},
		{"other call", `Advancement a { criteria = { x = { trigger = minecraft:impossible } }, display = { title = Text("|") } }`, nil, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s, off := assetSession(t, tc.src)
			items, sp := s.Completion("main.mms", off)
			got := labels(items)
			if tc.want == nil {
				for _, it := range items {
					if it.Kind == sema.CandFile {
						t.Fatalf("offered files: %v", got)
					}
				}
				return
			}
			if strings.Join(got, ",") != strings.Join(tc.want, ",") {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if items[0].Kind != sema.CandFile {
				t.Errorf("kind = %v", items[0].Kind)
			}
			src := s.R.Unit("main.mms").File.Src
			if txt := string(src[sp.Start:sp.End]); txt != tc.text {
				t.Errorf("replaces %q, want %q", txt, tc.text)
			}
		})
	}
}

func TestTriggerApplies(t *testing.T) {
	cases := []struct {
		src, ch string
		want    bool
	}{
		{`PaintingVariant p { width = 1, height = 1, asset_id = Texture("|") }`, `"`, true},
		{"PaintingVariant p { width = 1, height = 1, asset_id = Texture(\"art/|\n", "/", true},
		{`Advancement a { display = { title = "|" } }`, `"`, false},
		{"Function f {\n  /|\n}", "/", true},
		{"Function f {\n  let n = /|\n}", "/", true},
		{"Function f {\n  let a = 4\n  let n = a /|\n}", "/", false},
		{"Function f {\n  let b = a == /|\n}", "/", false},
		{"Global g = 4 /| 2", "/", false},
		{"Global g = 4 .| 2", ".", true},
		{"Blueprint B {\n  emit {\n    Objective ${|}: dummy\n  }\n}", "{", true},
		{"Blueprint B {\n  let k = \"a${|}\"\n}", "{", true},
		{"Blueprint B {\n  let k = {|}\n}", "{", false},
	}
	for _, tc := range cases {
		s, off := assetSession(t, tc.src)
		if got := s.TriggerApplies("main.mms", off, tc.ch); got != tc.want {
			t.Errorf("%q: got %v", tc.src, got)
		}
	}
}

// TestTextureMetaCompletion offers the sections of a .png.mcmeta file as
// Texture's named arguments, and their fields inside.
func TestTextureMetaCompletion(t *testing.T) {
	cases := []struct{ src, want string }{
		{`Model m { parent = block/cube_all, textures = { all = Texture("art/sunset.png", |) } }`, "animation"},
		{`Model m { parent = block/cube_all, textures = { all = Texture("art/sunset.png", g|) } }`, "gui"},
		{`Model m { parent = block/cube_all, textures = { all = Texture("art/sunset.png", animation = { | }) } }`, "frametime"},
		{`Model m { parent = block/cube_all, textures = { all = Texture("art/sunset.png", animation = { frametime = 2 }, |) } }`, "texture"},
	}
	for _, tc := range cases {
		s, off := assetSession(t, tc.src)
		items, _ := s.Completion("main.mms", off)
		if got := labels(items); !slices.Contains(got, tc.want) {
			t.Errorf("%s: got %v, want %s", tc.src, got, tc.want)
		}
	}
}

// TestTextureSlotCompletion offers the slots of a model's parent chain
// in its textures body, those its faces show first.
func TestTextureSlotCompletion(t *testing.T) {
	s, off := assetSession(t, `Model m { parent = block/cube_column, textures = { side = block/oak_log, | } }`)
	items, _ := s.Completion("main.mms", off)
	got := labels(items)
	if len(got) < 2 || got[0] != "end" || slices.Contains(got, "side") || !slices.Contains(got, "particle") {
		t.Errorf("got %v", got)
	}
}

func TestTextureHover(t *testing.T) {
	src := `Model m { parent = block/cube_all, textures = { all = Tex|ture("art/lamp.png", animation = { frametime = 3 }) } }`
	off := strings.Index(src, "|")
	src = src[:off] + src[off+1:]
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.Files = fstest.MapFS{"art/lamp.png": {Data: append([]byte("\x89PNG\r\n\x1a\n\x00\x00\x00\x0dIHDR"), 0, 0, 0, 16, 0, 0, 0, 64)}}
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	s := &Session{P: p, R: r}
	md, _ := s.Hover("main.mms", off)
	if !strings.Contains(md, "16×64 image") || !strings.Contains(md, "4 frames of 16×16, one loop every 12 ticks (0.6s)") {
		t.Errorf("hover: %s", md)
	}
}

func TestRetextureSlotCompletion(t *testing.T) {
	for src, want := range map[string]string{
		`Retexture minecraft:oak_log { textures = { | } }`:                                 "end side",
		`Retexture minecraft:oak_log { textures = { side = Texture("art/lamp.png"), | } }`: "end",
		`Retexture minecraft:oak_stairs { textures = { | } }`:                              "bottom side top",
	} {
		s, off := assetSession(t, src)
		items, _ := s.Completion("main.mms", off)
		got := labels(items)
		if n := len(strings.Fields(want)); len(got) < n || strings.Join(got[:n], " ") != want || !slices.Contains(got, "particle") {
			t.Errorf("%s: got %v", src, got)
		}
	}
}
