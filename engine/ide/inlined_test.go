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
