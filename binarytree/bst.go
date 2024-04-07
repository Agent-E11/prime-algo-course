package binarytree

import "slices"

func (t *BinaryNode[T]) isBSTRange(mn T, mx T, ignoreMn bool, ignoreMx bool) bool {

	if t.Value > mx && !ignoreMx {
		return false
	}
	if t.Value <= mn && !ignoreMn {
		return false
	}

	if !t.Left.isBSTRange(mn, t.Value, ignoreMn, false) {
		return false
	}
	if !t.Right.isBSTRange(t.Value, mx, false, ignoreMx) {
		return false
	}

	return true
}

func (t *BinaryNode[T]) IsBinarySearchTree() bool {
	return slices.IsSorted(t.InOrderSearch())
}
