package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("rotating strings")
	fmt.Println(tester("waterbottle", "erbottlewat", true))
}

func rotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return strings.Contains(a+a, b)
}

func tester(a, b string, expected bool) string {
	actual := rotation(a, b)
	if actual == expected {
		return "pass"
	}
	fmt.Println("actual", actual, "expected", expected)
	return "fail"
}
