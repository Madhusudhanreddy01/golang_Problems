package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// Create a map to keep track of the elements seen so far
	elements := make(map[int]bool)

	// Iterate through the slice and check if each element is already in the map
	for _, element := range nums {
		if elements[element] {
			return true
		}
		elements[element] = true
	}

	// If we reach this point, it means that no duplicates were found
	return false
}

func main() {
	// Read the slice from the user
	var slice []int
	fmt.Print("Enter the elements of the slice: ")
	for {
		var element int
		_, err := fmt.Scan(&element)
		if err != nil {
			break
		}
		slice = append(slice, element)
	}

	// Check if the slice contains duplicate elements
	if containsDuplicate(slice) {
		fmt.Println("The slice contains duplicate elements.")
	} else {
		fmt.Println("The slice does not contain duplicate elements.")
	}
}
