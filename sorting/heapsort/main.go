package main

import "fmt"

func main() {
	fmt.Println("heapsort")
	fmt.Println(heapsort([]int{3, 1, 0, 8, 5, 2, 4, 9, 7, 6, 1, 3}))
}

func heapsort(array []int) []int {
	fmt.Println("intput: ", array)
	buildHeap(array)
	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		maxHeapify(array[:i-1], 0)
	}
	return array
}

func buildHeap(array []int) {
	start := len(array)/2 - 1
	for i := start; i >= 0; i-- {
		maxHeapify(array, i)
	}
}

func maxHeapify(array []int, i int) {
	largest := i
	left := left(i)
	right := right(i)

	if left < len(array) && array[left] > array[i] {
		largest = left
	}
	if right < len(array) && array[right] > array[largest] {
		largest = right
	}
	if largest != i {
		array[largest], array[i] = array[i], array[largest]
		maxHeapify(array, largest)
	}
}

func parent(i int) int {
	i++
	return i/2 - 1
}

func left(i int) int {
	i++
	return 2*i - 1
}

func right(i int) int {
	i++
	return 2 * i
}
