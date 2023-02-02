package main

import "fmt"

func main() {
	var length, value, target int

	fmt.Println("Enter the length of array")
	fmt.Scan(&length)

	fmt.Println("Enter the array elemnts using space")
	nums := []int{}

	for i := 0; i < length; i++ {
		fmt.Scan(&value)
		nums = append(nums, value)
	}

	fmt.Println("Enter target")
	fmt.Scan(&target)

	res := bruteForce(nums, target)
	fmt.Println(res)
}

func bruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums)-1; j++ {
			if nums[i]+nums[j+1] == target {
				return []int{i, j + 1}
			}
		}
	}

	return []int{}
}
