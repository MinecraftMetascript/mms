package lib

func NewStack[K any]() *Stack[K] {
	return &Stack[K]{
		items: make([]K, 0),
	}
}

type Stack[K any] struct {
	items []K
}

func (s *Stack[K]) Push(item K) {
	s.items = append(s.items, item)
}

func (s *Stack[K]) Pop() K {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[K]) Peek() K {
	return s.items[len(s.items)-1]
}

func (s *Stack[K]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[K]) Size() int {
	return len(s.items)
}
