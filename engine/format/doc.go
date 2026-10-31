// Package format is the MMS formatter: one style, no options except line
// width. It prints the AST through a Wadler-style document IR and a
// printer that breaks a group only when it does not fit.
package format

import "strings"

// doc is a node of the document IR.
type doc interface{}

type (
	text     string
	line     struct{ soft bool } // a space (or nothing if soft) when flat, a newline when broken
	hardLine struct{}            // always a newline; forces enclosing groups to break
	group    struct{ docs []doc }
	indent   struct{ docs []doc }
	// ifBreak prints broken in a broken group and flat otherwise.
	ifBreak struct{ broken, flat doc }
	concat  []doc
	// forceBreak prints nothing but makes every enclosing group break; a
	// trailing line comment uses it because it must end its line.
	forceBreak struct{}
)

var (
	ln   = line{}
	soft = line{soft: true}
	hard = hardLine{}
)

func cat(ds ...doc) doc    { return concat(ds) }
func grp(ds ...doc) doc    { return group{ds} }
func ind(ds ...doc) doc    { return indent{ds} }
func txt(s string) doc     { return text(s) }
func whenBroken(d doc) doc { return ifBreak{broken: d, flat: text("")} }

// join puts sep between docs.
func join(sep doc, ds []doc) doc {
	out := make(concat, 0, 2*len(ds))
	for i, d := range ds {
		if i > 0 {
			out = append(out, sep)
		}
		out = append(out, d)
	}
	return out
}

type mode int

const (
	flat mode = iota
	broken
)

type cmd struct {
	indent int
	mode   mode
	d      doc
}

// render prints a document at the given width with two-space indents.
func render(d doc, width int) string {
	var sb strings.Builder
	col := 0
	stack := []cmd{{0, broken, d}}
	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		switch d := c.d.(type) {
		case nil:
		case text:
			sb.WriteString(string(d))
			col += len([]rune(string(d)))
		case concat:
			for i := len(d) - 1; i >= 0; i-- {
				stack = append(stack, cmd{c.indent, c.mode, d[i]})
			}
		case indent:
			for i := len(d.docs) - 1; i >= 0; i-- {
				stack = append(stack, cmd{c.indent + 2, c.mode, d.docs[i]})
			}
		case group:
			m := flat
			if c.mode == broken && (hasHard(concat(d.docs)) || !fits(concat(d.docs), width-col, stack)) {
				m = broken
			}
			if c.mode == flat {
				m = flat
			}
			for i := len(d.docs) - 1; i >= 0; i-- {
				stack = append(stack, cmd{c.indent, m, d.docs[i]})
			}
		case ifBreak:
			if c.mode == broken {
				stack = append(stack, cmd{c.indent, c.mode, d.broken})
			} else {
				stack = append(stack, cmd{c.indent, c.mode, d.flat})
			}
		case line:
			if c.mode == flat {
				if !d.soft {
					sb.WriteByte(' ')
					col++
				}
				continue
			}
			newline(&sb, c.indent)
			col = c.indent
		case hardLine:
			newline(&sb, c.indent)
			col = c.indent
		}
	}
	return sb.String()
}

func newline(sb *strings.Builder, indent int) {
	// Trim trailing spaces left by a preceding separator.
	s := sb.String()
	trimmed := strings.TrimRight(s, " ")
	if len(trimmed) != len(s) {
		sb.Reset()
		sb.WriteString(trimmed)
	}
	sb.WriteByte('\n')
	sb.WriteString(strings.Repeat(" ", indent))
}

// fits reports whether d printed flat, followed by the rest of the current
// line from the stack, fits in the remaining width.
func fits(d doc, remaining int, rest []cmd) bool {
	work := []cmd{{0, flat, d}}
	restIdx := len(rest) - 1
	for remaining >= 0 {
		if len(work) == 0 {
			if restIdx < 0 {
				return true
			}
			work = append(work, rest[restIdx])
			restIdx--
			continue
		}
		c := work[len(work)-1]
		work = work[:len(work)-1]
		switch d := c.d.(type) {
		case text:
			remaining -= len([]rune(string(d)))
		case concat:
			for i := len(d) - 1; i >= 0; i-- {
				work = append(work, cmd{c.indent, c.mode, d[i]})
			}
		case indent:
			for i := len(d.docs) - 1; i >= 0; i-- {
				work = append(work, cmd{c.indent, c.mode, d.docs[i]})
			}
		case group:
			for i := len(d.docs) - 1; i >= 0; i-- {
				work = append(work, cmd{c.indent, c.mode, d.docs[i]})
			}
		case ifBreak:
			if c.mode == broken {
				work = append(work, cmd{c.indent, c.mode, d.broken})
			} else {
				work = append(work, cmd{c.indent, c.mode, d.flat})
			}
		case line:
			if c.mode == broken {
				return true
			}
			if !d.soft {
				remaining--
			}
		case hardLine, forceBreak:
			return true
		}
	}
	return false
}

// hasHard reports whether a document contains a hard line, which forces
// every enclosing group to break.
func hasHard(d doc) bool {
	switch d := d.(type) {
	case hardLine, forceBreak:
		return true
	case concat:
		for _, x := range d {
			if hasHard(x) {
				return true
			}
		}
	case group:
		return hasHard(concat(d.docs))
	case indent:
		return hasHard(concat(d.docs))
	case ifBreak:
		return hasHard(d.broken)
	}
	return false
}
