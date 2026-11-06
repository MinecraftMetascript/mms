package ide

import (
	"strings"
	"testing"
)

const inlinedSrc = `Template ring(p: Id<particle>, points: int = 4): Commands {
  for i in 0..points - 1 {
    /particle ${p} ^${i} ^1 ^0
  }
}
Function f {
  ring(flame, 2)
}
`

func TestInlinedTemplate(t *testing.T) {
	s, _ := session(t, inlinedSrc+"|")
	var found bool
	for _, h := range s.InlayHints("main.mms") {
		if h.Label != "inlined" {
			continue
		}
		found = true
		if want := strings.Index(inlinedSrc, "ring(flame, 2)") + len("ring(flame, 2)"); h.Off != want {
			t.Errorf("inlay at %d, want %d", h.Off, want)
		}
		if !strings.Contains(h.Tooltip, "/particle flame ^0 ^1 ^0\n/particle flame ^1 ^1 ^0") {
			t.Errorf("tooltip: %q", h.Tooltip)
		}
	}
	if !found {
		t.Fatalf("no inlined hint: %+v", s.InlayHints("main.mms"))
	}
	md, _ := s.Hover("main.mms", strings.Index(inlinedSrc, "ring(flame")+1)
	if !strings.Contains(md, "```mms\n/particle flame ^0 ^1 ^0\n/particle flame ^1 ^1 ^0\n```") {
		t.Errorf("hover: %s", md)
	}
}

func TestTypeHints(t *testing.T) {
	src := "const max_wave = 10\nconst boss = @e[tag=boss, limit=1]\nconst mobs = [zombie, skeleton]\nconst names = [\"a\", \"b\"]\nconst typed: int = 3\nFunction f {\n  let n = 0\n  let s = \"x\"\n  let ok = n > 1\n  let t: int = 2\n}\n"
	s, _ := session(t, src+"|")
	got := map[string]string{}
	for _, h := range s.InlayHints("main.mms") {
		line := src[:h.Off]
		line = line[strings.LastIndexByte(line, '\n')+1:]
		got[strings.TrimSpace(line)] = h.Label
	}
	for k, want := range map[string]string{
		"const max_wave": ": int", "const boss": ": Selector", "const names": ": [string]",
		"let n": ": int", "let s": ": string", "let ok": ": bool",
	} {
		if got[k] != want {
			t.Errorf("%s: hint %q, want %q (all %v)", k, got[k], want, got)
		}
	}
	for _, k := range []string{"const typed: int", "let t: int", "const mobs"} {
		if _, ok := got[k]; ok {
			t.Errorf("%s has a hint", k)
		}
	}
}

func TestExpand(t *testing.T) {
	src := inlinedSrc + "Template twice(n: int): int = n * 2\nDensityFunction d = twice(3)\nCustomItem ruby { base = diamond, name = \"Ruby\" }\n"
	s, _ := session(t, src+"|")
	x, ok := s.Expand("main.mms", strings.Index(src, "ring(flame"))
	if !ok || x.Title != "Template ring" || x.Text != "/particle flame ^0 ^1 ^0\n/particle flame ^1 ^1 ^0" {
		t.Errorf("ring: %+v", x)
	}
	x, ok = s.Expand("main.mms", strings.Index(src, "twice(3)"))
	if !ok || x.Text != "6" {
		t.Errorf("twice: %+v", x)
	}
	x, ok = s.Expand("main.mms", strings.Index(src, "ruby"))
	if !ok || x.Title != "CustomItem ruby" || !strings.Contains(x.Text, "ruby") {
		t.Errorf("ruby: %+v", x)
	}
	t.Log(x.Text)
	if _, ok := s.Expand("main.mms", strings.Index(src, "Function f")); ok {
		t.Error("expansion on a plain declaration")
	}
}
