package sema

import (
	"strings"
	"unicode"
)

// Snake converts a PascalCase name to snake_case: XzScale -> xz_scale,
// YAbove -> y_above.
func Snake(s string) string {
	rs := []rune(s)
	var sb strings.Builder
	for i, r := range rs {
		if unicode.IsUpper(r) {
			prevLower := i > 0 && (unicode.IsLower(rs[i-1]) || unicode.IsDigit(rs[i-1]))
			nextLower := i > 0 && i+1 < len(rs) && unicode.IsUpper(rs[i-1]) && unicode.IsLower(rs[i+1])
			if prevLower || nextLower {
				sb.WriteByte('_')
			}
			sb.WriteRune(unicode.ToLower(r))
			continue
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

// Pascal converts snake_case (or a path's last segment) to PascalCase:
// y_above -> YAbove.
func Pascal(s string) string {
	if i := strings.LastIndexAny(s, ":/"); i >= 0 {
		s = s[i+1:]
	}
	var sb strings.Builder
	for _, part := range strings.Split(s, "_") {
		if part == "" {
			continue
		}
		rs := []rune(part)
		rs[0] = unicode.ToUpper(rs[0])
		sb.WriteString(string(rs))
	}
	return sb.String()
}

// isPascal reports whether a name starts with an upper-case letter, the
// convention for helpers and declaration kinds.
func isPascal(s string) bool {
	for _, r := range s {
		return unicode.IsUpper(r)
	}
	return false
}

// splitID splits "ns:path" into its parts; ns is "" when unqualified.
func splitID(id string) (ns, path string) {
	if i := strings.IndexByte(id, ':'); i >= 0 {
		return id[:i], id[i+1:]
	}
	return "", id
}
