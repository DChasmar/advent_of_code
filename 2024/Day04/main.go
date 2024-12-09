package main

import (
	"fmt"
	"log"

	"2024/common"
)

func main() {
	filePath := "input.txt"
	grid, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	result1 := searchGrid1(grid)
	fmt.Println(result1)

	result2 := searchGrid2(grid)
	fmt.Println(result2)
}

func searchGrid1(grid [][]string) int {
	sum := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == "X" {
				westable := j > 2
				eastable := j < len(row)-3
				northable := i > 2
				southable := i < len(grid)-3
				if westable {
					sum += west1(grid, i, j)
					if northable {
						sum += northwest1(grid, i, j)
					}
					if southable {
						sum += southwest1(grid, i, j)
					}
				}
				if eastable {
					sum += east1(grid, i, j)
					if northable {
						sum += northeast1(grid, i, j)
					}
					if southable {
						sum += southeast1(grid, i, j)
					}
				}
				if northable {
					sum += north1(grid, i, j)
				}
				if southable {
					sum += south1(grid, i, j)
				}
			}
		}
	}
	return sum
}

func west1(grid [][]string, i, j int) int {
	if grid[i][j-1] == "M" && grid[i][j-2] == "A" && grid[i][j-3] == "S" {
		return 1
	}
	return 0
}

func east1(grid [][]string, i, j int) int {
	if grid[i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S" {
		return 1
	}
	return 0
}

func north1(grid [][]string, i, j int) int {
	if grid[i-1][j] == "M" && grid[i-2][j] == "A" && grid[i-3][j] == "S" {
		return 1
	}
	return 0
}

func south1(grid [][]string, i, j int) int {
	if grid[i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S" {
		return 1
	}
	return 0
}

func southwest1(grid [][]string, i, j int) int {
	if grid[i+1][j-1] == "M" && grid[i+2][j-2] == "A" && grid[i+3][j-3] == "S" {
		return 1
	}
	return 0
}

func southeast1(grid [][]string, i, j int) int {
	if grid[i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S" {
		return 1
	}
	return 0
}

func northwest1(grid [][]string, i, j int) int {
	if grid[i-1][j-1] == "M" && grid[i-2][j-2] == "A" && grid[i-3][j-3] == "S" {
		return 1
	}
	return 0
}

func northeast1(grid [][]string, i, j int) int {
	if grid[i-1][j+1] == "M" && grid[i-2][j+2] == "A" && grid[i-3][j+3] == "S" {
		return 1
	}
	return 0
}

func searchGrid2(grid [][]string) int {
	sum := 0
	for i := 1; i < len(grid)-1; i++ {
		row := grid[i]
		for j := 1; j < len(row)-1; j++ {
			cell := row[j]
			if cell == "A" {
				s := grid[i+1][j+1] + grid[i+1][j-1] + grid[i-1][j-1] + grid[i-1][j+1]
				if s == "MMSS" || s == "SSMM" || s == "MSSM" || s == "SMMS" {
					sum++
				}
			}
		}
	}
	return sum
}

func readAndProcessFile(filePath string) ([][]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	grid := common.LinesToGrid(lines)

	return grid, nil
}
