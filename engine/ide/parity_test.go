package ide

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
	"github.com/minecraftmetascript/mms/engine/syntax/parser"
)

// TestCompletionParity guards completion across real code. At every
// third word of the examples it cuts the word to its first letter, as
// if it were being typed, and checks that completion offers the word.
// It does so twice: with the rest of the line kept, and with the rest of
// the line gone, which is how a line being typed looks. The hit rates per
// kind of position, such as a field name or a call in a list, must not
// fall below testdata/completion.tsv. MMS_UPDATE=1 rewrites that file
// after a change that raises them.
func TestCompletionParity(t *testing.T) {
	if testing.Short() {
		t.Skip("reads every example")
	}
	type tally struct{ hit, n int }
	got := map[string]*tally{}
	add := func(key string, hit bool) {
		if got[key] == nil {
			got[key] = &tally{}
		}
		got[key].n++
		if hit {
			got[key].hit++
		}
	}
	var misses []string
	dirs, _ := filepath.Glob("../../examples/*")
	for _, dir := range dirs {
		cfg, err := project.LoadConfig(dir)
		if err != nil {
			continue
		}
		paths, err := cfg.SourceFiles(dir)
		if err != nil {
			t.Fatal(err)
		}
		texts := map[string]string{}
		p := project.New(spec.MustDefault(), cfg)
		p.Files = os.DirFS(dir)
		for _, rel := range paths {
			b, err := os.ReadFile(filepath.Join(dir, rel))
			if err != nil {
				t.Fatal(err)
			}
			texts[rel] = string(b)
			p.SetFile(rel, string(b))
		}
		r0, err := p.Analyze(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for _, rel := range paths {
			u := r0.Unit(rel)
			if !strings.HasSuffix(rel, ".mms") || u == nil {
				continue
			}
			src := u.File.Src
			n := 0
			for _, tk := range u.File.Tokens {
				if tk.Hidden || tk.Synthetic || tk.Sp.Len() < 2 {
					continue
				}
				if tk.Kind != parser.MMSLexerIDENT && tk.Kind != parser.MMSLexerRESLOC && tk.Kind != parser.MMSLexerTNAME {
					continue
				}
				if n++; n%3 != 0 {
					continue
				}
				ctx := positionKind(u.File, tk.Sp.Start)
				if ctx == "" {
					continue // a name being declared: nothing to offer
				}
				word := string(src[tk.Sp.Start:tk.Sp.End])
				for _, mode := range []string{"word", "line"} {
					rest := string(src[tk.Sp.End:])
					if mode == "line" {
						if nl := strings.Index(rest, "\n"); nl >= 0 {
							rest = rest[nl:]
						} else {
							rest = ""
						}
					}
					p.SetFile(rel, string(src[:tk.Sp.Start+1])+rest)
					r, err := p.Analyze(context.Background())
					if err != nil {
						t.Fatal(err)
					}
					items, _ := (&Session{P: p, R: r}).Completion(rel, tk.Sp.Start+1)
					hit := offers(items, word)
					add(mode+"\t"+ctx, hit)
					add(mode+"\tall", hit)
					if !hit {
						line, _ := u.File.LineCol(tk.Sp.Start)
						misses = append(misses, fmt.Sprintf("%s %s/%s:%d %s (%s)", mode, filepath.Base(dir), rel, line, word, ctx))
					}
				}
				p.SetFile(rel, texts[rel])
			}
		}
	}

	baseline := "testdata/completion.tsv"
	if os.Getenv("MMS_UPDATE") != "" {
		keys := make([]string, 0, len(got))
		for k := range got {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var sb strings.Builder
		sb.WriteString("# mode\tposition\thits\tchecked (written by TestCompletionParity with MMS_UPDATE=1)\n")
		for _, k := range keys {
			fmt.Fprintf(&sb, "%s\t%d\t%d\n", k, got[k].hit, got[k].n)
		}
		if err := os.WriteFile(baseline, []byte(sb.String()), 0o644); err != nil {
			t.Fatal(err)
		}
		return
	}
	b, err := os.ReadFile(baseline)
	if err != nil {
		t.Fatalf("%v; run with MMS_UPDATE=1 to write it", err)
	}
	failed := false
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		f := strings.Split(line, "\t")
		if len(f) != 4 {
			continue
		}
		key := f[0] + "\t" + f[1]
		var hit, n int
		fmt.Sscan(f[2], &hit)
		fmt.Sscan(f[3], &n)
		g := got[key]
		if g == nil || n == 0 {
			continue // the examples changed; MMS_UPDATE=1 takes the new counts
		}
		// A rate may fall by a point, since the examples change too.
		if float64(g.hit)/float64(g.n) < float64(hit)/float64(n)-0.01 {
			t.Errorf("%s %s: completion offers the word at %d of %d positions, down from %d of %d", f[0], f[1], g.hit, g.n, hit, n)
			failed = true
		}
	}
	if failed {
		t.Logf("missed:\n%s", strings.Join(misses, "\n"))
	}
}

// positionKind names the kind of position at off by its innermost three
// nodes, such as "Block>FieldStmt>Name" for a field name, or "" for a
// name being declared.
func positionKind(f *syntax.File, off int) string {
	path := syntax.PathAt(f, off)
	var kinds []string
	for _, n := range path {
		kinds = append(kinds, strings.TrimPrefix(reflect.TypeOf(n).String(), "*syntax."))
	}
	if len(kinds) > 3 {
		kinds = kinds[len(kinds)-3:]
	}
	k := strings.Join(kinds, ">")
	for _, decl := range []string{"Decl>Name", "LetStmt>Name", "Param>Name", "ConstStmt>Name", "ForStmt>Name"} {
		if strings.HasSuffix(k, decl) {
			return ""
		}
	}
	return k
}

// offers reports whether items hold word, or the same id written another
// way: `stone` for minecraft:stone, `logs` for #logs.
func offers(items []Item, word string) bool {
	for _, it := range items {
		l := strings.Trim(it.Label, `"`)
		if l == word || strings.HasSuffix(l, ":"+word) || strings.HasSuffix(word, ":"+l) || strings.TrimPrefix(l, "#") == word {
			return true
		}
	}
	return false
}
