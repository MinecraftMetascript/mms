package lib

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FileTreeLike struct {
	Name     string                   `json:"name"`
	IsDir    bool                     `json:"isDir"`
	Content  string                   `json:"content,omitempty"`
	Children map[string]*FileTreeLike `json:"children,omitempty"`
	Data     any                      `json:"data"`
}

func (ft *FileTreeLike) MkDir(name string, data any) *FileTreeLike {
	parts := strings.Split(name, "/")
	if len(parts) == 1 {
		if _, ok := ft.Children[name]; ok {
			return ft.Children[name]
		}
		newDir := NewDirLike(name, data)
		ft.AddChild(newDir)
		return newDir
	}
	current := ft
	for _, part := range parts[:len(parts)-1] {
		if part == "" {
			continue
		}
		if _, ok := current.Children[part]; !ok {
			current = current.MkDir(part, nil)
		}
	}
	return current.MkDir(parts[len(parts)-1], data)
}

func (ft *FileTreeLike) MkFile(name string, content string, data any) *FileTreeLike {
	parts := strings.Split(name, "/")
	if len(parts) == 1 {
		if _, ok := ft.Children[name]; ok {
			return ft.Children[name]
		}
		newFile := NewFileLike(name, content, data)
		ft.AddChild(newFile)
		return newFile
	}
	current := ft
	for _, part := range parts[:len(parts)-1] {
		if part == "" {
			continue
		}
		if _, ok := current.Children[part]; !ok {
			current = current.MkDir(part, nil)
		}
	}
	return current.MkFile(parts[len(parts)-1], content, data)
}

func (ft *FileTreeLike) Get(name string) *FileTreeLike {
	if name == "" {
		return ft
	}
	parts := strings.Split(name, "/")
	current := ft
	for _, part := range parts {
		if part == "" {
			continue
		}
		if current.Children == nil {
			return nil
		}
		if child, ok := current.Children[part]; ok {
			current = child
		} else {
			return nil
		}
	}
	return current
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

func (ft *FileTreeLike) PrintDebug() {
	data, err := json.MarshalIndent(ft, "", "  ")
	if err != nil {
		fmt.Println("ERROR marshalling:", err)
	} else {
		fmt.Println(string(data))
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
