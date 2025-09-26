package ast

import "github.com/antlr4-go/antlr/v4"

type NodeFactory interface {
	// Match returns true if the context should be handled by the factory
	Match(ctx antlr.ParserRuleContext) bool

	ConstructNode(ctx antlr.ParserRuleContext) Node
}

type FactoryRegistry interface {
	Register(n NodeFactory)
	Construct(ctx antlr.ParserRuleContext)
}

type factoryRegistryImpl struct {
	factories []NodeFactory
}

func (f *factoryRegistryImpl) Register(n NodeFactory) {
	f.factories = append(f.factories, n)
}
func (f *factoryRegistryImpl) Construct(ctx antlr.ParserRuleContext) Node {
	for _, n := range f.factories {
		if n.Match(ctx) {
			return n.ConstructNode(ctx)
		}
	}
	return nil
}

var Registry = factoryRegistryImpl{
	factories: []NodeFactory{},
}
