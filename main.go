package main

import (
	//	"fmt"

	"fmt"

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

	fmt.Println(n.IsBinarySearchTree())
	fmt.Println(n.InOrderSearch())
}
