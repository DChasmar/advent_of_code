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
	keys, locks := processGrids(data)

	// fmt.Println("Keys:")
	// for _, key := range keys {
	// 	printGrid(key)
	// }

	// fmt.Println("Locks:")
	// for _, lock := range locks {
	// 	printGrid(lock)
	// }

	fmt.Println("No. of Keys:", len(keys))
	fmt.Println("No. of Locks:", len(locks))

	var keyCounts [][5]int
	for _, key := range keys {
		counts := [5]int{}
		for i := 1; i < len(key)-1; i++ {
			for j := 0; j < 5; j++ {
				if key[i][j] == "#" {
					counts[j]++
				}
			}
		}
		keyCounts = append(keyCounts, counts)
	}

	var lockCounts [][5]int
	for _, lock := range locks {
		counts := [5]int{}
		for i := 1; i < len(lock)-1; i++ {
			for j := 0; j < 5; j++ {
				if lock[i][j] == "#" {
					counts[j]++
				}
			}
		}
		lockCounts = append(lockCounts, counts)
	}

	total := 0

	for _, key := range keyCounts {
		for _, lock := range lockCounts {
			over := false
			for k := 0; k < 5; k++ {
				if key[k]+lock[k] > 5 {
					over = true
					break
				}
			}
			if !over {
				total++
			}
		}
	}
	fmt.Println("Total:", total)
}

func processGrids(lines []string) ([][][]string, [][][]string) {
	var keys, locks [][][]string
	var currentGrid [][]string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if len(currentGrid) > 0 {
				if isTopRowAll(currentGrid, '.') {
					keys = append(keys, currentGrid)
				} else if isTopRowAll(currentGrid, '#') {
					locks = append(locks, currentGrid)
				}
				currentGrid = nil
			}
		} else {
			currentGrid = append(currentGrid, strings.Split(line, ""))
		}
	}

	// Process the last grid if the file does not end with an empty line
	if len(currentGrid) > 0 {
		if isTopRowAll(currentGrid, '.') {
			keys = append(keys, currentGrid)
		} else if isTopRowAll(currentGrid, '#') {
			locks = append(locks, currentGrid)
		}
	}

	return keys, locks
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func isTopRowAll(grid [][]string, char rune) bool {
	if len(grid) == 0 {
		return false
	}
	for _, cell := range grid[0] {
		if cell != string(char) {
			return false
		}
	}
	return true
}

// func printGrid(grid [][]string) {
// 	for _, row := range grid {
// 		fmt.Println(strings.Join(row, ""))
// 	}
// 	fmt.Println()
// }
