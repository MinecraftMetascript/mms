package project

type File struct {
	path    string
	content string
}

func (f *File) Parse() error {
	parser := NewParser(f.content, f.path)

	parser.parser.File()
	return nil
}

func (f *File) Path() string    { return f.path }
func (f *File) Content() string { return f.content }
