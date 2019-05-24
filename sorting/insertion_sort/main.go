package main

import "fmt"

func main() {
	fmt.Println("insertion sort")
	fmt.Println(insertionSort([]int{3, 1, 0, 9, 5, 2, 4, 9, 7, 6, 1, 3}))
}

func insertionSort(array []int) []int {
	fmt.Println("input: ", array)
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}
	}
	return array
}
