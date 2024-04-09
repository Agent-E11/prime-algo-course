package binarytree

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal[int](nil, nil) {
		t.Fatalf("expected nil pointers to be equal")
	}

	if !Equal(
		&BinaryNode[int]{},
		&BinaryNode[int]{},
	) {
		t.Fatalf("expected zero values to be equal")
	}

	if !Equal(
		&BinaryNode[int]{
			Value: 0,
			Left: &BinaryNode[int]{
				Value: 1,
			},
			Right: &BinaryNode[int]{
				Value: 2,
			},
		},
		&BinaryNode[int]{
			Value: 0,
			Left: &BinaryNode[int]{
				Value: 1,
			},
			Right: &BinaryNode[int]{
				Value: 2,
			},
		},
	) {
		t.Fatalf("expected trees to be equal")
	}

	tree0 := &BinaryNode[int]{
		Value: 0,
	}
	tree0.Left = &BinaryNode[int]{
		Value: 1,
		Parent: tree0,
	}

	tree1 := &BinaryNode[int]{
		Value: 0,
	}
	tree1.Left = &BinaryNode[int]{
		Value: 1,
		Parent: tree1,
	}

	if !Equal(tree0, tree1) {
		t.Fatalf("expected trees with parents to be equal")
	}

	tree0.Left.Parent = nil

	if Equal(tree0, tree1) {
		t.Fatalf("expected trees with one nil parent to be not equal")
	}

	tree0.Left.Parent = &BinaryNode[int]{
		Value: 2,
	}

	if Equal(tree0, tree1) {
		t.Fatalf("expected trees with one nil parent and one malformed parent to be not equal")
	}

	tree1.Left.Parent = &BinaryNode[int]{
		Value: 3,
	}

	// Equal() should not distinguish between different malformed parents.
	// It should only check whether both are malformed, or both are nil,
	// or both are properly formed.
	if !Equal(tree0, tree1) {
		t.Fatalf("expected trees with both malformed parents to be equal")
	}
}

func TestNew(t *testing.T) {
	if !Equal(New[int](), &BinaryNode[int]{}) {
		t.Fatalf("expected %v, got %v", &BinaryNode[int]{}, New[int]())
	}
	if !Equal(New[string](), &BinaryNode[string]{}) {
		t.Fatalf("expected %v, got %v", &BinaryNode[string]{}, New[string]())
	}
	if !Equal(New[float32](), &BinaryNode[float32]{}) {
		t.Fatalf("expected %v, got %v", &BinaryNode[float32]{}, New[float32]())
	}
}

func TestFromValue(t *testing.T) {
	
}

func TestFrom(t *testing.T) {
	root := &BinaryNode[int]{
		Value: 1,
		Parent: nil,
		Left: nil,
		Right: nil,
	}

	root.Left = &BinaryNode[int]{
		Value: 2,
		Parent: root,
	}
	root.Right = &BinaryNode[int]{
		Value: 3,
		Parent: root,
	}
}
