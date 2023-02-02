package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	fmt.Printf("The length of the last word is %d\n", lengthOfLastWord(input))
}
