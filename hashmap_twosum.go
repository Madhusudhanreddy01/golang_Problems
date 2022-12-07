package main

import "fmt"

func main() {
	var length, value, target int

	fmt.Println("Enter the length of array")
	fmt.Scan(&length)

	fmt.Println("Enter the array elemnts using space")
	nums := []int{}

	for j := 0; j < length; j++ {
		fmt.Scan(&value)
		nums = append(nums, value)
	}

	fmt.Println("Enter target")
	fmt.Scan(&target)

	res := usingMap(nums, target)
	fmt.Println(res)
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
