package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/day1"
)

func main() {
	s := day1.NewStack[int]()

	v, ok := s.Peek()

	fmt.Println("value, ok:", v, ok)

	for i := range 5 {
		fmt.Println("Adding", i)
		s.Push(i)
	}

	i := 0
	for {
		if i % 3 == 0 || i % 5 == 0 {
			fmt.Printf("(%d)(%d) Adding\n", i, s.Length)
			s.Push(i)
		} else {
			v, ok := s.Pop()
			if !ok {
				break
			}
			fmt.Printf("(%d)(%d) Next item: %d\n", i, s.Length, v)
		}

		i++
	}
	fmt.Println("Done")
}
