package ide

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

const docLib = `/**
 * Adds two numbers.
 *
 * @param a the first
 * @param b the second
 * @returns the sum
 */
Function add(a: int, b: int): int {
  /** The running total. */
  let x = a + b
  return x
}

/** Total kills. */
Global total: int = 0

/** Points per player. */
Objective pts = dummy

/** @deprecated use add */
Function old_reward {
  /give @s diamond
}

/** Ridges for the terrain. */
Noise ridges = { first_octave = -1, amplitudes = [1] }
`

const docMain = `Function main {
  total += add(1, 2)
  pts[@s] = total
  /function test:old_reward
  /scoreboard players set @s pts 1
  /say ${total}
}
DensityFunction d = Noise(ridges)
`

// docSession builds a two-file project: declarations with doc comments in
// lib.mms, uses of them in main.mms.
func docSession(t *testing.T, main string) *Session {
	t.Helper()
	p := project.New(spec.MustDefault(), project.Config{Namespace: "test"})
	p.SetFile("lib.mms", docLib)
	p.SetFile("main.mms", main)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return &Session{P: p, R: r}
}

// offOf is the rune offset of the first `at` in src, plus skip.
func offOf(src, at string, skip int) int {
	i := strings.Index(src, at)
	if i < 0 {
		panic("missing " + at)
	}
	return len([]rune(src[:i])) + skip
}

func TestDocHover(t *testing.T) {
	s := docSession(t, docMain)
	cases := []struct {
		name, file, at string
		skip           int
		want           []string
	}{
		{"global use", "main.mms", "total +=", 1, []string{"Global test:total: int", "Total kills."}},
		{"function call", "main.mms", "add(1", 1, []string{"Function test:add(a: int, b: int): int", "Adds two numbers.", "- `a`: the first", "**Returns** the sum"}},
		{"objective score", "main.mms", "pts[", 1, []string{"Objective test:pts", "Points per player."}},
		{"function in command", "main.mms", "test:old_reward", 6, []string{"Function test:old_reward", "**Deprecated.** use add"}},
		{"objective in command", "main.mms", "@s pts", 4, []string{"Objective test:pts", "Points per player."}},
		{"interpolation", "main.mms", "${total}", 3, []string{"Global test:total", "Total kills."}},
		{"resource reference", "main.mms", "Noise(ridges)", 7, []string{"Noise test:ridges", "Ridges for the terrain."}},
		{"declaration", "lib.mms", "add(a", 1, []string{"Function test:add(a: int, b: int): int", "Adds two numbers."}},
		{"parameter declaration", "lib.mms", "a: int, b", 0, []string{"(parameter) a: int", "the first"}},
		{"parameter use", "lib.mms", "a + b", 4, []string{"(parameter) b: int", "the second"}},
		{"let declaration", "lib.mms", "x = a", 0, []string{"let x: int", "The running total."}},
		{"let use", "lib.mms", "return x", 7, []string{"let x: int", "The running total."}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			src := docMain
			if tc.file == "lib.mms" {
				src = docLib
			}
			md, _ := s.Hover(tc.file, offOf(src, tc.at, tc.skip))
			for _, w := range tc.want {
				if !strings.Contains(md, w) {
					t.Errorf("hover lacks %q:\n%s", w, md)
				}
			}
		})
	}
}

// TestCommandReferences checks go to definition and references for ids
// written inside commands, which have no syntax nodes of their own.
func TestCommandReferences(t *testing.T) {
	s := docSession(t, docMain)
	off := offOf(docMain, "test:old_reward", 6)
	defs := s.Definition("main.mms", off)
	if len(defs) != 1 || defs[0].Path != "lib.mms" || defs[0].Span.Start != offOf(docLib, "old_reward {", 0) {
		t.Errorf("definition = %+v", defs)
	}
	refs := s.References("lib.mms", offOf(docLib, "pts =", 0), false)
	if len(refs) != 2 {
		t.Errorf("objective references = %+v", refs)
	}
	edits, msg := s.Rename("main.mms", off, "legacy_reward")
	if msg != "" {
		t.Fatal(msg)
	}
	var texts []string
	for _, e := range edits {
		texts = append(texts, e.Path+":"+e.Text)
	}
	if got := strings.Join(texts, " "); got != "lib.mms:legacy_reward main.mms:test:legacy_reward" {
		t.Errorf("rename edits = %s", got)
	}
}

// TestCommandHoverAfterMove checks that references inside commands survive
// when their declaration is reused at a new offset.
func TestCommandHoverAfterMove(t *testing.T) {
	s := docSession(t, docMain)
	moved := "DensityFunction first = 1\n\n" + docMain
	s.P.SetFile("main.mms", moved)
	r, err := s.P.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	s.R = r
	md, _ := s.Hover("main.mms", offOf(moved, "test:old_reward", 6))
	if !strings.Contains(md, "Function test:old_reward") {
		t.Errorf("hover after move = %q", md)
	}
}

func TestDocSignatureAndCompletion(t *testing.T) {
	main := "Function main {\n  let y = add(1, 2)\n}\nDensityFunction d = Noise()\n"
	s := docSession(t, main)
	h, ok := s.Signature("main.mms", offOf(main, "2)", 0))
	if !ok || h.Label != "add(a: int, b: int): int" || h.Active != 1 || h.Doc != "Adds two numbers." || h.ParamDocs[1] != "the second" {
		t.Errorf("signature = %+v, %v", h, ok)
	}
	items, _ := s.Completion("main.mms", offOf(main, "Noise()", 6))
	found := false
	for _, it := range items {
		if it.Label == "ridges" {
			found = true
			if it.Doc != "Ridges for the terrain." {
				t.Errorf("completion doc = %q", it.Doc)
			}
		}
	}
	if !found {
		t.Errorf("no completion for ridges in %v", labels(items))
	}
	cmd := "Function main {\n  /function test:\n}\n"
	s = docSession(t, cmd)
	items, _ = s.Completion("main.mms", offOf(cmd, ":\n", 1))
	found = false
	for _, it := range items {
		if it.Label == "test:old_reward" {
			found = true
			if !it.Deprecated || !strings.Contains(it.Doc, "Deprecated") {
				t.Errorf("command completion = %+v", it)
			}
		}
	}
	if !found {
		t.Errorf("no command completion for test:old_reward in %v", labels(items))
	}
}
