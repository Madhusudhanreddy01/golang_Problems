package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumOfDigits(n int) int {
	// If the number is less than 10, return it
	if n < 10 {
		return n
	}

	// Split the number into its digits and add them together using recursion
	return sumOfDigits(n/10) + n%10
}

func main() {
	// Prompt the user for a number
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert the input string to an integer
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	// Find the sum of the digits of the number
	sum := sumOfDigits(n)

	// Print the result
	fmt.Println("The sum of the digits of", n, "is", sum)
}
