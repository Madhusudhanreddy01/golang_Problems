package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// Edge case: all digits are 9
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}

func main() {
	var digits []int
	fmt.Print("Enter the digits: ")
	fmt.Scan(&digits)
	fmt.Println(plusOne(digits))
}
