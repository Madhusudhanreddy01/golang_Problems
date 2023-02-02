package main

import "fmt"

func main() {
	var length, i, j int

	fmt.Print("Enter the Array Size = ")
	fmt.Scan(&length)

	actArr := make([]int, length)
	revArr := make([]int, length)

	fmt.Print("Enter the Array Items  = ")
	for i = 0; i < length; i++ {
		fmt.Scan(&actArr[i])
	}
	j = 0
	for i = length - 1; i >= 0; i-- {
		revArr[j] = actArr[i]
		j++
	}

	fmt.Println("\nThe Result of a Reverse Array = ", revArr)
}

// package main

// import "fmt"

// func reverseArray(actArr []int, actsize int) {
//     revArr := make([]int, actsize)
//     j := 0
//     for i := actsize - 1; i >= 0; i-- {
//         revArr[j] = actArr[i]
//         j++
//     }

//     fmt.Println("\nThe Result of a Reverse Array = ", revArr)
// }
// func main() {
//     var actsize, i int

//     fmt.Print("Enter the Even Array Size = ")
//     fmt.Scan(&actsize)

//     actArr := make([]int, actsize)

//     fmt.Print("Enter the Even Array Items  = ")
//     for i = 0; i < actsize; i++ {
//         fmt.Scan(&actArr[i])
//     }

//     reverseArray(actArr, actsize)
// }
