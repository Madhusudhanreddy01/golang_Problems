// with user input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// romanToInt converts a Roman numeral to an integer.
func romanToInt(s string) int {
	// Create a map that maps Roman numerals to integers.
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize a result variable to 0.
	result := 0

	// Iterate through the runes in the Roman numeral string.
	for i, r := range s {
		// If the current rune has a lower value than the next rune, subtract it from the result.
		if i < len(s)-1 && romanToIntMap[r] < romanToIntMap[rune(s[i+1])] {
			result -= romanToIntMap[r]
		} else {
			// Otherwise, add it to the result.
			result += romanToIntMap[r]
		}
	}

	// Return the result as the final integer value.
	return result
}

func main() {
	// Create a reader to read user input from the command line.
	reader := bufio.NewReader(os.Stdin)

	// Read the Roman numeral string from the user.
	fmt.Print("Enter a Roman numeral: ")
	romanNumeral, _ := reader.ReadString('\n')
	romanNumeral = strings.TrimSpace(romanNumeral)

	// Convert the Roman numeral to an integer and print the result.
	intValue := romanToInt(romanNumeral)
	fmt.Printf("The integer value of %s is %d.\n", romanNumeral, intValue)

	// -------------------------------------------------
	fmt.Println("--------------------------------------------------------")
	fmt.Println("The with-out user input RomantoInteger is:")
	fmt.Println(romanToInteger("MCMIV"))
}

// ----------------------------------------------------------------------------------------------

// with-out user input
// package main

// import (
// 	"fmt"
// )

// romanToInt converts a Roman numeral to an integer.
func romanToInteger(s string) int {
	// Create a map that maps Roman numerals to integers.
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize a result variable to 0.
	result := 0

	// Iterate through the runes in the Roman numeral string.
	for i, r := range s {
		// If the current rune has a lower value than the next rune, subtract it from the result.
		if i < len(s)-1 && romanToIntMap[r] < romanToIntMap[rune(s[i+1])] {
			result -= romanToIntMap[r]
		} else {
			// Otherwise, add it to the result.
			result += romanToIntMap[r]
		}
	}

	// Return the result as the final integer value.
	return result
}

// func main() {
// 	// Test the romanToInt function.
// 	fmt.Println(romanToInt("MCMIV")) // Output: 1904
// }
