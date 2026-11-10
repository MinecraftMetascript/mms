package ide

import (
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
)

// Hovering any schema type's name in a type says what it is: its shape,
// not only that it comes from the game's schema.
func TestSchemaTypeHoverMatrix(t *testing.T) {
	var names []string
	for name := range spec.MustDefault().ShortNames() {
		if len(name) > 1 && !strings.HasPrefix(name, "<") {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	var sb strings.Builder
	sb.WriteString("Type All = {\n")
	offs := map[string]int{}
	for i, n := range names {
		sb.WriteString("  f")
		sb.WriteString(strings.Repeat("x", i%3))
		sb.WriteString(itoa(i) + ": ")
		offs[n] = len([]rune(sb.String())) + 1
		sb.WriteString(n + "\n")
	}
	sb.WriteString("}\n")
	s, _ := session(t, sb.String())
	var bare []string
	checked := 0
	for _, n := range names {
		md, _ := s.Hover("main.mms", offs[n])
		if !strings.Contains(md, "/reference/types/") {
			continue // a resource kind, described by its own page
		}
		checked++
		if strings.Contains(md, "A type from the game's data schema.") || strings.TrimSpace(strings.TrimPrefix(md, "```mms\n"+n+"\n```")) == "" {
			bare = append(bare, n)
		}
	}
	if checked < 1000 {
		t.Errorf("only %d of %d schema type names hovered with a reference link", checked, len(names))
	}
	if len(bare) > 0 {
		t.Errorf("%d schema types hover without a shape: %s", len(bare), strings.Join(bare, ", "))
	}
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	var d []byte
	for ; i > 0; i /= 10 {
		d = append([]byte{byte('0' + i%10)}, d...)
	}
	return string(d)
}

// In a typed const, a field's name hovers with its type, and where its
// value goes completion offers each way of writing one.
func TestTypedConstFieldDiscoverable(t *testing.T) {
	s, off := session(t, "const c: ClimateParameters = {\n  tempe|rature = 1\n}\n")
	if md, _ := s.Hover("main.mms", off); !strings.Contains(md, "temperature: ClimateParameter") || !strings.Contains(md, "`0` or `[-2, 2]`") {
		t.Errorf("hover of a field name = %q", md)
	}
	s, off = session(t, "const c: ClimateParameters = {\n  temperature = |\n}\n")
	items, _ := s.Completion("main.mms", off)
	if got := labels(items); !has(got, "0", "[-2, 2]") {
		t.Errorf("completion of a ClimateParameter = %v", got)
	}
	s, off = session(t, "const s: BiomeSource = MultiNoise(|)\n")
	if h, ok := s.Signature("main.mms", off); !ok || !strings.Contains(h.Label, "biomes?:") {
		t.Errorf("MultiNoise signature = %q, want biomes listed", h.Label)
	}
}
