// Package prelude holds the MMS source that ships with the language: the
// declaration kinds the compiler builds in Go, the intrinsics, and the
// blueprints written in MMS itself, such as CustomItem. A project loads
// it before its own files, the way a library will be loaded.
package prelude

import (
	"embed"
	"io/fs"
	"sort"
	"strings"
	"sync"

	"github.com/minecraftmetascript/mms/engine/syntax"
)

//go:embed *.mms
var files embed.FS

// Prefix starts the path of every prelude file, which no project file
// can have.
const Prefix = "@prelude/"

// File is one prelude source file.
type File struct {
	Path, Text string
}

// Files returns the prelude's files, by path.
func Files() []File {
	names, _ := fs.Glob(files, "*.mms")
	sort.Strings(names)
	out := make([]File, 0, len(names))
	for _, n := range names {
		b, _ := files.ReadFile(n)
		out = append(out, File{Path: Prefix + n, Text: string(b)})
	}
	return out
}

// Is reports whether path names a prelude file.
func Is(path string) bool { return strings.HasPrefix(path, Prefix) }

// Parsed is a prelude file, parsed.
type Parsed struct {
	Path string
	File *syntax.File
}

// Parse returns the prelude's files, parsed once and shared: callers
// must not change them.
var Parse = sync.OnceValue(func() []Parsed {
	var out []Parsed
	for _, f := range Files() {
		out = append(out, Parsed{Path: f.Path, File: syntax.Parse(f.Text)})
	}
	return out
})

// Kind is a declaration kind the prelude declares with `extern Blueprint`.
type Kind struct {
	Name, Doc string
}

// Kinds returns the kinds the compiler builds in Go, by name.
var Kinds = sync.OnceValue(func() map[string]Kind {
	out := map[string]Kind{}
	for _, f := range Parse() {
		for _, d := range f.File.Decls {
			if d.Extern && d.Kind != nil && d.Kind.Name == "Blueprint" && d.Name != nil {
				out[d.Name.Text] = Kind{Name: d.Name.Text, Doc: d.Doc}
			}
		}
	}
	return out
})
