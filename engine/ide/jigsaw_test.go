package ide

import (
	"context"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/minecraftmetascript/mms/engine/nbt"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

func TestStructureHoverListsJigsaws(t *testing.T) {
	data, err := nbt.Encode(map[string]any{
		"size":    []any{int32(5), int32(3), int32(5)},
		"palette": []any{map[string]any{"Name": "minecraft:jigsaw", "Properties": map[string]any{"orientation": "south_up"}}},
		"blocks": []any{map[string]any{"pos": []any{int32(2), int32(0), int32(4)}, "state": int32(0),
			"nbt": map[string]any{"id": "minecraft:jigsaw", "pool": "test:tops", "name": "test:bottom", "target": "test:top"}}},
	})
	if err != nil {
		t.Fatal(err)
	}
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test", Output: "build"})
	p.Files = fstest.MapFS{"s/base.nbt": {Data: data}}
	src := `TemplatePool p = { fallback = minecraft:empty, elements = [{ weight = 1, element = { element_type = minecraft:single_pool_element, location = Structure("s/base.nbt"), projection = rigid, processors = minecraft:empty } }] }`
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	s := &Session{P: p, R: r}
	md, _ := s.Hover("main.mms", strings.Index(src, "Structure")+2)
	if !strings.Contains(md, "size 5×3×5") || !strings.Contains(md, "| 2 0 4 | south up | `test:bottom` | `test:tops` | `test:top` |") {
		t.Fatalf("hover:\n%s", md)
	}
}
