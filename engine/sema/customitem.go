package sema

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// Custom items: one declaration for a vanilla item with its own look and
// name, which otherwise takes a Model, an ItemModel, a Lang entry and the
// same item_model and item_name components repeated at every use.
//
//	CustomItem ember_shard = {
//	  base = amethyst_shard
//	  name = "Ember Shard"
//	  texture = Texture("art/ember_shard.png")
//	}
//
// It writes models/item/ember_shard.json (a flat item/generated model),
// items/ember_shard.json and an item.<ns>.ember_shard translation merged
// into the language's lang file. Its name then stands for the base item
// with those components: in item stacks (recipe results), in loot entries
// and in ${...} in commands.

const customItemKind = "CustomItem"

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

// CustomItemGenerates reports which files a custom item declaration
// writes, from its syntax alone so that the symbol table can reserve
// their names: the ItemModel unless item_model names one, and the Model
// when it is drawn from a texture. An item with none of texture, model
// and item_model keeps its base item's look and writes neither.
func CustomItemGenerates(d *syntax.Decl) (itemModel, model bool) {
	fields := customItemFields(d)
	im, hasIM := fields["item_model"]
	_, hasModel := fields["model"]
	_, hasTexture := fields["texture"]
	if hasIM {
		_, isObj := im.(*syntax.Object)
		return isObj, false
	}
	return hasModel || hasTexture, !hasModel && hasTexture
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

var customTypes sync.Map // *spec.Pack -> *spec.Type

// customItemType is the struct a CustomItem body is checked against,
// built from the pack's own types so that ids, textures and components
// are checked, completed and documented as they are elsewhere.
func customItemType(p *spec.Pack) *spec.Type {
	if t, ok := customTypes.Load(p); ok {
		return t.(*spec.Type)
	}
	var itemID, parent, texture *spec.Type
	if st := p.Resolve(p.Types[tItemStack]); st != nil {
		for _, m := range st.Members {
			if idOf(m).Present {
				itemID = m
			}
		}
	}
	if m := p.Resolve(p.Types["::java::assets::model::Model"]); m != nil {
		for _, f := range m.Fields {
			switch f.Key {
			case "parent":
				parent = f.Type
			case "textures":
				if tt := p.Resolve(f.Type); tt != nil {
					for _, g := range tt.Fields {
						if g.KeyType != nil {
							texture = g.Type
						}
					}
				}
			}
		}
	}
	if itemID == nil {
		itemID = idType("item")
	}
	if parent == nil {
		parent = idType("model")
	}
	if texture == nil {
		texture = idType("texture")
	}
	str := &spec.Type{Kind: "string"}
	t := &spec.Type{Kind: "struct", Fields: []*spec.Field{
		{Kind: "pair", Key: "base", Type: itemID,
			Desc: "The vanilla item this one is made from. It keeps the base item's behaviour; recipes and item tests that ask for an item id match the base."},
		{Kind: "pair", Key: "name", Optional: true, Type: &spec.Type{Kind: "union", Members: []*spec.Type{
			str,
			{Kind: "struct", Fields: []*spec.Field{{Kind: "pair", KeyType: str, Type: str}}},
		}}, Desc: "The item's name: text for `en_us`, or one text per language such as `{ en_us = \"Ember Shard\", de_de = \"Glutsplitter\" }`. It becomes an `item.<namespace>.<name>` entry in the lang files and the item's `item_name` component."},
		{Kind: "pair", Key: "texture", Optional: true, Type: texture,
			Desc: "The item's picture, usually `Texture(\"art/name.png\")`, which is copied to textures/item/<name>.png. It makes a flat item model, like most vanilla items."},
		{Kind: "pair", Key: "parent", Optional: true, Type: parent,
			Desc: "The parent of the generated model: `item/generated` (the default) for flat items, `item/handheld` for tools held like a sword."},
		{Kind: "pair", Key: "model", Optional: true, Type: parent,
			Desc: "An existing model to show instead of generating one from `texture`, such as a `Model` declared in the project."},
		{Kind: "pair", Key: "item_model", Optional: true, Type: &spec.Type{Kind: "union", Members: []*spec.Type{
			idType("item_definition"),
			{Kind: "reference", Path: "::java::assets::item_definition::ItemDefinition"},
		}}, Desc: "An item model definition to use instead of the generated one: the id of an `ItemModel`, or its body, such as `{ model = Select(...) }` for an item that changes its look."},
		{Kind: "pair", Key: "components", Optional: true, Type: &spec.Type{Kind: "reference", Path: "::java::world::component::DataComponentPatch"},
			Desc: "More item components, such as `max_stack_size` or `rarity`. Every use of the item gets them."},
	}}
	got, _ := customTypes.LoadOrStore(p, t)
	return got.(*spec.Type)
}

func idType(registry string) *spec.Type {
	lit, _ := json.Marshal(registry)
	return &spec.Type{Kind: "string", Attributes: []spec.Attribute{{Name: "id", Value: &spec.Type{Kind: "literal", Literal: &spec.Literal{Kind: "string", Value: lit}}}}}
}

// CustomItem checks a CustomItem declaration. id is its "ns:path".
func (c *Checker) CustomItem(d *syntax.Decl, id string) *CustomItemOut {
	const usage = "CustomItem needs a body: CustomItem ember_shard = { base = amethyst_shard, name = \"Ember Shard\", texture = Texture(\"art/ember_shard.png\") }"
	body := customItemBody(d)
	if body == nil || d.Name == nil {
		sp := d.Kind.Sp
		if d.Value != nil {
			sp = d.Value.Span()
			c.errorf("mms/type-mismatch", sp, "%s", usage)
			return nil
		}
		c.errorf("mms/missing-value", sp, "%s", usage)
		return nil
	}
	ns, path := splitID(id)
	prevTex, prevOf := c.textureName, c.customOf
	c.textureName, c.customOf = "item/"+path, id
	v := c.check(&syntax.Object{Body: body, Sp: body.Sp}, tview{T: customItemType(c.Pack), Name: customItemKind})
	c.textureName, c.customOf = prevTex, prevOf
	o, ok := v.(*jsonv.Object)
	if !ok {
		return nil
	}
	fields := customItemFields(d)
	errAt := func(key, format string, args ...any) {
		sp := d.Name.Sp
		if e := fields[key]; e != nil {
			sp = e.Span()
		}
		c.errorf("mms/custom-item", sp, format, args...)
	}
	str := func(key string) string {
		s, _ := o.Get(key).(jsonv.String)
		return string(s)
	}
	out := &CustomItemOut{ModelPath: "item/" + path}

	texture, model, parent, itemModel := o.Get("texture"), str("model"), str("parent"), o.Get("item_model")
	switch {
	case itemModel != nil && (texture != nil || model != "" || parent != ""):
		for _, k := range []string{"texture", "model", "parent"} {
			if fields[k] != nil {
				errAt(k, "item_model replaces the generated models, so '%s' has no effect; leave out one or the other", k)
				break
			}
		}
	case model != "" && (texture != nil || parent != ""):
		k := "texture"
		if texture == nil {
			k = "parent"
		}
		errAt(k, "model names the model to show, so '%s' has no effect; leave out one or the other", k)
	case parent != "" && texture == nil:
		errAt("parent", "parent is the parent of the model drawn from texture; add a texture, or leave it out")
	}

	itemModelID := ""
	if itemModel != nil || model != "" || texture != nil {
		itemModelID = id
	}
	switch im := itemModel.(type) {
	case jsonv.String:
		itemModelID = string(im)
	case *jsonv.Object:
		out.ItemModel = im
	default:
		modelID := model
		if modelID == "" && texture != nil {
			if parent == "" {
				parent = "minecraft:item/generated"
			}
			m := &jsonv.Object{}
			m.Set("parent", jsonv.String(parent))
			tex := &jsonv.Object{}
			tex.Set("layer0", texture)
			m.Set("textures", tex)
			out.Model = m
			modelID = ns + ":" + out.ModelPath
		}
		if modelID != "" {
			ref := &jsonv.Object{}
			ref.Set("type", jsonv.String("minecraft:model"))
			ref.Set("model", jsonv.String(modelID))
			def := &jsonv.Object{}
			def.Set("model", ref)
			out.ItemModel = def
		}
	}

	switch n := o.Get("name").(type) {
	case jsonv.String:
		out.Names = map[string]string{"en_us": string(n)}
	case *jsonv.Object:
		out.Names = map[string]string{}
		for _, m := range n.Members {
			if s, ok := m.Value.(jsonv.String); ok {
				if !langCode.MatchString(m.Key) {
					c.errorf("mms/custom-item", fields["name"].Span(), "'%s' is not a language code such as en_us or de_de", m.Key)
					continue
				}
				out.Names[m.Key] = string(s)
			}
		}
	}
	if len(out.Names) > 0 {
		out.LangKey = CustomItemLangKey(id)
	}

	user, _ := o.Get("components").(*jsonv.Object)
	comps := &jsonv.Object{}
	if itemModelID != "" && !hasComponent(user, "item_model") {
		comps.Set("item_model", jsonv.String(itemModelID))
	}
	if out.LangKey != "" && !hasComponent(user, "item_name") {
		// The fallback shows when the resource pack is not loaded.
		name := &jsonv.Object{}
		name.Set("translate", jsonv.String(out.LangKey))
		fallback, ok := out.Names["en_us"]
		if !ok {
			fallback = out.Names[out.Languages()[0]]
		}
		name.Set("fallback", jsonv.String(fallback))
		comps.Set("item_name", name)
	}
	if user != nil {
		comps.Members = append(comps.Members, user.Members...)
	}
	stack := &jsonv.Object{}
	if b := str("base"); b != "" {
		stack.Set("id", jsonv.String(b))
	}
	stack.Set("components", comps)
	out.Stack = stack
	return out
}

var langCode = regexp.MustCompile(`^[a-z]{2,3}_[a-z0-9]{2,3}$`)

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
	if s := c.Syms.Find("item", ns+":"+path); s == nil || s.Kind.MMS != customItemKind {
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
	sub.Files, sub.Deps, sub.parent, sub.customOf = c.Files, nil, c, s.ID()
	if s.Unit != nil && s.Unit.File != nil {
		sub.Src = s.Unit.File.Src
	}
	out := sub.CustomItem(s.Decl, s.ID())
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
		if s == nil || s.Kind.MMS != customItemKind {
			return nil, nil
		}
		s = c.customItemRef(ref)
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
		return &syntax.ResLoc{Text: text, Sp: ref.Span()}, s
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
	return s != nil && s.Kind.MMS == customItemKind
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

func quoteSNBT(s string) string {
	return `"` + strings.NewReplacer(`\`, `\\`, `"`, `\"`).Replace(s) + `"`
}
