package main

import (
	"fmt"
	"log"

	"2024/common"
)

const (
	Horizontal = iota
	Vertical
)

type Position = [2]int

type Memo struct {
	s     string
	depth byte
}

var dp = map[Memo]int{}       // memoization
var numStart = Position{3, 2} // Numeric keypad start position
var dirStart = Position{0, 2} // Directional keypad start position

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	var buttons [][]byte
	for _, line := range data {
		buttons = append(buttons, common.LineToByteArray(line))
	}

	numeric := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"X", "0", "A"},
	}

	numPosMap := make(map[string]Position)
	for i, row := range numeric {
		for j, button := range row {
			numPosMap[button] = Position{i, j}
		}
	}

	directional := [][]string{
		{"X", "^", "A"},
		{"<", "v", ">"},
	}

	dirPosMap := make(map[string]Position)
	for i, row := range directional {
		for j, button := range row {
			dirPosMap[button] = Position{i, j}
		}
	}

	result1 := solve(buttons, 2)
	fmt.Println("Result 1: ", result1)

	result2 := solve(buttons, 25)
	fmt.Println("Result 2: ", result2)
}

func solve(codes [][]byte, bots byte) int {
	var res int
	var codeInt int
	for _, code := range codes {
		var path [][]byte
		codeInt = common.ByteArrayToInt(code[:len(code)-1])

		prev := numStart
		for _, c := range code {
			curr := ctopn(c)
			path = append(path, shortestSeq(prev, curr, true))
			prev = curr
		}

		var pathLen int
		for _, code := range path {
			pathLen += dfs(Memo{string(code), bots})
		}
		res += pathLen * codeInt
	}

	return res
}

// Char to Pos on numeric keypad
func ctopn(c byte) Position {
	if c == 'A' {
		return numStart
	}
	if c == '0' {
		return Position{3, 1}
	}
	row := 2 - ((c - '0' - 1) / 3)
	col := (c - '0' - 1) % 3
	return Position{int(row), int(col)}
}

// Char to Pos on directional keypad
func ctopd(d byte) Position {
	switch d {
	case '^':
		return Position{0, 1}
	case '<':
		return Position{1, 0}
	case 'v':
		return Position{1, 1}
	case '>':
		return Position{1, 2}
	default:
		return Position{0, 2}
	}
}

func pathWriter(off, dir int) []byte {
	var path []byte
	var c byte
	if dir == Horizontal {
		if off < 0 {
			c = '>'
		} else {
			c = '<'
		}
	} else {
		if off < 0 {
			c = 'v'
		} else {
			c = '^'
		}
	}
	for range common.Abs(off) {
		path = append(path, c)
	}
	return path
}

func shortestSeq(src, dst Position, isNumPad bool) []byte {
	var path []byte

	dr := src[0] - dst[0]
	dc := src[1] - dst[1]

	movesV := pathWriter(dr, Vertical)
	movesH := pathWriter(dc, Horizontal)

	var onGap bool
	if isNumPad {
		onGap = (src[0] == 3 && dst[1] == 0) || (src[1] == 0 && dst[0] == 3)
	} else {
		onGap = (src[1] == 0 && dst[0] == 0) || (src[0] == 0 && dst[1] == 0)
	}

	goingLeft := dst[1] < src[1]

	if goingLeft != onGap {
		movesV, movesH = movesH, movesV
	}

	path = append(append([]byte{}, movesV...), movesH...)
	path = append(path, 'A')
	return path
}

func dfs(memo Memo) int {
	if v, ok := dp[memo]; ok {
		return v
	}
	if memo.depth == 0 {
		return len(memo.s)
	}

	var total int
	var path [][]byte
	prev := dirStart
	for _, c := range memo.s {
		pos := ctopd(byte(c))
		path = append(path, shortestSeq(prev, pos, false))
		prev = pos
	}

	for i := 0; i < len(path); i++ {
		newMemo := Memo{string(path[i]), memo.depth - 1}
		total += dfs(newMemo)
	}
	dp[memo] = total
	return total
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// Part 1: 157942 is too high
// Part 1: 147942... just a test... is too low
// Part 1: 159082 is too high... of course
// Part 1: 152382 is incorrect
// Part 1: 152942 is correct (for some reason v> and ^> are more efficient than >v and >^)
