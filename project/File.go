package project

import (
	"log"
	"slices"

	"github.com/minecraftmetascript/mms/lang/ast"
)

type File struct {
	path        string
	content     string
	diagnostics ast.Diagnostics
	astNodes    map[ast.SourceLocation]ast.Node
}

func (f *File) ingestNodes(node ast.Node) {
	if node == nil {
		return
	}
	if node.GetLocation() == nil {
		return
	}
	f.astNodes[*node.GetLocation()] = node

	for _, child := range node.Children() {
		f.ingestNodes(child)
	}
}

func (f *File) Parse() (map[string]*ast.Namespace, error) {
	parser := NewParser(f.content, f.path)

	parser.parser.File()

	// Freeze diagnostics
	f.diagnostics = *parser.diagnostics

	for _, block := range parser.blocks {
		f.ingestNodes(block)
	}

	return parser.namespaces, nil
}

func (f *File) Path() string    { return f.path }
func (f *File) Content() string { return f.content }
func (f *File) Diagnostics() []ast.Diagnostic {
	return f.diagnostics.Get()
}

func (f *File) Update(content string) (map[string]*ast.Namespace, error) {
	f.content = content
	return f.Parse()
}

func (f *File) AST() map[ast.SourceLocation]ast.Node {
	return f.astNodes
}

func (f *File) NodesAtPosition(pos ast.Location) []ast.Node {
	candidates := make([]ast.Node, 0)
	for loc, v := range f.astNodes {
		if loc.ContainsLocation(pos) {
			candidates = append(candidates, v)
		}
	}

	slices.SortStableFunc(candidates, func(a, b ast.Node) int {
		aLen := a.GetLocation().Stop.Index - a.GetLocation().Start.Index
		bLen := b.GetLocation().Stop.Index - b.GetLocation().Start.Index

		return aLen - bLen
	})

	return candidates
}

func (f *File) NodesInRange(start, stop ast.Location) []ast.Node {
	candidates := make([]ast.Node, 0)

	for loc, v := range f.astNodes {
		log.Printf("Node at %s", loc)
		if loc.ContainsLocation(start) || loc.ContainsLocation(stop) {
			log.Printf("  Matched")
			candidates = append(candidates, v)
		}
		log.Printf("  Done")
	}
	slices.SortStableFunc(candidates, func(a, b ast.Node) int {
		aLen := a.GetLocation().Stop.Index - a.GetLocation().Start.Index
		bLen := b.GetLocation().Stop.Index - b.GetLocation().Start.Index

		return aLen - bLen
	})

	return candidates
}
