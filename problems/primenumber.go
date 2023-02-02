// package main

// import (
// 	"fmt"
// )

// func PrimeNumbers(x int) {
// 	if x != 100 {
// 		count := 0
// 		for i := 1; 1 < x+1; i++ {
// 			if x%i == 0 {
// 				count++
// 			}
// 		}
// 		if count == 2 {
// 			fmt.Print(x, " ")
// 		}
// 		PrimeNumbers(x + 1)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	// Iterate over the numbers from 1 to 100
// 	for i := 1; i <= 100; i++ {
// 		// Assume the number is prime
// 		isPrime := true

// 		// Check if the number is prime
// 		for j := 2; j < i; j++ {
// 			if i%j == 0 {
// 				// If the number is divisible by any number between 2 and itself, it is not prime
// 				isPrime = false
// 				break
// 			}
// 		}

// 		// If the number is prime, print it
// 		if isPrime {
// 			fmt.Println(i)
// 		}
// 	}
// }

package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 2; i < 20; i++ {
		if isPrime(i) {
			fmt.Println(i, "is prime")
		}
	}
}
