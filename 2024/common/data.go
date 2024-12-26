package common

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func NumberOccurrencesDict(s []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range s {
		dict[num]++
	}
	return dict
}

func IntArrayToSet(s []int) map[int]bool {
	set := make(map[int]bool)
	for _, num := range s {
		set[num] = true
	}
	return set
}

func FilterArray(array [][]int, indexesToFilter map[int]bool) [][]int {
	filtered := [][]int{}
	for i, item := range array {
		if _, ok := indexesToFilter[i]; !ok {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func ArraysOfArrayLessOneValue(array []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(array); i++ {
		subArray := make([]int, 0, len(array)-1)
		subArray = append(subArray, array[:i]...)
		subArray = append(subArray, array[i+1:]...)
		result = append(result, subArray)
	}
	return result
}

func LinesToGrid(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = []string(strings.Split(line, ""))
	}
	return grid
}

func LinesToIntGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = LineToIntArray(line)
	}
	return grid
}

func LineToIntArray(line string) []int {
	array := make([]int, 0, len(line))
	for _, char := range line {
		array = append(array, int(char-'0'))
	}
	return array
}

func LineToIntArraySplitSpace(line string) []int {
	parts := strings.Fields(line)
	arr := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("failed to convert %s to int: %v", part, err)
		}
		arr[i] = num
	}
	return arr
}

func GridToTxt(grid [][]string, filename string) {
	HEIGHT := len(grid)
	WIDTH := len(grid[0])
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			file.WriteString(grid[y][x])
		}
		file.WriteString("\n")
	}
	fmt.Println("Grid exported to", filename)
}

func LineToArray(line string) []string {
	return strings.Split(line, "")
}

func LineToByteArray(line string) []byte {
	return []byte(line)
}

func ByteArrayToInt(b []byte) int {
	num, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatalf("failed to convert %s to int: %v", string(b), err)
	}
	return num
}

func LineToInt(line string) int {
	num, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("failed to convert %s to int: %v", line, err)
	}
	return num
}
