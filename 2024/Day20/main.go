package main

import (
	"container/heap"
	"fmt"
	"log"

	"2024/common"
)

type Position struct {
	row int
	col int
}

type Node struct {
	pos   Position
	score int
	cost  int
	index int
	path  []Position
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

	grid := common.LinesToGrid(data)

	var LENGTH int
	if len(grid) == len(grid[0]) {
		LENGTH = len(grid)
	} else {
		log.Fatalf("grid is not square")
	}
	// Initialize the start and end positions
	start, end, dotCount, hashCount := findLettersAndCount(grid, "S", "E", LENGTH)
	if start.row == -1 || end.row == -1 {
		log.Fatalf("start or end not found")
	}

	fmt.Println(start, end)
	fmt.Println(dotCount, hashCount)

	initShortest, path := aStarWithPath(grid, start, end)
	fmt.Println(initShortest)
	fmt.Println(path[:10])
	fmt.Println(path[len(path)-10:])
	fmt.Println(len(path))

	// Create a set of the path positions
	pathSet := make(map[Position]int)
	for index, p := range path {
		pathSet[p] = index
	}

	// This solution is very slow; how could it be optimized?
	// The shortest (and only) path passes through every dot in the grid***
	// result1 := solve1(grid, start, end, initShortest, LENGTH, pathSet)
	// fmt.Println("")
	// fmt.Println(result1)

	// pathGrid := make([][]string, LENGTH)
	// for i := 0; i < LENGTH; i++ {
	// 	pathGrid[i] = make([]string, LENGTH)
	// 	for j := 0; j < LENGTH; j++ {
	// 		if pathSet[Position{i, j}] > 0 && grid[i][j] == "." {
	// 			pathGrid[i][j] = "O"
	// 		} else {
	// 			pathGrid[i][j] = grid[i][j]
	// 		}
	// 	}
	// }
	// common.GridToTxt(pathGrid, filePath[:len(filePath)-4]+"_output.txt")

	// fmt.Println(pathSet)
	result2 := solve2(grid, path, pathSet)
	fmt.Println(result2)

	// Part 2: 186961 is too low
}

func solve2(grid [][]string, path []Position, pathSet map[Position]int) int {
	count := 0
	for i := 0; i < len(path); i++ {
		count += findGridPositionsWithin20(path[i], pathSet, len(grid))
	}
	return count
}

func findGridPositionsWithin20(start Position, pathSet map[Position]int, gridLength int) int {
	leap := 20
	count := 0
	for i := -leap; i <= leap; i++ {
		for j := -leap + abs(i); j <= leap-abs(i); j++ {
			row, col := start.row+i, start.col+j
			if row < 1 || row >= gridLength-1 || col < 1 || col >= gridLength-1 {
				continue
			}
			pos := Position{row, col}
			if val, ok := pathSet[pos]; ok {
				diff := abs(i) + abs(j)
				skipped := val - pathSet[start] - diff
				if skipped > 99 {
					count++
				}
			}
		}
	}
	return count
}

func solve1(grid [][]string, start, end Position, shortest int, LENGTH int) int {
	count := 0
	for i := 1; i < LENGTH-1; i++ {
		fmt.Print(i, " ")
		for j := 1; j < LENGTH-1; j++ {
			if surrounded(grid, Position{i, j}) {
				continue
			}
			if grid[i][j] == "#" {
				grid[i][j] = "."
				newShortest := aStar(grid, start, end)
				if newShortest+99 < shortest {
					count++
				}
				grid[i][j] = "#"
			}
		}
	}

	return count
}

func aStar(grid [][]string, start, end Position) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{pos: start, score: 0, cost: 0})

	memo := make(map[Position]int)
	memo[start] = 0

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		pos := node.pos

		if pos.row == end.row && pos.col == end.col {
			return node.cost
		}

		for _, d := range directions {
			newPos := Position{row: pos.row + d[0], col: pos.col + d[1]}
			if newPos.row < 0 || newPos.row >= len(grid) || newPos.col < 0 || newPos.col >= len(grid[0]) || grid[newPos.row][newPos.col] == "#" {
				continue
			}

			newCost := node.cost + 1

			if c, ok := memo[newPos]; !ok || newCost < c {
				memo[newPos] = newCost
				heap.Push(&pq, &Node{pos: newPos, score: newCost + heuristic(newPos, end), cost: newCost})
			}
		}
	}

	log.Fatalf("no path found")
	return -1
}

func aStarWithPath(grid [][]string, start, end Position) (int, []Position) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{pos: start, score: 0, cost: 0, path: []Position{start}})

	memo := make(map[Position]int)
	memo[start] = 0

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		pos := node.pos

		if pos.row == end.row && pos.col == end.col {
			return node.cost, node.path
		}

		for _, d := range directions {
			newPos := Position{row: pos.row + d[0], col: pos.col + d[1]}
			if newPos.row < 0 || newPos.row >= len(grid) || newPos.col < 0 || newPos.col >= len(grid[0]) || grid[newPos.row][newPos.col] == "#" {
				continue
			}

			newCost := node.cost + 1

			if c, ok := memo[newPos]; !ok || newCost < c {
				memo[newPos] = newCost
				newPath := append([]Position{}, node.path...)
				newPath = append(newPath, newPos)
				heap.Push(&pq, &Node{pos: newPos, score: newCost + heuristic(newPos, end), cost: newCost, path: newPath})
			}
		}
	}

	log.Fatalf("no path found")
	return -1, nil
}

func surrounded(grid [][]string, pos Position) bool {
	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for _, d := range directions {
		newPos := Position{row: pos.row + d[0], col: pos.col + d[1]}
		if grid[newPos.row][newPos.col] != "#" {
			return false
		}
	}
	return true
}

func heuristic(a, b Position) int {
	return abs(a.row-b.row) + abs(a.col-b.col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findLettersAndCount(grid [][]string, start, end string, LENGTH int) (Position, Position, int, int) {
	var s, e Position
	s = Position{row: -1, col: -1}
	e = Position{row: -1, col: -1}
	dotCount := 0
	hashCount := 0
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < LENGTH; j++ {
			if grid[i][j] == start {
				s = Position{i, j}
			} else if grid[i][j] == end {
				e = Position{i, j}
			} else if grid[i][j] == "." {
				dotCount++
			} else if grid[i][j] == "#" {
				hashCount++
			} else {
				log.Fatalf("invalid character in grid")
			}
		}
	}
	return s, e, dotCount, hashCount
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
