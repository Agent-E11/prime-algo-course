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

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (l *LinkedList[T]) Prepend(item T) {
	node := &LNode[T]{
		value: item,
		prev:  nil,
		next:  l.head,
	}

	l.Length++

	if l.Length == 0 {
		l.head = node
		l.tail = node
		return
	}

	l.head.prev = node
	l.head = node
}

func (l *LinkedList[T]) Append(item T) {
	node := &LNode[T]{
		value: item,
		prev:  l.tail,
		next:  nil,
	}

	l.Length++

	if l.Length == 0 {
		l.head = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) InsertAt(item T, idx int) error {
	if idx > l.Length || idx < 0 {
		return OutOfBoundsErr
	}

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
	}

	// Set new node's pointers
	node.next = currNode
	node.prev = currNode.prev

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
