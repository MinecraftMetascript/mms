package project

import (
	"context"
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/internal/mcfn"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// world builds src as namespace demo and returns an interpreter that has
// run the load function.
func world(t *testing.T, src string) *mcfn.VM {
	t.Helper()
	p := New(spec.MustDefault(), Config{Namespace: "demo", Description: "test"})
	p.SetFile("main.mms", src)
	r, err := p.Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range r.Diagnostics("main.mms") {
		if d.Severity == syntax.SeverityError {
			t.Fatalf("%s: %s", d.Code, d.Message)
		}
	}
	return load(t, r, "demo")
}

// load builds an interpreter from the result's functions and runs ns's
// load function.
func load(t *testing.T, r *Result, ns string) *mcfn.VM {
	t.Helper()
	files := map[string][]string{}
	for _, f := range r.Emit() {
		rest, ok := strings.CutPrefix(f.Path, "data/")
		if !ok || !strings.HasSuffix(rest, ".mcfunction") {
			continue
		}
		ns, path, _ := strings.Cut(strings.TrimSuffix(rest, ".mcfunction"), "/function/")
		files[ns+":"+path] = strings.Split(strings.TrimSuffix(string(f.Data), "\n"), "\n")
	}
	vm := mcfn.New(files)
	if _, ok := files[ns+":mms_gen/load"]; ok {
		if _, _, err := vm.Run(ns + ":mms_gen/load"); err != nil {
			t.Fatal(err)
		}
	}
	return vm
}

func score(t *testing.T, vm *mcfn.VM, holder, obj string) int32 {
	t.Helper()
	v, ok := vm.Score(holder, obj)
	if !ok {
		t.Fatalf("%s has no score on %s", holder, obj)
	}
	return v
}

func call(t *testing.T, vm *mcfn.VM, id string) int32 {
	t.Helper()
	v, _, err := vm.Run(id)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func TestFunctionsRun(t *testing.T) {
	vm := world(t, `
Global total: int = 5
Global flag: bool = true
Objective kills

Function add(a: int, b: int): int {
  return a + b
}

Function fib(n: int): int {
  let a = 0
  let b = 1
  for i in 1..n {
    let t = a + b
    a = b
    b = t
  }
  return a
}

Function arith {
  total = (17 / -5) * 100 + 17 % -5 * 10 + -(3 - 10)
}

Function sign(x: int): int {
  if (x > 0) {
    return 1
  } else if (x < 0) {
    return -1
  }
  return 0
}

Function loops: int {
  let sum = 0
  let i = 0
  while (true) {
    i += 1
    if (i > 20) {
      break
    }
    if (i % 3 == 0) {
      continue
    }
    if (i == 17) {
      return sum
    }
    sum += i
  }
  return -1
}

Function nested(): int {
  let found = 0
  for x in 1..5 {
    for y in 1..5 {
      if (x * y == 12 && x < y) {
        found = x * 10 + y
        break
      }
    }
    if (found > 0) {
      break
    }
  }
  return found
}

Function logic(a: int, b: int): bool {
  return (a > 0 || b > 0) && !(a == b)
}

Function bump {
  kills["#me"] += add(2, 3) * 2
  if (flag) {
    flag = false
  }
}

OnLoad {
  total = total + fib(10)
}
`)
	if v := score(t, vm, "$total", "mms.demo"); v != 60 {
		t.Errorf("after load total = %d, want 5 + fib(10) = 60", v)
	}
	call(t, vm, "demo:arith")
	// floor(17 / -5) = -4; floorMod(17, -5) = -3; -(3 - 10) = 7
	if v := score(t, vm, "$total", "mms.demo"); v != -400-30+7 {
		t.Errorf("arith = %d, want %d", v, -400-30+7)
	}
	for x, want := range map[int32]int32{5: 1, -3: -1, 0: 0} {
		vm.Scores["mms.demo"]["#sign.x"] = x
		if v := call(t, vm, "demo:sign"); v != want {
			t.Errorf("sign(%d) = %d, want %d", x, v, want)
		}
	}
	// 1..16 without multiples of 3: 136 - (3+6+9+12+15) = 91
	if v := call(t, vm, "demo:loops"); v != 91 {
		t.Errorf("loops = %d, want 91", v)
	}
	if v := call(t, vm, "demo:nested"); v != 34 {
		t.Errorf("nested = %d, want 34", v)
	}
	for _, tc := range []struct{ a, b, want int32 }{{1, 0, 1}, {0, 0, 0}, {2, 2, 0}, {-1, 3, 1}, {-1, -2, 0}} {
		vm.Scores["mms.demo"]["#logic.a"] = tc.a
		vm.Scores["mms.demo"]["#logic.b"] = tc.b
		if v := call(t, vm, "demo:logic"); v != tc.want {
			t.Errorf("logic(%d, %d) = %d, want %d", tc.a, tc.b, v, tc.want)
		}
	}
	call(t, vm, "demo:bump")
	call(t, vm, "demo:bump")
	if v := score(t, vm, "#me", "kills"); v != 20 {
		t.Errorf("kills = %d, want 20", v)
	}
	if v := score(t, vm, "$flag", "mms.demo"); v != 0 {
		t.Errorf("flag = %d, want 0", v)
	}
	if v := score(t, vm, "#cf", "mms.demo"); v != 0 {
		t.Errorf("#cf = %d after running, want 0", v)
	}
}

func TestFunctionErrors(t *testing.T) {
	cases := map[string]string{
		"Function f(): int {\n  let x = 1\n}":                 "mms/missing-return",
		"Function f {\n  break\n}":                            "mms/misplaced",
		"Function f {\n  let x = 1\n  let x = 2\n}":           "mms/duplicate",
		"Function f {\n  let x: int = true\n}":                "mms/type-mismatch",
		"Function f {\n  y = 1\n}":                            "mms/unknown-name",
		"Function f {\n  if (1) {\n  }\n}":                    "mms/type-mismatch",
		"Function f {\n  g(1)\n}\nFunction g {\n  /say hi\n}": "mms/arity",
		"Function f {\n  let x = 1 / 0\n}":                    "mms/division-by-zero",
		"Function f {\n  1 + 2\n}":                            "mms/unused-value",
		"Function f(): int {\n  return true\n}":               "mms/type-mismatch",
		"Objective k\nFunction f {\n  k = 1\n}":               "mms/type-mismatch",
		"Global g: int = 1 + 2\nGlobal h: int = g":            "mms/not-constant",
		"Function f(s: selector) {\n  /say hi\n}":             "mms/unknown-type",
		"Function f {\n  let x = 3000000000\n}":               "mms/type-mismatch",
	}
	for src, want := range cases {
		p := New(spec.MustDefault(), Config{Namespace: "demo"})
		p.SetFile("main.mms", src)
		r, _ := p.Analyze(context.Background())
		var codes []string
		for _, d := range r.Diagnostics("main.mms") {
			if d.Severity == syntax.SeverityError {
				codes = append(codes, d.Code)
			}
		}
		if len(codes) == 0 || codes[0] != want {
			t.Errorf("%q: errors %v, want %s", src, codes, want)
		}
	}
}

// TestSignatureDeps checks that callers are rechecked when a callee's
// signature changes in another file.
func TestSignatureDeps(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("a.mms", "Function f {\n  g(1)\n}\n")
	p.SetFile("b.mms", "Function g {\n  /say hi\n}\n")
	errs := func() int {
		r, _ := p.Analyze(context.Background())
		n := 0
		for _, d := range r.Diagnostics("a.mms") {
			if d.Severity == syntax.SeverityError {
				n++
			}
		}
		return n
	}
	if errs() != 1 {
		t.Fatal("want an arity error first")
	}
	p.SetFile("b.mms", "Function g(x: int) {\n  /say hi\n}\n")
	if n := errs(); n != 0 {
		t.Fatalf("after adding the parameter, %d errors remain", n)
	}
}

func TestStringsAndInterpolation(t *testing.T) {
	vm := world(t, `
Global greeting: string = "hello"
Global n: int = 3
Objective kills

Function greet(who: string, times: int) {
  let msg = greeting + ", " + who
  for i in 1..times {
    /say ${msg} #${i}
  }
  if (who == "steve") {
    /say hi steve
  }
  if (who != greeting) {
    /say not hello
  }
  /say const ${2 * 21} ${true} ${n > 2}
}

Function run {
  greet("alex", 2)
  greet("steve", 1)
  greeting += "!"
  greet(greeting, 1)
}
`)
	call(t, vm, "demo:run")
	want := []string{
		"hello, alex #1", "hello, alex #2", "not hello", "const 42 true true",
		"hello, steve #1", "hi steve", "not hello", "const 42 true true",
		"hello!, hello! #1", "const 42 true true",
	}
	if strings.Join(vm.Said, "\n") != strings.Join(want, "\n") {
		t.Errorf("said:\n%s\nwant:\n%s", strings.Join(vm.Said, "\n"), strings.Join(want, "\n"))
	}
}

func TestRecursion(t *testing.T) {
	vm := world(t, `
Function fib(n: int): int {
  if (n < 2) {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}

Function fact(n: int): int {
  let here = n
  if (n <= 1) {
    return 1
  }
  let rest = fact(n - 1)
  return here * rest
}

Function is_even(n: int): bool {
  if (n == 0) {
    return true
  }
  return is_odd(n - 1)
}

Function is_odd(n: int): bool {
  if (n == 0) {
    return false
  }
  return is_even(n - 1)
}

Function shout(word: string, n: int) {
  if (n > 0) {
    shout(word + "!", n - 1)
    /say ${word}
  }
}
`)
	vm.Scores["mms.demo"]["#fib.n"] = 15
	if v := call(t, vm, "demo:fib"); v != 610 {
		t.Errorf("fib(15) = %d, want 610", v)
	}
	vm.Scores["mms.demo"]["#fact.n"] = 10
	if v := call(t, vm, "demo:fact"); v != 3628800 {
		t.Errorf("fact(10) = %d, want 3628800", v)
	}
	for n, want := range map[int32]int32{0: 1, 7: 0, 10: 1} {
		vm.Scores["mms.demo"]["#is_even.n"] = n
		if v := call(t, vm, "demo:is_even"); v != want {
			t.Errorf("is_even(%d) = %d, want %d", n, v, want)
		}
	}
	vm.Storage["demo:mms"] = map[string]any{"l": map[string]any{"shout": map[string]any{"word": "hey"}}}
	vm.Scores["mms.demo"]["#shout.n"] = 3
	call(t, vm, "demo:shout")
	if got := strings.Join(vm.Said, " "); got != "hey!! hey! hey" {
		t.Errorf("shout said %q", got)
	}
	if st, _ := vm.Storage["demo:mms"]["stack"].([]any); len(st) != 0 {
		t.Errorf("stack left with %d frames", len(st))
	}
}

// TestOptimizedOutput pins what the optimiser does to common shapes.
func TestOptimizedOutput(t *testing.T) {
	p := New(spec.MustDefault(), Config{Namespace: "demo"})
	p.SetFile("main.mms", `
Function inc {
  let x = 1
  x = x + 1
}
Function down(n: int): int {
  if (n <= 0) {
    return 0
  }
  let unused = n * 7
  return down(n - 1) + n
}
`)
	r, _ := p.Analyze(context.Background())
	got := dumpFiles(r.Emit())
	for _, want := range []string{
		"=== data/demo/function/inc.mcfunction\nscoreboard players set #inc.x mms.demo 1\nscoreboard players add #inc.x mms.demo 1\n",
		// Only n is needed after the recursive call, so only n is saved.
		"data modify storage demo:mms stack append value {}\nexecute store result storage demo:mms stack[-1].s0 int 1 run scoreboard players get #down.n mms.demo\nscoreboard players operation #down.n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
	if strings.Contains(got, "stack[-1].s1") {
		t.Errorf("saved more than n:\n%s", got)
	}
}

// TestExampleFunctions runs the function examples' demo functions.
func TestExampleFunctions(t *testing.T) {
	r, err := loadDir(t, "../../examples/functions-advanced").Analyze(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	vm := load(t, r, "adv")
	call(t, vm, "adv:demo")
	want := []string{
		`"Hello, Steve!"`,
		"fib(10) is 55",
		"gcd(84, 36) is 12",
		"2^10 is 1024",
		`"Primes: 2 3 5 7 11 13 17 19 23 29 (10 in total)"`, // tellraw keeps its JSON text
	}
	if strings.Join(vm.Said, "\n") != strings.Join(want, "\n") {
		t.Errorf("said:\n%s\nwant:\n%s", strings.Join(vm.Said, "\n"), strings.Join(want, "\n"))
	}
}
