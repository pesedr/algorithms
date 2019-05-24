package main

import "fmt"

func main() {
	fmt.Println("countingSort")
	fmt.Println(countingSort([]int{3, 1, 0, 8, 5, 2, 4, 9, 7, 6, 1, 3}, 9))
}

func countingSort(array []int, max int) []int {
	fmt.Println("starting array", array)
	res := make([]int, len(array))
	counter := make([]int, max+1)
	for _, val := range array {
		counter[val]++
	}
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}
	for _, val := range array {
		index := counter[val] - 1
		res[index] = val
		counter[val]--
	}
	return res
}
