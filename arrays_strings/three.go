package main

import "fmt"
import "github.com/golang-collections/go-datastructures/bitarray"

func main() {
	fmt.Println("is palindrome permutation")
	fmt.Println(tester("tacocat", true))
	fmt.Println(tester("abcab", true))
	fmt.Println(tester("abz", false))
}

func solution(input string) bool {
	bitVector := bitarray.NewBitArray(27)
	for _, val := range input {
		i := uint64(val - 97)
		bit, _ := bitVector.GetBit(i)
		if bit {
			bitVector.ClearBit(i)
		} else {
			bitVector.SetBit(i)
		}
	}
	sol := bitVector.ToNums()
	return len(sol) <= 1
}

func tester(input string, expected bool) string {
	if solution(input) == expected {
		return "pass"
	}
	return "fail"
}
