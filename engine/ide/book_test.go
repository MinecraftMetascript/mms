package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

const bookDecls = "Objective kills = \"dummy\"\nGlobal wave = 1\nFunction start {\n  /say go\n}\n"

// bookSession opens a project with main.mms and a book file, with | in
// the book text marking the cursor.
func bookSession(t *testing.T, book string) (*Session, int) {
	t.Helper()
	off := strings.Index(book, "|")
	book = book[:off] + book[off+1:]
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("main.mms", bookDecls+"Book b = { title = \"T\", author = \"A\", text = TextFile(\"b.mmsbook\") }\n")
	p.SetFile("b.mmsbook", book)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return &Session{P: p, R: r}, len([]rune(book[:off]))
}

func TestBookCompletion(t *testing.T) {
	cases := []struct {
		name, book string
		want       []string
		none       bool
	}{
		{"tag", "Hello {|", []string{"gold", "bold", "obfuscated", "key:", "sprite:", "/"}, false},
		{"second word", "{gold o|", []string{"obfuscated"}, false},
		{"key", "{key:j|", []string{"jump", "inventory"}, false},
		{"sprite", "{sprite:|", []string{"item/diamond"}, false},
		{"font", "{font:|", []string{"uniform", "alt"}, false},
		{"target", "[go](|", []string{"page:", "run:/", "hover:", "https://"}, false},
		{"page", "a\n---\nb [x](page:|", []string{"1", "2"}, false},
		{"command", "[go](run:/fun|", []string{"function"}, false},
		{"command id", "[go](run:/function |", []string{"test:start"}, false},
		{"expr", "${|", []string{"kills", "wave"}, false},
		{"plain text", "Hello th|ere", nil, true},
	}
	for _, tc := range cases {
		s, off := bookSession(t, tc.book)
		items, _ := s.Completion("b.mmsbook", off)
		got := labels(items)
		if tc.none && len(got) > 0 {
			t.Errorf("%s: want no completions, got %v", tc.name, got)
		}
		if !has(got, tc.want...) {
			t.Errorf("%s: completions %v lack %v", tc.name, got, tc.want)
		}
	}
	// The word being typed is what the item replaces.
	s, off := bookSession(t, "{key:ju|")
	_, rep := s.Completion("b.mmsbook", off)
	if rep.Start != 5 || rep.End != 7 {
		t.Errorf("replace = %+v", rep)
	}
	// '{' opens completion in book text, but not in MMS code.
	if !s.TriggerApplies("b.mmsbook", off, "{") {
		t.Error("{ does not trigger in a book file")
	}
}

func TestBookStringCompletion(t *testing.T) {
	src := "Book b = { title = \"T {\", author = \"A\", text = [\"Hello {gol|\"] }\nX = { a = 1 }\n"
	s, off := session(t, src)
	items, rep := s.Completion("main.mms", off)
	if !has(labels(items), "gold") {
		t.Errorf("completions = %v", labels(items))
	}
	if got := string([]rune(strings.ReplaceAll(src, "|", ""))[rep.Start:rep.End]); got != "gol" {
		t.Errorf("replaces %q", got)
	}
	// Not in the title.
	s, off = session(t, "Book b = { title = \"T {|\", author = \"A\" }\n")
	if items, _ := s.Completion("main.mms", off); len(items) > 0 {
		t.Errorf("title completions = %v", labels(items))
	}
	// Nor in MMS code after a brace.
	s, off = session(t, "LootTable l = {|}\n")
	if s.TriggerApplies("main.mms", off, "{") {
		t.Error("{ triggers in code")
	}
}

func TestBookHover(t *testing.T) {
	for _, tc := range []struct{ book, want string }{
		{"Press {key:ju|mp}", "key.jump"},
		{"{gold ob|fuscated}x{/}", "Scrambled"},
		{"[go](ru|n:/function test:start)", "Runs a command"},
		{"[go](page:1|)", "Turns to a page"},
	} {
		s, off := bookSession(t, tc.book)
		if h, _ := s.Hover("b.mmsbook", off); !strings.Contains(h, tc.want) {
			t.Errorf("%q: hover = %q, want %q", tc.book, h, tc.want)
		}
	}
}

func TestBookStringHover(t *testing.T) {
	s, off := session(t, "Book b = { title = \"T\", author = \"A\", text = [\"Press {key:inv|entory}\"] }\n")
	if h, _ := s.Hover("main.mms", off); !strings.Contains(h, "key.inventory") {
		t.Errorf("hover = %q", h)
	}
}

// A Book's text field offers TextFile(...), whose path lists the
// project's .mmsbook files.
func TestBookTextCompletion(t *testing.T) {
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("books/guide.mmsbook", "# Hi\n")
	p.SetFile("books/lore.mmsbook", "Old\n")
	for _, tc := range []struct {
		name, src string
		want      []string
	}{
		{"value", "Book b = {\n  text = |\n}\n", []string{"TextFile"}},
		{"path", "Book b = {\n  text = TextFile(\"|\")\n}\n", []string{"books/guide.mmsbook", "books/lore.mmsbook"}},
		{"unclosed path", "Book b = {\n  text = TextFile(\"bo|\n}\n", []string{"books/guide.mmsbook"}},
	} {
		off := strings.Index(tc.src, "|")
		src := tc.src[:off] + tc.src[off+1:]
		p.SetFile("main.mms", src)
		r, err := p.Analyze(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		s := &Session{P: p, R: r}
		items, _ := s.Completion("main.mms", len([]rune(src[:off])))
		if got := labels(items); !has(got, tc.want...) {
			t.Errorf("%s: completions %v lack %v", tc.name, got, tc.want)
		}
		if tc.name == "path" {
			if !s.TriggerApplies("main.mms", len([]rune(src[:off])), `"`) {
				t.Errorf("%s: \" does not open completion", tc.name)
			}
			for _, it := range items {
				if strings.HasSuffix(it.Label, ".mms") {
					t.Errorf("%s: offers %s", tc.name, it.Label)
				}
			}
		}
	}
}

// Markup completes in any string where a text component goes.
func TestTextStringCompletion(t *testing.T) {
	src := "Advancement a = { display = { title = \"Hi {gol|\", description = \"d\", icon = { id = stone }, background = \"minecraft:block/stone\" }, criteria = { x = Tick() } }\n"
	s, off := session(t, src)
	items, _ := s.Completion("main.mms", off)
	if !has(labels(items), "gold") {
		t.Errorf("completions = %v", labels(items))
	}
}

// Markup completes and hovers in a tell message.
func TestTellCompletion(t *testing.T) {
	src := "Function f {\n  tell(@a, \"Hi {gol|\")\n}\n"
	s, off := session(t, src)
	items, _ := s.Completion("main.mms", off)
	if !has(labels(items), "gold") {
		t.Errorf("completions = %v", labels(items))
	}
	s, off = session(t, "Function f {\n  actionbar(@s, \"Press {key:inv|entory}\")\n}\n")
	if h, _ := s.Hover("main.mms", off); !strings.Contains(h, "key.inventory") {
		t.Errorf("hover = %q", h)
	}
	s, off = session(t, "Function f {\n  te|\n}\n")
	if items, _ := s.Completion("main.mms", off); !has(labels(items), "tell") {
		t.Errorf("statement completions lack tell: %v", labels(items))
	}
}

func TestTellHover(t *testing.T) {
	s, off := session(t, "Function f {\n  te|ll(@a, \"hi\")\n}\n")
	if h, _ := s.Hover("main.mms", off); !strings.Contains(h, "tell(target: Selector, message: Text)") {
		t.Errorf("hover = %q", h)
	}
}

func TestTellSignature(t *testing.T) {
	s, off := session(t, "Function f {\n  title(@a, \"hi\", |)\n}\n")
	h, ok := s.Signature("main.mms", off)
	if !ok || h.Active != 2 || !strings.HasPrefix(h.Label, "title(") {
		t.Errorf("signature = %+v", h)
	}
}
