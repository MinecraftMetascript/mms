package project

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/minecraftmetascript/mms/engine/fn"
	"github.com/minecraftmetascript/mms/engine/internal/mcfn"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// TestResult is the outcome of one Test.
type TestResult struct {
	Name string // as declared
	ID   string // the function it compiles to
	Unit string
	Span syntax.Span // the Test's name
	Pass bool
	// Expect is the failing expect, from 1, or 0 when the test passed or
	// stopped on an error; ExpectSpan is where it is.
	Expect     int
	ExpectSpan syntax.Span
	Message    string
	// Said is what the test's commands said, through say and tellraw.
	Said []string
}

// Tests lists the project's Test declarations, in source order.
func (r *Result) Tests() []*Decl {
	var out []*Decl
	for _, d := range r.Order {
		if d.Kind.MMS == "Test" && d.D.Name != nil {
			out = append(out, d)
		}
	}
	return out
}

// RunTests runs the Tests that keep says to (all of them when keep is
// nil), each in a fresh world in which the pack has loaded. A project
// with errors runs none.
func (r *Result) RunTests(keep func(name string) bool) ([]TestResult, error) {
	if r.HasErrors() {
		return nil, fmt.Errorf("the project has errors; fix them to run its tests")
	}
	files := map[string][]string{}
	tags := map[string][]string{}
	for _, f := range r.emit(true) {
		rest, ok := strings.CutPrefix(f.Path, "data/")
		if !ok {
			continue
		}
		ns, rest, _ := strings.Cut(rest, "/")
		if path, ok := strings.CutPrefix(rest, "function/"); ok && strings.HasSuffix(path, ".mcfunction") {
			files[ns+":"+strings.TrimSuffix(path, ".mcfunction")] = strings.Split(strings.TrimSuffix(string(f.Data), "\n"), "\n")
		}
		if path, ok := strings.CutPrefix(rest, "tags/function/"); ok && strings.HasSuffix(path, ".json") {
			var t struct {
				Values []json.RawMessage `json:"values"`
			}
			if json.Unmarshal(f.Data, &t) == nil {
				for _, v := range t.Values {
					var id string
					if json.Unmarshal(v, &id) != nil {
						var o struct{ ID string }
						json.Unmarshal(v, &o)
						id = o.ID
					}
					tags[ns+":"+strings.TrimSuffix(path, ".json")] = append(tags[ns+":"+strings.TrimSuffix(path, ".json")], id)
				}
			}
		}
	}
	var out []TestResult
	for _, d := range r.Tests() {
		name := d.D.Name.Text
		if keep != nil && !keep(name) {
			continue
		}
		id := sema.TestID(d.NS, name)
		res := TestResult{Name: name, ID: id, Unit: d.Unit.Path, Span: d.D.Name.Sp}
		vm := mcfn.New(files)
		err := runTag(vm, tags, "minecraft:load", map[string]bool{})
		if err != nil {
			res.Message = "while the pack loaded: " + err.Error()
			res.Said = vm.Said
			out = append(out, res)
			continue
		}
		before := len(vm.Said)
		v, _, err := vm.Run(id)
		res.Said = vm.Said[before:]
		switch {
		case err != nil:
			res.Message = err.Error()
		case v == 0:
			res.Pass = true
		case int(v) <= len(d.Expects):
			e := d.Expects[v-1]
			res.Expect, res.ExpectSpan = int(v), e.Sp
			res.Message = "expected " + e.Text
			if e.Values {
				res.Message += ", but " + seen(vm, d.NS, e)
			}
		default:
			res.Message = fmt.Sprintf("the test returned %d", v)
		}
		out = append(out, res)
	}
	return out, nil
}

// seen says what the two sides of a failed comparison were.
func seen(vm *mcfn.VM, ns string, e sema.TestExpect) string {
	// A side's text and value: "x is 3", or "x has no score" when it
	// reads a score that is not set.
	side := func(text, key string) string {
		v, err := vm.Get(fn.Storage(ns), sema.ExpectStorage+"."+key)
		if err != nil {
			return text + " has no score"
		}
		return text + " is " + v.String()
	}
	lit := func(text string) bool { _, err := strconv.Atoi(text); return err == nil }
	switch {
	case lit(e.Right) && lit(e.Left):
		return "it does not hold"
	case lit(e.Right):
		return side(e.Left, "left")
	case lit(e.Left):
		return side(e.Right, "right")
	}
	return side(e.Left, "left") + " and " + side(e.Right, "right")
}

func runTag(vm *mcfn.VM, tags map[string][]string, tag string, seen map[string]bool) error {
	if seen[tag] {
		return nil
	}
	seen[tag] = true
	for _, id := range tags[tag] {
		if t, ok := strings.CutPrefix(id, "#"); ok {
			if err := runTag(vm, tags, t, seen); err != nil {
				return err
			}
			continue
		}
		if _, _, err := vm.Run(id); err != nil {
			return err
		}
	}
	return nil
}
