package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/minecraftmetascript/mms/engine/prelude"
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
			for _, rel := range d.Related {
				fmt.Fprintf(w, "    note: %s, at %s\n", rel.Message, relatedAt(r, rel))
			}
		}
	}
	for _, msg := range r.ConfigWarnings() {
		fmt.Fprintf(w, "%s: warning: %s\n", project.ConfigFile, msg)
	}
	// Files nothing uses have no line to point at.
	if paths, sev := r.UnusedAssets(); sev != syntax.SeverityHint || hints {
		name := map[syntax.Severity]string{syntax.SeverityError: "error", syntax.SeverityWarning: "warning"}[sev]
		if name == "" {
			name = "note"
		}
		for _, p := range paths {
			if sev == syntax.SeverityError {
				errs++
			}
			fmt.Fprintf(w, "%s: %s: %s [lint/unused-asset]\n", p, name, project.UnusedAssetMessage(p))
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
	world := fl.String("world", "", "install into a world: writes <world>/datapacks/<namespace>, and merges the resource pack into <world>/resourcepacks/resources.zip, keeping other projects' files")
	var reset resetFlag
	fl.Var(&reset, "reset-dimensions", "with -world: delete the saved chunks of the project's dimensions whose worldgen changed, so they generate again; or of the dimensions named, as -reset-dimensions=ns:a,ns:b. Old chunks go to <world>/"+project.ResetFolder+"; the world must be closed")
	if err := fl.Parse(args); err != nil {
		return 2
	}
	dir := "."
	if fl.NArg() > 0 {
		dir = fl.Arg(0)
	}
	if reset.set && *world == "" {
		fmt.Fprintln(stderr, "mms: -reset-dimensions needs -world")
		return 2
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
		before, err := project.ReadPackDir(target)
		if err != nil {
			return fail(err)
		}
		var resetIDs []string
		if reset.set {
			if resetIDs = reset.ids; resetIDs == nil {
				resetIDs = project.ChangedDimensions(before, data)
			}
			// Checked before anything is written, so the next build sees
			// the same changes and resets the same dimensions.
			if len(resetIDs) > 0 && worldInUse(*world) {
				fmt.Fprintf(stderr, "mms: %s is open in Minecraft; leave the world to reset %s, then build again (nothing written)\n", *world, strings.Join(resetIDs, ", "))
				return 1
			}
		}
		if err := project.WriteDir(target, data); err != nil {
			return fail(err)
		}
		fmt.Fprintf(stdout, "wrote %d files to %s\n", len(data), target)
		rz, note := worldResourcePack(*world)
		oldZip, statErr := os.ReadFile(rz)
		if statErr != nil && !os.IsNotExist(statErr) {
			return fail(statErr)
		}
		resChanged := false
		// With no resources, a pack already there still loses the files
		// this project put in it last time.
		if len(res) > 0 || statErr == nil {
			m, err := project.WriteWorldPack(rz, p.Config.Namespace, res)
			if err != nil {
				return fail(err)
			}
			if len(res) == 0 {
				fmt.Fprintf(stdout, "%s has no resources now; removed what it had in %s\n", p.Config.Namespace, rz)
			} else {
				fmt.Fprintf(stdout, "wrote resource pack %s (%d files)\n", rz, len(res))
			}
			if m.Adopted {
				fmt.Fprintln(stdout, "replaced the resource pack an earlier mms build left there")
			}
			if len(m.Others) > 0 {
				fmt.Fprintf(stdout, "kept the other packs in it: %s\n", strings.Join(m.Others, ", "))
			}
			for _, r := range m.Replaced {
				fmt.Fprintf(stderr, "mms: warning: overwrote %s\n", r)
			}
			fmt.Fprintln(stdout, note)
			resChanged = !bytes.Equal(oldZip, m.Zip)
		}
		if reset.set {
			done, err := project.ResetDimensions(*world, resetIDs)
			for _, d := range done {
				if len(d.Moved) == 0 {
					fmt.Fprintf(stdout, "reset %s: nothing to reset, it has not been visited yet\n", d.ID)
				} else {
					fmt.Fprintf(stdout, "reset %s: it generates again when visited; moved its old chunks to %s\n", d.ID, filepath.Join(*world, project.ResetFolder))
				}
			}
			if err != nil {
				return fail(err)
			}
			if len(resetIDs) == 0 {
				fmt.Fprintln(stdout, "reset: no dimension's worldgen changed")
			}
		}
		fmt.Fprintln(stdout, project.PlanApply(before, data, resChanged).Message())
		return 0
	}
	resOut := filepath.Join(dir, p.Config.ResourceOutput())
	if *out == "" {
		*out = filepath.Join(dir, p.Config.Output)
		// mms.json comes with the project, so its folders must stay inside
		// it even through a symlink; -out is the user's own choice.
		for _, o := range []string{*out, resOut} {
			if err := insideProject(dir, o); err != nil {
				return fail(err)
			}
		}
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
	cfg := project.Config{Namespace: *ns, Description: "A data pack built with MMS", Output: "build"}
	if err := cfg.Validate(); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	src := fmt.Sprintf("Namespace %s\n\n/** A gentle hill shape. */\nDensityFunction hills = Noise(minecraft:ridge, 0.25, 0) * 0.5\n\nOnLoad {\n  /say %s loaded\n}\n", *ns, *ns)
	if err := os.WriteFile(filepath.Join(dir, project.ConfigFile), cfg.File(), 0o644); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	made := project.ConfigFile + " and main.mms"
	main := filepath.Join(dir, "main.mms")
	if _, err := os.Stat(main); err == nil {
		made = project.ConfigFile + " (kept the existing main.mms)"
	} else if err := os.WriteFile(main, []byte(src), 0o644); err != nil {
		fmt.Fprintln(stderr, "mms:", err)
		return 1
	}
	fmt.Fprintf(stdout, "created %s; run `mms build %s`\n", made, dir)
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

// worldResourcePack returns where --world puts the resource pack and a
// line about it. Minecraft applies a world's bundled pack, which lives at
// <world>/resourcepacks/resources.zip, by itself when the world is played
// (not over LAN, and a server needs its own resource-pack setting).
func worldResourcePack(world string) (path, note string) {
	return filepath.Join(world, "resourcepacks", "resources.zip"),
		"Minecraft applies it when the world is opened"
}

// resetFlag is -reset-dimensions: alone it means the dimensions whose
// worldgen changed; with a value, the dimensions it lists.
type resetFlag struct {
	set bool
	ids []string
}

func (f *resetFlag) String() string { return strings.Join(f.ids, ",") }

func (f *resetFlag) IsBoolFlag() bool { return true }

func (f *resetFlag) Set(v string) error {
	f.set = true
	switch v {
	case "true", "changed":
		f.ids = nil
		return nil
	case "false", "":
		f.set = false
		return nil
	}
	f.ids = nil
	for _, id := range strings.Split(v, ",") {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if !strings.Contains(id, ":") {
			id = "minecraft:" + id
		}
		if err := project.CheckDimensionID(id); err != nil {
			return err
		}
		f.ids = append(f.ids, id)
	}
	if len(f.ids) == 0 {
		return fmt.Errorf("no dimension ids in %q", v)
	}
	return nil
}

// insideProject says why out, an output folder from mms.json, resolves to
// somewhere outside the project folder dir, if it does. A folder that does
// not exist yet is checked through its nearest existing parent.
func insideProject(dir, out string) error {
	root, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return err
	}
	p, rest := filepath.Clean(out), ""
	for {
		real, err := filepath.EvalSymlinks(p)
		if err == nil {
			p = filepath.Join(real, rest)
			break
		}
		if !os.IsNotExist(err) {
			return err
		}
		parent := filepath.Dir(p)
		if parent == p {
			break
		}
		rest = filepath.Join(filepath.Base(p), rest)
		p = parent
	}
	if rel, err := filepath.Rel(root, p); err != nil || rel == "." || !filepath.IsLocal(rel) {
		return fmt.Errorf("%s resolves to %s, outside the project; mms.json's output folders must stay inside it", out, p)
	}
	return nil
}

// relatedAt is path:line:col for a related location, in a project file
// or the prelude.
func relatedAt(r *project.Result, rel syntax.Related) string {
	if u := r.Unit(rel.Path); u != nil {
		line, col := u.File.LineCol(rel.Sp.Start)
		return fmt.Sprintf("%s:%d:%d", rel.Path, line, col)
	}
	for _, f := range prelude.Files() {
		if f.Path == rel.Path {
			line, col := (&syntax.File{Src: []rune(f.Text)}).LineCol(rel.Sp.Start)
			return fmt.Sprintf("%s:%d:%d", rel.Path, line, col)
		}
	}
	return rel.Path
}
