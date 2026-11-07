// Package near finds the likely intended name for a typo, so every "did
// you mean" in the engine suggests the same way.
package near

import (
	"sort"
	"strings"
)

// Closest returns the candidate nearest to s, or "" when none is close
// enough to be a likely typo: fewer than len(s)/3+2 edits, ignoring case.
// Ties go to the candidate that sorts first, so the answer does not depend
// on the order of cands.
func Closest(s string, cands []string) string {
	best, bestD := "", len([]rune(s))/3+2
	ls := strings.ToLower(s)
	for _, c := range cands {
		d := Distance(ls, strings.ToLower(c))
		if d < bestD || d == bestD && best != "" && c < best {
			best, bestD = c, d
		}
	}
	return best
}

// ClosestKey is Closest over a set's keys.
func ClosestKey[V any](s string, set map[string]V) string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return Closest(s, keys)
}

// Distance is the Levenshtein distance between a and b, in runes. Names
// longer than 64 runes are never close.
func Distance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	if len(ra) > 64 || len(rb) > 64 {
		return 99
	}
	prev := make([]int, len(rb)+1)
	cur := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev, cur = cur, prev
	}
	return prev[len(rb)]
}
