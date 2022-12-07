package main

import "fmt"

func main() {
	var length, target, value int

	fmt.Println("Enter the length of array")
	fmt.Scan(&length)

	// for i := 0; i < length; i++ {
	fmt.Println("Enter the array elemnts using space")
	nums := []int{}
	for j := 0; j < length; j++ {
		fmt.Scan(&value)
		nums = append(nums, value)
	}

	fmt.Println("Enter target")
	fmt.Scan(&target)

	res := bruteForce(nums, target)
	fmt.Println(res)
	// break
	// }
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

// for i, v := range nums {
// 	for j, x := range nums[i+1:] {
// 		if x == target-v {
// 			return []int{i, j + i + 1}
// 		}
// 	}
// }
