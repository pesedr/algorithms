package main

import "fmt"

func main() {
	fmt.Println("rotate matrix")
	rotateInPlace(
		[][]int{[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
			[]int{13, 14, 15, 16}})
		rotateInPlace(
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9}})
}

func rotateInPlace(matrix [][]int) {
	n := len(matrix)
	fmt.Println("old")
	printMatrix(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - layer - 1
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
	fmt.Println("new")
	printMatrix(matrix)
}

func rotateWithNewMatrix(matrix [][]int) {
	res := make([][]int, len(matrix))
	for i := range res {
		res[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {

			newi := len(matrix) - i - 1
			res[j][newi] = matrix[i][j]
		}
	}
	fmt.Println("old")
	printMatrix(matrix)

	fmt.Println("new")
	printMatrix(res)
}

func printMatrix(matrix [][]int) {
	for i := range matrix {
		for j := range matrix {
			fmt.Print(matrix[i][j])
			fmt.Printf("\t")
		}
		fmt.Println()
	}
}
