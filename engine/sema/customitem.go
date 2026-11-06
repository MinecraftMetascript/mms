package sema

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Item stacks a declaration's name stands for. CustomItem itself is a
// Blueprint in the prelude (prelude/items.mms), whose `as` conversions
// give its name a stack; what is here serves Book, whose name stands for
// a written book, and renders stacks as SNBT for commands.

// CustomItemOut is a checked custom item.
type CustomItemOut struct {
	// Model is the generated item model under models/, nil when the item
	// names its own model or item_model.
	Model jsonv.Value
	// ModelPath is Model's path under models/: "item/<name>".
	ModelPath string
	// ItemModel is the generated item model definition under items/, nil
	// when item_model names an existing one.
	ItemModel jsonv.Value
	// LangKey is the translation key the item_name component uses, and
	// Names its text by language; both are empty without a name.
	LangKey string
	Names   map[string]string
	// Stack is the item stack the name stands for: {id, components}.
	Stack *jsonv.Object
	// Overflows are a Book's pages that run past the bottom of the page.
	Overflows []BookOverflow
	// PageStarts are where each page of a Book's text starts, so an
	// editor can open its preview on the page at the cursor.
	PageStarts []BookPageStart
}

// Languages lists the item's languages, sorted.
func (o *CustomItemOut) Languages() []string {
	var out []string
	for l := range o.Names {
		out = append(out, l)
	}
	sort.Strings(out)
	return out
}

// CustomItemLangKey is the translation key a custom item's name gets.
func CustomItemLangKey(id string) string {
	ns, path := splitID(id)
	return "item." + ns + "." + strings.ReplaceAll(path, "/", ".")
}

func customItemBody(d *syntax.Decl) *syntax.Block {
	if d.Body != nil {
		return d.Body
	}
	if o, ok := d.Value.(*syntax.Object); ok {
		return o.Body
	}
	return nil
}

func customItemFields(d *syntax.Decl) map[string]syntax.Expr {
	out := map[string]syntax.Expr{}
	if b := customItemBody(d); b != nil {
		for _, it := range b.Items {
			if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil {
				out[f.Key.Text] = f.Value
			}
		}
	}
	return out
}

func idType(registry string) *spec.Type {
	lit, _ := json.Marshal(registry)
	return &spec.Type{Kind: "string", Attributes: []spec.Attribute{{Name: "id", Value: &spec.Type{Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: lit}}}}}
}

func hasComponent(o *jsonv.Object, name string) bool {
	if o == nil {
		return false
	}
	return o.Get(name) != nil || o.Get("minecraft:"+name) != nil
}

// customItemSig is what declarations that use a custom item depend on:
// its whole source text.
func customItemSig(s *Symbol) string {
	if s.Unit == nil || s.Unit.File == nil {
		return ""
	}
	src := s.Unit.File.Src
	if s.Decl.Sp.End > len(src) || s.Decl.Sp.Start < 0 {
		return ""
	}
	h := sha256.Sum256([]byte(string(src[s.Decl.Sp.Start:s.Decl.Sp.End])))
	return hex.EncodeToString(h[:8])
}

// customItemRef returns the custom item an id expression names, or nil.
// It records the lookup, and the item's text, for incremental checking.
func (c *Checker) customItemRef(e syntax.Expr) *Symbol {
	var text string
	switch x := e.(type) {
	case *syntax.Ident:
		text = x.Name
	case *syntax.ResLoc:
		text = x.Text
	default:
		return nil
	}
	ns, path := splitID(text)
	if ns == "" {
		ns = c.NS
	}
	if s := c.Syms.Find("item", ns+":"+path); s == nil || !isItemKind(s.Kind.MMS) {
		return nil
	}
	s := c.findSig("item", ns+":"+path)
	in := c.info(e)
	in.Symbol, in.ID = s, s.ID()
	in.Bare, in.Folder = BareID, "item"
	return s
}

// customStack returns the item stack a custom item stands for, checked
// in its own declaration's namespace. Problems in the custom item are
// reported on its declaration, not here.
func (c *Checker) customStack(s *Symbol, at syntax.Span) *jsonv.Object {
	for p := c; p != nil; p = p.parent {
		if p.customOf == s.ID() {
			c.errorf("mms/custom-item", at, "custom item '%s' contains itself", s.ID())
			return nil
		}
	}
	sub := NewChecker(c.Pack, c.Kinds, c.Syms, s.NS)
	// The use depends on everything the item's check looks up, such as
	// another custom item in its components.
	sub.Files, sub.Texts, sub.Deps, sub.parent, sub.customOf = c.Files, c.Texts, c.Deps, c, s.ID()
	if s.Unit != nil && s.Unit.File != nil {
		sub.Src = s.Unit.File.Src
	}
	out := sub.Book(s.Decl, s.ID())
	if out == nil {
		return nil
	}
	return out.Stack
}

// customStackComponents is a custom item's components.
func customStackComponents(st *jsonv.Object) *jsonv.Object {
	comps, _ := st.Get("components").(*jsonv.Object)
	if comps == nil {
		comps = &jsonv.Object{}
	}
	return comps
}

// withCustomField rewrites `{ ..., key = <custom item>, ... }` or
// `Call(<custom item>, ...)` / `Call(key = <custom item>)` so that the
// custom item becomes its base id. It returns the rewritten expression
// and the item, or nil when e does not use a custom item there.
func (c *Checker) withCustomField(e syntax.Expr, key string) (syntax.Expr, *Symbol) {
	base := func(ref syntax.Expr) (syntax.Expr, *Symbol) {
		s := c.Syms.Find("item", c.idText(ref))
		if s == nil || !isItemKind(s.Kind.MMS) {
			return nil, nil
		}
		s = c.customItemRef(ref)
		return &syntax.ResLoc{Text: customItemBase(s), Sp: ref.Span()}, s
	}
	switch x := e.(type) {
	case *syntax.Object:
		for i, it := range x.Body.Items {
			f, ok := it.(*syntax.FieldStmt)
			if !ok || f.Key == nil || f.Key.Text != key || !isID(f.Value) {
				continue
			}
			nb, s := base(f.Value)
			if s == nil {
				return nil, nil
			}
			items := append([]syntax.Stmt(nil), x.Body.Items...)
			cp := *f
			cp.Value = nb
			items[i] = &cp
			blk := *x.Body
			blk.Items = items
			return &syntax.Object{Body: &blk, Sp: x.Sp}, s
		}
	case *syntax.Call:
		for i, a := range x.Args {
			if (a.Name == nil && i != 0) || (a.Name != nil && a.Name.Text != key) || !isID(a.Value) {
				continue
			}
			nb, s := base(a.Value)
			if s == nil {
				return nil, nil
			}
			args := append([]*syntax.Arg(nil), x.Args...)
			cp := *a
			cp.Value = nb
			args[i] = &cp
			call := *x
			call.Args = args
			return &call, s
		}
	}
	return nil, nil
}

// customItemBase is the namespaced id of a custom item's base item.
// Bases are vanilla items, so an unqualified one is in minecraft.
func customItemBase(s *Symbol) string {
	if s.Kind.MMS == bookKind {
		return "minecraft:written_book"
	}
	id, _ := customItemFields(s.Decl)["base"].(syntax.Expr)
	text := "minecraft:air"
	switch b := id.(type) {
	case *syntax.Ident:
		text = b.Name
	case *syntax.ResLoc:
		text = b.Text
	}
	if !strings.Contains(text, ":") {
		text = "minecraft:" + text
	}
	return text
}

// idText is the namespaced id an Ident or ResLoc names in this
// namespace.
func (c *Checker) idText(e syntax.Expr) string {
	var text string
	switch x := e.(type) {
	case *syntax.Ident:
		text = x.Name
	case *syntax.ResLoc:
		text = x.Text
	}
	if !strings.Contains(text, ":") {
		text = c.NS + ":" + text
	}
	return text
}

// isCustomItem reports whether an id expression names a custom item.
// The lookup is recorded, so that adding the item later checks e again.
func (c *Checker) isCustomItem(e syntax.Expr) bool {
	if !isID(e) {
		return false
	}
	s := c.find("item", c.idText(e))
	return s != nil && isItemKind(s.Kind.MMS)
}

// customStackRule lowers a custom item in an item stack: `ember_shard`,
// `ember_shard * 4` or `{ id = ember_shard, count = 4 }`. The stack gets
// the base id and the custom item's components; components written at
// the use come after them, so they win.
func customStackRule(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	if isID(e) {
		return c.check(mkObject(e.Span(), "id", e), v)
	}
	ne, s := c.withCustomField(e, "id")
	if s == nil {
		return nil
	}
	st := c.customStack(s, e.Span())
	out := c.check(ne, v)
	o, ok := out.(*jsonv.Object)
	if !ok || st == nil {
		return out
	}
	comps := &jsonv.Object{Members: append([]jsonv.Member(nil), customStackComponents(st).Members...)}
	if user, ok := o.Get("components").(*jsonv.Object); ok {
		for _, m := range user.Members {
			comps.Set(m.Key, m.Value)
		}
	}
	o.Set("components", comps)
	return o
}

// customEntryRule lowers a custom item in a loot entry: `ember_shard`,
// `Item(ember_shard, ...)` or `Item(name = ember_shard, ...)`. The entry
// drops the base item, and a set_components modifier, first in the list
// so later modifiers can change it, adds the custom item's components.
func customEntryRule(c *Checker, e syntax.Expr, v tview) jsonv.Value {
	if isID(e) {
		return c.check(mkObject(e.Span(), "type", word("item", e.Span()), "name", e), v)
	}
	ne, s := c.withCustomField(e, "name")
	if s == nil {
		return nil
	}
	st := c.customStack(s, e.Span())
	out := c.check(ne, v)
	o, ok := out.(*jsonv.Object)
	if !ok || st == nil {
		return out
	}
	set := &jsonv.Object{}
	set.Set("type", jsonv.String("minecraft:set_components"))
	set.Set("components", customStackComponents(st))
	mods, _ := o.Get("modifier").(jsonv.Array)
	o.Set("modifier", append(jsonv.Array{set}, mods...))
	return o
}

// usesCustomField reports whether e is an object or call whose key
// field names a custom item.
func (c *Checker) usesCustomField(e syntax.Expr, key string) bool {
	switch x := e.(type) {
	case *syntax.Object:
		for _, it := range x.Body.Items {
			if f, ok := it.(*syntax.FieldStmt); ok && f.Key != nil && f.Key.Text == key {
				return c.isCustomItem(f.Value)
			}
		}
	case *syntax.Call:
		if id, ok := x.Fn.(*syntax.Ident); !ok || id.Name != "Item" {
			return false
		}
		for i, a := range x.Args {
			if (a.Name == nil && i == 0) || (a.Name != nil && a.Name.Text == key) {
				return c.isCustomItem(a.Value)
			}
		}
	}
	return false
}

// StackSNBT writes an item stack as a command's item argument:
// id[component=value,...].
func StackSNBT(st *jsonv.Object) string {
	id, _ := st.Get("id").(jsonv.String)
	var sb strings.Builder
	sb.WriteString(strings.TrimPrefix(string(id), "minecraft:"))
	if comps, ok := st.Get("components").(*jsonv.Object); ok && len(comps.Members) > 0 {
		sb.WriteByte('[')
		for i, m := range comps.Members {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(m.Key)
			sb.WriteByte('=')
			writeSNBT(&sb, m.Value)
		}
		sb.WriteByte(']')
	}
	return sb.String()
}

var bareSNBTKey = regexp.MustCompile(`^[A-Za-z0-9_\-.+]+$`)

func writeSNBT(sb *strings.Builder, v jsonv.Value) {
	switch x := v.(type) {
	case *jsonv.Object:
		sb.WriteByte('{')
		for i, m := range x.Members {
			if i > 0 {
				sb.WriteByte(',')
			}
			if bareSNBTKey.MatchString(m.Key) {
				sb.WriteString(m.Key)
			} else {
				sb.WriteString(quoteSNBT(m.Key))
			}
			sb.WriteByte(':')
			writeSNBT(sb, m.Value)
		}
		sb.WriteByte('}')
	case jsonv.Array:
		sb.WriteByte('[')
		for i, e := range x {
			if i > 0 {
				sb.WriteByte(',')
			}
			writeSNBT(sb, e)
		}
		sb.WriteByte(']')
	case jsonv.String:
		sb.WriteString(quoteSNBT(string(x)))
	case jsonv.Number:
		sb.WriteString(string(x))
	case jsonv.Bool:
		if x {
			sb.WriteString("true")
		} else {
			sb.WriteString("false")
		}
	}
}

// quoteSNBT quotes a string for SNBT, escaping control characters so
// that a newline in a name cannot split the command line.
func quoteSNBT(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch {
		case r == '\\':
			b.WriteString(`\\`)
		case r == '"':
			b.WriteString(`\"`)
		case r == '\n':
			b.WriteString(`\n`)
		case r == '\t':
			b.WriteString(`\t`)
		case r < 0x20 || (r >= 0x7f && r <= 0x9f):
			fmt.Fprintf(&b, `\u%04x`, r)
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
	return b.String()
}
