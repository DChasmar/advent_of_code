package main

import (
	"fmt"
	"log"
	"regexp"
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

	// fmt.Println(len(s)) // There are 6 lines in the input file
	result1 := solve1(s)
	fmt.Println(result1)
	result2 := solve2(s)
	fmt.Println(result2)
}

func tupleAtIndex(s string, index int) (string, bool) {
	// Define the regular expression for a tuple
	re := regexp.MustCompile(`\(\d+,\d+\)`)

	// Ensure the index is within the valid range
	if index < 0 || index >= len(s) {
		return "", false
	}

	// Extract the substring from the index to the end of the string
	substring := s[index:]

	// Find the first match of the tuple in the substring
	match := re.FindStringIndex(substring)

	// Check if the match starts at the given index
	if match != nil && match[0] == 0 {
		// Extract the matched tuple
		matchedTuple := substring[match[0]:match[1]]
		// Extract the \d+,\d+ part from the matched tuple
		innerMatch := regexp.MustCompile(`\d+,\d+`).FindString(matchedTuple)
		return innerMatch, true
	}

	return "", false
}

func stringToTuple(s string) ([]int, bool) {
	tuple := strings.Split(s, ",")
	intTuple := make([]int, len(tuple))
	for i, v := range tuple {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, false
		}
		intTuple[i] = num
	}
	return intTuple, true
}

func lCheck(s string, index int) ([]int, bool) {
	if s[index] == 'l' {
		if s[index-1] == 'u' && s[index-2] == 'm' {
			index++
			tuple, ok := tupleAtIndex(s, index)
			if ok {
				newTuple, ok := stringToTuple(tuple)
				if ok {
					return newTuple, true
				}
			}
		}
	}
	return nil, false
}

func findL(s string) [][]int {
	var tuples [][]int
	index := 2
	for index < len(s) {
		tuple, ok := lCheck(s, index)
		if ok {
			tuples = append(tuples, tuple)
		}
		index++
	}
	return tuples
}

func findD(s string, index int) (bool, bool) {
	if s[index] == 'd' && len(s)-index >= 4 && s[index+1] == 'o' && s[index+2] == '(' && s[index+3] == ')' {
		return true, true
	} else if s[index] == 'd' && len(s)-index >= 7 && s[index+1] == 'o' && s[index+2] == 'n' && s[index+3] == '\'' && s[index+4] == 't' && s[index+5] == '(' && s[index+6] == ')' {
		return true, false
	}
	return false, false
}

func tupleProduct(t []int) int {
	product := 1
	for _, v := range t {
		product *= v
	}
	return product
}

func solve1(s []string) int {
	allTuples := [][]int{}
	for _, line := range s {
		lineTuples := findL(line)
		allTuples = append(allTuples, lineTuples...)
	}
	sum := 0
	for _, tuple := range allTuples {
		sum += tupleProduct(tuple)
	}
	return sum
}

func solve2(s []string) int {
	can := true
	allTuples := [][]int{}
	for _, line := range s {
		index := 2
		for index < len(line) {
			if can {
				tuple, ok := lCheck(line, index)
				if ok {
					allTuples = append(allTuples, tuple)
				}
			}
			change, newCan := findD(line, index)
			if change {
				can = newCan
			}
			index++
		}
	}

	sum := 0
	for _, tuple := range allTuples {
		sum += tupleProduct(tuple)
	}
	return sum
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
