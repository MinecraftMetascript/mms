package sema

import (
	"strings"
	"testing"
)

func linesOf(text string, bold bool) int {
	var gs []bookGlyph
	for _, r := range text {
		gs = append(gs, bookGlyph{w: glyphAdvance(r, bold), space: r == ' ', newline: r == '\n'})
	}
	n, _ := layoutPage(gs)
	return n
}

func TestLayoutPage(t *testing.T) {
	for _, tc := range []struct {
		text  string
		bold  bool
		lines int
	}{
		{strings.Repeat("a", 19), false, 1}, // 19 * 6 = 114 pixels fit
		{strings.Repeat("a", 20), false, 2},
		{strings.Repeat("a", 16), true, 1}, // bold glyphs are a pixel wider
		{strings.Repeat("a", 17), true, 2},
		{strings.Repeat("i", 57), false, 1}, // narrow letters fit more
		{"aaaaaaaaaa aaaaaaaaaa", false, 2}, // the second word moves down whole
		{"a\nb\n", false, 2},                // a final newline adds no line
		{strings.Repeat("a\n", 13) + "a", false, 14},
	} {
		if got := linesOf(tc.text, tc.bold); got != tc.lines {
			t.Errorf("%q bold=%v: %d lines, want %d", tc.text, tc.bold, got, tc.lines)
		}
	}
	var gs []bookGlyph
	for i := 0; i < 16; i++ {
		gs = append(gs, bookGlyph{w: 6}, bookGlyph{newline: true})
	}
	if n, cut := layoutPage(gs); n != 16 || cut != 28 {
		t.Errorf("16 lines: n=%d cut=%d, want 16 and 28", n, cut)
	}
}
