package main

import (
	"fmt"
	"log"

	"2024/common"
)

type Key struct {
	a, b, c, d int
}

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	var nums []int
	var nums2 []int
	for _, line := range data {
		num := common.LineToInt(line)
		nums = append(nums, num)
		nums2 = append(nums2, num)
	}

	result1 := solve1(nums, 2000)
	fmt.Println("Part 1:", result1)

	result2 := solve2(nums2, 2000)
	fmt.Println("Part 2:", result2)
	// 4231 is too high
	// 2678 is too high
	// 1578 is too high
}

func solve2(nums []int, times int) int {
	lenNums := len(nums)
	values := make(map[Key][]int)

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		lastOnes := val % 10
		lastFour := make([]int, 0, 4)
		for j := 0; j < times; j++ {
			val = process(val)
			ones := val % 10
			diff := ones - lastOnes
			if diff < -9 || diff > 9 {
				log.Fatalf("diff is out of range: %d", diff)
			}
			lastFour = append(lastFour, diff)
			if len(lastFour) > 4 {
				lastFour = lastFour[1:]
			}
			if len(lastFour) == 4 {
				key := Key{lastFour[0], lastFour[1], lastFour[2], lastFour[3]}
				if _, ok := values[key]; !ok {
					values[key] = make([]int, lenNums)
				}
				if values[key][i] == 0 {
					values[key][i] = ones
				}
			}
			lastOnes = ones
			if ones > 9 || ones < 0 {
				log.Fatalf("ones is out of range: %d", ones)
			}
			if lastOnes > 9 || lastOnes < 0 {
				log.Fatalf("lastOnes is out of range: %d", lastOnes)
			}
		}
	}

	// Calculate the sum of the values of the four registers
	maxSum := 0
	for key, value := range values {
		sum := common.SumIntArray(value)
		if sum > maxSum {
			fmt.Println(key, sum)
			maxSum = sum
		}
	}

	return maxSum
}

func solve1(nums []int, times int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < times; j++ {
			nums[i] = process(nums[i])
		}
	}
	return common.SumIntArray(nums)
}

func process(num int) int {
	num = step1(num)
	num = step2(num)
	num = step3(num)
	return num
}

func step1(num int) int {
	new := num * 64
	num = mix(new, num)
	return prune(num)
}

func step2(num int) int {
	new := num / 32
	num = mix(new, num)
	return prune(num)
}

func step3(num int) int {
	new := num * 2048
	num = mix(new, num)
	return prune(num)
}

func mix(a, b int) int {
	return common.BitwiseXOR(a, b)
}

func prune(a int) int {
	return a % 16777216
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
