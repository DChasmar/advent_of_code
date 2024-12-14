package main

import (
	"fmt"
	"log"

	"2024/common"
)

func main() {
	filePath := "input.txt"
	s, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	arr := common.LineToIntArray(s[0])

	result1 := solve1(arr)
	fmt.Println(result1)

	result2 := solve2(arr)
	fmt.Println(result2)
}

func solve2(arr []int) int {
	spaces := make([]int, (len(arr)-1)/2)
	values := make([]int, (len(arr)+1)/2)

	for i := range arr {
		if i%2 == 0 {
			values[i/2] = arr[i]
		} else {
			spaces[(i-1)/2] = arr[i]
		}
	}

	furthest := make(map[int]int)
	for i := len(values) - 1; i >= 0; i-- {
		val := values[i]
		if val > 0 {
			if _, exists := furthest[val]; !exists {
				furthest[val] = i
			}
		}
		every := true
		for j := 1; j <= 9; j++ {
			if _, exists := furthest[j]; !exists {
				every = false
				break
			}
		}
		if every {
			break
		}
	}

	i := 0
	checksum := 0

	for index := range len(arr) {
		if index%2 == 0 {
			valuesIndex := index / 2
			for range arr[index] {
				if values[valuesIndex] != 0 {
					new := (index / 2) * i
					checksum += new
				}
				i++
			}
		} else {
			spacesIndex := (index - 1) / 2
			for spaces[spacesIndex] > 0 {
				freeSpace := spaces[spacesIndex]
				amountToFill, farIndex := chooseFurthest(furthest, freeSpace, spacesIndex)
				if amountToFill == 0 || farIndex == 0 {
					i += spaces[spacesIndex]
					spaces[spacesIndex] = 0
				} else {
					for range amountToFill {
						new := farIndex * i
						checksum += new
						i++
					}
					spaces[spacesIndex] -= amountToFill
					values[farIndex] -= amountToFill
					furthest = updateFurthest(furthest, values, amountToFill, farIndex)
				}
			}
		}
	}
	// fmt.Println(furthest)
	// fmt.Println(values[:20])
	// fmt.Println(spaces[:20])
	// fmt.Println(i)
	return checksum
}

func chooseFurthest(furthest map[int]int, freeSpace int, spacesIndex int) (int, int) {
	amountToFill := 0
	farIndex := 0
	for i := freeSpace; i > 0; i-- {
		if val, exists := furthest[i]; exists {
			if exists && val > spacesIndex && val > farIndex {
				amountToFill = i
				farIndex = val
			}
		}
	}
	return amountToFill, farIndex
}

func updateFurthest(furthest map[int]int, values []int, amountToFill int, farIndex int) map[int]int {
	newFarIndex := farIndex
	for newFarIndex > 0 {
		newFarIndex--
		if values[newFarIndex] == amountToFill {
			break
		}
	}
	furthest[amountToFill] = newFarIndex
	return furthest
}

func solve1(arr []int) int {
	i := 0
	j := len(arr) - 1
	far := arr[j]
	count := sum(arr)

	checksum := 0

	done := false

	for index, val := range arr {
		for range val {
			if index%2 == 0 {
				new := (index / 2) * i
				checksum += new
			} else {
				j, far = moveForward(arr, j, far)
				new := (j / 2) * i
				checksum += new
			}
			i++
			if i >= count {
				done = true
			}
			if done {
				break
			}
		}
		if done {
			break
		}
	}
	return checksum
}

func moveForward(arr []int, j int, far int) (int, int) {
	if far > 0 {
		far--
	} else {
		j -= 2
		far = arr[j]
		far--
	}
	return j, far
}

func sum(arr []int) int {
	sum := 0
	for index, val := range arr {
		if index%2 == 0 {
			sum += val
		}
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

// Part 2:
// 33398771123038 is too high
// 10862485047414 is too high
// 6415666220005 is ...
