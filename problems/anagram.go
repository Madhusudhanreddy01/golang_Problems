package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	// Check that the strings are the same length
	if len(s1) != len(s2) {
		return false
	}

	// Create a map to hold the frequencies of the characters in the first string
	freq1 := make(map[rune]int)
	for _, c := range s1 {
		freq1[c]++
	}

	// Check that the characters in the second string have the same frequencies as in the first string
	for _, c := range s2 {
		if freq1[c] == 0 {
			return false
		}
		freq1[c]--
	}

	// If all the characters in the second string have the same frequencies as in the first string, the strings are anagrams
	return true
}

func main() {
	// Prompt the user for the first string
	fmt.Print("Enter the first string: ")
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	// Prompt the user for the second string
	fmt.Print("Enter the second string: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	// Check if the strings are anagrams
	if isAnagram(input1, input2) {
		fmt.Println("The strings are anagrams.")
	} else {
		fmt.Println("The strings are not anagrams.")
	}
}

// ----------------------------------------------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func isAnagram(s1, s2 string) bool {
// 	// Check that the strings are the same length
// 	if len(s1) != len(s2) {
// 		return false
// 	}

// 	// Sort the strings
// 	s1Sorted := sort.Sort([]rune(s1))
// 	s2Sorted := sort.Sort([]rune(s2))

// 	// Check that the sorted strings are equal
// 	if s1Sorted == s2Sorted {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	// Prompt the user for the first string
// 	fmt.Print("Enter the first string: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input1, _ := reader.ReadString('\n')
// 	input1 = strings.TrimSpace(input1)

// 	// Prompt the user for the second string
// 	fmt.Print("Enter the second string: ")
// 	input2, _ := reader.ReadString('\n')
// 	input2 = strings.TrimSpace(input2)

// 	// Check if the strings are anagrams
// 	if isAnagram(input1, input2) {
// 		fmt.Println("The strings are anagrams.")
// 	} else {
// 		fmt.Println("The strings are not anagrams.")
// 	}
// }
