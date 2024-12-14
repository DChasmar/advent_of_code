package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"2024/common"
)

type Coordinates struct {
	A [2]int
	B [2]int
	P [2]int
}

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	objects := buildData(data)

	result1 := solve1(objects)
	fmt.Println(result1)

	// Increment the prize positions by 10000000000000 for part 2
	for i := range objects {
		objects[i].P[0] += 10000000000000
		objects[i].P[1] += 10000000000000
	}

	result2 := solve2(objects)
	fmt.Println(result2)
}

func solve2(objects []Coordinates) int {
	sum := 0
	for index, obj := range objects {
		sum += calc2(obj, index)
	}
	return sum
}

func calc2(obj Coordinates, index int) int {
	mx, bx := getEqLine(obj.A[0], obj.B[0], obj.P[0])
	my, by := getEqLine(obj.A[1], obj.B[1], obj.P[1])
	fmt.Println(mx, bx, my, by)
	if bx == -1 || by == -1 {
		fmt.Println("fail first", index)
		return 0
	}
	first, lcm := firstZeroZero(mx, bx, my, by)
	if first == -1 {
		// fmt.Println("fail second", index)
		return 0
	}
	pressesX1, pressesY1 := findNumberOfBPresses(obj, first)
	if pressesX1 == pressesY1 {
		return calcTokens(first, pressesX1)
	}
	pressesX2, pressesY2 := findNumberOfBPresses(obj, first+lcm)
	abs1 := common.AbsDifference(pressesX1, pressesY1)
	abs2 := common.AbsDifference(pressesX2, pressesY2)
	if abs1 < abs2 {
		// fmt.Println("fail third", index)
		return 0
	}
	slope := abs1 - abs2
	if abs1%slope == 0 {
		factors := abs1 / slope
		fmt.Println("factors", factors, slope, abs1, abs2)
		Apresses := first + factors*lcm
		good, Bpresses := checkPossible(obj, Apresses)
		if good {
			fmt.Println("good", index, Apresses, Bpresses)
			return calcTokens(Apresses, Bpresses)
		}
	}
	// fmt.Println("fail fourth", index)
	return 0
}

func findNumberOfBPresses(obj Coordinates, Apresses int) (int, int) {
	return (obj.P[0] - obj.A[0]*Apresses) / obj.B[0], (obj.P[1] - obj.A[1]*Apresses) / obj.B[1]
}

func firstZeroZero(mx, bx, my, by int) (int, int) {
	// Calculate the least common multiple (LCM) of bx and by
	lcm := common.LCM(mx, my)

	firstOccurence := common.FirstOccurence(mx, bx, my, by)
	if firstOccurence == -1 {
		return -1, -1
	} else {
		return firstOccurence, lcm
	}
}

func getEqLine(a, b, p int) (int, int) {
	Apresses := 0
	remIdx := -1
	remSet := make(map[int]bool)
	for Apresses <= 1000 {
		rem := (p - a*Apresses) % b
		if rem < 0 {
			rem += b
		}
		if remSet[rem] {
			return Apresses, remIdx
		}
		if rem == 0 {
			if remIdx == -1 {
				remIdx = Apresses
			}
		}
		remSet[rem] = true
		Apresses++
	}
	return -1, -1
}

func solve1(objects []Coordinates) int {
	sum := 0
	for _, obj := range objects {
		sum += calc1(obj)
	}
	return sum
}

func calc1(obj Coordinates) int {
	Apresses := 0
	for Apresses < 101 && obj.A[0]*Apresses <= obj.P[0] && obj.A[1]*Apresses <= obj.P[1] {
		good, Bpresses := checkPossible(obj, Apresses)
		if good {
			return calcTokens(Apresses, Bpresses)
		}
		Apresses++
	}
	return 0
}

func checkPossible(obj Coordinates, Apresses int) (bool, int) {
	noRem := (obj.P[0]-obj.A[0]*Apresses)%obj.B[0] == 0 && (obj.P[1]-obj.A[1]*Apresses)%obj.B[1] == 0
	eq := (obj.P[0]-obj.A[0]*Apresses)/obj.B[0] == (obj.P[1]-obj.A[1]*Apresses)/obj.B[1]
	if noRem && eq {
		return true, (obj.P[0] - obj.A[0]*Apresses) / obj.B[0]
	} else {
		return false, 0
	}
}

func calcTokens(Apresses int, Bpresses int) int {
	return 3*Apresses + Bpresses
}

func buildData(lines []string) []Coordinates {
	var objects []Coordinates
	var current Coordinates

	re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)|Button B: X\+(\d+), Y\+(\d+)|Prize: X=(\d+), Y=(\d+)`)

	for _, line := range lines {
		if line == "" {
			objects = append(objects, current)
			current = Coordinates{}
			continue
		}

		matches := re.FindStringSubmatch(line)
		if matches != nil {
			if matches[1] != "" {
				current.A[0], _ = strconv.Atoi(matches[1])
				current.A[1], _ = strconv.Atoi(matches[2])
			} else if matches[3] != "" {
				current.B[0], _ = strconv.Atoi(matches[3])
				current.B[1], _ = strconv.Atoi(matches[4])
			} else if matches[5] != "" {
				current.P[0], _ = strconv.Atoi(matches[5])
				current.P[1], _ = strconv.Atoi(matches[6])
			}
		}
	}

	// Append the last object if the file does not end with a blank line
	if current != (Coordinates{}) {
		objects = append(objects, current)
	}

	return objects
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// Part 2: 135969187403144 is too high
// Part 2: 112856557376557 is too high
