package main

import (
	"fmt"
	"log"
	"strings"

	"2024/common"
)

const (
	tByte = byte('t')
)

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}
	graph := createGraph(data)

	result1 := solve1(graph)
	fmt.Println("Result 1: ", result1)

	resultStr, resultInt := solve2(graph)
	// Split resultStr every 2 characters with a comma...
	finalStr := ""
	for i := 0; i < len(resultStr); i += 2 {
		finalStr += resultStr[i : i+2]
		if i+2 < len(resultStr) {
			finalStr += ","
		}
	}
	fmt.Println("Result 2: ", finalStr, resultInt)
}

func solve1(graph map[string][]string) int {
	set := make(map[string]bool)
	for k, v := range graph {
		if []byte(k)[0] != tByte {
			continue
		}
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				if graph[v[i]] != nil && graph[v[j]] != nil {
					if containsStr(graph[v[i]], v[j]) {
						strArr := []string{k, v[i], v[j]}
						strArr = sortStringArray(strArr)
						str := joinStringArray(strArr, "-")
						set[str] = true
					}
				}
			}
		}
	}
	return len(set)
}

func solve2(graph map[string][]string) (string, int) {
	tracker := make(map[string]int)
	omit := 0
	for {
		fmt.Println(omit)
		for k, v := range graph {
			combos := arrCombinations(v, omit)
			for _, combo := range combos {
				strArr := append([]string{k}, combo...)
				strArr = sortStringArray(strArr)
				str := joinStringArray(strArr, "")
				tracker[str]++
			}
		}
		for k, v := range tracker {
			if len(k)/2 == v {
				return k, v
			}
		}
		omit++
	}
}

func arrCombinations(arr []string, omit int) [][]string {
	var result [][]string
	omitCombinations := combinations(len(arr), omit)

	for _, combo := range omitCombinations {
		var temp []string
		for i := 0; i < len(arr); i++ {
			if containsInt(combo, i) {
				temp = append(temp, arr[i])
			}
		}
		result = append(result, temp)
	}
	return result
}

func combinations(nums, omit int) [][]int {
	var result [][]int
	var current []int
	combine(nums, omit, 0, current, &result)
	return result
}

func combine(nums, omit, start int, current []int, result *[][]int) {
	if len(current) == nums-omit {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i < nums; i++ {
		current = append(current, i)
		combine(nums, omit, i+1, current, result)
		current = current[:len(current)-1]
	}
}

func containsStr(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func containsInt(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func sortStringArray(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func joinStringArray(arr []string, char string) string {
	return strings.Join(arr, char)
}

func createGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range lines {
		// Split the line into two-character strings
		parts := strings.Split(line, "-")
		if len(parts) != 2 || len(parts[0]) != 2 || len(parts[1]) != 2 {
			log.Fatalf("invalid line format: %s", line)
		}
		first := parts[0]
		second := parts[1]

		// Add the relationship to the graph
		graph[first] = append(graph[first], second)
		graph[second] = append(graph[second], first)
	}
	return graph
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
