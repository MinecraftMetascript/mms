package lsp

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/minecraftmetascript/mms/engine/project"
)

// workspace is one MMS project the server checks: the files under a folder
// holding mms.json, analyzed apart from every other project, so examples
// side by side in one editor workspace do not see each other.
type workspace struct {
	// dir holds mms.json; "" for the default project.
	dir string
	// root is where the project's files are read from: dir, or for the
	// default project the editor's folder if it has no mms.json ("" in
	// the browser).
	root string
	// extra are more folders whose loose files the default project holds.
	extra []string
	proj  *project.Project
	res   *project.Result
}

// contains reports whether the .mms file at path belongs to w on disk.
func (w *workspace) contains(path string) bool {
	if !strings.HasSuffix(path, ".mms") {
		return false
	}
	for _, root := range append([]string{w.root}, w.extra...) {
		if root == "" {
			continue
		}
		rel, err := filepath.Rel(root, path)
		if err == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			return true
		}
	}
	return false
}

// configDir is the nearest folder from dir upward that holds mms.json, or
// "" if there is none.
func configDir(dir string) string {
	for dir != "" {
		if st, err := os.Stat(filepath.Join(dir, project.ConfigFile)); err == nil && !st.IsDir() {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

// discover builds the projects from the editor's folders: one per mms.json
// found in or above them, and the default project for .mms files under a
// folder no mms.json covers. Earlier projects are dropped.
func (s *Server) discover(defCfg project.Config) {
	s.projects = map[string]*workspace{}
	s.owners = map[string]*workspace{}
	def := &workspace{proj: project.New(s.pack, defCfg)}
	s.projects[""] = def
	var loose []string
	for _, root := range s.roots {
		if dir := configDir(root); dir != "" {
			s.load(dir) // the folder is inside a project
		} else {
			loose = append(loose, root)
		}
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				if path != root && (strings.HasPrefix(d.Name(), ".") || d.Name() == "node_modules") {
					return filepath.SkipDir
				}
				return nil
			}
			if d.Name() == project.ConfigFile {
				s.load(filepath.Dir(path))
			}
			return nil
		})
	}
	if len(loose) > 0 {
		// Asset helpers in loose files read from the first such folder.
		def.root, def.extra = loose[0], loose[1:]
		def.proj.Files = os.DirFS(def.root)
	}
	for _, root := range loose {
		// SourceFiles leaves out folders with their own mms.json.
		paths, err := defCfg.SourceFiles(root)
		if err != nil {
			s.logf("%v", err)
		}
		for _, rel := range paths {
			s.readFile(def, filepath.Join(root, rel))
		}
	}
}

// load returns the project for the folder dir holding mms.json, reading
// its config and files the first time.
func (s *Server) load(dir string) *workspace {
	if w := s.projects[dir]; w != nil {
		return w
	}
	cfg, err := project.LoadConfig(dir)
	if err != nil {
		s.logf("%v", err)
		cfg = project.DefaultConfig()
	}
	w := &workspace{dir: dir, root: dir, proj: project.New(s.pack, cfg)}
	// Asset helpers such as Texture("a.png") read project files.
	w.proj.Files = os.DirFS(dir)
	s.projects[dir] = w
	paths, err := cfg.SourceFiles(dir)
	if err != nil {
		s.logf("%v", err)
	}
	for _, rel := range paths {
		s.readFile(w, filepath.Join(dir, rel))
	}
	return w
}

func (s *Server) readFile(w *workspace, path string) {
	uri := pathURI(path)
	if s.owners[uri] != nil {
		return
	}
	if b, err := os.ReadFile(path); err == nil {
		w.proj.SetFile(uri, string(b))
		s.owners[uri] = w
	}
}

// at is the project holding the file uri, or nil if no project has it.
func (s *Server) at(uri string) *workspace { return s.owners[uri] }

// projectFor is the project the file uri belongs to, which it joins when
// it is opened: the one of the nearest mms.json above it, loaded if need
// be, or else the default project.
func (s *Server) projectFor(uri string) *workspace {
	if w := s.owners[uri]; w != nil {
		return w
	}
	if path := uriPath(uri); path != "" {
		if dir := configDir(filepath.Dir(path)); dir != "" {
			return s.load(dir)
		}
	}
	return s.projects[""]
}

// setFile puts text in the project uri belongs to and returns it.
func (s *Server) setFile(uri, text string) *workspace {
	w := s.projectFor(uri)
	w.proj.SetFile(uri, text)
	s.owners[uri] = w
	return w
}

// removeFile takes uri out of its project and returns the project.
func (s *Server) removeFile(uri string) *workspace {
	w := s.owners[uri]
	if w == nil {
		return nil
	}
	w.proj.RemoveFile(uri)
	delete(s.owners, uri)
	return w
}

// reload rebuilds every project from disk after an mms.json changed,
// appeared or went away, keeping the text of open and managed files.
func (s *Server) reload(ctx context.Context) {
	texts := map[string]string{}
	keep := func(uri string) {
		if w := s.owners[uri]; w != nil {
			texts[uri], _ = w.proj.Text(uri)
		}
	}
	for uri := range s.open {
		keep(uri)
	}
	for uri := range s.managed {
		keep(uri)
	}
	s.discover(s.projects[""].proj.Config)
	for uri, t := range texts {
		s.setFile(uri, t)
	}
	s.last = s.owners[s.lastURI]
	s.analyzeAll(ctx)
}

// analyzeAll re-checks every project.
func (s *Server) analyzeAll(ctx context.Context) {
	for _, w := range s.projects {
		s.analyze(ctx, w)
	}
}
