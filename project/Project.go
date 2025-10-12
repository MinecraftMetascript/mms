package project

import (
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lib"
	"github.com/samber/lo"
)

type Project struct {
	files map[string]*File

	symbols map[string]*ast.Namespace
}

func NewProject() *Project {
	return &Project{
		files:   make(map[string]*File),
		symbols: make(map[string]*ast.Namespace),
	}
}

func (p *Project) AddFile(path, content string) (*File, error) {
	f := &File{
		path:     path,
		content:  content,
		astNodes: make(map[ast.SourceLocation]ast.Node),
	}
	if p.files[path] != nil {
		for _, decls := range p.symbols {
			for name, decl := range decls.AllDecls() {
				if decl == nil || decl.GetLocation() == nil {
					// Symbol has unknown location, leave it untouched
					continue
				}
				if decl.GetLocation().Filename == path {
					decls.Delete(name)
				}
			}
		}

	}
	p.files[path] = f
	declarations, err := f.Parse()

	for ns, decls := range declarations {
		if existing, ok := p.symbols[ns]; ok {
			existing.Merge(decls)
			p.symbols[ns] = existing
		} else {
			p.symbols[ns] = decls
		}
	}

	return f, err
}

func (p *Project) BuildFsLike(root string) *lib.FileTreeLike {
	rootDir := lib.NewDirLike(root, nil)

	for ns, decls := range p.symbols {
		nsDir := rootDir.MkDir(ns, nil)
		for name, decl := range decls.AllDecls() {
			declFs := decl.ToFileTreeLike(name)
			if declFs != nil {
				nsDir.AddChild(declFs)
			}
		}
	}
	return rootDir
}

func (p *Project) File(path string) *File {
	if p.files[path] != nil {
		return p.files[path]
	}
	return nil
}

func (p *Project) Files() map[string]*File {
	return p.files
}

func (p *Project) Symbols() map[string]*ast.Namespace {
	return p.symbols
}

func (p *Project) AllDiagnostics() []ast.Diagnostic {
	return lo.Reduce(lo.Values(p.Files()), func(agg []ast.Diagnostic, item *File, index int) []ast.Diagnostic {
		return slices.Concat(agg, item.Diagnostics())
	}, make([]ast.Diagnostic, 0))
}
