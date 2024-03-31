package day1

type Node[T any] struct {
	value T
	next *Node[T]
}

type Queue[T any] struct {
	Length int
	head *Node[T]
	tail *Node[T]
}

// Create a new empty Queue
func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		Length: 0,
		head: nil,
		tail: nil,
	}
}

// Add a new item to the end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.Length++

	node := &Node[T]{
		value: item,
		next: nil,
	}

	if q.tail == nil {
		q.head = node
		q.tail = node
		
		return
	}

	q.tail.next = node
	q.tail = node
}

// Remove and return the item in the front of the queue.
// If the queue is empty, then the ok value will be false
func (q *Queue[T]) Deque() (value T, ok bool) {

	// If there is no head, return nothing
	if q.head == nil {
		ok = false
		return
	}

	q.Length--

	// Save the current head
	head := q.head
	// Set the head to the next node, if there is only one node currently, then the next node will be nil
	q.head = q.head.next

	// Set the previous head's next to nil (should be done automatically)
	head.next = nil

	value = head.value
	ok = true

	return
}

// Look at the next item in the queue without removing it.
// If the queue is empty, then the ok value will be false
func (q *Queue[T]) Peek() (value T, ok bool) {

	// If there is no head, return nothing
	if q.head == nil {
		ok = false
		return
	}

	value = q.head.value
	ok = true

	return
}
