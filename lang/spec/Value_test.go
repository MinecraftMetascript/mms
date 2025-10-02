package spec

import (
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestExtractPrefixAtPosition(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		line       uint32
		character  uint32
		wantPrefix string
	}{
		{
			name:       "simple prefix",
			source:     "Foo",
			line:       0,
			character:  2,
			wantPrefix: "fo",
		},
		{
			name:       "prefix with namespace",
			source:     "minecraft:stone",
			line:       0,
			character:  10,
			wantPrefix: "minecraft:",
		},
		{
			name:       "no prefix at start",
			source:     "Foo",
			line:       0,
			character:  0,
			wantPrefix: "",
		},
		{
			name:       "prefix after whitespace",
			source:     "  Bar",
			line:       0,
			character:  4,
			wantPrefix: "ba",
		},
		{
			name:       "prefix after paren",
			source:     "Noise(Fi",
			line:       0,
			character:  8,
			wantPrefix: "fi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := protocol.Position{
				Line:      tt.line,
				Character: tt.character,
			}
			got := ExtractPrefixAtPosition(tt.source, pos)
			if got.Text != tt.wantPrefix {
				t.Errorf("ExtractPrefixAtPosition() = %q, want %q", got, tt.wantPrefix)
			}
		})
	}
}

func TestFilterCompletionsByPrefix(t *testing.T) {
	items := []protocol.CompletionItem{
		{Label: "Foo"},
		{Label: "Bar"},
		{Label: "Fizz"},
		{Label: "FooBar"},
	}

	tests := []struct {
		name      string
		prefix    string
		wantCount int
		wantItems []string
	}{
		{
			name:      "empty prefix returns all",
			prefix:    "",
			wantCount: 4,
			wantItems: []string{"Foo", "Bar", "Fizz", "FooBar"},
		},
		{
			name:      "filter by F",
			prefix:    "f",
			wantCount: 3,
			wantItems: []string{"Foo", "Fizz", "FooBar"},
		},
		{
			name:      "filter by Fo",
			prefix:    "fo",
			wantCount: 2,
			wantItems: []string{"Foo", "FooBar"},
		},
		{
			name:      "filter by B",
			prefix:    "b",
			wantCount: 1,
			wantItems: []string{"Bar"},
		},
		{
			name:      "no matches",
			prefix:    "xyz",
			wantCount: 0,
			wantItems: []string{},
		},
		{
			name:      "case insensitive",
			prefix:    "FO",
			wantCount: 2,
			wantItems: []string{"Foo", "FooBar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterCompletionsByPrefix(items, tt.prefix)
			if len(got) != tt.wantCount {
				t.Errorf("FilterCompletionsByPrefix() returned %d items, want %d", len(got), tt.wantCount)
			}

			for _, wantLabel := range tt.wantItems {
				found := false
				for _, item := range got {
					if item.Label == wantLabel {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("FilterCompletionsByPrefix() missing expected item %q", wantLabel)
				}
			}
		})
	}
}
