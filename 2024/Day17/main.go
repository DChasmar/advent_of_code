package main

import (
	"fmt"
	"log"
	"math"
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

	registers, ops := buildData(data)
	result := solve1(&registers, ops)
	s := joinIntArrWithCommas(result)
	fmt.Println(s)
	// Trial and error to find the correct value for A... needs to be fixed
	result2Alt := solve2Alt(ops)
	fmt.Println(result2Alt)
	// result2 := solve2(ops)
	// fmt.Println(result2)
}

func solve2Alt(ops []int) int {
	a := 267265166221856
	for i := 0; i < 1000; i++ {
		registers := Registers{A: a + i, B: 0, C: 0}
		output := solve1(&registers, ops)
		if endOfArrsMatch(output, ops, 16) {
			fmt.Println("output: ", output, "a", a+i)
			return a + i
		}

	}
	return -1
}

func solve2(ops []int) int {
	a := int(math.Pow(8, 15))
	power := 14

	equal := false
	allEqual := false

	count := 0

	for !allEqual {
		first := true
		for !equal {
			registers := Registers{A: a, B: 0, C: 0}
			output := solve1(&registers, ops)

			if first {
				first = false
				fmt.Printf("initA: %d, power: %d, output: %v\n", a, power, output)
			}

			if len(output) != 16 {
				fmt.Println("NOT 16", "output: ", output, "a", a)
			}
			if endOfArrsMatch(output, ops, 16-power) {
				fmt.Println("output: ", output, "a", a)
				power = max(0, power-1)
				equal = true
				if power == 0 {
					fmt.Println("output: ", output)
					allEqual = true
				}
			} else {
				if power == 7 {
					fmt.Println("output: ", output, "a", a)
					count++
				}
				a += int(math.Pow(8, float64(power)))
				if count > 200 {
					log.Fatalln("count > 200")
				}
			}
		}
		equal = false
	}
	return a
}

func endOfArrsMatch(arr1, arr2 []int, matches int) bool {
	if len(arr1) < matches || len(arr2) < matches {
		return false
	}
	for i := 0; i < matches; i++ {
		if arr1[len(arr1)-1-i] != arr2[len(arr2)-1-i] {
			return false
		}
	}
	return true
}

func solve1(registers *Registers, ops []int) []int {
	outs := []int{}
	i := 0
	for i < len(ops) {
		opcode := ops[i]
		operand := ops[i+1]
		switch opcode {
		case 0:
			adv(operand, registers)
		case 1:
			bxl(operand, registers)
		case 2:
			bst(operand, registers)
		case 3:
			if registers.A == 0 {
				i += 2
			} else if i == operand {
				fmt.Println("i == operand")
				i += 2
			} else {
				i = operand
			}
		case 4:
			bxc(registers)
		case 5:
			outs = append(outs, out(operand, registers))
		case 6:
			bdv(operand, registers)
		case 7:
			cdv(operand, registers)
		}
		if opcode != 3 {
			i += 2
		}
	}

	return outs
}

func adv(operand int, registers *Registers) {
	combo := combo(operand, *registers)
	val := registers.A / int(math.Pow(2, float64(combo)))
	registers.A = val
}

func bxl(operand int, registers *Registers) {
	bitwiseXOR := bitwiseXOR(registers.B, operand)
	registers.B = bitwiseXOR
}

func bst(operand int, registers *Registers) {
	combo := combo(operand, *registers)
	val := combo % 8
	registers.B = val
}

func bxc(registers *Registers) {
	bitwiseXOR := bitwiseXOR(registers.B, registers.C)
	registers.B = bitwiseXOR
}

func out(operand int, registers *Registers) int {
	count := 0
	combo := combo(operand, *registers)
	val := combo % 8
	count++
	if count == 8 {
		log.Fatalln("count == 8")
	}
	return val
}

func bdv(operand int, registers *Registers) {
	combo := combo(operand, *registers)
	val := registers.A / int(math.Pow(2, float64(combo)))
	registers.B = val
}

func cdv(operand int, registers *Registers) {
	combo := combo(operand, *registers)
	val := registers.A / int(math.Pow(2, float64(combo)))
	registers.C = val
}

func bitwiseXOR(num1, num2 int) int {
	return num1 ^ num2
}

func combo(opcode int, registers Registers) int {
	switch opcode {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return registers.A
	case 5:
		return registers.B
	case 6:
		return registers.C
	case 7:
		fmt.Println(opcode)
		fmt.Println(registers)
		log.Fatalln("invalid opcode")
		return -1
	}
	log.Fatalln("opcode not 0-7")
	return -1
}

func joinIntArrWithCommas(arr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
}

func buildData(lines []string) (Registers, []int) {
	registers := Registers{}
	ops := []int{}

	for _, line := range lines {
		if strings.HasPrefix(line, "Register A:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					registers.A = value
				}
			}
		} else if strings.HasPrefix(line, "Register B:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					registers.B = value
				}
			}
		} else if strings.HasPrefix(line, "Register C:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					registers.C = value
				}
			}
		} else if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				instructions := strings.Split(strings.TrimSpace(parts[1]), ",")
				for _, instr := range instructions {
					value, err := strconv.Atoi(strings.TrimSpace(instr))
					if err == nil {
						ops = append(ops, value)
					}
				}
			}
		}
	}

	return registers, ops
}

func readAndProcessFile(filePath string) ([]string, error) {
	lines, err := common.ReadFileLines(filePath)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
