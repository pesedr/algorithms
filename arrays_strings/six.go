package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(tester("aabcccccaaa", "a2b1c5a3"))
	fmt.Println(tester("aabbcdd", "aabbcdd"))
}

func compress(input string) string {
	var count = 1
	var previous int32
	var res = make([]rune, 0, len(input)+1)

	for key, val := range input {
		if key == 0 {
			count = 1
			previous = val
			continue
		} else if val != previous {
			num := []rune(strconv.Itoa(count))
			res = append(res, previous, num[0])
			if len(res) >= len(input) {
				return input
			}
			count = 1
			previous = val
		} else {
			count++
		}
		if key == len(input)-1 {
			num := []rune(strconv.Itoa(count))
			res = append(res, previous, num[0])

		}
	}
	if len(res) >= len(input) {
		return input
	}
	return string(res)
}

func tester(input, expected string) string {
	actual := compress(input)
	if actual != expected {
		fmt.Println("actual", actual, "expected", expected)
		return "fail"
	}
	return "pass"
}
