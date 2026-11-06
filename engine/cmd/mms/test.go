package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"regexp"
)

// runTest runs the project's Test declarations in the test interpreter.
func runTest(args []string, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("test", flag.ContinueOnError)
	fl.SetOutput(stderr)
	run := fl.String("run", "", "run only the tests whose names match this regular expression")
	verbose := fl.Bool("v", false, "print every test, and what each test's commands said")
	if err := fl.Parse(args); err != nil {
		return 2
	}
	dir := "."
	if fl.NArg() > 0 {
		dir = fl.Arg(0)
	}
	var keep func(string) bool
	if *run != "" {
		re, err := regexp.Compile(*run)
		if err != nil {
			fmt.Fprintln(stderr, "mms: -run:", err)
			return 2
		}
		keep = re.MatchString
	}
	p, err := load(dir)
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	r, err := p.Analyze(context.Background())
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	if n := report(stderr, r, false); n > 0 {
		fmt.Fprintf(stderr, "%d %s; no tests run\n", n, pluralize(n, "error"))
		return 1
	}
	res, err := r.RunTests(keep)
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	failed := 0
	for _, t := range res {
		u := r.Unit(t.Unit)
		if t.Pass {
			if *verbose {
				fmt.Fprintf(stdout, "ok    %s\n", t.Name)
			}
		} else {
			failed++
			sp := t.Span
			if t.Expect > 0 {
				sp = t.ExpectSpan
			}
			line, col := u.File.LineCol(sp.Start)
			fmt.Fprintf(stdout, "FAIL  %s\n    %s:%d:%d: %s\n", t.Name, t.Unit, line, col, t.Message)
		}
		if *verbose || !t.Pass {
			for _, s := range t.Said {
				fmt.Fprintf(stdout, "    said: %s\n", s)
			}
		}
	}
	switch {
	case len(res) == 0:
		fmt.Fprintln(stdout, "no tests")
	case failed > 0:
		fmt.Fprintf(stdout, "%d of %d %s failed\n", failed, len(res), pluralize(len(res), "test"))
		return 1
	default:
		fmt.Fprintf(stdout, "ok: %d %s passed\n", len(res), pluralize(len(res), "test"))
	}
	return 0
}
