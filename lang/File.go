package lang

import (
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type File struct {
	Content     string
	Path        string
	Project     *Project
	Diagnostics []traversal.Diagnostic
}

func (f *File) Parse() {
	parser := traversal.NewParser(f.Content, f.Project.GlobalScope, &f.Diagnostics)
	_ = parser.Parse()
	// Aggregate into project-level diagnostics if present
	if f.Project != nil {
		f.Project.Diagnostics = append(f.Project.Diagnostics, f.Diagnostics...)
	}
}
