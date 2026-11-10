package mcp

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"regexp"
	"sort"
	"strings"

	"github.com/minecraftmetascript/mms/engine/prelude"
)

// docFiles is a copy of the website's hand-written pages, site/ in the
// repository, which go:embed cannot reach from this module. TestDocsInSync
// keeps the copy the same as site/ and, run with -update, writes it.
//
//go:embed docs
var docFiles embed.FS

// docGroups is the reading order of the pages' folders, as the website's
// sidebar has them.
var docGroups = []string{"start", "guides", "concepts", "reference/language"}

// docPage is one page: Path is its path under /docs on the website.
type docPage struct {
	Path, Title, Summary, Text string
}

var heading = regexp.MustCompile(`(?m)^#\s+(.+)$`)

func docPages() []docPage {
	var out []docPage
	_ = fs.WalkDir(docFiles, "docs", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".md") {
			return err
		}
		b, _ := docFiles.ReadFile(p)
		text := string(b)
		pg := docPage{Path: strings.TrimSuffix(strings.TrimPrefix(p, "docs/"), ".md"), Text: text}
		pg.Title = pg.Path
		if m := heading.FindStringSubmatch(text); m != nil {
			pg.Title = m[1]
		}
		for _, para := range strings.Split(heading.ReplaceAllString(text, ""), "\n\n") {
			para = strings.TrimSpace(para)
			if para != "" && !strings.HasPrefix(para, "```") && !strings.HasPrefix(para, "|") && !strings.HasPrefix(para, "<") && !strings.HasPrefix(para, "- ") {
				pg.Summary = firstSentence(para)
				break
			}
		}
		out = append(out, pg)
		return nil
	})
	rank := func(p docPage) int {
		for i, g := range docGroups {
			if strings.HasPrefix(p.Path, g+"/") {
				return i
			}
		}
		return len(docGroups)
	}
	sort.SliceStable(out, func(i, j int) bool {
		if ri, rj := rank(out[i]), rank(out[j]); ri != rj {
			return ri < rj
		}
		return out[i].Path < out[j].Path
	})
	return out
}

func (s *Server) docs(_ context.Context, args json.RawMessage) (string, error) {
	var a struct {
		Page   string `json:"page"`
		Search string `json:"search"`
	}
	if err := parse(args, &a); err != nil {
		return "", err
	}
	pages := docPages()
	var b strings.Builder
	if a.Page != "" {
		want := strings.TrimSuffix(strings.Trim(strings.TrimPrefix(a.Page, "/docs"), "/"), ".md")
		for _, p := range pages {
			if p.Path == want {
				return p.Text, nil
			}
		}
		fmt.Fprintf(&b, "no page %q; the pages are:\n", a.Page)
		for _, p := range pages {
			fmt.Fprintf(&b, "  %s\n", p.Path)
		}
		return "", fmt.Errorf("%s", b.String())
	}
	if a.Search != "" {
		q := strings.ToLower(a.Search)
		hits := 0
		const most = 60
		for _, p := range pages {
			section, printed := "", ""
			wrotePage := false
			for _, line := range strings.Split(p.Text, "\n") {
				if strings.HasPrefix(line, "#") {
					section = strings.TrimSpace(strings.TrimLeft(line, "#"))
				}
				if !strings.Contains(strings.ToLower(line), q) {
					continue
				}
				if hits++; hits > most {
					continue
				}
				if !wrotePage {
					fmt.Fprintf(&b, "%s\n", p.Path)
					wrotePage = true
				}
				if section != printed {
					fmt.Fprintf(&b, "  § %s\n", section)
					printed = section
				}
				fmt.Fprintf(&b, "      %s\n", strings.TrimSpace(line))
			}
		}
		if hits == 0 {
			return fmt.Sprintf("No page mentions %q. The lookup tool searches the reference: kinds, helpers, Templates and lint rules.\n", a.Search), nil
		}
		if hits > most {
			fmt.Fprintf(&b, "… and %d more lines\n", hits-most)
		}
		b.WriteString("\nRead a page with page.\n")
		return b.String(), nil
	}
	b.WriteString("MMS documentation pages. Read one with page; read guides/resources and guides/functions before writing MMS.\n\n")
	for _, p := range pages {
		fmt.Fprintf(&b, "- %s: %s. %s\n", p.Path, p.Title, p.Summary)
	}
	b.WriteString("\nThe reference (every kind, helper, Template and lint rule) is in the lookup tool. The prelude's source, which defines CustomItem, Book, tell and the rest in MMS itself, is in the resources mms://prelude/<file>.\n")
	return b.String(), nil
}

// Resources are the docs pages and the prelude's source files.
const (
	docsScheme    = "mms://docs/"
	preludeScheme = "mms://prelude/"
)

func resourceList() []map[string]any {
	out := []map[string]any{}
	for _, p := range docPages() {
		out = append(out, map[string]any{
			"uri": docsScheme + p.Path, "name": p.Path, "title": p.Title, "description": p.Summary, "mimeType": "text/markdown",
		})
	}
	for _, f := range prelude.Files() {
		name := strings.TrimPrefix(f.Path, prelude.Prefix)
		out = append(out, map[string]any{
			"uri": preludeScheme + name, "name": "prelude/" + name, "title": "Prelude: " + name,
			"description": "MMS source of the prelude, which every project can use.", "mimeType": "text/plain",
		})
	}
	return out
}

func resourceMIME(uri string) string {
	if strings.HasPrefix(uri, docsScheme) {
		return "text/markdown"
	}
	return "text/plain"
}

func readResource(uri string) (string, bool) {
	if path, ok := strings.CutPrefix(uri, docsScheme); ok {
		for _, p := range docPages() {
			if p.Path == path {
				return p.Text, true
			}
		}
		return "", false
	}
	if name, ok := strings.CutPrefix(uri, preludeScheme); ok {
		for _, f := range prelude.Files() {
			if f.Path == prelude.Prefix+name {
				return f.Text, true
			}
		}
	}
	return "", false
}
