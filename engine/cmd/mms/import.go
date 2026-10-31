package main

import (
	"archive/zip"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/minecraftmetascript/mms/engine/importer"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// runImport turns a data pack folder or zip into an MMS project: mms.json
// plus one formatted .mms file per namespace and resource folder.
func runImport(args []string, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("import", flag.ContinueOnError)
	fl.SetOutput(stderr)
	fl.Usage = func() {
		fmt.Fprintln(stderr, "usage: mms import [-force] <pack folder or .zip> [project folder]")
		fl.PrintDefaults()
	}
	force := fl.Bool("force", false, "write into a folder that already has mms.json")
	if err := fl.Parse(args); err != nil {
		return 2
	}
	if fl.NArg() < 1 || fl.NArg() > 2 {
		fl.Usage()
		return 2
	}
	src, out := fl.Arg(0), "."
	if fl.NArg() == 2 {
		out = fl.Arg(1)
	}
	if _, err := os.Stat(filepath.Join(out, project.ConfigFile)); err == nil && !*force {
		fmt.Fprintf(stderr, "mms: %s already exists (use -force to import anyway)\n", filepath.Join(out, project.ConfigFile))
		return 1
	}
	fsys, closer, err := openPack(src)
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	defer closer()
	sp, err := spec.Default()
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	res, err := importer.Pack(fsys, sema.NewKinds(sp), sp)
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	if err := res.Tidy(context.Background(), sp); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	for _, f := range res.Files {
		text := f.Text
		dst := filepath.Join(out, filepath.FromSlash(f.Path))
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err == nil {
			err = os.WriteFile(dst, []byte(text), 0o644)
		}
		if err != nil {
			fmt.Fprintln(stderr, "mms:", err)
			return 1
		}
	}
	cfg := project.Config{Namespace: res.Namespace, Description: res.Description, Output: "build"}
	if cfg.Description == "" {
		cfg.Description = "Imported with MMS"
	}
	b, _ := json.MarshalIndent(cfg, "", "  ")
	if err := os.WriteFile(filepath.Join(out, project.ConfigFile), append(b, '\n'), 0o644); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	for _, s := range res.Skipped {
		fmt.Fprintln(stderr, "skipped", s)
	}
	fmt.Fprintf(stdout, "imported %d resources into %d files in %s; run `mms check %s`\n", res.Count, len(res.Files), out, out)
	return 0
}

// openPack opens a data pack folder or zip. A zip whose pack.mcmeta sits
// in a single top-level folder is opened at that folder.
func openPack(src string) (fs.FS, func(), error) {
	st, err := os.Stat(src)
	if err != nil {
		return nil, nil, err
	}
	if st.IsDir() {
		return os.DirFS(src), func() {}, nil
	}
	zr, err := zip.OpenReader(src)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: not a folder or zip: %w", src, err)
	}
	var fsys fs.FS = zr
	for _, f := range zr.File {
		if path.Base(f.Name) == "pack.mcmeta" && strings.Count(f.Name, "/") == 1 {
			if sub, err := fs.Sub(zr, path.Dir(f.Name)); err == nil {
				fsys = sub
			}
			break
		}
	}
	return fsys, func() { zr.Close() }, nil
}
