package ide

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// jigsawLabels completes at the ^ inserted after the first `after` in
// the jigsaw example.
func jigsawLabels(t *testing.T, after, insert string) []string {
	t.Helper()
	b, err := os.ReadFile("../../examples/jigsaw/dungeon.mms")
	if err != nil {
		t.Fatal(err)
	}
	src := string(b)
	i := strings.Index(src, after)
	if i < 0 {
		t.Fatalf("no %q in the example", after)
	}
	i += len(after)
	src = src[:i] + insert + src[i:]
	at := strings.Index(src, "^")
	src = src[:at] + src[at+1:]
	p := project.New(spec.MustDefault(), project.Config{Namespace: "dungeon"})
	p.Files = os.DirFS("../../examples/jigsaw")
	p.SetFile("dungeon.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	items, _ := (&Session{P: p, R: r}).Completion("dungeon.mms", len([]rune(src[:at])))
	return labels(items)
}

// Prelude Templates and Blueprints complete like built-in helpers: Piece
// where a pool entry goes, its structure files as its path, a
// JigsawStructure's own pools as pool ids.
func TestPreludeCompletion(t *testing.T) {
	has := func(what string, got []string, want ...string) {
		t.Helper()
		all := " " + strings.Join(got, " ") + " "
		for _, w := range want {
			if !strings.Contains(all, " "+w+" ") {
				t.Errorf("%s lacks %s: %s", what, w, strings.TrimSpace(all))
			}
		}
	}
	lacks := func(what string, got []string, unwanted ...string) {
		t.Helper()
		all := " " + strings.Join(got, " ") + " "
		for _, w := range unwanted {
			if strings.Contains(all, " "+w+" ") {
				t.Errorf("%s offers %s: %s", what, w, strings.TrimSpace(all))
			}
		}
	}
	has("a pool entry", jigsawLabels(t, "    start = [", "^"), "Piece", "WeightedElement")
	has("a pool entry being typed", jigsawLabels(t, "        Piece(\"halls/hall\", 3),\n", "        Pi^\n"), "Piece")
	has("a Piece path", jigsawLabels(t, "    start = [Piece(\"", "^"), "start", "halls/hall", "rooms/vault")
	got := jigsawLabels(t, "    start = [Piece(", "^")
	has("Piece's first argument", got, `"halls/hall"`)
	lacks("Piece's first argument", got, "Piece", "cos", "book_pages")
	has("start", jigsawLabels(t, "  start = ", "^"), "start", "halls", "rooms", "caps")
	has("fallback", jigsawLabels(t, "      fallback = ", "^"), "caps")
	// A line being typed, with the lines below it still there.
	has("an unclosed pool list", jigsawLabels(t, "    rooms = { fallback = caps, elements = [", "Pi^\n"), "Piece")
}

// A word at the start of a statement before an assignment, and the
// target of an assignment, complete from the names in scope; Templates
// that give a value complete in code, and Commands Templates as
// statements.
func TestCodeCompletion(t *testing.T) {
	for _, tc := range []struct{ src, want string }{
		{"Global barrel_id: int = 0\nFunction f {\n  let id = 0\n  id = b^\n  id = 2\n}\n", "barrel_id"},
		{"Global next_id: int = 0\nFunction f {\n  /say hi\n  n^ += 1\n}\n", "next_id"},
		{"Template fib(n: int): int = n\nTest t {\n  expect fi^(0) == 0\n}\n", "fib"},
		{"Template ring(x: string): Commands = run {\n  /say ${x}\n}\nFunction f {\n  ri^\n}\n", "ring"},
		{"LootTable t {\n  pools = [Pool(1, [Item(na^\n  ]\n}\n", "name"},
		{"CustomItem ruby { base = diamond }\nRecipe r = Shapeless([ru^], diamond)\n", "ruby"},
	} {
		got := labelsAt(t, tc.src)
		if !strings.Contains(" "+strings.Join(got, " ")+" ", " "+tc.want+" ") {
			t.Errorf("%q lacks %s: %v", tc.src, tc.want, got)
		}
	}
	got := strings.Join(labelsAt(t, "Template fib(n: int): int = n\nFunction f {\n  let x = fi^\n}\n"), " ")
	if strings.Contains(" "+got+" ", " Piece ") || strings.Contains(" "+got+" ", " cos ") {
		t.Errorf("code offers Templates it cannot use: %s", got)
	}
}
