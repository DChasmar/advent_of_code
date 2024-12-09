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
	array1, array2, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	result1 := solve1(array1, array2)
	fmt.Println(result1)
	result2 := solve2(array1, array2)
	fmt.Println(result2)
}

func solve2(array1, array2 []int) int {
	dict := common.NumberOccurrencesDict(array2)

	sum := 0

	for _, num := range array1 {
		if dict[num] != 0 {
			sum += num * dict[num]
		}
	}

	return sum
}

func solve1(array1, array2 []int) int {
	common.SortArray(array1)
	common.SortArray(array2)

	sum := 0
	for i := 0; i < len(array1); i++ {
		sum += common.AbsDifference(array1[i], array2[i])
	}

	return sum
}

func readAndProcessFile(filePath string) ([]int, []int, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, nil, err
	}

	var array1, array2 []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("unexpected number of fields in line: %v", line)
		}

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert %v to int: %w", numbers[0], err)
		}

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert %v to int: %w", numbers[1], err)
		}

		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	return array1, array2, nil
}
