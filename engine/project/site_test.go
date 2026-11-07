package project

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// siteSample is a fenced code block in one of the website's pages.
type siteSample struct {
	path string // relative to the repository root
	line int    // of the opening fence
	info string // the fence's words after "mms", such as "error"
	src  string
}

// siteSamples returns the ```mms blocks of every page under site/. A
// block fenced as plain ``` is not checked: it may be a fragment, a
// shell session or output.
func siteSamples(t *testing.T) []siteSample {
	t.Helper()
	var out []siteSample
	root := filepath.Join("..", "..")
	err := filepath.WalkDir(filepath.Join(root, "site"), func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, path)
		var cur *siteSample
		var body []string
		for i, line := range strings.Split(string(b), "\n") {
			trimmed := strings.TrimSpace(line)
			switch {
			case cur != nil && trimmed == "```":
				cur.src = strings.Join(body, "\n") + "\n"
				out = append(out, *cur)
				cur, body = nil, nil
			case cur != nil:
				body = append(body, line)
			case trimmed == "```mms" || strings.HasPrefix(trimmed, "```mms "):
				cur = &siteSample{path: filepath.ToSlash(rel), line: i + 1, info: strings.TrimSpace(strings.TrimPrefix(trimmed, "```mms"))}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// TestSiteSamples checks every ```mms block on the website as a file of
// its own, so the docs break the build when the language moves on
// without them. A block fenced ```mms error shows a mistake on purpose
// and must report at least one error. ```mms in=examples/jigsaw reads
// asset files, such as structure templates, from that folder.
func TestSiteSamples(t *testing.T) {
	samples := siteSamples(t)
	if len(samples) == 0 {
		t.Fatal("no ```mms samples found under site/")
	}
	pack := spec.MustDefault()
	for _, s := range samples {
		p := New(pack, DefaultConfig())
		wantErr := false
		bad := false
		for _, w := range strings.Fields(s.info) {
			switch {
			case w == "error":
				wantErr = true
			case strings.HasPrefix(w, "in="):
				p.Files = os.DirFS(filepath.Join("..", "..", filepath.FromSlash(strings.TrimPrefix(w, "in="))))
			default:
				t.Errorf("%s:%d: unknown fence word %q: use ```mms, ```mms error or ```mms in=<folder>", s.path, s.line, w)
				bad = true
			}
		}
		if bad {
			continue
		}
		p.SetFile("main.mms", s.src)
		r, err := p.Analyze(context.Background())
		if err != nil {
			t.Errorf("%s:%d: %v", s.path, s.line, err)
			continue
		}
		var errs []string
		for _, u := range r.Units {
			for _, d := range r.Diagnostics(u.Path) {
				if d.Severity == syntax.SeverityError {
					errs = append(errs, d.Code+" "+d.Message)
				}
			}
		}
		if wantErr {
			if len(errs) == 0 {
				t.Errorf("%s:%d: a ```mms error sample reports no error", s.path, s.line)
			}
			continue
		}
		for _, e := range errs {
			t.Errorf("%s:%d: %s", s.path, s.line, e)
		}
	}
}
