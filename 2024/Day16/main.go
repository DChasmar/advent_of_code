package main

import (
	"container/heap"
	"fmt"
	"log"
	"strings"

	"2024/common"
)

type Position struct {
	row int
	col int
	dir int // 0 = up, 1 = right, 2 = down, 3 = left
}

type Node struct {
	pos   Position
	score int
	cost  int
	index int
	path  []Position // Add a path field to track the path taken
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	grid := buildData(data)
	optimalScore := solve1(grid)
	fmt.Println(optimalScore)

	// Call solve2 and print the result
	positions := solve2(grid, optimalScore)
	fmt.Println(len(positions))

	newGrid := make([][]string, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	for row := range newGrid {
		for col := range newGrid[row] {
			if positionsContain(positions, Position{row: row, col: col}) {
				newGrid[row][col] = "0"
			}
		}
	}

	fileName := filePath[:len(filePath)-4] + "_output.txt"
	common.GridToTxt(newGrid, fileName)

	posSet := make(map[int]bool)
	for _, pos := range positions {
		posSet[pos.row*1000+pos.col] = true
	}
	fmt.Println(len(posSet))
}

func positionsContain(positions []Position, pos Position) bool {
	for _, p := range positions {
		if p.row == pos.row && p.col == pos.col {
			return true
		}
	}
	return false
}

func solve1(grid [][]string) int {
	rowS, colS := findS(grid)
	rowE, colE := findE(grid)
	startingPos := Position{row: rowS, col: colS, dir: 1}
	return aStar(grid, startingPos, Position{row: rowE, col: colE})
}

func solve2(grid [][]string, optimal int) []Position {
	rowS, colS := findS(grid)
	rowE, colE := findE(grid)
	startingPos := Position{row: rowS, col: colS, dir: 1}
	return aStarWithPath(grid, startingPos, Position{row: rowE, col: colE}, optimal)
}

func aStar(grid [][]string, start, end Position) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{pos: start, score: 0, cost: 0})

	memo := make(map[Position]int)
	memo[start] = 0

	directions := []Position{
		{row: -1, col: 0, dir: 0}, // up
		{row: 0, col: 1, dir: 1},  // right
		{row: 1, col: 0, dir: 2},  // down
		{row: 0, col: -1, dir: 3}, // left
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		pos := node.pos

		if pos.row == end.row && pos.col == end.col {
			return node.cost
		}

		for _, d := range directions {
			newPos := Position{row: pos.row + d.row, col: pos.col + d.col, dir: d.dir}
			if newPos.row < 0 || newPos.row >= len(grid) || newPos.col < 0 || newPos.col >= len(grid[0]) || grid[newPos.row][newPos.col] == "#" {
				continue
			}

			newCost := node.cost + 1
			if pos.dir != d.dir {
				newCost += 1000
			}

			if c, ok := memo[newPos]; !ok || newCost < c {
				memo[newPos] = newCost
				heap.Push(&pq, &Node{pos: newPos, score: newCost + heuristic(newPos, end), cost: newCost})
			}
		}
	}

	return -1
}

func aStarWithPath(grid [][]string, start, end Position, optimal int) []Position {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{pos: start, score: 0, cost: 0, path: []Position{start}})

	memo := make(map[Position]int)
	memo[start] = 0

	fmt.Println(start)
	fmt.Println(end)

	directions := []Position{
		{row: -1, col: 0, dir: 0}, // up
		{row: 0, col: 1, dir: 1},  // right
		{row: 1, col: 0, dir: 2},  // down
		{row: 0, col: -1, dir: 3}, // left
	}

	var allPaths [][]Position
	lowestCost := -1

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		pos := node.pos

		if pos.row == end.row && pos.col == end.col {
			if lowestCost == -1 || node.cost == lowestCost {
				lowestCost = node.cost
				allPaths = append(allPaths, node.path)
			} else if node.cost < lowestCost {
				lowestCost = node.cost
				allPaths = [][]Position{node.path}
			}
			continue
		}

		for _, d := range directions {
			newPos := Position{row: pos.row + d.row, col: pos.col + d.col, dir: d.dir}
			if newPos.row < 0 || newPos.row >= len(grid) || newPos.col < 0 || newPos.col >= len(grid[0]) || grid[newPos.row][newPos.col] == "#" {
				continue
			}

			newCost := node.cost + 1
			if pos.dir != d.dir {
				newCost += 1000
			}

			if newCost > optimal {
				continue
			}

			if c, ok := memo[newPos]; !ok || newCost <= c {
				memo[newPos] = newCost
				newPath := append([]Position{}, node.path...)
				newPath = append(newPath, newPos)
				heap.Push(&pq, &Node{pos: newPos, score: newCost + heuristic(newPos, end), cost: newCost, path: newPath})
			}
		}
	}

	visitedPositions := make(map[Position]bool)
	for _, path := range allPaths {
		for _, pos := range path {
			visitedPositions[pos] = true
		}
	}

	var result []Position
	for pos := range visitedPositions {
		result = append(result, pos)
	}

	fmt.Println(lowestCost)
	fmt.Println(len(result))

	return result
}

func heuristic(pos, end Position) int {
	return abs(pos.row-end.row) + abs(pos.col-end.col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findS(grid [][]string) (int, int) {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func findE(grid [][]string) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "E" {
				return i, j
			}
		}
	}
	return -1, -1
}

func buildData(lines []string) [][]string {
	var grid [][]string

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// Part 2: 517 is too low

// Original solution for Part 1

// package main

// import (
// 	"fmt"
// 	"log"
// 	"strings"

// 	"2024/common"
// )

// type Position struct {
// 	row int
// 	col int
// 	dir int // 0 = up, 1 = right, 2 = down, 3 = left
// }

// func main() {
// 	filePath := "input.txt"
// 	data, err := readAndProcessFile(filePath)
// 	if err != nil {
// 		log.Fatalf("failed to process file: %v", err)
// 	}

// 	// Example 1 optimal score: 7036
// 	// Example 2 optimal score: 11048

// 	grid := buildData(data)
// 	result1 := solve1(grid)
// 	fmt.Println(result1)
// }

// func solve1(grid [][]string) int {
// 	memo := make(map[Position]int)
// 	rowS, colS := findS(grid)
// 	startingPos := Position{row: rowS, col: colS, dir: 1}
// 	min := 1000000000
// 	move(grid, startingPos, 0, memo, &min)
// 	return min
// }

// func move(grid [][]string, pos Position, score int, memo map[Position]int, min *int) {
// 	if grid[pos.row][pos.col] == "E" {
// 		if score < *min {
// 			*min = score
// 		}
// 		return
// 	}

// 	// up
// 	if pos.row-1 >= 0 && grid[pos.row-1][pos.col] != "#" && pos.dir != 2 {
// 		newPos := Position{row: pos.row - 1, col: pos.col, dir: 0}
// 		if _, ok := memo[newPos]; !ok || score < memo[newPos] {
// 			memo[newPos] = score
// 			if pos.dir == 0 {
// 				move(grid, newPos, score+1, memo, min)
// 			} else {
// 				move(grid, newPos, score+1001, memo, min)
// 			}
// 		}
// 	}

// 	// right
// 	if pos.col+1 < len(grid[pos.row]) && grid[pos.row][pos.col+1] != "#" && pos.dir != 3 {
// 		newPos := Position{row: pos.row, col: pos.col + 1, dir: 1}
// 		if _, ok := memo[newPos]; !ok || score < memo[newPos] {
// 			memo[newPos] = score
// 			if pos.dir == 1 {
// 				move(grid, newPos, score+1, memo, min)
// 			} else {
// 				move(grid, newPos, score+1001, memo, min)
// 			}
// 		}
// 	}

// 	// down
// 	if pos.row+1 < len(grid) && grid[pos.row+1][pos.col] != "#" && pos.dir != 0 {
// 		newPos := Position{row: pos.row + 1, col: pos.col, dir: 2}
// 		if _, ok := memo[newPos]; !ok || score < memo[newPos] {
// 			memo[newPos] = score
// 			if pos.dir == 2 {
// 				move(grid, newPos, score+1, memo, min)
// 			} else {
// 				move(grid, newPos, score+1001, memo, min)
// 			}
// 		}
// 	}

// 	// left
// 	if pos.col-1 >= 0 && grid[pos.row][pos.col-1] != "#" && pos.dir != 1 {
// 		newPos := Position{row: pos.row, col: pos.col - 1, dir: 3}
// 		if _, ok := memo[newPos]; !ok || score < memo[newPos] {
// 			memo[newPos] = score
// 			if pos.dir == 3 {
// 				move(grid, newPos, score+1, memo, min)
// 			} else {
// 				move(grid, newPos, score+1001, memo, min)
// 			}
// 		}
// 	}

// }

// func findS(grid [][]string) (int, int) {
// 	for i := len(grid) - 1; i >= 0; i-- {
// 		for j := 0; j < len(grid[i]); j++ {
// 			if grid[i][j] == "S" {
// 				return i, j
// 			}
// 		}
// 	}
// 	return -1, -1
// }

// func buildData(lines []string) [][]string {
// 	var grid [][]string

// 	for _, line := range lines {
// 		grid = append(grid, strings.Split(line, ""))
// 	}

// 	return grid
// }

// func readAndProcessFile(filePath string) ([]string, error) {
// 	lines, err := common.ReadFileLines(filePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return lines, nil
// }
