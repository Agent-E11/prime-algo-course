// Binary Tree
package binarytree

func PreOrderSearch(head Node[int]) []int {
	return *preWalkBinaryTree(&head, &[]int{})
}

func preWalkBinaryTree(curr *Node[int], path *[]int) *[]int {
	if curr == nil {
		return path
	}

	*path = append(*path, curr.Value)

	preWalkBinaryTree(curr.Left, path)
	preWalkBinaryTree(curr.Right, path)

	return path
}
