package day1

import "math"

func BinarySearch(haystack []int, needle int) bool {
	// lo is inclusive, hi is exclusive
	lo := 0.0
	hi := float64(len(haystack))

	for {
		// Get the index and value in the middle of lo and hi
		mid := math.Floor(lo + (hi - lo)/2)
		val := haystack[int(mid)]

		if val == needle {
			// If it is found, just return true
			return true
		} else if val > needle {
			// If the middle value is greater than the target value
			// then we can ignore all values above mid (excluding mid, because
			// hi is exclusive)
			hi = mid
		} else {
			// If the middle value is less than the target value (we know it is
			// not equal), then we can ignore all values below mid (add one to
			// mid to exclude it, because lo is inclusive)
			lo = mid + 1
		}

		// If lo is no longer lower than hi, exit the loop
		if !(lo < hi) { break }
	}

	// Value was not found
	return false
}
