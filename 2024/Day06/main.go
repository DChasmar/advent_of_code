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

	result1 := solve1(grid)
	fmt.Println(len(result1))

	result2 := solve2(grid, result1)
	fmt.Println(len(result2))
}

func findGuard(grid [][]string) []int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" {
				return []int{i, j}
			}
		}
	}
	log.Fatalf("guard not found")
	return nil
}

func solve1(grid [][]string) map[int]bool {
	dir := "N"
	index := findGuard(grid)

	set := make(map[int]bool)
	set[index[0]*1000+index[1]] = true

	for index[0] > 0 && index[0] < len(grid)-1 && index[1] > 0 && index[1] < len(grid[index[0]])-1 {
		if dir == "N" {
			if grid[index[0]-1][index[1]] == "#" {
				dir = "E"
			} else {
				index[0]--
				set[index[0]*1000+index[1]] = true
			}
		} else if dir == "E" {
			if grid[index[0]][index[1]+1] == "#" {
				dir = "S"
			} else {
				index[1]++
				set[index[0]*1000+index[1]] = true
			}
		} else if dir == "S" {
			if grid[index[0]+1][index[1]] == "#" {
				dir = "W"
			} else {
				index[0]++
				set[index[0]*1000+index[1]] = true
			}
		} else if dir == "W" {
			if grid[index[0]][index[1]-1] == "#" {
				dir = "N"
			} else {
				index[1]--
				set[index[0]*1000+index[1]] = true
			}
		}
	}
	return set
}

func solve2(grid [][]string, set1 map[int]bool) map[int]bool {
	ogIndex := findGuard(grid)
	ogVal := ogIndex[0]*1000 + ogIndex[1]
	set1[ogVal] = false

	masterSet := make(map[int]bool)

	for k := range set1 {
		blockIndex := []int{k / 1000, k % 1000}
		char := grid[blockIndex[0]][blockIndex[1]]
		if char == "^" || char == "#" {
			continue
		}
		newGrid := make([][]string, len(grid))
		for i := range grid {
			newGrid[i] = make([]string, len(grid[i]))
			copy(newGrid[i], grid[i])
		}
		newGrid[blockIndex[0]][blockIndex[1]] = "#"

		index := make([]int, len(ogIndex))
		copy(index, ogIndex)
		dir := "N"

		setN := make(map[int]bool)
		setE := make(map[int]bool)
		setW := make(map[int]bool)
		setS := make(map[int]bool)

		for index[0] > 0 && index[0] < len(grid)-1 && index[1] > 0 && index[1] < len(grid[index[0]])-1 {
			if dir == "N" {
				if newGrid[index[0]-1][index[1]] == "#" {
					dir = "E"
				} else {
					index[0]--
					newVal := index[0]*1000 + index[1]
					if setN[newVal] {
						masterSet[k] = true
						break
					} else {
						setN[newVal] = true
					}
				}
			} else if dir == "E" {
				if newGrid[index[0]][index[1]+1] == "#" {
					dir = "S"
				} else {
					index[1]++
					newVal := index[0]*1000 + index[1]
					if setE[newVal] {
						masterSet[k] = true
						break
					} else {
						setE[newVal] = true
					}
				}
			} else if dir == "S" {
				if newGrid[index[0]+1][index[1]] == "#" {
					dir = "W"
				} else {
					index[0]++
					newVal := index[0]*1000 + index[1]
					if setS[newVal] {
						masterSet[k] = true
						break
					} else {
						setS[newVal] = true
					}
				}
			} else if dir == "W" {
				if newGrid[index[0]][index[1]-1] == "#" {
					dir = "N"
				} else {
					index[1]--
					newVal := index[0]*1000 + index[1]
					if setW[newVal] {
						masterSet[k] = true
						break
					} else {
						setW[newVal] = true
					}
				}
			}
		}
	}
	return masterSet
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
