package main

import (
	"fmt"
	"os"
	"strconv"
)

const gridSize = 9

type cell struct {
	row, col int
}

func main() {
	if len(os.Args) != gridSize+1 {
		fmt.Println("Error")
		return
	}

	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
		for j, char := range os.Args[i+1] {
			if char == '.' {
				grid[i][j] = 0
			} else {
				num, err := strconv.Atoi(string(char))
				if err != nil || num < 1 || num > 9 {
					fmt.Println("Error")
					return
				}
				grid[i][j] = num
			}
		}
	}

	if solveSudoku(grid) {
		printSudoku(grid)
	} else {
		fmt.Println("Error")
	}
}

func solveSudoku(grid [][]int) bool {
	empty := findEmptyCell(grid)
	if empty.row == -1 {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValidMove(grid, empty.row, empty.col, num) {
			grid[empty.row][empty.col] = num
			if solveSudoku(grid) {
				return true
			}
			grid[empty.row][empty.col] = 0
		}
	}
	return false
}

func printSudoku(grid [][]int) {
	for _, row := range grid {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

func findEmptyCell(grid [][]int) cell {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 0 {
				return cell{row, col}
			}
		}
	}
	return cell{-1, -1}
}

func isValidMove(grid [][]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

func usedInRow(grid [][]int, row, num int) bool {
	for col := 0; col < gridSize; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func usedInCol(grid [][]int, col, num int) bool {
	for row := 0; row < gridSize; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(grid [][]int, startRow, startCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+startRow][col+startCol] == num {
				return true
			}
		}
	}
	return false
}
