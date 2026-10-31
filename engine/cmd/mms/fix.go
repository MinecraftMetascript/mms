package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/minecraftmetascript/mms/engine/format"
)

// runFix applies every safe lint fix in the project and writes the
// changed files, formatted.
func runFix(args []string, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("fix", flag.ContinueOnError)
	fl.SetOutput(stderr)
	if err := fl.Parse(args); err != nil {
		return 2
	}
	dir := "."
	if fl.NArg() > 0 {
		dir = fl.Arg(0)
	}
	p, err := load(dir)
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	changed, n, err := p.FixAll(context.Background())
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	paths := make([]string, 0, len(changed))
	for path := range changed {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		text := changed[path]
		if out, err := format.Source(text); err == nil {
			text = out
		}
		if err := os.WriteFile(filepath.Join(dir, path), []byte(text), 0o644); err != nil {
			fmt.Fprintln(stderr, "mms:", err)
			return 1
		}
		fmt.Fprintln(stdout, "fixed", path)
	}
	fmt.Fprintf(stdout, "%d %s fixed in %d %s\n", n, pluralize(n, "problem"), len(paths), pluralize(len(paths), "file"))
	return 0
}
