package day1

type SNode[T any] struct {
	value T
	prev *SNode[T]
}

type Stack[T any] struct {
	Length int
	head *SNode[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		Length: 0,
		head: nil,
	}
}

// TODO:

func (s *Stack[T]) Push(item T) {
	s.Length++

	node := &SNode[T]{
		value: item,
		prev: s.head,
	}

	s.head = node
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	
	if s.head == nil {
		ok = false
		return
	}

	s.Length--

	if s.Length == 0 {
		head := s.head
		s.head = nil

		value = head.value
		ok = true
		return
	}

	head := s.head

	s.head = s.head.prev

	value = head.value
	ok = true
	return
}

func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.head == nil {
		ok = false
		return
	}

	value = s.head.value
	ok = true
	return
}
