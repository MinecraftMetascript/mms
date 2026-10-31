package project

import (
	"context"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// ApplyFixes applies the first safe fix of each diagnostic in ds (sorted
// by position) to src, skipping fixes that overlap an earlier one. It returns the new
// text and how many fixes were applied.
func ApplyFixes(src []rune, ds []syntax.Diagnostic) (string, int) {
	// Diagnostics come sorted by position, so a fix overlaps an earlier
	// one exactly when it starts before the furthest end so far.
	var edits []syntax.TextEdit
	n, reach := 0, -1
	for _, d := range ds {
		for _, f := range d.Fixes {
			if !f.Safe {
				continue
			}
			lo, hi := len(src), -1
			for _, e := range f.Edits {
				lo, hi = min(lo, e.Sp.Start), max(hi, e.Sp.End)
			}
			if lo <= reach {
				break
			}
			edits = append(edits, f.Edits...)
			reach = max(reach, hi)
			n++
			break
		}
	}
	sort.SliceStable(edits, func(i, j int) bool { return edits[i].Sp.Start < edits[j].Sp.Start })
	var sb strings.Builder
	at := 0
	for _, e := range edits {
		sb.WriteString(string(src[at:e.Sp.Start]))
		sb.WriteString(e.Text)
		at = e.Sp.End
	}
	sb.WriteString(string(src[at:]))
	return sb.String(), n
}

// FixAll applies safe fixes to the given files (every file when none are
// given) until none are left, since a fix can expose another, for at most
// a few rounds. It returns the files it changed, with their new text, and
// the number of fixes applied.
func (p *Project) FixAll(ctx context.Context, paths ...string) (map[string]string, int, error) {
	if len(paths) == 0 {
		paths = p.Paths()
	}
	changed := map[string]string{}
	total := 0
	for round := 0; round < 5; round++ {
		r, err := p.Analyze(ctx)
		if err != nil {
			return changed, total, err
		}
		n := 0
		for _, path := range paths {
			u := r.Unit(path)
			if u == nil {
				continue
			}
			text, k := ApplyFixes(u.File.Src, r.Diagnostics(path))
			if k > 0 {
				p.SetFile(path, text)
				changed[path] = text
				n += k
			}
		}
		if n == 0 {
			break
		}
		total += n
	}
	return changed, total, nil
}
