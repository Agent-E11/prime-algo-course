package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/heap"
)

func main() {
	h := heap.New[int]()

	for _, i := range []int{8, 2, 3, 7, 4, 0, 2, 9, 7, 5, 4, 0, 3, 4, 8, 0, 6, 4, 9, 2, 3, 4, 9, 2, 3, 4, 8, 0, 1, 0, 1} {
		h.Insert(i)
		h.Print()
	}

	fmt.Println(h.Delete())

	h.Print()
}
