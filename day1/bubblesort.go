package day1

import "fmt"

// Sort array in-place using the bubble sort method
func BubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
			fmt.Println("Values:", *arr)
		}
	}
}
