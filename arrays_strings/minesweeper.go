package main

import "fmt"

func main() {
	fmt.Println("minesweeper")
	minesweep([]string{"X000", "X000", "X000", "0000"})
}

func minesweep(board []string) {
	res := make([][]int, len(board)+2)
	for i := range res {
		res[i] = make([]int, len(board)+2)
	}
	for i, row := range board {
		for j, square := range row {
			if square == 'X' {
				res[i][j]++
				res[i][j+1]++
				res[i][j+2]++
				res[i+1][j]++
				res[i+1][j+1] += 9
				res[i+1][j+2]++
				res[i+2][j]++
				res[i+2][j+1]++
				res[i+2][j+2]++
			}
		}
	}
	printBoard(res, len(board))

}

func printBoard(res [][]int, length int) {

	for i := 1; i <= length; i++ {
		for j := 1; j <= length; j++ {
			if res[i][j] >= 9 {
				fmt.Print("X ")
			} else {
				fmt.Print(res[i][j], " ")
			}
		}
		fmt.Printf("\n")
	}

}
