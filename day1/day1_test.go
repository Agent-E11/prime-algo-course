package day1_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/prime-algo-course/day1"
)

func TestLinearSearch(t *testing.T) {
	haystack := []int{0, 1, 1, 2, 3, 5, 8, 13}

	if !day1.LinearSearch(haystack, 1) {
		t.Fatalf("expected to find 1 in list")
	}

	if day1.LinearSearch(haystack, 20) {
		t.Fatalf("expected to not find 20 in list")
	}
}

func TestBinarySearch(t *testing.T) {
	haystack := []int{0, 1, 1, 2, 3, 5, 8, 13}

	if !day1.BinarySearch(haystack, 1) {
		t.Fatalf("expected to find 1 in list")
	}

	if day1.LinearSearch(haystack, 20) {
		t.Fatalf("expected to not find 20 in list")
	}
}

func TestTwoCrystalBalls(t *testing.T) {
	i0 := day1.TwoCrystalBalls(
		[]bool{false, false, false, false, false, false, false, false, true},
	)

	if i0 != 8 {
		t.Fatalf("expected i0 to be 8")
	}

	i1 := day1.TwoCrystalBalls(
		[]bool{false, false, false, false, false, true, true, true, true},
	)

	if i1 != 5 {
		t.Fatalf("expected i0 to be 5")
	}

	i2 := day1.TwoCrystalBalls(
		[]bool{false, false, false, false, false, false, false, false, false},
	)

	if i2 != -1 {
		t.Fatalf("expected i0 to be -1 (not found)")
	}
}

func TestBubbleSort(t *testing.T) {
	arr0 := []int{2, 10, 1, 6, 7, 5, 8, 3, 4, 9}
	expect0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	day1.BubbleSort(&arr0)

	if !slices.Equal(arr0, expect0) {
		t.Fatalf("expected sorted array `%v`, got `%v`", expect0, arr0)
	}

	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	expect1 := []int{1, 2, 3, 4, 5, 6, 7}
	day1.BubbleSort(&arr1)

	if !slices.Equal(arr0, expect0) {
		t.Fatalf("expected sorted array `%v`, got `%v`", expect1, arr1)
	}
}
