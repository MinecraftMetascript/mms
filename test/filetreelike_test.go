package test

import (
	"github.com/minecraftmetascript/mms/lib"
	"testing"
)

func TestFileTreeLike_MkDir(t *testing.T) {
	tree := &lib.FileTreeLike{}

	dir := tree.MkDir("test_dir", nil)
	if dir == nil {
		t.Error("MkDir returned nil")
		return
	}

	if dir.Name != "test_dir" {
		t.Errorf("Expected dir name 'test_dir', got '%s'", dir.Name)
	}
}

func TestFileTreeLike_MkFile(t *testing.T) {
	tree := &lib.FileTreeLike{}

	file := tree.MkFile("test.txt", "content", nil)
	if file == nil {
		t.Error("MkFile returned nil")
		return
	}

	if file.Name != "test.txt" {
		t.Errorf("Expected file name 'test.txt', got '%s'", file.Name)
	}

	if file.Content != "content" {
		t.Errorf("Expected content 'content', got '%s'", file.Content)
	}
}

func TestFileTreeLike_Nested(t *testing.T) {
	tree := &lib.FileTreeLike{}

	dir := tree.MkDir("parent", nil)
	child := dir.MkDir("child", nil)
	file := child.MkFile("test.txt", "content", nil)

	if file == nil {
		t.Error("Nested file creation failed")
	}

	if child.Name != "child" {
		t.Errorf("Expected child dir name 'child', got '%s'", child.Name)
	}
}

func TestFileTreeLike_DirPath(t *testing.T) {
	tree := lib.NewDirLike("root", nil)
	tree.MkDir("a/b/c", nil)
	if tree.Get("a/b/c").Name != "c" {
		t.Error("Path failed")
	}
}

func TestFileTreeLike_FilePath(t *testing.T) {
	tree := lib.NewDirLike("root", nil)
	tree.MkFile("a/b/c/test.txt", "content", nil)
	file := tree.Get("a/b/c/test.txt")
	if file == nil || file.Name != "test.txt" || file.Content != "content" {
		t.Error("Path failed")
	}

}
