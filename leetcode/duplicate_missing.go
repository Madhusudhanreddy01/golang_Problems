package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDuplicateAndMissing(nums []int) (int, int) {
	// Create a map to hold the frequencies of the numbers
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// Find the duplicate and missing numbers
	var duplicate, missing int
	for i := 1; i <= len(nums); i++ {
		if freq[i] == 0 {
			missing = i
		} else if freq[i] == 2 {
			duplicate = i
		}
	}

	return duplicate, missing
}

func main() {
	// Prompt the user for a list of numbers
	fmt.Print("Enter a list of numbers separated by spaces: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Split the input string into a slice of integers
	var nums []int
	for _, s := range strings.Split(input, " ") {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		nums = append(nums, n)
	}

	// Find the duplicate and missing numbers
	duplicate, missing := findDuplicateAndMissing(nums)

	// Print the result
	fmt.Println("The duplicate number is", duplicate)
	fmt.Println("The missing number is", missing)
}
