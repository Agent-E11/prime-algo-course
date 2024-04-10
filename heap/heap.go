package heap

import (
	"cmp"
	"fmt"
	"math"
)

type Heap[T cmp.Ordered] struct {
	Length int
	data []T
}

func New[T cmp.Ordered]() Heap[T] {
	return Heap[T]{
		Length: 0,
		data: []T{},
	}
}

func (h *Heap[T]) Insert(value T) {
	h.data = append(h.data, value)

	h.heapifyUp(h.Length)

	h.Length++
}

func (h *Heap[T]) Delete() (value T, ok bool) {
	if h.Length == 0 {
		ok = false
		return
	}

	value = h.data[0]

	h.Length--

	if h.Length == 0 {
		ok = true
		h.data = []T{}
		return
	}

	h.data[0] = h.data[h.Length]

	h.data = h.data[:h.Length]

	h.heapifyDown(0)

	ok = true

	return
}

func (h *Heap[T]) Print() {
	fmt.Println(h.data)

	var curr float64 = 0

	for i, v := range h.data {

		log := math.Floor(math.Log2(float64(i) + 1))
		if curr < log {
			fmt.Println()
			curr = log
		}
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println()
}

func (h *Heap[T]) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx >= len(h.data) || rIdx >= len(h.data){
		fmt.Println("returning")
		return
	}

	fmt.Printf("lIdx: %v\n", lIdx)
	fmt.Printf("rIdx: %v\n", rIdx)

	lVal := h.data[lIdx]
	rVal := h.data[rIdx]
	val := h.data[idx]

	if lVal >= rVal && val > rVal {
		h.data[idx] = rVal
		h.data[rIdx] = val

		h.heapifyDown(rIdx)
	} else if rVal > lVal && val > lVal {
		h.data[idx] = lVal
		h.data[lIdx] = val

		h.heapifyDown(lIdx)
	}
}

func (h *Heap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := parent(idx)
	parentVal := h.data[p]
	val := h.data[idx]

	if parentVal > val {
		h.data[idx] = parentVal
		h.data[p] = val
		h.heapifyUp(p)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx * 2 + 1
}

func rightChild(idx int) int {
	return idx * 2 + 2
}
