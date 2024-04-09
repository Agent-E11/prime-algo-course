package binarytree

func (t *Node[T]) InOrderSearch() []T {
	return *t.inWalkBinaryTree(&[]T{})
}

func (t *Node[T]) inWalkBinaryTree(path *[]T) *[]T {
	if t == nil {
		return path
	}

	if t.Left != nil {
		t.Left.inWalkBinaryTree(path)
	}
	*path = append(*path, t.Value)
	if t.Right != nil {
		t.Right.inWalkBinaryTree(path)
	}

	return path
}
