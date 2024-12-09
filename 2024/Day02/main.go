package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"2024/common"
)

func main() {
	filePath := "input.txt"
	arrays, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	result1 := solve1(arrays)
	fmt.Println(len(result1))

	result2 := solve2(arrays, result1)

	fmt.Println(len(result2))
	fmt.Println(len(result1) + len(result2))
}

func solve2(arrays [][]int, result1 []int) []int {
	set1 := common.IntArrayToSet(result1)
	filtered := common.FilterArray(arrays, set1)

	var rejectArrs [][][]int
	for i := 0; i < len(filtered); i++ {
		rejectArrs = append(rejectArrs, common.ArraysOfArrayLessOneValue(filtered[i]))
	}

	safe2 := []int{}

	for i := 0; i < len(rejectArrs); i++ {
		for j := 0; j < len(rejectArrs[i]); j++ {
			valid := checkArray(rejectArrs[i][j])
			if valid {
				safe2 = append(safe2, i)
				break
			}
		}
	}

	return safe2
}

func checkArray(arr []int) bool {
	sign := arr[0] > arr[1]
	valid := true
	for j := 0; j < len(arr)-1; j++ {
		newSign := arr[j] > arr[j+1]
		newAbs := common.AbsDifference(arr[j], arr[j+1])
		if sign != newSign || newAbs < 1 || newAbs > 3 {
			valid = false
			break
		}
	}
	return valid
}

func solve1(arrs [][]int) []int {
	safe := []int{}
	for i := 0; i < len(arrs); i++ {
		valid := checkArray(arrs[i])
		if valid {
			safe = append(safe, i)
		}
	}
	return safe
}

func readAndProcessFile(filePath string) ([][]int, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	var result [][]int

	for _, line := range lines {
		numbers := strings.Fields(line)
		var numArray []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("failed to convert %v to int: %w", numStr, err)
			}
			numArray = append(numArray, num)
		}
		result = append(result, numArray)
	}

	return result, nil
}
