package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isAscending(sentence string) bool {
	// extract numbers from sentence
	words := strings.Fields(sentence)
	var numbers []int
	for _, word := range words {
		if n, err := strconv.Atoi(word); err == nil {
			numbers = append(numbers, n)
		}
	}

	// check if numbers are in ascending order
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] >= numbers[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	if isAscending(sentence) {
		fmt.Println("Numbers are in ascending order.")
	} else {
		fmt.Println("Numbers are not in ascending order.")
	}
}
