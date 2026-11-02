package sema

// Widths of the vanilla default font, used to find where a written book
// page runs out of room. A page shows bookLines lines of bookWidth pixels;
// text past the last line is cut off without a warning in game.

const (
	bookWidth = 114
	bookLines = 14
)

// asciiAdvance is how far each printable ASCII character moves the pen in
// the default font, its glyph width plus one pixel of spacing.
var asciiAdvance = func() [128]int {
	var w [128]int
	for i := range w {
		w[i] = 6
	}
	set := func(n int, chars string) {
		for _, r := range chars {
			w[r] = n
		}
	}
	set(2, "!',.:;i|")
	set(3, "`l")
	set(4, " \"()*I[]t{}")
	set(5, "<>fk")
	set(7, "@~")
	return w
}()

// glyphAdvance is a character's advance in pixels. Characters outside
// ASCII come from other bitmaps or the unifont fallback, so their width
// is a guess.
func glyphAdvance(r rune, bold bool) int {
	w := 6
	switch {
	case r < 128:
		w = asciiAdvance[r]
	case r >= 0x2E80:
		w = 9 // CJK and other wide scripts
	}
	if bold {
		w++
	}
	return w
}

// textAdvance is the width of a string in pixels.
func textAdvance(s string, bold bool) int {
	n := 0
	for _, r := range s {
		n += glyphAdvance(r, bold)
	}
	return n
}

// keyLabels are the default key names keybind components show, for
// measuring them. Keys not listed are measured as one letter.
var keyLabels = map[string]string{
	"key.jump": "Space", "key.sneak": "Left Shift", "key.sprint": "Left Control",
	"key.attack": "Left Button", "key.use": "Right Button", "key.pickItem": "Middle Button",
	"key.playerlist": "Tab", "key.togglePerspective": "F5", "key.screenshot": "F2",
	"key.fullscreen": "F11", "key.smoothCamera": "Not Bound", "key.spectatorOutlines": "Not Bound",
	"key.socialInteractions": "P", "key.saveToolbarActivator": "C", "key.loadToolbarActivator": "X",
}

func keyAdvance(key string, bold bool) int {
	if l, ok := keyLabels[key]; ok {
		return textAdvance(l, bold)
	}
	return glyphAdvance('W', bold)
}

// bookGlyph is one unit of page layout: a character, or a component such
// as a sprite or a score, which never breaks across lines.
type bookGlyph struct {
	w       int
	space   bool
	newline bool
	at      bookPos
}

// bookPos is where a glyph came from, for pointing a warning at it.
type bookPos struct {
	src *bookSrc
	off int // rune offset in src.text, -1 when unknown
}

// layoutPage wraps a page's glyphs the way the book screen does, breaking
// lines at spaces, and returns how many lines it takes and the glyph that
// starts the first line past the bottom of the page (index -1 when it fits).
func layoutPage(gs []bookGlyph) (lines, cut int) {
	lines, cut = 1, -1
	x, lastSpace := 0, -1
	start := 0 // index of the first glyph of the current line
	newLine := func(at int) {
		lines++
		if lines == bookLines+1 && cut < 0 {
			cut = at
		}
		start, lastSpace = at, -1
	}
	for i := 0; i < len(gs); i++ {
		g := gs[i]
		if g.newline {
			newLine(i + 1)
			x = 0
			continue
		}
		if x+g.w <= bookWidth {
			x += g.w
			if g.space {
				lastSpace = i
			}
			continue
		}
		switch {
		case g.space:
			// The space ends the line and is not drawn.
			newLine(i + 1)
			x = 0
		case lastSpace >= start:
			// Move the last word to the next line.
			ls := lastSpace
			newLine(ls + 1)
			x = 0
			for j := ls + 1; j <= i; j++ {
				x += gs[j].w
			}
		default:
			newLine(i)
			x = g.w
		}
	}
	// A page's last line can be empty, after a final newline.
	if n := len(gs); n > 0 && gs[n-1].newline {
		lines--
		if lines <= bookLines {
			cut = -1
		}
	}
	return lines, cut
}
