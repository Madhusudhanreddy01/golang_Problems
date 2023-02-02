package main

import (
	"fmt"
	"sort"
)

func main() {
	// numbers := []int{1, 2, -2, 0, -1}
	numbers := []int{30, 12, -53, 27, -80, 5}

	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)
}
