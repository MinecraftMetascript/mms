package lib

type FileTreeLike struct {
	Name     string                   `json:"name"`
	IsDir    bool                     `json:"isDir"`
	Content  string                   `json:"content,omitempty"`
	Children map[string]*FileTreeLike `json:"children,omitempty"`
	Data     any                      `json:"data"`
}

func (ft *FileTreeLike) MkDir(name string, data any) *FileTreeLike {
	if _, ok := ft.Children[name]; ok {
		return ft.Children[name]
	}
	newDir := NewDirLike(name, data)
	ft.AddChild(newDir)
	return newDir
}

func (ft *FileTreeLike) MkFile(name string, content string, data any) *FileTreeLike {
	if _, ok := ft.Children[name]; ok {
		return ft.Children[name]
	}
	newFile := NewFileLike(name, content, data)
	ft.AddChild(newFile)
	return newFile
}

func (ft *FileTreeLike) AddChild(child *FileTreeLike) {
	if ft.IsDir {
		ft.Children[child.Name] = child
	}
}

func (ft *FileTreeLike) SetContent(content string) {
	if !ft.IsDir {
		ft.Content = content
	}
}

func NewFileLike(name string, content string, data any) *FileTreeLike {
	return &FileTreeLike{
		Name:    name,
		IsDir:   false,
		Content: content,
		Data:    data,
	}
}

func NewDirLike(name string, data any) *FileTreeLike {
	return &FileTreeLike{
		Name:     name,
		IsDir:    true,
		Children: make(map[string]*FileTreeLike),
		Data:     data,
	}
}
