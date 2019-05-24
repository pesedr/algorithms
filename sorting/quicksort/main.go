package main

import "fmt"

func main() {
	fmt.Println("quicksort")
	fmt.Println(quicksort([]int{1, 0, 8, 5, 2, 4, 9, 7, 6, 3}))

}

func quicksort(array []int) []int {
	fmt.Println("input", array)
	sorter(array, 0, len(array)-1)
	return array
}

func sorter(array []int, start, finish int) {
	if start < finish {
		pivot := partition(array, start, finish)
		sorter(array, start, pivot-1)
		sorter(array, pivot+1, finish)
	}
}

func partition(array []int, start, finish int) int {
	pivot := array[finish]
	i := start - 1
	for j := start; j < finish; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[finish] = array[finish], array[i+1]
	return i + 1
}
