package main

import (
	"fmt"
)

func PrimeNumbers(x int) {
	if x != 100 {
		count := 0
		for i := 1; 1 < x+1; i++ {
			if x%i == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(x, " ")
		}
		PrimeNumbers(x + 1)
	}
}
