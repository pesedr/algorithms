package main

import "fmt"

func main() {
	fmt.Println("binarySearch")
	fmt.Println(binarySearch(23, []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98}))
	fmt.Println(binarySearch(50, []int{10, 11, 12, 16, 18, 23, 29, 33, 48, 54, 57, 68, 77, 84, 98}))
}

func binarySearch(key int, sorted []int) int {
	var low = 0
	var high = len(sorted) - 1
	for low <= high {
		var mid = low + (high-low)/2
		if key < sorted[mid] {
			high = mid - 1
		} else if key > sorted[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
