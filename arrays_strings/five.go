package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tester("pale", "paless", "IMPOSSIBRU"))
	fmt.Println(tester("pale", "ple", "DELETE"))
	fmt.Println(tester("pale", "pales", "INSERT"))
	fmt.Println(tester("pale", "bale", "FLIP"))
	fmt.Println(tester("pale", "bake", "IMPOSSIBRU"))
}

func oneOrZeroEdits(a, b string) string {
	if math.Abs(float64(len(a)-len(b))) > 1 {
		return "IMPOSSIBRU"
	}

	var flag bool
	var i, j int
	sameLength := len(a) == len(b)

	for i < len(a) && j < len(b) {
		if a[i] != b[j] {
			if flag {
				return "IMPOSSIBRU"
			}
			flag = true
			if sameLength {
				j++
			}
		} else {
			j++
		}
		i++
	}

	if sameLength {
		return "FLIP"
	}
	return min(a, b)
}

func min(a, b string) string {
	if len(a) > len(b) {
		return "DELETE"
	}
	return "INSERT"
}

func tester(inputA, inputB, expected string) string {
	actual := oneOrZeroEdits(inputA, inputB)
	if actual == expected {
		return "pass"
	}

	fmt.Println("actual", actual, "expected", expected)
	return "fail"
}
