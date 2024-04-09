package binarytree

import "cmp"

// Check if two Node's are equal. This function will check if the two
// trees coresponding to the nodes have the same values, and the same
// structure.
//
// The only exeption to this is that the nodes in the trees can differ in
// regards to parents. If two node's parents are "malformed", in that a.Left =
// b, but b.Parent = c, as long as the two nodes that are being compared both
// have malformed parents, then they can be equal. However, the root nodes are
// will not be checked for malformed or nil parents
//
// Example:
//
// a.Left = b,
// b.Parent = c
// 
// x.Left = y,
// y.Parent = z
//
// b can still be equal to y, because they are both malformed.
func Equal[T cmp.Ordered](a *Node[T], b *Node[T]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	if a.Left != nil && b.Left != nil {
		aLeftStatus := -1
		bLeftStatus := -1
		if a.Left.Parent == a {
			aLeftStatus = 0
		} else if a.Left.Parent == nil {
			aLeftStatus = 1
		} else {
			aLeftStatus = 2
		}
		if b.Left.Parent == b {
			bLeftStatus = 0
		} else if b.Left.Parent == nil {
			bLeftStatus = 1
		} else {
			bLeftStatus = 2
		}

		if aLeftStatus != bLeftStatus {
			return false
		}
	}
	if a.Right != nil && b.Right != nil {
		aRightStatus := -1
		bRightStatus := -1
		if a.Right.Parent == a {
			aRightStatus = 0
		} else if a.Right.Parent == nil {
			aRightStatus = 1
		} else {
			aRightStatus = 2
		}
		if b.Right.Parent == b {
			bRightStatus = 0
		} else if b.Right.Parent == nil {
			bRightStatus = 1
		} else {
			bRightStatus = 2
		}

		if aRightStatus != bRightStatus {
			return false
		}
	}

	return Equal(a.Left, b.Left) && Equal(a.Right, b.Right)
}
