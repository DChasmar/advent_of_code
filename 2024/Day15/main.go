package main

import (
	"fmt"
	"log"
	"strings"

	"2024/common"
)

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	grid, directions := buildData(data)

	result1 := solve1(grid, directions)
	fmt.Println(result1)

	ogGrid, directions := buildData(data)
	grid2 := transformGrid(ogGrid)
	// common.GridToTxt(grid2, "transformedExample.txt")
	result2 := solve2(grid2, directions)
	// common.GridToTxt(grid2, "transformedFinal.txt")
	fmt.Println(result2)
}

func transformGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]string, len(row)*2)
		for j, cell := range row {
			switch cell {
			case "#":
				newGrid[i][j*2] = "#"
				newGrid[i][j*2+1] = "#"
			case ".":
				newGrid[i][j*2] = "."
				newGrid[i][j*2+1] = "."
			case "@":
				newGrid[i][j*2] = "@"
				newGrid[i][j*2+1] = "."
			case "O":
				newGrid[i][j*2] = "["
				newGrid[i][j*2+1] = "]"
			}
		}
	}
	return newGrid
}

func solve2(grid [][]string, directions []string) int {
	row, col := findRobot(grid)
	for _, dir := range directions {
		switch dir {
		case "^":
			if grid[row-1][col] == "." {
				grid[row-1][col] = "@"
				grid[row][col] = "."
				row -= 1
			} else {
				if grid[row-1][col] == "]" {
					newCols := [][]int{{col}}
					newCols = append(newCols, []int{col - 1})
					ranges, ok := up2(grid, row, newCols)
					if ok {
						moveUp2(grid, row, ranges)
						row -= 1
					}
				} else if grid[row-1][col] == "[" {
					newCols := [][]int{{col}}
					newCols = append(newCols, []int{col})
					ranges, ok := up2(grid, row, newCols)
					if ok {
						moveUp2(grid, row, ranges)
						row -= 1
					}
				}
			}
		case "v":
			if grid[row+1][col] == "." {
				grid[row+1][col] = "@"
				grid[row][col] = "."
				row += 1
			} else {
				if grid[row+1][col] == "]" {
					newCols := [][]int{{col}}
					newCols = append(newCols, []int{col - 1})
					ranges, ok := down2(grid, row, newCols)
					if ok {
						moveDown2(grid, row, ranges)
						row += 1
					}
				} else if grid[row+1][col] == "[" {
					newCols := [][]int{{col}}
					newCols = append(newCols, []int{col})
					ranges, ok := down2(grid, row, newCols)
					if ok {
						moveDown2(grid, row, ranges)
						row += 1
					}
				}
			}
		case "<":
			emptyCol, ok := left2(grid, row, col-1)
			if ok {
				moveLeft2(grid, row, emptyCol)
				col -= 1
			}
		case ">":
			emptyCol, ok := right2(grid, row, col+1)
			if ok {
				moveRight2(grid, row, emptyCol)
				col += 1
			}
		}
	}

	return countSum(grid, "[")
}

func moveDown2(grid [][]string, row int, cols [][]int) {
	for i := len(cols) - 1; i >= 0; i-- {
		for _, col := range cols[i] {
			if grid[row+i][col] == "@" {
				grid[row+1][col] = "@"
				grid[row][col] = "."
			} else {
				newRow := row + i
				grid[newRow+1][col] = grid[newRow][col]
				grid[newRow+1][col+1] = grid[newRow][col+1]
				grid[newRow][col] = "."
				grid[newRow][col+1] = "."
			}
		}
	}
}

func down2(grid [][]string, row int, cols [][]int) ([][]int, bool) {
	length := len(cols)
	lastIndex := len(cols) - 1
	newArr := []int{}
	allDots := true
	for _, col := range cols[lastIndex] {
		if grid[row+length][col] == "#" || (grid[row+length][col] != "@" && grid[row+length][col+1] == "#") {
			return cols, false
		} else if grid[row+length][col] == "]" {
			if !common.IntArrContains(newArr, col-1) {
				newArr = append(newArr, col-1)
				allDots = false
			}
		} else if grid[row+length][col] == "[" {
			if !common.IntArrContains(newArr, col) {
				newArr = append(newArr, col)
				allDots = false
			}
		}
		if grid[row+length][col+1] == "[" && grid[row+length][col] != "@" {
			if !common.IntArrContains(newArr, col+1) {
				newArr = append(newArr, col+1)
				allDots = false
			}
		}
	}
	if allDots {
		return cols, true
	} else {
		cols = append(cols, newArr)
		return down2(grid, row, cols)
	}
}

func moveUp2(grid [][]string, row int, cols [][]int) {
	for i := len(cols) - 1; i >= 0; i-- {
		for _, col := range cols[i] {
			if grid[row-i][col] == "@" {
				grid[row-1][col] = "@"
				grid[row][col] = "."
			} else {
				newRow := row - i
				grid[newRow-1][col] = grid[newRow][col]
				grid[newRow-1][col+1] = grid[newRow][col+1]
				grid[newRow][col] = "."
				grid[newRow][col+1] = "."
			}
		}
	}
}

func up2(grid [][]string, row int, cols [][]int) ([][]int, bool) {
	length := len(cols)
	lastIndex := len(cols) - 1
	newArr := []int{}
	allDots := true
	for _, col := range cols[lastIndex] {
		if grid[row-length][col] == "#" || (grid[row-length][col] != "@" && grid[row-length][col+1] == "#") {
			return cols, false
		} else if grid[row-length][col] == "]" {
			if !common.IntArrContains(newArr, col-1) {
				newArr = append(newArr, col-1)
				allDots = false
			}
		} else if grid[row-length][col] == "[" {
			if !common.IntArrContains(newArr, col) {
				newArr = append(newArr, col)
				allDots = false
			}
		}
		if grid[row-length][col+1] == "[" && grid[row-length][col] != "@" {
			if !common.IntArrContains(newArr, col+1) {
				newArr = append(newArr, col+1)
				allDots = false
			}
		}
	}
	if allDots {
		return cols, true
	} else {
		cols = append(cols, newArr)
		return up2(grid, row, cols)
	}
}

func moveLeft2(grid [][]string, row, col int) {
	if grid[row][col+1] == "@" {
		grid[row][col] = "@"
		grid[row][col+1] = "."
		return
	} else {
		grid[row][col] = grid[row][col+1]
		moveLeft2(grid, row, col+1)
	}
}

func moveRight2(grid [][]string, row, col int) {
	if grid[row][col-1] == "@" {
		grid[row][col] = "@"
		grid[row][col-1] = "."
		return
	} else {
		grid[row][col] = grid[row][col-1]
		moveRight2(grid, row, col-1)
	}
}

func left2(grid [][]string, row, col int) (int, bool) {
	if grid[row][col] == "#" {
		return -1, false
	}
	if grid[row][col] == "]" {
		return left2(grid, row, col-2)
	} else if grid[row][col] == "." {
		return col, true
	} else {
		common.GridToTxt(grid, "failedOutputLeft.txt")
		log.Fatalf("unexpected value: %v, at row: %d, col %d", grid[row][col], row, col)
		return -1, false
	}
}

func right2(grid [][]string, row, col int) (int, bool) {
	if grid[row][col] == "#" {
		return -1, false
	}
	if grid[row][col] == "[" {
		return right2(grid, row, col+2)
	} else if grid[row][col] == "." {
		return col, true
	} else {
		common.GridToTxt(grid, "failedOutputRight.txt")
		log.Fatalf("unexpected value: %v, at row: %d, col %d", grid[row][col], row, col)
		return -1, false
	}
}

func solve1(grid [][]string, directions []string) int {
	row, col := findRobot(grid)
	for _, dir := range directions {
		switch dir {
		case "^":
			emptyRow, ok := up(grid, row-1, col)
			if ok {
				moveUp(grid, emptyRow, col)
				row -= 1
			}
		case "v":
			emptyRow, ok := down(grid, row+1, col)
			if ok {
				moveDown(grid, emptyRow, col)
				row += 1
			}
		case "<":
			emptyCol, ok := left(grid, row, col-1)
			if ok {
				moveLeft(grid, row, emptyCol)
				col -= 1
			}
		case ">":
			emptyCol, ok := right(grid, row, col+1)
			if ok {
				moveRight(grid, row, emptyCol)
				col += 1
			}
		}
	}

	return countSum(grid, "O")
}

func countSum(grid [][]string, val string) int {
	count := 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == val {
				count += 100*r + c
			}
		}
	}
	return count
}

func moveUp(grid [][]string, row, col int) {
	if grid[row+1][col] == "@" {
		grid[row][col] = "@"
		grid[row+1][col] = "."
		return
	} else {
		grid[row][col] = "O"
		moveUp(grid, row+1, col)
	}
}

func moveDown(grid [][]string, row, col int) {
	if grid[row-1][col] == "@" {
		grid[row][col] = "@"
		grid[row-1][col] = "."
		return
	} else {
		grid[row][col] = "O"
		moveDown(grid, row-1, col)
	}
}

func moveLeft(grid [][]string, row, col int) {
	if grid[row][col+1] == "@" {
		grid[row][col] = "@"
		grid[row][col+1] = "."
		return
	} else {
		grid[row][col] = "O"
		moveLeft(grid, row, col+1)
	}
}

func moveRight(grid [][]string, row, col int) {
	if grid[row][col-1] == "@" {
		grid[row][col] = "@"
		grid[row][col-1] = "."
		return
	} else {
		grid[row][col] = "O"
		moveRight(grid, row, col-1)
	}
}

func findRobot(grid [][]string) (int, int) {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == "@" {
				return row, col
			}
		}
	}
	log.Fatal("robot not found")
	return -1, -1
}

func up(grid [][]string, row, col int) (int, bool) {
	if grid[row][col] == "#" {
		return -1, false
	}
	if grid[row][col] == "O" {
		return up(grid, row-1, col)
	} else if grid[row][col] == "." {
		return row, true
	} else {
		log.Fatalf("unexpected value: %v, at row: %d, col %d", grid[row][col], row, col)
		return -1, false
	}
}

func down(grid [][]string, row, col int) (int, bool) {
	if grid[row][col] == "#" {
		return -1, false
	}
	if grid[row][col] == "O" {
		return down(grid, row+1, col)
	} else if grid[row][col] == "." {
		return row, true
	} else {
		log.Fatalf("unexpected value: %v, at row: %d, col %d", grid[row][col], row, col)
		return -1, false
	}
}

func left(grid [][]string, row, col int) (int, bool) {
	if grid[row][col] == "#" {
		return -1, false
	}
	if grid[row][col] == "O" {
		return left(grid, row, col-1)
	} else if grid[row][col] == "." {
		return col, true
	} else {
		log.Fatalf("unexpected value: %v, at row: %d, col %d", grid[row][col], row, col)
		return -1, false
	}
}

func right(grid [][]string, row, col int) (int, bool) {
	if grid[row][col] == "#" {
		return -1, false
	}
	if grid[row][col] == "O" {
		return right(grid, row, col+1)
	} else if grid[row][col] == "." {
		return col, true
	} else {
		log.Fatalf("unexpected value: %v, at row: %d, col %d", grid[row][col], row, col)
		return -1, false
	}
}

func buildData(lines []string) ([][]string, []string) {
	var grid [][]string
	var directions []string
	gridPart := true

	for _, line := range lines {
		if line == "" {
			gridPart = false
			continue
		}

		if gridPart {
			grid = append(grid, strings.Split(line, ""))
		} else {
			directions = append(directions, strings.Split(line, "")...)
		}
	}

	return grid, directions
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
