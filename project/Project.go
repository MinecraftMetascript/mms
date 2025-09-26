package project

type Project struct {
	files map[string]*File
}

func NewProject() *Project {
	return &Project{
		files: make(map[string]*File),
	}
}

func (p *Project) Parse() {}

func (p *Project) AddFile(path, content string) *File {
	f := &File{
		path:    path,
		content: content,
	}
	p.files[path] = f
	return f

}

func (p *Project) BuildFsLike(root string) error {
	return nil
}
