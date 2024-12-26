package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"2024/common"
)

type Registers struct {
	A int
	B int
	C int
}

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	bytes, err := buildData(data)

	if err != nil {
		log.Fatalf("failed to build data: %v", err)
	}

	var LENGTH int
	var NUMBER_OF_BYTES int
	if filePath == "example.txt" {
		LENGTH = 7
		NUMBER_OF_BYTES = 12
	} else if filePath == "input.txt" {
		LENGTH = 71
		NUMBER_OF_BYTES = 1024
	} else {
		LENGTH = -1
		NUMBER_OF_BYTES = -1
	}

	if LENGTH == -1 || NUMBER_OF_BYTES == -1 {
		log.Fatalf("invalid file path: %s", filePath)
	}

	grid := make([][]string, LENGTH)
	for i := 0; i < LENGTH; i++ {
		grid[i] = make([]string, LENGTH)
		for j := 0; j < LENGTH; j++ {
			grid[i][j] = "."
		}
	}

	for index, b := range bytes {
		if index == NUMBER_OF_BYTES {
			break
		}
		grid[b[1]][b[0]] = "X"
	}

	// common.GridToTxt(grid, "output.txt")

	var START []int
	var END []int
	if filePath == "example.txt" {
		START = []int{0, 0}
		END = []int{6, 6}
	} else if filePath == "input.txt" {
		START = []int{0, 0}
		END = []int{70, 70}
	} else {
		START = []int{-1, -1}
		END = []int{-1, -1}
	}

	if START[0] == -1 || START[1] == -1 || END[0] == -1 || END[1] == -1 {
		log.Fatalf("invalid file path: %s", filePath)
	}

	unobstructed := findUnobstructedCells(grid, START, END)

	// for y := range grid {
	// 	for x := range grid[y] {
	// 		if _, ok := unobstructed[[2]int{x, y}]; ok {
	// 			grid[y][x] = "O"
	// 		}
	// 	}
	// }

	// common.GridToTxt(grid, filePath[:len(filePath)-4]+"_output.txt")

	fmt.Println("Unobstructed cells:", len(unobstructed))
	for cell := range unobstructed {
		fmt.Println(cell, unobstructed[cell])
	}

	// Part 1 Summary:
	// After running this code and checking the output file...
	// we can see that two of the unobstructed cells are the fastest paths out...
	// one from the start and the other from the end...
	// get the manhattan distance between these two points (clean path)...

	// Part 2 Approach:
	// In analyzing the output file, we can identify two cells as being choke points...
	// There must be a clear path...
	// From the start to choke point 1...
	// From choke point 1 to choke point 2...
	// And from choke point 2 to the end...

	START2 := [2]int{0, 0}
	CHOKE1 := [2]int{12, 50}
	CHOKE2 := [2]int{42, 54}
	END2 := [2]int{70, 70}

	initialVisited1 := make(map[[2]int]bool)
	initialVisited1[[2]int{CHOKE1[0], CHOKE1[1] - 1}] = true
	initialVisited2 := make(map[[2]int]bool)
	initialVisited2[[2]int{CHOKE2[0], CHOKE2[1] - 1}] = true

	path1 := findOnePath(grid, START2, CHOKE1, make(map[[2]int]bool))
	path2 := findOnePath(grid, CHOKE1, CHOKE2, initialVisited1)
	path3 := findOnePath(grid, CHOKE2, END2, initialVisited2)

	// fmt.Println("Path from START2 to CHOKE1:", path1)
	// fmt.Println("Path from CHOKE1 to CHOKE2:", path2)
	// fmt.Println("Path from CHOKE2 to END2:", path3)

	for i := NUMBER_OF_BYTES; i < len(bytes); i++ {
		x, y := bytes[i][0], bytes[i][1]
		val := [2]int{x, y}
		grid[y][x] = "X"
		if contains(path1, val) {
			fmt.Println(i, val, "is in path1... searching for a new path...")
			path1 = findOnePath(grid, START2, CHOKE1, make(map[[2]int]bool))
			if len(path1) == 0 {
				fmt.Println("No path from START2 to CHOKE1... searching for a new path...")
				break
			}
		}
		if contains(path2, val) {
			fmt.Println(i, val, "is in path2... searching for a new path...")
			new_visited1 := make(map[[2]int]bool)
			new_visited1[[2]int{CHOKE1[0], CHOKE1[1] - 1}] = true
			path2 = findOnePath(grid, CHOKE1, CHOKE2, new_visited1)
			if len(path2) == 0 {
				fmt.Println("No path from CHOKE1 to CHOKE2... searching for a new path...")
				break
			}
		}
		if contains(path3, val) {
			fmt.Println(i, val, "is in path3... searching for a new path...")
			new_visited2 := make(map[[2]int]bool)
			new_visited2[[2]int{CHOKE2[0], CHOKE2[1] - 1}] = true
			path3 = findOnePath(grid, CHOKE2, END2, new_visited2)
			if len(path3) == 0 {
				fmt.Println("No path from CHOKE2 to END2... searching for a new path...")
				break
			}
		}
	}

	common.GridToTxt(grid, filePath[:len(filePath)-4]+"_output2.txt")
}

func contains(path [][2]int, val [2]int) bool {
	for _, v := range path {
		if v == val {
			return true
		}
	}
	return false
}

func findOnePath(grid [][]string, start, end [2]int, visited map[[2]int]bool) [][2]int {
	queue := [][][2]int{{{start[0], start[1]}}}
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		x, y := path[len(path)-1][0], path[len(path)-1][1]

		if x == end[0] && y == end[1] {
			return path
		}

		if visited[[2]int{x, y}] {
			continue
		}
		visited[[2]int{x, y}] = true

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newY][newX] == "." {
				newPath := append([][2]int{}, path...)
				newPath = append(newPath, [2]int{newX, newY})
				queue = append(queue, newPath)
			}
		}
	}

	fmt.Println("No path found")
	return [][2]int{}
}

func findUnobstructedCells(grid [][]string, start []int, end []int) map[[2]int]int {
	unobstructed := make(map[[2]int]int)
	visited := make(map[[2]int]bool)

	move1(start[0], start[1], 0, grid, &unobstructed, &visited)
	move1(end[0], end[1], 0, grid, &unobstructed, &visited)
	return unobstructed
}

func move1(x1, y1, steps1 int, grid [][]string, unobstructed *map[[2]int]int, visited *map[[2]int]bool) {
	LENGTH := len(grid)
	queue := [][3]int{{x1, y1, steps1}}

	for len(queue) > 0 {
		x, y, steps := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		upable := y-1 >= 0 && grid[y-1][x] == "."
		downable := y+1 < LENGTH && grid[y+1][x] == "."
		leftable := x-1 >= 0 && grid[y][x-1] == "."
		rightable := x+1 < LENGTH && grid[y][x+1] == "."

		if upable && downable && leftable && rightable {
			if val, exists := (*unobstructed)[[2]int{x, y}]; exists {
				(*unobstructed)[[2]int{x, y}] = min(steps, val)
			} else {
				(*unobstructed)[[2]int{x, y}] = steps
			}
			continue
		}

		(*visited)[[2]int{x, y}] = true

		// up
		if upable && (!(*visited)[[2]int{x, y - 1}] || (*unobstructed)[[2]int{x, y - 1}] > steps+1) {
			queue = append(queue, [3]int{x, y - 1, steps + 1})
		}

		// down
		if downable && (!(*visited)[[2]int{x, y + 1}] || (*unobstructed)[[2]int{x, y + 1}] > steps+1) {
			queue = append(queue, [3]int{x, y + 1, steps + 1})
		}

		// left
		if leftable && (!(*visited)[[2]int{x - 1, y}] || (*unobstructed)[[2]int{x - 1, y}] > steps+1) {
			queue = append(queue, [3]int{x - 1, y, steps + 1})
		}

		// right
		if rightable && (!(*visited)[[2]int{x + 1, y}] || (*unobstructed)[[2]int{x + 1, y}] > steps+1) {
			queue = append(queue, [3]int{x + 1, y, steps + 1})
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildData(data []string) ([][]int, error) {
	var matrix [][]int
	for _, line := range data {
		var row []int
		values := strings.Split(line, ",")
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return matrix, nil
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
