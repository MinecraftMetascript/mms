package ast

import "github.com/minecraftmetascript/mms/lib"

type Node interface{}

type HelpfulNode interface{}
type CompletableNode interface{}

type Symbol interface {
	Kind() string
	Export(name string) *lib.FileTreeLike
}
