package reverse_string

func NewStack[T any]() *Stack[T] {
	stack := Stack[T]{
		items: []*T{},
	}
	return &stack
}

type Stack[T any] struct {
	items []*T
}

func (s *Stack[T]) Push(item *T) *Stack[T] {
	s.items = append(s.items, item)
	return s
}

func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
