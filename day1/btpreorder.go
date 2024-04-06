// Binary Tree
package day1

type BinaryNode[T comparable] struct {
	Value T
	Parent *BinaryNode[T]
	Left *BinaryNode[T]
	Right *BinaryNode[T]
}

func PreOrderSearch(head BinaryNode[int]) []int {
	return *preWalkBinaryTree(&head, &[]int{})
}

func preWalkBinaryTree(curr *BinaryNode[int], path *[]int) *[]int {
	if curr == nil {
		return path
	}

	*path = append(*path, curr.Value)

	preWalkBinaryTree(curr.Left, path)
	preWalkBinaryTree(curr.Right, path)

	return path
}
