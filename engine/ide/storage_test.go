package ide

import (
	"strings"
	"testing"
)

const storageDecls = "/** The siege. */\nType State {\n  /** The wave number. */\n  wave: int = 0\n  boss: bool\n}\n/** Kept between reloads. */\nStorage game: State\n"

func TestStorageCompletion(t *testing.T) {
	s, off := session(t, storageDecls+"Function f {\n  game.w|\n}\n")
	items, _ := s.Completion("main.mms", off)
	if !has(labels(items), "wave") || !has(labels(items), "boss") {
		t.Errorf("field completions = %v", labels(items))
	}
	s, off = session(t, storageDecls+"Function f {\n  ga|\n}\n")
	if items, _ := s.Completion("main.mms", off); !has(labels(items), "game") {
		t.Errorf("statement completions lack game: %v", labels(items))
	}
	s, off = session(t, storageDecls+"Function f {\n  /say ${ga|}\n}\n")
	if items, _ := s.Completion("main.mms", off); !has(labels(items), "game") {
		t.Errorf("interpolation completions lack game: %v", labels(items))
	}
}

func TestStorageHover(t *testing.T) {
	s, off := session(t, storageDecls+"Function f {\n  ga|me.wave += 1\n}\n")
	if h, _ := s.Hover("main.mms", off); !strings.Contains(h, "Kept between reloads") {
		t.Errorf("hover on the storage = %q", h)
	}
	s, off = session(t, storageDecls+"Function f {\n  game.wa|ve += 1\n}\n")
	if h, _ := s.Hover("main.mms", off); !strings.Contains(h, "The wave number") || !strings.Contains(h, "int") {
		t.Errorf("hover on the field = %q", h)
	}
}
