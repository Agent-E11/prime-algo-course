package day1

type ArrayList[T any] struct {
	array []T
	length int
	capacity int
}

// TODO: Methods

func NewArrayList[T any]() ArrayList[T] {
	return ArrayList[T]{
		array: []T{},
		length: 0,
		capacity: 5,
	}
}

// Get
func (a *ArrayList[T]) Get(idx int) (value T, ok bool) {
	if idx >= a.length {
		ok = false
		return
	}

	value = a.array[idx]
	ok = true
	return
}

// Push

// Pop

// Enqueue

// Deque

// Insert

// Delete

// Length
