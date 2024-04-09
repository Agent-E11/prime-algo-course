package binarytree

import (
	"cmp"
	"fmt"
)

// TODO: Rename to "Node"

type BinaryNode[T cmp.Ordered] struct {
	Value T
	Parent *BinaryNode[T]
	Left *BinaryNode[T]
	Right *BinaryNode[T]
}

func New[T cmp.Ordered]() *BinaryNode[T] {
	return &BinaryNode[T]{}
}

func FromValue[T cmp.Ordered](value T) *BinaryNode[T] {
	return &BinaryNode[T]{Value: value}
}

func From[T cmp.Ordered](value T, left *BinaryNode[T], right *BinaryNode[T]) *BinaryNode[T] {
	return &BinaryNode[T]{
		Value: value,
		Left: left,
		Right: right,
	}
}

func (t *BinaryNode[T]) rawPrint(i int) {
	var opening, closing string
	switch i % 4 {
	case 0:
		opening = "("
		closing = ")"
	case 1:
		opening = "{"
		closing = "}"
	case 2:
		opening = "["
		closing = "]"
	case 3:
		opening = "<"
		closing = ">"
	}

	fmt.Print(t.Value)
	fmt.Printf(" %v", opening)

	if t.Left != nil {
		t.Left.rawPrint(i + 1)
	} else {
		fmt.Print("nil")
	}

	fmt.Print(", ")

	if t.Right != nil {
		t.Right.rawPrint(i + 1)
	} else {
		fmt.Print("nil")
	}

	fmt.Printf("%v", closing)
}

func (t *BinaryNode[T]) Print() {
	t.rawPrint(0)
	fmt.Println()
}
