package lrucache

type Node[T any] struct {
	Value T
	Next *Node[T]
	Prev *Node[T]
}

func newNode[V any](value V) *Node[V] {
	return &Node[V]{
		Value: value,
		Next: nil,
		Prev: nil,
	}
}

type LRU[K comparable, V any] struct {
	length int
	capacity int
	head *Node[V]
	tail *Node[V]
	lookup map[K]*Node[V]
	reverseLookup map[*Node[V]]K
}

func New[K comparable, V any](capacity int) LRU[K, V] {
	return LRU[K, V]{
		length: 0,
		capacity: capacity,
		head: nil,
		tail: nil,
		lookup: make(map[K]*Node[V]),
		reverseLookup: make(map[*Node[V]]K),
	}
}

func (lru *LRU[K, V]) Update(key K, value V) {
	// check if it exists
	node, ok := lru.lookup[key]
	if !ok {
		node = newNode(value)
		lru.length++
		lru.prepend(node)
		lru.trimCache()

		lru.lookup[key] = node
		lru.reverseLookup[node] = key
	} else {
		// If the node exists
		// Move it to the front
		lru.detach(node)
		lru.prepend(node)
		// Update its value
		node.Value = value
	}
}

func (lru *LRU[K, V]) Get(key K) (value V, ok bool) {
	// Check the cache for existence
	node, ok := lru.lookup[key]
	if !ok {
		return
	}

	// update the value and move it to the front
	lru.detach(node)
	lru.prepend(node)

	// return the value
	value = node.Value
	ok = true
	return
}

func (lru *LRU[K, V]) detach(node *Node[V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if lru.head == node {
		lru.head = lru.head.Next
	}

	if lru.tail == node {
		lru.tail = lru.tail.Prev
	}

	node.Next, node.Prev = nil, nil
}

func (lru *LRU[K, V]) prepend(node *Node[V]) {
	if lru.head == nil {
		lru.head, lru.tail = node, node
		return
	}

	node.Next = lru.head
	lru.head.Prev = node

	lru.head = node
}

func (lru *LRU[K, V]) trimCache() {
	if lru.length <= lru.capacity {
		return
	}

	tail := lru.tail

	lru.detach(lru.tail)

	key, ok := lru.reverseLookup[tail]
	if !ok {
		panic("trimCache: tail was not present in reverseLookup")
	}

	delete(lru.lookup, key)
	delete(lru.reverseLookup, tail)
	lru.length--
}
