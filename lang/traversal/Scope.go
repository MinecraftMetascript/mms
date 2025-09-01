package traversal

import (
	"errors"
	"fmt"
)

type Scope struct {
	symbols map[string]Symbol
	parent  *Scope
}

func NewScope() *Scope {
	return &Scope{
		symbols: make(map[string]Symbol),
		parent:  nil,
	}
}

func (s *Scope) NewScope() *Scope {
	return &Scope{
		symbols: make(map[string]Symbol),
		parent:  s,
	}
}

func (s *Scope) Get(name Reference) (Symbol, bool) {
	val, ok := s.symbols[name.String()]
	if !ok && s.parent != nil {
		if val, ok = s.parent.Get(name); !ok {
			return nil, false
		}
	}

	return val, ok
}

func (s *Scope) Register(value Symbol) error {
	if value == nil || value.GetValue() == nil {
		return errors.New(fmt.Sprintf("symbol is nil or has nil value: %v", value))
	}
	if _, ok := s.symbols[value.GetReference().String()]; ok {
		return errors.New(fmt.Sprintf("%s already exists", value.GetReference().String()))
	}
	s.symbols[value.GetReference().String()] = value
	return nil
}

func (s *Scope) Parent() *Scope {
	return s.parent
}

func (s *Scope) Symbols() map[string]Symbol {
	return s.symbols
}
