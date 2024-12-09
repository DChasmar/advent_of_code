package common

import "strings"

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
