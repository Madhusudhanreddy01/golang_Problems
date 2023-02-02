package main

import "fmt"

func main() {
	// fmt.Println("Main Function")
	// var testCase int

	// fmt.Println("Number of TestCases ")
	// fmt.Scan(&testCase)

	// for i := 0; i < length; i++ {
	// 	nums := []int{}
	// 	var length, target, val int
		// fmt.Println("------- Example :", i+1, "---------")
		// fmt.Println("Enter the length of array")
		// fmt.Scan(&length)

		fmt.Println("Enter the array elemnts using space")
		for j := 0; j < length; j++ {
			nums := []int{}
			var length, target, val int
			fmt.Scan(&val)
			nums = append(nums, val)
		}

		fmt.Println("Enter target")
		fmt.Scan(&target)

		res := bruteForce(nums, target)
		fmt.Println(res)
	}
}

func bruteForce(nums []int, target int) []int {
	for i, v := range nums {
		for j, x := range nums[i+1:] {
			if x == target-v {
				return []int{i, j + i + 1}
			}
		}
	}

	return []int{}
}
