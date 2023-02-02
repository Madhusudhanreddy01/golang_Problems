package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func main() {
	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sorted list of integers, separated by spaces:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Parse the input into a linked list
	var head *ListNode
	for _, s := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(s)
		node := &ListNode{Val: val, Next: head}
		head = node
	}

	// Remove duplicates from the list
	res := deleteDuplicates(head)

	// Print the modified list
	fmt.Println("Modified list:")
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
	fmt.Println()
}

// --------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	curr := head
// 	for curr.Next != nil {
// 		if curr.Val == curr.Next.Val {
// 			curr.Next = curr.Next.Next
// 		} else {
// 			curr = curr.Next
// 		}
// 	}
// 	return head
// }

// func main() {
// 	// Test case 1
// 	l1 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
// 	res1 := deleteDuplicates(l1)
// 	fmt.Println("Test case 1:")
// 	for res1 != nil {
// 		fmt.Print(res1.Val, " ")
// 		res1 = res1.Next
// 	}
// 	fmt.Println()

// 	// Test case 2
// 	l2 := &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}}
// 	res2 := deleteDuplicates(l2)
// 	fmt.Println("Test case 2:")
// 	for res2 != nil {
// 		fmt.Print(res2.Val, " ")
// 		res2 = res2.Next
// 	}
// 	fmt.Println()

// 	// Test case 3
// 	l3 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
// 	res3 := deleteDuplicates(l3)
// 	fmt.Println("Test case 3:")
// 	for res3 != nil {
// 		fmt.Print(res3.Val, " ")
// 		res3 = res3.Next
// 	}
// 	fmt.Println()
// }
