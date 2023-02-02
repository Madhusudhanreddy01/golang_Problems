package main

import (
	"fmt"
	"math"
)

// func IsPrime(value int) bool {
// 	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
// 		if value%i == 0 {
// 			return false
// 		}
// 	}
// 	return value > 1
// }

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	for i := 1; i <= 100; i++ {
		if IsPrimeSqrt(i) {
			fmt.Printf("%v ", i)
		}
	}
}
