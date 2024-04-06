package main

import (
//	"fmt"

	"github.com/agent-e11/prime-algo-course/binarytree"
)

func main() {
	n := binarytree.BinaryNode[int]{
		Value: 7,
		Left: &binarytree.BinaryNode[int]{
			Value: 23,
			Left: &binarytree.BinaryNode[int]{
				Value: 5,
				Left: nil,
				Right: &binarytree.BinaryNode[int]{
					Value: 4,
					Left: nil,
					Right: nil,
				},
			},
			Right: &binarytree.BinaryNode[int]{
				Value: 4,
				Left: nil,
				Right: nil,
			},
		},
		Right: &binarytree.BinaryNode[int]{
			Value: 3,
			Left: &binarytree.BinaryNode[int]{
				Value: 18,
				Left: nil,
				Right: nil,
			},
			Right: &binarytree.BinaryNode[int]{
				Value: 21,
				Left: &binarytree.BinaryNode[int]{
					Value: 18,
					Left: nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	n.Print()

//	fmt.Printf("binarytree.BreadthFirstSearch(&n, 1): %v\n", binarytree.BreadthFirstSearch(&n, 1))
//	fmt.Printf("binarytree.BreadthFirstSearch(&n, 3): %v\n", binarytree.BreadthFirstSearch(&n, 3))
//	fmt.Printf("binarytree.BreadthFirstSearch(&n, 4): %v\n", binarytree.BreadthFirstSearch(&n, 4))
//	fmt.Printf("binarytree.BreadthFirstSearch(&n, 6): %v\n", binarytree.BreadthFirstSearch(&n, 6))
//	fmt.Printf("binarytree.BreadthFirstSearch(&n, 9): %v\n", binarytree.BreadthFirstSearch(&n, 9))
}
