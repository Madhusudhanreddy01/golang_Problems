// package main

// import "fmt"

// func main() {
// 	var length, target, val int

// 	fmt.Println("Enter the length of array")
// 	fmt.Scan(&length)

// 	for i := 0; i < length; i++ {
// 		fmt.Println("Enter the array elemnts using space")
// 		nums := []int{}
// 		for j := 0; j < length; j++ {
// 			fmt.Scan(&val)
// 			nums = append(nums, val)
// 		}

// 		fmt.Println("Enter target")
// 		fmt.Scan(&target)

// 		res := bruteForce(nums, target)
// 		fmt.Println(res)
// 		break
// 	}
// }

// func bruteForce(nums []int, target int) []int {
// 	for i, v := range nums {
// 		for j, x := range nums[i+1:] {
// 			if x == target-v {
// 				return []int{i, j + i + 1}
// 			}
// 		}
// 	}

// 	return []int{}
// }

// -----------------------------------------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func twoSum(nums []int, target int) []int {
// 	var a []int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums)-1; j++ {
// 			if nums[i]+nums[j+1] == target {
// 				a = append(a, i, j+1)
// 			}
// 		}
// 	}
// 	return a
// }
// func readLine(reader *bufio.Reader) string {
// 	str, _, _ := reader.ReadLine()
// 	return strings.TrimRight(string(str), "\r\n")
// }
// func main() {
// 	//read input
// 	fmt.Print("How ,many number you append to slice: ")
// 	var q int
// 	fmt.Scanln(&q)
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
// 	//string[2 7 11 5]
// 	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
// 	var a []int
// 	for i := 0; i < q; i++ {
// 		//string to int
// 		aItemTemp, _ := strconv.Atoi(aTemp[i])
// 		aItem := int(aItemTemp)
// 		a = append(a, aItem)
// 	}
// 	//target
// 	fmt.Print("Target: ")
// 	var w int
// 	fmt.Scanln(&w)
// 	b := twoSum(a, w)
// 	fmt.Println(b)
// }

// -----------------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	var length, value, target int

// 	fmt.Println("Enter the length of array")
// 	fmt.Scan(&length)

// 	fmt.Println("Enter the array elemnts using space")
// 	nums := []int{}

// 	for i := 0; i < length; i++ {
// 		fmt.Scan(&value)
// 		nums = append(nums, value)
// 	}

// 	fmt.Println("Enter target")
// 	fmt.Scan(&target)

// 	res := bruteForce(nums, target)
// 	fmt.Println(res)
// }

// func bruteForce(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums)-1; j++ {
// 			if nums[i]+nums[j+1] == target {
// 				return []int{i, j + 1}
// 			}
// 		}
// 	}

// 	// for i, v := range nums {
// 	// 	for j, x := range nums[i+1:] {
// 	// 		if x == target-v {
// 	// 			return []int{i, j + i + 1}
// 	// 		}
// 	// 	}
// 	// }

// 	return []int{}
// }

// --------------------------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BruteForce(nums []int, target int) []int {
	IntSlice := new([]int)
	// using linear search chek
	for index1, Num1 := range nums {
		for index2, Num2 := range nums {
			//if both the number add == target and index not same then append to IntSlice
			if Num1+Num2 == target && index1 != index2 {
				*IntSlice = append(*IntSlice, index1)
			}
		}
	}
	return *IntSlice
}
func HashMap(nums []int, target int) []int {
	//Hashmap store the complement and index
	see := make(map[int]int)
	//result
	IntSlice := new([]int)
	//chek the complement if true append both index to IntSlice if not update to Hashmap
	for Cindex := 0; Cindex < len(nums); Cindex++ {
		r := target - nums[Cindex]
		index, ok := see[r]
		if ok {
			*IntSlice = append(*IntSlice, index, Cindex)
		}
		//update to Hashmap
		//complement=index
		see[nums[Cindex]] = Cindex
	}
	return *IntSlice

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	//using reader read the input by using readline
	fmt.Print("Please enter your Slice Numbers: ")
	//example:
	//Please enter your Slice Numbers:2 7 11 6
	val, _, _ := reader.ReadLine()
	//ReadLine return byte,bool,error

	//convert byte to string
	Readline := strings.TrimSpace(string(val))

	//split convert string to slice_string
	input1 := strings.Split(Readline, " ")
	var a []int
	for i := 0; i < len(input1); i++ {
		intnum, err := strconv.Atoi(input1[i])
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, intnum)
	}

	//convert byte to string
	fmt.Print("Please enter your target Number : ")
	var w int
	fmt.Scanln(&w)
	output := HashMap(a, w)
	fmt.Println("Hash map Two Sum is :", output)
	output2 := BruteForce(a, w)
	fmt.Println("Bruit_Force Two Sum is :", output2)

}
