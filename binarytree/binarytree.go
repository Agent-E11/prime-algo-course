package binarytree

import "fmt"

type BinaryNode[T comparable] struct {
	Value T
	Parent *BinaryNode[T]
	Left *BinaryNode[T]
	Right *BinaryNode[T]
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
