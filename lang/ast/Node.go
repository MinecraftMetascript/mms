package ast

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Node interface {
	Children() []Node
	GetLocation() *SourceLocation
}

type BaseNode struct {
	Location *SourceLocation `json:"location"`
}

func (n *BaseNode) GetLocation() *SourceLocation {
	return n.Location
}

func (n *BaseNode) Children() []Node {
	return []Node{}
}

type HelpfulNode interface {
	Node
	GetHelp() string
}
type CompletableNode interface {
	Complete(fileSource string, position protocol.Position, triggerChar *string, symbols map[string]*Namespace) []protocol.CompletionItem
}

type ExtractableNode interface {
	ExtractInlineSymbols() []Symbol
}
