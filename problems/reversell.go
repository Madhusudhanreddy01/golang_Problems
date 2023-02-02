package main

import "fmt"

func main() {
	var length, value int

	fmt.Println("Enter the length of array")
	fmt.Scan(&length)

	fmt.Println("Enter the array elemnts using space")
	nums := []int{}

	for i := 0; i < length; i++ {
		fmt.Scan(&value)
		nums = append(nums, value)
	}

	res := reverseList(head)
	fmt.Println(res)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
