package syntax

import (
	"strings"
	"unicode"
)

// cleanDoc turns the text of a /** */ comment into its content: the
// delimiters and the leading `*` of each line are removed, and the
// indentation that remains is kept relative to the least indented line, so
// indented code blocks and nested lists survive.
func cleanDoc(text string) string {
	text = strings.TrimSuffix(strings.TrimPrefix(text, "/**"), "*/")
	text = strings.ReplaceAll(text, "\r\n", "\n")
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		if i > 0 {
			t := strings.TrimLeftFunc(l, unicode.IsSpace)
			if rest, ok := strings.CutPrefix(t, "*"); ok {
				l = strings.TrimPrefix(rest, " ")
			}
		} else {
			l = strings.TrimPrefix(l, " ")
		}
		lines[i] = strings.TrimRightFunc(l, unicode.IsSpace)
	}
	// Drop the indentation every non-blank line after the first shares,
	// for comments written without leading stars.
	common := -1
	for _, l := range lines[1:] {
		if l == "" {
			continue
		}
		n := len(l) - len(strings.TrimLeft(l, " \t"))
		if common < 0 || n < common {
			common = n
		}
	}
	if common > 0 {
		for i := 1; i < len(lines); i++ {
			if len(lines[i]) >= common {
				lines[i] = lines[i][common:]
			}
		}
	}
	lines[0] = strings.TrimSpace(lines[0])
	return strings.Trim(strings.Join(lines, "\n"), "\n")
}

// Doc is a parsed doc comment: free Markdown text followed by optional
// JSDoc-style block tags such as `@param name text`.
type Doc struct {
	Text       string   // the description, Markdown
	Params     []DocTag // @param name text
	Returns    string   // @returns text (or @return)
	Deprecated *string  // @deprecated [text]; nil when absent
	See        []string // @see text
	Examples   []string // @example code
	Other      []DocTag // any other tag, Name being the tag
}

// DocTag is one `@tag` entry.
type DocTag struct {
	Name, Text string
}

// ParseDoc splits a cleaned doc comment into its description and tags. A
// tag starts a line with `@`; its text runs until the next tag. Tags inside
// fenced code blocks are left alone.
func ParseDoc(s string) Doc {
	var d Doc
	var desc []string
	var tag string
	var body []string
	flush := func() {
		if tag == "" {
			return
		}
		text := strings.Trim(strings.Join(body, "\n"), "\n")
		switch tag {
		case "param", "arg", "argument":
			name, rest := cutWord(text)
			name = strings.Trim(name, "[]")
			rest = strings.TrimPrefix(strings.TrimSpace(rest), "- ")
			d.Params = append(d.Params, DocTag{Name: name, Text: rest})
		case "returns", "return":
			d.Returns = text
		case "deprecated":
			d.Deprecated = &text
		case "see":
			d.See = append(d.See, text)
		case "example":
			d.Examples = append(d.Examples, text)
		default:
			d.Other = append(d.Other, DocTag{Name: tag, Text: text})
		}
		tag, body = "", nil
	}
	fence := false
	for _, l := range strings.Split(s, "\n") {
		trimmed := strings.TrimSpace(l)
		if strings.HasPrefix(trimmed, "```") {
			fence = !fence
		}
		if !fence && strings.HasPrefix(trimmed, "@") && len(trimmed) > 1 && unicode.IsLetter(rune(trimmed[1])) {
			flush()
			name, rest := cutWord(trimmed[1:])
			tag = strings.ToLower(name)
			if strings.TrimSpace(rest) != "" {
				body = append(body, strings.TrimSpace(rest))
			}
			continue
		}
		if tag != "" {
			if tag == "example" {
				body = append(body, l)
			} else {
				body = append(body, trimmed)
			}
			continue
		}
		desc = append(desc, l)
	}
	flush()
	d.Text = strings.Trim(strings.Join(desc, "\n"), "\n")
	return d
}

func cutWord(s string) (string, string) {
	s = strings.TrimSpace(s)
	if i := strings.IndexFunc(s, unicode.IsSpace); i >= 0 {
		return s[:i], s[i:]
	}
	return s, ""
}

// Param returns the @param text for name, or "".
func (d Doc) Param(name string) string {
	for _, p := range d.Params {
		if p.Name == name {
			return p.Text
		}
	}
	return ""
}

// Markdown renders the doc comment for hovers and completion details:
// the description, then the tags as short labelled sections.
func (d Doc) Markdown() string {
	var parts []string
	if d.Deprecated != nil {
		msg := "**Deprecated.**"
		if *d.Deprecated != "" {
			msg += " " + *d.Deprecated
		}
		parts = append(parts, msg)
	}
	if d.Text != "" {
		parts = append(parts, d.Text)
	}
	if len(d.Params) > 0 {
		var sb strings.Builder
		sb.WriteString("**Parameters**\n")
		for _, p := range d.Params {
			sb.WriteString("- `" + p.Name + "`")
			if p.Text != "" {
				sb.WriteString(": " + oneLine(p.Text))
			}
			sb.WriteString("\n")
		}
		parts = append(parts, strings.TrimRight(sb.String(), "\n"))
	}
	if d.Returns != "" {
		parts = append(parts, "**Returns** "+d.Returns)
	}
	for _, ex := range d.Examples {
		if !strings.Contains(ex, "```") {
			ex = "```mms\n" + ex + "\n```"
		}
		parts = append(parts, "**Example**\n"+ex)
	}
	for _, s := range d.See {
		parts = append(parts, "**See** "+s)
	}
	for _, t := range d.Other {
		line := "*@" + t.Name + "*"
		if t.Text != "" {
			line += " " + t.Text
		}
		parts = append(parts, line)
	}
	return strings.Join(parts, "\n\n")
}

func oneLine(s string) string { return strings.Join(strings.Fields(s), " ") }
