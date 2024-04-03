package day1

import (
	"errors"
)

var OutOfBoundsErr error = errors.New("Index out of bounds")

type LNode[T any] struct {
	value T
	prev  *LNode[T]
	next  *LNode[T]
}

type LinkedList[T any] struct {
	Length int
	head   *LNode[T]
	tail   *LNode[T]
}

// Create an empty LinkedList with the specified type
func NewLinkedList[T any]() LinkedList[T] {
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
	if l.Length == 0 {
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
	return
}

func (l *LinkedList[T]) RemoveAt(idx int) (value T, ok bool) {
	return
}

func (l *LinkedList[T]) Get(idx int) (value T, ok bool) {
	return
}
