package lang

import (
	"fmt"

	"github.com/minecraftmetascript/mms/lang/traversal"
	"github.com/minecraftmetascript/mms/lib"
)
import _ "github.com/minecraftmetascript/mms/lang/constructs/worldgen/surface_rules"

type Project struct {
	Files       map[string]*File
	GlobalScope *traversal.Scope
}

func (p *Project) Diagnostics() []traversal.Diagnostic {
	diagnostics := make([]traversal.Diagnostic, 0)
	for _, file := range p.Files {
		diagnostics = append(diagnostics, file.Diagnostics...)
	}
	return diagnostics
}

func NewProject() *Project {
	return &Project{
		Files:       make(map[string]*File),
		GlobalScope: traversal.NewScope(),
	}
}

func (p *Project) AddFile(path, content string) *File {
	if p.Files[path] != nil {
		out := p.Files[path]
		out.Content = content
	} else {
		p.Files[path] = &File{
			Content: content,
			Path:    path,
			Project: p,
		}
	}
	return p.Files[path]
}

func (p *Project) BuildFsLike(rootName string) *lib.FileTreeLike {
	root := lib.NewDirLike(rootName, nil)
	for name, symbol := range p.GlobalScope.Symbols() {
		ref := symbol.GetValue()
		err := ref.ExportSymbol(symbol, root)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error while exporting %s: %s", name, err.Error()))
		}
	}

	return root
}
