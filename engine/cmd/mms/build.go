package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/spec"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// load reads the project in dir: mms.json and every source file.
func load(dir string) (*project.Project, error) {
	cfg, err := project.LoadConfig(dir)
	if err != nil {
		return nil, err
	}
	pack, err := spec.Default()
	if err != nil {
		return nil, err
	}
	p := project.New(pack, cfg)
	p.Files = os.DirFS(dir)
	if cfg.Icon != "" {
		if st, err := os.Stat(filepath.Join(dir, filepath.FromSlash(cfg.Icon))); err != nil || st.IsDir() {
			return nil, fmt.Errorf("mms.json: icon: no file %q in the project", cfg.Icon)
		}
	}
	paths, err := cfg.SourceFiles(dir)
	if err != nil {
		return nil, err
	}
	for _, rel := range paths {
		b, err := os.ReadFile(filepath.Join(dir, rel))
		if err != nil {
			return nil, err
		}
		p.SetFile(rel, string(b))
	}
	return p, nil
}

// report prints diagnostics like a compiler: path:line:col, the message,
// and the source line with a caret. Hints are editor suggestions and are
// printed only when asked for. It returns the number of errors.
func report(w io.Writer, r *project.Result, hints bool) int {
	errs := 0
	for _, u := range r.Units {
		for _, d := range r.Diagnostics(u.Path) {
			if d.Severity == syntax.SeverityHint && !hints {
				continue
			}
			line, col := u.File.LineCol(d.Sp.Start)
			sev := "error"
			switch d.Severity {
			case syntax.SeverityWarning:
				sev = "warning"
			case syntax.SeverityInfo, syntax.SeverityHint:
				sev = "note"
			default:
				errs++
			}
			fmt.Fprintf(w, "%s:%d:%d: %s: %s [%s]\n", u.Path, line, col, sev, d.Message, d.Code)
			text := u.File.LineText(line)
			width := max(1, min(d.Sp.End-d.Sp.Start, len([]rune(text))-col+1))
			fmt.Fprintf(w, "    %s\n    %s%s\n", strings.ReplaceAll(text, "\t", " "), strings.Repeat(" ", col-1), strings.Repeat("^", width))
		}
	}
	return errs
}

func runCheck(args []string, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("check", flag.ContinueOnError)
	fl.SetOutput(stderr)
	hints := fl.Bool("hints", false, "also print hints (style suggestions `mms fix` can apply)")
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
	r, err := p.Analyze(context.Background())
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	if n := report(stderr, r, *hints); n > 0 {
		fmt.Fprintf(stderr, "%d %s\n", n, pluralize(n, "error"))
		return 1
	}
	fmt.Fprintf(stdout, "ok: %d %s\n", len(r.Units), pluralize(len(r.Units), "file"))
	return 0
}

func runBuild(args []string, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("build", flag.ContinueOnError)
	fl.SetOutput(stderr)
	out := fl.String("out", "", "output folder (default: mms.json output, or build); a resource pack goes next to it, in <out>-resources")
	zipPath := fl.String("zip", "", "write a zip archive instead of a folder; a resource pack goes in <name>-resources.zip")
	world := fl.String("world", "", "install into a world: writes <world>/datapacks/<namespace>, and the resource pack to the game folder's resourcepacks/<namespace>-resources.zip")
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
	r, err := p.Analyze(context.Background())
	if err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	if n := report(stderr, r, false); n > 0 {
		fmt.Fprintf(stderr, "%d %s; nothing written\n", n, pluralize(n, "error"))
		return 1
	}
	files := r.Emit()
	if err := project.ReadAssets(files, p.Files); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	data, res := project.SplitPacks(files)
	fail := func(err error) int {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	writeZip := func(path string, files []project.OutFile) error {
		b, err := project.Zip(files)
		if err == nil {
			err = os.WriteFile(path, b, 0o644)
		}
		return err
	}
	switch {
	case *zipPath != "":
		if err := writeZip(*zipPath, data); err != nil {
			return fail(err)
		}
		fmt.Fprintf(stdout, "wrote %s (%d files)\n", *zipPath, len(data))
		if len(res) > 0 {
			rz := strings.TrimSuffix(*zipPath, filepath.Ext(*zipPath)) + "-resources.zip"
			if err := writeZip(rz, res); err != nil {
				return fail(err)
			}
			fmt.Fprintf(stdout, "wrote resource pack %s (%d files)\n", rz, len(res))
		}
		return 0
	case *world != "":
		if _, err := os.Stat(filepath.Join(*world, "level.dat")); err != nil {
			fmt.Fprintf(stderr, "mms: %s does not look like a world folder (no level.dat)\n", *world)
			return 1
		}
		target := filepath.Join(*world, "datapacks", p.Config.Namespace)
		if err := project.WriteDir(target, data); err != nil {
			return fail(err)
		}
		fmt.Fprintf(stdout, "wrote %d files to %s\n", len(data), target)
		if len(res) > 0 {
			rz, enable := worldResourcePack(*world, p.Config.Namespace)
			if err := os.MkdirAll(filepath.Dir(rz), 0o755); err != nil {
				return fail(err)
			}
			if err := writeZip(rz, res); err != nil {
				return fail(err)
			}
			fmt.Fprintf(stdout, "wrote resource pack %s (%d files)\n", rz, len(res))
			fmt.Fprintln(stdout, enable)
		}
		return 0
	}
	resOut := filepath.Join(dir, p.Config.ResourceOutput())
	if *out == "" {
		*out = filepath.Join(dir, p.Config.Output)
	} else {
		resOut = filepath.Clean(*out) + "-resources"
	}
	if err := project.WriteDir(*out, data); err != nil {
		return fail(err)
	}
	fmt.Fprintf(stdout, "wrote %d files to %s\n", len(data), *out)
	if len(res) > 0 {
		if err := project.WriteDir(resOut, res); err != nil {
			return fail(err)
		}
		fmt.Fprintf(stdout, "wrote resource pack: %d files to %s\n", len(res), resOut)
	}
	return 0
}

func runInit(args []string, stdout, stderr io.Writer) int {
	fl := flag.NewFlagSet("init", flag.ContinueOnError)
	fl.SetOutput(stderr)
	ns := fl.String("namespace", "", "default namespace (default: the folder name)")
	if err := fl.Parse(args); err != nil {
		return 2
	}
	dir := "."
	if fl.NArg() > 0 {
		dir = fl.Arg(0)
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	if _, err := os.Stat(filepath.Join(dir, project.ConfigFile)); err == nil {
		fmt.Fprintf(stderr, "mms: %s already exists\n", filepath.Join(dir, project.ConfigFile))
		return 1
	}
	if *ns == "" {
		abs, _ := filepath.Abs(dir)
		*ns = sanitizeNS(filepath.Base(abs))
	}
	cfg := fmt.Sprintf("{\n  \"$schema\": %q,\n  \"namespace\": %q,\n  \"description\": \"A data pack built with MMS\",\n  \"output\": \"build\"\n}\n", project.SchemaURL, *ns)
	src := fmt.Sprintf("Namespace %s\n\n/** A gentle hill shape. */\nDensityFunction hills = Noise(minecraft:ridge, 0.25, 0) * 0.5\n\nOnLoad {\n  /say %s loaded\n}\n", *ns, *ns)
	if err := os.WriteFile(filepath.Join(dir, project.ConfigFile), []byte(cfg), 0o644); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	main := filepath.Join(dir, "main.mms")
	if _, err := os.Stat(main); os.IsNotExist(err) {
		os.WriteFile(main, []byte(src), 0o644)
	}
	fmt.Fprintf(stdout, "created %s and main.mms; run `mms build %s`\n", project.ConfigFile, dir)
	return 0
}

func sanitizeNS(s string) string {
	s = strings.ToLower(s)
	var sb strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '_', r == '-', r == '.':
			sb.WriteRune(r)
		default:
			sb.WriteByte('_')
		}
	}
	if sb.Len() == 0 {
		return "mms"
	}
	return sb.String()
}

func pluralize(n int, s string) string {
	if n == 1 {
		return s
	}
	return s + "s"
}

// worldResourcePack picks where --world puts the resource pack, and the
// line that says how to turn it on. A singleplayer world's own
// resources.zip is not applied (checked in 26.3), so a world inside a
// game folder's saves/ gets the pack in that game folder's resourcepacks/,
// under a fixed name so that enabling it once keeps it on across builds.
// Any other folder, such as a server's world, gets it beside the world.
func worldResourcePack(world, ns string) (path, enable string) {
	name := ns + "-resources.zip"
	abs, err := filepath.Abs(world)
	if err != nil {
		abs = world
	}
	saves := filepath.Dir(abs)
	if filepath.Base(saves) == "saves" {
		game := filepath.Dir(saves)
		return filepath.Join(game, "resourcepacks", name),
			"enable it once in Options > Resource Packs; it stays on after rebuilds, and F3+T reloads it in game"
	}
	return filepath.Join(filepath.Dir(abs), name),
		"this world is not in a game folder's saves/, so copy it to resourcepacks/ and enable it, or set it as the server's resource pack"
}
