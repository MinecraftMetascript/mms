package lang

import (
	"encoding/json"
	"os"

	"github.com/minecraftmetascript/mms/lang/worldgen/noise"
	"github.com/minecraftmetascript/mms/lang/worldgen/surface"
	"github.com/minecraftmetascript/mms/lib"
)

type MMSProject struct {
	files   map[string]*MMSFile
	symbols map[string]*lib.Namespace
}

func NewMMSProject() *MMSProject {
	namespaces := make(map[string]*lib.Namespace)
	return &MMSProject{
		files:   make(map[string]*MMSFile),
		symbols: namespaces,
	}
}

func (p *MMSProject) AddFile(filepath, content string) error {
	f := p.ParseFile(filepath, content)

	p.files[filepath] = f
	return nil
}

func (p *MMSProject) GetFile(filepath string) *MMSFile {
	for _, file := range p.files {
		if file.path == filepath {
			return file
		}
	}
	return nil
}

func (p *MMSProject) GetFiles() []MMSFile {
	return []MMSFile{}
}

func (p *MMSProject) ParseFiles(root string) error {
	stat, err := os.Stat(root)
	if err != nil {
		return err
	}
	files := make([]string, 0)
	if stat.IsDir() {
		files, err = lib.RecursiveFilesInDir(root)
		if err != nil {
			return err
		}
	} else {
		files = append(files, root)
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		err = p.AddFile(file, string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *MMSProject) Export() lib.FileTreeLike {
	root := &lib.FileTreeLike{
		Name:     "mms.dist",
		IsDir:    true,
		Children: map[string]*lib.FileTreeLike{},
	}
	for namespace, symbols := range p.symbols {
		namespaceDir := root.MkDir(namespace, nil)
		surface.Export(symbols, namespaceDir)
		noise.Export(symbols, namespaceDir)
	}

	return *root
}

func ProjectSymbols[T json.Marshaler](p *MMSProject) map[string]map[string]lib.Symbol[T] {
	out := make(map[string]map[string]lib.Symbol[T])
	for namespace, symbols := range p.symbols {
		namespaceSymbols := lib.AllOf[T](symbols)
		out[namespace] = namespaceSymbols
	}

	return out
}
