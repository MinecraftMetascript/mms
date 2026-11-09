package command

import (
	"strings"
	"testing"
)

func TestUsage(t *testing.T) {
	c := checker(t)
	for _, q := range []struct{ text, at, arg, want string }{
		{"data get entity @s Pos", "data", "", "data get (block|entity|storage)\ndata merge (block|entity|storage)\ndata modify (block|entity|storage)\ndata remove (block|entity|storage)"},
		{"data get entity @s Pos", "get", "", "data get block <target> [<path> [<scale>]]\ndata get entity <target> [<path> [<scale>]]\ndata get storage <target> [<path> [<scale>]]"},
		{"data get entity @s Pos", "@s", "<target: selector>", "data get entity <target> [<path> [<scale>]]"},
		{"execute as @a at @s run say hi", "at", "", "execute as <targets> at <targets> -> execute"},
		{"execute as @a run say hi", "say", "", "execute as <targets> run say <message>"},
		{"execute as @a run say hi", "run", "", "execute as <targets> run ..."},
		{"tp @s ~ ~ ~", "tp", "", "tp -> teleport\ntp <destination>\ntp <location>\ntp <targets> <destination>\ntp <targets> <location> [facing entity <facingEntity> [<facingAnchor>]]\ntp <targets> <location> [facing <facingLocation>]\ntp <targets> <location> [<rotation>]"},
		{"give @s diamond", "give", "", "give <targets> <item> [<count>]"},
		{"data bogus", "bogus", "", "data get (block|entity|storage)\ndata merge (block|entity|storage)\ndata modify (block|entity|storage)\ndata remove (block|entity|storage)"},
	} {
		off := len([]rune(q.text[:strings.Index(q.text, q.at)]))
		u, ok := c.Usage(q.text, off)
		if got := strings.Join(u.Lines, "\n"); !ok || got != q.want || u.Arg != q.arg {
			t.Errorf("Usage(%q) at %q = %v %q\n%s\nwant %q\n%s", q.text, q.at, ok, u.Arg, got, q.arg, q.want)
		}
	}
	for _, text := range []string{"bogus", "$say $(x)"} {
		if u, ok := c.Usage(text, 0); ok {
			t.Errorf("Usage(%q) = %+v, want none", text, u)
		}
	}
}

// TestUsageInterpolated checks that ${...} words still find their node.
func TestUsageInterpolated(t *testing.T) {
	c := checker(t)
	u, ok := c.Usage("give ${who} diamond", 14)
	if !ok || u.Arg != "<item: item>" || strings.Join(u.Lines, "\n") != "give <targets> <item> [<count>]" {
		t.Errorf("Usage = %+v, %v", u, ok)
	}
}
