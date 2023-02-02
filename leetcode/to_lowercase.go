package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 'a' - 'A'
		}
	}
	return string(b)
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	fmt.Println("Original string:", input)
	fmt.Println("Lowercase string:", toLowerCase(input))
}

// ------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Get input from user
// 	var input string
// 	fmt.Print("Enter a string: ")
// 	fmt.Scan(&input)

// 	// Convert input to lowercase
// 	lowercase := strings.ToLower(input)

// 	// Print lowercase string
// 	fmt.Println("Lowercase:", lowercase)
// }
