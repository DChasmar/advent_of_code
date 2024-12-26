package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"2024/common"
)

func main() {
	filePath := "input.txt"
	data, err := readAndProcessFile(filePath)
	if err != nil {
		log.Fatalf("failed to process file: %v", err)
	}

	values, steps, zSet := processLines(data)

	// Check for some common errors
	checks(steps)

	var initialXKeys []string
	var initialYKeys []string
	for k := range values {
		if strings.HasPrefix(k, "x") {
			initialXKeys = append(initialXKeys, k)
		} else if strings.HasPrefix(k, "y") {
			initialYKeys = append(initialYKeys, k)
		}
	}

	sort.Strings(initialXKeys)
	sort.Strings(initialYKeys)

	xBytes := byteSlices(initialXKeys, values)
	yBytes := byteSlices(initialYKeys, values)
	initZBytes := sumBytesSlices(xBytes, yBytes)

	fmt.Println("init Z", initZBytes)

	xNum := bytesToInt(xBytes)
	yNum := bytesToInt(yBytes)
	initZ := xNum + yNum

	fmt.Println(xBytes)
	fmt.Println(yBytes)

	fmt.Println("Initial Z:", initZ)

	solve1(values, steps, zSet)
	fmt.Println("Done")
	// Create a sorted list of zSet keys
	var zKeys []string
	for k := range zSet {
		zKeys = append(zKeys, k)
	}
	sort.Strings(zKeys)

	// fmt.Println("Z Keys:", zKeys)

	// zBytes := byteSlices(zKeys, values)

	// num := bytesToInt(zBytes)
	// fmt.Println("Z Bytes:", zBytes)
	// fmt.Println("Number:", num)
}

func checks(steps [][]string) {
	for _, step := range steps {
		if step[1] == "AND" {
			ok := findOp(steps, step[3], "OR")
			if !ok {
				fmt.Println("No OR found for", step)
			}
		}
		// if step[1] == "XOR" && !(strings.HasPrefix(step[0], "x") || strings.HasPrefix(step[2], "x") || strings.HasPrefix(step[3], "z")) {
		// 	fmt.Println("XOR found for", step)
		// }
		// if strings.HasPrefix(step[3], "z") && step[1] != "XOR" {
		// 	fmt.Println("Z found for", step)
		// }
		if step[1] == "XOR" && (strings.HasPrefix(step[0], "x") || strings.HasPrefix(step[2], "x")) {
			ok := findOp(steps, step[3], "XOR")
			if !ok {
				fmt.Println("No XOR found for", step)
			}
		}
	}
}

// z27, z18, z22, hbs, dhq, jcp, pdg, kfp
// sorted: dhq,hbs,jcp,kfp,pdg,z18,z22,z27

func findOp(steps [][]string, val, op string) bool {
	for _, step := range steps {
		if step[1] == op && (step[0] == val || step[2] == val) {
			return true
		}
	}
	// fmt.Println("No", op, "found for", val)
	return false
}

func sumBytesSlices(b1 []byte, b2 []byte) []byte {
	maxLength := max(len(b1), len(b2))
	b3 := make([]byte, maxLength)
	carry := byte(0)

	for i := 0; i < maxLength; i++ {
		var sum int // Use int to avoid overflow during addition
		if i < len(b1) {
			sum += int(b1[i])
		}
		if i < len(b2) {
			sum += int(b2[i])
		}
		sum += int(carry)
		if sum > 255 {
			carry = 1
			sum -= 256
		} else {
			carry = 0
		}
		b3[i] = byte(sum)                                                                                             // Convert back to byte
		fmt.Printf("Index %d: b1 = %d, b2 = %d, sum = %d, carry = %d, b3 = %d\n", i, b1[i], b2[i], sum, carry, b3[i]) // Debugging statement
	}

	if carry > 0 {
		b3 = append(b3, carry)
	}

	return b3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bytesToInt(zBytes []byte) int {
	num := 0
	for i, b := range zBytes {
		num += int(b) << (8 * i)
	}
	return num
}

func byteSlices(keys []string, values map[string]*byte) []byte {
	// Initialize a slice of bytes to store the bits
	zBytes := make([]byte, (len(keys)+7)/8) // +7 to round up to the nearest byte

	// Set the bits in the slice of bytes
	for _, k := range keys {
		if v, exists := values[k]; exists && v != nil {
			if *v == 1 {
				// Extract the numeric part of the key
				numPart := k[1:]
				bitPos, err := strconv.Atoi(numPart)
				if err != nil {
					log.Fatalf("invalid key format: %s", k)
				}
				byteIndex := bitPos / 8
				bitIndex := bitPos % 8
				zBytes[byteIndex] |= (1 << bitIndex)
				// fmt.Printf("Setting bit %d (byte %d, bit %d) for key %s\n", bitPos, byteIndex, bitIndex, k) // Debugging statement
			}
		}
	}
	return zBytes
}

func solve1(values map[string]*byte, steps [][]string, zSet map[string]bool) {
	newZSet := make(map[string]bool)
	for len(newZSet) < len(zSet) {
		for _, step := range steps {
			// Check for nil pointers
			if values[step[0]] == nil || values[step[2]] == nil {
				continue
			}
			// Process the step
			left := *values[step[0]]
			right := *values[step[2]]
			var result byte
			switch step[1] {
			case "AND":
				result = left & right
			case "OR":
				result = left | right
			case "XOR":
				result = left ^ right
			}

			// Assign the result to the destination
			values[step[3]] = &result

			// Track unique strings that begin with 'z'
			if strings.HasPrefix(step[3], "z") {
				newZSet[step[3]] = true
			}
		}
	}
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func processLines(lines []string) (map[string]*byte, [][]string, map[string]bool) {
	values := make(map[string]*byte)
	var steps [][]string
	zSet := make(map[string]bool)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, ":") {
			// Process lines like "x00: 1"
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				log.Fatalf("invalid line format: %s", line)
			}
			key := strings.TrimSpace(parts[0])
			valueStr := strings.TrimSpace(parts[1])
			var value byte
			if valueStr == "1" {
				value = 1
			} else if valueStr == "0" {
				value = 0
			} else {
				log.Fatalf("invalid value: %s", valueStr)
			}
			values[key] = &value
		} else if strings.Contains(line, "->") {
			// Process lines like "vdt OR tnw -> bfw"
			parts := strings.Split(line, "->")
			if len(parts) != 2 {
				log.Fatalf("invalid line format: %s", line)
			}
			left := strings.TrimSpace(parts[0])
			right := strings.TrimSpace(parts[1])
			leftParts := strings.Fields(left)
			if len(leftParts) != 3 {
				log.Fatalf("invalid left part format: %s", left)
			}
			step := append(leftParts, right)
			steps = append(steps, step)

			// Assign nil to index 0, 2, 3 values in step if they do not already exist in values
			for _, idx := range []int{0, 2, 3} {
				if _, exists := values[step[idx]]; !exists {
					values[step[idx]] = nil
				}
			}

			// Track unique strings that begin with 'z'
			for _, idx := range []int{0, 2, 3} {
				if strings.HasPrefix(step[idx], "z") {
					zSet[step[idx]] = true
				}
			}
		} else {
			log.Fatalf("unrecognized line format: %s", line)
		}
	}

	return values, steps, zSet
}
