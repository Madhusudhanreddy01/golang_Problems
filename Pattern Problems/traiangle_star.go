package main

import "fmt"

func main() {
	var rows int
	fmt.Println("Enter the value: ")
	fmt.Scanln(&rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
