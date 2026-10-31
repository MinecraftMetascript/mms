package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/minecraftmetascript/mms/engine/format"
)

// runFmt formats files in place, or reports which ones are unformatted with
// -check. With no paths it formats stdin to stdout. Directories are walked
// for .mms files. It returns the process exit code.
func runFmt(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("fmt", flag.ContinueOnError)
	fl.SetOutput(stderr)
	check := fl.Bool("check", false, "list files that are not formatted and exit 1, without writing")
	if err := fl.Parse(args); err != nil {
		return 2
	}
	if fl.NArg() == 0 {
		src, err := io.ReadAll(stdin)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		out, err := format.Source(string(src))
		if err != nil {
			fmt.Fprintln(stderr, "<stdin>:", err)
			return 1
		}
		if *check {
			if out != string(src) {
				fmt.Fprintln(stdout, "<stdin>")
				return 1
			}
			return 0
		}
		io.WriteString(stdout, out)
		return 0
	}

	code := 0
	for _, path := range collect(fl.Args(), stderr, &code) {
		src, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(stderr, err)
			code = 1
			continue
		}
		out, err := format.Source(string(src))
		if err != nil {
			if errors.Is(err, format.ErrUnsafe) {
				err = fmt.Errorf("%w (please report this)", err)
			}
			fmt.Fprintf(stderr, "%s: %v\n", path, err)
			code = 1
			continue
		}
		if out == string(src) {
			continue
		}
		if *check {
			fmt.Fprintln(stdout, path)
			code = 1
			continue
		}
		if err := os.WriteFile(path, []byte(out), 0o644); err != nil {
			fmt.Fprintln(stderr, err)
			code = 1
		}
	}
	return code
}

// collect expands directories into the .mms files under them.
func collect(args []string, stderr io.Writer, code *int) []string {
	var paths []string
	for _, a := range args {
		info, err := os.Stat(a)
		if err != nil {
			fmt.Fprintln(stderr, err)
			*code = 1
			continue
		}
		if !info.IsDir() {
			paths = append(paths, a)
			continue
		}
		filepath.WalkDir(a, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintln(stderr, err)
				*code = 1
				return nil
			}
			if d.IsDir() && p != a && strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			if !d.IsDir() && strings.HasSuffix(p, ".mms") {
				paths = append(paths, p)
			}
			return nil
		})
	}
	return paths
}
