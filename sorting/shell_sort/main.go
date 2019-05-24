package main

import "fmt"

func main() {
	fmt.Println("insertion sort")
	fmt.Println(insertionSort([]int{3, 1, 0, 9, 5, 2, 4, 9, 7, 6, 1, 3}))
}

func insertionSort(array []int) []int {
	var h = 1
	for h < len(array)/3 {
		h = 3*h + 1
	}
	fmt.Println("h", h)
	for h >= 1 {
		for i := h; i < len(array); i++ {
			fmt.Println(array)
			for j := i; j >= h && array[j] < array[j-h]; j -= h {
				array[j-h], array[j] = array[j], array[j-h]
			}
		}
		h = h / 3
	}
	return array
}
