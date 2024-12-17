package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"2024/common"
)

type Coordinates struct {
	p [2]int
	v [2]int
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

	newObjects := buildData(data)
	result2 := solve2(newObjects)
	fmt.Println(result2)
}

func solve2(objects []Coordinates) int {
	WIDTH := 101
	HEIGHT := 103
	minDistance := 10000000
	distanceTotal := 0
	for i := 0; i < 200000; i++ {
		set := make(map[int]bool)
		for object := range objects {
			objects[object].p[0] += objects[object].v[0]
			objects[object].p[1] += objects[object].v[1]
			newX, newY := getInRange(objects[object].p[0], objects[object].p[1], WIDTH, HEIGHT)
			distanceTotal += euclideanDistance(newX, newY, (WIDTH-1)/2, (HEIGHT-1)/2)
			val := newX*1000 + newY
			set[val] = true
		}
		if distanceTotal < minDistance {
			fmt.Println("Distance:", distanceTotal)
			minDistance = distanceTotal
			if i > 1000 {
				fileName := fmt.Sprintf("grid_%d.txt", i)
				displayGrid(set, WIDTH, HEIGHT, fileName)
			}
		}
		distanceTotal = 0
	}
	return -1
}

func euclideanDistance(x1, y1, x2, y2 int) int {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}

// func average(arr []int) int {
// 	sum := 0
// 	for _, val := range arr {
// 		sum += val
// 	}
// 	return sum / len(arr)
// }

func displayGrid(set map[int]bool, WIDTH, HEIGHT int, filename string) {
	grid := make([][]string, HEIGHT)
	for y := 0; y < HEIGHT; y++ {
		grid[y] = make([]string, WIDTH)
		for x := 0; x < WIDTH; x++ {
			grid[y][x] = "."
		}
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if set[x*1000+y] {
				grid[y][x] = "#"
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			file.WriteString(grid[y][x])
		}
		file.WriteString("\n")
	}
	fmt.Println("Grid exported to", filename)
}

func solve1(objects []Coordinates) int {
	SECONDS := 100
	WIDTH := 101
	HEIGHT := 103
	quad1 := make(map[int]int)
	quad2 := make(map[int]int)
	quad3 := make(map[int]int)
	quad4 := make(map[int]int)
	for object := range objects {
		newX := objects[object].p[0] + objects[object].v[0]*SECONDS
		newY := objects[object].p[1] + objects[object].v[1]*SECONDS
		newX, newY = getInRange(newX, newY, WIDTH, HEIGHT)
		quad := findQuad(newX, newY, WIDTH, HEIGHT)
		val := newX*1000 + newY
		switch quad {
		case 1:
			quad1[val]++
		case 2:
			quad2[val]++
		case 3:
			quad3[val]++
		case 4:
			quad4[val]++
		}
	}
	product := 1
	for _, quad := range []map[int]int{quad1, quad2, quad3, quad4} {
		product *= getMapSum(quad)
	}
	return product
}

func getMapSum(quad map[int]int) int {
	sum := 0
	for _, val := range quad {
		sum += val
	}
	return sum
}

func findQuad(x, y, WIDTH, HEIGHT int) int {
	xDiv := (WIDTH - 1) / 2
	yDiv := (HEIGHT - 1) / 2
	if x >= 0 && x < xDiv && y >= 0 && y < yDiv {
		return 1
	} else if x > xDiv && x < WIDTH && y >= 0 && y < yDiv {
		return 2
	} else if x > xDiv && x < WIDTH && y > yDiv && y < HEIGHT {
		return 3
	} else if x >= 0 && x < xDiv && y > yDiv && y < HEIGHT {
		return 4
	}
	return 0
}

func getInRange(x, y, WIDTH, HEIGHT int) (int, int) {
	if x >= 0 {
		x = x % WIDTH
	} else {
		x = WIDTH + x%WIDTH
		if x == WIDTH {
			x = 0
		}
	}
	if y >= 0 {
		y = y % HEIGHT
	} else {
		y = HEIGHT + y%HEIGHT
		if y == HEIGHT {
			y = 0
		}
	}
	return x, y
}

func buildData(lines []string) []Coordinates {
	var objects []Coordinates

	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if matches != nil {
			var current Coordinates
			current.p[0], _ = strconv.Atoi(matches[1])
			current.p[1], _ = strconv.Atoi(matches[2])
			current.v[0], _ = strconv.Atoi(matches[3])
			current.v[1], _ = strconv.Atoi(matches[4])
			objects = append(objects, current)
		}
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
