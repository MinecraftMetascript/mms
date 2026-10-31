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
	if _, ok := files["demo:mms_gen/load"]; ok {
		if _, _, err := vm.Run("demo:mms_gen/load"); err != nil {
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
