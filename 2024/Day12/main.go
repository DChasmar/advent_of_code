package main

import (
	"fmt"
	"log"

	"2024/common"
)

func main() {
	filePath := "input.txt"
	s, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	grid := common.LinesToGrid(s)

	areas := getAreas(grid)
	result1 := solve1(areas)

	fmt.Println(result1)

	result2 := solve2(areas)
	fmt.Println(result2)
}

func solve2(areas map[int]map[int]bool) int {
	sum := 0
	for _, area := range areas {
		peri := calcPeri2(area)
		// fmt.Println(len(area), peri)
		sum += len(area) * peri
	}
	return sum
}

func calcPeri2(area map[int]bool) int {
	aboveSet := make(map[int]map[int]bool)
	belowSet := make(map[int]map[int]bool)
	leftSet := make(map[int]map[int]bool)
	rightSet := make(map[int]map[int]bool)
	for key := range area {
		row := key / 1000
		col := key % 1000
		if row-1 < 0 || !area[gridVal(row-1, col)] {
			rowVal := periMinus(row)
			if aboveSet[rowVal] == nil {
				aboveSet[rowVal] = make(map[int]bool)
			}
			aboveSet[rowVal][col] = true
		}
		if !area[gridVal(row+1, col)] {
			rowVal := periPlus(row)
			if belowSet[rowVal] == nil {
				belowSet[rowVal] = make(map[int]bool)
			}
			belowSet[rowVal][col] = true
		}
		if col-1 < 0 || !area[gridVal(row, col-1)] {
			colVal := periMinus(col)
			if leftSet[colVal] == nil {
				leftSet[colVal] = make(map[int]bool)
			}
			leftSet[colVal][row] = true
		}
		if !area[gridVal(row, col+1)] {
			colVal := periPlus(col)
			if rightSet[colVal] == nil {
				rightSet[colVal] = make(map[int]bool)
			}
			rightSet[colVal][row] = true
		}
	}
	sides := periSides2(aboveSet) + periSides2(belowSet) + periSides2(leftSet) + periSides2(rightSet)
	return sides
}

func periSides2(periSet map[int]map[int]bool) int {
	sides := 0
	for _, sideSet := range periSet {
		for len(sideSet) > 0 {
			for key := range sideSet {
				removeAdjacent(sideSet, key)
				sides++
			}
		}
	}
	return sides
}

func removeAdjacent(set map[int]bool, value int) {
	if set[value] {
		delete(set, value)
		if set[value-1] {
			removeAdjacent(set, value-1)
		}
		if set[value+1] {
			removeAdjacent(set, value+1)
		}
	}
}

func getAreas(grid [][]string) map[int]map[int]bool {
	used := make(map[int]bool)
	areas := make(map[int]map[int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if !used[gridVal(row, col)] {
				used[gridVal(row, col)] = true
				area := make(map[int]bool)
				area[gridVal(row, col)] = true
				searchGrid(grid, row, col, grid[row][col], used, area)
				areas[gridVal(row, col)] = area
			}
		}
	}
	return areas
}

func solve1(areas map[int]map[int]bool) int {
	sum := 0
	for _, area := range areas {
		peri := calcPeri1(area)
		sum += len(area) * peri
	}
	return sum
}

func calcPeri1(area map[int]bool) int {
	peri := 0
	for key := range area {
		row := key / 1000
		col := key % 1000

		if row-1 < 0 || !area[gridVal(row-1, col)] {
			peri++
		}

		if !area[gridVal(row+1, col)] {
			peri++
		}

		if col-1 < 0 || !area[gridVal(row, col-1)] {
			peri++
		}

		if !area[gridVal(row, col+1)] {
			peri++
		}
	}
	return peri
}

func searchGrid(grid [][]string, row int, col int, val string, used map[int]bool, area map[int]bool) {
	// Check Up
	if row-1 >= 0 && grid[row-1][col] == val && !used[gridVal(row-1, col)] {
		used[gridVal(row-1, col)] = true
		area[gridVal(row-1, col)] = true
		searchGrid(grid, row-1, col, val, used, area)
	}

	// Check Down
	if row+1 < len(grid) && grid[row+1][col] == val && !used[gridVal(row+1, col)] {
		used[gridVal(row+1, col)] = true
		area[gridVal(row+1, col)] = true
		searchGrid(grid, row+1, col, val, used, area)
	}

	// Check Left
	if col-1 >= 0 && grid[row][col-1] == val && !used[gridVal(row, col-1)] {
		used[gridVal(row, col-1)] = true
		area[gridVal(row, col-1)] = true
		searchGrid(grid, row, col-1, val, used, area)
	}

	// Check Right
	if col+1 < len(grid[0]) && grid[row][col+1] == val && !used[gridVal(row, col+1)] {
		used[gridVal(row, col+1)] = true
		area[gridVal(row, col+1)] = true

		searchGrid(grid, row, col+1, val, used, area)
	}
}

func gridVal(row int, col int) int {
	return row*1000 + col
}

func periMinus(val int) int {
	return (2*(val+2) - 1)
}

func periPlus(val int) int {
	return (2*(val+2) + 1)
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// Part 1: 1413014 is too low
// Part 1: 1445979 is too low
// Part 1: 1446042 is too low
// Part 2: 895486 is too low
