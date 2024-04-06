package day1

import (
	"errors"
	"fmt"
)

var OutOfBoundsErr error = errors.New("Index out of bounds")

type LNode[T comparable] struct {
	value T
	prev  *LNode[T]
	next  *LNode[T]
}

type LinkedList[T comparable] struct {
	Length int
	head   *LNode[T]
	tail   *LNode[T]
}

// Create an empty LinkedList with the specified type
func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

// Add a new node to the front of the list, containing the item
func (l *LinkedList[T]) Prepend(item T) {
	// Create a new node that points forward to the current head
	node := &LNode[T]{
		value: item,
		prev:  nil,
		next:  l.head,
	}

	l.Length++

	// If the list is empty,
	// set the head and tail to the new node and return
	if l.Length == 0 {
		l.head = node
		l.tail = node
		return
	}

	// Set the previous head to point back to the new node
	l.head.prev = node
	// Move the head to the new node
	l.head = node
}

// Add a new node to the back of the list, containing the item
func (l *LinkedList[T]) Append(item T) {
	// Create a new node that points back to the current tail
	node := &LNode[T]{
		value: item,
		prev:  l.tail,
		next:  nil,
	}

	l.Length++

	// If the list is empty,
	// set the head and tail to the new node and return
	if l.Length == 1 {
		l.head = node
		l.tail = node
		return
	}

	// Set the previous tail to point forward to the new node
	l.tail.next = node
	// Move the tail to the new node
	l.tail = node
}

// Insert a new node containing item, so that it is now at the specified index.
// If the index is 0, then it is the same as list.Prepend(item). If the index is
// len(list), then it is the same as list.Append(item). If the index is greater
// than len(list), or less than 0, then day1.OutOfBoundsErr will be returned
func (l *LinkedList[T]) InsertAt(item T, idx int) error {
	// Ensure index is in bounds
	if idx > l.Length || idx < 0 {
		return OutOfBoundsErr
	}

	// If the index is the beginning, or one off the end,
	// prepend or append respectively
	if idx == l.Length {
		l.Append(item)
		return nil
	} else if idx == 0 {
		l.Prepend(item)
		return nil
	}

	l.Length++

	// Walk to node at the index
	currNode := l.head
	for range idx {
		currNode = currNode.next
	}

	node := &LNode[T]{
		value: item,
		// Set new node's pointers
		next: currNode,
		prev: currNode.prev,
	}

	// Change old nodes' pointers
	node.next.prev = node
	node.prev.next = node

	return nil
}

func (l *LinkedList[T]) Remove(item T) (value T, ok bool) {
	currNode := l.head

	// Loop over the entire list
	for i := 0; currNode != nil && i < l.Length; i++ {
		if currNode.value == item {
			break
		}

		currNode = currNode.next
	}

	// Did not find the element
	if currNode == nil {
		ok = false
		return
	}

	value, ok = l.removeNode(currNode)

	return
}

func (l *LinkedList[T]) RemoveAt(idx int) (value T, ok bool) {
	node := l.getAt(idx)

	if node == nil {
		ok = false
		return
	}

	value, ok = l.removeNode(node)

	return
}

func (l *LinkedList[T]) Get(idx int) (value T, ok bool) {
	node := l.getAt(idx)
	if node == nil {
		ok = false
		return
	}
	value = node.value
	ok = true
	return
}

func (l *LinkedList[T]) Println() {
	currNode := l.head

	fmt.Print("[ ")
	for currNode != nil {
		fmt.Printf("%v, ", currNode.value)
		currNode = currNode.next
	}
	fmt.Print("]\n")
}

func (l *LinkedList[T]) Debug() {
	currNode := l.head
	idx := 0
	fmt.Println("Linked List, length", l.Length)
	fmt.Println("[")
	for currNode != nil {
		fmt.Printf("  %v: { (%p)\n", idx, currNode)
		fmt.Printf("    v: %v\n", currNode.value)
		fmt.Printf("    p: %p %v\n", currNode.prev, currNode.prev)
		fmt.Printf("    n: %p %v\n", currNode.next, currNode.next)
		fmt.Print ("  },\n")

		currNode = currNode.next
		idx++
	}
	fmt.Println("]")
}

func (l *LinkedList[T]) removeNode(node *LNode[T]) (value T, ok bool) {
	l.Length--

	if l.Length == 0 {
		value = l.head.value
		ok = true
		l.head, l.tail = nil, nil
		return
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node == l.head {
		l.head = node.next
	}
	if node == l.tail {
		l.tail = node.prev
	}

	node.next, node.prev = nil, nil

	value = node.value
	ok = true

	return
}

func (l *LinkedList[T]) getAt(idx int) *LNode[T] {
	if idx >= l.Length || idx < 0 {
		return nil
	}

	currNode := l.head
	for range idx {
		currNode = currNode.next
	}

	return currNode
}
