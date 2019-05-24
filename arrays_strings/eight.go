package main

import "fmt"

func main() {
	fmt.Println("nullify columns and rows")
	nullify([][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}})
}

func nullify(matrix [][]int) {
	fmt.Println("old")
	printMatrix(matrix)
	for i, row := range matrix {
		for j, square := range row {
			if square == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < len(matrix); j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}

	fmt.Println("new")
	printMatrix(matrix)
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
