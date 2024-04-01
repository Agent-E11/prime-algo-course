package day1

func qs(arr *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx - 1)
	qs(arr, pivotIdx + 1, hi)
}

func partition(arr *[]int, lo int, hi int) int {
	
	pivot := (*arr)[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if (*arr)[i] <= pivot {
			idx++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[idx]
			(*arr)[idx] = tmp
		}
	}

	// Move the index onto the pivot location
	idx++

	(*arr)[hi] = (*arr)[idx]
	(*arr)[idx] = pivot

	return idx
}

// Sort an array in-place using the quicksort method
func Quicksort(arr *[]int) {
	qs(arr, 0, len(*arr)-1)

}
