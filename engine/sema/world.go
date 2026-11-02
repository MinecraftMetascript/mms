package sema

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// World access from function code: conditions on the world
// (`Block(~ ~-1 ~, stone)`, `Entity(@a[distance=..8])`), NBT reads and
// writes (`@s::Health`, `siege:path::queue[0].x`), commands as values
// (`let n = /execute if entity @a`) and the small built-in functions
// (Count, Abs, Min, Max). Everything here lowers to one or two commands;
// nothing hides a loop.

// condHelpers are the execute conditions that function code can use as
// bool values and as execute blocks. They share the helper machinery
// with As/At: the arguments are written into the subcommand in the
// order the command takes them, and the chain is checked against the
// command tree.
var condHelpers = []*ExecHelper{
	{Name: "Block", Sub: "if block", Params: "pos, block", min: 2, cond: true, Insert: "Block(${1:~ ~ ~}, ${2:stone})",
		Doc: "Holds when the block at pos is the given block, block state or #tag (`execute if block`): `Block(~ ~-1 ~, stone)`, `Block(~ ~ ~, #minecraft:logs)`."},
	{Name: "Blocks", Sub: "if blocks", Params: "start, end, destination, all | masked", min: 4, cond: true, Insert: "Blocks(${1:~ ~ ~}, ${2:~ ~ ~}, ${3:~ ~ ~}, ${4:all})",
		Doc: "Holds when the blocks from start to end match those at destination (`execute if blocks`). `masked` ignores air in the source."},
	{Name: "Entity", Sub: "if entity", Params: "targets", min: 1, cond: true, Insert: "Entity(${1:@a})",
		Doc: "Holds when the selector matches at least one entity (`execute if entity`). `Count(@a)` gives the number."},
	{Name: "Biome", Sub: "if biome", Params: "pos, biome", min: 2, cond: true, Insert: "Biome(${1:~ ~ ~}, ${2:plains})",
		Doc: "Holds when the biome at pos is the given biome or #tag (`execute if biome`)."},
	{Name: "Dimension", Sub: "if dimension", Params: "dimension", min: 1, cond: true, Insert: "Dimension(${1:overworld})",
		Doc: "Holds in the given dimension (`execute if dimension`)."},
	{Name: "Loaded", Sub: "if loaded", Params: "pos", min: 1, cond: true, Insert: "Loaded(${1:~ ~ ~})",
		Doc: "Holds when the chunk at pos is loaded (`execute if loaded`)."},
	{Name: "Data", Sub: "if data", Params: "target::path | target, path", min: 1, cond: true, nsArg: true, Insert: "Data(${1:@s::Pos})",
		Doc: "Holds when the NBT path exists (`execute if data`): `Data(@s::Inventory[0])`, `Data(siege:path, queue[0])`, `Data(~ ~-1 ~, Items)`. The target is an entity, a storage or a block, told apart by how it is written.",
		render: func(args []execArg) string {
			if len(args) == 1 {
				return "if data " + args[0].text
			}
			return "if data " + args[0].kind() + " " + joinArgs(args)
		}},
	{Name: "Items", Sub: "if items", Params: "target, slot, item", min: 3, cond: true, Insert: "Items(${1:@s}, ${2:weapon.mainhand}, ${3:diamond_sword})",
		Doc: "Holds when the slot of an entity or a block holds the item (`execute if items`): `Items(@s, weapon.mainhand, diamond_sword)`, `Items(~ ~ ~, \"container.*\", #logs)`.",
		render: func(args []execArg) string {
			return "if items " + args[0].kind() + " " + joinArgs(args)
		}},
	{Name: "Predicate", Sub: "if predicate", Params: "predicate", min: 1, cond: true, nsArg: true, Insert: "Predicate(${1:name})",
		Doc: "Holds when the predicate passes (`execute if predicate`). A bare name is one of this project's predicates."},
	{Name: "Function", Sub: "if function", Params: "function", min: 1, cond: true, nsArg: true, Insert: "Function(${1:name})",
		Doc: "Runs the function and holds when it returns a value other than 0 (`execute if function`). A bare name is one of this project's functions."},
}

func init() {
	ExecHelpers = append(ExecHelpers, condHelpers...)
	for _, h := range condHelpers {
		execByName[h.Name] = h
	}
}

// kind is the data target kind an argument names: an entity for a
// selector, a block for coordinates, otherwise a storage.
func (a execArg) kind() string {
	switch {
	case a.selector:
		return "entity"
	case a.coord:
		return "block"
	}
	return "storage"
}

// condExpr lowers an execute chain used as a value, such as
// `Entity(@a)` or `At(@p).Block(~ ~-1 ~, stone)`, to a bool. A chain
// that only tests is the condition itself; one that also changes the
// context (as, at, positioned...) is run once into a flag, so the
// statement it guards keeps its own context.
func (f *fc) condExpr(pre *[]fn.Op, e syntax.Expr) *val {
	if f.src == nil {
		f.c.errorf("mms/not-constant", e.Span(), "a global's initial value must be a constant")
		return nil
	}
	parts, ok := f.execParts(e)
	if !ok {
		return nil
	}
	steps, _ := execChain(e)
	last := steps[len(steps)-1]
	lastPart := parts[len(parts)-1]
	if h := execByName[last.text()]; (h == nil || !h.cond) && !strings.HasPrefix(lastPart, "if ") && !strings.HasPrefix(lastPart, "unless ") {
		f.c.errorf("mms/type-mismatch", last.nameSpan(), "a condition ends with a test, such as Entity(@a) or Block(~ ~-1 ~, stone); %s only changes where the test runs", last.text())
		return nil
	}
	allTests := true
	for _, p := range parts {
		if !strings.HasPrefix(p, "if ") && !strings.HasPrefix(p, "unless ") {
			allTests = false
		}
	}
	if allTests {
		return &val{typ: tBool, cond: &fn.Cond{Parts: parts}}
	}
	t := f.temp()
	*pre = append(*pre,
		fn.Set{Dst: t, Src: fn.Const(0)},
		fn.Exec{Parts: parts, Body: &fn.Block{Ops: []fn.Op{fn.Set{Dst: t, Src: fn.Const(1)}}}})
	return &val{typ: tBool, cond: &fn.Cond{Parts: []string{fmt.Sprintf("if score %s matches 1", t)}}}
}

// nbtRef is a resolved `target::path`: an NBT location on an entity, in
// a storage or in a block.
type nbtRef struct {
	kind   string // "entity", "storage" or "block"
	target string
	path   string
	sp     syntax.Span
}

func (r *nbtRef) String() string { return r.kind + " " + r.target + " " + r.path }

// nbtRef resolves `@s::Pos[1]`, `siege:path::queue[0].x` or
// `~ ~-1 ~::Items`. isNBT is false when e is not an NBT path at all; a
// malformed path is reported and returns a nil ref with isNBT true.
func (f *fc) nbtRef(e syntax.Expr) (ref *nbtRef, isNBT bool) {
	var steps []syntax.Expr
	x := e
	for {
		switch n := x.(type) {
		case *syntax.Member:
			steps = append(steps, n)
			x = n.X
			continue
		case *syntax.Index:
			steps = append(steps, n)
			x = n.X
			continue
		case *syntax.Paren:
			x = n.X
			continue
		case *syntax.NBT:
			return f.resolveNBT(n, steps, e.Span()), true
		}
		return nil, false
	}
}

func (f *fc) resolveNBT(n *syntax.NBT, steps []syntax.Expr, sp syntax.Span) *nbtRef {
	c := f.c
	if f.src == nil {
		c.errorf("mms/not-constant", sp, "a global's initial value must be a constant")
		return nil
	}
	if n.Key == nil {
		c.errorf("mms/missing-name", n.Sp, "expected an NBT path after '::'")
		return nil
	}
	ref := &nbtRef{sp: sp}
	switch t := n.X.(type) {
	case *syntax.Selector:
		ref.kind, ref.target = "entity", "@"+t.Name
	case *syntax.Index:
		if _, ok := t.X.(*syntax.Selector); !ok {
			c.errorf("mms/type-mismatch", t.Sp, "an NBT path starts with a selector (@s::Pos), a storage (siege:path::x) or a block position (~ ~-1 ~::Items)")
			return nil
		}
		ref.kind, ref.target = "entity", compactSelector(string(f.src[t.Sp.Start:t.Sp.End]))
	case *syntax.ResLoc:
		ref.kind, ref.target = "storage", t.Text
		if !strings.Contains(t.Text, ":") {
			ref.target = f.ns + ":" + t.Text
		}
	case *syntax.Ident:
		if _, isVar := f.lookup(t.Name); isVar || f.c.find("Global", f.ns+":"+t.Name) != nil {
			c.errorf("mms/type-mismatch", t.Sp, "'%s' is a variable, which has no NBT; a bare name before :: is a storage in this namespace", t.Name)
			return nil
		}
		ref.kind, ref.target = "storage", f.ns+":"+t.Name
	case *syntax.Coord:
		ref.kind, ref.target = "block", strings.Join(strings.Fields(t.Text), " ")
	default:
		c.errorf("mms/type-mismatch", n.X.Span(), "an NBT path starts with a selector (@s::Pos), a storage (siege:path::x) or a block position (~ ~-1 ~::Items)")
		return nil
	}
	// A quoted key is the path as written, so paths the expression
	// syntax cannot spell, such as Inventory[{Slot:0b}].id, still work.
	ref.path = n.Key.Text
	for i := len(steps) - 1; i >= 0; i-- {
		switch s := steps[i].(type) {
		case *syntax.Member:
			if s.Name == nil {
				c.errorf("mms/missing-name", s.Sp, "expected a name after '.'")
				return nil
			}
			ref.path += "." + s.Name.Text
		case *syntax.Index:
			if len(s.Args) != 1 || s.Args[0].Name != nil {
				c.errorf("mms/unsupported", s.Sp, "an NBT list index is one number: Pos[1]")
				return nil
			}
			idx, ok := constIndex(s.Args[0].Value)
			if !ok {
				c.errorf("mms/unsupported", s.Args[0].Sp, "an NBT list index is a whole number, such as [0] or [-1]; for other selectors write the path in quotes: @s::\"Inventory[{Slot:0b}].id\"")
				return nil
			}
			ref.path += "[" + idx + "]"
		}
	}
	errs := len(c.Diags)
	c.commandText("data get "+ref.String(), func(start, end int) syntax.Span { return sp })
	if hasErrorsSince(c.Diags, errs) {
		return nil
	}
	return ref
}

// constIndex reads a literal list index: 1 or -1.
func constIndex(e syntax.Expr) (string, bool) {
	switch x := e.(type) {
	case *syntax.Number:
		if !strings.ContainsAny(x.Text, ".eE") {
			return x.Text, true
		}
	case *syntax.Unary:
		if n, ok := x.X.(*syntax.Number); ok && x.Op == "-" && !strings.ContainsAny(n.Text, ".eE") {
			return "-" + n.Text, true
		}
	}
	return "", false
}

// nbtRead lowers an NBT read to an int: `data get`, which floors
// decimals. scale, when not empty, multiplies first.
func (f *fc) nbtRead(pre *[]fn.Op, ref *nbtRef, scale string) *val {
	t := f.temp()
	cmd := "data get " + ref.String()
	if scale != "" {
		cmd += " " + scale
	}
	*pre = append(*pre, fn.Store{Dst: t, Cmd: cmd})
	return &val{typ: tInt, op: fn.Of(t)}
}

// scaledRead folds `@s::Pos[1] * 100` into the read's scale, so the
// decimals survive: data get floors after scaling.
func (f *fc) scaledRead(pre *[]fn.Op, e *syntax.Binary) (*val, bool) {
	if e.Op != "*" {
		return nil, false
	}
	for _, side := range [][2]syntax.Expr{{e.L, e.R}, {e.R, e.L}} {
		n, ok := side[1].(*syntax.Number)
		if !ok {
			continue
		}
		ref, isNBT := f.nbtRef(side[0])
		if !isNBT {
			continue
		}
		if ref == nil {
			return nil, true
		}
		if _, err := strconv.ParseInt(n.Text, 10, 32); err != nil {
			f.c.errorf("mms/type-mismatch", n.Sp, "scores are whole numbers from -2147483648 to 2147483647")
			return nil, true
		}
		return f.nbtRead(pre, ref, n.Text), true
	}
	return nil, false
}

// nbtTypes are the typed casts allowed on the right of an NBT assignment,
// with the SNBT suffix of a literal and the `execute store` type.
var nbtTypes = map[string]struct{ suffix, store string }{
	"Byte":   {"b", "byte"},
	"Short":  {"s", "short"},
	"Int":    {"", "int"},
	"Long":   {"L", "long"},
	"Float":  {"f", "float"},
	"Double": {"d", "double"},
}

// nbtAssign lowers `target::path = value` and `target::path += value`.
// Ints and bools are written as int (a bool as a byte) unless the value
// is wrapped in a cast: `@s::Health = Float(20)`.
func (f *fc) nbtAssign(out *fn.Block, ref *nbtRef, s *syntax.AssignStmt) {
	c := f.c
	if s.Value == nil {
		return
	}
	value := s.Value
	typ := ""
	if call, ok := value.(*syntax.Call); ok {
		if id, ok := call.Fn.(*syntax.Ident); ok && nbtTypes[id.Name].store != "" {
			if len(call.Args) != 1 || call.Args[0].Name != nil || call.Args[0].Value == nil {
				c.errorf("mms/arity", call.ArgsSp, "%s takes one value", id.Name)
				return
			}
			c.info(id).Help = "```mms\n" + id.Name + "(value)\n```\nWrites the value to NBT as a " + nbtTypes[id.Name].store + "."
			typ = id.Name
			value = call.Args[0].Value
		}
	}
	if s.Op != "=" {
		if typ != "" {
			c.errorf("mms/unsupported", s.Value.Span(), "%s only goes with =; write %s = %s(%s %s ...)", typ, "path", typ, "path", strings.TrimSuffix(s.Op, "="))
			return
		}
		// Read, compute, write back.
		cur := f.nbtRead(&out.Ops, ref, "")
		rv := f.intExpr(&out.Ops, s.Value)
		if rv == nil {
			return
		}
		t := cur.op.S
		out.Add(fn.Arith{Dst: t, Op: s.Op, Src: *rv})
		out.Add(fn.Cmd{Text: fmt.Sprintf("execute store result %s int 1 run scoreboard players get %s", ref.String(), t)})
		return
	}
	switch v := value.(type) {
	case *syntax.List, *syntax.Object:
		c.errorf("mms/unsupported", v.Span(), "lists and compounds are not written from function code yet; use /data modify for now")
		return
	}
	if src, isNBT := f.nbtRef(value); isNBT {
		if src == nil {
			return
		}
		if typ != "" {
			c.errorf("mms/unsupported", value.Span(), "a copy between NBT paths keeps the source's type; drop the %s", typ)
			return
		}
		out.Add(fn.Cmd{Text: "data modify " + ref.String() + " set from " + src.String()})
		return
	}
	v := f.expr(&out.Ops, value)
	if v == nil {
		return
	}
	nt := nbtTypes[typ]
	switch {
	case v.typ == tString:
		if typ != "" {
			c.errorf("mms/type-mismatch", value.Span(), "%s needs a number, but this is a string", typ)
			return
		}
		if v.str != nil {
			out.Add(fn.Cmd{Text: "data modify " + ref.String() + " set value " + fn.QuoteSNBT(*v.str)})
		} else {
			out.Add(fn.Cmd{Text: "data modify " + ref.String() + " set from storage " + v.path})
		}
		return
	case v.typ == tBool && typ == "":
		nt = nbtTypes["Byte"]
	case typ == "":
		nt = nbtTypes["Int"]
	}
	op := f.operand(&out.Ops, v)
	if op.IsConst {
		out.Add(fn.Cmd{Text: fmt.Sprintf("data modify %s set value %d%s", ref.String(), op.Const, nt.suffix)})
		return
	}
	out.Add(fn.Cmd{Text: fmt.Sprintf("execute store result %s %s 1 run scoreboard players get %s", ref.String(), nt.store, op.S)})
}

// cmdVal is a command used as a value. It runs when the value is needed,
// storing its result for an int or its success for a bool.
type cmdVal struct {
	text string
}

// commandExpr checks `/command` on the right of `=`.
func (f *fc) commandExpr(e *syntax.CommandExpr) *val {
	if strings.Contains(e.Text, "${") {
		f.c.errorf("mms/unsupported", e.Sp, "${...} is not supported in a command used as a value; put the value in a variable first")
		return nil
	}
	errs := len(f.c.Diags)
	f.c.Command(&syntax.CommandStmt{Text: e.Text, Sp: e.Sp})
	if hasErrorsSince(f.c.Diags, errs) {
		return nil
	}
	return &val{typ: tInt, cmd: &cmdVal{text: e.Text}}
}

// successCond runs a command value for its success.
func (f *fc) successCond(v *val) *fn.Cond {
	t := f.temp()
	return &fn.Cond{Pre: []fn.Op{fn.Store{Dst: t, Success: true, Cmd: v.cmd.text}}, Parts: []string{fmt.Sprintf("if score %s matches 1", t)}}
}

// asBool turns a command value into a bool where one is expected; other
// values come back unchanged.
func (f *fc) asBool(v *val) *val {
	if v == nil || v.cmd == nil {
		return v
	}
	return &val{typ: tBool, cond: f.successCond(v)}
}

// Built-in functions of function code.
var builtins = map[string]string{
	"Count": "Count(targets): int\n\nThe number of entities the selector matches (`execute if entity`).",
	"Abs":   "Abs(x): int\n\nThe absolute value of x.",
	"Min":   "Min(a, b, ...): int\n\nThe smallest of the values.",
	"Max":   "Max(a, b, ...): int\n\nThe largest of the values.",
}

// BuiltinCandidates lists the built-in functions, for completion.
func BuiltinCandidates() []Candidate {
	out := []Candidate{
		{Label: "Count", Kind: CandCallable, Detail: "int", Doc: builtins["Count"], Insert: "Count(${1:@a})", Rank: 40},
		{Label: "Abs", Kind: CandCallable, Detail: "int", Doc: builtins["Abs"], Insert: "Abs($1)", Rank: 41},
		{Label: "Min", Kind: CandCallable, Detail: "int", Doc: builtins["Min"], Insert: "Min($1)", Rank: 42},
		{Label: "Max", Kind: CandCallable, Detail: "int", Doc: builtins["Max"], Insert: "Max($1)", Rank: 43},
	}
	for i, h := range condHelpers {
		out = append(out, Candidate{Label: h.Name, Kind: CandCallable, Detail: "bool", Doc: h.help(), Insert: h.Insert, Rank: 20 + i})
	}
	return out
}

// isBuiltin reports whether a call names a built-in, which wins over a
// Function of the same name only when none is declared.
func (f *fc) isBuiltin(e *syntax.Call) (string, bool) {
	id, ok := e.Fn.(*syntax.Ident)
	if !ok {
		return "", false
	}
	if _, ok := builtins[id.Name]; !ok {
		if nbtTypes[id.Name].store != "" {
			return id.Name, true
		}
		return "", false
	}
	if f.c.find("function", f.ns+":"+id.Name) != nil {
		return "", false
	}
	return id.Name, true
}

// builtin lowers Count, Abs, Min and Max.
func (f *fc) builtin(pre *[]fn.Op, name string, e *syntax.Call) *val {
	c := f.c
	id := e.Fn.(*syntax.Ident)
	if nt, ok := nbtTypes[name]; ok {
		c.errorf("mms/misplaced", id.Sp, "%s writes a %s to NBT, so it goes on the right of an NBT assignment: @s::Health = %s(20)", name, nt.store, name)
		return nil
	}
	c.info(id).Help = "```mms\n" + strings.SplitN(builtins[name], "\n", 2)[0] + "\n```\n" + strings.SplitN(builtins[name], "\n\n", 2)[1]
	for _, a := range e.Args {
		if a.Name != nil {
			c.errorf("mms/unsupported", a.Name.Sp, "%s takes its arguments by position", name)
			return nil
		}
	}
	switch name {
	case "Count":
		if len(e.Args) != 1 || e.Args[0].Value == nil {
			c.errorf("mms/arity", e.ArgsSp, "Count takes one selector: Count(@a)")
			return nil
		}
		if f.src == nil {
			c.errorf("mms/not-constant", e.Sp, "a global's initial value must be a constant")
			return nil
		}
		sel := e.Args[0].Value
		switch x := sel.(type) {
		case *syntax.Selector:
		case *syntax.Index:
			if _, ok := x.X.(*syntax.Selector); !ok {
				c.errorf("mms/type-mismatch", sel.Span(), "Count takes a selector: Count(@e[tag=x])")
				return nil
			}
		default:
			c.errorf("mms/type-mismatch", sel.Span(), "Count takes a selector: Count(@e[tag=x])")
			return nil
		}
		sp := sel.Span()
		text := "execute if entity " + string(f.src[sp.Start:sp.End])
		errs := len(c.Diags)
		c.commandText(text, func(start, end int) syntax.Span {
			const base = len("execute if entity ")
			if start >= base {
				s := sp.Start + start - base
				en := sp.Start + end - base
				if en > sp.End {
					en = sp.End
				}
				if en <= s {
					en = s + 1
				}
				return syntax.Span{Start: s, End: en}
			}
			return sp
		})
		if hasErrorsSince(c.Diags, errs) {
			return nil
		}
		t := f.temp()
		*pre = append(*pre, fn.Store{Dst: t, Cmd: text})
		return &val{typ: tInt, op: fn.Of(t)}
	case "Abs":
		if len(e.Args) != 1 || e.Args[0].Value == nil {
			c.errorf("mms/arity", e.ArgsSp, "Abs takes one int: Abs(x)")
			return nil
		}
		v := f.intExpr(pre, e.Args[0].Value)
		if v == nil {
			return nil
		}
		if v.IsConst {
			if v.Const < 0 {
				return &val{typ: tInt, op: fn.Const(wrap(-int64(v.Const)))}
			}
			return &val{typ: tInt, op: *v}
		}
		t := f.temp()
		*pre = append(*pre, fn.Set{Dst: t, Src: *v},
			fn.If{Cond: fn.Cond{Parts: []string{fmt.Sprintf("if score %s matches ..-1", t)}}, Then: &fn.Block{Ops: []fn.Op{fn.Arith{Dst: t, Op: "*=", Src: fn.Const(-1)}}}})
		return &val{typ: tInt, op: fn.Of(t)}
	}
	// Min and Max.
	if len(e.Args) < 2 {
		c.errorf("mms/arity", e.ArgsSp, "%s takes two or more ints: %s(a, b)", name, name)
		return nil
	}
	var ops []fn.Operand
	for _, a := range e.Args {
		if a.Value == nil {
			return nil
		}
		v := f.intExpr(pre, a.Value)
		if v == nil {
			return nil
		}
		ops = append(ops, *v)
	}
	op := "<"
	if name == "Max" {
		op = ">"
	}
	// Fold the constants into one.
	var rest []fn.Operand
	var k *int32
	for _, o := range ops {
		if !o.IsConst {
			rest = append(rest, o)
			continue
		}
		if k == nil || (name == "Min" && o.Const < *k) || (name == "Max" && o.Const > *k) {
			v := o.Const
			k = &v
		}
	}
	if k != nil {
		rest = append(rest, fn.Const(*k))
	}
	if len(rest) == 1 {
		return &val{typ: tInt, op: rest[0]}
	}
	t := f.temp()
	*pre = append(*pre, fn.Set{Dst: t, Src: rest[0]})
	for _, o := range rest[1:] {
		*pre = append(*pre, fn.Arith{Dst: t, Op: op, Src: o})
	}
	return &val{typ: tInt, op: fn.Of(t)}
}

// unsetCompare rewrites a comparison of an objective score against a
// constant so that an unset score counts as 0: `execute if score`
// fails on an unset score, so a test whose accepted range holds 0 is
// written as the complement with `unless`. ok is false when the test
// is fine as it is.
func unsetCompare(s fn.Score, op string, v int64) (parts []string, always, ok bool) {
	unless := func(rng string) string { return fmt.Sprintf("unless score %s matches %s", s, rng) }
	switch op {
	case "==":
		if v == 0 {
			return []string{unless("..-1"), unless("1..")}, false, true
		}
	case "<":
		if v > 0 {
			return []string{unless(fmt.Sprintf("%d..", v))}, false, true
		}
	case "<=":
		if v == 1<<31-1 {
			return nil, true, true
		}
		if v >= 0 {
			return []string{unless(fmt.Sprintf("%d..", v+1))}, false, true
		}
	case ">":
		if v < 0 {
			return []string{unless(fmt.Sprintf("..%d", v))}, false, true
		}
	case ">=":
		if v == -1<<31 {
			return nil, true, true
		}
		if v <= 0 {
			return []string{unless(fmt.Sprintf("..%d", v-1))}, false, true
		}
	}
	return nil, false, false
}
