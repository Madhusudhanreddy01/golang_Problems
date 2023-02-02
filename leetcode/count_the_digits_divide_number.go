package main

import (
	"fmt"
	"strconv"
)

func countDigits(n int) int {
	var count int
	num := strconv.Itoa(n)
	for i := 0; i < len(num); i++ {
		digit, _ := strconv.Atoi(string(num[i]))
		if digit != 0 && n%digit == 0 {
			count++
		}
	}
	return count
}

func main() {
	// get user input
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Println("Number of digits that divide the number:", countDigits(n))
}
