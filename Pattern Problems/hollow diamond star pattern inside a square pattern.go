package main

import "fmt"

func main() {

	var i, j, k, l, row int

	fmt.Print("Enter Hollow Diamond inside a Square row = ")
	fmt.Scanln(&row)

	fmt.Println("Hollow Diamond inside a Square Star Pattern")

	for i = 1; i <= row; i++ {
		for j = i; j <= row; j++ {
			fmt.Printf("*")
		}
		for k = 1; k <= 2*i-2; k++ {
			fmt.Printf(" ")
		}
		for l = i; l <= row; l++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i = 1; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("*")
		}
		for k = 2*i - 2; k < 2*row-2; k++ {
			fmt.Printf(" ")
		}
		for l = 1; l <= i; l++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
