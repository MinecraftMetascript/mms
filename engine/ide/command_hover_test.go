package ide

import (
	"strings"
	"testing"
)

// TestCommandSignatureHover checks that a raw command's words show the
// command's signature, and that ids inside commands keep their own hover.
func TestCommandSignatureHover(t *testing.T) {
	main := "Function main {\n  /data get entity @s Pos\n  /function test:old_reward\n  let n = /data get entity @s Health\n}\n"
	s := docSession(t, main)
	cases := []struct {
		at   string
		skip int
		want string
	}{
		{"/data get entity @s Pos", 0, "```\n/data get (block|entity|storage)\n/data merge"},
		{"/data get entity @s Pos", 2, "/data modify (block|entity|storage)"},
		{"get entity @s Pos", 1, "```\n/data get block <target> [<path> [<scale>]]\n"},
		{"@s Pos", 0, "```\n/data get entity <target> [<path> [<scale>]]\n```\n`<target: selector>`"},
		{"/data get entity @s Health", 1, "/data get (block|entity|storage)"},
		{"/function", 2, "/function <name> [with block <source> [<path>]]\n"},
		{"test:old_reward", 6, "Function test:old_reward"},
	}
	for _, c := range cases {
		md, sp := s.Hover("main.mms", offOf(main, c.at, c.skip))
		if !strings.Contains(md, c.want) {
			t.Errorf("hover at %q+%d = %q, want %q", c.at, c.skip, md, c.want)
		}
		if sp.Len() == 0 {
			t.Errorf("hover at %q+%d has an empty span", c.at, c.skip)
		}
	}
	if md, _ := s.Hover("main.mms", offOf(main, "data", 0)); !strings.HasPrefix(md, "```") {
		t.Errorf("hover = %q", md)
	}
}
