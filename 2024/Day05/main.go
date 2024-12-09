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

	orders, preDict, postDict := getData(s)

	validOrders := getValidOrderIndices(orders, postDict)

	result1 := solve1(orders, validOrders)
	fmt.Println(result1)

	validSet := make(map[int]bool)
	for _, index := range validOrders {
		validSet[index] = true
	}
	result2 := solve2(orders, validSet, preDict)
	fmt.Println(result2)
}

func solve2(orders [][]string, validSet map[int]bool, preDict map[string][]string) int {
	good := 0
	for i, order := range orders {
		if validSet[i] {
			continue
		}
		set := make(map[string]bool)
		for _, item := range order {
			set[item] = true
		}
		reorderedOrder := reorderOrder(order, set, preDict)
		middle, ok := getMiddleNumber(reorderedOrder)
		if ok {
			good += middle
		}
	}
	return good
}

func reorderOrder(order []string, set map[string]bool, preDict map[string][]string) []string {
	reorderedOrder := make([]string, 0, len(order))

	for len(set) > 0 {
		for _, item := range order {
			if _, exists := set[item]; exists {
				valid := true
				for _, pre := range preDict[item] {
					if set[pre] {
						valid = false
						break
					}
				}
				if valid {
					delete(set, item)
					reorderedOrder = append(reorderedOrder, item)
				}
			}
		}
	}
	return reorderedOrder
}

func solve1(orders [][]string, validIndices []int) int {
	good := 0
	for _, i := range validIndices {
		middle, ok := getMiddleNumber(orders[i])
		if ok {
			good += middle
		}
	}
	return good
}

func getValidOrderIndices(orders [][]string, postDict map[string][]string) []int {
	var validIndices []int
	for i, order := range orders {
		set := make(map[string]bool)
		valid := true
		for _, item := range order {
			set[item] = true
			for _, post := range postDict[item] {
				if set[post] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			validIndices = append(validIndices, i)
		}
	}
	return validIndices
}

func getData(s []string) ([][]string, map[string][]string, map[string][]string) {
	preDict := make(map[string][]string)
	postDict := make(map[string][]string)
	var orders [][]string

	for _, line := range s {
		if len(line) > 2 && line[2] == '|' {
			instruction := strings.Split(line, "|")
			if len(instruction) == 2 {
				preDict[instruction[1]] = append(preDict[instruction[1]], instruction[0])
				postDict[instruction[0]] = append(postDict[instruction[0]], instruction[1])
			}
		} else if len(line) > 2 && line[2] == ',' {
			orders = append(orders, strings.Split(line, ","))
		}
	}

	return orders, preDict, postDict
}

func getMiddleNumber(s []string) (int, bool) {
	if len(s) < 3 {
		return 0, false
	}

	num, err := strconv.Atoi(s[len(s)/2])
	if err != nil {
		return 0, false
	}

	return num, true
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
