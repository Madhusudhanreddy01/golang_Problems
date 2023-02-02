package main

import "fmt"

func main() {
	var rows int
	fmt.Println("Enter the value: ")
	fmt.Scanln(&rows)

	for i := 1; i <= rows; i++ {

		for j := 1; j <= i; j++ {
			// fmt.Printf("%c ", 'A'+j-1)
			fmt.Printf("%d ", j)
		}
		fmt.Println()

	}

}
