package spec

import (
	"strconv"
	"strings"
)

// CompareVersions compares release versions such as "1.21.5" and "26.3" as
// dotted integers. Missing parts count as zero, so "26.3" equals "26.3.0".
func CompareVersions(a, b string) int {
	as, bs := strings.Split(a, "."), strings.Split(b, ".")
	for i := 0; i < len(as) || i < len(bs); i++ {
		x, y := part(as, i), part(bs, i)
		if x != y {
			if x < y {
				return -1
			}
			return 1
		}
	}
	return 0
}

func part(ps []string, i int) int {
	if i >= len(ps) {
		return 0
	}
	n, _ := strconv.Atoi(ps[i])
	return n
}

// InVersion reports whether attributes gate a node into the given version:
// `since` is inclusive and `until` is exclusive (decision D4).
func InVersion(attrs []Attribute, version string) bool {
	for _, a := range attrs {
		v, ok := a.LiteralString()
		if !ok {
			continue
		}
		switch a.Name {
		case "since":
			if CompareVersions(version, v) < 0 {
				return false
			}
		case "until":
			if CompareVersions(version, v) >= 0 {
				return false
			}
		}
	}
	return true
}
