package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		fmt.Printf("%c", input[i]+1)
	}
	fmt.Println()

	for _, char := range input {
		fmt.Printf("%c", char+1)
	}
	fmt.Println()
}

// ----------------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	input := "hello"
// 	l := ""
// 	for _, i := range input {
// 		l += string(i + 1)

// 	}
// 	fmt.Println(l)
// }
