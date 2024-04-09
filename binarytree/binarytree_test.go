package binarytree

import (
	"testing"
)

func TestEqual(t *testing.T) {
	if !Equal[int](nil, nil) {
		t.Fatalf("expected nil pointers to be equal")
	}

	if !Equal(
		&Node[int]{},
		&Node[int]{},
	) {
		t.Fatalf("expected zero values to be equal")
	}

	if !Equal(
		&Node[int]{
			Value: 0,
			Left: &Node[int]{
				Value: 1,
			},
			Right: &Node[int]{
				Value: 2,
			},
		},
		&Node[int]{
			Value: 0,
			Left: &Node[int]{
				Value: 1,
			},
			Right: &Node[int]{
				Value: 2,
			},
		},
	) {
		t.Fatalf("expected trees to be equal")
	}

	tree0 := &Node[int]{
		Value: 0,
	}
	tree0.Left = &Node[int]{
		Value: 1,
		Parent: tree0,
	}

	tree1 := &Node[int]{
		Value: 0,
	}
	tree1.Left = &Node[int]{
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

	tree0.Left.Parent = &Node[int]{
		Value: 2,
	}

	if Equal(tree0, tree1) {
		t.Fatalf("expected trees with one nil parent and one malformed parent to be not equal")
	}

	tree1.Left.Parent = &Node[int]{
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
	if !Equal(New[int](), &Node[int]{}) {
		t.Fatalf("expected %v, got %v", &Node[int]{}, New[int]())
	}
	if !Equal(New[string](), &Node[string]{}) {
		t.Fatalf("expected %v, got %v", &Node[string]{}, New[string]())
	}
	if !Equal(New[float32](), &Node[float32]{}) {
		t.Fatalf("expected %v, got %v", &Node[float32]{}, New[float32]())
	}
}

func TestFromValue(t *testing.T) {
	if !Equal(FromValue(1), &Node[int]{Value: 1}) {
		t.Fatalf("expected %v, got %v", &Node[int]{Value:1}, FromValue(1))
	}
	if !Equal(FromValue("thing"), &Node[string]{Value: "thing"}) {
		t.Fatalf("expected %v, got %v", &Node[string]{Value:"thing"}, FromValue("thing"))
	}
	if !Equal(FromValue[float32](3.14), &Node[float32]{Value: 3.14}) {
		t.Fatalf("expected %v, got %v", &Node[float32]{Value:3.14}, FromValue(3.14))
	}
}

func TestFrom(t *testing.T) {
	tree0 := &Node[int]{
		Value: 1,
		Parent: nil,
		Left: nil,
		Right: nil,
	}

	tree0.Left = &Node[int]{
		Value: 2,
		Parent: tree0,
	}
	tree0.Right = &Node[int]{
		Value: 3,
		Parent: tree0,
	}

	tree1 := From(1,
		From(2, nil, nil),
		From(3, nil, nil),
	)

	if !Equal[int](tree0, tree1) {
		t.Fatalf("expected trees to be equal")
	}
}
