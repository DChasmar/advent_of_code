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

	grid := common.LinesToGrid(s)

	dict := make(map[string][][]int)

	for r, row := range grid {
		for c, cell := range row {
			if cell != "." {
				dict[cell] = append(dict[cell], []int{r, c})
			}
		}
	}

	set1 := make(map[int]bool)

	for val := range dict {
		antiCoords1 := getAntinodes1(dict[val])
		for _, coord := range antiCoords1 {
			if coord[0] >= 0 && coord[0] < len(grid) && coord[1] >= 0 && coord[1] < len(grid[0]) {
				set1[coord[0]*1000+coord[1]] = true
			}
		}
	}

	fmt.Println(len(set1))

	set2 := make(map[int]bool)

	for val := range dict {
		antiCoords2 := getAntinodes2(dict[val], grid)
		for _, coord := range antiCoords2 {
			set2[coord[0]*1000+coord[1]] = true
		}
	}

	fmt.Println(len(set2))
}

func getAntinodes1(coords [][]int) [][]int {
	antinodes := [][]int{}
	for i := range coords {
		for j := range coords {
			xDiff := coords[i][0] - coords[j][0]
			yDiff := coords[i][1] - coords[j][1]
			if xDiff != 0 || yDiff != 0 {
				antinodes = append(antinodes, []int{coords[i][0] + xDiff, coords[i][1] + yDiff})
				antinodes = append(antinodes, []int{coords[j][0] - xDiff, coords[j][1] - yDiff})
			}
		}
	}
	return antinodes
}

func getAntinodes2(coords [][]int, grid [][]string) [][]int {
	antinodes := [][]int{}
	for i := range coords {
		for j := range coords {
			xDiff := coords[i][0] - coords[j][0]
			yDiff := coords[i][1] - coords[j][1]
			ax := coords[i][0] + xDiff
			ay := coords[i][1] + yDiff
			if xDiff != 0 || yDiff != 0 {
				antinodes = append(antinodes, []int{coords[i][0], coords[i][1]})
				antinodes = append(antinodes, []int{coords[j][0], coords[j][1]})
				for ax >= 0 && ax < len(grid) && ay >= 0 && ay < len(grid[0]) {
					antinodes = append(antinodes, []int{ax, ay})
					ax += xDiff
					ay += yDiff
				}
				ax = coords[j][0] - xDiff
				ay = coords[j][1] - yDiff
				for ax >= 0 && ax < len(grid) && ay >= 0 && ay < len(grid[0]) {
					antinodes = append(antinodes, []int{ax, ay})
					ax -= xDiff
					ay -= yDiff
				}
			}
		}
	}
	return antinodes
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
