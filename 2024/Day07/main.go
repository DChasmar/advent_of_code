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
	s, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}
	combos := buildNumDict(s)
	fmt.Println(len(combos))

	sum := 0
	goodCombosSet := make(map[int]bool)

	for index, numDict := range combos {
		for num, nums := range numDict {
			if canReachTarget1(nums, num) {
				sum += num
				goodCombosSet[index] = true
			}
		}
	}
	fmt.Println(sum)

	altSum := 0
	altGoodCombosSet := make(map[int]bool)

	for index, numDict := range combos {
		if _, ok := goodCombosSet[index]; ok {
			continue
		}
		for num, nums := range numDict {
			if canReachTarget2(nums, num) {
				altSum += num
				altGoodCombosSet[index] = true
				break
			}
		}
	}

	fmt.Println(sum + altSum)
}

func concatenate(leftNum, rightNum int) int {
	str := strconv.Itoa(leftNum) + strconv.Itoa(rightNum)
	num, ok := strconv.Atoi(str)
	if ok != nil {
		log.Fatalf("failed to convert %s to int: %v", str, ok)
	}
	return num
}

func canReachTarget2(numbers []int, target int) bool {
	return search2(numbers, 0, numbers[0], target)
}

func search2(numbers []int, index int, currentValue int, target int) bool {
	// Base case: if we've used all numbers, check if we've reached the target
	if currentValue > target {
		return false
	}
	if index == len(numbers)-1 {
		return currentValue == target
	}

	// Recursive case: try both adding and multiplying the next number
	nextIndex := index + 1
	nextNumber := numbers[nextIndex]

	// Try adding the next number
	if search2(numbers, nextIndex, currentValue+nextNumber, target) {
		return true
	}

	// Try multiplying the next number
	if search2(numbers, nextIndex, currentValue*nextNumber, target) {
		return true
	}

	// Try concatenating the next number
	if search2(numbers, nextIndex, concatenate(currentValue, nextNumber), target) {
		return true
	}

	// If neither adding nor multiplying leads to the target, return false
	return false
}

func canReachTarget1(numbers []int, target int) bool {
	return search1(numbers, 0, numbers[0], target)
}

func search1(numbers []int, index int, currentValue int, target int) bool {
	// Base case: if we've used all numbers, check if we've reached the target
	if currentValue > target {
		return false
	}
	if index == len(numbers)-1 {
		return currentValue == target
	}

	// Recursive case: try both adding and multiplying the next number
	nextIndex := index + 1
	nextNumber := numbers[nextIndex]

	// Try adding the next number
	if search1(numbers, nextIndex, currentValue+nextNumber, target) {
		return true
	}

	// Try multiplying the next number
	if search1(numbers, nextIndex, currentValue*nextNumber, target) {
		return true
	}

	// If neither adding nor multiplying leads to the target, return false
	return false
}

func buildNumDict(s []string) []map[int][]int {
	combos := []map[int][]int{}
	for _, line := range s {
		str := strings.Split(line, ": ")
		num, ok := strconv.Atoi(str[0])
		if ok != nil {
			log.Fatalf("failed to convert %s to int: %v", str[0], ok)
		}
		numsArr := strings.Split(str[1], " ")
		nums := make([]int, len(numsArr))
		for i, n := range numsArr {
			num, ok := strconv.Atoi(n)
			if ok != nil {
				log.Fatalf("failed to convert %s to int: %v", n, ok)
			}
			nums[i] = num
		}
		numDict := make(map[int][]int)
		numDict[num] = nums
		combos = append(combos, numDict)
	}
	return combos
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// Part 1: 2664459966384 is too low
// Part 1: 2664460013123 is correct
// Part 2: 31304694714407 is too low
