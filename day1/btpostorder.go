package day1

func PostOrderSearch(head BinaryNode[int]) []int {
	return *postWalkBinaryTree(&head, &[]int{})
}

func postWalkBinaryTree(curr *BinaryNode[int], path *[]int) *[]int {
	if curr == nil {
		return path
	}

	postWalkBinaryTree(curr.Left, path)
	postWalkBinaryTree(curr.Right, path)
	*path = append(*path, curr.Value)

	return path
}
