package main

import (
	"fmt"
	"log"

	"2024/common"
)

func main() {
	filePath := "example.txt"
	s, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	grid := common.LinesToIntGrid(s)

	result1 := solve1(grid)
	fmt.Println(result1)

	result2 := solve2(grid)
	fmt.Println(result2)

}

func solve2(grid [][]int) int {
	sum := 0
	for i, row := range grid {
		for j, val := range row {
			if val == 0 {
				sum += findNine2(grid, val, i, j)
			}
		}
	}
	return sum
}

func findNine2(grid [][]int, val int, row int, col int) int {
	if val == 9 {
		return 1
	}

	count := 0

	// Check Up
	if row > 0 && grid[row-1][col] == val+1 {
		count += findNine2(grid, val+1, row-1, col)
	}

	// Check Down
	if row < len(grid)-1 && grid[row+1][col] == val+1 {
		count += findNine2(grid, val+1, row+1, col)
	}

	// Check Left
	if col > 0 && grid[row][col-1] == val+1 {
		count += findNine2(grid, val+1, row, col-1)
	}

	// Check Right
	if col < len(grid[0])-1 && grid[row][col+1] == val+1 {
		count += findNine2(grid, val+1, row, col+1)
	}

	return count
}

func solve1(grid [][]int) int {
	sum := 0
	for i, row := range grid {
		for j, val := range row {
			if val == 0 {
				set := make(map[int]bool)
				findNine1(grid, val, i, j, set)
				delete(set, -1) // Remove the initial -1 entry if present
				sum += len(set)
			}
		}
	}
	return sum
}

func findNine1(grid [][]int, val int, row int, col int, set map[int]bool) {
	if val == 9 {
		set[row*1000+col] = true
		return
	}

	// Check Up
	if row > 0 && grid[row-1][col] == val+1 {
		findNine1(grid, val+1, row-1, col, set)
	}

	// Check Down
	if row < len(grid)-1 && grid[row+1][col] == val+1 {
		findNine1(grid, val+1, row+1, col, set)
	}

	// Check Left
	if col > 0 && grid[row][col-1] == val+1 {
		findNine1(grid, val+1, row, col-1, set)
	}

	// Check Right
	if col < len(grid[0])-1 && grid[row][col+1] == val+1 {
		findNine1(grid, val+1, row, col+1, set)
	}
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
