package types_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
)

var update = flag.Bool("update", false, "rewrite the golden files")

// TestGolden analyzes each case in testdata (`-- name.mms --` sections)
// and compares every diagnostic with the case's golden file.
func TestGolden(t *testing.T) {
	cases, _ := filepath.Glob("testdata/*.txt")
	for _, path := range cases {
		t.Run(filepath.Base(path), func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			p := project.New(spec.MustDefault(), project.Config{Namespace: "demo"})
			var names []string
			var name string
			var body strings.Builder
			flush := func() {
				if name != "" {
					p.SetFile(name, body.String())
					names = append(names, name)
				}
			}
			for _, line := range strings.SplitAfter(string(data), "\n") {
				if strings.HasPrefix(line, "-- ") {
					flush()
					name = strings.Fields(line)[1]
					body.Reset()
					continue
				}
				body.WriteString(line)
			}
			flush()
			r, err := p.Analyze(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			var b strings.Builder
			for _, n := range names {
				src := r.Unit(n).File.Src
				for _, d := range r.Diagnostics(n) {
					line, col := 1, 1
					for _, c := range src[:d.Sp.Start] {
						if c == '\n' {
							line, col = line+1, 1
						} else {
							col++
						}
					}
					fmt.Fprintf(&b, "%s:%d:%d %s: %s [%q]\n", n, line, col, d.Code, d.Message, string(src[d.Sp.Start:min(d.Sp.End, len(src))]))
				}
			}
			got := b.String()
			golden := strings.TrimSuffix(path, ".txt") + ".golden"
			if *update {
				os.WriteFile(golden, []byte(got), 0o644)
				return
			}
			want, _ := os.ReadFile(golden)
			if got != string(want) {
				t.Errorf("differs from %s (rerun with -update):\n%s", golden, got)
			}
		})
	}
}
