package main

import "fmt"

func main() {
	fmt.Println("selection sort")
	fmt.Println(selectionSort([]int{6, 4, 10, 9, 7, 7, 8, 10, 8, 9, 10}))
}

func selectionSort(array []int) []int {
	fmt.Println("input: ", array)
	for key := range array {
		fmt.Println(array)
		for j := key + 1; j < len(array); j++ {
			if array[j] < array[key] {
				array[key], array[j] = array[j], array[key]
			}
		}
	}
	return array
}
