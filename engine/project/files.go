package project

import (
	"io/fs"
	"path"
	"sort"
	"strings"
)

// maxAssetFiles caps how many files AssetFiles lists, so a project
// folder holding a whole game install does not stall completion.
const maxAssetFiles = 2000

// AssetFiles lists the files in fsys, a project folder, whose extension
// is ext (compared without case), as sorted slash-separated paths. Like
// SourceFiles it leaves out the build outputs, hidden folders,
// node_modules and folders with their own mms.json, which are separate
// projects.
func (c Config) AssetFiles(fsys fs.FS, ext string) []string {
	if fsys == nil {
		return nil
	}
	var out []string
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			if p == "." {
				return err
			}
			return nil
		}
		if d.IsDir() {
			if p == "." {
				return nil
			}
			if p == c.Output || p == c.ResourceOutput() || strings.HasPrefix(d.Name(), ".") || d.Name() == "node_modules" {
				return fs.SkipDir
			}
			if st, err := fs.Stat(fsys, path.Join(p, ConfigFile)); err == nil && !st.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if strings.EqualFold(path.Ext(p), ext) {
			out = append(out, p)
			if len(out) >= maxAssetFiles {
				return fs.SkipAll
			}
		}
		return nil
	})
	sort.Strings(out)
	return out
}
