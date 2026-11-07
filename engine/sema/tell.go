package sema

import (
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/prelude"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// The text commands of a function body: tell, title and actionbar, from
// the prelude. Their message is text markup, written into the command as
// a text component.

// ShownText is a message a text command shows, for an editor's chat
// preview: Kind is "chat", "title" or "actionbar". Sp is the command, or
// the whole of File when the message is read from that file.
type ShownText struct {
	Kind              string
	Message, Subtitle jsonv.Value
	File              string
	Sp                syntax.Span
}

// textCommands are the prelude's text commands, by name.
var textCommands = map[string]bool{"tell": true, "title": true, "actionbar": true}

// textCommandDecls are the prelude declarations of the text commands.
func textCommandDecls() []*syntax.Decl {
	var out []*syntax.Decl
	for _, f := range prelude.Parse() {
		for _, d := range f.File.Decls {
			if d.Kind != nil && d.Kind.Name == "Template" && textCommands[d.NameText()] {
				out = append(out, d)
			}
		}
	}
	return out
}

// textCommandHead is a text command's signature, as its hover shows it.
func textCommandHead(d *syntax.Decl) string {
	var ps []string
	for _, p := range d.Params.Params {
		ps = append(ps, p.Name.Text+": "+syntax.TypeText(p.Type))
	}
	return d.NameText() + "(" + strings.Join(ps, ", ") + ")"
}

// TextCommandCandidates lists the text commands where a statement goes in
// function code, with their prelude docs.
func TextCommandCandidates() []Candidate {
	var out []Candidate
	for _, d := range textCommandDecls() {
		name := d.NameText()
		out = append(out, Candidate{Label: name, Kind: CandCallable, Detail: textCommandHead(d),
			Doc: DocMarkdown(d.Doc), Insert: name + "(${1:@s}, \"$2\")", Rank: 40 + len(out)})
	}
	return out
}

// selectorText is a selector written as a command takes it, such as
// @e[type=zombie,limit=1], or "" when e is not a selector.
func selectorText(e syntax.Expr) string {
	switch e := e.(type) {
	case *syntax.Selector:
		return "@" + e.Name
	case *syntax.Index:
		if _, ok := e.X.(*syntax.Selector); ok {
			return compactSelector(format.Expr(e))
		}
	}
	return ""
}

// textCommand lowers tell(target, message), title(target, message,
// subtitle = ...) or actionbar(target, message).
func (f *fc) textCommand(pre *[]fn.Op, name string, e *syntax.Call) {
	c := f.c
	for _, d := range textCommandDecls() {
		if d.NameText() == name {
			c.info(e.Fn).Help = "```mms\n" + textCommandHead(d) + "\n```\n" + DocMarkdown(d.Doc)
		}
	}
	params := []string{"target", "message"}
	if name == "title" {
		params = append(params, "subtitle")
	}
	args := map[string]syntax.Expr{}
	pos := 0
	for _, a := range e.Args {
		key := ""
		if a.Name != nil {
			key = a.Name.Text
		} else if pos < len(params) {
			key = params[pos]
			pos++
		}
		known := false
		for _, p := range params {
			known = known || p == key
		}
		switch {
		case !known && a.Name != nil:
			c.errorf("mms/unknown-arg", a.Name.Sp, "%s has no parameter '%s'", name, key)
			return
		case !known:
			c.errorf("mms/too-many-args", a.Sp, "%s takes %d arguments", name, len(params))
			return
		case args[key] != nil:
			c.errorf("mms/duplicate-arg", a.Sp, "'%s' is given twice", key)
			return
		}
		args[key] = a.Value
	}
	for _, p := range params[:2] {
		if args[p] == nil {
			c.errorf("mms/missing-arg", e.ArgsSp, "%s needs a %s: %s", name, p, textCommandExample[name])
			return
		}
	}
	target := selectorText(args["target"])
	if target == "" {
		c.errorf("mms/type-mismatch", args["target"].Span(), "expected a selector such as @a or @s")
		return
	}
	message, file, n := f.textArg(pre, args["message"])
	if message == nil {
		return
	}
	shown := ShownText{Kind: name, Message: message, File: file, Sp: e.Sp}
	if name == "tell" {
		shown.Kind = "chat"
	}
	if file != "" {
		shown.Sp = syntax.Span{Start: 0, End: n}
	}
	switch name {
	case "tell":
		*pre = append(*pre, fn.Cmd{Text: "tellraw " + target + " " + componentSNBT(message)})
	case "actionbar":
		*pre = append(*pre, fn.Cmd{Text: "title " + target + " actionbar " + componentSNBT(message)})
	case "title":
		// The subtitle shows with the next title, so it goes first.
		if sub := args["subtitle"]; sub != nil && !isNone(sub) {
			st, _, _ := f.textArg(pre, sub)
			if st == nil {
				return
			}
			shown.Subtitle = st
			*pre = append(*pre, fn.Cmd{Text: "title " + target + " subtitle " + componentSNBT(st)})
		}
		*pre = append(*pre, fn.Cmd{Text: "title " + target + " title " + componentSNBT(message)})
	}
	c.Shown = append(c.Shown, shown)
}

var textCommandExample = map[string]string{
	"tell":      `tell(@a, "Hello")`,
	"title":     `title(@a, "Hello", subtitle = "and welcome")`,
	"actionbar": `actionbar(@s, "Hello")`,
}

func isNone(e syntax.Expr) bool {
	id, ok := e.(*syntax.Ident)
	return ok && id.Name == "none"
}

// textArg reads a text argument: a string of markup, or TextFile(...).
// For a project file read in the editor, it also returns the file and
// its length.
func (f *fc) textArg(pre *[]fn.Op, e syntax.Expr) (jsonv.Value, string, int) {
	c := f.c
	switch e := e.(type) {
	case *syntax.String:
		in := c.info(e)
		in.Markup = true
		if !strings.ContainsAny(e.Value, markupRunes) {
			return jsonv.String(e.Value), "", 0
		}
		src := c.stringSrc(e)
		src.noPages = true
		v, accidental := c.textMarkupIn(src, f, pre)
		in.Accidental = accidental
		return v, "", 0
	case *syntax.Call:
		if id, ok := e.Fn.(*syntax.Ident); ok && id.Name == "TextFile" {
			src := c.bookFile(e)
			if src == nil {
				return nil, "", 0
			}
			src.noPages = true
			v, _ := c.textMarkupIn(src, f, pre)
			file := src.extern
			if file == "" {
				file = src.file
			}
			return v, file, len(src.text)
		}
	}
	c.errorf("mms/type-mismatch", e.Span(), `expected text: a string of markup such as "{gold}Hello", or TextFile("path")`)
	return nil, "", 0
}

// componentSNBT writes a text component as a command takes it:
// {text:"Hi",color:"gold"}.
func componentSNBT(v jsonv.Value) string {
	var b strings.Builder
	var write func(v jsonv.Value)
	write = func(v jsonv.Value) {
		switch v := v.(type) {
		case *jsonv.Object:
			b.WriteByte('{')
			for i, m := range v.Members {
				if i > 0 {
					b.WriteByte(',')
				}
				if plainKey(m.Key) {
					b.WriteString(m.Key)
				} else {
					b.WriteString(fn.QuoteSNBT(m.Key))
				}
				b.WriteByte(':')
				write(m.Value)
			}
			b.WriteByte('}')
		case jsonv.Array:
			b.WriteByte('[')
			for i, x := range v {
				if i > 0 {
					b.WriteByte(',')
				}
				write(x)
			}
			b.WriteByte(']')
		case jsonv.String:
			b.WriteString(fn.QuoteSNBT(string(v)))
		case jsonv.Number:
			b.WriteString(string(v))
		case jsonv.Bool:
			b.WriteString(strconv.FormatBool(bool(v)))
		}
	}
	// ["", x] is x: one run needs no list.
	if a, ok := v.(jsonv.Array); ok && len(a) == 2 && a[0] == jsonv.String("") {
		v = a[1]
	}
	write(v)
	return b.String()
}

func plainKey(k string) bool {
	if k == "" {
		return false
	}
	for _, r := range k {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_') {
			return false
		}
	}
	return true
}

// TextCommandSignature is signature help for a text command: its label,
// parameters and doc, or "" when name is not one.
func TextCommandSignature(name string) (label string, params []string, doc string) {
	for _, d := range textCommandDecls() {
		if d.NameText() != name {
			continue
		}
		for _, p := range d.Params.Params {
			params = append(params, p.Name.Text+": "+syntax.TypeText(p.Type))
		}
		return textCommandHead(d), params, syntax.ParseDoc(d.Doc).Text
	}
	return "", nil, ""
}
