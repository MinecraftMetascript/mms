package ide

import (
	"strings"
	"testing"
)

func TestBuildTimeHover(t *testing.T) {
	cases := []struct{ name, src, want string }{
		{"loop variable", "Const biomes = [\"meadow\", \"forest\"]\nfor b| in biomes {\n  Function \"f_${b}\" {\n    /say hi\n  }\n}\n", "meadow"},
		{"loop list", "Const biomes = [\"meadow\", \"forest\"]\nfor b in bio|mes {\n  Function \"f_${b}\" {\n    /say hi\n  }\n}\n", "Const"},
		{"const use in body", "Const n = 4\nFunction f {\n  let x = n|\n}\n", "Value: `4`"},
		{"loop variable in body", "Const biomes = [\"meadow\", \"forest\"]\nfor b in biomes {\n  BiomeTag \"t_${b}\" = [b|]\n}\n", "meadow"},
		{"unrolled variable", "Function f {\n  for a| in [\"x\", \"y\"] {\n    /say ${a}\n  }\n}\n", "loop variable"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s, off := session(t, tc.src)
			h, _ := s.Hover("main.mms", off)
			if !strings.Contains(h, tc.want) {
				t.Fatalf("hover = %q, want %q", h, tc.want)
			}
		})
	}
}

func TestCompletionInLoopBody(t *testing.T) {
	s, off := session(t, "Const biomes = [\"meadow\"]\nfor b in biomes {\n  LootTable \"l_${b}\" {\n    |\n  }\n}\n")
	items, _ := s.Completion("main.mms", off)
	if !has(labels(items), "pools") {
		t.Fatalf("completion in a loop body: %v", labels(items))
	}
}

func TestSymbolsInLoops(t *testing.T) {
	s, _ := session(t, "for b in [\"a\"] {\n  Function \"f_${b}\" {\n    /say hi\n  }\n}\n")
	syms := s.Symbols("main.mms")
	if len(syms) != 1 || syms[0].Name != "f_${b}" {
		t.Fatalf("symbols = %+v", syms)
	}
}
