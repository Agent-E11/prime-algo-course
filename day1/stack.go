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

// Push

// Pop

// Peek
