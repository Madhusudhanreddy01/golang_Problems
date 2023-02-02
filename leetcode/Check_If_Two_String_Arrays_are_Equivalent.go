// with user input
package main

import (
	"fmt"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	// Concatenate the strings in the arrays
	str1 := ""
	str2 := ""
	for _, w := range word1 {
		str1 += w
	}
	for _, w := range word2 {
		str2 += w
	}
	// Compare the concatenated strings
	return str1 == str2
}

func main() {
	var word1, word2 []string
	fmt.Println("Enter the elements of the first array, separated by commas:")
	fmt.Scan(&word1)
	fmt.Println("Enter the elements of the second array, separated by commas:")
	fmt.Scan(&word2)
	fmt.Println(arrayStringsAreEqual(word1, word2))
}

// without-userinput

// package main

// import "fmt"

// func arrayStringsAreEqual(word1 []string, word2 []string) bool {
// 	// Concatenate the strings in the arrays
// 	str1 := ""
// 	str2 := ""
// 	for _, w := range word1 {
// 		str1 += w
// 	}
// 	for _, w := range word2 {
// 		str2 += w
// 	}

// 	// Compare the concatenated strings
// 	return str1 == str2
// }

// func main() {
// 	word1 := []string{"Hello", "world"}
// 	word2 := []string{"Hello", "world"}
// 	fmt.Println(arrayStringsAreEqual(word1, word2))
// }

// -----------------------------------------------------------------

// package main

// import (
//     "fmt"
// )

// func main() {
//     // Get user input for first array
//     var array1 []string
//     fmt.Println("Enter the elements of the first array, separated by commas:")
//     fmt.Scan(&array1)

//     // Get user input for second array
//     var array2 []string
//     fmt.Println("Enter the elements of the second array, separated by commas:")
//     fmt.Scan(&array2)

//     // Compare the arrays
//     if len(array1) != len(array2) {
//         fmt.Println("The arrays are not equivalent.")
//     } else {
//         for i := 0; i < len(array1); i++ {
//             if array1[i] != array2[i] {
//                 fmt.Println("The arrays are not equivalent.")
//                 return
//             }
//         }
//         fmt.Println("The arrays are equivalent.")
//     }
// }
