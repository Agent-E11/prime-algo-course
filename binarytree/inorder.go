package binarytree

func InOrderSearch(head BinaryNode[int]) []int {
	return *inWalkBinaryTree(&head, &[]int{})
}

func inWalkBinaryTree(curr *BinaryNode[int], path *[]int) *[]int {
	if curr == nil {
		return path
	}

	inWalkBinaryTree(curr.Left, path)
	*path = append(*path, curr.Value)
	inWalkBinaryTree(curr.Right, path)

	return path
}
