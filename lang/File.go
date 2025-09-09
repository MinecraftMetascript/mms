package lang

import (
	"github.com/minecraftmetascript/mms/lang/grammar"
	"github.com/minecraftmetascript/mms/lang/traversal"
)

type File struct {
	Content     string
	Path        string
	Project     *Project
	Diagnostics []traversal.Diagnostic
	Script      *grammar.ScriptContext
}

func (f *File) Parse() error {
	err := f.Project.GlobalScope.PurgeFile(f.Path)
	f.Diagnostics = make([]traversal.Diagnostic, 0)
	if err != nil {
		return err
	}
	parser := traversal.NewParser(f.Content, f.Path, f.Project.GlobalScope, &f.Diagnostics)
	script, err := parser.Parse()
	if err != nil {
		return err
	}
	f.Script = script

	return nil
}
