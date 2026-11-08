package syntax

import "testing"

func TestUnquote(t *testing.T) {
	for raw, want := range map[string]string{
		`"plain"`:              "plain",
		`"tab\there"`:          "tab\there",
		`"a\nb"`:               "a\nb",
		`"say \"hi\""`:         `say "hi"`,
		`'it\'s'`:              "it's",
		`"back\\slash"`:        `back\slash`,
		`"é"`:                  "é",
		`"tab\there \{x}"`:     "tab\there \\{x}",
		`"'quoted' \{x}"`:      `'quoted' \{x}`,
		`"5 \* 3"`:             `5 \* 3`,
		`"cost \${n}"`:         `cost \${n}`,
		`'"double" in single'`: `"double" in single`,
	} {
		if got := unquote(raw); got != want {
			t.Errorf("unquote(%s) = %q, want %q", raw, got, want)
		}
	}
}
