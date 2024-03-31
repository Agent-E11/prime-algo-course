package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/day1"
)

func main() {
	q := day1.NewQueue[int]()

	v, ok := q.Peek()

	fmt.Println("value, ok:", v, ok)

	for i := range 10 {
		fmt.Println("Adding", i)
		q.Enqueue(i)
	}

	for {
		v, ok := q.Deque()
		if !ok {
			break
		}
		fmt.Println("Next item:", v)
	}
	fmt.Println("Done")
}
