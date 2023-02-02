package main

import (
	"fmt"
)

func main() {
	var a, b, c = [5]int{}, [5]int{}, [5]int{}

	fmt.Print("\nEnter the first array:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Printf("\nEnter the second array:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&b[i])
	}

	for i := 0; i < 5; i++ {
		c[i] = a[i] + b[i]
		fmt.Printf("\nthird array element at index %d is: %d", i, c[i])

	}
}
