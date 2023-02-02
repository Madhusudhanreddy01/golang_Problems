package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// [1,2,2,3,5,6]
	//[1,2,3] [2,5,6]
	//     L     R
	Left := 0
	Right := 0
	var a []int
	for Left < m && Right < n {
		if nums1[Left] < nums2[Right] {
			a = append(a, nums1[Left])
			Left++
		} else {
			a = append(a, nums2[Right])
			Right++
		}
	}
	if Left < m {
		a = append(a, nums1[Left:]...)
	}
	if Right < n {
		a = append(a, nums2[Right:]...)
	}

}
func main() {
	arr := []int{1, 2, 3}
	arr2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(arr, m, arr2, n)
	var a []int
	// arr := []int{10, 20, 30, 40, 50, 60}
	a = append(a, arr[2:]...)
	fmt.Println(a)

}
