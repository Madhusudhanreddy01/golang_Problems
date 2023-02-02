package main

import (
	"fmt"
)

func numCommonFactors(a int, b int) int {
	smaller := min(a, b)
	count := 0
	for i := 1; i <= smaller; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var a, b int
	fmt.Println("Enter the two integers : ")
	fmt.Scan(&a, &b)
	result := numCommonFactors(a, b)
	fmt.Printf("The number of common factors between %d and %d is: %d\n", a, b, result)
}
