package main

import (
	"fmt"

	"github.com/agent-e11/prime-algo-course/day1"
)

func main() {
	n := day1.BinaryNode[int]{
		Value: 7,
		Left: &day1.BinaryNode[int]{
			Value: 23,
			Left: &day1.BinaryNode[int]{
				Value: 5,
				Left: nil,
				Right: nil,
			},
			Right: &day1.BinaryNode[int]{
				Value: 4,
				Left: nil,
				Right: nil,
			},
		},
		Right: &day1.BinaryNode[int]{
			Value: 3,
			Left: &day1.BinaryNode[int]{
				Value: 18,
				Left: nil,
				Right: nil,
			},
			Right: &day1.BinaryNode[int]{
				Value: 21,
				Left: nil,
				Right: nil,
			},
		},
	}

	fmt.Printf("day1.PreOrderSearch(n): %v\n", day1.PreOrderSearch(n))
	fmt.Printf("day1.InOrderSearch(n): %v\n", day1.InOrderSearch(n))
	fmt.Printf("day1.PostOrderSearch(n): %v\n", day1.PostOrderSearch(n))
}
