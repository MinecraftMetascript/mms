package project

import (
	"log"

	"github.com/minecraftmetascript/mms/lang/ast"
	"github.com/minecraftmetascript/mms/lib"
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
		// TODO: Ensure that we are properly deleting all symbols from an
		// existing file before we re-parse it.
		for _, decls := range p.symbols {
			for name, decl := range decls.AllDecls() {
				if decl.GetLocation().Filename == path {
					decls.Delete(name)
				} else {
					log.Printf("Not deleting %s from %s (it is located in %s)", name, path, decl.GetNameLocation().Filename)
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
			nsDir.Merge(declFs)
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
