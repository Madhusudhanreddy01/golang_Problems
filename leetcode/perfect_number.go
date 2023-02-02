// package main

// import "fmt"

// func main() {
// 	// Get the number to check
// 	var num int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&num)

// 	// Initialize the sum of divisors to 0
// 	sum := 0

// 	// Find the divisors and add them to the sum
// 	for i := 1; i <= num/2; i++ {
// 		if num%i == 0 {
// 			sum += i
// 		}
// 	}

// 	// Check if the sum of divisors is equal to the number
// 	if sum == num {
// 		fmt.Println(num, "is a perfect number")
// 	} else {
// 		fmt.Println(num, "is not a perfect number")
// 	}
// }

// -----------------------------------------------------------------------------------------

package main

import "fmt"

func main() {
	// Get the number to check
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check if the number is a perfect number
	if checkPerfectNumber(num) {
		fmt.Println(num, "is a perfect number")
	} else {
		fmt.Println(num, "is not a perfect number")
	}
}

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sum := 1
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}
