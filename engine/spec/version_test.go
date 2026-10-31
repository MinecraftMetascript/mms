package spec

import "testing"

func TestCompareVersions(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"26.3", "26.3", 0},
		{"26.3", "26.3.0", 0},
		{"1.21.5", "26.2", -1},
		{"26.3", "26.2", 1},
		{"1.21.10", "1.21.9", 1},
	}
	for _, c := range cases {
		if got := CompareVersions(c.a, c.b); got != c.want {
			t.Errorf("CompareVersions(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestInVersion(t *testing.T) {
	lit := func(v string) *Type {
		return &Type{Kind: "literal", Literal: &Literal{Kind: "string", Value: []byte(`"` + v + `"`)}}
	}
	since := func(v string) []Attribute { return []Attribute{{Name: "since", Value: lit(v)}} }
	until := func(v string) []Attribute { return []Attribute{{Name: "until", Value: lit(v)}} }
	if !InVersion(since("26.2"), "26.3") || !InVersion(since("26.3"), "26.3") || InVersion(since("26.4"), "26.3") {
		t.Error("since is inclusive")
	}
	if InVersion(until("26.2"), "26.3") || InVersion(until("26.3"), "26.3") || !InVersion(until("26.4"), "26.3") {
		t.Error("until is exclusive")
	}
}
