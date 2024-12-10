package stack

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
