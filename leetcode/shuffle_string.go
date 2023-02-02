package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	n := len(s)
	result := make([]byte, n)
	for i, v := range indices {
		result[v] = s[i]
	}
	return string(result)
}

func main() {
	fmt.Println("Enter a string:")
	var input string
	fmt.Scan(&input)

	fmt.Println("Enter the indices array (separated by spaces):")
	var indices []int
	for {
		var index int
		_, err := fmt.Scan(&index)
		if err != nil {
			break
		}
		indices = append(indices, index)
	}

	fmt.Println("Shuffled string:", restoreString(input, indices))
}
