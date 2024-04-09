package binarytree

import (
	"cmp"
	"fmt"
	"strings"
)

type Node[T cmp.Ordered] struct {
	Value T
	Parent *Node[T]
	Left *Node[T]
	Right *Node[T]
}

func New[T cmp.Ordered]() *Node[T] {
	return &Node[T]{}
}

func FromValue[T cmp.Ordered](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func From[T cmp.Ordered](value T, left *Node[T], right *Node[T]) (node *Node[T]) {

	node = &Node[T]{
		Value: value,
	}

	if left != nil {
		left.Parent = node
	}
	if right != nil {
		right.Parent = node
	}

	node.Left, node.Right = left, right

	return
}

func (t *Node[T]) rawPrint(i int) {
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

func (t *Node[T]) Print() {
	t.rawPrint(0)
	fmt.Println()
}

func (t *Node[T]) rawDebug(i int) {
	indent := strings.Repeat("  ", i)

	fmt.Printf("%v (%p) {\n", t.Value, t)

	fmt.Print(indent + "  ", "Parent: ")

	if t.Parent != nil {
		fmt.Printf("%p,\n", t.Parent)
	} else {
		fmt.Print("nil,\n")
	}

	fmt.Print(indent + "  ", "Left: ")

	if t.Left != nil {
		t.Left.rawDebug(i + 1)
		fmt.Println()
	} else {
		fmt.Print("nil,\n")
	}

	fmt.Print(indent + "  ", "Right: ")

	if t.Right != nil {
		t.Right.rawDebug(i + 1)
		fmt.Println()
	} else {
		fmt.Print("nil,\n")
	}

	fmt.Printf(indent + "}")
	if i > 0 { fmt.Print(",") }
}

func (t *Node[T]) Debug() {
	t.rawDebug(0)
	fmt.Println()
}
