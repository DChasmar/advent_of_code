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

	patterns, designs, err := buildData(data)
	if err != nil {
		log.Fatalf("failed to build data: %v", err)
	}

	result1 := solve1(patterns, designs)
	fmt.Println("Result 1:", result1)

	result2 := solve2(patterns, designs)
	fmt.Println("Result 2:", result2)

}

func solve2(patterns []string, designs []string) int {
	count := 0
	for _, design := range designs {
		count += findPatterns(design, patterns)
	}
	return count
}

func findPatterns(design string, patterns []string) int {
	n := len(design)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for _, pattern := range patterns {
			m := len(pattern)
			if i >= m && design[i-m:i] == pattern {
				dp[i] += dp[i-m]
			}
		}
	}
	return dp[n]
}

func solve1(patterns []string, designs []string) int {
	count := 0
	memo := make(map[string]bool)
	for _, pattern := range patterns {
		memo[pattern] = true
	}
	for _, design := range designs {
		count += findMatch(design, memo)
	}
	return count
}

func findMatch(design string, memo map[string]bool) int {
	if memo[design] {
		return 1
	}
	for i := 1; i < len(design); i++ {
		if memo[design[:i]] && findMatch(design[i:], memo) == 1 {
			return 1
		}
	}

	return 0
}

func buildData(data []string) ([]string, []string, error) {
	if len(data) == 0 {
		return nil, nil, fmt.Errorf("no data to process")
	}

	// Process the first line
	firstLineArray := strings.Split(data[0], ", ")

	// Process the remaining lines into a single array
	var remainingLinesArray []string
	remainingLinesArray = append(remainingLinesArray, data[2:]...) // Skip the empty line after the first line

	return firstLineArray, remainingLinesArray, nil
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
