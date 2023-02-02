package main

import "fmt"

func main() {
	fmt.Println("Main Function")
	var testCase int

	fmt.Println("Number of TestCases ")
	fmt.Scan(&testCase)

	for i := 0; i < testCase; i++ {
		nums := []int{}
		var length, target, val int
		fmt.Println("------- Example :", i+1, "---------")
		fmt.Println("Enter the length of array")
		fmt.Scan(&length)

		fmt.Println("Enter the array elemnts using space")
		for j := 0; j < length; j++ {
			fmt.Scan(&val)
			nums = append(nums, val)
		}

		fmt.Println("Enter target")
		fmt.Scan(&target)

		// res := bruteForce(nums, target)
		// fmt.Println(res)
		res := usingMap(nums, target)
		fmt.Println(res)
	}
}
func usingMap(nums []int, target int) []int {
	temp := map[int]int{}

	for i, v := range nums {
		if val, ok := temp[target-v]; ok {
			return []int{val, i}
		}
		temp[v] = i
	}

	return []int{}
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
