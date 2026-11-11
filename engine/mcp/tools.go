package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/format"
	"github.com/minecraftmetascript/mms/engine/ide"
	"github.com/minecraftmetascript/mms/engine/importer"
	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/project"
	"github.com/minecraftmetascript/mms/engine/sema"
	"github.com/minecraftmetascript/mms/engine/syntax"
)

// tool is one MCP tool. run returns the text the model reads; an error
// is shown to the model as the tool's failure.
type tool struct {
	name, title, doc string
	props            map[string]any
	required         []string
	// writes marks a tool that can change files on disk.
	writes bool
	run    func(ctx context.Context, args json.RawMessage) (string, error)
}

func (t *tool) describe() map[string]any {
	schema := map[string]any{"type": "object", "properties": t.props, "additionalProperties": false}
	if len(t.required) > 0 {
		schema["required"] = t.required
	}
	return map[string]any{
		"name":        t.name,
		"title":       t.title,
		"description": t.doc,
		"inputSchema": schema,
		"annotations": map[string]any{
			"title":           t.title,
			"readOnlyHint":    !t.writes,
			"destructiveHint": false,
			"idempotentHint":  true,
			"openWorldHint":   false,
		},
	}
}

func str(doc string) map[string]any     { return map[string]any{"type": "string", "description": doc} }
func boolean(doc string) map[string]any { return map[string]any{"type": "boolean", "description": doc} }
func num(doc string) map[string]any     { return map[string]any{"type": "integer", "description": doc} }

// with returns the dir and source properties every project tool takes,
// plus its own.
func with(more map[string]any) map[string]any {
	out := map[string]any{
		"dir":    str("The project folder on disk: the one holding mms.json, or any folder of .mms files. Give this or source."),
		"source": str("The text of one .mms file, checked on its own as main.mms in the namespace `mms`. Give this or dir. It cannot use asset files."),
	}
	for k, v := range more {
		out[k] = v
	}
	return out
}

// target is where a project tool reads from.
type target struct {
	Dir    string `json:"dir"`
	Source string `json:"source"`
}

// sourcePath names the file a source argument is checked as.
const sourcePath = "main.mms"

func (t target) open() (*project.Project, error) {
	switch {
	case t.Dir != "" && t.Source != "":
		return nil, errors.New("give dir or source, not both")
	case t.Dir != "":
		st, err := os.Stat(t.Dir)
		if err != nil {
			return nil, err
		}
		if !st.IsDir() {
			return nil, fmt.Errorf("%s is a file; dir is the project folder", t.Dir)
		}
		return project.Load(t.Dir)
	case t.Source != "":
		return nil, nil
	}
	return nil, errors.New("give dir (a project folder) or source (the text of a .mms file)")
}

// load opens the target as a project, analysed.
func (s *Server) load(ctx context.Context, t target) (*project.Project, *project.Result, error) {
	p, err := t.open()
	if err != nil {
		return nil, nil, err
	}
	if p == nil {
		p = project.New(s.pack, project.DefaultConfig())
		p.SetFile(sourcePath, t.Source)
	}
	r, err := p.Analyze(ctx)
	if err != nil {
		return nil, nil, err
	}
	return p, r, nil
}

func parse(args json.RawMessage, into any) error {
	dec := json.NewDecoder(strings.NewReader(string(args)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(into); err != nil {
		return fmt.Errorf("bad arguments: %w", err)
	}
	return nil
}

func (s *Server) toolList() []*tool {
	return []*tool{
		{
			name: "check", title: "Check a project",
			doc: "Check an MMS project, or one file's text, and list every problem: where it is, the source line, and the fixes MMS can make. " +
				"Run it after every edit. A problem marked `safe fix` goes away with the fix tool.",
			props: with(map[string]any{"hints": boolean("Also list hints: style suggestions the fix tool can apply. Default false.")}),
			run:   s.check,
		},
		{
			name: "explain", title: "Show compiled output",
			doc: "Show what MMS code compiles to. Without name, list every file the project writes and the declaration each comes from. " +
				"With name, print the JSON or .mcfunction files that one declaration writes, and what its Blueprint instance or Template calls expand to.",
			props: with(map[string]any{"name": str("A declaration's name or id (`greet`, `demo:greet`), or part of an output path (`loot_table/ruby`).")}),
			run:   s.explain,
		},
		{
			name: "test", title: "Run tests",
			doc: "Run the project's Test declarations in MMS's built-in command interpreter and report each failing expect, with what the test's commands said. " +
				"The interpreter covers scoreboards, storage, functions, entities and tags, not the whole game.",
			props: with(map[string]any{"run": str("Run only the tests whose names match this regular expression.")}),
			run:   s.test,
		},
		{
			name: "fix", title: "Apply safe fixes",
			doc:    "Apply every safe lint fix. With dir, the changed files are written, formatted, and listed. With source, the fixed text is returned.",
			props:  with(nil),
			writes: true,
			run:    s.fix,
		},
		{
			name: "fmt", title: "Format",
			doc:    "Format MMS code. With dir, every .mms file of the project is formatted in place and the changed ones are listed. With source, the formatted text is returned.",
			props:  with(nil),
			writes: true,
			run:    s.fmt,
		},
		{
			name: "convert_json", title: "Convert vanilla JSON to MMS",
			doc: "Turn the JSON of one data pack or resource pack file into the MMS declaration that writes it, using shorthand where MMS has it. " +
				"Useful when you know how a file looks in vanilla JSON and not how to write it in MMS.",
			props: map[string]any{
				"kind": str("The declaration kind (`LootTable`, `Recipe`, `BlockTag`) or the file's folder in a pack (`loot_table`, `recipe`, `tags/block`, `worldgen/biome`)."),
				"name": str("The new declaration's name, with or without a namespace."),
				"json": str("The file's JSON text."),
				"dir":  str("The project the declaration is for, so that its names and namespace are taken into account. Optional."),
			},
			required: []string{"kind", "name", "json"},
			run:      s.convertJSON,
		},
		{
			name: "lookup", title: "Look up the language reference",
			doc: "Search the MMS reference: declaration kinds, helpers (such as `Clamp` or `SetCount`), the prelude's Templates, Blueprints and Types (such as `tell`, `CustomItem`), named types, and lint rules. " +
				"An exact name gives the full entry with parameters and an example; anything else lists the entries whose name or description matches.",
			props: map[string]any{
				"query": str("A name or a word to search for, such as `CustomItem`, `noise` or `unused`."),
				"in": map[string]any{"type": "string", "enum": lookupCategories,
					"description": "Search only one part of the reference. Default: all of it."},
				"limit": num("The most entries to list. Default 30."),
			},
			required: []string{"query"},
			run:      s.lookup,
		},
		{
			name: "vanilla", title: "Look up vanilla game data",
			doc: "Look up Minecraft's own data for the version MMS targets. With no arguments, list the registries (block, item, sound_event, worldgen/biome and so on) and the folders of vanilla data pack files. " +
				"With registry, list its ids, narrowed by query. With query alone, search every registry. With file, print one vanilla data pack file as JSON. " +
				"For a block, an exact id also gives its block state properties.",
			props: map[string]any{
				"registry": str("A registry name, such as `block`, `item`, `sound_event` or `worldgen/biome`."),
				"query":    str("Text an id must contain, such as `copper`."),
				"file":     str("A vanilla data pack file, as folder/id: `loot_table/chests/simple_dungeon`, `worldgen/noise_settings/overworld`, `tags/block/logs`."),
				"limit":    num("The most ids to list. Default 100."),
			},
			run: s.vanilla,
		},
		{
			name: "biomes", title: "Explain a dimension's biomes",
			doc: "Sample a Dimension's biome source the way the editors' biome map does, and explain it: the biomes in an area and their share, the climate at its centre with the nearest multi-noise entries and the parameters each one misses, " +
				"and with biome, where that biome wins or, if it wins nowhere, where it comes closest, what beats it there, and which climate parameters never reach its ranges. " +
				"Needs a noise generator with its biomes listed inline: a multi-noise preset (`preset = overworld`) cannot be sampled. Runs the editors' preview code, so it needs Node.js (or Bun or Deno).",
			props: with(map[string]any{
				"dimension": str("The Dimension's id or name. Optional when the project has one."),
				"biome":     str("A biome to explain, such as `minecraft:cherry_grove` or one of the project's."),
				"seed":      str("The world seed: a number, or text hashed the way the game does. Default 0."),
				"x":         num("The centre's x. Default 0."),
				"z":         num("The centre's z. Default 0."),
				"y":         num("The height biomes are sampled at. Default 64."),
				"radius":    num("Half the side of the square sampled, in blocks: 128×128 samples cover it. Default 2048."),
			}),
			run: func(ctx context.Context, args json.RawMessage) (string, error) {
				return s.preview(ctx, "biomes", args, true)
			},
		},
		{
			name: "density", title: "Sample worldgen density",
			doc: "Sample a noise, density function, noise settings' final density, or a terrain's top blocks, through the editors' preview code. " +
				"mode area: values over a top-down square at y (range, mean, share above 0, and a table). mode column: values down one column, with the heights where the value crosses 0 (for a final density: the ground and cave edges). " +
				"mode surface (noise settings or a Dimension): the top block of each column, with surface rules and aquifers applied: heights and which blocks show. " +
				"Ids may be the project's or vanilla's (`minecraft:overworld`). Needs Node.js (or Bun or Deno).",
			props: with(map[string]any{
				"id":     str("The noise, density function, noise settings or Dimension, by id or name."),
				"kind":   map[string]any{"type": "string", "enum": []string{"noise", "density_function", "noise_settings", "dimension"}, "description": "Which kind id is, when several have that name."},
				"mode":   map[string]any{"type": "string", "enum": []string{"area", "column", "surface"}, "description": "What to sample. Default area, or surface for a Dimension."},
				"seed":   str("The world seed. Default 0."),
				"x":      num("The centre's (or column's) x. Default 0."),
				"z":      num("The centre's (or column's) z. Default 0."),
				"y":      num("area: the height sampled (default 64). surface: show the top block at or below this height (default: the world's top)."),
				"radius": num("area and surface: half the side of the square, in blocks. Default 256 for area, 32 for surface (at most 64: surface builds every chunk)."),
				"size":   num("area and surface: the table's rows and columns, 2 to 32. Default 8."),
				"from":   num("column: the lowest y. Default -64."),
				"to":     num("column: the highest y. Default 320."),
				"step":   num("column: blocks between samples. Default 8."),
			}),
			required: []string{"id"},
			run: func(ctx context.Context, args json.RawMessage) (string, error) {
				return s.preview(ctx, "density", args, true)
			},
		},
		{
			name: "jigsaw", title: "Simulate a jigsaw structure",
			doc: "Grow a JigsawStructure with the simulator the editors' Simulation tab uses (best effort: no terrain matching). " +
				"With one run: every piece placed (depth, pool, rotation, origin, parent), each jigsaw block that placed nothing and why, and the pieces its pools list that never appeared. " +
				"With runs: how many seeds place each piece, pieces per run, and the jigsaw blocks that most often stay empty. Needs Node.js (or Bun or Deno).",
			props: with(map[string]any{
				"structure": str("The JigsawStructure's id or name. Optional when the project has one."),
				"seed":      num("The first seed. Default 0."),
				"runs":      num("How many seeds to run, from seed up, 1 to 200. Default 1."),
				"depth":     num("Grow this many pieces deep instead of the structure's size, 0 to 20."),
			}),
			run: func(ctx context.Context, args json.RawMessage) (string, error) {
				return s.preview(ctx, "jigsaw", args, false)
			},
		},
		{
			name: "docs", title: "Read the MMS docs",
			doc: "Read the MMS documentation. With no arguments, list the pages. With page, return that page as Markdown. With search, list the lines of every page that contain the text, under their headings.",
			props: map[string]any{
				"page":   str("A page path from the list, such as `guides/functions`."),
				"search": str("Text to find in the pages, such as `Storage`."),
			},
			run: s.docs,
		},
	}
}

func (s *Server) check(ctx context.Context, args json.RawMessage) (string, error) {
	var a struct {
		target
		Hints bool `json:"hints"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	_, r, err := s.load(ctx, a.target)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	c := report(&b, r, a.Hints)
	if c.total() == 0 {
		fmt.Fprintf(&b, "ok: no problems in %s\n", count(len(r.Units), "file"))
		return b.String(), nil
	}
	fmt.Fprintf(&b, "%s, %s", count(c.errors, "error"), count(c.warnings, "warning"))
	if c.notes > 0 {
		fmt.Fprintf(&b, ", %s", count(c.notes, "note"))
	}
	fmt.Fprintf(&b, " in %s\n", count(len(r.Units), "file"))
	if c.safe > 0 {
		fmt.Fprintf(&b, "%s a safe fix: the fix tool applies them\n", countHas(c.safe, "problem"))
	}
	return b.String(), nil
}

type counts struct{ errors, warnings, notes, safe int }

func (c counts) total() int { return c.errors + c.warnings + c.notes }

// report writes a result's problems the way a compiler prints them, each
// followed by its fixes and the other places that explain it.
func report(b *strings.Builder, r *project.Result, hints bool) counts {
	var c counts
	at := func(path string, off int) string {
		if u := r.Unit(path); u != nil {
			line, col := u.File.LineCol(off)
			return fmt.Sprintf("%s:%d:%d", path, line, col)
		}
		return path
	}
	for _, u := range r.Units {
		for _, d := range r.Diagnostics(u.Path) {
			if d.Severity == syntax.SeverityHint && !hints {
				continue
			}
			sev := "error"
			switch d.Severity {
			case syntax.SeverityWarning:
				sev = "warning"
				c.warnings++
			case syntax.SeverityInfo, syntax.SeverityHint:
				sev = "note"
				c.notes++
			default:
				c.errors++
			}
			line, col := u.File.LineCol(d.Sp.Start)
			fmt.Fprintf(b, "%s:%d:%d: %s: %s [%s]\n", u.Path, line, col, sev, d.Message, d.Code)
			text := u.File.LineText(line)
			width := max(1, min(d.Sp.End-d.Sp.Start, len([]rune(text))-col+1))
			fmt.Fprintf(b, "    %s\n    %s%s\n", strings.ReplaceAll(text, "\t", " "), strings.Repeat(" ", col-1), strings.Repeat("^", width))
			safe := false
			for _, f := range d.Fixes {
				label := "fix"
				if f.Safe {
					label, safe = "safe fix", true
				}
				fmt.Fprintf(b, "  %s: %s\n", label, f.Title)
				for _, e := range f.Edits {
					l1, c1 := u.File.LineCol(e.Sp.Start)
					l2, c2 := u.File.LineCol(e.Sp.End)
					switch {
					case e.Sp.Start == e.Sp.End:
						fmt.Fprintf(b, "      at %d:%d insert %q\n", l1, c1, e.Text)
					case e.Text == "":
						fmt.Fprintf(b, "      delete %d:%d to %d:%d\n", l1, c1, l2, c2)
					default:
						fmt.Fprintf(b, "      replace %d:%d to %d:%d with %q\n", l1, c1, l2, c2, e.Text)
					}
				}
			}
			if safe {
				c.safe++
			}
			for _, rel := range d.Related {
				fmt.Fprintf(b, "  see %s: %s\n", at(rel.Path, rel.Sp.Start), rel.Message)
			}
		}
	}
	// Files nothing uses have no line to point at.
	if paths, sev := r.UnusedAssets(); sev != syntax.SeverityHint || hints {
		for _, p := range paths {
			name := "note"
			switch sev {
			case syntax.SeverityError:
				name = "error"
				c.errors++
			case syntax.SeverityWarning:
				name = "warning"
				c.warnings++
			default:
				c.notes++
			}
			fmt.Fprintf(b, "%s: %s: %s [lint/unused-asset]\n", p, name, project.UnusedAssetMessage(p))
		}
	}
	return c
}

func count(n int, word string) string {
	if n == 1 {
		return "1 " + word
	}
	return fmt.Sprintf("%d %ss", n, word)
}

func countHas(n int, word string) string {
	if n == 1 {
		return "1 " + word + " has"
	}
	return fmt.Sprintf("%d %ss have", n, word)
}

func (s *Server) explain(ctx context.Context, args json.RawMessage) (string, error) {
	var a struct {
		target
		Name string `json:"name"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	p, r, err := s.load(ctx, a.target)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	if r.HasErrors() {
		b.WriteString("The project has errors, so some of this may be missing or wrong; run check.\n\n")
	}
	files := r.Emit()
	from := func(f project.OutFile) string {
		u := r.Unit(f.Unit)
		if u == nil {
			return ""
		}
		line, _ := u.File.LineCol(f.Span.Start)
		return fmt.Sprintf("%s:%d", f.Unit, line)
	}
	pack := func(f project.OutFile) string {
		if f.Pack == sema.ResourcePack {
			return "resource pack: "
		}
		return ""
	}
	if a.Name == "" {
		fmt.Fprintf(&b, "%s written:\n", count(len(files), "file"))
		for _, f := range files {
			fmt.Fprintf(&b, "  %s%s", pack(f), f.Path)
			switch {
			case f.Asset != "":
				fmt.Fprintf(&b, "  (copied from %s)", f.Asset)
			case from(f) != "":
				fmt.Fprintf(&b, "  (from %s)", from(f))
			default:
				b.WriteString("  (generated)")
			}
			b.WriteByte('\n')
		}
		b.WriteString("\nPass name to print the files of one declaration.\n")
		return b.String(), nil
	}

	// The declarations the name means, then the files written from
	// inside them: a file's span is its declaration's, or for a hoisted
	// inline value the value's.
	var decls []*project.Decl
	for _, d := range append(append([]*project.Decl(nil), r.Order...), r.Instances...) {
		if d.D == nil || d.D.Name == nil || d.Unit == nil {
			continue
		}
		id := d.ID()
		if id == a.Name || d.D.Name.Text == a.Name || strings.HasSuffix(id, ":"+a.Name) {
			decls = append(decls, d)
		}
	}
	inside := func(unit string, off int) bool {
		for _, d := range decls {
			if d.Unit.Path == unit && d.D.Sp.Contains(off) {
				return true
			}
		}
		return false
	}
	shown := 0
	for _, f := range files {
		if f.Asset != "" {
			continue
		}
		// Part of an output path is the fallback, for a name no
		// declaration has.
		if len(decls) > 0 {
			if f.Unit == "" || !inside(f.Unit, f.Span.Start) {
				continue
			}
		} else if !strings.Contains(f.Path, a.Name) {
			continue
		}
		shown++
		fmt.Fprintf(&b, "=== %s%s", pack(f), f.Path)
		if w := from(f); w != "" {
			fmt.Fprintf(&b, "  (from %s)", w)
		}
		fmt.Fprintf(&b, "\n%s\n\n", strings.TrimRight(string(f.Data), "\n"))
	}
	sess := &ide.Session{P: p, R: r}
	for _, u := range r.Units {
		for _, x := range sess.Expansions(u.Path) {
			if !inside(u.Path, x.Sp.Start) {
				continue
			}
			shown++
			line, _ := u.File.LineCol(x.Sp.Start)
			fmt.Fprintf(&b, "=== %s at %s:%d expands to\n%s\n\n", x.Title, u.Path, line, x.Text)
		}
	}
	if shown == 0 {
		var names []string
		for _, d := range r.Order {
			if d.D != nil && d.D.Name != nil && d.Unit != nil {
				names = append(names, d.Kind.Name+" "+d.ID())
			}
		}
		fmt.Fprintf(&b, "Nothing is written by a declaration named %q, and no output path contains it.\n", a.Name)
		if len(names) > 0 {
			sort.Strings(names)
			if len(names) > 60 {
				names = append(names[:60], "…")
			}
			fmt.Fprintf(&b, "The project declares:\n  %s\n", strings.Join(names, "\n  "))
		}
	}
	return b.String(), nil
}

func (s *Server) test(ctx context.Context, args json.RawMessage) (string, error) {
	var a struct {
		target
		Run string `json:"run"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	var keep func(string) bool
	if a.Run != "" {
		re, err := regexp.Compile(a.Run)
		if err != nil {
			return "", fmt.Errorf("run: %w", err)
		}
		keep = re.MatchString
	}
	_, r, err := s.load(ctx, a.target)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	if r.HasErrors() {
		c := report(&b, r, false)
		fmt.Fprintf(&b, "%s; no tests run\n", count(c.errors, "error"))
		return "", errors.New(b.String())
	}
	res, err := r.RunTests(keep)
	if err != nil {
		return "", err
	}
	failed := 0
	for _, t := range res {
		if t.Pass {
			fmt.Fprintf(&b, "ok    %s\n", t.Name)
		} else {
			failed++
			sp := t.Span
			if t.Expect > 0 {
				sp = t.ExpectSpan
			}
			line, col := 0, 0
			if u := r.Unit(t.Unit); u != nil {
				line, col = u.File.LineCol(sp.Start)
			}
			fmt.Fprintf(&b, "FAIL  %s\n    %s:%d:%d: %s\n", t.Name, t.Unit, line, col, t.Message)
		}
		for _, said := range t.Said {
			fmt.Fprintf(&b, "    said: %s\n", said)
		}
	}
	switch {
	case len(res) == 0 && a.Run != "":
		fmt.Fprintf(&b, "no tests match %q\n", a.Run)
	case len(res) == 0:
		b.WriteString("no tests: the project has no Test declarations (see docs page reference/language/tests)\n")
	case failed > 0:
		fmt.Fprintf(&b, "%d of %s failed\n", failed, count(len(res), "test"))
	default:
		fmt.Fprintf(&b, "ok: %s passed\n", count(len(res), "test"))
	}
	return b.String(), nil
}

func (s *Server) fix(ctx context.Context, args json.RawMessage) (string, error) {
	var a target
	if err := parse(args, &a); err != nil {
		return "", err
	}
	p, err := a.open()
	if err != nil {
		return "", err
	}
	if p == nil {
		p = project.New(s.pack, project.DefaultConfig())
		p.SetFile(sourcePath, a.Source)
	}
	changed, n, err := p.FixAll(ctx)
	if err != nil {
		return "", err
	}
	if a.Dir == "" {
		text := a.Source
		if t, ok := changed[sourcePath]; ok {
			text = formatted(t)
		}
		return fmt.Sprintf("%s fixed. The text is now:\n\n%s", count(n, "problem"), text), nil
	}
	paths := make([]string, 0, len(changed))
	for path := range changed {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	var b strings.Builder
	for _, path := range paths {
		if err := os.WriteFile(filepath.Join(a.Dir, filepath.FromSlash(path)), []byte(formatted(changed[path])), 0o644); err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "fixed %s\n", path)
	}
	fmt.Fprintf(&b, "%s fixed in %s\n", count(n, "problem"), count(len(paths), "file"))
	return b.String(), nil
}

func formatted(src string) string {
	if out, err := format.Source(src); err == nil {
		return out
	}
	return src
}

func (s *Server) fmt(_ context.Context, args json.RawMessage) (string, error) {
	var a target
	if err := parse(args, &a); err != nil {
		return "", err
	}
	if a.Dir == "" {
		if a.Source == "" {
			return "", errors.New("give dir (a project folder) or source (the text of a .mms file)")
		}
		out, err := format.Source(a.Source)
		if err != nil {
			return "", fmt.Errorf("cannot format: %w", err)
		}
		return out, nil
	}
	if a.Source != "" {
		return "", errors.New("give dir or source, not both")
	}
	cfg, err := project.LoadConfig(a.Dir)
	if err != nil {
		return "", err
	}
	paths, err := cfg.SourceFiles(a.Dir)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	n, failed := 0, 0
	for _, rel := range paths {
		if !strings.HasSuffix(rel, ".mms") {
			continue
		}
		full := filepath.Join(a.Dir, filepath.FromSlash(rel))
		src, err := os.ReadFile(full)
		if err != nil {
			return "", err
		}
		out, err := format.Source(string(src))
		if err != nil {
			failed++
			fmt.Fprintf(&b, "%s: cannot format: %v\n", rel, err)
			continue
		}
		if out == string(src) {
			continue
		}
		if err := os.WriteFile(full, []byte(out), 0o644); err != nil {
			return "", err
		}
		n++
		fmt.Fprintf(&b, "formatted %s\n", rel)
	}
	fmt.Fprintf(&b, "%s formatted\n", count(n, "file"))
	if failed > 0 {
		return "", errors.New(b.String())
	}
	return b.String(), nil
}

func (s *Server) convertJSON(ctx context.Context, args json.RawMessage) (string, error) {
	var a struct {
		Kind string `json:"kind"`
		Name string `json:"name"`
		JSON string `json:"json"`
		Dir  string `json:"dir"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	kinds := sema.NewKinds(s.pack)
	k, ok := kinds.Lookup(a.Kind)
	if !ok {
		k, ok = kinds.ForFolder(strings.Trim(a.Kind, "/"))
	}
	if !ok || !k.IsResource() {
		return "", fmt.Errorf("%q is not a kind that writes a JSON file; lookup lists the kinds, and the folder of each", a.Kind)
	}
	if a.Name == "" {
		return "", errors.New("name is empty")
	}
	v, err := jsonv.Parse([]byte(a.JSON))
	if err != nil {
		return "", fmt.Errorf("not valid JSON: %w", err)
	}
	p := project.New(s.pack, project.DefaultConfig())
	if a.Dir != "" {
		if p, err = project.Load(a.Dir); err != nil {
			return "", err
		}
	}
	ns := p.Config.Namespace
	id := a.Name
	if !strings.Contains(id, ":") {
		id = ns + ":" + id
	}
	var src string
	if k.TagOf != "" {
		src = importer.TagDecl(k, ns, id, v)
	} else {
		src = importer.Decl(k, ns, id, v)
	}
	return importer.TidyIn(ctx, p, ns, src)
}
