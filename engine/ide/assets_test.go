package ide

import (
	"context"
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

func TestAssetPathCompletion(t *testing.T) {
	pngs := []string{"art/Moon.PNG", "art/sub/deeper/far.png", "art/sunset.png"}
	cases := []struct {
		name, src string
		want      []string
		text      string // the text the item replaces
	}{
		{"empty string", `PaintingVariant p = { width = 1, height = 1, asset_id = Texture("|") }`, pngs, ""},
		{"partial", `PaintingVariant p = { width = 1, height = 1, asset_id = Texture("art/su|nset.png") }`, pngs, "art/sunset.png"},
		{"unterminated", "PaintingVariant p = { width = 1, height = 1, asset_id = Texture(\"art/|\n", pngs, "art/"},
		{"sound", `JukeboxSong s = { sound_event = Sound("|"), description = "x", length_in_seconds = 1, comparator_output = 1 }`, []string{"music/tune.ogg"}, ""},
		{"structure", `Function f { /say hi }` + "\nTemplatePool p = { fallback = minecraft:empty, elements = [{ weight = 1, element = { element_type = minecraft:single_pool_element, location = Structure(\"|\"), projection = rigid, processors = minecraft:empty } }] }", []string{"builds/house.nbt"}, ""},
		{"second argument", `PaintingVariant p = { width = 1, height = 1, asset_id = Texture("art/sunset.png", "|") }`, nil, ""},
		{"other call", `Advancement a = { criteria = { x = { trigger = minecraft:impossible } }, display = { title = Text("|") } }`, nil, ""},
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
		{`PaintingVariant p = { width = 1, height = 1, asset_id = Texture("|") }`, `"`, true},
		{"PaintingVariant p = { width = 1, height = 1, asset_id = Texture(\"art/|\n", "/", true},
		{`Advancement a = { display = { title = "|" } }`, `"`, false},
		{"Function f {\n  /|\n}", "/", true},
		{"Global g = 4 /| 2", "/", false},
		{"Global g = 4 .| 2", ".", true},
	}
	for _, tc := range cases {
		s, off := assetSession(t, tc.src)
		if got := s.TriggerApplies("main.mms", off, tc.ch); got != tc.want {
			t.Errorf("%q: got %v", tc.src, got)
		}
	}
}
