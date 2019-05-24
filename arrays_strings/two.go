package main

import "fmt"

func main() {
	fmt.Println(test("abc", "cba", true))
	fmt.Println(test("abc", "abd", false))
}

func permutation(a, b string) bool {
	mapA := make(map[int32]int)

	for _, val := range a {
		mapA[val]++
	}
	for _, val := range b {
		mapA[val]--
		if mapA[val] < 0 {
			return false
		}
	}

	return true
}

func test(a, b string, expected bool) string {
	ans := permutation(a, b)
	if ans == expected {
		return "pass"
	}
	return "fail"
}
