package day1

func CompareBinaryTrees[T comparable](a *BinaryNode[T], b *BinaryNode[T]) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return CompareBinaryTrees(a.Left, b.Left) && CompareBinaryTrees(a.Right, b.Right)
}
