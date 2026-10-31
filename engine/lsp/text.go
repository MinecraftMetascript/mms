package lsp

import (
	"unicode/utf8"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Position is an LSP position: a zero-based line and a UTF-16 column.
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// Range is an LSP range.
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// lines indexes a text's line starts, in runes, for position conversion.
type lines struct {
	src    []rune
	starts []int
}

func newLines(text string) *lines {
	src := []rune(text)
	l := &lines{src: src, starts: []int{0}}
	for i, r := range src {
		if r == '\n' {
			l.starts = append(l.starts, i+1)
		}
	}
	return l
}

// pos converts a rune offset to an LSP position.
func (l *lines) pos(off int) Position {
	if off > len(l.src) {
		off = len(l.src)
	}
	lo, hi := 0, len(l.starts)-1
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if l.starts[mid] <= off {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	col := 0
	for _, r := range l.src[l.starts[lo]:off] {
		col += utf16Len(r)
	}
	return Position{Line: lo, Character: col}
}

// offset converts an LSP position to a rune offset, clamping to the line.
func (l *lines) offset(p Position) int {
	if p.Line < 0 {
		return 0
	}
	if p.Line >= len(l.starts) {
		return len(l.src)
	}
	off := l.starts[p.Line]
	for col := 0; off < len(l.src) && l.src[off] != '\n'; off++ {
		n := utf16Len(l.src[off])
		if col+n > p.Character {
			break
		}
		col += n
	}
	return off
}

func (l *lines) rng(sp syntax.Span) Range {
	return Range{Start: l.pos(sp.Start), End: l.pos(sp.End)}
}

func utf16Len(r rune) int {
	if r >= 0x10000 && utf8.ValidRune(r) {
		return 2
	}
	return 1
}
