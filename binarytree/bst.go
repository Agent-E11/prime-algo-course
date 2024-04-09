package binarytree

import "slices"

func (t *Node[T]) isBSTRange(mn T, mx T, ignoreMn bool, ignoreMx bool) bool {

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

func (t *Node[T]) IsBinarySearchTree() bool {
	return slices.IsSorted(t.InOrderSearch())
}

func (t *Node[T]) BSTInsert(value T) {
	curr := t

	for {
		if curr.Value < value {
			if curr.Right == nil {
				curr.Right = &Node[T]{
					Value: value,
					Parent: curr,
				}
				return
			}
			curr = curr.Right
		} else {
			if curr.Left == nil {
				curr.Left = &Node[T]{
					Value: value,
					Parent: curr,
				}
				return
			}
			curr = curr.Left
		}
	}
}
