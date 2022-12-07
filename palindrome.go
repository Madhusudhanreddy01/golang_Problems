package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		palindrome(i)
	}
}

func palindrome(num int) {
	temp := num

	res := 0
	for num > 0 {
		rem := num % 10
		res = res*10 + rem
		num = num / 10

	}
	if temp == res {
		fmt.Println(temp)
	}
}

//Another way

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	fmt.Println(palindrome(1, 100))

// }
// func palindrome(a, b int) []string {
// 	slice := []string{}
// 	for i := a; i < b+1; i++ {
// 		I := strconv.Itoa(i)
// 		j := []rune(I)
// 		var result []rune
// 		for i := len(j) - 1; i >= 0; i-- {
// 			result = append(result, j[i])
// 		}
// 		if I == string(result) {
// 			slice = append(slice, string(result))
// 		}

// 	}
// 	return slice
// }
