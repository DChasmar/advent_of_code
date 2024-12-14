package common

import (
	"fmt"
)

func AbsDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Function to calculate the least common multiple (LCM)
func FirstOccurence(m1, b1, m2, b2 int) int {
	test1 := b1
	test2 := b2
	counter := 0
	for test1 != test2 {
		if test1 < test2 {
			test1 += m1
		} else {
			test2 += m2
		}
		counter++
		if counter > 1000 {
			fmt.Println("fail LCM", test1, m1, b1, m2, b2)
			return -1
		}
	}
	return test1
}
